---
type: pull_request
number: 1994
title: "Update module golang.org/x/crypto to v0.46.0 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-12-22T08:40:10Z
updated: 2025-12-22T12:36:59Z
closed: 2025-12-22T12:36:57Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-crypto-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1994
---

# Pull Request #1994: Update module golang.org/x/crypto to v0.46.0 - autoclosed

**Author**: @red-hat-konflux
**Created**: December 22, 2025 at 08:40 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-crypto-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| golang.org/x/crypto | `v0.43.0` -> `v0.46.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/golang.org%2fx%2fcrypto/v0.46.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/golang.org%2fx%2fcrypto/v0.43.0/v0.46.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

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

### Comment by @red-hat-konflux on December 22, 2025 at 08:40 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 5 additional dependencies were updated


Details:


| **Package**          | **Change**                                           |
| :------------------- | :--------------------------------------------------- |
| `golang.org/x/net`   | `v0.46.1-0.20251013234738-63d1a5100f82` -> `v0.47.0` |
| `golang.org/x/sync`  | `v0.18.0` -> `v0.19.0`                               |
| `golang.org/x/sys`   | `v0.38.0` -> `v0.39.0`                               |
| `golang.org/x/text`  | `v0.31.0` -> `v0.32.0`                               |
| `golang.org/x/tools` | `v0.38.0` -> `v0.39.0`                               |

### Comment by @jira-linking on December 22, 2025 at 08:40 AM UTC

Commits missing Jira IDs:
98db848cbb3ddebc16d71743292558a68061b05e


### Comment by @codecov-commenter on December 22, 2025 at 08:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1994?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`5df8d34`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5df8d34f2ffe7572d21ed4ad6e74e0d240978ece?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`98db848`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/98db848cbb3ddebc16d71743292558a68061b05e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1994   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1994/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1994/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1994?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1994*
