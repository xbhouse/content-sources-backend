---
type: pull_request
number: 618
title: "fix(Reducers): SPM-1073 clear details page reducers"
state: merged
author: mkholjuraev
created: 2021-08-05T14:30:08Z
updated: 2022-05-17T08:50:23Z
closed: 2021-08-05T14:56:04Z
merged: 2021-08-05T14:56:04Z
base_branch: master
head_branch: clear-details-pages
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/618
---

# Pull Request #618: fix(Reducers): SPM-1073 clear details page reducers

**Author**: @mkholjuraev
**Created**: August 05, 2021 at 02:30 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `clear-details-pages`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on August 05, 2021 at 02:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#618](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (da581f4) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/c7a3f08a905899161720461c9724a8b75c8bd4be?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c7a3f08) will **decrease** coverage by `0.09%`.
> The diff coverage is `33.33%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #618      +/-   ##
==========================================
- Coverage   52.42%   52.33%   -0.10%     
==========================================
  Files          79       79              
  Lines        1894     1905      +11     
  Branches      413      415       +2     
==========================================
+ Hits          993      997       +4     
- Misses        820      827       +7     
  Partials       81       81              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/store/Reducers/AdvisorySystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5U3lzdGVtc1N0b3JlLmpz) | `42.85% <0.00%> (-7.15%)` | :arrow_down: |
| [src/store/Reducers/PackageSystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhY2thZ2VTeXN0ZW1zU3RvcmUuanM=) | `46.15% <0.00%> (-8.40%)` | :arrow_down: |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `67.18% <100.00%> (+1.05%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [c7a3f08...da581f4](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/618?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 05, 2021 at 03:11 PM UTC

:tada: This PR is included in version 1.28.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.28.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.28.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/618*
