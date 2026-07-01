---
type: pull_request
number: 1614
title: "RHINENG-14719: set system template based on rhsm environments"
state: merged
author: psegedy
created: 2025-03-26T15:09:46Z
updated: 2025-03-28T09:44:14Z
closed: 2025-03-28T09:44:14Z
merged: 2025-03-28T09:44:14Z
base_branch: master
head_branch: candlepinenv
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1614
---

# Pull Request #1614: RHINENG-14719: set system template based on rhsm environments

**Author**: @psegedy
**Created**: March 26, 2025 at 03:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `candlepinenv`

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

### Comment by @jira-linking on March 26, 2025 at 03:09 PM UTC

Commits missing Jira IDs:
31dfd2f5f72206645f3872debdb034d770556d98
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-14719


### Comment by @codecov-commenter on March 26, 2025 at 03:21 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `79.41176%` with `7 lines` in your changes missing coverage. Please review.
> Project coverage is 58.15%. Comparing base [(`bec69c6`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bec69c6691b776dec3fc234a9181ff9d42383ac4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`801d16a`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/801d16aab5bb5fe97df1025ab166e0da83ed45b2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [listener/rhsm.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?src=pr&el=tree&filepath=listener%2Frhsm.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvcmhzbS5nbw==) | 88.46% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 57.14% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?src=pr&el=tree&filepath=base%2Fdatabase%2Ftesting.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1614      +/-   ##
==========================================
+ Coverage   58.08%   58.15%   +0.06%     
==========================================
  Files         145      146       +1     
  Lines       11271    11304      +33     
==========================================
+ Hits         6547     6574      +27     
- Misses       4139     4143       +4     
- Partials      585      587       +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.15% <79.41%> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1614?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Commented on March 27, 2025 at 08:50 AM UTC

### Review by @MichaelMraka - Commented on March 27, 2025 at 04:28 PM UTC

### Review by @MichaelMraka - Approved on March 28, 2025 at 08:55 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1614*
