---
type: pull_request
number: 1095
title: "Add OS and template columns and OS filter to Package detail table"
state: merged
author: leSamo
created: 2023-07-20T09:17:41Z
updated: 2023-08-14T13:21:50Z
closed: 2023-07-21T12:58:50Z
merged: 2023-07-21T12:58:50Z
base_branch: master
head_branch: package-detail-filters
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1095
---

# Pull Request #1095: Add OS and template columns and OS filter to Package detail table

**Author**: @leSamo
**Created**: July 20, 2023 at 09:17 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `package-detail-filters`

## Description

https://issues.redhat.com/browse/SPM-1951

---

## Discussion

### Comment by @codecov-commenter on July 20, 2023 at 09:24 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`33.33`**% and project coverage change: **`-1.66`** :warning:
> Comparison is base [(`64549d6`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/64549d6b590761cdd83cd6334b2ddbb068a5b438?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.70% compared to head [(`98c6fc1`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.04%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1095      +/-   ##
==========================================
- Coverage   62.70%   61.04%   -1.66%     
==========================================
  Files         119      119              
  Lines        2984     3068      +84     
  Branches      765      795      +30     
==========================================
+ Hits         1871     1873       +2     
- Misses       1113     1195      +82     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `59.45% <0.00%> (-7.21%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.90% <100.00%> (+0.16%)` | :arrow_up: |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `57.14% <100.00%> (+0.44%)` | :arrow_up: |

... and [1 file with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1095?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on August 07, 2023 at 02:35 PM UTC

:tada: This PR is included in version 1.63.11 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.11)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on July 21, 2023 at 12:48 PM UTC

Good job! LGTM :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1095*
