---
type: pull_request
number: 1546
title: "RHINENG-14936: fix wrong type casting in db"
state: merged
author: MichaelMraka
created: 2024-12-17T11:08:15Z
updated: 2025-01-06T10:04:05Z
closed: 2025-01-06T10:04:04Z
merged: 2025-01-06T10:04:04Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1546
---

# Pull Request #1546: RHINENG-14936: fix wrong type casting in db

**Author**: @MichaelMraka
**Created**: December 17, 2024 at 11:08 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

## Description

inventory_id::text IN (?) which prevents use of indexes correct cast is the other way: inventory_id IN (?::uuid)

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

### Comment by @jira-linking on December 17, 2024 at 11:08 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-14936


### Comment by @codecov-commenter on December 17, 2024 at 01:18 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1546?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 57.91%. Comparing base [(`e7ddd40`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e7ddd40f3eb3d8d43cfeb4bd7d4f075e3de0ec40?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`8c4d5e4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8c4d5e4ef469970696f906897a2f0c6f7a28c9c0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

> :exclamation:  There is a different number of reports uploaded between BASE (e7ddd40) and HEAD (8c4d5e4). Click for more details.
> 
> <details><summary>HEAD has 1 upload less than BASE</summary>
>
>| Flag | BASE (e7ddd40) | HEAD (8c4d5e4) |
>|------|------|------|
>|unittests|2|1|
></details>

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1546      +/-   ##
==========================================
- Coverage   63.67%   57.91%   -5.76%     
==========================================
  Files         116      143      +27     
  Lines        9742    11124    +1382     
==========================================
+ Hits         6203     6443     +240     
- Misses       2987     4102    +1115     
- Partials      552      579      +27     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1546/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1546/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.91% <100.00%> (-5.76%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1546?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on December 18, 2024 at 09:10 AM UTC

/retest

---

## Reviews

### Review by @psegedy - Approved on January 06, 2025 at 10:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1546*
