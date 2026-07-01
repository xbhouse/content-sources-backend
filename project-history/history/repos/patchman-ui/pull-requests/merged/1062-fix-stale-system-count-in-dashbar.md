---
type: pull_request
number: 1062
title: "Fix stale system count in Dashbar"
state: merged
author: leSamo
created: 2023-05-25T10:19:16Z
updated: 2026-04-01T20:36:44Z
closed: 2023-05-29T10:07:56Z
merged: 2023-05-29T10:07:56Z
base_branch: master
head_branch: fix-stale-count
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1062
---

# Pull Request #1062: Fix stale system count in Dashbar

**Author**: @leSamo
**Created**: May 25, 2023 at 10:19 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-stale-count`

## Description

https://issues.redhat.com/browse/SPM-2054

### After
![Screenshot from 2023-05-25 12-16-06](https://github.com/RedHatInsights/patchman-ui/assets/8426204/2667ee11-c1a3-4e14-a11f-d8c3ff42c917)

- stale system count used to display 0 instead of the correct amount, this was fixed by adding staleness param to the API request
- improved the responsivity of the dashbar
- fixed camel case `SystemsStatusreport` -> `SystemsStatusReport`

---

## Discussion

### Comment by @codecov-commenter on May 25, 2023 at 10:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1062?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 61.92%. Comparing base ([`546dd7e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/546dd7e53b01ef76706e6f371978f0203c23b437?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`109c604`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/109c604573facc4a15a4a463e0326732ab562d16?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 375 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1062      +/-   ##
==========================================
- Coverage   63.03%   61.92%   -1.12%     
==========================================
  Files         119      119              
  Lines        2949     3044      +95     
  Branches      754      795      +41     
==========================================
+ Hits         1859     1885      +26     
- Misses       1090     1159      +69     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1062?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @Fewwy - Approved on May 29, 2023 at 09:56 AM UTC

Fixed! LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1062*
