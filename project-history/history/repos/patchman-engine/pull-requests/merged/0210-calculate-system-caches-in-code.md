---
type: pull_request
number: 210
title: "Calculate system caches in code"
state: merged
author: semtexzv
created: 2020-04-03T09:19:21Z
updated: 2021-03-16T09:28:05Z
closed: 2020-04-08T09:30:59Z
merged: 2020-04-08T09:30:59Z
base_branch: master
head_branch: system-caches
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/210
---

# Pull Request #210: Calculate system caches in code

**Author**: @semtexzv
**Created**: April 03, 2020 at 09:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `system-caches`

## Description

- Does not use database function for calculating system caches
- Updates last_evaluation when updating caches, removing one write per evaluation

---

## Discussion

### Comment by @josef-hak on April 03, 2020 at 02:01 PM UTC

@semtexzv tests failed on wrong go format.

### Comment by @codecov-io on April 06, 2020 at 12:37 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=h1) Report
> Merging [#210](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fff25116fcc35fb5d573258efc2a757d29e66e0f&el=desc) will **increase** coverage by `0.00%`.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #210   +/-   ##
=======================================
  Coverage   62.99%   62.99%           
=======================================
  Files          46       46           
  Lines        1978     1997   +19     
=======================================
+ Hits         1246     1258   +12     
- Misses        588      591    +3     
- Partials      144      148    +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.42% <66.66%> (-0.70%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=footer). Last update [fff2511...7a28f0c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/210?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on April 06, 2020 at 09:24 AM UTC

### Review by @semtexzv - Commented on April 06, 2020 at 11:35 AM UTC

### Review by @semtexzv - Commented on April 07, 2020 at 09:12 AM UTC

### Review by @semtexzv - Commented on April 07, 2020 at 09:12 AM UTC

### Review by @josef-hak - Approved on April 08, 2020 at 09:30 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/210*
