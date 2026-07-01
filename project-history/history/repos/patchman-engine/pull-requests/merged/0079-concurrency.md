---
type: pull_request
number: 79
title: "Concurrency"
state: merged
author: semtexzv
created: 2020-01-28T14:52:42Z
updated: 2021-03-16T09:27:04Z
closed: 2020-01-29T15:38:05Z
merged: 2020-01-29T15:38:05Z
base_branch: master
head_branch: concurrency
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/79
---

# Pull Request #79: Concurrency

**Author**: @semtexzv
**Created**: January 28, 2020 at 02:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `concurrency`

## Description

Change the concurrency profile in listener
- We'll have 8 listeners
- Each listener will have 8 consumers,
- Scheduled on 8 hardware threads


This should allow us to process 64 messages at a time, and allow us to schedule goroutines on hardware threads, so the whole pipeline won't be blocked on waiting for database.

---

## Discussion

### Comment by @codecov-io on January 29, 2020 at 12:30 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=h1) Report
> Merging [#79](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/25f0c2d3bc2ed46efc07aa24d54899a2fa57bae1?src=pr&el=desc) will **decrease** coverage by `0.07%`.
> The diff coverage is `33.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #79      +/-   ##
==========================================
- Coverage   57.49%   57.42%   -0.08%     
==========================================
  Files          34       34              
  Lines        1327     1334       +7     
==========================================
+ Hits          763      766       +3     
- Misses        474      478       +4     
  Partials       90       90
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `83.78% <100%> (ø)` | :arrow_up: |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `47.82% <27.27%> (-12.18%)` | :arrow_down: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `22.22% <0%> (ø)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `75.81% <0%> (ø)` | :arrow_up: |
| [evaluator/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL21ldHJpY3MuZ28=) | `100% <0%> (ø)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `67.7% <0%> (+0.68%)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `36.84% <0%> (+7.43%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=footer). Last update [25f0c2d...12552bc](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/79?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on January 29, 2020 at 03:37 PM UTC

ok, let's try it

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/79*
