---
type: pull_request
number: 1474
title: "fix(test): update remediation plan test to reflect new remediations UI"
state: merged
author: xbhouse
created: 2026-01-19T16:42:36Z
updated: 2026-01-20T14:49:05Z
closed: 2026-01-20T14:49:05Z
merged: 2026-01-20T14:49:05Z
base_branch: master
head_branch: fix-tests
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1474
---

# Pull Request #1474: fix(test): update remediation plan test to reflect new remediations UI

**Author**: @xbhouse
**Created**: January 19, 2026 at 04:42 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-tests`

## Description

# Description

The remediations UI was updated, this PR fixes the related UI test to reflect those updates.

# How to test the PR

UI/SystemsTests passes (specifically "Create a new remediation plan") 


---

## Discussion

### Comment by @codecov-commenter on January 19, 2026 at 04:54 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1474?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`aa080f7`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/aa080f744f79156dd86345080fd7f2eb7780f80e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2ea3e4a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2ea3e4a525cd269f7b07add2fa218dcd0b9f27eb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1474   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1474?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Commented on January 19, 2026 at 07:49 PM UTC

### Review by @katarinazaprazna - Approved on January 19, 2026 at 07:53 PM UTC

Thanks for taking the time to update the test to match the new UI! I’ve left one small suggestion regarding the `ExecutionHistoryTab`. Aside from that, this is ready to roll 🚢

### Review by @xbhouse - Commented on January 19, 2026 at 08:39 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1474*
