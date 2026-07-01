---
type: pull_request
number: 78
title: "chore(deps): update module github.com/prometheus/common to v0.67.5 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-11-18T16:44:02Z
updated: 2026-01-07T08:42:37Z
closed: 2026-01-07T08:42:37Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-prometheus-common-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/78
---

# Pull Request #78: chore(deps): update module github.com/prometheus/common to v0.67.5 - autoclosed

**Author**: @red-hat-konflux
**Created**: November 18, 2025 at 04:44 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | `v0.67.2` -> `v0.67.5` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fcommon/v0.67.5?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fcommon/v0.67.2/v0.67.5?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.67.5`](https://redirect.github.com/prometheus/common/releases/tag/v0.67.5)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.4...v0.67.5)

#### What's Changed

- build(deps): bump golang.org/x/oauth2 from 0.32.0 to 0.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;871](https://redirect.github.com/prometheus/common/pull/871)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;866](https://redirect.github.com/prometheus/common/pull/866)
- build(deps): bump golang.org/x/net from 0.46.0 to 0.48.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;872](https://redirect.github.com/prometheus/common/pull/872)
- build(deps): bump google.golang.org/protobuf from 1.36.10 to 1.36.11 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;870](https://redirect.github.com/prometheus/common/pull/870)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.67.4...v0.67.5>

### [`v0.67.4`](https://redirect.github.com/prometheus/common/releases/tag/v0.67.4): / 2025-11-18

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.3...v0.67.4)

#### What's Changed

- chore: clean up golangci-lint configuration by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;782](https://redirect.github.com/prometheus/common/pull/782)
- chore: 'omitempty' to Oauth2 fields with type Secret to avoid requiring them by [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) in [#&#8203;864](https://redirect.github.com/prometheus/common/pull/864)
- chore: Add omitempty tag to all config fields by [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) in [#&#8203;865](https://redirect.github.com/prometheus/common/pull/865)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.67.3...v0.67.4>

### [`v0.67.3`](https://redirect.github.com/prometheus/common/releases/tag/v0.67.3): / 2025-11-18

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.2...v0.67.3)

#### What's Changed

- Support JWT Profile for Authorization Grant (RFC 7523 3.1) by [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) in [#&#8203;862](https://redirect.github.com/prometheus/common/pull/862)
- Config: remove outdated comment about HTTP/2 issues by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;863](https://redirect.github.com/prometheus/common/pull/863)

#### New Contributors

- [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) made their first contribution in [#&#8203;862](https://redirect.github.com/prometheus/common/pull/862)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.67.2...v0.67.3>

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

## Discussion

### Comment by @red-hat-konflux on January 05, 2026 at 04:40 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 3 additional dependencies were updated


Details:


| **Package**                  | **Change**               |
| :--------------------------- | :----------------------- |
| `golang.org/x/net`           | `v0.47.0` -> `v0.48.0`   |
| `golang.org/x/oauth2`        | `v0.32.0` -> `v0.34.0`   |
| `google.golang.org/protobuf` | `v1.36.10` -> `v1.36.11` |

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/78*
