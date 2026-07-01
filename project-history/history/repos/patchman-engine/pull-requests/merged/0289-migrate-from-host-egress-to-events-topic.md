---
type: pull_request
number: 289
title: "Migrate from host-egress to events topic"
state: merged
author: semtexzv
created: 2020-07-16T10:19:34Z
updated: 2021-03-16T09:25:59Z
closed: 2020-07-21T08:32:52Z
merged: 2020-07-21T08:32:52Z
base_branch: master
head_branch: listener-topic
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/289
---

# Pull Request #289: Migrate from host-egress to events topic

**Author**: @semtexzv
**Created**: July 16, 2020 at 10:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener-topic`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on July 16, 2020 at 10:28 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=h1) Report
> Merging [#289](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/acae0c620c69af0f54ebc8c28e97e9703cdccb29&el=desc) will **decrease** coverage by `0.28%`.
> The diff coverage is `22.22%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #289      +/-   ##
==========================================
- Coverage   62.20%   61.92%   -0.29%     
==========================================
  Files          51       51              
  Lines        2408     2403       -5     
==========================================
- Hits         1498     1488      -10     
- Misses        722      727       +5     
  Partials      188      188              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `70.58% <ø> (-3.10%)` | :arrow_down: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `46.83% <12.50%> (-3.85%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.14% <100.00%> (-0.45%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=footer). Last update [acae0c6...4328b2c](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/289?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on July 20, 2020 at 10:29 AM UTC

Please resolve the conflicts.

### Review by @josef-hak - Approved on July 21, 2020 at 08:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/289*
