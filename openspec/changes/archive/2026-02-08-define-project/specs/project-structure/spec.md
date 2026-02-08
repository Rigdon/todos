# Project Structure

## ADDED Requirements

### Requirement: Project Roots
The project MUST have valid configuration files for the chosen stack.

#### Scenario: Backend Initialization
- **WHEN** the project is initialized
- **THEN** a `go.mod` file exists in the root
- **AND** the module name matches the project directory

#### Scenario: Frontend Initialization
- **WHEN** the project is initialized
- **THEN** a `package.json` file exists in the root
- **AND** it contains `typescript` as a dependency

### Requirement: Context Configuration
The OpenSpec configuration MUST reflect the project identity.

#### Scenario: Config Context
- **WHEN** `openspec/config.yaml` is read
- **THEN** it contains a `context` section
- **AND** the context mentions "Go", "TypeScript", and "SQLite3"
