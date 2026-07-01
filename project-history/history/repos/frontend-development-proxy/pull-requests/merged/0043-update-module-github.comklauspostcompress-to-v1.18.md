---
type: pull_request
number: 43
title: "Update module github.com/klauspost/compress to v1.18.2"
state: merged
author: red-hat-konflux
created: 2025-11-06T04:18:47Z
updated: 2025-12-03T08:54:15Z
closed: 2025-12-03T08:54:15Z
merged: 2025-12-03T08:54:15Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-klauspost-compress-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/43
---

# Pull Request #43: Update module github.com/klauspost/compress to v1.18.2

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 04:18 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-klauspost-compress-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/klauspost/compress](https://redirect.github.com/klauspost/compress) | `v1.18.0` -> `v1.18.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fklauspost%2fcompress/v1.18.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fklauspost%2fcompress/v1.18.0/v1.18.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>klauspost/compress (github.com/klauspost/compress)</summary>

### [`v1.18.2`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.2)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.18.1...v1.18.2)

#### What's Changed

- Fix invalid encoding on level 9 with single value input by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1115](https://redirect.github.com/klauspost/compress/pull/1115)
- flate: reduce stateless allocations by [@&#8203;RXamzin](https://redirect.github.com/RXamzin) in [#&#8203;1106](https://redirect.github.com/klauspost/compress/pull/1106)
- build(deps): bump github/codeql-action from 3.30.5 to 4.31.2 in the github-actions group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1111](https://redirect.github.com/klauspost/compress/pull/1111)

`v1.18.1` is marked "retracted" due to invalid flate/zip/gzip encoding.

#### New Contributors

- [@&#8203;RXamzin](https://redirect.github.com/RXamzin) made their first contribution in [#&#8203;1106](https://redirect.github.com/klauspost/compress/pull/1106)

**Full Changelog**: <https://github.com/klauspost/compress/compare/v1.18.1...v1.18.2>

### [`v1.18.1`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.1)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.18.0...v1.18.1)

#### What's Changed

- zstd: Fix incorrect buffer size in dictionary encodes by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1059](https://redirect.github.com/klauspost/compress/pull/1059)
- s2: check for cap, not len of buffer in EncodeBetter/Best by [@&#8203;vdarulis](https://redirect.github.com/vdarulis) in [#&#8203;1080](https://redirect.github.com/klauspost/compress/pull/1080)
- zstd: Add simple zstd EncodeTo/DecodeTo functions by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1079](https://redirect.github.com/klauspost/compress/pull/1079)
- zlib: Avoiding extra allocation in zlib.reader.Reset by [@&#8203;travelpolicy](https://redirect.github.com/travelpolicy) in [#&#8203;1086](https://redirect.github.com/klauspost/compress/pull/1086)
- gzhttp: remove redundant err check in zstdReader by [@&#8203;ryanfowler](https://redirect.github.com/ryanfowler) in [#&#8203;1090](https://redirect.github.com/klauspost/compress/pull/1090)
- Run modernize. Deprecate Go 1.22 by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1095](https://redirect.github.com/klauspost/compress/pull/1095)
- flate: Simplify matchlen by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1101](https://redirect.github.com/klauspost/compress/pull/1101)
- flate: Add examples by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1102](https://redirect.github.com/klauspost/compress/pull/1102)
- flate: Use exact sizes for huffman tables by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1103](https://redirect.github.com/klauspost/compress/pull/1103)
- flate: Faster load+store by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1104](https://redirect.github.com/klauspost/compress/pull/1104)
- Add notice to S2 about MinLZ by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [#&#8203;1065](https://redirect.github.com/klauspost/compress/pull/1065)

#### New Contributors

- [@&#8203;wooffie](https://redirect.github.com/wooffie) made their first contribution in [#&#8203;1069](https://redirect.github.com/klauspost/compress/pull/1069)
- [@&#8203;vdarulis](https://redirect.github.com/vdarulis) made their first contribution in [#&#8203;1080](https://redirect.github.com/klauspost/compress/pull/1080)
- [@&#8203;travelpolicy](https://redirect.github.com/travelpolicy) made their first contribution in [#&#8203;1086](https://redirect.github.com/klauspost/compress/pull/1086)
- [@&#8203;ryanfowler](https://redirect.github.com/ryanfowler) made their first contribution in [#&#8203;1090](https://redirect.github.com/klauspost/compress/pull/1090)

**Full Changelog**: <https://github.com/klauspost/compress/compare/v1.18.0...v1.18.1>

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

### Review by @Hyperkid123 - Approved on December 03, 2025 at 08:54 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/43*
