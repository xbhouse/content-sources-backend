---
type: pull_request
number: 1306
title: "chore: Declarative column merging"
state: merged
author: leSamo
created: 2025-03-13T16:50:53Z
updated: 2025-03-21T12:55:49Z
closed: 2025-03-21T12:55:48Z
merged: 2025-03-21T12:55:48Z
base_branch: master
head_branch: declarative-column-merging
labels: ["refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1306
---

# Pull Request #1306: chore: Declarative column merging

**Author**: @leSamo
**Created**: March 13, 2025 at 04:50 PM UTC
**Status**: Merged
**Labels**: `refactor`
**Base**: `master` ← **Head**: `declarative-column-merging`

## Description

This is a necessary prerequisite for Column management implementation. This allows us to simply change properties of table columns retrieved from Inventory (display name, workspaces (groups), tags, last seen).

Obsolete function `systemsColumnsMerger` will be removed together with old template pages.

---

## Discussion

### Comment by @codecov-commenter on March 13, 2025 at 05:18 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `44.44444%` with `5 lines` in your changes missing coverage. Please review.
> Project coverage is 62.37%. Comparing base [(`6070f6e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6070f6e23ef224bdad51f5b413640395e7ecbced?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5467545`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5467545611af3103ed690188b1cdd650af4566c4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsListAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/AdvisorySystems/AdvisorySystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisorySystems%2FAdvisorySystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zVGFibGUuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&filepath=src%2FSmartComponents%2FPackageSystems%2FPackageSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Systems/SystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNUYWJsZS5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1306      +/-   ##
==========================================
- Coverage   62.61%   62.37%   -0.24%     
==========================================
  Files         126      126              
  Lines        3322     3328       +6     
  Branches      835      836       +1     
==========================================
- Hits         2080     2076       -4     
- Misses       1242     1252      +10     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.37% <44.44%> (-0.24%)` | :arrow_down: |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.37% <44.44%> (-0.24%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1306?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>🚀 New features to boost your workflow: </summary>

- ❄ [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- 📦 [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @bastilian - Approved on March 21, 2025 at 11:06 AM UTC

Nice refactor! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1306*
