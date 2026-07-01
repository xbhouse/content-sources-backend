---
type: pull_request
number: 1473
title: "RHINENG-11562: rename resources env vars"
state: merged
author: psegedy
created: 2024-09-04T11:46:13Z
updated: 2024-10-08T14:04:10Z
closed: 2024-10-08T14:04:10Z
merged: 2024-10-08T14:04:10Z
base_branch: master
head_branch: resources
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1473
---

# Pull Request #1473: RHINENG-11562: rename resources env vars

**Author**: @psegedy
**Created**: September 04, 2024 at 11:46 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `resources`

## Description

set default values according to prod environment

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

### Comment by @jira-linking on September 04, 2024 at 11:46 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11562


### Comment by @codecov-commenter on September 04, 2024 at 11:51 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1473?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.03%. Comparing base [(`bbba975`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bbba97512313e51e37d43891fbc447a3b6b12bea?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`8427f98`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8427f9889933fe5a10cdef3d6eed3bff02317d13?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1473   +/-   ##
=======================================
  Coverage   65.03%   65.03%           
=======================================
  Files         114      114           
  Lines        7880     7880           
=======================================
  Hits         5125     5125           
  Misses       2216     2216           
  Partials      539      539           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1473/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1473/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.03% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1473?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on October 02, 2024 at 04:59 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on October 04, 2024 at 08:16 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1473*
