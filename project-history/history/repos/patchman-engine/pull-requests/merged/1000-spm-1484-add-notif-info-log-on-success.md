---
type: pull_request
number: 1000
title: "SPM-1484: Add notif info log on success"
state: merged
author: michalslomczynski
created: 2022-06-30T09:12:53Z
updated: 2022-06-30T10:09:58Z
closed: 2022-06-30T10:09:58Z
merged: 2022-06-30T10:09:58Z
base_branch: master
head_branch: notif-log
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1000
---

# Pull Request #1000: SPM-1484: Add notif info log on success

**Author**: @michalslomczynski
**Created**: June 30, 2022 at 09:12 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `notif-log`

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

### Comment by @jira-linking on June 30, 2022 at 09:12 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1484


### Comment by @codecov-commenter on June 30, 2022 at 09:20 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1000](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (24baa92) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a7633c7d91193601cfd85372327af02fb713959b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a7633c7) will **decrease** coverage by `0.00%`.
> The diff coverage is `100.00%`.

```diff
@@            Coverage Diff             @@
##           master    #1000      +/-   ##
==========================================
- Coverage   60.98%   60.97%   -0.01%     
==========================================
  Files          95       95              
  Lines        5372     5374       +2     
==========================================
+ Hits         3276     3277       +1     
- Misses       1665     1666       +1     
  Partials      431      431              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.97% <100.00%> (-0.01%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `72.54% <100.00%> (+0.54%)` | :arrow_up: |
| [listener/metrics.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvbWV0cmljcy5nbw==) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [dac129d...24baa92](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1000?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on June 30, 2022 at 09:40 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1000*
