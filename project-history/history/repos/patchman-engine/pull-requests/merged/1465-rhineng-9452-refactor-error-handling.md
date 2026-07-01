---
type: pull_request
number: 1465
title: "RHINENG-9452: refactor error handling"
state: merged
author: psegedy
created: 2024-08-16T15:53:42Z
updated: 2024-08-26T07:48:25Z
closed: 2024-08-26T07:48:25Z
merged: 2024-08-26T07:48:25Z
base_branch: master
head_branch: errors
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1465
---

# Pull Request #1465: RHINENG-9452: refactor error handling

**Author**: @psegedy
**Created**: August 16, 2024 at 03:53 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `errors`

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

### Comment by @jira-linking on August 16, 2024 at 03:53 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-9452


### Comment by @codecov-commenter on August 19, 2024 at 02:50 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `60.56338%` with `28 lines` in your changes missing coverage. Please review.
> Project coverage is 64.98%. Comparing base [(`36039a2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/36039a2209a740d751a8106fe779cfe530c6d21d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`bbb08c0`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bbb08c0694c3aab0322f679ebcb964b5f661f088?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 67.79% | [14 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/mqueue/mqueue.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&filepath=base%2Fmqueue%2Fmqueue.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | 25.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&filepath=evaluator%2Fevaluate.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/evaluate\_baseline.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&filepath=evaluator%2Fevaluate_baseline.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Jhc2VsaW5lLmdv) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [evaluator/remediations.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&filepath=evaluator%2Fremediations.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | 66.66% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/authentication.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Fauthentication.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1465      +/-   ##
==========================================
- Coverage   65.00%   64.98%   -0.02%     
==========================================
  Files         114      114              
  Lines        7810     7812       +2     
==========================================
  Hits         5077     5077              
+ Misses       2208     2206       -2     
- Partials      525      529       +4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `64.98% <60.56%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1465?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @niyazRedhat on August 20, 2024 at 09:52 AM UTC

/retest

### Comment by @niyazRedhat on August 21, 2024 at 11:11 AM UTC

/retest

---

## Reviews

### Review by @Dugowitch - Approved on August 20, 2024 at 09:27 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1465*
