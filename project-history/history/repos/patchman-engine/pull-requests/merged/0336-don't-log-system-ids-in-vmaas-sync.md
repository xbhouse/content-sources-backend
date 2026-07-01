---
type: pull_request
number: 336
title: "Don't log system ids in vmaas sync"
state: merged
author: semtexzv
created: 2020-09-15T03:40:19Z
updated: 2021-03-16T09:28:53Z
closed: 2020-10-02T07:32:59Z
merged: 2020-10-02T07:32:59Z
base_branch: master
head_branch: vmaas-sync-log
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/336
---

# Pull Request #336: Don't log system ids in vmaas sync

**Author**: @semtexzv
**Created**: September 15, 2020 at 03:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `vmaas-sync-log`

## Description

Last recalc we spent 15 minutes just logging all of the systems and sending messages

---

## Discussion

### Comment by @codecov-commenter on September 15, 2020 at 03:48 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=h1) Report
> Merging [#336](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b0bb5fcbc99869111a60167b903c296616ec23b8?el=desc) will **decrease** coverage by `0.01%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #336      +/-   ##
==========================================
- Coverage   63.30%   63.29%   -0.02%     
==========================================
  Files          54       54              
  Lines        2679     2678       -1     
==========================================
- Hits         1696     1695       -1     
  Misses        772      772              
  Partials      211      211              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/send\_messages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zZW5kX21lc3NhZ2VzLmdv) | `60.60% <ø> (-1.16%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=footer). Last update [b0bb5fc...3283136](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/336?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on September 16, 2020 at 07:30 AM UTC

@semtexzv yes, but it's debug log. Wouldn't it better to switch to "INFO" log level in prod rather than removing it?

---

## Reviews

### Review by @josef-hak - Changes Requested on September 18, 2020 at 11:23 AM UTC

### Review by @josef-hak - Commented on October 02, 2020 at 07:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/336*
