---
type: pull_request
number: 312
title: "[WIP] Basic dynamic implementation of packages endpoint"
state: merged
author: semtexzv
created: 2020-08-24T14:22:15Z
updated: 2021-03-16T09:26:14Z
closed: 2020-09-01T08:10:21Z
merged: 2020-09-01T08:10:21Z
base_branch: master
head_branch: acc-packages
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/312
---

# Pull Request #312: [WIP] Basic dynamic implementation of packages endpoint

**Author**: @semtexzv
**Created**: August 24, 2020 at 02:22 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `acc-packages`

## Description

First implementation of per-account package data, currently calculating everything dynamically. If performance will be a problem, we'll move to pre-calculated table.

---

## Discussion

### Comment by @codecov-commenter on August 26, 2020 at 03:07 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=h1) Report
> Merging [#312](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/06f1f4a0bf42675367147f40e97519c3bf8b3369?el=desc) will **increase** coverage by `0.11%`.
> The diff coverage is `79.31%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #312      +/-   ##
==========================================
+ Coverage   61.69%   61.81%   +0.11%     
==========================================
  Files          53       54       +1     
  Lines        2600     2629      +29     
==========================================
+ Hits         1604     1625      +21     
- Misses        792      797       +5     
- Partials      204      207       +3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `77.77% <77.77%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `62.50% <100.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `75.38% <0.00%> (-1.18%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=footer). Last update [06f1f4a...32cacdd](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/312?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on August 27, 2020 at 01:17 PM UTC

lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/312*
