## Why

Currently, archiving an OpenSpec change is a manual process that involves merging delta specs and moving folders. Doing this inside a feature PR creates review noise, and doing it manually after merge is prone to human error. Automating this ensures documentation is always consistent with the code in `main` without manual intervention.

## What Changes

- **Automated Workflow**: A new GitHub Action will run when a PR is merged to `main`.
- **Change Detection**: The workflow will automatically identify the active OpenSpec change from the PR.
- **Auto-Archiving**: The workflow will execute `openspec archive` and commit the results back to `main`.

## Capabilities

### New Capabilities
- `auto-archiving`: Automated management of the OpenSpec lifecycle.

### Modified Capabilities
<!-- No requirement changes to existing capabilities. -->

## Impact

- **DX**: Developers no longer need to manually archive changes or worry about spec synchronization conflicts.
- **CI**: A new workflow will run on PR merges.
