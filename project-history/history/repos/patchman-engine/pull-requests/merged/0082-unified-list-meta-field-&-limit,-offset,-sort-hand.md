---
type: pull_request
number: 82
title: "Unified list meta field & limit, offset, sort handling"
state: merged
author: semtexzv
created: 2020-01-29T15:59:45Z
updated: 2021-03-16T09:27:11Z
closed: 2020-02-03T08:52:04Z
merged: 2020-02-03T08:52:04Z
base_branch: master
head_branch: list-meta
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/82
---

# Pull Request #82: Unified list meta field & limit, offset, sort handling

**Author**: @semtexzv
**Created**: January 29, 2020 at 03:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `list-meta`

## Description

- Unifies different `meta` fields.
- Handles limit, offset, sort in common function

---

## Discussion

### Comment by @codecov-io on January 30, 2020 at 07:23 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=h1) Report
> Merging [#82](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1485e60e47c32fe088c10ecf36c9f3504051763c?src=pr&el=desc) will **decrease** coverage by `0.79%`.
> The diff coverage is `86.53%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##           master      #82     +/-   ##
=========================================
- Coverage   57.23%   56.44%   -0.8%     
=========================================
  Files          34       34             
  Lines        1410     1343     -67     
=========================================
- Hits          807      758     -49     
+ Misses        506      494     -12     
+ Partials       97       91      -6
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `81.63% <100%> (+1.11%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `92.68% <100%> (+5.91%)` | :arrow_up: |
| [manager/controllers/test\_utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100% <100%> (ø)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `91.42% <100%> (+5.71%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `78.57% <100%> (+4.88%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `83.58% <81.08%> (-2.53%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=footer). Last update [1485e60...73c7856](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/82?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on February 03, 2020 at 08:49 AM UTC

looks good to me, thanks for generalization

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/82*
