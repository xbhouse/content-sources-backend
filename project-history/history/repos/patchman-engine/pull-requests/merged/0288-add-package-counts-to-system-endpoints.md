---
type: pull_request
number: 288
title: "Add package counts to system endpoints"
state: merged
author: semtexzv
created: 2020-07-16T09:53:21Z
updated: 2021-03-16T09:26:08Z
closed: 2020-08-12T08:08:44Z
merged: 2020-08-12T08:08:44Z
base_branch: master
head_branch: pkg-counts
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/288
---

# Pull Request #288: Add package counts to system endpoints

**Author**: @semtexzv
**Created**: July 16, 2020 at 09:53 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-counts`

## Description

Adds pre-calculated package counts to `system_platform`.

---

## Discussion

### Comment by @codecov-commenter on August 03, 2020 at 03:05 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=h1) Report
> Merging [#288](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/56f040a828c66977c25913e52837d26550e40c73&el=desc) will **increase** coverage by `0.14%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #288      +/-   ##
==========================================
+ Coverage   61.83%   61.98%   +0.14%     
==========================================
  Files          52       52              
  Lines        2602     2612      +10     
==========================================
+ Hits         1609     1619      +10     
  Misses        787      787              
  Partials      206      206              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <ø> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `67.65% <100.00%> (+0.89%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=footer). Last update [56f040a...a81a0c4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/288?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on July 21, 2020 at 09:48 AM UTC

### Review by @semtexzv - Commented on July 22, 2020 at 07:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/288*
