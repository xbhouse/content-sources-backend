---
type: pull_request
number: 1050
title: "fix: ungate the templates route"
state: merged
author: mkholjuraev
created: 2023-05-10T10:27:42Z
updated: 2026-04-02T14:22:27Z
closed: 2023-05-10T12:25:29Z
merged: 2023-05-10T12:25:29Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1050
---

# Pull Request #1050: fix: ungate the templates route

**Author**: @mkholjuraev
**Created**: May 10, 2023 at 10:27 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: # (issue)

This un gates the templates routes completely.


# How to test the PR

1. Run the PR locally
2. open C.R.C with insights-qa acc on stage-beta env
3. navigate to /templates route from the left sidebar
4. observer that you are not redirected to /advisories route
5. Now open a new session on stage-stable env
6. navigate to /templates route as previously done
7. observe that you are not redirected to /advisories
8. The same procedure applies to /templates/:templateName route as well 

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

### Comment by @codecov-commenter on May 10, 2023 at 10:39 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1050?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `1 line` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.16%. Comparing base ([`10fc592`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/10fc5924c2bac8450335f1a2b7e8c0c141a8b3f1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`3a4b8db`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/3a4b8dbc8800f479176789bbb70ab810e134dcc9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 385 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1050?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1050?src=pr&el=tree&filepath=src%2FRoutes.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1050?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1050      +/-   ##
==========================================
+ Coverage   63.09%   63.16%   +0.06%     
==========================================
  Files         119      119              
  Lines        2962     2959       -3     
  Branches      763      761       -2     
==========================================
  Hits         1869     1869              
+ Misses       1093     1090       -3     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1050?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on May 10, 2023 at 12:25 PM UTC

@AsToNlele yup, the flag also can be removed

### Comment by @mkholjuraev on May 10, 2023 at 05:14 PM UTC

:tada: This PR is included in version 1.63.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on May 10, 2023 at 12:21 PM UTC

Works well thank you! Can the feature flag be removed completely or is this a quick fix?

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1050*
