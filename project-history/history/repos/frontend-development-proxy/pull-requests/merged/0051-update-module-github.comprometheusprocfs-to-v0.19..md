---
type: pull_request
number: 51
title: "Update module github.com/prometheus/procfs to v0.19.2"
state: merged
author: red-hat-konflux
created: 2025-11-07T12:18:43Z
updated: 2025-12-04T08:19:57Z
closed: 2025-12-04T08:19:57Z
merged: 2025-12-04T08:19:57Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-prometheus-procfs-0.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/51
---

# Pull Request #51: Update module github.com/prometheus/procfs to v0.19.2

**Author**: @red-hat-konflux
**Created**: November 07, 2025 at 12:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-prometheus-procfs-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/procfs](https://redirect.github.com/prometheus/procfs) | `v0.16.1` -> `v0.19.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fprocfs/v0.19.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fprocfs/v0.16.1/v0.19.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

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

### [`v0.17.0`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.17.0)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.16.1...v0.17.0)

#### What's Changed

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;718](https://redirect.github.com/prometheus/procfs/pull/718)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;721](https://redirect.github.com/prometheus/procfs/pull/721)
- btrfs: correct allocation ratios for raid1c\[34] by [@&#8203;SimSaladin](https://redirect.github.com/SimSaladin) in [#&#8203;722](https://redirect.github.com/prometheus/procfs/pull/722)
- build(deps): bump golang.org/x/sync from 0.13.0 to 0.14.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;724](https://redirect.github.com/prometheus/procfs/pull/724)
- build(deps): bump golang.org/x/sys from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;723](https://redirect.github.com/prometheus/procfs/pull/723)
- Supports collection of process shared memory by [@&#8203;SilenceAdele](https://redirect.github.com/SilenceAdele) in [#&#8203;719](https://redirect.github.com/prometheus/procfs/pull/719)
- build(deps): bump golang.org/x/sync from 0.14.0 to 0.15.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;732](https://redirect.github.com/prometheus/procfs/pull/732)
- nvme: Add ControllerID output by [@&#8203;ShashwatHiregoudar](https://redirect.github.com/ShashwatHiregoudar) in [#&#8203;731](https://redirect.github.com/prometheus/procfs/pull/731)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;727](https://redirect.github.com/prometheus/procfs/pull/727)
- sysfs: Add support to collect link status for  PCIe devices by [@&#8203;naoki9911](https://redirect.github.com/naoki9911) in [#&#8203;728](https://redirect.github.com/prometheus/procfs/pull/728)
- nfs/parse.go: fix ClientV4Stats' GetDeviceInfo/LayoutGet -- values were swapped by [@&#8203;johnleslie](https://redirect.github.com/johnleslie) in [#&#8203;726](https://redirect.github.com/prometheus/procfs/pull/726)
- Fix linting issue by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;733](https://redirect.github.com/prometheus/procfs/pull/733)
- feat(mdstat): recognize reshape status by [@&#8203;tamcore](https://redirect.github.com/tamcore) in [#&#8203;679](https://redirect.github.com/prometheus/procfs/pull/679)
- Nvidia/Mellanox expose ROCE ECN information on sysfs on the path by [@&#8203;dasturiasArista](https://redirect.github.com/dasturiasArista) in [#&#8203;695](https://redirect.github.com/prometheus/procfs/pull/695)
- added zswap, zswapped, secpagetables, filehugepages, hugetlb and unaccepted to meminfo by [@&#8203;navidys](https://redirect.github.com/navidys) in [#&#8203;655](https://redirect.github.com/prometheus/procfs/pull/655)
- Parse StartCode, EndCode, and StartStack in `Proc.Stat()` by [@&#8203;pgimalac](https://redirect.github.com/pgimalac) in [#&#8203;659](https://redirect.github.com/prometheus/procfs/pull/659)
- Add node\_guid to infiniband class by [@&#8203;di3go-sona](https://redirect.github.com/di3go-sona) in [#&#8203;670](https://redirect.github.com/prometheus/procfs/pull/670)
- Fix linting issues by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;734](https://redirect.github.com/prometheus/procfs/pull/734)

#### New Contributors

- [@&#8203;SimSaladin](https://redirect.github.com/SimSaladin) made their first contribution in [#&#8203;722](https://redirect.github.com/prometheus/procfs/pull/722)
- [@&#8203;SilenceAdele](https://redirect.github.com/SilenceAdele) made their first contribution in [#&#8203;719](https://redirect.github.com/prometheus/procfs/pull/719)
- [@&#8203;ShashwatHiregoudar](https://redirect.github.com/ShashwatHiregoudar) made their first contribution in [#&#8203;731](https://redirect.github.com/prometheus/procfs/pull/731)
- [@&#8203;naoki9911](https://redirect.github.com/naoki9911) made their first contribution in [#&#8203;728](https://redirect.github.com/prometheus/procfs/pull/728)
- [@&#8203;johnleslie](https://redirect.github.com/johnleslie) made their first contribution in [#&#8203;726](https://redirect.github.com/prometheus/procfs/pull/726)
- [@&#8203;tamcore](https://redirect.github.com/tamcore) made their first contribution in [#&#8203;679](https://redirect.github.com/prometheus/procfs/pull/679)
- [@&#8203;navidys](https://redirect.github.com/navidys) made their first contribution in [#&#8203;655](https://redirect.github.com/prometheus/procfs/pull/655)
- [@&#8203;pgimalac](https://redirect.github.com/pgimalac) made their first contribution in [#&#8203;659](https://redirect.github.com/prometheus/procfs/pull/659)
- [@&#8203;di3go-sona](https://redirect.github.com/di3go-sona) made their first contribution in [#&#8203;670](https://redirect.github.com/prometheus/procfs/pull/670)

**Full Changelog**: <https://github.com/prometheus/procfs/compare/v0.16.1...v0.17.0>

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

### Review by @Hyperkid123 - Approved on December 04, 2025 at 08:19 AM UTC

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/51*
