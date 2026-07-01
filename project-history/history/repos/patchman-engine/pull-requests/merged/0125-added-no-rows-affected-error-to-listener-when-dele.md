---
type: pull_request
number: 125
title: "Added no rows affected error to listener when deleting"
state: merged
author: josef-hak
created: 2020-02-12T07:23:33Z
updated: 2020-02-12T15:43:37Z
closed: 2020-02-12T15:31:19Z
merged: 2020-02-12T15:31:19Z
base_branch: master
head_branch: delete-no-rows
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/125
---

# Pull Request #125: Added no rows affected error to listener when deleting

**Author**: @josef-hak
**Created**: February 12, 2020 at 07:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `delete-no-rows`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 12, 2020 at 12:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=h1) Report
> Merging [#125](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2ff84324ddbd41a3f5194f9a463355af4047cb3b?src=pr&el=desc) will **decrease** coverage by `0.35%`.
> The diff coverage is `0%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #125      +/-   ##
==========================================
- Coverage   57.17%   56.81%   -0.36%     
==========================================
  Files          40       40              
  Lines        1581     1591      +10     
==========================================
  Hits          904      904              
- Misses        570      577       +7     
- Partials      107      110       +3
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0% <ø> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `30.43% <0%> (-6.41%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.69% <0%> (-1.86%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=footer). Last update [2ff8432...fb5ac5c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/125?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Commented on February 12, 2020 at 09:48 AM UTC

### Review by @josef-hak - Commented on February 12, 2020 at 12:05 PM UTC

### Review by @beav - Approved on February 12, 2020 at 12:56 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/125*
