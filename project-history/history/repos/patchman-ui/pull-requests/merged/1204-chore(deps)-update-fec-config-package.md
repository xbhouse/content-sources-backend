---
type: pull_request
number: 1204
title: "chore(deps): Update FEC config package"
state: merged
author: leSamo
created: 2024-10-10T09:25:15Z
updated: 2024-10-25T19:12:05Z
closed: 2024-10-16T13:59:23Z
merged: 2024-10-16T13:59:23Z
base_branch: master
head_branch: update-fec
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1204
---

# Pull Request #1204: chore(deps): Update FEC config package

**Author**: @leSamo
**Created**: October 10, 2024 at 09:25 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-fec`

## Description

https://issues.redhat.com/browse/RHINENG-12002

Cypress had to be update to v13 so that it's compatible with webpack-dev-server v5.

---

## Discussion

### Comment by @codecov-commenter on October 10, 2024 at 09:35 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1204?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`616db83`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/616db833065f1ebb1d3b5f90090a08aece3acbe9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`9acb63f`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9acb63fa6bebf8c630b7d8b3744a20f104cc3e07?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1204   +/-   ##
=======================================
  Coverage   64.01%   64.01%           
=======================================
  Files         124      124           
  Lines        3235     3235           
  Branches      831      831           
=======================================
  Hits         2071     2071           
  Misses       1164     1164           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1204?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on October 14, 2024 at 02:42 PM UTC

/retest

### Comment by @gkarat on October 16, 2024 at 01:58 PM UTC

Everything should be fine. To the exception of some flaky ci.int tests. It's OK to merge it

### Comment by @mkholjuraev on October 25, 2024 at 07:12 PM UTC

:tada: This PR is included in version 1.68.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.68.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on October 10, 2024 at 11:42 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1204*
