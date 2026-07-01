---
type: pull_request
number: 443
title: "Cherry-picks without db updates"
state: merged
author: josef-hak
created: 2020-12-10T14:49:41Z
updated: 2020-12-10T16:07:12Z
closed: 2020-12-10T16:05:22Z
merged: 2020-12-10T16:05:22Z
base_branch: master
head_branch: no-db-updates
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/443
---

# Pull Request #443: Cherry-picks without db updates

**Author**: @josef-hak
**Created**: December 10, 2020 at 02:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `no-db-updates`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on December 10, 2020 at 03:58 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=h1) Report
> Merging [#443](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=desc) (d718067) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/6e5243d6d364e6d741f1812ed6fcf92068079fa9?el=desc) (6e5243d) will **increase** coverage by `0.95%`.
> The diff coverage is `92.17%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #443      +/-   ##
==========================================
+ Coverage   61.91%   62.87%   +0.95%     
==========================================
  Files          59       59              
  Lines        2626     2707      +81     
==========================================
+ Hits         1626     1702      +76     
- Misses        766      769       +3     
- Partials      234      236       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `72.22% <70.37%> (-2.78%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.94% <85.71%> (+0.45%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.40% <100.00%> (ø)` | |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `66.66% <100.00%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `84.61% <100.00%> (+4.61%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `78.12% <100.00%> (+0.58%)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.42% <100.00%> (+0.51%)` | :arrow_up: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `82.75% <100.00%> (+13.52%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `76.59% <100.00%> (+2.17%)` | :arrow_up: |
| ... and [9 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=footer). Last update [6e5243d...d718067](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/443?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/443*
