---
type: pull_request
number: 1420
title: "chore: unify linting and formatting configs"
state: merged
author: katarinazaprazna
created: 2025-10-24T13:30:05Z
updated: 2025-11-05T21:12:38Z
closed: 2025-11-04T14:34:09Z
merged: 2025-11-04T14:34:09Z
base_branch: master
head_branch: unify-linting-and-formatting-configs
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1420
---

# Pull Request #1420: chore: unify linting and formatting configs

**Author**: @katarinazaprazna
**Created**: October 24, 2025 at 01:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `unify-linting-and-formatting-configs`

## Description

# Description

> Planning to merge after [RHINENG-18569](https://github.com/RedHatInsights/patchman-ui/pull/1428) to prevent merge conflicts.

### Configuration Alignment

- I've unified the development configuration to match the content-sources-frontend repository for easier maintenance across both projects
- Prettier was added and configured to handle all code formatting, allowing ESLint to focus strictly on code quality rules alone. This change allowed us to remove the redundant EditorConfig file and clean up `/* global ... */` comments by moving environment settings into the main ESLint file

### Tool Modernization and Scope

- ESLint was bumped from v7 to v9 (Flat Config Format)
- For styles, Prettier now handles SCSS formatting, and Stylelint handles SCSS quality rules
- I narrowed ESLint's scope to only lint the `/src` directory, excluding the `/config` folder to prevent issues with older conventions

### Git

- Adjusted `commitlint` config to enforce the 50/72 rule for commit message formatting, aligning with the content-sources-frontend repository

# How to test the PR

- Verify that the app builds without any errors.
- Ensure that all automated tests pass.
- Ensure that output of running `npm run lint:js:fix`, `npm run lint:sass` and `npm run prettier:check` will be a message confirming that all linting and formatting rules have been satisfied.

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [x] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 01, 2025 at 07:37 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `53.28244%` with `612 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.36%. Comparing base ([`2511619`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/251161969b050dec92d2387b8b2d618e1a9ef102?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`575daa5`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/575daa51e78487553bccec80f7177f33d182ebf7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetDetail%2FPatchSetDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | 0.00% | [107 Missing and 13 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/PatchSet/PatchSet.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSet%2FPatchSet.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | 0.00% | [63 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FReviewSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | 0.00% | [47 Missing and 4 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemDetail%2FInventoryDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | 0.00% | [24 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FRoutes.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | 0.00% | [24 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Modals/AssignSystemsModal.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FModals%2FAssignSystemsModal.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvQXNzaWduU3lzdGVtc01vZGFsLmpz) | 50.00% | [23 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2FPatchSetWizard.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | 0.00% | [18 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/App.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FApp.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | 0.00% | [15 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FPackageSystems%2FPackageSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | 68.42% | [15 Missing and 3 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2FWizardAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | 39.28% | [12 Missing and 5 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| ... and [40 more](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1420      +/-   ##
==========================================
- Coverage   62.93%   62.36%   -0.57%     
==========================================
  Files         128      128              
  Lines        3399     3329      -70     
  Branches      906      901       -5     
==========================================
- Hits         2139     2076      -63     
+ Misses       1141     1132       -9     
- Partials      119      121       +2     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1420?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @sentry on November 05, 2025 at 09:12 PM UTC

## Issues attributed to commits in this pull request
This pull request was merged and Sentry observed the following issues:

* ‼️ [**TypeError: can't access property "forEach" of unde...**](https://red-hat-it.sentry.io/issues/7001351594/?referrer=github-pr-bot) in `Prod`


---

## Reviews

### Review by @TenSt - Approved on October 30, 2025 at 10:40 AM UTC

lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1420*
