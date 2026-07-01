---
type: pull_request
number: 1516
title: "HMS-9755: remove code related to patch sets"
state: merged
author: Dugowitch
created: 2026-02-18T17:49:51Z
updated: 2026-02-24T14:46:05Z
closed: 2026-02-24T14:45:48Z
merged: 2026-02-24T14:45:48Z
base_branch: master
head_branch: hms9755
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1516
---

# Pull Request #1516: HMS-9755: remove code related to patch sets

**Author**: @Dugowitch
**Created**: February 18, 2026 at 05:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `hms9755`

## Description

# Description
Content templates have replaced patch sets and the APIs that patch sets used were removed on the backend. Patch sets have been disabled in the patch UI using a feature flag, so no UI changes are expected. This PR removes all code associated with patch sets. 

Associated Jira ticket: #HMS-9755

# How to test the PR
Check that the UI looks and works the same.

# Before the change
Patch sets code present, but patch sets disabled by feature flag.

# After the change
No patch sets code. No UI changes.

# Checklist:
- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work

---

## Discussion

### Comment by @codecov-commenter on February 18, 2026 at 06:22 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `50.00000%` with `2 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 73.67%. Comparing base ([`6fb82f8`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6fb82f8c94950b1c0440f115e45f410a5a957313?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`680d837`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/680d837a89bfddc44393ea953e65459cd0ebbea0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 7 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNUYWJsZS5qcw==) | 50.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/hooks/useFeatureFlag.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?src=pr&el=tree&filepath=src%2FUtilities%2Fhooks%2FuseFeatureFlag.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9ob29rcy91c2VGZWF0dXJlRmxhZy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@             Coverage Diff             @@
##           master    #1516       +/-   ##
===========================================
+ Coverage   62.53%   73.67%   +11.13%     
===========================================
  Files         127       97       -30     
  Lines        3310     2359      -951     
  Branches      899      671      -228     
===========================================
- Hits         2070     1738      -332     
+ Misses       1120      551      -569     
+ Partials      120       70       -50     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1516?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on February 19, 2026 at 01:32 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1516*
