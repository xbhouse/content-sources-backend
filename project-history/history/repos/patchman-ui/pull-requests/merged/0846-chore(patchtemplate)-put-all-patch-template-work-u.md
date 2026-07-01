---
type: pull_request
number: 846
title: "chore(PatchTemplate): put all patch template work under a feature flag and fix wizard title bug"
state: merged
author: mkholjuraev
created: 2022-07-19T08:44:19Z
updated: 2022-08-01T14:54:47Z
closed: 2022-07-21T11:07:42Z
merged: 2022-07-21T11:07:42Z
base_branch: master
head_branch: feature-flag
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/846
---

# Pull Request #846: chore(PatchTemplate): put all patch template work under a feature flag and fix wizard title bug

**Author**: @mkholjuraev
**Created**: July 19, 2022 at 08:44 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `feature-flag`

## Description

Introduces a feature flag again for all patch template work and fixes the wizard title bug.

Please verify:

1. When you click on the 'Templates' nav, the page gets redirected to the Advisories page(landing page).
2. Systems page has no 'Assign to patch set' button, Patch template column, and row actions do not include any action related to patch templates.
3. Inventory details header has no patch template name and no actions in the dropdown.
4. On The last step in the wizard after a patch template has been created/edited/assigned, the wizard title dynamically changes. Please refer video below:
[screencast-2022-07-19-132703.webm](https://user-images.githubusercontent.com/59481011/179709041-021c8902-e726-42a6-b081-2451d842e9e9.webm)




---

## Discussion

### Comment by @codecov-commenter on July 19, 2022 at 09:57 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#846](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (91b37b9) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d69d4cb11443ddee62a96f219cedd39e3d46224a?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d69d4cb) will **decrease** coverage by `0.04%`.
> The diff coverage is `45.45%`.

> :exclamation: Current head 91b37b9 differs from pull request most recent head 0463094. Consider uploading reports for the commit 0463094 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #846      +/-   ##
==========================================
- Coverage   72.03%   71.99%   -0.05%     
==========================================
  Files         108      108              
  Lines        2485     2492       +7     
  Branches      633      635       +2     
==========================================
+ Hits         1790     1794       +4     
- Misses        635      638       +3     
  Partials       60       60              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `28.94% <0.00%> (-2.49%)` | :arrow_down: |
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <ø> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `26.66% <66.66%> (+1.66%)` | :arrow_up: |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `83.87% <100.00%> (+1.11%)` | :arrow_up: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `81.66% <100.00%> (+1.66%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [d69d4cb...0463094](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/846?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 01, 2022 at 02:54 PM UTC

:tada: This PR is included in version 1.48.7 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.7)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.7)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on July 21, 2022 at 10:41 AM UTC

LGTM! Thank you @mkholjuraev!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/846*
