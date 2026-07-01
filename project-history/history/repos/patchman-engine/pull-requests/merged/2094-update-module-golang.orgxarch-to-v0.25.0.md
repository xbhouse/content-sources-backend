---
type: pull_request
number: 2094
title: "Update module golang.org/x/arch to v0.25.0"
state: merged
author: red-hat-konflux
created: 2026-03-09T09:30:17Z
updated: 2026-03-09T14:13:59Z
closed: 2026-03-09T11:41:55Z
merged: 2026-03-09T11:41:55Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-arch-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2094
---

# Pull Request #2094: Update module golang.org/x/arch to v0.25.0

**Author**: @red-hat-konflux
**Created**: March 09, 2026 at 09:30 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-arch-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| golang.org/x/arch | `v0.24.0` -> `v0.25.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/golang.org%2fx%2farch/v0.25.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/golang.org%2fx%2farch/v0.24.0/v0.25.0?slim=true) |

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

### Comment by @github-actions on March 09, 2026 at 09:30 AM UTC

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

### Comment by @codecov-commenter on March 09, 2026 at 09:37 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2094?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.40%. Comparing base ([`69a4c6b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/69a4c6bada80314a01e0a1e763b05917ff1e34f5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`3fa563c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/3fa563c84b4f35131c1057643048ee601fc27401?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2094   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2094/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2094/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.40% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2094?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @TenSt - Approved on March 09, 2026 at 10:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2094*
