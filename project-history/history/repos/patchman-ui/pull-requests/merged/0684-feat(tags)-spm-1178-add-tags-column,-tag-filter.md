---
type: pull_request
number: 684
title: "feat(Tags): SPM-1178 add tags column, tag filter"
state: merged
author: mkholjuraev
created: 2021-11-02T09:58:27Z
updated: 2022-05-17T08:49:57Z
closed: 2021-11-29T12:18:55Z
merged: 2021-11-29T12:18:55Z
base_branch: master
head_branch: tag-filter
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/684
---

# Pull Request #684: feat(Tags): SPM-1178 add tags column, tag filter

**Author**: @mkholjuraev
**Created**: November 02, 2021 at 09:58 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `tag-filter`

## Description

*No description provided*

---

## Discussion

### Comment by @mkholjuraev on November 02, 2021 at 09:07 PM UTC

The global tag filter and local tag filter in the inventory table are unified, giving the privilege to the local tag.

### Comment by @codecov-commenter on November 24, 2021 at 11:59 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#684](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fb8c83c) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/30d2cd7e04d4b27150e8f42d2ebcc17e2041e320?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (30d2cd7) will **decrease** coverage by `0.83%`.
> The diff coverage is `32.72%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #684      +/-   ##
==========================================
- Coverage   83.02%   82.18%   -0.84%     
==========================================
  Files          83       84       +1     
  Lines        1879     1914      +35     
  Branches      487      499      +12     
==========================================
+ Hits         1560     1573      +13     
- Misses        282      300      +18     
- Partials       37       41       +4     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/App.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/SystemsListAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | `90.90% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `82.81% <ø> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `86.89% <0.00%> (-1.22%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `80.31% <10.00%> (-6.01%)` | :arrow_down: |
| [src/store/Reducers/GlobalFilterStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0dsb2JhbEZpbHRlclN0b3JlLmpz) | `40.00% <40.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `80.70% <50.00%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `95.34% <100.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `95.00% <100.00%> (ø)` | |
| [src/store/ActionTypes.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL0FjdGlvblR5cGVzLmpz) | `100.00% <100.00%> (ø)` | |
| ... and [2 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [30d2cd7...fb8c83c](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/684?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on November 24, 2021 at 01:09 PM UTC

<img width="1145" alt="obrazek" src="https://user-images.githubusercontent.com/6286045/143244274-3b433d48-fd36-4f1f-be80-0cc25086059e.png">

I see the Tags column is a bit misaligned. Btw is change of the order of columns intended? Also, OS columns is visible twice

### Comment by @jiridostal on November 24, 2021 at 01:13 PM UTC

<img width="1153" alt="obrazek" src="https://user-images.githubusercontent.com/6286045/143244949-f09f9ee0-3a32-4132-84df-4895791d5bdd.png">

Here I see there is only one Tag available per row, but when I click it I see 6 tags

### Comment by @jiridostal on November 29, 2021 at 12:34 PM UTC

:tada: This PR is included in version 1.38.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.38.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.38.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/684*
