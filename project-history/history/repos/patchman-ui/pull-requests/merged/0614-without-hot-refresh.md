---
type: pull_request
number: 614
title: "Without hot refresh"
state: merged
author: mkholjuraev
created: 2021-08-05T10:20:13Z
updated: 2022-05-17T08:50:25Z
closed: 2021-08-05T10:28:31Z
merged: 2021-08-05T10:28:31Z
base_branch: master
head_branch: without-hot-refresh
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/614
---

# Pull Request #614: Without hot refresh

**Author**: @mkholjuraev
**Created**: August 05, 2021 at 10:20 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `without-hot-refresh`

## Description

The PR includes 2 fixes for the release: 
1. Broken page after reload because of difference between how installed_evra is decoded from URL params. installed_evra gets decoded as an array, while version filter handles installed_evra as a string. Now I have put a condition to handle the type accordingly.
2. System advisories and System packages were not persisting parameters as we were clearing the store on unmount. I have removed clearing the store.

---

## Discussion

### Comment by @codecov-commenter on August 05, 2021 at 10:27 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#614](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (08872fa) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9020387d822a9f8f92222c58a47d6de9660ef99c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9020387) will **decrease** coverage by `0.23%`.
> The diff coverage is `0.00%`.

> :exclamation: Current head 08872fa differs from pull request most recent head 984b757. Consider uploading reports for the commit 984b757 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #614      +/-   ##
==========================================
- Coverage   52.56%   52.33%   -0.24%     
==========================================
  Files          79       79              
  Lines        1889     1886       -3     
  Branches      407      411       +4     
==========================================
- Hits          993      987       -6     
- Misses        816      818       +2     
- Partials       80       81       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [.../PresentationalComponents/Filters/VersionFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1ZlcnNpb25GaWx0ZXIuanM=) | `0.00% <0.00%> (ø)` | |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `95.34% <ø> (-0.21%)` | :arrow_down: |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `100.00% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `76.95% <0.00%> (-0.36%)` | :arrow_down: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `82.73% <0.00%> (-1.21%)` | :arrow_down: |
| [src/store/Reducers/SystemAdvisoryListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbUFkdmlzb3J5TGlzdFN0b3JlLmpz) | `78.94% <ø> (-2.01%)` | :arrow_down: |
| [src/store/Reducers/SystemPackageListStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbVBhY2thZ2VMaXN0U3RvcmUuanM=) | `33.33% <ø> (+3.33%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9020387...984b757](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/614?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 05, 2021 at 10:43 AM UTC

:tada: This PR is included in version 1.28.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.28.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.28.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/614*
