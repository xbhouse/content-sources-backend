---
type: pull_request
number: 1945
title: "Update module github.com/prometheus/common to v0.67.4"
state: merged
author: red-hat-konflux
created: 2025-11-24T04:44:20Z
updated: 2025-12-01T08:46:20Z
closed: 2025-12-01T08:45:55Z
merged: 2025-12-01T08:45:55Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-common-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1945
---

# Pull Request #1945: Update module github.com/prometheus/common to v0.67.4

**Author**: @red-hat-konflux
**Created**: November 24, 2025 at 04:44 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | `v0.67.2` -> `v0.67.4` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fcommon/v0.67.4?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fcommon/v0.67.2/v0.67.4?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.67.4`](https://redirect.github.com/prometheus/common/releases/tag/v0.67.4): / 2025-11-18

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.3...v0.67.4)

#### What's Changed

- chore: clean up golangci-lint configuration by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;782](https://redirect.github.com/prometheus/common/pull/782)
- chore: 'omitempty' to Oauth2 fields with type Secret to avoid requiring them by [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) in [#&#8203;864](https://redirect.github.com/prometheus/common/pull/864)
- chore: Add omitempty tag to all config fields by [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) in [#&#8203;865](https://redirect.github.com/prometheus/common/pull/865)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.67.3...v0.67.4>

### [`v0.67.3`](https://redirect.github.com/prometheus/common/releases/tag/v0.67.3): / 2025-11-18

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.67.2...v0.67.3)

#### What's Changed

- Support JWT Profile for Authorization Grant (RFC 7523 3.1) by [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) in [#&#8203;862](https://redirect.github.com/prometheus/common/pull/862)
- Config: remove outdated comment about HTTP/2 issues by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;863](https://redirect.github.com/prometheus/common/pull/863)

#### New Contributors

- [@&#8203;JorTurFer](https://redirect.github.com/JorTurFer) made their first contribution in [#&#8203;862](https://redirect.github.com/prometheus/common/pull/862)

**Full Changelog**: <https://github.com/prometheus/common/compare/v0.67.2...v0.67.3>

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

### Comment by @codecov-commenter on December 01, 2025 at 04:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1945?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.83%. Comparing base ([`bd9ea7c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bd9ea7c3eec3c02fe4efcd77a10e3da511890970?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f3d9029`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f3d9029fecea51955c3ee663b47e01a6ec2cf0e2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1945   +/-   ##
=======================================
  Coverage   58.83%   58.83%           
=======================================
  Files         131      131           
  Lines        8407     8407           
=======================================
  Hits         4946     4946           
  Misses       2927     2927           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1945/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1945/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.83% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1945?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1945*
