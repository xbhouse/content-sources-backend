---
type: pull_request
number: 178
title: "Parse filter values"
state: merged
author: semtexzv
created: 2020-03-18T10:18:36Z
updated: 2021-03-16T09:27:46Z
closed: 2020-03-18T15:47:10Z
merged: 2020-03-18T15:47:10Z
base_branch: master
head_branch: parse-filters
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/178
---

# Pull Request #178: Parse filter values

**Author**: @semtexzv
**Created**: March 18, 2020 at 10:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `parse-filters`

## Description

We used to rely on raw string values being accepted by postgresql as query parameters.

This PR implements parsing values in filter query parameters so we catch and properly report invalid values.

---

## Discussion

### Comment by @codecov-io on March 18, 2020 at 10:27 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=h1) Report
> Merging [#178](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/be898216a6cc21d147b9f419c7d4ecd250974d07&el=desc) will **decrease** coverage by `0.86%`.
> The diff coverage is `43.75%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #178      +/-   ##
==========================================
- Coverage   65.14%   64.28%   -0.87%     
==========================================
  Files          44       44              
  Lines        1779     1820      +41     
==========================================
+ Hits         1159     1170      +11     
- Misses        488      513      +25     
- Partials      132      137       +5     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/query.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9xdWVyeS5nbw==) | `64.19% <31.57%> (-29.29%)` | :arrow_down: |
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `70.49% <58.33%> (-4.06%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.40% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=footer). Last update [be89821...9d64b5d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/178?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on March 18, 2020 at 03:46 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/178*
