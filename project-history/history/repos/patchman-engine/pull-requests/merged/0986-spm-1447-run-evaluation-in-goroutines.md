---
type: pull_request
number: 986
title: "SPM-1447: run evaluation in goroutines"
state: merged
author: psegedy
created: 2022-06-22T15:02:45Z
updated: 2022-06-28T13:37:58Z
closed: 2022-06-28T13:37:58Z
merged: 2022-06-28T13:37:58Z
base_branch: master
head_branch: evaluate_concurrent
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/986
---

# Pull Request #986: SPM-1447: run evaluation in goroutines

**Author**: @psegedy
**Created**: June 22, 2022 at 03:02 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `evaluate_concurrent`

## Description

It looks like it's helping, but just a bit. Recalc time went down from ~1h to ~40min (once I got 11 min, so I'm not sure...) when I used 50 goroutines in ephemeral env. IMO we need to see evals/sec metrics in some other environments.

The PR runs `Evaluate` function in Goroutines limited by MAX_EVAL_GOROUTINES env var.

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @codecov-commenter on June 23, 2022 at 04:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#986](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (96990d4) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0c8da08ab6959218f49ed17b69be999c67a9ae4a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0c8da08) will **increase** coverage by `0.34%`.
> The diff coverage is `66.88%`.

```diff
@@            Coverage Diff             @@
##           master     #986      +/-   ##
==========================================
+ Coverage   60.56%   60.90%   +0.34%     
==========================================
  Files          94       95       +1     
  Lines        5219     5303      +84     
==========================================
+ Hits         3161     3230      +69     
- Misses       1646     1649       +3     
- Partials      412      424      +12     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.90% <66.88%> (+0.34%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50.00% <ø> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `43.80% <0.00%> (+0.60%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.60% <59.79%> (+1.16%)` | :arrow_up: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.77% <77.77%> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.50% <100.00%> (+10.83%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `81.74% <100.00%> (+0.29%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [502a9ae...96990d4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/986?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on June 24, 2022 at 08:20 AM UTC

I am curious results of benchmark in <CPU_CORES, 1000> range of goroutines.

### Comment by @psegedy on June 24, 2022 at 09:34 AM UTC

well, I tried that only for recalc and it's not really easy to benchmark. You need to generate a lot of data, you need data synced in vmaas, then you need to trigger recalc which sometimes runs 11min and other times a lot longer... In the end its still on generated data and it does not really be close to reality. Maybe we could benchmark it with uploads, but again we would need thousands of systems with different packages, repos, and uploaded to different accounts to really hit the eval in batches (1 kafka message can contain multiple systems). Also, we would need an environment with fully synced vmaas for the proper benchmarking.

After all, we can't run very high number of goroutines, we are limited by memory and vmaas cannot send too many requests to vmaas.

The main issue these goroutines should solve is when you get a kafka message with a lot of systems (default max is 4000). It looks like we are seeing this scenario in prod with recalc messages. So, instead of processing these messages from 1 kafka message sequentially, we can try to evaluate them concurrently (but still we cannot halt vmaas) which should help.

### Comment by @psegedy on June 28, 2022 at 01:37 PM UTC

🙏 

---

## Reviews

### Review by @MichaelMraka - Approved on June 28, 2022 at 11:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/986*
