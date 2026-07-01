---
type: pull_request
number: 1091
title: "fix(SOM-2082): properly update URL params in useGetEntities hook"
state: merged
author: mkholjuraev
created: 2023-07-17T10:47:22Z
updated: 2026-04-02T02:59:09Z
closed: 2023-07-20T09:52:57Z
merged: 2023-07-20T09:52:57Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1091
---

# Pull Request #1091: fix(SOM-2082): properly update URL params in useGetEntities hook

**Author**: @mkholjuraev
**Created**: July 17, 2023 at 10:47 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: # [(issue)](https://issues.redhat.com/browse/SPM-2082)

Fixes the missing global filter chip after page refresh on Systems and advisory systems pages and also removes redundant trash code in Hooks.js


# How to test the PR

1. Navigate to the system all page.
2. Select SAP from the global filter.
3. Refresh the page
4. Check workloads SAP filter after refreshing the page

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

### Comment by @codecov-commenter on July 19, 2023 at 02:28 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1091?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `1 line` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.70%. Comparing base ([`64549d6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/64549d6b590761cdd83cd6334b2ddbb068a5b438?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`efae9c3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/efae9c3be3a1358578a9de730f1aed80da8e24fa?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 348 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1091?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1091?src=pr&el=tree&filepath=src%2FUtilities%2FHooks.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1091?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1091   +/-   ##
=======================================
  Coverage   62.70%   62.70%           
=======================================
  Files         119      119           
  Lines        2984     2984           
  Branches      765      765           
=======================================
  Hits         1871     1871           
  Misses       1113     1113           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1091?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on July 20, 2023 at 07:27 AM UTC

/retest

### Comment by @mkholjuraev on July 20, 2023 at 10:11 AM UTC

:tada: This PR is included in version 1.63.9 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.9)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on July 19, 2023 at 12:14 PM UTC

This works great, good job!
Just fix the `package-lock` and the tests should pass

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1091*
