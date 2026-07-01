---
type: pull_request
number: 2003
title: "Update module google.golang.org/grpc to v1.78.0"
state: merged
author: red-hat-konflux
created: 2025-12-29T08:41:43Z
updated: 2026-01-05T08:42:07Z
closed: 2026-01-05T04:53:06Z
merged: 2026-01-05T04:53:06Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2003
---

# Pull Request #2003: Update module google.golang.org/grpc to v1.78.0

**Author**: @red-hat-konflux
**Created**: December 29, 2025 at 08:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.77.0` -> `v1.78.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.78.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.77.0/v1.78.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grpc/grpc-go (google.golang.org/grpc)</summary>

### [`v1.78.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.78.0): Release 1.78.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.77.0...v1.78.0)

### Behavior Changes

- client: Reject target URLs containing unbracketed colons in the hostname in Go version 1.26+. ([#&#8203;8716](https://redirect.github.com/grpc/grpc-go/issues/8716))
  - Special Thanks: [@&#8203;neild](https://redirect.github.com/neild)

### New Features

- stats/otel: Add backend service label to wrr metrics as part of A89. ([#&#8203;8737](https://redirect.github.com/grpc/grpc-go/issues/8737))
- stats/otel: Add subchannel metrics (without the disconnection reason) to eventually replace the pickfirst metrics. ([#&#8203;8738](https://redirect.github.com/grpc/grpc-go/issues/8738))
- client: Wait for all pending goroutines to complete when closing a graceful switch balancer. ([#&#8203;8746](https://redirect.github.com/grpc/grpc-go/issues/8746))
  - Special Thanks: [@&#8203;twz123](https://redirect.github.com/twz123)

### Bug Fixes

- transport/client : Return status code `Unknown` on malformed grpc-status. ([#&#8203;8735](https://redirect.github.com/grpc/grpc-go/issues/8735))
- client: Add `experimental.AcceptCompressors` so callers can restrict the `grpc-accept-encoding` header advertised for a call. ([#&#8203;8718](https://redirect.github.com/grpc/grpc-go/issues/8718))
  - Special Thanks: [@&#8203;iblancasa](https://redirect.github.com/iblancasa)
- xds: Fix a bug in `StringMatcher` where regexes would match incorrectly when ignore\_case is set to true. ([#&#8203;8723](https://redirect.github.com/grpc/grpc-go/issues/8723))
- xds/resolver:
  - Drop previous route resources and report an error when no matching virtual host is found.
  - Only log LDS/RDS configuration errors following a successful update and retain the last valid resource to prevent transient failures. ([#&#8203;8711](https://redirect.github.com/grpc/grpc-go/issues/8711))
- client:
  - Change connectivity state to CONNECTING when creating the name resolver (as part of exiting IDLE).
  - Change connectivity state to TRANSIENT\_FAILURE if name resolver creation fails (as part of exiting IDLE).
  - Change connectivity state to IDLE after idle timeout expires even when current state is TRANSIENT\_FAILURE.
  - Fix a bug that resulted in `OnFinish` call option not being invoked for RPCs where stream creation failed. ([#&#8203;8710](https://redirect.github.com/grpc/grpc-go/issues/8710))
- xdsclient: Fix a race in the xdsClient that could lead to resource-not-found errors. ([#&#8203;8627](https://redirect.github.com/grpc/grpc-go/issues/8627))

### Performance Improvements

- mem: Round up to nearest 4KiB for pool allocations larger than 1MiB. ([#&#8203;8705](https://redirect.github.com/grpc/grpc-go/issues/8705))
  - Special Thanks: [@&#8203;cjc25](https://redirect.github.com/cjc25)

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on December 29, 2025 at 08:41 AM UTC

Commits missing Jira IDs:
86e45729db5e255f6a2fe7adfef4405f807f15ac


### Comment by @codecov-commenter on December 29, 2025 at 08:47 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2003?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`5b0e9f0`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5b0e9f03335f093590483be68cc4ad725da266c3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`86e4572`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/86e45729db5e255f6a2fe7adfef4405f807f15ac?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2003   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2003/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2003/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2003?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2003*
