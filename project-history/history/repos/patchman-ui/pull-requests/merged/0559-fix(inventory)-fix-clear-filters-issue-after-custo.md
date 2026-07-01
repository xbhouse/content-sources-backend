---
type: pull_request
number: 559
title: "fix(Inventory): fix clear filters issue after custom inventory fetch"
state: merged
author: mkholjuraev
created: 2021-06-02T10:53:13Z
updated: 2021-08-09T06:56:59Z
closed: 2021-06-02T12:29:18Z
merged: 2021-06-02T12:29:18Z
base_branch: master
head_branch: clear-filters
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/559
---

# Pull Request #559: fix(Inventory): fix clear filters issue after custom inventory fetch

**Author**: @mkholjuraev
**Created**: June 02, 2021 at 10:53 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `clear-filters`

## Description

Systems, advisory details page breaks after clicking 'Clear filters' button due to initial undefined data 

---

## Discussion

### Comment by @codecov-commenter on June 02, 2021 at 10:57 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#559](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d00048d) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9edd005c6e2c579cf58a149dbf8f25a547bc8ccf?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (9edd005) will **decrease** coverage by `0.02%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #559      +/-   ##
==========================================
- Coverage   54.16%   54.13%   -0.03%     
==========================================
  Files          78       78              
  Lines        1837     1838       +1     
  Branches      392      394       +2     
==========================================
  Hits          995      995              
  Misses        780      780              
- Partials       62       63       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...resentationalComponents/Snippets/AdvisoriesIcon.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9BZHZpc29yaWVzSWNvbi5qcw==) | `50.00% <0.00%> (-16.67%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [9edd005...d00048d](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/559?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 02, 2021 at 12:37 PM UTC

:tada: This PR is included in version 1.20.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.20.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.20.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/559*
