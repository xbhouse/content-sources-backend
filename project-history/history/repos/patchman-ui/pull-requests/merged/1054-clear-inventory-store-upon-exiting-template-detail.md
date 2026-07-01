---
type: pull_request
number: 1054
title: "Clear Inventory store upon exiting Template detail page"
state: merged
author: leSamo
created: 2023-05-10T16:38:34Z
updated: 2023-06-05T21:42:11Z
closed: 2023-05-11T12:06:28Z
merged: 2023-05-11T12:06:28Z
base_branch: master
head_branch: template-detail-clear-inventory
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1054
---

# Pull Request #1054: Clear Inventory store upon exiting Template detail page

**Author**: @leSamo
**Created**: May 10, 2023 at 04:38 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `template-detail-clear-inventory`

## Description

Fixes https://issues.redhat.com/browse/SPM-2060

---

## Discussion

### Comment by @codecov-commenter on May 10, 2023 at 04:50 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1054?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`-0.81`** :warning:
> Comparison is base [(`1855269`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1855269343fc8b8241c9e9968dd4991d4e5c0598?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.09% compared to head [(`2cde557`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1054?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.29%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1054      +/-   ##
==========================================
- Coverage   63.09%   62.29%   -0.81%     
==========================================
  Files         119      119              
  Lines        2962     3005      +43     
  Branches      763      772       +9     
==========================================
+ Hits         1869     1872       +3     
- Misses       1093     1133      +40     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1054?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1054?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |

... and [2 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1054/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1054?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on May 11, 2023 at 11:17 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1054*
