---
type: pull_request
number: 2062
title: "Update module github.com/klauspost/compress to v1.18.4"
state: merged
author: red-hat-konflux
created: 2026-02-16T08:55:41Z
updated: 2026-02-16T12:52:58Z
closed: 2026-02-16T09:01:48Z
merged: 2026-02-16T09:01:48Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-klauspost-compress-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2062
---

# Pull Request #2062: Update module github.com/klauspost/compress to v1.18.4

**Author**: @red-hat-konflux
**Created**: February 16, 2026 at 08:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-klauspost-compress-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/klauspost/compress](https://redirect.github.com/klauspost/compress) | `v1.18.3` -> `v1.18.4` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fklauspost%2fcompress/v1.18.4?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fklauspost%2fcompress/v1.18.3/v1.18.4?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>klauspost/compress (github.com/klauspost/compress)</summary>

### [`v1.18.4`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.4)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.18.3...v1.18.4)

#### What's Changed

- gzhttp: Add zstandard to server handler wrapper by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1121](https://redirect.github.com/klauspost/compress/pull/1121)
- zstd: Add ResetWithOptions to encoder/decoder by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1122](https://redirect.github.com/klauspost/compress/pull/1122)
- gzhttp: preserve qvalue when extra parameters follow in Accept-Encoding by [@&#8203;analytically](https://redirect.github.com/analytically) in [#&#8203;1116](https://redirect.github.com/klauspost/compress/pull/1116)

#### New Contributors

- [@&#8203;analytically](https://redirect.github.com/analytically) made their first contribution in [#&#8203;1116](https://redirect.github.com/klauspost/compress/pull/1116)
- [@&#8203;ethaizone](https://redirect.github.com/ethaizone) made their first contribution in [#&#8203;1124](https://redirect.github.com/klauspost/compress/pull/1124)
- [@&#8203;zwass](https://redirect.github.com/zwass) made their first contribution in [#&#8203;1125](https://redirect.github.com/klauspost/compress/pull/1125)

**Full Changelog**: <https://github.com/klauspost/compress/compare/v1.18.2...v1.18.4>

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

### Comment by @github-actions on February 16, 2026 at 08:55 AM UTC

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

### Comment by @codecov-commenter on February 16, 2026 at 09:02 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2062?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.38%. Comparing base ([`41b42fc`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/41b42fca5fb05859a77b7af0fd2a429a83615557?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a686d20`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a686d2095b2d563125b747de1573bfc1381923dc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 8 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2062   +/-   ##
=======================================
  Coverage   59.38%   59.38%           
=======================================
  Files         134      134           
  Lines        8679     8679           
=======================================
  Hits         5154     5154           
  Misses       2978     2978           
  Partials      547      547           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2062/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2062/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.38% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2062?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2062*
