---
type: pull_request
number: 14
title: "Update module github.com/antlr4-go/antlr/v4 to v4.13.1"
state: merged
author: red-hat-konflux
created: 2025-08-03T22:27:24Z
updated: 2025-11-07T11:08:22Z
closed: 2025-11-07T11:08:21Z
merged: 2025-11-07T11:08:21Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-antlr4-go-antlr-v4-4.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/14
---

# Pull Request #14: Update module github.com/antlr4-go/antlr/v4 to v4.13.1

**Author**: @red-hat-konflux
**Created**: August 03, 2025 at 10:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-antlr4-go-antlr-v4-4.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/antlr4-go/antlr/v4](https://redirect.github.com/antlr4-go/antlr) | `v4.13.0` -> `v4.13.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fantlr4-go%2fantlr%2fv4/v4.13.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fantlr4-go%2fantlr%2fv4/v4.13.0/v4.13.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>antlr4-go/antlr (github.com/antlr4-go/antlr/v4)</summary>

### [`v4.13.1`](https://redirect.github.com/antlr4-go/antlr/releases/tag/v4.13.1): Release 4.13.1

[Compare Source](https://redirect.github.com/antlr4-go/antlr/compare/v4.13.0...v4.13.1)

The 4.13.1 release has:

- Some minor performance improvements
- An optional build tag -tags antlr.nomutex which allows a build without
  using mutexes, when the user knows they will not use it with multiple
  go routines. (Note, somewhat experimental)

Note that I did not remove the sort.Slices calls. They will not be experimental for long
and are implemented efficently. Please make a fork if you must use it without this
dependency. Maybe review whether you need to do so for the sake of one call.

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuOTAuMS1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYWluIiwibGFiZWxzIjpbXX0=-->


---

## Reviews

### Review by @Hyperkid123 - Approved on November 07, 2025 at 11:08 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/14*
