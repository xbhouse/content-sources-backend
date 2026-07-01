---
type: pull_request
number: 2016
title: "Update golang.org/x/exp digest to 716be56"
state: merged
author: red-hat-konflux
created: 2026-01-19T08:23:07Z
updated: 2026-01-22T11:36:14Z
closed: 2026-01-22T11:36:14Z
merged: 2026-01-22T11:36:14Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-exp-digest
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2016
---

# Pull Request #2016: Update golang.org/x/exp digest to 716be56

**Author**: @red-hat-konflux
**Created**: January 19, 2026 at 08:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-exp-digest`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| golang.org/x/exp | require | digest | `944ab1f` -> `716be56` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

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

### Comment by @red-hat-konflux on January 19, 2026 at 08:23 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 4 additional dependencies were updated


Details:


| **Package**           | **Change**             |
| :-------------------- | :--------------------- |
| `golang.org/x/crypto` | `v0.46.0` -> `v0.47.0` |
| `golang.org/x/mod`    | `v0.31.0` -> `v0.32.0` |
| `golang.org/x/net`    | `v0.48.0` -> `v0.49.0` |
| `golang.org/x/tools`  | `v0.40.0` -> `v0.41.0` |

### Comment by @jira-linking on January 19, 2026 at 08:23 AM UTC

Commits missing Jira IDs:
e3bd81663724439a8b87a3d637940fbdb35d35b9


### Comment by @codecov-commenter on January 19, 2026 at 08:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2016?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`d0195ac`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d0195ac37be0e97415dbb6b2136930d924aa3359?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e3bd816`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e3bd81663724439a8b87a3d637940fbdb35d35b9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2016   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2016/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2016/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2016?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @TenSt on January 22, 2026 at 11:32 AM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on January 21, 2026 at 12:46 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2016*
