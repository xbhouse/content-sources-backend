---
type: pull_request
number: 1334
title: "RHINENG-2610: use latest_repo_change timestamp for repo based re-eval"
state: merged
author: psegedy
created: 2023-11-03T15:39:16Z
updated: 2026-04-06T03:00:25Z
closed: 2023-11-06T12:19:49Z
merged: 2023-11-06T12:19:49Z
base_branch: master
head_branch: repobased
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1334
---

# Pull Request #1334: RHINENG-2610: use latest_repo_change timestamp for repo based re-eval

**Author**: @psegedy
**Created**: November 03, 2023 at 03:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `repobased`

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

### Comment by @jira-linking on November 03, 2023 at 03:39 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-2610


### Comment by @codecov-commenter on April 06, 2026 at 03:00 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `85.71429%` with `1 line` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 61.93%. Comparing base ([`655d3c8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/655d3c852d2e8fcb2ed93bd13c62c7924268c8f6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`aa55af6`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/aa55af67af5b53b0679ab98bf74cf552c856aec6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1668 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [tasks/vmaas\_sync/repo\_based.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334?src=pr&el=tree&filepath=tasks%2Fvmaas_sync%2Frepo_based.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-dGFza3Mvdm1hYXNfc3luYy9yZXBvX2Jhc2VkLmdv) | 83.33% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1334      +/-   ##
==========================================
+ Coverage   61.81%   61.93%   +0.12%     
==========================================
  Files         107      107              
  Lines        6714     6762      +48     
==========================================
+ Hits         4150     4188      +38     
- Misses       2034     2040       +6     
- Partials      530      534       +4     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `61.93% <85.71%> (+0.12%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1334?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on November 03, 2023 at 07:33 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1334*
