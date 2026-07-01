---
type: pull_request
number: 185
title: "Fix cache calculation"
state: merged
author: semtexzv
created: 2020-03-24T11:18:23Z
updated: 2020-03-25T09:30:09Z
closed: 2020-03-25T09:30:08Z
merged: 2020-03-25T09:30:08Z
base_branch: master
head_branch: fix-caches-calc
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/185
---

# Pull Request #185: Fix cache calculation

**Author**: @semtexzv
**Created**: March 24, 2020 at 11:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-caches-calc`

## Description

There were several issues which caused invalid caches.
- Not locking `advisory_account_data` in between 2 updates when evaluation could result in an inconsistencies when evaluating 2 package profiles with same advisories from same account
- Not locking `system_platform` rows when performing system culling tasks
- Not locking `system_platform` when inserting from listener

---

## Discussion

### Comment by @codecov-io on March 24, 2020 at 11:28 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=h1) Report
> Merging [#185](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/eefd7a084cc4263249ed505bc15e88d41451f248&el=desc) will **decrease** coverage by `1.80%`.
> The diff coverage is `40.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #185      +/-   ##
==========================================
- Coverage   65.14%   63.34%   -1.81%     
==========================================
  Files          44       46       +2     
  Lines        1836     1904      +68     
==========================================
+ Hits         1196     1206      +10     
- Misses        501      557      +56     
- Partials      139      141       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.44% <93.10%> (+1.83%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `75.86% <100.00%> (-0.28%)` | :arrow_down: |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `24.00% <0.00%> (-0.33%)` | :arrow_down: |
| [manager/controllers/test\_utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100.00% <0.00%> (ø)` | |
| [manager/controllers/health.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9oZWFsdGguZ28=) | | |
| [base/core/gintesting.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2dpbnRlc3RpbmcuZ28=) | `83.33% <0.00%> (ø)` | |
| [base/core/probes.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL3Byb2Jlcy5nbw==) | `45.45% <0.00%> (ø)` | |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <0.00%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=footer). Last update [eefd7a0...1e73dd4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/185?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on March 24, 2020 at 11:41 AM UTC

Great, thanks for fixes. I've just added some notes.

### Review by @semtexzv - Commented on March 25, 2020 at 09:04 AM UTC

### Review by @semtexzv - Commented on March 25, 2020 at 09:05 AM UTC

### Review by @semtexzv - Commented on March 25, 2020 at 09:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/185*
