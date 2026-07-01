---
type: pull_request
number: 183
title: "Added liveness and readiness probes"
state: merged
author: josef-hak
created: 2020-03-24T10:30:57Z
updated: 2020-03-24T14:32:37Z
closed: 2020-03-24T13:05:00Z
merged: 2020-03-24T13:05:00Z
base_branch: master
head_branch: probes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/183
---

# Pull Request #183: Added liveness and readiness probes

**Author**: @josef-hak
**Created**: March 24, 2020 at 10:30 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `probes`

## Description

- added probes and used in components
- added tests

---

## Discussion

### Comment by @codecov-io on March 24, 2020 at 10:40 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=h1) Report
> Merging [#183](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/eefd7a084cc4263249ed505bc15e88d41451f248&el=desc) will **decrease** coverage by `0.24%`.
> The diff coverage is `65.38%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #183      +/-   ##
==========================================
- Coverage   65.14%   64.89%   -0.25%     
==========================================
  Files          44       46       +2     
  Lines        1836     1849      +13     
==========================================
+ Hits         1196     1200       +4     
- Misses        501      509       +8     
- Partials      139      140       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [manager/controllers/test\_utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy90ZXN0X3V0aWxzLmdv) | `100.00% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `24.00% <0.00%> (-0.33%)` | :arrow_down: |
| [base/core/probes.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL3Byb2Jlcy5nbw==) | `45.45% <45.45%> (ø)` | |
| [base/core/gintesting.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2dpbnRlc3RpbmcuZ28=) | `83.33% <83.33%> (ø)` | |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `66.66% <100.00%> (+4.16%)` | :arrow_up: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `66.66% <100.00%> (+4.16%)` | :arrow_up: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=footer). Last update [eefd7a0...cda8ffd](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/183?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on March 24, 2020 at 11:49 AM UTC

:+1: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/183*
