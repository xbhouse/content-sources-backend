---
type: pull_request
number: 160
title: "Fix test data IDs & make tests run under different users"
state: merged
author: semtexzv
created: 2020-03-03T13:55:45Z
updated: 2021-03-16T09:27:37Z
closed: 2020-03-05T09:08:04Z
merged: 2020-03-05T09:08:04Z
base_branch: master
head_branch: tests-rights
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/160
---

# Pull Request #160: Fix test data IDs & make tests run under different users

**Author**: @semtexzv
**Created**: March 03, 2020 at 01:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `tests-rights`

## Description

Gorm interprets id of 0 as nonexistant, always inserting, by changing
all ID fields to start at 1 we fix this issue. Built in sequences
always start at 1 so those should be fine.

Also, this commit adds support for running tests in different components
under user accounts, under which they will be running in production.
This change should help us catch cases, where we forget to add access
rights after changes.

Also, adds missing rights to  vmaas_sync so that systen culling can be performed.

---

## Discussion

### Comment by @codecov-io on March 03, 2020 at 03:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=h1) Report
> Merging [#160](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b00d42a536f296936ae2fb3e0831263a45675b80?src=pr&el=desc) will **decrease** coverage by `0.36%`.
> The diff coverage is `18.18%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #160      +/-   ##
==========================================
- Coverage   65.44%   65.08%   -0.37%     
==========================================
  Files          43       43              
  Lines        1609     1618       +9     
==========================================
  Hits         1053     1053              
- Misses        443      452       +9     
  Partials      113      113
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `9.67% <0%> (-3.96%)` | :arrow_down: |
| [manager/controllers/test\_utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100% <100%> (ø)` | :arrow_up: |
| [manager/controllers/system\_delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fZGVsZXRlLmdv) | `54.54% <100%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=footer). Last update [b00d42a...ea72f00](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/160?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on March 03, 2020 at 03:14 PM UTC

Good, just please check my two notes and we can merge it.

### Review by @josef-hak - Approved on March 05, 2020 at 09:07 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/160*
