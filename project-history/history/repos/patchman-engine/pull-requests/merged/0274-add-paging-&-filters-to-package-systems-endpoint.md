---
type: pull_request
number: 274
title: "Add paging & filters to package systems endpoint"
state: merged
author: semtexzv
created: 2020-06-17T07:02:30Z
updated: 2021-03-16T09:28:30Z
closed: 2020-06-22T10:41:36Z
merged: 2020-06-22T10:41:36Z
base_branch: master
head_branch: pkgdata-fixes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/274
---

# Pull Request #274: Add paging & filters to package systems endpoint

**Author**: @semtexzv
**Created**: June 17, 2020 at 07:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkgdata-fixes`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on June 19, 2020 at 10:13 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=h1) Report
> Merging [#274](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ae465739da754dcb1e40adf5ccb69e40d330da2b&el=desc) will **decrease** coverage by `0.11%`.
> The diff coverage is `86.95%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #274      +/-   ##
==========================================
- Coverage   62.15%   62.03%   -0.12%     
==========================================
  Files          49       51       +2     
  Lines        2328     2376      +48     
==========================================
+ Hits         1447     1474      +27     
- Misses        704      719      +15     
- Partials      177      183       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `67.85% <84.21%> (+1.19%)` | :arrow_up: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `64.00% <100.00%> (+6.85%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `52.63% <0.00%> (-9.87%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.11% <0.00%> (-0.85%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `64.54% <0.00%> (-0.55%)` | :arrow_down: |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `90.32% <0.00%> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `62.50% <0.00%> (ø)` | |
| [base/mqueue/message.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.10% <0.00%> (+2.65%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=footer). Last update [ae46573...22513bf](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/274?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on June 18, 2020 at 07:55 AM UTC

Lint fail.

### Review by @josef-hak - Changes Requested on June 22, 2020 at 08:59 AM UTC

Good, just one note.

### Review by @semtexzv - Commented on June 22, 2020 at 10:24 AM UTC

### Review by @josef-hak - Approved on June 22, 2020 at 10:41 AM UTC

Great, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/274*
