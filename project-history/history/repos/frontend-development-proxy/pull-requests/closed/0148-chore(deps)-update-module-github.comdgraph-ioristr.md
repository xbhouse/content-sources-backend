---
type: pull_request
number: 148
title: "chore(deps): update module github.com/dgraph-io/ristretto to v2 - autoclosed"
state: closed
author: red-hat-konflux
created: 2026-06-23T17:05:35Z
updated: 2026-06-26T20:35:13Z
closed: 2026-06-26T20:35:13Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-dgraph-io-ristretto-2.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/148
---

# Pull Request #148: chore(deps): update module github.com/dgraph-io/ristretto to v2 - autoclosed

**Author**: @red-hat-konflux
**Created**: June 23, 2026 at 05:05 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-dgraph-io-ristretto-2.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/dgraph-io/ristretto](https://redirect.github.com/dgraph-io/ristretto) | `v0.2.0` → `v2.4.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fdgraph-io%2fristretto/v2.4.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fdgraph-io%2fristretto/v0.2.0/v2.4.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>dgraph-io/ristretto (github.com/dgraph-io/ristretto)</summary>

### [`v2.4.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v240---2026-01-21)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v2.3.0...v2.4.0)

##### Added

- Implement public `Cache.IterValues()` method ([#&#8203;475](https://redirect.github.com/dgraph-io/ristretto/issues/475))
- Allow custom key types with underlying types in Key constraint ([#&#8203;478](https://redirect.github.com/dgraph-io/ristretto/issues/478))

##### Fixed

- Fix compilation on 32-bit archs ([#&#8203;465](https://redirect.github.com/dgraph-io/ristretto/issues/465))

**Full Changelog**: <https://github.com/dgraph-io/ristretto/compare/v2.3.0...v2.4.0>

### [`v2.3.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v240---2026-01-21)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v2.2.0...v2.3.0)

##### Added

- Implement public `Cache.IterValues()` method ([#&#8203;475](https://redirect.github.com/dgraph-io/ristretto/issues/475))
- Allow custom key types with underlying types in Key constraint ([#&#8203;478](https://redirect.github.com/dgraph-io/ristretto/issues/478))

##### Fixed

- Fix compilation on 32-bit archs ([#&#8203;465](https://redirect.github.com/dgraph-io/ristretto/issues/465))

**Full Changelog**: <https://github.com/dgraph-io/ristretto/compare/v2.3.0...v2.4.0>

### [`v2.2.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v230---2025-08-19)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v2.1.0...v2.2.0)

##### Added

- Add public `Cache.RemainingCost()` method ([#&#8203;448](https://redirect.github.com/dgraph-io/ristretto/issues/448))
- Add support for uint keys ([#&#8203;463](https://redirect.github.com/dgraph-io/ristretto/issues/463))

##### Fixed

- Fix typo: ffor → for ([#&#8203;456](https://redirect.github.com/dgraph-io/ristretto/issues/456))
- Correct grammar in error message ([#&#8203;461](https://redirect.github.com/dgraph-io/ristretto/issues/461))

**Full Changelog**: <https://github.com/dgraph-io/ristretto/compare/v2.2.0...v2.3.0>

### [`v2.1.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v220---2025-03-30)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v2.0.1...v2.1.0)

##### Changed

- Remove dependency: github.com/pkg/errors ([#&#8203;443](https://redirect.github.com/dgraph-io/ristretto/issues/443))

##### Fixed

- Switch from using a sync.WaitGroup to closing a channel of struct{} ([#&#8203;442](https://redirect.github.com/dgraph-io/ristretto/issues/442))

**Full Changelog**: <https://github.com/dgraph-io/ristretto/compare/v2.1.0...v2.2.0>

### [`v2.0.1`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v210---2025-01-09)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v2.0.0...v2.0.1)

##### Added

- Add `ShouldUpdate()` function in config ([#&#8203;427](https://redirect.github.com/dgraph-io/ristretto/issues/427))

##### Fixed

- Fix memory leak while cleaning up expiration map ([#&#8203;429](https://redirect.github.com/dgraph-io/ristretto/issues/429))
- Execute `m.Unlock` in defer in store.go ([#&#8203;425](https://redirect.github.com/dgraph-io/ristretto/issues/425))

**Full Changelog**: <https://github.com/dgraph-io/ristretto/compare/v2.0.1...v2.1.0>

### [`v2.0.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v201---2024-12-11)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v1.0.1...v2.0.0)

**Fixed**

- Wait for goroutines to finish ([#&#8203;423](https://redirect.github.com/dgraph-io/ristretto/issues/423))
- Bump golang.org/x/sys from 0.27.0 to 0.28.0 in the minor group ([#&#8203;421](https://redirect.github.com/dgraph-io/ristretto/issues/421))
- Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 in the minor group ([#&#8203;420](https://redirect.github.com/dgraph-io/ristretto/issues/420))
- Bump golang.org/x/sys from 0.26.0 to 0.27.0 in the minor group ([#&#8203;419](https://redirect.github.com/dgraph-io/ristretto/issues/419))

**Full Changelog**: <https://github.com/dgraph-io/ristretto/compare/v2.0.0...v2.0.1>

### [`v1.0.1`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v101)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v1.0.0...v1.0.1)

**This release is deprecated**

### [`v1.0.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v100)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v0.2.0...v1.0.0)

**This release is deprecated**

</details>

---

### Configuration

📅 **Schedule**: (UTC)

- Branch creation
  - At any time (no schedule defined)
- Automerge
  - At any time (no schedule defined)

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQzLjIxMC4wLXJwbSIsInRhcmdldEJyYW5jaCI6Im1haW4iLCJsYWJlbHMiOltdfQ==-->


---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/148*
