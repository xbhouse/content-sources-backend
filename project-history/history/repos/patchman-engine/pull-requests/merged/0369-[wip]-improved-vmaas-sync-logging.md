---
type: pull_request
number: 369
title: "[WIP] Improved vmaas-sync logging"
state: merged
author: josef-hak
created: 2020-10-08T11:18:58Z
updated: 2020-10-29T11:40:40Z
closed: 2020-10-08T12:26:05Z
merged: 2020-10-08T12:26:05Z
base_branch: master
head_branch: log-advisory-sync
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/369
---

# Pull Request #369: [WIP] Improved vmaas-sync logging

**Author**: @josef-hak
**Created**: October 08, 2020 at 11:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `log-advisory-sync`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on October 08, 2020 at 11:27 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=h1) Report
> Merging [#369](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/7feb84aa490002f5dabb05fd84c7aedbafc5a02f?el=desc) will **decrease** coverage by `0.02%`.
> The diff coverage is `75.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #369      +/-   ##
==========================================
- Coverage   61.73%   61.70%   -0.03%     
==========================================
  Files          57       57              
  Lines        2587     2588       +1     
==========================================
  Hits         1597     1597              
- Misses        758      759       +1     
  Partials      232      232              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `61.66% <50.00%> (+0.32%)` | :arrow_up: |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `79.68% <100.00%> (+0.65%)` | :arrow_up: |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `45.45% <0.00%> (-1.43%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `67.39% <0.00%> (-0.27%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=footer). Last update [7feb84a...7ab8bd6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/369?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on October 08, 2020 at 12:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/369*
