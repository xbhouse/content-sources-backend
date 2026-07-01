---
type: pull_request
number: 1421
title: "feat(RHINENG-21349): Pass in sentry vars via konflux"
state: merged
author: adonispuente
created: 2025-10-27T16:39:23Z
updated: 2025-10-30T14:49:31Z
closed: 2025-10-30T14:49:31Z
merged: 2025-10-30T14:49:31Z
base_branch: master
head_branch: sourcemaps
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1421
---

# Pull Request #1421: feat(RHINENG-21349): Pass in sentry vars via konflux

**Author**: @adonispuente
**Created**: October 27, 2025 at 04:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `sourcemaps`

## Description

This PR allows konflux to access the secret variables needed to initialize sentry properly and upload sourcemaps to sentry in the image builds we later serve to production
[RHINENG-21349](https://issues.redhat.com/browse/RHINENG-21349) 

---

## Discussion

### Comment by @codecov-commenter on October 27, 2025 at 04:43 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1421?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.67%. Comparing base ([`4186869`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/41868696288797029de54c51b9fe5022b2779ba8?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c7a2d6a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c7a2d6aa3f30c641f42b1a1f62f855b527725cc0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1421   +/-   ##
=======================================
  Coverage   62.67%   62.67%           
=======================================
  Files         126      126           
  Lines        3306     3306           
  Branches      863      857    -6     
=======================================
  Hits         2072     2072           
  Misses       1113     1113           
  Partials      121      121           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1421?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on October 30, 2025 at 10:25 AM UTC

LGTM, thank you!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1421*
