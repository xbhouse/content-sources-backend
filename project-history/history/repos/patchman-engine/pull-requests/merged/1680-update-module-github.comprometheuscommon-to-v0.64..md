---
type: pull_request
number: 1680
title: "Update module github.com/prometheus/common to v0.64.0"
state: merged
author: red-hat-konflux
created: 2025-06-15T10:35:02Z
updated: 2025-06-15T10:40:48Z
closed: 2025-06-15T10:40:48Z
merged: 2025-06-15T10:40:48Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-common-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1680
---

# Pull Request #1680: Update module github.com/prometheus/common to v0.64.0

**Author**: @red-hat-konflux
**Created**: June 15, 2025 at 10:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | indirect | minor | `v0.63.0` -> `v0.64.0` |

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.64.0`](https://redirect.github.com/prometheus/common/releases/tag/v0.64.0)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.63.0...v0.64.0)

#### What's Changed

-   Add deprecation notice to otlptranslator by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [https://github.com/prometheus/common/pull/773](https://redirect.github.com/prometheus/common/pull/773)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/common/pull/774](https://redirect.github.com/prometheus/common/pull/774)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/common/pull/775](https://redirect.github.com/prometheus/common/pull/775)
-   Update Go by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [https://github.com/prometheus/common/pull/770](https://redirect.github.com/prometheus/common/pull/770)
-   chore: Upgrade golangci-lint to v2 by [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) in [https://github.com/prometheus/common/pull/779](https://redirect.github.com/prometheus/common/pull/779)
-   build(deps): bump golang.org/x/net from 0.37.0 to 0.38.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/777](https://redirect.github.com/prometheus/common/pull/777)
-   build(deps): bump google.golang.org/protobuf from 1.36.5 to 1.36.6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/776](https://redirect.github.com/prometheus/common/pull/776)
-   promslog: Use the default timezone (again) by [@&#8203;beorn7](https://redirect.github.com/beorn7) in [https://github.com/prometheus/common/pull/739](https://redirect.github.com/prometheus/common/pull/739)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/common/pull/787](https://redirect.github.com/prometheus/common/pull/787)
-   build(deps): bump github.com/prometheus/client_model from 0.6.1 to 0.6.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/784](https://redirect.github.com/prometheus/common/pull/784)
-   build(deps): bump golang.org/x/oauth2 from 0.28.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/785](https://redirect.github.com/prometheus/common/pull/785)
-   build(deps): bump golang.org/x/net from 0.38.0 to 0.39.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/786](https://redirect.github.com/prometheus/common/pull/786)
-   refactor(promslog): make `NewNopLogger()` wrapper around `New()` by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [https://github.com/prometheus/common/pull/783](https://redirect.github.com/prometheus/common/pull/783)
-   build(deps): bump golang.org/x/oauth2 from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/common/pull/788](https://redirect.github.com/prometheus/common/pull/788)

#### New Contributors

-   [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) made their first contribution in [https://github.com/prometheus/common/pull/779](https://redirect.github.com/prometheus/common/pull/779)

**Full Changelog**: https://github.com/prometheus/common/compare/v0.63.0...v0.64.0

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4yNjQuMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4yNjQuMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->

## Summary by Sourcery

Build:
- Bump github.com/prometheus/common from v0.63.0 to v0.64.0 in go.mod

---

## Discussion

### Comment by @jira-linking on June 15, 2025 at 10:35 AM UTC

Commits missing Jira IDs:
b4992aaec972415988e5cde5bbcc5f49986b7c4d


### Comment by @sourcery-ai on June 15, 2025 at 10:35 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

Upgrade the indirect dependency on github.com/prometheus/common from v0.63.0 to v0.64.0 by updating the version in go.mod and regenerating go.sum to align checksums.

#### Class Diagram: Key Changes in github.com/prometheus/common v0.64.0

```mermaid
classDiagram
  namespace prometheus_common_v0_64_0 {
    class otlptranslator {
      <<deprecated>>
      # Specific methods/attributes are not detailed in PR summary
      # This component is deprecated in v0.64.0
    }
  }

  class promslog {
      +New(...any) Logger
      +NewNopLogger() Logger
      # NewNopLogger() was refactored internally in v0.64.0
      # promslog also reverted to using default timezone for logging
  }
```

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Prometheus/common module version bumped | <ul><li>Modify the version entry for github.com/prometheus/common in go.mod</li><li>Run `go mod tidy` to update go.sum with new checksums</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on June 15, 2025 at 10:40 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1680?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 57.29%. Comparing base [(`25b6cc2`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/25b6cc296f0db1522ad5af313d6ee9d39fb257bd?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`b4992aa`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b4992aaec972415988e5cde5bbcc5f49986b7c4d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1680   +/-   ##
=======================================
  Coverage   57.29%   57.29%           
=======================================
  Files         138      138           
  Lines       10790    10790           
=======================================
  Hits         6182     6182           
  Misses       4048     4048           
  Partials      560      560           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1680/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1680/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.29% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1680?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1680*
