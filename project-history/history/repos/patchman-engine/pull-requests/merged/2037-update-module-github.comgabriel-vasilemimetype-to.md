---
type: pull_request
number: 2037
title: "Update module github.com/gabriel-vasile/mimetype to v1.4.13"
state: merged
author: red-hat-konflux
created: 2026-02-02T04:47:48Z
updated: 2026-02-02T08:48:52Z
closed: 2026-02-02T04:53:20Z
merged: 2026-02-02T04:53:20Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gabriel-vasile-mimetype-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2037
---

# Pull Request #2037: Update module github.com/gabriel-vasile/mimetype to v1.4.13

**Author**: @red-hat-konflux
**Created**: February 02, 2026 at 04:47 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gabriel-vasile-mimetype-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/gabriel-vasile/mimetype](https://redirect.github.com/gabriel-vasile/mimetype) | `v1.4.12` -> `v1.4.13` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgabriel-vasile%2fmimetype/v1.4.13?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgabriel-vasile%2fmimetype/v1.4.12/v1.4.13?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>gabriel-vasile/mimetype (github.com/gabriel-vasile/mimetype)</summary>

### [`v1.4.13`](https://redirect.github.com/gabriel-vasile/mimetype/releases/tag/v1.4.13): Support for .hlp, .inf, .fm, .bufr

[Compare Source](https://redirect.github.com/gabriel-vasile/mimetype/compare/v1.4.12...v1.4.13)

#### What's Changed

- ndjson: fix inputs truncated on the second line; fix [#&#8203;744](https://redirect.github.com/gabriel-vasile/mimetype/issues/744) in [#&#8203;745](https://redirect.github.com/gabriel-vasile/mimetype/pull/745)
- bmp: harden detection against false-positives in [#&#8203;746](https://redirect.github.com/gabriel-vasile/mimetype/pull/746)
- os2: add support for .hlp and .inf in [#&#8203;747](https://redirect.github.com/gabriel-vasile/mimetype/pull/747)
- ttf: harden detection in [#&#8203;750](https://redirect.github.com/gabriel-vasile/mimetype/pull/750)
- ttf: use ints instead of string for better performance in [#&#8203;751](https://redirect.github.com/gabriel-vasile/mimetype/pull/751)
- framemaker: add support in [#&#8203;752](https://redirect.github.com/gabriel-vasile/mimetype/pull/752)
- bufr: add support in [#&#8203;754](https://redirect.github.com/gabriel-vasile/mimetype/pull/754)
- Extend: ensure MIME string normalization by [@&#8203;yzqzss](https://redirect.github.com/yzqzss) in [#&#8203;756](https://redirect.github.com/gabriel-vasile/mimetype/pull/756)
- m3u: add x-mpegurl alias by [@&#8203;AltayAkkus](https://redirect.github.com/AltayAkkus) in [#&#8203;755](https://redirect.github.com/gabriel-vasile/mimetype/pull/755)

#### New Contributors

- [@&#8203;yzqzss](https://redirect.github.com/yzqzss) made their first contribution in [#&#8203;756](https://redirect.github.com/gabriel-vasile/mimetype/pull/756)
- [@&#8203;AltayAkkus](https://redirect.github.com/AltayAkkus) made their first contribution in [#&#8203;755](https://redirect.github.com/gabriel-vasile/mimetype/pull/755)

**Full Changelog**: <https://github.com/gabriel-vasile/mimetype/compare/v1.4.12...v1.4.13>

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

### Comment by @jira-linking on February 02, 2026 at 04:47 AM UTC

Commits missing Jira IDs:
30a431511cd1dc83267710d267dc6afa0d1335ed


### Comment by @codecov-commenter on February 02, 2026 at 04:53 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2037?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.25%. Comparing base ([`8250aa8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/8250aa872cd37966db751390e4f9591f28ff07ff?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`30a4315`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/30a431511cd1dc83267710d267dc6afa0d1335ed?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2037   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2037/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2037/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2037?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2037*
