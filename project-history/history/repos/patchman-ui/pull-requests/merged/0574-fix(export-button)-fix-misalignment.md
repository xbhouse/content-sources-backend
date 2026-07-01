---
type: pull_request
number: 574
title: "fix(Export button): fix misalignment"
state: merged
author: mkholjuraev
created: 2021-06-25T12:06:11Z
updated: 2026-04-03T07:05:37Z
closed: 2021-06-28T13:24:39Z
merged: 2021-06-28T13:24:39Z
base_branch: master
head_branch: align-remediation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/574
---

# Pull Request #574: fix(Export button): fix misalignment

**Author**: @mkholjuraev
**Created**: June 25, 2021 at 12:06 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `align-remediation`

## Description

The problem is related to how FEC sets min-width to the button provided by dedicatedActions. This is a temporary fix. 

![image](https://user-images.githubusercontent.com/59481011/123422470-8a8e0d80-d5be-11eb-91e7-529c8c3eed2f.png)


---

## Discussion

### Comment by @codecov-commenter on June 25, 2021 at 12:09 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/574?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 53.72%. Comparing base ([`320e15a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/320e15a0f4499299f10c92578d0517359421d055?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`4b0ccb3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4b0ccb3755b655b9ba314352f168e9015385c119?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 906 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master     #574   +/-   ##
=======================================
  Coverage   53.72%   53.72%           
=======================================
  Files          77       77           
  Lines        1798     1798           
  Branches      380      380           
=======================================
  Hits          966      966           
  Misses        761      761           
  Partials       71       71           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/574?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @jiridostal on June 28, 2021 at 01:33 PM UTC

:tada: This PR is included in version 1.21.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.21.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.21.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/574*
