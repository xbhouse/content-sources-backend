---
type: pull_request
number: 1763
title: "Update module github.com/prometheus/client_golang to v1.23.0"
state: merged
author: red-hat-konflux
created: 2025-08-03T12:49:34Z
updated: 2025-08-03T17:40:12Z
closed: 2025-08-03T17:40:10Z
merged: 2025-08-03T17:40:10Z
base_branch: security-compliance
head_branch: konflux/mintmaker/security-compliance/github.com-prometheus-client_golang-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1763
---

# Pull Request #1763: Update module github.com/prometheus/client_golang to v1.23.0

**Author**: @red-hat-konflux
**Created**: August 03, 2025 at 12:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `security-compliance` ← **Head**: `konflux/mintmaker/security-compliance/github.com-prometheus-client_golang-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/client_golang](https://redirect.github.com/prometheus/client_golang) | `v1.22.0` -> `v1.23.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fclient_golang/v1.23.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fclient_golang/v1.22.0/v1.23.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/client_golang (github.com/prometheus/client_golang)</summary>

### [`v1.23.0`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.23.0): - 2025-07-30

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.22.0...v1.23.0)

- \[CHANGE] Minimum required Go version is now 1.23, only the two latest Go versions are supported from now on. [#&#8203;1812](https://redirect.github.com/prometheus/client_golang/issues/1812)
- \[FEATURE] Add WrapCollectorWith and WrapCollectorWithPrefix [#&#8203;1766](https://redirect.github.com/prometheus/client_golang/issues/1766)
- \[FEATURE] Add exemplars for native histograms [#&#8203;1686](https://redirect.github.com/prometheus/client_golang/issues/1686)
- \[ENHANCEMENT] exp/api: Bubble up status code from writeResponse [#&#8203;1823](https://redirect.github.com/prometheus/client_golang/issues/1823)
- \[ENHANCEMENT] collector/go: Update runtime metrics for Go v1.23 and v1.24 [#&#8203;1833](https://redirect.github.com/prometheus/client_golang/issues/1833)
- \[BUGFIX] exp/api: client prompt return on context cancellation [#&#8203;1729](https://redirect.github.com/prometheus/client_golang/issues/1729)

<details>
<summary>All Changes</summary>
* Update example test by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1770
* build(deps): bump golang.org/x/net from 0.34.0 to 0.36.0 in /tutorials/whatsup by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1776
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1771
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1778
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1779
* build(deps): bump google.golang.org/protobuf from 1.36.5 to 1.36.6 in /exp by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1782
* build(deps): bump github.com/prometheus/common from 0.62.0 to 0.63.0 in /exp by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1781
* build(deps): bump github.com/prometheus/common from 0.62.0 to 0.63.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1783
* build(deps): bump google.golang.org/protobuf from 1.36.5 to 1.36.6 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1784
* build(deps): bump github.com/prometheus/procfs from 0.15.1 to 0.16.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1786
* chore: Upgrade golangci-lint to v2 by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1789
* build(deps): bump the github-actions group across 1 directory with 3 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1790
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1791
* Add `WrapCollectorWith` and `WrapCollectorWithPrefix` by @&#8203;colega in https://github.com/prometheus/client_golang/pull/1766
* feat(github-actions): add Go file change detection to golangci-lint workflow by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1794
* chore(ci): Fix flaky tests by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1795
* build(deps): bump golang.org/x/net from 0.36.0 to 0.38.0 in /tutorials/whatsup by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1799
* test(registry): Add goleak-based goroutine leak detection by @&#8203;surinkim in https://github.com/prometheus/client_golang/pull/1797
* build(deps): bump go.uber.org/goleak from 1.2.0 to 1.3.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1806
* build(deps): bump the github-actions group with 2 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1804
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1809
* Add exemplars for native histograms by @&#8203;shivanthzen in https://github.com/prometheus/client_golang/pull/1686
* build(deps): bump golang.org/x/sys from 0.30.0 to 0.32.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1807
* build(deps): bump github.com/prometheus/client_model from 0.6.1 to 0.6.2 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1805
* build(deps): bump github.com/prometheus/procfs from 0.16.0 to 0.16.1 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1808
* build(deps): bump golang.org/x/net from 0.35.0 to 0.38.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1800
* Update supported Go versions by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1812
* Cleaup Go modules by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1813
* fix: client prompt return on context cancellation by @&#8203;umegbewe in https://github.com/prometheus/client_golang/pull/1729
* Simplify buf binary install by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1814
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1815
* build(deps): bump the github-actions group with 5 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1817
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1821
* exp/api: Bubble up status code from writeResponse by @&#8203;saswatamcode in https://github.com/prometheus/client_golang/pull/1823
* build(deps): bump github.com/prometheus/common from 0.64.0 to 0.65.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1827
* build(deps): bump github.com/prometheus/common from 0.64.0 to 0.65.0 in /exp by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1828
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1831
* Update runtime metrics for Go v1.23 and v1.24 by @&#8203;aknuds1 in https://github.com/prometheus/client_golang/pull/1833
* Fix `errNotImplemented` reference by @&#8203;aknuds1 in https://github.com/prometheus/client_golang/pull/1835
* build(deps): bump the github-actions group with 3 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1826
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1832
* Cut v1.23.0-rc.0 by @&#8203;vesari in https://github.com/prometheus/client_golang/pull/1837
* cut v1.23.0-rc.1 by @&#8203;vesari in https://github.com/prometheus/client_golang/pull/1842

</details>

#### New Contributors
* @&#8203;surinkim made their first contributi[https://github.com/prometheus/client_golang/pull/1797](https://redirect.github.com/prometheus/client_golang/pull/1797)l/1797
* @&#8203;umegbewe made their first contributi[https://github.com/prometheus/client_golang/pull/1729](https://redirect.github.com/prometheus/client_golang/pull/1729)l/1729
* @&#8203;aknuds1 made their first contributi[https://github.com/prometheus/client_golang/pull/1833](https://redirect.github.com/prometheus/client_golang/pull/1833)l/1833

**Full Changelog**: https://github.com/prometheus/client\_golang/compare/v1.22.0...v1.23.0

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuNy4wLXJwbSIsInRhcmdldEJyYW5jaCI6InNlY3VyaXR5LWNvbXBsaWFuY2UiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on August 03, 2025 at 12:49 PM UTC

Commits missing Jira IDs:
3c7705e377513c91628e1cd6b7c83c8180dcc835


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1763*
