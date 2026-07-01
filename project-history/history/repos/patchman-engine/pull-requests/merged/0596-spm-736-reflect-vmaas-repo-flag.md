---
type: pull_request
number: 596
title: "SPM-736: Reflect Vmaas repo flag"
state: merged
author: MichaelMraka
created: 2021-04-07T16:30:26Z
updated: 2021-04-08T16:22:59Z
closed: 2021-04-08T16:22:58Z
merged: 2021-04-08T16:22:58Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/596
---

# Pull Request #596: SPM-736: Reflect Vmaas repo flag

**Author**: @MichaelMraka
**Created**: April 07, 2021 at 04:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

+ updated vmaas client code

---

## Discussion

### Comment by @codecov-io on April 08, 2021 at 08:17 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=h1) Report
> Merging [#596](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=desc) (ccdad39) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/370339dd49f2f296e6da225c3c8a55ba8bd8c2ea?el=desc) (370339d) will **decrease** coverage by `0.03%`.
> The diff coverage is `84.72%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #596      +/-   ##
==========================================
- Coverage   59.05%   59.01%   -0.04%     
==========================================
  Files          70       70              
  Lines        3158     3155       -3     
==========================================
- Hits         1865     1862       -3     
- Misses       1023     1024       +1     
+ Partials      270      269       -1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.01% <84.72%> (-0.04%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-YmFzZS91dGlscy90ZXN0aW5nLmdv) | `59.09% <0.00%> (-9.34%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `62.81% <50.00%> (-0.38%)` | :arrow_down: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `63.86% <85.18%> (+2.19%)` | :arrow_up: |
| [vmaas\_sync/repo\_based.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | `68.25% <88.88%> (-0.50%)` | :arrow_down: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `77.66% <90.90%> (ø)` | |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `63.81% <100.00%> (ø)` | |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `71.42% <100.00%> (ø)` | |
| [vmaas\_sync/pkg\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9wa2dfc3luYy5nbw==) | `76.92% <100.00%> (-0.50%)` | :arrow_down: |
| [vmaas\_sync/repo\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy9yZXBvX3N5bmMuZ28=) | `33.33% <100.00%> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596/diff?src=pr&el=tree#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `34.40% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=footer). Last update [370339d...ccdad39](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/596?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Approved on April 08, 2021 at 08:58 AM UTC

Looks pretty good, just please check two things I mentioned.

### Review by @MichaelMraka - Commented on April 08, 2021 at 10:39 AM UTC

### Review by @MichaelMraka - Commented on April 08, 2021 at 11:07 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/596*
