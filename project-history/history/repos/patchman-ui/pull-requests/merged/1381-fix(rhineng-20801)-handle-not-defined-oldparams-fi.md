---
type: pull_request
number: 1381
title: "fix(RHINENG-20801): Handle not defined oldParams filter"
state: merged
author: LightOfHeaven1994
created: 2025-09-17T09:13:32Z
updated: 2025-09-17T14:36:34Z
closed: 2025-09-17T14:36:34Z
merged: 2025-09-17T14:36:34Z
base_branch: master
head_branch: RHINENG-20801
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1381
---

# Pull Request #1381: fix(RHINENG-20801): Handle not defined oldParams filter

**Author**: @LightOfHeaven1994
**Created**: September 17, 2025 at 09:13 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-20801`

## Description

This PR fixes navigation from Dashboard to Patchman Advisories 
```
Uncaught TypeError: can't access property "filter", oldParams is undefined
```

### How to test:

1. Open Dashboard page and find Patch card
2. Click on any of these links [`Security advisories`, `But fixes`, `Enhancements`]
3. Observe that Advisories page is open without any errors and filter is applied
<img width="1238" height="476" alt="image" src="https://github.com/user-attachments/assets/a5775279-1514-4c00-b41b-ecea38b28b63" />



---

## Discussion

### Comment by @codecov-commenter on September 17, 2025 at 09:15 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1381?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.57%. Comparing base ([`52cf79e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/52cf79e786bd483283b87d14cae409de9539f037?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b69de63`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b69de634109a814cdd1528cf2fc540f5bbfbd9e9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1381   +/-   ##
=======================================
  Coverage   62.57%   62.57%           
=======================================
  Files         126      126           
  Lines        3337     3337           
  Branches      866      872    +6     
=======================================
  Hits         2088     2088           
  Misses       1128     1128           
  Partials      121      121           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1381?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @bastilian - Commented on September 17, 2025 at 11:07 AM UTC

### Review by @LightOfHeaven1994 - Commented on September 17, 2025 at 11:48 AM UTC

### Review by @LightOfHeaven1994 - Commented on September 17, 2025 at 01:33 PM UTC

### Review by @bastilian - Approved on September 17, 2025 at 02:31 PM UTC

LGTM! nice fix! Thank you @LightOfHeaven1994!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1381*
