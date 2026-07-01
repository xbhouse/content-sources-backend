---
type: pull_request
number: 2084
title: "Update module github.com/prometheus/procfs to v0.20.1"
state: closed
author: red-hat-konflux
created: 2026-03-02T05:23:14Z
updated: 2026-03-05T17:11:36Z
closed: 2026-03-05T15:27:37Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-procfs-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2084
---

# Pull Request #2084: Update module github.com/prometheus/procfs to v0.20.1

**Author**: @red-hat-konflux
**Created**: March 02, 2026 at 05:23 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-procfs-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/prometheus/procfs](https://redirect.github.com/prometheus/procfs) | `v0.19.2` -> `v0.20.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fprocfs/v0.20.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fprocfs/v0.19.2/v0.20.1?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/procfs (github.com/prometheus/procfs)</summary>

### [`v0.20.1`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.20.1)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.20.0...v0.20.1)

#### What's Changed

- nvme: Parse NVMe namespace details  by [@&#8203;ShashwatHiregoudar](https://redirect.github.com/ShashwatHiregoudar) in [#&#8203;765](https://redirect.github.com/prometheus/procfs/pull/765)
- Fix bcachefs parsing by [@&#8203;ananthb](https://redirect.github.com/ananthb) in [#&#8203;789](https://redirect.github.com/prometheus/procfs/pull/789)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;788](https://redirect.github.com/prometheus/procfs/pull/788)
- Update sysfs/class\_thermal: continue on EINVAL in parseClassThermalZone to ignore only invalid thermal zones which raise "invalid argument" by [@&#8203;ccastiglione-reply](https://redirect.github.com/ccastiglione-reply) in [#&#8203;763](https://redirect.github.com/prometheus/procfs/pull/763)

#### New Contributors

- [@&#8203;ccastiglione-reply](https://redirect.github.com/ccastiglione-reply) made their first contribution in [#&#8203;763](https://redirect.github.com/prometheus/procfs/pull/763)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.20.0...v0.20.1>

### [`v0.20.0`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.20.0)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.19.2...v0.20.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;747](https://redirect.github.com/prometheus/procfs/pull/747)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;766](https://redirect.github.com/prometheus/procfs/pull/766)
- build(deps): bump golang.org/x/sync from 0.17.0 to 0.19.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;773](https://redirect.github.com/prometheus/procfs/pull/773)
- build(deps): bump golang.org/x/sys from 0.37.0 to 0.39.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;772](https://redirect.github.com/prometheus/procfs/pull/772)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;774](https://redirect.github.com/prometheus/procfs/pull/774)
- Fix /proc/interrupts by [@&#8203;ffyuanda](https://redirect.github.com/ffyuanda) in [#&#8203;775](https://redirect.github.com/prometheus/procfs/pull/775)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;778](https://redirect.github.com/prometheus/procfs/pull/778)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;779](https://redirect.github.com/prometheus/procfs/pull/779)
- Migrate to GitHub actions by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;780](https://redirect.github.com/prometheus/procfs/pull/780)
- build(deps): bump golang.org/x/sys from 0.40.0 to 0.41.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;782](https://redirect.github.com/prometheus/procfs/pull/782)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;785](https://redirect.github.com/prometheus/procfs/pull/785)
- bcachefs support by [@&#8203;ananthb](https://redirect.github.com/ananthb) in [#&#8203;750](https://redirect.github.com/prometheus/procfs/pull/750)
- build(deps): bump actions/checkout from 6.0.1 to 6.0.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;781](https://redirect.github.com/prometheus/procfs/pull/781)
- feat: parse capabilities in /proc/pid/status by [@&#8203;biscout42](https://redirect.github.com/biscout42) in [#&#8203;784](https://redirect.github.com/prometheus/procfs/pull/784)
- class\_cooling\_device: ignore EINVAL (etc) when reading files. by [@&#8203;malcolmr](https://redirect.github.com/malcolmr) in [#&#8203;783](https://redirect.github.com/prometheus/procfs/pull/783)
- Add type and name to the DRM parser class by [@&#8203;Deezzir](https://redirect.github.com/Deezzir) in [#&#8203;672](https://redirect.github.com/prometheus/procfs/pull/672)

#### New Contributors

- [@&#8203;ffyuanda](https://redirect.github.com/ffyuanda) made their first contribution in [#&#8203;775](https://redirect.github.com/prometheus/procfs/pull/775)
- [@&#8203;ananthb](https://redirect.github.com/ananthb) made their first contribution in [#&#8203;750](https://redirect.github.com/prometheus/procfs/pull/750)
- [@&#8203;biscout42](https://redirect.github.com/biscout42) made their first contribution in [#&#8203;784](https://redirect.github.com/prometheus/procfs/pull/784)
- [@&#8203;malcolmr](https://redirect.github.com/malcolmr) made their first contribution in [#&#8203;783](https://redirect.github.com/prometheus/procfs/pull/783)
- [@&#8203;Deezzir](https://redirect.github.com/Deezzir) made their first contribution in [#&#8203;672](https://redirect.github.com/prometheus/procfs/pull/672)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.19.2...v0.20.0>

</details>

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

### Comment by @red-hat-konflux on March 02, 2026 at 05:23 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- The `go` directive was updated for compatibility reasons


Details:


| **Package** | **Change**            |
| :---------- | :-------------------- |
| `go`        | `1.24.10` -> `1.25.0` |

### Comment by @github-actions on March 02, 2026 at 05:23 AM UTC

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

Because you closed this PR without merging, Renovate will ignore this update (`v0.20.1`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2084*
