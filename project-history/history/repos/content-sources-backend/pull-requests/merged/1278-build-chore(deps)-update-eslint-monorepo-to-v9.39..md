---
type: pull_request
number: 1278
title: "Build chore(deps): update eslint monorepo to v9.39.0"
state: merged
author: red-hat-konflux
created: 2025-11-03T08:19:47Z
updated: 2025-11-03T12:17:33Z
closed: 2025-11-03T10:01:44Z
merged: 2025-11-03T10:01:44Z
base_branch: main
head_branch: konflux/mintmaker/main/eslint-monorepo
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1278
---

# Pull Request #1278: Build chore(deps): update eslint monorepo to v9.39.0

**Author**: @red-hat-konflux
**Created**: November 03, 2025 at 08:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/eslint-monorepo`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [@eslint/js](https://eslint.org) ([source](https://redirect.github.com/eslint/eslint/tree/HEAD/packages/js)) | [`9.38.0` -> `9.39.0`](https://renovatebot.com/diffs/npm/@eslint%2fjs/9.38.0/9.39.0) | [![age](https://developer.mend.io/api/mc/badges/age/npm/@eslint%2fjs/9.39.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/@eslint%2fjs/9.38.0/9.39.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |
| [eslint](https://eslint.org) ([source](https://redirect.github.com/eslint/eslint)) | [`9.38.0` -> `9.39.0`](https://renovatebot.com/diffs/npm/eslint/9.38.0/9.39.0) | [![age](https://developer.mend.io/api/mc/badges/age/npm/eslint/9.39.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/eslint/9.38.0/9.39.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>eslint/eslint (@&#8203;eslint/js)</summary>

### [`v9.39.0`](https://redirect.github.com/eslint/eslint/releases/tag/v9.39.0)

[Compare Source](https://redirect.github.com/eslint/eslint/compare/v9.38.0...v9.39.0)

#### Features

- [`cc57d87`](https://redirect.github.com/eslint/eslint/commit/cc57d87a3f119e9d39c55e044e526ae067fa31ce) feat: update error loc to key in `no-dupe-class-members` ([#&#8203;20259](https://redirect.github.com/eslint/eslint/issues/20259)) (Tanuj Kanti)
- [`126552f`](https://redirect.github.com/eslint/eslint/commit/126552fcf35da3ddcefa527db06dabc54c04041c) feat: update error location in `for-direction` and `no-dupe-args` ([#&#8203;20258](https://redirect.github.com/eslint/eslint/issues/20258)) (Tanuj Kanti)
- [`167d097`](https://redirect.github.com/eslint/eslint/commit/167d0970d3802a66910e9820f31dcd717fab0b2a) feat: update `complexity` rule to highlight only static block header ([#&#8203;20245](https://redirect.github.com/eslint/eslint/issues/20245)) (jaymarvelz)

#### Bug Fixes

- [`15f5c7c`](https://redirect.github.com/eslint/eslint/commit/15f5c7c168d0698683943f51dd617f14a5e6815c) fix: forward traversal `step.args` to visitors ([#&#8203;20253](https://redirect.github.com/eslint/eslint/issues/20253)) (jaymarvelz)
- [`5a1a534`](https://redirect.github.com/eslint/eslint/commit/5a1a534e877f7c4c992885867f923df307c3929d) fix: allow JSDoc comments in object-shorthand rule ([#&#8203;20167](https://redirect.github.com/eslint/eslint/issues/20167)) (Nitin Kumar)
- [`e86b813`](https://redirect.github.com/eslint/eslint/commit/e86b813eb660f1a5adc8e143a70d9b683cd12362) fix: Use more types from [@&#8203;eslint/core](https://redirect.github.com/eslint/core) ([#&#8203;20257](https://redirect.github.com/eslint/eslint/issues/20257)) (Nicholas C. Zakas)
- [`927272d`](https://redirect.github.com/eslint/eslint/commit/927272d1f0d5683b029b729d368a96527f283323) fix: correct `Scope` typings ([#&#8203;20198](https://redirect.github.com/eslint/eslint/issues/20198)) (jaymarvelz)
- [`37f76d9`](https://redirect.github.com/eslint/eslint/commit/37f76d9c539bb6fc816fedb7be4486b71a58620a) fix: use `AST.Program` type for Program node ([#&#8203;20244](https://redirect.github.com/eslint/eslint/issues/20244)) (Francesco Trotta)
- [`ae07f0b`](https://redirect.github.com/eslint/eslint/commit/ae07f0b3334ebd22ae2e7b09bca5973b96aa9768) fix: unify timing report for concurrent linting ([#&#8203;20188](https://redirect.github.com/eslint/eslint/issues/20188)) (jaymarvelz)
- [`b165d47`](https://redirect.github.com/eslint/eslint/commit/b165d471be6062f4475b972155b02654a974a0e9) fix: correct `Rule` typings ([#&#8203;20199](https://redirect.github.com/eslint/eslint/issues/20199)) (jaymarvelz)
- [`fb97cda`](https://redirect.github.com/eslint/eslint/commit/fb97cda70d87286a7dbd2457f578ef578d6905e8) fix: improve error message for missing fix function in suggestions ([#&#8203;20218](https://redirect.github.com/eslint/eslint/issues/20218)) (jaymarvelz)

#### Documentation

- [`d3e81e3`](https://redirect.github.com/eslint/eslint/commit/d3e81e30ee6be5a21151b7a17ef10a714b6059c0) docs: Always recommend to include a files property ([#&#8203;20158](https://redirect.github.com/eslint/eslint/issues/20158)) (Percy Ma)
- [`0f0385f`](https://redirect.github.com/eslint/eslint/commit/0f0385f1404dcadaba4812120b1ad02334dbd66a) docs: use consistent naming recommendation ([#&#8203;20250](https://redirect.github.com/eslint/eslint/issues/20250)) (Alex M. Spieslechner)
- [`a3b1456`](https://redirect.github.com/eslint/eslint/commit/a3b145609ac649fac837c8c0515cbb2a9321ca40) docs: Update README (GitHub Actions Bot)
- [`cf5f2dd`](https://redirect.github.com/eslint/eslint/commit/cf5f2dd58dd98084a21da04fe7b9054b9478d552) docs: fix correct tag of `no-useless-constructor` ([#&#8203;20255](https://redirect.github.com/eslint/eslint/issues/20255)) (Tanuj Kanti)
- [`10b995c`](https://redirect.github.com/eslint/eslint/commit/10b995c8e5473de8d66d3cd99d816e046f35e3ec) docs: add TS options and examples for `nofunc` in `no-use-before-define` ([#&#8203;20249](https://redirect.github.com/eslint/eslint/issues/20249)) (Tanuj Kanti)
- [`2584187`](https://redirect.github.com/eslint/eslint/commit/2584187e4a305ea7a98e1a5bd4dca2a60ad132f8) docs: remove repetitive word in comment ([#&#8203;20242](https://redirect.github.com/eslint/eslint/issues/20242)) (reddaisyy)
- [`637216b`](https://redirect.github.com/eslint/eslint/commit/637216bd4f2aae7c928ad04a4e40eecffb50c9e5) docs: update CLI flags migration instructions ([#&#8203;20238](https://redirect.github.com/eslint/eslint/issues/20238)) (jaymarvelz)
- [`e7cda3b`](https://redirect.github.com/eslint/eslint/commit/e7cda3bdf1bdd664e6033503a3315ad81736b200) docs: Update README (GitHub Actions Bot)
- [`7b9446f`](https://redirect.github.com/eslint/eslint/commit/7b9446f7cc2054aa2cdf8e6225f4ac15a03671a8) docs: handle empty flags sections on the feature flags page ([#&#8203;20222](https://redirect.github.com/eslint/eslint/issues/20222)) (sethamus)

#### Chores

- [`dfe3c1b`](https://redirect.github.com/eslint/eslint/commit/dfe3c1b2034228765c48c8a445554223767dd16d) chore: update `@eslint/js` version to 9.39.0 ([#&#8203;20270](https://redirect.github.com/eslint/eslint/issues/20270)) (Francesco Trotta)
- [`2375a6d`](https://redirect.github.com/eslint/eslint/commit/2375a6de8263393c129d41cac1b407b40111a73c) chore: package.json update for [@&#8203;eslint/js](https://redirect.github.com/eslint/js) release (Jenkins)
- [`a1f4e52`](https://redirect.github.com/eslint/eslint/commit/a1f4e52d67c94bef61edd1607dcd130047c1baf0) chore: update `@eslint` dependencies ([#&#8203;20265](https://redirect.github.com/eslint/eslint/issues/20265)) (Francesco Trotta)
- [`c7d3229`](https://redirect.github.com/eslint/eslint/commit/c7d32298482752eeac9fb46378d4f1ea095f3836) chore: update dependency [@&#8203;eslint/core](https://redirect.github.com/eslint/core) to ^0.17.0 ([#&#8203;20256](https://redirect.github.com/eslint/eslint/issues/20256)) (renovate\[bot])
- [`27549bc`](https://redirect.github.com/eslint/eslint/commit/27549bc774c7c2dc5c569070a3e87c62f602bf7d) chore: update fuzz testing to not error if code sample minimizer fails ([#&#8203;20252](https://redirect.github.com/eslint/eslint/issues/20252)) (Milos Djermanovic)
- [`a1370ee`](https://redirect.github.com/eslint/eslint/commit/a1370ee40e9d8e0e41843f3278cd745fc1ad543f) ci: bump actions/setup-node from 5 to 6 ([#&#8203;20230](https://redirect.github.com/eslint/eslint/issues/20230)) (dependabot\[bot])
- [`9e7fad4`](https://redirect.github.com/eslint/eslint/commit/9e7fad4a1867709060686d03e0ec1d0d69671cfb) chore: add script to auto-generate eslint:recommended configuration ([#&#8203;20208](https://redirect.github.com/eslint/eslint/issues/20208)) (唯然)

</details>

<details>
<summary>eslint/eslint (eslint)</summary>

### [`v9.39.0`](https://redirect.github.com/eslint/eslint/compare/v9.38.0...ac3a60dffc29d8d4d5031621bc062e77f891532a)

[Compare Source](https://redirect.github.com/eslint/eslint/compare/v9.38.0...v9.39.0)

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "on Monday after 3am and before 10am" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about these updates again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @swadeley - Approved on November 03, 2025 at 10:01 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1278*
