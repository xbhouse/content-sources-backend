---
type: pull_request
number: 226
title: "Added Kafka readers closing using WaitGroup"
state: merged
author: josef-hak
created: 2020-04-16T09:35:42Z
updated: 2020-04-17T12:23:10Z
closed: 2020-04-17T12:21:32Z
merged: 2020-04-17T12:21:32Z
base_branch: master
head_branch: go-wait-group
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/226
---

# Pull Request #226: Added Kafka readers closing using WaitGroup

**Author**: @josef-hak
**Created**: April 16, 2020 at 09:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `go-wait-group`

## Description

- also added testing upload loop

---

## Discussion

### Comment by @codecov-io on April 16, 2020 at 02:37 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=h1) Report
> Merging [#226](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/f0b799e193f0b82c9a75e7231d2d76b21c3e416a&el=desc) will **decrease** coverage by `0.17%`.
> The diff coverage is `53.84%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #226      +/-   ##
==========================================
- Coverage   62.83%   62.65%   -0.18%     
==========================================
  Files          46       46              
  Lines        2034     2046      +12     
==========================================
+ Hits         1278     1282       +4     
- Misses        608      616       +8     
  Partials      148      148              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.22% <33.33%> (-0.54%)` | :arrow_down: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `72.22% <42.85%> (-9.03%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `70.96% <69.23%> (-3.11%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=footer). Last update [f0b799e...0e7ccfe](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/226?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on April 16, 2020 at 10:36 AM UTC

### Review by @josef-hak - Commented on April 16, 2020 at 11:19 AM UTC

### Review by @semtexzv - Commented on April 16, 2020 at 11:40 AM UTC

### Review by @josef-hak - Commented on April 16, 2020 at 11:49 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/226*
