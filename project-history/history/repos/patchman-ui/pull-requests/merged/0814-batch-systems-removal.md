---
type: pull_request
number: 814
title: "Batch systems removal"
state: merged
author: mkholjuraev
created: 2022-06-03T11:05:02Z
updated: 2024-04-03T09:23:26Z
closed: 2022-06-13T10:52:27Z
merged: 2022-06-13T10:52:27Z
base_branch: master
head_branch: batch-systems-removal
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/814
---

# Pull Request #814: Batch systems removal

**Author**: @mkholjuraev
**Created**: June 03, 2022 at 11:05 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `batch-systems-removal`

## Description

This PR is to resolve https://issues.redhat.com/browse/SPM-1488 by enabling users remove one or more systems from one or more patch sets. The action can be taken through the 'Remove from Patch Set' action in the table toolbas and the table row actions.

Expectations:

1. When you try to remove a system without patch set assigned, 'Remove from Patch Set' row action should be disabled.
2. When you try to remove a single system without patch set assigned, a modal should pop with proper message indicating there is no effect and disabled Remove button.
3. When you select a single system with a patch set assigned, modal should open and you should be able to remove.
4. If you select systems with both patch set assigned and not assigned, a modal with the message about  _Do you want to remove the SomeNumber selected system from assigned Patch sets?_ and _There is 1 system you are trying to remove that is not assigned to any existing Patch Set. This action will not affect it._

---

## Discussion

### Comment by @mkholjuraev on June 11, 2022 at 07:13 PM UTC

@gabipodolnikova thanks for the review. All suggestions are improvement and I have applied them. The refresh bug is also fixed hopefully.

### Comment by @codecov-commenter on June 11, 2022 at 07:19 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#814](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (57bee00) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b5aa43743497516864bf50c2eb2e74a2c01d21ab?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b5aa437) will **increase** coverage by `0.66%`.
> The diff coverage is `87.27%`.

```diff
@@            Coverage Diff             @@
##           master     #814      +/-   ##
==========================================
+ Coverage   70.99%   71.65%   +0.66%     
==========================================
  Files         101      103       +2     
  Lines        2444     2480      +36     
  Branches      632      640       +8     
==========================================
+ Hits         1735     1777      +42     
+ Misses        642      640       -2     
+ Partials       67       63       -4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PatchSetWizard/PatchSetWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9QYXRjaFNldFdpemFyZC5qcw==) | `7.69% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `68.75% <58.33%> (-0.99%)` | :arrow_down: |
| [src/SmartComponents/Modals/UnassignSystemsModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvVW5hc3NpZ25TeXN0ZW1zTW9kYWwuanM=) | `95.23% <94.11%> (-4.77%)` | :arrow_down: |
| [src/SmartComponents/Modals/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvSGVscGVycy5qcw==) | `100.00% <100.00%> (ø)` | |
| [...c/SmartComponents/Modals/useUnassignSystemsHook.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvdXNlVW5hc3NpZ25TeXN0ZW1zSG9vay5qcw==) | `100.00% <100.00%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `83.87% <100.00%> (+0.53%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `84.00% <100.00%> (+2.18%)` | :arrow_up: |
| [src/Utilities/unitTestingUtilities.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91bml0VGVzdGluZ1V0aWxpdGllcy5qcw==) | `29.72% <100.00%> (+10.97%)` | :arrow_up: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `55.45% <0.00%> (+0.90%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `79.47% <0.00%> (+2.64%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [b5aa437...57bee00](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/814?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 13, 2022 at 11:09 AM UTC

:tada: This PR is included in version 1.48.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gabipodolnikova - Changes Requested on June 10, 2022 at 01:23 PM UTC

Just two minor things.
And also leaving what I already sent you here too:
When removing patch set from the toolbar it does not refresh the page.
When assigning either from row actions or a toolbar it also does not refresh the page.

### Review by @mkholjuraev - Commented on June 11, 2022 at 07:13 PM UTC

### Review by @mkholjuraev - Commented on June 11, 2022 at 07:14 PM UTC

### Review by @gabipodolnikova - Approved on June 13, 2022 at 10:50 AM UTC

Looks good, thank you! :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/814*
