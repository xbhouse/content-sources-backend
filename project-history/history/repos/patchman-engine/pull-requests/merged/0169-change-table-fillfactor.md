---
type: pull_request
number: 169
title: "Change table fillfactor"
state: merged
author: semtexzv
created: 2020-03-13T08:32:03Z
updated: 2021-03-16T09:27:41Z
closed: 2020-03-16T10:04:41Z
merged: 2020-03-16T10:04:41Z
base_branch: master
head_branch: table-opt
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/169
---

# Pull Request #169: Change table fillfactor

**Author**: @semtexzv
**Created**: March 13, 2020 at 08:32 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `table-opt`

## Description

Should allow us to get less dead tuples due to [HOT](https://malisper.me/postgres-heap-only-tuples/) updates & more aggressive autovacuuming.

Will consume more disk space, but should significantly reduce performance degradation on insert heavy tables.

After the migration is applied, `VACUUM FULL` on the tables will need to be performed, and it can't be run from within the migrations.

---

## Discussion

### Comment by @codecov-io on March 13, 2020 at 08:40 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=h1) Report
> Merging [#169](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b0b9c52b614228ee4d9fae8b4c83a4d5caa97592&el=desc) will **increase** coverage by `0.23%`.
> The diff coverage is `84.78%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #169      +/-   ##
==========================================
+ Coverage   64.99%   65.23%   +0.23%     
==========================================
  Files          44       44              
  Lines        1757     1772      +15     
==========================================
+ Hits         1142     1156      +14     
- Misses        485      486       +1     
  Partials      130      130              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `62.50% <ø> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.01% <84.44%> (+1.69%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.61% <100.00%> (-0.12%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=footer). Last update [b4eacff...715686b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/169?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on March 16, 2020 at 10:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/169*
