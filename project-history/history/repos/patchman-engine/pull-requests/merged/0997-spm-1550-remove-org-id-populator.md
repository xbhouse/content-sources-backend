---
type: pull_request
number: 997
title: "SPM-1550: remove org-id-populator"
state: merged
author: psegedy
created: 2022-06-29T14:21:00Z
updated: 2022-06-30T09:29:55Z
closed: 2022-06-30T09:29:54Z
merged: 2022-06-30T09:29:54Z
base_branch: master
head_branch: remove-org-id-populator
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/997
---

# Pull Request #997: SPM-1550: remove org-id-populator

**Author**: @psegedy
**Created**: June 29, 2022 at 02:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `remove-org-id-populator`

## Description

it is not needed anymore, we've translated all accounts in production

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

### Comment by @jira-linking on June 29, 2022 at 02:21 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1550


### Comment by @codecov-commenter on June 29, 2022 at 02:30 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/997?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#997](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/997?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7084e09) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/a7633c7d91193601cfd85372327af02fb713959b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a7633c7) will **not change** coverage.
> The diff coverage is `n/a`.

```diff
@@           Coverage Diff           @@
##           master     #997   +/-   ##
=======================================
  Coverage   60.98%   60.98%           
=======================================
  Files          95       95           
  Lines        5372     5372           
=======================================
  Hits         3276     3276           
  Misses       1665     1665           
  Partials      431      431           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.98% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/997?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/997?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [e2fb82b...7084e09](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/997?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on June 29, 2022 at 03:55 PM UTC

good catch

---

## Reviews

### Review by @michalslomczynski - Commented on June 29, 2022 at 02:37 PM UTC

Is that needed?
https://github.com/RedHatInsights/patchman-engine/blob/e2fb82b52f2591f83d865bd15d5c61ea4ba502d2/deploy/clowdapp.yaml#L573-L582

### Review by @MichaelMraka - Approved on June 30, 2022 at 08:20 AM UTC

### Review by @michalslomczynski - Approved on June 30, 2022 at 09:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/997*
