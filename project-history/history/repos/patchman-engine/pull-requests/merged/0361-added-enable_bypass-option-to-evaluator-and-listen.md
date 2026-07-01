---
type: pull_request
number: 361
title: "Added ENABLE_BYPASS option to evaluator and listener"
state: merged
author: josef-hak
created: 2020-10-01T15:39:50Z
updated: 2020-10-29T11:33:12Z
closed: 2020-10-02T07:15:21Z
merged: 2020-10-02T07:15:21Z
base_branch: master
head_branch: eval-bypass
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/361
---

# Pull Request #361: Added ENABLE_BYPASS option to evaluator and listener

**Author**: @josef-hak
**Created**: October 01, 2020 at 03:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `eval-bypass`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on October 01, 2020 at 05:51 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=h1) Report
> Merging [#361](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fd945f58e68459686e7a9b3b035259175b7217e4?el=desc) will **decrease** coverage by `0.16%`.
> The diff coverage is `33.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #361      +/-   ##
==========================================
- Coverage   62.44%   62.27%   -0.17%     
==========================================
  Files          56       56              
  Lines        2524     2534      +10     
==========================================
+ Hits         1576     1578       +2     
- Misses        724      730       +6     
- Partials      224      226       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `67.66% <20.00%> (-0.66%)` | :arrow_down: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.76% <33.33%> (-3.33%)` | :arrow_down: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `80.00% <100.00%> (+1.05%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=footer). Last update [fd945f5...0d2fca5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/361?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on October 02, 2020 at 07:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/361*
