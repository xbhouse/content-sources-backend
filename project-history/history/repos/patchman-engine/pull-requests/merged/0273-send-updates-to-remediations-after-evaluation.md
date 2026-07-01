---
type: pull_request
number: 273
title: "Send updates to remediations after evaluation"
state: merged
author: semtexzv
created: 2020-06-17T05:19:48Z
updated: 2021-03-16T09:28:29Z
closed: 2020-06-22T08:35:41Z
merged: 2020-06-22T08:35:41Z
base_branch: master
head_branch: remediations-updates
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/273
---

# Pull Request #273: Send updates to remediations after evaluation

**Author**: @semtexzv
**Created**: June 17, 2020 at 05:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `remediations-updates`

## Description

Depends on https://github.com/RedHatInsights/e2e-deploy/pull/1853

---

## Discussion

### Comment by @codecov-commenter on June 17, 2020 at 05:27 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=h1) Report
> Merging [#273](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ae465739da754dcb1e40adf5ccb69e40d330da2b&el=desc) will **decrease** coverage by `0.34%`.
> The diff coverage is `37.93%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #273      +/-   ##
==========================================
- Coverage   62.15%   61.81%   -0.35%     
==========================================
  Files          49       51       +2     
  Lines        2328     2362      +34     
==========================================
+ Hits         1447     1460      +13     
- Misses        704      720      +16     
- Partials      177      182       +5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/message.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `52.63% <0.00%> (-9.87%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.11% <25.00%> (-0.85%)` | :arrow_down: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `62.50% <62.50%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `64.54% <0.00%> (-0.55%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `90.32% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=footer). Last update [ae46573...a04db68](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/273?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on June 18, 2020 at 07:48 AM UTC

Thanks, just some ideas and questions added.

### Review by @semtexzv - Commented on June 19, 2020 at 10:06 AM UTC

### Review by @josef-hak - Approved on June 22, 2020 at 08:35 AM UTC

Ok, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/273*
