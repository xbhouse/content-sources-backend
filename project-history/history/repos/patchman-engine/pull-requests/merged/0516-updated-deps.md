---
type: pull_request
number: 516
title: "updated deps"
state: merged
author: josef-hak
created: 2021-02-02T13:27:26Z
updated: 2021-02-04T10:49:51Z
closed: 2021-02-03T11:13:14Z
merged: 2021-02-03T11:13:14Z
base_branch: master
head_branch: deps
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/516
---

# Pull Request #516: updated deps

**Author**: @josef-hak
**Created**: February 02, 2021 at 01:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `deps`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 02, 2021 at 02:44 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=h1) Report
> Merging [#516](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=desc) (ac867d0) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/40a3041d1bc52cc2f960e62a5067fcc7d6f70c9e?el=desc) (40a3041) will **decrease** coverage by `1.20%`.
> The diff coverage is `61.76%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #516      +/-   ##
==========================================
- Coverage   61.51%   60.31%   -1.21%     
==========================================
  Files          61       66       +5     
  Lines        2770     2870     +100     
==========================================
+ Hits         1704     1731      +27     
- Misses        815      891      +76     
+ Partials      251      248       -3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `96.07% <ø> (+3.62%)` | :arrow_up: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `71.05% <ø> (+3.55%)` | :arrow_up: |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `60.60% <ø> (+3.46%)` | :arrow_up: |
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `74.19% <ø> (+4.49%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.35% <ø> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `31.57% <50.00%> (ø)` | |
| [manager/middlewares/rbac.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9taWRkbGV3YXJlcy9yYmFjLmdv) | `59.52% <59.09%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.49% <100.00%> (+0.08%)` | :arrow_up: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <100.00%> (+4.15%)` | :arrow_up: |
| ... and [6 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=footer). Last update [929cd2c...ac867d0](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/516?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/516*
