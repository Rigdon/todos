# Implementation Tasks

## 1. Domain & Persistence

- [x] 1.1 Create `internal/task` package and define `Task` struct & `Repository` interface
- [x] 1.2 Implement `SQLiteRepository` in `internal/task/sqlite.go` with schema migration
- [x] 1.3 Add Repository integration tests

## 2. HTTP Handlers

- [x] 2.1 Implement `Handler` struct and dependency injection in `internal/task/handler.go`
- [x] 2.2 Implement `Create` and `List` handlers with JSON encoding/decoding
- [x] 2.3 Implement `Update` and `Delete` handlers
- [x] 2.4 Add Handler unit tests

## 3. Wiring & Verification

- [x] 3.1 Initialize SQLite and mount routes in `main.go`
- [x] 3.2 Verify all endpoints with `curl` (manual check)
