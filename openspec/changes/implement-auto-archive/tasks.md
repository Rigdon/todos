## 1. Workflow Implementation

- [ ] 1.1 Create `.github/workflows/openspec-archive.yml` with `pull_request` closed trigger
- [ ] 1.2 Implement change detection logic to find the active change directory
- [ ] 1.3 Add `openspec archive` command with `--yes` flag
- [ ] 1.4 Configure `git-auto-commit-action` to push changes back to `main` with `[skip ci]`

## 2. Verification

- [ ] 2.1 Verify workflow syntax using `act` (if available) or online validator
- [ ] 2.2 Verify change detection logic with a local script test
