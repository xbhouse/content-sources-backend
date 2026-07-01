---
type: pull_request
number: 343
title: "Add support for nested query maps"
state: merged
author: semtexzv
created: 2020-09-23T08:52:30Z
updated: 2021-03-16T09:28:57Z
closed: 2020-09-29T12:20:26Z
merged: 2020-09-29T12:20:26Z
base_branch: master
head_branch: nested-query
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/343
---

# Pull Request #343: Add support for nested query maps

**Author**: @semtexzv
**Created**: September 23, 2020 at 08:52 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `nested-query`

## Description

Needed in order to support deeply nested filters syntax.

---

## Discussion

### Comment by @codecov-commenter on September 28, 2020 at 07:29 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=h1) Report
> Merging [#343](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1f5521fad47f9b35accfb571fa05c40d75a98544?el=desc) will **increase** coverage by `0.94%`.
> The diff coverage is `85.18%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #343      +/-   ##
==========================================
+ Coverage   61.38%   62.33%   +0.94%     
==========================================
  Files          54       56       +2     
  Lines        2380     2506     +126     
==========================================
+ Hits         1461     1562     +101     
- Misses        706      725      +19     
- Partials      213      219       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.77% <85.18%> (+4.44%)` | :arrow_up: |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28.73% <0.00%> (-0.68%)` | :arrow_down: |
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `85.18% <0.00%> (ø)` | |
| [vmaas\_sync/metrics\_db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2RiLmdv) | `60.71% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=footer). Last update [1f5521f...13e477d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/343?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/343*
