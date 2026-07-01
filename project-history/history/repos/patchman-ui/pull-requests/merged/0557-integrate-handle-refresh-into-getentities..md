---
type: pull_request
number: 557
title: "Integrate handle refresh into getEntities."
state: merged
author: Hyperkid123
created: 2021-06-02T09:30:06Z
updated: 2021-06-02T12:37:55Z
closed: 2021-06-02T10:15:01Z
merged: 2021-06-02T10:15:01Z
base_branch: master
head_branch: debug-inventory
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/557
---

# Pull Request #557: Integrate handle refresh into getEntities.

**Author**: @Hyperkid123
**Created**: June 02, 2021 at 09:30 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `debug-inventory`

## Description

cc @mkholjuraev 

This PR fixes an issue when the inventory table hangs in a loading state.

The problem was caused by duplicate API calls by the `getEntities` and `onRefresh` callbacks. Where possible, I've integrated the onRefresh into the `getEntities` callback which will prevent double calls and hanging loading state in the entities reducer.

@mkholjuraev Please double-check that nothing broke with this change.

---

## Discussion

### Comment by @codecov-commenter on June 02, 2021 at 09:42 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#557](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d4a382f) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/43b4525bdf1904ca00ae5d212536f2174e03bd01?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (43b4525) will **decrease** coverage by `0.03%`.
> The diff coverage is `52.94%`.

> :exclamation: Current head d4a382f differs from pull request most recent head c09828c. Consider uploading reports for the commit c09828c to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #557      +/-   ##
==========================================
- Coverage   54.19%   54.16%   -0.04%     
==========================================
  Files          78       78              
  Lines        1834     1837       +3     
  Branches      392      392              
==========================================
+ Hits          994      995       +1     
- Misses        778      780       +2     
  Partials       62       62              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `85.63% <100.00%> (+0.55%)` | :arrow_up: |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `94.23% <100.00%> (-0.32%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [43b4525...c09828c](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/557?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 02, 2021 at 12:37 PM UTC

:tada: This PR is included in version 1.20.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.20.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.20.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on June 02, 2021 at 10:14 AM UTC

Thanks for the PR, I have faced a small issue with 'Clear filters'. But it is an easy to fix I believe. 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/557*
