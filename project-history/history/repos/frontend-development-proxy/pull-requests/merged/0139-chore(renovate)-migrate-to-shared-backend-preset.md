---
type: pull_request
number: 139
title: "chore(renovate): migrate to shared backend preset"
state: merged
author: platex-rehor-bot
created: 2026-05-26T18:48:43Z
updated: 2026-05-26T21:44:17Z
closed: 2026-05-26T21:32:17Z
merged: 2026-05-26T21:32:17Z
base_branch: main
head_branch: bot/RHCLOUD-47826
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/139
---

# Pull Request #139: chore(renovate): migrate to shared backend preset

**Author**: @platex-rehor-bot
**Created**: May 26, 2026 at 06:48 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `bot/RHCLOUD-47826`

## Description

### Description

Migrate `frontend-development-proxy` from its current inline Renovate configuration to the shared backend preset from `RedHatInsights/shared-workflows`.

[RHCLOUD-47826](https://issues.redhat.com/browse/RHCLOUD-47826)

**What changed:**
- Replaced inline `renovate.json` (deprecated `config:base`, custom package rules, tekton schedule) with the shared backend preset
- The shared preset already includes: Go module grouping, Docker image grouping, UBI9 base image pinning, Tekton schedule, semantic commits

---

### How to test locally

No local testing needed — this is a Renovate configuration change. Verify by checking that Renovate picks up the new preset and creates PRs according to the shared backend rules.

---

### Anything reviewers should know?

- The deprecated `config:base` preset is replaced by the shared preset (which uses `config:recommended`)
- Custom package rules (major disabled, minor/patch grouping) are superseded by the shared preset's rules
- `baseBranches` override removed — Renovate defaults to the repo's default branch
- Prerequisite RHCLOUD-47802 (create shared presets) is already merged

---

### Checklist
- [x] Tests: new/updated tests cover the change
- [x] API: spec updated if endpoints changed (or N/A)
- [x] Migrations: backwards compatible if schema changed (or N/A)
- [x] Dependencies: no known impact to dependent services
- [x] Security: reviewed against [secure coding checklist](https://github.com/RedHatInsights/secure-coding-checklist) (or N/A)

### AI disclosure
Assisted by: Claude Code

[RHCLOUD-47826]: https://redhat.atlassian.net/browse/RHCLOUD-47826?atlOrigin=eyJpIjoiNWRkNTljNzYxNjVmNDY3MDlhMDU5Y2ZhYzA5YTRkZjUiLCJwIjoiZ2l0aHViLWNvbS1KU1cifQ

---

## Reviews

### Review by @karelhala - Approved on May 26, 2026 at 09:32 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/139*
