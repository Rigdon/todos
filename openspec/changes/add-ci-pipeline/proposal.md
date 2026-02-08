# Why
To ensure code quality and security, we need automated checks for every pull request. Currently, testing is manual.

# What Changes
- Add GitHub Actions workflow for Unit Testing & Building (`.github/workflows/ci.yaml`) - *Rolled out in PR 1*
- Add GitHub Actions workflow for Dependency Review (`.github/workflows/dependency-review.yaml`) - *Rolled out in PR 2*

# Capabilities

## New Capabilities
- `continuous-integration`: Automates `go test`, `go build`, and potentially linting on every push/PR.
- `dependency-review`: Automates scanning of dependencies for vulnerabilities on PRs.

## Modified Capabilities
None.

# Impact
- All future PRs will be blocked if checks fail.
- Requires `GITHUB_TOKEN` permissions for dependency review.
