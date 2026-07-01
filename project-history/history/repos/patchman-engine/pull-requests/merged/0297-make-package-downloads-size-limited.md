---
type: pull_request
number: 297
title: "Make package downloads size limited"
state: merged
author: semtexzv
created: 2020-07-28T14:31:25Z
updated: 2021-03-16T09:26:05Z
closed: 2020-07-29T08:55:33Z
merged: 2020-07-29T08:55:33Z
base_branch: master
head_branch: pkg-download-limit
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/297
---

# Pull Request #297: Make package downloads size limited

**Author**: @semtexzv
**Created**: July 28, 2020 at 02:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-download-limit`

## Description

By repeatedly slicing the packages array. Should only happen when we have too many packages in a page of advisories

---

## Discussion

### Comment by @codecov-commenter on July 29, 2020 at 08:55 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=h1) Report
> Merging [#297](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bd5c8be8072e7f5c21760c1db9b8ea96fdc60a37&el=desc) will **increase** coverage by `0.01%`.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #297      +/-   ##
==========================================
+ Coverage   61.76%   61.78%   +0.01%     
==========================================
  Files          52       52              
  Lines        2584     2593       +9     
==========================================
+ Hits         1596     1602       +6     
- Misses        784      786       +2     
- Partials      204      205       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.29% <40.00%> (-0.47%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `65.28% <70.00%> (+1.18%)` | :arrow_up: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `71.83% <100.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `32.05% <100.00%> (+0.88%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=footer). Last update [bd5c8be...90b5131](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/297?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on July 28, 2020 at 03:08 PM UTC

### Review by @josef-hak - Changes Requested on July 28, 2020 at 03:14 PM UTC

pipeline failed

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/297*
