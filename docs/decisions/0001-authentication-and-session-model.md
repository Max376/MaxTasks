# ADR-0001: Authentication and session model

- Status: Accepted
- Date: 2026-09-15
- Related issues: #5

## Context

MaxTasks needs authentication for the MVP so that each task list and task is
owned by the signed-in user. The primary client is a React browser application
that is also installable as an Android PWA. The Go API and PostgreSQL database
run together in the self-hosted deployment.

The design must protect passwords and browser sessions, support logout and
forced revocation, avoid account-enumeration leaks, and remain small enough to
operate in a self-hosted installation. The design covers the MVP browser/PWA
client. Native clients and third-party integrations are outside this decision.

## Decision

MaxTasks will use server-side sessions backed by PostgreSQL and an opaque,
cryptographically random session cookie.

### Passwords

- Store only a password hash in the user record; never store or log a plaintext
  password.
- Hash passwords with Argon2id. The initial implementation baseline is a
  16-byte random salt, 32-byte output, three iterations, 64 MiB memory, and one
  lane. Parameters must be benchmarked on the deployment target and kept in a
  versioned password-hash format so they can be increased later.
- Normalize email addresses for uniqueness and lookup, while preserving the
  original display form only if the product later needs it.

### Sessions

- Generate at least 32 bytes of cryptographically secure random token data at
  sign-in. Send the token only in the `maxtasks_session` cookie and store a
  one-way hash of the token in PostgreSQL.
- Store the user ID, creation time, last-seen time, absolute expiry, and
  revocation time for each session. Use a 30-day absolute lifetime and a
  seven-day idle lifetime for the MVP.
- Rotate the session after successful sign-in and other authentication-boundary
  changes. Revoke the current session on logout, and provide a server-side
  path to revoke all sessions after a password change or suspected compromise.
- Set `HttpOnly`, `Path=/`, and `SameSite=Lax` on the cookie. Set `Secure` in
  production and whenever HTTPS is used. Local HTTP development may disable
  `Secure` through an explicit development configuration.

### API and browser behavior

- The browser sends credentials with same-origin API requests. A deployed
  reverse proxy should serve the frontend and API from one origin. Local
  development may use a configured allowlist for the frontend origin.
- Credentialed CORS, when needed locally, must name the configured origin and
  must never use `*` with credentials.
- Validate the `Origin` header for state-changing requests and require a CSRF
  token for cookie-authenticated state changes. Authentication implementation
  work will define the concrete CSRF endpoint and middleware contract.
- Every protected endpoint loads the session from PostgreSQL and authorizes
  the requested resource against the session's user ID. An expired or revoked
  session is treated as unauthenticated.

### Error behavior

- Sign-in and password-recovery failures use the same generic response for an
  unknown email, a wrong password, an expired session, and a revoked session.
  Responses, timing-sensitive logs, and metrics must not reveal whether an
  account exists.
- Do not include password hashes, session tokens, database errors, or stack
  traces in API responses or normal application logs.

## Alternatives considered

### JSON Web Tokens in the browser

JWT access tokens would make stateless API verification easy, but browser token
storage and refresh-token handling increase exposure to XSS and add complexity
to immediate logout and revocation. Server-side sessions provide the required
revocation behavior with the existing PostgreSQL dependency.

### Server-side sessions stored only in process memory

In-memory sessions are simple for one process, but they disappear on restart
and do not work reliably when the API is scaled to multiple containers. The
PostgreSQL-backed session store matches the self-hosted architecture and keeps
the behavior predictable across restarts.

### External identity provider or OAuth-only login

An external provider would reduce password handling in MaxTasks, but it adds a
service dependency and setup burden to a self-hosted portfolio project. It can
be added later through a separate ADR without changing the local account
session contract.

## Consequences

PostgreSQL becomes the source of truth for active sessions, so logout,
password-change revocation, and forced sign-out work immediately. The API must
query or cache session state on protected requests and periodically remove
expired rows. Cookie authentication gives the browser a safer default than
JavaScript-readable tokens, while the CSRF checks become mandatory for
state-changing routes.

The Go service must include Argon2id and session middleware, configurable
cookie/CORS settings, generic authentication errors, and tests for expiry,
revocation, ownership, and account-enumeration behavior. A future native client
or third-party integration may require a separate token or OAuth decision.
