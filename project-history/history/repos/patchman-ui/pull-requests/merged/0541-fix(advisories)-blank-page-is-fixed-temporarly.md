---
type: pull_request
number: 541
title: "fix(Advisories): blank page is fixed temporarly"
state: merged
author: mkholjuraev
created: 2021-05-20T10:06:29Z
updated: 2021-05-20T10:27:36Z
closed: 2021-05-20T10:18:38Z
merged: 2021-05-20T10:18:38Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/541
---

# Pull Request #541: fix(Advisories): blank page is fixed temporarly

**Author**: @mkholjuraev
**Created**: May 20, 2021 at 10:06 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on May 20, 2021 at 10:10 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#541](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (623e6eb) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/02269ea6cffac1e317b4be358b1c2656f72fc395?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (02269ea) will **decrease** coverage by `0.83%`.
> The diff coverage is `12.50%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #541      +/-   ##
==========================================
- Coverage   55.09%   54.26%   -0.84%     
==========================================
  Files          78       78              
  Lines        1826     1830       +4     
  Branches      388      392       +4     
==========================================
- Hits         1006      993      -13     
- Misses        759      775      +16     
- Partials       61       62       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/PresentationalComponents/Snippets/NoAccess.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9Ob0FjY2Vzcy5qcw==) | `25.00% <0.00%> (-25.00%)` | :arrow_down: |
| [src/Routes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | `32.25% <0.00%> (-1.08%)` | :arrow_down: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `70.37% <0.00%> (ø)` | |
| [src/store/Reducers/HelperReducers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0hlbHBlclJlZHVjZXJzLmpz) | `84.61% <50.00%> (-15.39%)` | :arrow_down: |
| [src/store/Reducers/AdvisoryDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5RGV0YWlsU3RvcmUuanM=) | `85.71% <0.00%> (-14.29%)` | :arrow_down: |
| [src/store/Reducers/SystemsListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNMaXN0U3RvcmUuanM=) | `64.70% <0.00%> (-11.77%)` | :arrow_down: |
| [src/store/Reducers/AdvisoryListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5TGlzdFN0b3JlLmpz) | `78.94% <0.00%> (-10.53%)` | :arrow_down: |
| [src/store/Reducers/AdvisorySystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5U3lzdGVtc1N0b3JlLmpz) | `78.94% <0.00%> (-10.53%)` | :arrow_down: |
| [src/store/Reducers/SystemAdvisoryListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbUFkdmlzb3J5TGlzdFN0b3JlLmpz) | `80.95% <0.00%> (-9.53%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [02269ea...623e6eb](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/541?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on May 20, 2021 at 10:27 AM UTC

:tada: This PR is included in version 1.19.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.19.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.19.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/541*
