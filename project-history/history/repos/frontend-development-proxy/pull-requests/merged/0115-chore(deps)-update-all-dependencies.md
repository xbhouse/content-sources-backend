---
type: pull_request
number: 115
title: "chore(deps): update all dependencies"
state: merged
author: red-hat-konflux
created: 2026-03-10T17:35:55Z
updated: 2026-03-11T12:14:54Z
closed: 2026-03-11T12:14:50Z
merged: 2026-03-11T12:14:50Z
base_branch: main
head_branch: konflux/mintmaker/main/all-deps
labels: ["dependencies"]
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/115
---

# Pull Request #115: chore(deps): update all dependencies

**Author**: @red-hat-konflux
**Created**: March 10, 2026 at 05:35 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `konflux/mintmaker/main/all-deps`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/googleapis/gax-go/v2](https://redirect.github.com/googleapis/gax-go) | `v2.17.0` -> `v2.18.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgoogleapis%2fgax-go%2fv2/v2.18.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgoogleapis%2fgax-go%2fv2/v2.17.0/v2.18.0?slim=true) |
| [google.golang.org/api](https://redirect.github.com/googleapis/google-api-go-client) | `v0.270.0` -> `v0.271.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fapi/v0.271.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fapi/v0.270.0/v0.271.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>googleapis/gax-go (github.com/googleapis/gax-go/v2)</summary>

### [`v2.18.0`](https://redirect.github.com/googleapis/gax-go/releases/tag/v2.18.0): v2: v2.18.0

[Compare Source](https://redirect.github.com/googleapis/gax-go/compare/v2.17.0...v2.18.0)

##### Features

- move gax-go to use 1.25 as the lower bound of support ([#&#8203;469](https://redirect.github.com/googleapis/gax-go/issues/469)) ([01594ca5](https://redirect.github.com/googleapis/gax-go/commit/01594ca5))

- add callctx telemetry helpers ([#&#8203;472](https://redirect.github.com/googleapis/gax-go/issues/472)) ([fa319ffc](https://redirect.github.com/googleapis/gax-go/commit/fa319ffc))

</details>

<details>
<summary>googleapis/google-api-go-client (google.golang.org/api)</summary>

### [`v0.271.0`](https://redirect.github.com/googleapis/google-api-go-client/releases/tag/v0.271.0)

[Compare Source](https://redirect.github.com/googleapis/google-api-go-client/compare/v0.270.0...v0.271.0)

##### Features

- **all:** Auto-regenerate discovery clients ([#&#8203;3532](https://redirect.github.com/googleapis/google-api-go-client/issues/3532)) ([ccff5b3](https://redirect.github.com/googleapis/google-api-go-client/commit/ccff5b35c0d730214473de122dcb96b110be0029))

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://redirect.github.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6WyJkZXBlbmRlbmNpZXMiXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on March 10, 2026 at 05:35 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                                 | **Change**                                                                   |
| :------------------------------------------ | :--------------------------------------------------------------------------- |
| `google.golang.org/genproto/googleapis/api` | `v0.0.0-20260128011058-8636f8732409` -> `v0.0.0-20260217215200-42d3e9bedb6d` |

---

## Reviews

### Review by @Hyperkid123 - Approved on March 11, 2026 at 12:14 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/115*
