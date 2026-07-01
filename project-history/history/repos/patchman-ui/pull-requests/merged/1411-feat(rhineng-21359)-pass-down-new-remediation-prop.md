---
type: pull_request
number: 1411
title: "feat(RHINENG-21359): Pass down new remediation prop"
state: merged
author: adonispuente
created: 2025-10-14T20:31:13Z
updated: 2025-10-17T19:04:57Z
closed: 2025-10-17T19:04:57Z
merged: 2025-10-17T19:04:56Z
base_branch: master
head_branch: newProp
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1411
---

# Pull Request #1411: feat(RHINENG-21359): Pass down new remediation prop

**Author**: @adonispuente
**Created**: October 14, 2025 at 08:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `newProp`

## Description

This just adds in a new boolean to the button so that if the user has perms but nothing is selected, the button is still disabled and prompts the user to select more items. 
Cannot be merged until https://github.com/RedHatInsights/insights-remediations-frontend/pull/668 is merged 
https://issues.redhat.com/browse/RHINENG-20641

---

## Discussion

### Comment by @codecov-commenter on October 14, 2025 at 08:34 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1411?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.67%. Comparing base ([`ff6c3bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ff6c3bd60630e54c1138882defb52b5267317912?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2adbb71`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2adbb71af2482581d7a57e193f73fb6d2b92bd00?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1411      +/-   ##
==========================================
+ Coverage   62.66%   62.67%   +0.01%     
==========================================
  Files         126      126              
  Lines        3305     3306       +1     
  Branches      857      857              
==========================================
+ Hits         2071     2072       +1     
  Misses       1113     1113              
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1411?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @adonispuente on October 17, 2025 at 04:11 PM UTC

@hugohezel  Let me verify again, locally im not able to reproduce what youre seeing for some reason

### Comment by @adonispuente on October 17, 2025 at 05:36 PM UTC

/retest

---

## Reviews

### Review by @hugohezel - Dismissed on October 16, 2025 at 06:01 PM UTC

Code-wise looks good. Couldn't verify functionality due kessel related error on stage - will do in ON_QA phase.
LGTM, good work @adonispuente!

### Review by @hugohezel - Commented on October 17, 2025 at 03:22 PM UTC

### Review by @hugohezel - Commented on October 17, 2025 at 03:45 PM UTC

Selecting Advisories doesn't enable the button.
<img width="2508" height="825" alt="Image" src="https://github.com/user-attachments/assets/25aad607-7fc8-4f41-b238-cc256c408ebb" />

### Review by @hugohezel - Approved on October 17, 2025 at 04:59 PM UTC

My previous comment was due error in dev env that blocked seeing code changes.
LGTM.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1411*
