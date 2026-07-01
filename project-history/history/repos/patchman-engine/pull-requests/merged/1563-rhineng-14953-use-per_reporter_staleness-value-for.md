---
type: pull_request
number: 1563
title: "RHINENG-14953: use per_reporter_staleness value for last_upload"
state: merged
author: psegedy
created: 2025-01-29T18:11:00Z
updated: 2025-02-05T11:07:44Z
closed: 2025-02-05T11:07:44Z
merged: 2025-02-05T11:07:44Z
base_branch: master
head_branch: last_upload
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1563
---

# Pull Request #1563: RHINENG-14953: use per_reporter_staleness value for last_upload

**Author**: @psegedy
**Created**: January 29, 2025 at 06:11 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `last_upload`

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

### Comment by @jira-linking on January 29, 2025 at 06:11 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-14953


### Comment by @codecov-commenter on January 30, 2025 at 05:27 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `72.72727%` with `3 lines` in your changes missing coverage. Please review.
> Project coverage is 57.97%. Comparing base [(`ccbb6bb`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ccbb6bb4ef52bcf137388e606e277fe3ab9611d9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5258bc9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5258bc9133a69128db99879c8cbcf09a9ca605c1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 72.72% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1563   +/-   ##
=======================================
  Coverage   57.96%   57.97%           
=======================================
  Files         143      143           
  Lines       11131    11140    +9     
=======================================
+ Hits         6452     6458    +6     
- Misses       4100     4102    +2     
- Partials      579      580    +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.97% <72.72%> (+<0.01%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1563?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on February 03, 2025 at 12:43 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on January 31, 2025 at 11:57 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1563*
