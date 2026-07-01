---
type: pull_request
number: 2150
title: "Update module github.com/getkin/kin-openapi to v0.135.0"
state: merged
author: red-hat-konflux
created: 2026-04-13T01:29:57Z
updated: 2026-04-13T05:17:44Z
closed: 2026-04-13T01:30:21Z
merged: 2026-04-13T01:30:21Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2150
---

# Pull Request #2150: Update module github.com/getkin/kin-openapi to v0.135.0

**Author**: @red-hat-konflux
**Created**: April 13, 2026 at 01:29 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/getkin/kin-openapi](https://redirect.github.com/getkin/kin-openapi) | `v0.134.0` → `v0.135.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgetkin%2fkin-openapi/v0.135.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgetkin%2fkin-openapi/v0.134.0/v0.135.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>getkin/kin-openapi (github.com/getkin/kin-openapi)</summary>

### [`v0.135.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.135.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.134.0...v0.135.0)

#### What's Changed

- openapi3: strip **origin** from Encodings and ServerVariables maps by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1132](https://redirect.github.com/getkin/kin-openapi/pull/1132)
- fix: update yaml3 to prevent panic on empty mapping node in sequence by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1133](https://redirect.github.com/getkin/kin-openapi/pull/1133)
- openapi3: strip **origin** from extension values to prevent spurious diffs by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1137](https://redirect.github.com/getkin/kin-openapi/pull/1137)
- openapi3: strip **origin** from slices in example values by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1138](https://redirect.github.com/getkin/kin-openapi/pull/1138)
- fix: bump yaml and yaml3 to v0.0.4 by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1136](https://redirect.github.com/getkin/kin-openapi/pull/1136)
- openapi3: OriginTree approach for origin tracking — separate pass, no inline stripping by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1142](https://redirect.github.com/getkin/kin-openapi/pull/1142)
- README: drop go-openapi/spec3 by [@&#8203;zonescape](https://redirect.github.com/zonescape) in [#&#8203;1143](https://redirect.github.com/getkin/kin-openapi/pull/1143)
- fix: bump yaml3+yaml to v0.0.9 to fix -root schema origin by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1144](https://redirect.github.com/getkin/kin-openapi/pull/1144)
- openapi3: call ReadFromURIFunc before checking IsExternalRefsAllowed by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1146](https://redirect.github.com/getkin/kin-openapi/pull/1146)
- fix: use location.String() instead of location.Path for origin file tracking by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [#&#8203;1145](https://redirect.github.com/getkin/kin-openapi/pull/1145)
- refactor: Replace sort usage with slices package by [@&#8203;jedevc](https://redirect.github.com/jedevc) in [#&#8203;1147](https://redirect.github.com/getkin/kin-openapi/pull/1147)

#### New Contributors

- [@&#8203;zonescape](https://redirect.github.com/zonescape) made their first contribution in [#&#8203;1143](https://redirect.github.com/getkin/kin-openapi/pull/1143)
- [@&#8203;jedevc](https://redirect.github.com/jedevc) made their first contribution in [#&#8203;1147](https://redirect.github.com/getkin/kin-openapi/pull/1147)

**Full Changelog**: <https://github.com/getkin/kin-openapi/compare/v0.134.0...v0.135.0>

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

Update OpenAPI tooling dependencies to their latest compatible versions.

Enhancements:
- Bump github.com/getkin/kin-openapi from v0.134.0 to v0.135.0.
- Update indirect dependencies github.com/oasdiff/yaml and github.com/oasdiff/yaml3 to v0.0.9.

---

## Discussion

### Comment by @red-hat-konflux on April 13, 2026 at 01:29 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 2 additional dependencies were updated


Details:


| **Package**                | **Change**           |
| :------------------------- | :------------------- |
| `github.com/oasdiff/yaml`  | `v0.0.1` -> `v0.0.9` |
| `github.com/oasdiff/yaml3` | `v0.0.1` -> `v0.0.9` |

### Comment by @sourcery-ai on April 13, 2026 at 01:30 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Bumps github.com/getkin/kin-openapi from v0.134.0 to v0.135.0 and aligns its indirect yaml dependencies accordingly.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Update kin-openapi dependency to the latest minor version and sync its yaml-related indirect dependencies. | <ul><li>Change required version of github.com/getkin/kin-openapi from v0.134.0 to v0.135.0 in go.mod</li><li>Update indirect github.com/oasdiff/yaml from v0.0.1 to v0.0.9 in go.mod</li><li>Update indirect github.com/oasdiff/yaml3 from v0.0.1 to v0.0.9 in go.mod</li><li>Refresh go.sum entries to match the new module versions</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @github-actions on April 13, 2026 at 01:30 AM UTC

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

### Comment by @github-actions on April 13, 2026 at 01:30 AM UTC

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

### Comment by @codecov-commenter on April 13, 2026 at 01:35 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2150?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.21%. Comparing base ([`689b238`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/689b238305003aae37e72d47bac6024921795aa3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`4095896`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/40958968939b28f011d46142129719d81a0f113a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2150   +/-   ##
=======================================
  Coverage   59.21%   59.21%           
=======================================
  Files         134      134           
  Lines        8639     8639           
=======================================
  Hits         5116     5116           
  Misses       2988     2988           
  Partials      535      535           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2150/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2150/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.21% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2150?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on April 13, 2026 at 01:30 AM UTC

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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2150*
