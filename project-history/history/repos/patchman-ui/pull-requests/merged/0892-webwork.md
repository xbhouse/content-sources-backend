---
type: pull_request
number: 892
title: "Webwork"
state: merged
author: mkholjuraev
created: 2022-10-25T11:48:03Z
updated: 2024-04-03T09:21:54Z
closed: 2022-11-28T19:40:24Z
merged: 2022-11-28T19:40:24Z
base_branch: master
head_branch: webwork
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/892
---

# Pull Request #892: Webwork

**Author**: @mkholjuraev
**Created**: October 25, 2022 at 11:48 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `webwork`

## Description

Resolves: # RHIF-106
This a separate PR to improve paginated Remediation issues fetching by

1. Moving all cumbersome jobs into a web worker
2. creating grouped API requests with time interval
3. limit and offset are introduced into the API requests to speed up the SQL query.
4. disables remediation button, after it is already clicked and the module is being loading 

Please verify that there is a correct number of API requests to remediate all selected entities and that pagination works as it should.

---

## Discussion

### Comment by @codecov-commenter on November 04, 2022 at 12:14 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.26**% // Head: **70.52**% // Decreases project coverage by **`-1.74%`** :warning:
> Coverage data is based on head [(`7f1afba`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`daae4a7`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/daae4a7abe4c099ffefb14dc7c0fdb41c53633e3?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 3.50% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #892      +/-   ##
==========================================
- Coverage   72.26%   70.52%   -1.75%     
==========================================
  Files         109      109              
  Lines        2564     2585      +21     
  Branches      655      665      +10     
==========================================
- Hits         1853     1823      -30     
- Misses        649      692      +43     
- Partials       62       70       +8     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.13% <ø> (-0.32%)` | :arrow_down: |
| [src/Utilities/RemediationPairs.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SZW1lZGlhdGlvblBhaXJzLmpz) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `56.48% <ø> (+2.17%)` | :arrow_up: |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `82.85% <100.00%> (+0.24%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `80.76% <0.00%> (-3.24%)` | :arrow_down: |
| [src/Utilities/useRemediationDataProvider.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VSZW1lZGlhdGlvbkRhdGFQcm92aWRlci5qcw==) | | |
| [...ntationalComponents/Header/HeaderBreadcrumbs.cy.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyQnJlYWRjcnVtYnMuY3kuanM=) | `0.00% <0.00%> (ø)` | |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `93.61% <0.00%> (+0.13%)` | :arrow_up: |
| ... and [3 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/892?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 09, 2022 at 12:23 PM UTC

@RedHatInsights/team-interact I am ready to do a pair reviewing if this helps ;).

### Comment by @jiridostal on November 28, 2022 at 07:58 PM UTC

:tada: This PR is included in version 1.57.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.57.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.57.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on November 14, 2022 at 10:33 AM UTC

### Review by @mkholjuraev - Commented on November 14, 2022 at 10:53 AM UTC

### Review by @mkholjuraev - Commented on November 14, 2022 at 11:51 AM UTC

### Review by @bastilian - Commented on November 14, 2022 at 11:57 AM UTC

### Review by @mkholjuraev - Commented on November 14, 2022 at 12:01 PM UTC

### Review by @mkholjuraev - Commented on November 14, 2022 at 12:04 PM UTC

### Review by @bastilian - Commented on November 14, 2022 at 12:08 PM UTC

### Review by @bastilian - Commented on November 14, 2022 at 12:09 PM UTC

### Review by @mkholjuraev - Commented on November 14, 2022 at 12:17 PM UTC

### Review by @gkarat - Approved on November 22, 2022 at 11:48 AM UTC

lgtm! tested with the steps you provided, works ok for accounts with lesser number of advisories/systems. a few comments on the code: I think there are places where you can work with data more efficiently by reducing loops

### Review by @mkholjuraev - Commented on November 22, 2022 at 12:30 PM UTC

### Review by @mkholjuraev - Commented on November 22, 2022 at 12:33 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/892*
