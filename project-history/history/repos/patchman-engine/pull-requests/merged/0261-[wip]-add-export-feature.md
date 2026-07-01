---
type: pull_request
number: 261
title: "[WIP] Add export feature"
state: merged
author: semtexzv
created: 2020-05-20T16:10:29Z
updated: 2020-05-25T11:20:38Z
closed: 2020-05-25T11:20:38Z
merged: 2020-05-25T11:20:38Z
base_branch: master
head_branch: export
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/261
---

# Pull Request #261: [WIP] Add export feature

**Author**: @semtexzv
**Created**: May 20, 2020 at 04:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `export`

## Description

- Added endpoints for exporting data
- Endpoints do not support paging, sorting, or filters(TBD)
- Supporting 'application/json' & 'text/csv' content types
- Flat object structure for easy manipulation

---

## Discussion

### Comment by @codecov-commenter on May 20, 2020 at 04:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=h1) Report
> Merging [#261](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/398b6ee4e53f7984d20e6fab2e8e0d8b105fe7bf&el=desc) will **increase** coverage by `0.06%`.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #261      +/-   ##
==========================================
+ Coverage   62.59%   62.66%   +0.06%     
==========================================
  Files          47       47              
  Lines        2163     2215      +52     
==========================================
+ Hits         1354     1388      +34     
- Misses        652      665      +13     
- Partials      157      162       +5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `74.41% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.69% <44.44%> (-3.09%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `84.21% <66.66%> (-7.69%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `80.39% <74.07%> (-8.90%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=footer). Last update [398b6ee...53a5b21](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/261?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on May 21, 2020 at 12:38 PM UTC

It seems good to me. I've just added some questions and ideas.

### Review by @semtexzv - Commented on May 25, 2020 at 09:57 AM UTC

### Review by @semtexzv - Commented on May 25, 2020 at 10:00 AM UTC

### Review by @semtexzv - Commented on May 25, 2020 at 10:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/261*
