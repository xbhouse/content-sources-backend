---
type: pull_request
number: 2164
title: "Update grafana/grafana Docker tag to v13"
state: merged
author: red-hat-konflux
created: 2026-04-20T01:25:02Z
updated: 2026-04-20T09:33:51Z
closed: 2026-04-20T05:30:14Z
merged: 2026-04-20T05:30:14Z
base_branch: master
head_branch: konflux/mintmaker/master/major-grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2164
---

# Pull Request #2164: Update grafana/grafana Docker tag to v13

**Author**: @red-hat-konflux
**Created**: April 20, 2026 at 01:25 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/major-grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | major | `12.4.3` → `13.0.1` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v13.0.1`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1301-2026-04-17)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.4.3...v13.0.1)

##### Features and enhancements

- **Dashboard:** Preserve timezone user-preference when converting V1 → V2 [#&#8203;122673](https://redirect.github.com/grafana/grafana/pull/122673), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Provisioning:** Include dashboard validation errors in pull request comments [#&#8203;122433](https://redirect.github.com/grafana/grafana/pull/122433), [@&#8203;gttrigger](https://redirect.github.com/gttrigger)

##### Bug fixes

- **Unified storage:** Skip migrations if dualwrite state shows they were already migrated [#&#8203;122880](https://redirect.github.com/grafana/grafana/pull/122880), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)

<!-- 13.0.1 END -->

<!-- 12.4.3 START -->

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

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2164?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.13%. Comparing base ([`f8ba638`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f8ba63821d7f0e4433dd9378aa87e1aeb77aa29a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`68bdf48`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/68bdf484c22a92a90815afa546e606935c70d7c3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #2164      +/-   ##
==========================================
+ Coverage   59.09%   59.13%   +0.03%     
==========================================
  Files         134      134              
  Lines        8738     8738              
==========================================
+ Hits         5164     5167       +3     
+ Misses       3031     3028       -3     
  Partials      543      543              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2164/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2164/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.13% <ø> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2164?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2164*
