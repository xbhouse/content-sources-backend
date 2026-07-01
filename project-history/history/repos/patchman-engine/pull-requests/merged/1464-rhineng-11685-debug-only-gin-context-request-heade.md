---
type: pull_request
number: 1464
title: "RHINENG-11685: debug only gin context request header"
state: merged
author: psegedy
created: 2024-08-16T12:39:48Z
updated: 2024-08-20T09:10:40Z
closed: 2024-08-20T09:10:40Z
merged: 2024-08-20T09:10:40Z
base_branch: master
head_branch: debug_gzip
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1464
---

# Pull Request #1464: RHINENG-11685: debug only gin context request header

**Author**: @psegedy
**Created**: August 16, 2024 at 12:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `debug_gzip`

## Description

current log fails with
```
Unable to read entry, failed to marshal fields to JSON, json: unsupported type: func() (io.ReadCloser, error)
Failed to fire hook: failed to marshal fields to JSON, json: unsupported type: func() (io.ReadCloser, error)
```
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

### Comment by @jira-linking on August 16, 2024 at 12:39 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11685


### Comment by @codecov-commenter on August 16, 2024 at 06:45 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1464?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.04%. Comparing base [(`36039a2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/36039a2209a740d751a8106fe779cfe530c6d21d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`3395e30`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/3395e30e4ebcbfa9b7a0c97d5d706725b8bf5b35?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1464      +/-   ##
==========================================
+ Coverage   65.00%   65.04%   +0.03%     
==========================================
  Files         114      114              
  Lines        7810     7810              
==========================================
+ Hits         5077     5080       +3     
+ Misses       2208     2205       -3     
  Partials      525      525              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1464/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1464/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.04% <100.00%> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1464?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @Dugowitch - Approved on August 20, 2024 at 08:08 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1464*
