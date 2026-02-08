# Continuous Integration

## Purpose
To ensure code quality and stability by automatically running tests and verifying builds on every pull request and push to main.

## Requirements

### Requirement: Unit Testing on PR
The system SHALL run unit tests on every Pull Request.

#### Scenario: tests pass
- **WHEN** a PR is opened or updated
- **AND** all unit tests pass (`go test ./...`)
- **THEN** the CI check MUST be marked as "success"

#### Scenario: tests fail
- **WHEN** a PR is opened or updated
- **AND** any unit test fails
- **THEN** the CI check MUST be marked as "failure"
- **AND** the PR merge MUST be blocked

### Requirement: Build Verification
The system SHALL verify the code builds successfully.

#### Scenario: build success
- **WHEN** a PR is opened or updated
- **AND** `go build ./...` succeeds
- **THEN** the CI check MUST be marked as "success"

#### Scenario: build failure
- **WHEN** a PR is opened or updated
- **AND** `go build ./...` fails
- **THEN** the CI check MUST be marked as "failure"
- **AND** the PR merge MUST be blocked
