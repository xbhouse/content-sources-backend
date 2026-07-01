---
type: pull_request
number: 797
title: "fix(SPM-1492): system details page refreshs after patch-set change"
state: merged
author: mkholjuraev
created: 2022-05-12T09:45:53Z
updated: 2024-04-03T09:21:46Z
closed: 2022-05-17T08:37:55Z
merged: 2022-05-17T08:37:55Z
base_branch: master
head_branch: refresh-system-detail
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/797
---

# Pull Request #797: fix(SPM-1492): system details page refreshs after patch-set change

**Author**: @mkholjuraev
**Created**: May 12, 2022 at 09:45 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `refresh-system-detail`

## Description

Resolves: https://issues.redhat.com/browse/SPM-1492. I am not able to do API request other than the type GET, theoratically just adding `patchSetState.shouldRefresh` to the the hook dependency should do the trick, most probably. Thus, please run this PR locally to make sure it does the fix.

To reproduce:
1. Assign system to a patch set via action on system detail page
2. The page continues to display old patch set name a needs to be refreshed to display new patch set name

---

## Discussion

### Comment by @codecov-commenter on May 12, 2022 at 11:55 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#797](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7f1897a) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/03ef9f193bf4ec4e2dc500ed03b32e10bf967fec?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (03ef9f1) will **decrease** coverage by `0.35%`.
> The diff coverage is `7.14%`.

> :exclamation: Current head 7f1897a differs from pull request most recent head f9ff150. Consider uploading reports for the commit f9ff150 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #797      +/-   ##
==========================================
- Coverage   70.83%   70.47%   -0.36%     
==========================================
  Files         101      101              
  Lines        2393     2405      +12     
  Branches      618      623       +5     
==========================================
  Hits         1695     1695              
- Misses        631      641      +10     
- Partials       67       69       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <0.00%> (-0.31%)` | :arrow_down: |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `76.92% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `26.92% <7.69%> (-19.75%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [03ef9f1...f9ff150](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/797?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 13, 2022 at 11:44 AM UTC

@MichalSajdik thanks for having a look. Yeap, I forgot to checkout to a new branch from master branch. Issue is resolved.

### Comment by @mkholjuraev on May 17, 2022 at 08:37 AM UTC

@MichalSajdik thanks!

### Comment by @jiridostal on May 17, 2022 at 08:59 AM UTC

:tada: This PR is included in version 1.47.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @MichalSajdik - Commented on May 13, 2022 at 09:30 AM UTC

I noticed that in https://github.com/RedHatInsights/patchman-ui/pull/798
you have similiar changes with different commit id's. I think it should be consistent

### Review by @MichalSajdik - Changes Requested on May 13, 2022 at 09:31 AM UTC

the comment above

### Review by @MichalSajdik - Approved on May 17, 2022 at 08:30 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/797*
