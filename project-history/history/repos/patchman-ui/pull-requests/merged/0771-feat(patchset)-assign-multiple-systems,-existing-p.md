---
type: pull_request
number: 771
title: "feat(PatchSet): assign multiple systems, existing patch-set selection\u2026"
state: merged
author: mkholjuraev
created: 2022-04-11T11:23:23Z
updated: 2022-04-11T12:28:14Z
closed: 2022-04-11T12:11:56Z
merged: 2022-04-11T12:11:56Z
base_branch: master
head_branch: integrate-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/771
---

# Pull Request #771: feat(PatchSet): assign multiple systems, existing patch-set selection…

**Author**: @mkholjuraev
**Created**: April 11, 2022 at 11:23 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `integrate-patch-set`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on April 11, 2022 at 11:42 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#771](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8fcb7e8) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/83d34923f9043e20fd03ad342efa50aec0cb38cf?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (83d3492) will **increase** coverage by `0.02%`.
> The diff coverage is `60.60%`.

> :exclamation: Current head 8fcb7e8 differs from pull request most recent head a69a1e1. Consider uploading reports for the commit a69a1e1 to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #771      +/-   ##
==========================================
+ Coverage   71.13%   71.15%   +0.02%     
==========================================
  Files         101      101              
  Lines        2352     2371      +19     
  Branches      604      608       +4     
==========================================
+ Hits         1673     1687      +14     
- Misses        617      622       +5     
  Partials       62       62              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/PatchSetsReducer.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhdGNoU2V0c1JlZHVjZXIuanM=) | `26.31% <0.00%> (-3.10%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `66.66% <50.00%> (ø)` | |
| [...s/PatchSetWizard/InputFields/SelectExistingSets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9JbnB1dEZpZWxkcy9TZWxlY3RFeGlzdGluZ1NldHMuanM=) | `89.74% <84.21%> (-6.42%)` | :arrow_down: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `75.00% <100.00%> (+12.50%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [83d3492...a69a1e1](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/771?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on April 11, 2022 at 12:28 PM UTC

:tada: This PR is included in version 1.44.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.44.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.44.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/771*
