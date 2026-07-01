---
type: pull_request
number: 47
title: "Update module github.com/Masterminds/semver/v3 to v3.4.0"
state: merged
author: red-hat-konflux
created: 2025-11-06T12:21:53Z
updated: 2025-11-07T09:03:52Z
closed: 2025-11-07T09:03:50Z
merged: 2025-11-07T09:03:50Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-masterminds-semver-v3-3.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/47
---

# Pull Request #47: Update module github.com/Masterminds/semver/v3 to v3.4.0

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 12:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-masterminds-semver-v3-3.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/Masterminds/semver/v3](https://redirect.github.com/Masterminds/semver) | `v3.3.0` -> `v3.4.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fMasterminds%2fsemver%2fv3/v3.4.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fMasterminds%2fsemver%2fv3/v3.3.0/v3.4.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>Masterminds/semver (github.com/Masterminds/semver/v3)</summary>

### [`v3.4.0`](https://redirect.github.com/Masterminds/semver/releases/tag/v3.4.0)

[Compare Source](https://redirect.github.com/Masterminds/semver/compare/v3.3.1...v3.4.0)

There are a few changes in this release to highlight:

1. `Constraints` now has a property `IncludePrerelease`. When set to true the `Check` and `Validate` methods will include prereleases.
2. When an AND group has one constraint with a prerelease but more than one constraint then prereleases will be included. For example, `>1.0.0-beta.1 < 2`. In the past this would not have included prereleases because each constraint needed to have a prerelease. Now, only one constraint needs to have a prerelease. This is considered a long standing bug fix. Note, this does not carry across OR groups. For example, `>1.0.0-beta.1 < 2 || > 3`. In this case, prereleases will not be included when evaluating against `>3`.
3. `NewVersion` coercion with leading "0"'s is restored. This can be disabled by setting the package level property `CoerceNewVersion` to `false`.

#### What's Changed

- fix the CodeQL link by [@&#8203;dmitris](https://redirect.github.com/dmitris) in [#&#8203;257](https://redirect.github.com/Masterminds/semver/pull/257)
- Restore detailed errors when failed to parse with NewVersion by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;262](https://redirect.github.com/Masterminds/semver/pull/262)
- updating go version tested with by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;263](https://redirect.github.com/Masterminds/semver/pull/263)
- Restore the ability to have leading 0's with NewVersion by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;266](https://redirect.github.com/Masterminds/semver/pull/266)
- Handle pre-releases on all in an and group by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;267](https://redirect.github.com/Masterminds/semver/pull/267)
- Add property to include prereleases by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;268](https://redirect.github.com/Masterminds/semver/pull/268)
- Updating the error message handling by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;269](https://redirect.github.com/Masterminds/semver/pull/269)
- Update the release notes and readme for new version by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;270](https://redirect.github.com/Masterminds/semver/pull/270)

#### New Contributors

- [@&#8203;dmitris](https://redirect.github.com/dmitris) made their first contribution in [#&#8203;257](https://redirect.github.com/Masterminds/semver/pull/257)

**Full Changelog**: <https://github.com/Masterminds/semver/compare/v3.3.1...v3.4.0>

### [`v3.3.1`](https://redirect.github.com/Masterminds/semver/releases/tag/v3.3.1)

[Compare Source](https://redirect.github.com/Masterminds/semver/compare/v3.3.0...v3.3.1)

#### What's Changed

- Fix for allowing some version that were invalid by [@&#8203;mattfarina](https://redirect.github.com/mattfarina) in [#&#8203;253](https://redirect.github.com/Masterminds/semver/pull/253)

**Full Changelog**: <https://github.com/Masterminds/semver/compare/v3.3.0...v3.3.1>

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

### Review by @Hyperkid123 - Approved on November 07, 2025 at 09:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/47*
