---
type: pull_request
number: 111
title: "Added update check to listener"
state: merged
author: josef-hak
created: 2020-02-10T10:01:06Z
updated: 2020-02-10T13:21:13Z
closed: 2020-02-10T13:17:37Z
merged: 2020-02-10T13:17:37Z
base_branch: master
head_branch: listener-check-update
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/111
---

# Pull Request #111: Added update check to listener

**Author**: @josef-hak
**Created**: February 10, 2020 at 10:01 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener-check-update`

## Description

- Added updated rows count. Throw error when this number == 0

---

## Discussion

### Comment by @codecov-io on February 10, 2020 at 10:12 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=h1) Report
> Merging [#111](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0750df6510633a58aa424f903a3a7b7206d36d36?src=pr&el=desc) will **decrease** coverage by `0.07%`.
> The diff coverage is `25%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #111      +/-   ##
==========================================
- Coverage   54.15%   54.08%   -0.08%     
==========================================
  Files          38       38              
  Lines        1479     1481       +2     
==========================================
  Hits          801      801              
- Misses        581      582       +1     
- Partials       97       98       +1
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `63.63% <25%> (-1.07%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=footer). Last update [0750df6...45535f5](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/111?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @beav - Approved on February 10, 2020 at 12:46 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/111*
