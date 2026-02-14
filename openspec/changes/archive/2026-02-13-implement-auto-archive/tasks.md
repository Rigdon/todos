## 1. Workflow Implementation

- [x] 1.1 Create `.github/workflows/openspec-archive.yml` with `pull_request` closed trigger
- [x] 1.2 Implement change detection logic to find the active change directory
- [x] 1.3 Add `openspec archive` command with `--yes` flag
- [x] 1.4 Configure `git-auto-commit-action` to push changes back to `main` with `[skip ci]`

## 2. Verification

- [x] 2.1 Verify workflow syntax using `act` (if available) or online validator
- [x] 2.2 Verify change detection logic with a local script test
