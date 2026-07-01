---
type: pull_request
number: 365
title: "Remove unused tags code"
state: merged
author: semtexzv
created: 2020-10-06T14:34:27Z
updated: 2021-03-16T09:29:31Z
closed: 2020-10-07T11:58:36Z
merged: 2020-10-07T11:58:36Z
base_branch: master
head_branch: unused-tags
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/365
---

# Pull Request #365: Remove unused tags code

**Author**: @semtexzv
**Created**: October 06, 2020 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `unused-tags`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on October 06, 2020 at 02:44 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=h1) Report
> Merging [#365](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fee97a7f77d067ac041bc68090e9d9db66b2f89d?el=desc) will **decrease** coverage by `0.54%`.
> The diff coverage is `40.67%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #365      +/-   ##
==========================================
- Coverage   62.27%   61.73%   -0.55%     
==========================================
  Files          56       57       +1     
  Lines        2534     2587      +53     
==========================================
+ Hits         1578     1597      +19     
- Misses        730      758      +28     
- Partials      226      232       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.66% <ø> (ø)` | |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <ø> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.79% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `58.33% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `31.25% <0.00%> (-1.22%)` | :arrow_down: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `51.61% <0.00%> (-7.01%)` | :arrow_down: |
| [vmaas\_sync/metrics\_cyndi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2N5bmRpLmdv) | `40.42% <40.42%> (ø)` | |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=footer). Last update [4c932b5...31e5d7a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/365?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/365*
