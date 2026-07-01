---
type: pull_request
number: 809
title: "SPM-1146: Update vmaas address reading in Clowder"
state: merged
author: josef-hak
created: 2021-09-02T11:28:21Z
updated: 2021-09-02T12:02:58Z
closed: 2021-09-02T12:02:58Z
merged: 2021-09-02T12:02:58Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/809
---

# Pull Request #809: SPM-1146: Update vmaas address reading in Clowder

**Author**: @josef-hak
**Created**: September 02, 2021 at 11:28 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

vmaas "service" fields values changed ("webapp" > "webapp-service") etc.

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

### Comment by @codecov-commenter on September 02, 2021 at 11:43 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#809](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7b50445) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/dd1d6f2b2b89ed2def3a4cf41e1686bd1dc97f59?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (dd1d6f2) will **not change** coverage.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #809   +/-   ##
=======================================
  Coverage   57.68%   57.68%           
=======================================
  Files          81       81           
  Lines        3890     3890           
=======================================
  Hits         2244     2244           
  Misses       1327     1327           
  Partials      319      319           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.68% <0.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/utils/config.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9jb25maWcuZ28=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [dd1d6f2...7b50445](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/809?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on September 02, 2021 at 11:49 AM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/809*
