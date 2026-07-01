---
type: pull_request
number: 2134
title: "Update grafana/grafana Docker tag to v12.4.2"
state: merged
author: red-hat-konflux
created: 2026-03-30T01:24:58Z
updated: 2026-03-30T05:13:55Z
closed: 2026-03-30T01:25:09Z
merged: 2026-03-30T01:25:09Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2134
---

# Pull Request #2134: Update grafana/grafana Docker tag to v12.4.2

**Author**: @red-hat-konflux
**Created**: March 30, 2026 at 01:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | patch | `12.4.1` -> `12.4.2` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v12.4.2`](https://redirect.github.com/grafana/grafana/blob/HEAD/CHANGELOG.md#1242-2026-03-25)

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v12.4.1...v12.4.2)

##### Features and enhancements

- **Analytics tab:** Improve voice over accessibility (Enterprise)
- **Dashboards a11y:** Do not open time zonemenu on focus [#&#8203;120388](https://redirect.github.com/grafana/grafana/pull/120388), [@&#8203;idastambuk](https://redirect.github.com/idastambuk)
- **Dashboards:** Resolve display names by identity in version history [#&#8203;120273](https://redirect.github.com/grafana/grafana/pull/120273), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Plugins:** Forward AWS SDK credential chain env vars to external AWS plugins [#&#8203;120209](https://redirect.github.com/grafana/grafana/pull/120209), [@&#8203;kevinwcyu](https://redirect.github.com/kevinwcyu)
- **Public Dashboards:** Prevent unintended CRUD operations from different orgs [#&#8203;120457](https://redirect.github.com/grafana/grafana/pull/120457), [@&#8203;mmandrus](https://redirect.github.com/mmandrus)

##### Bug fixes

- **IAM:** Handle NULL team\_member.external column to fix dashboard loading [#&#8203;120179](https://redirect.github.com/grafana/grafana/pull/120179), [@&#8203;difro](https://redirect.github.com/difro)
- **Plugins:** Fix installer IsDisabled condition [#&#8203;120568](https://redirect.github.com/grafana/grafana/pull/120568), [@&#8203;andresmgot](https://redirect.github.com/andresmgot)
- **Plugins:** Forward PLUGIN\_UNIX\_SOCKET\_DIR to plugin processes to fix tmp dir in restricted environments [#&#8203;120275](https://redirect.github.com/grafana/grafana/pull/120275), [@&#8203;HarshadaGawas05](https://redirect.github.com/HarshadaGawas05)
- **Security:** Fix CVE-2026-27876
- **Security:** Fix CVE-2026-27877
- **Security:** Fix CVE-2026-28375
- **Security:** Fix CVE-2026-27879
- **Security:** Fix CVE-2026-27880
- **Security:** Fix CVE-2026-27876
- **Security:** Fix CVE-2026-27876
- **Security:** Fix CVE-2026-33375

<!-- 12.4.2 END -->

<!-- 12.3.6 START -->

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

### Comment by @github-actions on March 30, 2026 at 01:25 AM UTC

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

### Comment by @codecov-commenter on March 30, 2026 at 01:30 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2134?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`c4888aa`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c4888aa8434a265eb6e68e683ee57a42a5193cd9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0a64c9c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0a64c9c23137aea629d7aadede212d1324d91169?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2134   +/-   ##
=======================================
  Coverage   59.18%   59.18%           
=======================================
  Files         134      134           
  Lines        8624     8624           
=======================================
  Hits         5104     5104           
  Misses       2985     2985           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2134/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2134/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2134?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2134*
