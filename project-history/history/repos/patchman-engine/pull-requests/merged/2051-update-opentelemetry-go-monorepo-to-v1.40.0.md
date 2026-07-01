---
type: pull_request
number: 2051
title: "Update opentelemetry-go monorepo to v1.40.0"
state: merged
author: red-hat-konflux
created: 2026-02-09T08:57:20Z
updated: 2026-02-13T00:49:42Z
closed: 2026-02-12T23:14:21Z
merged: 2026-02-12T23:14:21Z
base_branch: master
head_branch: konflux/mintmaker/master/opentelemetry-go-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2051
---

# Pull Request #2051: Update opentelemetry-go monorepo to v1.40.0

**Author**: @red-hat-konflux
**Created**: February 09, 2026 at 08:57 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/opentelemetry-go-monorepo`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [go.opentelemetry.io/otel](https://redirect.github.com/open-telemetry/opentelemetry-go) | `v1.39.0` -> `v1.40.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/go.opentelemetry.io%2fotel/v1.40.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/go.opentelemetry.io%2fotel/v1.39.0/v1.40.0?slim=true) |
| [go.opentelemetry.io/otel/metric](https://redirect.github.com/open-telemetry/opentelemetry-go) | `v1.39.0` -> `v1.40.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/go.opentelemetry.io%2fotel%2fmetric/v1.40.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/go.opentelemetry.io%2fotel%2fmetric/v1.39.0/v1.40.0?slim=true) |
| [go.opentelemetry.io/otel/trace](https://redirect.github.com/open-telemetry/opentelemetry-go) | `v1.39.0` -> `v1.40.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/go.opentelemetry.io%2fotel%2ftrace/v1.40.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/go.opentelemetry.io%2fotel%2ftrace/v1.39.0/v1.40.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>open-telemetry/opentelemetry-go (go.opentelemetry.io/otel)</summary>

### [`v1.40.0`](https://redirect.github.com/open-telemetry/opentelemetry-go/compare/v1.39.0...v1.40.0)

[Compare Source](https://redirect.github.com/open-telemetry/opentelemetry-go/compare/v1.39.0...v1.40.0)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about these updates again.

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

### Comment by @codecov-commenter on February 09, 2026 at 09:03 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2051?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.38%. Comparing base ([`4a40e1f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4a40e1fa1d9ed51010000801926331363a55b73c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`699c06b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/699c06ba060a02350f9c1416197f212690b5784a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2051   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2051/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2051/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.38% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2051?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @github-actions on February 12, 2026 at 11:08 PM UTC

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

---

## Reviews

### Review by @TenSt - Approved on February 12, 2026 at 11:07 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2051*
