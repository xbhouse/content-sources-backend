---
type: pull_request
number: 103
title: "Fixed duration observing in duration metrics"
state: merged
author: josef-hak
created: 2020-02-06T13:07:02Z
updated: 2020-02-06T16:00:10Z
closed: 2020-02-06T14:26:58Z
merged: 2020-02-06T14:26:58Z
base_branch: master
head_branch: fix-observer-duration
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/103
---

# Pull Request #103: Fixed duration observing in duration metrics

**Author**: @josef-hak
**Created**: February 06, 2020 at 01:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-observer-duration`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on February 06, 2020 at 01:16 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=h1) Report
> Merging [#103](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/412d8b822bfc5b9f3692e9993802831dae86a465?src=pr&el=desc) will **decrease** coverage by `0.11%`.
> The diff coverage is `37.5%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #103      +/-   ##
==========================================
- Coverage   54.36%   54.24%   -0.12%     
==========================================
  Files          35       36       +1     
  Lines        1398     1401       +3     
==========================================
  Hits          760      760              
- Misses        547      550       +3     
  Partials       91       91
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `39.04% <0%> (-0.38%)` | :arrow_down: |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `0% <0%> (ø)` | :arrow_up: |
| [base/utils/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9tZXRyaWNzLmdv) | `0% <0%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `70.49% <100%> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `36.84% <100%> (ø)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `68.68% <100%> (ø)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=footer). Last update [412d8b8...1459891](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/103?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @semtexzv - Approved on February 06, 2020 at 02:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/103*
