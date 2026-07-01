---
type: pull_request
number: 1915
title: "Update module github.com/prometheus/procfs to v0.19.2"
state: merged
author: red-hat-konflux
created: 2025-11-03T08:19:14Z
updated: 2025-11-03T12:17:28Z
closed: 2025-11-03T08:28:35Z
merged: 2025-11-03T08:28:34Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-procfs-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1915
---

# Pull Request #1915: Update module github.com/prometheus/procfs to v0.19.2

**Author**: @red-hat-konflux
**Created**: November 03, 2025 at 08:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-procfs-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/procfs](https://redirect.github.com/prometheus/procfs) | `v0.17.0` -> `v0.19.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fprocfs/v0.19.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fprocfs/v0.17.0/v0.19.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>prometheus/procfs (github.com/prometheus/procfs)</summary>

### [`v0.19.2`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.19.2)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.19.1...v0.19.2)

#### What's Changed

- chore: Migrate tests to cmp package by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;760](https://redirect.github.com/prometheus/procfs/pull/760)
- Fix: Use base16 to convert pci sriov\_vf\_device by [@&#8203;jj-asama](https://redirect.github.com/jj-asama) in [#&#8203;762](https://redirect.github.com/prometheus/procfs/pull/762)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.19.1...v0.19.2>

### [`v0.19.1`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.19.1)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.19.0...v0.19.1)

#### What's Changed

- meminfo: Fix ZswappedBytes by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;759](https://redirect.github.com/prometheus/procfs/pull/759)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.19.0...v0.19.1>

### [`v0.19.0`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.19.0)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.18.0...v0.19.0)

#### What's Changed

- Add FS handler for mountinfo by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;757](https://redirect.github.com/prometheus/procfs/pull/757)
- sysfs: add link\_layer property to InfiniBandPort by [@&#8203;thomasbarrett](https://redirect.github.com/thomasbarrett) in [#&#8203;700](https://redirect.github.com/prometheus/procfs/pull/700)
- feat: expose MD raid component devices by [@&#8203;robbat2](https://redirect.github.com/robbat2) in [#&#8203;674](https://redirect.github.com/prometheus/procfs/pull/674)

#### New Contributors

- [@&#8203;robbat2](https://redirect.github.com/robbat2) made their first contribution in [#&#8203;674](https://redirect.github.com/prometheus/procfs/pull/674)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.18.0...v0.19.0>

### [`v0.18.0`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.18.0)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.17.0...v0.18.0)

#### What's Changed

- chore: clean up golangci-lint configuration by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;720](https://redirect.github.com/prometheus/procfs/pull/720)
- chore: enable errorlint linter by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;742](https://redirect.github.com/prometheus/procfs/pull/742)
- Modify proc\_statm\_test notes by [@&#8203;SilenceAdele](https://redirect.github.com/SilenceAdele) in [#&#8203;735](https://redirect.github.com/prometheus/procfs/pull/735)
- feat: Add hung\_task\_detect\_count by [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) in [#&#8203;738](https://redirect.github.com/prometheus/procfs/pull/738)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;736](https://redirect.github.com/prometheus/procfs/pull/736)
- net\_dev\_snmp6: directory traversal by [@&#8203;lzap](https://redirect.github.com/lzap) in [#&#8203;743](https://redirect.github.com/prometheus/procfs/pull/743)
- sysfs/class\_sas\_phy: Continue on EINVAL in parseSASPhy by [@&#8203;hrtbrock](https://redirect.github.com/hrtbrock) in [#&#8203;749](https://redirect.github.com/prometheus/procfs/pull/749)
- Add compact metrics to vmstat by [@&#8203;cbensimon](https://redirect.github.com/cbensimon) in [#&#8203;754](https://redirect.github.com/prometheus/procfs/pull/754)
- Update supported Go versions. by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;755](https://redirect.github.com/prometheus/procfs/pull/755)
- chore: enable several rules from go-critic by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [#&#8203;741](https://redirect.github.com/prometheus/procfs/pull/741)
- build(deps): bump golang.org/x/sync from 0.15.0 to 0.17.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;751](https://redirect.github.com/prometheus/procfs/pull/751)
- build(deps): bump golang.org/x/sys from 0.33.0 to 0.36.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;752](https://redirect.github.com/prometheus/procfs/pull/752)
- \[PROM-50] Update copyright headers by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;756](https://redirect.github.com/prometheus/procfs/pull/756)
- add sriov, power state and numa node info in PciDevice by [@&#8203;jj-asama](https://redirect.github.com/jj-asama) in [#&#8203;748](https://redirect.github.com/prometheus/procfs/pull/748)
- Fix Proc.Limits limit name matching by [@&#8203;inkel](https://redirect.github.com/inkel) in [#&#8203;667](https://redirect.github.com/prometheus/procfs/pull/667)
- add netfilter queue support by [@&#8203;KonstantinKuklin](https://redirect.github.com/KonstantinKuklin) in [#&#8203;677](https://redirect.github.com/prometheus/procfs/pull/677)

#### New Contributors

- [@&#8203;lzap](https://redirect.github.com/lzap) made their first contribution in [#&#8203;743](https://redirect.github.com/prometheus/procfs/pull/743)
- [@&#8203;hrtbrock](https://redirect.github.com/hrtbrock) made their first contribution in [#&#8203;749](https://redirect.github.com/prometheus/procfs/pull/749)
- [@&#8203;cbensimon](https://redirect.github.com/cbensimon) made their first contribution in [#&#8203;754](https://redirect.github.com/prometheus/procfs/pull/754)
- [@&#8203;jj-asama](https://redirect.github.com/jj-asama) made their first contribution in [#&#8203;748](https://redirect.github.com/prometheus/procfs/pull/748)
- [@&#8203;inkel](https://redirect.github.com/inkel) made their first contribution in [#&#8203;667](https://redirect.github.com/prometheus/procfs/pull/667)
- [@&#8203;KonstantinKuklin](https://redirect.github.com/KonstantinKuklin) made their first contribution in [#&#8203;677](https://redirect.github.com/prometheus/procfs/pull/677)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.17.0...v0.18.0>

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on November 03, 2025 at 08:24 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1915?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 58.96%. Comparing base ([`c2ea903`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/c2ea903c9ea6f813bdce71e02d54c381887dccd1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d1fe35a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d1fe35a97b983820efbdb3de1458a2f8bb8da3bb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1915   +/-   ##
=======================================
  Coverage   58.96%   58.96%           
=======================================
  Files         131      131           
  Lines        8407     8407           
=======================================
  Hits         4957     4957           
  Misses       2916     2916           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1915/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1915/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.96% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1915?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1915*
