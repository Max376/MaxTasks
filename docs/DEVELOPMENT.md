# Development guide

## Prerequisites

- Go 1.27 or newer
- Node.js and npm
- Docker Desktop with Docker Compose
- Git

Check the installed tools with:

```powershell
go version
node --version
npm --version
docker compose version
```

## Repository layout

```text
backend/       Go API service
frontend/      React and TypeScript web client
migrations/    PostgreSQL schema migrations
docs/          Project and technical documentation
.github/       GitHub issue and pull-request workflow
```

## Start the local Docker runtime

Docker Compose runs the frontend, Go API, and PostgreSQL together. Create a local environment file once, then start the complete stack from the repository root:

```powershell
Copy-Item .env.example .env
# Edit .env and replace POSTGRES_PASSWORD with a local-only value.
docker compose up --build
```

Open the frontend at `http://localhost:5173`. The API startup endpoint is available at `http://localhost:8080/`, and PostgreSQL is exposed on `localhost:5432` for local tools. The host ports can be changed in `.env` with `FRONTEND_PORT`, `API_PORT`, and `POSTGRES_PORT`.

The `postgres_data` named volume keeps PostgreSQL data when containers are stopped or recreated:

```powershell
docker compose down
docker compose up
```

To remove the database data as well, explicitly remove the named volume:

```powershell
docker compose down --volumes
```

Never commit `.env`; it is ignored by Git. Compose builds the API's internal `DATABASE_URL` from the PostgreSQL variables, so the database hostname inside the Docker network remains `db`.

## Apply database migrations

Migrations are versioned SQL files in `migrations/` and run with the pinned
`migrate/migrate` image. Start the database and apply all pending migrations
from the repository root:

```powershell
docker compose up -d db
docker compose --profile tools run --rm migrate
```

The migration service waits for PostgreSQL to become healthy and records the
applied version in `schema_migrations`. It is safe to run the command again;
already-applied migrations are skipped. The schema, relationships, indexes,
and rollback file are documented in [`migrations/README.md`](../migrations/README.md).

To start the complete stack after the schema is applied:

```powershell
docker compose up -d api frontend
```

For a clean local database, remove the development volume and repeat the
commands above. This deletes only local data:

```powershell
docker compose down --volumes
docker compose up -d db
docker compose --profile tools run --rm migrate
```

## Troubleshooting Docker

- Check the rendered, non-secret Compose configuration with `docker compose --env-file .env.example config`.
- Check the migration service configuration with `docker compose --env-file .env.example --profile tools config --quiet`.
- If a migration fails, inspect the database logs with `docker compose logs db`; fix the migration before retrying it.
- Check service state with `docker compose ps` and inspect logs with `docker compose logs api`, `docker compose logs frontend`, or `docker compose logs db`.
- If a host port is already in use, change the corresponding `*_PORT` value in `.env`, then run `docker compose up --build` again.
- If the database was initialized with old credentials, stop the stack with `docker compose down --volumes` and start it again. This deletes only the local development database volume.
- If Docker Desktop is not running, start it and retry `docker compose up --build`.

## Start the current foundation without Docker

From the repository root, enter the backend module and start the API:

```powershell
Set-Location backend
go run ./cmd/server
```

In a second terminal, install and start the web client:

```powershell
Set-Location frontend
npm install
npm run dev
```

The API listens on `http://localhost:8080` by default. The frontend development server prints its local URL after startup.

## Quality checks

Run these checks before opening a pull request:

```powershell
Set-Location backend
go test ./...
go build ./...
Set-Location ..
Set-Location frontend
npm run lint
npm run build
```

Use `.env.example` as the starting point for local configuration. Never commit real credentials or production data.
