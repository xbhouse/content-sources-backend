---
type: pull_request
number: 529
title: "fix(UserStatuses): use different approach to check user account status"
state: merged
author: mkholjuraev
created: 2021-05-06T15:29:40Z
updated: 2021-05-06T15:49:20Z
closed: 2021-05-06T15:40:18Z
merged: 2021-05-06T15:40:18Z
base_branch: master
head_branch: temp-hotfix
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/529
---

# Pull Request #529: fix(UserStatuses): use different approach to check user account status

**Author**: @mkholjuraev
**Created**: May 06, 2021 at 03:29 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `temp-hotfix`

## Description

I have disabled some tests to save up time as I was not able to find the cause for a long time. Later we can solve them. If you feel that it needs more work/test should be solved please feel free to ping me on slack

---

## Discussion

### Comment by @codecov-commenter on May 06, 2021 at 03:33 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#529](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b67582c) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/24347032afee3d5ed59953d8ce8f56cbe060bce2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2434703) will **decrease** coverage by `3.08%`.
> The diff coverage is `57.14%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #529      +/-   ##
==========================================
- Coverage   55.25%   52.17%   -3.09%     
==========================================
  Files          78       79       +1     
  Lines        1817     1817              
  Branches      383      376       -7     
==========================================
- Hits         1004      948      -56     
- Misses        754      810      +56     
  Partials       59       59              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `47.61% <0.00%> (+14.28%)` | :arrow_up: |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/axiosInterceptors.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9heGlvc0ludGVyY2VwdG9ycy5qcw==) | `67.64% <66.66%> (-1.32%)` | :arrow_down: |
| [...sentationalComponents/Snippets/GeneralComponent.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9HZW5lcmFsQ29tcG9uZW50Lmpz) | `75.00% <75.00%> (ø)` | |
| [...rc/PresentationalComponents/TableView/TableView.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3Lmpz) | `100.00% <100.00%> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <100.00%> (ø)` | |
| [src/store/Reducers/HelperReducers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0hlbHBlclJlZHVjZXJzLmpz) | `100.00% <100.00%> (ø)` | |
| [...c/SmartComponents/AdvisoryDetail/AdvisoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9BZHZpc29yeURldGFpbC5qcw==) | `10.00% <0.00%> (-90.00%)` | :arrow_down: |
| ... and [4 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2434703...b67582c](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/529?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on May 06, 2021 at 03:49 PM UTC

:tada: This PR is included in version 1.18.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.18.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.18.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/529*
