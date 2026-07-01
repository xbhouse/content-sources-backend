---
type: pull_request
number: 385
title: "fix(InventoryEntitiesReducer): remove extra store utilization"
state: closed
author: mkholjuraev
created: 2021-01-11T13:57:54Z
updated: 2021-08-09T06:55:07Z
closed: 2021-01-13T14:02:59Z
base_branch: master
head_branch: remove-extra-getState
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/385
---

# Pull Request #385: fix(InventoryEntitiesReducer): remove extra store utilization

**Author**: @mkholjuraev
**Created**: January 11, 2021 at 01:57 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `remove-extra-getState`

## Description

fixes: 'You may not call store.getState() while the reducer is executing. The reducer has already received the state as an argument. Pass it down from the top reducer instead of reading it from the store.' issue related to fetching an already existing store inside InventoryEntitiesReducer.  The issue occurred after updating frontend-components-utilities to 2.2.9.

---

## Discussion

### Comment by @codecov-io on January 11, 2021 at 02:32 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=h1) Report
> Merging [#385](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=desc) (c513ab0) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/6312ae1a81bc9d988083de29e65d3d7edf802a99?el=desc) (6312ae1) will **decrease** coverage by `0.02%`.
> The diff coverage is `75.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #385      +/-   ##
==========================================
- Coverage   72.80%   72.77%   -0.03%     
==========================================
  Files          69       69              
  Lines        1239     1238       -1     
  Branches      161      160       -1     
==========================================
- Hits          902      901       -1     
  Misses        283      283              
  Partials       54       54              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `4.00% <0.00%> (ø)` | |
| [src/store/Reducers/InventoryEntitiesReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385/diff?src=pr&el=tree#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0ludmVudG9yeUVudGl0aWVzUmVkdWNlci5qcw==) | `59.09% <100.00%> (-1.78%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=footer). Last update [6312ae1...c513ab0](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/385?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on January 13, 2021 at 02:02 PM UTC

Closing as it was solved in https://github.com/RedHatInsights/patchman-ui/pull/389

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/385*
