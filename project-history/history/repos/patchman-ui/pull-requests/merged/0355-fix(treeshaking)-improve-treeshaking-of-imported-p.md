---
type: pull_request
number: 355
title: "fix(treeshaking): Improve treeshaking of imported patch detail"
state: merged
author: karelhala
created: 2020-12-09T15:45:24Z
updated: 2026-04-04T15:57:08Z
closed: 2020-12-10T08:42:09Z
merged: 2020-12-10T08:42:09Z
base_branch: master
head_branch: better-treeshaking
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/355
---

# Pull Request #355: fix(treeshaking): Improve treeshaking of imported patch detail

**Author**: @karelhala
**Created**: December 09, 2020 at 03:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `better-treeshaking`

## Description

### Imported patch detail bloats the app bundle

When importing patch detail the final build is bloated by approx 4MB, this PR fixes such issue by improving how PF and FEC files are imported. When finally imported the build increased only by 0.4MB.

---

## Discussion

### Comment by @karelhala on December 09, 2020 at 03:50 PM UTC

ping @mkholjuraev @jiridostal 

### Comment by @jiridostal on December 10, 2020 at 08:42 AM UTC

Looks great! Thanks

### Comment by @codecov-commenter on April 04, 2026 at 03:57 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/355?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 73.62%. Comparing base ([`5e22483`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5e224831c521b9be543e1940007beedda05193e5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1385b31`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1385b31dc46376c9cb8235bcf385af720459ee1a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1125 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master     #355   +/-   ##
=======================================
  Coverage   73.62%   73.62%           
=======================================
  Files          68       68           
  Lines        1202     1202           
  Branches      159      159           
=======================================
  Hits          885      885           
  Misses        264      264           
  Partials       53       53           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/355?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/355*
