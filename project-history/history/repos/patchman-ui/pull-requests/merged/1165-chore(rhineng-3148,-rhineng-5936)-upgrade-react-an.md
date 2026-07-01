---
type: pull_request
number: 1165
title: "chore(RHINENG-3148, RHINENG-5936): Upgrade React and PatternFly"
state: merged
author: gkarat
created: 2024-02-28T13:22:03Z
updated: 2024-03-13T13:32:35Z
closed: 2024-03-12T11:05:03Z
merged: 2024-03-12T11:05:03Z
base_branch: master
head_branch: pf-react-migration
labels: ["dependencies", "released", ":rage1:"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1165
---

# Pull Request #1165: chore(RHINENG-3148, RHINENG-5936): Upgrade React and PatternFly

**Author**: @gkarat
**Created**: February 28, 2024 at 01:22 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `released`, `:rage1:`
**Base**: `master` ← **Head**: `pf-react-migration`

## Description

Implements https://issues.redhat.com/browse/RHINENG-3148, https://issues.redhat.com/browse/RHINENG-5936.

This upgrades React to use version 18, PatternFly to use version 5, and other related packages that affected with this upgrade to avoid conflicting peer dependencies. This makes sure the unit and cypress test are running fine.

## How to test

Make sure the application is built both for development and production environment. Ensure you are able to run it locally, the components (tables, dropdown, modals, empty states, etc.) are not broken and visually are up-to-date with PF5 or stay the same.

---

## Discussion

### Comment by @codecov-commenter on March 11, 2024 at 10:13 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `45.94595%` with `20 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.15%. Comparing base [(`b9294b4`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b9294b43f0a317f802aae0b672e08e838e1102b9?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`95200f7`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...Components/PatchSetWizard/steps/RequestProgress.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXF1ZXN0UHJvZ3Jlc3MuanM=) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...ntationalComponents/Header/HeaderBreadcrumbs.cy.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyQnJlYWRjcnVtYnMuY3kuanM=) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/PatchSetWizard/InputFields/NameField.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9OYW1lRmllbGQuanM=) | 81.25% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [.../PresentationalComponents/Filters/VersionFilter.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1ZlcnNpb25GaWx0ZXIuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...tationalComponents/Snippets/DescriptionWithLink.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9EZXNjcmlwdGlvbldpdGhMaW5rLmpz) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Modals/DeleteSetModal.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvRGVsZXRlU2V0TW9kYWwuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...nts/PatchSetWizard/InputFields/DescriptionField.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9EZXNjcmlwdGlvbkZpZWxkLmpz) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1165      +/-   ##
==========================================
+ Coverage   63.11%   63.15%   +0.04%     
==========================================
  Files         127      127              
  Lines        3215     3224       +9     
  Branches      822      826       +4     
==========================================
+ Hits         2029     2036       +7     
- Misses       1186     1188       +2     
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1165?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 13, 2024 at 01:32 PM UTC

:tada: This PR is included in version 1.67.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on March 08, 2024 at 11:20 AM UTC

Looks good to me! No major issues, only icon colors seem to be affected a bit

### Review by @gkarat - Commented on March 11, 2024 at 10:04 AM UTC

### Review by @gkarat - Commented on March 11, 2024 at 10:04 AM UTC

### Review by @gkarat - Commented on March 11, 2024 at 10:15 AM UTC

### Review by @mkholjuraev - Approved on March 12, 2024 at 10:24 AM UTC

LGTM! I have tested again and did not spot any more issues. 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1165*
