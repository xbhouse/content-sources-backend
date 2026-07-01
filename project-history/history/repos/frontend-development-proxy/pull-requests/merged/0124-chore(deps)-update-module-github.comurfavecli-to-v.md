---
type: pull_request
number: 124
title: "chore(deps): update module github.com/urfave/cli to v3"
state: merged
author: red-hat-konflux
created: 2026-03-16T17:22:39Z
updated: 2026-03-17T10:40:23Z
closed: 2026-03-17T10:40:21Z
merged: 2026-03-17T10:40:21Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-urfave-cli-3.x
labels: ["dependencies"]
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/124
---

# Pull Request #124: chore(deps): update module github.com/urfave/cli to v3

**Author**: @red-hat-konflux
**Created**: March 16, 2026 at 05:22 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-urfave-cli-3.x`

## Description

> **Note:** This PR body was truncated due to platform limits.

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/urfave/cli](https://redirect.github.com/urfave/cli) | `v1.22.17` -> `v3.7.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2furfave%2fcli/v3.7.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2furfave%2fcli/v1.22.17/v3.7.0?slim=true) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>urfave/cli (github.com/urfave/cli)</summary>

### [`v3.7.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.7.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.6.2...v3.7.0)

#### What's Changed

- Fix: use the correct type name in the tracef message by [@&#8203;icholy](https://redirect.github.com/icholy) in [#&#8203;2251](https://redirect.github.com/urfave/cli/pull/2251)
- chore(deps): bump mkdocs-git-revision-date-localized-plugin from 1.5.0 to 1.5.1 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2252](https://redirect.github.com/urfave/cli/pull/2252)
- Fix:(issue\_2254) Fix incorrect handling of arg after short option token by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2255](https://redirect.github.com/urfave/cli/pull/2255)
- feat: ShellComplete for fish by [@&#8203;marcusramberg](https://redirect.github.com/marcusramberg) in [#&#8203;2256](https://redirect.github.com/urfave/cli/pull/2256)
- Fix: propagate MutuallyExclusiveFlags persistent flags to subcommands by [@&#8203;siutsin](https://redirect.github.com/siutsin) in [#&#8203;2266](https://redirect.github.com/urfave/cli/pull/2266)
- feat: support dynamic fish completion by [@&#8203;Maks1mS](https://redirect.github.com/Maks1mS) in [#&#8203;2270](https://redirect.github.com/urfave/cli/pull/2270)
- fix(help): show GLOBAL OPTIONS for leaf subcommands when HideHelpCommand is true by [@&#8203;TimSoethout](https://redirect.github.com/TimSoethout) in [#&#8203;2269](https://redirect.github.com/urfave/cli/pull/2269)

#### New Contributors

- [@&#8203;marcusramberg](https://redirect.github.com/marcusramberg) made their first contribution in [#&#8203;2256](https://redirect.github.com/urfave/cli/pull/2256)
- [@&#8203;siutsin](https://redirect.github.com/siutsin) made their first contribution in [#&#8203;2266](https://redirect.github.com/urfave/cli/pull/2266)
- [@&#8203;TimSoethout](https://redirect.github.com/TimSoethout) made their first contribution in [#&#8203;2269](https://redirect.github.com/urfave/cli/pull/2269)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.6.2...v3.7.0>

### [`v3.6.2`](https://redirect.github.com/urfave/cli/releases/tag/v3.6.2)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.6.1...v3.6.2)

#### What's Changed

- chore(deps): bump actions/checkout from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2230](https://redirect.github.com/urfave/cli/pull/2230)
- chore(deps): bump mkdocs-material from 9.6.23 to 9.7.0 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2231](https://redirect.github.com/urfave/cli/pull/2231)
- Improve test coverage by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2236](https://redirect.github.com/urfave/cli/pull/2236)
- Add more tests to improve code coverage by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2237](https://redirect.github.com/urfave/cli/pull/2237)
- Fix:(issue\_2238) Dont process flags for completion command by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2239](https://redirect.github.com/urfave/cli/pull/2239)
- Fix:(issue\_2228) Fix for default command by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2232](https://redirect.github.com/urfave/cli/pull/2232)
- chore(deps): bump mkdocs-material from 9.7.0 to 9.7.1 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2242](https://redirect.github.com/urfave/cli/pull/2242)
- Docs: correct typo in migration guide by [@&#8203;kzygmans](https://redirect.github.com/kzygmans) in [#&#8203;2243](https://redirect.github.com/urfave/cli/pull/2243)
- Fix:(issue\_2244) Dont check req flags for help and completion commands by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2245](https://redirect.github.com/urfave/cli/pull/2245)
- Only show separator if command has usage text by [@&#8203;mikecluck](https://redirect.github.com/mikecluck) in [#&#8203;2247](https://redirect.github.com/urfave/cli/pull/2247)

#### New Contributors

- [@&#8203;kzygmans](https://redirect.github.com/kzygmans) made their first contribution in [#&#8203;2243](https://redirect.github.com/urfave/cli/pull/2243)
- [@&#8203;mikecluck](https://redirect.github.com/mikecluck) made their first contribution in [#&#8203;2247](https://redirect.github.com/urfave/cli/pull/2247)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.6.1...v3.6.2>

### [`v3.6.1`](https://redirect.github.com/urfave/cli/releases/tag/v3.6.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.6.0...v3.6.1)

#### What's Changed

- chore(deps): bump golangci/golangci-lint-action from 8 to 9 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2222](https://redirect.github.com/urfave/cli/pull/2222)
- feat: add ability to override usage text of default help command by [@&#8203;Maks1mS](https://redirect.github.com/Maks1mS) in [#&#8203;2196](https://redirect.github.com/urfave/cli/pull/2196)
- Fix:(issue\_2223) Fix incorrect processing of empty value after = by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2224](https://redirect.github.com/urfave/cli/pull/2224)

#### New Contributors

- [@&#8203;Maks1mS](https://redirect.github.com/Maks1mS) made their first contribution in [#&#8203;2196](https://redirect.github.com/urfave/cli/pull/2196)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.6.0...v3.6.1>

### [`v3.6.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.6.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.5.0...v3.6.0)

#### What's Changed

- support parallel running of commands by [@&#8203;oprudkyi](https://redirect.github.com/oprudkyi) in [#&#8203;2215](https://redirect.github.com/urfave/cli/pull/2215)
- Fix:(issue\_2208) Fix local flag by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2211](https://redirect.github.com/urfave/cli/pull/2211)
- chore(deps): bump the python-packages group with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2219](https://redirect.github.com/urfave/cli/pull/2219)
- Call actions on flags set from env by [@&#8203;malclocke](https://redirect.github.com/malclocke) in [#&#8203;2221](https://redirect.github.com/urfave/cli/pull/2221)

#### New Contributors

- [@&#8203;malclocke](https://redirect.github.com/malclocke) made their first contribution in [#&#8203;2221](https://redirect.github.com/urfave/cli/pull/2221)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.5.0...v3.6.0>

### [`v3.5.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.5.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.4.1...v3.5.0)

#### What's Changed

- Update mkdocs reqs by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;2190](https://redirect.github.com/urfave/cli/pull/2190)
- Allow the user to stop processing flags after seeing N args by [@&#8203;adrian-thurston](https://redirect.github.com/adrian-thurston) in [#&#8203;2163](https://redirect.github.com/urfave/cli/pull/2163)
- chore(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2194](https://redirect.github.com/urfave/cli/pull/2194)
- chore(deps): bump mkdocs-material from 9.6.16 to 9.6.18 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2195](https://redirect.github.com/urfave/cli/pull/2195)
- chore(deps): bump actions/setup-go from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2198](https://redirect.github.com/urfave/cli/pull/2198)
- chore(deps): bump actions/setup-node from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2199](https://redirect.github.com/urfave/cli/pull/2199)
- chore(deps): bump actions/setup-python from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2200](https://redirect.github.com/urfave/cli/pull/2200)
- chore(deps): bump github.com/stretchr/testify from 1.11.0 to 1.11.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2197](https://redirect.github.com/urfave/cli/pull/2197)
- chore(deps): bump mkdocs-material from 9.6.18 to 9.6.19 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2201](https://redirect.github.com/urfave/cli/pull/2201)
- chore(deps): bump mkdocs-material from 9.6.19 to 9.6.20 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2202](https://redirect.github.com/urfave/cli/pull/2202)
- feat: add name of argument into error message when parsing fails by [@&#8203;oprudkyi](https://redirect.github.com/oprudkyi) in [#&#8203;2203](https://redirect.github.com/urfave/cli/pull/2203)
- chore(deps): bump mkdocs-material from 9.6.20 to 9.6.21 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2204](https://redirect.github.com/urfave/cli/pull/2204)
- add space between arguments usage by [@&#8203;dimfu](https://redirect.github.com/dimfu) in [#&#8203;2207](https://redirect.github.com/urfave/cli/pull/2207)
- chore(deps): bump mkdocs-material from 9.6.21 to 9.6.22 in the python-packages group by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2213](https://redirect.github.com/urfave/cli/pull/2213)
- Fix: Make DefaultText behaviour consistent by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2214](https://redirect.github.com/urfave/cli/pull/2214)

#### New Contributors

- [@&#8203;adrian-thurston](https://redirect.github.com/adrian-thurston) made their first contribution in [#&#8203;2163](https://redirect.github.com/urfave/cli/pull/2163)
- [@&#8203;oprudkyi](https://redirect.github.com/oprudkyi) made their first contribution in [#&#8203;2203](https://redirect.github.com/urfave/cli/pull/2203)
- [@&#8203;dimfu](https://redirect.github.com/dimfu) made their first contribution in [#&#8203;2207](https://redirect.github.com/urfave/cli/pull/2207)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.4.1...v3.5.0>

### [`v3.4.1`](https://redirect.github.com/urfave/cli/releases/tag/v3.4.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.4.0...v3.4.1)

#### What's Changed

- Use recommended GitHub Actions runner labels by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;2181](https://redirect.github.com/urfave/cli/pull/2181)
- chore(deps): bump actions/checkout from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;2179](https://redirect.github.com/urfave/cli/pull/2179)
- Document that `v3` series is recommended for new development by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;2186](https://redirect.github.com/urfave/cli/pull/2186)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.4.0...v3.4.1>

### [`v3.4.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.4.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.9...v3.4.0)

#### What's Changed

- Export help display functions as variables to allow custom help display logic by [@&#8203;almas-x](https://redirect.github.com/almas-x) in [#&#8203;2150](https://redirect.github.com/urfave/cli/pull/2150)
- Invoke OnUsageError when missing required flags by [@&#8203;MohitPanchariya](https://redirect.github.com/MohitPanchariya) in [#&#8203;2161](https://redirect.github.com/urfave/cli/pull/2161)
- Fix:(issue\_2169) Allow trim space for string slice flags by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2171](https://redirect.github.com/urfave/cli/pull/2171)
- Add example of flag groups to docs by [@&#8203;jllovet](https://redirect.github.com/jllovet) in [#&#8203;2178](https://redirect.github.com/urfave/cli/pull/2178)
- Add installation instructions for gfmrun by [@&#8203;jllovet](https://redirect.github.com/jllovet) in [#&#8203;2177](https://redirect.github.com/urfave/cli/pull/2177)
- Ensure public vars reference public types by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;2180](https://redirect.github.com/urfave/cli/pull/2180)

#### New Contributors

- [@&#8203;almas-x](https://redirect.github.com/almas-x) made their first contribution in [#&#8203;2150](https://redirect.github.com/urfave/cli/pull/2150)
- [@&#8203;MohitPanchariya](https://redirect.github.com/MohitPanchariya) made their first contribution in [#&#8203;2161](https://redirect.github.com/urfave/cli/pull/2161)
- [@&#8203;jllovet](https://redirect.github.com/jllovet) made their first contribution in [#&#8203;2178](https://redirect.github.com/urfave/cli/pull/2178)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.9...v3.4.0>

### [`v3.3.9`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.9)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.8...v3.3.9)

#### What's Changed

- Fix typos in documentation for customizations and full API example by [@&#8203;amarjit03](https://redirect.github.com/amarjit03) in [#&#8203;2165](https://redirect.github.com/urfave/cli/pull/2165)
- Update advanced.md by [@&#8203;thetillhoff](https://redirect.github.com/thetillhoff) in [#&#8203;2170](https://redirect.github.com/urfave/cli/pull/2170)

#### New Contributors

- [@&#8203;amarjit03](https://redirect.github.com/amarjit03) made their first contribution in [#&#8203;2165](https://redirect.github.com/urfave/cli/pull/2165)
- [@&#8203;thetillhoff](https://redirect.github.com/thetillhoff) made their first contribution in [#&#8203;2170](https://redirect.github.com/urfave/cli/pull/2170)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.8...v3.3.9>

### [`v3.3.8`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.8)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.7...v3.3.8)

#### What's Changed

- Remove "alpha" wording around `v3` series by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;2155](https://redirect.github.com/urfave/cli/pull/2155)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.7...v3.3.8>

### [`v3.3.7`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.7)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.6...v3.3.7)

#### What's Changed

- fix: add missing `IsLocal` for BoolWithInverseFlag by [@&#8203;huiyifyj](https://redirect.github.com/huiyifyj) in [#&#8203;2151](https://redirect.github.com/urfave/cli/pull/2151)
- Fix OnUsageError Trigger When Error Is Caused by Mutually Exclusive Flags by [@&#8203;Ali-Doustkani](https://redirect.github.com/Ali-Doustkani) in [#&#8203;2152](https://redirect.github.com/urfave/cli/pull/2152)

#### New Contributors

- [@&#8203;Ali-Doustkani](https://redirect.github.com/Ali-Doustkani) made their first contribution in [#&#8203;2152](https://redirect.github.com/urfave/cli/pull/2152)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.6...v3.3.7>

### [`v3.3.6`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.6)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.5...v3.3.6)

#### What's Changed

- Fish completions with identically named sub-commands now work by [@&#8203;bittrance](https://redirect.github.com/bittrance) in [#&#8203;2130](https://redirect.github.com/urfave/cli/pull/2130)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.5...v3.3.6>

### [`v3.3.5`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.5)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.4...v3.3.5)

#### What's Changed

- Fix:(issue\_2137) Ensure default value for bool with inverse flag is h… by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2138](https://redirect.github.com/urfave/cli/pull/2138)
- Fix:(issue\_2131) Show help text for BoolWithInverseFlag by [@&#8203;Juneezee](https://redirect.github.com/Juneezee) in [#&#8203;2142](https://redirect.github.com/urfave/cli/pull/2142)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.4...v3.3.5>

### [`v3.3.4`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.4)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.3...v3.3.4)

#### What's Changed

- Fix Docs(issue\_2125) Add PathFlag to StringFlag migration by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2136](https://redirect.github.com/urfave/cli/pull/2136)
- fix: remove extraneous space from subcommand help template by [@&#8203;G-Rath](https://redirect.github.com/G-Rath) in [#&#8203;2140](https://redirect.github.com/urfave/cli/pull/2140)
- Fix:(issue\_2135) Correct formatting of default subcommand USAGE text by [@&#8203;zzspoon](https://redirect.github.com/zzspoon) in [#&#8203;2139](https://redirect.github.com/urfave/cli/pull/2139)

#### New Contributors

- [@&#8203;zzspoon](https://redirect.github.com/zzspoon) made their first contribution in [#&#8203;2139](https://redirect.github.com/urfave/cli/pull/2139)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.3...v3.3.4>

### [`v3.3.3`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.3)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.2...v3.3.3)

#### What's Changed

- Simpler top-level context detection for fish completions by [@&#8203;bittrance](https://redirect.github.com/bittrance) in [#&#8203;2121](https://redirect.github.com/urfave/cli/pull/2121)
- Fish completion inside hidden commands by [@&#8203;bittrance](https://redirect.github.com/bittrance) in [#&#8203;2122](https://redirect.github.com/urfave/cli/pull/2122)
- chore(deps): bump golangci/golangci-lint-action from 7 to 8 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2123](https://redirect.github.com/urfave/cli/pull/2123)
- fix: off-by-one in timestamp parsing by [@&#8203;nickajacks1](https://redirect.github.com/nickajacks1) in [#&#8203;2127](https://redirect.github.com/urfave/cli/pull/2127)
- Fish completions tests invokes setup by [@&#8203;bittrance](https://redirect.github.com/bittrance) in [#&#8203;2124](https://redirect.github.com/urfave/cli/pull/2124)
- Fix docs by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2128](https://redirect.github.com/urfave/cli/pull/2128)

#### New Contributors

- [@&#8203;nickajacks1](https://redirect.github.com/nickajacks1) made their first contribution in [#&#8203;2127](https://redirect.github.com/urfave/cli/pull/2127)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.2...v3.3.3>

### [`v3.3.2`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.2)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.1...v3.3.2)

#### What's Changed

- Add docs for advanced value source by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2119](https://redirect.github.com/urfave/cli/pull/2119)
- docs: add renames to v3 migration doc by [@&#8203;G-Rath](https://redirect.github.com/G-Rath) in [#&#8203;2111](https://redirect.github.com/urfave/cli/pull/2111)

#### New Contributors

- [@&#8203;G-Rath](https://redirect.github.com/G-Rath) made their first contribution in [#&#8203;2111](https://redirect.github.com/urfave/cli/pull/2111)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.1...v3.3.2>

### [`v3.3.1`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.3.0...v3.3.1)

#### What's Changed

- Avoid suggesting files in fish command completions. by [@&#8203;bittrance](https://redirect.github.com/bittrance) in [#&#8203;2114](https://redirect.github.com/urfave/cli/pull/2114)
- Cleanup docs by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2118](https://redirect.github.com/urfave/cli/pull/2118)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.3.0...v3.3.1>

### [`v3.3.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.3.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.2.0...v3.3.0)

#### What's Changed

- Add v3 issue template by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2104](https://redirect.github.com/urfave/cli/pull/2104)
- Fix:(issue\_2105) Ensure fish completion works by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2106](https://redirect.github.com/urfave/cli/pull/2106)
- add test for MutuallyExclusiveFlags with After by [@&#8203;bystones](https://redirect.github.com/bystones) in [#&#8203;2107](https://redirect.github.com/urfave/cli/pull/2107)
- use correct context in After function with subcommand by [@&#8203;bystones](https://redirect.github.com/bystones) in [#&#8203;2108](https://redirect.github.com/urfave/cli/pull/2108)
- Enable to customize completion commands by [@&#8203;suzuki-shunsuke](https://redirect.github.com/suzuki-shunsuke) in [#&#8203;2103](https://redirect.github.com/urfave/cli/pull/2103)
- Fish completions no longer suggest subcommands that have already been picked by [@&#8203;bittrance](https://redirect.github.com/bittrance) in [#&#8203;2117](https://redirect.github.com/urfave/cli/pull/2117)
- feat: adds support for explicit `float32` and `float64` by [@&#8203;ldez](https://redirect.github.com/ldez) in [#&#8203;2112](https://redirect.github.com/urfave/cli/pull/2112)

#### New Contributors

- [@&#8203;bystones](https://redirect.github.com/bystones) made their first contribution in [#&#8203;2107](https://redirect.github.com/urfave/cli/pull/2107)
- [@&#8203;bittrance](https://redirect.github.com/bittrance) made their first contribution in [#&#8203;2117](https://redirect.github.com/urfave/cli/pull/2117)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.2.0...v3.3.0>

### [`v3.2.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.2.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.1.1...v3.2.0)

**Breaking change IntFlag now uses int type and not int64. Please change to using Int64Flag for int64 types. Similar behavior for UintFlag as well. See <https://pkg.go.dev/github.com/urfave/cli/v3> for a full list of flag types. See [#&#8203;2094](https://redirect.github.com/urfave/cli/issues/2094)  for full patch for this**

#### What's Changed

- chore: Bump golangci-lint to v2 by [@&#8203;mrueg](https://redirect.github.com/mrueg) in [#&#8203;2083](https://redirect.github.com/urfave/cli/pull/2083)
- Fix docs for shell completions by [@&#8203;antimatter96](https://redirect.github.com/antimatter96) in [#&#8203;2090](https://redirect.github.com/urfave/cli/pull/2090)
- docs: improve migration guides render by [@&#8203;ldez](https://redirect.github.com/ldez) in [#&#8203;2091](https://redirect.github.com/urfave/cli/pull/2091)
- docs: improve migration guide v3 by [@&#8203;ldez](https://redirect.github.com/ldez) in [#&#8203;2093](https://redirect.github.com/urfave/cli/pull/2093)
- feat!: add more integers and unsigned integers type flags by [@&#8203;somebadcode](https://redirect.github.com/somebadcode) in [#&#8203;2094](https://redirect.github.com/urfave/cli/pull/2094)
- PR-2094: Fix docs by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2099](https://redirect.github.com/urfave/cli/pull/2099)
- Fix:(PR-2094) Update docs for new types by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2100](https://redirect.github.com/urfave/cli/pull/2100)
- Fix:(issue\_2056) Add cmd.XXXArgs() functions for retrieving args by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2088](https://redirect.github.com/urfave/cli/pull/2088)
- Add docs for arg types by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2101](https://redirect.github.com/urfave/cli/pull/2101)

#### New Contributors

- [@&#8203;antimatter96](https://redirect.github.com/antimatter96) made their first contribution in [#&#8203;2090](https://redirect.github.com/urfave/cli/pull/2090)
- [@&#8203;ldez](https://redirect.github.com/ldez) made their first contribution in [#&#8203;2091](https://redirect.github.com/urfave/cli/pull/2091)
- [@&#8203;somebadcode](https://redirect.github.com/somebadcode) made their first contribution in [#&#8203;2094](https://redirect.github.com/urfave/cli/pull/2094)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.1.1...v3.2.0>

### [`v3.1.1`](https://redirect.github.com/urfave/cli/compare/v3.1.0...v3.1.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v3.1.0...v3.1.1)

### [`v3.1.0`](https://redirect.github.com/urfave/cli/releases/tag/v3.1.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.7...v3.1.0)

#### What's Changed

- go.mod: Require go1.22 by [@&#8203;mrueg](https://redirect.github.com/mrueg) in [#&#8203;2026](https://redirect.github.com/urfave/cli/pull/2026)
- Fix:(issue\_2030) Add support for trailing hypen for short options by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2031](https://redirect.github.com/urfave/cli/pull/2031)
- Run Before actions after setting up subcommand by [@&#8203;fjl](https://redirect.github.com/fjl) in [#&#8203;2028](https://redirect.github.com/urfave/cli/pull/2028)
- The example have some problem in api by [@&#8203;jokemanfire](https://redirect.github.com/jokemanfire) in [#&#8203;2039](https://redirect.github.com/urfave/cli/pull/2039)
- Rename "Bash Completions" to "Shell Completions" by [@&#8203;abitrolly](https://redirect.github.com/abitrolly) in [#&#8203;2044](https://redirect.github.com/urfave/cli/pull/2044)
- Support root level map keys in map sources by [@&#8203;lukasbindreiter](https://redirect.github.com/lukasbindreiter) in [#&#8203;2047](https://redirect.github.com/urfave/cli/pull/2047)
- while print flag , the placeholder if need but not set. by [@&#8203;jokemanfire](https://redirect.github.com/jokemanfire) in [#&#8203;2043](https://redirect.github.com/urfave/cli/pull/2043)
- Add dependabot by [@&#8203;mrueg](https://redirect.github.com/mrueg) in [#&#8203;2025](https://redirect.github.com/urfave/cli/pull/2025)
- Bump github.com/stretchr/testify from 1.9.0 to 1.10.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2054](https://redirect.github.com/urfave/cli/pull/2054)
- Bump golangci/golangci-lint-action from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2053](https://redirect.github.com/urfave/cli/pull/2053)
- Bump codecov/codecov-action from 4 to 5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;2052](https://redirect.github.com/urfave/cli/pull/2052)
- Fix:(issue\_2032) Support for post parse config loading by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2033](https://redirect.github.com/urfave/cli/pull/2033)
- Fix:(issue\_2066) Remove dependency on golang flag library by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2074](https://redirect.github.com/urfave/cli/pull/2074)
- Fix:(issue\_1891) Roll out v3 docs by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2080](https://redirect.github.com/urfave/cli/pull/2080)
- Fix:(issue\_2077) Make sure onUsageError is invoked for command when a… by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2081](https://redirect.github.com/urfave/cli/pull/2081)

#### New Contributors

- [@&#8203;mrueg](https://redirect.github.com/mrueg) made their first contribution in [#&#8203;2026](https://redirect.github.com/urfave/cli/pull/2026)
- [@&#8203;jokemanfire](https://redirect.github.com/jokemanfire) made their first contribution in [#&#8203;2039](https://redirect.github.com/urfave/cli/pull/2039)
- [@&#8203;lukasbindreiter](https://redirect.github.com/lukasbindreiter) made their first contribution in [#&#8203;2047](https://redirect.github.com/urfave/cli/pull/2047)
- [@&#8203;dependabot](https://redirect.github.com/dependabot) made their first contribution in [#&#8203;2054](https://redirect.github.com/urfave/cli/pull/2054)

**Full Changelog**: <https://github.com/urfave/cli/compare/v3.0.0-beta1.01...v3.1.0>

### [`v2.27.7`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.7)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.6...v2.27.7)

#### What's Changed

- Update dependencies in v2 series by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;2159](https://redirect.github.com/urfave/cli/pull/2159)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.6...v2.27.7>

### [`v2.27.6`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.6)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.5...v2.27.6)

#### What's Changed

- Use usage template in subcommand help by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1986](https://redirect.github.com/urfave/cli/pull/1986)
- Docs: Update cli.yml by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2015](https://redirect.github.com/urfave/cli/pull/2015)
- Fix:(issue\_2069) Add sep for string slice by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;2070](https://redirect.github.com/urfave/cli/pull/2070)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.5...v2.27.6>

### [`v2.27.5`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.5)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.4...v2.27.5)

#### What's Changed

- docs(flag): add `UseShortOptionHandling` description by [@&#8203;BlackHole1](https://redirect.github.com/BlackHole1) in [#&#8203;1956](https://redirect.github.com/urfave/cli/pull/1956)
- \[Backport] Fix: Use $0 env var to correctly retrieve the current active shell by [@&#8203;asahasrabuddhe](https://redirect.github.com/asahasrabuddhe) in [#&#8203;1970](https://redirect.github.com/urfave/cli/pull/1970)
- Update dependencies in v2-maint by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1980](https://redirect.github.com/urfave/cli/pull/1980)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.4...v2.27.5>

### [`v2.27.4`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.4)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.3...v2.27.4)

#### What's Changed

- Fix:(issue\_1962) Fix tests failing on 32 bit architectures by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1963](https://redirect.github.com/urfave/cli/pull/1963)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.3...v2.27.4>

### [`v2.27.3`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.3)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.2...v2.27.3)

#### What's Changed

- v2 Docs: Mention value from env as default value by [@&#8203;sj14](https://redirect.github.com/sj14) in [#&#8203;1910](https://redirect.github.com/urfave/cli/pull/1910)
- Bump github.com/xrash/smetrics dependency by [@&#8203;elezar](https://redirect.github.com/elezar) in [#&#8203;1911](https://redirect.github.com/urfave/cli/pull/1911)
- fix: disable bash completion if double dash is included in arguments (v2) by [@&#8203;suzuki-shunsuke](https://redirect.github.com/suzuki-shunsuke) in [#&#8203;1938](https://redirect.github.com/urfave/cli/pull/1938)
- Fix improper whitespace formatting in usageTemplate, AppHelpTemplate … by [@&#8203;caeret](https://redirect.github.com/caeret) in [#&#8203;1947](https://redirect.github.com/urfave/cli/pull/1947)

#### New Contributors

- [@&#8203;sj14](https://redirect.github.com/sj14) made their first contribution in [#&#8203;1910](https://redirect.github.com/urfave/cli/pull/1910)
- [@&#8203;elezar](https://redirect.github.com/elezar) made their first contribution in [#&#8203;1911](https://redirect.github.com/urfave/cli/pull/1911)
- [@&#8203;caeret](https://redirect.github.com/caeret) made their first contribution in [#&#8203;1947](https://redirect.github.com/urfave/cli/pull/1947)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.2...v2.27.3>

### [`v2.27.2`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.2)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.1...v2.27.2)

#### What's Changed

- Cleanup: Remove unnecessary intermediate variables by [@&#8203;1ambd4](https://redirect.github.com/1ambd4) in [#&#8203;1857](https://redirect.github.com/urfave/cli/pull/1857)
- Docs:(issue\_1866) Fix documentation on filepath vs env preference by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1867](https://redirect.github.com/urfave/cli/pull/1867)
- Fix:(issue\_1860) Remove hidden flags from flag categories by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1868](https://redirect.github.com/urfave/cli/pull/1868)
- Fix repeated \[arguments...] in usage template in v2 by [@&#8203;edeustua](https://redirect.github.com/edeustua) in [#&#8203;1872](https://redirect.github.com/urfave/cli/pull/1872)
- Update dependencies, actions steps, and usage for v2-maint by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1888](https://redirect.github.com/urfave/cli/pull/1888)

#### New Contributors

- [@&#8203;1ambd4](https://redirect.github.com/1ambd4) made their first contribution in [#&#8203;1857](https://redirect.github.com/urfave/cli/pull/1857)
- [@&#8203;edeustua](https://redirect.github.com/edeustua) made their first contribution in [#&#8203;1872](https://redirect.github.com/urfave/cli/pull/1872)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.1...v2.27.2>

### [`v2.27.1`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.27.0...v2.27.1)

#### What's Changed

- v2: Add build tag urfave\_cli\_no\_suggest by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1847](https://redirect.github.com/urfave/cli/pull/1847)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.27.0...v2.27.1>

### [`v2.27.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.27.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.26.0...v2.27.0)

#### What's Changed

- v2 Add integration with golangci-lint by [@&#8203;skelouse](https://redirect.github.com/skelouse) in [#&#8203;1830](https://redirect.github.com/urfave/cli/pull/1830)
- v2: GitHub Actions: upgrade Go, upgrade actions by [@&#8203;dolmen](https://redirect.github.com/dolmen) in [#&#8203;1848](https://redirect.github.com/urfave/cli/pull/1848)
- Feat:(issue\_1797) Add Args for app/cmd/subcmd to avoid argument... be… by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1829](https://redirect.github.com/urfave/cli/pull/1829)
- Fix:(issue\_1850) Add RunAction for uint/uint64 slice flags by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1851](https://redirect.github.com/urfave/cli/pull/1851)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.26.0...v2.27.0>

### [`v2.26.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.26.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.7...v2.26.0)

#### What's Changed

- Bash completion nits by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1762](https://redirect.github.com/urfave/cli/pull/1762)
- Chore: Rename mkdocs requirements file name by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1776](https://redirect.github.com/urfave/cli/pull/1776)
- Fix:(issue\_1787) Add fix for commands not listed when hide help comma… by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1788](https://redirect.github.com/urfave/cli/pull/1788)
- Fix nil HelpFlag panic (v2) by [@&#8203;wxiaoguang](https://redirect.github.com/wxiaoguang) in [#&#8203;1795](https://redirect.github.com/urfave/cli/pull/1795)
- Always get 0 for a nested int64 value in v2.25.7 by [@&#8203;stephenfire](https://redirect.github.com/stephenfire) in [#&#8203;1799](https://redirect.github.com/urfave/cli/pull/1799)
- Helper messages for documenting build process by [@&#8203;abitrolly](https://redirect.github.com/abitrolly) in [#&#8203;1800](https://redirect.github.com/urfave/cli/pull/1800)
- fix: check duplicated sub command name and alias by [@&#8203;linrl3](https://redirect.github.com/linrl3) in [#&#8203;1805](https://redirect.github.com/urfave/cli/pull/1805)
- Fix:(issue\_1689) Have consistent behavior for default text in man and… by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1825](https://redirect.github.com/urfave/cli/pull/1825)
- Fix linting issues by [@&#8203;skelouse](https://redirect.github.com/skelouse) in [#&#8203;1696](https://redirect.github.com/urfave/cli/pull/1696)

#### New Contributors

- [@&#8203;stephenfire](https://redirect.github.com/stephenfire) made their first contribution in [#&#8203;1799](https://redirect.github.com/urfave/cli/pull/1799)
- [@&#8203;linrl3](https://redirect.github.com/linrl3) made their first contribution in [#&#8203;1805](https://redirect.github.com/urfave/cli/pull/1805)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.7...v2.26.0>

### [`v2.25.7`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.7)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.6...v2.25.7)

#### What's Changed

- Fix: fix v2 broken tests by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1757](https://redirect.github.com/urfave/cli/pull/1757)
- Fix:(issue\_1755) Ensure that timestamp flag destination is set correctly by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1756](https://redirect.github.com/urfave/cli/pull/1756)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.6...v2.25.7>

### [`v2.25.6`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.6)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.5...v2.25.6)

#### What's Changed

- Fix:(issue\_1668) Add test case for sub command of sub command completion by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1747](https://redirect.github.com/urfave/cli/pull/1747)
- Update dependencies for v2 by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1749](https://redirect.github.com/urfave/cli/pull/1749)
- Document slice flags as part of examples (v2) by [@&#8203;carhartl](https://redirect.github.com/carhartl) in [#&#8203;1751](https://redirect.github.com/urfave/cli/pull/1751)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.5...v2.25.6>

### [`v2.25.5`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.5)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.4...v2.25.5)

#### What's Changed

- Fix:(issue\_1737) Set bool count by taking care of num of aliases by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1740](https://redirect.github.com/urfave/cli/pull/1740)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.4...v2.25.5>

### [`v2.25.4`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.4)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.3...v2.25.4)

#### What's Changed

- Bug/fix issue 1703 by [@&#8203;jojje](https://redirect.github.com/jojje) in [#&#8203;1728](https://redirect.github.com/urfave/cli/pull/1728)
- Fix:(issue\_1734) Show categories for subcommands by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1735](https://redirect.github.com/urfave/cli/pull/1735)
- Fix:(issue\_1610). Keep RunAsSubcommand behaviour as before by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1736](https://redirect.github.com/urfave/cli/pull/1736)
- Fix:(issue\_1731) Add fix for checking if aliases are set by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1732](https://redirect.github.com/urfave/cli/pull/1732)
- Fix func name referenced in doc comment by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1738](https://redirect.github.com/urfave/cli/pull/1738)

#### New Contributors

- [@&#8203;jojje](https://redirect.github.com/jojje) made their first contribution in [#&#8203;1728](https://redirect.github.com/urfave/cli/pull/1728)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.3...v2.25.4>

### [`v2.25.3`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.3)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.2...v2.25.3)

#### What's Changed

- Fix `incorrectTypeForFlagError` for unknowns by [@&#8203;danhunsaker](https://redirect.github.com/danhunsaker) in [#&#8203;1708](https://redirect.github.com/urfave/cli/pull/1708)

#### New Contributors

- [@&#8203;danhunsaker](https://redirect.github.com/danhunsaker) made their first contribution in [#&#8203;1708](https://redirect.github.com/urfave/cli/pull/1708)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.2...v2.25.3>

### [`v2.25.2`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.2)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.1...v2.25.2)

#### What's Changed

- Fix missing required flag error uses flag name and not alias by [@&#8203;nirhaas](https://redirect.github.com/nirhaas) in [#&#8203;1709](https://redirect.github.com/urfave/cli/pull/1709)
- Remove redundant variable declarations by [@&#8203;huiyifyj](https://redirect.github.com/huiyifyj) in [#&#8203;1714](https://redirect.github.com/urfave/cli/pull/1714)
- Fix:(issue 1689) Match markdown output with help output by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1723](https://redirect.github.com/urfave/cli/pull/1723)

#### New Contributors

- [@&#8203;nirhaas](https://redirect.github.com/nirhaas) made their first contribution in [#&#8203;1709](https://redirect.github.com/urfave/cli/pull/1709)
- [@&#8203;huiyifyj](https://redirect.github.com/huiyifyj) made their first contribution in [#&#8203;1714](https://redirect.github.com/urfave/cli/pull/1714)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.1...v2.25.2>

### [`v2.25.1`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.25.0...v2.25.1)

#### What's Changed

- Shift tested Go versions in v2-maint by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1692](https://redirect.github.com/urfave/cli/pull/1692)
- Preserve separator spec on subcommands by [@&#8203;thschmitt](https://redirect.github.com/thschmitt) in [#&#8203;1710](https://redirect.github.com/urfave/cli/pull/1710)

#### New Contributors

- [@&#8203;thschmitt](https://redirect.github.com/thschmitt) made their first contribution in [#&#8203;1710](https://redirect.github.com/urfave/cli/pull/1710)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.25.0...v2.25.1>

### [`v2.25.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.25.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.24.4...v2.25.0)

#### What's Changed

- Drop support for Go versions before 1.18 by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1667](https://redirect.github.com/urfave/cli/pull/1667)
- add Integer type casting when loading yaml file by [@&#8203;kjeom](https://redirect.github.com/kjeom) in [#&#8203;1669](https://redirect.github.com/urfave/cli/pull/1669)

#### New Contributors

- [@&#8203;kjeom](https://redirect.github.com/kjeom) made their first contribution in [#&#8203;1669](https://redirect.github.com/urfave/cli/pull/1669)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.24.4...v2.25.0>

### [`v2.24.4`](https://redirect.github.com/urfave/cli/releases/tag/v2.24.4)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.24.3...v2.24.4)

#### What's Changed

- Deprecation of package ioutil in Go 1.16 by [@&#8203;ArangoGutierrez](https://redirect.github.com/ArangoGutierrez) in [#&#8203;1678](https://redirect.github.com/urfave/cli/pull/1678)
- Backport \[v2]:  Fix some issues in bash autocompletion by [@&#8203;MrNaif2018](https://redirect.github.com/MrNaif2018) in [#&#8203;1676](https://redirect.github.com/urfave/cli/pull/1676)

#### New Contributors

- [@&#8203;ArangoGutierrez](https://redirect.github.com/ArangoGutierrez) made their first contribution in [#&#8203;1678](https://redirect.github.com/urfave/cli/pull/1678)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.24.3...v2.24.4>

### [`v2.24.3`](https://redirect.github.com/urfave/cli/releases/tag/v2.24.3)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.24.2...v2.24.3)

#### What's Changed

- Make trim space optional by [@&#8203;palsivertsen](https://redirect.github.com/palsivertsen) in [#&#8203;1675](https://redirect.github.com/urfave/cli/pull/1675)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.24.2...v2.24.3>

### [`v2.24.2`](https://redirect.github.com/urfave/cli/releases/tag/v2.24.2)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.24.1...v2.24.2)

#### What's Changed

- Update README badges for v2 by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1664](https://redirect.github.com/urfave/cli/pull/1664)
- Target two most recent Go versions in v2 by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1666](https://redirect.github.com/urfave/cli/pull/1666)
- flag: remove dependencies on shared variables by [@&#8203;zllovesuki](https://redirect.github.com/zllovesuki) in [#&#8203;1671](https://redirect.github.com/urfave/cli/pull/1671)
- Show non categorized flags with categorized on help by [@&#8203;skelouse](https://redirect.github.com/skelouse) in [#&#8203;1673](https://redirect.github.com/urfave/cli/pull/1673)

#### New Contributors

- [@&#8203;zllovesuki](https://redirect.github.com/zllovesuki) made their first contribution in [#&#8203;1671](https://redirect.github.com/urfave/cli/pull/1671)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.24.1...v2.24.2>

### [`v2.24.1`](https://redirect.github.com/urfave/cli/releases/tag/v2.24.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.24.0...v2.24.1)

#### What's Changed

- Fix v2 docs by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1655](https://redirect.github.com/urfave/cli/pull/1655)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.24.0...v2.24.1>

### [`v2.24.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.24.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.10...v2.24.0)

#### What's Changed

- Fix:(issue\_1592) Add support for float64slice, uint, int, int64 for a… by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1647](https://redirect.github.com/urfave/cli/pull/1647)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.10...v2.24.0>

### [`v2.23.10`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.10)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.9...v2.23.10)

#### What's Changed

- Doc:(issue\_1593) Add flag category topic in docs (v2-maint) by [@&#8203;ovcharenko-di](https://redirect.github.com/ovcharenko-di) in [#&#8203;1653](https://redirect.github.com/urfave/cli/pull/1653)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.9...v2.23.10>

### [`v2.23.9`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.9)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.8...v2.23.9)

#### What's Changed

- Fixes [#&#8203;1648](https://redirect.github.com/urfave/cli/issues/1648) by [@&#8203;palsivertsen](https://redirect.github.com/palsivertsen) in [#&#8203;1649](https://redirect.github.com/urfave/cli/pull/1649)

#### New Contributors

- [@&#8203;palsivertsen](https://redirect.github.com/palsivertsen) made their first contribution in [#&#8203;1649](https://redirect.github.com/urfave/cli/pull/1649)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.8...v2.23.9>

### [`v2.23.8`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.8)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.7...v2.23.8)

#### What's Changed

- Fix:(issue\_1277) Remove default text for version/help flags by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1633](https://redirect.github.com/urfave/cli/pull/1633)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.7...v2.23.8>

### [`v2.23.7`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.7)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.6...v2.23.7)

#### What's Changed

- Fix:(issue\_1617) Fix Bash completion for subcommands by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1618](https://redirect.github.com/urfave/cli/pull/1618)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.6...v2.23.7>

### [`v2.23.6`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.6)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.5...v2.23.6)

#### What's Changed

- Disable slice flag separator by [@&#8203;feedmeapples](https://redirect.github.com/feedmeapples) in [#&#8203;1588](https://redirect.github.com/urfave/cli/pull/1588)
- Fix:(issue\_1591) Use AppHelpTemplate instead of subcommand help template by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1608](https://redirect.github.com/urfave/cli/pull/1608)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.5...v2.23.6>

### [`v2.23.5`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.5)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.4...v2.23.5)

#### What's Changed

- Update x/text to 0.3.8 by [@&#8203;dirkmueller](https://redirect.github.com/dirkmueller) in [#&#8203;1571](https://redirect.github.com/urfave/cli/pull/1571)
- Update github actions events for v2-maint branch by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1574](https://redirect.github.com/urfave/cli/pull/1574)
- Update dependencies in v2 series by [@&#8203;meatballhat](https://redirect.github.com/meatballhat) in [#&#8203;1573](https://redirect.github.com/urfave/cli/pull/1573)

#### New Contributors

- [@&#8203;dirkmueller](https://redirect.github.com/dirkmueller) made their first contribution in [#&#8203;1571](https://redirect.github.com/urfave/cli/pull/1571)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.4...v2.23.5>

### [`v2.23.4`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.4)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.3...v2.23.4)

#### What's Changed

- Chore: Add altsrc flag definition generation by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1553](https://redirect.github.com/urfave/cli/pull/1553)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.3...v2.23.4>

### [`v2.23.3`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.3)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.2...v2.23.3)

#### What's Changed

- Fix: Avoid duplication of help commands by [@&#8203;remiposo](https://redirect.github.com/remiposo) in [#&#8203;1565](https://redirect.github.com/urfave/cli/pull/1565)
- Feature:(issue\_1550):Add support Int64Slice by read toml config file by [@&#8203;Edelweiss-Snow](https://redirect.github.com/Edelweiss-Snow) in [#&#8203;1551](https://redirect.github.com/urfave/cli/pull/1551)

#### New Contributors

- [@&#8203;remiposo](https://redirect.github.com/remiposo) made their first contribution in [#&#8203;1565](https://redirect.github.com/urfave/cli/pull/1565)
- [@&#8203;Edelweiss-Snow](https://redirect.github.com/Edelweiss-Snow) made their first contribution in [#&#8203;1551](https://redirect.github.com/urfave/cli/pull/1551)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.2...v2.23.3>

Note. This is considered a minor release even though it has a new "feature" i.e support for int64slice for alstrc flags. The int64slice is verbatim copy of existing code and doesnt include any new behaviour compared to other altsrc flags.

### [`v2.23.2`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.2)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.1...v2.23.2)

#### What's Changed

- Fix:(issue\_1114) Add SkipFlagParsing to app to allow -- by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1562](https://redirect.github.com/urfave/cli/pull/1562)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.1...v2.23.2>

### [`v2.23.1`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.1)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.23.0...v2.23.1)

#### What's Changed

- FIx: Allow ext flags to be opt-out by default rather than opt-in by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1561](https://redirect.github.com/urfave/cli/pull/1561)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.23.0...v2.23.1>

### [`v2.23.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.23.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.22.0...v2.23.0)

#### What's Changed

- Feature:(issue\_269) Allow external package flag definitions by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1540](https://redirect.github.com/urfave/cli/pull/1540)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.22.0...v2.23.0>

### [`v2.22.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.22.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.21.0...v2.22.0)

#### What's Changed

- Feature:(issue\_1090): Add unwrap for ExitCoder by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1545](https://redirect.github.com/urfave/cli/pull/1545)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.21.0...v2.22.0>

### [`v2.21.0`](https://redirect.github.com/urfave/cli/releases/tag/v2.21.0)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.20.5...v2.21.0)

#### What's Changed

- Feature: (issue\_1451) customized slice flag separator by [@&#8203;FGYFFFF](https://redirect.github.com/FGYFFFF) in [#&#8203;1546](https://redirect.github.com/urfave/cli/pull/1546)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.20.5...v2.21.0>

### [`v2.20.5`](https://redirect.github.com/urfave/cli/releases/tag/v2.20.5)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.20.4...v2.20.5)

#### What's Changed

- Fix:(issue\_1548) Check root before run default cmd by [@&#8203;smalnote](https://redirect.github.com/smalnote) in [#&#8203;1549](https://redirect.github.com/urfave/cli/pull/1549)

#### New Contributors

- [@&#8203;smalnote](https://redirect.github.com/smalnote) made their first contribution in [#&#8203;1549](https://redirect.github.com/urfave/cli/pull/1549)

**Full Changelog**: <https://github.com/urfave/cli/compare/v2.20.4...v2.20.5>

### [`v2.20.4`](https://redirect.github.com/urfave/cli/releases/tag/v2.20.4)

[Compare Source](https://redirect.github.com/urfave/cli/compare/v2.20.3...v2.20.4)

#### What's Changed

- Spelling by [@&#8203;jsoref](https://redirect.github.com/jsoref) in [#&#8203;1543](https://redirect.github.com/urfave/cli/pull/1543)
- Documentation:(issue\_786) Add docs for flag validation by [@&#8203;dearchap](https://redirect.github.com/dearchap) in [#&#8203;1544](https://redirect.github.com/urfave/c

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6WyJkZXBlbmRlbmNpZXMiXX0=-->


---

## Reviews

### Review by @Hyperkid123 - Approved on March 17, 2026 at 10:40 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/124*
