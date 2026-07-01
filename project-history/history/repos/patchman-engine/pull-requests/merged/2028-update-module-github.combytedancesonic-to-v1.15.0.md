---
type: pull_request
number: 2028
title: "Update module github.com/bytedance/sonic to v1.15.0"
state: merged
author: red-hat-konflux
created: 2026-01-26T08:43:43Z
updated: 2026-01-26T12:38:42Z
closed: 2026-01-26T08:49:29Z
merged: 2026-01-26T08:49:29Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-bytedance-sonic-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2028
---

# Pull Request #2028: Update module github.com/bytedance/sonic to v1.15.0

**Author**: @red-hat-konflux
**Created**: January 26, 2026 at 08:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-bytedance-sonic-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/bytedance/sonic](https://redirect.github.com/bytedance/sonic) | `v1.14.2` -> `v1.15.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fbytedance%2fsonic/v1.15.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fbytedance%2fsonic/v1.14.2/v1.15.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>bytedance/sonic (github.com/bytedance/sonic)</summary>

### [`v1.15.0`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.15.0)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.14.2...v1.15.0)

#### What's Changed

- chore: benchmark with msgpack by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;881](https://redirect.github.com/bytedance/sonic/pull/881)
- opt: ast node set by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;883](https://redirect.github.com/bytedance/sonic/pull/883)
- chore: update Go CI by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [#&#8203;886](https://redirect.github.com/bytedance/sonic/pull/886)
- feat: support Go 1.26 by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [#&#8203;898](https://redirect.github.com/bytedance/sonic/pull/898)
- chore: update go mod by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [#&#8203;899](https://redirect.github.com/bytedance/sonic/pull/899)

**Full Changelog**: <https://github.com/bytedance/sonic/compare/v1.14.2...v1.15.0>

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

### Comment by @red-hat-konflux on January 26, 2026 at 08:43 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                         | **Change**           |
| :---------------------------------- | :------------------- |
| `github.com/bytedance/sonic/loader` | `v0.4.0` -> `v0.5.0` |

### Comment by @jira-linking on January 26, 2026 at 08:43 AM UTC

Commits missing Jira IDs:
0c79bda879bc0ac59d484b3c402532ecbf75cf7a


### Comment by @codecov-commenter on January 26, 2026 at 08:50 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2028?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`03a8100`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/03a8100da6dfe3aaade5da9eac2b1abd29a588f1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0c79bda`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0c79bda879bc0ac59d484b3c402532ecbf75cf7a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 8 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2028   +/-   ##
=======================================
  Coverage   59.18%   59.18%           
=======================================
  Files         133      133           
  Lines        8599     8599           
=======================================
  Hits         5089     5089           
  Misses       2967     2967           
  Partials      543      543           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2028/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2028/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2028?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2028*
