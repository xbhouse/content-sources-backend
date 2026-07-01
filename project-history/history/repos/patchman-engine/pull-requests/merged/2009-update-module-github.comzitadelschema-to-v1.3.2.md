---
type: pull_request
number: 2009
title: "Update module github.com/zitadel/schema to v1.3.2"
state: merged
author: red-hat-konflux
created: 2026-01-12T04:43:30Z
updated: 2026-01-12T08:39:34Z
closed: 2026-01-12T04:51:32Z
merged: 2026-01-12T04:51:32Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-zitadel-schema-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2009
---

# Pull Request #2009: Update module github.com/zitadel/schema to v1.3.2

**Author**: @red-hat-konflux
**Created**: January 12, 2026 at 04:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-zitadel-schema-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/zitadel/schema](https://redirect.github.com/zitadel/schema) | `v1.3.1` -> `v1.3.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fzitadel%2fschema/v1.3.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fzitadel%2fschema/v1.3.1/v1.3.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>zitadel/schema (github.com/zitadel/schema)</summary>

### [`v1.3.2`](https://redirect.github.com/zitadel/schema/releases/tag/v1.3.2)

[Compare Source](https://redirect.github.com/zitadel/schema/compare/v1.3.1...v1.3.2)

##### Bug Fixes

- add test ([95c7b8f](https://redirect.github.com/zitadel/schema/commit/95c7b8f17212cce5d25cddfad47212c2199c2961))
- decode error message ([1b632a9](https://redirect.github.com/zitadel/schema/commit/1b632a9ad4c13e208228f9298efcc8116f581b4b))
- fix assertion test ([595717f](https://redirect.github.com/zitadel/schema/commit/595717facd9c466dad08f298c70b75dab4fea18d))
- if default element type of value are setted in slice , raise error ([9c25058](https://redirect.github.com/zitadel/schema/commit/9c250587ea34a64ce051d25649ec5ae8f93eec25))
- indirection through nil pointer to embedded struct ([#&#8203;211](https://redirect.github.com/zitadel/schema/issues/211)) ([c3913e4](https://redirect.github.com/zitadel/schema/commit/c3913e416b2fdbe4754081550fcded3e44afef49))
- test data ([2b94d1c](https://redirect.github.com/zitadel/schema/commit/2b94d1c4c5d23a1786b09e4a562aafd79a4bcf5b))

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

### Comment by @jira-linking on January 12, 2026 at 04:43 AM UTC

Commits missing Jira IDs:
275e598b7b6ca28f7710564297d8dff87e8701a1


### Comment by @codecov-commenter on January 12, 2026 at 04:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2009?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`dc69f88`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dc69f88a01d76a6123769bbebc8dceb71212763e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`275e598`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/275e598b7b6ca28f7710564297d8dff87e8701a1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2009   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2009/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2009/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2009?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2009*
