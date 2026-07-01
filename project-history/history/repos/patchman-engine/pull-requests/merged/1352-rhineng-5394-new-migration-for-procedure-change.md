---
type: pull_request
number: 1352
title: "RHINENG-5394: new migration for procedure change"
state: merged
author: psegedy
created: 2023-12-07T10:48:01Z
updated: 2023-12-07T10:59:10Z
closed: 2023-12-07T10:59:10Z
merged: 2023-12-07T10:59:10Z
base_branch: master
head_branch: migration3
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1352
---

# Pull Request #1352: RHINENG-5394: new migration for procedure change

**Author**: @psegedy
**Created**: December 07, 2023 at 10:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `migration3`

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

### Comment by @jira-linking on December 07, 2023 at 10:48 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-5394


### Comment by @codecov-commenter on December 07, 2023 at 10:54 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `20 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`25f765e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/25f765ec1438368c4e3d9d1d2d9094362619b9ac?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.82% compared to head [(`4fff9c0`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.03%.
> Report is 34 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 85.13% | [6 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/system\_packages\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXNfZXhwb3J0Lmdv) | 66.66% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/utils.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1352      +/-   ##
==========================================
+ Coverage   61.82%   62.03%   +0.20%     
==========================================
  Files         108      108              
  Lines        6814     6811       -3     
==========================================
+ Hits         4213     4225      +12     
+ Misses       2062     2052      -10     
+ Partials      539      534       -5     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.03% <81.98%> (+0.20%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1352?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on December 07, 2023 at 10:49 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1352*
