---
type: pull_request
number: 1026
title: "Add status filter to Advisory systems table"
state: merged
author: leSamo
created: 2023-04-17T14:26:56Z
updated: 2023-05-09T16:05:39Z
closed: 2023-04-18T17:19:21Z
merged: 2023-04-18T17:19:21Z
base_branch: master
head_branch: advisory-status-filter
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1026
---

# Pull Request #1026: Add status filter to Advisory systems table

**Author**: @leSamo
**Created**: April 17, 2023 at 02:26 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisory-status-filter`

## Description

Part of https://issues.redhat.com/browse/SPM-1949

---

## Discussion

### Comment by @codecov-commenter on April 17, 2023 at 02:38 PM UTC

## [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`100.00`**% and project coverage change: **`+0.10`** :tada:
> Comparison is base [(`9d76343`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9d7634343f8ab378f07d99164d86346ea5a2db0d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.12% compared to head [(`dd68712`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 64.22%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1026      +/-   ##
==========================================
+ Coverage   64.12%   64.22%   +0.10%     
==========================================
  Files         116      117       +1     
  Lines        2832     2840       +8     
  Branches      733      734       +1     
==========================================
+ Hits         1816     1824       +8     
  Misses       1016     1016              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `84.90% <ø> (ø)` | |
| [...tationalComponents/Filters/AdvisoryStatusFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL0Fkdmlzb3J5U3RhdHVzRmlsdGVyLmpz) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report in Codecov by Sentry](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/1026?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 09, 2023 at 04:05 PM UTC

:tada: This PR is included in version 1.63.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on April 18, 2023 at 01:52 PM UTC

LGTM! works as expected.

### Review by @Fewwy - Approved on April 18, 2023 at 02:19 PM UTC

LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1026*
