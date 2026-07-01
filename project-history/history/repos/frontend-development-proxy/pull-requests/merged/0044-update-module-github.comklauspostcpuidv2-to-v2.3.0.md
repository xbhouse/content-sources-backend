---
type: pull_request
number: 44
title: "Update module github.com/klauspost/cpuid/v2 to v2.3.0"
state: merged
author: red-hat-konflux
created: 2025-11-06T04:18:50Z
updated: 2025-11-07T09:14:22Z
closed: 2025-11-07T09:14:17Z
merged: 2025-11-07T09:14:17Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-klauspost-cpuid-v2-2.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/44
---

# Pull Request #44: Update module github.com/klauspost/cpuid/v2 to v2.3.0

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 04:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-klauspost-cpuid-v2-2.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/klauspost/cpuid/v2](https://redirect.github.com/klauspost/cpuid) | `v2.2.9` -> `v2.3.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fklauspost%2fcpuid%2fv2/v2.3.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fklauspost%2fcpuid%2fv2/v2.2.9/v2.3.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>klauspost/cpuid (github.com/klauspost/cpuid/v2)</summary>

### [`v2.3.0`](https://redirect.github.com/klauspost/cpuid/releases/tag/v2.3.0)

[Compare Source](https://redirect.github.com/klauspost/cpuid/compare/v2.2.11...v2.3.0)

#### What's Changed

- Add in PMU parsing from CPUID leaf 0xA for Intel processors by [@&#8203;echiugoog](https://redirect.github.com/echiugoog) in [#&#8203;165](https://redirect.github.com/klauspost/cpuid/pull/165)
- Add SGXPQC detection by [@&#8203;ozhuraki](https://redirect.github.com/ozhuraki) in [#&#8203;163](https://redirect.github.com/klauspost/cpuid/pull/163)
- Detect AMD TSA mitigations by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;166](https://redirect.github.com/klauspost/cpuid/pull/166)

#### New Contributors

- [@&#8203;echiugoog](https://redirect.github.com/echiugoog) made their first contribution in [#&#8203;165](https://redirect.github.com/klauspost/cpuid/pull/165)

**Full Changelog**: <https://github.com/klauspost/cpuid/compare/v2.2.11...v2.3.0>

### [`v2.2.11`](https://redirect.github.com/klauspost/cpuid/releases/tag/v2.2.11)

[Compare Source](https://redirect.github.com/klauspost/cpuid/compare/v2.2.10...v2.2.11)

#### What's Changed

- Add AMXTRANSPOSE detection by [@&#8203;ozhuraki](https://redirect.github.com/ozhuraki) in [#&#8203;156](https://redirect.github.com/klauspost/cpuid/pull/156)
- Add SM3, SM4 detection on x86 by [@&#8203;ozhuraki](https://redirect.github.com/ozhuraki) in [#&#8203;157](https://redirect.github.com/klauspost/cpuid/pull/157)
- darwin/arm64: fix SIMD detection and improve ARM feature probing by [@&#8203;HippoBaro](https://redirect.github.com/HippoBaro) in [#&#8203;160](https://redirect.github.com/klauspost/cpuid/pull/160)
- fix: Fix division by zero in physicalCores on intel by [@&#8203;skartikey](https://redirect.github.com/skartikey) in [#&#8203;162](https://redirect.github.com/klauspost/cpuid/pull/162)

#### New Contributors

- [@&#8203;HippoBaro](https://redirect.github.com/HippoBaro) made their first contribution in [#&#8203;160](https://redirect.github.com/klauspost/cpuid/pull/160)
- [@&#8203;skartikey](https://redirect.github.com/skartikey) made their first contribution in [#&#8203;162](https://redirect.github.com/klauspost/cpuid/pull/162)

**Full Changelog**: <https://github.com/klauspost/cpuid/compare/v2.2.10...v2.2.11>

### [`v2.2.10`](https://redirect.github.com/klauspost/cpuid/releases/tag/v2.2.10)

[Compare Source](https://redirect.github.com/klauspost/cpuid/compare/v2.2.9...v2.2.10)

#### What's Changed

- Add rest of fields from ID\_AA64ISAR0\_EL1 by [@&#8203;craciunoiuc](https://redirect.github.com/craciunoiuc) in [#&#8203;152](https://redirect.github.com/klauspost/cpuid/pull/152)
- Add AMXCOMPLEX detection by [@&#8203;ozhuraki](https://redirect.github.com/ozhuraki) in [#&#8203;153](https://redirect.github.com/klauspost/cpuid/pull/153)
- Add AMXTF32 detection by [@&#8203;ozhuraki](https://redirect.github.com/ozhuraki) in [#&#8203;154](https://redirect.github.com/klauspost/cpuid/pull/154)
- Update versions & releaser by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;155](https://redirect.github.com/klauspost/cpuid/pull/155)

#### New Contributors

- [@&#8203;craciunoiuc](https://redirect.github.com/craciunoiuc) made their first contribution in [#&#8203;152](https://redirect.github.com/klauspost/cpuid/pull/152)

**Full Changelog**: <https://github.com/klauspost/cpuid/compare/v2.2.9...v2.2.10>

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

### Review by @Hyperkid123 - Approved on November 07, 2025 at 09:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/44*
