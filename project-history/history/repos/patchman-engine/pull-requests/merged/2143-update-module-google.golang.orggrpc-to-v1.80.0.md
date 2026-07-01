---
type: pull_request
number: 2143
title: "Update module google.golang.org/grpc to v1.80.0"
state: merged
author: red-hat-konflux
created: 2026-04-06T01:22:20Z
updated: 2026-04-06T05:18:07Z
closed: 2026-04-06T01:22:38Z
merged: 2026-04-06T01:22:38Z
base_branch: master
head_branch: konflux/mintmaker/master/google.golang.org-grpc-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2143
---

# Pull Request #2143: Update module google.golang.org/grpc to v1.80.0

**Author**: @red-hat-konflux
**Created**: April 06, 2026 at 01:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/google.golang.org-grpc-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [google.golang.org/grpc](https://redirect.github.com/grpc/grpc-go) | `v1.79.3` → `v1.80.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/google.golang.org%2fgrpc/v1.80.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/google.golang.org%2fgrpc/v1.79.3/v1.80.0?slim=true) |

---

### Release Notes

<details>
<summary>grpc/grpc-go (google.golang.org/grpc)</summary>

### [`v1.80.0`](https://redirect.github.com/grpc/grpc-go/releases/tag/v1.80.0): Release 1.80.0

[Compare Source](https://redirect.github.com/grpc/grpc-go/compare/v1.79.3...v1.80.0)

### Behavior Changes

- balancer: log a warning if a balancer is registered with uppercase letters, as balancer names should be lowercase. In a future release, balancer names will be treated as case-insensitive; see [#&#8203;5288](https://redirect.github.com/grpc/grpc-go/issues/5288) for details. ([#&#8203;8837](https://redirect.github.com/grpc/grpc-go/issues/8837))
- xds: update resource error handling and re-resolution logic ([#&#8203;8907](https://redirect.github.com/grpc/grpc-go/issues/8907))
  - Re-resolve all `LOGICAL_DNS` clusters simultaneously when re-resolution is requested.
  - Fail all in-flight RPCs immediately upon receipt of listener or route resource errors, instead of allowing them to complete.

### Bug Fixes

- xds: support the LB policy configured in `LOGICAL_DNS` cluster resources instead of defaulting to `pick_first`. ([#&#8203;8733](https://redirect.github.com/grpc/grpc-go/issues/8733))
- credentials/tls: perform per-RPC authority validation against the leaf certificate instead of the entire peer certificate chain. ([#&#8203;8831](https://redirect.github.com/grpc/grpc-go/issues/8831))
- xds: enabling A76 ring hash endpoint keys no longer causes EDS resources with invalid proxy metadata to be NACKed when HTTP CONNECT (gRFC A86) is disabled. ([#&#8203;8875](https://redirect.github.com/grpc/grpc-go/issues/8875))
- xds: validate that the sum of endpoint weights in a locality does not exceed the maximum `uint32` value. ([#&#8203;8899](https://redirect.github.com/grpc/grpc-go/issues/8899))
  - Special Thanks: [@&#8203;RAVEYUS](https://redirect.github.com/RAVEYUS)
- xds: fix incorrect proto field access in the weighted round robin (WRR) configuration where `blackout_period` was used instead of `weight_expiration_period`. ([#&#8203;8915](https://redirect.github.com/grpc/grpc-go/issues/8915))
  - Special Thanks: [@&#8203;gregbarasch](https://redirect.github.com/gregbarasch)
- xds/rbac: handle addresses with ports in IP matchers. ([#&#8203;8990](https://redirect.github.com/grpc/grpc-go/issues/8990))

### New Features

- ringhash: enable gRFC A76 (endpoint hash keys and request hash headers) by default. ([#&#8203;8922](https://redirect.github.com/grpc/grpc-go/issues/8922))

### Performance Improvements

- credentials/alts: pool write buffers to reduce memory allocations and usage. ([#&#8203;8919](https://redirect.github.com/grpc/grpc-go/issues/8919))
- grpc: enable the use of pooled write buffers for buffering HTTP/2 frame writes by default. This reduces memory usage when connections are idle. Use the [WithSharedWriteBuffer](https://pkg.go.dev/google.golang.org/grpc#WithSharedWriteBuffer) dial option or the [SharedWriteBuffer](https://pkg.go.dev/google.golang.org/grpc#SharedWriteBuffer) server option to disable this feature. ([#&#8203;8957](https://redirect.github.com/grpc/grpc-go/issues/8957))
- xds/priority: stop caching child LB policies removed from the configuration. This will help reduce memory and cpu usage when localities are constantly switching between priorities. ([#&#8203;8997](https://redirect.github.com/grpc/grpc-go/issues/8997))
- mem: add a faster tiered buffer pool; use the experimental [mem.NewBinaryTieredBufferPool](https://pkg.go.dev/google.golang.org/grpc/mem@master#NewBinaryTieredBufferPool) function to create such pools. ([#&#8203;8775](https://redirect.github.com/grpc/grpc-go/issues/8775))

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

Build:
- Bump google.golang.org/grpc module from v1.79.3 to v1.80.0.

---

## Discussion

### Comment by @github-actions on April 06, 2026 at 01:22 AM UTC

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

### Comment by @sourcery-ai on April 06, 2026 at 01:22 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Updates google.golang.org/grpc from v1.79.3 to v1.80.0 in go.mod (and go.sum accordingly), pulling in upstream behavior changes, bug fixes, performance improvements, and new features from grpc-go 1.80.0.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Bump grpc-go dependency to v1.80.0 and refresh dependency checksums. | <ul><li>Update google.golang.org/grpc requirement in go.mod from v1.79.3 to v1.80.0.</li><li>Regenerate go.sum entries to match the new grpc-go version and its transitive dependencies.</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on April 06, 2026 at 01:28 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2143?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.16%. Comparing base ([`cb12aaf`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/cb12aafbbe4e667512446715890a62fce7bfbb23?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`ba66d12`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/ba66d12871f9d1f738ae6919218982f4c9f4a063?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2143   +/-   ##
=======================================
  Coverage   59.16%   59.16%           
=======================================
  Files         134      134           
  Lines        8628     8628           
=======================================
  Hits         5105     5105           
  Misses       2987     2987           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2143/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2143/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2143?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on April 06, 2026 at 01:24 AM UTC

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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2143*
