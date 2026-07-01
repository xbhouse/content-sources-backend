---
type: pull_request
number: 1043
title: "SPM-1682: views apis without advisories/systems"
state: merged
author: psegedy
created: 2022-08-02T11:51:15Z
updated: 2022-08-03T11:52:09Z
closed: 2022-08-03T11:52:08Z
merged: 2022-08-03T11:52:08Z
base_branch: master
head_branch: view_advisory_system
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1043
---

# Pull Request #1043: SPM-1682: views apis without advisories/systems

**Author**: @psegedy
**Created**: August 02, 2022 at 11:51 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `view_advisory_system`

## Description

NOTE: I haven't tested it since my docker-compose stopped working (probably after upgrade to qemu-7.0.0)
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

### Comment by @jira-linking on August 02, 2022 at 11:51 AM UTC

Commits missing Jira IDs:
6247ce0bd3025f90f87edc50e13a3e834e1802c4
Referenced Jiras:
https://issues.redhat.com/browse/SPM-1682


### Comment by @codecov-commenter on August 02, 2022 at 11:58 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1043](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (35f9974) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ba8d228d4a20d4402182068f745348ca5112a1fa?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ba8d228) will **decrease** coverage by `0.11%`.
> The diff coverage is `45.45%`.

```diff
@@            Coverage Diff             @@
##           master    #1043      +/-   ##
==========================================
- Coverage   61.25%   61.13%   -0.12%     
==========================================
  Files          95       95              
  Lines        5448     5463      +15     
==========================================
+ Hits         3337     3340       +3     
- Misses       1682     1694      +12     
  Partials      429      429              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.13% <45.45%> (-0.12%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/systems\_advisories\_view.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zX2Fkdmlzb3JpZXNfdmlldy5nbw==) | `56.52% <42.85%> (-17.68%)` | :arrow_down: |
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | `22.85% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ac8dbe4...35f9974](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1043?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on August 02, 2022 at 12:44 PM UTC

I assume that it didn't break anything since PR check tests passed

---

## Reviews

### Review by @MichaelMraka - Changes Requested on August 03, 2022 at 07:56 AM UTC

### Review by @psegedy - Commented on August 03, 2022 at 10:10 AM UTC

### Review by @MichaelMraka - Approved on August 03, 2022 at 10:18 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1043*
