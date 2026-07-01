---
type: pull_request
number: 1327
title: "Countermeasures from threat modeling"
state: merged
author: psegedy
created: 2023-10-16T16:05:29Z
updated: 2026-04-01T06:56:51Z
closed: 2023-10-19T13:26:43Z
merged: 2023-10-19T13:26:43Z
base_branch: master
head_branch: threat
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1327
---

# Pull Request #1327: Countermeasures from threat modeling

**Author**: @psegedy
**Created**: October 16, 2023 at 04:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `threat`

## Description

- [x] fine tune limits
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

### Comment by @jira-linking on October 16, 2023 at 04:05 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2433
https://issues.redhat.com/browse/RHINENG-2426


### Comment by @codecov-commenter on October 16, 2023 at 04:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `15.78947%` with `16 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.04%. Comparing base ([`48e75b1`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/48e75b1cc85a2a73cbcd64db9a6558400f381d27?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f819487`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f81948734dabda30dfa26eb19def9b2ad122a7a5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1678 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [manager/middlewares/limits.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327?src=pr&el=tree&filepath=manager%2Fmiddlewares%2Flimits.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9saW1pdHMuZ28=) | 0.00% | [16 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1327      +/-   ##
==========================================
- Coverage   62.19%   62.04%   -0.15%     
==========================================
  Files         107      107              
  Lines        6705     6737      +32     
==========================================
+ Hits         4170     4180      +10     
- Misses       2007     2029      +22     
  Partials      528      528              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.04% <15.78%> (-0.15%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1327?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on October 18, 2023 at 12:30 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1327*
