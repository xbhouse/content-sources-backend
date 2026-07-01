---
type: pull_request
number: 2209
title: "Update module github.com/getkin/kin-openapi to v0.139.0"
state: merged
author: red-hat-konflux
created: 2026-05-25T05:37:19Z
updated: 2026-05-25T09:36:35Z
closed: 2026-05-25T05:37:34Z
merged: 2026-05-25T05:37:33Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2209
---

# Pull Request #2209: Update module github.com/getkin/kin-openapi to v0.139.0

**Author**: @red-hat-konflux
**Created**: May 25, 2026 at 05:37 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/getkin/kin-openapi](https://redirect.github.com/getkin/kin-openapi) | `v0.138.0` → `v0.139.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgetkin%2fkin-openapi/v0.139.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgetkin%2fkin-openapi/v0.138.0/v0.139.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>getkin/kin-openapi (github.com/getkin/kin-openapi)</summary>

### [`v0.139.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.139.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.138.0...v0.139.0)

#### What's Changed

- feat(openapi3): batch-convert long-tail RequiredFieldError sites by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1170](https://redirect.github.com/getkin/kin-openapi/pull/1170)
- feat(openapi3): typed validation error clusters (combined: [#&#8203;1171](https://redirect.github.com/getkin/kin-openapi/issues/1171)-[#&#8203;1179](https://redirect.github.com/getkin/kin-openapi/issues/1179)) by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1180](https://redirect.github.com/getkin/kin-openapi/pull/1180)
- openapi3gen: skip component export for anonymous types by [@&#8203;0-don](https://redirect.github.com/0-don) in [#&#8203;1163](https://redirect.github.com/getkin/kin-openapi/pull/1163)
- feat: migrate to oasdiff/yaml v0.1.0 single Unmarshal API + enable DisableTimestamps by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1181](https://redirect.github.com/getkin/kin-openapi/pull/1181)
- openapi3: typed context errors for Validate() wrapper chain by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1183](https://redirect.github.com/getkin/kin-openapi/pull/1183)
- openapi3: track Origin on the document root (T) by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1184](https://redirect.github.com/getkin/kin-openapi/pull/1184)
- openapi3: tests flakiness corrected by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [#&#8203;1159](https://redirect.github.com/getkin/kin-openapi/pull/1159)
- openapi3: aggregate independent validation errors via EnableMultiError by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1185](https://redirect.github.com/getkin/kin-openapi/pull/1185)
- openapi3: fix validation of duplicated path templates by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1189](https://redirect.github.com/getkin/kin-openapi/pull/1189)
- openapi3: type the remaining bare-error validation sites by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1187](https://redirect.github.com/getkin/kin-openapi/pull/1187)

**Full Changelog**: <https://github.com/getkin/kin-openapi/compare/v0.138.0...v0.139.0>

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

Update OpenAPI-related dependencies to their latest compatible versions.

Enhancements:
- Bump github.com/getkin/kin-openapi from v0.138.0 to v0.139.0 to pick up the latest OpenAPI validation and generation improvements.
- Refresh indirect yaml parsing dependencies github.com/oasdiff/yaml and github.com/oasdiff/yaml3 to newer releases aligned with the updated kin-openapi version.

---

## Discussion

### Comment by @red-hat-konflux on May 25, 2026 at 05:37 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 2 additional dependencies were updated


Details:


| **Package**                | **Change**             |
| :------------------------- | :--------------------- |
| `github.com/oasdiff/yaml`  | `v0.0.9` -> `v0.1.0`   |
| `github.com/oasdiff/yaml3` | `v0.0.12` -> `v0.0.13` |

### Comment by @sourcery-ai on May 25, 2026 at 05:37 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Bumps github.com/getkin/kin-openapi from v0.138.0 to v0.139.0 and aligns its transitive yaml dependencies to the versions expected by the new release, updating go.mod and corresponding go.sum entries.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Upgrade kin-openapi dependency to v0.139.0 and refresh related yaml library versions. | <ul><li>Update github.com/getkin/kin-openapi requirement from v0.138.0 to v0.139.0 in go.mod.</li><li>Bump indirect github.com/oasdiff/yaml from v0.0.9 to v0.1.0 and github.com/oasdiff/yaml3 from v0.0.12 to v0.0.13 in go.mod to match kin-openapi’s new requirements.</li><li>Regenerate go.sum entries to reflect the new versions and checksums of kin-openapi and yaml dependencies.</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on May 25, 2026 at 05:43 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2209?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.23%. Comparing base ([`adb0815`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/adb0815953c59af048c75f0180633226f56c1f89?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a24ec84`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/a24ec84d49d2ca8b24bbc6ef0a205a82f7399da9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2209   +/-   ##
=======================================
  Coverage   59.23%   59.23%           
=======================================
  Files         137      137           
  Lines        8795     8795           
=======================================
  Hits         5210     5210           
  Misses       3037     3037           
  Partials      548      548           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2209/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2209/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.23% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2209?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on May 25, 2026 at 05:37 AM UTC

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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2209*
