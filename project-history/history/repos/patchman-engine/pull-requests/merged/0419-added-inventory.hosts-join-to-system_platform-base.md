---
type: pull_request
number: 419
title: "added inventory.hosts join to system_platform-based queries in endpoints"
state: merged
author: josef-hak
created: 2020-11-09T16:29:02Z
updated: 2020-11-27T15:22:23Z
closed: 2020-11-10T15:06:43Z
merged: 2020-11-10T15:06:43Z
base_branch: master
head_branch: inventory-hosts-join
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/419
---

# Pull Request #419: added inventory.hosts join to system_platform-based queries in endpoints

**Author**: @josef-hak
**Created**: November 09, 2020 at 04:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `inventory-hosts-join`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on November 09, 2020 at 05:38 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=h1) Report
> Merging [#419](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=desc) (b5006ed) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/eff67fc1016cb664c5199e346b74e97891c012fe?el=desc) (eff67fc) will **increase** coverage by `0.28%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #419      +/-   ##
==========================================
+ Coverage   62.46%   62.74%   +0.28%     
==========================================
  Files          59       59              
  Lines        2656     2676      +20     
==========================================
+ Hits         1659     1679      +20     
  Misses        765      765              
  Partials      232      232              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.94% <ø> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.42% <100.00%> (+0.13%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.59% <100.00%> (+1.04%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `78.04% <100.00%> (+1.73%)` | :arrow_up: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `91.42% <100.00%> (+0.51%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `73.17% <100.00%> (+1.37%)` | :arrow_up: |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `91.42% <100.00%> (+0.80%)` | :arrow_up: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `66.66% <100.00%> (+2.77%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `90.32% <100.00%> (+1.03%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=footer). Last update [eff67fc...b5006ed](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/419?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on November 10, 2020 at 12:31 PM UTC

### Review by @josef-hak - Commented on November 10, 2020 at 01:04 PM UTC

### Review by @semtexzv - Approved on November 10, 2020 at 01:34 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/419*
