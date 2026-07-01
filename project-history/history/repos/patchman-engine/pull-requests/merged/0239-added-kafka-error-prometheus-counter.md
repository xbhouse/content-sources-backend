---
type: pull_request
number: 239
title: "Added kafka error Prometheus counter"
state: merged
author: josef-hak
created: 2020-04-24T08:10:40Z
updated: 2020-04-28T10:11:05Z
closed: 2020-04-28T09:12:58Z
merged: 2020-04-28T09:12:58Z
base_branch: master
head_branch: metrics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/239
---

# Pull Request #239: Added kafka error Prometheus counter

**Author**: @josef-hak
**Created**: April 24, 2020 at 08:10 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `metrics`

## Description

- Added prometheus counters counting Kafka connection read/write errors.
- Fixed TestRunServer test.
- Fixed and unified Kafka connection errors logging.

---

## Discussion

### Comment by @codecov-io on April 24, 2020 at 10:46 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=h1) Report
> Merging [#239](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ce2016635616534deaa1df88cd42995c92cb7061&el=desc) will **increase** coverage by `0.01%`.
> The diff coverage is `70.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #239      +/-   ##
==========================================
+ Coverage   62.76%   62.78%   +0.01%     
==========================================
  Files          46       47       +1     
  Lines        2095     2104       +9     
==========================================
+ Hits         1315     1321       +6     
- Misses        626      629       +3     
  Partials      154      154              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbWV0cmljcy5nbw==) | `20.00% <20.00%> (ø)` | |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `73.84% <85.71%> (+2.87%)` | :arrow_up: |
| [base/core/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL2NvbmZpZy5nbw==) | `100.00% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=footer). Last update [ce20166...fc4473d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/239?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on April 28, 2020 at 08:30 AM UTC

### Review by @josef-hak - Commented on April 28, 2020 at 08:38 AM UTC

### Review by @semtexzv - Approved on April 28, 2020 at 09:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/239*
