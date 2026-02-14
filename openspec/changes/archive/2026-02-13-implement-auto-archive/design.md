## Context

Currently, archiving an OpenSpec change is a manual process performed after a PR is merged. This involves running `openspec archive` locally and pushing a new commit. This is prone to human error (forgetting to archive, archiving incomplete changes) and toil.

## Goals / Non-Goals

**Goals:**
- Automatically archive OpenSpec changes when a PR is merged to `main`.
- Sync delta specs to the main library as part of the archive process.
- Commit the archive results back to the repository without triggering recursive workflows.

**Non-Goals:**
- Automating the `apply` or `sync` steps *before* merge (sync happens at archive time).
- validating specs in this workflow (validation should happen on PR open/synchronize).

## Decisions

### 1. Workflow Trigger
- **Decision**: Trigger on `pull_request` event with type `closed`.
- **Condition**: `if: github.event.pull_request.merged == true`
- **Rationale**: We only want to archive when code is actually merged into the main branch.

### 2. Change Detection
- **Decision**: Use a shell script to identify the modified change directory.
- **Logic**: Look for changes in `openspec/changes/*` that are NOT in `openspec/changes/archive/`.
- **Constraint**: Assumes one change per PR (standard convention). If multiple are found, fail or process the first.

### 3. Execution Environment
- **Tool**: `@fission-ai/openspec` CLI.
- **Command**: `openspec archive <change-name> --yes`
- **Permissions**: `contents: write` (needed to push the archive commit).

### 4. Commit Strategy
- **Action**: `stefanzweifel/git-auto-commit-action`
- **Message**: `docs(openspec): auto-archive <change-name> [skip ci]`
- **Rationale**: `[skip ci]` is critical to prevent the archive commit from triggering a new build/test cycle.

## Risks / Trade-offs

- **Risk**: Recursive workflow runs.
  - **Mitigation**: Use `[skip ci]` in the commit message.
- **Risk**: Malformed specs being archived.
  - **Mitigation**: Rely on the existing `api-lint` or separate `spec-lint` workflow to block the PR if specs are invalid.
