---
type: pull_request
number: 391
title: "fix(database): Fix ERROR message"
state: merged
author: AlesKas
created: 2020-10-15T08:47:39Z
updated: 2020-11-16T11:35:45Z
closed: 2020-10-29T10:00:22Z
merged: 2020-10-29T10:00:22Z
base_branch: master
head_branch: SPM-500
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/391
---

# Pull Request #391: fix(database): Fix ERROR message

**Author**: @AlesKas
**Created**: October 15, 2020 at 08:47 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `SPM-500`

## Description

log_min_duration_statement can only be set by superuser, we can either grant db_admin superuser privileges, or set this in DB creation, when user is postgres

---

## Discussion

### Comment by @codecov-io on October 19, 2020 at 06:21 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391?src=pr&el=h1) Report
> Merging [#391](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/db87b9094b00a4376eaa4dba0f21a9b7b707a82f?el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #391   +/-   ##
=======================================
  Coverage   61.62%   61.62%           
=======================================
  Files          58       58           
  Lines        2603     2603           
=======================================
  Hits         1604     1604           
  Misses        764      764           
  Partials      235      235           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391?src=pr&el=footer). Last update [db87b90...aa62c1d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/391?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on October 19, 2020 at 10:31 AM UTC

### Review by @AlesKas - Commented on October 19, 2020 at 11:09 AM UTC

### Review by @semtexzv - Commented on October 19, 2020 at 07:04 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/391*
