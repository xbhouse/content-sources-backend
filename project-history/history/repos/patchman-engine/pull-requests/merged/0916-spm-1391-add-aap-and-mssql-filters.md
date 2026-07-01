---
type: pull_request
number: 916
title: "SPM-1391: Add AAP and MSSQL filters"
state: merged
author: michalslomczynski
created: 2022-03-07T08:26:17Z
updated: 2022-03-09T14:40:27Z
closed: 2022-03-07T15:17:42Z
merged: 2022-03-07T15:17:42Z
base_branch: master
head_branch: aap-mysql-filters
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/916
---

# Pull Request #916: SPM-1391: Add AAP and MSSQL filters

**Author**: @michalslomczynski
**Created**: March 07, 2022 at 08:26 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `aap-mysql-filters`

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

### Comment by @codecov-commenter on March 07, 2022 at 12:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#916](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8c42ded) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/824a0d57ae31ad0146efb9f6fc1cdcb806a84427?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (824a0d5) will **increase** coverage by `0.18%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #916      +/-   ##
==========================================
+ Coverage   59.35%   59.54%   +0.18%     
==========================================
  Files          90       90              
  Lines        4709     4731      +22     
==========================================
+ Hits         2795     2817      +22     
  Misses       1528     1528              
  Partials      386      386              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `59.54% <100.00%> (+0.18%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `84.59% <100.00%> (+1.19%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9cfb10a...8c42ded](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/916?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @michalslomczynski on March 07, 2022 at 01:52 PM UTC

Go-critic was complaining on function complexities thus style changes and function split.

---

## Reviews

### Review by @josef-hak - Approved on March 07, 2022 at 03:13 PM UTC

/lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/916*
