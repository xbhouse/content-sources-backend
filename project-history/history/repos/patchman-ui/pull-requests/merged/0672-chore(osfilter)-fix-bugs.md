---
type: pull_request
number: 672
title: "chore(osFilter): fix bugs"
state: merged
author: mkholjuraev
created: 2021-10-15T13:06:35Z
updated: 2022-05-17T08:49:51Z
closed: 2021-10-18T07:59:02Z
merged: 2021-10-18T07:59:02Z
base_branch: master
head_branch: os-version-filter
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/672
---

# Pull Request #672: chore(osFilter): fix bugs

**Author**: @mkholjuraev
**Created**: October 15, 2021 at 01:06 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `os-version-filter`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on October 15, 2021 at 01:21 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#672](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a0fa151) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/41d452d63e4e06cbe135526f467980eba844656e?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (41d452d) will **decrease** coverage by `0.02%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #672      +/-   ##
==========================================
- Coverage   83.10%   83.08%   -0.03%     
==========================================
  Files          83       83              
  Lines        1853     1850       -3     
  Branches      456      458       +2     
==========================================
- Hits         1540     1537       -3     
  Misses        276      276              
  Partials       37       37              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/AdvisoryType/AdvisoryType.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVR5cGUvQWR2aXNvcnlUeXBlLmpz) | `100.00% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `87.11% <ø> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |
| [src/store/Reducers/AdvisorySystemsStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5U3lzdGVtc1N0b3JlLmpz) | `100.00% <ø> (ø)` | |
| [...resentationalComponents/Filters/OsVersionFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL09zVmVyc2lvbkZpbHRlci5qcw==) | `90.47% <100.00%> (-0.44%)` | :arrow_down: |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `95.34% <100.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [41d452d...a0fa151](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/672?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on October 20, 2021 at 09:48 AM UTC

:tada: This PR is included in version 1.35.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.35.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.35.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on October 18, 2021 at 07:58 AM UTC

Merging to make a UI release.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/672*
