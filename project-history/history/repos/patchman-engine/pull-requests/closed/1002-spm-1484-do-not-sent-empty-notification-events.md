---
type: pull_request
number: 1002
title: "SPM-1484: Do not sent empty notification events"
state: closed
author: michalslomczynski
created: 2022-06-30T11:25:21Z
updated: 2022-07-01T08:09:42Z
closed: 2022-06-30T11:38:32Z
base_branch: master
head_branch: adv-notif
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1002
---

# Pull Request #1002: SPM-1484: Do not sent empty notification events

**Author**: @michalslomczynski
**Created**: June 30, 2022 at 11:25 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `adv-notif`

## Description

It prevents from sending notification with empty events slice.
I wrote check like this with the first PR but listened blindly to golint and removed it by mistake :disappointed: 

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

### Comment by @jira-linking on June 30, 2022 at 11:25 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/SPM-1484


### Comment by @codecov-commenter on June 30, 2022 at 11:32 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#1002](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d3b2924) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/bf619f46f8c3ad992a623e23e98cb50c81230a93?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bf619f4) will **not change** coverage.
> The diff coverage is `0.00%`.

```diff
@@           Coverage Diff           @@
##           master    #1002   +/-   ##
=======================================
  Coverage   60.97%   60.97%           
=======================================
  Files          95       95           
  Lines        5374     5374           
=======================================
  Hits         3277     3277           
  Misses       1666     1666           
  Partials      431      431           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `60.97% <0.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/notifications.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL25vdGlmaWNhdGlvbnMuZ28=) | `72.54% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [511f291...d3b2924](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1002?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on June 30, 2022 at 11:38 AM UTC

I am looking for a reason why we see Patch events without email action from time to time but its not it :slightly_frowning_face: 
https://console.stage.redhat.com/settings/notifications/eventlog?application=rhel.patch
```
New advisory | Patch - Red Hat Enterprise Linux | No actions | 11 minutes ago
``` 
Closing.



---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1002*
