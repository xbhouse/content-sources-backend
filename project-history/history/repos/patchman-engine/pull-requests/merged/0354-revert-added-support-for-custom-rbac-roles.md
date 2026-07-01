---
type: pull_request
number: 354
title: "Revert \"added support for custom rbac roles\""
state: merged
author: josef-hak
created: 2020-09-30T08:16:51Z
updated: 2020-10-01T17:35:13Z
closed: 2020-09-30T10:22:23Z
merged: 2020-09-30T10:22:23Z
base_branch: master
head_branch: revert-custom-rbac
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/354
---

# Pull Request #354: Revert "added support for custom rbac roles"

**Author**: @josef-hak
**Created**: September 30, 2020 at 08:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `revert-custom-rbac`

## Description

This reverts commit 6687d81e

---

## Discussion

### Comment by @codecov-commenter on September 30, 2020 at 08:25 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=h1) Report
> Merging [#354](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8b2f4004d5861751bcfe28658c564a1dc856e2e5?el=desc) will **increase** coverage by `0.04%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #354      +/-   ##
==========================================
+ Coverage   62.33%   62.37%   +0.04%     
==========================================
  Files          56       56              
  Lines        2506     2509       +3     
==========================================
+ Hits         1562     1565       +3     
  Misses        725      725              
  Partials      219      219              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `69.29% <0.00%> (+0.27%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=footer). Last update [8b2f400...1b722b7](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/354?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on September 30, 2020 at 10:22 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/354*
