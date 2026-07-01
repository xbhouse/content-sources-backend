---
type: pull_request
number: 983
title: "SPM-1432: fix empty strings in schema"
state: merged
author: MichaelMraka
created: 2022-06-20T15:26:36Z
updated: 2022-06-27T09:03:21Z
closed: 2022-06-27T09:03:21Z
merged: 2022-06-27T09:03:20Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/983
---

# Pull Request #983: SPM-1432: fix empty strings in schema

**Author**: @MichaelMraka
**Created**: June 20, 2022 at 03:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

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

### Comment by @MichaelMraka on June 20, 2022 at 03:27 PM UTC

supersedes #951 

### Comment by @psegedy on June 20, 2022 at 04:48 PM UTC

/retest

### Comment by @psegedy on June 21, 2022 at 08:27 AM UTC

/retest

### Comment by @psegedy on June 21, 2022 at 10:34 AM UTC

please rebase it since it conflicts with already merged #978
```
11:00:13 evaluator/evaluate.go:178:22: invalid operation: system.VmaasJSON == "" (mismatched types *string and untyped string)
```

### Comment by @psegedy on June 22, 2022 at 03:48 PM UTC

hm I still see an error with `solution` in logs. It is in `previous` vmaas-sync container logs ~, so maybe it the error which happened before schema migration?~ It happens after vmaas-sync receives message from vmaas, roughly 10 min after database-admin finishes migrations

>```time="2022-06-22T15:18:25Z" level=fatal msg="vmaas data sync failed" err="Failed to sync advisories: Erratas page download and process failed: Storing advisories: Storing advisories: ERROR: new row for relation \"advisory_metadata\" violates check constraint \"advisory_metadata_solution_check\" (SQLSTATE 23514)"```


### Comment by @codecov-commenter on June 23, 2022 at 09:08 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#983](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (be22c0d) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/0c8da08ab6959218f49ed17b69be999c67a9ae4a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (0c8da08) will **increase** coverage by `0.11%`.
> The diff coverage is `67.17%`.

```diff
@@            Coverage Diff             @@
##           master     #983      +/-   ##
==========================================
+ Coverage   60.56%   60.68%   +0.11%     
==========================================
  Files          94       95       +1     
  Lines        5219     5293      +74     
==========================================
+ Hits         3161     3212      +51     
- Misses       1646     1656      +10     
- Partials      412      425      +13     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.68% <67.17%> (+0.11%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `50.00% <ø> (ø)` | |
| [base/utils/openapi.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9vcGVuYXBpLmdv) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/advisory\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9kZXRhaWwuZ28=) | `77.77% <ø> (ø)` | |
| [vmaas\_sync/vmaas\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy92bWFhc19zeW5jLmdv) | `43.80% <0.00%> (+0.60%)` | :arrow_up: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `65.06% <55.38%> (-2.38%)` | :arrow_down: |
| [base/utils/vmaas.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | `77.77% <77.77%> (ø)` | |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | `52.50% <100.00%> (+10.83%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `81.81% <100.00%> (+0.36%)` | :arrow_up: |
| [vmaas\_sync/advisory\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9hZHZpc29yeV9zeW5jLmdv) | `63.19% <100.00%> (ø)` | |
| [vmaas\_sync/package\_sync.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dm1hYXNfc3luYy9wYWNrYWdlX3N5bmMuZ28=) | `62.00% <100.00%> (+0.77%)` | :arrow_up: |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ea4d2c3...be22c0d](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/983?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Commented on June 21, 2022 at 10:44 AM UTC

### Review by @psegedy - Commented on June 21, 2022 at 10:51 AM UTC

### Review by @psegedy - Changes Requested on June 23, 2022 at 04:40 PM UTC

### Review by @psegedy - Commented on June 23, 2022 at 04:42 PM UTC

### Review by @psegedy - Approved on June 24, 2022 at 03:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/983*
