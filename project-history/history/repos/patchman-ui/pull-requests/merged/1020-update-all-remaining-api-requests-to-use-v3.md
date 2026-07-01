---
type: pull_request
number: 1020
title: "Update all remaining API requests to use v3"
state: merged
author: leSamo
created: 2023-04-11T14:46:06Z
updated: 2026-04-04T07:02:53Z
closed: 2023-04-12T18:50:05Z
merged: 2023-04-12T18:50:05Z
base_branch: master
head_branch: api-v3-last
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1020
---

# Pull Request #1020: Update all remaining API requests to use v3

**Author**: @leSamo
**Created**: April 11, 2023 at 02:46 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `api-v3-last`

## Description

There should be no change in functionality.

---

## Discussion

### Comment by @codecov-commenter on April 11, 2023 at 02:56 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1020?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `40.00000%` with `6 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 64.30%. Comparing base ([`3e941e0`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/3e941e09b7852c766f6b6e9f3911e4e2de84a58b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`9b1a0ff`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9b1a0ff74c5a299b742268e9254c98c715d97a2e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 419 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1020?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/api.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1020?src=pr&el=tree&filepath=src%2FUtilities%2Fapi.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | 40.00% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1020?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1020   +/-   ##
=======================================
  Coverage   64.30%   64.30%           
=======================================
  Files         116      116           
  Lines        2821     2821           
  Branches      725      725           
=======================================
  Hits         1814     1814           
  Misses       1007     1007           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1020?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on April 17, 2023 at 06:40 PM UTC

:tada: This PR is included in version 1.63.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on April 12, 2023 at 05:12 PM UTC

LGTM codewise! not testing locally as the changes are obvious.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1020*
