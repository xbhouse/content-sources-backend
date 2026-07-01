---
type: pull_request
number: 1874
title: "Update module github.com/go-playground/validator/v10 to v10.28.0"
state: merged
author: red-hat-konflux
created: 2025-10-06T08:15:48Z
updated: 2026-04-06T14:14:35Z
closed: 2025-10-06T08:29:07Z
merged: 2025-10-06T08:29:07Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1874
---

# Pull Request #1874: Update module github.com/go-playground/validator/v10 to v10.28.0

**Author**: @red-hat-konflux
**Created**: October 06, 2025 at 08:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/go-playground/validator/v10](https://redirect.github.com/go-playground/validator) | `v10.27.0` -> `v10.28.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-playground%2fvalidator%2fv10/v10.28.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-playground%2fvalidator%2fv10/v10.27.0/v10.28.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-playground/validator (github.com/go-playground/validator/v10)</summary>

### [`v10.28.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.28.0): Release 10.28.0

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.27.0...v10.28.0)

#### What's Changed

- Update workflow\.yml to support 2 most recent major versions by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [#&#8203;1417](https://redirect.github.com/go-playground/validator/pull/1417)
- Bump actions/checkout from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1456](https://redirect.github.com/go-playground/validator/pull/1456)
- Go 1.25 support by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [#&#8203;1459](https://redirect.github.com/go-playground/validator/pull/1459)
- Bump github.com/gabriel-vasile/mimetype from 1.4.8 to 1.4.10 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1463](https://redirect.github.com/go-playground/validator/pull/1463)
- Bump golang.org/x/text from 0.22.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1464](https://redirect.github.com/go-playground/validator/pull/1464)
- Bump actions/setup-go from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1465](https://redirect.github.com/go-playground/validator/pull/1465)
- Bump golang.org/x/crypto from 0.33.0 to 0.42.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1467](https://redirect.github.com/go-playground/validator/pull/1467)
- fix: should panic when define duplicate field param in `required_if` by [@&#8203;duyquang6](https://redirect.github.com/duyquang6) in [#&#8203;1468](https://redirect.github.com/go-playground/validator/pull/1468)
- Fixed missing keys from returned errors in map validation by [@&#8203;gelozr](https://redirect.github.com/gelozr) in [#&#8203;1284](https://redirect.github.com/go-playground/validator/pull/1284)
- Added https\_url tag by [@&#8203;ahmedkamalio](https://redirect.github.com/ahmedkamalio) in [#&#8203;1461](https://redirect.github.com/go-playground/validator/pull/1461)
- docs: add description for 'port' validator by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [#&#8203;1435](https://redirect.github.com/go-playground/validator/pull/1435)
- Add alphaspace validator by [@&#8203;takaaa220](https://redirect.github.com/takaaa220) in [#&#8203;1343](https://redirect.github.com/go-playground/validator/pull/1343)

#### New Contributors

- [@&#8203;duyquang6](https://redirect.github.com/duyquang6) made their first contribution in [#&#8203;1468](https://redirect.github.com/go-playground/validator/pull/1468)
- [@&#8203;gelozr](https://redirect.github.com/gelozr) made their first contribution in [#&#8203;1284](https://redirect.github.com/go-playground/validator/pull/1284)
- [@&#8203;ahmedkamalio](https://redirect.github.com/ahmedkamalio) made their first contribution in [#&#8203;1461](https://redirect.github.com/go-playground/validator/pull/1461)
- [@&#8203;takaaa220](https://redirect.github.com/takaaa220) made their first contribution in [#&#8203;1343](https://redirect.github.com/go-playground/validator/pull/1343)

**Full Changelog**: <https://github.com/go-playground/validator/compare/v10.27.0...v10.28.0>

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

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on October 06, 2025 at 08:15 AM UTC

Commits missing Jira IDs:
ff0f3489625b13938242a836f5100dd8f16e3443


### Comment by @codecov-commenter on October 06, 2025 at 08:21 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1874?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.44%. Comparing base ([`43cc7e8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/43cc7e82ecf5394f27affff68d56d601271406f7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ff0f348`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ff0f3489625b13938242a836f5100dd8f16e3443?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 507 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1874   +/-   ##
=======================================
  Coverage   54.44%   54.44%           
=======================================
  Files         138      138           
  Lines       10735    10735           
=======================================
  Hits         5845     5845           
  Misses       4350     4350           
  Partials      540      540           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1874/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1874/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.44% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1874?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1874*
