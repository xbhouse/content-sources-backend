---
type: pull_request
number: 623
title: "feat(remediation): SPM-1050 redesign remediation button"
state: merged
author: jiridostal
created: 2021-08-10T08:58:36Z
updated: 2022-08-02T08:37:12Z
closed: 2021-08-11T13:25:26Z
merged: 2021-08-11T13:25:26Z
base_branch: master
head_branch: remediation-button
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/623
---

# Pull Request #623: feat(remediation): SPM-1050 redesign remediation button

**Author**: @jiridostal
**Created**: August 10, 2021 at 08:58 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remediation-button`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on August 10, 2021 at 01:47 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#623](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (35bdbb8) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/abf56509c0766d90796c4f22ee9707257113f091?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (abf5650) will **decrease** coverage by `0.02%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #623      +/-   ##
==========================================
- Coverage   52.38%   52.36%   -0.03%     
==========================================
  Files          79       79              
  Lines        1909     1906       -3     
  Branches      417      414       -3     
==========================================
- Hits         1000      998       -2     
+ Misses        828      827       -1     
  Partials       81       81              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...rtComponents/Remediation/PatchRemediationButton.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9QYXRjaFJlbWVkaWF0aW9uQnV0dG9uLmpz) | `100.00% <ø> (+16.66%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [abf5650...35bdbb8](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/623?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on August 10, 2021 at 11:23 PM UTC

LGTM! I will merge once I resolve the Jest testing failures problem.

### Comment by @jiridostal on August 12, 2021 at 07:49 AM UTC

:tada: This PR is included in version 1.29.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.29.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.29.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 10, 2021 at 11:22 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/623*
