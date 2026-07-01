---
type: pull_request
number: 1927
title: "Update golang.org/x/exp digest to 944ab1f"
state: merged
author: red-hat-konflux
created: 2025-11-17T04:41:04Z
updated: 2025-12-22T08:46:06Z
closed: 2025-12-22T08:45:25Z
merged: 2025-12-22T08:45:25Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-exp-digest
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1927
---

# Pull Request #1927: Update golang.org/x/exp digest to 944ab1f

**Author**: @red-hat-konflux
**Created**: November 17, 2025 at 04:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-exp-digest`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| golang.org/x/exp | require | digest | `a4bb9ff` -> `944ab1f` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

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

### Comment by @red-hat-konflux on November 17, 2025 at 04:41 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 7 additional dependencies were updated


Details:


| **Package**           | **Change**                                           |
| :-------------------- | :--------------------------------------------------- |
| `golang.org/x/crypto` | `v0.43.0` -> `v0.46.0`                               |
| `golang.org/x/mod`    | `v0.30.0` -> `v0.31.0`                               |
| `golang.org/x/net`    | `v0.46.1-0.20251013234738-63d1a5100f82` -> `v0.48.0` |
| `golang.org/x/sync`   | `v0.18.0` -> `v0.19.0`                               |
| `golang.org/x/sys`    | `v0.38.0` -> `v0.39.0`                               |
| `golang.org/x/text`   | `v0.31.0` -> `v0.32.0`                               |
| `golang.org/x/tools`  | `v0.38.0` -> `v0.40.0`                               |

### Comment by @jira-linking on December 06, 2025 at 12:41 AM UTC

Commits missing Jira IDs:
8ebfa9f70fb109b361bf8ec7f2025b3e307ede88


### Comment by @codecov-commenter on December 22, 2025 at 08:45 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1927?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`5df8d34`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5df8d34f2ffe7572d21ed4ad6e74e0d240978ece?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`8ebfa9f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8ebfa9f70fb109b361bf8ec7f2025b3e307ede88?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1927   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1927/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1927/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1927?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1927*
