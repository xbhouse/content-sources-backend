---
type: pull_request
number: 2135
title: "Update module github.com/gin-contrib/timeout to v1.2.1"
state: merged
author: red-hat-konflux
created: 2026-03-30T01:26:16Z
updated: 2026-03-30T05:13:53Z
closed: 2026-03-30T01:26:33Z
merged: 2026-03-30T01:26:33Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gin-contrib-timeout-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2135
---

# Pull Request #2135: Update module github.com/gin-contrib/timeout to v1.2.1

**Author**: @red-hat-konflux
**Created**: March 30, 2026 at 01:26 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gin-contrib-timeout-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/gin-contrib/timeout](https://redirect.github.com/gin-contrib/timeout) | `v1.1.0` -> `v1.2.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgin-contrib%2ftimeout/v1.2.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgin-contrib%2ftimeout/v1.1.0/v1.2.1?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>gin-contrib/timeout (github.com/gin-contrib/timeout)</summary>

### [`v1.2.1`](https://redirect.github.com/gin-contrib/timeout/releases/tag/v1.2.1)

[Compare Source](https://redirect.github.com/gin-contrib/timeout/compare/v1.2.0...v1.2.1)

#### Changelog

##### Bug fixes

- [`2d46de1`](https://redirect.github.com/gin-contrib/timeout/commit/2d46de11bf5f41b9e34ec2c178c2d2b2421b7eb6): fix: resolve data race on gin.Context index field ([#&#8203;84](https://redirect.github.com/gin-contrib/timeout/issues/84)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`4eb7613`](https://redirect.github.com/gin-contrib/timeout/commit/4eb761355c81a65f9f5a35c969e2d7590faa21a1): fix(writer): allow multiple WriteHeader calls for static file serving ([#&#8203;85](https://redirect.github.com/gin-contrib/timeout/issues/85)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Build process updates

- [`14337c7`](https://redirect.github.com/gin-contrib/timeout/commit/14337c77cc76636b1bf308304453526e343f784f): ci(actions): bump codecov/codecov-action from v5 to v6 ([@&#8203;appleboy](https://redirect.github.com/appleboy))

### [`v1.2.0`](https://redirect.github.com/gin-contrib/timeout/releases/tag/v1.2.0)

[Compare Source](https://redirect.github.com/gin-contrib/timeout/compare/v1.1.0...v1.2.0)

#### Changelog

##### Features

- [`9c9ea6f`](https://redirect.github.com/gin-contrib/timeout/commit/9c9ea6f127a193c26cc262b8699749de31a02301): feat: track response body size in Writer and expose Size method ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Enhancements

- [`9cbff66`](https://redirect.github.com/gin-contrib/timeout/commit/9cbff66ba7adf73f8f86e225b1bc0c7090741489): chore: update dependencies and refresh lockfiles to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`8998fe0`](https://redirect.github.com/gin-contrib/timeout/commit/8998fe07a8ab5745504fcac8253deb0ea5cc835a): chore: update Go modules and dependencies to latest versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`5780699`](https://redirect.github.com/gin-contrib/timeout/commit/578069934732fe8ca687f44ce8114c37e82fcaf9): chore: remove bearer.yml workflow ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`d48bec6`](https://redirect.github.com/gin-contrib/timeout/commit/d48bec603691f6edc0d225fe830361562a6a5c87): chore(deps): bump actions/checkout from 4 to 6 ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`c09c917`](https://redirect.github.com/gin-contrib/timeout/commit/c09c917f0465e40fc23adc770eadb7d5ede31468): chore: drop Go 1.23 support, require Go 1.24+ ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`c818cbf`](https://redirect.github.com/gin-contrib/timeout/commit/c818cbfaf7998315865e8868893cb8af2a7e735c): chore(deps): upgrade quic-go to v0.57.1 ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`5de17bd`](https://redirect.github.com/gin-contrib/timeout/commit/5de17bd480ecc8176e34ab0a8c4ad88984383ba8): chore(ci): update golangci-lint to v2.6 ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`71d1ccc`](https://redirect.github.com/gin-contrib/timeout/commit/71d1ccc35a1b322efaa0f2c8eb87d0398adaf14a): chore(deps): bump actions/cache from 4 to 5 ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`0c2a03b`](https://redirect.github.com/gin-contrib/timeout/commit/0c2a03b2b10f0d891cb302eb02836d1f531641ba): chore(ci): upgrade trivy-action from 0.33.1 to 0.35.0 ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`85b6248`](https://redirect.github.com/gin-contrib/timeout/commit/85b6248b6890a09a371f37a0d11a925d492ae149): chore(deps): upgrade gin to v1.12.0 and update CI Go versions ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`c5a777e`](https://redirect.github.com/gin-contrib/timeout/commit/c5a777e0d42b18cebbe7d6c60407467bd3b7450f): chore(deps): upgrade golang.org/x/text to v0.35.0 ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Refactor

- [`68ce8b8`](https://redirect.github.com/gin-contrib/timeout/commit/68ce8b88cc56e5a87541747c12ad1e7349b4c80e): refactor: update Go modules and refactor timeout handling in examples ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`b2e7eb5`](https://redirect.github.com/gin-contrib/timeout/commit/b2e7eb5baf810003e0959fd086afe7fa8f31c254): refactor: refactor panic recovery handling for improved clarity and reuse ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`04d4cff`](https://redirect.github.com/gin-contrib/timeout/commit/04d4cffb8e3291d8cc69a7a1a9b3610682c8ba9d): refactor: refactor and enhance tests for improved clarity and coverage ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Build process updates

- [`888c691`](https://redirect.github.com/gin-contrib/timeout/commit/888c6911b6605bc0961217ca3df79b4d8dd25f97): ci: enable race detection in CI test pipeline ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`2ec65b7`](https://redirect.github.com/gin-contrib/timeout/commit/2ec65b7ca32ade3689557254d33604c22e34b87c): ci: add CI matrix for race and non-race Go test runs ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`6e93116`](https://redirect.github.com/gin-contrib/timeout/commit/6e93116ba8b1fb1b7eff1716c99665891f6d4b7f): ci: update CI workflow for Go 1.25 and job name consistency ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`94f1626`](https://redirect.github.com/gin-contrib/timeout/commit/94f16261a1623d274416eecca216e5fa8c26578c): ci: integrate Trivy security scanning workflow and reporting ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`7d57230`](https://redirect.github.com/gin-contrib/timeout/commit/7d572307e8e5784f0496bf5adb2e609a816b82fb): ci(workflow): bump goreleaser/goreleaser-action from v6 to v7 ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Documentation updates

- [`5a95b3e`](https://redirect.github.com/gin-contrib/timeout/commit/5a95b3e049166fea22e228c32b5198626339f7de): docs: revise and clarify example usage and documentation ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`042f29e`](https://redirect.github.com/gin-contrib/timeout/commit/042f29ecf4c5111d94c4261ef693f25b0e78fb22): docs: remove Gitter chat badge from documentation ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`c83a0a8`](https://redirect.github.com/gin-contrib/timeout/commit/c83a0a872948fb22107b1cf1edb7ef5b025b32ea): docs: revamp documentation with expanded guides and usage examples ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Others

- [`a1ae889`](https://redirect.github.com/gin-contrib/timeout/commit/a1ae889173dbef3bea2b251d782ce31772ead7c2): test: improve BufferPool robustness and testing clarity ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`41886ff`](https://redirect.github.com/gin-contrib/timeout/commit/41886ff9b39bbcb7b5d198a7e727a39fd2db636f): test: improve timeout option handling and response validation in tests ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`58b1a9e`](https://redirect.github.com/gin-contrib/timeout/commit/58b1a9e8404f2282a65af4cf8044541012917524): style: standardize code formatting and enhance readability ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`1343c93`](https://redirect.github.com/gin-contrib/timeout/commit/1343c93cc737d814588c58b0ffd0949585328f06): Add Go 1.26 to GitHub Actions test matrix ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`e5dd31e`](https://redirect.github.com/gin-contrib/timeout/commit/e5dd31e5074da26639f0db78c4760a7da1c63b63): Update golangci-lint version to v2.9 ([@&#8203;appleboy](https://redirect.github.com/appleboy))

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

## Summary by Sourcery

Update the gin-contrib/timeout dependency to the latest 1.2.1 release.

New Features:
- Inherit upstream support for tracking response body size via the timeout middleware writer.

Bug Fixes:
- Pull in upstream fixes for a gin.Context data race and improved handling of multiple WriteHeader calls when serving static files.

Enhancements:
- Adopt upstream dependency, CI, and documentation improvements included in gin-contrib/timeout v1.2.x.

---

## Discussion

### Comment by @github-actions on March 30, 2026 at 01:26 AM UTC

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

### Comment by @sourcery-ai on March 30, 2026 at 01:26 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Updates the gin-contrib/timeout middleware dependency from v1.1.0 to v1.2.1 in go.mod (and corresponding entries in go.sum), pulling in upstream bug fixes, new response size tracking, and internal refactors without changing local application code.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Bump gin-contrib/timeout dependency to v1.2.1 and refresh module checksums. | <ul><li>Update required version of github.com/gin-contrib/timeout from v1.1.0 to v1.2.1 in go.mod.</li><li>Regenerate go.sum entries to match the new timeout module version and its transitive dependencies.</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on March 30, 2026 at 01:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2135?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.18%. Comparing base ([`c4888aa`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c4888aa8434a265eb6e68e683ee57a42a5193cd9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`c26991f`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c26991f306463daa2061112d1179be28eba518cf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2135   +/-   ##
=======================================
  Coverage   59.18%   59.18%           
=======================================
  Files         134      134           
  Lines        8624     8624           
=======================================
  Hits         5104     5104           
  Misses       2985     2985           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2135/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2135/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.18% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2135?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on March 30, 2026 at 01:29 AM UTC

Hey - I've left some high level feedback:

- Given that gin-contrib/timeout v1.2.x drops older Go support and requires a newer Go toolchain, please confirm that your CI and deployment environments are configured to use a compatible Go version to avoid build/runtime issues.

<details>
<summary>Prompt for AI Agents</summary>

~~~markdown
Please address the comments from this code review:

## Overall Comments
- Given that gin-contrib/timeout v1.2.x drops older Go support and requires a newer Go toolchain, please confirm that your CI and deployment environments are configured to use a compatible Go version to avoid build/runtime issues.
~~~

</details>

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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2135*
