---
type: pull_request
number: 1452
title: "HMS-9727: setup Nightly pipeline (stage)"
state: merged
author: katarinazaprazna
created: 2025-12-12T13:45:39Z
updated: 2026-01-06T11:06:15Z
closed: 2026-01-06T10:57:50Z
merged: 2026-01-06T10:57:50Z
base_branch: master
head_branch: setup-nightly-pipeline-stage
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1452
---

# Pull Request #1452: HMS-9727: setup Nightly pipeline (stage)

**Author**: @katarinazaprazna
**Created**: December 12, 2025 at 01:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `setup-nightly-pipeline-stage`

## Description

# Description
- Create a new GitHub workflow to run nightly (integration) tests on stage environment
- Port all integration test helpers from the content-sources repository and refactor them for clarity and simplicity
- Add a simple integration test to verify that our test setup is valid

Associated Jira ticket: https://issues.redhat.com/browse/HMS-9727

# Dependencies

- [x] Disable old workflows in GitHub Actions after PR merge (test.yml, playwright.yml, nightly-stage-test.yml renamed)
- [x] Clean up old workflow runs from Actions tab
- [x] Require `ci-checks` instead of `test` in required Status checks

Reminder: Do these after your PR with the renamed files is merged to master

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [x] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on December 12, 2025 at 01:48 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1452?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`27875c6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/27875c65535f2411f83b444b0a13b81cef5a068c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e66f107`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e66f10789886406a53eb8aaf161f95721bbe2b3b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1452   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      899    +7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1452/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1452/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `?` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1452/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `?` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1452?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Commented on December 15, 2025 at 12:37 PM UTC

### Review by @TenSt - Approved on December 23, 2025 at 10:18 AM UTC

lgtm

### Review by @dominikvagner - Approved on January 06, 2026 at 09:34 AM UTC

looks good! ✨ great job! 👏🏼🚀 

nitpick: I don't think the change of the job name from `test` to `ci-checks` is necessary if it breaks the branch rulesets, but those can be changed 💭😅 

_had some trouble getting the test to pass yesterday, but I guess it was some office networking shenanigans again, as it work today_ 🤷🏼😆  

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1452*
