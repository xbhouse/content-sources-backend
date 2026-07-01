---
type: pull_request
number: 364
title: "Better check to skip writing package updates, reduce loaded columns"
state: merged
author: semtexzv
created: 2020-10-06T14:31:16Z
updated: 2021-03-16T09:29:32Z
closed: 2020-10-08T08:58:08Z
merged: 2020-10-08T08:58:08Z
base_branch: master
head_branch: skip-check
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/364
---

# Pull Request #364: Better check to skip writing package updates, reduce loaded columns

**Author**: @semtexzv
**Created**: October 06, 2020 at 02:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `skip-check`

## Description

Should improve performance by loading less data and overwriting less data.

---

## Discussion

### Comment by @codecov-io on October 07, 2020 at 07:27 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@4c932b5`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `43.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master     #364   +/-   ##
=========================================
  Coverage          ?   61.66%           
=========================================
  Files             ?       57           
  Lines             ?     2585           
  Branches          ?        0           
=========================================
  Hits              ?     1594           
  Misses            ?      759           
  Partials          ?      232           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `45.45% <0.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.66% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.79% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `58.33% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `31.25% <0.00%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `51.61% <0.00%> (ø)` | |
| ... and [4 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364?src=pr&el=footer). Last update [4c932b5...5101917](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/364?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on October 07, 2020 at 12:02 PM UTC

### Review by @semtexzv - Commented on October 07, 2020 at 12:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/364*
