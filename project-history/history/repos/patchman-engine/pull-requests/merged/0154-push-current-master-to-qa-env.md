---
type: pull_request
number: 154
title: "Push current master to QA env"
state: merged
author: semtexzv
created: 2020-03-02T09:29:00Z
updated: 2020-03-16T10:23:18Z
closed: 2020-03-02T10:00:10Z
merged: 2020-03-02T10:00:10Z
base_branch: stable
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/154
---

# Pull Request #154: Push current master to QA env

**Author**: @semtexzv
**Created**: March 02, 2020 at 09:29 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `stable` ← **Head**: `master`

## Description

Changes:

- RBAC implementation, temporarily disabled
- Multiple evaluators
- Manifests
- Host-egress queue
- Updates to cache calculation

---

## Discussion

### Comment by @codecov-io on March 16, 2020 at 10:23 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=h1) Report
> Merging [#154](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=desc) into [stable](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/76df208e8abbe76bdd6b0f800b5d90ac54aec382?src=pr&el=desc) will **decrease** coverage by `0.54%`.
> The diff coverage is `74.02%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           stable     #154      +/-   ##
==========================================
- Coverage    64.2%   63.65%   -0.55%     
==========================================
  Files          41       43       +2     
  Lines        1595     1615      +20     
==========================================
+ Hits         1024     1028       +4     
- Misses        450      466      +16     
  Partials      121      121
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.5% <ø> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `30.43% <ø> (ø)` | :arrow_up: |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0% <0%> (ø)` | :arrow_up: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `77.5% <100%> (+2.5%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `91.89% <100%> (ø)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `81.25% <100%> (-7.22%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `78.57% <100%> (ø)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `45.45% <100%> (+45.45%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.17% <100%> (ø)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `76.31% <100%> (+0.64%)` | :arrow_up: |
| ... and [11 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=footer). Last update [76df208...e830f0e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/154?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on March 02, 2020 at 10:00 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/154*
