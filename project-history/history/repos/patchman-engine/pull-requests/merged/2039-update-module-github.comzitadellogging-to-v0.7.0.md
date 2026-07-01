---
type: pull_request
number: 2039
title: "Update module github.com/zitadel/logging to v0.7.0"
state: merged
author: red-hat-konflux
created: 2026-02-02T08:48:50Z
updated: 2026-02-02T12:48:19Z
closed: 2026-02-02T08:54:30Z
merged: 2026-02-02T08:54:29Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-zitadel-logging-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2039
---

# Pull Request #2039: Update module github.com/zitadel/logging to v0.7.0

**Author**: @red-hat-konflux
**Created**: February 02, 2026 at 08:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-zitadel-logging-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/zitadel/logging](https://redirect.github.com/zitadel/logging) | `v0.6.2` -> `v0.7.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fzitadel%2flogging/v0.7.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fzitadel%2flogging/v0.6.2/v0.7.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>zitadel/logging (github.com/zitadel/logging)</summary>

### [`v0.7.0`](https://redirect.github.com/zitadel/logging/releases/tag/v0.7.0)

[Compare Source](https://redirect.github.com/zitadel/logging/compare/v0.6.2...v0.7.0)

##### Features

- add deprecation message ([#&#8203;45](https://redirect.github.com/zitadel/logging/issues/45)) ([75f2cfb](https://redirect.github.com/zitadel/logging/commit/75f2cfbd2481d11408c52467dc58b2c669ea34fd))

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

### Comment by @red-hat-konflux on February 02, 2026 at 08:48 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- The `go` directive was updated for compatibility reasons


Details:


| **Package** | **Change**            |
| :---------- | :-------------------- |
| `go`        | `1.24.4` -> `1.24.10` |

### Comment by @jira-linking on February 02, 2026 at 08:48 AM UTC

Commits missing Jira IDs:
db72a6f71ba08e480d60e5737ee5cad6e5ffc0ae


### Comment by @codecov-commenter on February 02, 2026 at 08:54 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2039?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.25%. Comparing base ([`92592c9`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/92592c95755f9762d74e9f8cc68fac318de73333?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`db72a6f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/db72a6f71ba08e480d60e5737ee5cad6e5ffc0ae?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2039   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2039/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2039/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2039?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2039*
