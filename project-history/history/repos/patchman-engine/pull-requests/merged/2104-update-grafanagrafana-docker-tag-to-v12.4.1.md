---
type: pull_request
number: 2104
title: "Update grafana/grafana Docker tag to v12.4.1"
state: merged
author: red-hat-konflux
created: 2026-03-16T05:35:47Z
updated: 2026-03-16T09:21:45Z
closed: 2026-03-16T09:17:58Z
merged: 2026-03-16T09:17:58Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2104
---

# Pull Request #2104: Update grafana/grafana Docker tag to v12.4.1

**Author**: @red-hat-konflux
**Created**: March 16, 2026 at 05:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | patch | `12.4.0` -> `12.4.1` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.4.1`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1241-2026-03-09)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.4.0...v12.4.1)

##### Features and enhancements

- **AccessControl:** Invalidate scope resolver cache on datasource deletion [#&#8203;118741](https://redirect.github.com/grafana/grafana/pull/118741), [@&#8203;mihai-turdean](https://redirect.github.com/mihai-turdean)
- **Go:** Update to 1.25.8 [#&#8203;119693](https://redirect.github.com/grafana/grafana/pull/119693), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Rendering:** Add support for custom CA certs in Image Renderer [#&#8203;118859](https://redirect.github.com/grafana/grafana/pull/118859), [@&#8203;mrevutskyi](https://redirect.github.com/mrevutskyi)

##### Bug fixes

- **AccessControl:** Fix test utility for datasource deletion permissions cleanup (Enterprise)
- **Alerting:** Change scope for testing new receivers to use supported resource type. [#&#8203;118495](https://redirect.github.com/grafana/grafana/pull/118495), [@&#8203;yuri-tceretian](https://redirect.github.com/yuri-tceretian)
- **Alerting:** Fix CollateAlertRuleGroup migration for MariaDB compatibility [#&#8203;119028](https://redirect.github.com/grafana/grafana/pull/119028), [@&#8203;alexander-akhmetov](https://redirect.github.com/alexander-akhmetov)

<!-- 12.4.1 END -->

<!-- 12.3.5 START -->

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

### Comment by @github-actions on March 16, 2026 at 05:36 AM UTC

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

### Comment by @codecov-commenter on March 16, 2026 at 05:41 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2104?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.40%. Comparing base ([`4c3a143`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4c3a143acd46537677d4ab4dccf70bad45d86a8e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`cc99063`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/cc99063132c52364b4bef685b7b5aaf1ab09e445?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2104   +/-   ##
=======================================
  Coverage   59.40%   59.40%           
=======================================
  Files         134      134           
  Lines        8707     8707           
=======================================
  Hits         5172     5172           
  Misses       2989     2989           
  Partials      546      546           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2104/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2104/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.40% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2104?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 16, 2026 at 09:17 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2104*
