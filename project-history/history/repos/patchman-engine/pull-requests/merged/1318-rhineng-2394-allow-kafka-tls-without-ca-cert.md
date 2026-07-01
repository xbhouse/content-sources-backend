---
type: pull_request
number: 1318
title: "RHINENG-2394: allow Kafka TLS without CA cert"
state: merged
author: psegedy
created: 2023-10-04T09:02:16Z
updated: 2026-03-31T07:03:17Z
closed: 2023-10-04T12:35:45Z
merged: 2023-10-04T12:35:45Z
base_branch: master
head_branch: fix_kafka
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1318
---

# Pull Request #1318: RHINENG-2394: allow Kafka TLS without CA cert

**Author**: @psegedy
**Created**: October 04, 2023 at 09:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_kafka`

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

### Comment by @jira-linking on October 04, 2023 at 09:02 AM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2394


### Comment by @codecov-commenter on October 04, 2023 at 09:09 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `63.63636%` with `4 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.30%. Comparing base ([`7e95993`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/7e959935f2ea1a6d0f18c0b8b1d3a64ee0dd92ca?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2fcfa9f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2fcfa9faaf1631e53123dc33f4697526534d96e9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1693 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/mqueue/mqueue\_impl\_gokafka.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318?src=pr&el=tree&filepath=base%2Fmqueue%2Fmqueue_impl_gokafka.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | 55.55% | [3 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1318      +/-   ##
==========================================
+ Coverage   62.24%   62.30%   +0.06%     
==========================================
  Files         106      106              
  Lines        6687     6690       +3     
==========================================
+ Hits         4162     4168       +6     
+ Misses       1998     1995       -3     
  Partials      527      527              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.30% <63.63%> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1318?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @psegedy on October 04, 2023 at 12:03 PM UTC

/retest


---

## Reviews

### Review by @MichaelMraka - Approved on October 04, 2023 at 11:20 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1318*
