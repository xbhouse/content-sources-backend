---
type: pull_request
number: 243
title: "Add backoff specifically for vmaas returning 503"
state: merged
author: semtexzv
created: 2020-04-28T10:55:31Z
updated: 2021-03-16T09:28:16Z
closed: 2020-05-04T13:10:44Z
merged: 2020-05-04T13:10:44Z
base_branch: master
head_branch: eval-vmaas-backoff
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/243
---

# Pull Request #243: Add backoff specifically for vmaas returning 503

**Author**: @semtexzv
**Created**: April 28, 2020 at 10:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `eval-vmaas-backoff`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on May 04, 2020 at 11:54 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=h1) Report
> Merging [#243](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/4833405cbea5c78405d0840cfa7b7801ffa94750&el=desc) will **increase** coverage by `0.03%`.
> The diff coverage is `50.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #243      +/-   ##
==========================================
+ Coverage   62.88%   62.92%   +0.03%     
==========================================
  Files          47       47              
  Lines        2099     2109      +10     
==========================================
+ Hits         1320     1327       +7     
- Misses        627      629       +2     
- Partials      152      153       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.18% <50.00%> (-0.12%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=footer). Last update [4833405...6fa3e7a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/243?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on May 04, 2020 at 01:10 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/243*
