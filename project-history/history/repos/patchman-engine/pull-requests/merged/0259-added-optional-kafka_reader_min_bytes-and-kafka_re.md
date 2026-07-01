---
type: pull_request
number: 259
title: "Added optional KAFKA_READER_MIN_BYTES and KAFKA_READER_MAX_BYTES vars"
state: merged
author: josef-hak
created: 2020-05-20T07:30:46Z
updated: 2020-05-27T18:43:52Z
closed: 2020-05-20T10:21:01Z
merged: 2020-05-20T10:21:01Z
base_branch: master
head_branch: kafka-max-bytes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/259
---

# Pull Request #259: Added optional KAFKA_READER_MIN_BYTES and KAFKA_READER_MAX_BYTES vars

**Author**: @josef-hak
**Created**: May 20, 2020 at 07:30 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafka-max-bytes`

## Description

There were some issues due to Kafka settings in perf-testing environment. Kafka reader config update solve this problem. I've added options to change Kafka reader config using env vars.

---

## Discussion

### Comment by @codecov-commenter on May 20, 2020 at 07:42 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=h1) Report
> Merging [#259](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/30ecdf9491e692b29f860888a83f2d2d06838195&el=desc) will **increase** coverage by `0.03%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #259      +/-   ##
==========================================
+ Coverage   62.19%   62.22%   +0.03%     
==========================================
  Files          47       47              
  Lines        2161     2163       +2     
==========================================
+ Hits         1344     1346       +2     
  Misses        660      660              
  Partials      157      157              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `74.62% <100.00%> (+0.78%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=footer). Last update [30ecdf9...7c28bb5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/259?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on May 20, 2020 at 09:47 AM UTC

### Review by @josef-hak - Commented on May 20, 2020 at 10:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/259*
