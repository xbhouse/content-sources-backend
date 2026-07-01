---
type: pull_request
number: 1294
title: "Build fix(deps): update dependency dotenv to v17"
state: merged
author: red-hat-konflux
created: 2025-11-10T08:21:00Z
updated: 2025-11-10T20:32:00Z
closed: 2025-11-10T17:35:16Z
merged: 2025-11-10T17:35:16Z
base_branch: main
head_branch: konflux/mintmaker/main/dotenv-17.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1294
---

# Pull Request #1294: Build fix(deps): update dependency dotenv to v17

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 08:21 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/dotenv-17.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [dotenv](https://redirect.github.com/motdotla/dotenv) | [`^16.4.7` -> `^17.0.0`](https://renovatebot.com/diffs/npm/dotenv/16.6.1/17.2.3) | [![age](https://developer.mend.io/api/mc/badges/age/npm/dotenv/17.2.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/npm/dotenv/16.6.1/17.2.3?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>motdotla/dotenv (dotenv)</summary>

### [`v17.2.3`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1723-2025-09-29)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v17.2.2...v17.2.3)

##### Changed

- Fixed typescript error definition ([#&#8203;912](https://redirect.github.com/motdotla/dotenv/pull/912))

### [`v17.2.2`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1722-2025-09-02)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v17.2.1...v17.2.2)

##### Added

- 🙏 A big thank you to new sponsor [Tuple.app](https://tuple.app/dotenv) - *the premier screen sharing app for developers on macOS and Windows.* Go check them out. It's wonderful and generous of them to give back to open source by sponsoring dotenv. Give them some love back.

### [`v17.2.1`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1721-2025-07-24)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v17.2.0...v17.2.1)

##### Changed

- Fix clickable tip links by removing parentheses ([#&#8203;897](https://redirect.github.com/motdotla/dotenv/pull/897))

### [`v17.2.0`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1720-2025-07-09)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v17.1.0...v17.2.0)

##### Added

- Optionally specify `DOTENV_CONFIG_QUIET=true` in your environment or `.env` file to quiet the runtime log ([#&#8203;889](https://redirect.github.com/motdotla/dotenv/pull/889))
- Just like dotenv any `DOTENV_CONFIG_` environment variables take precedence over any code set options like `({quiet: false})`

```ini
```

### [`v17.1.0`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1710-2025-07-07)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v17.0.1...v17.1.0)

##### Added

- Add additional security and configuration tips to the runtime log ([#&#8203;884](https://redirect.github.com/motdotla/dotenv/pull/884))
- Dim the tips text from the main injection information text

```js
const TIPS = [
  '🔐 encrypt with dotenvx: https://dotenvx.com',
  '🔐 prevent committing .env to code: https://dotenvx.com/precommit',
  '🔐 prevent building .env in docker: https://dotenvx.com/prebuild',
  '🛠️  run anywhere with `dotenvx run -- yourcommand`',
  '⚙️  specify custom .env file path with { path: \'/custom/path/.env\' }',
  '⚙️  enable debug logging with { debug: true }',
  '⚙️  override existing env vars with { override: true }',
  '⚙️  suppress all logs with { quiet: true }',
  '⚙️  write to custom object with { processEnv: myObject }',
  '⚙️  load multiple .env files with { path: [\'.env.local\', \'.env\'] }'
]
```

### [`v17.0.1`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1701-2025-07-01)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v17.0.0...v17.0.1)

##### Changed

- Patched injected log to count only populated/set keys to process.env ([#&#8203;879](https://redirect.github.com/motdotla/dotenv/pull/879))

### [`v17.0.0`](https://redirect.github.com/motdotla/dotenv/blob/HEAD/CHANGELOG.md#1700-2025-06-27)

[Compare Source](https://redirect.github.com/motdotla/dotenv/compare/v16.6.1...v17.0.0)

##### Changed

- Default `quiet` to false - informational (file and keys count) runtime log message shows by default ([#&#8203;875](https://redirect.github.com/motdotla/dotenv/pull/875))

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

### Review by @swadeley - Approved on November 10, 2025 at 04:16 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1294*
