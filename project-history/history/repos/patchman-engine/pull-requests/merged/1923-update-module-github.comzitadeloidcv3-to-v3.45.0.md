---
type: pull_request
number: 1923
title: "Update module github.com/zitadel/oidc/v3 to v3.45.0"
state: merged
author: red-hat-konflux
created: 2025-11-10T08:20:12Z
updated: 2025-11-10T12:17:49Z
closed: 2025-11-10T10:17:48Z
merged: 2025-11-10T10:17:48Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1923
---

# Pull Request #1923: Update module github.com/zitadel/oidc/v3 to v3.45.0

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 08:20 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/zitadel/oidc/v3](https://redirect.github.com/zitadel/oidc) | `v3.44.0` -> `v3.45.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fzitadel%2foidc%2fv3/v3.45.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fzitadel%2foidc%2fv3/v3.44.0/v3.45.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>zitadel/oidc (github.com/zitadel/oidc/v3)</summary>

### [`v3.45.0`](https://redirect.github.com/zitadel/oidc/releases/tag/v3.45.0)

[Compare Source](https://redirect.github.com/zitadel/oidc/compare/v3.44.0...v3.45.0)

##### Features

- **rp:** add WithPKCEFromDisocvery ([#&#8203;776](https://redirect.github.com/zitadel/oidc/issues/776)) ([e3169b6](https://redirect.github.com/zitadel/oidc/commit/e3169b695fdd837fc98a824c918248095e85dedf)), closes [#&#8203;506](https://redirect.github.com/zitadel/oidc/issues/506) [#&#8203;506](https://redirect.github.com/zitadel/oidc/issues/506)

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

### Comment by @codecov-commenter on November 10, 2025 at 08:25 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1923?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.90%. Comparing base ([`f287d51`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f287d51f7f21f3378c4f5a6a10444c8350708b69?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`dfe7b66`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dfe7b66fa1d6cb8aea37bb2d0307ead56f1f1e47?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1923   +/-   ##
=======================================
  Coverage   58.90%   58.90%           
=======================================
  Files         131      131           
  Lines        8421     8421           
=======================================
  Hits         4960     4960           
  Misses       2927     2927           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1923/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1923/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.90% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1923?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1923*
