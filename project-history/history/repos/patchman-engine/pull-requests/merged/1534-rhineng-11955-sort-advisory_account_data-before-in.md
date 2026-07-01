---
type: pull_request
number: 1534
title: "RHINENG-11955: sort advisory_account_data before insert"
state: merged
author: psegedy
created: 2024-11-19T14:19:14Z
updated: 2024-11-20T12:23:24Z
closed: 2024-11-20T12:22:49Z
merged: 2024-11-20T12:22:49Z
base_branch: master
head_branch: aad_insert_deadlock
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1534
---

# Pull Request #1534: RHINENG-11955: sort advisory_account_data before insert

**Author**: @psegedy
**Created**: November 19, 2024 at 02:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `aad_insert_deadlock`

## Description

fix deadlocks when inserting to db in when processing multiple systems

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

### Comment by @jira-linking on November 19, 2024 at 02:19 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11955


### Comment by @codecov-commenter on November 19, 2024 at 02:34 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `40.00000%` with `3 lines` in your changes missing coverage. Please review.
> Project coverage is 63.45%. Comparing base [(`5a1ff8f`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5a1ff8fd121bba83bbce29633f84da6d9e6cd5a2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`ab5bea0`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ab5bea0823c6efb0798f295ef032740c49a2588f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534?src=pr&el=tree&filepath=evaluator%2Fevaluate_advisories.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | 40.00% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1534      +/-   ##
==========================================
- Coverage   63.46%   63.45%   -0.02%     
==========================================
  Files         114      114              
  Lines        9632     9637       +5     
==========================================
+ Hits         6113     6115       +2     
- Misses       2978     2980       +2     
- Partials      541      542       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.45% <40.00%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1534?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

----
🚨 Try these New Features:

- [Flaky Tests Detection](https://docs.codecov.com/docs/test-result-ingestion-beta) - Detect and resolve failed and flaky tests

---

## Reviews

### Review by @MichaelMraka - Approved on November 19, 2024 at 02:27 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1534*
