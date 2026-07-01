---
type: pull_request
number: 1482
title: "chore(CI): Add Konflux to commitlint exceptions"
state: merged
author: LightOfHeaven1994
created: 2026-01-29T08:50:39Z
updated: 2026-01-29T16:57:23Z
closed: 2026-01-29T16:57:23Z
merged: 2026-01-29T16:57:22Z
base_branch: master
head_branch: commit-lint-konflux-exception
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1482
---

# Pull Request #1482: chore(CI): Add Konflux to commitlint exceptions

**Author**: @LightOfHeaven1994
**Created**: January 29, 2026 at 08:50 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `commit-lint-konflux-exception`

## Description

Commit lint is being hard on PRs that are opened by konflux bot. Sometime ago we have configured automerge for konflux PRs that are not breaking important pipelines (such as konflux-on-pull-request, npm tests). But currently commit-lint pipeline prevents automerges (see example https://github.com/RedHatInsights/patchman-ui/pull/1471) by being to strict to bot commits

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on January 29, 2026 at 08:53 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1482?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`8a3a6f9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/8a3a6f98a4c10bdb78ef1d0d05ce2726c37c475e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`35809e9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/35809e9c3bb37e1626b014cfa24073356e2fd7d4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1482   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1482?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1482*
