---
type: pull_request
number: 1328
title: "Konflux build pipeline service account migration"
state: merged
author: red-hat-konflux
created: 2025-04-30T05:26:29Z
updated: 2025-04-30T09:35:16Z
closed: 2025-04-30T09:35:16Z
merged: 2025-04-30T09:35:16Z
base_branch: master
head_branch: konflux-sa-migration-patchman-ui
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1328
---

# Pull Request #1328: Konflux build pipeline service account migration

**Author**: @red-hat-konflux
**Created**: April 30, 2025 at 05:26 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux-sa-migration-patchman-ui`

## Description


## Build pipeline Service Account migration

This PR changes Service Account used by build pipeline from "appstudio-pipeline" to dedicated to the Component Service Account.
Please merge the Service Account update to avoid broken builds when deprected "appstudio-pipeline" Service Account is removed.


---

## Discussion

### Comment by @codecov-commenter on April 30, 2025 at 05:32 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1328?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 100.00%. Comparing base [(`c7a507c`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c7a507caf43daf015a8b88ade10d69bb14c74735?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5eb05ca`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5eb05caf363fc9b41c1399b25a660a1b6112fb5c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##            master     #1328   +/-   ##
=========================================
  Coverage   100.00%   100.00%           
=========================================
  Files            1         1           
  Lines            6         6           
  Branches         2         2           
=========================================
  Hits             6         6           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1328/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1328/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1328/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1328?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @LightOfHeaven1994 on April 30, 2025 at 09:22 AM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1328*
