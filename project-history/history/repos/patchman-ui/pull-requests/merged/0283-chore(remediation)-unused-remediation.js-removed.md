---
type: pull_request
number: 283
title: "chore(Remediation) unused Remediation.js removed"
state: merged
author: mkholjuraev
created: 2020-09-14T11:23:12Z
updated: 2021-08-09T06:55:18Z
closed: 2020-09-15T12:12:19Z
merged: 2020-09-15T12:12:19Z
base_branch: master
head_branch: delete-remediation.js
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/283
---

# Pull Request #283: chore(Remediation) unused Remediation.js removed

**Author**: @mkholjuraev
**Created**: September 14, 2020 at 11:23 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `delete-remediation.js`

## Description

 While I was working on unit tests for **SmartComponents/Remediations**, I realized that Remediation.js is not used anymore. So, if it is actually unnecessary, let us delete it.

---

## Discussion

### Comment by @codecov-commenter on September 14, 2020 at 11:33 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283?src=pr&el=h1) Report
> Merging [#283](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/06a89f0acca7f84ee64f99ddc78ead0c5b43fadc?el=desc) will **increase** coverage by `0.33%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #283      +/-   ##
==========================================
+ Coverage   66.46%   66.80%   +0.33%     
==========================================
  Files          57       56       -1     
  Lines         999      994       -5     
  Branches      118      118              
==========================================
  Hits          664      664              
+ Misses        296      291       -5     
  Partials       39       39              
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283?src=pr&el=footer). Last update [06a89f0...612af78](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/283?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on September 15, 2020 at 12:27 PM UTC

:tada: This PR is included in version 0.24.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.24.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.24.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/283*
