---
type: pull_request
number: 1461
title: "fix(SystemDetail, Systems): Fix broken link to content templates (HMS-9991)"
state: closed
author: ochosi
created: 2026-01-08T07:29:52Z
updated: 2026-01-08T08:01:48Z
closed: 2026-01-08T08:01:47Z
base_branch: master
head_branch: fix-content-template-links
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1461
---

# Pull Request #1461: fix(SystemDetail, Systems): Fix broken link to content templates (HMS-9991)

**Author**: @ochosi
**Created**: January 08, 2026 at 07:29 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix-content-template-links`

## Description

# Description

Associated Jira ticket: HMS-9991

The` /details` target does not exist and leads users back to the overview of templates. By removing it, the user gets redirected to the default location, which currently is `/systems`.
By not explicitly making it systems, we retain flexibility to change the default page later without having to update these links.

<img width="1237" height="702" alt="image" src="https://github.com/user-attachments/assets/75d584a0-ef4a-4b99-9525-4c1c32372726" />


# How to test the PR

Navigate to the Systems page in Content and select a single system.
Click on the link to the associated content template. You will be redirected to the overview of templates as opposed to the details of the single, selected template.

# Before the change

Users get redirected to the list of templates.

# After the change

Users get to the template in question.

# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on January 08, 2026 at 07:36 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1461?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `0%` with `3 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`b98d888`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b98d888b805a79e216602241d58113564c337cf2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ca49983`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ca4998302d20e4a829aa2ee6f8e27dd01dbd3291?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1461?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1461?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsListAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1461?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1461   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      892           
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1461?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @ochosi on January 08, 2026 at 08:01 AM UTC

Closed in favor of #1460 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1461*
