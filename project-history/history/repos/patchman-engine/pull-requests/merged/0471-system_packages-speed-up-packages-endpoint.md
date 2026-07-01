---
type: pull_request
number: 471
title: "System_packages: speed up packages endpoint"
state: merged
author: semtexzv
created: 2021-01-15T10:55:59Z
updated: 2021-03-16T09:25:16Z
closed: 2021-01-29T15:12:25Z
merged: 2021-01-29T15:12:25Z
base_branch: master
head_branch: packages-opt
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/471
---

# Pull Request #471: System_packages: speed up packages endpoint

**Author**: @semtexzv
**Created**: January 15, 2021 at 10:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `packages-opt`

## Description

The indexes in packages table forced us to do 2 joins before aggregation. this adds name_id into the table, and allows us perform aggregation before joins, which should greatly speed up the packages endpoint.


---

## Discussion

### Comment by @codecov-io on January 27, 2021 at 02:41 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=h1) Report
> Merging [#471](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=desc) (b34d1d3) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/edd262979dd72e6b94a49e0084a244c11fa3345e?el=desc) (edd2629) will **increase** coverage by `0.22%`.
> The diff coverage is `62.50%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #471      +/-   ##
==========================================
+ Coverage   61.51%   61.74%   +0.22%     
==========================================
  Files          61       61              
  Lines        2770     2763       -7     
==========================================
+ Hits         1704     1706       +2     
+ Misses        815      812       -3     
+ Partials      251      245       -6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.49% <100.00%> (+0.08%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <100.00%> (+4.15%)` | :arrow_up: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `89.28% <100.00%> (+6.52%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.35% <0.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `60.60% <0.00%> (+3.46%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.05% <0.00%> (+3.55%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.07% <0.00%> (+3.62%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=footer). Last update [edd2629...b34d1d3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/471?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on January 15, 2021 at 11:04 AM UTC

Good, I hope it will help us. But now there are some conflicts to resolve.

### Review by @semtexzv - Commented on January 21, 2021 at 02:04 PM UTC

### Review by @semtexzv - Commented on January 21, 2021 at 02:06 PM UTC

### Review by @josef-hak - Changes Requested on January 28, 2021 at 03:17 PM UTC

### Review by @josef-hak - Approved on January 28, 2021 at 04:19 PM UTC

Great, but we should merge and deploy "culling fixes PR" before this one.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/471*
