---
type: pull_request
number: 2188
title: "Update module google.golang.org/grpc to v1.81.0"
state: merged
author: red-hat-konflux
created: 2026-05-11T01:37:39Z
updated: 2026-05-11T05:39:46Z
closed: 2026-05-11T01:38:05Z
merged: 2026-05-11T01:38:05Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2188
---

# Pull Request #2188: Update module google.golang.org/grpc to v1.81.0

**Author**: @red-hat-konflux
**Created**: May 11, 2026 at 01:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.80.0` → `v1.81.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.81.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.80.0/v1.81.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>grpc/grpc-go (google.golang.org/grpc)</summary>

### [`v1.81.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.81.0): Release 1.81.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.80.0...v1.81.0)

### Behavior Changes

- balancer/rls: Switch gauge metrics to asynchronous emission (once per collection cycle) to reduce telemetry noise and align with other gRPC language implementations. ([#&#8203;8808](https://redirect.github.com/grpc/grpc-go/issues/8808))

### Dependencies

- Minimum supported Go version is now 1.25. ([#&#8203;8969](https://redirect.github.com/grpc/grpc-go/issues/8969))

### Bug Fixes

- xds: Use the leaf cluster's security config for the TLS handshake instead of the aggregate cluster's config. ([#&#8203;8956](https://redirect.github.com/grpc/grpc-go/issues/8956))
- transport: Send a `RST_STREAM` when receiving an `END_STREAM` when the stream is not already half-closed. ([#&#8203;8832](https://redirect.github.com/grpc/grpc-go/issues/8832))
- xds: Fix ADS resource name validation to prevent a panic. ([#&#8203;8970](https://redirect.github.com/grpc/grpc-go/issues/8970))

### New Features

- grpc/stats: Add support for custom labels in per-call metrics ([gRFC A108](https://redirect.github.com/grpc/proposal/blob/master/A108-otel-custom-per-call-label.md)). ([#&#8203;9008](https://redirect.github.com/grpc/grpc-go/issues/9008))
- xds: Add support for Server Name Indication (SNI) and SAN validation ([gRFC A101](https://redirect.github.com/grpc/proposal/blob/master/A101-SNI-setting-and-SNI-SAN-validation.md)). Disabled by default. To enable, set `GRPC_EXPERIMENTAL_XDS_SNI=true` environment variable. ([#&#8203;9016](https://redirect.github.com/grpc/grpc-go/issues/9016))
- xds: Add support to control which fields get propagated from ORCA backend metric reports to LRS load reports ([gRFC A85](https://redirect.github.com/grpc/proposal/blob/master/A85-lrs-custom-metrics-changes.md)). Disabled by default. To enable, set `GRPC_EXPERIMENTAL_XDS_ORCA_LRS_PROPAGATION=true`. ([#&#8203;9005](https://redirect.github.com/grpc/grpc-go/issues/9005))
- xds: Add metrics to track xDS client connectivity and cached resource state ([gRFC A78](https://redirect.github.com/grpc/proposal/blob/master/A78-grpc-metrics-wrr-pf-xds.md)). ([#&#8203;8807](https://redirect.github.com/grpc/grpc-go/issues/8807))
- stats/otel: Enhance `grpc.subchannel.disconnections` metric by adding disconnection reason to the `grpc.disconnect_error` label ([gRFC A94](https://redirect.github.com/grpc/proposal/blob/master/A94-subchannel-otel-metrics.md)). This provides granular insights into why subchannels are closing. ([#&#8203;8973](https://redirect.github.com/grpc/grpc-go/issues/8973))
- mem: Add `mem.Buffer.Slice()` API to slice the buffer like a slice. ([#&#8203;8977](https://redirect.github.com/grpc/grpc-go/issues/8977))
  - Special Thanks: [@&#8203;ash2k](https://redirect.github.com/ash2k)

### Performance Improvements

- alts: Pool read buffers to lower memory utilization when sockets are unreadable. ([#&#8203;8964](https://redirect.github.com/grpc/grpc-go/issues/8964))
- transport: Pool HTTP/2 framer read buffers to reduce idle memory consumption. Currently limited to Linux for ALTS and non-encrypted transports (TCP, Unix). To disable, set `GRPC_GO_EXPERIMENTAL_HTTP_FRAMER_READ_BUFFER_POOLING=false` and report any issues. ([#&#8203;9032](https://redirect.github.com/grpc/grpc-go/issues/9032))

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->

## Summary by Sourcery

Update gRPC and OpenTelemetry dependencies to their latest compatible versions.

New Features:
- Adopt gRPC v1.81.0 for improved metrics, xDS capabilities, and performance enhancements provided by upstream.
- Leverage OpenTelemetry v1.43.0 libraries for updated telemetry and tracing support.

---

## Discussion

### Comment by @red-hat-konflux on May 11, 2026 at 01:37 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 3 additional dependencies were updated


Details:


| **Package**                       | **Change**             |
| :-------------------------------- | :--------------------- |
| `go.opentelemetry.io/otel`        | `v1.41.0` -> `v1.43.0` |
| `go.opentelemetry.io/otel/metric` | `v1.41.0` -> `v1.43.0` |
| `go.opentelemetry.io/otel/trace`  | `v1.41.0` -> `v1.43.0` |

### Comment by @sourcery-ai on May 11, 2026 at 01:37 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Updates the gRPC dependency to v1.81.0 and aligns indirect OpenTelemetry libraries with versions compatible with the new gRPC release.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Bump gRPC library to v1.81.0 and refresh related OpenTelemetry dependencies. | <ul><li>Update google.golang.org/grpc requirement from v1.80.0 to v1.81.0 in go.mod</li><li>Update go.opentelemetry.io/otel, /metric, and /trace indirect dependencies from v1.41.0 to v1.43.0 in go.mod</li><li>Refresh go.sum entries to reflect the new dependency versions</li></ul> | `go.mod`<br/>`go.sum` |

</details>

---

<details>
<summary>Tips and commands</summary>

#### Interacting with Sourcery

- **Trigger a new review:** Comment `@sourcery-ai review` on the pull request.
- **Continue discussions:** Reply directly to Sourcery's review comments.
- **Generate a GitHub issue from a review comment:** Ask Sourcery to create an
  issue from a review comment by replying to it. You can also reply to a
  review comment with `@sourcery-ai issue` to create an issue from it.
- **Generate a pull request title:** Write `@sourcery-ai` anywhere in the pull
  request title to generate a title at any time. You can also comment
  `@sourcery-ai title` on the pull request to (re-)generate the title at any time.
- **Generate a pull request summary:** Write `@sourcery-ai summary` anywhere in
  the pull request body to generate a PR summary at any time exactly where you
  want it. You can also comment `@sourcery-ai summary` on the pull request to
  (re-)generate the summary at any time.
- **Generate reviewer's guide:** Comment `@sourcery-ai guide` on the pull
  request to (re-)generate the reviewer's guide at any time.
- **Resolve all Sourcery comments:** Comment `@sourcery-ai resolve` on the
  pull request to resolve all Sourcery comments. Useful if you've already
  addressed all the comments and don't want to see them anymore.
- **Dismiss all Sourcery reviews:** Comment `@sourcery-ai dismiss` on the pull
  request to dismiss all existing Sourcery reviews. Especially useful if you
  want to start fresh with a new review - don't forget to comment
  `@sourcery-ai review` to trigger a new review!

#### Customizing Your Experience

Access your [dashboard](https://app.sourcery.ai) to:
- Enable or disable review features such as the Sourcery-generated pull request
  summary, the reviewer's guide, and others.
- Change the review language.
- Add, remove or edit custom review instructions.
- Adjust other review settings.

#### Getting Help

- [Contact our support team](mailto:support@sourcery.ai) for questions or feedback.
- Visit our [documentation](https://docs.sourcery.ai) for detailed guides and information.
- Keep in touch with the Sourcery team by following us on [X/Twitter](https://x.com/SourceryAI), [LinkedIn](https://www.linkedin.com/company/sourcery-ai/) or [GitHub](https://github.com/sourcery-ai).

</details>

<!-- Generated by sourcery-ai[bot]: end review_guide -->

### Comment by @codecov-commenter on May 11, 2026 at 01:43 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2188?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.21%. Comparing base ([`e059ec7`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e059ec7c73fc78d0d48727a61d38047224306c28?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`57a6abf`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/57a6abf02eddca1ef4c3d607d3d4ee079b9603f7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2188   +/-   ##
=======================================
  Coverage   59.21%   59.21%           
=======================================
  Files         136      136           
  Lines        8781     8781           
=======================================
  Hits         5200     5200           
  Misses       3035     3035           
  Partials      546      546           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2188/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2188/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.21% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2188?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on May 11, 2026 at 01:38 AM UTC

Hey - I've reviewed your changes and they look great!

***

<details>
<summary>Sourcery is free for open source - if you like our reviews please consider sharing them ✨</summary>

- [X](https://twitter.com/intent/tweet?text=I%20just%20got%20an%20instant%20code%20review%20from%20%40SourceryAI%2C%20and%20it%20was%20brilliant%21%20It%27s%20free%20for%20open%20source%20and%20has%20a%20free%20trial%20for%20private%20code.%20Check%20it%20out%20https%3A//sourcery.ai)
- [Mastodon](https://mastodon.social/share?text=I%20just%20got%20an%20instant%20code%20review%20from%20%40SourceryAI%2C%20and%20it%20was%20brilliant%21%20It%27s%20free%20for%20open%20source%20and%20has%20a%20free%20trial%20for%20private%20code.%20Check%20it%20out%20https%3A//sourcery.ai)
- [LinkedIn](https://www.linkedin.com/sharing/share-offsite/?url=https://sourcery.ai)
- [Facebook](https://www.facebook.com/sharer/sharer.php?u=https://sourcery.ai)

</details>

<sub>
Help me be more useful! Please click 👍 or 👎 on each comment and I'll use the feedback to improve your reviews.
</sub>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2188*
