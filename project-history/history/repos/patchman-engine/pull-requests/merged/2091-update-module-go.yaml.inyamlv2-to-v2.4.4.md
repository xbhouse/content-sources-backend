---
type: pull_request
number: 2091
title: "Update module go.yaml.in/yaml/v2 to v2.4.4"
state: merged
author: red-hat-konflux
created: 2026-03-09T05:24:18Z
updated: 2026-03-09T14:13:59Z
closed: 2026-03-09T09:36:20Z
merged: 2026-03-09T09:36:20Z
base_branch: master
head_branch: konflux/mintmaker/master/go.yaml.in-yaml-v2-2.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2091
---

# Pull Request #2091: Update module go.yaml.in/yaml/v2 to v2.4.4

**Author**: @red-hat-konflux
**Created**: March 09, 2026 at 05:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/go.yaml.in-yaml-v2-2.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [go.yaml.in/yaml/v2](https://redirect.github.com/yaml/go-yaml) | `v2.4.3` -> `v2.4.4` | ![age](https://developer.mend.io/api/mc/badges/age/go/go.yaml.in%2fyaml%2fv2/v2.4.4?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/go.yaml.in%2fyaml%2fv2/v2.4.3/v2.4.4?slim=true) |

---

### Release Notes

<details>
<summary>yaml/go-yaml (go.yaml.in/yaml/v2)</summary>

### [`v2.4.4`](https://redirect.github.com/yaml/go-yaml/compare/v2.4.3...v2.4.4)

[Compare Source](https://redirect.github.com/yaml/go-yaml/compare/v2.4.3...v2.4.4)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

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

### Comment by @github-actions on March 09, 2026 at 05:24 AM UTC

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

### Comment by @codecov-commenter on March 09, 2026 at 05:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2091?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.40%. Comparing base ([`03c3b13`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/03c3b13c1ac6dbe72458652777567fa9be377eae?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`deba491`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/deba4910fed7491fa4102b3ba5592dfef4818406?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2091   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2091/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2091/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.40% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2091?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 09, 2026 at 09:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2091*
