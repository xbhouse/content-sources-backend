---
type: pull_request
number: 1078
title: "Update module github.com/prometheus/procfs to v0.16.0"
state: closed
author: red-hat-konflux
created: 2025-04-13T18:06:04Z
updated: 2025-10-07T12:16:28Z
closed: 2025-04-14T15:06:19Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-prometheus-procfs-0.x
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1078
---

# Pull Request #1078: Update module github.com/prometheus/procfs to v0.16.0

**Author**: @red-hat-konflux
**Created**: April 13, 2025 at 06:06 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-prometheus-procfs-0.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/prometheus/procfs](https://redirect.github.com/prometheus/procfs) | indirect | minor | `v0.15.1` -> `v0.16.0` |

---

### Release Notes

<details>
<summary>prometheus/procfs (github.com/prometheus/procfs)</summary>

### [`v0.16.0`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.16.0)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.15.1...v0.16.0)

#### What's Changed

-   enhancement: Expose CPU online status by [@&#8203;rexagod](https://redirect.github.com/rexagod) in [https://github.com/prometheus/procfs/pull/644](https://redirect.github.com/prometheus/procfs/pull/644)
-   enable gofmt and goimports linter by [@&#8203;mmorel-35](https://redirect.github.com/mmorel-35) in [https://github.com/prometheus/procfs/pull/646](https://redirect.github.com/prometheus/procfs/pull/646)
-   chore: deprecate `NewTCPx` methods by [@&#8203;rexagod](https://redirect.github.com/rexagod) in [https://github.com/prometheus/procfs/pull/640](https://redirect.github.com/prometheus/procfs/pull/640)
-   Add support for SELinux AVC statistics by [@&#8203;cgzones](https://redirect.github.com/cgzones) in [https://github.com/prometheus/procfs/pull/599](https://redirect.github.com/prometheus/procfs/pull/599)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/647](https://redirect.github.com/prometheus/procfs/pull/647)
-   Fix parsing NSpids field in /proc/{PID}/status by [@&#8203;timuralp](https://redirect.github.com/timuralp) in [https://github.com/prometheus/procfs/pull/648](https://redirect.github.com/prometheus/procfs/pull/648)
-   Bump golang.org/x/sys from 0.20.0 to 0.21.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/649](https://redirect.github.com/prometheus/procfs/pull/649)
-   Bump golang.org/x/sys from 0.21.0 to 0.22.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/656](https://redirect.github.com/prometheus/procfs/pull/656)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/652](https://redirect.github.com/prometheus/procfs/pull/652)
-   Bump golang.org/x/sys from 0.22.0 to 0.24.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/661](https://redirect.github.com/prometheus/procfs/pull/661)
-   Bump golang.org/x/sync from 0.7.0 to 0.8.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/660](https://redirect.github.com/prometheus/procfs/pull/660)
-   Add Disk IO stats and ext4 FS stats by [@&#8203;mshahzeb](https://redirect.github.com/mshahzeb) in [https://github.com/prometheus/procfs/pull/651](https://redirect.github.com/prometheus/procfs/pull/651)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/657](https://redirect.github.com/prometheus/procfs/pull/657)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/662](https://redirect.github.com/prometheus/procfs/pull/662)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/666](https://redirect.github.com/prometheus/procfs/pull/666)
-   Forbid print statements by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [https://github.com/prometheus/procfs/pull/668](https://redirect.github.com/prometheus/procfs/pull/668)
-   Update supported Go versions by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [https://github.com/prometheus/procfs/pull/669](https://redirect.github.com/prometheus/procfs/pull/669)
-   NetIPSocketSummary: Fix typo in godoc: UPD→UDP by [@&#8203;Miciah](https://redirect.github.com/Miciah) in [https://github.com/prometheus/procfs/pull/681](https://redirect.github.com/prometheus/procfs/pull/681)
-   Bump golang.org/x/sys from 0.25.0 to 0.26.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/682](https://redirect.github.com/prometheus/procfs/pull/682)
-   feat(blockdevice): added sysblockdevicesize method and test by [@&#8203;fs185143](https://redirect.github.com/fs185143) in [https://github.com/prometheus/procfs/pull/658](https://redirect.github.com/prometheus/procfs/pull/658)
-   Bump golang.org/x/sync from 0.8.0 to 0.9.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/683](https://redirect.github.com/prometheus/procfs/pull/683)
-   Bump golang.org/x/sys from 0.26.0 to 0.28.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/688](https://redirect.github.com/prometheus/procfs/pull/688)
-   feat: Read PCIE AER counters class/net by [@&#8203;dasturiasArista](https://redirect.github.com/dasturiasArista) in [https://github.com/prometheus/procfs/pull/686](https://redirect.github.com/prometheus/procfs/pull/686)
-   Bump golang.org/x/sync from 0.9.0 to 0.10.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/690](https://redirect.github.com/prometheus/procfs/pull/690)
-   Enhancement: Implement classes to read drm subsystem by [@&#8203;jritter](https://redirect.github.com/jritter) in [https://github.com/prometheus/procfs/pull/654](https://redirect.github.com/prometheus/procfs/pull/654)
-   Bump golang.org/x/sys from 0.28.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/697](https://redirect.github.com/prometheus/procfs/pull/697)
-   fix: correct linters-settings typo in .golangci.yml by [@&#8203;thevilledev](https://redirect.github.com/thevilledev) in [https://github.com/prometheus/procfs/pull/696](https://redirect.github.com/prometheus/procfs/pull/696)
-   Add support for per-interface SNMP6 stats by [@&#8203;Vascko](https://redirect.github.com/Vascko) in [https://github.com/prometheus/procfs/pull/693](https://redirect.github.com/prometheus/procfs/pull/693)
-   build(deps): bump golang.org/x/sync from 0.10.0 to 0.11.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/701](https://redirect.github.com/prometheus/procfs/pull/701)
-   build(deps): bump github.com/google/go-cmp from 0.6.0 to 0.7.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/702](https://redirect.github.com/prometheus/procfs/pull/702)
-   build(deps): bump golang.org/x/sys from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/703](https://redirect.github.com/prometheus/procfs/pull/703)

#### New Contributors

-   [@&#8203;cgzones](https://redirect.github.com/cgzones) made their first contribution in [https://github.com/prometheus/procfs/pull/599](https://redirect.github.com/prometheus/procfs/pull/599)
-   [@&#8203;timuralp](https://redirect.github.com/timuralp) made their first contribution in [https://github.com/prometheus/procfs/pull/648](https://redirect.github.com/prometheus/procfs/pull/648)
-   [@&#8203;mshahzeb](https://redirect.github.com/mshahzeb) made their first contribution in [https://github.com/prometheus/procfs/pull/651](https://redirect.github.com/prometheus/procfs/pull/651)
-   [@&#8203;Miciah](https://redirect.github.com/Miciah) made their first contribution in [https://github.com/prometheus/procfs/pull/681](https://redirect.github.com/prometheus/procfs/pull/681)
-   [@&#8203;fs185143](https://redirect.github.com/fs185143) made their first contribution in [https://github.com/prometheus/procfs/pull/658](https://redirect.github.com/prometheus/procfs/pull/658)
-   [@&#8203;dasturiasArista](https://redirect.github.com/dasturiasArista) made their first contribution in [https://github.com/prometheus/procfs/pull/686](https://redirect.github.com/prometheus/procfs/pull/686)
-   [@&#8203;jritter](https://redirect.github.com/jritter) made their first contribution in [https://github.com/prometheus/procfs/pull/654](https://redirect.github.com/prometheus/procfs/pull/654)
-   [@&#8203;thevilledev](https://redirect.github.com/thevilledev) made their first contribution in [https://github.com/prometheus/procfs/pull/696](https://redirect.github.com/prometheus/procfs/pull/696)
-   [@&#8203;Vascko](https://redirect.github.com/Vascko) made their first contribution in [https://github.com/prometheus/procfs/pull/693](https://redirect.github.com/prometheus/procfs/pull/693)

**Full Changelog**: https://github.com/prometheus/procfs/compare/v0.15.1...v0.16.0

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYWluIiwibGFiZWxzIjpbXX0=-->


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1078*
