---
type: pull_request
number: 1959
title: "Update module google.golang.org/grpc to v1.77.0"
state: merged
author: red-hat-konflux
created: 2025-12-01T08:41:44Z
updated: 2025-12-06T04:39:58Z
closed: 2025-12-06T00:48:00Z
merged: 2025-12-06T00:48:00Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1959
---

# Pull Request #1959: Update module google.golang.org/grpc to v1.77.0

**Author**: @red-hat-konflux
**Created**: December 01, 2025 at 08:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.75.1` -> `v1.77.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.77.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.75.1/v1.77.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grpc/grpc-go (google.golang.org/grpc)</summary>

### [`v1.77.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.77.0): Release 1.77.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.76.0...v1.77.0)

### API Changes

- mem: Replace the `Reader` interface with a struct for better performance and maintainability. ([#&#8203;8669](https://redirect.github.com/grpc/grpc-go/issues/8669))

### Behavior Changes

- balancer/pickfirst: Remove support for the old `pick_first` LB policy via the environment variable `GRPC_EXPERIMENTAL_ENABLE_NEW_PICK_FIRST=false`. The new `pick_first` has been the default since `v1.71.0`. ([#&#8203;8672](https://redirect.github.com/grpc/grpc-go/issues/8672))

### Bug Fixes

- xdsclient: Fix a race condition in the ADS stream implementation that could result in `resource-not-found` errors, causing the gRPC client channel to move to `TransientFailure`. ([#&#8203;8605](https://redirect.github.com/grpc/grpc-go/issues/8605))
- client: Ignore HTTP status header for gRPC streams. ([#&#8203;8548](https://redirect.github.com/grpc/grpc-go/issues/8548))
- client: Set a read deadline when closing a transport to prevent it from blocking indefinitely on a broken connection. ([#&#8203;8534](https://redirect.github.com/grpc/grpc-go/issues/8534))
  - Special Thanks: [@&#8203;jgold2-stripe](https://redirect.github.com/jgold2-stripe)
- client: Fix a bug where default port 443 was not automatically added to addresses without a specified port when sent to a proxy.
  - Setting environment variable `GRPC_EXPERIMENTAL_ENABLE_DEFAULT_PORT_FOR_PROXY_TARGET=false` disables this change; please file a bug if any problems are encountered as we will remove this option soon. ([#&#8203;8613](https://redirect.github.com/grpc/grpc-go/issues/8613))
- balancer/pickfirst: Fix a bug where duplicate addresses were not being ignored as intended. ([#&#8203;8611](https://redirect.github.com/grpc/grpc-go/issues/8611))
- server: Fix a bug that caused overcounting of channelz metrics for successful and failed streams. ([#&#8203;8573](https://redirect.github.com/grpc/grpc-go/issues/8573))
  - Special Thanks: [@&#8203;hugehoo](https://redirect.github.com/hugehoo)
- balancer/pickfirst: When configured, shuffle addresses in resolver updates that lack endpoints. Since gRPC automatically adds endpoints to resolver updates, this bug only affects custom LB policies that delegate to `pick_first` but don't set endpoints. ([#&#8203;8610](https://redirect.github.com/grpc/grpc-go/issues/8610))
- mem: Clear large buffers before re-using. ([#&#8203;8670](https://redirect.github.com/grpc/grpc-go/issues/8670))

### Performance Improvements

- transport: Reduce heap allocations to reduce time spent in garbage collection. ([#&#8203;8624](https://redirect.github.com/grpc/grpc-go/issues/8624), [#&#8203;8630](https://redirect.github.com/grpc/grpc-go/issues/8630), [#&#8203;8639](https://redirect.github.com/grpc/grpc-go/issues/8639), [#&#8203;8668](https://redirect.github.com/grpc/grpc-go/issues/8668))
- transport: Avoid copies when reading and writing Data frames. ([#&#8203;8657](https://redirect.github.com/grpc/grpc-go/issues/8657), [#&#8203;8667](https://redirect.github.com/grpc/grpc-go/issues/8667))
- mem: Avoid clearing newly allocated buffers. ([#&#8203;8670](https://redirect.github.com/grpc/grpc-go/issues/8670))

### New Features

- outlierdetection: Add metrics specified in [gRFC A91](https://redirect.github.com/grpc/proposal/blob/master/A91-outlier-detection-metrics.md). ([#&#8203;8644](https://redirect.github.com/grpc/grpc-go/issues/8644))
  - Special Thanks: [@&#8203;davinci26](https://redirect.github.com/davinci26), [@&#8203;PardhuKonakanchi](https://redirect.github.com/PardhuKonakanchi)
- stats/opentelemetry: Add support for optional label `grpc.lb.backend_service` in per-call metrics ([#&#8203;8637](https://redirect.github.com/grpc/grpc-go/issues/8637))
- xds: Add support for JWT Call Credentials as specified in [gRFC A97](https://redirect.github.com/grpc/proposal/blob/master/A97-xds-jwt-call-creds.md). Set environment variable `GRPC_EXPERIMENTAL_XDS_BOOTSTRAP_CALL_CREDS=true` to enable this feature. ([#&#8203;8536](https://redirect.github.com/grpc/grpc-go/issues/8536))
  - Special Thanks: [@&#8203;dimpavloff](https://redirect.github.com/dimpavloff)
- experimental/stats: Add support for up/down counters. ([#&#8203;8581](https://redirect.github.com/grpc/grpc-go/issues/8581))

### [`v1.76.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.76.0): Release 1.76.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.75.1...v1.76.0)

### Dependencies

- Minimum supported Go version is now 1.24 ([#&#8203;8509](https://redirect.github.com/grpc/grpc-go/issues/8509))
  - Special Thanks: [@&#8203;kevinGC](https://redirect.github.com/kevinGC)

### Bug Fixes

- client: Return status `INTERNAL` when a server sends zero response messages for a unary or client-streaming RPC. ([#&#8203;8523](https://redirect.github.com/grpc/grpc-go/issues/8523))
- client: Fail RPCs with status `INTERNAL` instead of `UNKNOWN` upon receiving http headers with status 1xx and  `END_STREAM` flag set. ([#&#8203;8518](https://redirect.github.com/grpc/grpc-go/issues/8518))
  - Special Thanks: [@&#8203;vinothkumarr227](https://redirect.github.com/vinothkumarr227)
- pick\_first: Fix race condition that could cause pick\_first to get stuck in `IDLE` state on backend address change. ([#&#8203;8615](https://redirect.github.com/grpc/grpc-go/issues/8615))

### New Features

- credentials: Add `credentials/jwt` package providing file-based JWT PerRPCCredentials (A97). ([#&#8203;8431](https://redirect.github.com/grpc/grpc-go/issues/8431))
  - Special Thanks: [@&#8203;dimpavloff](https://redirect.github.com/dimpavloff)

### Performance Improvements

- client: Improve HTTP/2 header size estimate to reduce re-allocations. ([#&#8203;8547](https://redirect.github.com/grpc/grpc-go/issues/8547))
- encoding/proto: Avoid redundant message size calculation when marshaling. ([#&#8203;8569](https://redirect.github.com/grpc/grpc-go/issues/8569))
  - Special Thanks: [@&#8203;rs-unity](https://redirect.github.com/rs-unity)

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on December 01, 2025 at 08:41 AM UTC

### ℹ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**        | **Change**                                           |
| :----------------- | :--------------------------------------------------- |
| `golang.org/x/net` | `v0.46.0` -> `v0.46.1-0.20251013234738-63d1a5100f82` |

### Comment by @codecov-commenter on December 01, 2025 at 08:47 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1959?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.84%. Comparing base ([`1a02243`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1a02243fe750b65111b1349abdc8bcdc3a350958?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`024aca7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/024aca736f3f1ccf1603217e0e0795d3fe3e97bc?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1959   +/-   ##
=======================================
  Coverage   58.84%   58.84%           
=======================================
  Files         131      131           
  Lines        8436     8436           
=======================================
  Hits         4964     4964           
  Misses       2937     2937           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1959/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1959/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.84% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1959?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @jira-linking on December 06, 2025 at 12:42 AM UTC

Commits missing Jira IDs:
024aca736f3f1ccf1603217e0e0795d3fe3e97bc


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1959*
