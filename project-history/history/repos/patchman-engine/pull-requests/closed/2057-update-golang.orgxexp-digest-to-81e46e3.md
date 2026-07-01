---
type: pull_request
number: 2057
title: "Update golang.org/x/exp digest to 81e46e3"
state: closed
author: red-hat-konflux
created: 2026-02-16T04:54:08Z
updated: 2026-03-05T12:08:07Z
closed: 2026-03-05T12:07:55Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-exp-digest
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2057
---

# Pull Request #2057: Update golang.org/x/exp digest to 81e46e3

**Author**: @red-hat-konflux
**Created**: February 16, 2026 at 04:54 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-exp-digest`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| golang.org/x/exp | require | digest | `716be56` -> `81e46e3` |

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

### Comment by @red-hat-konflux on February 16, 2026 at 04:54 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 5 additional dependencies were updated
- The `go` directive was updated for compatibility reasons


Details:


| **Package**           | **Change**             |
| :-------------------- | :--------------------- |
| `go`                  | `1.24.10` -> `1.25.0`  |
| `golang.org/x/crypto` | `v0.47.0` -> `v0.48.0` |
| `golang.org/x/mod`    | `v0.32.0` -> `v0.33.0` |
| `golang.org/x/net`    | `v0.49.0` -> `v0.50.0` |
| `golang.org/x/text`   | `v0.33.0` -> `v0.34.0` |
| `golang.org/x/tools`  | `v0.41.0` -> `v0.42.0` |

### Comment by @github-actions on February 16, 2026 at 04:54 AM UTC

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

### Comment by @patchkez on February 17, 2026 at 09:25 AM UTC

/retest

### Comment by @patchkez on February 17, 2026 at 09:36 AM UTC

/retest

### Comment by @red-hat-konflux on February 18, 2026 at 06:08 PM UTC

### Edited/Blocked Notification

Renovate will not automatically rebase this PR, because it does not recognize the last commit author and assumes somebody else may have edited the PR.

You can manually request rebase by checking the rebase/retry box above.

 ⚠️ **Warning**: custom changes will be lost.

### Comment by @TenSt on March 05, 2026 at 12:08 PM UTC

Closed in favor of https://github.com/RedHatInsights/patchman-engine/pull/2089

---

## Reviews

### Review by @TenSt - Dismissed on February 16, 2026 at 05:31 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2057*
