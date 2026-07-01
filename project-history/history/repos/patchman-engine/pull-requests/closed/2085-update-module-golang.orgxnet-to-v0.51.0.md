---
type: pull_request
number: 2085
title: "Update module golang.org/x/net to v0.51.0"
state: closed
author: red-hat-konflux
created: 2026-03-02T09:18:17Z
updated: 2026-03-05T17:11:38Z
closed: 2026-03-05T15:27:44Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-net-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2085
---

# Pull Request #2085: Update module golang.org/x/net to v0.51.0

**Author**: @red-hat-konflux
**Created**: March 02, 2026 at 09:18 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-net-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| golang.org/x/net | `v0.49.0` -> `v0.51.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/golang.org%2fx%2fnet/v0.51.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/golang.org%2fx%2fnet/v0.49.0/v0.51.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on March 02, 2026 at 09:18 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- The `go` directive was updated for compatibility reasons


Details:


| **Package** | **Change**            |
| :---------- | :-------------------- |
| `go`        | `1.24.10` -> `1.25.0` |

### Comment by @github-actions on March 02, 2026 at 09:18 AM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** ⚪ **NONE**

No SC Environment-specific impacts detected in this PR.

<details>
<summary>What was checked</summary>

This PR was automatically scanned for:
- Database migrations
- ClowdApp configuration changes
- Kessel integration changes
- AWS service integrations (S3, RDS, ElastiCache)
- Kafka topic changes
- Secrets management changes
- External dependencies
</details>

### Comment by @TenSt on March 05, 2026 at 03:27 PM UTC

Closing in favor of https://github.com/RedHatInsights/patchman-engine/pull/2089

### Comment by @red-hat-konflux on March 05, 2026 at 05:11 PM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`v0.51.0`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2085*
