---
type: pull_request
number: 352
title: "Use account ID instead of name in most queries"
state: merged
author: semtexzv
created: 2020-09-29T11:35:25Z
updated: 2021-03-16T09:29:00Z
closed: 2020-09-30T13:04:59Z
merged: 2020-09-30T13:04:59Z
base_branch: master
head_branch: use-acc-ids
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/352
---

# Pull Request #352: Use account ID instead of name in most queries

**Author**: @semtexzv
**Created**: September 29, 2020 at 11:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `use-acc-ids`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on September 30, 2020 at 12:47 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=h1) Report
> Merging [#352](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/e62f86cddf9f2f33249a59bb89fdf1f00cdb67ab?el=desc) will **decrease** coverage by `0.08%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #352      +/-   ##
==========================================
- Coverage   62.64%   62.55%   -0.09%     
==========================================
  Files          56       56              
  Lines        2500     2494       -6     
==========================================
- Hits         1566     1560       -6     
  Misses        715      715              
  Partials      219      219              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/core/gintesting.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2dpbnRlc3RpbmcuZ28=) | `87.50% <100.00%> (ø)` | |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `86.66% <100.00%> (-0.35%)` | :arrow_down: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `74.41% <100.00%> (-0.59%)` | :arrow_down: |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `68.96% <100.00%> (-1.04%)` | :arrow_down: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `76.00% <100.00%> (ø)` | |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `72.50% <100.00%> (ø)` | |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `53.12% <100.00%> (ø)` | |
| [manager/controllers/system\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGV0YWlsLmdv) | `89.65% <100.00%> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `58.33% <100.00%> (-1.13%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `82.97% <100.00%> (-0.36%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=footer). Last update [e62f86c...2822e1f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/352?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on September 29, 2020 at 11:41 AM UTC

### Review by @semtexzv - Commented on September 30, 2020 at 10:22 AM UTC

### Review by @semtexzv - Commented on September 30, 2020 at 10:27 AM UTC

### Review by @josef-hak - Approved on September 30, 2020 at 12:49 PM UTC

lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/352*
