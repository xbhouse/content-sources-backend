---
type: pull_request
number: 81
title: "Update module github.com/go-jose/go-jose/v4 to v4.1.3"
state: merged
author: red-hat-konflux
created: 2025-11-24T12:42:40Z
updated: 2025-12-03T08:31:25Z
closed: 2025-12-03T08:31:25Z
merged: 2025-12-03T08:31:25Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-go-jose-go-jose-v4-4.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/81
---

# Pull Request #81: Update module github.com/go-jose/go-jose/v4 to v4.1.3

**Author**: @red-hat-konflux
**Created**: November 24, 2025 at 12:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-go-jose-go-jose-v4-4.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/go-jose/go-jose/v4](https://redirect.github.com/go-jose/go-jose) | `v4.1.2` -> `v4.1.3` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-jose%2fgo-jose%2fv4/v4.1.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-jose%2fgo-jose%2fv4/v4.1.2/v4.1.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-jose/go-jose (github.com/go-jose/go-jose/v4)</summary>

### [`v4.1.3`](https://redirect.github.com/go-jose/go-jose/releases/tag/v4.1.3)

[Compare Source](https://redirect.github.com/go-jose/go-jose/compare/v4.1.2...v4.1.3)

This release drops Go 1.23 support as that Go release is no longer supported. With that, we can drop `x/crypto` and no longer have any external dependencies in go-jose outside of the standard library!

This release fixes a bug where a critical b64 header was ignored if in an unprotected header. It is now rejected instead of ignored.

#### What's Changed

- Remove Go 1.23 support by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;205](https://redirect.github.com/go-jose/go-jose/pull/205)
- Reject JWS with an unprotected critical b64 header by [@&#8203;mcpherrinm](https://redirect.github.com/mcpherrinm) in [#&#8203;210](https://redirect.github.com/go-jose/go-jose/pull/210)

**Full Changelog**: <https://github.com/go-jose/go-jose/compare/v4.1.2...v4.1.3>

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @Hyperkid123 - Approved on December 03, 2025 at 08:31 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/81*
