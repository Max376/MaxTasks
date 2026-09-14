# MaxTasks API

The backend is a Go HTTP service. It currently exposes a small startup endpoint while the API and database work is implemented in the Foundation and MVP milestones.

## Local commands

From `backend/`:

```powershell
go run ./cmd/server
go test ./...
go build ./...
gofmt -w .
```

Set `API_PORT` to change the listening port. The default is `8080`.
