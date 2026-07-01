---
type: pull_request
number: 603
title: "fix(select): Don't ignore query params in bulk select"
state: merged
author: jiridostal
created: 2021-07-22T12:08:30Z
updated: 2022-08-02T08:37:21Z
closed: 2021-07-23T10:02:29Z
merged: 2021-07-23T10:02:29Z
base_branch: master
head_branch: query-params-ignored
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/603
---

# Pull Request #603: fix(select): Don't ignore query params in bulk select

**Author**: @jiridostal
**Created**: July 22, 2021 at 12:08 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `query-params-ignored`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-commenter on July 22, 2021 at 12:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#603](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e9cdf0a) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/7d83083a80b9b55aa2297409324f018ce6edf25c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7d83083) will **not change** coverage.
> The diff coverage is `0.00%`.

> :exclamation: Current head e9cdf0a differs from pull request most recent head 819dfeb. Consider uploading reports for the commit 819dfeb to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #603   +/-   ##
=======================================
  Coverage   52.70%   52.70%           
=======================================
  Files          77       77           
  Lines        1848     1848           
  Branches      399      399           
=======================================
  Hits          974      974           
  Misses        787      787           
  Partials       87       87           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [7d83083...819dfeb](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/603?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on July 23, 2021 at 10:11 AM UTC

:tada: This PR is included in version 1.26.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.26.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.26.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on July 23, 2021 at 10:02 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/603*
