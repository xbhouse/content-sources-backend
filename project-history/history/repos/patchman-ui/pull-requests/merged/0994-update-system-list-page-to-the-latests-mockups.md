---
type: pull_request
number: 994
title: "Update System list page to the latests mockups"
state: merged
author: leSamo
created: 2023-03-17T02:36:36Z
updated: 2023-05-08T18:16:35Z
closed: 2023-03-17T16:29:02Z
merged: 2023-03-17T16:29:02Z
base_branch: master
head_branch: fix-system-list
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/994
---

# Pull Request #994: Update System list page to the latests mockups

**Author**: @leSamo
**Created**: March 17, 2023 at 02:36 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-system-list`

## Description

# Description

Associated Jira ticket: [SPM-1956](https://issues.redhat.com/browse/SPM-1956) [SPM-1959](https://issues.redhat.com/browse/SPM-1959)

Mockups: https://www.sketch.com/s/598ac9a1-16ce-47b1-8515-e79928a2b2fb/a/Om3zKYq

- changed /systems API from v2 to v3
  - this adds new `baseline_id` parameter to the response from which link to Template detail can be correctly created
  - simplifies OS parameters into one
- System list table
  - replace "Applicable advisories" column with "Installable advisories" column
  - rename "Packages" column to "Installed packages"
  - change the order of the columns according to the mockups
  - move "Assign to a template" button into the toolbar kebab
- Advisory systems table
  - removed packages and applicable advisories columns

---

## Discussion

### Comment by @codecov-commenter on March 17, 2023 at 03:31 AM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> :exclamation: No coverage uploaded for pull request base (`master@b26f944`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#section-missing-base-commit).
> Patch coverage: 43.75% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##             master     #994   +/-   ##
=========================================
  Coverage          ?   68.75%           
=========================================
  Files             ?      116           
  Lines             ?     2820           
  Branches          ?      720           
=========================================
  Hits              ?     1939           
  Misses            ?      881           
  Partials          ?        0           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `84.61% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `61.11% <22.22%> (ø)` | |
| [src/Utilities/SystemsHelpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | `26.66% <50.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `83.82% <100.00%> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `67.81% <100.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `58.18% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/994?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on March 28, 2023 at 08:57 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 17, 2023 at 03:15 PM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/994*
