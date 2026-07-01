---
type: pull_request
number: 299
title: "Package data perf improvements"
state: merged
author: semtexzv
created: 2020-07-30T08:27:44Z
updated: 2021-03-16T09:26:07Z
closed: 2020-07-30T09:29:36Z
merged: 2020-07-30T09:29:36Z
base_branch: master
head_branch: pkgdata-perf
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/299
---

# Pull Request #299: Package data perf improvements

**Author**: @semtexzv
**Created**: July 30, 2020 at 08:27 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkgdata-perf`

## Description

- Changed how we query for packages during evaluation
- Added prometheus metrics for additional evaluation steps

---

## Discussion

### Comment by @codecov-commenter on July 30, 2020 at 08:35 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=h1) Report
> Merging [#299](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/41240e54c711e6aa991da539be010e8851893a7a&el=desc) will **increase** coverage by `0.07%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #299      +/-   ##
==========================================
+ Coverage   61.76%   61.83%   +0.07%     
==========================================
  Files          52       52              
  Lines        2597     2602       +5     
==========================================
+ Hits         1604     1609       +5     
  Misses        787      787              
  Partials      206      206              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.75% <100.00%> (+0.46%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=footer). Last update [41240e5...72e5d4e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/299?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on July 30, 2020 at 08:36 AM UTC

### Review by @semtexzv - Commented on July 30, 2020 at 08:39 AM UTC

### Review by @josef-hak - Approved on July 30, 2020 at 09:29 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/299*
