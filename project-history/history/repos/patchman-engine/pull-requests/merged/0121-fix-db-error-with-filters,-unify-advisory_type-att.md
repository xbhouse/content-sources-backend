---
type: pull_request
number: 121
title: "Fix db error with  filters, unify advisory_type attr & filter"
state: merged
author: semtexzv
created: 2020-02-11T12:57:05Z
updated: 2020-02-11T13:07:26Z
closed: 2020-02-11T13:07:26Z
merged: 2020-02-11T13:07:26Z
base_branch: master
head_branch: api-filters
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/121
---

# Pull Request #121: Fix db error with  filters, unify advisory_type attr & filter

**Author**: @semtexzv
**Created**: February 11, 2020 at 12:57 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `api-filters`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 11, 2020 at 01:07 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=h1) Report
> Merging [#121](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/3d4ad65d1b16b03e16f2753800320d46ff6f1910?src=pr&el=desc) will **decrease** coverage by `0.09%`.
> The diff coverage is `44.44%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##           master     #121     +/-   ##
=========================================
- Coverage   55.31%   55.21%   -0.1%     
=========================================
  Files          40       40             
  Lines        1580     1581      +1     
=========================================
- Hits          874      873      -1     
- Misses        603      604      +1     
- Partials      103      104      +1
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `92.68% <100%> (ø)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.35% <33.33%> (-1.86%)` | :arrow_down: |
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `77.77% <50%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=footer). Last update [3d4ad65...ab9598e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/121?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/121*
