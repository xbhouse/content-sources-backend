---
type: pull_request
number: 130
title: "docs: add agent readiness documentation"
state: merged
author: platex-rehor-bot
created: 2026-04-15T11:17:39Z
updated: 2026-04-15T14:08:58Z
closed: 2026-04-15T14:08:58Z
merged: 2026-04-15T14:08:58Z
base_branch: main
head_branch: bot/RHCLOUD-46723
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/130
---

# Pull Request #130: docs: add agent readiness documentation

**Author**: @platex-rehor-bot
**Created**: April 15, 2026 at 11:17 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `bot/RHCLOUD-46723`

## Description

## Summary

- Added `AGENTS.md` with project overview, tech stack, conventions, common pitfalls, and docs index
- Added `CLAUDE.md` with `@AGENTS.md` import and build/test/lint commands
- Added `CONTRIBUTING.md` with development workflow, commit conventions, and code style guidelines
- Added `docs/ARCHITECTURE.md` documenting system design, request flow, and key design decisions
- Added `docs/security-guidelines.md` covering JWT handling, identity headers, proxy config, and container security
- Added `docs/testing-guidelines.md` with Go unit test patterns, integration testing, and route generation testing
- Updated `README.md` with links to all new documentation

All changes are additive — no existing content was modified or removed.

RHCLOUD-46723

## Test plan

- [ ] Verify all documentation files are accurate against the current codebase
- [ ] Verify README.md links point to correct files
- [ ] Verify Docker build still succeeds with no code changes

🤖 Generated with [Claude Code](https://claude.com/claude-code)

---

## Discussion

### Comment by @platex-rehor-bot on April 15, 2026 at 01:13 PM UTC

The commit on this PR is unsigned. The bot needs to rebase to re-sign it with GPG before this can be merged.

### Comment by @platex-rehor-bot on April 15, 2026 at 01:40 PM UTC

Re-signed the commit via `git rebase --force-rebase HEAD~1`. The commit is now GPG-signed. Force pushed to update the PR.

---

## Reviews

### Review by @Hyperkid123 - Approved on April 15, 2026 at 02:08 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/130*
