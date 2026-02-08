# Context
The project is currently empty. We need to initialize the standard toolchains.

# Goals
- Initialize Go module `todos`
- Initialize Node package with TypeScript
- Configure OpenSpec with project context

# Decisions

## Decision 1: Stack Initialization
- **Go**: Use `go mod init todos` to match the directory name.
- **Node**: Use `npm init -y` followed by `npm install -D typescript @types/node`.
- **Database**: SQLite3 will be a dependency of the Go backend (likely `mattn/go-sqlite3` or modern CGO-free alternatives like `modernc.org/sqlite` later, but for now just the concept).

## Decision 2: OpenSpec Context
We will prepend the context string to `openspec/config.yaml` to ensure AI agents know the preferred stack immediately.
