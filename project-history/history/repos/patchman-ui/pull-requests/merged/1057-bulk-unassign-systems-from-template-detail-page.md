---
type: pull_request
number: 1057
title: "Bulk unassign systems from template detail page"
state: merged
author: leSamo
created: 2023-05-19T13:41:19Z
updated: 2023-06-05T21:41:59Z
closed: 2023-05-23T12:31:55Z
merged: 2023-05-23T12:31:55Z
base_branch: master
head_branch: bulk-unassing-template-detail
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1057
---

# Pull Request #1057: Bulk unassign systems from template detail page

**Author**: @leSamo
**Created**: May 19, 2023 at 01:41 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `bulk-unassing-template-detail`

## Description

https://issues.redhat.com/browse/SPM-1943

---

## Discussion

### Comment by @codecov-commenter on May 19, 2023 at 01:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`7.14`**% and project coverage change: **`-0.22`** :warning:
> Comparison is base [(`546dd7e`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/546dd7e53b01ef76706e6f371978f0203c23b437?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 63.03% compared to head [(`8fe9231`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.82%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1057      +/-   ##
==========================================
- Coverage   63.03%   62.82%   -0.22%     
==========================================
  Files         119      119              
  Lines        2949     2959      +10     
  Branches      754      757       +3     
==========================================
  Hits         1859     1859              
- Misses       1090     1100      +10     
```


| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `44.44% <0.00%> (ø)` | |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/useOnSelect.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VPblNlbGVjdC5qcw==) | `98.57% <50.00%> (-1.43%)` | :arrow_down: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1057?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on May 23, 2023 at 10:29 AM UTC

tested locally. the bulk select is now visible, and I could remove more systems. lgtm, thank you 👍🏼 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1057*
