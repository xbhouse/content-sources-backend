---
type: pull_request
number: 79
title: "chore(deps): update module github.com/dgraph-io/badger/v2 to v4"
state: merged
author: red-hat-konflux
created: 2025-11-18T16:51:18Z
updated: 2026-02-05T08:20:07Z
closed: 2026-02-05T08:20:04Z
merged: 2026-02-05T08:20:04Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-dgraph-io-badger-v2-4.x
labels: ["dependencies"]
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/79
---

# Pull Request #79: chore(deps): update module github.com/dgraph-io/badger/v2 to v4

**Author**: @red-hat-konflux
**Created**: November 18, 2025 at 04:51 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-dgraph-io-badger-v2-4.x`

## Description

> **Note:** This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/dgraph-io/badger/v2](https://redirect.github.com/dgraph-io/badger) | `v2.2007.4` -> `v4.9.1` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fdgraph-io%2fbadger%2fv2/v4.9.1?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fdgraph-io%2fbadger%2fv2/v2.2007.4/v4.9.1?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>dgraph-io/badger (github.com/dgraph-io/badger/v2)</summary>

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

### [`v4.1.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v4.0.1`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v4.0.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v3.2103.5`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v3.2103.4`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v3.2103.3`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v3.2103.2`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v3.2103.1`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.1...v2.0.2

[2.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.0...v2.0.1

[2.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.6.0...v2.0.0

[1.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.5...v1.6.0

[1.5.5]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.3...v1.5.5

[1.5.3]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.2...v1.5.3

[1.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.1...v1.5.2

[1.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.5.0...v1.5.1

[1.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.4.0...v1.5.0

[1.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.3.0...v1.4.0

[1.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.2.0...v1.3.0

[1.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.1...v1.2.0

[1.1.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.1.0...v1.1.1

[1.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.1...v1.1.0

[1.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v1.0.0...v1.0.1

### [`v3.2103.0`](https://redirect.github.com/dgraph-io/badger/blob/HEAD/CHANGELOG.md#101---2017-11-06)

[Compare Source](https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0)

- Fix an uint16 overflow when resizing key slice

[4.9.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.8.0...v4.9.0

[4.8.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.7.0...v4.8.0

[4.7.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.6.0...v4.7.0

[4.6.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.2...v4.6.0

[4.5.2]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.1...v4.5.2

[4.5.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.5.0...v4.5.1

[4.5.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.4.0...v4.5.0

[4.4.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.1...v4.4.0

[4.3.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.3.0...v4.3.1

[4.3.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.2.0...v4.3.0

[4.2.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.1.0...v4.2.0

[4.1.0]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.1...v4.1.0

[4.0.1]: https://redirect.github.com/dgraph-io/badger/compare/v4.0.0...v4.0.1

[4.0.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.5...v4.0.0

[3.2103.5]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.4...v3.2103.5

[3.2103.4]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.3...v3.2103.4

[3.2103.3]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.2...v3.2103.3

[3.2103.2]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.1...v3.2103.2

[3.2103.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2103.0...v3.2103.1

[3.2103.0]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.1...v3.2103.0

[3.2011.1]: https://redirect.github.com/dgraph-io/badger/compare/v3.2011.0...v3.2011.1

[3.2011.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.4...v3.2011.0

[2.2007.4]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.3...v2.2007.4

[2.2007.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.2...v2.2007.3

[2.2007.2]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.1...v2.2007.2

[2.2007.1]: https://redirect.github.com/dgraph-io/badger/compare/v2.2007.0...v2.2007.1

[2.2007.0]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.3...v2.2007.0

[2.0.3]: https://redirect.github.com/dgraph-io/badger/compare/v2.0.2...v2.0.3

[2.0.2]: https://redirect.github.com/dgraph-io/badger/co

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6WyJkZXBlbmRlbmNpZXMiXX0=-->


---

## Discussion

### Comment by @red-hat-konflux on November 18, 2025 at 04:51 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 1 additional dependency was updated


Details:


| **Package**                     | **Change**              |
| :------------------------------ | :---------------------- |
| `github.com/klauspost/compress` | `v1.17.11` -> `v1.18.0` |

---

## Reviews

### Review by @Hyperkid123 - Approved on February 05, 2026 at 08:19 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/79*
