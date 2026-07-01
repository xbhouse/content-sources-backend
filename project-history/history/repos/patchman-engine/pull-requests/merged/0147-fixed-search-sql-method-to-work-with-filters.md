---
type: pull_request
number: 147
title: "Fixed search sql method to work with filters"
state: merged
author: josef-hak
created: 2020-02-24T16:42:46Z
updated: 2020-02-27T08:35:37Z
closed: 2020-02-25T13:54:14Z
merged: 2020-02-25T13:54:14Z
base_branch: master
head_branch: filter_search
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/147
---

# Pull Request #147: Fixed search sql method to work with filters

**Author**: @josef-hak
**Created**: February 24, 2020 at 04:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `filter_search`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 24, 2020 at 04:52 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=h1) Report
> Merging [#147](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b699c9230a979daddff7fc42ffdeaf8eaf6a5bb7?src=pr&el=desc) will **decrease** coverage by `0.04%`.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #147      +/-   ##
==========================================
- Coverage   64.98%   64.94%   -0.05%     
==========================================
  Files          42       42              
  Lines        1628     1626       -2     
==========================================
- Hits         1058     1056       -2     
  Misses        450      450              
  Partials      120      120
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.17% <100%> (-0.49%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=footer). Last update [b699c92...2ec44fc](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/147?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on February 25, 2020 at 01:54 PM UTC

I think the new approach will cause more overhead, but we can change it to generate separate conditions for each column if that ever becomes a problem.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/147*
