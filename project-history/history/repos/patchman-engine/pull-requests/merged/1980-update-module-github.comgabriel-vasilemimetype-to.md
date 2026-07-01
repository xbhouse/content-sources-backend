---
type: pull_request
number: 1980
title: "Update module github.com/gabriel-vasile/mimetype to v1.4.12"
state: merged
author: red-hat-konflux
created: 2025-12-15T08:42:04Z
updated: 2025-12-15T12:38:35Z
closed: 2025-12-15T08:56:57Z
merged: 2025-12-15T08:56:57Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gabriel-vasile-mimetype-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1980
---

# Pull Request #1980: Update module github.com/gabriel-vasile/mimetype to v1.4.12

**Author**: @red-hat-konflux
**Created**: December 15, 2025 at 08:42 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gabriel-vasile-mimetype-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/gabriel-vasile/mimetype](https://redirect.github.com/gabriel-vasile/mimetype) | `v1.4.11` -> `v1.4.12` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgabriel-vasile%2fmimetype/v1.4.12?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgabriel-vasile%2fmimetype/v1.4.11/v1.4.12?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>gabriel-vasile/mimetype (github.com/gabriel-vasile/mimetype)</summary>

### [`v1.4.12`](https://redirect.github.com/gabriel-vasile/mimetype/releases/tag/v1.4.12): RFC822, GRIB, Zlib support

[Compare Source](https://redirect.github.com/gabriel-vasile/mimetype/compare/v1.4.11...v1.4.12)

#### What's Changed

- zip+json: add benchmarks for better performance tracking of pathological inputs in [#&#8203;730](https://redirect.github.com/gabriel-vasile/mimetype/pull/730)
- zip+json: performance improvements for pathological cases in [#&#8203;732](https://redirect.github.com/gabriel-vasile/mimetype/pull/732)
- Fix integer overflow panic on 32bit architectures in [#&#8203;733](https://redirect.github.com/gabriel-vasile/mimetype/pull/733)
- ci: add more linters and fix their warnings in [#&#8203;734](https://redirect.github.com/gabriel-vasile/mimetype/pull/734)
- jar: manifest must be first in [#&#8203;735](https://redirect.github.com/gabriel-vasile/mimetype/pull/735)
- rfc822: add support in [#&#8203;740](https://redirect.github.com/gabriel-vasile/mimetype/pull/740)
- grib: add support in [#&#8203;742](https://redirect.github.com/gabriel-vasile/mimetype/pull/742)
- zlib: add support in [#&#8203;743](https://redirect.github.com/gabriel-vasile/mimetype/pull/743)

**Full Changelog**: <https://github.com/gabriel-vasile/mimetype/compare/v1.4.11...v1.4.12>

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

### Comment by @jira-linking on December 15, 2025 at 08:42 AM UTC

Commits missing Jira IDs:
cff66925db7a629c9f070ae34323576a20a755f6


### Comment by @codecov-commenter on December 15, 2025 at 08:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1980?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`4539026`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4539026948e76ff035c11d5b933d5c48cfaa5106?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`cff6692`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/cff66925db7a629c9f070ae34323576a20a755f6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1980   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1980/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1980/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1980?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1980*
