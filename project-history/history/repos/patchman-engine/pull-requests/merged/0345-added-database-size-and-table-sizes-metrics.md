---
type: pull_request
number: 345
title: "Added database size and table sizes metrics"
state: merged
author: josef-hak
created: 2020-09-23T14:28:27Z
updated: 2020-09-30T06:56:10Z
closed: 2020-09-25T11:25:06Z
merged: 2020-09-25T11:25:06Z
base_branch: master
head_branch: db-metrics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/345
---

# Pull Request #345: Added database size and table sizes metrics

**Author**: @josef-hak
**Created**: September 23, 2020 at 02:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `db-metrics`

## Description

also added database processes metrics got using `pg_stat_activity`

---

## Discussion

### Comment by @codecov-commenter on September 23, 2020 at 07:17 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=h1) Report
> Merging [#345](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1f5521fad47f9b35accfb571fa05c40d75a98544?el=desc) will **decrease** coverage by `0.05%`.
> The diff coverage is `54.83%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #345      +/-   ##
==========================================
- Coverage   61.38%   61.32%   -0.06%     
==========================================
  Files          54       55       +1     
  Lines        2380     2410      +30     
==========================================
+ Hits         1461     1478      +17     
- Misses        706      716      +10     
- Partials      213      216       +3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.73% <0.00%> (-0.68%)` | :arrow_down: |
| [vmaas\_sync/metrics\_db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2RiLmdv) | `60.71% <60.71%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=footer). Last update [1f5521f...aaae1a3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/345?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on September 25, 2020 at 11:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/345*
