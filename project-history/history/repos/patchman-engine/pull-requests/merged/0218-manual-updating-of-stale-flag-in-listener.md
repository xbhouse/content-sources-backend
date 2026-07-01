---
type: pull_request
number: 218
title: "Manual updating of stale flag in listener"
state: merged
author: semtexzv
created: 2020-04-08T16:10:04Z
updated: 2021-03-16T09:28:04Z
closed: 2020-04-14T07:45:56Z
merged: 2020-04-14T07:45:56Z
base_branch: master
head_branch: fix-stale-flags
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/218
---

# Pull Request #218: Manual updating of stale flag in listener

**Author**: @semtexzv
**Created**: April 08, 2020 at 04:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-stale-flags`

## Description

The `on_system_timestamp_update` did not call the trigger to update cached counts after updating the `stale` flag.
This changes so that listener updates this flag manually.

---

## Discussion

### Comment by @codecov-io on April 09, 2020 at 12:33 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=h1) Report
> Merging [#218](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1f0511854e00faec203e40a0069a6f9ccb544d21&el=desc) will **decrease** coverage by `0.11%`.
> The diff coverage is `75.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #218      +/-   ##
==========================================
- Coverage   63.02%   62.90%   -0.12%     
==========================================
  Files          46       46              
  Lines        1993     1995       +2     
==========================================
- Hits         1256     1255       -1     
- Misses        590      592       +2     
- Partials      147      148       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `74.73% <100.00%> (-0.83%)` | :arrow_down: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `61.76% <0.00%> (-1.10%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `73.58% <0.00%> (-0.49%)` | :arrow_down: |
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `83.78% <0.00%> (-0.43%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `65.62% <0.00%> (-0.36%)` | :arrow_down: |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `54.54% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=footer). Last update [1f05118...26109f0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/218?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on April 09, 2020 at 01:17 PM UTC

### Review by @semtexzv - Commented on April 14, 2020 at 07:26 AM UTC

### Review by @semtexzv - Commented on April 14, 2020 at 07:27 AM UTC

### Review by @josef-hak - Approved on April 14, 2020 at 07:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/218*
