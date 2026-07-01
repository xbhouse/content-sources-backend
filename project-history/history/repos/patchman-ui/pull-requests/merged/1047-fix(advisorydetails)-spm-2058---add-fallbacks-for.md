---
type: pull_request
number: 1047
title: "fix(AdvisoryDetails): SPM-2058 - Add fallbacks for osName"
state: merged
author: bastilian
created: 2023-05-09T13:03:02Z
updated: 2026-03-31T12:19:51Z
closed: 2023-05-09T14:03:35Z
merged: 2023-05-09T14:03:35Z
base_branch: master
head_branch: SPM-2058
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1047
---

# Pull Request #1047: fix(AdvisoryDetails): SPM-2058 - Add fallbacks for osName

**Author**: @bastilian
**Created**: May 09, 2023 at 01:03 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `SPM-2058`

## Description

# Description

Associated Jira ticket: # SPM-2058 (issue)

Please include a summary of the change, what this fixes/creates/improves.


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

### Comment by @codecov-commenter on May 09, 2023 at 01:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1047?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `1 line` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 63.09%. Comparing base ([`bf9cf45`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bf9cf45381a02889ec9f93d81614ea77c8ac3469?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c3cbb36`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c3cbb3652783fdc053ed6b6b2f4ff15e5d36e46d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 382 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1047?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1047?src=pr&el=tree&filepath=src%2FUtilities%2FDataMappers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1047?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1047   +/-   ##
=======================================
  Coverage   63.09%   63.09%           
=======================================
  Files         119      119           
  Lines        2962     2962           
  Branches      763      763           
=======================================
  Hits         1869     1869           
  Misses       1093     1093           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1047?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on May 09, 2023 at 04:05 PM UTC

:tada: This PR is included in version 1.63.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on May 09, 2023 at 01:18 PM UTC

LGTM! I would gladly support consolidating and unifying OS field handling both in the engine and front end to avoid such cases. 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1047*
