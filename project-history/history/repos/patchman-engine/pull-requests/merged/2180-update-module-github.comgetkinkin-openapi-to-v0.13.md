---
type: pull_request
number: 2180
title: "Update module github.com/getkin/kin-openapi to v0.138.0"
state: merged
author: red-hat-konflux
created: 2026-05-04T01:33:45Z
updated: 2026-05-11T05:45:38Z
closed: 2026-05-11T05:39:52Z
merged: 2026-05-11T05:39:52Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2180
---

# Pull Request #2180: Update module github.com/getkin/kin-openapi to v0.138.0

**Author**: @red-hat-konflux
**Created**: May 04, 2026 at 01:33 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/getkin/kin-openapi](https://redirect.github.com/getkin/kin-openapi) | `v0.135.0` → `v0.138.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgetkin%2fkin-openapi/v0.138.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgetkin%2fkin-openapi/v0.135.0/v0.138.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>getkin/kin-openapi (github.com/getkin/kin-openapi)</summary>

### [`v0.138.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.138.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.137.0...v0.138.0)

#### What's Changed

- openapi3gen: clear nullable on exported component bodies by [@&#8203;0-don](https://redirect.github.com/0-don) in [#&#8203;1164](https://redirect.github.com/getkin/kin-openapi/pull/1164)
- openapi3: add test for issue [#&#8203;927](https://redirect.github.com/getkin/kin-openapi/issues/927) (nullable not respected on $ref schemas) by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1165](https://redirect.github.com/getkin/kin-openapi/pull/1165)
- test: move public-API tests to external \_test packages by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1168](https://redirect.github.com/getkin/kin-openapi/pull/1168)
- feat(openapi3): add per-type validation errors with cluster wrappers by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1166](https://redirect.github.com/getkin/kin-openapi/pull/1166)
- feat(openapi3conv): canonicalization pass for 3.0 -> 3.x by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1162](https://redirect.github.com/getkin/kin-openapi/pull/1162)
- openapi3conv: test Upgrade on many documents by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1169](https://redirect.github.com/getkin/kin-openapi/pull/1169)

**Full Changelog**: <https://github.com/getkin/kin-openapi/compare/v0.137.0...v0.138.0>

### [`v0.137.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.137.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.136.0...v0.137.0)

#### What's Changed

- revert to go 1.25 and revert [`cc4f8d9`](https://redirect.github.com/getkin/kin-openapi/commit/cc4f8d99) by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1161](https://redirect.github.com/getkin/kin-openapi/pull/1161)

**Full Changelog**: <https://github.com/getkin/kin-openapi/compare/v0.136.0...v0.137.0>

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

### Comment by @red-hat-konflux on May 04, 2026 at 01:33 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                | **Change**            |
| :------------------------- | :-------------------- |
| `github.com/oasdiff/yaml3` | `v0.0.9` -> `v0.0.12` |

### Comment by @sourcery-ai on May 04, 2026 at 01:33 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

This PR updates the OpenAPI parsing/validation dependency stack by bumping github.com/getkin/kin-openapi from v0.135.0 to v0.137.0 and aligning related indirect dependencies in go.mod/go.sum.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Upgrade kin-openapi dependency and align indirect OpenAPI/yaml/jsonschema libraries. | <ul><li>Bump github.com/getkin/kin-openapi from v0.135.0 to v0.137.0 in go.mod to pick up OpenAPI 3.1 support, determinism fixes, and other upstream changes.</li><li>Update indirect github.com/oasdiff/yaml3 from v0.0.9 to v0.0.12 to match the new kin-openapi requirements.</li><li>Add indirect dependency github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 pulled in by the newer kin-openapi version.</li><li>Refresh go.sum entries accordingly for the updated and new modules.</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on May 04, 2026 at 01:42 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2180?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.21%. Comparing base ([`50aa148`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/50aa148738565d420ac7a53261e3f88b83921bcf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2b0f067`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2b0f0679d520149afd8290a1a6b2e1c8da8c6d5f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2180   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2180/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2180/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.21% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2180?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on May 04, 2026 at 01:35 AM UTC

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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2180*
