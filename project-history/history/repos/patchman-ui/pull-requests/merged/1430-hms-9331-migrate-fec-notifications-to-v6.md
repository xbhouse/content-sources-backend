---
type: pull_request
number: 1430
title: "HMS-9331: migrate fec-notifications to v6"
state: merged
author: xbhouse
created: 2025-11-04T22:19:25Z
updated: 2025-11-10T19:23:36Z
closed: 2025-11-10T19:23:36Z
merged: 2025-11-10T19:23:35Z
base_branch: master
head_branch: 9331
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1430
---

# Pull Request #1430: HMS-9331: migrate fec-notifications to v6

**Author**: @xbhouse
**Created**: November 04, 2025 at 10:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `9331`

## Description

# Description

Migrates fec-notifications to v6

Associated Jira ticket: # (issue)

[HMS-9331](https://issues.redhat.com/browse/HMS-9331)

# How to test the PR

1. Try any action that creates a toast notification (like exporting advisories, packages, systems or creating a plan remediation playbook)
2. Toast notifications should work as before, duplicate close button should be removed

# Before the change

<img width="720" height="183" alt="Screenshot From 2025-11-05 12-14-11" src="https://github.com/user-attachments/assets/f0d1fa38-7932-469d-90d2-655fe4c338fa" />

# After the change

<img width="719" height="174" alt="Screenshot From 2025-11-05 12-16-58" src="https://github.com/user-attachments/assets/af7a1cb5-505e-4ea8-9f25-b3ddf60603b2" />

# Dependent work link

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 04, 2025 at 10:30 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `40.00000%` with `15 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.35%. Comparing base ([`012e8fb`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/012e8fb301346433809fc4e61da4879d24c73544?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d837074`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d83707490bd8b7482c91368a2fcbb1d148cbf8ee?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSet%2FPatchSet.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetDetail%2FPatchSetDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/hooks/useRemediationDataProvider.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FuseRemediationDataProvider.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VSZW1lZGlhdGlvbkRhdGFQcm92aWRlci5qcw==) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Modals/AssignSystemsModal.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FSmartComponents%2FModals%2FAssignSystemsModal.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvQXNzaWduU3lzdGVtc01vZGFsLmpz) | 33.33% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/PatchSetWizard/steps/RequestProgress.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FRequestProgress.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXF1ZXN0UHJvZ3Jlc3MuanM=) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/Modals/useUnassignSystemsHook.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FSmartComponents%2FModals%2FuseUnassignSystemsHook.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvdXNlVW5hc3NpZ25TeXN0ZW1zSG9vay5qcw==) | 75.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rtComponents/Remediation/AsyncRemediationButton.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&filepath=src%2FSmartComponents%2FRemediation%2FAsyncRemediationButton.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9Bc3luY1JlbWVkaWF0aW9uQnV0dG9uLmpz) | 50.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1430      +/-   ##
==========================================
- Coverage   62.36%   62.35%   -0.01%     
==========================================
  Files         128      128              
  Lines        3329     3323       -6     
  Branches      894      894              
==========================================
- Hits         2076     2072       -4     
+ Misses       1132     1130       -2     
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1430?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Approved on November 10, 2025 at 11:03 AM UTC

Everything in this PR is clean and works as expected. Thank you for the contribution ❇️ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1430*
