# Dependency Review

## ADDED Requirements

### Requirement: Dependency Vulnerability Scanning
The system SHALL scan dependencies for known vulnerabilities on every Pull Request.

#### Scenario: no vulnerabilities
- **WHEN** a PR is opened or updated
- **AND** no known vulnerabilities are found in dependencies
- **THEN** the check MUST pass

#### Scenario: vulnerability found
- **WHEN** a PR introduces a dependency with a known vulnerability
- **THEN** the check MUST fail
- **AND** the PR merge MUST be blocked
