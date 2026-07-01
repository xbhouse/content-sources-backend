---
type: pull_request
number: 778
title: "feat(SPM-1311): assign systems to a patch-set in System details page"
state: merged
author: mkholjuraev
created: 2022-04-19T22:10:31Z
updated: 2022-04-20T10:30:26Z
closed: 2022-04-20T10:12:26Z
merged: 2022-04-20T10:12:26Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/778
---

# Pull Request #778: feat(SPM-1311): assign systems to a patch-set in System details page

**Author**: @mkholjuraev
**Created**: April 19, 2022 at 10:10 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

This adds assign patch set action in System details header actions dropdown.

---

## Discussion

### Comment by @codecov-commenter on April 19, 2022 at 10:28 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#778](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (14163eb) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9c7891ef381e8744da58c80c29974c97ddc1ae34?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9c7891e) will **increase** coverage by `0.79%`.
> The diff coverage is `100.00%`.

```diff
@@            Coverage Diff             @@
##           master     #778      +/-   ##
==========================================
+ Coverage   70.74%   71.53%   +0.79%     
==========================================
  Files         101      101              
  Lines        2389     2396       +7     
  Branches      617      620       +3     
==========================================
+ Hits         1690     1714      +24     
+ Misses        632      615      -17     
  Partials       67       67              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `85.18% <100.00%> (+5.18%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `77.97% <0.00%> (+1.78%)` | :arrow_up: |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `53.33% <0.00%> (+6.66%)` | :arrow_up: |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `60.00% <0.00%> (+52.00%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9c7891e...14163eb](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/778?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on April 20, 2022 at 10:30 AM UTC

:tada: This PR is included in version 1.46.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.46.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.46.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/778*
