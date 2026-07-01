---
type: pull_request
number: 1587
title: "chore: update readme and rules for agents"
state: merged
author: xbhouse
created: 2026-04-21T21:38:04Z
updated: 2026-05-14T13:13:22Z
closed: 2026-05-14T13:13:22Z
merged: 2026-05-14T13:13:22Z
base_branch: master
head_branch: readme-agent-updates
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1587
---

# Pull Request #1587: chore: update readme and rules for agents

**Author**: @xbhouse
**Created**: April 21, 2026 at 09:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `readme-agent-updates`

## Description

# Description

- Adds initial setup for agents with the same config from [content-sources-frontend](https://github.com/content-services/content-sources-frontend/pull/963). Moves the existing Playwright rule so it's auto-loaded by  both Claude and Cursor
- Adds a rule and some initial guidelines to help agents help us with migrating to TypeScript
- Updates README

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on April 21, 2026 at 09:41 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1587?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.21%. Comparing base ([`391b812`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/391b8128417497b2aeba732dcbeff7dbfc9964ec?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a0dd27d`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a0dd27dd9e00fecf520bf729be8551dbd4b39f6c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1587   +/-   ##
=======================================
  Coverage   76.21%   76.21%           
=======================================
  Files         103      103           
  Lines        3187     3187           
  Branches      693      697    +4     
=======================================
  Hits         2429     2429           
  Misses        678      678           
  Partials       80       80           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1587?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Commented on April 29, 2026 at 05:54 AM UTC

### Review by @katarinazaprazna - Dismissed on April 29, 2026 at 05:57 AM UTC

Thanks for the changes!

I left one small suggestion. Feel free to consider it if you think it adds value, but otherwise, awesome work on this :)

### Review by @xbhouse - Commented on April 29, 2026 at 05:23 PM UTC

### Review by @katarinazaprazna - Approved on May 12, 2026 at 08:36 PM UTC

Thank you! ❤️ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1587*
