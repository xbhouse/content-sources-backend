---
type: pull_request
number: 1144
title: "Disable template edit actions when missing permission"
state: merged
author: leSamo
created: 2023-11-24T13:11:08Z
updated: 2023-12-13T11:16:11Z
closed: 2023-12-04T12:16:02Z
merged: 2023-12-04T12:16:02Z
base_branch: master
head_branch: fix-rbac-template
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1144
---

# Pull Request #1144: Disable template edit actions when missing permission

**Author**: @leSamo
**Created**: November 24, 2023 at 01:11 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-rbac-template`

## Description

Associated Jira ticket: [RHINENG-3143](https://issues.redhat.com/browse/RHINENG-3143)

When user is missing permissions to edit templates:
- Disable `Apply to systems` button in `NoAppliedSystems` empty state in Template detail page
- Disable edit and delete template buttons in Template detail page header dropdown
- Disable template add and remove template buttons is System detail header dropdown
- Disable add and remove template table actions (both bulk and row) in System list table

I tried to add tooltip why the button is disabled where available, but due to the limitations of external components it was not always possible

---

## Discussion

### Comment by @codecov-commenter on November 24, 2023 at 01:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: `12 lines` in your changes are missing coverage. Please review.
> Comparison is base [(`eaf8d12`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/eaf8d12a537a0c90797bcd96051bf4c6ce1b1f75?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.49% compared to head [(`397b633`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.39%.
> Report is 1 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | 0.00% | [7 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | 33.33% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1144      +/-   ##
==========================================
- Coverage   63.49%   63.39%   -0.10%     
==========================================
  Files         122      122              
  Lines        3087     3101      +14     
  Branches      799      808       +9     
==========================================
+ Hits         1960     1966       +6     
- Misses       1127     1135       +8     
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1144?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on December 13, 2023 at 11:15 AM UTC

:tada: This PR is included in version 1.65.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on November 28, 2023 at 11:55 AM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1144*
