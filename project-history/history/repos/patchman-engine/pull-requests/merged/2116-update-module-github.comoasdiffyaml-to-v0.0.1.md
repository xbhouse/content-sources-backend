---
type: pull_request
number: 2116
title: "Update module github.com/oasdiff/yaml to v0.0.1"
state: merged
author: red-hat-konflux
created: 2026-03-23T09:15:09Z
updated: 2026-03-23T13:16:42Z
closed: 2026-03-23T09:15:31Z
merged: 2026-03-23T09:15:31Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-oasdiff-yaml-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2116
---

# Pull Request #2116: Update module github.com/oasdiff/yaml to v0.0.1

**Author**: @red-hat-konflux
**Created**: March 23, 2026 at 09:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-oasdiff-yaml-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/oasdiff/yaml](https://redirect.github.com/oasdiff/yaml) | `v0.0.0-20250309154309-f31be36b4037` -> `v0.0.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2foasdiff%2fyaml/v0.0.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2foasdiff%2fyaml/v0.0.0-20250309154309-f31be36b4037/v0.0.1?slim=true) |

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

### Comment by @red-hat-konflux on March 23, 2026 at 09:15 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                | **Change**                                       |
| :------------------------- | :----------------------------------------------- |
| `github.com/oasdiff/yaml3` | `v0.0.0-20250309153720-d2182401db90` -> `v0.0.1` |

### Comment by @github-actions on March 23, 2026 at 09:15 AM UTC

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

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2116*
