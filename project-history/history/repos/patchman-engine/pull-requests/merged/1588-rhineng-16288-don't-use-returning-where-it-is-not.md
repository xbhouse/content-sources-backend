---
type: pull_request
number: 1588
title: "RHINENG-16288: don't use RETURNING where it is not necessary"
state: merged
author: psegedy
created: 2025-02-27T16:05:27Z
updated: 2026-04-02T12:44:53Z
closed: 2025-03-03T15:13:58Z
merged: 2025-03-03T15:13:58Z
base_branch: master
head_branch: returning
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1588
---

# Pull Request #1588: RHINENG-16288: don't use RETURNING where it is not necessary

**Author**: @psegedy
**Created**: February 27, 2025 at 04:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `returning`

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

### Comment by @jira-linking on February 27, 2025 at 04:05 PM UTC

Commits missing Jira IDs:
d470541175ccd3f0589bfb418fc38202d3362fae
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16288


### Comment by @codecov-commenter on February 27, 2025 at 05:50 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `82.60870%` with `4 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 57.96%. Comparing base ([`4c40ca6`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4c40ca6fd4e6832a90c448123e84f166e7b36b98?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d470541`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d470541175ccd3f0589bfb418fc38202d3362fae?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 970 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/types/timestamp.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588?src=pr&el=tree&filepath=base%2Ftypes%2Ftimestamp.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS90eXBlcy90aW1lc3RhbXAuZ28=) | 20.00% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1588      +/-   ##
==========================================
- Coverage   57.98%   57.96%   -0.02%     
==========================================
  Files         143      143              
  Lines       11172    11184      +12     
==========================================
+ Hits         6478     6483       +5     
- Misses       4112     4118       +6     
- Partials      582      583       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.96% <82.60%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1588?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on February 27, 2025 at 04:20 PM UTC

### Review by @MichaelMraka - Approved on March 03, 2025 at 01:34 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1588*
