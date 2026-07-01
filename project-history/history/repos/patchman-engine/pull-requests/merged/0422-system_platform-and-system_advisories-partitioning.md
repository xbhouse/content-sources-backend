---
type: pull_request
number: 422
title: "system_platform and system_advisories partitioning fix"
state: merged
author: MichaelMraka
created: 2020-11-13T19:17:57Z
updated: 2020-11-18T11:46:58Z
closed: 2020-11-18T11:46:58Z
merged: 2020-11-18T11:46:58Z
base_branch: master
head_branch: pr4
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/422
---

# Pull Request #422: system_platform and system_advisories partitioning fix

**Author**: @MichaelMraka
**Created**: November 13, 2020 at 07:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr4`

## Description

Before merging this request fix CI database 
`update schema_migrations set version = 42, dirty = false;`

---

## Discussion

### Comment by @codecov-io on November 18, 2020 at 10:58 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422?src=pr&el=h1) Report
> Merging [#422](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422?src=pr&el=desc) (3d6e5ba) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/c106d5b8431146206f040332c9b6049466ecb8b5?el=desc) (c106d5b) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #422   +/-   ##
=======================================
  Coverage   62.75%   62.75%           
=======================================
  Files          59       59           
  Lines        2698     2698           
=======================================
  Hits         1693     1693           
  Misses        769      769           
  Partials      236      236           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422?src=pr&el=footer). Last update [c106d5b...3d6e5ba](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/422?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on November 18, 2020 at 09:08 AM UTC

### Review by @MichaelMraka - Commented on November 18, 2020 at 10:09 AM UTC

### Review by @semtexzv - Commented on November 18, 2020 at 09:56 AM UTC

Please remove the down migrations in the commit for sequential running

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/422*
