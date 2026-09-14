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

## Start the current foundation

Start the API from the repository root:

```powershell
go run ./backend/cmd/server
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
go test ./backend/...
go build ./backend/...
Set-Location frontend
npm run lint
npm run build
```

Use `.env.example` as the starting point for local configuration. Never commit real credentials or production data.
