---
type: pull_request
number: 2036
title: "Update grafana/grafana Docker tag to v12.3.2"
state: merged
author: red-hat-konflux
created: 2026-02-02T04:47:37Z
updated: 2026-02-02T08:48:53Z
closed: 2026-02-02T04:53:05Z
merged: 2026-02-02T04:53:05Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2036
---

# Pull Request #2036: Update grafana/grafana Docker tag to v12.3.2

**Author**: @red-hat-konflux
**Created**: February 02, 2026 at 04:47 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | patch | `12.3.1` -> `12.3.2` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.3.2`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1232-2026-01-27)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.3.1...v12.3.2)

##### Features and enhancements

- **API:** Add missing scope check on dashboards [#&#8203;116888](https://redirect.github.com/grafana/grafana/pull/116888), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
- **Avatar:** Require sign-in, remove queue, respect timeout [#&#8203;116893](https://redirect.github.com/grafana/grafana/pull/116893), [@&#8203;macabu](https://redirect.github.com/macabu)
- **ElasticSearch:** Update annotation time-range properties [#&#8203;115566](https://redirect.github.com/grafana/grafana/pull/115566), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Explore:** Reset legend when a new query is run [#&#8203;116590](https://redirect.github.com/grafana/grafana/pull/116590), [@&#8203;ifrost](https://redirect.github.com/ifrost)
- **Go:** Update to 1.25.6 [#&#8203;116396](https://redirect.github.com/grafana/grafana/pull/116396), [@&#8203;macabu](https://redirect.github.com/macabu)

##### Bug fixes

- **Alerting:** Fix a race condition panic in ResetStateByRuleUID [#&#8203;115680](https://redirect.github.com/grafana/grafana/pull/115680), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)
- **Alerting:** Fix data source recording rules editor [#&#8203;116303](https://redirect.github.com/grafana/grafana/pull/116303), [@&#8203;soniaAguilarPeiron](https://redirect.github.com/soniaAguilarPeiron)

<!-- 12.3.2 END -->

<!-- 12.2.4 START -->

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on February 02, 2026 at 04:47 AM UTC

Commits missing Jira IDs:
ae22820b3dfbdb3d05e983fd944a7b916066eca6


### Comment by @codecov-commenter on February 02, 2026 at 04:53 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2036?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.25%. Comparing base ([`8250aa8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8250aa872cd37966db751390e4f9591f28ff07ff?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ae22820`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ae22820b3dfbdb3d05e983fd944a7b916066eca6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2036   +/-   ##
=======================================
  Coverage   59.25%   59.25%           
=======================================
  Files         134      134           
  Lines        8615     8615           
=======================================
  Hits         5105     5105           
  Misses       2967     2967           
  Partials      543      543           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2036/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2036/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2036?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2036*
