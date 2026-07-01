---
type: pull_request
number: 1470
title: "HMS-9728: add integration test for template assignment"
state: merged
author: katarinazaprazna
created: 2026-01-15T18:17:05Z
updated: 2026-01-28T09:48:34Z
closed: 2026-01-28T09:48:34Z
merged: 2026-01-28T09:48:34Z
base_branch: master
head_branch: systems-templates-integration-test
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1470
---

# Pull Request #1470: HMS-9728: add integration test for template assignment

**Author**: @katarinazaprazna
**Created**: January 15, 2026 at 06:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `systems-templates-integration-test`

## Description

# Description

I’ve added a new integration test, `CanAssignTemplateToSystem.spec.ts`, to cover the end-to-end template assignment flow. This walks through registering a system with a unique identity, creating a patch template, and assigning it, finally verifying that the Systems UI updates the "Template" column correctly.

I introduced a `cleanupTemplates` helper to ensure test data is properly removed and updated the cleanup fixture to prevent premature teardown. I also added a `disableTrackingAndConsent` utility to block external scripts like TrustArc and Amplitude, which helps reduce console noise.

Associated Jira ticket: https://issues.redhat.com/browse/HMS-9728

## Other

Added a condition to track failed task statuses. This prevents infinite loops and test hangs by ensuring the polling mechanism recognizes both success and failure as terminal states. This addresses observed issues where failing tasks caused the test suite to hang indefinitely.

I ended up using the Patch API for template assignment rather than `rhc`. I found that `rhc` wouldn't apply the template if it detected the system was already registered, which forced me to disconnect and reconnect the system just to make it work. Switching to the API for assignment (while keeping `rhc` for registration) was much cleaner, though it took a bit of refactoring to get right.


# How to test the PR

1. Run the specific test and ensure it passes successfully
2. Verify that no stale templates with the prefix `TemplateAssignmentTest-` remain in the environment after the test completes

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on January 15, 2026 at 06:20 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1470?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`6993b74`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6993b74a4145748d8c19db615774ed2879ce0832?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`8eb5c43`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/8eb5c435b39e1c1b14b3d7c23b44e7c7a4dbf76c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1470   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1470?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @marusak - Commented on January 16, 2026 at 07:23 AM UTC

### Review by @katarinazaprazna - Commented on January 16, 2026 at 11:20 AM UTC

### Review by @dominikvagner - Commented on January 20, 2026 at 10:02 AM UTC

looks great, just two smaller comments/thoughts on logging 💭 
good job! thanks! 👏🏼🚀

### Review by @marusak - Commented on January 20, 2026 at 01:29 PM UTC

### Review by @marusak - Commented on January 20, 2026 at 01:29 PM UTC

### Review by @marusak - Commented on January 20, 2026 at 01:30 PM UTC

### Review by @marusak - Commented on January 20, 2026 at 01:32 PM UTC

### Review by @marusak - Commented on January 20, 2026 at 01:34 PM UTC

### Review by @marusak - Commented on January 20, 2026 at 01:35 PM UTC

### Review by @katarinazaprazna - Commented on January 20, 2026 at 07:24 PM UTC

### Review by @katarinazaprazna - Commented on January 20, 2026 at 07:40 PM UTC

### Review by @katarinazaprazna - Commented on January 20, 2026 at 08:03 PM UTC

### Review by @marusak - Commented on January 27, 2026 at 12:14 PM UTC

### Review by @TenSt - Approved on January 27, 2026 at 01:03 PM UTC

lgtm!
There are some minor comments from @marusak about navigating to the same page where you are right now, but otherwise all looks good, great job!

### Review by @katarinazaprazna - Commented on January 27, 2026 at 02:18 PM UTC

### Review by @dominikvagner - Approved on January 27, 2026 at 02:19 PM UTC

ack from me too! ✅ 
_(the extra navigation mentioned doesn't seem necessary, but doesn't break anything)_

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1470*
