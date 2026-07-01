---
type: pull_request
number: 1369
title: "RHINENG-7684: create new schema and APIs for templates"
state: merged
author: MichaelMraka
created: 2024-02-13T13:05:28Z
updated: 2026-04-02T00:00:26Z
closed: 2024-03-20T13:33:19Z
merged: 2024-03-20T13:33:19Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1369
---

# Pull Request #1369: RHINENG-7684: create new schema and APIs for templates

**Author**: @MichaelMraka
**Created**: February 13, 2024 at 01:05 PM UTC
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

### Comment by @jira-linking on February 13, 2024 at 01:05 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-7684


### Comment by @codecov-commenter on March 05, 2024 at 03:01 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `60.47431%` with `100 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.59%. Comparing base ([`17ca27e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/17ca27efdaf680406800c2587413cf90ae69ce17?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`900d8d5`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/900d8d5be58135e785fc4e763e035e17ff2fb634?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1474 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/database/testing.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&filepath=base%2Fdatabase%2Ftesting.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | 0.00% | [34 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zLmdv) | 66.66% | [20 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems\_update.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_update.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX3VwZGF0ZS5nbw==) | 58.62% | [16 Missing and 8 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/templates.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplates.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZXMuZ28=) | 83.60% | [6 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems\_export.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_export.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX2V4cG9ydC5nbw==) | 64.28% | [3 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/controllers/template\_systems\_delete.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&filepath=manager%2Fcontrollers%2Ftemplate_systems_delete.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy90ZW1wbGF0ZV9zeXN0ZW1zX2RlbGV0ZS5nbw==) | 72.72% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1369      +/-   ##
==========================================
- Coverage   63.68%   63.59%   -0.10%     
==========================================
  Files         107      112       +5     
  Lines        6521     6774     +253     
==========================================
+ Hits         4153     4308     +155     
- Misses       1879     1958      +79     
- Partials      489      508      +19     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.59% <60.47%> (-0.10%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1369?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @psegedy - Commented on February 29, 2024 at 04:54 PM UTC

### Review by @MichaelMraka - Commented on March 05, 2024 at 03:04 PM UTC

### Review by @psegedy - Commented on March 06, 2024 at 02:16 PM UTC

### Review by @MichaelMraka - Commented on March 06, 2024 at 04:04 PM UTC

### Review by @MichaelMraka - Commented on March 06, 2024 at 04:52 PM UTC

### Review by @MichaelMraka - Commented on March 07, 2024 at 09:42 AM UTC

### Review by @MichaelMraka - Commented on March 07, 2024 at 09:43 AM UTC

### Review by @MichaelMraka - Commented on March 07, 2024 at 09:43 AM UTC

### Review by @psegedy - Approved on March 19, 2024 at 11:45 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1369*
