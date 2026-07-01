---
type: pull_request
number: 1258
title: "Build chore(deps): update dependency eslint-config-prettier to v10"
state: merged
author: red-hat-konflux
created: 2025-10-27T08:15:45Z
updated: 2025-10-27T12:14:13Z
closed: 2025-10-27T09:26:57Z
merged: 2025-10-27T09:26:57Z
base_branch: main
head_branch: konflux/mintmaker/main/eslint-config-prettier-10.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1258
---

# Pull Request #1258: Build chore(deps): update dependency eslint-config-prettier to v10

**Author**: @red-hat-konflux
**Created**: October 27, 2025 at 08:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/eslint-config-prettier-10.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [eslint-config-prettier](https://redirect.github.com/prettier/eslint-config-prettier) | [`^9.1.0` -> `^10.0.0`](https://renovatebot.com/diffs/npm/eslint-config-prettier/9.1.2/10.1.8) | [![age](https://developer.mend.io/api/mc/badges/age/npm/eslint-config-prettier/10.1.8?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/eslint-config-prettier/9.1.2/10.1.8?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>prettier/eslint-config-prettier (eslint-config-prettier)</summary>

### [`v10.1.8`](https://redirect.github.com/prettier/eslint-config-prettier/releases/tag/v10.1.8)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.1.5...v10.1.8)

republish latest version

**Full Changelog**: <https://github.com/prettier/eslint-config-prettier/compare/v10.1.5...v10.1.8>

### [`v10.1.5`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1015)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.1.4...v10.1.5)

##### Patch Changes

- [#&#8203;332](https://redirect.github.com/prettier/eslint-config-prettier/pull/332) [`60fef02`](https://redirect.github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4) Thanks [@&#8203;JounQin](https://redirect.github.com/JounQin)! - chore: add `funding` field into `package.json`

### [`v10.1.4`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1014)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.1.3...v10.1.4)

##### Patch Changes

- [#&#8203;328](https://redirect.github.com/prettier/eslint-config-prettier/pull/328) [`94b4799`](https://redirect.github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850) Thanks [@&#8203;silvenon](https://redirect.github.com/silvenon)! - fix(cli): do not crash on no rules configured

### [`v10.1.3`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1013)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.1.2...v10.1.3)

##### Patch Changes

- [#&#8203;325](https://redirect.github.com/prettier/eslint-config-prettier/pull/325) [`4e95a1d`](https://redirect.github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df) Thanks [@&#8203;pilikan](https://redirect.github.com/pilikan)! - fix: this package is `commonjs`, align its types correctly

### [`v10.1.2`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1012)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.1.1...v10.1.2)

##### Patch Changes

- [#&#8203;321](https://redirect.github.com/prettier/eslint-config-prettier/pull/321) [`a8768bf`](https://redirect.github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0) Thanks [@&#8203;Fdawgs](https://redirect.github.com/Fdawgs)! - chore(package): add homepage for some 3rd-party registry - see [#&#8203;321](https://redirect.github.com/prettier/eslint-config-prettier/pull/321) for more details

### [`v10.1.1`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1011)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.1.0...v10.1.1)

##### Patch Changes

- [#&#8203;309](https://redirect.github.com/prettier/eslint-config-prettier/pull/309) [`eb56a5e`](https://redirect.github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d) Thanks [@&#8203;JounQin](https://redirect.github.com/JounQin)! - fix: separate the `/flat` entry for compatibility

  For flat config users, the previous `"eslint-config-prettier"` entry still works, but `"eslint-config-prettier/flat"` adds a new `name` property for [config-inspector](https://eslint.org/blog/2024/04/eslint-config-inspector/), we just can't add it for the default entry for compatibility.

  See also [#&#8203;308](https://redirect.github.com/prettier/eslint-config-prettier/issues/308)

  ```ts
  // before
  import eslintConfigPrettier from "eslint-config-prettier";

  // after
  import eslintConfigPrettier from "eslint-config-prettier/flat";
  ```

### [`v10.1.0`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1010)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.0.3...v10.1.0)

##### Minor Changes

- [#&#8203;306](https://redirect.github.com/prettier/eslint-config-prettier/pull/306) [`56e2e34`](https://redirect.github.com/prettier/eslint-config-prettier/commit/56e2e3466391d0fdfc200e42130309c687aaab53) Thanks [@&#8203;JounQin](https://redirect.github.com/JounQin)! - feat: migrate to exports field

### [`v10.0.3`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1003)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.0.2...v10.0.3)

##### Patch Changes

- [#&#8203;294](https://redirect.github.com/prettier/eslint-config-prettier/pull/294) [`8dbbd6d`](https://redirect.github.com/prettier/eslint-config-prettier/commit/8dbbd6d70b8a56cdfa4ea4e185d3699d5729b38e) Thanks [@&#8203;FloEdelmann](https://redirect.github.com/FloEdelmann)! - feat: add name to config

- [#&#8203;280](https://redirect.github.com/prettier/eslint-config-prettier/pull/280) [`cba5737`](https://redirect.github.com/prettier/eslint-config-prettier/commit/cba57377e4c86d20d17042d6999eabc754fddc03) Thanks [@&#8203;zanminkian](https://redirect.github.com/zanminkian)! - feat: add declaration file

### [`v10.0.2`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1002)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.0.1...v10.0.2)

##### Patch Changes

- [#&#8203;299](https://redirect.github.com/prettier/eslint-config-prettier/pull/299) [`e750edc`](https://redirect.github.com/prettier/eslint-config-prettier/commit/e750edc530c816e0b3ffabfab1f4e46532bccbfe) Thanks [@&#8203;Fdawgs](https://redirect.github.com/Fdawgs)! - chore(package): explicitly declare js module type

### [`v10.0.1`](https://redirect.github.com/prettier/eslint-config-prettier/releases/tag/v10.0.1)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/v10.0.0...v10.0.1)

### eslint-config-prettier

#### 10.0.1

#### What's Changed

- chore: migrate to changeset for automatically releasing by [@&#8203;JounQin](https://redirect.github.com/JounQin) in [#&#8203;278](https://redirect.github.com/prettier/eslint-config-prettier/pull/278)
- add support for `@stylistic/eslint-plugin` by [@&#8203;abrahamguo](https://redirect.github.com/abrahamguo) in [#&#8203;272](https://redirect.github.com/prettier/eslint-config-prettier/pull/272)

#### New Contributors

- [@&#8203;JounQin](https://redirect.github.com/JounQin) made their first contribution in [#&#8203;278](https://redirect.github.com/prettier/eslint-config-prettier/pull/278)
- [@&#8203;abrahamguo](https://redirect.github.com/abrahamguo) made their first contribution in [#&#8203;272](https://redirect.github.com/prettier/eslint-config-prettier/pull/272)

**Full Changelog**: <https://github.com/prettier/eslint-config-prettier/compare/v9.1.0...v10.0.1>

### [`v10.0.0`](https://redirect.github.com/prettier/eslint-config-prettier/blob/HEAD/CHANGELOG.md#1000)

[Compare Source](https://redirect.github.com/prettier/eslint-config-prettier/compare/62a15bbfc953d798da2cc8a46e78897da0beac22...v10.0.0)

##### Major Changes

- [#&#8203;272](https://redirect.github.com/prettier/eslint-config-prettier/pull/272) [`5be64be`](https://redirect.github.com/prettier/eslint-config-prettier/commit/5be64bef68c3a9bf7202f591f54ffec02572e46b) Thanks [@&#8203;abrahamguo](https://redirect.github.com/abrahamguo)! - add support for [@&#8203;stylistic](https://redirect.github.com/stylistic) formatting rules

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "on Monday after 3am and before 10am" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @swadeley - Approved on October 27, 2025 at 09:26 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1258*
