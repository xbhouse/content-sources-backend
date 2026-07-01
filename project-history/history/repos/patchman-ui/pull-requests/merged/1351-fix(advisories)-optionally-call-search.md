---
type: pull_request
number: 1351
title: "fix(Advisories): Optionally call search"
state: merged
author: bastilian
created: 2025-07-09T20:04:06Z
updated: 2025-07-10T07:58:28Z
closed: 2025-07-10T07:58:28Z
merged: 2025-07-10T07:58:28Z
base_branch: master
head_branch: fix_search_error
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1351
---

# Pull Request #1351: fix(Advisories): Optionally call search

**Author**: @bastilian
**Created**: July 09, 2025 at 08:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix_search_error`

## Description

# Description

Associated Jira ticket: #-


This prevents the page erroring in case the `queryParams` are undefined. 


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on July 09, 2025 at 08:07 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1351?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 62.52%. Comparing base [(`5db3ffe`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5db3ffedae86db8f2ba9fed1c7d0477ca60cb077?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`bcedf77`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bcedf774deb07e5087333c7a0d18a7d716949e16?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@             Coverage Diff              @@
##            master    #1351       +/-   ##
============================================
- Coverage   100.00%   62.52%   -37.48%     
============================================
  Files            1      126      +125     
  Lines            6     3333     +3327     
  Branches         2      869      +867     
============================================
+ Hits             6     2084     +2078     
- Misses           0     1128     +1128     
- Partials         0      121      +121     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1351?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @adonispuente on July 09, 2025 at 09:47 PM UTC

/retest

---

## Reviews

### Review by @adonispuente - Approved on July 09, 2025 at 09:45 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1351*
