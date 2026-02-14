## Why

The current CI workflows use mutable tags (e.g., `v4`) for GitHub Actions, which poses a security risk if a tag is maliciously updated. Additionally, the API linting step does not use caching, slowing down feedback loops. Finally, there is no automated system to keep dependencies (Go, npm, Actions) up to date.

## What Changes

- **Pin Actions**: All GitHub Actions will be pinned to specific Git commit SHAs.
- **Cache API Lint**: The `api-lint` workflow will be configured to cache `npm` dependencies.
- **Automated Updates**: Dependabot will be configured to automatically open PRs for dependency updates.

## Capabilities

### New Capabilities
- `ci-hardening`: Improvements to CI security and performance.

### Modified Capabilities
<!-- No requirement changes to existing capabilities. -->

## Impact

- **CI/CD**: Workflows will be more secure and precise. API linting will be faster.
- **Maintenance**: Automated PRs for dependency updates will require regular review.
