---
type: pull_request
number: 1630
title: "fix(test): remove workaround for multi-user UI tests"
state: merged
author: xbhouse
created: 2026-05-20T14:34:54Z
updated: 2026-05-21T15:31:09Z
closed: 2026-05-21T15:31:09Z
merged: 2026-05-21T15:31:09Z
base_branch: master
head_branch: remove-test-workaround
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1630
---

# Pull Request #1630: fix(test): remove workaround for multi-user UI tests

**Author**: @xbhouse
**Created**: May 20, 2026 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `remove-test-workaround`

## Description

# Description

- Reverts the workaround we had that split users into different UI projects to prevent auth state conflicts
- The conflicts were caused by the frontend-development-proxy caching parallel responses by URL, and this was [fixed](https://github.com/RedHatInsights/frontend-development-proxy/pull/137) in the latest proxy image

# How to test the PR

UI tests pass with no auth-related failures

---

## Discussion

### Comment by @codecov-commenter on May 20, 2026 at 02:38 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1630?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.58%. Comparing base ([`098cb46`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/098cb4621aeaf3e7299f556518d2e2e8e04df4f3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f76c46d`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f76c46dca53e8bb69e3443a12c2c077ab9458d0f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1630   +/-   ##
=======================================
  Coverage   77.58%   77.58%           
=======================================
  Files         103      103           
  Lines        3266     3266           
  Branches      729      729           
=======================================
  Hits         2534     2534           
  Misses        655      655           
  Partials       77       77           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1630?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @xbhouse on May 21, 2026 at 03:30 PM UTC

merging, ui test failures are unrelated, i'm looking into those [here](https://github.com/RedHatInsights/patchman-ui/pull/1627)

---

## Reviews

### Review by @swadeley - Approved on May 20, 2026 at 05:18 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1630*
