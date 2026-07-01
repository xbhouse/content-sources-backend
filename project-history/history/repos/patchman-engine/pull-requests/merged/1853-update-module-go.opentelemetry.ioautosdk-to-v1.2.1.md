---
type: pull_request
number: 1853
title: "Update module go.opentelemetry.io/auto/sdk to v1.2.1"
state: merged
author: red-hat-konflux
created: 2025-09-21T04:22:44Z
updated: 2025-10-08T16:16:26Z
closed: 2025-09-22T12:23:07Z
merged: 2025-09-22T12:23:07Z
base_branch: master
head_branch: konflux/mintmaker/master/go.opentelemetry.io-auto-sdk-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1853
---

# Pull Request #1853: Update module go.opentelemetry.io/auto/sdk to v1.2.1

**Author**: @red-hat-konflux
**Created**: September 21, 2025 at 04:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/go.opentelemetry.io-auto-sdk-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [go.opentelemetry.io/auto/sdk](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation) | `v1.1.0` -> `v1.2.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/go.opentelemetry.io%2fauto%2fsdk/v1.2.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/go.opentelemetry.io%2fauto%2fsdk/v1.1.0/v1.2.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>open-telemetry/opentelemetry-go-instrumentation (go.opentelemetry.io/auto/sdk)</summary>

### [`v1.2.1`](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/blob/HEAD/CHANGELOG.md#goopentelemetryioautosdk-v121---2025-09-15)

##### Fixed

- Fix `uint32` bounding on 32 bit architectures in the `go.opentelemetry.io/auto/sdk` module. ([#&#8203;2810](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2810))

### [`v1.2.0`](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/blob/HEAD/CHANGELOG.md#v0230v120---2025-09-10)

<!-- markdownlint-disable MD028 -->

> \[!NOTE]
> This is the last release version that will support building the auto-instrumentation CLI using Go 1.23.
> The next release will require development to be done using Go >= 1.24.

<!-- markdownlint-enable MD028 -->

##### Added

- Cache offsets for `golang.org/x/net` `0.42.0`. ([#&#8203;2503](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2503))
- Cache offsets for `google.golang.org/grpc` `1.74.2`. ([#&#8203;2546](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2546))
- Cache offsets for `google.golang.org/grpc` `1.76.0-dev`. ([#&#8203;2596](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2596))
- Allow configuration of the resource using a [resource.Detector]. ([#&#8203;2598](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2598))
  - The `WithResourceDetector` function is added to `go.opentelemetry.io/auto/pipeline/otelsdk`.
  - The `WithEnv` function is updated to parse the `OTEL_RESOURCE_DETECTOR` environment variable.
    Values are expected to be a comma-separated list of resource detector IDs registered with the [`autodetect` package].
- Cache offsets for Go `1.23.12`. ([#&#8203;2603](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2603))
- Cache offsets for Go `1.24.6`. ([#&#8203;2603](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2603))
- Cache offsets for `golang.org/x/net` `0.43.0`. ([#&#8203;2615](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2615))
- Cache offsets for Go `1.25.0`. ([#&#8203;2651](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2651))
- Cache offsets for `google.golang.org/grpc` `1.75.0`. ([#&#8203;2686](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2686))
- Cache offsets for `github.com/segmentio/kafka-go` `0.4.49`. ([#&#8203;2699](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2699))
- Cache offsets for `go.opentelemetry.io/otel` `v1.38.0`. ([#&#8203;2726](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2726))
- Cache offsets for Go `1.24.7`. ([#&#8203;2747](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2747))
- Cache offsets for Go `1.25.1`. ([#&#8203;2747](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2747))
- Cache offsets for `golang.org/x/net` `0.44.0`. ([#&#8203;2773](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2773))
- Cache offsets for `google.golang.org/grpc` `1.72.3`. ([#&#8203;2787](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2787))
- Cache offsets for `google.golang.org/grpc` `1.73.1`. ([#&#8203;2787](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2787))
- Cache offsets for `google.golang.org/grpc` `1.74.3`. ([#&#8203;2787](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2787))
- Cache offsets for `google.golang.org/grpc` `1.75.1`. ([#&#8203;2787](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2787))

[resource.Detector]: https://pkg.go.dev/go.opentelemetry.io/otel/sdk/resource#Detector

[`autodetect` package]: https://pkg.go.dev/go.opentelemetry.io/contrib/detectors/autodetect

##### Changed

- Upgrade `go.opentelemetry.io/auto` semconv to `v1.37.0`. ([#&#8203;2763](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2763))
- Upgrade `go.opentelemetry.io/auto/sdk` semconv to `v1.37.0`. ([#&#8203;2763](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2763))

##### Fixed

- Add `telemetry.distro.version` resource attribute to the `otelsdk` handler. ([#&#8203;2383](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2383))
- `active_spans_by_span_ptr` eBPF map used in the traceglobal probe changed to LRU. ([#&#8203;2509](https://redirect.github.com/open-telemetry/opentelemetry-go-instrumentation/pull/2509))

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on September 21, 2025 at 04:22 AM UTC

Commits missing Jira IDs:
d73c4d83d9d46b7f691647fa2007b4bdf0fee764


### Comment by @codecov-commenter on September 22, 2025 at 12:23 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1853?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.53%. Comparing base ([`2700693`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/270069310a271e11705707b6991489d8d865dda2?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d73c4d8`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d73c4d83d9d46b7f691647fa2007b4bdf0fee764?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 8 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1853   +/-   ##
=======================================
  Coverage   54.53%   54.53%           
=======================================
  Files         138      138           
  Lines       10743    10743           
=======================================
  Hits         5859     5859           
  Misses       4346     4346           
  Partials      538      538           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1853/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1853/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.53% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1853?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1853*
