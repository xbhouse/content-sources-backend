---
type: pull_request
number: 1484
title: "test(HMS-10053): add plan remediations test from advisories"
state: merged
author: TenSt
created: 2026-01-30T13:07:23Z
updated: 2026-02-03T10:22:41Z
closed: 2026-02-03T10:22:41Z
merged: 2026-02-03T10:22:41Z
base_branch: master
head_branch: stepan/HMS-10053-extend-plan-remediations-tests-from-advisories
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1484
---

# Pull Request #1484: test(HMS-10053): add plan remediations test from advisories

**Author**: @TenSt
**Created**: January 30, 2026 at 01:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `stepan/HMS-10053-extend-plan-remediations-tests-from-advisories`

## Description

# Description

Associated Jira ticket: #HMS-10053

- add plan remediations test from advisories
- refactor remediations cleanup

# How to test the PR

`Advisories Tests` should pass.

---

## Discussion

### Comment by @codecov-commenter on January 30, 2026 at 01:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1484?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`dd44879`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/dd4487908be3aae17a78999b38c1d065e3d02907?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`aaa4665`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/aaa4665267d48e1668bb68980c9dd852f215c4ae?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1484   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      892           
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1484?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @TenSt on January 30, 2026 at 05:02 PM UTC

The tests were working fine locally, then failed in CI after PR was created and now they fail locally too with the same error. I'm trying to investigate what is happening.

### Comment by @TenSt on January 30, 2026 at 05:08 PM UTC

I've added `scrollIntoViewIfNeeded()` function for the failing step and the skeleton table is there:
<img width="1280" height="720" alt="test-failed-1" src="https://github.com/user-attachments/assets/156c89bc-3b4c-48a0-b8a2-000b03172f72" />


---

## Reviews

### Review by @xbhouse - Commented on January 30, 2026 at 08:23 PM UTC

### Review by @xbhouse - Commented on January 30, 2026 at 09:00 PM UTC

### Review by @TenSt - Commented on February 02, 2026 at 03:58 PM UTC

### Review by @TenSt - Commented on February 02, 2026 at 04:45 PM UTC

### Review by @xbhouse - Approved on February 02, 2026 at 09:23 PM UTC

lgtm! nice job 🥳 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1484*
