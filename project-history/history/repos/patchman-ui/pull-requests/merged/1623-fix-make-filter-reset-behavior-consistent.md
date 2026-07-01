---
type: pull_request
number: 1623
title: "fix:  make filter reset behavior consistent "
state: merged
author: dominikvagner
created: 2026-05-13T13:49:23Z
updated: 2026-05-18T09:12:16Z
closed: 2026-05-18T09:11:13Z
merged: 2026-05-18T09:11:13Z
base_branch: master
head_branch: fix-reset-filter-behavior-3
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1623
---

# Pull Request #1623: fix:  make filter reset behavior consistent 

**Author**: @dominikvagner
**Created**: May 13, 2026 at 01:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-reset-filter-behavior-3`

## Description

This fixes some of the incorrect behavior with filter clear or reset
buttons. They should now only be shown when the state differs from
default and use correct terms for restoring and clearing (depends on
default filters).

---

## Discussion

### Comment by @codecov-commenter on May 13, 2026 at 01:54 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `98.16514%` with `2 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 77.58%. Comparing base ([`ae76011`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ae76011a19c7c569b10a09035ee7be80a692f265?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`93bef4b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/93bef4b9cfa4b34b3e2b58165b6bda764096da00?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 98.18% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/hooks/Hooks.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FHooks.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy9Ib29rcy5qcw==) | 92.30% | [0 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1623      +/-   ##
==========================================
+ Coverage   76.21%   77.58%   +1.37%     
==========================================
  Files         103      103              
  Lines        3187     3266      +79     
  Branches      693      729      +36     
==========================================
+ Hits         2429     2534     +105     
+ Misses        678      655      -23     
+ Partials       80       77       -3     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1623?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on May 14, 2026 at 10:52 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1623*
