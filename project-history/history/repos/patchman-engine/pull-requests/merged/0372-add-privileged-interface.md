---
type: pull_request
number: 372
title: "Add privileged interface"
state: merged
author: semtexzv
created: 2020-10-09T14:24:33Z
updated: 2020-11-24T13:12:10Z
closed: 2020-11-24T13:12:10Z
merged: 2020-11-24T13:12:10Z
base_branch: master
head_branch: admin
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/372
---

# Pull Request #372: Add privileged interface

**Author**: @semtexzv
**Created**: October 09, 2020 at 02:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `admin`

## Description

Use platform-go-middleware for parsing identity, Add a new api group accepting only associate auth, which comes from turnpike.

---

## Discussion

### Comment by @codecov-io on November 19, 2020 at 09:13 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=h1) Report
> Merging [#372](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=desc) (7a2a8ba) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/520d79c816d335bbbeab6d6585b2609a78d6c148?el=desc) (520d79c) will **increase** coverage by `0.78%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #372      +/-   ##
==========================================
+ Coverage   61.97%   62.76%   +0.78%     
==========================================
  Files          59       59              
  Lines        2630     2691      +61     
==========================================
+ Hits         1630     1689      +59     
  Misses        766      766              
- Partials      234      236       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/identity.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9pZGVudGl0eS5nbw==) | `50.00% <100.00%> (+6.25%)` | :arrow_up: |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `72.22% <0.00%> (-2.78%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.89% <0.00%> (-0.60%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <0.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.36% <0.00%> (+0.44%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.23% <0.00%> (+0.69%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <0.00%> (+0.70%)` | :arrow_up: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `91.42% <0.00%> (+0.80%)` | :arrow_up: |
| [manager/controllers/systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2V4cG9ydC5nbw==) | `79.16% <0.00%> (+0.90%)` | :arrow_up: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `91.17% <0.00%> (+1.52%)` | :arrow_up: |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=footer). Last update [520d79c...7a2a8ba](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/372?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on November 19, 2020 at 02:39 PM UTC

looks good, I think I found one forgotten print, thanks

### Review by @semtexzv - Commented on November 23, 2020 at 12:09 PM UTC

### Review by @josef-hak - Approved on November 23, 2020 at 12:57 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/372*
