---
type: pull_request
number: 1424
title: "chore: Update dependabot to ignore jest, pf and rh cloud services"
state: merged
author: avitova
created: 2025-10-29T13:43:42Z
updated: 2026-01-07T21:12:44Z
closed: 2026-01-07T21:12:44Z
merged: 2026-01-07T21:12:44Z
base_branch: master
head_branch: dependabot
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1424
---

# Pull Request #1424: chore: Update dependabot to ignore jest, pf and rh cloud services

**Author**: @avitova
**Created**: October 29, 2025 at 01:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `dependabot`

## Description

# Description

I copied the settings of [Dependabot](https://github.com/content-services/content-sources-frontend/blob/main/.github/dependabot.yml) for Content. The patternfly, react-router dom and rh cloud services use higher versions. I think this should also be updated in content, as you already migrated pf, react-router-dom and rh services to higher versions than you ignore.

---

## Discussion

### Comment by @codecov-commenter on October 29, 2025 at 01:46 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1424?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`f71f815`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f71f815cd82d91bd29576b6ab1a5282e542fc73e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`854015a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/854015a65dd1ad0ac1add9750c5f8931c4b46f70?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1424   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1424?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @marusak on October 30, 2025 at 02:09 PM UTC

In Content we also have https://github.com/content-services/content-sources-frontend/blob/main/.github/dependabot.yml#L35 which makes sure, we get one PR for everything and not bunch of small ones. Should we also add something like that?

### Comment by @avitova on November 06, 2025 at 02:17 PM UTC

yessir, you ask for it, you have it

### Comment by @avitova on November 20, 2025 at 06:00 PM UTC

Yes, thank you!:)

### Comment by @TenSt on December 23, 2025 at 10:21 AM UTC

Hi @avitova 👋 
Could you please update this PR and merge it?

---

## Reviews

### Review by @LightOfHeaven1994 - Approved on November 20, 2025 at 09:26 AM UTC

LGTM, there is an issue with commit lint so `test` pipeline fails. Can you fix that? 🙏 

### Review by @TenSt - Approved on December 11, 2025 at 02:06 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1424*
