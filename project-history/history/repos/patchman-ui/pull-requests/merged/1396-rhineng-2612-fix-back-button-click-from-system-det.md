---
type: pull_request
number: 1396
title: "RHINENG-2612: fix back button click from system detail page"
state: merged
author: Dugowitch
created: 2025-10-08T16:13:49Z
updated: 2025-10-13T08:02:15Z
closed: 2025-10-13T08:02:11Z
merged: 2025-10-13T08:02:11Z
base_branch: master
head_branch: RHINENG-2612
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1396
---

# Pull Request #1396: RHINENG-2612: fix back button click from system detail page

**Author**: @Dugowitch
**Created**: October 08, 2025 at 04:13 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-2612`

## Description

# Description
Browser back button did not work as expected. First click would just remove path segment instead of returning to the previous page. It was caused by pushes to history even on internal state change. The fix uses `replace: true`, so browser back button click should always take the user back to the systems list page (or some other page they came from).

# How to test the PR
1. Navigate to Content > Systems > a system detail
2. optionally apply some filters, change page, ...
3. click the browser back button

Check that navigation with browser back/forward buttons works as expected in all cases.

# Before the change
History would be polluted with unnecessary pushes and user would need to click the back button twice or even several times to get back to the system list page. 

# After the change
A single browser back button click takes the user back to the system list page.

# Checklist:
- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on October 08, 2025 at 04:20 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1396?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.66%. Comparing base ([`9936e95`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9936e9536be6a431c6db572fdf78b7674c23ec62?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`526e95e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/526e95e15aa104569468cd241948c52696e1cfa6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1396      +/-   ##
==========================================
+ Coverage   62.46%   62.66%   +0.19%     
==========================================
  Files         126      126              
  Lines        3312     3305       -7     
  Branches      857      857              
==========================================
+ Hits         2069     2071       +2     
+ Misses       1122     1113       -9     
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1396?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @Dugowitch on October 09, 2025 at 06:09 PM UTC

/retest

---

## Reviews

### Review by @bastilian - Approved on October 10, 2025 at 12:04 PM UTC

Nice! LGTM! Thank you for fixing this @Dugowitch!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1396*
