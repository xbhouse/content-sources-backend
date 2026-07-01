---
type: pull_request
number: 327
title: "Change kafka implementation"
state: merged
author: semtexzv
created: 2020-09-09T15:07:19Z
updated: 2021-03-16T09:28:55Z
closed: 2020-09-25T11:45:25Z
merged: 2020-09-25T11:45:25Z
base_branch: master
head_branch: kafka-lib
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/327
---

# Pull Request #327: Change kafka implementation

**Author**: @semtexzv
**Created**: September 09, 2020 at 03:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafka-lib`

## Description

Changes kafka implementation to use new library. 

Do not merge, it should be beneficial to compare performance of evaluator before and after.

---

## Discussion

### Comment by @codecov-commenter on September 23, 2020 at 10:55 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=h1) Report
> Merging [#327](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1f5521fad47f9b35accfb571fa05c40d75a98544?el=desc) will **decrease** coverage by `0.17%`.
> The diff coverage is `69.23%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #327      +/-   ##
==========================================
- Coverage   61.38%   61.21%   -0.18%     
==========================================
  Files          54       54              
  Lines        2380     2377       -3     
==========================================
- Hits         1461     1455       -6     
- Misses        706      709       +3     
  Partials      213      213              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/message.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWV0cmljcy5nbw==) | `0.00% <ø> (-33.34%)` | :arrow_down: |
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `0.00% <0.00%> (-100.00%)` | :arrow_down: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `63.63% <ø> (+9.79%)` | :arrow_up: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `49.09% <ø> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.34% <50.00%> (+0.31%)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `87.50% <50.00%> (+8.55%)` | :arrow_up: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `68.33% <73.80%> (+1.66%)` | :arrow_up: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <100.00%> (ø)` | |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=footer). Last update [1f5521f...f27dec8](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/327?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on September 23, 2020 at 06:33 PM UTC

Main goal is to simplify the application, leave a lot of partition handling to the library. I don't want to complicate things further by using 2 libraries at the same time

### Comment by @josef-hak on September 23, 2020 at 06:49 PM UTC

@semtexzv ok, so just please process my "wg" note, I can approve it then.

---

## Reviews

### Review by @josef-hak - Commented on September 23, 2020 at 11:27 AM UTC

It looks good. But as you mentioned perf. testing... wouldn't it be useful to keep old implementation usable as well and allow switching between these two frameworks using env variable?

### Review by @josef-hak - Approved on September 24, 2020 at 10:45 AM UTC

Ok, I hope it will work. Thanks for big upgrade.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/327*
