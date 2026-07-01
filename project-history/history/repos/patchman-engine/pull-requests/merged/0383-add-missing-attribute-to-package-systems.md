---
type: pull_request
number: 383
title: "Add missing attribute to package systems"
state: merged
author: semtexzv
created: 2020-10-13T12:43:13Z
updated: 2021-03-16T09:29:42Z
closed: 2020-10-14T16:26:13Z
merged: 2020-10-14T16:26:13Z
base_branch: master
head_branch: fix-package-systems
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/383
---

# Pull Request #383: Add missing attribute to package systems

**Author**: @semtexzv
**Created**: October 13, 2020 at 12:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-package-systems`

## Description

We had missing attribute on this endpoint. Needed for package systems table.

---

## Discussion

### Comment by @codecov-io on October 14, 2020 at 04:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=h1) Report
> Merging [#383](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ad05dcbf18bfb96a9a04f7b263fe537a15d5ad99?el=desc) will **increase** coverage by `0.02%`.
> The diff coverage is `83.09%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #383      +/-   ##
==========================================
+ Coverage   61.60%   61.63%   +0.02%     
==========================================
  Files          57       58       +1     
  Lines        2589     2601      +12     
==========================================
+ Hits         1595     1603       +8     
- Misses        761      763       +2     
- Partials      233      235       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.84% <ø> (ø)` | |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `75.00% <20.00%> (-4.69%)` | :arrow_down: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `69.23% <69.23%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `95.91% <100.00%> (+9.25%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <100.00%> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <100.00%> (ø)` | |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `79.31% <100.00%> (+2.38%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <100.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.65% <100.00%> (+0.64%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=footer). Last update [93a5e8b...2a00679](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/383?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on October 14, 2020 at 02:21 PM UTC

can be merged after passed tests

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/383*
