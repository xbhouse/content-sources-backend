---
type: pull_request
number: 1457
title: "HMS-9725: Add test to verify exporting feature on patch"
state: merged
author: mayurilahane
created: 2025-12-19T03:03:47Z
updated: 2026-01-09T20:57:10Z
closed: 2026-01-09T20:57:10Z
merged: 2026-01-09T20:57:10Z
base_branch: master
head_branch: mlahane/9725
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1457
---

# Pull Request #1457: HMS-9725: Add test to verify exporting feature on patch

**Author**: @mayurilahane
**Created**: December 19, 2025 at 03:03 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `mlahane/9725`

## Description

# Description

- Export data into CSV and JSON.
- Test exporting ability on the "Advisories", "Packages" 
and "Systems" pages.
- Validate the data structure and some of the contents.
- Used "papaparse" to parse the CSV files and 
added it to the dependencies.
- Fixed advisories navigation to use exact matching. -

https://issues.redhat.com/browse/HMS-9725


---

## Discussion

### Comment by @mayurilahane on December 19, 2025 at 04:10 AM UTC

/retest

### Comment by @codecov-commenter on December 19, 2025 at 04:14 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1457?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`1bf4fc5`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1bf4fc5f4a51be3a5422e9efd903d4b8cb56051a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f00fd24`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f00fd24ac8dace4be0e6bc4e3e83634456732102?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1457   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      899      892    -7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1457?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Changes Requested on December 22, 2025 at 06:49 PM UTC

There are some minor comments, but otherwise all is good, great job!

### Review by @dominikvagner - Commented on January 06, 2026 at 10:00 AM UTC

### Review by @dominikvagner - Commented on January 08, 2026 at 10:15 AM UTC

2 small comments with possible improvements 🔧  
otherwise, this looks good! ✨ 👏🏼 

### Review by @TenSt - Approved on January 08, 2026 at 01:27 PM UTC

lgtm

### Review by @mayurilahane - Commented on January 08, 2026 at 06:28 PM UTC

### Review by @mayurilahane - Commented on January 08, 2026 at 06:28 PM UTC

### Review by @dominikvagner - Approved on January 09, 2026 at 09:33 AM UTC

nice work! thanks! ✨ 
ack! 🚀 ✅  

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1457*
