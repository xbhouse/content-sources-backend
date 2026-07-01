---
type: pull_request
number: 1089
title: "fix(SPM-2117): remove await inside react hook"
state: merged
author: mkholjuraev
created: 2023-07-13T13:56:23Z
updated: 2026-04-04T08:45:38Z
closed: 2023-07-13T15:07:19Z
merged: 2023-07-13T15:07:19Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1089
---

# Pull Request #1089: fix(SPM-2117): remove await inside react hook

**Author**: @mkholjuraev
**Created**: July 13, 2023 at 01:56 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

# Description

Associated Jira ticket: # https://issues.redhat.com/browse/SPM-2117

This fixes the frozen systems table on Patch.


# How to test the PR

1. Run the PR locally
2. navigate to systems page
3. try selecting a row
4. observe that the page is not frozen
5. try searing somethin
6. observer again that page is not frozen
7. Play with the page

# Before the change


# After the change


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

### Comment by @codecov-commenter on July 13, 2023 at 02:05 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1089?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `66.66667%` with `2 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.79%. Comparing base ([`83f93b4`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/83f93b41b50694e95fbf1bdbf44e76751ce28271?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`7f53e27`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7f53e272e17e3be9df915c3142cda055efaae2b4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 357 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1089?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1089?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FStatusReports%2FSystemsStatusReport.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | 66.66% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1089?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1089      +/-   ##
==========================================
- Coverage   62.84%   62.79%   -0.05%     
==========================================
  Files         119      119              
  Lines        2971     2973       +2     
  Branches      763      763              
==========================================
  Hits         1867     1867              
- Misses       1104     1106       +2     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1089?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @adonispuente on July 13, 2023 at 02:42 PM UTC

/retest

### Comment by @mkholjuraev on July 13, 2023 at 03:26 PM UTC

:tada: This PR is included in version 1.63.6 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.6)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on July 13, 2023 at 02:45 PM UTC

When testing I am unable to reproduce the bug, great job! LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1089*
