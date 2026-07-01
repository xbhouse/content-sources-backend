---
type: pull_request
number: 763
title: "Fix systems select"
state: merged
author: mkholjuraev
created: 2022-04-06T11:27:22Z
updated: 2022-05-17T08:50:47Z
closed: 2022-04-14T13:30:39Z
merged: 2022-04-14T13:30:39Z
base_branch: master
head_branch: fix-systems-select
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/763
---

# Pull Request #763: Fix systems select

**Author**: @mkholjuraev
**Created**: April 06, 2022 at 11:27 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-systems-select`

## Description

Resolves: https://issues.redhat.com/browse/SPM-1394.

It contains a commit to fix the bulk select bug and another commit to adjust test cases and test improvements

---

## Discussion

### Comment by @codecov-commenter on April 06, 2022 at 11:32 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#763](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e2d9dd5) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/6684e006a8332c8af5c7105aff40d2e1e39e0fe5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (6684e00) will **decrease** coverage by `0.08%`.
> The diff coverage is `100.00%`.

```diff
@@            Coverage Diff             @@
##           master     #763      +/-   ##
==========================================
- Coverage   70.82%   70.74%   -0.09%     
==========================================
  Files         101      101              
  Lines        2389     2389              
  Branches      617      617              
==========================================
- Hits         1692     1690       -2     
  Misses        632      632              
- Partials       65       67       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `69.44% <100.00%> (+2.77%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `75.34% <0.00%> (-1.37%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [6684e00...e2d9dd5](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/763?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on April 14, 2022 at 01:45 PM UTC

:tada: This PR is included in version 1.45.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.45.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.45.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @MichalSajdik - Approved on April 14, 2022 at 01:30 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/763*
