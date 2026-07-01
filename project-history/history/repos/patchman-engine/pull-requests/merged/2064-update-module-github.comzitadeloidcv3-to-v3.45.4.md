---
type: pull_request
number: 2064
title: "Update module github.com/zitadel/oidc/v3 to v3.45.4"
state: merged
author: red-hat-konflux
created: 2026-02-16T08:56:12Z
updated: 2026-02-16T12:53:00Z
closed: 2026-02-16T09:02:57Z
merged: 2026-02-16T09:02:57Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2064
---

# Pull Request #2064: Update module github.com/zitadel/oidc/v3 to v3.45.4

**Author**: @red-hat-konflux
**Created**: February 16, 2026 at 08:56 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-zitadel-oidc-v3-3.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/zitadel/oidc/v3](https://redirect.github.com/zitadel/oidc) | `v3.45.3` -> `v3.45.4` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fzitadel%2foidc%2fv3/v3.45.4?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fzitadel%2foidc%2fv3/v3.45.3/v3.45.4?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>zitadel/oidc (github.com/zitadel/oidc/v3)</summary>

### [`v3.45.4`](https://redirect.github.com/zitadel/oidc/releases/tag/v3.45.4)

[Compare Source](https://redirect.github.com/zitadel/oidc/compare/v3.45.3...v3.45.4)

##### Bug Fixes

- grant type in device code validation error message ([#&#8203;845](https://redirect.github.com/zitadel/oidc/issues/845)) ([c238329](https://redirect.github.com/zitadel/oidc/commit/c238329a7906e042fa14242a018d801ab77206d9))

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

### Comment by @codecov-commenter on February 16, 2026 at 09:02 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2064?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.38%. Comparing base ([`41b42fc`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/41b42fca5fb05859a77b7af0fd2a429a83615557?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`bb94b9c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bb94b9c556b87e26f42fd076b3fc72d83ad25c1b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 8 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2064   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2064/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2064/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.38% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2064?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2064*
