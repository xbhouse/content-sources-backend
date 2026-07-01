---
type: pull_request
number: 1740
title: "Update module github.com/swaggo/swag to v1.16.5"
state: merged
author: red-hat-konflux
created: 2025-07-20T06:53:10Z
updated: 2025-07-20T06:58:50Z
closed: 2025-07-20T06:58:48Z
merged: 2025-07-20T06:58:48Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-swaggo-swag-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1740
---

# Pull Request #1740: Update module github.com/swaggo/swag to v1.16.5

**Author**: @red-hat-konflux
**Created**: July 20, 2025 at 06:53 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-swaggo-swag-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/swaggo/swag](https://redirect.github.com/swaggo/swag) | indirect | patch | `v1.16.4` -> `v1.16.5` |

---

### Release Notes

<details>
<summary>swaggo/swag (github.com/swaggo/swag)</summary>

### [`v1.16.5`](https://redirect.github.com/swaggo/swag/releases/tag/v1.16.5)

[Compare Source](https://redirect.github.com/swaggo/swag/compare/v1.16.4...v1.16.5)

#### What's Changed

-   Added support for [@&#8203;tag](https://redirect.github.com/tag).x- attributes for tags ([#&#8203;1784](https://redirect.github.com/swaggo/swag/issues/1784)) by [@&#8203;Ponywka](https://redirect.github.com/Ponywka) in [https://github.com/swaggo/swag/pull/1785](https://redirect.github.com/swaggo/swag/pull/1785)
-   feat: Add x-enum-descriptions to generated Swagger documentation for Enum by [@&#8203;wakamenod](https://redirect.github.com/wakamenod) in [https://github.com/swaggo/swag/pull/1878](https://redirect.github.com/swaggo/swag/pull/1878)
-   fix: use '&&' for security pair(AND) by [@&#8203;kkkiio](https://redirect.github.com/kkkiio) in [https://github.com/swaggo/swag/pull/1659](https://redirect.github.com/swaggo/swag/pull/1659)
-   feat: ParseComment error to contain the comment by [@&#8203;stokito](https://redirect.github.com/stokito) in [https://github.com/swaggo/swag/pull/1777](https://redirect.github.com/swaggo/swag/pull/1777)
-   support generate var-declared function doc by [@&#8203;book987](https://redirect.github.com/book987) in [https://github.com/swaggo/swag/pull/1657](https://redirect.github.com/swaggo/swag/pull/1657)
-   Fix compare original and formatted by [@&#8203;0daryo](https://redirect.github.com/0daryo) in [https://github.com/swaggo/swag/pull/1915](https://redirect.github.com/swaggo/swag/pull/1915)
-   Transfer golang type to swagger type with format reserved by [@&#8203;sdghchj](https://redirect.github.com/sdghchj) in [https://github.com/swaggo/swag/pull/1944](https://redirect.github.com/swaggo/swag/pull/1944)
-   Fix format by goimports by [@&#8203;0daryo](https://redirect.github.com/0daryo) in [https://github.com/swaggo/swag/pull/1927](https://redirect.github.com/swaggo/swag/pull/1927)
-   chore(deps): bump golang.org/x/crypto from 0.21.0 to 0.31.0 in /example/object-map-example by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in[https://github.com/swaggo/swag/pull/1945](https://redirect.github.com/swaggo/swag/pull/1945)5
-   fix: [@&#8203;name](https://redirect.github.com/name) for recursion by [@&#8203;njacob1001](https://redirect.github.com/njacob1001) in [https://github.com/swaggo/swag/pull/1948](https://redirect.github.com/swaggo/swag/pull/1948)
-   Fix typo error in README by [@&#8203;yashisrani](https://redirect.github.com/yashisrani) in [https://github.com/swaggo/swag/pull/1954](https://redirect.github.com/swaggo/swag/pull/1954)
-   Extension: collectionFormat in struct tag by [@&#8203;sdghchj](https://redirect.github.com/sdghchj) in [https://github.com/swaggo/swag/pull/1989](https://redirect.github.com/swaggo/swag/pull/1989)
-   chore: accept event streaming by [@&#8203;miguelhrocha](https://redirect.github.com/miguelhrocha) in [https://github.com/swaggo/swag/pull/1992](https://redirect.github.com/swaggo/swag/pull/1992)
-   Non-empty stderr for `go list` is not an error in itself by [@&#8203;atercattus](https://redirect.github.com/atercattus) in [https://github.com/swaggo/swag/pull/1981](https://redirect.github.com/swaggo/swag/pull/1981)
-   Updated golang/x/text v0.21, x/tools v0.21 to address CVE-2024-45338 in net v0.23 by [@&#8203;blame19](https://redirect.github.com/blame19) in [https://github.com/swaggo/swag/pull/1962](https://redirect.github.com/swaggo/swag/pull/1962)
-   Remove redundant `(default: false)` for parseFuncBody flag by [@&#8203;nikpivkin](https://redirect.github.com/nikpivkin) in [https://github.com/swaggo/swag/pull/1961](https://redirect.github.com/swaggo/swag/pull/1961)
-   chore(deps): bump golang.org/x/crypto from 0.21.0 to 0.31.0 in /example/celler by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in[https://github.com/swaggo/swag/pull/1993](https://redirect.github.com/swaggo/swag/pull/1993)3
-   fix(test): include `rune` test in enums tests by [@&#8203;sharunkumar](https://redirect.github.com/sharunkumar) in [https://github.com/swaggo/swag/pull/2025](https://redirect.github.com/swaggo/swag/pull/2025)
-   feat: json:omitempty marks field as optional by [@&#8203;andyatmiami](https://redirect.github.com/andyatmiami) in [https://github.com/swaggo/swag/pull/2041](https://redirect.github.com/swaggo/swag/pull/2041)

#### New Contributors

-   [@&#8203;wakamenod](https://redirect.github.com/wakamenod) made their first contribution in [https://github.com/swaggo/swag/pull/1878](https://redirect.github.com/swaggo/swag/pull/1878)
-   [@&#8203;stokito](https://redirect.github.com/stokito) made their first contribution in [https://github.com/swaggo/swag/pull/1777](https://redirect.github.com/swaggo/swag/pull/1777)
-   [@&#8203;book987](https://redirect.github.com/book987) made their first contribution in [https://github.com/swaggo/swag/pull/1657](https://redirect.github.com/swaggo/swag/pull/1657)
-   [@&#8203;njacob1001](https://redirect.github.com/njacob1001) made their first contribution in [https://github.com/swaggo/swag/pull/1948](https://redirect.github.com/swaggo/swag/pull/1948)
-   [@&#8203;yashisrani](https://redirect.github.com/yashisrani) made their first contribution in [https://github.com/swaggo/swag/pull/1954](https://redirect.github.com/swaggo/swag/pull/1954)
-   [@&#8203;miguelhrocha](https://redirect.github.com/miguelhrocha) made their first contribution in [https://github.com/swaggo/swag/pull/1992](https://redirect.github.com/swaggo/swag/pull/1992)
-   [@&#8203;atercattus](https://redirect.github.com/atercattus) made their first contribution in [https://github.com/swaggo/swag/pull/1981](https://redirect.github.com/swaggo/swag/pull/1981)
-   [@&#8203;blame19](https://redirect.github.com/blame19) made their first contribution in [https://github.com/swaggo/swag/pull/1962](https://redirect.github.com/swaggo/swag/pull/1962)
-   [@&#8203;sharunkumar](https://redirect.github.com/sharunkumar) made their first contribution in [https://github.com/swaggo/swag/pull/2025](https://redirect.github.com/swaggo/swag/pull/2025)
-   [@&#8203;andyatmiami](https://redirect.github.com/andyatmiami) made their first contribution in [https://github.com/swaggo/swag/pull/2041](https://redirect.github.com/swaggo/swag/pull/2041)

**Full Changelog**: https://github.com/swaggo/swag/compare/v1.16.4...v1.16.5

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

Bump github.com/swaggo/swag to v1.16.5 and add golang.org/x/mod v0.25.0 as an indirect dependency

Build:
- Update github.com/swaggo/swag from v1.16.4 to v1.16.5
- Add golang.org/x/mod v0.25.0 as an indirect dependency

---

## Discussion

### Comment by @jira-linking on July 20, 2025 at 06:53 AM UTC

Commits missing Jira IDs:
1d00b99801e6f7e77a205f14a83527ad238e76f3


### Comment by @sourcery-ai on July 20, 2025 at 06:53 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

This PR bumps the indirect dependency github.com/swaggo/swag from v1.16.4 to v1.16.5, introduces a new indirect module golang.org/x/mod, and regenerates go.sum to sync checksums.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Upgrade swaggo/swag to v1.16.5 | <ul><li>Update version in go.mod from v1.16.4 to v1.16.5</li></ul> | `go.mod` |
| Synchronize indirect dependencies | <ul><li>Add golang.org/x/mod v0.25.0 as an indirect requirement</li><li>Regenerate go.sum to reflect updated module checksums</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on July 20, 2025 at 06:58 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1740?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 54.77%. Comparing base [(`12f612f`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/12f612f26f6fc0a811e0ee946bae32fecff337c6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`1d00b99`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1d00b99801e6f7e77a205f14a83527ad238e76f3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 7 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1740   +/-   ##
=======================================
  Coverage   54.77%   54.77%           
=======================================
  Files         139      139           
  Lines       10752    10752           
=======================================
  Hits         5889     5889           
  Misses       4333     4333           
  Partials      530      530           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1740/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1740/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.77% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1740?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1740*
