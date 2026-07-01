---
type: pull_request
number: 252
title: "fix(select): disallow selection of non-updatable packages"
state: merged
author: jiridostal
created: 2020-08-27T08:57:56Z
updated: 2026-04-01T07:18:36Z
closed: 2020-08-28T13:20:45Z
merged: 2020-08-28T13:20:45Z
base_branch: master
head_branch: select-update-only
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/252
---

# Pull Request #252: fix(select): disallow selection of non-updatable packages

**Author**: @jiridostal
**Created**: August 27, 2020 at 08:57 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `select-update-only`

## Description

Allow to select only package that have an update

---

## Discussion

### Comment by @codecov-commenter on August 27, 2020 at 09:08 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/252?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 51.42%. Comparing base ([`593cc8f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/593cc8f7f38365facb518c0ff096cf2483c37512?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`62b17b3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/62b17b3e21ddf7a826db5392c603042a06eb186e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1269 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/252?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/252?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemPackages%2FSystemPackages.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | 0.00% | [1 Missing and 2 partials :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/252?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master     #252      +/-   ##
==========================================
+ Coverage   48.34%   51.42%   +3.07%     
==========================================
  Files          55       55              
  Lines         937     1021      +84     
  Branches      101      128      +27     
==========================================
+ Hits          453      525      +72     
- Misses        443      450       +7     
- Partials       41       46       +5     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/252?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @jiridostal on August 28, 2020 at 01:28 PM UTC

:tada: This PR is included in version 0.19.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.19.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.19.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on August 28, 2020 at 01:20 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/252*
