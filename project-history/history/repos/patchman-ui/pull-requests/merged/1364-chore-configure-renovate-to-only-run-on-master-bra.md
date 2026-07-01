---
type: pull_request
number: 1364
title: "chore: configure renovate to only run on master branch"
state: merged
author: ryemorris
created: 2025-08-07T17:45:00Z
updated: 2025-08-18T08:32:31Z
closed: 2025-08-18T08:32:31Z
merged: 2025-08-18T08:32:31Z
base_branch: master
head_branch: chore/add-renovate-basebranches
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1364
---

# Pull Request #1364: chore: configure renovate to only run on master branch

**Author**: @ryemorris
**Created**: August 07, 2025 at 05:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `chore/add-renovate-basebranches`

## Description

Currently, mintmaker will try to update all onboarded components but this is unnecessary for SC branches and just clutters the repo PR list. Only allowing renovate to run on main/master should help.

---

## Discussion

### Comment by @codecov-commenter on August 07, 2025 at 05:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1364?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.54%. Comparing base ([`4f8d3ef`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4f8d3ef640e59918ea81931a8a0df5bf00a4d90c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`605257c`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/605257c6b78c51761b0281551b074b5a7d3345d1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1364   +/-   ##
=======================================
  Coverage   62.54%   62.54%           
=======================================
  Files         126      126           
  Lines        3335     3335           
  Branches      865      865           
=======================================
  Hits         2086     2086           
  Misses       1128     1128           
  Partials      121      121           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1364/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1364/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.54% <ø> (?)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1364/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (?)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1364/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.54% <ø> (?)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1364?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on August 18, 2025 at 08:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1364*
