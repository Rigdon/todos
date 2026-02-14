## Context

The current CI/CD pipeline uses mutable tags (e.g., `v4`, `v5`) for GitHub Actions. This creates a security vulnerability where a compromised tag could execute malicious code. Additionally, the `api-lint` workflow re-downloads npm dependencies on every run, slowing down feedback. We also lack automated dependency updates, leading to stale and potentially vulnerable dependencies.

## Goals / Non-Goals

**Goals:**
- Pin all GitHub Actions to immutable full-length commit SHAs.
- Enable caching for `npm` dependencies in the `api-lint` workflow.
- Configure automated weekly dependency updates for Actions, Go modules, and npm using Dependabot.

**Non-Goals:**
- Migrating to a different CI provider.
- changing the actual linting or testing logic/tools.

## Decisions

### 1. Pinning Strategy
- **Decision**: Use full-length 40-character SHA-1 hashes for all GitHub Actions.
- **Rationale**: SHA-1 hashes are immutable. Tags can be moved. Pinning to SHA guarantees the code running in CI is exactly what was audited (or at least what was present at the time of pinning).
- **Specific SHAs**:
  - `actions/checkout`: `34e114876b0b11c390a56381ad16ebd13914f8d5` (v4.3.1)
  - `actions/setup-node`: `49933ea5288caeca8642d1e84afbd3f7d6820020` (v4.4.0)
  - `actions/setup-go`: `40f1582b2485089dde7abd97c1529aa768e1baff` (v5.6.0)
  - `actions/dependency-review-action`: `3c4e3dcb1aa7874d2c16be7d79418e9b7efd6261` (v4.8.2)

### 2. Caching Strategy
- **Decision**: Use `actions/setup-node` built-in caching.
- **Rationale**: `actions/setup-node` provides a `cache: 'npm'` option that automatically handles caching the `~/.npm` directory based on `package-lock.json`. This is simpler than manually using `actions/cache`.

### 3. Automated Updates
- **Decision**: Use GitHub Dependabot.
- **Rationale**: Native integration with GitHub, supports all our ecosystems (npm, go, github-actions), and is free for public repositories (and included in GitHub Enterprise).

## Risks / Trade-offs

- **Risk**: Pinned actions effectively "freeze" the CI. Security patches in actions won't be applied automatically.
  - **Mitigation**: Configure Dependabot to open PRs for GitHub Action updates. This turns implicit updates (moving tags) into explicit, reviewable PRs.
- **Risk**: Cache corruption or staleness.
  - **Mitigation**: `actions/setup-node` uses the lockfile as a cache key. Any change to `package-lock.json` busts the cache.
