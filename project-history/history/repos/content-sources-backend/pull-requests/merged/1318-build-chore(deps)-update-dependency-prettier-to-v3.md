---
type: pull_request
number: 1318
title: "Build: chore(deps): update dependency prettier to v3.7.3"
state: merged
author: red-hat-konflux
created: 2025-12-01T04:39:21Z
updated: 2025-12-01T10:46:11Z
closed: 2025-12-01T10:46:09Z
merged: 2025-12-01T10:46:09Z
base_branch: main
head_branch: konflux/mintmaker/main/prettier-3.x-lockfile
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1318
---

# Pull Request #1318: Build: chore(deps): update dependency prettier to v3.7.3

**Author**: @red-hat-konflux
**Created**: December 01, 2025 at 04:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/prettier-3.x-lockfile`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [prettier](https://prettier.io) ([source](https://redirect.github.com/prettier/prettier)) | [`3.7.1` -> `3.7.3`](https://renovatebot.com/diffs/npm/prettier/3.7.1/3.7.3) | [![age](https://developer.mend.io/api/mc/badges/age/npm/prettier/3.7.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/prettier/3.7.1/3.7.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>prettier/prettier (prettier)</summary>

### [`v3.7.3`](https://redirect.github.com/prettier/prettier/blob/HEAD/CHANGELOG.md#373)

[Compare Source](https://redirect.github.com/prettier/prettier/compare/3.7.2...3.7.3)

[diff](https://redirect.github.com/prettier/prettier/compare/3.7.2...3.7.3)

##### API: Fix `prettier.getFileInfo()` change that breaks VSCode extension ([#&#8203;18375](https://redirect.github.com/prettier/prettier/pull/18375) by [@&#8203;fisker](https://redirect.github.com/fisker))

An internal refactor accidentally broke the VSCode extension plugin loading.

### [`v3.7.2`](https://redirect.github.com/prettier/prettier/blob/HEAD/CHANGELOG.md#372)

[Compare Source](https://redirect.github.com/prettier/prettier/compare/3.7.1...3.7.2)

[diff](https://redirect.github.com/prettier/prettier/compare/3.7.1...3.7.2)

##### JavaScript: Fix string print when switching quotes ([#&#8203;18351](https://redirect.github.com/prettier/prettier/pull/18351) by [@&#8203;fisker](https://redirect.github.com/fisker))

<!-- prettier-ignore -->

```jsx
// Input
console.log("A descriptor\\'s .kind must be \"method\" or \"field\".")

// Prettier 3.7.1
console.log('A descriptor\\'s .kind must be "method" or "field".');

// Prettier 3.7.2
console.log('A descriptor\\\'s .kind must be "method" or "field".');
```

##### JavaScript: Preserve quote for embedded HTML attribute values ([#&#8203;18352](https://redirect.github.com/prettier/prettier/pull/18352) by [@&#8203;kovsu](https://redirect.github.com/kovsu))

<!-- prettier-ignore -->

```tsx
// Input
const html = /* HTML */ ` <div class="${styles.banner}"></div> `;

// Prettier 3.7.1
const html = /* HTML */ ` <div class=${styles.banner}></div> `;

// Prettier 3.7.2
const html = /* HTML */ ` <div class="${styles.banner}"></div> `;
```

##### TypeScript: Fix comment in empty type literal ([#&#8203;18364](https://redirect.github.com/prettier/prettier/pull/18364) by [@&#8203;fisker](https://redirect.github.com/fisker))

<!-- prettier-ignore -->

```tsx
// Input
export type XXX = {
  // tbd
};

// Prettier 3.7.1
export type XXX = { // tbd };

// Prettier 3.7.2
export type XXX = {
  // tbd
};
```

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

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Reviews

### Review by @TenSt - Approved on December 01, 2025 at 10:46 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1318*
