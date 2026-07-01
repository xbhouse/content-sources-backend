---
type: pull_request
number: 228
title: "added manager termination handling"
state: merged
author: josef-hak
created: 2020-04-17T08:59:28Z
updated: 2020-04-17T12:22:16Z
closed: 2020-04-17T12:21:08Z
merged: 2020-04-17T12:21:08Z
base_branch: master
head_branch: manager-term
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/228
---

# Pull Request #228: added manager termination handling

**Author**: @josef-hak
**Created**: April 17, 2020 at 08:59 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `manager-term`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on April 17, 2020 at 11:52 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=h1) Report
> Merging [#228](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f0b799e193f0b82c9a75e7231d2d76b21c3e416a&el=desc) will **decrease** coverage by `0.00%`.
> The diff coverage is `58.82%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #228      +/-   ##
==========================================
- Coverage   62.83%   62.82%   -0.01%     
==========================================
  Files          46       46              
  Lines        2034     2047      +13     
==========================================
+ Hits         1278     1286       +8     
- Misses        608      611       +3     
- Partials      148      150       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/debug\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9kZWJ1Z19hcGkuZ28=) | `0.00% <0.00%> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `35.55% <0.00%> (ø)` | |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9naW4uZ28=) | `21.62% <61.53%> (+21.62%)` | :arrow_up: |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `66.66% <100.00%> (ø)` | |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `66.66% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=footer). Last update [f0b799e...d5c7f66](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/228?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on April 17, 2020 at 10:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/228*
