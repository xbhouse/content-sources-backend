---
type: pull_request
number: 1627
title: "RHINENG-15782: deprecate baselines"
state: merged
author: psegedy
created: 2025-04-04T12:56:21Z
updated: 2025-04-04T13:53:01Z
closed: 2025-04-04T13:53:01Z
merged: 2025-04-04T13:53:01Z
base_branch: master
head_branch: deprecate_baselines
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1627
---

# Pull Request #1627: RHINENG-15782: deprecate baselines

**Author**: @psegedy
**Created**: April 04, 2025 at 12:56 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `deprecate_baselines`

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

### Comment by @jira-linking on April 04, 2025 at 12:56 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-15782


### Comment by @codecov-commenter on April 04, 2025 at 01:02 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `10.00000%` with `9 lines` in your changes missing coverage. Please review.
> Project coverage is 58.19%. Comparing base [(`36ae164`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/36ae164a29cbdcd23da23e06a6e936fbff2f8086?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`8795f51`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8795f5157070b02cdd8d77f0b8b19c4aaae94cdb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/deprecations/deprecations.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?src=pr&el=tree&filepath=base%2Fdeprecations%2Fdeprecations.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kZXByZWNhdGlvbnMvZGVwcmVjYXRpb25zLmdv) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [manager/routes/routes.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?src=pr&el=tree&filepath=manager%2Froutes%2Froutes.go&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9yb3V0ZXMvcm91dGVzLmdv) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1627      +/-   ##
==========================================
- Coverage   58.27%   58.19%   -0.08%     
==========================================
  Files         146      146              
  Lines       11321    11330       +9     
==========================================
- Hits         6597     6594       -3     
- Misses       4142     4154      +12     
  Partials      582      582              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.19% <10.00%> (-0.08%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1627?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on April 04, 2025 at 01:18 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1627*
