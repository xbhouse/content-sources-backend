---
type: pull_request
number: 631
title: "SPM-1113 Fix inconsistent spacing on toolbars/tables"
state: merged
author: jiridostal
created: 2021-08-17T10:35:17Z
updated: 2022-08-02T08:34:34Z
closed: 2021-08-19T08:53:31Z
merged: 2021-08-19T08:53:31Z
base_branch: master
head_branch: spacing-fixed
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/631
---

# Pull Request #631: SPM-1113 Fix inconsistent spacing on toolbars/tables

**Author**: @jiridostal
**Created**: August 17, 2021 at 10:35 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `spacing-fixed`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on August 17, 2021 at 10:51 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#631](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (622a14f) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/2cdc00f1f3b5105ba43917cbf12ca26b192854f3?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2cdc00f) will **increase** coverage by `0.08%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #631      +/-   ##
==========================================
+ Coverage   52.86%   52.94%   +0.08%     
==========================================
  Files          78       78              
  Lines        1922     1919       -3     
  Branches      421      418       -3     
==========================================
  Hits         1016     1016              
+ Misses        826      823       -3     
  Partials       80       80              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2cdc00f...622a14f](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/631?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 19, 2021 at 08:45 AM UTC

I think it's a plugin called 'organize-imports' :-)

### Comment by @jiridostal on August 19, 2021 at 11:12 AM UTC

:tada: This PR is included in version 1.30.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.30.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.30.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 17, 2021 at 11:22 PM UTC

I am wondering if you have a special prettifier that sorts file includes alphabetically. All the imports are sorted and it is good.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/631*
