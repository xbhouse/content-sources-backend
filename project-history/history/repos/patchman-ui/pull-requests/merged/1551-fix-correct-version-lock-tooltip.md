---
type: pull_request
number: 1551
title: "fix: correct version lock tooltip"
state: merged
author: katarinazaprazna
created: 2026-03-18T17:05:55Z
updated: 2026-03-19T08:49:28Z
closed: 2026-03-19T08:49:28Z
merged: 2026-03-19T08:49:28Z
base_branch: master
head_branch: fix-tooltip-version-locked-systems
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1551
---

# Pull Request #1551: fix: correct version lock tooltip

**Author**: @katarinazaprazna
**Created**: March 18, 2026 at 05:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-tooltip-version-locked-systems`

## Description

# Description
Twin PR to https://github.com/content-services/content-sources-frontend/pull/911 for consistency across applications.

- Replaced unsupported prop value `inlineFlex` with a supported equivalent `flexDefault`.
- Updated tooltip message for clarity and consistency. "This system is locked to version X" (instead of "Your RHEL version is locked at version X")

# How to test the PR

- Verify the version lock tooltip displays correctly for version-locked systems
- Confirm the lock icon appears properly aligned next to the OS name
- Run tests to ensure they pass

# After the change
<img width="2405" height="1565" alt="Screenshot From 2026-03-18 17-56-29" src="https://github.com/user-attachments/assets/818dde86-42c1-439a-b859-8b1f99803ca6" />

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work

---

## Discussion

### Comment by @codecov-commenter on March 18, 2026 at 05:09 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1551?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 72.34%. Comparing base ([`ccc7f0e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ccc7f0e96fe9300cec6ca6644b09384af2b02530?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1f6fe41`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1f6fe41ae34b0a693a2b7914297da2d9853b2ac4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1551      +/-   ##
==========================================
+ Coverage   72.30%   72.34%   +0.03%     
==========================================
  Files          99       99              
  Lines        2405     2408       +3     
  Branches      677      678       +1     
==========================================
+ Hits         1739     1742       +3     
  Misses        586      586              
  Partials       80       80              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1551?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on March 19, 2026 at 08:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1551*
