---
type: pull_request
number: 463
title: "Bulk delete"
state: merged
author: semtexzv
created: 2021-01-12T10:37:30Z
updated: 2021-01-12T13:54:11Z
closed: 2021-01-12T13:54:10Z
merged: 2021-01-12T13:54:10Z
base_branch: master
head_branch: bulk-delete
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/463
---

# Pull Request #463: Bulk delete

**Author**: @semtexzv
**Created**: January 12, 2021 at 10:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bulk-delete`

## Description

- Add stored procedure for bulk delete
- Update delete_culled_systems to use this method
- Manually delete yupana systems in batches of 5000

---

## Discussion

### Comment by @codecov-io on January 12, 2021 at 10:52 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463?src=pr&el=h1) Report
> Merging [#463](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463?src=pr&el=desc) (77ad6a5) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e01f5f46dc811256c45e42e9a059763bf2d943b5?el=desc) (e01f5f4) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #463   +/-   ##
=======================================
  Coverage   62.74%   62.74%           
=======================================
  Files          61       61           
  Lines        2781     2781           
=======================================
  Hits         1745     1745           
  Misses        792      792           
  Partials      244      244           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463?src=pr&el=footer). Last update [e01f5f4...1e31e48](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/463?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on January 12, 2021 at 12:07 PM UTC

### Review by @josef-hak - Commented on January 12, 2021 at 12:08 PM UTC

@semtexzv don't you want to fix mark_systems_stale() method once we know that we have a problem with that?

### Review by @josef-hak - Commented on January 12, 2021 at 01:09 PM UTC

### Review by @semtexzv - Commented on January 12, 2021 at 01:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/463*
