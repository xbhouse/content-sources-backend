---
type: pull_request
number: 73
title: "Fixed error (500) response on not allowed filter"
state: merged
author: josef-hak
created: 2020-01-23T16:44:16Z
updated: 2020-01-30T17:25:45Z
closed: 2020-01-28T08:36:22Z
merged: 2020-01-28T08:36:22Z
base_branch: master
head_branch: sort_bad_request
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/73
---

# Pull Request #73: Fixed error (500) response on not allowed filter

**Author**: @josef-hak
**Created**: January 23, 2020 at 04:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sort_bad_request`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 27, 2020 at 07:39 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=h1) Report
> Merging [#73](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a5eb8735cd8049ad393d80f15cccae2570318487?src=pr&el=desc) will **decrease** coverage by `0.01%`.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #73      +/-   ##
==========================================
- Coverage   57.57%   57.55%   -0.02%     
==========================================
  Files          29       32       +3     
  Lines        1221     1303      +82     
==========================================
+ Hits          703      750      +47     
- Misses        436      466      +30     
- Partials       82       87       +5
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.56% <100%> (+4.47%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.15% <100%> (+5.76%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.71% <100%> (+4.76%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `79.72% <100%> (+4.05%)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `44.44% <0%> (-8.89%)` | :arrow_down: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `0% <0%> (-3.39%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `69.31% <0%> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `44.44% <0%> (ø)` | :arrow_up: |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28% <0%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `66.66% <0%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=footer). Last update [a5eb873...65743f0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/73?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on January 27, 2020 at 10:00 AM UTC

### Review by @semtexzv - Commented on January 27, 2020 at 10:01 AM UTC

### Review by @josef-hak - Commented on January 27, 2020 at 10:04 AM UTC

### Review by @semtexzv - Commented on January 27, 2020 at 10:34 AM UTC

### Review by @semtexzv - Approved on January 27, 2020 at 10:37 AM UTC

Missclick. Will merge after test cases are added.

### Review by @josef-hak - Commented on January 27, 2020 at 12:40 PM UTC

### Review by @semtexzv - Approved on January 28, 2020 at 08:36 AM UTC

:+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/73*
