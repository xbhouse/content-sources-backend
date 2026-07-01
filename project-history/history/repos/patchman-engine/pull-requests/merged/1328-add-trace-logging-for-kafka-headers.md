---
type: pull_request
number: 1328
title: "add trace logging for kafka headers"
state: merged
author: psegedy
created: 2023-10-18T12:00:19Z
updated: 2026-04-02T20:07:11Z
closed: 2023-10-18T12:30:58Z
merged: 2023-10-18T12:30:58Z
base_branch: master
head_branch: kafka_headers_log
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1328
---

# Pull Request #1328: add trace logging for kafka headers

**Author**: @psegedy
**Created**: October 18, 2023 at 12:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kafka_headers_log`

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

### Comment by @jira-linking on October 18, 2023 at 12:00 PM UTC

Commits missing Jira IDs:
77b2b59630d1d6466d5fdb9cd4ba56b27ad5d92c


### Comment by @codecov-commenter on October 18, 2023 at 12:07 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.11%. Comparing base ([`48e75b1`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/48e75b1cc85a2a73cbcd64db9a6558400f381d27?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`77b2b59`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/77b2b59630d1d6466d5fdb9cd4ba56b27ad5d92c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1678 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/mqueue/mqueue\_impl\_gokafka.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328?src=pr&el=tree&filepath=base%2Fmqueue%2Fmqueue_impl_gokafka.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | 0.00% | [2 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1328      +/-   ##
==========================================
- Coverage   62.19%   62.11%   -0.09%     
==========================================
  Files         107      107              
  Lines        6705     6725      +20     
==========================================
+ Hits         4170     4177       +7     
- Misses       2007     2019      +12     
- Partials      528      529       +1     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.11% <0.00%> (-0.09%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1328?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on October 18, 2023 at 12:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1328*
