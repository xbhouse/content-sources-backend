---
type: pull_request
number: 186
title: "Remove null characters from input"
state: merged
author: semtexzv
created: 2020-03-25T09:38:59Z
updated: 2021-03-16T09:27:53Z
closed: 2020-03-25T13:00:31Z
merged: 2020-03-25T13:00:31Z
base_branch: master
head_branch: remove-nulls
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/186
---

# Pull Request #186: Remove null characters from input

**Author**: @semtexzv
**Created**: March 25, 2020 at 09:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `remove-nulls`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on March 25, 2020 at 12:29 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=h1) Report
> Merging [#186](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/96b1018bf3b6cad05e513d5a2c7d6255df6d438d&el=desc) will **decrease** coverage by `1.78%`.
> The diff coverage is `50.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #186      +/-   ##
==========================================
- Coverage   64.89%   63.11%   -1.79%     
==========================================
  Files          46       46              
  Lines        1849     1922      +73     
==========================================
+ Hits         1200     1213      +13     
- Misses        509      567      +58     
- Partials      140      142       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/query.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9xdWVyeS5nbw==) | `82.71% <0.00%> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `76.40% <100.00%> (ø)` | |
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `6.38% <0.00%> (-1.31%)` | :arrow_down: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `71.83% <0.00%> (-0.75%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `75.86% <0.00%> (-0.28%)` | :arrow_down: |
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `26.15% <0.00%> (+1.15%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.44% <0.00%> (+1.83%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=footer). Last update [96b1018...c1b7085](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/186?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on March 25, 2020 at 09:58 AM UTC

### Review by @josef-hak - Approved on March 25, 2020 at 01:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/186*
