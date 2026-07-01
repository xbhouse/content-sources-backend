---
type: pull_request
number: 1167
title: "chore(deps): Update axios to fix CVE"
state: merged
author: leSamo
created: 2024-03-01T02:11:30Z
updated: 2024-03-13T13:32:26Z
closed: 2024-03-04T16:22:54Z
merged: 2024-03-04T16:22:54Z
base_branch: master
head_branch: update-axios
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1167
---

# Pull Request #1167: chore(deps): Update axios to fix CVE

**Author**: @leSamo
**Created**: March 01, 2024 at 02:11 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-axios`

## Description

Fixes https://issues.redhat.com/browse/RHINENG-8635

Axios did have an [error handler refactoring](https://github.com/axios/axios/pull/3645) in [0.27.0](https://github.com/axios/axios/releases/tag/v0.27.0), therefore I had to change the accessor to `isAxiosError`.

---

## Discussion

### Comment by @codecov-commenter on March 01, 2024 at 02:19 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1167?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 62.84%. Comparing base [(`3678dbd`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/3678dbd32092734424c7033abc52752826e89d8d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`610e546`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1167?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1167   +/-   ##
=======================================
  Coverage   62.84%   62.84%           
=======================================
  Files         127      127           
  Lines        3192     3192           
  Branches      817      817           
=======================================
  Hits         2006     2006           
  Misses       1186     1186           
```



</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1167?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on March 01, 2024 at 11:12 AM UTC

/retest

### Comment by @leSamo on March 01, 2024 at 12:25 PM UTC

/retest

### Comment by @leSamo on March 01, 2024 at 01:30 PM UTC

/retest

### Comment by @leSamo on March 04, 2024 at 02:43 PM UTC

/retest

### Comment by @mkholjuraev on March 13, 2024 at 01:31 PM UTC

:tada: This PR is included in version 1.67.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on March 01, 2024 at 08:44 AM UTC

LGTM!

### Review by @leSamo - Commented on March 01, 2024 at 11:12 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1167*
