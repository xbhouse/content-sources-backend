---
type: pull_request
number: 1967
title: "Update module github.com/klauspost/compress to v1.18.2"
state: merged
author: red-hat-konflux
created: 2025-12-08T04:37:07Z
updated: 2025-12-08T08:41:35Z
closed: 2025-12-08T04:48:08Z
merged: 2025-12-08T04:48:08Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-klauspost-compress-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1967
---

# Pull Request #1967: Update module github.com/klauspost/compress to v1.18.2

**Author**: @red-hat-konflux
**Created**: December 08, 2025 at 04:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-klauspost-compress-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/klauspost/compress](https://redirect.github.com/klauspost/compress) | `v1.18.1` -> `v1.18.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fklauspost%2fcompress/v1.18.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fklauspost%2fcompress/v1.18.1/v1.18.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>klauspost/compress (github.com/klauspost/compress)</summary>

### [`v1.18.2`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.2)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.18.1...v1.18.2)

#### What's Changed

- flate: Fix invalid encoding on level 9 with single value input by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1115](https://redirect.github.com/klauspost/compress/pull/1115)
- flate: reduce stateless allocations by [@&#8203;RXamzin](https://redirect.github.com/RXamzin) in [#&#8203;1106](https://redirect.github.com/klauspost/compress/pull/1106)
- build(deps): bump github/codeql-action from 3.30.5 to 4.31.2 in the github-actions group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1111](https://redirect.github.com/klauspost/compress/pull/1111)

`v1.18.1` is marked "retracted" due to invalid flate/zip/gzip encoding.

#### New Contributors

- [@&#8203;RXamzin](https://redirect.github.com/RXamzin) made their first contribution in [#&#8203;1106](https://redirect.github.com/klauspost/compress/pull/1106)

**Full Changelog**: <https://github.com/klauspost/compress/compare/v1.18.1...v1.18.2>

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

### Comment by @jira-linking on December 08, 2025 at 04:37 AM UTC

Commits missing Jira IDs:
38f8cafb460b0245458562fd86b6c69ed461b886


### Comment by @codecov-commenter on December 08, 2025 at 04:42 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1967?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.84%. Comparing base ([`c25f971`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c25f9719a58b2dc920748401ffd96bb4c4ed2047?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`38f8caf`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/38f8cafb460b0245458562fd86b6c69ed461b886?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1967      +/-   ##
==========================================
+ Coverage   58.81%   58.84%   +0.02%     
==========================================
  Files         131      131              
  Lines        8436     8436              
==========================================
+ Hits         4962     4964       +2     
+ Misses       2939     2937       -2     
  Partials      535      535              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1967/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1967/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.84% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1967?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1967*
