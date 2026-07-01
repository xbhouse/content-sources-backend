---
type: pull_request
number: 2027
title: "Update module github.com/zitadel/oidc/v3 to v3.45.3"
state: merged
author: red-hat-konflux
created: 2026-01-26T08:43:29Z
updated: 2026-01-26T12:38:43Z
closed: 2026-01-26T08:49:00Z
merged: 2026-01-26T08:49:00Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2027
---

# Pull Request #2027: Update module github.com/zitadel/oidc/v3 to v3.45.3

**Author**: @red-hat-konflux
**Created**: January 26, 2026 at 08:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/zitadel/oidc/v3](https://redirect.github.com/zitadel/oidc) | `v3.45.1` -> `v3.45.3` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fzitadel%2foidc%2fv3/v3.45.3?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fzitadel%2foidc%2fv3/v3.45.1/v3.45.3?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>zitadel/oidc (github.com/zitadel/oidc/v3)</summary>

### [`v3.45.3`](https://redirect.github.com/zitadel/oidc/releases/tag/v3.45.3)

[Compare Source](https://redirect.github.com/zitadel/oidc/compare/v3.45.2...v3.45.3)

##### Bug Fixes

- **tracing:** renamed tracer in op ([#&#8203;809](https://redirect.github.com/zitadel/oidc/issues/809)) ([d105faf](https://redirect.github.com/zitadel/oidc/commit/d105fafa376430fb408e7f3777857ae1eedf575e)), closes [#&#8203;808](https://redirect.github.com/zitadel/oidc/issues/808)

### [`v3.45.2`](https://redirect.github.com/zitadel/oidc/releases/tag/v3.45.2)

[Compare Source](https://redirect.github.com/zitadel/oidc/compare/v3.45.1...v3.45.2)

##### Bug Fixes

- consistently handle string-valued boolean fields from non-compliant OIDC providers ([#&#8203;791](https://redirect.github.com/zitadel/oidc/issues/791)) ([b4dca67](https://redirect.github.com/zitadel/oidc/commit/b4dca67d3c70b50162ae6b40acc2f124cc591257)), closes [#&#8203;139](https://redirect.github.com/zitadel/oidc/issues/139)

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

### Comment by @jira-linking on January 26, 2026 at 08:43 AM UTC

Commits missing Jira IDs:
0734fd7778e95cd721161aa54c7dbb723356f044


### Comment by @codecov-commenter on January 26, 2026 at 08:49 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2027?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`03a8100`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/03a8100da6dfe3aaade5da9eac2b1abd29a588f1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0734fd7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0734fd7778e95cd721161aa54c7dbb723356f044?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2027   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2027/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2027/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2027?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2027*
