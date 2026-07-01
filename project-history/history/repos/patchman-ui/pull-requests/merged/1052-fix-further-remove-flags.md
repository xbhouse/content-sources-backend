---
type: pull_request
number: 1052
title: "fix: further remove flags"
state: merged
author: mkholjuraev
created: 2023-05-10T14:50:32Z
updated: 2026-04-05T19:34:55Z
closed: 2023-05-10T16:51:03Z
merged: 2023-05-10T16:51:03Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1052
---

# Pull Request #1052: fix: further remove flags

**Author**: @mkholjuraev
**Created**: May 10, 2023 at 02:50 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: # (issue)

Removes further flags used for templates.

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

### Comment by @codecov-commenter on May 10, 2023 at 03:02 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `57.14286%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.12%. Comparing base ([`7e5344b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7e5344b1afd19e909b1a1623f1918b03aa39c0dc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f4a68d0`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f4a68d0598093a895d8f9e54ba108c9d7260fb9b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 385 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisorySystems%2FAdvisorySystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Systems/Systems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?src=pr&el=tree&filepath=src%2FUtilities%2FSystemsHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | 50.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1052      +/-   ##
==========================================
- Coverage   63.16%   63.12%   -0.04%     
==========================================
  Files         119      119              
  Lines        2959     2945      -14     
  Branches      761      751      -10     
==========================================
- Hits         1869     1859      -10     
+ Misses       1090     1086       -4     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1052?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on May 10, 2023 at 05:14 PM UTC

:tada: This PR is included in version 1.63.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Approved on May 10, 2023 at 04:23 PM UTC

LGTM! Thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1052*
