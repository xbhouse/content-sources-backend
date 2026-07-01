---
type: pull_request
number: 1645
title: "test: reduce filter test runtime"
state: merged
author: xbhouse
created: 2026-05-27T20:07:28Z
updated: 2026-05-28T12:10:52Z
closed: 2026-05-28T12:10:52Z
merged: 2026-05-28T12:10:52Z
base_branch: master
head_branch: reduce-test-time
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1645
---

# Pull Request #1645: test: reduce filter test runtime

**Author**: @xbhouse
**Created**: May 27, 2026 at 08:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `reduce-test-time`

## Description

# Description

Removes unnecessary waits from the filter test helpers. This should reduce the runtime of the  "Patch Filters" > "Filter types on Advisories page" UI test.  

# How to test the PR

UI filter tests shouldn't time out 

---

## Discussion

### Comment by @codecov-commenter on May 27, 2026 at 08:10 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1645?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.58%. Comparing base ([`1143e8b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1143e8b0401f2805164b2c2676f6b8d1aa224838?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ae5ac1f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ae5ac1feef8a5677e89052c27a80db8cc8b795dd?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1645   +/-   ##
=======================================
  Coverage   77.58%   77.58%           
=======================================
  Files         103      103           
  Lines        3266     3266           
  Branches      734      729    -5     
=======================================
  Hits         2534     2534           
  Misses        655      655           
  Partials       77       77           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1645/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1645/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `77.58% <ø> (?)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1645/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `77.58% <ø> (?)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1645?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @swadeley on May 28, 2026 at 11:46 AM UTC

Hi @xbhouse , " UI/FilterTypePatch" still fails but that is because the base system which is in the Workspace called "TestSpace" is not in the results, in the table. But it is there when the test is still at the filter by operating stage of the test. 

### Comment by @xbhouse on May 28, 2026 at 12:07 PM UTC

> Hi @xbhouse , " UI/FilterTypePatch" still fails but that is because the base system which is in the Workspace called "TestSpace" is not in the results, in the table. But it is there when the test is still at the filter by operating stage of the test.

that's expected @swadeley, the workspace filtering should be fixed once Inventory reverts the breaking changes. this PR only affects the "Filter types on Advisories page" test (reduces the runtime of that test from about 4 minutes to 1.5 minutes)  and we shouldn't see timeouts like [this](https://github.com/RedHatInsights/patchman-ui/actions/runs/26517766967/job/78130640145) anymore

---

## Reviews

### Review by @swadeley - Approved on May 28, 2026 at 11:37 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1645*
