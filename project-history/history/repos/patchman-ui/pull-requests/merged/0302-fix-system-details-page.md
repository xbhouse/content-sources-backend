---
type: pull_request
number: 302
title: "Fix system details page"
state: merged
author: mkholjuraev
created: 2020-09-29T10:03:28Z
updated: 2021-08-09T06:55:14Z
closed: 2020-09-29T11:27:31Z
merged: 2020-09-29T11:27:31Z
base_branch: master
head_branch: fix-system-details-page
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/302
---

# Pull Request #302: Fix system details page

**Author**: @mkholjuraev
**Created**: September 29, 2020 at 10:03 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-system-details-page`

## Description

Fixes: [the infinite re - rendering issue](https://ansible.slack.com/archives/CPS5FH07Q/p1601014601017800?thread_ts=1601014582.017700&cid=CPS5FH07Q).

I noticed that SystemAdvisories component is rendering infinite times. I believe that there was some changes in dependencies and all setStates inside fetchInventory function is causing race condition. I deleted the last setInventoryWrapper and it solved the issue. But I am not sure if this is the optimal solution.   

---

## Discussion

### Comment by @codecov-commenter on September 29, 2020 at 10:44 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=h1) Report
> Merging [#302](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/924bd1367e1d443d382b3aed2a11335d3f39473e?el=desc) will **increase** coverage by `0.20%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #302      +/-   ##
==========================================
+ Coverage   67.57%   67.77%   +0.20%     
==========================================
  Files          56       56              
  Lines         996      993       -3     
  Branches      119      118       -1     
==========================================
  Hits          673      673              
+ Misses        285      283       -2     
+ Partials       38       37       -1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=footer). Last update [924bd13...2ef0aa7](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/302?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 29, 2020 at 11:37 AM UTC

:tada: This PR is included in version 0.24.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.24.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.24.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/302*
