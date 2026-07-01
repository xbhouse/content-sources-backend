---
type: pull_request
number: 1529
title: "HMS-10056: Add tests for satellite managed systems"
state: merged
author: Dugowitch
created: 2026-03-02T15:07:45Z
updated: 2026-03-09T12:17:39Z
closed: 2026-03-09T12:17:35Z
merged: 2026-03-09T12:17:35Z
base_branch: master
head_branch: hms10056
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1529
---

# Pull Request #1529: HMS-10056: Add tests for satellite managed systems

**Author**: @Dugowitch
**Created**: March 02, 2026 at 03:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `hms10056`

## Description

# Description

Associated Jira ticket: #HMS-10056

Add a satellite-managed system and add a test that checks:
* template column label is "Managed by Satellite"
* there are corresponding alerts in system's details
* the satellite-managed system cannot be assigned to a template

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on March 02, 2026 at 04:12 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1529?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 73.67%. Comparing base ([`5d8712e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5d8712e00ac5c67116070693893ad0b6e4bec7b3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a7dc55c`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a7dc55c917115cc1fd32d04cb475749ae03eb562?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1529   +/-   ##
=======================================
  Coverage   73.67%   73.67%           
=======================================
  Files          97       97           
  Lines        2359     2359           
  Branches      666      666           
=======================================
  Hits         1738     1738           
  Misses        551      551           
  Partials       70       70           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1529?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @Dugowitch on March 04, 2026 at 04:25 PM UTC

/retest

### Comment by @Dugowitch on March 06, 2026 at 02:03 PM UTC

/retest

---

## Reviews

### Review by @dominikvagner - Approved on March 09, 2026 at 12:06 PM UTC

looks good! ✨  
ack ✅ thx!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1529*
