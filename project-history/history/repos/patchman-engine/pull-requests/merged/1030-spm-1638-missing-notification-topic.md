---
type: pull_request
number: 1030
title: "SPM-1638: Missing notification topic"
state: merged
author: Mischulee
created: 2022-07-20T09:08:56Z
updated: 2022-07-20T18:30:06Z
closed: 2022-07-20T18:30:06Z
merged: 2022-07-20T18:30:06Z
base_branch: master
head_branch: up-to-date
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1030
---

# Pull Request #1030: SPM-1638: Missing notification topic

**Author**: @Mischulee
**Created**: July 20, 2022 at 09:08 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `up-to-date`

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

### Comment by @jira-linking on July 20, 2022 at 09:09 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1638


### Comment by @codecov-commenter on July 20, 2022 at 09:17 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1030](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ebfee22) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/fd466045239b9c4ed8afae2e6294685d32eee94a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fd46604) will **decrease** coverage by `0.01%`.
> The diff coverage is `n/a`.

```diff
@@            Coverage Diff             @@
##           master    #1030      +/-   ##
==========================================
- Coverage   61.25%   61.24%   -0.02%     
==========================================
  Files          95       95              
  Lines        5413     5413              
==========================================
- Hits         3316     3315       -1     
- Misses       1673     1674       +1     
  Partials      424      424              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.24% <ø> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `68.50% <0.00%> (-0.98%)` | :arrow_down: |
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `76.92% <0.00%> (+3.84%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b6ea455...ebfee22](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1030?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on July 20, 2022 at 06:28 PM UTC

thank you!

---

## Reviews

### Review by @psegedy - Commented on July 20, 2022 at 11:41 AM UTC

looks good, but I noticed that payload tracked topic is missing too 😄 would you add it as well? It's `platform.payload-status` 

### Review by @michalslomczynski - Approved on July 20, 2022 at 02:17 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1030*
