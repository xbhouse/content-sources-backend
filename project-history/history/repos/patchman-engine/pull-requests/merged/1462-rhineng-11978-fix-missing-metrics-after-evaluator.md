---
type: pull_request
number: 1462
title: "RHINENG-11978: fix missing metrics after evaluator refactor"
state: merged
author: Dugowitch
created: 2024-08-15T10:44:42Z
updated: 2024-09-26T10:10:30Z
closed: 2024-08-20T09:11:06Z
merged: 2024-08-20T09:11:06Z
base_branch: master
head_branch: RHINENG-11978
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1462
---

# Pull Request #1462: RHINENG-11978: fix missing metrics after evaluator refactor

**Author**: @Dugowitch
**Created**: August 15, 2024 at 10:44 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-11978`

## Description

## Metric Changes Overview
- introduce `advisories-load` and `advisories-evaluate` instead of the missing `advisories-processing`
- similarly to `lazy-package-save`, add `advisories-lazy-save`
- fix missing `installable`, `applicable`, and `fixed` counter updates

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

### Comment by @jira-linking on August 15, 2024 at 10:44 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-11978


### Comment by @codecov-commenter on August 15, 2024 at 10:49 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `88.23529%` with `2 lines` in your changes missing coverage. Please review.
> Project coverage is 65.06%. Comparing base [(`5adb8a6`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5adb8a64539f5cf6a16142f26a0d817334bb0131?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`ddbcbb4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ddbcbb40b21d9b945cd1ea9a1a46a3e5890de155?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [evaluator/evaluate\_advisories.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462?src=pr&el=tree&filepath=evaluator%2Fevaluate_advisories.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX2Fkdmlzb3JpZXMuZ28=) | 88.23% | [1 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1462      +/-   ##
==========================================
+ Coverage   65.01%   65.06%   +0.05%     
==========================================
  Files         114      114              
  Lines        7804     7821      +17     
==========================================
+ Hits         5074     5089      +15     
- Misses       2205     2206       +1     
- Partials      525      526       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.06% <88.23%> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1462?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Approved on August 19, 2024 at 08:53 AM UTC

looks good but let's merge it after https://github.com/RedHatInsights/patchman-engine/pull/1464

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1462*
