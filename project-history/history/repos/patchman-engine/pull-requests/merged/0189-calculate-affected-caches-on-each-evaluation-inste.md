---
type: pull_request
number: 189
title: "Calculate affected caches on each evaluation instead of modifying values"
state: merged
author: semtexzv
created: 2020-03-25T20:02:10Z
updated: 2021-03-16T09:27:57Z
closed: 2020-03-30T10:05:59Z
merged: 2020-03-30T10:05:59Z
base_branch: master
head_branch: absolute-caches
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/189
---

# Pull Request #189: Calculate affected caches on each evaluation instead of modifying values

**Author**: @semtexzv
**Created**: March 25, 2020 at 08:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `absolute-caches`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on March 30, 2020 at 08:21 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=h1) Report
> Merging [#189](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b027a777932c1a9f031ad6af7d24434eaad4e61a&el=desc) will **decrease** coverage by `0.37%`.
> The diff coverage is `58.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #189      +/-   ##
==========================================
- Coverage   63.13%   62.75%   -0.38%     
==========================================
  Files          46       46              
  Lines        1923     1917       -6     
==========================================
- Hits         1214     1203      -11     
- Misses        567      571       +4     
- Partials      142      143       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.55% <70.00%> (-1.90%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=footer). Last update [b027a77...b714bcb](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/189?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on March 27, 2020 at 12:07 PM UTC

Great, it could simplify our code quite a lot and make evaluation more robust. I hope it will be fast enough too. I found some things to improve it.

### Review by @semtexzv - Commented on March 27, 2020 at 05:00 PM UTC

### Review by @semtexzv - Commented on March 27, 2020 at 05:01 PM UTC

### Review by @semtexzv - Commented on March 27, 2020 at 05:03 PM UTC

### Review by @semtexzv - Commented on March 27, 2020 at 09:12 PM UTC

### Review by @semtexzv - Commented on March 27, 2020 at 09:28 PM UTC

### Review by @josef-hak - Changes Requested on March 30, 2020 at 08:00 AM UTC

Just one typo.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/189*
