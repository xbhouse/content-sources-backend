---
type: pull_request
number: 2026
title: "Update module github.com/sirupsen/logrus to v1.9.4"
state: merged
author: red-hat-konflux
created: 2026-01-26T08:43:09Z
updated: 2026-02-02T09:45:10Z
closed: 2026-02-02T09:45:09Z
merged: 2026-02-02T09:45:09Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-sirupsen-logrus-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2026
---

# Pull Request #2026: Update module github.com/sirupsen/logrus to v1.9.4

**Author**: @red-hat-konflux
**Created**: January 26, 2026 at 08:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-sirupsen-logrus-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/sirupsen/logrus](https://redirect.github.com/sirupsen/logrus) | `v1.9.3` -> `v1.9.4` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fsirupsen%2flogrus/v1.9.4?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fsirupsen%2flogrus/v1.9.3/v1.9.4?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>sirupsen/logrus (github.com/sirupsen/logrus)</summary>

### [`v1.9.4`](https://redirect.github.com/sirupsen/logrus/releases/tag/v1.9.4)

[Compare Source](https://redirect.github.com/sirupsen/logrus/compare/v1.9.3...v1.9.4)

#### Notable changes

- go.mod: update minimum supported go version to v1.17 [#&#8203;1460](https://redirect.github.com/sirupsen/logrus/pull/1460)
- go.mod: bump up dependencies  [#&#8203;1460](https://redirect.github.com/sirupsen/logrus/pull/1460)
- Touch-up godoc and add "doc" links.
- README: fix links, grammar, and update examples.
- Add GNU/Hurd support [#&#8203;1364](https://redirect.github.com/sirupsen/logrus/pull/1364)
- Add WASI wasip1 support [#&#8203;1388](https://redirect.github.com/sirupsen/logrus/pull/1388)
- Remove uses of deprecated `ioutil` package [#&#8203;1472](https://redirect.github.com/sirupsen/logrus/pull/1472)
- CI: update actions and golangci-lint [#&#8203;1459](https://redirect.github.com/sirupsen/logrus/pull/1459)
- CI: remove appveyor, add macOS  [#&#8203;1460](https://redirect.github.com/sirupsen/logrus/pull/1460)

**Full Changelog**: <https://github.com/sirupsen/logrus/compare/v1.9.3...v1.9.4>

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
b4a9cdcbe041db7b420a45e24b5c51c7ef47b480


### Comment by @codecov-commenter on January 26, 2026 at 08:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2026?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.25%. Comparing base ([`92592c9`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/92592c95755f9762d74e9f8cc68fac318de73333?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b4a9cdc`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b4a9cdcbe041db7b420a45e24b5c51c7ef47b480?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2026   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2026/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2026/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2026?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2026*
