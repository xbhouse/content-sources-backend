---
type: pull_request
number: 823
title: "fix(InventoryDetail): ESSNTL-2698 handle disconnected systems"
state: merged
author: mkholjuraev
created: 2022-06-18T17:00:09Z
updated: 2026-04-04T06:09:08Z
closed: 2022-06-21T09:26:31Z
merged: 2022-06-21T09:26:31Z
base_branch: master
head_branch: inventory-details
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/823
---

# Pull Request #823: fix(InventoryDetail): ESSNTL-2698 handle disconnected systems

**Author**: @mkholjuraev
**Created**: June 18, 2022 at 05:00 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `inventory-details`

## Description

Resolves: [ESSNTL-2698](https://issues.redhat.com/browse/ESSNTL-2698).

To reproduce:

1. run this app against stage env
2. visit patch systems page
3. click on any system that is disconnected from insights-client (this kind of systems are marked with disconnected icon in the last seen column).
4. System details page should have special empty component for disonnected systems.  

---

## Discussion

### Comment by @codecov-commenter on June 18, 2022 at 05:07 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `71.42857%` with `6 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 71.69%. Comparing base ([`bf4c632`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bf4c632dbc877772f93f9b99ab152c4e7830af33?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c354945`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c354945cfe1ec7fad1a4b434cb2059e1ee3d09b3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 636 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemDetail%2FInventoryDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | 83.33% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/store/Reducers/SystemDetailStore.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?src=pr&el=tree&filepath=src%2Fstore%2FReducers%2FSystemDetailStore.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbURldGFpbFN0b3JlLmpz) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master     #823      +/-   ##
==========================================
- Coverage   71.91%   71.69%   -0.23%     
==========================================
  Files         103      103              
  Lines        2482     2487       +5     
  Branches      642      647       +5     
==========================================
- Hits         1785     1783       -2     
- Misses        636      642       +6     
- Partials       61       62       +1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/823?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @jiridostal on June 21, 2022 at 09:43 AM UTC

:tada: This PR is included in version 1.48.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on June 21, 2022 at 09:26 AM UTC

I tried reviewing the PR. Looks good to me :). 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/823*
