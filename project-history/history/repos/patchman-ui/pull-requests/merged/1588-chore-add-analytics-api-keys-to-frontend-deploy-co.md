---
type: pull_request
number: 1588
title: "chore: add analytics API keys to frontend deploy config"
state: merged
author: dajohnso
created: 2026-04-22T18:28:40Z
updated: 2026-05-08T09:06:35Z
closed: 2026-05-08T09:06:35Z
merged: 2026-05-08T09:06:35Z
base_branch: master
head_branch: add-analytics-keys
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1588
---

# Pull Request #1588: chore: add analytics API keys to frontend deploy config

**Author**: @dajohnso
**Created**: April 22, 2026 at 06:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `add-analytics-keys`

## Description

## Summary
- Adds the `analytics` block with `APIKey` and `APIKeyDev` to the `module` section in `deploy/frontend.yml`
- Aligns with the pattern used by other frontends (inventory, advisor, tasks, ocp-advisor)
- Fixes RHINENG-25987
## Test plan
- [ ] Verify the YAML is valid and the deploy CRD parses correctly
- [ ] Confirm analytics events are sent in stage with the correct API key

---

## Discussion

### Comment by @codecov-commenter on May 01, 2026 at 09:03 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1588?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.38%. Comparing base ([`d0b7ce9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d0b7ce998053563899f78535b6780930ecda244e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c49a708`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c49a7089a04b148eb6b24c023369d766d825019b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1588   +/-   ##
=======================================
  Coverage   76.38%   76.38%           
=======================================
  Files         103      103           
  Lines        3185     3185           
  Branches      691      692    +1     
=======================================
  Hits         2433     2433           
  Misses        674      674           
  Partials       78       78           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1588?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on May 08, 2026 at 08:58 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1588*
