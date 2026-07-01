---
type: pull_request
number: 353
title: "Uploal-evaluation delay metric"
state: merged
author: josef-hak
created: 2020-09-29T15:02:24Z
updated: 2020-10-01T17:36:45Z
closed: 2020-09-30T08:29:18Z
merged: 2020-09-30T08:29:18Z
base_branch: master
head_branch: eval-metric-delay
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/353
---

# Pull Request #353: Uploal-evaluation delay metric

**Author**: @josef-hak
**Created**: September 29, 2020 at 03:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `eval-metric-delay`

## Description

- added evaluation delay metric
- added metrics help strings
- fixed bug in docker-compose.yml (not creating kafka output topic)

---

## Discussion

### Comment by @codecov-commenter on September 29, 2020 at 03:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=h1) Report
> Merging [#353](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8b2f4004d5861751bcfe28658c564a1dc856e2e5?el=desc) will **increase** coverage by `0.30%`.
> The diff coverage is `85.71%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #353      +/-   ##
==========================================
+ Coverage   62.33%   62.64%   +0.30%     
==========================================
  Files          56       56              
  Lines        2506     2500       -6     
==========================================
+ Hits         1562     1566       +4     
+ Misses        725      715      -10     
  Partials      219      219              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `32.46% <0.00%> (+3.73%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.38% <100.00%> (+0.36%)` | :arrow_up: |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `66.66% <100.00%> (+4.16%)` | :arrow_up: |
| [vmaas\_sync/metrics\_db.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzX2RiLmdv) | `59.25% <100.00%> (-1.46%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=footer). Last update [8b2f400...3d44c81](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/353?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @MichaelMraka - Approved on September 30, 2020 at 08:28 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/353*
