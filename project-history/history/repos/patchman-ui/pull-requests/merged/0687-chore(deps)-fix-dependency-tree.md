---
type: pull_request
number: 687
title: "chore(deps): fix dependency tree"
state: merged
author: mkholjuraev
created: 2021-11-11T13:13:20Z
updated: 2022-05-17T08:49:56Z
closed: 2021-11-12T08:11:36Z
merged: 2021-11-12T08:11:36Z
base_branch: master
head_branch: fix-deps
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/687
---

# Pull Request #687: chore(deps): fix dependency tree

**Author**: @mkholjuraev
**Created**: November 11, 2021 at 01:13 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-deps`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on November 11, 2021 at 07:32 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#687](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e54a4fd) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/769cbb3722c4415ce887cd103fc1e10b80e2fbf0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (769cbb3) will **decrease** coverage by `0.00%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #687      +/-   ##
==========================================
- Coverage   83.02%   83.02%   -0.01%     
==========================================
  Files          83       83              
  Lines        1862     1879      +17     
  Branches      466      487      +21     
==========================================
+ Hits         1546     1560      +14     
- Misses        279      282       +3     
  Partials       37       37              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/store/Actions/Actions.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvbnMvQWN0aW9ucy5qcw==) | `88.52% <ø> (+3.90%)` | :arrow_up: |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `93.47% <0.00%> (-2.08%)` | :arrow_down: |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `82.81% <0.00%> (-1.94%)` | :arrow_down: |
| [src/Utilities/axiosInterceptors.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9heGlvc0ludGVyY2VwdG9ycy5qcw==) | `58.82% <0.00%> (-1.79%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `86.32% <0.00%> (-0.52%)` | :arrow_down: |
| [src/SmartComponents/Packages/Packages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlcy9QYWNrYWdlcy5qcw==) | `100.00% <0.00%> (ø)` | |
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `98.03% <0.00%> (+0.03%)` | :arrow_up: |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `95.00% <0.00%> (+0.08%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `88.11% <0.00%> (+0.25%)` | :arrow_up: |
| [...resentationalComponents/Filters/OsVersionFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL09zVmVyc2lvbkZpbHRlci5qcw==) | `90.90% <0.00%> (+0.43%)` | :arrow_up: |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [769cbb3...e54a4fd](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/687?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on November 12, 2021 at 08:27 AM UTC

:tada: This PR is included in version 1.37.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.37.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.37.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/687*
