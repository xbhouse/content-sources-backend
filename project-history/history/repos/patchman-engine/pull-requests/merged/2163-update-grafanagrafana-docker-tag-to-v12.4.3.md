---
type: pull_request
number: 2163
title: "Update grafana/grafana Docker tag to v12.4.3"
state: merged
author: red-hat-konflux
created: 2026-04-20T01:24:44Z
updated: 2026-04-20T05:30:10Z
closed: 2026-04-20T01:25:05Z
merged: 2026-04-20T01:25:05Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2163
---

# Pull Request #2163: Update grafana/grafana Docker tag to v12.4.3

**Author**: @red-hat-konflux
**Created**: April 20, 2026 at 01:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | patch | `12.4.2` → `12.4.3` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.4.3`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1243-2026-04-14)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.4.2...v12.4.3)

##### Features and enhancements

- **Analytics:** Keep internal dashboard id [#&#8203;121417](https://redirect.github.com/grafana/grafana/pull/121417), [@&#8203;ryantxu](https://redirect.github.com/ryantxu)
- **Go:** Update to 1.25.9 [#&#8203;122095](https://redirect.github.com/grafana/grafana/pull/122095), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Reporting:** Correctly apply appSubURL to report settings requests (Enterprise)

##### Bug fixes

- **Alerting:** Document Grafana HA Alertmanager cluster metrics prefix change in Grafana 12.4 [#&#8203;121481](https://redirect.github.com/grafana/grafana/pull/121481), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)

<!-- 12.4.3 END -->

<!-- 13.0.0 START -->

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on April 20, 2026 at 01:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2163?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.13%. Comparing base ([`ec8a02a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ec8a02a3d869947287f93004fd7844060b4226ae?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d74f15e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d74f15ee5c233dcb4101b5cbdf7147151b9eb31f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2163   +/-   ##
=======================================
  Coverage   59.13%   59.13%           
=======================================
  Files         134      134           
  Lines        8738     8738           
=======================================
  Hits         5167     5167           
  Misses       3028     3028           
  Partials      543      543           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2163/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2163/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.13% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2163?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2163*
