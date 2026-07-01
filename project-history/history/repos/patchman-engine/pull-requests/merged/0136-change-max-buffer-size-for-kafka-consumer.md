---
type: pull_request
number: 136
title: "Change max buffer size for kafka consumer"
state: merged
author: semtexzv
created: 2020-02-18T13:00:59Z
updated: 2020-02-18T14:07:16Z
closed: 2020-02-18T14:07:16Z
merged: 2020-02-18T14:07:16Z
base_branch: master
head_branch: buffer-size
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/136
---

# Pull Request #136: Change max buffer size for kafka consumer

**Author**: @semtexzv
**Created**: February 18, 2020 at 01:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `buffer-size`

## Description

The issues in QA environment were probably caused by loading too many messages.

+ 8 consumers  = 8 x 10MB = 80MB in raw buffers
+ deserialized = 2 x 80MB = 160MB in just the data size.

With GOGC being 100 by default, go will allocate 100% of previous memory snapshot before forcing a GC. which will result in 320 MB peak ram. 

320MB + our base memory experienced by listener with no load = 320MB + 30MB = 350MB, just what we experienced.


With max buffer of 1MB per listener:

8 x 1MB = 8MB ...  max buffers
2x 8MB = 16MB ... buffers + deserialized
2x 16MB = 32MB .... Peak before GC

32MB + 30 MB = 80MB , we should be under 128 MB resource limit.


---

## Discussion

### Comment by @codecov-io on February 18, 2020 at 01:10 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=h1) Report
> Merging [#136](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/eb4f5c118991b1bf33ad181816da6f0edcad73d6?src=pr&el=desc) will **not change** coverage.
> The diff coverage is `100%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #136   +/-   ##
=======================================
  Coverage   64.46%   64.46%           
=======================================
  Files          41       41           
  Lines        1604     1604           
=======================================
  Hits         1034     1034           
  Misses        449      449           
  Partials      121      121
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `75% <100%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=footer). Last update [eb4f5c1...1d20026](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/136?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on February 18, 2020 at 02:07 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/136*
