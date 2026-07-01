---
type: pull_request
number: 2181
title: "Update module github.com/bytedance/sonic to v1.15.1"
state: merged
author: red-hat-konflux
created: 2026-05-04T01:33:56Z
updated: 2026-05-04T05:34:58Z
closed: 2026-05-04T01:34:13Z
merged: 2026-05-04T01:34:13Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-bytedance-sonic-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2181
---

# Pull Request #2181: Update module github.com/bytedance/sonic to v1.15.1

**Author**: @red-hat-konflux
**Created**: May 04, 2026 at 01:33 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-bytedance-sonic-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/bytedance/sonic](https://redirect.github.com/bytedance/sonic) | `v1.15.0` → `v1.15.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fbytedance%2fsonic/v1.15.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fbytedance%2fsonic/v1.15.0/v1.15.1?slim=true) |

---

### Release Notes

<details>
<summary>bytedance/sonic (github.com/bytedance/sonic)</summary>

### [`v1.15.1`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.15.1)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.15.0...v1.15.1)

#### What's Changed

- refactor: improve code formatting and remove unused lookup\_small\_key … by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;903](https://redirect.github.com/bytedance/sonic/pull/903)
- ci: add macOS ARM (Apple Silicon) runners to workflows by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;911](https://redirect.github.com/bytedance/sonic/pull/911)
- chore: use go fmt format by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;913](https://redirect.github.com/bytedance/sonic/pull/913)
- fix(decoder): memory corruption when decode prefilled interface by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;914](https://redirect.github.com/bytedance/sonic/pull/914)
- fix(decoder): align jit string-tag mismatch with encoding/json by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;917](https://redirect.github.com/bytedance/sonic/pull/917)
- revert: drop integer range mismatch and related tests by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;922](https://redirect.github.com/bytedance/sonic/pull/922)
- fix(rt): align map IndirectElem semantics across Go versions by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [#&#8203;924](https://redirect.github.com/bytedance/sonic/pull/924)
- feat:(encoder) not omit zero value for `omitempty` tag by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [#&#8203;927](https://redirect.github.com/bytedance/sonic/pull/927)
- opt: unify JIT funcs in single moduledata on `Pretouch` by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [#&#8203;932](https://redirect.github.com/bytedance/sonic/pull/932)
- chore: update loader v0.5.1 by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [#&#8203;933](https://redirect.github.com/bytedance/sonic/pull/933)
- docs: update README Go 1.26 support by [@&#8203;equationzhao](https://redirect.github.com/equationzhao) in [#&#8203;931](https://redirect.github.com/bytedance/sonic/pull/931)

#### New Contributors

- [@&#8203;equationzhao](https://redirect.github.com/equationzhao) made their first contribution in [#&#8203;931](https://redirect.github.com/bytedance/sonic/pull/931)

**Full Changelog**: <https://github.com/bytedance/sonic/compare/v1.15.0...v1.15.1>

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

Update JSON serialization dependency to the latest patch release.

New Features:
- Adopt github.com/bytedance/sonic v1.15.1 for JSON handling.

Enhancements:
- Refresh indirect github.com/bytedance/sonic/loader dependency to v0.5.1 for compatibility with the updated sonic version.

---

## Discussion

### Comment by @red-hat-konflux on May 04, 2026 at 01:33 AM UTC

### ℹ️ Artifact update notice

##### File name: go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                         | **Change**           |
| :---------------------------------- | :------------------- |
| `github.com/bytedance/sonic/loader` | `v0.5.0` -> `v0.5.1` |

### Comment by @sourcery-ai on May 04, 2026 at 01:34 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

<details>
<summary>Reviewer's guide (collapsed on small PRs)</summary>

## Reviewer's Guide

Updates the github.com/bytedance/sonic dependency and its loader submodule to the latest patch version, with no code changes outside go.mod/go.sum.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Bump github.com/bytedance/sonic and its loader submodule to v1.15.1 for bug fixes and compatibility updates. | <ul><li>Update main sonic module requirement from v1.15.0 to v1.15.1 in go.mod</li><li>Update indirect sonic/loader requirement from v0.5.0 to v0.5.1 in go.mod</li><li>Regenerate go.sum entries to match the new sonic and loader versions</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on May 04, 2026 at 01:41 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2181?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.16%. Comparing base ([`afed5b4`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/afed5b4d1084176c61e7af999e479f908f4c9894?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`777fae0`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/777fae01bec5673f7b1047c459d59dd183212f2b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2181   +/-   ##
=======================================
  Coverage   59.16%   59.16%           
=======================================
  Files         136      136           
  Lines        8759     8759           
=======================================
  Hits         5182     5182           
  Misses       3032     3032           
  Partials      545      545           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2181/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2181/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2181?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @sourcery-ai - Commented on May 04, 2026 at 01:34 AM UTC

Hey - I've left some high level feedback:

- The sonic update includes a behavioral change around `omitempty` tags and decoder fixes; consider scanning any code paths that rely on sonic’s encoding/decoding semantics (especially zero-value omission and interface decoding) to confirm they still match expectations with v1.15.1.

<details>
<summary>Prompt for AI Agents</summary>

~~~markdown
Please address the comments from this code review:

## Overall Comments
- The sonic update includes a behavioral change around `omitempty` tags and decoder fixes; consider scanning any code paths that rely on sonic’s encoding/decoding semantics (especially zero-value omission and interface decoding) to confirm they still match expectations with v1.15.1.
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

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2181*
