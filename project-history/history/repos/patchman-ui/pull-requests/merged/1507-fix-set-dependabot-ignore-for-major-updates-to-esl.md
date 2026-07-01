---
type: pull_request
number: 1507
title: "fix: Set dependabot ignore for major updates to eslint"
state: merged
author: swadeley
created: 2026-02-16T12:04:22Z
updated: 2026-02-16T14:57:22Z
closed: 2026-02-16T14:57:22Z
merged: 2026-02-16T14:57:22Z
base_branch: master
head_branch: swadeley/hold_eslint_to_major_version_9
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1507
---

# Pull Request #1507: fix: Set dependabot ignore for major updates to eslint

**Author**: @swadeley
**Created**: February 16, 2026 at 12:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `swadeley/hold_eslint_to_major_version_9`

## Description


# Description

    Hold Eslint version to 9.x because with ESLint 10, eslint-plugin-react breaks for every file
     (e.g. src/App.tsx) because of the context.getFilename API change, not only for Playwright tests.


# How to test the PR

Tests pass




---

## Discussion

### Comment by @codecov-commenter on February 16, 2026 at 12:10 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1507?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`cedc0ca`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/cedc0ca36801773c2b26c53c05cb886027336837?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`de9308b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/de9308b64dae76824722b485a9b574b0d8804712?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1507   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1507?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on February 16, 2026 at 02:38 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1507*
