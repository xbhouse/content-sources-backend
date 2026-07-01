---
type: pull_request
number: 262
title: "Implement filters for export endpoints"
state: merged
author: semtexzv
created: 2020-05-25T11:21:56Z
updated: 2021-03-16T09:28:25Z
closed: 2020-05-28T08:21:46Z
merged: 2020-05-28T08:21:46Z
base_branch: master
head_branch: export-filter
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/262
---

# Pull Request #262: Implement filters for export endpoints

**Author**: @semtexzv
**Created**: May 25, 2020 at 11:21 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `export-filter`

## Description

- Adds a way to filter exported data, format is preserved between API and export endpoints

---

## Discussion

### Comment by @codecov-commenter on May 25, 2020 at 12:03 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=h1) Report
> Merging [#262](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/18b24f817ac8d631559ab0a2df4db97cf25c5f4e&el=desc) will **decrease** coverage by `0.16%`.
> The diff coverage is `47.36%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #262      +/-   ##
==========================================
- Coverage   62.66%   62.50%   -0.17%     
==========================================
  Files          47       47              
  Lines        2215     2232      +17     
==========================================
+ Hits         1388     1395       +7     
- Misses        665      671       +6     
- Partials      162      166       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `73.45% <40.00%> (-3.25%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `77.77% <50.00%> (-2.62%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `81.96% <60.00%> (-2.25%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=footer). Last update [18b24f8...632738e](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/262?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on May 27, 2020 at 01:53 PM UTC

Great, looks good. Just one code-style question and I am not sure whether it's not possible to keep or increase coverage.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/262*
