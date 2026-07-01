---
type: pull_request
number: 1173
title: "chore: use paginated requests to fetch all template names"
state: merged
author: mkholjuraev
created: 2024-03-12T11:44:06Z
updated: 2026-03-31T14:40:45Z
closed: 2024-03-19T12:55:17Z
merged: 2024-03-19T12:55:17Z
base_branch: master
head_branch: paginated-requests-template-names
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1173
---

# Pull Request #1173: chore: use paginated requests to fetch all template names

**Author**: @mkholjuraev
**Created**: March 12, 2024 at 11:44 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `paginated-requests-template-names`

## Description

# Description

Associated Jira ticket: # (issue)

Fetches all template names using paginated requests


# How to test the PR

1. Visit the templates page
2. Try creating a new template
3. Observe that 
     - the next button is disabled unless all taken/reserved template names are fetched from the API 
     - spinner icon is displayed while the requests are made
     - once all requests are resolved, the name input is validated
4. Visit some template detail page
5. Try adding more systems to an existing template with systems assigned
6. Observe that the already assigned systems are marked as selected

# Before the change


# After the change


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

### Comment by @codecov-commenter on March 18, 2024 at 11:35 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `59.45946%` with `15 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.48%. Comparing base ([`005bc44`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/005bc444703987f19cf07611592572f8dce00ca7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`8c066d7`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/8c066d721b162dfd5e308467724d58648f181ce4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 289 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FReviewSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | 8.33% | [11 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2FWizardAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | 75.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2FPatchSetWizard.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...ts/PatchSetWizard/steps/ConfigurationStepFields.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FConfigurationStepFields.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9Db25maWd1cmF0aW9uU3RlcEZpZWxkcy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1173      +/-   ##
==========================================
+ Coverage   63.15%   63.48%   +0.33%     
==========================================
  Files         127      127              
  Lines        3224     3226       +2     
  Branches      826      826              
==========================================
+ Hits         2036     2048      +12     
+ Misses       1188     1178      -10     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1173?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on April 03, 2024 at 01:00 PM UTC

:tada: This PR is included in version 1.67.3 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Changes Requested on March 15, 2024 at 03:06 PM UTC

When creating a new template, the systems table is stuck on loading
![image](https://github.com/RedHatInsights/patchman-ui/assets/20592948/0c7240a9-88a2-43c4-9ea1-92685d40124f)
The request with systems is received but is not displaying the systems


### Review by @AsToNlele - Approved on March 19, 2024 at 12:49 PM UTC

LGTM, thanks for the fix :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1173*
