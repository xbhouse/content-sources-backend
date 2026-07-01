---
type: pull_request
number: 1568
title: "chore: update @currents/playwright to address security compromise"
state: merged
author: katarinazaprazna
created: 2026-03-31T14:10:19Z
updated: 2026-03-31T14:20:03Z
closed: 2026-03-31T14:20:03Z
merged: 2026-03-31T14:20:03Z
base_branch: master
head_branch: bump-playwright-version
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1568
---

# Pull Request #1568: chore: update @currents/playwright to address security compromise

**Author**: @katarinazaprazna
**Created**: March 31, 2026 at 02:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `bump-playwright-version`

## Description

# Description

- Upgrade `@currents/playwright` from `1.22.1` to `1.22.3` to address supply chain security compromise.
- The maintainers have published version `1.22.3` as a clean, safe release.

CHANGELOG.md: https://www.npmjs.com/package/@currents/playwright/v/1.22.3?activeTab=code

Confirmed with the maintainers:
<img width="1253" height="752" alt="Screenshot From 2026-03-31 15-54-28" src="https://github.com/user-attachments/assets/90c751b2-f540-4e52-ba07-4b08ee48ddfe" />

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on March 31, 2026 at 02:13 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1568?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`f0de562`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f0de562ff60ce1facb0e33840281cddec75502eb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b39cabb`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b39cabb84c2c129f1d15c7a72cc5ca511becef8a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1568   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      685      686    +1     
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1568?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @dominikvagner - Approved on March 31, 2026 at 02:14 PM UTC

ack! thx! ✅ 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1568*
