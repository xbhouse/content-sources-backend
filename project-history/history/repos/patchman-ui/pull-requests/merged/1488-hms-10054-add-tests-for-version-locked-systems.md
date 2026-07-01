---
type: pull_request
number: 1488
title: "HMS-10054: add tests for version locked systems"
state: merged
author: xbhouse
created: 2026-02-05T19:28:07Z
updated: 2026-02-12T17:18:26Z
closed: 2026-02-12T17:18:26Z
merged: 2026-02-12T17:18:26Z
base_branch: master
head_branch: 10054
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1488
---

# Pull Request #1488: HMS-10054: add tests for version locked systems

**Author**: @xbhouse
**Created**: February 05, 2026 at 07:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `10054`

## Description

# Description

Associated Jira ticket: [HMS-10054](https://issues.redhat.com/browse/HMS-10054)

- Adds an archive for a version-locked system
- Adds a test to verify version-locked systems are displayed in patch with the appropriate tooltip and that templates cannot be assigned to these types of systems
- Adds a helper to create a template and modifies the existing CanAssignTemplateToSystem test to use this

# How to test the PR

Tests pass

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

### Comment by @codecov-commenter on February 05, 2026 at 07:51 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1488?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`19600df`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/19600df8f43d5bdb0b9ed9bd5566d2f48b64b71a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a16d6b9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a16d6b9b5f17059b5cba3ecfe87bb1a1634af0de?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1488   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      892           
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1488?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Approved on February 08, 2026 at 08:46 PM UTC

Bryttanie, thanks for the updates! The changes look great. It looks like those test failures are just noise from the current Patch issues. As soon as that clears up and we get a green build, let’s ship this :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1488*
