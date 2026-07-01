---
type: pull_request
number: 1712
title: "Update module github.com/go-playground/validator/v10 to v10.27.0"
state: merged
author: red-hat-konflux
created: 2025-07-06T04:53:09Z
updated: 2025-07-06T04:58:55Z
closed: 2025-07-06T04:58:53Z
merged: 2025-07-06T04:58:53Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1712
---

# Pull Request #1712: Update module github.com/go-playground/validator/v10 to v10.27.0

**Author**: @red-hat-konflux
**Created**: July 06, 2025 at 04:53 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/go-playground/validator/v10](https://redirect.github.com/go-playground/validator) | indirect | minor | `v10.26.0` -> `v10.27.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-playground/validator (github.com/go-playground/validator/v10)</summary>

### [`v10.27.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.27.0): Release 10.27.0

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.26.0...v10.27.0)

#### What's Changed

-   Fix Release version badge on README page by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1406](https://redirect.github.com/go-playground/validator/pull/1406)
-   fix russian E.164 error message by [@&#8203;prigornitskiy](https://redirect.github.com/prigornitskiy) in [https://github.com/go-playground/validator/pull/1349](https://redirect.github.com/go-playground/validator/pull/1349)
-   chore: remove unnecessary statement by [@&#8203;qshuai](https://redirect.github.com/qshuai) in [https://github.com/go-playground/validator/pull/1200](https://redirect.github.com/go-playground/validator/pull/1200)
-   Re-enable several linters by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1412](https://redirect.github.com/go-playground/validator/pull/1412)
-   add support to tag validateFn by [@&#8203;peczenyj](https://redirect.github.com/peczenyj) in [https://github.com/go-playground/validator/pull/1363](https://redirect.github.com/go-playground/validator/pull/1363)
-   Bump golang.org/x/crypto from 0.33.0 to 0.35.0 in /\_examples/validate_fn by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/go-playground/validator/pull/1418](https://redirect.github.com/go-playground/validator/pull/1418)
-   Bump golang.org/x/net from 0.34.0 to 0.38.0 in /\_examples/validate_fn by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/go-playground/validator/pull/1419](https://redirect.github.com/go-playground/validator/pull/1419)
-   Align required_without with the contract stated in the documentation by [@&#8203;jmfrees](https://redirect.github.com/jmfrees) in [https://github.com/go-playground/validator/pull/1422](https://redirect.github.com/go-playground/validator/pull/1422)
-   Add translation example by [@&#8203;cxlblm](https://redirect.github.com/cxlblm) in [https://github.com/go-playground/validator/pull/1394](https://redirect.github.com/go-playground/validator/pull/1394)
-   doc(errors): mention RegisterTagNameFunc for FieldError.Field by [@&#8203;khan-ajamal](https://redirect.github.com/khan-ajamal) in [https://github.com/go-playground/validator/pull/1358](https://redirect.github.com/go-playground/validator/pull/1358)
-   Bump golangci/golangci-lint-action from 7 to 8 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/go-playground/validator/pull/1425](https://redirect.github.com/go-playground/validator/pull/1425)
-   feat(translation): add en translation for urn_rfc2141 by [@&#8203;ryanmalesic](https://redirect.github.com/ryanmalesic) in [https://github.com/go-playground/validator/pull/1431](https://redirect.github.com/go-playground/validator/pull/1431)
-   fix: panics when private field is validated by [@&#8203;ykalchevskiy](https://redirect.github.com/ykalchevskiy) in [https://github.com/go-playground/validator/pull/1423](https://redirect.github.com/go-playground/validator/pull/1423)
-   Fix: support validation for map values with struct types by [@&#8203;JunaidIslam2105](https://redirect.github.com/JunaidIslam2105) in [https://github.com/go-playground/validator/pull/1433](https://redirect.github.com/go-playground/validator/pull/1433)
-   Omitzero does not work with slice and map bug by [@&#8203;JunaidIslam2105](https://redirect.github.com/JunaidIslam2105) in [https://github.com/go-playground/validator/pull/1436](https://redirect.github.com/go-playground/validator/pull/1436)
-   Fix: Validator panics when 'nil' is used along with required if for slices and maps by [@&#8203;JunaidIslam2105](https://redirect.github.com/JunaidIslam2105) in [https://github.com/go-playground/validator/pull/1442](https://redirect.github.com/go-playground/validator/pull/1442)
-   docs: typos by [@&#8203;eqsdxr](https://redirect.github.com/eqsdxr) in [https://github.com/go-playground/validator/pull/1440](https://redirect.github.com/go-playground/validator/pull/1440)
-   fix: make "file://" fail `url` validation by [@&#8203;bfabio](https://redirect.github.com/bfabio) in [https://github.com/go-playground/validator/pull/1444](https://redirect.github.com/go-playground/validator/pull/1444)
-   disable way too aggressive and disagreeable linters by [@&#8203;deankarn](https://redirect.github.com/deankarn) in [https://github.com/go-playground/validator/pull/1445](https://redirect.github.com/go-playground/validator/pull/1445)
-   use golangci lint file for disables by [@&#8203;deankarn](https://redirect.github.com/deankarn) in [https://github.com/go-playground/validator/pull/1447](https://redirect.github.com/go-playground/validator/pull/1447)

#### New Contributors

-   [@&#8203;prigornitskiy](https://redirect.github.com/prigornitskiy) made their first contribution in [https://github.com/go-playground/validator/pull/1349](https://redirect.github.com/go-playground/validator/pull/1349)
-   [@&#8203;qshuai](https://redirect.github.com/qshuai) made their first contribution in [https://github.com/go-playground/validator/pull/1200](https://redirect.github.com/go-playground/validator/pull/1200)
-   [@&#8203;peczenyj](https://redirect.github.com/peczenyj) made their first contribution in [https://github.com/go-playground/validator/pull/1363](https://redirect.github.com/go-playground/validator/pull/1363)
-   [@&#8203;jmfrees](https://redirect.github.com/jmfrees) made their first contribution in [https://github.com/go-playground/validator/pull/1422](https://redirect.github.com/go-playground/validator/pull/1422)
-   [@&#8203;cxlblm](https://redirect.github.com/cxlblm) made their first contribution in [https://github.com/go-playground/validator/pull/1394](https://redirect.github.com/go-playground/validator/pull/1394)
-   [@&#8203;khan-ajamal](https://redirect.github.com/khan-ajamal) made their first contribution in [https://github.com/go-playground/validator/pull/1358](https://redirect.github.com/go-playground/validator/pull/1358)
-   [@&#8203;ryanmalesic](https://redirect.github.com/ryanmalesic) made their first contribution in [https://github.com/go-playground/validator/pull/1431](https://redirect.github.com/go-playground/validator/pull/1431)
-   [@&#8203;ykalchevskiy](https://redirect.github.com/ykalchevskiy) made their first contribution in [https://github.com/go-playground/validator/pull/1423](https://redirect.github.com/go-playground/validator/pull/1423)
-   [@&#8203;JunaidIslam2105](https://redirect.github.com/JunaidIslam2105) made their first contribution in [https://github.com/go-playground/validator/pull/1433](https://redirect.github.com/go-playground/validator/pull/1433)
-   [@&#8203;eqsdxr](https://redirect.github.com/eqsdxr) made their first contribution in [https://github.com/go-playground/validator/pull/1440](https://redirect.github.com/go-playground/validator/pull/1440)
-   [@&#8203;bfabio](https://redirect.github.com/bfabio) made their first contribution in [https://github.com/go-playground/validator/pull/1444](https://redirect.github.com/go-playground/validator/pull/1444)

**Full Changelog**: https://github.com/go-playground/validator/compare/v10.26.0...v10.27.0

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

Update the validator module to v10.27.0 to bring in upstream minor enhancements, new translations, and various bug fixes, including panic prevention and improved validation behavior.

New Features:
- Add English translation for urn_rfc2141 validation.
- Support tagging custom validation functions with `validateFn`.
- Include a translation example in the validation examples.

Bug Fixes:
- Prevent panics when validating private fields and when using `required` with nil slices or maps.
- Fix `omitzero` behavior for slices and maps.
- Enable validation of map values with struct types.
- Correct Russian E.164 error message.
- Ensure `file://` URLs correctly fail URL validation.

Enhancements:
- Re-enable several linters and update lint configuration.
- Fix the release version badge in the README.
- Align `required_without` logic with the documentation and remove unnecessary statements.

Build:
- Bump golangci/golangci-lint-action from v7 to v8.

Documentation:
- Document `RegisterTagNameFunc` usage for `FieldError.Field`.

---

## Discussion

### Comment by @jira-linking on July 06, 2025 at 04:53 AM UTC

Commits missing Jira IDs:
1c98ab117b86c49611da06b7c9bf93d070c5436e


### Comment by @sourcery-ai on July 06, 2025 at 04:53 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

This PR performs a minor version bump of the go-playground/validator/v10 indirect dependency from v10.26.0 to v10.27.0 by updating module metadata and checksum files.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Update go-playground/validator/v10 module version | <ul><li>Modify require statement in go.mod to v10.27.0</li><li>Regenerate go.sum entries to match updated version</li></ul> | `go.mod`<br/>`go.sum` |

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

### Comment by @codecov-commenter on July 06, 2025 at 04:58 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1712?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 57.05%. Comparing base [(`faa5730`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/faa5730f0a5ff664e61dd1bc8496056b6f5111d6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`1c98ab1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1c98ab117b86c49611da06b7c9bf93d070c5436e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1712   +/-   ##
=======================================
  Coverage   57.05%   57.05%           
=======================================
  Files         139      139           
  Lines       10807    10807           
=======================================
  Hits         6166     6166           
  Misses       4081     4081           
  Partials      560      560           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1712/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1712/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.05% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1712?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1712*
