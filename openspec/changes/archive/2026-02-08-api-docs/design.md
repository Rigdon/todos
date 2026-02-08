## Context

The project currently has a REST API exposed by the Go backend but lacks formal documentation. This makes it hard for frontend developers or external consumers to know the contract.

## Goals / Non-Goals

**Goals:**
- Create a single source of truth for the API contract (`openapi.yaml`).
- Ensure the spec file itself is syntactically valid and follows best practices via CI linting.

**Non-Goals:**
- Automated code generation from spec (at this stage).
- Automated contract testing (verifying implementation matches spec).

## Decisions

### 1. Spec format and location
- **Decision**: Hand-authored OpenAPI 3.0 YAML file at repository root (`/openapi.yaml`).
- **Rationale**: 
  - Manual authoring is simple and sufficient for a small API.
  - YAML is more readable than JSON.
  - Root location is standard and easy to find.
- **Alternatives**: 
  - Generating from code (swag/swaggo): Requires adding tags to code, which is intrusive for a first pass.
  - `docs/` folder: also viable, but root is more visible.

### 2. Linting Tool
- **Decision**: Use `redocly/cli` via GitHub Actions.
- **Rationale**:
  - Fast, standard, and has good default rules (`recommended` config).
  - Easy to run via `npx` or detailed action.
- **Alternatives**:
  - `stoplight/spectral`: Also good, but Redocly is very popular for OpenAPI specifically.

## Risks / Trade-offs

- **Risk**: Spec drifts from implementation.
- **Mitigation**: manual review required on PRs. Future work could add contract testing (e.g. `dredd` or similar).

## Verification Plan

### Automated Tests
- The new CI step will test the validity of `openapi.yaml`.
- We can manually trigger a failure by committing an invalid file to a branch and push.

### Manual Verification
- View the rendered spec (e.g. via VS Code extension or pasting into editor.swagger.io) to ensure it looks correct.
