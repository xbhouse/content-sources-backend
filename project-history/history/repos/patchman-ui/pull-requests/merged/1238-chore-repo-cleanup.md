---
type: pull_request
number: 1238
title: "chore: Repo cleanup"
state: merged
author: leSamo
created: 2025-01-10T17:03:09Z
updated: 2025-01-15T10:12:05Z
closed: 2025-01-15T10:12:05Z
merged: 2025-01-15T10:12:05Z
base_branch: master
head_branch: cleanup
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1238
---

# Pull Request #1238: chore: Repo cleanup

**Author**: @leSamo
**Created**: January 10, 2025 at 05:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `cleanup`

## Description

- Remove Travis deploy step -- Remove obsolete files and constants related to Travis deploy step which is not performed since containerized builds migration
- Remove firebase, sitemap and sonarqube config files -- files inherited from repository template -- haven't been used for at least 4 years
- Remove Husky -- remove pre-commit test run, to increase consistency with other Interact repos and remove problems with performing commits through VS Code
- Remove rollup packages -- rollup is not used anymore
- Remove semantic release files -- semantic releases were removed around 2 years ago
- Remove redundant svg file inherited from template repository
- Remove old Messages file -- it hasn't been touched in the last 4 years, messages are loaded from a different file

---

## Discussion

### Comment by @codecov-commenter on January 10, 2025 at 05:43 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1238?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.60%. Comparing base [(`803c2dc`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/803c2dc38ad595aad223ca53862d9693ae5caf37?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`b00df34`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b00df349fcb25495944203b14ff45945dfeae442?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1238   +/-   ##
=======================================
  Coverage   63.60%   63.60%           
=======================================
  Files         124      124           
  Lines        3267     3267           
  Branches      860      860           
=======================================
  Hits         2078     2078           
  Misses       1189     1189           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1238?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @gkarat on January 14, 2025 at 10:51 AM UTC

/retest

---

## Reviews

### Review by @gkarat - Approved on January 14, 2025 at 10:49 AM UTC

Just one thing regarding intl. Overall, looks nice and can be merged.

### Review by @leSamo - Commented on January 14, 2025 at 06:03 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1238*
