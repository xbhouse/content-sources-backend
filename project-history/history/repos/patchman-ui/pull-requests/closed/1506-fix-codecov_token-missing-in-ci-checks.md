---
type: pull_request
number: 1506
title: "fix: CODECOV_TOKEN missing in CI checks"
state: closed
author: swadeley
created: 2026-02-16T09:32:45Z
updated: 2026-02-16T14:21:40Z
closed: 2026-02-16T14:21:40Z
base_branch: master
head_branch: swadeley/add_code_coverage_token_get_step
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1506
---

# Pull Request #1506: fix: CODECOV_TOKEN missing in CI checks

**Author**: @swadeley
**Created**: February 16, 2026 at 09:32 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `swadeley/add_code_coverage_token_get_step`

## Description

# Description

CODECOV_TOKEN missing in CI checks

# How to test the PR

ci-check passes




---

## Discussion

### Comment by @codecov-commenter on February 16, 2026 at 09:38 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1506?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`cedc0ca`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/cedc0ca36801773c2b26c53c05cb886027336837?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`489eb04`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/489eb046620b2fc6b5f432a12a2743b764127226?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1506   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      899    +7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1506?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @swadeley on February 16, 2026 at 12:50 PM UTC

This did not help, in "Run codecov/codecov-action"  I still see:

```
==> CLI integrity verified

 -> Token of length 0 detected
```

This is a problem in PR checks made by dependabot,

Strangely, it says :

`info - 2026-02-16 09:37:59,372 -- Process Upload complete`

### Comment by @swadeley on February 16, 2026 at 02:21 PM UTC

Set token in "Dependabot secrets" instead.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1506*
