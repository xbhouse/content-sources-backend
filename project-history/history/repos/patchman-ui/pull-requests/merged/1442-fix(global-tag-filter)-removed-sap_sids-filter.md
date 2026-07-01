---
type: pull_request
number: 1442
title: "fix(Global tag filter): removed sap_sids filter"
state: merged
author: computercamplove
created: 2025-11-27T13:12:25Z
updated: 2025-12-05T11:51:22Z
closed: 2025-12-05T11:51:22Z
merged: 2025-12-05T11:51:22Z
base_branch: master
head_branch: remove-sap-sids-filter
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1442
---

# Pull Request #1442: fix(Global tag filter): removed sap_sids filter

**Author**: @computercamplove
**Created**: November 27, 2025 at 01:12 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `remove-sap-sids-filter`

## Description

# Description

Associated Jira ticket: # (https://issues.redhat.com/browse/RHINENG-21539)

while updating Global tag filter with new workaloads filed it was decided that we no longer needed SAP SIDS filter - it was removed https://github.com/RedHatInsights/insights-chrome/pull/3300

It cause issue in Content systems page while applying tag filter via Global tag Filter


# How to test the PR

1. Navigate to Content systems page (or any page where you use global tag filter)
2. Apply any tag option via Global tag filter
3. Open Network tab

# Before the change
Verify against production


it apply tag value into `sap_sids` filter:
<img width="438" height="62" alt="image" src="https://github.com/user-attachments/assets/908474b2-960d-4d7d-923d-39f4a16257ca" />
<img width="2955" height="716" alt="image" src="https://github.com/user-attachments/assets/f26d62fd-3cb7-4859-9186-d56359da3123" />

# After the change
No `sap_sids` filter should applied on while requesting filtered `/systems`
(in `stage` it will return 0 systems because of current bug https://redhat-internal.slack.com/archives/C05KTHJU4R4/p1764241977728679, please verify it in `production` env)

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

### Comment by @codecov-commenter on November 27, 2025 at 01:15 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `40.00000%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.38%. Comparing base ([`bed59d0`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bed59d0fcef92861faf034c206f07dbd714afa9c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`39f781b`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/39f781b622df1c465e1b36eb063018e48fa12873?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 50.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/App.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?src=pr&el=tree&filepath=src%2FApp.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL0FwcC5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1442      +/-   ##
==========================================
+ Coverage   62.35%   62.38%   +0.03%     
==========================================
  Files         128      128              
  Lines        3323     3318       -5     
  Branches      894      892       -2     
==========================================
- Hits         2072     2070       -2     
+ Misses       1130     1128       -2     
+ Partials      121      120       -1     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1442?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @computercamplove on November 27, 2025 at 02:46 PM UTC

@dominikvagner i am going to draft for now this PR - we decided to revert last change https://github.com/RedHatInsights/insights-chrome/pull/3305 - i need to make sure my change works with reverted PR in insights-chrome 

---

## Reviews

### Review by @TenSt - Approved on December 04, 2025 at 11:49 AM UTC

lgtm

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1442*
