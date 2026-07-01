---
type: pull_request
number: 1357
title: "RHINENG-3098: deprecate limit=-1 and limit>100"
state: merged
author: psegedy
created: 2023-12-15T11:14:00Z
updated: 2026-04-03T06:02:20Z
closed: 2023-12-15T13:36:05Z
merged: 2023-12-15T13:36:05Z
base_branch: master
head_branch: limit_offset
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1357
---

# Pull Request #1357: RHINENG-3098: deprecate limit=-1 and limit>100

**Author**: @psegedy
**Created**: December 15, 2023 at 11:14 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `limit_offset`

## Description

- set Deprecation, Sunset, and Warning headers if limit=-1 or limit>100 is used
- respond with HTTP status 400 if limit=-1 or limit>100 is used after 2024-03-01
- LIMIT_PAGE_SIZE=false can be used to bypass limit of the limit parameter

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

### Comment by @jira-linking on December 15, 2023 at 11:14 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-3098


### Comment by @codecov-commenter on December 15, 2023 at 12:37 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `25.00000%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.13%. Comparing base ([`b713b5d`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b713b5d3fe502b2048738defe3980f9f4e3ca4da?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ad10c46`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ad10c46b49ae4295794911937a708860d1ebfa41?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1581 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/gin.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357?src=pr&el=tree&filepath=base%2Futils%2Fgin.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9naW4uZ28=) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1357      +/-   ##
==========================================
+ Coverage   61.98%   62.13%   +0.14%     
==========================================
  Files         108      108              
  Lines        6811     6801      -10     
==========================================
+ Hits         4222     4226       +4     
+ Misses       2055     2041      -14     
  Partials      534      534              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.13% <25.00%> (+0.14%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1357?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @psegedy on December 15, 2023 at 01:09 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on December 15, 2023 at 01:13 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1357*
