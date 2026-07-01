---
type: pull_request
number: 1638
title: "RHINENG-17345: assign template according to system's environment in candlepin"
state: merged
author: MichaelMraka
created: 2025-04-17T13:30:09Z
updated: 2026-04-01T04:37:32Z
closed: 2025-04-25T08:10:14Z
merged: 2025-04-25T08:10:13Z
base_branch: master
head_branch: pr1
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1638
---

# Pull Request #1638: RHINENG-17345: assign template according to system's environment in candlepin

**Author**: @MichaelMraka
**Created**: April 17, 2025 at 01:30 PM UTC
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

### Comment by @jira-linking on April 17, 2025 at 01:30 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-17345


### Comment by @codecov-commenter on April 17, 2025 at 01:35 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `50.00000%` with `57 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 58.21%. Comparing base ([`a49be49`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a49be497cce42d81fd24b0b1b92e7c45bc961fdb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`5cf05d5`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5cf05d53c5fd1ff62179f7375dacb1fd106f3ffc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 872 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/candlepin/candlepin.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&filepath=base%2Fcandlepin%2Fcandlepin.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jYW5kbGVwaW4vY2FuZGxlcGluLmdv) | 37.83% | [21 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [platform/candlepin.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&filepath=platform%2Fcandlepin.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-cGxhdGZvcm0vY2FuZGxlcGluLmdv) | 0.00% | [15 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/upload.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&filepath=listener%2Fupload.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | 71.05% | [7 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [listener/rhsm.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&filepath=listener%2Frhsm.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvcmhzbS5nbw==) | 57.89% | [5 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1638      +/-   ##
==========================================
+ Coverage   58.16%   58.21%   +0.05%     
==========================================
  Files         146      146              
  Lines       11333    11397      +64     
==========================================
+ Hits         6592     6635      +43     
- Misses       4159     4175      +16     
- Partials      582      587       +5     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.21% <50.00%> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1638?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @jlsherrill - Commented on April 22, 2025 at 07:02 PM UTC

### Review by @MichaelMraka - Commented on April 23, 2025 at 08:47 AM UTC

### Review by @psegedy - Commented on April 23, 2025 at 02:38 PM UTC

### Review by @MichaelMraka - Commented on April 24, 2025 at 09:18 AM UTC

### Review by @psegedy - Approved on April 24, 2025 at 08:00 PM UTC

### Review by @jlsherrill - Approved on April 24, 2025 at 08:12 PM UTC

haven't tested this, but it looks correct to me.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1638*
