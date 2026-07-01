---
type: pull_request
number: 75
title: "chore(deps): update module google.golang.org/grpc to v1.78.0 - autoclosed"
state: closed
author: red-hat-konflux
created: 2025-11-14T16:43:13Z
updated: 2026-01-07T08:42:58Z
closed: 2026-01-07T08:42:58Z
base_branch: main
head_branch: konflux/mintmaker/main/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/75
---

# Pull Request #75: chore(deps): update module google.golang.org/grpc to v1.78.0 - autoclosed

**Author**: @red-hat-konflux
**Created**: November 14, 2025 at 04:43 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.76.0` -> `v1.78.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.78.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.76.0/v1.78.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

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

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @red-hat-konflux on November 18, 2025 at 04:39 AM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 5 additional dependencies were updated


Details:


| **Package**                                 | **Change**                                                                   |
| :------------------------------------------ | :--------------------------------------------------------------------------- |
| `go.opentelemetry.io/auto/sdk`              | `v1.1.0` -> `v1.2.1`                                                         |
| `go.opentelemetry.io/otel`                  | `v1.37.0` -> `v1.38.0`                                                       |
| `go.opentelemetry.io/otel/metric`           | `v1.37.0` -> `v1.38.0`                                                       |
| `go.opentelemetry.io/otel/trace`            | `v1.37.0` -> `v1.38.0`                                                       |
| `google.golang.org/genproto/googleapis/api` | `v0.0.0-20251020155222-88f65dc88635` -> `v0.0.0-20251029180050-ab9386a59fda` |

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/75*
