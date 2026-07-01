---
type: pull_request
number: 708
title: "SPM-806: Add \"reboot_required\" info to advisory"
state: merged
author: josef-hak
created: 2021-06-18T13:33:00Z
updated: 2021-09-20T13:43:51Z
closed: 2021-09-09T20:16:16Z
merged: 2021-09-09T20:16:16Z
base_branch: master
head_branch: reboot-req-db
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/708
---

# Pull Request #708: SPM-806: Add "reboot_required" info to advisory

**Author**: @josef-hak
**Created**: June 18, 2021 at 01:33 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `reboot-req-db`

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

### Comment by @codecov-commenter on June 18, 2021 at 01:39 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#708](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7ae3d3f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/d9d41025917d58131d60a313df5eabbfa027a0ae?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d9d4102) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #708   +/-   ##
=======================================
  Coverage   56.45%   56.45%           
=======================================
  Files          77       77           
  Lines        3592     3592           
=======================================
  Hits         2028     2028           
  Misses       1267     1267           
  Partials      297      297           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/system\_advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllcy5nbw==) | `70.27% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [d9d4102...7ae3d3f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/708?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on June 21, 2021 at 09:27 AM UTC

/retest


---

## Reviews

### Review by @MichaelMraka - Commented on June 21, 2021 at 09:32 AM UTC

### Review by @josef-hak - Commented on June 21, 2021 at 01:00 PM UTC

### Review by @MichaelMraka - Approved on June 21, 2021 at 01:06 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/708*
