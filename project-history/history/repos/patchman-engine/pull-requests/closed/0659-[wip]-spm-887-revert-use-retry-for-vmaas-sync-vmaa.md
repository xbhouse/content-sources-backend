---
type: pull_request
number: 659
title: "[WIP] SPM-887 Revert use retry for vmaas-sync vmaas calls"
state: closed
author: josef-hak
created: 2021-05-04T15:28:24Z
updated: 2021-08-26T18:42:22Z
closed: 2021-05-04T17:37:26Z
base_branch: master
head_branch: rev-retry
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/659
---

# Pull Request #659: [WIP] SPM-887 Revert use retry for vmaas-sync vmaas calls

**Author**: @josef-hak
**Created**: May 04, 2021 at 03:28 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `rev-retry`

## Description

- failing in evaluator evaluate.go:289
- interface conversion: interface {} is nil, not

This reverts commit 60f734e680fdd91bd44b6ad39229c10e96736a4f.


---

## Discussion

### Comment by @codecov-commenter on May 04, 2021 at 03:37 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#659](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (11229a6) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/60f734e680fdd91bd44b6ad39229c10e96736a4f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (60f734e) will **increase** coverage by `0.20%`.
> The diff coverage is `60.86%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #659      +/-   ##
==========================================
+ Coverage   58.11%   58.31%   +0.20%     
==========================================
  Files          73       72       -1     
  Lines        3347     3313      -34     
==========================================
- Hits         1945     1932      -13     
+ Misses       1121     1099      -22     
- Partials      281      282       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.31% <60.86%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `38.73% <ø> (-1.09%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `63.72% <47.05%> (-2.10%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `60.93% <100.00%> (-1.19%)` | :arrow_down: |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `72.72% <100.00%> (-0.66%)` | :arrow_down: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `71.42% <100.00%> (-2.65%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [60f734e...11229a6](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/659?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on May 04, 2021 at 05:37 PM UTC

~~~
"err":"interface conversion: interface {} is nil, not *vmaas.UpdatesV2Response","level":"error","message":"Panicked","stack":"goroutine 33 [running]:|runtime/debug.Stack(0xc002d0f4e0, 0x1174240, 0xc0021fc750)|\t/usr/lib/golang/src/runtime/debug/stack.go:24 +0x9d|app/base/utils.LogPanics(0x1)|\t/go/src/app/base/utils/core.go:120 +0x6e|panic(0x1174240, 0xc0021fc750)|\t/usr/lib/golang/src/runtime/panic.go:969 +0x166|app/evaluator.callVMaas(0x1609620, 0xc0003e3600, 0xc002ac71a0, 0x0, 0x0, 0x0)|\t/go/src/app/evaluator/evaluate.go:289 +0x33b|app/evaluator.evaluateWithVmaas(0x1609620, 0xc0003e3600, 0xc0005cab70, 0xc002ac71a0, 0xc001e424e0, 0xc002ac71a0, 0xc001e424e0, 0x0)|\t/go/src/app/evaluator/evaluate.go:114 +0x43|app/evaluator.evaluateInDatabase(0x1609620, 0xc0003e3600, 0x3, 0xc002f240f0, 0x24, 0xc0006ae860, 0x0, 0x0, 0x0, 0x0)|\t/go/src/app/evaluator/evaluate.go:105 +0x1a7|app/evaluator.Evaluate(0x1609620, 0xc0003e3600, 0x3, 0xc002f240f0, 0x24, 0xc0006ae860, 0xc00003c0ab, 0x6, 0x0, 0x0)|\t/go/src/app/evaluator/evaluate.go:75 +0x190|app/evaluator.evaluateHandler(0x0, 0x0, 0x0, 0xc0006ae860, 0x0, 0x3, 0x0, 0x0, 0xc0003e2200, 0x1, ...)|\t/go/src/app/evaluator/evaluate.go:321 +0xd8|app/base/mqueue.MakeMessageHandler.func1(0xc0005d68a0, 0x19, 0x1b, 0x133b, 0x0, 0x0, 0x0, 0xc002c06210, 0xb0, 0xb0, ...)|\t/go/src/app/base/mqueue/event.go:49 +0x204|app/base/mqueue.MakeRetryingHandler.func1(0xc0005d68a0, 0x19, 0x1b, 0x133b, 0x0, 0x0, 0x0, 0xc002c06210, 0xb0, 0xb0, ...)|\t/go/src/app/base/mqueue/mqueue.go:99 +0x2a2|app/base/mqueue.(*readerImpl).HandleMessages(0xc0005b81a0, 0xc00229c010)|\t/go/src/app/base/mqueue/mqueue.go:120 +0xc1|app/base/mqueue.runReader(0xc000502060, 0xc00003a39b, 0x19, 0x1345d48, 0xc00229c010)|\t/go/src/app/base/mqueue/mqueue.go:141 +0xbc|created by app/base/mqueue.SpawnReader|\t/go/src/app/base/mqueue/mqueue.go:146 +0x7e|"}
~~~

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/659*
