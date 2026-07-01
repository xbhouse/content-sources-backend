---
type: pull_request
number: 366
title: "Update delete_system to take partitioning into account"
state: merged
author: semtexzv
created: 2020-10-07T09:50:49Z
updated: 2021-03-16T09:29:32Z
closed: 2020-10-07T12:59:58Z
merged: 2020-10-07T12:59:57Z
base_branch: master
head_branch: delete-part
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/366
---

# Pull Request #366: Update delete_system to take partitioning into account

**Author**: @semtexzv
**Created**: October 07, 2020 at 09:50 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `delete-part`

## Description

Probable cause of recent database deadlock

---

## Discussion

### Comment by @codecov-io on October 07, 2020 at 12:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@4c932b5`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `40.67%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master     #366   +/-   ##
=========================================
  Coverage          ?   61.73%           
=========================================
  Files             ?       57           
  Lines             ?     2587           
  Branches          ?        0           
=========================================
  Hits              ?     1597           
  Misses            ?      758           
  Partials          ?      232           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.66% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.79% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `58.33% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `31.25% <0.00%> (ø)` | |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `51.61% <0.00%> (ø)` | |
| [vmaas\_sync/metrics\_cyndi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2N5bmRpLmdv) | `40.42% <40.42%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366?src=pr&el=footer). Last update [4c932b5...92fe870](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/366?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on October 07, 2020 at 12:05 PM UTC

Tests passed, looks good. But I have not tested it more. @MichaelMraka could you check it too?

### Review by @MichaelMraka - Approved on October 07, 2020 at 12:20 PM UTC

### Review by @MichaelMraka - Approved on October 07, 2020 at 12:27 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/366*
