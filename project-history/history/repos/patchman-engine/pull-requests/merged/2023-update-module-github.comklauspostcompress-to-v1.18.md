---
type: pull_request
number: 2023
title: "Update module github.com/klauspost/compress to v1.18.3"
state: merged
author: red-hat-konflux
created: 2026-01-26T04:40:32Z
updated: 2026-01-26T08:43:14Z
closed: 2026-01-26T04:46:18Z
merged: 2026-01-26T04:46:18Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-klauspost-compress-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2023
---

# Pull Request #2023: Update module github.com/klauspost/compress to v1.18.3

**Author**: @red-hat-konflux
**Created**: January 26, 2026 at 04:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-klauspost-compress-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/klauspost/compress](https://redirect.github.com/klauspost/compress) | `v1.18.2` -> `v1.18.3` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fklauspost%2fcompress/v1.18.3?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fklauspost%2fcompress/v1.18.2/v1.18.3?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>klauspost/compress (github.com/klauspost/compress)</summary>

### [`v1.18.3`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.3)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.18.2...v1.18.3)

Downstream CVE-2025-61728

See [golang/go#77102](https://redirect.github.com/golang/go/issues/77102)

**Full Changelog**: <https://github.com/klauspost/compress/compare/v1.18.2...v1.18.3>

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

### Comment by @jira-linking on January 26, 2026 at 04:40 AM UTC

Commits missing Jira IDs:
ef4e8e3cdb6d28fe39d2378ae639140de0dda7bc


### Comment by @codecov-commenter on January 26, 2026 at 04:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2023?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`8c737f8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8c737f8e3aabc913f375d0a771c5a44d749ee447?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ef4e8e3`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ef4e8e3cdb6d28fe39d2378ae639140de0dda7bc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2023   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2023/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2023/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2023?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2023*
