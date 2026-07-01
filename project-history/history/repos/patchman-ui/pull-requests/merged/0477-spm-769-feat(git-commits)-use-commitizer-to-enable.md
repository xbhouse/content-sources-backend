---
type: pull_request
number: 477
title: "SPM-769  feat(git commits): use commitizer to enable conventionam commit naming"
state: merged
author: mkholjuraev
created: 2021-03-25T15:51:00Z
updated: 2021-08-09T06:56:23Z
closed: 2021-04-06T10:05:52Z
merged: 2021-04-06T10:05:52Z
base_branch: master
head_branch: install-cz
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/477
---

# Pull Request #477: SPM-769  feat(git commits): use commitizer to enable conventionam commit naming

**Author**: @mkholjuraev
**Created**: March 25, 2021 at 03:51 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `install-cz`

## Description

This is a test PR from using commitizer for the ticket https://issues.redhat.com/browse/SPM-769

[cz-cli](https://github.com/commitizen) has several steps to finish the commit e.g type, scope, footer .... I thought some of them are unnecessaryy, thus I used [cz-customizable](https://github.com/leonardoanalista/cz-customizable) with custom config which can be improved according to needs later. 

---

## Discussion

### Comment by @codecov-io on March 25, 2021 at 04:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477?src=pr&el=h1) Report
> Merging [#477](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477?src=pr&el=desc) (d99a7c3) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/353e91bbb1057c277191735fa9e71deb75259037?el=desc) (353e91b) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #477   +/-   ##
=======================================
  Coverage   54.31%   54.31%           
=======================================
  Files          78       78           
  Lines        1541     1541           
  Branches      179      179           
=======================================
  Hits          837      837           
  Misses        643      643           
  Partials       61       61           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477?src=pr&el=footer). Last update [353e91b...d99a7c3](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/477?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on April 06, 2021 at 10:16 AM UTC

:tada: This PR is included in version 1.15.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.15.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.15.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Changes Requested on March 29, 2021 at 12:28 PM UTC

Looks and works really great, just some minor improvements!

### Review by @mkholjuraev - Commented on March 30, 2021 at 12:26 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/477*
