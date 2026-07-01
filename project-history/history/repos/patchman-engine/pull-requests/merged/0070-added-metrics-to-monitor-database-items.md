---
type: pull_request
number: 70
title: "Added metrics to monitor database items"
state: merged
author: josef-hak
created: 2020-01-22T10:16:25Z
updated: 2020-01-30T17:25:03Z
closed: 2020-01-27T10:40:11Z
merged: 2020-01-27T10:40:11Z
base_branch: master
head_branch: manager_metrics
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/70
---

# Pull Request #70: Added metrics to monitor database items

**Author**: @josef-hak
**Created**: January 22, 2020 at 10:16 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `manager_metrics`

## Description

- systems stats
- advisories stats
- system advisories stats


---

## Discussion

### Comment by @codecov-io on January 22, 2020 at 11:44 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=h1) Report
> Merging [#70](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a5eb8735cd8049ad393d80f15cccae2570318487?src=pr&el=desc) will **decrease** coverage by `1.7%`.
> The diff coverage is `27.63%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #70      +/-   ##
==========================================
- Coverage   57.57%   55.86%   -1.71%     
==========================================
  Files          29       30       +1     
  Lines        1221     1287      +66     
==========================================
+ Hits          703      719      +16     
- Misses        436      478      +42     
- Partials       82       90       +8
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `0% <0%> (-3.39%)` | :arrow_down: |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28% <28%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `48.33% <0%> (-5%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=footer). Last update [a5eb873...0c997eb](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/70?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on January 24, 2020 at 09:34 AM UTC

I don't thinks these metrics should be in manager. I'd move these VMaaS sync & focus on reporting metrics only about endpoints here.

### Comment by @josef-hak on January 24, 2020 at 02:21 PM UTC

@semtexzv ok, I moved new metrics to be served from vmaas_sync. We already have monitoring of manager endpoints.

---

## Reviews

### Review by @semtexzv - Approved on January 27, 2020 at 10:40 AM UTC

No issues here

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/70*
