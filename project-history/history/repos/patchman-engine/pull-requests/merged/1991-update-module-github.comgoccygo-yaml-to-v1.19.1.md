---
type: pull_request
number: 1991
title: "Update module github.com/goccy/go-yaml to v1.19.1"
state: merged
author: red-hat-konflux
created: 2025-12-22T04:42:08Z
updated: 2025-12-22T08:40:26Z
closed: 2025-12-22T04:48:12Z
merged: 2025-12-22T04:48:12Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-goccy-go-yaml-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1991
---

# Pull Request #1991: Update module github.com/goccy/go-yaml to v1.19.1

**Author**: @red-hat-konflux
**Created**: December 22, 2025 at 04:42 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-goccy-go-yaml-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/goccy/go-yaml](https://redirect.github.com/goccy/go-yaml) | `v1.19.0` -> `v1.19.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgoccy%2fgo-yaml/v1.19.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgoccy%2fgo-yaml/v1.19.0/v1.19.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>goccy/go-yaml (github.com/goccy/go-yaml)</summary>

### [`v1.19.1`](https://redirect.github.com/goccy/go-yaml/releases/tag/v1.19.1): 1.19.1

[Compare Source](https://redirect.github.com/goccy/go-yaml/compare/v1.19.0...v1.19.1)

#### What's Changed

- Fix decoding of integer keys of map type by [@&#8203;goccy](https://redirect.github.com/goccy) in [#&#8203;829](https://redirect.github.com/goccy/go-yaml/pull/829)
- Support line comment for flow sequence or flow map by [@&#8203;goccy](https://redirect.github.com/goccy) in [#&#8203;834](https://redirect.github.com/goccy/go-yaml/pull/834)

**Full Changelog**: <https://github.com/goccy/go-yaml/compare/v1.19.0...v1.19.1>

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

### Comment by @jira-linking on December 22, 2025 at 04:42 AM UTC

Commits missing Jira IDs:
e3ed909a74ac19052034b3ce183ed055375c192d


### Comment by @codecov-commenter on December 22, 2025 at 04:47 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1991?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`e90b3ef`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e90b3ef44c967d50ee185921e94081b25e639c62?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e3ed909`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e3ed909a74ac19052034b3ce183ed055375c192d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1991   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1991/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1991/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1991?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1991*
