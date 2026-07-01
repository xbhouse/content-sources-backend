---
type: pull_request
number: 1493
title: "chore(kessel): react-kessel-access-check"
state: merged
author: mtclinton
created: 2026-02-09T21:38:43Z
updated: 2026-02-10T14:46:44Z
closed: 2026-02-10T14:46:44Z
merged: 2026-02-10T14:46:44Z
base_branch: master
head_branch: RHINENG-23945
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1493
---

# Pull Request #1493: chore(kessel): react-kessel-access-check

**Author**: @mtclinton
**Created**: February 09, 2026 at 09:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-23945`

## Description

Add @project-kessel/react-kessel-access-check package for fetching Kessel access checks.

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

### Comment by @katarinazaprazna on February 09, 2026 at 10:22 PM UTC

Could you please shorten the commit title? The `ci-checks` are failing as the title cannot be longer than 50 chars. Thanks!

### Comment by @mtclinton on February 09, 2026 at 11:19 PM UTC

/retest

### Comment by @codecov-commenter on February 09, 2026 at 11:49 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1493?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`87ec4bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/87ec4bdb48b1c3c0e5e99802e10cc4341592d4c7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`3724a5d`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/3724a5dc8c84999106cacc9765409d092f5b69ad?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1493   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      899      892    -7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1493?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mtclinton on February 10, 2026 at 01:35 PM UTC

Thanks for the help and review on this! Merging

---

## Reviews

### Review by @dominikvagner - Dismissed on February 10, 2026 at 08:48 AM UTC

the Playwright failures are unrelated, they are caused by an issue elsewhere 👍🏼 [[1](https://issues.redhat.com/browse/RHINENG-23892)]

### Review by @katarinazaprazna - Approved on February 10, 2026 at 11:11 AM UTC

Thank you! ✅

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1493*
