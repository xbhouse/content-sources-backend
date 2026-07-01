---
type: pull_request
number: 1957
title: "Update module github.com/goccy/go-yaml to v1.19.0"
state: merged
author: red-hat-konflux
created: 2025-12-01T04:41:23Z
updated: 2025-12-01T08:41:48Z
closed: 2025-12-01T04:49:11Z
merged: 2025-12-01T04:49:11Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-goccy-go-yaml-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1957
---

# Pull Request #1957: Update module github.com/goccy/go-yaml to v1.19.0

**Author**: @red-hat-konflux
**Created**: December 01, 2025 at 04:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-goccy-go-yaml-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/goccy/go-yaml](https://redirect.github.com/goccy/go-yaml) | `v1.18.0` -> `v1.19.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgoccy%2fgo-yaml/v1.19.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgoccy%2fgo-yaml/v1.18.0/v1.19.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>goccy/go-yaml (github.com/goccy/go-yaml)</summary>

### [`v1.19.0`](https://redirect.github.com/goccy/go-yaml/releases/tag/v1.19.0): 1.19.0

[Compare Source](https://redirect.github.com/goccy/go-yaml/compare/v1.18.0...v1.19.0)

#### What's Changed

- Revert "feat: Dont make copies of structs for validation" by [@&#8203;shuheiktgw](https://redirect.github.com/shuheiktgw) in [#&#8203;763](https://redirect.github.com/goccy/go-yaml/pull/763)
- Add decode option that allows specific field prefixes by [@&#8203;cpuguy83](https://redirect.github.com/cpuguy83) in [#&#8203;795](https://redirect.github.com/goccy/go-yaml/pull/795)
- Normalize CR and CRLF in multi-line strings by [@&#8203;shuheiktgw](https://redirect.github.com/shuheiktgw) in [#&#8203;754](https://redirect.github.com/goccy/go-yaml/pull/754)
- Support non string map keys by [@&#8203;shuheiktgw](https://redirect.github.com/shuheiktgw) in [#&#8203;756](https://redirect.github.com/goccy/go-yaml/pull/756)
- Skip directive in path operations by [@&#8203;shuheiktgw](https://redirect.github.com/shuheiktgw) in [#&#8203;758](https://redirect.github.com/goccy/go-yaml/pull/758)
- Add indentation to flow values on new lines by [@&#8203;shuheiktgw](https://redirect.github.com/shuheiktgw) in [#&#8203;759](https://redirect.github.com/goccy/go-yaml/pull/759)
- Add support for RawMessage, similar to json.RawMessage by [@&#8203;thanethomson](https://redirect.github.com/thanethomson) in [#&#8203;790](https://redirect.github.com/goccy/go-yaml/pull/790)

#### New Contributors

- [@&#8203;cpuguy83](https://redirect.github.com/cpuguy83) made their first contribution in [#&#8203;795](https://redirect.github.com/goccy/go-yaml/pull/795)
- [@&#8203;thanethomson](https://redirect.github.com/thanethomson) made their first contribution in [#&#8203;790](https://redirect.github.com/goccy/go-yaml/pull/790)

**Full Changelog**: <https://github.com/goccy/go-yaml/compare/v1.18.0...v1.19.0>

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

### Comment by @codecov-commenter on December 01, 2025 at 04:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1957?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.83%. Comparing base ([`29b134c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/29b134cd85fb4993e8bc7ab6e6d006aedb8ff205?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`7e7715f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/7e7715f66081b1c7a6bcfaa5a0e5cff3542ceced?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1957      +/-   ##
==========================================
+ Coverage   58.79%   58.83%   +0.03%     
==========================================
  Files         131      131              
  Lines        8407     8407              
==========================================
+ Hits         4943     4946       +3     
+ Misses       2930     2927       -3     
  Partials      534      534              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1957/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1957/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.83% <ø> (+0.03%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1957?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1957*
