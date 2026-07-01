---
type: pull_request
number: 245
title: "master -> stable"
state: merged
author: josef-hak
created: 2020-05-04T07:00:55Z
updated: 2020-05-04T07:42:18Z
closed: 2020-05-04T07:42:18Z
merged: 2020-05-04T07:42:18Z
base_branch: stable
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/245
---

# Pull Request #245: master -> stable

**Author**: @josef-hak
**Created**: May 04, 2020 at 07:00 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `stable` ← **Head**: `master`

## Description

- vmaas call errors logging
- kafka errors fixed log formatting
- kafka error counters
- updated default list size (from 25 to 20)
- added indexes to db schema
- ignored unknown advisories in vmaas response during evaluation


---

## Discussion

### Comment by @codecov-io on May 04, 2020 at 07:12 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=h1) Report
> Merging [#245](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=desc) into [stable](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/68533bdc574b470e3a123da640568d76f1c721db&el=desc) will **increase** coverage by `0.21%`.
> The diff coverage is `72.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           stable     #245      +/-   ##
==========================================
+ Coverage   62.67%   62.88%   +0.21%     
==========================================
  Files          46       47       +1     
  Lines        2095     2099       +4     
==========================================
+ Hits         1313     1320       +7     
+ Misses        629      627       -2     
+ Partials      153      152       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `35.55% <0.00%> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWV0cmljcy5nbw==) | `20.00% <20.00%> (ø)` | |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `16.36% <33.33%> (-0.88%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.30% <81.81%> (+1.08%)` | :arrow_up: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `73.84% <86.66%> (+2.87%)` | :arrow_up: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `73.68% <100.00%> (+1.46%)` | :arrow_up: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `31.16% <100.00%> (+0.90%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=footer). Last update [68533bd...4833405](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/245?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on May 04, 2020 at 07:37 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/245*
