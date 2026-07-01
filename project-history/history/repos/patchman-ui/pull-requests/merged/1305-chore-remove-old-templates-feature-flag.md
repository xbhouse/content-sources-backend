---
type: pull_request
number: 1305
title: "chore: Remove old templates feature flag"
state: merged
author: leSamo
created: 2025-03-07T13:26:19Z
updated: 2025-03-13T12:30:47Z
closed: 2025-03-13T12:30:47Z
merged: 2025-03-13T12:30:47Z
base_branch: master
head_branch: remove-template-ff
labels: ["refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1305
---

# Pull Request #1305: chore: Remove old templates feature flag

**Author**: @leSamo
**Created**: March 07, 2025 at 01:26 PM UTC
**Status**: Merged
**Labels**: `refactor`
**Base**: `master` ← **Head**: `remove-template-ff`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on March 07, 2025 at 01:34 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `53.33333%` with `7 lines` in your changes missing coverage. Please review.
> Project coverage is 62.61%. Comparing base [(`6c5a411`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6c5a4115a569d93172475c21eaed687fdf843843?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`f709f22`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f709f22bfd96b8fabd2466b6469cf025b60d8ef0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemDetail%2FInventoryDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FReviewSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/AdvisorySystems/AdvisorySystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisorySystems%2FAdvisorySystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zVGFibGUuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Systems/SystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNUYWJsZS5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1305      +/-   ##
==========================================
+ Coverage   62.04%   62.61%   +0.56%     
==========================================
  Files         126      126              
  Lines        3367     3322      -45     
  Branches      868      835      -33     
==========================================
- Hits         2089     2080       -9     
+ Misses       1278     1242      -36     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.61% <53.33%> (+0.56%)` | :arrow_up: |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.61% <53.33%> (+0.56%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1305?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>🚀 New features to boost your workflow: </summary>

- ❄ [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- 📦 [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @leSamo on March 08, 2025 at 04:21 PM UTC

/retest

### Comment by @leSamo on March 08, 2025 at 04:32 PM UTC

/retest

### Comment by @leSamo on March 08, 2025 at 06:18 PM UTC

/retest

### Comment by @leSamo on March 10, 2025 at 11:08 AM UTC

/retest

### Comment by @leSamo on March 10, 2025 at 10:06 PM UTC

/retest

### Comment by @leSamo on March 12, 2025 at 09:16 AM UTC

/retest

---

## Reviews

### Review by @gkarat - Commented on March 12, 2025 at 10:30 AM UTC

### Review by @leSamo - Commented on March 12, 2025 at 12:50 PM UTC

### Review by @leSamo - Commented on March 12, 2025 at 12:53 PM UTC

### Review by @gkarat - Approved on March 13, 2025 at 10:00 AM UTC

LGTM, thanks @leSamo 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1305*
