---
type: pull_request
number: 296
title: "Fix multiple installed packages handling"
state: merged
author: semtexzv
created: 2020-07-28T13:50:29Z
updated: 2021-03-16T09:26:05Z
closed: 2020-07-28T14:17:57Z
merged: 2020-07-28T14:17:57Z
base_branch: master
head_branch: pkgdata-multi-fix
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/296
---

# Pull Request #296: Fix multiple installed packages handling

**Author**: @semtexzv
**Created**: July 28, 2020 at 01:50 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkgdata-multi-fix`

## Description

- Old implementation still assumed one entry per package name.
- There was a bug in formatting NEVRAs


---

## Discussion

### Comment by @codecov-commenter on July 28, 2020 at 01:59 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=h1) Report
> Merging [#296](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/15afc83e3b70b09e9c6730f45725c898304e1fde&el=desc) will **decrease** coverage by `0.93%`.
> The diff coverage is `38.29%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #296      +/-   ##
==========================================
- Coverage   62.69%   61.76%   -0.94%     
==========================================
  Files          52       52              
  Lines        2571     2584      +13     
==========================================
- Hits         1612     1596      -16     
- Misses        757      784      +27     
- Partials      202      204       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `47.05% <20.00%> (-2.95%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.76% <46.87%> (-6.58%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=footer). Last update [15afc83...afa7eae](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/296?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on July 28, 2020 at 02:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/296*
