---
type: pull_request
number: 370
title: "Release 1.3.7"
state: merged
author: josef-hak
created: 2020-10-08T15:01:35Z
updated: 2020-10-09T13:22:12Z
closed: 2020-10-09T13:22:12Z
merged: 2020-10-09T13:22:12Z
base_branch: stable
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/370
---

# Pull Request #370: Release 1.3.7

**Author**: @josef-hak
**Created**: October 08, 2020 at 03:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `stable` ← **Head**: `master`

## Description

Fixed partitioning deletes
Better logging for vmaas-sync
...

---

## Discussion

### Comment by @codecov-io on October 08, 2020 at 04:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=h1) Report
> Merging [#370](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=desc) into [stable](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0910451d7a75d38a8e81fdbb18ca8dd5ff8562f4?el=desc) will **decrease** coverage by `0.13%`.
> The diff coverage is `73.52%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           stable     #370      +/-   ##
==========================================
- Coverage   61.73%   61.59%   -0.14%     
==========================================
  Files          57       57              
  Lines        2587     2588       +1     
==========================================
- Hits         1597     1594       -3     
- Misses        758      761       +3     
- Partials      232      233       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <ø> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `45.45% <0.00%> (-1.43%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.66% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.79% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `58.33% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <ø> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `61.66% <50.00%> (+0.32%)` | :arrow_up: |
| ... and [5 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=footer). Last update [0910451...41dd6d1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/370?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on October 09, 2020 at 09:01 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/370*
