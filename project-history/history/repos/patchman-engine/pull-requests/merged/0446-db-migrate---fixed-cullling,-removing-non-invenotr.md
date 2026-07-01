---
type: pull_request
number: 446
title: "db migrate - Fixed cullling, removing non-invenotry system batch"
state: merged
author: josef-hak
created: 2020-12-11T12:44:41Z
updated: 2021-02-04T10:51:03Z
closed: 2021-01-08T09:50:40Z
merged: 2021-01-08T09:50:40Z
base_branch: master
head_branch: fix-culling
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/446
---

# Pull Request #446: db migrate - Fixed cullling, removing non-invenotry system batch

**Author**: @josef-hak
**Created**: December 11, 2020 at 12:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-culling`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on December 11, 2020 at 02:21 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=h1) Report
> Merging [#446](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=desc) (7a417c0) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a6cab408b1f32b1636734059f6595d254b78fc1c?el=desc) (a6cab40) will **increase** coverage by `0.01%`.
> The diff coverage is `59.09%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #446      +/-   ##
==========================================
+ Coverage   62.53%   62.55%   +0.01%     
==========================================
  Files          60       60              
  Lines        2725     2745      +20     
==========================================
+ Hits         1704     1717      +13     
- Misses        781      788       +7     
  Partials      240      240              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `30.86% <0.00%> (-0.39%)` | :arrow_down: |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `33.33% <60.00%> (+33.33%)` | :arrow_up: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `31.32% <100.00%> (+0.83%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=footer). Last update [a6cab40...7a417c0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/446?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on January 08, 2021 at 09:49 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/446*
