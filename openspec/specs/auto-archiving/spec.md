# Auto Archiving

## Purpose
To ensure that OpenSpec changes are automatically archived and specs are synchronized when code is merged, reducing manual overhead and maintaining a clean change history.

## Requirements

### Requirement: Automated Archiving Trigger
The system SHALL automatically initiate the OpenSpec archiving process when a Pull Request is successfully merged into the main branch.

#### Scenario: Merge to Main
- **WHEN** a Pull Request is merged to `main`
- **AND** the PR contains changes to `openspec/changes/<change-name>/`
- **THEN** an archiving workflow triggers

### Requirement: Change Identification
The system SHALL identify the specific OpenSpec change associated with a merged Pull Request to ensure the correct artifacts are archived.

#### Scenario: Single Change Identification
- **WHEN** the archiving workflow runs
- **THEN** it scans for modified files in `openspec/changes/`
- **AND** identifies the change directory name
- **AND** ignores the `openspec/changes/archive/` directory

### Requirement: Archive Persistence
The system SHALL commit the results of the archiving process back to the repository to maintain a verifiable history of specification changes.

#### Scenario: Commit and Push
- **WHEN** the `openspec archive` command completes successfully
- **THEN** the system commits the changes (updated global specs and moved archive folder) to `main`
- **AND** the commit message includes `[skip ci]` to prevent recursive workflow triggers
