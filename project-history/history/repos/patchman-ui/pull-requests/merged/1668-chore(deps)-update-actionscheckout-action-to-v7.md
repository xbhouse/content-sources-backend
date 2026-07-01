---
type: pull_request
number: 1668
title: "chore(deps): update actions/checkout action to v7"
state: merged
author: red-hat-konflux
created: 2026-06-22T18:26:54Z
updated: 2026-06-22T21:57:06Z
closed: 2026-06-22T19:43:16Z
merged: 2026-06-22T19:43:16Z
base_branch: master
head_branch: konflux/mintmaker/master/actions-checkout-7.x
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1668
---

# Pull Request #1668: chore(deps): update actions/checkout action to v7

**Author**: @red-hat-konflux
**Created**: June 22, 2026 at 06:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/actions-checkout-7.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [actions/checkout](https://redirect.github.com/actions/checkout) | action | major | `v6` → `v7` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>actions/checkout (actions/checkout)</summary>

### [`v7`](https://redirect.github.com/actions/checkout/blob/HEAD/CHANGELOG.md#v700)

[Compare Source](https://redirect.github.com/actions/checkout/compare/v6...v7)

- Block checking out fork PR for pull\_request\_target and workflow\_run by [@&#8203;aiqiaoy](https://redirect.github.com/aiqiaoy) in [#&#8203;2454](https://redirect.github.com/actions/checkout/pull/2454)
- Bump actions/publish-immutable-action from 0.0.3 to 0.0.4 in the minor-actions-dependencies group across 1 directory by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2458](https://redirect.github.com/actions/checkout/pull/2458)
- Bump flatted from 3.3.1 to 3.4.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2460](https://redirect.github.com/actions/checkout/pull/2460)
- Bump js-yaml from 4.1.0 to 4.2.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2461](https://redirect.github.com/actions/checkout/pull/2461)
- Bump [@&#8203;actions/core](https://redirect.github.com/actions/core) and [@&#8203;actions/tool-cache](https://redirect.github.com/actions/tool-cache) and Remove uuid by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2459](https://redirect.github.com/actions/checkout/pull/2459)
- upgrade module to esm and update dependencies by [@&#8203;aiqiaoy](https://redirect.github.com/aiqiaoy) in [#&#8203;2463](https://redirect.github.com/actions/checkout/pull/2463)
- Bump the minor-npm-dependencies group across 1 directory with 3 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2462](https://redirect.github.com/actions/checkout/pull/2462)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on June 22, 2026 at 06:30 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1668?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.27%. Comparing base ([`49d0fb1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/49d0fb1e3781c62e9029b9593a6cc546021cf53b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2a5cc13`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2a5cc134ec23a5ba5af2d7e35b5089f1dd08000d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 11 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1668   +/-   ##
=======================================
  Coverage   77.27%   77.27%           
=======================================
  Files         103      103           
  Lines        3287     3287           
  Branches      740      735    -5     
=======================================
  Hits         2540     2540           
  Misses        668      668           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1668?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @github-actions - Approved on June 22, 2026 at 06:27 PM UTC

This PR from Konflux has been automatically approved.

### Review by @TenSt - Approved on June 22, 2026 at 07:42 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1668*
