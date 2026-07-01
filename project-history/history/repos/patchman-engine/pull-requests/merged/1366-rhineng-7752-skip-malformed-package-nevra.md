---
type: pull_request
number: 1366
title: "RHINENG-7752: skip malformed package nevra"
state: merged
author: psegedy
created: 2024-01-26T13:24:37Z
updated: 2024-01-29T10:10:27Z
closed: 2024-01-29T10:10:27Z
merged: 2024-01-29T10:10:27Z
base_branch: master
head_branch: hotfixmaster
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1366
---

# Pull Request #1366: RHINENG-7752: skip malformed package nevra

**Author**: @psegedy
**Created**: January 26, 2024 at 01:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `hotfixmaster`

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

### Comment by @jira-linking on January 26, 2024 at 01:24 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-7752


### Comment by @codecov-commenter on January 26, 2024 at 01:29 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `14 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`79e2763`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/79e2763cbf3c87eb58344005ffdc7782415d54e1?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.87% compared to head [(`dae17b7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.98%.
> Report is 3 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/vmaas.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy92bWFhcy5nbw==) | 23.07% | [10 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/utils/rpm.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9ycG0uZ28=) | 50.00% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 50.00% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1366      +/-   ##
==========================================
+ Coverage   61.87%   61.98%   +0.10%     
==========================================
  Files         108      108              
  Lines        6815     6821       +6     
==========================================
+ Hits         4217     4228      +11     
+ Misses       2062     2058       -4     
+ Partials      536      535       -1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.98% <41.66%> (+0.10%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1366?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Commented on January 26, 2024 at 04:37 PM UTC

### Review by @MichaelMraka - Approved on January 26, 2024 at 05:57 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1366*
