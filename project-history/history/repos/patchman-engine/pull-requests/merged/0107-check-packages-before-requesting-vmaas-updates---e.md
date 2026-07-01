---
type: pull_request
number: 107
title: "Check packages before requesting VMaaS updates - evaluation"
state: merged
author: josef-hak
created: 2020-02-07T09:55:03Z
updated: 2020-02-10T10:08:22Z
closed: 2020-02-07T10:25:04Z
merged: 2020-02-07T10:25:04Z
base_branch: master
head_branch: eval_empty_pkg_list
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/107
---

# Pull Request #107: Check packages before requesting VMaaS updates - evaluation

**Author**: @josef-hak
**Created**: February 07, 2020 at 09:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `eval_empty_pkg_list`

## Description

- do not call VMaaS when package_list len is 0
- added new error states to evaluation
- added evaluation parts duration measurements

---

## Discussion

### Comment by @codecov-io on February 07, 2020 at 10:04 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=h1) Report
> Merging [#107](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/90aabf72799f66050016faaa4a7e63ecc838856a?src=pr&el=desc) will **increase** coverage by `0.63%`.
> The diff coverage is `74.6%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #107      +/-   ##
==========================================
+ Coverage   54.24%   54.88%   +0.63%     
==========================================
  Files          36       36              
  Lines        1401     1443      +42     
==========================================
+ Hits          760      792      +32     
- Misses        550      557       +7     
- Partials       91       94       +3
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `0% <0%> (ø)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.55% <75.8%> (+1.06%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=footer). Last update [90aabf7...c597344](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/107?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on February 07, 2020 at 10:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/107*
