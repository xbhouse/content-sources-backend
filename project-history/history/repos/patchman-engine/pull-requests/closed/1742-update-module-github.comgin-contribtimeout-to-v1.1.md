---
type: pull_request
number: 1742
title: "Update module github.com/gin-contrib/timeout to v1.1.0"
state: closed
author: red-hat-konflux
created: 2025-07-20T09:25:43Z
updated: 2025-07-28T10:34:36Z
closed: 2025-07-28T08:42:59Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gin-contrib-timeout-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1742
---

# Pull Request #1742: Update module github.com/gin-contrib/timeout to v1.1.0

**Author**: @red-hat-konflux
**Created**: July 20, 2025 at 09:25 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gin-contrib-timeout-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/gin-contrib/timeout](https://redirect.github.com/gin-contrib/timeout) | `v1.0.2` -> `v1.1.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgin-contrib%2ftimeout/v1.1.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgin-contrib%2ftimeout/v1.0.2/v1.1.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>gin-contrib/timeout (github.com/gin-contrib/timeout)</summary>

### [`v1.1.0`](https://redirect.github.com/gin-contrib/timeout/releases/tag/v1.1.0)

[Compare Source](https://redirect.github.com/gin-contrib/timeout/compare/v1.0.2...v1.1.0)

#### Changelog

##### Features

- [`ff7b0b4`](https://redirect.github.com/gin-contrib/timeout/commit/ff7b0b426c913268be0afdf4c4fada38e0186e73): feat: refactor and enhance panic capture and debug response handling ([#&#8203;76](https://redirect.github.com/gin-contrib/timeout/issues/76)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`33ca6b7`](https://redirect.github.com/gin-contrib/timeout/commit/33ca6b767283efbf08966f0fabe24d7074970a84): feat: add Gin timeout and admin middleware example with tests ([#&#8203;78](https://redirect.github.com/gin-contrib/timeout/issues/78)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Bug fixes

- [`c5c090f`](https://redirect.github.com/gin-contrib/timeout/commit/c5c090f895efdaeccba86f94c553d82bb3aae31a): fix(response): conflict when handler completed and concurrent map writes ([#&#8203;55](https://redirect.github.com/gin-contrib/timeout/issues/55)) ([@&#8203;demouth](https://redirect.github.com/demouth))
- [`ba241c4`](https://redirect.github.com/gin-contrib/timeout/commit/ba241c4c3e55f26e2e2e4faaadfc03928907904b): fix: prevent middleware execution after request timeout ([#&#8203;74](https://redirect.github.com/gin-contrib/timeout/issues/74)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`998b6fe`](https://redirect.github.com/gin-contrib/timeout/commit/998b6fe375eb676a78db17279d33db5957f95c58): fix(AbortWithStatus): the middleware is unable to successfully set the http headers once calling gin.Context.AbortWithStatus ([#&#8203;72](https://redirect.github.com/gin-contrib/timeout/issues/72)) ([@&#8203;zhyee](https://redirect.github.com/zhyee))
- [`0e9cc40`](https://redirect.github.com/gin-contrib/timeout/commit/0e9cc404ce156e74feb0840f44f43859d401c995): fix: refactor Gin context and timeout handling for safer concurrency ([#&#8203;77](https://redirect.github.com/gin-contrib/timeout/issues/77)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Enhancements

- [`0de674d`](https://redirect.github.com/gin-contrib/timeout/commit/0de674d1491a716807da28e138a67bd87769a2ac): chore: update Go version to 1.23 and refresh dependencies ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`91d5013`](https://redirect.github.com/gin-contrib/timeout/commit/91d501322aa3cfe22ba2becedac16716fb407c3c): chore: refactor golangci-lint configuration and update dependencies ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Build process updates

- [`b6869b3`](https://redirect.github.com/gin-contrib/timeout/commit/b6869b3210c7391f4c50f58c73cd14efc348f932): ci: update CI pipeline and testing configurations ([@&#8203;appleboy](https://redirect.github.com/appleboy))

##### Documentation updates

- [`8a050ef`](https://redirect.github.com/gin-contrib/timeout/commit/8a050ef7029876be574edd4a37008bce71079866): docs: improve buffer management with comments and performance notes ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`8ae792b`](https://redirect.github.com/gin-contrib/timeout/commit/8ae792b55ebc72ef554e9ac5006aac46d365095f): docs(example): add Gin server example with concurrent timeout handling ([#&#8203;75](https://redirect.github.com/gin-contrib/timeout/issues/75)) ([@&#8203;appleboy](https://redirect.github.com/appleboy))
- [`80a9f69`](https://redirect.github.com/gin-contrib/timeout/commit/80a9f69ade39f76f0505794bab506e540629ed46): docs: revise and expand documentation for middleware usage and setup ([@&#8203;appleboy](https://redirect.github.com/appleboy))

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4yNjQuMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on July 20, 2025 at 09:25 AM UTC

Commits missing Jira IDs:
9d3dc5a0d34e2f309c0497ec855783afa2af69ff


### Comment by @sourcery-ai on July 20, 2025 at 09:25 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

This PR bumps the github.com/gin-contrib/timeout module from v1.0.2 to v1.1.0 by updating the require directive in go.mod and regenerating go.sum to include the new checksums.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Bump timeout dependency version | <ul><li>Update require entry in go.mod from v1.0.2 to v1.1.0</li></ul> | `go.mod` |
| Refresh checksum file | <ul><li>Regenerate go.sum to reflect the new module version and its transitive dependencies</li></ul> | `go.sum` |

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

### Comment by @red-hat-konflux on July 28, 2025 at 10:34 AM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`v1.1.0`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1742*
