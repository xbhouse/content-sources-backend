---
type: pull_request
number: 49
title: "Fix indexing error in advisory_sync"
state: merged
author: semtexzv
created: 2020-01-10T09:01:45Z
updated: 2021-03-16T09:26:50Z
closed: 2020-01-10T09:18:49Z
merged: 2020-01-10T09:18:49Z
base_branch: master
head_branch: fix-advisory-sync
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/49
---

# Pull Request #49: Fix indexing error in advisory_sync

**Author**: @semtexzv
**Created**: January 10, 2020 at 09:01 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-advisory-sync`

## Description

An indexing error resulted in vmaas_sync container crashing after each sync.

---

## Discussion

### Comment by @codecov-io on January 10, 2020 at 09:08 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=h1) Report
> Merging [#49](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ff6dc5ec4ec07879349dbc2feec90eaa385a8cce?src=pr&el=desc) will **not change** coverage.
> The diff coverage is `0%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=tree)

```diff
@@          Coverage Diff           @@
##           master     #49   +/-   ##
======================================
  Coverage    56.2%   56.2%           
======================================
  Files          28      28           
  Lines         943     943           
======================================
  Hits          530     530           
  Misses        355     355           
  Partials       58      58
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `36.98% <0%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=footer). Last update [ff6dc5e...0eb48eb](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/49?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on January 10, 2020 at 09:18 AM UTC

Ok, I hope it will solve our problem :smile: 


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/49*
