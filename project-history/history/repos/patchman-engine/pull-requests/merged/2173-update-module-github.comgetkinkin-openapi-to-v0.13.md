---
type: pull_request
number: 2173
title: "Update module github.com/getkin/kin-openapi to v0.136.0"
state: merged
author: red-hat-konflux
created: 2026-04-27T01:37:39Z
updated: 2026-04-27T09:49:49Z
closed: 2026-04-27T05:40:20Z
merged: 2026-04-27T05:40:20Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2173
---

# Pull Request #2173: Update module github.com/getkin/kin-openapi to v0.136.0

**Author**: @red-hat-konflux
**Created**: April 27, 2026 at 01:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/getkin/kin-openapi](https://redirect.github.com/getkin/kin-openapi) | `v0.135.0` → `v0.136.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgetkin%2fkin-openapi/v0.136.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgetkin%2fkin-openapi/v0.135.0/v0.136.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>getkin/kin-openapi (github.com/getkin/kin-openapi)</summary>

### [`v0.136.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.136.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.135.0...v0.136.0)

#### What's Changed

- openapi3: stop injecting contentless default in NewResponses() by [@&#8203;0-don](https://redirect.github.com/0-don) in [#&#8203;1148](https://redirect.github.com/getkin/kin-openapi/pull/1148)
- openapi3: standardize Origin json tag to "-" across all types by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1149](https://redirect.github.com/getkin/kin-openapi/pull/1149)
- Update usage message in cmd/validate by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1150](https://redirect.github.com/getkin/kin-openapi/pull/1150)
- openapi3: fix determinism when handling discriminator mappings by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1151](https://redirect.github.com/getkin/kin-openapi/pull/1151)
- feat: bump Go to 1.26 by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1152](https://redirect.github.com/getkin/kin-openapi/pull/1152)
- openapi3: use componentNames for deterministic visitings by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1153](https://redirect.github.com/getkin/kin-openapi/pull/1153)
- feat: add OpenAPI 3.1 support by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1125](https://redirect.github.com/getkin/kin-openapi/pull/1125)
- openapi3: add JoinFunc for custom $ref path resolution by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1154](https://redirect.github.com/getkin/kin-openapi/pull/1154)
- Add many many tests from ApisGuruOpenapiDirectory by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1155](https://redirect.github.com/getkin/kin-openapi/pull/1155)
- openapi3: remove map-iteration order leaks causing flaky tests by [@&#8203;cloudnativeninja](https://redirect.github.com/cloudnativeninja) in [#&#8203;1158](https://redirect.github.com/getkin/kin-openapi/pull/1158)
- openapi2conv: nil-guard components lookup in FromV3SchemaRef by [@&#8203;SAY-5](https://redirect.github.com/SAY-5) in [#&#8203;1156](https://redirect.github.com/getkin/kin-openapi/pull/1156)
- Address various lint errors by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1157](https://redirect.github.com/getkin/kin-openapi/pull/1157)

#### New Contributors

- [@&#8203;0-don](https://redirect.github.com/0-don) made their first contribution in [#&#8203;1148](https://redirect.github.com/getkin/kin-openapi/pull/1148)
- [@&#8203;cloudnativeninja](https://redirect.github.com/cloudnativeninja) made their first contribution in [#&#8203;1158](https://redirect.github.com/getkin/kin-openapi/pull/1158)
- [@&#8203;SAY-5](https://redirect.github.com/SAY-5) made their first contribution in [#&#8203;1156](https://redirect.github.com/getkin/kin-openapi/pull/1156)

**Full Changelog**: <https://github.com/getkin/kin-openapi/compare/v0.135.0...v0.136.0>

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


---

## Discussion

### Comment by @red-hat-konflux on April 27, 2026 at 01:37 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated
- The `go` directive was updated for compatibility reasons


Details:


| **Package**                | **Change**            |
| :------------------------- | :-------------------- |
| `go`                       | `1.25.7` -> `1.26`    |
| `github.com/oasdiff/yaml3` | `v0.0.9` -> `v0.0.12` |

### Comment by @sourcery-ai on April 27, 2026 at 01:37 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Bumps kin-openapi to v0.136.0, aligns transitive dependencies, and updates the module’s Go toolchain version to 1.26 so it matches the new kin-openapi requirements.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Update kin-openapi dependency and align related transitive deps | <ul><li>Bump github.com/getkin/kin-openapi from v0.135.0 to v0.136.0 in go.mod</li><li>Update indirect github.com/oasdiff/yaml3 from v0.0.9 to v0.0.12 to satisfy new kin-openapi graph</li><li>Add indirect dependency github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 required by the updated kin-openapi</li><li>Refresh go.sum entries to match the new dependency versions</li></ul> | `go.mod`<br/>`go.sum` |
| Raise module Go version to satisfy kin-openapi’s minimum version | <ul><li>Increase go directive from 1.25.7 to 1.26 in go.mod so the build uses a Go toolchain compatible with the new kin-openapi release</li></ul> | `go.mod` |

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

---

## Reviews

### Review by @sourcery-ai - Commented on April 27, 2026 at 01:38 AM UTC

Hey - I've left some high level feedback:

- Since the Go version is bumped to 1.26, double-check that your CI/CD images, local toolchains, and any build configs (e.g., Dockerfiles, Makefiles) are aligned to use Go 1.26 to avoid mismatched toolchain issues.
- After updating kin-openapi and related indirect dependencies (yaml3, jsonschema/v6), consider running `go mod tidy` and verifying that only actually needed indirect dependencies remain to keep go.mod/go.sum minimal and consistent.

<details>
<summary>Prompt for AI Agents</summary>

~~~markdown
Please address the comments from this code review:

## Overall Comments
- Since the Go version is bumped to 1.26, double-check that your CI/CD images, local toolchains, and any build configs (e.g., Dockerfiles, Makefiles) are aligned to use Go 1.26 to avoid mismatched toolchain issues.
- After updating kin-openapi and related indirect dependencies (yaml3, jsonschema/v6), consider running `go mod tidy` and verifying that only actually needed indirect dependencies remain to keep go.mod/go.sum minimal and consistent.
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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2173*
