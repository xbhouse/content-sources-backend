---
type: pull_request
number: 482
title: "Query cleanup"
state: merged
author: semtexzv
created: 2021-01-20T14:22:01Z
updated: 2021-03-16T09:25:21Z
closed: 2021-01-26T06:21:08Z
merged: 2021-01-26T06:21:08Z
base_branch: master
head_branch: query-cleanup
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/482
---

# Pull Request #482: Query cleanup

**Author**: @semtexzv
**Created**: January 20, 2021 at 02:22 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `query-cleanup`

## Description

Cleanup the SQL queries used across the manager.

---

## Discussion

### Comment by @codecov-io on January 25, 2021 at 03:48 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=h1) Report
> Merging [#482](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=desc) (b0cf1a9) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4eb8b558dee6fec3067e59d6bd06a4250b3575f4?el=desc) (4eb8b55) will **decrease** coverage by `0.59%`.
> The diff coverage is `69.56%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #482      +/-   ##
==========================================
- Coverage   62.63%   62.03%   -0.60%     
==========================================
  Files          61       61              
  Lines        2783     2771      -12     
==========================================
- Hits         1743     1719      -24     
- Misses        796      808      +12     
  Partials      244      244              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.75% <100.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.22% <100.00%> (-0.14%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <100.00%> (ø)` | |
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `82.97% <100.00%> (-2.21%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `74.28% <100.00%> (-3.22%)` | :arrow_down: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `89.65% <100.00%> (-1.53%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <100.00%> (ø)` | |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `88.00% <100.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `62.85% <100.00%> (-2.94%)` | :arrow_down: |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=footer). Last update [4eb8b55...b0cf1a9](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/482?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on January 25, 2021 at 09:29 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/482*
