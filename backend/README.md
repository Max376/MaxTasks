# MaxTasks API

The backend is a Go HTTP service. It exposes process health and database
readiness endpoints for local development and deployment monitoring.

## Health and readiness

`GET /health` checks that the API process can accept requests. It returns
`200 OK` with this JSON shape, even when PostgreSQL is unavailable:

```json
{"service":"maxtasks-api","status":"ok"}
```

`GET /ready` checks the configured PostgreSQL connection with a short timeout.
It returns `200 OK` when the database responds:

```json
{"service":"maxtasks-api","status":"ready","checks":{"database":"ok"}}
```

When `DATABASE_URL` is missing or PostgreSQL cannot be reached, it returns
`503 Service Unavailable` with the same documented shape and generic status
values. Connection details and internal errors are never included in the
response:

```json
{"service":"maxtasks-api","status":"not_ready","checks":{"database":"unavailable"}}
```

## Local commands

From `backend/`:

```powershell
go run ./cmd/server
go test ./...
go build ./...
gofmt -w .
```

Set `API_PORT` to change the listening port. The default is `8080`. Set
`DATABASE_URL` to configure PostgreSQL; when it is absent, `/health` remains
available while `/ready` reports `503 Service Unavailable`.
