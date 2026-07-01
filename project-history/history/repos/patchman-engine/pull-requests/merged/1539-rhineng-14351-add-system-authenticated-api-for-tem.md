---
type: pull_request
number: 1539
title: "RHINENG-14351: add system authenticated API for templates"
state: merged
author: MichaelMraka
created: 2024-12-03T14:46:14Z
updated: 2024-12-11T09:37:25Z
closed: 2024-12-11T09:37:25Z
merged: 2024-12-11T09:37:25Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1539
---

# Pull Request #1539: RHINENG-14351: add system authenticated API for templates

**Author**: @MichaelMraka
**Created**: December 03, 2024 at 02:46 PM UTC
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

### Comment by @jira-linking on December 03, 2024 at 02:46 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-14351


### Comment by @codecov-commenter on December 03, 2024 at 02:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `67.16418%` with `22 lines` in your changes missing coverage. Please review.
> Project coverage is 63.68%. Comparing base [(`a845aa7`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a845aa73282e977f7381bf58866feb507946061c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`2abe489`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2abe489cba38716ffbcfdce8fabc0baea34dc0b1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 2 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [.../controllers/template\_subscribed\_systems\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_subscribed_systems_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zdWJzY3JpYmVkX3N5c3RlbXNfdXBkYXRlLmdv) | 65.38% | [13 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/middlewares/authentication.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Fauthentication.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | 73.33% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1539      +/-   ##
==========================================
+ Coverage   63.33%   63.68%   +0.35%     
==========================================
  Files         115      116       +1     
  Lines        9673     9740      +67     
==========================================
+ Hits         6126     6203      +77     
+ Misses       3003     2986      -17     
- Partials      544      551       +7     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.68% <67.16%> (+0.35%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1539?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @psegedy - Commented on December 04, 2024 at 03:13 PM UTC

### Review by @psegedy - Approved on December 04, 2024 at 03:16 PM UTC

i think it looks good

### Review by @MichaelMraka - Commented on December 05, 2024 at 01:04 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1539*
