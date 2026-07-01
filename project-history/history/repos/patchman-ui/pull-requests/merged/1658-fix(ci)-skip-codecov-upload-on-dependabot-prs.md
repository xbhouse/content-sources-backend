---
type: pull_request
number: 1658
title: "fix(ci): skip codecov upload on dependabot PRs"
state: merged
author: swadeley
created: 2026-06-11T18:04:12Z
updated: 2026-06-11T18:49:12Z
closed: 2026-06-11T18:49:11Z
merged: 2026-06-11T18:49:11Z
base_branch: master
head_branch: swadeley/skip-codecov-for-dependabot
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1658
---

# Pull Request #1658: fix(ci): skip codecov upload on dependabot PRs

**Author**: @swadeley
**Created**: June 11, 2026 at 06:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `swadeley/skip-codecov-for-dependabot`

## Description

## Summary
- Skip the Codecov upload step when the workflow is triggered by Dependabot
- Dependabot PRs cannot access regular GitHub Actions secrets, so uploads fail with a misleading "Repository not found" error even though v7 works on master pushes
- Coverage is still uploaded when changes merge to master

## Test plan
- [ ] Verify Dependabot PR CI passes (lint, tests, build) without the codecov step
- [ ] Verify master push CI still uploads coverage successfully


Made with [Cursor](https://cursor.com)

---

## Discussion

### Comment by @codecov-commenter on June 11, 2026 at 06:07 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1658?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.27%. Comparing base ([`4b4cbff`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4b4cbff289fc6268a7d0fff584b3416fa5d5c9fb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1e645e2`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1e645e22be0e7a24ff41d40d12a95715643121bc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1658      +/-   ##
==========================================
- Coverage   77.57%   77.27%   -0.31%     
==========================================
  Files         103      103              
  Lines        3278     3287       +9     
  Branches      731      735       +4     
==========================================
- Hits         2543     2540       -3     
- Misses        658      668      +10     
- Partials       77       79       +2     
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1658?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @xbhouse - Commented on June 11, 2026 at 06:38 PM UTC

### Review by @xbhouse - Approved on June 11, 2026 at 06:39 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1658*
