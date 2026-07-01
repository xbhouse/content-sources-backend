---
type: pull_request
number: 26
title: "Update module github.com/spf13/pflag to v1.0.10"
state: merged
author: red-hat-konflux
created: 2025-10-20T12:18:13Z
updated: 2025-10-21T06:56:38Z
closed: 2025-10-21T06:56:35Z
merged: 2025-10-21T06:56:35Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-spf13-pflag-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/26
---

# Pull Request #26: Update module github.com/spf13/pflag to v1.0.10

**Author**: @red-hat-konflux
**Created**: October 20, 2025 at 12:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-spf13-pflag-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/spf13/pflag](https://redirect.github.com/spf13/pflag) | `v1.0.5` -> `v1.0.10` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fspf13%2fpflag/v1.0.10?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fspf13%2fpflag/v1.0.5/v1.0.10?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>spf13/pflag (github.com/spf13/pflag)</summary>

### [`v1.0.10`](https://redirect.github.com/spf13/pflag/releases/tag/v1.0.10)

[Compare Source](https://redirect.github.com/spf13/pflag/compare/v1.0.9...v1.0.10)

#### What's Changed

- fix deprecation comment for (FlagSet.)ParseErrorsWhitelist by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [#&#8203;447](https://redirect.github.com/spf13/pflag/pull/447)
- remove uses of errors.Is, which requires go1.13, move go1.16/go1.21 tests to separate file by [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) in [#&#8203;448](https://redirect.github.com/spf13/pflag/pull/448)

#### New Contributors

- [@&#8203;thaJeztah](https://redirect.github.com/thaJeztah) made their first contribution in [#&#8203;447](https://redirect.github.com/spf13/pflag/pull/447)

**Full Changelog**: <https://github.com/spf13/pflag/compare/v1.0.9...v1.0.10>

### [`v1.0.9`](https://redirect.github.com/spf13/pflag/releases/tag/v1.0.9)

[Compare Source](https://redirect.github.com/spf13/pflag/compare/v1.0.8...v1.0.9)

#### What's Changed

- fix: Restore ParseErrorsWhitelist name for now by [@&#8203;tomasaschan](https://redirect.github.com/tomasaschan) in [#&#8203;446](https://redirect.github.com/spf13/pflag/pull/446)

**Full Changelog**: <https://github.com/spf13/pflag/compare/v1.0.8...v1.0.9>

### [`v1.0.8`](https://redirect.github.com/spf13/pflag/releases/tag/v1.0.8)

[Compare Source](https://redirect.github.com/spf13/pflag/compare/v1.0.7...v1.0.8)

#### :warning: Breaking Change

This version, while only a patch bump, includes a (very minor) breaking change: the `flag.ParseErrorsWhitelist` struct and corresponding `FlagSet.parseErrorsWhitelist` field have been renamed to `ParseErrorsAllowlist`.

This should result in compilation errors in any code that uses these fields, which can be fixed by adjusting the names at call sites. There is no change in semantics or behavior of the struct or field referred to by these names. If your code compiles without errors after bumping to/past v1.0.8, you are not affected by this change.

The breaking change was reverted in v1.0.9, by means of re-introducing the old names with deprecation warnings. The plan is still to remove them in a future release, so if your code does depend on the old names, please change them to use the new names at your earliest convenience.

#### What's Changed

- Remove Redundant "Unknown-Flag" Error by [@&#8203;vaguecoder](https://redirect.github.com/vaguecoder) in [#&#8203;364](https://redirect.github.com/spf13/pflag/pull/364)
- Switching from whitelist to Allowlist terminology by [@&#8203;dubrie](https://redirect.github.com/dubrie) in [#&#8203;261](https://redirect.github.com/spf13/pflag/pull/261)
- Omit zero time.Time default from usage line by [@&#8203;mologie](https://redirect.github.com/mologie) in [#&#8203;438](https://redirect.github.com/spf13/pflag/pull/438)
- implement CopyToGoFlagSet by [@&#8203;pohly](https://redirect.github.com/pohly) in [#&#8203;330](https://redirect.github.com/spf13/pflag/pull/330)
- flag: Emulate stdlib behavior and do not print ErrHelp by [@&#8203;tmc](https://redirect.github.com/tmc) in [#&#8203;407](https://redirect.github.com/spf13/pflag/pull/407)
- Print Default Values of String-to-String in Sorted Order by [@&#8203;vaguecoder](https://redirect.github.com/vaguecoder) in [#&#8203;365](https://redirect.github.com/spf13/pflag/pull/365)
- fix: Don't print ErrHelp in ParseAll by [@&#8203;tomasaschan](https://redirect.github.com/tomasaschan) in [#&#8203;443](https://redirect.github.com/spf13/pflag/pull/443)
- Reset args on re-parse even if empty by [@&#8203;tomasaschan](https://redirect.github.com/tomasaschan) in [#&#8203;444](https://redirect.github.com/spf13/pflag/pull/444)

#### New Contributors

- [@&#8203;vaguecoder](https://redirect.github.com/vaguecoder) made their first contribution in [#&#8203;364](https://redirect.github.com/spf13/pflag/pull/364)
- [@&#8203;dubrie](https://redirect.github.com/dubrie) made their first contribution in [#&#8203;261](https://redirect.github.com/spf13/pflag/pull/261)
- [@&#8203;mologie](https://redirect.github.com/mologie) made their first contribution in [#&#8203;438](https://redirect.github.com/spf13/pflag/pull/438)
- [@&#8203;pohly](https://redirect.github.com/pohly) made their first contribution in [#&#8203;330](https://redirect.github.com/spf13/pflag/pull/330)
- [@&#8203;tmc](https://redirect.github.com/tmc) made their first contribution in [#&#8203;407](https://redirect.github.com/spf13/pflag/pull/407)
- [@&#8203;tomasaschan](https://redirect.github.com/tomasaschan) made their first contribution in [#&#8203;443](https://redirect.github.com/spf13/pflag/pull/443)

**Full Changelog**: <https://github.com/spf13/pflag/compare/v1.0.7...v1.0.8>

### [`v1.0.7`](https://redirect.github.com/spf13/pflag/releases/tag/v1.0.7)

[Compare Source](https://redirect.github.com/spf13/pflag/compare/v1.0.6...v1.0.7)

#### What's Changed

- Fix defaultIsZeroValue check for generic Value types by [@&#8203;MidnightRocket](https://redirect.github.com/MidnightRocket) in [#&#8203;422](https://redirect.github.com/spf13/pflag/pull/422)
- feat: Use structs for errors returned by pflag. by [@&#8203;eth-p](https://redirect.github.com/eth-p) in [#&#8203;425](https://redirect.github.com/spf13/pflag/pull/425)
- Fix typos by [@&#8203;co63oc](https://redirect.github.com/co63oc) in [#&#8203;428](https://redirect.github.com/spf13/pflag/pull/428)
- fix [#&#8203;423](https://redirect.github.com/spf13/pflag/issues/423) : Add helper function and some documentation to parse shorthand go test flags. by [@&#8203;valdar](https://redirect.github.com/valdar) in [#&#8203;424](https://redirect.github.com/spf13/pflag/pull/424)
- add support equivalent to golang flag.TextVar(), also fixes the test failure as described in [#&#8203;368](https://redirect.github.com/spf13/pflag/issues/368) by [@&#8203;hujun-open](https://redirect.github.com/hujun-open) in [#&#8203;418](https://redirect.github.com/spf13/pflag/pull/418)
- add support for Func() and BoolFunc() [#&#8203;426](https://redirect.github.com/spf13/pflag/issues/426) by [@&#8203;LeGEC](https://redirect.github.com/LeGEC) in [#&#8203;429](https://redirect.github.com/spf13/pflag/pull/429)
- fix: correct argument length check in FlagSet.Parse by [@&#8203;ShawnJeffersonWang](https://redirect.github.com/ShawnJeffersonWang) in [#&#8203;409](https://redirect.github.com/spf13/pflag/pull/409)
- fix usage message for func flags, fix arguments order by [@&#8203;LeGEC](https://redirect.github.com/LeGEC) in [#&#8203;431](https://redirect.github.com/spf13/pflag/pull/431)
- Add support for time.Time flags by [@&#8203;max-frank](https://redirect.github.com/max-frank) in [#&#8203;348](https://redirect.github.com/spf13/pflag/pull/348)

#### New Contributors

- [@&#8203;MidnightRocket](https://redirect.github.com/MidnightRocket) made their first contribution in [#&#8203;422](https://redirect.github.com/spf13/pflag/pull/422)
- [@&#8203;eth-p](https://redirect.github.com/eth-p) made their first contribution in [#&#8203;425](https://redirect.github.com/spf13/pflag/pull/425)
- [@&#8203;co63oc](https://redirect.github.com/co63oc) made their first contribution in [#&#8203;428](https://redirect.github.com/spf13/pflag/pull/428)
- [@&#8203;valdar](https://redirect.github.com/valdar) made their first contribution in [#&#8203;424](https://redirect.github.com/spf13/pflag/pull/424)
- [@&#8203;hujun-open](https://redirect.github.com/hujun-open) made their first contribution in [#&#8203;418](https://redirect.github.com/spf13/pflag/pull/418)
- [@&#8203;LeGEC](https://redirect.github.com/LeGEC) made their first contribution in [#&#8203;429](https://redirect.github.com/spf13/pflag/pull/429)
- [@&#8203;ShawnJeffersonWang](https://redirect.github.com/ShawnJeffersonWang) made their first contribution in [#&#8203;409](https://redirect.github.com/spf13/pflag/pull/409)
- [@&#8203;max-frank](https://redirect.github.com/max-frank) made their first contribution in [#&#8203;348](https://redirect.github.com/spf13/pflag/pull/348)

**Full Changelog**: <https://github.com/spf13/pflag/compare/v1.0.6...v1.0.7>

### [`v1.0.6`](https://redirect.github.com/spf13/pflag/releases/tag/v1.0.6)

[Compare Source](https://redirect.github.com/spf13/pflag/compare/v1.0.5...v1.0.6)

#### What's Changed

- Add exported functions to preserve `pkg/flag` compatibility by [@&#8203;mckern](https://redirect.github.com/mckern) in [#&#8203;220](https://redirect.github.com/spf13/pflag/pull/220)
- remove dead code for checking error nil by [@&#8203;yashbhutwala](https://redirect.github.com/yashbhutwala) in [#&#8203;282](https://redirect.github.com/spf13/pflag/pull/282)
- Add IPNetSlice and unit tests by [@&#8203;rpothier](https://redirect.github.com/rpothier) in [#&#8203;170](https://redirect.github.com/spf13/pflag/pull/170)
- allow for blank ip addresses by [@&#8203;duhruh](https://redirect.github.com/duhruh) in [#&#8203;316](https://redirect.github.com/spf13/pflag/pull/316)
- add github actions by [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) in [#&#8203;419](https://redirect.github.com/spf13/pflag/pull/419)

#### New Contributors

- [@&#8203;mckern](https://redirect.github.com/mckern) made their first contribution in [#&#8203;220](https://redirect.github.com/spf13/pflag/pull/220)
- [@&#8203;yashbhutwala](https://redirect.github.com/yashbhutwala) made their first contribution in [#&#8203;282](https://redirect.github.com/spf13/pflag/pull/282)
- [@&#8203;rpothier](https://redirect.github.com/rpothier) made their first contribution in [#&#8203;170](https://redirect.github.com/spf13/pflag/pull/170)
- [@&#8203;duhruh](https://redirect.github.com/duhruh) made their first contribution in [#&#8203;316](https://redirect.github.com/spf13/pflag/pull/316)
- [@&#8203;sagikazarmark](https://redirect.github.com/sagikazarmark) made their first contribution in [#&#8203;419](https://redirect.github.com/spf13/pflag/pull/419)

**Full Changelog**: <https://github.com/spf13/pflag/compare/v1.0.5...v1.0.6>

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

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @Hyperkid123 on October 21, 2025 at 06:29 AM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/26*
