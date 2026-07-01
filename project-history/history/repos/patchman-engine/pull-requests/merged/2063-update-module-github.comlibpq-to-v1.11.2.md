---
type: pull_request
number: 2063
title: "Update module github.com/lib/pq to v1.11.2"
state: merged
author: red-hat-konflux
created: 2026-02-16T08:55:59Z
updated: 2026-02-16T12:52:59Z
closed: 2026-02-16T09:02:06Z
merged: 2026-02-16T09:02:06Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-lib-pq-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2063
---

# Pull Request #2063: Update module github.com/lib/pq to v1.11.2

**Author**: @red-hat-konflux
**Created**: February 16, 2026 at 08:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-lib-pq-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/lib/pq](https://redirect.github.com/lib/pq) | `v1.11.1` -> `v1.11.2` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2flib%2fpq/v1.11.2?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2flib%2fpq/v1.11.1/v1.11.2?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>lib/pq (github.com/lib/pq)</summary>

### [`v1.11.2`](https://redirect.github.com/lib/pq/blob/HEAD/CHANGELOG.md#v1112-2025-02-10)

[Compare Source](https://redirect.github.com/lib/pq/compare/v1.11.1...v1.11.2)

This fixes two regressions:

- Don't send startup parameters if there is no value, improving compatibility
  with Supavisor ([#&#8203;1260]).

- Don't send `dbname` as a startup parameter if `database=[..]` is used in the
  connection string. It's recommended to use dbname=, as database= is not a
  libpq option, and only worked by accident previously. ([#&#8203;1261])

[#&#8203;1260]: https://redirect.github.com/lib/pq/pull/1260

[#&#8203;1261]: https://redirect.github.com/lib/pq/pull/1261

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

### Comment by @github-actions on February 16, 2026 at 08:56 AM UTC

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

### Comment by @codecov-commenter on February 16, 2026 at 09:03 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2063?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.38%. Comparing base ([`41b42fc`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/41b42fca5fb05859a77b7af0fd2a429a83615557?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c3f2c0e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c3f2c0e0ec6c810d8500b119477c00961dc4b7ee?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 11 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2063   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2063/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2063/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.38% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2063?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2063*
