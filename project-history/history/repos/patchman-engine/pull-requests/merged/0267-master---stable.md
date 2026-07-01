---
type: pull_request
number: 267
title: "master -> stable"
state: merged
author: josef-hak
created: 2020-06-09T09:28:12Z
updated: 2020-06-09T11:47:07Z
closed: 2020-06-09T11:47:06Z
merged: 2020-06-09T11:47:06Z
base_branch: stable
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/267
---

# Pull Request #267: master -> stable

**Author**: @josef-hak
**Created**: June 09, 2020 at 09:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `stable` ← **Head**: `master`

## Description

- we have some changes here we can ship to production

---

## Discussion

### Comment by @codecov-commenter on June 09, 2020 at 11:19 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=h1) Report
> Merging [#267](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=desc) into [stable](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/30ecdf9491e692b29f860888a83f2d2d06838195&el=desc) will **increase** coverage by `0.13%`.
> The diff coverage is `59.34%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           stable     #267      +/-   ##
==========================================
+ Coverage   62.19%   62.32%   +0.13%     
==========================================
  Files          47       47              
  Lines        2161     2243      +82     
==========================================
+ Hits         1344     1398      +54     
- Misses        660      679      +19     
- Partials      157      166       +9     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `14.28% <0.00%> (-2.08%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.08% <ø> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `74.41% <ø> (ø)` | |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `90.00% <ø> (-0.33%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `73.45% <42.10%> (-6.34%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `81.96% <64.00%> (-9.93%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `77.77% <70.00%> (-11.51%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `74.62% <100.00%> (+0.78%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.57% <100.00%> (+0.38%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.59% <100.00%> (+4.08%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=footer). Last update [30ecdf9...70511e2](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/267?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on June 09, 2020 at 11:39 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/267*
