---
type: pull_request
number: 286
title: "Support granular package data"
state: merged
author: semtexzv
created: 2020-07-13T08:18:12Z
updated: 2021-01-22T12:32:02Z
closed: 2020-07-28T12:38:35Z
merged: 2020-07-28T12:38:35Z
base_branch: master
head_branch: pkgdata
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/286
---

# Pull Request #286: Support granular package data

**Author**: @semtexzv
**Created**: July 13, 2020 at 08:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkgdata`

## Description

This combines several PRs.

Should result in functional package data API, which provides per-package details along with fast `/packages/:id/systems` and `/systems/:id/packages` implementations.

Contains:
https://github.com/RedHatInsights/patchman-engine/pull/282
https://github.com/RedHatInsights/patchman-engine/pull/283
https://github.com/RedHatInsights/patchman-engine/pull/284
https://github.com/RedHatInsights/patchman-engine/pull/285

---

## Discussion

### Comment by @codecov-commenter on July 20, 2020 at 12:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=h1) Report
> Merging [#286](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8d41b285cb11dd06e1c8140f387046326abc3925&el=desc) will **increase** coverage by `0.40%`.
> The diff coverage is `70.61%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #286      +/-   ##
==========================================
+ Coverage   62.28%   62.69%   +0.40%     
==========================================
  Files          51       52       +1     
  Lines        2437     2571     +134     
==========================================
+ Hits         1518     1612      +94     
- Misses        729      757      +28     
- Partials      190      202      +12     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `50.00% <0.00%> (-8.34%)` | :arrow_down: |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `31.37% <0.00%> (ø)` | |
| [base/database/batch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9iYXRjaC5nbw==) | `65.28% <33.33%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `64.10% <54.54%> (-0.45%)` | :arrow_down: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `62.50% <62.50%> (ø)` | |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `80.30% <80.30%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.33% <80.70%> (+1.21%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `70.96% <100.00%> (+3.11%)` | :arrow_up: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `31.16% <100.00%> (ø)` | |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=footer). Last update [8d41b28...d92f219](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/286?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on July 21, 2020 at 09:09 AM UTC

@semtexzv, why did you create this feature branch in origin repo?

### Comment by @semtexzv on July 21, 2020 at 09:16 AM UTC

The feature was too big and individual PRs left master in inconsistent state.

---

## Reviews

### Review by @josef-hak - Changes Requested on July 20, 2020 at 02:01 PM UTC

conflicts

### Review by @josef-hak - Changes Requested on July 21, 2020 at 08:57 AM UTC

It's huge PR, many features added at once (schama upgrade, syncing, data exposing). I would expect more tests added for that. But anyway in-repo testing dataset is too small to show eventual issues. In case of such big update I would do perf-testing for that branch before merging it to master.

### Review by @semtexzv - Commented on July 22, 2020 at 08:22 AM UTC

### Review by @semtexzv - Commented on July 22, 2020 at 08:22 AM UTC

### Review by @semtexzv - Commented on July 27, 2020 at 09:25 AM UTC

### Review by @josef-hak - Commented on July 27, 2020 at 03:53 PM UTC

### Review by @semtexzv - Commented on July 28, 2020 at 08:27 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/286*
