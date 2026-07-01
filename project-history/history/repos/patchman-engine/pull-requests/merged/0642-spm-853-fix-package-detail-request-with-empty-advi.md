---
type: pull_request
number: 642
title: "SPM-853 Fix package detail request with empty advisory_id"
state: merged
author: josef-hak
created: 2021-04-28T08:02:34Z
updated: 2021-05-03T12:54:58Z
closed: 2021-04-28T09:12:42Z
merged: 2021-04-28T09:12:42Z
base_branch: master
head_branch: pkg-adv
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/642
---

# Pull Request #642: SPM-853 Fix package detail request with empty advisory_id

**Author**: @josef-hak
**Created**: April 28, 2021 at 08:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pkg-adv`

## Description

We don't store advisory_id in package info in new approach - syncing from vmaas/pkgtree. Anyway this info (package - advisory relation) was not shown anywhere.

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

### Comment by @codecov-commenter on April 28, 2021 at 08:30 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#642](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (856378a) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/88af2c869fa15705d01147df9d891a496c48857d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (88af2c8) will **not change** coverage.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #642   +/-   ##
=======================================
  Coverage   58.90%   58.90%           
=======================================
  Files          72       72           
  Lines        3283     3283           
=======================================
  Hits         1934     1934           
  Misses       1067     1067           
  Partials      282      282           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.90% <0.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [88af2c8...856378a](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/642?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @semtexzv - Approved on April 28, 2021 at 08:44 AM UTC

### Review by @MichaelMraka - Approved on April 28, 2021 at 08:48 AM UTC

### Review by @josef-hak - Commented on April 28, 2021 at 09:20 AM UTC

### Review by @MichaelMraka - Commented on April 28, 2021 at 09:25 AM UTC

### Review by @josef-hak - Commented on April 28, 2021 at 10:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/642*
