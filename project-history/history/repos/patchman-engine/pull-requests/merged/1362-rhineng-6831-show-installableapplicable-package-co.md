---
type: pull_request
number: 1362
title: "RHINENG-6831: show installable/applicable package counts in API"
state: merged
author: MichaelMraka
created: 2024-01-16T09:56:07Z
updated: 2024-01-16T15:31:44Z
closed: 2024-01-16T15:31:44Z
merged: 2024-01-16T15:31:44Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1362
---

# Pull Request #1362: RHINENG-6831: show installable/applicable package counts in API

**Author**: @MichaelMraka
**Created**: January 16, 2024 at 09:56 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

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

### Comment by @jira-linking on January 16, 2024 at 09:56 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-6831


### Comment by @codecov-commenter on January 16, 2024 at 10:01 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `36 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`77fa1d6`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/77fa1d63baf39adda8f97b4045e1dc817cfe7a91?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.13% compared to head [(`a5449ab`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.94%.
> Report is 4 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_packages.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | 75.47% | [23 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [10 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1362      +/-   ##
==========================================
- Coverage   62.13%   61.94%   -0.20%     
==========================================
  Files         108      108              
  Lines        6811     6813       +2     
==========================================
- Hits         4232     4220      -12     
- Misses       2043     2058      +15     
+ Partials      536      535       -1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.94% <70.00%> (-0.20%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1362?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on January 16, 2024 at 03:31 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1362*
