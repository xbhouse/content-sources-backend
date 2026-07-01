---
type: pull_request
number: 1626
title: "RHINENG-16397: fix nil pointer dereference when candlepin call fails"
state: merged
author: psegedy
created: 2025-04-03T16:13:02Z
updated: 2026-04-06T04:24:52Z
closed: 2025-04-04T09:59:46Z
merged: 2025-04-04T09:59:46Z
base_branch: master
head_branch: templatebug
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1626
---

# Pull Request #1626: RHINENG-16397: fix nil pointer dereference when candlepin call fails

**Author**: @psegedy
**Created**: April 03, 2025 at 04:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `templatebug`

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

### Comment by @jira-linking on April 03, 2025 at 04:13 PM UTC

Commits missing Jira IDs:
9b71acf4acf6fb8b0ee5560a70fec02a9b77d781
Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-16397


### Comment by @codecov-commenter on April 03, 2025 at 04:43 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `55.55556%` with `8 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 58.27%. Comparing base ([`d90de5c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d90de5c730ea2b0e4721ef52cb7d0b0298e3555c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`9b71acf`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9b71acf4acf6fb8b0ee5560a70fec02a9b77d781?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 906 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [platform/candlepin.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626?src=pr&el=tree&filepath=platform%2Fcandlepin.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-cGxhdGZvcm0vY2FuZGxlcGluLmdv) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1626      +/-   ##
==========================================
+ Coverage   58.15%   58.27%   +0.11%     
==========================================
  Files         146      146              
  Lines       11304    11321      +17     
==========================================
+ Hits         6574     6597      +23     
+ Misses       4143     4142       -1     
+ Partials      587      582       -5     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.27% <55.55%> (+0.11%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1626?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on April 04, 2025 at 07:34 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1626*
