---
type: pull_request
number: 572
title: "fix(Headers): isLoading is set correctly"
state: merged
author: mkholjuraev
created: 2021-06-23T07:52:12Z
updated: 2021-08-09T06:56:52Z
closed: 2021-06-24T07:38:04Z
merged: 2021-06-24T07:38:04Z
base_branch: master
head_branch: fix-detail-headers
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/572
---

# Pull Request #572: fix(Headers): isLoading is set correctly

**Author**: @mkholjuraev
**Created**: June 23, 2021 at 07:52 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-detail-headers`

## Description

Continuously loading headers are fixed in advisory detail and package detail page.

---

## Discussion

### Comment by @codecov-commenter on June 23, 2021 at 07:57 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#572](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7415525) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/1850af60bb305669e16f87728523f744954ebf62?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1850af6) will **not change** coverage.
> The diff coverage is `n/a`.

> :exclamation: Current head 7415525 differs from pull request most recent head 1da16c8. Consider uploading reports for the commit 1da16c8 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #572   +/-   ##
=======================================
  Coverage   53.72%   53.72%           
=======================================
  Files          77       77           
  Lines        1798     1798           
  Branches      380      380           
=======================================
  Hits          966      966           
  Misses        761      761           
  Partials       71       71           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PackageDetail/PackageDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlRGV0YWlsL1BhY2thZ2VEZXRhaWwuanM=) | `0.00% <ø> (ø)` | |
| [src/store/Reducers/AdvisoryDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL0Fkdmlzb3J5RGV0YWlsU3RvcmUuanM=) | `85.71% <ø> (ø)` | |
| [src/store/Reducers/PackageDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1BhY2thZ2VEZXRhaWxTdG9yZS5qcw==) | `42.85% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1850af6...1da16c8](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/572?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on June 24, 2021 at 07:46 AM UTC

:tada: This PR is included in version 1.21.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.21.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.21.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/572*
