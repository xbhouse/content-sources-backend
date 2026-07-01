---
type: pull_request
number: 1598
title: "fix: handle empty relations and reduce test flakiness"
state: merged
author: xbhouse
created: 2026-05-01T13:49:09Z
updated: 2026-05-08T13:43:21Z
closed: 2026-05-08T13:43:21Z
merged: 2026-05-08T13:43:21Z
base_branch: master
head_branch: empty-relations
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1598
---

# Pull Request #1598: fix: handle empty relations and reduce test flakiness

**Author**: @xbhouse
**Created**: May 01, 2026 at 01:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `empty-relations`

## Description

# Description

- Adjusts permission check to wait for relations to load fully before granting/denying access
- Updates playwright UI test config and AdvisoriesTests to reduce flakiness 
  - First issue: When multiple users fetch workspaces at the same time, this seems to cause auth leakage across sessions so workspaces for the wrong user might be fetched to determine access. To prevent this, this PR isolates UI test projects by user and runs each project serially
  - Second issue: Since AdvisoriesTests creates a remediation plan for an advisory, any systems the advisory is associated with are included in the request to remediations. If a system in the request has already been deleted in inventory, it will still be in patch for around 10 minutes. Retrying will almost always result in this kind of failure, so this PR disables retries for this test
 
# How to test the PR

- The "not authorized" page shouldn't flash briefly on initial app load or refresh
- There shouldn't be any test failures related to rejected access

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on May 01, 2026 at 01:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1598?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `5 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 76.34%. Comparing base ([`e420017`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e4200178def57545354ad3aec2b95014f19b7cf0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0b2d354`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/0b2d354799a46f5e25eed4983975af0e98bbb1ed?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1598?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/hooks/usePermissionCheck.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1598?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FusePermissionCheck.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VQZXJtaXNzaW9uQ2hlY2suanM=) | 0.00% | [3 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1598?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1598      +/-   ##
==========================================
- Coverage   76.38%   76.34%   -0.05%     
==========================================
  Files         103      103              
  Lines        3185     3187       +2     
  Branches      692      693       +1     
==========================================
  Hits         2433     2433              
- Misses        674      675       +1     
- Partials       78       79       +1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1598?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @xbhouse on May 07, 2026 at 08:23 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on May 08, 2026 at 08:56 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1598*
