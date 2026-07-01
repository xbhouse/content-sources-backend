---
type: pull_request
number: 242
title: "Removed unknown advisories storing during evaluation"
state: merged
author: josef-hak
created: 2020-04-28T09:42:29Z
updated: 2020-05-04T11:29:21Z
closed: 2020-04-30T12:14:35Z
merged: 2020-04-30T12:14:35Z
base_branch: master
head_branch: rm-unknown-advisories
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/242
---

# Pull Request #242: Removed unknown advisories storing during evaluation

**Author**: @josef-hak
**Created**: April 28, 2020 at 09:42 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rm-unknown-advisories`

## Description

- Unknown advisories received from vmaas are ignored during evaluation
- Added Prometheus counter to track receiving of unknown advisories

---

## Discussion

### Comment by @codecov-io on April 28, 2020 at 09:58 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=h1) Report
> Merging [#242](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/383112b641f7854aa38c0829f38f3393d0a70363&el=desc) will **increase** coverage by `0.11%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #242      +/-   ##
==========================================
+ Coverage   62.76%   62.88%   +0.11%     
==========================================
  Files          47       47              
  Lines        2103     2099       -4     
==========================================
  Hits         1320     1320              
+ Misses        629      627       -2     
+ Partials      154      152       -2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.30% <100.00%> (+1.08%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=footer). Last update [383112b...d9430ae](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/242?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on April 30, 2020 at 12:14 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/242*
