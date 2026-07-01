---
type: pull_request
number: 886
title: "Fix remediations"
state: merged
author: mkholjuraev
created: 2022-10-13T12:30:20Z
updated: 2022-10-13T20:23:19Z
closed: 2022-10-13T20:07:44Z
merged: 2022-10-13T20:07:44Z
base_branch: master
head_branch: fix-remediations
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/886
---

# Pull Request #886: Fix remediations

**Author**: @mkholjuraev
**Created**: October 13, 2022 at 12:30 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-remediations`

## Description

This fixes the remediation on the Systems page after it broke because of the new pagination enabled on `/view/advisories/systems `endpoint. Now remediation issue pairs are fetched from` /view/systems/advisories`, then passed to the Remediation module. 

To test:

1. Visit the systems page
2. Select one or more systems
3. Click on remediation
4. Observe that the request to fetch remediation pairs is made to /view/systems/advisories
5. Remediation modal does not brake
6. You can remediate selected system advisories

---

## Discussion

### Comment by @codecov-commenter on October 13, 2022 at 12:39 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.12**% // Head: **71.67**% // Decreases project coverage by **`-0.44%`** :warning:
> Coverage data is based on head [(`3c68469`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`90468d4`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/90468d4561d2b6d81cb062c1d5c2728d22d1915d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 29.16% of modified lines in pull request are covered.

> :exclamation: Current head 3c68469 differs from pull request most recent head f16c22c. Consider uploading reports for the commit f16c22c to get more accurate results

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #886      +/-   ##
==========================================
- Coverage   72.12%   71.67%   -0.45%     
==========================================
  Files         108      108              
  Lines        2500     2521      +21     
  Branches      644      650       +6     
==========================================
+ Hits         1803     1807       +4     
- Misses        636      651      +15     
- Partials       61       63       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `43.47% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `79.71% <0.00%> (-2.38%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.24% <7.69%> (-3.03%)` | :arrow_down: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `56.89% <25.00%> (-1.14%)` | :arrow_down: |
| [src/Utilities/useRemediationDataProvider.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VSZW1lZGlhdGlvbkRhdGFQcm92aWRlci5qcw==) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/886?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on October 13, 2022 at 08:23 PM UTC

:tada: This PR is included in version 1.52.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.52.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.52.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on October 13, 2022 at 07:21 PM UTC

PR works as intended, great job! LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/886*
