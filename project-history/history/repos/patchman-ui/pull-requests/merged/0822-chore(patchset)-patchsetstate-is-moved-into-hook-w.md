---
type: pull_request
number: 822
title: "chore(PatchSet): patchSetState is moved into hook with some refactoring"
state: merged
author: mkholjuraev
created: 2022-06-16T12:06:08Z
updated: 2024-04-03T09:23:18Z
closed: 2022-06-29T10:57:43Z
merged: 2022-06-29T10:57:43Z
base_branch: master
head_branch: refactor-patch-set-state
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/822
---

# Pull Request #822: chore(PatchSet): patchSetState is moved into hook with some refactoring

**Author**: @mkholjuraev
**Created**: June 16, 2022 at 12:06 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `refactor-patch-set-state`

## Description

Refactoring the usage of state, components of PatchSetWizard and UnassignSystemsModal in SmartComponents.

- Wrapper component created to hold both PatchSetWizard and UnassignModal
- The state is moved into a shared hook with helper functions. 
- Test coverage for ^


---

## Discussion

### Comment by @codecov-commenter on June 20, 2022 at 08:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#822](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e73e2e2) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/bf4c632dbc877772f93f9b99ab152c4e7830af33?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bf4c632) will **decrease** coverage by `0.04%`.
> The diff coverage is `52.77%`.

```diff
@@            Coverage Diff             @@
##           master     #822      +/-   ##
==========================================
- Coverage   71.91%   71.87%   -0.05%     
==========================================
  Files         103      105       +2     
  Lines        2482     2482              
  Branches      642      643       +1     
==========================================
- Hits         1785     1784       -1     
  Misses        636      636              
- Partials       61       62       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/usePatchSetState.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VQYXRjaFNldFN0YXRlLmpz) | `58.33% <58.33%> (ø)` | |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `84.00% <66.66%> (+0.12%)` | :arrow_up: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `70.83% <66.66%> (+2.08%)` | :arrow_up: |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `84.00% <66.66%> (ø)` | |
| [...ionalComponents/PatchSetWrapper/PatchSetWrapper.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9QYXRjaFNldFdyYXBwZXIvUGF0Y2hTZXRXcmFwcGVyLmpz) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/axiosInterceptors.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9heGlvc0ludGVyY2VwdG9ycy5qcw==) | `50.00% <0.00%> (-8.83%)` | :arrow_down: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `54.62% <0.00%> (-0.83%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [bf4c632...e73e2e2](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/822?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 30, 2022 at 01:38 PM UTC

:tada: This PR is included in version 1.48.5 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.5)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/822*
