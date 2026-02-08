# Why
The application currently has no backend logic. We need a REST API to manage tasks (create, read, update, delete) to support the frontend.

# What Changes
- Create a new Go package `internal/task` for task logic.
- Implement HTTP handlers for:
  - `GET /tasks` (List)
  - `POST /tasks` (Create)
  - `PUT /tasks/{id}` (Update)
  - `DELETE /tasks/{id}` (Delete)
- Connect handlers to SQLite database.

# Capabilities

## New Capabilities
- `task-management`: Core API for managing tasks.

## Modified Capabilities
<!-- None -->

# Impact
- New `internal/task` package.
- `main.go`: Mounts new routes.
