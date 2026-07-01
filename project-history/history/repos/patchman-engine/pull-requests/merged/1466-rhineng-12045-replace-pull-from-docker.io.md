---
type: pull_request
number: 1466
title: "RHINENG-12045: replace pull from docker.io"
state: merged
author: psegedy
created: 2024-08-19T14:35:45Z
updated: 2024-08-20T09:45:00Z
closed: 2024-08-20T09:45:00Z
merged: 2024-08-20T09:45:00Z
base_branch: master
head_branch: dockerio
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1466
---

# Pull Request #1466: RHINENG-12045: replace pull from docker.io

**Author**: @psegedy
**Created**: August 19, 2024 at 02:35 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `dockerio`

## Description

We're getting Too many requests on internal network because we cannot use docker.io account. Replace docker.io images with alternatives
- replace postgresql with the postgresql image used by clowder
- mirror cp-kafka to quay
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

### Comment by @jira-linking on August 19, 2024 at 02:35 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-12045


### Comment by @codecov-commenter on August 19, 2024 at 02:41 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1466?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.00%. Comparing base [(`36039a2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/36039a2209a740d751a8106fe779cfe530c6d21d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`f79d8de`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f79d8de8b78e6a6b769832675e632d6edd73d8d1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1466   +/-   ##
=======================================
  Coverage   65.00%   65.00%           
=======================================
  Files         114      114           
  Lines        7810     7810           
=======================================
  Hits         5077     5077           
  Misses       2208     2208           
  Partials      525      525           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1466/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1466/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1466?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on August 20, 2024 at 09:44 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1466*
