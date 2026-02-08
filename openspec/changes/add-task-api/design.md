# Context
No backend logic exists. We are bootstrapping the API layer.

# Goals
- Implement `Task` domain model (`id`, `title`, `status`, `created_at`)
- Implement RESTful HTTP handlers
- Implement SQLite persistence
- Dependency Injection for testability

# Decisions

## Decision 1: Package Structure
We will use a domain-oriented package `internal/task`.
- **Why**: Keeps domain logic, persistence, and transport concerns cohesive but decoupled via interfaces.
- **Alternatives**: Flat structure (everything in `main.go`) - rejected as it doesn't scale.

## Decision 2: Persistence
- **Interface**: `task.Repository` interface defined in `internal/task`.
- **Implementation**: `internal/task/sqlite.go` using `database/sql`.
- **Schema**:
  ```sql
  CREATE TABLE tasks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
  );
  ```

## Decision 3: JSON Handling
- Use standard `encoding/json`.
- Strict decoding (disallow unknown fields) to prevent client errors going unnoticed.

## Decision 4: Routing
- Use Go 1.22+ `net/http` enhanced routing (`mux.HandleFunc("POST /tasks", ...)`).
- **Why**: Standard library is now sufficient, no need for `chi` or `gorilla/mux` dependencies yet.

# Risks
- **Risk**: SQLite CGO requirement.
- **Mitigation**: Will likely need `CGO_ENABLED=1`. If this becomes painful we can switch to `modernc.org/sqlite` later. For now, assuming standard environment.
