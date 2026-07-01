---
type: pull_request
number: 610
title: "fix(Clear Filters): SPM-1020 redundant API call is removed"
state: merged
author: mkholjuraev
created: 2021-07-29T23:07:22Z
updated: 2021-08-02T12:31:03Z
closed: 2021-08-02T12:07:55Z
merged: 2021-08-02T12:07:55Z
base_branch: master
head_branch: remove-filter
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/610
---

# Pull Request #610: fix(Clear Filters): SPM-1020 redundant API call is removed

**Author**: @mkholjuraev
**Created**: July 29, 2021 at 11:07 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-filter`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on July 29, 2021 at 11:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#610](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c87071a) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3a9269334fb567432faa5b3a868a7accfe9a4b37?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3a92693) will **not change** coverage.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #610   +/-   ##
=======================================
  Coverage   52.81%   52.81%           
=======================================
  Files          79       79           
  Lines        1880     1880           
  Branches      402      402           
=======================================
  Hits          993      993           
  Misses        807      807           
  Partials       80       80           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `0.00% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [3a92693...c87071a](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/610?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on July 30, 2021 at 09:45 AM UTC

Seems like there are some styling issues 
BEFORE:
![image](https://user-images.githubusercontent.com/6286045/127635202-d2c44a35-9591-4ff4-953e-d23173a6202d.png)


AFTER:
![image](https://user-images.githubusercontent.com/6286045/127635267-2700a2ff-38a6-431d-9377-53b9d727fb21.png)


### Comment by @jiridostal on August 02, 2021 at 12:30 PM UTC

:tada: This PR is included in version 1.28.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.28.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.28.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/610*
