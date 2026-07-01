---
type: pull_request
number: 1642
title: "HMS-10728: add 'Expand All' option to table header"
state: merged
author: Dugowitch
created: 2026-05-27T09:38:37Z
updated: 2026-06-12T07:53:49Z
closed: 2026-06-11T17:46:52Z
merged: 2026-06-11T17:46:52Z
base_branch: master
head_branch: ui-review
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1642
---

# Pull Request #1642: HMS-10728: add 'Expand All' option to table header

**Author**: @Dugowitch
**Created**: May 27, 2026 at 09:38 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `ui-review`

## Description

# Description
Associated Jira ticket: [HMS-10728](https://redhat.atlassian.net/browse/HMS-10728)

Add "Expand all" option to the table header for tables that have expandable rows.

# How to test the PR
Check that tables with expandable rows have "expand all" option in the table header and that it works as expected. Check that other tables do not have this option.

# Before the change
<img width="1558" height="153" alt="image" src="https://github.com/user-attachments/assets/b8a12bab-ce81-4e96-a7df-dc95088767e7" />

# After the change
<img width="1558" height="153" alt="image" src="https://github.com/user-attachments/assets/a260c29a-3009-490e-8d30-c05565f4f548" />

# Checklist:
- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


[HMS-10728]: https://redhat.atlassian.net/browse/HMS-10728?atlOrigin=eyJpIjoiNWRkNTljNzYxNjVmNDY3MDlhMDU5Y2ZhYzA5YTRkZjUiLCJwIjoiZ2l0aHViLWNvbS1KU1cifQ

---

## Discussion

### Comment by @codecov-commenter on May 27, 2026 at 09:41 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `58.82353%` with `14 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 77.27%. Comparing base ([`4b4cbff`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4b4cbff289fc6268a7d0fff584b3416fa5d5c9fb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`6a77167`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6a7716754127d5e8766a4896d687f476b8dfd6b6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisories%2FAdvisories.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | 46.15% | [6 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemAdvisories%2FSystemAdvisories.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | 0.00% | [6 Missing and 1 partial :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1642      +/-   ##
==========================================
- Coverage   77.57%   77.27%   -0.31%     
==========================================
  Files         103      103              
  Lines        3278     3287       +9     
  Branches      731      735       +4     
==========================================
- Hits         2543     2540       -3     
- Misses        658      668      +10     
- Partials       77       79       +2     
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1642?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @swadeley on May 28, 2026 at 07:32 PM UTC

Hi @Dugowitch , rebase please.

### Comment by @swadeley on June 04, 2026 at 09:21 AM UTC

Hi, new dropdown arrow is present and works as described:

<img width="231" height="60" alt="image" src="https://github.com/user-attachments/assets/0bf6575b-13a8-4cb9-bfda-4b67d643c789" />


### Comment by @TenSt on June 11, 2026 at 12:32 PM UTC

/retest

### Comment by @TenSt on June 11, 2026 at 12:50 PM UTC

/retest

### Comment by @swadeley on June 11, 2026 at 02:38 PM UTC

 /ok-to-test

### Comment by @swadeley on June 11, 2026 at 05:40 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on June 04, 2026 at 09:21 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1642*
