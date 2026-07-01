---
type: pull_request
number: 650
title: "chore(deps): vulnerabilities are remediated"
state: merged
author: mkholjuraev
created: 2021-09-14T09:16:42Z
updated: 2022-05-17T08:49:45Z
closed: 2021-09-14T12:21:03Z
merged: 2021-09-14T12:21:03Z
base_branch: master
head_branch: update
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/650
---

# Pull Request #650: chore(deps): vulnerabilities are remediated

**Author**: @mkholjuraev
**Created**: September 14, 2021 at 09:16 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update`

## Description

1.  SPM-1141: CVE-2021-3749 is remediated by updating Axios version from 0.21.1 to 0.21.4
2. SPM-1093: CVE-2021-32803 is remediated by updating deps in the dependency tree using 'npm audit fix'
3. SPM-1094: CVE-2021-32804 is remediated by updating deps in the dependency tree using 'npm audit fix'

---

## Discussion

### Comment by @codecov-commenter on September 14, 2021 at 09:24 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#650](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (98fdc7b) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/a5e8710494ce59435c95d72261b50a02349ed273?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a5e8710) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #650   +/-   ##
=======================================
  Coverage   52.10%   52.10%           
=======================================
  Files          81       81           
  Lines        1967     1967           
  Branches      436      436           
=======================================
  Hits         1025     1025           
  Misses        857      857           
  Partials       85       85           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [a5e8710...98fdc7b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/650?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on September 24, 2021 at 12:44 PM UTC

:tada: This PR is included in version 1.33.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.33.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.33.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/650*
