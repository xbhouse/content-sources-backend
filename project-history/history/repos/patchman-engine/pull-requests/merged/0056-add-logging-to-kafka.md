---
type: pull_request
number: 56
title: "Add logging to kafka"
state: merged
author: semtexzv
created: 2020-01-14T09:26:22Z
updated: 2021-03-16T09:26:52Z
closed: 2020-01-14T13:43:53Z
merged: 2020-01-14T13:43:53Z
base_branch: master
head_branch: listener-logging
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/56
---

# Pull Request #56: Add logging to kafka

**Author**: @semtexzv
**Created**: January 14, 2020 at 09:26 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener-logging`

## Description

Add logging to listener in order to diagnose restarts.

---

## Discussion

### Comment by @codecov-io on January 14, 2020 at 09:32 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=h1) Report
> Merging [#56](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/949d233a2a1ad7d2683d69cc67f3c0c59d73706f?src=pr&el=desc) will **decrease** coverage by `0.04%`.
> The diff coverage is `50%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #56      +/-   ##
==========================================
- Coverage   58.16%   58.12%   -0.05%     
==========================================
  Files          29       29              
  Lines        1114     1120       +6     
==========================================
+ Hits          648      651       +3     
- Misses        392      395       +3     
  Partials       74       74
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `50% <50%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=footer). Last update [949d233...86cf0f5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/56?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @beav on January 14, 2020 at 01:15 PM UTC

messages now appear:

```
time="2020-01-14T13:14:31Z" level=error msg="Syncing [2 3 %!d(string=main@f63682231057 (github.com/segmentio/kafka-go)-fa34335c-e8e4-
4752-9c0e-4124991f4beb)] assignments for generation %!d(MISSING) as member %!s(MISSING)" type=kafka
```

---

## Reviews

### Review by @beav - Approved on January 14, 2020 at 01:15 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/56*
