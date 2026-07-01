---
type: pull_request
number: 630
title: "Don't display advisory link for EPEL on advisories page"
state: merged
author: jiridostal
created: 2021-08-17T09:47:57Z
updated: 2022-08-02T08:42:48Z
closed: 2021-08-19T09:48:30Z
merged: 2021-08-19T09:48:30Z
base_branch: master
head_branch: epel-advisories
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/630
---

# Pull Request #630: Don't display advisory link for EPEL on advisories page

**Author**: @jiridostal
**Created**: August 17, 2021 at 09:47 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `epel-advisories`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on August 17, 2021 at 09:54 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#630](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8dd4486) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/812186a5012640cc79f336f1af54159864f2401c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (812186a) will **decrease** coverage by `0.00%`.
> The diff coverage is `54.54%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #630      +/-   ##
==========================================
- Coverage   52.81%   52.80%   -0.01%     
==========================================
  Files          78       78              
  Lines        1920     1922       +2     
  Branches      421      422       +1     
==========================================
+ Hits         1014     1015       +1     
- Misses        826      827       +1     
  Partials       80       80              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | `83.33% <ø> (-1.67%)` | :arrow_down: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `76.92% <50.00%> (-0.15%)` | :arrow_down: |
| [...tationalComponents/Snippets/DescriptionWithLink.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9EZXNjcmlwdGlvbldpdGhMaW5rLmpz) | `85.71% <66.66%> (+2.38%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [812186a...8dd4486](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/630?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on August 19, 2021 at 11:12 AM UTC

:tada: This PR is included in version 1.30.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.30.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.30.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 17, 2021 at 11:12 PM UTC

### Review by @mkholjuraev - Approved on August 19, 2021 at 08:57 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/630*
