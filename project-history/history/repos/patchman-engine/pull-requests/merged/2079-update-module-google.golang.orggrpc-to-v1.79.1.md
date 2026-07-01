---
type: pull_request
number: 2079
title: "Update module google.golang.org/grpc to v1.79.1"
state: merged
author: red-hat-konflux
created: 2026-02-23T08:54:17Z
updated: 2026-02-23T12:53:26Z
closed: 2026-02-23T09:00:04Z
merged: 2026-02-23T09:00:04Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2079
---

# Pull Request #2079: Update module google.golang.org/grpc to v1.79.1

**Author**: @red-hat-konflux
**Created**: February 23, 2026 at 08:54 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.78.0` -> `v1.79.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.79.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.78.0/v1.79.1?slim=true) |

---

### Release Notes

<details>
<summary>grpc/grpc-go (google.golang.org/grpc)</summary>

### [`v1.79.1`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.79.1): Release 1.79.1

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.79.0...v1.79.1)

### Bug Fixes

- grpc: Remove the `-dev` suffix from the User-Agent header. ([#&#8203;8902](https://redirect.github.com/grpc/grpc-go/pull/8902))

### [`v1.79.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.79.0): Release 1.79.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.78.0...v1.79.0)

### API Changes

- mem: Add experimental API `SetDefaultBufferPool` to change the default buffer pool. ([#&#8203;8806](https://redirect.github.com/grpc/grpc-go/issues/8806))
  - Special Thanks: [@&#8203;vanja-p](https://redirect.github.com/vanja-p)
- experimental/stats: Update `MetricsRecorder` to require embedding the new `UnimplementedMetricsRecorder` (a no-op struct) in all implementations for forward compatibility. ([#&#8203;8780](https://redirect.github.com/grpc/grpc-go/issues/8780))

### Behavior Changes

- balancer/weightedtarget: Remove handling of `Addresses` and only handle `Endpoints` in resolver updates. ([#&#8203;8841](https://redirect.github.com/grpc/grpc-go/issues/8841))

### New Features

- experimental/stats: Add support for asynchronous gauge metrics through the new `AsyncMetricReporter` and `RegisterAsyncReporter` APIs. ([#&#8203;8780](https://redirect.github.com/grpc/grpc-go/issues/8780))
- pickfirst: Add support for weighted random shuffling of endpoints, as described in [gRFC A113](https://redirect.github.com/grpc/proposal/pull/535).
  - This is enabled by default, and can be turned off using the environment variable `GRPC_EXPERIMENTAL_PF_WEIGHTED_SHUFFLING`. ([#&#8203;8864](https://redirect.github.com/grpc/grpc-go/issues/8864))
- xds: Implement `:authority` rewriting, as specified in [gRFC A81](https://redirect.github.com/grpc/proposal/blob/master/A81-xds-authority-rewriting.md). ([#&#8203;8779](https://redirect.github.com/grpc/grpc-go/issues/8779))
- balancer/randomsubsetting: Implement the `random_subsetting` LB policy, as specified in [gRFC A68](https://redirect.github.com/grpc/proposal/blob/master/A68-random-subsetting.md). ([#&#8203;8650](https://redirect.github.com/grpc/grpc-go/issues/8650))
  - Special Thanks: [@&#8203;marek-szews](https://redirect.github.com/marek-szews)

### Bug Fixes

- credentials/tls: Fix a bug where the port was not stripped from the authority override before validation. ([#&#8203;8726](https://redirect.github.com/grpc/grpc-go/issues/8726))
  - Special Thanks: [@&#8203;Atul1710](https://redirect.github.com/Atul1710)
- xds/priority: Fix a bug causing delayed failover to lower-priority clusters when a higher-priority cluster is stuck in `CONNECTING` state. ([#&#8203;8813](https://redirect.github.com/grpc/grpc-go/issues/8813))
- health: Fix a bug where health checks failed for clients using legacy compression options (`WithDecompressor` or `RPCDecompressor`). ([#&#8203;8765](https://redirect.github.com/grpc/grpc-go/issues/8765))
  - Special Thanks: [@&#8203;sanki92](https://redirect.github.com/sanki92)
- transport: Fix an issue where the HTTP/2 server could skip header size checks when terminating a stream early. ([#&#8203;8769](https://redirect.github.com/grpc/grpc-go/issues/8769))
  - Special Thanks: [@&#8203;joybestourous](https://redirect.github.com/joybestourous)
- server: Propagate status detail headers, if available, when terminating a stream during request header processing. ([#&#8203;8754](https://redirect.github.com/grpc/grpc-go/issues/8754))
  - Special Thanks: [@&#8203;joybestourous](https://redirect.github.com/joybestourous)

### Performance Improvements

- credentials/alts: Optimize read buffer alignment to reduce copies. ([#&#8203;8791](https://redirect.github.com/grpc/grpc-go/issues/8791))
- mem: Optimize pooling and creation of `buffer` objects.  ([#&#8203;8784](https://redirect.github.com/grpc/grpc-go/issues/8784))
- transport: Reduce slice re-allocations by reserving slice capacity. ([#&#8203;8797](https://redirect.github.com/grpc/grpc-go/issues/8797))

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

### Comment by @github-actions on February 23, 2026 at 08:54 AM UTC

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

### Comment by @codecov-commenter on February 23, 2026 at 09:00 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2079?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.45%. Comparing base ([`63c7cd4`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/63c7cd4e4febb7d71785737cf2068f05298b1a8a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`465a84b`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/465a84ba1a8c66c47e1b971a53933520baabde0f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2079   +/-   ##
=======================================
  Coverage   59.45%   59.45%           
=======================================
  Files         134      134           
  Lines        8699     8699           
=======================================
  Hits         5172     5172           
  Misses       2981     2981           
  Partials      546      546           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2079/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2079/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2079?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2079*
