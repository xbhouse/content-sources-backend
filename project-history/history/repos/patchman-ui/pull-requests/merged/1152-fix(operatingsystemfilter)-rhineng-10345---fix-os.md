---
type: pull_request
number: 1152
title: "fix(OperatingSystemFilter): RHINENG-10345 - Fix OS inventory filter conversion"
state: merged
author: bastilian
created: 2023-12-12T21:10:41Z
updated: 2026-04-01T16:49:12Z
closed: 2024-07-03T09:49:08Z
merged: 2024-07-03T09:49:08Z
base_branch: master
head_branch: RHINENG-2912-2
labels: ["bug", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1152
---

# Pull Request #1152: fix(OperatingSystemFilter): RHINENG-10345 - Fix OS inventory filter conversion

**Author**: @bastilian
**Created**: December 12, 2023 at 09:10 PM UTC
**Status**: Merged
**Labels**: `bug`, `released`
**Base**: `master` ← **Head**: `RHINENG-2912-2`

## Description

This fixes issues with syncing the state and URL of filters in the Systems table.


# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

Please include steps to test your PR.

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

### Comment by @codecov-commenter on December 12, 2023 at 09:20 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `64.10256%` with `14 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 64.00%. Comparing base ([`7f7bc89`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7f7bc89d38149928a4ac7cffa35423cb22a529a7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e7c00af`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e7c00afc4a1dab9c7c33d4aafb847f0c3dd3ea61?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 280 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNUYWJsZS5qcw==) | 36.36% | [7 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/SystemsHelpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?src=pr&el=tree&filepath=src%2FUtilities%2FSystemsHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9TeXN0ZW1zSGVscGVycy5qcw==) | 16.66% | [5 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 90.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1152      +/-   ##
==========================================
- Coverage   64.14%   64.00%   -0.14%     
==========================================
  Files         124      124              
  Lines        3207     3234      +27     
  Branches      818      830      +12     
==========================================
+ Hits         2057     2070      +13     
- Misses       1150     1164      +14     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1152?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @mkholjuraev on June 25, 2024 at 12:02 PM UTC

@bastilian can you please have a look at this PR. The OS filter is broken in both stage and prod and we are receiving more and more complaints unfortunately.

### Comment by @mkholjuraev on July 01, 2024 at 09:41 AM UTC

/retest

### Comment by @niyazRedhat on July 01, 2024 at 10:12 AM UTC

/retest

### Comment by @mtclinton on July 01, 2024 at 12:32 PM UTC

/retest

### Comment by @mkholjuraev on July 03, 2024 at 10:10 AM UTC

:tada: This PR is included in version 1.67.7 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.7)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Commented on June 04, 2024 at 09:16 AM UTC

Looks good! I have tried to test it, couldn't find issues and the OS filter now operates as expected. There is one thing I would like to ask for in terms of tests before we can merge...

### Review by @bastilian - Commented on June 04, 2024 at 09:31 AM UTC

### Review by @gkarat - Commented on June 04, 2024 at 09:56 AM UTC

### Review by @bastilian - Commented on June 04, 2024 at 09:59 AM UTC

### Review by @mkholjuraev - Commented on June 04, 2024 at 11:38 AM UTC

### Review by @bastilian - Commented on June 04, 2024 at 11:41 AM UTC

### Review by @mkholjuraev - Commented on June 04, 2024 at 11:46 AM UTC

### Review by @bastilian - Commented on June 04, 2024 at 11:47 AM UTC

### Review by @bastilian - Commented on June 04, 2024 at 11:49 AM UTC

### Review by @bastilian - Commented on June 04, 2024 at 11:52 AM UTC

### Review by @bastilian - Commented on June 04, 2024 at 11:53 AM UTC

### Review by @mkholjuraev - Approved on July 01, 2024 at 09:41 AM UTC

LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1152*
