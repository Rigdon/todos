# CI Hardening

## Purpose
To improve the security and performance of the CI pipeline by ensuring immutable dependencies, efficient caching, and automated updates.

## Requirements

### Requirement: CI Security Pinning
The system's CI/CD workflows SHALL use immutable commit SHAs for all third-party GitHub Actions to prevent supply chain attacks via mutable tags.

#### Scenario: Verify Action References
- **WHEN** a workflow configuration is inspected
- **THEN** all `uses` directives for third-party actions contain a 40-character commit SHA
- **AND** do not use version tags (e.g., `@v4`)

### Requirement: API Lint Caching
The API linting workflow SHALL utilize caching to improve execution speed by avoiding redundant dependency downloads.

#### Scenario: Linting with Cache
- **WHEN** the `api-lint` workflow runs
- **AND** `package-lock.json` has not changed since the last successful run
- **THEN** `npm` dependencies are restored from the cache
- **AND** `npm ci` executes faster than a fresh install

### Requirement: Automated Dependency Updates
The system SHALL automatically check for and propose updates to dependencies on a weekly basis.

#### Scenario: Weekly Dependabot Run
- **WHEN** the scheduled time arrives (e.g., weekly)
- **THEN** Dependabot scans for outdated dependencies in:
    - GitHub Actions (.github/workflows)
    - Go modules (go.mod)
    - npm packages (package.json)
- **AND** opens pull requests for any available updates
