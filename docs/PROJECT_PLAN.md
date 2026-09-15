# MaxTasks Project Plan

## 1. Project vision

MaxTasks is a self-hosted, API-first task manager for web and Android. It provides a focused task-list experience inspired by Google Tasks while demonstrating production-minded software development with React, Go, PostgreSQL, Docker, automated testing, and continuous integration.

The project should be easy to run locally, easy to understand from its documentation, and reliable enough to demonstrate as a portfolio project.

## 2. Goals

- Provide a simple, responsive task-management experience.
- Make the application installable on Android as a Progressive Web App (PWA).
- Expose the core functionality through a documented REST API.
- Support self-hosted deployment with Docker Compose.
- Demonstrate secure data access, database migrations, testing, and CI.
- Document decisions and setup steps so another developer can run and evaluate the project.

## 3. Non-goals for the first release

The first release will not attempt to reproduce every feature of Google Tasks. The following items are intentionally deferred:

- Native Android and iOS applications
- Team workspaces and task assignment
- Calendar integrations
- Real-time collaboration
- Notifications and reminders
- Importing data from external task managers

These can be considered after the MVP is stable and documented.

## 4. MVP scope

### User experience

- Register and sign in.
- Create, rename, and delete task lists.
- Create, edit, complete, and delete tasks.
- Set an optional due date.
- Reorder tasks within a list.
- View the application on desktop and mobile layouts.
- Install the web application as a PWA on Android.

### API

- REST endpoints for authentication, lists, and tasks.
- Validation with clear error responses.
- Authentication and authorization on every user-owned resource.
- OpenAPI documentation with example requests and responses.
- A health-check endpoint for deployment monitoring.

### Operations

- Local development through Docker Compose.
- PostgreSQL schema managed with versioned migrations.
- Automated tests for important backend and frontend behavior.
- GitHub Actions for formatting, linting, tests, and builds.
- Documented backup and restore procedure.

## 5. Proposed architecture

```text
React + TypeScript PWA
          |
          | HTTPS / JSON REST API
          v
        Go API
          |
          v
      PostgreSQL
```

The frontend is responsible for presentation and client-side interaction. The Go service owns business rules, authentication, authorization, validation, and persistence. PostgreSQL is the source of truth for application data. Docker Compose provides a repeatable local and self-hosted runtime.

## 6. Milestones

### Milestone 0 — Documentation and repository foundation

- [x] Add the project description and plan.
- [ ] Decide the initial license.
- [ ] Add contribution guidelines and code of conduct.
- [ ] Create GitHub labels and milestones.
- [ ] Add issue and pull-request templates.
- [ ] Record the initial architecture decision.

### Milestone 1 — Technical foundation

- [ ] Create the Go backend module.
- [ ] Create the React and TypeScript frontend.
- [ ] Add Docker Compose for the application and PostgreSQL.
- [ ] Add configuration through environment variables.
- [x] Add database migrations.
- [ ] Add `/health` and readiness checks.

### Milestone 2 — Core task management

- [ ] Implement the user and session model.
- [ ] Implement task-list CRUD operations.
- [ ] Implement task CRUD operations.
- [ ] Implement completion and ordering.
- [ ] Add authorization tests for user-owned data.

### Milestone 3 — Usable web and Android experience

- [ ] Build the main task-list interface.
- [ ] Add responsive mobile layouts.
- [ ] Add loading, empty, validation, and error states.
- [ ] Add PWA manifest and service-worker support.
- [ ] Verify installation and basic use on Android.

### Milestone 4 — Quality and public API

- [ ] Publish the OpenAPI specification.
- [ ] Add API examples and local API usage instructions.
- [ ] Add integration and end-to-end tests.
- [ ] Add GitHub Actions CI.
- [ ] Add structured logging and safe error handling.

### Milestone 5 — Release and portfolio presentation

- [ ] Deploy a live demo with non-sensitive sample data.
- [ ] Add screenshots and an architecture diagram to the README.
- [ ] Document backup and restore.
- [ ] Create a tagged MVP release.
- [ ] Record known limitations and the next roadmap items.

## 7. Delivery workflow

Each feature should be tracked as a GitHub issue with a clear outcome and acceptance criteria. Work should be developed on a short-lived branch and merged through a pull request. Pull requests should explain the change, link the issue, and include the relevant test or verification result.

Recommended branch naming:

```text
feature/<short-name>
fix/<short-name>
docs/<short-name>
```

Recommended issue format:

```text
## Goal

What user or engineering outcome should this issue deliver?

## Acceptance criteria

- [ ] Criterion one
- [ ] Criterion two

## Notes

Relevant technical decisions, screenshots, or references.
```

## 8. Definition of done

A feature is done when:

- Its behavior matches the acceptance criteria.
- Invalid input and expected failure cases are handled.
- User-owned data is protected by authorization checks.
- Automated tests cover the important behavior.
- Documentation is updated when setup, API behavior, or user behavior changes.
- The project still passes formatting, linting, tests, and builds in CI.

## 9. Documentation map

The repository should keep documentation close to the code:

- `README.md` — project overview, screenshots, quick start, and demo link.
- `docs/PROJECT_PLAN.md` — scope, milestones, workflow, and delivery criteria.
- `docs/decisions/` — architecture decision records.
- `docs/api/` — OpenAPI specification and API examples.
- `docs/operations/` — deployment, backup, restore, and troubleshooting guides.

## 10. Risks and decisions to revisit

- Authentication design must balance portfolio value with implementation size.
- Offline editing and synchronization should only be added after the online MVP is reliable.
- A public demo must use isolated sample data and must not expose development credentials.
- Database backups should be encrypted and stored separately from normal source-code history.

## 11. Current status

The repository is implementing the technical foundation: a Dockerized Go API,
React frontend, and PostgreSQL database with a versioned initial migration.
