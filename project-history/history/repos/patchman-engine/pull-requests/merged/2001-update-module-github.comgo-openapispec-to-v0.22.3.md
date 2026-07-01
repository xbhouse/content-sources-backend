---
type: pull_request
number: 2001
title: "Update module github.com/go-openapi/spec to v0.22.3"
state: merged
author: red-hat-konflux
created: 2025-12-29T04:48:29Z
updated: 2025-12-29T08:42:03Z
closed: 2025-12-29T04:53:51Z
merged: 2025-12-29T04:53:51Z
base_branch: master
head_branch: konflux/mintmaker/master/go-openapi
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2001
---

# Pull Request #2001: Update module github.com/go-openapi/spec to v0.22.3

**Author**: @red-hat-konflux
**Created**: December 29, 2025 at 04:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/go-openapi`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/go-openapi/spec](https://redirect.github.com/go-openapi/spec) | `v0.22.2` -> `v0.22.3` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-openapi%2fspec/v0.22.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-openapi%2fspec/v0.22.2/v0.22.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-openapi/spec (github.com/go-openapi/spec)</summary>

### [`v0.22.3`](https://redirect.github.com/go-openapi/spec/releases/tag/v0.22.3)

[Compare Source](https://redirect.github.com/go-openapi/spec/compare/v0.22.2...v0.22.3)

#### [0.22.3](https://redirect.github.com/go-openapi/spec/tree/v0.22.3) - 2025-12-24

**Full Changelog**: <https://github.com/go-openapi/spec/compare/v0.22.2...v0.22.3>

1 commits in this release.

***

##### <!-- 01 -->Fixed bugs

- fix: fixed key escaping in OrderedItems marshaling by [@&#8203;fredbi](https://redirect.github.com/fredbi) in [#&#8203;246](https://redirect.github.com/go-openapi/spec/pull/246) [...](https://redirect.github.com/go-openapi/spec/commit/3b2ff60674feba6b4c7c62628f6860999b185409)

***

##### People who contributed to this release

- [@&#8203;fredbi](https://redirect.github.com/fredbi)

***

**[spec](https://redirect.github.com/go-openapi/spec) license terms**

[![License][license-badge]][license-url]

[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg

[license-url]: https://redirect.github.com/go-openapi/spec/?tab=Apache-2.0-1-ov-file#readme

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

### Comment by @jira-linking on December 29, 2025 at 04:48 AM UTC

Commits missing Jira IDs:
ee2a345c8c5b18f83cf99e5785b3e72d5f1f779e


### Comment by @codecov-commenter on December 29, 2025 at 04:53 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2001?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`a722b8e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a722b8ecaef32eaf14c2b22d6e2df60d43f85fb2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ee2a345`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ee2a345c8c5b18f83cf99e5785b3e72d5f1f779e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2001   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2001/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2001/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2001?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2001*
