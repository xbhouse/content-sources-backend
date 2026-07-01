---
type: pull_request
number: 77
title: "chore(deps): update module github.com/dgraph-io/badger to v4"
state: open
author: red-hat-konflux
created: 2025-11-15T09:11:50Z
updated: 2026-06-27T13:19:55Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-dgraph-io-badger-4.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/77
---

# Pull Request #77: chore(deps): update module github.com/dgraph-io/badger to v4

**Author**: @red-hat-konflux
**Created**: November 15, 2025 at 09:11 AM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-dgraph-io-badger-4.x`

## Description

> ℹ️ **Note**
> 
> This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/dgraph-io/badger](https://redirect.github.com/dgraph-io/badger) | `v1.6.2` → `v4.9.2` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fdgraph-io%2fbadger/v4.9.2?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fdgraph-io%2fbadger/v1.6.2/v4.9.2?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>dgraph-io/badger (github.com/dgraph-io/badger)</summary>

### [`v4.9.2`](https://redirect.github.com/dgraph-io/badger/releases/tag/v4.9.2)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.9.1...v4.9.2)

#### What's Changed

- chore(core): remove unused event log by [@&#8203;xqqp](https://redirect.github.com/xqqp) in [#&#8203;2257](https://redirect.github.com/dgraph-io/badger/pull/2257)
- fix(cd): upload build artifacts to GitHub Release by [@&#8203;matthewmcneely](https://redirect.github.com/matthewmcneely) in [#&#8203;2273](https://redirect.github.com/dgraph-io/badger/pull/2273)
- fix: Prevent NPE on sync with inmemory DB by [@&#8203;alpe](https://redirect.github.com/alpe) in [#&#8203;2264](https://redirect.github.com/dgraph-io/badger/pull/2264)
- perf: skip lsm lookup for expired entries during value log rewrite by [@&#8203;lamb007](https://redirect.github.com/lamb007) in [#&#8203;2269](https://redirect.github.com/dgraph-io/badger/pull/2269)
- chore: update jemalloc to 5.3.1 by [@&#8203;solracsf](https://redirect.github.com/solracsf) in [#&#8203;2275](https://redirect.github.com/dgraph-io/badger/pull/2275)
- fix : read only user should be able to open db on ReadOnly mode  by [@&#8203;Naelpuissant](https://redirect.github.com/Naelpuissant) in [#&#8203;2268](https://redirect.github.com/dgraph-io/badger/pull/2268)
- fix(compaction): clamp baseLevel to >=1 to prevent L0->L0 panic on la… by [@&#8203;shiva-istari](https://redirect.github.com/shiva-istari) in [#&#8203;2296](https://redirect.github.com/dgraph-io/badger/pull/2296)

#### New Contributors

- [@&#8203;xqqp](https://redirect.github.com/xqqp) made their first contribution in [#&#8203;2257](https://redirect.github.com/dgraph-io/badger/pull/2257)
- [@&#8203;alpe](https://redirect.github.com/alpe) made their first contribution in [#&#8203;2264](https://redirect.github.com/dgraph-io/badger/pull/2264)
- [@&#8203;lamb007](https://redirect.github.com/lamb007) made their first contribution in [#&#8203;2269](https://redirect.github.com/dgraph-io/badger/pull/2269)
- [@&#8203;Naelpuissant](https://redirect.github.com/Naelpuissant) made their first contribution in [#&#8203;2268](https://redirect.github.com/dgraph-io/badger/pull/2268)
- [@&#8203;shiva-istari](https://redirect.github.com/shiva-istari) made their first contribution in [#&#8203;2296](https://redirect.github.com/dgraph-io/badger/pull/2296)

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.9.1...v4.9.2>

### [`v4.9.1`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#491---2026-02-04)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.9.0...v4.9.1)

**Fixed**

- fix(aix): add aix directory synchronization support ([#&#8203;2115](https://redirect.github.com/dgraph-io/badger/issues/2115))
- fix: correct the comment on value size in skl.node ([#&#8203;2250](https://redirect.github.com/dgraph-io/badger/issues/2250))

**Tests**

- test: add checksum tests for package y ([#&#8203;2246](https://redirect.github.com/dgraph-io/badger/issues/2246))

**Chores**

- chore(ci): update arm runner label ([#&#8203;2248](https://redirect.github.com/dgraph-io/badger/issues/2248))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.9.0...v4.9.1>

### [`v4.9.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#491---2026-02-04)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0)

**Fixed**

- fix(aix): add aix directory synchronization support ([#&#8203;2115](https://redirect.github.com/dgraph-io/badger/issues/2115))
- fix: correct the comment on value size in skl.node ([#&#8203;2250](https://redirect.github.com/dgraph-io/badger/issues/2250))

**Tests**

- test: add checksum tests for package y ([#&#8203;2246](https://redirect.github.com/dgraph-io/badger/issues/2246))

**Chores**

- chore(ci): update arm runner label ([#&#8203;2248](https://redirect.github.com/dgraph-io/badger/issues/2248))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.9.0...v4.9.1>

### [`v4.8.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#490---2025-12-15)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0)

**Fixed**

- fix(y): y.SafeCopy shall always return empty slice rather than nil ([#&#8203;2245](https://redirect.github.com/dgraph-io/badger/issues/2245))
  > **WARNING** SafeCopy now returns an empty slice rather than nil. For those using our `y` utility
  > package, this could be a breaking change. This has implications for empty slices stored in
  > badger, specifically, upon retrieval the value stored with the key will be equal to what was set
  > (an empty \[]byte). See [#&#8203;2067](https://redirect.github.com/dgraph-io/badger/issues/2067) for more details.
- fix: test.sh error ([#&#8203;2225](https://redirect.github.com/dgraph-io/badger/issues/2225))
- fix: typo of abandoned ([#&#8203;2222](https://redirect.github.com/dgraph-io/badger/issues/2222))

**Docs**

- add doc for encryption at rest ([#&#8203;2240](https://redirect.github.com/dgraph-io/badger/issues/2240))
- move docs pages in the repo ([#&#8203;2232](https://redirect.github.com/dgraph-io/badger/issues/2232))

**Chores**

- chore(ci): restrict Dgraph test to core packages only ([#&#8203;2242](https://redirect.github.com/dgraph-io/badger/issues/2242))
- chore: update README.md with correct links and badges ([#&#8203;2239](https://redirect.github.com/dgraph-io/badger/issues/2239))
- chore: change renovate to maintain backwards compatible go version ([#&#8203;2236](https://redirect.github.com/dgraph-io/badger/issues/2236))
- chore: configure renovate to leave go version as declared ([#&#8203;2235](https://redirect.github.com/dgraph-io/badger/issues/2235))
- chore(deps): Update actions (major) ([#&#8203;2229](https://redirect.github.com/dgraph-io/badger/issues/2229))
- chore(deps): Update actions/checkout action to v5 ([#&#8203;2221](https://redirect.github.com/dgraph-io/badger/issues/2221))
- chore(deps): Update go minor and patch ([#&#8203;2218](https://redirect.github.com/dgraph-io/badger/issues/2218))
- chore: update the trunk conf file ([#&#8203;2217](https://redirect.github.com/dgraph-io/badger/issues/2217))
- chore(deps): Update dependency node to v22 ([#&#8203;2219](https://redirect.github.com/dgraph-io/badger/issues/2219))
- chore(deps): Update go minor and patch ([#&#8203;2212](https://redirect.github.com/dgraph-io/badger/issues/2212))

**CI**

- move to GitHub Actions runners

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.8.0...v4.8.1>

### [`v4.7.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#480---2025-07-15)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0)

**Features**

- feat(stream): Update stream framework with new alternate keyToList function ([#&#8203;2211](https://redirect.github.com/dgraph-io/badger/issues/2211))

**Fixed**

- fix: crash loop on missing manifest tables ([#&#8203;2198](https://redirect.github.com/dgraph-io/badger/issues/2198))

**Chores**

- chore(deps): Update module golang.org/x/sys to v0.34.0 ([#&#8203;2210](https://redirect.github.com/dgraph-io/badger/issues/2210))
- chore(deps): Update go minor and patch ([#&#8203;2208](https://redirect.github.com/dgraph-io/badger/issues/2208))
- chore(deps): Update go minor and patch ([#&#8203;2204](https://redirect.github.com/dgraph-io/badger/issues/2204))
- chore(deps): Update go minor and patch ([#&#8203;2202](https://redirect.github.com/dgraph-io/badger/issues/2202))
- chore(deps): Update go minor and patch ([#&#8203;2200](https://redirect.github.com/dgraph-io/badger/issues/2200))
- chore(deps): Update module golang.org/x/sys to v0.33.0 ([#&#8203;2195](https://redirect.github.com/dgraph-io/badger/issues/2195))
- chore(deps): Update go minor and patch ([#&#8203;2189](https://redirect.github.com/dgraph-io/badger/issues/2189))
- Compile with jemalloc v5.3.0 ([#&#8203;2191](https://redirect.github.com/dgraph-io/badger/issues/2191))

**CI**

- Update trunk.yml
- move Trunk to action

**Docs**

- docs: add new badge ([#&#8203;2194](https://redirect.github.com/dgraph-io/badger/issues/2194))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0>

### [`v4.6.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#470---2025-04-08)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0)

**Chores**

- chore(deps): remove dependency on github.com/pkg/errors ([#&#8203;2184](https://redirect.github.com/dgraph-io/badger/issues/2184))
- chore(deps): Update go minor and patch ([#&#8203;2187](https://redirect.github.com/dgraph-io/badger/issues/2187))
- chore(deps): Update go minor and patch ([#&#8203;2181](https://redirect.github.com/dgraph-io/badger/issues/2181))
- chore(deps): Update module golang.org/x/sys to v0.31.0 ([#&#8203;2179](https://redirect.github.com/dgraph-io/badger/issues/2179))

**Fixed**

- fix broken badge ([#&#8203;2186](https://redirect.github.com/dgraph-io/badger/issues/2186))

**Docs**

- Update README.md
- doc: add Blink Labs projects to the using Badger list ([#&#8203;2183](https://redirect.github.com/dgraph-io/badger/issues/2183))
- doc: add FlowG to "Projects Using Badger" section of the README ([#&#8203;2180](https://redirect.github.com/dgraph-io/badger/issues/2180))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0>

### [`v4.5.2`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#460---2025-02-26)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2)

**Chores**

- chore(deps): Migrate from OpenCensus to OpenTelemetry ([#&#8203;2169](https://redirect.github.com/dgraph-io/badger/issues/2169))
- chore(deps): Update go minor and patch ([#&#8203;2177](https://redirect.github.com/dgraph-io/badger/issues/2177))
- chore(deps): Update module github.com/spf13/cobra to v1.9.0 ([#&#8203;2174](https://redirect.github.com/dgraph-io/badger/issues/2174))
- chore: add editor config
- update .gitignore ([#&#8203;2176](https://redirect.github.com/dgraph-io/badger/issues/2176))

**Fixed**

- fix: remove accidentally uploaded binary `badger-darwin-arm64` ([#&#8203;2175](https://redirect.github.com/dgraph-io/badger/issues/2175))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0>

### [`v4.5.1`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#452---2025-02-14)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1)

**Chores**

- chore(deps): Update go minor and patch ([#&#8203;2168](https://redirect.github.com/dgraph-io/badger/issues/2168))
- chore(deps): bump minimum Go support to 1.22 ([#&#8203;2171](https://redirect.github.com/dgraph-io/badger/issues/2171))
- chore: migrate docs to centralized docs repo ([#&#8203;2166](https://redirect.github.com/dgraph-io/badger/issues/2166))
- chore: align repo conventions ([#&#8203;2158](https://redirect.github.com/dgraph-io/badger/issues/2158))
- chore(deps): bump the patch group with 2 updates ([#&#8203;2156](https://redirect.github.com/dgraph-io/badger/issues/2156))
- chore(deps): bump github.com/google/flatbuffers from 24.12.23+incompatible to 25.1.21+incompatible
  ([#&#8203;2153](https://redirect.github.com/dgraph-io/badger/issues/2153))
- chore(deps): bump golangci/golangci-lint-action from 6.1.1 to 6.2.0 in the actions group ([#&#8203;2154](https://redirect.github.com/dgraph-io/badger/issues/2154))
- Update renovate.json
- Update trunk.yaml
- enable Trivy

**Fixed**

- update docs link in error message ([#&#8203;2170](https://redirect.github.com/dgraph-io/badger/issues/2170))
- Revert "Update badgerpb4.pb.go" ([#&#8203;2172](https://redirect.github.com/dgraph-io/badger/issues/2172))

**Docs**

- Update README.md
- Added my project that uses Badger database ([#&#8203;2157](https://redirect.github.com/dgraph-io/badger/issues/2157))
- Create SECURITY.md

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2>

### [`v4.5.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#451---2025-01-21)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0)

- chore(deps): bump google.golang.org/protobuf from 1.36.2 to 1.36.3 in the patch group ([#&#8203;2150](https://redirect.github.com/dgraph-io/badger/issues/2150))
- bump github.com/dgraph-io/ristretto/v2 from 2.0.1 to 2.1.0 in the minor group ([#&#8203;2151](https://redirect.github.com/dgraph-io/badger/issues/2151))
- feat(info): print total size of listed keys ([#&#8203;2149](https://redirect.github.com/dgraph-io/badger/issues/2149))
- chore(deps): bump google.golang.org/protobuf from 1.36.1 to 1.36.2 in the patch group ([#&#8203;2146](https://redirect.github.com/dgraph-io/badger/issues/2146))
- chore(deps): bump the minor group with 2 updates ([#&#8203;2147](https://redirect.github.com/dgraph-io/badger/issues/2147))
- fix(info): print Total BloomFilter Size with totalBloomFilter instead of totalIndex ([#&#8203;2145](https://redirect.github.com/dgraph-io/badger/issues/2145))
- chore(deps): bump the minor group with 2 updates ([#&#8203;2141](https://redirect.github.com/dgraph-io/badger/issues/2141))
- chore(deps): bump google.golang.org/protobuf from 1.36.0 to 1.36.1 in the patch group ([#&#8203;2140](https://redirect.github.com/dgraph-io/badger/issues/2140))
- chore(deps): bump google.golang.org/protobuf from 1.35.2 to 1.36.0 in the minor group ([#&#8203;2139](https://redirect.github.com/dgraph-io/badger/issues/2139))
- chore(deps): bump github.com/dgraph-io/ristretto/v2 from 2.0.0 to 2.0.1 in the patch group ([#&#8203;2136](https://redirect.github.com/dgraph-io/badger/issues/2136))
- chore(deps): bump golang.org/x/net from 0.31.0 to 0.32.0 in the minor group ([#&#8203;2137](https://redirect.github.com/dgraph-io/badger/issues/2137))
- chore(deps): bump the minor group with 2 updates ([#&#8203;2135](https://redirect.github.com/dgraph-io/badger/issues/2135))
- docs: Add pagination explanation to docs ([#&#8203;2134](https://redirect.github.com/dgraph-io/badger/issues/2134))
- Fix build for GOARCH=wasm with GOOS=js or GOOS=wasip1 ([#&#8203;2048](https://redirect.github.com/dgraph-io/badger/issues/2048))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1>

### [`v4.4.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#450---2024-11-29)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0)

- fix the cd pipeline by [@&#8203;mangalaman93](https://redirect.github.com/mangalaman93) in [#&#8203;2127](https://redirect.github.com/dgraph-io/badger/pull/2127)
- chore(deps): bump the minor group with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2128](https://redirect.github.com/dgraph-io/badger/pull/2128)
- chore(deps): bump github.com/stretchr/testify from 1.9.0 to 1.10.0 in the minor group by
  [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2130](https://redirect.github.com/dgraph-io/badger/pull/2130)
- upgrade protobuf library by [@&#8203;shivaji-kharse](https://redirect.github.com/shivaji-kharse) in [#&#8203;2131](https://redirect.github.com/dgraph-io/badger/pull/2131)

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0>

### [`v4.3.1`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#440---2024-10-26)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1)

- retract v4.3.0 due to [#&#8203;2121](https://redirect.github.com/dgraph-io/badger/issues/2121) and [#&#8203;2113](https://redirect.github.com/dgraph-io/badger/issues/2113), upgrade to Go v1.23, use ristretto v2 in
  [#&#8203;2122](https://redirect.github.com/dgraph-io/badger/pull/2122)
- Allow stream custom maxsize per batch in [#&#8203;2063](https://redirect.github.com/dgraph-io/badger/pull/2063)
- chore(deps): bump github.com/klauspost/compress from 1.17.10 to 1.17.11 in the patch group in
  [#&#8203;2120](https://redirect.github.com/dgraph-io/badger/pull/2120)
- fix: sentinel errors should not have stack traces in [#&#8203;2042](https://redirect.github.com/dgraph-io/badger/pull/2042)
- chore(deps): bump the minor group with 2 updates in [#&#8203;2119](https://redirect.github.com/dgraph-io/badger/pull/2119)

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0>

### [`v4.3.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#431---2024-10-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0)

- chore: update docs links by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in [#&#8203;2097](https://redirect.github.com/dgraph-io/badger/pull/2097)
- chore(deps): bump golang.org/x/sys from 0.24.0 to 0.25.0 in the minor group by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2100](https://redirect.github.com/dgraph-io/badger/pull/2100)
- chore(deps): bump golang.org/x/net from 0.28.0 to 0.29.0 in the minor group by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2106](https://redirect.github.com/dgraph-io/badger/pull/2106)
- fix: fix reverse iterator broken by seek by [@&#8203;harshil-goel](https://redirect.github.com/harshil-goel) in
  [#&#8203;2109](https://redirect.github.com/dgraph-io/badger/pull/2109)
- chore(deps): bump github.com/klauspost/compress from 1.17.9 to 1.17.10 in the patch group by
  [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2114](https://redirect.github.com/dgraph-io/badger/pull/2114)
- chore(deps): bump github.com/dgraph-io/ristretto from 0.1.2-0.20240116140435-c67e07994f91 to 1.0.0
  by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2112](https://redirect.github.com/dgraph-io/badger/pull/2112)

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1>

### [`v4.2.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#430---2024-08-29)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0)

> **Warning** The tag v4.3.0 has been retracted due to an issue go.sum. Use v4.3.1 (see [#&#8203;2121](https://redirect.github.com/dgraph-io/badger/issues/2121) and
> [#&#8203;2113](https://redirect.github.com/dgraph-io/badger/issues/2113))

**Fixes**

- chore(changelog): add a missed entry in CHANGELOG for v4.2.0 by [@&#8203;mangalaman93](https://redirect.github.com/mangalaman93) in
  [#&#8203;1988](https://redirect.github.com/dgraph-io/badger/pull/1988)
- update README with project KVS using badger by [@&#8203;tauraamui](https://redirect.github.com/tauraamui) in
  [#&#8203;1989](https://redirect.github.com/dgraph-io/badger/pull/1989)
- fix edge case for watermark when index is zero by [@&#8203;mangalaman93](https://redirect.github.com/mangalaman93) in
  [#&#8203;1999](https://redirect.github.com/dgraph-io/badger/pull/1999)
- upgrade spf13/cobra to version v1.7.0 by [@&#8203;mangalaman93](https://redirect.github.com/mangalaman93) in
  [#&#8203;2001](https://redirect.github.com/dgraph-io/badger/pull/2001)
- chore: update readme by [@&#8203;joshua-goldstein](https://redirect.github.com/joshua-goldstein) in [#&#8203;2011](https://redirect.github.com/dgraph-io/badger/pull/2011)
- perf: upgrade compress package test and benchmark. by [@&#8203;siddhant2001](https://redirect.github.com/siddhant2001) in
  [#&#8203;2009](https://redirect.github.com/dgraph-io/badger/pull/2009)
- fix(Transactions): Fix resource consumption on empty write transaction by [@&#8203;Zach-Johnson](https://redirect.github.com/Zach-Johnson) in
  [#&#8203;2018](https://redirect.github.com/dgraph-io/badger/pull/2018)
- chore(deps): bump golang.org/x/net from 0.7.0 to 0.17.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2017](https://redirect.github.com/dgraph-io/badger/pull/2017)
- perf(compactor): optimize allocations: use buffer for priorities by [@&#8203;deff7](https://redirect.github.com/deff7) in
  [#&#8203;2006](https://redirect.github.com/dgraph-io/badger/pull/2006)
- fix(Transaction): discard empty transactions on CommitWith by [@&#8203;Wondertan](https://redirect.github.com/Wondertan) in
  [#&#8203;2031](https://redirect.github.com/dgraph-io/badger/pull/2031)
- fix(levelHandler): use lock for levelHandler sort tables instead of rlock by [@&#8203;xgzlucario](https://redirect.github.com/xgzlucario) in
  [#&#8203;2034](https://redirect.github.com/dgraph-io/badger/pull/2034)
- Docs: update README with project LLS using badger by [@&#8203;Boc-chi-no](https://redirect.github.com/Boc-chi-no) in
  [#&#8203;2032](https://redirect.github.com/dgraph-io/badger/pull/2032)
- chore: MaxTableSize has been renamed to BaseTableSize by [@&#8203;mitar](https://redirect.github.com/mitar) in
  [#&#8203;2038](https://redirect.github.com/dgraph-io/badger/pull/2038)
- Update CODEOWNERS by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in [#&#8203;2043](https://redirect.github.com/dgraph-io/badger/pull/2043)
- Chore(): add Stale Action by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in [#&#8203;2070](https://redirect.github.com/dgraph-io/badger/pull/2070)
- Update ristretto and refactor for use of generics by [@&#8203;paralin](https://redirect.github.com/paralin) in
  [#&#8203;2047](https://redirect.github.com/dgraph-io/badger/pull/2047)
- chore: Remove obsolete comment by [@&#8203;mitar](https://redirect.github.com/mitar) in [#&#8203;2039](https://redirect.github.com/dgraph-io/badger/pull/2039)
- chore(Docs): Update jQuery 3.2.1 to 3.7.1 by [@&#8203;kokizzu](https://redirect.github.com/kokizzu) in
  [#&#8203;2023](https://redirect.github.com/dgraph-io/badger/pull/2023)
- chore(deps): bump the go\_modules group with 3 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2074](https://redirect.github.com/dgraph-io/badger/pull/2074)
- docs(): update docs path by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in [#&#8203;2076](https://redirect.github.com/dgraph-io/badger/pull/2076)
- perf: fix operation in seek by [@&#8203;harshil-goel](https://redirect.github.com/harshil-goel) in [#&#8203;2077](https://redirect.github.com/dgraph-io/badger/pull/2077)
- Add lakeFS to README.md by [@&#8203;N-o-Z](https://redirect.github.com/N-o-Z) in [#&#8203;2078](https://redirect.github.com/dgraph-io/badger/pull/2078)
- chore(): add Dependabot by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in [#&#8203;2080](https://redirect.github.com/dgraph-io/badger/pull/2080)
- chore(deps): bump golangci/golangci-lint-action from 4 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2083](https://redirect.github.com/dgraph-io/badger/pull/2083)
- chore(deps): bump actions/upload-artifact from 3 to 4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2081](https://redirect.github.com/dgraph-io/badger/pull/2081)
- chore(deps): bump github/codeql-action from 2 to 3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2082](https://redirect.github.com/dgraph-io/badger/pull/2082)
- chore(deps): bump the minor group with 7 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2089](https://redirect.github.com/dgraph-io/badger/pull/2089)
- Action Manager by [@&#8203;madhu72](https://redirect.github.com/madhu72) in [#&#8203;2050](https://redirect.github.com/dgraph-io/badger/pull/2050)
- chore(deps): bump golang.org/x/sys from 0.23.0 to 0.24.0 in the minor group by [@&#8203;dependabot](https://redirect.github.com/dependabot) in
  [#&#8203;2091](https://redirect.github.com/dgraph-io/badger/pull/2091)
- chore(deps): bump github.com/golang/protobuf from 1.5.3 to 1.5.4 in the patch group by [@&#8203;dependabot](https://redirect.github.com/dependabot)
  in [#&#8203;2090](https://redirect.github.com/dgraph-io/badger/pull/2090)
- chore: fix some comments by [@&#8203;dufucun](https://redirect.github.com/dufucun) in [#&#8203;2092](https://redirect.github.com/dgraph-io/badger/pull/2092)
- chore(deps): bump github.com/google/flatbuffers from 1.12.1 to 24.3.25+incompatible by [@&#8203;dependabot](https://redirect.github.com/dependabot)
  in [#&#8203;2084](https://redirect.github.com/dgraph-io/badger/pull/2084)

**CI**

- ci: change cron frequency to fix ghost jobs by [@&#8203;joshua-goldstein](https://redirect.github.com/joshua-goldstein) in
  [#&#8203;2010](https://redirect.github.com/dgraph-io/badger/pull/2010)
- fix(CI): Update to pull\_request trigger by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in
  [#&#8203;2056](https://redirect.github.com/dgraph-io/badger/pull/2056)
- ci/cd optimization by [@&#8203;ryanfoxtyler](https://redirect.github.com/ryanfoxtyler) in [#&#8203;2051](https://redirect.github.com/dgraph-io/badger/pull/2051)
- fix(cd): fixed cd pipeline by [@&#8203;harshil-goel](https://redirect.github.com/harshil-goel) in [#&#8203;2093](https://redirect.github.com/dgraph-io/badger/pull/2093)
- fix(cd): change name by [@&#8203;harshil-goel](https://redirect.github.com/harshil-goel) in [#&#8203;2094](https://redirect.github.com/dgraph-io/badger/pull/2094)
- fix(cd): added more debug things to cd by [@&#8203;harshil-goel](https://redirect.github.com/harshil-goel) in
  [#&#8203;2095](https://redirect.github.com/dgraph-io/badger/pull/2095)
- fix(cd): removing some debug items by [@&#8203;harshil-goel](https://redirect.github.com/harshil-goel) in
  [#&#8203;2096](https://redirect.github.com/dgraph-io/badger/pull/2096)

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0>

### [`v4.1.0`](https://redirect.github.com/dgraph-io/badger/releases/tag/v4.1.0): Badger v4.1.0

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0)

This release adds support for incremental stream writer.  We also do some cleanup in the docs and resolve some CI issues for community PR's.  We resolve high and medium CVE's and fix [#&#8203;1833](https://redirect.github.com/dgraph-io/badger/issues/1833).

##### Features

- feat(stream): add support for incremental stream writer ([#&#8203;1722](https://redirect.github.com/dgraph-io/badger/issues/1722)) ([#&#8203;1874](https://redirect.github.com/dgraph-io/badger/issues/1874))

##### Fixes

- chore: upgrade xxhash from v1.1.0 to v2.1.2 ([#&#8203;1910](https://redirect.github.com/dgraph-io/badger/issues/1910)) (fixes [#&#8203;1833](https://redirect.github.com/dgraph-io/badger/issues/1833))

##### Security

- chore(deps): bump golang.org/x/net from 0.0.0-20201021035429-f5854403a974 to 0.7.0 ([#&#8203;1885](https://redirect.github.com/dgraph-io/badger/issues/1885))

##### CVE's

- [CVE-2021-31525](https://redirect.github.com/dgraph-io/badger/security/dependabot/7)
- [CVE-2022-41723](https://redirect.github.com/dgraph-io/badger/security/dependabot/4)
- [CVE-2022-27664](https://redirect.github.com/dgraph-io/badger/security/dependabot/5)
- [CVE-2021-33194](https://redirect.github.com/dgraph-io/badger/security/dependabot/9)
- [CVE-2022-41723](https://redirect.github.com/dgraph-io/badger/security/dependabot/13)
- [CVE-2021-33194](https://redirect.github.com/dgraph-io/badger/security/dependabot/16)
- [CVE-2021-38561](https://redirect.github.com/dgraph-io/badger/security/dependabot/8)

##### Chores

- fix(docs): update README ([#&#8203;1915](https://redirect.github.com/dgraph-io/badger/issues/1915))
- cleanup sstable file after tests ([#&#8203;1912](https://redirect.github.com/dgraph-io/badger/issues/1912))
- chore(ci): add dgraph regression tests ([#&#8203;1908](https://redirect.github.com/dgraph-io/badger/issues/1908))
- docs: fix the default value in docs ([#&#8203;1909](https://redirect.github.com/dgraph-io/badger/issues/1909))
- chore: update URL for unsupported manifest version error ([#&#8203;1905](https://redirect.github.com/dgraph-io/badger/issues/1905))
- docs(README): add raft-badger to projects using badger ([#&#8203;1902](https://redirect.github.com/dgraph-io/badger/issues/1902))
- sync the docs with README with projects using badger ([#&#8203;1903](https://redirect.github.com/dgraph-io/badger/issues/1903))
- fix: update code comments for WithNumCompactors ([#&#8203;1900](https://redirect.github.com/dgraph-io/badger/issues/1900))
- docs: add loggie to projects using badger ([#&#8203;1882](https://redirect.github.com/dgraph-io/badger/issues/1882))
- chore(memtable): refactor code for memtable flush ([#&#8203;1866](https://redirect.github.com/dgraph-io/badger/issues/1866))
- resolve coveralls issue for community PR's ([#&#8203;1892](https://redirect.github.com/dgraph-io/badger/issues/1892), [#&#8203;1894](https://redirect.github.com/dgraph-io/badger/issues/1894), [#&#8203;1896](https://redirect.github.com/dgraph-io/badger/issues/1896))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0>

### [`v4.0.1`](https://redirect.github.com/dgraph-io/badger/releases/tag/v4.0.1): Badger DB v4.0.1

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1)

This release fixes a bug in the maxHeaderSize parameter that could lead to panics. We introduce an external magic number to keep track of external dependencies. We bump up the minimum required Go version to 1.19. No changes were made to the format of data on disk.  This is a major release because we are making a switch to SemVer in order to make it easier for the community to understand when breaking API and data format changes are made.

> **Warning**
> The tag v4.0.0 has been retracted due to a bug in publisher.
> Use v4.0.1 (see [#&#8203;1889](https://redirect.github.com/dgraph-io/badger/issues/1889))

##### Fixed

- fix(pb): fix generated protos [#&#8203;1888](https://redirect.github.com/dgraph-io/badger/issues/1888)
- fix(publisher): initialize the atomic variable [#&#8203;1889](https://redirect.github.com/dgraph-io/badger/issues/1889)
- fix: update maxHeaderSize [#&#8203;1877](https://redirect.github.com/dgraph-io/badger/issues/1877)
- feat(externalMagic): Introduce external magic number ([#&#8203;1745](https://redirect.github.com/dgraph-io/badger/issues/1745)) [#&#8203;1852](https://redirect.github.com/dgraph-io/badger/issues/1852)
- fix(bench): bring in benchmark fixes from main [#&#8203;1863](https://redirect.github.com/dgraph-io/badger/issues/1863)

##### Chores

- upgrade go to 1.19 [#&#8203;1868](https://redirect.github.com/dgraph-io/badger/issues/1868)
- moving from CalVer to SemVer
- chore(cd): tag based deployments [#&#8203;1887](https://redirect.github.com/dgraph-io/badger/issues/1887)
- chore(ci): fail fast when testing [#&#8203;1890](https://redirect.github.com/dgraph-io/badger/issues/1890)
- enable linters (gosimple, govet, lll, unused, staticcheck, errcheck, ineffassign, gofmt) [#&#8203;1871](https://redirect.github.com/dgraph-io/badger/issues/1871) [#&#8203;1870](https://redirect.github.com/dgraph-io/badger/issues/1870) [#&#8203;1876](https://redirect.github.com/dgraph-io/badger/issues/1876)
- remove dependency on io/ioutil [#&#8203;1879](https://redirect.github.com/dgraph-io/badger/issues/1879)
- various doc and comment fixes [#&#8203;1857](https://redirect.github.com/dgraph-io/badger/issues/1857)

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.1>

### [`v4.0.0`](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0)

### [`v3.2103.5`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2103.5): Badger DB v3.2103.5

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5)

We release Badger CLI tool binaries for amd64 and now arm64.  This release does not involve any core code changes to Badger.  We added a CD job for building Badger for arm64.

### [`v3.2103.4`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2103.4): Badger DB v3.2103.4

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4)

This patches an issue that could lead to manifest corruption.  Fix was merged in [#&#8203;1756](https://redirect.github.com/dgraph-io/badger/issues/1756).  Addresses [this issue](https://discuss.dgraph.io/t/badgerdb-manifest-corruption-issue-solution/15915) on Discuss and[ this issue](https://redirect.github.com/dgraph-io/badger/issues/1819) on Badger.  We also bring the release branch to parity with main by updating the CI/CD jobs, Readme, Codeowners, PR and issue templates, etc.

##### Fixed

- fix(manifest): fix manifest corruption due to race condition in concurrent compactions ([#&#8203;1756](https://redirect.github.com/dgraph-io/badger/issues/1756))

##### Chores

- Add CI/CD jobs to release branch
- Add PR and Issue templates to release branch
- Update Codeowners in release branch
- Update Readme in release branch

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4>

### [`v3.2103.3`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2103.3): BadgerDB v3.2103.3

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3)

This is a minor patch release that fixes arm64 related issues.  The issues in the `z` package in Ristretto were resolved in Ristretto v0.1.1.

##### Fixed

- fix(arm64): bump ristretto v0.1.0 --> v0.1.1 ([#&#8203;1806](https://redirect.github.com/dgraph-io/badger/issues/1806))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3>

### [`v3.2103.2`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2103.2): BadgerDB v3.2103.2

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2)

This patch release contains:

##### Fixed

- fix(compact): close vlog after the compaction at L0 has been completed ([#&#8203;1752](https://redirect.github.com/dgraph-io/badger/issues/1752))
- fix(builder): put the upper limit on reallocation ([#&#8203;1748](https://redirect.github.com/dgraph-io/badger/issues/1748))
- deps: Bump github.com/google/flatbuffers to v1.12.1 ([#&#8203;1746](https://redirect.github.com/dgraph-io/badger/issues/1746))
- fix(levels): Avoid a deadlock when acquiring read locks in levels ([#&#8203;1744](https://redirect.github.com/dgraph-io/badger/issues/1744))
- fix(pubsub): avoid deadlock in publisher and subscriber ([#&#8203;1749](https://redirect.github.com/dgraph-io/badger/issues/1749)) ([#&#8203;1751](https://redirect.github.com/dgraph-io/badger/issues/1751))

**Full Changelog**: <https://github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2>

### [`v3.2103.1`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2103.1): BadgerDB v3.2103.1

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1)

This release removes CGO dependency opf badger by using Klauspost's ZSTD instead of Datadog's ZSTD. Also, this has some of the fixes.

##### Fixed

- fix(compaction): copy over the file ID when building tables [#&#8203;1713](https://redirect.github.com/dgraph-io/badger/issues/1713)
- fix: Fix conflict detection for managed DB ([#&#8203;1716](https://redirect.github.com/dgraph-io/badger/issues/1716))
- fix(pendingWrites): don't skip the pending entries with version=0 ([#&#8203;1721](https://redirect.github.com/dgraph-io/badger/issues/1721))

##### Features

- feat(zstd): replace datadog's zstd with Klauspost's zstd ([#&#8203;1709](https://redirect.github.com/dgraph-io/badger/issues/1709))

### [`v3.2103.0`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2103.0): BadgerDB v3.2103.0

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0)

##### Breaking

- Subscribe: Add option to subscribe with holes in prefixes. ([#&#8203;1658](https://redirect.github.com/dgraph-io/badger/issues/1658))

##### Fixed

- fix(compaction): Remove compaction backoff mechanism ([#&#8203;1686](https://redirect.github.com/dgraph-io/badger/issues/1686))
- Add a name to mutexes to make them unexported ([#&#8203;1678](https://redirect.github.com/dgraph-io/badger/issues/1678))
- fix(merge-operator): don't read the deleted keys ([#&#8203;1675](https://redirect.github.com/dgraph-io/badger/issues/1675))
- fix(discard): close the discard stats file on db close ([#&#8203;1672](https://redirect.github.com/dgraph-io/badger/issues/1672))
- fix(iterator): fix iterator when data does not exist in read only mode ([#&#8203;1670](https://redirect.github.com/dgraph-io/badger/issues/1670))
- fix(badger): Do not reuse variable across badger commands ([#&#8203;1624](https://redirect.github.com/dgraph-io/badger/issues/1624))
- fix(dropPrefix): check properly if the key is present in a table  ([#&#8203;1623](https://redirect.github.com/dgraph-io/badger/issues/1623))

##### Performance

- Opt(Stream): Optimize how we deduce key ranges for iteration ([#&#8203;1687](https://redirect.github.com/dgraph-io/badger/issues/1687))
- Increase value threshold from 1 KB to 1 MB ([#&#8203;1664](https://redirect.github.com/dgraph-io/badger/issues/1664))
- opt(DropPrefix): check if there exist some data to drop before dropping prefixes ([#&#8203;1621](https://redirect.github.com/dgraph-io/badger/issues/1621))

##### Features

- feat(options): allow special handling and checking when creating options from superflag ([#&#8203;1688](https://redirect.github.com/dgraph-io/badger/issues/1688))
- overwrite default Options from SuperFlag string ([#&#8203;1663](https://redirect.github.com/dgraph-io/badger/issues/1663))
- Support SinceTs in iterators ([#&#8203;1653](https://redirect.github.com/dgraph-io/badger/issues/1653))
- feat(info): Add a flag to parse and print DISCARD file ([#&#8203;1662](https://redirect.github.com/dgraph-io/badger/issues/1662))
- feat(vlog): making vlog threshold dynamic [`6ce3b7c`](https://redirect.github.com/dgraph-io/badger/commit/6ce3b7c) ([#&#8203;1635](https://redirect.github.com/dgraph-io/badger/issues/1635))
- feat(options): add NumGoroutines option for default Stream.numGo ([#&#8203;1656](https://redirect.github.com/dgraph-io/badger/issues/1656))
- feat(Trie): Working prefix match with holes ([#&#8203;1654](https://redirect.github.com/dgraph-io/badger/issues/1654))
- feat: add functionality to ban a prefix ([#&#8203;1638](https://redirect.github.com/dgraph-io/badger/issues/1638))
- feat(compaction): Support Lmax to Lmax compaction ([#&#8203;1615](https://redirect.github.com/dgraph-io/badger/issues/1615))

##### New APIs

- Badger.DB
  - BanNamespace
  - BannedNamespaces
  - Ranges
- Badger.Options
  - FromSuperFlag
  - WithNumGoRoutines
  - WithNamespaceOffset
  - WithVLogPercentile
- Badger.Trie
  - AddMatch
  - DeleteMatch
- Badger.Table
  - StaleDataSize
- Badger.Table.Builder
  - AddStaleKey
- Badger.InitDiscardStats

##### Removed APIs

- Badger.DB
  - KeySplits
- Badger.Options
  - SkipVlog

##### Changed APIs

- Badger.DB
  - Subscribe
- Badger.Options
  - WithValueThreshold

### [`v3.2011.1`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2011.1): BadgerDB v3.2011.1

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1)

fix(compaction): Set base level correctly after stream ([#&#8203;1631](https://redirect.github.com/dgraph-io/badger/issues/1631)) ([#&#8203;1651](https://redirect.github.com/dgraph-io/badger/issues/1651))
fix: update ristretto and use filepath ([#&#8203;1649](https://redirect.github.com/dgraph-io/badger/issues/1649)) ([#&#8203;1652](https://redirect.github.com/dgraph-io/badger/issues/1652))
fix(badger): Do not reuse variable across badger commands ([#&#8203;1624](https://redirect.github.com/dgraph-io/badger/issues/1624)) ([#&#8203;1650](https://redirect.github.com/dgraph-io/badger/issues/1650))
fix(build): fix 32-bit build ([#&#8203;1627](https://redirect.github.com/dgraph-io/badger/issues/1627)) ([#&#8203;1646](https://redirect.github.com/dgraph-io/badger/issues/1646))
fix(table): always sync SST to disk ([#&#8203;1625](https://redirect.github.com/dgraph-io/badger/issues/1625)) ([#&#8203;1645](https://redirect.github.com/dgraph-io/badger/issues/1645))

### [`v3.2011.0`](https://redirect.github.com/dgraph-io/badger/releases/tag/v3.2011.0): BadgerDB v3.2011.0

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0)

This release is not backward compatible with Badger v2.x.x

##### Breaking:

- opt(compactions): Improve compaction performance ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
- Change how Badger handles WAL ([#&#8203;1555](https://redirect.github.com/dgraph-io/badger/issues/1555))
- feat(index): Use flatbuffers instead of protobuf ([#&#8203;1546](https://redirect.github.com/dgraph-io/badger/issues/1546))

##### Fixed:

- Fix(GC): Set bits correctly for moved keys ([#&#8203;1619](https://redirect.github.com/dgraph-io/badger/issues/1619))
- Fix(tableBuilding): reduce scope of valuePointer ([#&#8203;1617](https://redirect.github.com/dgraph-io/badger/issues/1617))
- Fix(compaction): fix table size estimation on compaction ([#&#8203;1613](https://redirect.github.com/dgraph-io/badger/issues/1613))
- Fix(OOM): Reuse pb.KVs in Stream ([#&#8203;1609](https://redirect.github.com/dgraph-io/badger/issues/1609))
- Fix race condition in L0StallMs variable ([#&#8203;1605](https://redirect.github.com/dgraph-io/badger/issues/1605))
- Fix(stream): Stop produceKVs on error ([#&#8203;1604](https://redirect.github.com/dgraph-io/badger/issues/1604))
- Fix(skiplist): Remove z.Buffer from skiplist ([#&#8203;1600](https://redirect.github.com/dgraph-io/badger/issues/1600))
- Fix(readonly): fix the file opening mode ([#&#8203;1592](https://redirect.github.com/dgraph-io/badger/issues/1592))
- Fix: Disable CompactL0OnClose by default ([#&#8203;1586](https://redirect.github.com/dgraph-io/badger/issues/1586))
- Fix(compaction): Don't drop data when split overlaps with top tables ([#&#8203;1587](https://redirect.github.com/dgraph-io/badger/issues/1587))
- Fix(subcompaction): Close builder before throttle.Done ([#&#8203;1582](https://redirect.github.com/dgraph-io/badger/issues/1582))
- Fix(table): Add onDisk size ([#&#8203;1569](https://redirect.github.com/dgraph-io/badger/issues/1569))
- Fix(Stream): Only send done markers if told to do so
- Fix(value log GC): Fix a bug which caused value log files to not be GCed.
- Fix segmentation fault when cache sizes are small. ([#&#8203;1552](https://redirect.github.com/dgraph-io/badger/issues/1552))
- Fix(builder): Too many small tables when compression is enabled ([#&#8203;1549](https://redirect.github.com/dgraph-io/badger/issues/1549))
- Fix integer overflow error when building for 386 ([#&#8203;1541](https://redirect.github.com/dgraph-io/badger/issues/1541))
- Fix(writeBatch): Avoid deadlock in commit callback ([#&#8203;1529](https://redirect.github.com/dgraph-io/badger/issues/1529))
- Fix(db): Handle nil logger ([#&#8203;1534](https://redirect.github.com/dgraph-io/badger/issues/1534))
- Fix(maxVersion): Use choosekey instead of KeyToList ([#&#8203;1532](https://redirect.github.com/dgraph-io/badger/issues/1532))
- Fix(Backup/Restore): Keep all versions ([#&#8203;1462](https://redirect.github.com/dgraph-io/badger/issues/1462))
- Fix(build): Fix nocgo builds. ([#&#8203;1493](https://redirect.github.com/dgraph-io/badger/issues/1493))
- Fix(cleanup): Avoid truncating in value.Open on error ([#&#8203;1465](https://redirect.github.com/dgraph-io/badger/issues/1465))
- Fix(compaction): Don't use cache for table compaction ([#&#8203;1467](https://redirect.github.com/dgraph-io/badger/issues/1467))
- Fix(compaction): Use separate compactors for L0, L1 ([#&#8203;1466](https://redirect.github.com/dgraph-io/badger/issues/1466))
- Fix(options): Do not implicitly enable cache ([#&#8203;1458](https://redirect.github.com/dgraph-io/badger/issues/1458))
- Fix(cleanup): Do not close cache before compaction ([#&#8203;1464](https://redirect.github.com/dgraph-io/badger/issues/1464))
- Fix(replay): Update head for LSM entires also ([#&#8203;1456](https://redirect.github.com/dgraph-io/badger/issues/1456))
- fix(levels): Cleanup builder resources on building an empty table ([#&#8203;1414](https://redirect.github.com/dgraph-io/badger/issues/1414))

##### Performance

- perf(GC): Remove move keys ([#&#8203;1539](https://redirect.github.com/dgraph-io/badger/issues/1539))
- Keep the cheaper parts of the index within table struct. ([#&#8203;1608](https://redirect.github.com/dgraph-io/badger/issues/1608))
- Opt(stream): Use z.Buffer to stream data ([#&#8203;1606](https://redirect.github.com/dgraph-io/badger/issues/1606))
- opt(builder): Use z.Allocator for building tables ([#&#8203;1576](https://redirect.github.com/dgraph-io/badger/issues/1576))
- opt(memory): Use z.Calloc for allocating KVList  ([#&#8203;1563](https://redirect.github.com/dgraph-io/badger/issues/1563))
- opt: Small memory usage optimizations ([#&#8203;1562](https://redirect.github.com/dgraph-io/badger/issues/1562))
- KeySplits checks tables and memtables when number of splits is small. ([#&#8203;1544](https://redirect.github.com/dgraph-io/badger/issues/1544))
- perf: Reduce memory usage by better struct packing ([#&#8203;1528](https://redirect.github.com/dgraph-io/badger/issues/1528))
- perf(tableIterator): Don't do next on NewIterator ([#&#8203;1512](https://redirect.github.com/dgraph-io/badger/issues/1512))
- Improvements: Manual Memory allocation via Calloc ([#&#8203;1459](https://redirect.github.com/dgraph-io/badger/issues/1459))
- Various bug fixes: Break up list and run DropAll func ([#&#8203;1439](https://redirect.github.com/dgraph-io/badger/issues/1439))
- Add a limit to the size of the batches sent over a stream. ([#&#8203;1412](https://redirect.github.com/dgraph-io/badger/issues/1412))
- Commit does not panic after Finish, instead returns an error ([#&#8203;1396](https://redirect.github.com/dgraph-io/badger/issues/1396))
- levels: Compaction incorrectly drops some delete markers ([#&#8203;1422](https://redirect.github.com/dgraph-io/badger/issues/1422))
- Remove vlog file if bootstrap, syncDir or mmap fails ([#&#8203;1434](https://redirect.github.com/dgraph-io/badger/issues/1434))

##### Features:

- Use opencensus for tracing ([#&#8203;1566](https://redirect.github.com/dgraph-io/badger/issues/1566))
- Export functions from Key Registry ([#&#8203;1561](https://redirect.github.com/dgraph-io/badger/issues/1561))
- Allow sizes of block and index caches to be updated. ([#&#8203;1551](https://redirect.github.com/dgraph-io/badger/issues/1551))
- Add metric for number of tables being compacted ([#&#8203;1554](https://redirect.github.com/dgraph-io/badger/issues/1554))
- feat(info): Show index and bloom filter size ([#&#8203;1543](https://redirect.github.com/dgraph-io/badger/issues/1543))
- feat(db): Add db.MaxVersion API ([#&#8203;1526](https://redirect.github.com/dgraph-io/badger/issues/1526))
- Expose DB options in Badger. ([#&#8203;1521](https://redirect.github.com/dgraph-io/badger/issues/1521))
- Feature: Add a Calloc based Buffer ([#&#8203;1471](https://redirect.github.com/dgraph-io/badger/issues/1471))
- Add command to stream contents of DB into another DB. ([#&#8203;1463](https://redirect.github.com/dgraph-io/badger/issues/1463))
- Expose NumAlloc metrics via expvar ([#&#8203;1470](https://redirect.github.com/dgraph-io/badger/issues/1470))
- Support fully disabling the bloom filter ([#&#8203;1319](https://redirect.github.com/dgraph-io/badger/issues/1319))
- Add --enc-key flag in badger info tool ([#&#8203;1441](https://redirect.github.com/dgraph-io/badger/issues/1441))

##### New APIs

- Badger.DB
  - CacheMaxCost ([#&#8203;1551](https://redirect.github.com/dgraph-io/badger/issues/1551))
  - Levels ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - LevelsToString  ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - Opts ([#&#8203;1521](https://redirect.github.com/dgraph-io/badger/issues/1521))
- Badger.Options
  - WithBaseLevelSize ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - WithBaseTableSize ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - WithMemTableSize ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
- Badger.KeyRegistry
  - DataKey ([#&#8203;1561](https://redirect.github.com/dgraph-io/badger/issues/1561))
  - LatestDataKey ([#&#8203;1561](https://redirect.github.com/dgraph-io/badger/issues/1561))

##### Removed APIs

- Badger.Options
  - WithKeepL0InMemory ([#&#8203;1555](https://redirect.github.com/dgraph-io/badger/issues/1555))
  - WithLevelOneSize  ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - WithLoadBloomsOnOpen ([#&#8203;1555](https://redirect.github.com/dgraph-io/badger/issues/1555))
  - WithLogRotatesToFlush  ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - WithMaxTableSize ([#&#8203;1574](https://redirect.github.com/dgraph-io/badger/issues/1574))
  - WithTableLoadingMode ([#&#8203;1555](https://redirect.github.com/dgraph-io/badger/issues/1555))
  - WithTruncate ([#&#8203;1555](https://redirect.github.com/dgraph-io/badger/issues/1555))
  - WithValueLogLoadingMode ([#&#8203;1555](https://redirect.github.com/dgraph-io/badger/issues/1555))

### [`v2.2007.4`](https://redirect.github.com/dgraph-io/badger/releases/tag/v2.2007.4): BadgerDB v2.2007.4

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4)

##### Fixed

- Fix build on Plan 9 ([#&#8203;1451](https://redirect.github.com/dgraph-io/badger/issues/1451))  ([#&#8203;1508](https://redirect.github.com/dgraph-io/badger/issues/1508)) ([#&#8203;1738](https://redirect.github.com/dgraph-io/badger/issues/1738))

##### Features

- feat(zstd): backport replacement of DataDog's zstd with Klauspost's zstd ([#&#8203;1736](https://redirect.github.com/dgraph-io/badger/issues/1736))

### [`v2.2007.3`](https://redirect.github.com/dgraph-io/badger/releases/tag/v2.2007.3): BadgerDB v2.2007.3

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3)

This patch release contains:

##### Fixed

- fix(maxVersion): Use choosekey instead of KeyToList ([#&#8203;1532](https://redirect.github.com/dgraph-io/badger/issues/1532)) [#&#8203;1533](https://redirect.github.com/dgraph-io/badger/issues/1533)
- fix(flatten): Add --num\_versions flag ([#&#8203;1518](https://redirect.github.com/dgraph-io/badger/issues/1518)) [#&#8203;1520](https://redirect.github.com/dgraph-io/badger/issues/1520)
- fix(build): Fix integer overflow on 32-bit architectures [#&#8203;1558](https://redirect.github.com/dgraph-io/badger/issues/1558)
- fix(pb): avoid protobuf warning due to common filename ([#&#8203;1519](https://redirect.github.com/dgraph-io/badger/issues/1519))

##### Features

- Add command to stream contents of DB into another DB. ([#&#8203;1486](https://redirect.github.com/dgraph-io/badger/issues/1486))

##### New APIs

- DB.StreamDB
- DB.MaxVersion

### [`v2.2007.2`](https://redirect.github.com/dgraph-io/badger/releases/tag/v2.2007.2): BadgerDB v2.2007.2

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2)

##### Fixed

- Compaction: Use separate compactors for L0, L1 ([#&#8203;1466](https://redirect.github.com/dgraph-io/badger/issues/1466))
- Rework Block and Index cache ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
- Add IsClosed method ([#&#8203;1478](https://redirect.github.com/dgraph-io/badger/issues/1478))
- Cleanup: Avoid truncating in vlog.Open on error ([#&#8203;1465](https://redirect.github.com/dgraph-io/badger/issues/1465))
- Cleanup: Do not close cache before compactions ([#&#8203;1464](https://redirect.github.com/dgraph-io/badger/issues/1464))

##### New APIs

- Badger.DB
  - BlockCacheMetrics ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
  - IndexCacheMetrics ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
- Badger.Option
  - WithBlockCacheSize ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
  - WithIndexCacheSize ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))

##### Removed APIs \[Breaking Changes]

- Badger.DB
  - DataCacheMetrics ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
  - BfCacheMetrics ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
- Badger.Option
  - WithMaxCacheSize ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
  - WithMaxBfCacheSize ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
  - WithKeepBlockIndicesInCache ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))
  - WithKeepBlocksInCache ([#&#8203;1473](https://redirect.github.com/dgraph-io/badger/issues/1473))

### [`v2.2007.1`](https://redirect.github.com/dgraph-io/badger/releases/tag/v2.2007.1): BadgerDB v2.2007.1

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1)

##### Fixed

- Remove vlog file if bootstrap, syncDir or mmap fails ([#&#8203;1434](https://redirect.github.com/dgraph-io/badger/issues/1434))
- levels: Compaction incorrectly drops some delete markers ([#&#8203;1422](https://redirect.github.com/dgraph-io/badger/issues/1422))
- Replay: Update head for LSM entires also ([#&#8203;1456](https://redirect.github.com/dgraph-io/badger/issues/1456))

### [`v2.2007.0`](https://redirect.github.com/dgraph-io/badger/releases/tag/v2.2007.0): BadgerDB v2.2007.0

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0)

##### Fixed

- Add a limit to the size of the batches sent over a stream. ([#&#8203;1412](https://redirect.github.com/dgraph-io/badger/issues/1412))
- Fix Sequence generates duplicate values ([#&#8203;1281](https://redirect.github.com/dgraph-io/badger/issues/1281))
- Fix race condition in DoesNotHave ([#&#8203;1287](https://redirect.github.com/dgraph-io/badger/issues/1287))
- Fail fast if cgo is disabled and compression is ZSTD ([#&#8203;1284](https://redirect.github.com/dgraph-io/badger/issues/1284))
- Proto: make badger/v2 compatible with v1 ([#&#8203;1293](https://redirect.github.com/dgraph-io/badger/issues/1293))
- Proto: Rename dgraph.badger.v2.pb to badgerpb2 ([#&#8203;1314](https://redirect.github.com/dgraph-io/badger/issues/1314))
- Handle duplicates in ManagedWriteBatch ([#&#8203;1315](https://redirect.github.com/dgraph-io/badger/issues/1315))
- Ensure `bitValuePointer` flag is cleared for LSM entry values written to LSM ([#&#8203;1313](https://redirect.github.com/dgraph-io/badger/issues/1313))
- DropPrefix: Return error on blocked writes ([#&#8203;1329](https://redirect.github.com/dgraph-io/badger/issues/1329))
- Confirm `badgerMove` entry required before rewrite ([#&#8203;1302](https://redirect.github.com/dgraph-io/badger/issues/1302))
- Drop move keys when its key prefix is dropped ([#&#8203;1331](https://redirect.github.com/dgraph-io/badger/issues/1331))
- Iterator: Always add key to txn.reads ([#&#8203;1328](https://redirect.github.com/dgraph-io/badger/issues/1328))
- Restore: Account for value size as well ([#&#8203;1358](https://redirect.github.com/dgraph-io/badger/issues/1358))
- Compaction: Expired keys and delete markers are never purged ([#&#8203;1354](https://redirect.github.com/dgraph-io/badger/issues/1354))
- GC: Consider size of value while rewriting ([#&#8203;1357](https://redirect.github.com/dgraph-io/badger/issues/1357))
- Force KeepL0InMemory to be true when InMemory is true ([#&#8203;1375](https://redirect.github.com/dgraph-io/badger/issues/1375))
- Rework DB.DropPrefix ([#&#8

> ✂ **Note**
> 
> PR body was truncated to here.


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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQzLjIxMC4wLXJwbSIsInRhcmdldEJyYW5jaCI6Im1haW4iLCJsYWJlbHMiOltdfQ==-->


---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/77*
