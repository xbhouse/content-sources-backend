---
type: pull_request
number: 1582
title: "Test: add prefix to test names"
state: merged
author: swadeley
created: 2026-04-16T18:19:23Z
updated: 2026-04-17T08:17:22Z
closed: 2026-04-17T08:17:22Z
merged: 2026-04-17T08:17:22Z
base_branch: master
head_branch: swadeley/fix_system_names_use_prefix
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1582
---

# Pull Request #1582: Test: add prefix to test names

**Author**: @swadeley
**Created**: April 16, 2026 at 06:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `swadeley/fix_system_names_use_prefix`

## Description

# Description

Adding a prefix and using the name returned by the system fixture for test names.
To ensure a system found is the one created in that specific test run.



---

## Discussion

### Comment by @codecov-commenter on April 16, 2026 at 06:22 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1582?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`1cc5a1b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1cc5a1bd34f9bf664e0314dcb0e5a78954814c4a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`bc2334f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bc2334f743d2e63ba4517334b6c6566c3a4ad028?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1582   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      686      686           
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1582/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1582/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `?` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1582/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `?` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1582?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @xbhouse - Commented on April 16, 2026 at 06:40 PM UTC

### Review by @swadeley - Commented on April 16, 2026 at 06:50 PM UTC

### Review by @xbhouse - Approved on April 16, 2026 at 07:36 PM UTC

thank you!! tests pass locally (with the workspace test case commented out). ack 🚀 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1582*
