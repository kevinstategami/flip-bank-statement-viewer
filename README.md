# Bank Statement Viewer - Backend (Go)

Simple backend for take-home test: upload CSV, compute balance, and list non-successful transactions.

## Run locally

1. Clone / copy files
2. `go run cmd/server/main.go`
3. Server runs at `http://localhost:8080`

## Endpoints

- `POST /upload` - multipart form with key `file` (CSV)
- `GET /balance` - returns `{ "balance": "12345" }`
- `GET /issues` - returns JSON array of FAILED or PENDING transactions

## CSV format

`timestamp, name, type, amount, status, description`

Use `sample.csv` as example.

## Tests

Run unit tests:
```bash
go test ./...
