## Context

The Todos application currently lacks a graphical user interface, relying on command-line tools for interaction. This change introduces a basic web frontend to improve usability. The backend is written in Go, and we strictly use SQLite for persistence.

## Goals / Non-Goals

**Goals:**
- Provide a user-friendly interface for managing tasks.
- Keep the implementation simple and lightweight.
- Use the existing backend API without modification (except for serving static files).

**Non-Goals:**
- Implementing user authentication (out of scope for now).
- Using a complex frontend framework (React, Vue, etc.).
- Real-time updates (WebSockets).

## Decisions

### 1. Frontend Technology Stack
- **Decision**: Use Vanilla TypeScript and CSS.
- **Rationale**: For a simple "basic" UI, a full framework like React or Angular introduces unnecessary complexity and build overhead. Vanilla TS provides type safety while keeping the bundle size small and the codebase easy to understand.
- **Alternatives Considered**: React, Vue, Svelte. Rejected due to complexity for this specific scope.

### 2. Serving Static Assets
- **Decision**: Serve static files directly from the Go backend using `http.FileServer`.
- **Rationale**: Simplifies deployment and development. No need for a separate frontend server (e.g., Nginx) for this basic version.
- **Alternatives Considered**: Separate frontend server. Rejected as it complicates the architecture unnecessarily.

### 3. DOM Manipulation
- **Decision**: Direct DOM manipulation using TypeScript.
- **Rationale**: Efficient for small-scale updates.
- **Risk Mitigation**: Will use helper functions to create elements safely and avoid XSS vulnerabilities (e.g., setting `textContent` instead of `innerHTML`).

## Risks / Trade-offs

- **Risk**: Vanilla JS/TS can become "spaghetti code" as complexity grows.
  - **Mitigation**: Structure the code into logical modules (e.g., `api.ts`, `dom.ts`) even without a framework.
- **Risk**: Lack of component reusability compared to frameworks.
  - **Mitigation**: Create simple factory functions for repeatable UI elements (tasks).
- **Risk**: Security (XSS).
  - **Mitigation**: Strictly strictly avoid `innerHTML` for user-generated content.
