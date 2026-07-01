---
type: pull_request
number: 1629
title: "Update module github.com/bytedance/sonic to v1.13.2"
state: merged
author: red-hat-konflux
created: 2025-04-06T06:00:23Z
updated: 2025-04-15T09:35:41Z
closed: 2025-04-15T09:35:41Z
merged: 2025-04-15T09:35:41Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-bytedance-sonic-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1629
---

# Pull Request #1629: Update module github.com/bytedance/sonic to v1.13.2

**Author**: @red-hat-konflux
**Created**: April 06, 2025 at 06:00 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-bytedance-sonic-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/bytedance/sonic](https://redirect.github.com/bytedance/sonic) | require | minor | `v1.12.7` -> `v1.13.2` |

---

### Release Notes

<details>
<summary>bytedance/sonic (github.com/bytedance/sonic)</summary>

### [`v1.13.2`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.13.2)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.13.1...v1.13.2)

#### What's Changed

-   fix(amd64) bugs when parsing into `nil-iface`  by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/776](https://redirect.github.com/bytedance/sonic/pull/776)
-   fix(aarch64): bugs when parsing into `[]byte` by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/776](https://redirect.github.com/bytedance/sonic/pull/776)
-   fix(aarch64): unmarshal into `[]json.Number` by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/779](https://redirect.github.com/bytedance/sonic/pull/779)

**Full Changelog**: https://github.com/bytedance/sonic/compare/v1.13.1...v1.13.2

### [`v1.13.1`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.13.1)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.13.0...v1.13.1)

#### What's Changed

-   fix(aarch64): avoid using the boundary pointer of an array by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/763](https://redirect.github.com/bytedance/sonic/pull/763)
-   fix: Write version warning to stderr by [@&#8203;roysc](https://redirect.github.com/roysc) in [https://github.com/bytedance/sonic/pull/761](https://redirect.github.com/bytedance/sonic/pull/761)
-   feat: support golang 1.24 by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/741](https://redirect.github.com/bytedance/sonic/pull/741)
-   fix(compat): duplicate description by [@&#8203;y1jiong](https://redirect.github.com/y1jiong) in [https://github.com/bytedance/sonic/pull/767](https://redirect.github.com/bytedance/sonic/pull/767)
-   chore: add notice for go1.24.0 on README by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [https://github.com/bytedance/sonic/pull/766](https://redirect.github.com/bytedance/sonic/pull/766)
-   fix: loader version by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/770](https://redirect.github.com/bytedance/sonic/pull/770)

#### New Contributors

-   [@&#8203;roysc](https://redirect.github.com/roysc) made their first contribution in [https://github.com/bytedance/sonic/pull/761](https://redirect.github.com/bytedance/sonic/pull/761)
-   [@&#8203;y1jiong](https://redirect.github.com/y1jiong) made their first contribution in [https://github.com/bytedance/sonic/pull/767](https://redirect.github.com/bytedance/sonic/pull/767)

**Full Changelog**: https://github.com/bytedance/sonic/compare/v1.12.10...v1.13.1

### [`v1.13.0`](https://redirect.github.com/bytedance/sonic/compare/v1.12.10...v1.13.0)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.12.10...v1.13.0)

### [`v1.12.10`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.12.10)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.12.9...v1.12.10)

#### Bugfix

-   fix(aarch64): support parse map that key is float by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/748](https://redirect.github.com/bytedance/sonic/pull/748)
-   fix(aarch64): bug when unmarshal into prefilled interface slice by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/751](https://redirect.github.com/bytedance/sonic/pull/751)
-   fix: bug when marshal nil interface{} that contains not-indirect value by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/756](https://redirect.github.com/bytedance/sonic/pull/756)
-   fix(loader): invalid binary length in stackmap by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/754](https://redirect.github.com/bytedance/sonic/pull/754)
-   fix(aarch64): bug when unmarshal into non-nil interface by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/759](https://redirect.github.com/bytedance/sonic/pull/759)
-   fix bug when serde with unsupported type by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/757](https://redirect.github.com/bytedance/sonic/pull/757)

**Full Changelog**: https://github.com/bytedance/sonic/compare/v1.12.9...v1.12.10

### [`v1.12.9`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.12.9)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.12.8...v1.12.9)

#### What's Changed

##### BugFix( for aarch64)

-   fix crash cause by `sigprof` , not use`X28` register in arm64 of native C  \[PR: [#&#8203;746](https://redirect.github.com/bytedance/sonic/issues/746), Issue: [#&#8203;743](https://redirect.github.com/bytedance/sonic/issues/743)]
-   fix bug when parsing invalid json that contains unescaped control chars  \[PR: [#&#8203;745](https://redirect.github.com/bytedance/sonic/issues/745), Issue:  [#&#8203;739](https://redirect.github.com/bytedance/sonic/issues/739)]
-   fix bug when parsing large json  \[PR: [#&#8203;745](https://redirect.github.com/bytedance/sonic/issues/745), Issue:  [#&#8203;744](https://redirect.github.com/bytedance/sonic/issues/744)]

##### Chore

-   chore: fix ci by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/740](https://redirect.github.com/bytedance/sonic/pull/740)
-   Update codeql-analysis.yml by [@&#8203;willem-bd](https://redirect.github.com/willem-bd) in [https://github.com/bytedance/sonic/pull/736](https://redirect.github.com/bytedance/sonic/pull/736)
-   chore: add issue and pr template by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/742](https://redirect.github.com/bytedance/sonic/pull/742)

#### New Contributors

-   [@&#8203;willem-bd](https://redirect.github.com/willem-bd) made their first contribution in [https://github.com/bytedance/sonic/pull/736](https://redirect.github.com/bytedance/sonic/pull/736)

**Full Changelog**: https://github.com/bytedance/sonic/compare/v1.12.8...v1.12.9

### [`v1.12.8`](https://redirect.github.com/bytedance/sonic/releases/tag/v1.12.8)

[Compare Source](https://redirect.github.com/bytedance/sonic/compare/v1.12.7...v1.12.8)

#### What's Changed

##### Feature

-   feat:(decoder) add option CaseSensitive by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [https://github.com/bytedance/sonic/pull/709](https://redirect.github.com/bytedance/sonic/pull/709)

##### Fix

-   fix:(loader) avoid race on `lastmoduledatap` with go plugin (only effects on byted-tango) by [@&#8203;AsterDY](https://redirect.github.com/AsterDY) in [https://github.com/bytedance/sonic/pull/707](https://redirect.github.com/bytedance/sonic/pull/707)
-   fix: update base64x version to avoid illegal instruction in sse by [@&#8203;liuq19](https://redirect.github.com/liuq19) in [https://github.com/bytedance/sonic/pull/734](https://redirect.github.com/bytedance/sonic/pull/734)

**Full Changelog**: https://github.com/bytedance/sonic/compare/v1.12.7...v1.12.8

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on April 06, 2025 at 06:00 AM UTC

Commits missing Jira IDs:
0a586bafec9d3eb19211a3ab9653799b4503749d


### Comment by @codecov-commenter on April 06, 2025 at 06:05 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1629?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.16%. Comparing base [(`6b415da`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/6b415da523ae7ed8b76398f4b41c7fd0e032085f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`0a586ba`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/0a586bafec9d3eb19211a3ab9653799b4503749d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1629   +/-   ##
=======================================
  Coverage   58.16%   58.16%           
=======================================
  Files         146      146           
  Lines       11333    11333           
=======================================
  Hits         6592     6592           
  Misses       4159     4159           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1629/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1629/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1629?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1629*
