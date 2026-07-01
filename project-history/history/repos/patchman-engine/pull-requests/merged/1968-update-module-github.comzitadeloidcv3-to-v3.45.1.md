---
type: pull_request
number: 1968
title: "Update module github.com/zitadel/oidc/v3 to v3.45.1"
state: merged
author: red-hat-konflux
created: 2025-12-08T04:37:20Z
updated: 2025-12-29T00:35:55Z
closed: 2025-12-28T20:51:02Z
merged: 2025-12-28T20:51:02Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1968
---

# Pull Request #1968: Update module github.com/zitadel/oidc/v3 to v3.45.1

**Author**: @red-hat-konflux
**Created**: December 08, 2025 at 04:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/zitadel/oidc/v3](https://redirect.github.com/zitadel/oidc) | `v3.45.0` -> `v3.45.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fzitadel%2foidc%2fv3/v3.45.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fzitadel%2foidc%2fv3/v3.45.0/v3.45.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>zitadel/oidc (github.com/zitadel/oidc/v3)</summary>

### [`v3.45.1`](https://redirect.github.com/zitadel/oidc/releases/tag/v3.45.1)

[Compare Source](https://redirect.github.com/zitadel/oidc/compare/v3.45.0...v3.45.1)

##### Bug Fixes

- **rp:** don't ignore JWKS parsing errors ([#&#8203;771](https://redirect.github.com/zitadel/oidc/issues/771)) ([a3f3428](https://redirect.github.com/zitadel/oidc/commit/a3f34289fafbece952fb5e76bd26435c743afcee)), closes [#&#8203;541](https://redirect.github.com/zitadel/oidc/issues/541)

</details>

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on December 08, 2025 at 04:37 AM UTC

Commits missing Jira IDs:
77e6ddd064125419449293b1a2a24ff930a712ef


### Comment by @codecov-commenter on December 22, 2025 at 04:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1968?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`deb9cf6`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/deb9cf64184a89ca320feb0d18475b1426ebdd8b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`77e6ddd`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/77e6ddd064125419449293b1a2a24ff930a712ef?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1968   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1968/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1968/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1968?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1968*
