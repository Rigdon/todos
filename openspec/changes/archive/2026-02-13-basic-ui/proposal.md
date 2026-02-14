## Why

Currently, interacting with the Todos application requires using command-line tools like `curl`. This is inefficient for managing tasks. A visual interface is needed to make the application user-friendly and allows for easier task management.

## What Changes

- Create a `public` directory to host static assets.
- Add `index.html` as the entry point for the application.
- Add `style.css` for basic styling.
- Add `app.ts` (and compile to `app.js`) for frontend logic to interact with the existing API.
- Update `main.go` to serve static files from the `public` directory.
- Update `package.json` to include a build script for TypeScript.

## Capabilities

### New Capabilities
- `basic-ui`: A web-based interface for listing, creating, updating, and deleting tasks.

### Modified Capabilities
<!-- No existing capabilities are changing requirements. -->

## Impact

- **Backend**: `main.go` will be modified to include a file server handler.
- **Frontend**: New files in `public/`.
- **Build**: New TypeScript compilation step.
