---
type: pull_request
number: 1445
title: "feat(test): initialize playwright tests"
state: merged
author: dominikvagner
created: 2025-12-01T13:59:23Z
updated: 2025-12-08T08:34:42Z
closed: 2025-12-08T08:34:42Z
merged: 2025-12-08T08:34:41Z
base_branch: master
head_branch: kickstart-playwright-1
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1445
---

# Pull Request #1445: feat(test): initialize playwright tests

**Author**: @dominikvagner
**Created**: December 01, 2025 at 01:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `kickstart-playwright-1`

## Description

# Description
This initializes playwright tests in this repo. Adds all the necessary configs, setups a CI pipeline to run these tests on every PR. Implements the needed fixtures and helpers for testing and adds 2 simple tests as examples.

Associated Jira ticket: [#HMS-9438](https://issues.redhat.com/browse/HMS-9438)

---
Checklist:
- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on December 01, 2025 at 02:02 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1445?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`bed59d0`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bed59d0fcef92861faf034c206f07dbd714afa9c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`549a81f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/549a81f1dc2ef3c2b4512038ff3b3171f8439719?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1445      +/-   ##
==========================================
+ Coverage   62.35%   62.53%   +0.18%     
==========================================
  Files         128      127       -1     
  Lines        3323     3310      -13     
  Branches      894      892       -2     
==========================================
- Hits         2072     2070       -2     
+ Misses       1130     1120      -10     
+ Partials      121      120       -1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1445?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Commented on December 03, 2025 at 02:14 PM UTC

Overall looks great! Some comments below, but otherwise great job.

### Review by @dominikvagner - Commented on December 04, 2025 at 10:06 AM UTC

### Review by @dominikvagner - Commented on December 04, 2025 at 10:07 AM UTC

### Review by @TenSt - Approved on December 04, 2025 at 10:22 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1445*
