---
type: pull_request
number: 991
title: "Template detail empty state button launches wizard"
state: merged
author: leSamo
created: 2023-03-14T00:07:00Z
updated: 2023-05-08T18:16:38Z
closed: 2023-03-17T16:33:16Z
merged: 2023-03-17T16:33:16Z
base_branch: master
head_branch: wizard-edit-create
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/991
---

# Pull Request #991: Template detail empty state button launches wizard

**Author**: @leSamo
**Created**: March 14, 2023 at 12:07 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `wizard-edit-create`

## Description

# Description

Associated Jira ticket: [SPM-1893](https://issues.redhat.com/browse/SPM-1893)

- if template has 0 systems it shows an empty state on template detail page, the button in the empty state now correctly opens wizard; previously it didn't do anything
- removed unused 'assign' template wizard type
- editing a template from template detail page now has correct "edit" verb instead of "create" in the titles

---

## Discussion

### Comment by @codecov-commenter on March 14, 2023 at 12:10 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/991?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> :exclamation: No coverage uploaded for pull request base (`master@b26f944`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#section-missing-base-commit).
> Patch coverage: 0.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##             master     #991   +/-   ##
=========================================
  Coverage          ?   69.07%           
=========================================
  Files             ?      116           
  Lines             ?     2804           
  Branches          ?      717           
=========================================
  Hits              ?     1937           
  Misses            ?      867           
  Partials          ?        0           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/991?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/991?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/991?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/991?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `39.13% <0.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/991?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 28, 2023 at 08:57 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 17, 2023 at 03:03 PM UTC

Wizard now opens as expected! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/991*
