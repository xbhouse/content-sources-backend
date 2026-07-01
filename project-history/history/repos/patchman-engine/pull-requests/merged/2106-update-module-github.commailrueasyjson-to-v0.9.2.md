---
type: pull_request
number: 2106
title: "Update module github.com/mailru/easyjson to v0.9.2"
state: merged
author: red-hat-konflux
created: 2026-03-16T09:21:41Z
updated: 2026-03-16T13:25:01Z
closed: 2026-03-16T09:27:14Z
merged: 2026-03-16T09:27:14Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-mailru-easyjson-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2106
---

# Pull Request #2106: Update module github.com/mailru/easyjson to v0.9.2

**Author**: @red-hat-konflux
**Created**: March 16, 2026 at 09:21 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-mailru-easyjson-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/mailru/easyjson](https://redirect.github.com/mailru/easyjson) | `v0.9.1` -> `v0.9.2` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fmailru%2feasyjson/v0.9.2?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fmailru%2feasyjson/v0.9.1/v0.9.2?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>mailru/easyjson (github.com/mailru/easyjson)</summary>

### [`v0.9.2`](https://redirect.github.com/mailru/easyjson/releases/tag/v0.9.2)

[Compare Source](https://redirect.github.com/mailru/easyjson/compare/v0.9.1...v0.9.2)

#### What's Changed

- Bugfix: do not generate invalid JSON when float value is +Inf, -Inf or NaN by [@&#8203;dmitrybarsukov](https://redirect.github.com/dmitrybarsukov) in [#&#8203;421](https://redirect.github.com/mailru/easyjson/pull/421)
- fix null after MarshalText work by [@&#8203;AndreiBerezin](https://redirect.github.com/AndreiBerezin) in [#&#8203;423](https://redirect.github.com/mailru/easyjson/pull/423)
- Support for json:",omitzero" tag by [@&#8203;nsd20463](https://redirect.github.com/nsd20463) in [#&#8203;429](https://redirect.github.com/mailru/easyjson/pull/429)

**Full Changelog**: <https://github.com/mailru/easyjson/compare/v0.9.1...v0.9.2>

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

### Comment by @github-actions on March 16, 2026 at 09:21 AM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** ⚪ **NONE**

No SC Environment-specific impacts detected in this PR.

<details>
<summary>What was checked</summary>

This PR was automatically scanned for:
- Database migrations
- ClowdApp configuration changes
- Kessel integration changes
- AWS service integrations (S3, RDS, ElastiCache)
- Kafka topic changes
- Secrets management changes
- External dependencies
</details>

### Comment by @codecov-commenter on March 16, 2026 at 09:27 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2106?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.40%. Comparing base ([`6ad27f0`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6ad27f01b48e813ccf7fb5aad13ef9fd7ac461d0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b3cca80`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b3cca80f784f0a445f891fb28d831accd1ad37f1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #2106      +/-   ##
==========================================
+ Coverage   59.37%   59.40%   +0.02%     
==========================================
  Files         134      134              
  Lines        8707     8707              
==========================================
+ Hits         5170     5172       +2     
+ Misses       2991     2989       -2     
  Partials      546      546              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2106/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2106/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.40% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2106?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 16, 2026 at 09:24 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2106*
