---
type: pull_request
number: 772
title: "Integrate patch set"
state: merged
author: mkholjuraev
created: 2022-04-12T09:44:11Z
updated: 2022-04-12T10:51:17Z
closed: 2022-04-12T10:36:19Z
merged: 2022-04-12T10:36:19Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/772
---

# Pull Request #772: Integrate patch set

**Author**: @mkholjuraev
**Created**: April 12, 2022 at 09:44 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on April 12, 2022 at 09:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#772](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (eafad20) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/83d34923f9043e20fd03ad342efa50aec0cb38cf?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (83d3492) will **decrease** coverage by `0.06%`.
> The diff coverage is `52.63%`.

> :exclamation: Current head eafad20 differs from pull request most recent head 5bd6fd8. Consider uploading reports for the commit 5bd6fd8 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #772      +/-   ##
==========================================
- Coverage   71.13%   71.06%   -0.07%     
==========================================
  Files         101      101              
  Lines        2352     2374      +22     
  Branches      604      609       +5     
==========================================
+ Hits         1673     1687      +14     
- Misses        617      624       +7     
- Partials       62       63       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | `4.54% <0.00%> (-0.34%)` | :arrow_down: |
| [src/store/Reducers/PatchSetsReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | `26.31% <0.00%> (-3.10%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `66.66% <50.00%> (ø)` | |
| [...s/PatchSetWizard/InputFields/SelectExistingSets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9TZWxlY3RFeGlzdGluZ1NldHMuanM=) | `89.74% <84.21%> (-6.42%)` | :arrow_down: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `75.00% <100.00%> (+12.50%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [83d3492...5bd6fd8](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/772?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on April 12, 2022 at 10:51 AM UTC

:tada: This PR is included in version 1.44.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.44.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.44.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/772*
