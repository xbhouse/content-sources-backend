---
type: pull_request
number: 2025
title: "Update module github.com/segmentio/kafka-go to v0.4.50"
state: merged
author: red-hat-konflux
created: 2026-01-26T08:42:57Z
updated: 2026-01-26T12:38:42Z
closed: 2026-01-26T08:48:21Z
merged: 2026-01-26T08:48:21Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-segmentio-kafka-go-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2025
---

# Pull Request #2025: Update module github.com/segmentio/kafka-go to v0.4.50

**Author**: @red-hat-konflux
**Created**: January 26, 2026 at 08:42 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-segmentio-kafka-go-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/segmentio/kafka-go](https://redirect.github.com/segmentio/kafka-go) | `v0.4.49` -> `v0.4.50` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fsegmentio%2fkafka-go/v0.4.50?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fsegmentio%2fkafka-go/v0.4.49/v0.4.50?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>segmentio/kafka-go (github.com/segmentio/kafka-go)</summary>

### [`v0.4.50`](https://redirect.github.com/segmentio/kafka-go/releases/tag/v0.4.50)

[Compare Source](https://redirect.github.com/segmentio/kafka-go/compare/v0.4.49...v0.4.50)

#### What's Changed

- feat: support describegroups v5 by [@&#8203;petedannemann](https://redirect.github.com/petedannemann) in [#&#8203;1421](https://redirect.github.com/segmentio/kafka-go/pull/1421)

**Full Changelog**: <https://github.com/segmentio/kafka-go/compare/v0.4.49...v0.4.50>

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
6e60efe256e06a2c0159079e90b37e9c41eb6ad6


### Comment by @codecov-commenter on January 26, 2026 at 08:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2025?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`03a8100`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/03a8100da6dfe3aaade5da9eac2b1abd29a588f1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`6e60efe`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6e60efe256e06a2c0159079e90b37e9c41eb6ad6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2025   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2025/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2025/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2025?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2025*
