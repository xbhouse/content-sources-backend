---
type: pull_request
number: 97
title: "System Culling"
state: merged
author: semtexzv
created: 2020-02-03T10:20:20Z
updated: 2020-02-07T14:29:31Z
closed: 2020-02-07T14:29:31Z
merged: 2020-02-07T14:29:31Z
base_branch: master
head_branch: system-culling
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/97
---

# Pull Request #97: System Culling

**Author**: @semtexzv
**Created**: February 03, 2020 at 10:20 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `system-culling`

## Description

- Add system culling timestamps, load them from inventory
- Add periodic jobs to vmaas_sync to delete culled systems and mark stale systems

---

## Discussion

### Comment by @codecov-io on February 03, 2020 at 10:58 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=h1) Report
> Merging [#97](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ae92cd4b20384fa1590a226ee0b06ca9dec70668?src=pr&el=desc) will **increase** coverage by `0.12%`.
> The diff coverage is `51.16%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #97      +/-   ##
==========================================
+ Coverage   54.36%   54.48%   +0.12%     
==========================================
  Files          35       37       +2     
  Lines        1398     1470      +72     
==========================================
+ Hits          760      801      +41     
- Misses        547      572      +25     
- Partials       91       97       +6
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/system\_culling.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9zeXN0ZW1fY3VsbGluZy5nbw==) | `0% <0%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `0% <0%> (ø)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `64.7% <61.11%> (-3.99%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `39.04% <0%> (-0.38%)` | :arrow_down: |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `0% <0%> (ø)` | :arrow_up: |
| [base/utils/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9tZXRyaWNzLmdv) | `0% <0%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `71.55% <0%> (+1.06%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=footer). Last update [ae92cd4...642918f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/97?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on February 06, 2020 at 02:48 PM UTC

thanks, some ideas to improve it

### Review by @josef-hak - Approved on February 07, 2020 at 02:29 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/97*
