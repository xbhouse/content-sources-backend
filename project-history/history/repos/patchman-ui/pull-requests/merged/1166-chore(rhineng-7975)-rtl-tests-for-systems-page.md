---
type: pull_request
number: 1166
title: "chore(RHINENG-7975): RTL tests for systems page"
state: merged
author: mkholjuraev
created: 2024-02-29T21:28:36Z
updated: 2026-04-06T19:51:32Z
closed: 2024-03-01T16:27:08Z
merged: 2024-03-01T16:27:08Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1166
---

# Pull Request #1166: chore(RHINENG-7975): RTL tests for systems page

**Author**: @mkholjuraev
**Created**: February 29, 2024 at 09:28 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: # (issue)

This is another PR to rewrite enzyme tests to RTL. This is for the Systems page where InventoryTable is used. The PR has a meaningful mocked inventory table component with buttons to trigger functions to open assign, unassign template modals, and remediation wizard. With this, we can test that those models properly get opened without issues. 

Furthermore, props that change the inventory's behavior are tested by asserting their proper state in the tests.

I will create further PRs to have covered other pages where InventroyTable is used


# How to test the PR

Please include steps to test your PR.

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

### Comment by @codecov-commenter on February 29, 2024 at 09:39 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1166?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.87%. Comparing base ([`3678dbd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/3678dbd32092734424c7033abc52752826e89d8d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`059602a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/059602aae99e241ab5c895a5f7996ff342963782?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 302 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1166      +/-   ##
==========================================
+ Coverage   62.84%   62.87%   +0.03%     
==========================================
  Files         127      127              
  Lines        3192     3192              
  Branches      817      817              
==========================================
+ Hits         2006     2007       +1     
+ Misses       1186     1185       -1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1166?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @leSamo on February 29, 2024 at 11:39 PM UTC

/retest

### Comment by @mkholjuraev on March 13, 2024 at 01:32 PM UTC

:tada: This PR is included in version 1.67.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Approved on February 29, 2024 at 11:41 PM UTC

Looks good, could not find any problems!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1166*
