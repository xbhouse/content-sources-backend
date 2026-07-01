---
type: pull_request
number: 151
title: "Re-work the system_update trigger to fix skipped DELETEs"
state: merged
author: semtexzv
created: 2020-02-26T15:35:09Z
updated: 2021-03-16T09:27:33Z
closed: 2020-02-28T12:11:20Z
merged: 2020-02-28T12:11:20Z
base_branch: master
head_branch: fix-aad-counts
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/151
---

# Pull Request #151: Re-work the system_update trigger to fix skipped DELETEs

**Author**: @semtexzv
**Created**: February 26, 2020 at 03:35 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-aad-counts`

## Description

Attempts to fix issue when system_update trigger left rows with affected_systems of 0 in the table.

Most probable cause for this issue was the fact that queries within CTE are not evaluated sequentially, and subsequent queries do not see changes of previous queries on underlying tables.

This meant that changes performed by UPDATE (decrement) were not seen by DELETE, resulting in missed deletes.

New approach uses `to_update_advisories` temporary table with row locks as a source of truth, and performs updates, inserts, deletes based on information stored in it.

---

## Discussion

### Comment by @codecov-io on February 28, 2020 at 10:52 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=h1) Report
> Merging [#151](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f19666869158c5c7a9424e4ed283eaebb7aea7ac?src=pr&el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #151   +/-   ##
=======================================
  Coverage   64.94%   64.94%           
=======================================
  Files          42       42           
  Lines        1626     1626           
=======================================
  Hits         1056     1056           
  Misses        450      450           
  Partials      120      120
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `70.37% <0%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=footer). Last update [f196668...50f9de6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/151?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on February 27, 2020 at 09:07 AM UTC

Thanks, it's quite hard to express such a complex logic in pure SQL. One thing is whether it's working correctly, second and not less important one is keep the code understandable and clear 

### Review by @semtexzv - Commented on February 28, 2020 at 10:07 AM UTC

### Review by @semtexzv - Commented on February 28, 2020 at 10:19 AM UTC

### Review by @semtexzv - Commented on February 28, 2020 at 10:43 AM UTC

### Review by @MichaelMraka - Commented on February 28, 2020 at 11:14 AM UTC

### Review by @josef-hak - Approved on February 28, 2020 at 12:11 PM UTC

Thanks for updates. Ok, let's try it.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/151*
