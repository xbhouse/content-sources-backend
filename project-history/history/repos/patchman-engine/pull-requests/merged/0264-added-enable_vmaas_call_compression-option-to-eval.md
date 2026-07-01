---
type: pull_request
number: 264
title: "Added ENABLE_VMAAS_CALL_COMPRESSION option to evaluator"
state: merged
author: josef-hak
created: 2020-05-27T16:06:28Z
updated: 2020-06-01T10:41:10Z
closed: 2020-06-01T10:26:12Z
merged: 2020-06-01T10:26:12Z
base_branch: master
head_branch: gzip_request
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/264
---

# Pull Request #264: Added ENABLE_VMAAS_CALL_COMPRESSION option to evaluator

**Author**: @josef-hak
**Created**: May 27, 2020 at 04:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `gzip_request`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on May 27, 2020 at 04:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=h1) Report
> Merging [#264](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/18b24f817ac8d631559ab0a2df4db97cf25c5f4e&el=desc) will **decrease** coverage by `0.31%`.
> The diff coverage is `33.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #264      +/-   ##
==========================================
- Coverage   62.66%   62.34%   -0.32%     
==========================================
  Files          47       47              
  Lines        2215     2244      +29     
==========================================
+ Hits         1388     1399      +11     
- Misses        665      679      +14     
- Partials      162      166       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/core.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9jb3JlLmdv) | `14.28% <0.00%> (-2.08%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.57% <100.00%> (+0.38%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `73.45% <0.00%> (-3.25%)` | :arrow_down: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `77.77% <0.00%> (-2.62%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `81.96% <0.00%> (-2.25%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=footer). Last update [18b24f8...2c4e128](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/264?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on May 28, 2020 at 08:21 AM UTC

### Review by @josef-hak - Commented on May 28, 2020 at 11:21 AM UTC

### Review by @semtexzv - Approved on June 01, 2020 at 08:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/264*
