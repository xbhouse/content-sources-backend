---
type: pull_request
number: 59
title: "Update module github.com/spf13/cobra to v1.10.1"
state: merged
author: red-hat-konflux
created: 2025-11-10T16:28:15Z
updated: 2025-11-11T14:12:03Z
closed: 2025-11-11T14:12:00Z
merged: 2025-11-11T14:12:00Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-spf13-cobra-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/59
---

# Pull Request #59: Update module github.com/spf13/cobra to v1.10.1

**Author**: @red-hat-konflux
**Created**: November 10, 2025 at 04:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-spf13-cobra-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/spf13/cobra](https://redirect.github.com/spf13/cobra) | `v1.8.1` -> `v1.10.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fspf13%2fcobra/v1.10.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fspf13%2fcobra/v1.8.1/v1.10.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>spf13/cobra (github.com/spf13/cobra)</summary>

### [`v1.10.1`](https://redirect.github.com/spf13/cobra/releases/tag/v1.10.1)

[Compare Source](https://redirect.github.com/spf13/cobra/compare/v1.10.0...v1.10.1)

##### 🐛 Fix

- chore: upgrade pflags v1.0.9 by [@&#8203;jpmcb](https://redirect.github.com/jpmcb) in [#&#8203;2305](https://redirect.github.com/spf13/cobra/pull/2305)

v1.0.9 of pflags brought back `ParseErrorsWhitelist` and marked it as deprecated

**Full Changelog**: <https://github.com/spf13/cobra/compare/v1.10.0...v1.10.1>

### [`v1.10.0`](https://redirect.github.com/spf13/cobra/releases/tag/v1.10.0)

[Compare Source](https://redirect.github.com/spf13/cobra/compare/v1.9.1...v1.10.0)

#### What's Changed

##### 🚨 Attention!

- Bump pflag to 1.0.8 by [@&#8203;tomasaschan](https://redirect.github.com/tomasaschan) in [#&#8203;2303](https://redirect.github.com/spf13/cobra/pull/2303)

This version of `pflag` carried a breaking change: it renamed `ParseErrorsWhitelist` to `ParseErrorsAllowlist` which can break builds if both `pflag` and `cobra` are dependencies in your project.

- If you use both `pflag and `cobra`, upgrade `pflag`to 1.0.8 and`cobra`to`1.10.0\`
- ***or*** use the newer, fixed version of `pflag` v1.0.9 which keeps the deprecated `ParseErrorsWhitelist`

More details can be found here: [#&#8203;2303 (comment)](https://redirect.github.com/spf13/cobra/pull/2303#issuecomment-3242333515)

##### ✨ Features

- Flow context to command in SetHelpFunc by [@&#8203;Frassle](https://redirect.github.com/Frassle) in [#&#8203;2241](https://redirect.github.com/spf13/cobra/pull/2241)
- The default ShellCompDirective can be customized for a command and its subcommands by [@&#8203;albers](https://redirect.github.com/albers) in [#&#8203;2238](https://redirect.github.com/spf13/cobra/pull/2238)

##### 🐛 Fix

- Upgrade golangci-lint to v2, address findings by [@&#8203;scop](https://redirect.github.com/scop) in [#&#8203;2279](https://redirect.github.com/spf13/cobra/pull/2279)

##### 🪠 Testing

- Test with Go 1.24 by [@&#8203;harryzcy](https://redirect.github.com/harryzcy) in [#&#8203;2236](https://redirect.github.com/spf13/cobra/pull/2236)
- chore: Rm GitHub Action PR size labeler by [@&#8203;jpmcb](https://redirect.github.com/jpmcb) in [#&#8203;2256](https://redirect.github.com/spf13/cobra/pull/2256)

##### 📝 Docs

- Remove traling curlybrace by [@&#8203;yedayak](https://redirect.github.com/yedayak) in [#&#8203;2237](https://redirect.github.com/spf13/cobra/pull/2237)
- Update command.go by [@&#8203;styee](https://redirect.github.com/styee) in [#&#8203;2248](https://redirect.github.com/spf13/cobra/pull/2248)
- feat: Add security policy by [@&#8203;jpmcb](https://redirect.github.com/jpmcb) in [#&#8203;2253](https://redirect.github.com/spf13/cobra/pull/2253)
- Update Readme (Warp) by [@&#8203;ericdachen](https://redirect.github.com/ericdachen) in [#&#8203;2267](https://redirect.github.com/spf13/cobra/pull/2267)
- Add Periscope to the list of projects using Cobra by [@&#8203;anishathalye](https://redirect.github.com/anishathalye) in [#&#8203;2299](https://redirect.github.com/spf13/cobra/pull/2299)

#### New Contributors

- [@&#8203;harryzcy](https://redirect.github.com/harryzcy) made their first contribution in [#&#8203;2236](https://redirect.github.com/spf13/cobra/pull/2236)
- [@&#8203;yedayak](https://redirect.github.com/yedayak) made their first contribution in [#&#8203;2237](https://redirect.github.com/spf13/cobra/pull/2237)
- [@&#8203;Frassle](https://redirect.github.com/Frassle) made their first contribution in [#&#8203;2241](https://redirect.github.com/spf13/cobra/pull/2241)
- [@&#8203;styee](https://redirect.github.com/styee) made their first contribution in [#&#8203;2248](https://redirect.github.com/spf13/cobra/pull/2248)
- [@&#8203;ericdachen](https://redirect.github.com/ericdachen) made their first contribution in [#&#8203;2267](https://redirect.github.com/spf13/cobra/pull/2267)
- [@&#8203;albers](https://redirect.github.com/albers) made their first contribution in [#&#8203;2238](https://redirect.github.com/spf13/cobra/pull/2238)
- [@&#8203;anishathalye](https://redirect.github.com/anishathalye) made their first contribution in [#&#8203;2299](https://redirect.github.com/spf13/cobra/pull/2299)
- [@&#8203;tomasaschan](https://redirect.github.com/tomasaschan) made their first contribution in [#&#8203;2303](https://redirect.github.com/spf13/cobra/pull/2303)

**Full Changelog**: <https://github.com/spf13/cobra/compare/v1.9.1...v1.9.2>

### [`v1.9.1`](https://redirect.github.com/spf13/cobra/releases/tag/v1.9.1)

[Compare Source](https://redirect.github.com/spf13/cobra/compare/v1.9.0...v1.9.1)

##### 🐛 Fixes

- Fix CompletionFunc implementation by [@&#8203;ccoVeille](https://redirect.github.com/ccoVeille) in [#&#8203;2234](https://redirect.github.com/spf13/cobra/pull/2234)
- Revert "Make detection for test-binary more universal ([#&#8203;2173](https://redirect.github.com/spf13/cobra/issues/2173))" by [@&#8203;marckhouzam](https://redirect.github.com/marckhouzam) in [#&#8203;2235](https://redirect.github.com/spf13/cobra/pull/2235)

**Full Changelog**: <https://github.com/spf13/cobra/compare/v1.9.0...v1.9.1>

### [`v1.9.0`](https://redirect.github.com/spf13/cobra/releases/tag/v1.9.0)

[Compare Source](https://redirect.github.com/spf13/cobra/compare/v1.8.1...v1.9.0)

#### ✨ Features

- Allow linker to perform deadcode elimination for program using Cobra by [@&#8203;aarzilli](https://redirect.github.com/aarzilli) in [#&#8203;1956](https://redirect.github.com/spf13/cobra/pull/1956)
- Add default completion command even if there are no other sub-commands by [@&#8203;marckhouzam](https://redirect.github.com/marckhouzam) in [#&#8203;1559](https://redirect.github.com/spf13/cobra/pull/1559)
- Add CompletionWithDesc helper by [@&#8203;ccoVeille](https://redirect.github.com/ccoVeille) in [#&#8203;2231](https://redirect.github.com/spf13/cobra/pull/2231)

#### 🐛 Fixes

- Fix deprecation comment for Command.SetOutput by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [#&#8203;2172](https://redirect.github.com/spf13/cobra/pull/2172)
- Replace deprecated ioutil usage by [@&#8203;nirs](https://redirect.github.com/nirs) in [#&#8203;2181](https://redirect.github.com/spf13/cobra/pull/2181)
- Fix --version help and output for plugins by [@&#8203;nirs](https://redirect.github.com/nirs) in [#&#8203;2180](https://redirect.github.com/spf13/cobra/pull/2180)
- Allow to reset the templates to the default by [@&#8203;marckhouzam](https://redirect.github.com/marckhouzam) in [#&#8203;2229](https://redirect.github.com/spf13/cobra/pull/2229)

#### 🤖 Completions

- Make Powershell completion work in constrained mode by [@&#8203;lstemplinger](https://redirect.github.com/lstemplinger) in [#&#8203;2196](https://redirect.github.com/spf13/cobra/pull/2196)
- Improve detection for flags that accept multiple values by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [#&#8203;2210](https://redirect.github.com/spf13/cobra/pull/2210)
- add CompletionFunc type to help with completions by [@&#8203;ccoVeille](https://redirect.github.com/ccoVeille) in [#&#8203;2220](https://redirect.github.com/spf13/cobra/pull/2220)
- Add similar whitespace escape logic to bash v2 completions than in other completions by [@&#8203;kangasta](https://redirect.github.com/kangasta) in [#&#8203;1743](https://redirect.github.com/spf13/cobra/pull/1743)
- Print ActiveHelp for bash along other completions by [@&#8203;marckhouzam](https://redirect.github.com/marckhouzam) in [#&#8203;2076](https://redirect.github.com/spf13/cobra/pull/2076)
- fix(completions): Complete map flags multiple times by [@&#8203;gabe565](https://redirect.github.com/gabe565) in [#&#8203;2174](https://redirect.github.com/spf13/cobra/pull/2174)
- fix(bash): nounset unbound file filter variable on empty extension by [@&#8203;scop](https://redirect.github.com/scop) in [#&#8203;2228](https://redirect.github.com/spf13/cobra/pull/2228)

#### 🧪 Testing

- Test also with go 1.23 by [@&#8203;nirs](https://redirect.github.com/nirs) in [#&#8203;2182](https://redirect.github.com/spf13/cobra/pull/2182)
- Make detection for test-binary more universal by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [#&#8203;2173](https://redirect.github.com/spf13/cobra/pull/2173)

#### ✍🏼 Documentation

- docs: update README.md by [@&#8203;eltociear](https://redirect.github.com/eltociear) in [#&#8203;2197](https://redirect.github.com/spf13/cobra/pull/2197)
- Improve site formatting by [@&#8203;nirs](https://redirect.github.com/nirs) in [#&#8203;2183](https://redirect.github.com/spf13/cobra/pull/2183)
- doc: add Conduit by [@&#8203;raulb](https://redirect.github.com/raulb) in [#&#8203;2230](https://redirect.github.com/spf13/cobra/pull/2230)
- doc: azion project added to the list of CLIs that use cobra by [@&#8203;maxwelbm](https://redirect.github.com/maxwelbm) in [#&#8203;2198](https://redirect.github.com/spf13/cobra/pull/2198)
- Fix broken links in active\_help.md by [@&#8203;vuil](https://redirect.github.com/vuil) in [#&#8203;2202](https://redirect.github.com/spf13/cobra/pull/2202)
- chore: fix function name in comment by [@&#8203;zhuhaicity](https://redirect.github.com/zhuhaicity) in [#&#8203;2216](https://redirect.github.com/spf13/cobra/pull/2216)

#### 🔧 Dependency upgrades

- build(deps): bump github.com/cpuguy83/go-md2man/v2 from 2.0.5 to 2.0.6 by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [#&#8203;2206](https://redirect.github.com/spf13/cobra/pull/2206)
- Update to latest go-md2man by [@&#8203;mikelolasagasti](https://redirect.github.com/mikelolasagasti) in [#&#8203;2201](https://redirect.github.com/spf13/cobra/pull/2201)
- Upgrade `pflag` dependencies for v1.9.0 by [@&#8203;jpmcb](https://redirect.github.com/jpmcb) in [#&#8203;2233](https://redirect.github.com/spf13/cobra/pull/2233)

***

Thank you to all of our amazing contributors and all the great work that's been going into the completions feature!!

##### 👋🏼 New Contributors

- [@&#8203;gabe565](https://redirect.github.com/gabe565) made their first contribution in [#&#8203;2174](https://redirect.github.com/spf13/cobra/pull/2174)
- [@&#8203;maxwelbm](https://redirect.github.com/maxwelbm) made their first contribution in [#&#8203;2198](https://redirect.github.com/spf13/cobra/pull/2198)
- [@&#8203;lstemplinger](https://redirect.github.com/lstemplinger) made their first contribution in [#&#8203;2196](https://redirect.github.com/spf13/cobra/pull/2196)
- [@&#8203;vuil](https://redirect.github.com/vuil) made their first contribution in [#&#8203;2202](https://redirect.github.com/spf13/cobra/pull/2202)
- [@&#8203;mikelolasagasti](https://redirect.github.com/mikelolasagasti) made their first contribution in [#&#8203;2201](https://redirect.github.com/spf13/cobra/pull/2201)
- [@&#8203;zhuhaicity](https://redirect.github.com/zhuhaicity) made their first contribution in [#&#8203;2216](https://redirect.github.com/spf13/cobra/pull/2216)
- [@&#8203;ccoVeille](https://redirect.github.com/ccoVeille) made their first contribution in [#&#8203;2220](https://redirect.github.com/spf13/cobra/pull/2220)
- [@&#8203;kangasta](https://redirect.github.com/kangasta) made their first contribution in [#&#8203;1743](https://redirect.github.com/spf13/cobra/pull/1743)
- [@&#8203;aarzilli](https://redirect.github.com/aarzilli) made their first contribution in [#&#8203;1956](https://redirect.github.com/spf13/cobra/pull/1956)

**Full Changelog**: <https://github.com/spf13/cobra/compare/v1.8.1...v1.9.0>

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

### Review by @Hyperkid123 - Approved on November 11, 2025 at 02:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/59*
