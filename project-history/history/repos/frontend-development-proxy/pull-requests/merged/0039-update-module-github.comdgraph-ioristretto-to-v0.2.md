---
type: pull_request
number: 39
title: "Update module github.com/dgraph-io/ristretto to v0.2.0"
state: merged
author: red-hat-konflux
created: 2025-11-05T20:18:31Z
updated: 2025-11-18T09:49:48Z
closed: 2025-11-18T09:49:47Z
merged: 2025-11-18T09:49:47Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-dgraph-io-ristretto-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/39
---

# Pull Request #39: Update module github.com/dgraph-io/ristretto to v0.2.0

**Author**: @red-hat-konflux
**Created**: November 05, 2025 at 08:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-dgraph-io-ristretto-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/dgraph-io/ristretto](https://redirect.github.com/dgraph-io/ristretto) | `v0.1.0` -> `v0.2.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fdgraph-io%2fristretto/v0.2.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fdgraph-io%2fristretto/v0.1.0/v0.2.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>dgraph-io/ristretto (github.com/dgraph-io/ristretto)</summary>

### [`v0.2.0`](https://redirect.github.com/dgraph-io/ristretto/blob/HEAD/CHANGELOG.md#v020---2024-10-06)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v0.1.1...v0.2.0)

**Added**

- [fix: support compilation to wasip1 by @&#8203;achille-roussel](https://redirect.github.com/hypermodeinc/ristretto/pull/344)
- [add config for cleanup ticker duration by @&#8203;singhvikash11](https://redirect.github.com/hypermodeinc/ristretto/pull/342)

**Fixed**

- [docs(readme): Use new Wait method by @&#8203;angadn](https://redirect.github.com/hypermodeinc/ristretto/pull/327)
- [docs: format example on readme by @&#8203;rfyiamcool](https://redirect.github.com/hypermodeinc/ristretto/pull/339)
- [Fix flakes in TestDropUpdates by @&#8203;evanj](https://redirect.github.com/hypermodeinc/ristretto/pull/334)
- [docs(Cache): document Wait, clarify Get by @&#8203;evanj](https://redirect.github.com/hypermodeinc/ristretto/pull/333)
- [chore: fix typo error by @&#8203;proost](https://redirect.github.com/hypermodeinc/ristretto/pull/341)
- [remove glog dependency by @&#8203;jhawk28](https://redirect.github.com/hypermodeinc/ristretto/pull/350)
- [fix(OnEvict): Set missing Expiration field on evicted items by @&#8203;0x1ee7](https://redirect.github.com/hypermodeinc/ristretto/pull/345)
- [uint32 -> uint64 in slice methods by @&#8203;mocurin](https://redirect.github.com/hypermodeinc/ristretto/pull/323)
- [fix: cleanupTicker not being stopped by @&#8203;IlyaFloppy](https://redirect.github.com/hypermodeinc/ristretto/pull/343)

**Full Changelog**: <https://github.com/hypermodeinc/ristretto/compare/v0.1.1...v0.2.0>

### [`v0.1.1`](https://redirect.github.com/dgraph-io/ristretto/releases/tag/v0.1.1)

[Compare Source](https://redirect.github.com/dgraph-io/ristretto/compare/v0.1.0...v0.1.1)

#### [0.1.1] - 2022-10-12

[0.1.1]: https://redirect.github.com/dgraph-io/ristretto/compare/v0.1.0..v0.1.1

This release fixes certain arm64 build issues in the z package.  It also incorporates CI steps in our repository.

##### Changed

- [chore(docs): Include SpiceDB in the list of projects using Ristretto (#&#8203;285)](https://redirect.github.com/dgraph-io/ristretto/pull/311)

##### Added

- [Run CI Jobs via Github Actions #&#8203;304](https://redirect.github.com/dgraph-io/ristretto/pull/304)

##### Fixed

- [fix(build): update x/sys dependency](https://redirect.github.com/dgraph-io/ristretto/pull/308)
- [fix(z): Address inconsistent mremap return arguments with arm64](https://redirect.github.com/dgraph-io/ristretto/pull/309)
- [fix(z): runtime error: index out of range for !amd64 env #&#8203;287](https://redirect.github.com/dgraph-io/ristretto/pull/307)

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

### Review by @Hyperkid123 - Approved on November 18, 2025 at 09:49 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/39*
