---
type: pull_request
number: 1155
title: "feat(RHINENG-6147): introduce paginated requests to fetch applicable"
state: merged
author: mkholjuraev
created: 2024-01-09T21:54:31Z
updated: 2026-04-02T21:30:32Z
closed: 2024-01-29T13:36:51Z
merged: 2024-01-29T13:36:51Z
base_branch: master
head_branch: paginated-requests
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1155
---

# Pull Request #1155: feat(RHINENG-6147): introduce paginated requests to fetch applicable

**Author**: @mkholjuraev
**Created**: January 09, 2024 at 09:54 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `paginated-requests`

## Description

# Description

Associated Jira ticket: https://issues.redhat.com/browse/RHINENG-6147

TO BE REVIEWED AFTER: https://github.com/RedHatInsights/patchman-ui/pull/1154. This is intended to clean up hooks used in the codebase before introducing new ones for RHINENG-6147.

Introduces batched/paginated requests to fetch applicable advisories that are used to populate the remediation modal. This is required by the backend as they have disabled fetching everything using limit=-1. There should no be change in the functionality

# How to test the PR

1. Navigate to the systems page
2. click on some row action kebab
3. click on the 'Apply applicable advisories' action item
4. Observe that the remediation modal opens, the number of fetched issues matches with the total count and remedition can be generated

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

### Comment by @codecov-commenter on January 18, 2024 at 12:53 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `81.63265%` with `9 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.51%. Comparing base ([`47bcde9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/47bcde9a19b4a8549e61faf14573ea29981f9aa0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0fbdda1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/0fbdda144037605f4c1b14023994e63a4af0c5b3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 307 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/hooks/useFetchBatched.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FuseFetchBatched.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VGZXRjaEJhdGNoZWQuanM=) | 60.00% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsListAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | 88.23% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | 85.71% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1155      +/-   ##
==========================================
+ Coverage   63.24%   63.51%   +0.27%     
==========================================
  Files         122      124       +2     
  Lines        3107     3144      +37     
  Branches      809      817       +8     
==========================================
+ Hits         1965     1997      +32     
- Misses       1142     1147       +5     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1155?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on January 29, 2024 at 01:59 PM UTC

:tada: This PR is included in version 1.67.0 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on January 18, 2024 at 03:16 PM UTC

### Review by @bastilian - Approved on January 29, 2024 at 01:01 PM UTC

Nice work! Thank you @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1155*
