---
type: pull_request
number: 1053
title: "Remove redundant group column from package systems table"
state: merged
author: leSamo
created: 2023-05-10T16:30:33Z
updated: 2023-06-05T21:42:08Z
closed: 2023-05-15T19:02:01Z
merged: 2023-05-15T19:02:01Z
base_branch: master
head_branch: remove-group-column
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1053
---

# Pull Request #1053: Remove redundant group column from package systems table

**Author**: @leSamo
**Created**: May 10, 2023 at 04:30 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-group-column`

## Description

Removed this empty "group" column:
![Screenshot from 2023-05-10 18-28-24](https://github.com/RedHatInsights/patchman-ui/assets/8426204/72a18ae9-c4de-458d-8cde-8eabc6575371)


---

## Discussion

### Comment by @codecov-commenter on May 10, 2023 at 04:40 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1053?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.79`** :warning:
> Comparison is base [(`1855269`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1855269343fc8b8241c9e9968dd4991d4e5c0598?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.09% compared to head [(`eba95c8`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1053?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.31%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1053      +/-   ##
==========================================
- Coverage   63.09%   62.31%   -0.79%     
==========================================
  Files         119      119              
  Lines        2962     3004      +42     
  Branches      763      772       +9     
==========================================
+ Hits         1869     1872       +3     
- Misses       1093     1132      +39     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1053?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1053?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `90.74% <ø> (ø)` | |

... and [2 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1053/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1053?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on May 11, 2023 at 12:51 PM UTC

LGTM! :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1053*
