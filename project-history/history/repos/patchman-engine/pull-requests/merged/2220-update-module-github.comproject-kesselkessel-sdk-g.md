---
type: pull_request
number: 2220
title: "Update module github.com/project-kessel/kessel-sdk-go to v1.9.1"
state: merged
author: red-hat-konflux
created: 2026-06-01T02:34:43Z
updated: 2026-06-01T05:13:50Z
closed: 2026-06-01T02:34:57Z
merged: 2026-06-01T02:34:57Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-project-kessel-kessel-sdk-go-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2220
---

# Pull Request #2220: Update module github.com/project-kessel/kessel-sdk-go to v1.9.1

**Author**: @red-hat-konflux
**Created**: June 01, 2026 at 02:34 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-project-kessel-kessel-sdk-go-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/project-kessel/kessel-sdk-go](https://redirect.github.com/project-kessel/kessel-sdk-go) | `v1.9.0` → `v1.9.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fproject-kessel%2fkessel-sdk-go/v1.9.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fproject-kessel%2fkessel-sdk-go/v1.9.0/v1.9.1?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>project-kessel/kessel-sdk-go (github.com/project-kessel/kessel-sdk-go)</summary>

### [`v1.9.1`](https://redirect.github.com/project-kessel/kessel-sdk-go/releases/tag/v1.9.1)

[Compare Source](https://redirect.github.com/project-kessel/kessel-sdk-go/compare/v1.9.0...v1.9.1)

#### What's Changed

- Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.11-20260209202127-80ab13bee0bf.1 to 1.36.11-20260415201107-50325440f8f2.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;52](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/52)
- Bump github.com/zitadel/oidc/v3 from 3.45.5 to 3.47.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;56](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/56)
- Bump google.golang.org/grpc from 1.79.2 to 1.81.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;60](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/60)
- RHCLOUD-47739 - Update locking mechanism for concurrent token refreshes by [@&#8203;lennysgarage](https://redirect.github.com/lennysgarage) in [#&#8203;61](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/61)

**Full Changelog**: <https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.0...v1.9.1>

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
- Bump github.com/project-kessel/kessel-sdk-go from v1.9.0 to v1.9.1 along with associated indirect dependency updates for protovalidate and OIDC.

---

## Discussion

### Comment by @red-hat-konflux on June 01, 2026 at 02:34 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 2 additional dependencies were updated


Details:


| **Package**                                                  | **Change**                                                                           |
| :----------------------------------------------------------- | :----------------------------------------------------------------------------------- |
| `buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go` | `v1.36.11-20260209202127-80ab13bee0bf.1` -> `v1.36.11-20260415201107-50325440f8f2.1` |
| `github.com/zitadel/oidc/v3`                                 | `v3.45.5` -> `v3.47.5`                                                               |

### Comment by @sourcery-ai on June 01, 2026 at 02:34 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Bumps the kessel SDK dependency to v1.9.1 and aligns related indirect dependencies to the versions used by the SDK, with no application code changes.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Update direct dependency on kessel SDK to v1.9.1 and sync indirect dependencies | <ul><li>Bump github.com/project-kessel/kessel-sdk-go from v1.9.0 to v1.9.1 in go.mod</li><li>Update buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go indirect dependency to the newer generated version</li><li>Update github.com/zitadel/oidc/v3 indirect dependency to v3.47.5 to match upstream requirements</li><li>Regenerate go.sum entries to reflect the new dependency versions</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on June 01, 2026 at 02:40 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2220?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.03%. Comparing base ([`e988cde`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/e988cdefbf26896eac51eb34055306e4c3e02e2f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`10dccbf`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/10dccbf0fb50eb560b53f262d4d914e9899e7aa4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2220   +/-   ##
=======================================
  Coverage   59.03%   59.03%           
=======================================
  Files         137      137           
  Lines        8831     8831           
=======================================
  Hits         5213     5213           
  Misses       3070     3070           
  Partials      548      548           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2220/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2220/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.03% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2220?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on June 01, 2026 at 02:36 AM UTC

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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2220*
