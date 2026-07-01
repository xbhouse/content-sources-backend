---
type: pull_request
number: 94
title: "Enable to call vmaas_sync - sync manually"
state: merged
author: josef-hak
created: 2020-01-31T15:25:41Z
updated: 2020-02-06T16:00:11Z
closed: 2020-02-04T08:44:26Z
merged: 2020-02-04T08:44:26Z
base_branch: master
head_branch: sync_manual
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/94
---

# Pull Request #94: Enable to call vmaas_sync - sync manually

**Author**: @josef-hak
**Created**: January 31, 2020 at 03:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sync_manual`

## Description

- also added syncing duration Prometheus metric

---

## Discussion

### Comment by @codecov-io on January 31, 2020 at 03:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=h1) Report
> Merging [#94](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ec8b9b010b5dd1bc6348236b99303668795bcb28?src=pr&el=desc) will **decrease** coverage by `0.72%`.
> The diff coverage is `5.26%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #94      +/-   ##
==========================================
- Coverage   57.23%   56.51%   -0.73%     
==========================================
  Files          34       35       +1     
  Lines        1410     1428      +18     
==========================================
  Hits          807      807              
- Misses        506      524      +18     
  Partials       97       97
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `83.78% <ø> (ø)` | :arrow_up: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `40.19% <0%> (-0.81%)` | :arrow_down: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `0% <0%> (ø)` | :arrow_up: |
| [vmaas\_sync/debug\_api.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9kZWJ1Z19hcGkuZ28=) | `0% <0%> (ø)` | |
| [vmaas\_sync/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9tZXRyaWNzLmdv) | `28% <100%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=footer). Last update [ec8b9b0...b116069](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/94?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on February 04, 2020 at 08:42 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/94*
