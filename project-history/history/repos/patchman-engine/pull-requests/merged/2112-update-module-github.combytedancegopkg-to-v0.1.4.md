---
type: pull_request
number: 2112
title: "Update module github.com/bytedance/gopkg to v0.1.4"
state: merged
author: red-hat-konflux
created: 2026-03-23T05:15:29Z
updated: 2026-03-23T09:15:28Z
closed: 2026-03-23T05:15:45Z
merged: 2026-03-23T05:15:45Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-bytedance-gopkg-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2112
---

# Pull Request #2112: Update module github.com/bytedance/gopkg to v0.1.4

**Author**: @red-hat-konflux
**Created**: March 23, 2026 at 05:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-bytedance-gopkg-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/bytedance/gopkg](https://redirect.github.com/bytedance/gopkg) | `v0.1.3` -> `v0.1.4` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fbytedance%2fgopkg/v0.1.4?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fbytedance%2fgopkg/v0.1.3/v0.1.4?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>bytedance/gopkg (github.com/bytedance/gopkg)</summary>

### [`v0.1.4`](https://redirect.github.com/bytedance/gopkg/releases/tag/v0.1.4)

[Compare Source](https://redirect.github.com/bytedance/gopkg/compare/v0.1.3...v0.1.4)

#### What's Changed

- fix(lang/dirtmake): compat with tinygo by [@&#8203;xiaost](https://redirect.github.com/xiaost) in [#&#8203;253](https://redirect.github.com/bytedance/gopkg/pull/253)
- feat(circuitbreaker): cleanup breakers with no samples to prevent memory leak by [@&#8203;PureWhiteWu](https://redirect.github.com/PureWhiteWu) in [#&#8203;254](https://redirect.github.com/bytedance/gopkg/pull/254)
- fix: typos/grammar, and zset perf by [@&#8203;xiaost](https://redirect.github.com/xiaost) in [#&#8203;255](https://redirect.github.com/bytedance/gopkg/pull/255)
- chore: update code owners by [@&#8203;JasonChen233](https://redirect.github.com/JasonChen233) in [#&#8203;257](https://redirect.github.com/bytedance/gopkg/pull/257)
- chore(README): updated to reflect current status by [@&#8203;xiaost](https://redirect.github.com/xiaost) in [#&#8203;258](https://redirect.github.com/bytedance/gopkg/pull/258)
- chore: update codeowners by [@&#8203;xiaost](https://redirect.github.com/xiaost) in [#&#8203;261](https://redirect.github.com/bytedance/gopkg/pull/261)
- refactor: clean up dependencies by [@&#8203;xiaost](https://redirect.github.com/xiaost) in [#&#8203;262](https://redirect.github.com/bytedance/gopkg/pull/262)

#### New Contributors

- [@&#8203;JasonChen233](https://redirect.github.com/JasonChen233) made their first contribution in [#&#8203;257](https://redirect.github.com/bytedance/gopkg/pull/257)

**Full Changelog**: <https://github.com/bytedance/gopkg/compare/v0.1.3...v0.1.4>

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

### Comment by @github-actions on March 23, 2026 at 05:15 AM UTC

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

### Comment by @github-actions on March 23, 2026 at 05:15 AM UTC

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

### Comment by @codecov-commenter on March 23, 2026 at 05:21 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2112?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`ac3fb7f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ac3fb7ff6a173da2eb250d2166031628137f2d66?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`e646f0e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e646f0eba7d50319534bcda34e4ec04321e1b922?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #2112      +/-   ##
==========================================
+ Coverage   59.16%   59.18%   +0.02%     
==========================================
  Files         134      134              
  Lines        8622     8622              
==========================================
+ Hits         5101     5103       +2     
+ Misses       2986     2984       -2     
  Partials      535      535              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2112/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2112/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2112?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2112*
