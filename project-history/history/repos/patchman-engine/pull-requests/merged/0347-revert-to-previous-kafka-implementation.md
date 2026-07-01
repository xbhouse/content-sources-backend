---
type: pull_request
number: 347
title: "Revert to previous kafka implementation"
state: merged
author: semtexzv
created: 2020-09-25T12:27:55Z
updated: 2021-03-16T09:28:56Z
closed: 2020-09-25T13:30:52Z
merged: 2020-09-25T13:30:52Z
base_branch: master
head_branch: revert-kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/347
---

# Pull Request #347: Revert to previous kafka implementation

**Author**: @semtexzv
**Created**: September 25, 2020 at 12:27 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `revert-kafka`

## Description

This reverts commit b81fedfaa9756d63daea76f6374c8cc59b5e1d85.

Revert "Change kafka implementation"

This reverts commit 3108a0fd2afbf9c860ef27eaef7c71cde8ffb0c4.

---

## Discussion

### Comment by @codecov-commenter on September 25, 2020 at 12:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=h1) Report
> Merging [#347](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b81fedfaa9756d63daea76f6374c8cc59b5e1d85?el=desc) will **increase** coverage by `0.16%`.
> The diff coverage is `56.36%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #347      +/-   ##
==========================================
+ Coverage   61.68%   61.85%   +0.16%     
==========================================
  Files          56       56              
  Lines        2461     2464       +3     
==========================================
+ Hits         1518     1524       +6     
+ Misses        724      721       -3     
  Partials      219      219              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/message.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWV0cmljcy5nbw==) | `33.33% <ø> (+33.33%)` | :arrow_up: |
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `100.00% <ø> (+100.00%)` | :arrow_up: |
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `53.84% <0.00%> (-9.80%)` | :arrow_down: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `49.09% <ø> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.02% <40.00%> (-0.32%)` | :arrow_down: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `78.94% <40.00%> (-8.56%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `66.66% <61.53%> (-1.67%)` | :arrow_down: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `61.53% <100.00%> (ø)` | |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=footer). Last update [b81fedf...1e034fa](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/347?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/347*
