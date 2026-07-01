---
type: pull_request
number: 1990
title: "Update grafana/grafana Docker tag to v12.3.1"
state: merged
author: red-hat-konflux
created: 2025-12-22T04:41:50Z
updated: 2025-12-22T08:40:28Z
closed: 2025-12-22T04:47:09Z
merged: 2025-12-22T04:47:09Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1990
---

# Pull Request #1990: Update grafana/grafana Docker tag to v12.3.1

**Author**: @red-hat-konflux
**Created**: December 22, 2025 at 04:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | patch | `12.3.0` -> `12.3.1` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.3.1`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1231-2025-12-16)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.3.0...v12.3.1)

##### Features and enhancements

- **Alerting:** Update alerting dependency [#&#8203;114259](https://redirect.github.com/grafana/grafana/pull/114259), [@&#8203;moustafab](https://redirect.github.com/moustafab)
- **Azure:** Improved column handling in logs query builder [#&#8203;114841](https://redirect.github.com/grafana/grafana/pull/114841), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Include aggregate columns in logs builder [#&#8203;114835](https://redirect.github.com/grafana/grafana/pull/114835), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Dependencies:** Bump Go to v1.25.5 [#&#8203;114751](https://redirect.github.com/grafana/grafana/pull/114751), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Docs:** Clarify section title for repeating rows and tabs [#&#8203;115346](https://redirect.github.com/grafana/grafana/pull/115346), [@&#8203;imatwawana](https://redirect.github.com/imatwawana)
- **Plugins:** Add PluginContext to plugins when scenes is disabled [#&#8203;115064](https://redirect.github.com/grafana/grafana/pull/115064), [@&#8203;hugohaggmark](https://redirect.github.com/hugohaggmark)
- **QueryEditorRows:** Clear hideSeriesFrom override on query edit [#&#8203;114628](https://redirect.github.com/grafana/grafana/pull/114628), [@&#8203;Sergej-Vlasov](https://redirect.github.com/Sergej-Vlasov)

##### Bug fixes

- **Azure:** Fix `dcount` aggregation [#&#8203;114907](https://redirect.github.com/grafana/grafana/pull/114907), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Azure:** Fix `percentile` syntax [#&#8203;114707](https://redirect.github.com/grafana/grafana/pull/114707), [@&#8203;aangelisc](https://redirect.github.com/aangelisc)
- **Dashboards:** Fix empty space under time controls when a dashboard has a lot of variables [#&#8203;114730](https://redirect.github.com/grafana/grafana/pull/114730), [@&#8203;oscarkilhed](https://redirect.github.com/oscarkilhed)
- **Plugins:** Datasource breadcrumb link should link to settings tab [#&#8203;113910](https://redirect.github.com/grafana/grafana/pull/113910), [@&#8203;wbrowne](https://redirect.github.com/wbrowne)
- **Postgresql:** Fix variable interpolation logic when the variable has multiple values [#&#8203;114876](https://redirect.github.com/grafana/grafana/pull/114876), [@&#8203;itsmylife](https://redirect.github.com/itsmylife)

<!-- 12.3.1 END -->

<!-- 12.2.3 START -->

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on December 22, 2025 at 04:41 AM UTC

Commits missing Jira IDs:
da49d44ee0fc9b01c9d3854fd1c4c19ef0e308a9


### Comment by @codecov-commenter on December 22, 2025 at 04:47 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1990?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`e90b3ef`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e90b3ef44c967d50ee185921e94081b25e639c62?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`da49d44`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/da49d44ee0fc9b01c9d3854fd1c4c19ef0e308a9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1990   +/-   ##
=======================================
  Coverage   59.01%   59.01%           
=======================================
  Files         131      131           
  Lines        8493     8493           
=======================================
  Hits         5012     5012           
  Misses       2947     2947           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1990/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1990/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1990?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1990*
