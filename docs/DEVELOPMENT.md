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

## Troubleshooting Docker

- Check the rendered, non-secret Compose configuration with `docker compose --env-file .env.example config`.
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
