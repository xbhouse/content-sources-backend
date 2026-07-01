---
type: pull_request
number: 640
title: "fix(SystemtatusBar): regressions are resolved"
state: merged
author: mkholjuraev
created: 2021-08-25T11:44:31Z
updated: 2021-08-25T13:11:06Z
closed: 2021-08-25T12:55:41Z
merged: 2021-08-25T12:55:41Z
base_branch: master
head_branch: advisories
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/640
---

# Pull Request #640: fix(SystemtatusBar): regressions are resolved

**Author**: @mkholjuraev
**Created**: August 25, 2021 at 11:44 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisories`

## Description

Regressions are resolved:
1. Spinner is changed to Skeleton on load
2. the space between status cards are made 24 px
3. Icons are made 'md' size
4. status filter wording are made consistent with the mockups
5. double filters chips are fixed. now stale filter chips are only one fresh and one stale
6. status filters are sharable using URL params

https://user-images.githubusercontent.com/59481011/130784860-014d48d2-53cc-4b17-a604-35a795b7f587.mp4



---

## Discussion

### Comment by @codecov-commenter on August 25, 2021 at 11:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#640](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (bb9125a) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3685f0f9f998d5b492142d6a6a277217750f82f0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3685f0f) will **decrease** coverage by `0.10%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #640      +/-   ##
==========================================
- Coverage   52.11%   52.01%   -0.11%     
==========================================
  Files          81       81              
  Lines        1961     1963       +2     
  Branches      429      433       +4     
==========================================
- Hits         1022     1021       -1     
- Misses        855      857       +2     
- Partials       84       85       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/Filters/SystemStaleFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1N5c3RlbVN0YWxlRmlsdGVyLmpz) | `0.00% <0.00%> (ø)` | |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.13% <ø> (-0.11%)` | :arrow_down: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `53.19% <ø> (+0.55%)` | :arrow_up: |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| [src/store/Reducers/SystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNTdG9yZS5qcw==) | `42.85% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [3685f0f...bb9125a](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/640?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 25, 2021 at 12:55 PM UTC

LGTM!


### Comment by @jiridostal on August 25, 2021 at 01:11 PM UTC

:tada: This PR is included in version 1.31.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.31.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.31.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/640*
