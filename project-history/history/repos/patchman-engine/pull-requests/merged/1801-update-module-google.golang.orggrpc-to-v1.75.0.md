---
type: pull_request
number: 1801
title: "Update module google.golang.org/grpc to v1.75.0"
state: merged
author: red-hat-konflux
created: 2025-08-24T12:21:14Z
updated: 2025-08-24T12:28:58Z
closed: 2025-08-24T12:28:58Z
merged: 2025-08-24T12:28:57Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1801
---

# Pull Request #1801: Update module google.golang.org/grpc to v1.75.0

**Author**: @red-hat-konflux
**Created**: August 24, 2025 at 12:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.74.2` -> `v1.75.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.75.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.74.2/v1.75.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>grpc/grpc-go (google.golang.org/grpc)</summary>

### [`v1.75.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.75.0): Release 1.75.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.74.2...v1.75.0)

### Behavior Changes

- xds: Remove support for GRPC\_EXPERIMENTAL\_XDS\_FALLBACK environment variable. Fallback support can no longer be disabled. ([#&#8203;8482](https://redirect.github.com/grpc/grpc-go/issues/8482))
- stats: Introduce `DelayedPickComplete` event, a type alias of `PickerUpdated`. ([#&#8203;8465](https://redirect.github.com/grpc/grpc-go/issues/8465))
  - This (combined) event will now be emitted only once per call, when a transport is successfully selected for the attempt.
  - OpenTelemetry metrics will no longer have multiple "Delayed LB pick complete" events in Go, matching other gRPC languages.
  - A future release will delete the `PickerUpdated` symbol.
- credentials: Properly apply `grpc.WithAuthority` as the highest-priority option for setting authority, above the setting in the credentials themselves. ([#&#8203;8488](https://redirect.github.com/grpc/grpc-go/issues/8488))
  - Now that this `WithAuthority` is available, the credentials should not be used to override the authority.
- round\_robin: Randomize the order in which addresses are connected to in order to spread out initial RPC load between clients. ([#&#8203;8438](https://redirect.github.com/grpc/grpc-go/issues/8438))
- server: Return status code INTERNAL when a client sends more than one request in unary and server streaming RPC. ([#&#8203;8385](https://redirect.github.com/grpc/grpc-go/issues/8385))
  - This is a behavior change but also a bug fix to bring gRPC-Go in line with the gRPC spec.

### New Features

- dns: Add an environment variable (`GRPC_ENABLE_TXT_SERVICE_CONFIG`) to provide a way to disable TXT lookups in the DNS resolver (by setting it to `false`).  By default, TXT lookups are enabled, as they were previously. ([#&#8203;8377](https://redirect.github.com/grpc/grpc-go/issues/8377))

### Bug Fixes

- xds: Fix regression preventing empty node IDs in xDS bootstrap configuration. ([#&#8203;8476](https://redirect.github.com/grpc/grpc-go/issues/8476))
  - Special Thanks: [@&#8203;davinci26](https://redirect.github.com/davinci26)
- xds: Fix possible panic when certain invalid resources are encountered. ([#&#8203;8412](https://redirect.github.com/grpc/grpc-go/issues/8412))
  - Special Thanks: [@&#8203;wooffie](https://redirect.github.com/wooffie)
- xdsclient: Fix a rare panic caused by processing a response from a closed server. ([#&#8203;8389](https://redirect.github.com/grpc/grpc-go/issues/8389))
- stats: Fix metric unit formatting by enclosing non-standard units like `call` and `endpoint` in curly braces to comply with UCUM and gRPC OpenTelemetry guidelines. ([#&#8203;8481](https://redirect.github.com/grpc/grpc-go/issues/8481))
- xds: Fix possible panic when clusters are removed from the xds configuration. ([#&#8203;8428](https://redirect.github.com/grpc/grpc-go/issues/8428))
- xdsclient: Fix a race causing "resource doesn not exist" when rapidly subscribing and unsubscribing to the same resource. ([#&#8203;8369](https://redirect.github.com/grpc/grpc-go/issues/8369))
- client: When determining the authority, properly percent-encode (if needed, which is unlikely) when the target string omits the hostname and only specifies a port (`grpc.NewClient(":<port-number-or-name>")`). ([#&#8203;8488](https://redirect.github.com/grpc/grpc-go/issues/8488))

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuNy4wLXJwbSIsInRhcmdldEJyYW5jaCI6Im1hc3RlciIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @jira-linking on August 24, 2025 at 12:21 PM UTC

Commits missing Jira IDs:
cbea8ea65fd39d62261501f232b5670189046884


### Comment by @codecov-commenter on August 24, 2025 at 12:26 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1801?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.85%. Comparing base ([`bda1efa`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/bda1efa1b4e02f58939c1938b04bd9f788fff817?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`cbea8ea`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/cbea8ea65fd39d62261501f232b5670189046884?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1801   +/-   ##
=======================================
  Coverage   54.85%   54.85%           
=======================================
  Files         140      140           
  Lines       10875    10875           
=======================================
  Hits         5966     5966           
  Misses       4373     4373           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1801/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1801/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.85% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1801?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1801*
