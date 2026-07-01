---
type: pull_request
number: 1512
title: "RHINENG-13518: assign only system which are in candlepin"
state: merged
author: MichaelMraka
created: 2024-10-15T11:12:52Z
updated: 2026-03-31T04:39:08Z
closed: 2024-12-05T12:46:44Z
merged: 2024-12-05T12:46:44Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1512
---

# Pull Request #1512: RHINENG-13518: assign only system which are in candlepin

**Author**: @MichaelMraka
**Created**: October 15, 2024 at 11:12 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr1`

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

### Comment by @jira-linking on October 15, 2024 at 11:12 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-13518


### Comment by @codecov-commenter on October 15, 2024 at 11:18 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `63.33333%` with `11 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.43%. Comparing base ([`1a55ecd`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1a55ecd35986284ee354a18542e05d12445f7e49?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`76d0fbb`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/76d0fbb35f729b45e94fa37f61f8b0cd1087b905?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1086 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/controllers/template\_systems\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX3VwZGF0ZS5nbw==) | 56.00% | [6 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1512      +/-   ##
==========================================
- Coverage   63.45%   63.43%   -0.02%     
==========================================
  Files         114      114              
  Lines        9637     9651      +14     
==========================================
+ Hits         6115     6122       +7     
- Misses       2980     2985       +5     
- Partials      542      544       +2     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.43% <63.33%> (-0.02%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1512?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @psegedy on December 02, 2024 at 02:44 PM UTC

/retest

---

## Reviews

### Review by @psegedy - Commented on October 29, 2024 at 02:48 PM UTC

### Review by @psegedy - Approved on December 04, 2024 at 01:03 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1512*
