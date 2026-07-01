---
type: pull_request
number: 746
title: "Prod beta"
state: merged
author: mkholjuraev
created: 2022-03-10T11:01:09Z
updated: 2022-03-10T11:38:15Z
closed: 2022-03-10T11:38:15Z
merged: 2022-03-10T11:38:15Z
base_branch: prod-beta
head_branch: prod-beta
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/746
---

# Pull Request #746: Prod beta

**Author**: @mkholjuraev
**Created**: March 10, 2022 at 11:01 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `prod-beta` ← **Head**: `prod-beta`

## Description

Previous deployment failed doe to the Nodejs issue. I have cherry-picked fixes from master branch. 

---

## Discussion

### Comment by @codecov-commenter on March 10, 2022 at 11:30 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#746](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9146d1f) into [prod-beta](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/5faec371f72ae2580e53114e4d6c462aaee14d6c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (5faec37) will **decrease** coverage by `3.16%`.
> The diff coverage is `37.50%`.

> :exclamation: Current head 9146d1f differs from pull request most recent head 0dea269. Consider uploading reports for the commit 0dea269 to get more accurate results

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@              Coverage Diff              @@
##           prod-beta     #746      +/-   ##
=============================================
- Coverage      71.62%   68.45%   -3.17%     
=============================================
  Files             85       85              
  Lines           1896     1899       +3     
  Branches         491      491              
=============================================
- Hits            1358     1300      -58     
- Misses           484      539      +55     
- Partials          54       60       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `20.00% <0.00%> (-57.78%)` | :arrow_down: |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `72.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `53.68% <50.00%> (-2.70%)` | :arrow_down: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `83.56% <100.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `1.75% <0.00%> (-77.20%)` | :arrow_down: |
| [...c/SmartComponents/Remediation/RemediationWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbldpemFyZC5qcw==) | `66.66% <0.00%> (-33.34%)` | :arrow_down: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `76.92% <0.00%> (-3.85%)` | :arrow_down: |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `81.25% <0.00%> (+1.56%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [5faec37...0dea269](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/746?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/746*
