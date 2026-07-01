---
type: pull_request
number: 804
title: "fix(SPM-1488): remove systems from patch sets"
state: merged
author: mkholjuraev
created: 2022-05-19T12:27:59Z
updated: 2024-04-03T09:21:48Z
closed: 2022-05-31T07:00:03Z
merged: 2022-05-31T07:00:03Z
base_branch: master
head_branch: remove-sets
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/804
---

# Pull Request #804: fix(SPM-1488): remove systems from patch sets

**Author**: @mkholjuraev
**Created**: May 19, 2022 at 12:27 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-sets`

## Description

Resolves: https://issues.redhat.com/browse/SPM-1488. I am not able to make any other API requests except GET. Thus please make sure that API handles request nicely.

There are 2 use cases where users can remove system(s) from patch sets.

To reproduce from table row actions:
1. Go to Systems page
2. Click on 'Remove from patch set' row action of a system which has a set assigned.
3. The confirmation with the system number should appear
4. When users clicks on Remove button on the modal, system should be unassigned from the patch set

 To reproduce from table toolbar dropdown:
 1. Go to Systems page
 2. Select as many systems as you want
3. Click on 'Remove from patch set' button from the tabletoolbars dropdown.
4. The confirmation with the system number should appear
5. When users clicks on Remove button on the modal, system should be unassigned from the patch set

---

## Discussion

### Comment by @codecov-commenter on May 20, 2022 at 12:00 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#804](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5c5123b) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/884560c8ce28bd190956982f81f0b2af698bc502?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (884560c) will **increase** coverage by `0.43%`.
> The diff coverage is `89.65%`.

```diff
@@            Coverage Diff             @@
##           master     #804      +/-   ##
==========================================
+ Coverage   70.25%   70.69%   +0.43%     
==========================================
  Files         101      102       +1     
  Lines        2414     2440      +26     
  Branches      624      628       +4     
==========================================
+ Hits         1696     1725      +29     
+ Misses        648      647       -1     
+ Partials       70       68       -2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `73.33% <50.00%> (-3.59%)` | :arrow_down: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `54.54% <50.00%> (-0.09%)` | :arrow_down: |
| [src/SmartComponents/Modals/UnassignSystemsModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9Nb2RhbHMvVW5hc3NpZ25TeXN0ZW1zTW9kYWwuanM=) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `70.66% <100.00%> (+1.22%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `78.94% <100.00%> (+2.47%)` | :arrow_up: |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | `25.92% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `75.66% <0.00%> (+2.00%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [884560c...5c5123b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/804?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on May 26, 2022 at 06:12 AM UTC

@MichalSajdik :+1: 

### Comment by @jiridostal on May 31, 2022 at 07:16 AM UTC

:tada: This PR is included in version 1.47.5 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.5)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @MichalSajdik - Approved on May 24, 2022 at 08:36 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/804*
