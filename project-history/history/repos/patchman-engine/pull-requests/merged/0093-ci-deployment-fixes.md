---
type: pull_request
number: 93
title: "CI deployment fixes"
state: merged
author: semtexzv
created: 2020-01-31T11:57:56Z
updated: 2021-03-16T09:27:09Z
closed: 2020-01-31T13:45:33Z
merged: 2020-01-31T13:45:33Z
base_branch: master
head_branch: ci-deploy-fixes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/93
---

# Pull Request #93: CI deployment fixes

**Author**: @semtexzv
**Created**: January 31, 2020 at 11:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `ci-deploy-fixes`

## Description

Fixes which had to be performed in order to get large batch of PRs deployed on CI
Includes:

- Making existing migrations more resilient to schema mismatches
- Automatically deploying Database through migrations
- Fixes to openshift templates
- Fixes to evaluator component 

---

## Discussion

### Comment by @codecov-io on January 31, 2020 at 12:24 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=h1) Report
> Merging [#93](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/b7f6eec7009dde5a49bbbbf466a09a2786c8b190?src=pr&el=desc) will **decrease** coverage by `0.36%`.
> The diff coverage is `9.09%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #93      +/-   ##
==========================================
- Coverage    57.6%   57.23%   -0.37%     
==========================================
  Files          34       34              
  Lines        1401     1410       +9     
==========================================
  Hits          807      807              
- Misses        497      506       +9     
  Partials       97       97
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.76% <0%> (-0.43%)` | :arrow_down: |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `22.22% <0%> (-77.78%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `41% <33.33%> (-0.42%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=footer). Last update [b7f6eec...bf83060](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/93?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on January 31, 2020 at 01:45 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/93*
