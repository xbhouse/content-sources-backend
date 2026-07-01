---
type: pull_request
number: 1503
title: "fix: host count for remediation"
state: closed
author: swadeley
created: 2026-02-16T03:14:11Z
updated: 2026-02-16T07:49:11Z
closed: 2026-02-16T07:39:57Z
base_branch: master
head_branch: swadeley/fix_host_count_planned_remediations
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1503
---

# Pull Request #1503: fix: host count for remediation

**Author**: @swadeley
**Created**: February 16, 2026 at 03:14 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `swadeley/fix_host_count_planned_remediations`

## Description

# Description

Fix test to filter on system created for the test.

# How to test the PR

'Check actions, systems, and execution history tabs' test passes.

# Before the change

'Check actions, systems, and execution history tabs' test fails because many other systems from other tests are present in the table.

# After the change

'Check actions, systems, and execution history tabs' test passes.






---

## Discussion

### Comment by @codecov-commenter on February 16, 2026 at 03:16 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1503?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`cedc0ca`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/cedc0ca36801773c2b26c53c05cb886027336837?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`eef7060`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/eef70608f6120e2bf9083decb1ebf1994bcecac7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1503   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1503?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @swadeley on February 16, 2026 at 07:39 AM UTC

I will create new PR from my fork:

https://github.com/RedHatInsights/patchman-ui/pull/1505

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1503*
