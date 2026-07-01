---
type: pull_request
number: 2101
title: "Update golang.org/x/exp digest to 7ab1446"
state: merged
author: red-hat-konflux
created: 2026-03-16T05:34:30Z
updated: 2026-03-16T09:18:16Z
closed: 2026-03-16T09:18:14Z
merged: 2026-03-16T09:18:14Z
base_branch: master
head_branch: konflux/mintmaker/master/golang.org-x-exp-digest
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2101
---

# Pull Request #2101: Update golang.org/x/exp digest to 7ab1446

**Author**: @red-hat-konflux
**Created**: March 16, 2026 at 05:34 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/golang.org-x-exp-digest`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| golang.org/x/exp | require | digest | `3dfff04` -> `7ab1446` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

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

### Comment by @red-hat-konflux on March 16, 2026 at 05:34 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 5 additional dependencies were updated


Details:


| **Package**           | **Change**             |
| :-------------------- | :--------------------- |
| `golang.org/x/crypto` | `v0.48.0` -> `v0.49.0` |
| `golang.org/x/mod`    | `v0.33.0` -> `v0.34.0` |
| `golang.org/x/net`    | `v0.51.0` -> `v0.52.0` |
| `golang.org/x/text`   | `v0.34.0` -> `v0.35.0` |
| `golang.org/x/tools`  | `v0.42.0` -> `v0.43.0` |

### Comment by @github-actions on March 16, 2026 at 05:34 AM UTC

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

### Comment by @codecov-commenter on March 16, 2026 at 05:40 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2101?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.40%. Comparing base ([`4c3a143`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4c3a143acd46537677d4ab4dccf70bad45d86a8e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`825a7a6`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/825a7a61370d019c157395b95954c500768cd75d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2101   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2101/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2101/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.40% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2101?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2101*
