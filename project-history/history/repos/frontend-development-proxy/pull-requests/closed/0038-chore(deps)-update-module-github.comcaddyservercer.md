---
type: pull_request
number: 38
title: "chore(deps): update module github.com/caddyserver/certmagic to v0.25.1 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-10-27T12:15:08Z
updated: 2026-01-07T08:42:17Z
closed: 2026-01-07T08:42:15Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-caddyserver-certmagic-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/38
---

# Pull Request #38: chore(deps): update module github.com/caddyserver/certmagic to v0.25.1 - autoclosed

**Author**: @red-hat-konflux
**Created**: October 27, 2025 at 12:15 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-caddyserver-certmagic-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/caddyserver/certmagic](https://redirect.github.com/caddyserver/certmagic) | `v0.24.0` -> `v0.25.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fcaddyserver%2fcertmagic/v0.25.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fcaddyserver%2fcertmagic/v0.24.0/v0.25.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>caddyserver/certmagic (github.com/caddyserver/certmagic)</summary>

### [`v0.25.1`](https://redirect.github.com/caddyserver/certmagic/compare/v0.25.0...v0.25.1)

[Compare Source](https://redirect.github.com/caddyserver/certmagic/compare/v0.25.0...v0.25.1)

### [`v0.25.0`](https://redirect.github.com/caddyserver/certmagic/releases/tag/v0.25.0)

[Compare Source](https://redirect.github.com/caddyserver/certmagic/compare/v0.24.0...v0.25.0)

Adds support for disabling distributed solving, but still allows distributed solving of the http-01 challenge as long as the right ACME account is used.

#### What's Changed

- Implement precise lock lease renewal for storage backends that support lease renewal. by [@&#8203;zoltan-frm](https://redirect.github.com/zoltan-frm) in [#&#8203;347](https://redirect.github.com/caddyserver/certmagic/pull/347)

#### New Contributors

- [@&#8203;zoltan-frm](https://redirect.github.com/zoltan-frm) made their first contribution in [#&#8203;347](https://redirect.github.com/caddyserver/certmagic/pull/347)

**Full Changelog**: <https://github.com/caddyserver/certmagic/compare/v0.24.0...v0.25.0>

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

### Comment by @red-hat-konflux on November 07, 2025 at 12:17 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 5 additional dependencies were updated


Details:


| **Package**                      | **Change**             |
| :------------------------------- | :--------------------- |
| `github.com/caddyserver/zerossl` | `v0.1.3` -> `v0.1.4`   |
| `github.com/libdns/libdns`       | `v1.1.0` -> `v1.1.1`   |
| `github.com/miekg/dns`           | `v1.1.68` -> `v1.1.69` |
| `go.uber.org/zap`                | `v1.27.0` -> `v1.27.1` |
| `golang.org/x/net`               | `v0.47.0` -> `v0.48.0` |

### Comment by @Hyperkid123 on November 11, 2025 at 02:21 PM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/38*
