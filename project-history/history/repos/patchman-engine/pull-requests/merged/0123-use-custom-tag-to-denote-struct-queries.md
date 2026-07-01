---
type: pull_request
number: 123
title: "Use custom tag to denote struct queries"
state: merged
author: semtexzv
created: 2020-02-11T15:17:58Z
updated: 2021-03-16T09:27:24Z
closed: 2020-02-12T16:22:23Z
merged: 2020-02-12T16:22:23Z
base_branch: master
head_branch: api-filters
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/123
---

# Pull Request #123: Use custom tag to denote struct queries

**Author**: @semtexzv
**Created**: February 11, 2020 at 03:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `api-filters`

## Description

Reworks the filter & sort implementation, so that SQL column definitions are stored along with structs which will be retrieved from database

Also avoids loading whole models. Instead loads just structs which will be returned from the API.

---

## Discussion

### Comment by @codecov-io on February 11, 2020 at 03:27 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=h1) Report
> Merging [#123](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/90322f0586e4f4200481314af4bef4db044b8120?src=pr&el=desc) will **increase** coverage by `2.2%`.
> The diff coverage is `96.62%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##           master     #123     +/-   ##
=========================================
+ Coverage   55.21%   57.42%   +2.2%     
=========================================
  Files          40       41      +1     
  Lines        1581     1597     +16     
=========================================
+ Hits          873      917     +44     
+ Misses        604      572     -32     
- Partials      104      108      +4
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `78.57% <100%> (-3.07%)` | :arrow_down: |
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `77.77% <100%> (ø)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `91.89% <100%> (-0.8%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `88.88% <100%> (-2.54%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `77.17% <100%> (-1.18%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `75.67% <100%> (-3.4%)` | :arrow_down: |
| [base/database/query.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9xdWVyeS5nbw==) | `93.47% <93.47%> (ø)` | |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `68.57% <0%> (+29.52%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=footer). Last update [90322f0...9950156](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/123?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on February 12, 2020 at 12:12 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/123*
