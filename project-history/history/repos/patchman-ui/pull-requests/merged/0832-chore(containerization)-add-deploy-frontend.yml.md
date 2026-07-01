---
type: pull_request
number: 832
title: "chore(containerization): add deploy frontend.yml"
state: merged
author: mkholjuraev
created: 2022-06-23T19:47:11Z
updated: 2024-04-03T09:23:20Z
closed: 2022-06-24T13:12:41Z
merged: 2022-06-24T13:12:41Z
base_branch: master
head_branch: deploy-yml
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/832
---

# Pull Request #832: chore(containerization): add deploy frontend.yml

**Author**: @mkholjuraev
**Created**: June 23, 2022 at 07:47 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `deploy-yml`

## Description

This PR is intended to create [deployment config file](https://github.com/RedHatInsights/patchman-ui/pull/832). Some of config info is taken from: https://github.com/RedHatInsights/cloud-services-config/blob/ci-beta/chrome/fed-modules.json#L291.



---

## Discussion

### Comment by @mkholjuraev on June 23, 2022 at 08:07 PM UTC

@BlakeHolifield thank you for super quck review. I will fix all issues.

### Comment by @mkholjuraev on June 23, 2022 at 08:44 PM UTC

@BlakeHolifield can you please have another look whenever you are free?

### Comment by @codecov-commenter on June 24, 2022 at 01:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#832](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3889f59) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/64361cb7a135a7b6d00a061efb8551e25af90747?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (64361cb) will **decrease** coverage by `0.10%`.
> The diff coverage is `74.41%`.

```diff
@@            Coverage Diff             @@
##           master     #832      +/-   ##
==========================================
- Coverage   71.77%   71.67%   -0.11%     
==========================================
  Files         103      103              
  Lines        2480     2485       +5     
  Branches      642      646       +4     
==========================================
+ Hits         1780     1781       +1     
- Misses        638      642       +4     
  Partials       62       62              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <ø> (ø)` | |
| [src/store/Reducers/SystemDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbURldGFpbFN0b3JlLmpz) | `83.33% <0.00%> (-16.67%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `68.35% <58.33%> (-0.40%)` | :arrow_down: |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `82.85% <83.33%> (-1.02%)` | :arrow_down: |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/Reducers/InventoryEntitiesReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0ludmVudG9yeUVudGl0aWVzUmVkdWNlci5qcw==) | `92.30% <100.00%> (-0.29%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [c93c43f...3889f59](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/832?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 30, 2022 at 01:38 PM UTC

:tada: This PR is included in version 1.48.5 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.5)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @BlakeHolifield - Changes Requested on June 23, 2022 at 07:54 PM UTC

A couple small fixes, but looking good overall. 

### Review by @mkholjuraev - Commented on June 23, 2022 at 08:43 PM UTC

### Review by @mkholjuraev - Commented on June 23, 2022 at 08:43 PM UTC

### Review by @BlakeHolifield - Commented on June 24, 2022 at 12:30 PM UTC

Just a question here. Almost done!

### Review by @mkholjuraev - Commented on June 24, 2022 at 12:40 PM UTC

### Review by @BlakeHolifield - Approved on June 24, 2022 at 12:47 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/832*
