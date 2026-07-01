---
type: pull_request
number: 33
title: "Implement Filter parsing & SQL generation"
state: merged
author: semtexzv
created: 2019-12-09T14:05:45Z
updated: 2020-02-11T12:19:07Z
closed: 2020-02-11T12:19:07Z
merged: 2020-02-11T12:19:07Z
base_branch: master
head_branch: api-filters
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/33
---

# Pull Request #33: Implement Filter parsing & SQL generation

**Author**: @semtexzv
**Created**: December 09, 2019 at 02:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `api-filters`

## Description

Further work depending on #82 .

---

## Discussion

### Comment by @codecov-io on January 30, 2020 at 03:34 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=h1) Report
> Merging [#33](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0750df6510633a58aa424f903a3a7b7206d36d36?src=pr&el=desc) will **increase** coverage by `1.06%`.
> The diff coverage is `80.86%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #33      +/-   ##
==========================================
+ Coverage   54.15%   55.21%   +1.06%     
==========================================
  Files          38       40       +2     
  Lines        1479     1581     +102     
==========================================
+ Hits          801      873      +72     
- Misses        581      604      +23     
- Partials       97      104       +7
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `81.63% <100%> (ø)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `91.42% <100%> (ø)` | :arrow_up: |
| [manager/controllers/paging.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWdpbmcuZ28=) | `100% <100%> (ø)` | :arrow_up: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `92.68% <100%> (-0.18%)` | :arrow_down: |
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `79.06% <100%> (ø)` | :arrow_up: |
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `77.77% <77.77%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `80.2% <77.77%> (-3.38%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `62.8% <0%> (-1.9%)` | :arrow_down: |
| [base/utils/http.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9odHRwLmdv) | `0% <0%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=footer). Last update [0750df6...883a3f3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/33?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on February 10, 2020 at 04:12 PM UTC

Good, just a few ideas what to improve.

### Review by @semtexzv - Commented on February 11, 2020 at 10:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/33*
