---
type: pull_request
number: 57
title: "Update module github.com/spf13/cast to v1.10.0"
state: merged
author: red-hat-konflux
created: 2025-11-10T12:19:23Z
updated: 2025-11-10T15:19:16Z
closed: 2025-11-10T14:01:09Z
merged: 2025-11-10T14:01:09Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-spf13-cast-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/57
---

# Pull Request #57: Update module github.com/spf13/cast to v1.10.0

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 12:19 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-spf13-cast-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/spf13/cast](https://redirect.github.com/spf13/cast) | `v1.7.0` -> `v1.10.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fspf13%2fcast/v1.10.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fspf13%2fcast/v1.7.0/v1.10.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>spf13/cast (github.com/spf13/cast)</summary>

### [`v1.10.0`](https://redirect.github.com/spf13/cast/releases/tag/v1.10.0)

[Compare Source](https://redirect.github.com/spf13/cast/compare/v1.9.2...v1.10.0)

#### What's Changed

- build(deps): bump ossf/scorecard-action from 2.4.1 to 2.4.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;275](https://redirect.github.com/spf13/cast/pull/275)
- build(deps): bump github/codeql-action from 3.28.18 to 3.28.19 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;277](https://redirect.github.com/spf13/cast/pull/277)
- build(deps): bump github/codeql-action from 3.28.19 to 3.29.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;289](https://redirect.github.com/spf13/cast/pull/289)
- build(deps): bump github/codeql-action from 3.29.7 to 3.29.10 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;296](https://redirect.github.com/spf13/cast/pull/296)
- build(deps): bump actions/dependency-review-action from 4.7.1 to 4.7.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;295](https://redirect.github.com/spf13/cast/pull/295)
- build(deps): bump actions/checkout from 4.2.2 to 5.0.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;293](https://redirect.github.com/spf13/cast/pull/293)
- build(deps): bump github/codeql-action from 3.29.10 to 3.30.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;301](https://redirect.github.com/spf13/cast/pull/301)
- build(deps): bump actions/setup-go from 5.5.0 to 6.0.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;300](https://redirect.github.com/spf13/cast/pull/300)
- build(deps): bump actions/dependency-review-action from 4.7.2 to 4.7.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;298](https://redirect.github.com/spf13/cast/pull/298)
- Always return empty map instead of nil when conversion fails by [@&#8203;andig](https://redirect.github.com/andig) in [#&#8203;283](https://redirect.github.com/spf13/cast/pull/283)

#### New Contributors

- [@&#8203;andig](https://redirect.github.com/andig) made their first contribution in [#&#8203;283](https://redirect.github.com/spf13/cast/pull/283)

**Full Changelog**: <https://github.com/spf13/cast/compare/v1.9.2...v1.10.0>

### [`v1.9.2`](https://redirect.github.com/spf13/cast/releases/tag/v1.9.2)

[Compare Source](https://redirect.github.com/spf13/cast/compare/v1.9.1...v1.9.2)

#### What's Changed

- fix: float string to number parsing by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;276](https://redirect.github.com/spf13/cast/pull/276)

**Full Changelog**: <https://github.com/spf13/cast/compare/v1.9.1...v1.9.2>

### [`v1.9.1`](https://redirect.github.com/spf13/cast/releases/tag/v1.9.1)

[Compare Source](https://redirect.github.com/spf13/cast/compare/v1.9.0...v1.9.1)

#### What's Changed

- fix: indirection of typed nils by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;273](https://redirect.github.com/spf13/cast/pull/273)

**Full Changelog**: <https://github.com/spf13/cast/compare/v1.9.0...v1.9.1>

### [`v1.9.0`](https://redirect.github.com/spf13/cast/releases/tag/v1.9.0)

[Compare Source](https://redirect.github.com/spf13/cast/compare/v1.8.0...v1.9.0)

#### Notable new features 🎉

- Casting *from* type aliases is now supported for basic types
- Added generic functions: `To`/`ToE`, `Must`, `ToNumber`/`ToNumberE`
- Increased test coverage
- Converting float numbers from string is now supported

> \[!WARNING]
> Since cast now supports converting float values from strings, a related edge case behaves differently:
>
> In previous versions, attempting to convert an empty string to a float **resulted in an error**.
>
> Starting with this version, the same operation **no longer raises an error**.
>
> To maintain consistency with the rest of the library, an empty string now converts to the float value `0.0`.

#### What's Changed

- build(deps): bump github/codeql-action from 3.28.17 to 3.28.18 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;248](https://redirect.github.com/spf13/cast/pull/248)
- build(deps): bump actions/dependency-review-action from 4.6.0 to 4.7.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;247](https://redirect.github.com/spf13/cast/pull/247)
- build(deps): bump actions/setup-go from 5.4.0 to 5.5.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;245](https://redirect.github.com/spf13/cast/pull/245)
- refactor: move number parsing to generic functions by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;250](https://redirect.github.com/spf13/cast/pull/250)
- Improve ToString/ToStringE performance by [@&#8203;ganigeorgiev](https://redirect.github.com/ganigeorgiev) in [#&#8203;244](https://redirect.github.com/spf13/cast/pull/244)
- Split caste.go into smaller files by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;251](https://redirect.github.com/spf13/cast/pull/251)
- refactor: remove unused initial int conversion by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;253](https://redirect.github.com/spf13/cast/pull/253)
- Generate code to make maintenance easier by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;252](https://redirect.github.com/spf13/cast/pull/252)
- feat: add To and ToNumber functions by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;255](https://redirect.github.com/spf13/cast/pull/255)
- build(deps): bump golangci/golangci-lint-action from 7.0.0 to 8.0.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;243](https://redirect.github.com/spf13/cast/pull/243)
- Move tests next to their implementation by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;256](https://redirect.github.com/spf13/cast/pull/256)
- Refactor number tests by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;257](https://redirect.github.com/spf13/cast/pull/257)
- feat: return 0 when casting an empty string to a number by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;259](https://redirect.github.com/spf13/cast/pull/259)
- Support converting string float numbers to integer types by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;261](https://redirect.github.com/spf13/cast/pull/261)
- Test improvements by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;262](https://redirect.github.com/spf13/cast/pull/262)
- Improvements by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;263](https://redirect.github.com/spf13/cast/pull/263)
- refactor: return indirection result from indirect function by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;264](https://redirect.github.com/spf13/cast/pull/264)
- Slice improvements by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;265](https://redirect.github.com/spf13/cast/pull/265)
- refactor: move error message to a constant by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;267](https://redirect.github.com/spf13/cast/pull/267)
- chore: improve map cast functions by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;269](https://redirect.github.com/spf13/cast/pull/269)
- Resolve aliases by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;271](https://redirect.github.com/spf13/cast/pull/271)

#### New Contributors

- [@&#8203;ganigeorgiev](https://redirect.github.com/ganigeorgiev) made their first contribution in [#&#8203;244](https://redirect.github.com/spf13/cast/pull/244)

**Full Changelog**: <https://github.com/spf13/cast/compare/v1.8.0...v1.9.0>

### [`v1.8.0`](https://redirect.github.com/spf13/cast/releases/tag/v1.8.0)

[Compare Source](https://redirect.github.com/spf13/cast/compare/v1.7.1...v1.8.0)

#### What's Changed

- Updates by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;237](https://redirect.github.com/spf13/cast/pull/237)
- Bump actions/setup-go from 4.0.1 to 4.1.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;193](https://redirect.github.com/spf13/cast/pull/193)
- Generic by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;238](https://redirect.github.com/spf13/cast/pull/238)
- Add unsigned integer support in ToStringSliceE function by [@&#8203;nicklaus-dev](https://redirect.github.com/nicklaus-dev) in [#&#8203;200](https://redirect.github.com/spf13/cast/pull/200)
- Add function to cast `interface{}` to `[]float64` by [@&#8203;ste93cry](https://redirect.github.com/ste93cry) in [#&#8203;179](https://redirect.github.com/spf13/cast/pull/179)
- Add cast methods ToUintSlice by [@&#8203;nmvalera](https://redirect.github.com/nmvalera) in [#&#8203;236](https://redirect.github.com/spf13/cast/pull/236)
- build(deps): bump actions/dependency-review-action from 4.5.0 to 4.6.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;240](https://redirect.github.com/spf13/cast/pull/240)
- build(deps): bump github/codeql-action from 2.13.4 to 3.28.15 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;239](https://redirect.github.com/spf13/cast/pull/239)
- build(deps): bump github/codeql-action from 3.28.15 to 3.28.17 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;242](https://redirect.github.com/spf13/cast/pull/242)
- Add ToInt64Slice() and ToInt64SliceE() by [@&#8203;arui1628](https://redirect.github.com/arui1628) in [#&#8203;234](https://redirect.github.com/spf13/cast/pull/234)

#### New Contributors

- [@&#8203;nicklaus-dev](https://redirect.github.com/nicklaus-dev) made their first contribution in [#&#8203;200](https://redirect.github.com/spf13/cast/pull/200)
- [@&#8203;ste93cry](https://redirect.github.com/ste93cry) made their first contribution in [#&#8203;179](https://redirect.github.com/spf13/cast/pull/179)
- [@&#8203;nmvalera](https://redirect.github.com/nmvalera) made their first contribution in [#&#8203;236](https://redirect.github.com/spf13/cast/pull/236)
- [@&#8203;arui1628](https://redirect.github.com/arui1628) made their first contribution in [#&#8203;234](https://redirect.github.com/spf13/cast/pull/234)

**Full Changelog**: <https://github.com/spf13/cast/compare/v1.7.1...v1.8.0>

### [`v1.7.1`](https://redirect.github.com/spf13/cast/releases/tag/v1.7.1)

[Compare Source](https://redirect.github.com/spf13/cast/compare/v1.7.0...v1.7.1)

#### What's Changed

- Update README.md by [@&#8203;lesichkovm](https://redirect.github.com/lesichkovm) in [#&#8203;224](https://redirect.github.com/spf13/cast/pull/224)
- Fix ToUint64 issue with string input exceeding maximum int64 by [@&#8203;skyjerry](https://redirect.github.com/skyjerry) in [#&#8203;213](https://redirect.github.com/spf13/cast/pull/213)

#### New Contributors

- [@&#8203;lesichkovm](https://redirect.github.com/lesichkovm) made their first contribution in [#&#8203;224](https://redirect.github.com/spf13/cast/pull/224)
- [@&#8203;skyjerry](https://redirect.github.com/skyjerry) made their first contribution in [#&#8203;213](https://redirect.github.com/spf13/cast/pull/213)

**Full Changelog**: <https://github.com/spf13/cast/compare/v1.7.0...v1.7.1>

</details>

---

### Configuration

📅 **Schedule**: Branch creation - At any time (no schedule defined), Automerge - At any time (no schedule defined).

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

### Review by @Hyperkid123 - Approved on November 10, 2025 at 02:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/57*
