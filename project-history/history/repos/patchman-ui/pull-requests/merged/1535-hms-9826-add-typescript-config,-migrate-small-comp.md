---
type: pull_request
number: 1535
title: "HMS-9826: add typescript config, migrate small component to TS"
state: merged
author: xbhouse
created: 2026-03-06T19:04:30Z
updated: 2026-03-30T13:43:35Z
closed: 2026-03-30T13:43:35Z
merged: 2026-03-30T13:43:35Z
base_branch: master
head_branch: 9826
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1535
---

# Pull Request #1535: HMS-9826: add typescript config, migrate small component to TS

**Author**: @xbhouse
**Created**: March 06, 2026 at 07:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `9826`

## Description

# Description

Associated Jira ticket: [HMS-9826](https://issues.redhat.com/browse/HMS-9826)

This PR adds initial Typescript support and migrates the CvesModal SmartComponent as a first example

- Installs additional TS packages and updates Jest, Babel, and ESlint configs to support TS. Splits the ESlint config into separate sections for JS and TS, different configs are needed for JS and TS components and Playwright tests
- Changes the `noImplicitAny` option in the tsconfig to `false` to ease the initial migration, but we should still aim to define types when it's possible
- Migrates the CvesModal SmartComponent and its related unit tests, API, and Redux components to TS
- Migrates some utilities used by these components, while leaving others in JS to demonstrate that JS modules can still be consumed by TS components during an incremental migration

# How to test the PR

- Run the app and ensure everything still works as expected, especially the CvesModal on the AdvisoryDetail page
- Tests should pass

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


[HMS-9826]: https://redhat.atlassian.net/browse/HMS-9826?atlOrigin=eyJpIjoiNWRkNTljNzYxNjVmNDY3MDlhMDU5Y2ZhYzA5YTRkZjUiLCJwIjoiZ2l0aHViLWNvbS1KU1cifQ

---

## Discussion

### Comment by @codecov-commenter on March 06, 2026 at 07:26 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `84.44444%` with `7 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`ea57ed9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ea57ed91accae1f4f9badcb6d8bb17f3c13faba1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a982dda`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a982ddac691916b4c93c808dd84a61dfe42ac399?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/api/vulnerabilityApi.ts](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?src=pr&el=tree&filepath=src%2FUtilities%2Fapi%2FvulnerabilityApi.ts&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkvdnVsbmVyYWJpbGl0eUFwaS50cw==) | 14.28% | [6 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/AdvisoryDetail/CvesModal.tsx](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisoryDetail%2FCvesModal.tsx&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeURldGFpbC9DdmVzTW9kYWwudHN4) | 94.73% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1535      +/-   ##
==========================================
+ Coverage   71.94%   75.88%   +3.94%     
==========================================
  Files          98      103       +5     
  Lines        2445     3164     +719     
  Branches      686      686              
==========================================
+ Hits         1759     2401     +642     
- Misses        603      684      +81     
+ Partials       83       79       -4     
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1535?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @katarinazaprazna - Dismissed on March 18, 2026 at 06:44 PM UTC

Really nice work on the TS migration! I've left a few small suggestions on specific lines for some minor improvements, but nothing blocking :)

### Review by @xbhouse - Commented on March 26, 2026 at 04:17 PM UTC

### Review by @xbhouse - Commented on March 26, 2026 at 04:18 PM UTC

### Review by @katarinazaprazna - Approved on March 26, 2026 at 11:20 PM UTC

Thank you for making these updates! The changes look great :)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1535*
