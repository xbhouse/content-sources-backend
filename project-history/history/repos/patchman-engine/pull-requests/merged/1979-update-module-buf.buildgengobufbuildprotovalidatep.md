---
type: pull_request
number: 1979
title: "Update module buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go to v1.36.11-20251209175733-2a1774d88802.1"
state: merged
author: red-hat-konflux
created: 2025-12-15T04:40:05Z
updated: 2025-12-15T12:38:34Z
closed: 2025-12-15T08:48:44Z
merged: 2025-12-15T08:48:44Z
base_branch: master
head_branch: konflux/mintmaker/master/buf.build-gen-go-bufbuild-protovalidate-protocolbuffers-go-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1979
---

# Pull Request #1979: Update module buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go to v1.36.11-20251209175733-2a1774d88802.1

**Author**: @red-hat-konflux
**Created**: December 15, 2025 at 04:40 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/buf.build-gen-go-bufbuild-protovalidate-protocolbuffers-go-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go | `v1.36.10-20250912141014-52f32327d4b0.1` -> `v1.36.11-20251209175733-2a1774d88802.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/buf.build%2fgen%2fgo%2fbufbuild%2fprotovalidate%2fprotocolbuffers%2fgo/v1.36.11-20251209175733-2a1774d88802.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/buf.build%2fgen%2fgo%2fbufbuild%2fprotovalidate%2fprotocolbuffers%2fgo/v1.36.10-20250912141014-52f32327d4b0.1/v1.36.11-20251209175733-2a1774d88802.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on December 15, 2025 at 04:40 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                  | **Change**               |
| :--------------------------- | :----------------------- |
| `google.golang.org/protobuf` | `v1.36.10` -> `v1.36.11` |

### Comment by @jira-linking on December 15, 2025 at 04:40 AM UTC

Commits missing Jira IDs:
0071a74dbfbcc1b4a40623738061ffb5feff3219


### Comment by @codecov-commenter on December 15, 2025 at 04:45 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1979?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`4539026`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4539026948e76ff035c11d5b933d5c48cfaa5106?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0071a74`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0071a74dbfbcc1b4a40623738061ffb5feff3219?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1979   +/-   ##
=======================================
  Coverage   59.01%   59.01%           
=======================================
  Files         131      131           
  Lines        8493     8493           
=======================================
  Hits         5012     5012           
  Misses       2947     2947           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1979/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1979/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1979?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1979*
