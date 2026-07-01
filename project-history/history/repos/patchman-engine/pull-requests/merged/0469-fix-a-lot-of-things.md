---
type: pull_request
number: 469
title: "Fix a lot of things"
state: merged
author: semtexzv
created: 2021-01-13T10:07:30Z
updated: 2021-03-16T09:25:17Z
closed: 2021-01-29T12:53:16Z
merged: 2021-01-29T12:53:16Z
base_branch: master
head_branch: bulk-delete
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/469
---

# Pull Request #469: Fix a lot of things

**Author**: @semtexzv
**Created**: January 13, 2021 at 10:07 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bulk-delete`

## Description

- Was missing order by clause in locking select
- Created index for inventory_id in system_platform
- added max_systems parameter to mark_system_stale

---

## Discussion

### Comment by @codecov-io on January 13, 2021 at 11:41 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469?src=pr&el=h1) Report
> Merging [#469](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469?src=pr&el=desc) (1e980c3) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b8e7872205fcf02b4dd8d6e46bc18526bdf96057?el=desc) (b8e7872) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #469   +/-   ##
=======================================
  Coverage   62.72%   62.72%           
=======================================
  Files          61       61           
  Lines        2785     2785           
=======================================
  Hits         1747     1747           
  Misses        794      794           
  Partials      244      244           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469?src=pr&el=footer). Last update [b8e7872...1e980c3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/469?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on January 18, 2021 at 03:26 PM UTC

@semtexzv is this still actual?

### Comment by @josef-hak on January 28, 2021 at 02:35 PM UTC

@semtexzv `TestEvaluate` failed.

---

## Reviews

### Review by @josef-hak - Changes Requested on January 22, 2021 at 10:15 AM UTC

### Review by @josef-hak - Changes Requested on January 28, 2021 at 01:56 PM UTC

### Review by @josef-hak - Changes Requested on January 28, 2021 at 03:19 PM UTC

### Review by @semtexzv - Commented on January 28, 2021 at 03:52 PM UTC

### Review by @josef-hak - Approved on January 28, 2021 at 04:24 PM UTC

Just check the rest of my comments. Otherwise it looks good.

### Review by @semtexzv - Commented on January 29, 2021 at 10:43 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/469*
