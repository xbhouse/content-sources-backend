---
type: pull_request
number: 1958
title: "Update module github.com/xdg-go/scram to v1.2.0"
state: merged
author: red-hat-konflux
created: 2025-12-01T08:40:56Z
updated: 2025-12-01T12:44:38Z
closed: 2025-12-01T08:46:27Z
merged: 2025-12-01T08:46:27Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-xdg-go-scram-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1958
---

# Pull Request #1958: Update module github.com/xdg-go/scram to v1.2.0

**Author**: @red-hat-konflux
**Created**: December 01, 2025 at 08:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-xdg-go-scram-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/xdg-go/scram](https://redirect.github.com/xdg-go/scram) | `v1.1.2` -> `v1.2.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fxdg-go%2fscram/v1.2.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fxdg-go%2fscram/v1.1.2/v1.2.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>xdg-go/scram (github.com/xdg-go/scram)</summary>

### [`v1.2.0`](https://redirect.github.com/xdg-go/scram/blob/HEAD/CHANGELOG.md#v120---2025-11-24)

[Compare Source](https://redirect.github.com/xdg-go/scram/compare/v1.1.2...v1.2.0)

##### Added

- **Channel binding support for SCRAM-PLUS variants** (RFC 5929, RFC 9266)
- `GetStoredCredentialsWithError()` method that returns errors from PBKDF2
  key derivation instead of panicking.
- Support for Go 1.24+ stdlib `crypto/pbkdf2` package, which provides
  FIPS 140-3 compliance when using SHA-256 or SHA-512 hash functions.

##### Changed

- Minimum Go version bumped from 1.11 to 1.18.
- Migrated from `github.com/xdg-go/pbkdf2` to stdlib `crypto/pbkdf2` on
  Go 1.24+. Legacy Go versions (<1.24) continue using the external
  library via build tags for backward compatibility.
- Internal error handling improved for PBKDF2 key derivation failures.

##### Deprecated

- `GetStoredCredentials()` is deprecated in favor of
  `GetStoredCredentialsWithError()`. The old method panics on PBKDF2
  errors to maintain backward compatibility but will be removed in a
  future major version.

##### Notes

- FIPS 140-3 compliance is available on Go 1.24+ when using SCRAM-SHA-256
  or SCRAM-SHA-512 with appropriate salt lengths (≥16 bytes). SCRAM-SHA-1
  is not FIPS-approved.

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

### Comment by @codecov-commenter on December 01, 2025 at 08:46 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1958?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.83%. Comparing base ([`bd9ea7c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bd9ea7c3eec3c02fe4efcd77a10e3da511890970?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d7a925d`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d7a925d42f567e065393c5b4e1dc87555e842700?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 6 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1958   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1958/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1958/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.83% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1958?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1958*
