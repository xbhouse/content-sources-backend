---
type: pull_request
number: 400
title: "Re-partition system_package"
state: merged
author: semtexzv
created: 2020-10-21T09:33:28Z
updated: 2021-03-16T09:29:51Z
closed: 2020-10-23T11:29:35Z
merged: 2020-10-23T11:29:35Z
base_branch: master
head_branch: syspkg-part
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/400
---

# Pull Request #400: Re-partition system_package

**Author**: @semtexzv
**Created**: October 21, 2020 at 09:33 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `syspkg-part`

## Description

Seems that 16 partitions was not enough. This PR creates new table with 80 partitions, copies old data and drops the old one.


---

## Discussion

### Comment by @codecov-io on October 21, 2020 at 11:23 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400?src=pr&el=h1) Report
> Merging [#400](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/da1a4e8c12a5dd3dbd2dc0ed4343510dd0aff404?el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #400   +/-   ##
=======================================
  Coverage   61.87%   61.87%           
=======================================
  Files          59       59           
  Lines        2607     2607           
=======================================
  Hits         1613     1613           
  Misses        761      761           
  Partials      233      233           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400?src=pr&el=footer). Last update [da1a4e8...77aff54](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/400?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Changes Requested on October 21, 2020 at 10:57 AM UTC

all migration files should start with 3-digit

### Review by @semtexzv - Commented on October 21, 2020 at 11:08 AM UTC

### Review by @semtexzv - Commented on October 21, 2020 at 11:09 AM UTC

### Review by @josef-hak - Commented on October 22, 2020 at 09:25 AM UTC

### Review by @semtexzv - Commented on October 23, 2020 at 09:27 AM UTC

### Review by @semtexzv - Commented on October 23, 2020 at 11:02 AM UTC

### Review by @josef-hak - Approved on October 23, 2020 at 11:27 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/400*
