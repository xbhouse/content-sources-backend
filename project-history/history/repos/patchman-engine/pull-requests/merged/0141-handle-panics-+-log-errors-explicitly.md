---
type: pull_request
number: 141
title: "Handle panics + log errors explicitly"
state: merged
author: semtexzv
created: 2020-02-20T07:51:09Z
updated: 2021-03-16T09:28:11Z
closed: 2020-04-15T12:57:12Z
merged: 2020-04-15T12:57:12Z
base_branch: master
head_branch: panics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/141
---

# Pull Request #141: Handle panics + log errors explicitly

**Author**: @semtexzv
**Created**: February 20, 2020 at 07:51 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `panics`

## Description

This should allow us to see fatal errors in external logging such as cloudwatch.

This is only a temporary solution, since panics happening in other goroutines will still just tear down the whole process.


---

## Discussion

### Comment by @codecov-io on February 24, 2020 at 09:05 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=h1) Report
> Merging [#141](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/05dbc18547888c25def99b0fcd9a6ba4e2f79a73&el=desc) will **decrease** coverage by `0.08%`.
> The diff coverage is `50.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #141      +/-   ##
==========================================
- Coverage   62.90%   62.82%   -0.09%     
==========================================
  Files          46       46              
  Lines        1995     2012      +17     
==========================================
+ Hits         1255     1264       +9     
- Misses        592      599       +7     
- Partials      148      149       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `23.68% <0.00%> (-0.32%)` | :arrow_down: |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/utils/awscloudwatch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9hd3NjbG91ZHdhdGNoLmdv) | `19.35% <25.00%> (+1.49%)` | :arrow_up: |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `17.24% <63.63%> (+10.85%)` | :arrow_up: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `74.07% <100.00%> (+0.48%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=footer). Last update [05dbc18...eee32f1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/141?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on February 25, 2020 at 08:53 AM UTC

### Review by @josef-hak - Approved on April 15, 2020 at 12:57 PM UTC

ok, thanks also for log library unification

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/141*
