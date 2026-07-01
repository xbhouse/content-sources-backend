---
type: pull_request
number: 108
title: "Fix refresh_system_caches modifying all system_platforms"
state: merged
author: semtexzv
created: 2020-02-07T14:53:15Z
updated: 2021-03-16T09:27:20Z
closed: 2020-02-07T15:35:50Z
merged: 2020-02-07T15:35:50Z
base_branch: master
head_branch: fix-system-caches
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/108
---

# Pull Request #108: Fix refresh_system_caches modifying all system_platforms

**Author**: @semtexzv
**Created**: February 07, 2020 at 02:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-system-caches`

## Description

Fix `refresh_system_caches` always modyfying all `system_platform` rows.

---

## Discussion

### Comment by @codecov-io on February 07, 2020 at 03:35 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108?src=pr&el=h1) Report
> Merging [#108](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ddb79d7ea12ba01f6080f84f3b10e94d1779d6b0?src=pr&el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #108   +/-   ##
=======================================
  Coverage   54.15%   54.15%           
=======================================
  Files          38       38           
  Lines        1479     1479           
=======================================
  Hits          801      801           
  Misses        581      581           
  Partials       97       97
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108?src=pr&el=footer). Last update [ddb79d7...964422b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/108?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on February 07, 2020 at 03:35 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/108*
