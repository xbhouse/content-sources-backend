---
type: pull_request
number: 2008
title: "Update module github.com/prometheus/common to v0.67.5"
state: merged
author: red-hat-konflux
created: 2026-01-12T04:42:59Z
updated: 2026-01-12T08:39:34Z
closed: 2026-01-12T04:49:06Z
merged: 2026-01-12T04:49:06Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-common-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2008
---

# Pull Request #2008: Update module github.com/prometheus/common to v0.67.5

**Author**: @red-hat-konflux
**Created**: January 12, 2026 at 04:42 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | `v0.67.4` -> `v0.67.5` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fcommon/v0.67.5?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fcommon/v0.67.4/v0.67.5?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.67.5`](https://redirect.github.com/prometheus/common/releases/tag/v0.67.5)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.4...v0.67.5)

#### What's Changed

- build(deps): bump golang.org/x/oauth2 from 0.32.0 to 0.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;871](https://redirect.github.com/prometheus/common/pull/871)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;866](https://redirect.github.com/prometheus/common/pull/866)
- build(deps): bump golang.org/x/net from 0.46.0 to 0.48.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;872](https://redirect.github.com/prometheus/common/pull/872)
- build(deps): bump google.golang.org/protobuf from 1.36.10 to 1.36.11 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;870](https://redirect.github.com/prometheus/common/pull/870)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.67.4...v0.67.5>

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

### Comment by @jira-linking on January 12, 2026 at 04:43 AM UTC

Commits missing Jira IDs:
33346a4828d59f401c5d5b31f48aea0b72fe1e18


### Comment by @codecov-commenter on January 12, 2026 at 04:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2008?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`dc69f88`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dc69f88a01d76a6123769bbebc8dceb71212763e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`33346a4`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/33346a4828d59f401c5d5b31f48aea0b72fe1e18?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2008   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2008/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2008/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2008?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2008*
