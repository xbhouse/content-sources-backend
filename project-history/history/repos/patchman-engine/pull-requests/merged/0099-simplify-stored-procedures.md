---
type: pull_request
number: 99
title: "Simplify stored procedures"
state: merged
author: semtexzv
created: 2020-02-05T10:21:54Z
updated: 2021-03-16T09:27:13Z
closed: 2020-02-05T14:59:56Z
merged: 2020-02-05T14:59:56Z
base_branch: master
head_branch: simplify-sql
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/99
---

# Pull Request #99: Simplify stored procedures

**Author**: @semtexzv
**Created**: February 05, 2020 at 10:21 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `simplify-sql`

## Description

- Remove duplication between multiple `refresh_*` functions
- Simplify opt_out trigger
- Remove all tables and functions before testing migrations

---

## Discussion

### Comment by @codecov-io on February 05, 2020 at 10:31 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99?src=pr&el=h1) Report
> Merging [#99](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1c69a9a5379b462e3b37f6a1320b533c74aceae8?src=pr&el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master      #99   +/-   ##
=======================================
  Coverage   55.75%   55.75%           
=======================================
  Files          35       35           
  Lines        1363     1363           
=======================================
  Hits          760      760           
  Misses        512      512           
  Partials       91       91
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99?src=pr&el=footer). Last update [1c69a9a...a3bec8d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/99?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on February 05, 2020 at 02:59 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/99*
