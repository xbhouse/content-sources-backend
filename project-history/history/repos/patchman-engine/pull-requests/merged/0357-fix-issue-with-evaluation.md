---
type: pull_request
number: 357
title: "Fix issue with evaluation"
state: merged
author: semtexzv
created: 2020-10-01T09:04:15Z
updated: 2021-03-16T09:29:29Z
closed: 2020-10-01T10:41:07Z
merged: 2020-10-01T10:41:07Z
base_branch: master
head_branch: fix-eval-join
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/357
---

# Pull Request #357: Fix issue with evaluation

**Author**: @semtexzv
**Created**: October 01, 2020 at 09:04 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-eval-join`

## Description

We perform left join and in previous PR moved additional join params to where clase, which changed behavior of whole query.

---

## Discussion

### Comment by @codecov-commenter on October 01, 2020 at 10:10 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=h1) Report
> Merging [#357](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/79760775a975d2c7bed11cdeb733ecfb1dae4a69?el=desc) will **not change** coverage.
> The diff coverage is `66.66%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #357   +/-   ##
=======================================
  Coverage   62.52%   62.52%           
=======================================
  Files          56       56           
  Lines        2492     2492           
=======================================
  Hits         1558     1558           
  Misses        715      715           
  Partials      219      219           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.38% <66.66%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=footer). Last update [7976077...10cfd21](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/357?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on October 01, 2020 at 09:06 AM UTC

### Review by @semtexzv - Commented on October 01, 2020 at 10:01 AM UTC

### Review by @josef-hak - Commented on October 01, 2020 at 10:12 AM UTC

### Review by @josef-hak - Approved on October 01, 2020 at 10:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/357*
