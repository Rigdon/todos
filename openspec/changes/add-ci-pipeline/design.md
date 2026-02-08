## Context
Currently, tests must be run manually before pushing. There is no automated verification for PRs, leading to potential broken builds or regressions in `main`.

## Goals / Non-Goals
**Goals:**
- Automate `go test` and `go build` on every PR and push to main.
- prevent merging broken code.
- Detect vulnerable dependencies in PRs.

**Non-Goals:**
- Deployment automation (CD) - out of scope for now.
- Linting (golangci-lint) - can be added later.

## Decisions

### Decision 1: GitHub Actions
We will use GitHub Actions as the CI provider.
- **Why**: Native integration with GitHub, free for public repos (and reasonable tier for private), no external infra needed.

### Decision 2: Workflow Structure
We will create two separate workflows:
1.  `ci.yaml`: Runs checking application code (Build & Test).
    - Triggers: `push` to main, `pull_request`.
    - Job: `test` (Setup Go 1.22, Checkout, Test, Build).
2.  `dependency-review.yaml`: Runs security scans.
    - Triggers: `pull_request`.
    - Action: `actions/dependency-review-action`.

### Decision 3: Go Version
- Use `1.22` in `actions/setup-go` to match `go.mod`.

## Risks / Trade-offs
- **Risk**: Flaky tests blocking merges.
    - **Mitigation**: Ensure tests are robust. We can mark checks as optional in repository settings if needed temporarily.
