---
type: pull_request
number: 193
title: "Implement retrying for message handling"
state: closed
author: semtexzv
created: 2020-03-27T10:27:47Z
updated: 2020-03-31T08:16:39Z
closed: 2020-03-31T08:16:39Z
base_branch: master
head_branch: retry-experiments
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/193
---

# Pull Request #193: Implement retrying for message handling

**Author**: @semtexzv
**Created**: March 27, 2020 at 10:27 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `retry-experiments`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on March 27, 2020 at 10:37 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=h1) Report
> Merging [#193](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a9a85e5f4068284804d04ff796bdc98c5bf7c4e6&el=desc) will **increase** coverage by `0.04%`.
> The diff coverage is `67.64%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #193      +/-   ##
==========================================
+ Coverage   63.11%   63.15%   +0.04%     
==========================================
  Files          46       46              
  Lines        1922     1938      +16     
==========================================
+ Hits         1213     1224      +11     
- Misses        567      571       +4     
- Partials      142      143       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `76.00% <60.00%> (+0.13%)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `68.75% <66.66%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `75.00% <66.66%> (-3.05%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.25% <66.66%> (-0.20%)` | :arrow_down: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `83.33% <83.33%> (+0.72%)` | :arrow_up: |
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `62.85% <0.00%> (+1.09%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=footer). Last update [a9a85e5...b5f9ce3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/193?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on March 27, 2020 at 02:26 PM UTC

Good, just some cosmetics proposals. Feel free to merge it if you wish.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/193*
