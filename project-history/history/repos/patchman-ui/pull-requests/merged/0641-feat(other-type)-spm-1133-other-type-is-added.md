---
type: pull_request
number: 641
title: "feat(Other type): SPM-1133 \"Other\" type is added"
state: merged
author: mkholjuraev
created: 2021-08-25T13:16:26Z
updated: 2022-05-17T08:49:43Z
closed: 2021-08-26T08:03:13Z
merged: 2021-08-26T08:03:13Z
base_branch: master
head_branch: other-type
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/641
---

# Pull Request #641: feat(Other type): SPM-1133 "Other" type is added

**Author**: @mkholjuraev
**Created**: August 25, 2021 at 01:16 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `other-type`

## Description

Other type is added into the UI. Pages changed:

Advisories page: 

![otherAdvisory](https://user-images.githubusercontent.com/59481011/130797277-95be8da5-9c98-4c31-8f10-e7bd44796c7c.png)

Advisory details: 

![otherAdvisoryDetail](https://user-images.githubusercontent.com/59481011/130797333-e4ff874b-2545-4d3c-acdd-05000deed697.png)

Systems page: 

![otherSystems](https://user-images.githubusercontent.com/59481011/130797371-dac9d994-89ec-4f94-98dc-c0c58b27b8eb.png)


---

## Discussion

### Comment by @codecov-commenter on August 25, 2021 at 01:24 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#641](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (50c403e) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3685f0f9f998d5b492142d6a6a277217750f82f0?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3685f0f) will **decrease** coverage by `0.00%`.
> The diff coverage is `46.15%`.

> :exclamation: Current head 50c403e differs from pull request most recent head 59a907c. Consider uploading reports for the commit 59a907c to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #641      +/-   ##
==========================================
- Coverage   52.11%   52.10%   -0.01%     
==========================================
  Files          81       81              
  Lines        1961     1967       +6     
  Branches      429      436       +7     
==========================================
+ Hits         1022     1025       +3     
- Misses        855      857       +2     
- Partials       84       85       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/Filters/SystemStaleFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1N5c3RlbVN0YWxlRmlsdGVyLmpz) | `0.00% <0.00%> (ø)` | |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `53.19% <ø> (+0.55%)` | :arrow_up: |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| [src/store/Reducers/SystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbXNTdG9yZS5qcw==) | `42.85% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `72.22% <100.00%> (+1.06%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.33% <100.00%> (+0.10%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [3685f0f...59a907c](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/641?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 26, 2021 at 08:03 AM UTC

Looks good! merging

### Comment by @jiridostal on August 26, 2021 at 08:18 AM UTC

:tada: This PR is included in version 1.32.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.32.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.32.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/641*
