---
type: pull_request
number: 48
title: "Update module github.com/prometheus/client_golang to v1.23.2"
state: closed
author: red-hat-konflux
created: 2025-11-06T12:22:05Z
updated: 2025-11-11T16:31:16Z
closed: 2025-11-11T14:19:23Z
base_branch: main
head_branch: konflux/mintmaker/main/github.com-prometheus-client_golang-1.x
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/48
---

# Pull Request #48: Update module github.com/prometheus/client_golang to v1.23.2

**Author**: @red-hat-konflux
**Created**: November 06, 2025 at 12:22 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `konflux/mintmaker/main/github.com-prometheus-client_golang-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/client_golang](https://redirect.github.com/prometheus/client_golang) | `v1.19.1` -> `v1.23.2` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fclient_golang/v1.23.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fclient_golang/v1.19.1/v1.23.2?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/client_golang (github.com/prometheus/client_golang)</summary>

### [`v1.23.2`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.23.2): - 2025-09-05

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.23.1...v1.23.2)

This release is made to upgrade to prometheus/common v0.66.1, which drops the dependencies github.com/grafana/regexp and go.uber.org/atomic and replaces gopkg.in/yaml.v2 with go.yaml.in/yaml/v2 (a drop-in replacement). There are no functional changes.

<details>
<summary>All Changes</summary>

- \[release-1.23] Upgrade to prometheus/common\@&#8203;v0.66.1 by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;1869](https://redirect.github.com/prometheus/client_golang/pull/1869)
- \[release-1.23] Cut v1.23.2 by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;1870](https://redirect.github.com/prometheus/client_golang/pull/1870)

</details>

**Full Changelog**: <https://github.com/prometheus/client_golang/compare/v1.23.1...v1.23.2>

### [`v1.23.1`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.23.1): - 2025-09-04

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.23.0...v1.23.1)

This release is made to be compatible with a backwards incompatible API change in prometheus/common v0.66.0. There are no functional changes.

<details>
<summary>All Changes</summary>

- \[release-1.23] Upgrade to prometheus/common v0.66 by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;1866](https://redirect.github.com/prometheus/client_golang/pull/1866)
- \[release-1.23] Cut v1.23.1 by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;1867](https://redirect.github.com/prometheus/client_golang/pull/1867)

</details>

**Full Changelog**: <https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.1>

### [`v1.23.0`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.23.0): - 2025-07-30

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.22.0...v1.23.0)

- \[CHANGE] Minimum required Go version is now 1.23, only the two latest Go versions are supported from now on. [#&#8203;1812](https://redirect.github.com/prometheus/client_golang/issues/1812)
- \[FEATURE] Add WrapCollectorWith and WrapCollectorWithPrefix [#&#8203;1766](https://redirect.github.com/prometheus/client_golang/issues/1766)
- \[FEATURE] Add exemplars for native histograms [#&#8203;1686](https://redirect.github.com/prometheus/client_golang/issues/1686)
- \[ENHANCEMENT] exp/api: Bubble up status code from writeResponse [#&#8203;1823](https://redirect.github.com/prometheus/client_golang/issues/1823)
- \[ENHANCEMENT] collector/go: Update runtime metrics for Go v1.23 and v1.24 [#&#8203;1833](https://redirect.github.com/prometheus/client_golang/issues/1833)
- \[BUGFIX] exp/api: client prompt return on context cancellation [#&#8203;1729](https://redirect.github.com/prometheus/client_golang/issues/1729)

<details>
<summary>All Changes</summary>
* Update example test by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1770
* build(deps): bump golang.org/x/net from 0.34.0 to 0.36.0 in /tutorials/whatsup by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1776
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1771
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1778
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1779
* build(deps): bump google.golang.org/protobuf from 1.36.5 to 1.36.6 in /exp by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1782
* build(deps): bump github.com/prometheus/common from 0.62.0 to 0.63.0 in /exp by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1781
* build(deps): bump github.com/prometheus/common from 0.62.0 to 0.63.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1783
* build(deps): bump google.golang.org/protobuf from 1.36.5 to 1.36.6 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1784
* build(deps): bump github.com/prometheus/procfs from 0.15.1 to 0.16.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1786
* chore: Upgrade golangci-lint to v2 by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1789
* build(deps): bump the github-actions group across 1 directory with 3 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1790
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1791
* Add `WrapCollectorWith` and `WrapCollectorWithPrefix` by @&#8203;colega in https://github.com/prometheus/client_golang/pull/1766
* feat(github-actions): add Go file change detection to golangci-lint workflow by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1794
* chore(ci): Fix flaky tests by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1795
* build(deps): bump golang.org/x/net from 0.36.0 to 0.38.0 in /tutorials/whatsup by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1799
* test(registry): Add goleak-based goroutine leak detection by @&#8203;surinkim in https://github.com/prometheus/client_golang/pull/1797
* build(deps): bump go.uber.org/goleak from 1.2.0 to 1.3.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1806
* build(deps): bump the github-actions group with 2 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1804
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1809
* Add exemplars for native histograms by @&#8203;shivanthzen in https://github.com/prometheus/client_golang/pull/1686
* build(deps): bump golang.org/x/sys from 0.30.0 to 0.32.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1807
* build(deps): bump github.com/prometheus/client_model from 0.6.1 to 0.6.2 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1805
* build(deps): bump github.com/prometheus/procfs from 0.16.0 to 0.16.1 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1808
* build(deps): bump golang.org/x/net from 0.35.0 to 0.38.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1800
* Update supported Go versions by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1812
* Cleaup Go modules by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1813
* fix: client prompt return on context cancellation by @&#8203;umegbewe in https://github.com/prometheus/client_golang/pull/1729
* Simplify buf binary install by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1814
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1815
* build(deps): bump the github-actions group with 5 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1817
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1821
* exp/api: Bubble up status code from writeResponse by @&#8203;saswatamcode in https://github.com/prometheus/client_golang/pull/1823
* build(deps): bump github.com/prometheus/common from 0.64.0 to 0.65.0 by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1827
* build(deps): bump github.com/prometheus/common from 0.64.0 to 0.65.0 in /exp by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1828
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1831
* Update runtime metrics for Go v1.23 and v1.24 by @&#8203;aknuds1 in https://github.com/prometheus/client_golang/pull/1833
* Fix `errNotImplemented` reference by @&#8203;aknuds1 in https://github.com/prometheus/client_golang/pull/1835
* build(deps): bump the github-actions group with 3 updates by @&#8203;dependabot[bot] in https://github.com/prometheus/client_golang/pull/1826
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1832
* Cut v1.23.0-rc.0 by @&#8203;vesari in https://github.com/prometheus/client_golang/pull/1837
* cut v1.23.0-rc.1 by @&#8203;vesari in https://github.com/prometheus/client_golang/pull/1842

</details>

#### New Contributors
* @&#8203;surinkim made their first contributi[https://github.com/prometheus/client_golang/pull/1797](https://redirect.github.com/prometheus/client_golang/pull/1797)l/1797
* @&#8203;umegbewe made their first contributi[https://github.com/prometheus/client_golang/pull/1729](https://redirect.github.com/prometheus/client_golang/pull/1729)l/1729
* @&#8203;aknuds1 made their first contributi[https://github.com/prometheus/client_golang/pull/1833](https://redirect.github.com/prometheus/client_golang/pull/1833)l/1833

**Full Changelog**: <https://github.com/prometheus/client_golang/compare/v1.22.0...v1.23.0>

### [`v1.22.0`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.22.0): - 2025-04-07

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.21.1...v1.22.0)

:warning: This release contains potential breaking change if you use experimental `zstd` support introduce in [#&#8203;1496](https://redirect.github.com/prometheus/client_golang/issues/1496) :warning:

Experimental support for `zstd` on scrape was added, controlled by the request `Accept-Encoding` header.
It was enabled by default since version 1.20, but now you need to add a blank import to enable it.
The decision to make it opt-in by default was originally made because the Go standard library was expected to have default zstd support added soon,
[golang/go#62513](https://redirect.github.com/golang/go/issues/62513) however, the work took longer than anticipated and it will be postponed to upcoming major Go versions.

e.g.:

> ```go
> import (
>   _ "github.com/prometheus/client_golang/prometheus/promhttp/zstd"
> )
> ```

- \[FEATURE] prometheus: Add new CollectorFunc utility [#&#8203;1724](https://redirect.github.com/prometheus/client_golang/issues/1724)
- \[CHANGE] Minimum required Go version is now 1.22 (we also test client\_golang against latest go version - 1.24) [#&#8203;1738](https://redirect.github.com/prometheus/client_golang/issues/1738)
- \[FEATURE] api: `WithLookbackDelta` and `WithStats` options have been added to API client. [#&#8203;1743](https://redirect.github.com/prometheus/client_golang/issues/1743)
- \[CHANGE] :warning: promhttp: Isolate zstd support and klauspost/compress library use to promhttp/zstd package. [#&#8203;1765](https://redirect.github.com/prometheus/client_golang/issues/1765)

<details>
<summary> All Changes </summary>

- build(deps): bump golang.org/x/sys from 0.28.0 to 0.29.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1720](https://redirect.github.com/prometheus/client_golang/pull/1720)
- build(deps): bump google.golang.org/protobuf from 1.36.1 to 1.36.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1719](https://redirect.github.com/prometheus/client_golang/pull/1719)
- Update RELEASE.md by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1721](https://redirect.github.com/prometheus/client_golang/pull/1721)
- chore(docs): Add links for the upstream PRs by [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) in [#&#8203;1722](https://redirect.github.com/prometheus/client_golang/pull/1722)
- Added tips on releasing client and checking with k8s. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1723](https://redirect.github.com/prometheus/client_golang/pull/1723)
- feat: Add new CollectorFunc utility by [@&#8203;Saumya40-codes](https://redirect.github.com/Saumya40-codes) in [#&#8203;1724](https://redirect.github.com/prometheus/client_golang/pull/1724)
- build(deps): bump google.golang.org/protobuf from 1.36.3 to 1.36.4 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1725](https://redirect.github.com/prometheus/client_golang/pull/1725)
- build(deps): bump the github-actions group with 5 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1726](https://redirect.github.com/prometheus/client_golang/pull/1726)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1727](https://redirect.github.com/prometheus/client_golang/pull/1727)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1731](https://redirect.github.com/prometheus/client_golang/pull/1731)
- build(deps): bump golang.org/x/sys from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1739](https://redirect.github.com/prometheus/client_golang/pull/1739)
- build(deps): bump google.golang.org/protobuf from 1.36.4 to 1.36.5 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1740](https://redirect.github.com/prometheus/client_golang/pull/1740)
- Cleanup dependabot config by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;1741](https://redirect.github.com/prometheus/client_golang/pull/1741)
- Upgrade Golang version v1.24 by [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) in [#&#8203;1738](https://redirect.github.com/prometheus/client_golang/pull/1738)
- build(deps): bump the github-actions group with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1742](https://redirect.github.com/prometheus/client_golang/pull/1742)
- Merging 1.21 release back to main. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1744](https://redirect.github.com/prometheus/client_golang/pull/1744)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1745](https://redirect.github.com/prometheus/client_golang/pull/1745)
- Add support for undocumented query options for API by [@&#8203;mahendrapaipuri](https://redirect.github.com/mahendrapaipuri) in [#&#8203;1743](https://redirect.github.com/prometheus/client_golang/pull/1743)
- exp/api: Add experimental exp module; Add remote API with write client and handler. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1658](https://redirect.github.com/prometheus/client_golang/pull/1658)
- exp/api: Add accepted msg type validation to handler by [@&#8203;saswatamcode](https://redirect.github.com/saswatamcode) in [#&#8203;1750](https://redirect.github.com/prometheus/client_golang/pull/1750)
- build(deps): bump the github-actions group with 5 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1751](https://redirect.github.com/prometheus/client_golang/pull/1751)
- build(deps): bump github.com/klauspost/compress from 1.17.11 to 1.18.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1752](https://redirect.github.com/prometheus/client_golang/pull/1752)
- build(deps): bump github.com/google/go-cmp from 0.6.0 to 0.7.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1753](https://redirect.github.com/prometheus/client_golang/pull/1753)
- exp: Reset snappy buf by [@&#8203;saswatamcode](https://redirect.github.com/saswatamcode) in [#&#8203;1756](https://redirect.github.com/prometheus/client_golang/pull/1756)
- Merge release 1.21.1 to main. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1762](https://redirect.github.com/prometheus/client_golang/pull/1762)
- exp: Add dependabot config by [@&#8203;saswatamcode](https://redirect.github.com/saswatamcode) in [#&#8203;1754](https://redirect.github.com/prometheus/client_golang/pull/1754)
- build(deps): bump peter-evans/create-pull-request from 7.0.7 to 7.0.8 in the github-actions group by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1764](https://redirect.github.com/prometheus/client_golang/pull/1764)
- promhttp: Isolate zstd support and klauspost/compress library use to promhttp/zstd package by [@&#8203;liggitt](https://redirect.github.com/liggitt) in [#&#8203;1765](https://redirect.github.com/prometheus/client_golang/pull/1765)
- Cut 1.22.0-rc.0 by [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) in [#&#8203;1768](https://redirect.github.com/prometheus/client_golang/pull/1768)

</details>

#### New Contributors
* @&#8203;Saumya40-codes made their first contributi[https://github.com/prometheus/client_golang/pull/1724](https://redirect.github.com/prometheus/client_golang/pull/1724)l/1724
* @&#8203;mahendrapaipuri made their first contributi[https://github.com/prometheus/client_golang/pull/1743](https://redirect.github.com/prometheus/client_golang/pull/1743)l/1743
* @&#8203;liggitt made their first contributi[https://github.com/prometheus/client_golang/pull/1765](https://redirect.github.com/prometheus/client_golang/pull/1765)l/1765

**Full Changelog**: <https://github.com/prometheus/client_golang/compare/v1.21.1...v1.22.0-rc.0>

### [`v1.21.1`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.21.1): / 2025-03-04

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.21.0...v1.21.1)

This release addresses a performance regression introduced in [#&#8203;1661](https://redirect.github.com/prometheus/client_golang/issues/1661) -- thanks to all who [reported this quickly](https://redirect.github.com/prometheus/client_golang/issues/1748):
[@&#8203;chlunde](https://redirect.github.com/chlunde), [@&#8203;dethi](https://redirect.github.com/dethi), [@&#8203;aaronbee](https://redirect.github.com/aaronbee) [@&#8203;tsuna](https://redirect.github.com/tsuna) [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun)  💪🏽. This patch release also fixes the iOS build.

We will be hardening the release process even further ([#&#8203;1759](https://redirect.github.com/prometheus/client_golang/issues/1759), [#&#8203;1761](https://redirect.github.com/prometheus/client_golang/issues/1761)) to prevent this in future, sorry for the inconvenience!

The high concurrency optimization is planned to be eventually reintroduced, however in a much safer manner, potentially in a separate API.

- \[BUGFIX] prometheus: Revert of `Inc`, `Add` and `Observe` cumulative metric CAS optimizations ([#&#8203;1661](https://redirect.github.com/prometheus/client_golang/issues/1661)), causing regressions on low concurrency cases [#&#8203;1757](https://redirect.github.com/prometheus/client_golang/issues/1757)
- \[BUGFIX] prometheus: Fix GOOS=ios build, broken due to process\_collector\_\* wrong build tags. [#&#8203;1758](https://redirect.github.com/prometheus/client_golang/issues/1758)

<details>
<summary>All commits</summary>

- Revert "exponential backoff for CAS operations on floats" and cut 1.21.1 by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1757](https://redirect.github.com/prometheus/client_golang/pull/1757)
- Fix ios build for 1.21.1 by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1758](https://redirect.github.com/prometheus/client_golang/pull/1758)

</details>

**Full Changelog**: <https://github.com/prometheus/client_golang/compare/v1.21.0...v1.21.1>

### [`v1.21.0`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.21.0): / 2025-02-19

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.5...v1.21.0)

:warning: This release contains potential breaking change if you upgrade `github.com/prometheus/common` to 0.62+ together with client\_golang (and depend on the strict, legacy validation for the label names). New common version [changes `model.NameValidationScheme` global variable](https://redirect.github.com/prometheus/common/pull/724), which relaxes the validation of label names and metric name, allowing all UTF-8 characters. Typically, this should not break any user, unless your test or usage expects strict certain names to panic/fail on client\_golang metric registration, gathering or scrape. In case of problems change `model.NameValidationScheme` to old `model.LegacyValidation` value in your project `init` function. :warning:

- \[BUGFIX] gocollector: Fix help message for runtime/metric metrics. [#&#8203;1583](https://redirect.github.com/prometheus/client_golang/issues/1583)
- \[BUGFIX] prometheus: Fix `Desc.String()` method for no labels case. [#&#8203;1687](https://redirect.github.com/prometheus/client_golang/issues/1687)
- \[PERF] prometheus: Optimize popular `prometheus.BuildFQName` function; now up to 30% faster. [#&#8203;1665](https://redirect.github.com/prometheus/client_golang/issues/1665)
- \[PERF] prometheus: Optimize `Inc`, `Add` and `Observe` cumulative metrics; now up to 50% faster under high concurrent contention. [#&#8203;1661](https://redirect.github.com/prometheus/client_golang/issues/1661)
- \[CHANGE] Upgrade prometheus/common to 0.62.0 which changes `model.NameValidationScheme` global variable. [#&#8203;1712](https://redirect.github.com/prometheus/client_golang/issues/1712)
- \[CHANGE] Add support for Go 1.23. [#&#8203;1602](https://redirect.github.com/prometheus/client_golang/issues/1602)
- \[FEATURE] process\_collector: Add support for Darwin systems. [#&#8203;1600](https://redirect.github.com/prometheus/client_golang/issues/1600) [#&#8203;1616](https://redirect.github.com/prometheus/client_golang/issues/1616) [#&#8203;1625](https://redirect.github.com/prometheus/client_golang/issues/1625) [#&#8203;1675](https://redirect.github.com/prometheus/client_golang/issues/1675) [#&#8203;1715](https://redirect.github.com/prometheus/client_golang/issues/1715)
- \[FEATURE] api: Add ability to invoke `CloseIdleConnections` on api.Client using `api.Client.(CloseIdler).CloseIdleConnections()` casting. [#&#8203;1513](https://redirect.github.com/prometheus/client_golang/issues/1513)
- \[FEATURE] promhttp: Add `promhttp.HandlerOpts.EnableOpenMetricsTextCreatedSamples` option to create OpenMetrics \_created lines. Not recommended unless you want to use opt-in Created Timestamp feature. Community works on OpenMetrics 2.0 format that should make those lines obsolete (they increase cardinality significantly). [#&#8203;1408](https://redirect.github.com/prometheus/client_golang/issues/1408)
- \[FEATURE] prometheus: Add `NewConstNativeHistogram` function. [#&#8203;1654](https://redirect.github.com/prometheus/client_golang/issues/1654)

<details>
<summary> All commits </summary>
* Merge release-1.20 to main by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1582
* gocollector: Tiny fix for help message with runtime/metrics source. by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1583
* ci: bump dagger to the latest version by @&#8203;marcosnils in https://github.com/prometheus/client_golang/pull/1588
* Merge release-1.20 back to main by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1593
* Update linting by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1603
* Update supported Go versions by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1602
* build(deps): bump golang.org/x/sys from 0.22.0 to 0.24.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1611
* build(deps): bump github.com/prometheus/common from 0.55.0 to 0.57.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1612
* changed the name of all variables with min/max name by @&#8203;parthlaw in https://github.com/prometheus/client_golang/pull/1606
* Update Dagger and build. by @&#8203;SuperQ in https://github.com/prometheus/client_golang/pull/1610
* build(deps): bump github/codeql-action from 3.25.15 to 3.26.6 in the github-actions group across 1 directory by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1614
* examples: Improved GoCollector example. by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1589
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1615
* process_collector: fill in most statistics on macOS by @&#8203;mharbison72 in https://github.com/prometheus/client_golang/pull/1600
* :zap: http client defer CloseIdleConnections by @&#8203;cuisongliu in https://github.com/prometheus/client_golang/pull/1513
* Set allow-utf-8 in Format during tests to avoid escaping. by @&#8203;ywwg in https://github.com/prometheus/client_golang/pull/1618
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1622
* Merge Release 1.20 back to main by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1627
* examples: Add custom labels example by @&#8203;ying-jeanne in https://github.com/prometheus/client_golang/pull/1626
* Refactor default runtime metrics tests for Go collector so that default runtime metric set autogenerates by @&#8203;vesari in https://github.com/prometheus/client_golang/pull/1631
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1628
* process_xxx_memory statistics for macOS (cgo) by @&#8203;mharbison72 in https://github.com/prometheus/client_golang/pull/1616
* build(deps): bump github.com/klauspost/compress from 1.17.9 to 1.17.10 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1633
* build(deps): bump golang.org/x/sys from 0.24.0 to 0.25.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1632
* process_collector: Add Platform-Specific Describe for processCollector by @&#8203;ying-jeanne in https://github.com/prometheus/client_golang/pull/1625
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1635
* build(deps): bump the github-actions group with 4 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1634
* Optionally print OM created lines by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1408
* process_collector: merge wasip1 and js into a single implementation by @&#8203;ying-jeanne in https://github.com/prometheus/client_golang/pull/1644
* Merge release 1.20 to main by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1647
* Add Arianna as maintainer 💪 by @&#8203;ArthurSens in https://github.com/prometheus/client_golang/pull/1651
* test add headers round tripper by @&#8203;Manask322 in https://github.com/prometheus/client_golang/pull/1657
* build(deps): bump github.com/klauspost/compress from 1.17.10 to 1.17.11 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1668
* build(deps): bump golang.org/x/sys from 0.25.0 to 0.26.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1669
* build(deps): bump github.com/prometheus/common from 0.59.1 to 0.60.1 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1667
* build(deps): bump google.golang.org/protobuf from 1.34.2 to 1.35.1 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1670
* Optimize BuildFQName function by @&#8203;jkroepke in https://github.com/prometheus/client_golang/pull/1665
* fix: use injected now() instead of time.Now() in summary methods by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1672
* process_collector: avoid a compiler warning on macOS (fixes #&#8203;1660) by @&#8203;mharbison72 in https://github.com/prometheus/client_golang/pull/1675
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1674
* build(deps): bump the github-actions group across 1 directory with 3 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1678
* [chore]: enable perfsprint linter by @&#8203;mmorel-35 in https://github.com/prometheus/client_golang/pull/1676
* Duplicate of #&#8203;1662 by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1673
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1679
* chore: enable usestdlibvars linter by @&#8203;mmorel-35 in https://github.com/prometheus/client_golang/pull/1680
* Add: exponential backoff for CAS operations on floats by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1661
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1683
* [1617] Add ConstnativeHistogram by @&#8203;shivanthzen in https://github.com/prometheus/client_golang/pull/1654
* fix: replace fmt.Errorf with errors.New by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1689
* Add codeowners by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1688
* fix: add very small delay between observations in `TestHistogramAtomicObserve` by @&#8203;imorph in https://github.com/prometheus/client_golang/pull/1691
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1692
* Fix: handle nil variableLabels in Desc.String() method and add tests for nil label values by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1687
* examples: Follow best practices and established naming conventions by @&#8203;lilic in https://github.com/prometheus/client_golang/pull/1650
* setup OSSF Scorecard workflow by @&#8203;mmorel-35 in https://github.com/prometheus/client_golang/pull/1432
* build(deps): bump google.golang.org/protobuf from 1.35.1 to 1.35.2 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1697
* build(deps): bump golang.org/x/sys from 0.26.0 to 0.27.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1696
* build(deps): bump the github-actions group with 5 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1695
* update links to openmetrics to reference the v1.0.0 release by @&#8203;dashpole in https://github.com/prometheus/client_golang/pull/1699
* build(deps): bump google.golang.org/protobuf from 1.35.2 to 1.36.1 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1706
* build(deps): bump golang.org/x/sys from 0.27.0 to 0.28.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1705
* build(deps): bump the github-actions group with 5 updates by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1707
* build(deps): bump github.com/prometheus/common from 0.60.1 to 0.61.0 by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1704
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1703
* Synchronize common files from prometheus/prometheus by @&#8203;prombot in https://github.com/prometheus/client_golang/pull/1708
* Upgrade to prometheus/common 0.62.0 with breaking change by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1712
* build(deps): bump golang.org/x/net from 0.26.0 to 0.33.0 in /tutorials/whatsup by @&#8203;dependabot in https://github.com/prometheus/client_golang/pull/1713
* docs: Add RELEASE.md for the release process by @&#8203;kakkoyun in https://github.com/prometheus/client_golang/pull/1690
* tutorials/whatsup: Updated deps by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1716
* process collector: Fixed pedantic registry failures on darwin with cgo. by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1715
* Revert "ci: daggerize test and lint pipelines (#&#8203;1534)" by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1717
* Cut 1.21.0-rc.0 by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1718
* Cut 1.21 by @&#8203;bwplotka in https://github.com/prometheus/client_golang/pull/1737
</details>

#### New Contributors
* @&#8203;parthlaw made their first contributi[https://github.com/prometheus/client_golang/pull/1606](https://redirect.github.com/prometheus/client_golang/pull/1606)l/1606
* @&#8203;mharbison72 made their first contributi[https://github.com/prometheus/client_golang/pull/1600](https://redirect.github.com/prometheus/client_golang/pull/1600)l/1600
* @&#8203;cuisongliu made their first contributi[https://github.com/prometheus/client_golang/pull/1513](https://redirect.github.com/prometheus/client_golang/pull/1513)l/1513
* @&#8203;ying-jeanne made their first contributi[https://github.com/prometheus/client_golang/pull/1626](https://redirect.github.com/prometheus/client_golang/pull/1626)l/1626
* @&#8203;Manask322 made their first contributi[https://github.com/prometheus/client_golang/pull/1657](https://redirect.github.com/prometheus/client_golang/pull/1657)l/1657
* @&#8203;jkroepke made their first contributi[https://github.com/prometheus/client_golang/pull/1665](https://redirect.github.com/prometheus/client_golang/pull/1665)l/1665
* @&#8203;imorph made their first contributi[https://github.com/prometheus/client_golang/pull/1672](https://redirect.github.com/prometheus/client_golang/pull/1672)l/1672
* @&#8203;mmorel-35 made their first contributi[https://github.com/prometheus/client_golang/pull/1676](https://redirect.github.com/prometheus/client_golang/pull/1676)l/1676
* @&#8203;shivanthzen made their first contributi[https://github.com/prometheus/client_golang/pull/1654](https://redirect.github.com/prometheus/client_golang/pull/1654)l/1654
* @&#8203;dashpole made their first contributi[https://github.com/prometheus/client_golang/pull/1699](https://redirect.github.com/prometheus/client_golang/pull/1699)l/1699

**Full Changelog**: <https://github.com/prometheus/client_golang/compare/v1.20.5...v1.21.0>

### [`v1.20.5`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.20.5): / 2024-10-15

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.4...v1.20.5)

We decided to revert [the `testutil` change](https://redirect.github.com/prometheus/client_golang/pull/1424) that made our util functions less error-prone, but created a lot of work for our downstream users. Apologies for the pain! This revert should not cause any major breaking change, even if you already did the work--unless you depend on the [exact error message](https://redirect.github.com/grafana/mimir/pull/9624#issuecomment-2413401565).

Going forward, we plan to reinforce our release testing strategy [\[1\]](https://redirect.github.com/prometheus/client_golang/issues/1646),[\[2\]](https://redirect.github.com/prometheus/client_golang/issues/1648) and deliver an enhanced [`testutil` package/module](https://redirect.github.com/prometheus/client_golang/issues/1639) with more flexible and safer APIs.

Thanks to [@&#8203;dashpole](https://redirect.github.com/dashpole) [@&#8203;dgrisonnet](https://redirect.github.com/dgrisonnet) [@&#8203;kakkoyun](https://redirect.github.com/kakkoyun) [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) [@&#8203;vesari](https://redirect.github.com/vesari) [@&#8203;logicalhan](https://redirect.github.com/logicalhan) [@&#8203;krajorama](https://redirect.github.com/krajorama) [@&#8203;bwplotka](https://redirect.github.com/bwplotka) who helped in this patch release! 🤗

##### Changelog

\[BUGFIX] testutil: Reverted [#&#8203;1424](https://redirect.github.com/prometheus/client_golang/issues/1424); functions using compareMetricFamilies are (again) only failing if filtered metricNames are in the expected input. [#&#8203;1645](https://redirect.github.com/prometheus/client_golang/issues/1645)

### [`v1.20.4`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.20.4)

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.3...v1.20.4)

- \[BUGFIX] histograms: Fix a possible data race when appending exemplars vs metrics gather. [#&#8203;1623](https://redirect.github.com/prometheus/client_golang/issues/1623)

### [`v1.20.3`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.20.3)

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.2...v1.20.3)

- \[BUGFIX] histograms: Fix possible data race when appending exemplars. [#&#8203;1608](https://redirect.github.com/prometheus/client_golang/issues/1608)

### [`v1.20.2`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.20.2)

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.1...v1.20.2)

- \[BUGFIX] promhttp: Unset Content-Encoding header when data is uncompressed. [#&#8203;1596](https://redirect.github.com/prometheus/client_golang/issues/1596)

### [`v1.20.1`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.20.1)

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.20.0...v1.20.1)

This release contains the critical fix for the [issue](https://redirect.github.com/prometheus/client_golang/issues/1584). Thanks to [@&#8203;geberl](https://redirect.github.com/geberl), [@&#8203;CubicrootXYZ](https://redirect.github.com/CubicrootXYZ), [@&#8203;zetaab](https://redirect.github.com/zetaab) and [@&#8203;timofurrer](https://redirect.github.com/timofurrer) for helping us with the investigation!

- \[BUGFIX] process-collector: Fixed unregistered descriptor error when using process collector with PedanticRegistry on Linux machines. [#&#8203;1587](https://redirect.github.com/prometheus/client_golang/issues/1587)

### [`v1.20.0`](https://redirect.github.com/prometheus/client_golang/releases/tag/v1.20.0)

[Compare Source](https://redirect.github.com/prometheus/client_golang/compare/v1.19.1...v1.20.0)

Thanks everyone for contributions!

:warning: In this release we remove one (broken anyway, given Go runtime changes) metric and add three new (representing GOGC, GOMEMLIMIT and GOMAXPROCS flags) to the default `collectors.NewGoCollector()` collector. Given its popular usage, expect your binary to expose two additional metric.

#### Changes

- \[CHANGE] :warning: go-collector: Remove `go_memstat_lookups_total` metric which was always 0; Go runtime stopped sharing pointer lookup statistics. [#&#8203;1577](https://redirect.github.com/prometheus/client_golang/issues/1577)
- \[FEATURE] :warning: go-collector: Add 3 default metrics: `go_gc_gogc_percent`, `go_gc_gomemlimit_bytes` and `go_sched_gomaxprocs_threads` as those are recommended by the Go team. [#&#8203;1559](https://redirect.github.com/prometheus/client_golang/issues/1559)
- \[FEATURE] go-collector: Add more information to all metrics' HELP e.g. the exact `runtime/metrics` sourcing each metric (if relevant). [#&#8203;1568](https://redirect.github.com/prometheus/client_golang/issues/1568) [#&#8203;1578](https://redirect.github.com/prometheus/client_golang/issues/1578)
- \[FEATURE] testutil: Add CollectAndFormat method. [#&#8203;1503](https://redirect.github.com/prometheus/client_golang/issues/1503)
- \[FEATURE] histograms: Add support for exemplars in native histograms. [#&#8203;1471](https://redirect.github.com/prometheus/client_golang/issues/1471)
- \[FEATURE] promhttp: Add experimental support for `zstd` on scrape, controlled by the request `Accept-Encoding` header. [#&#8203;1496](https://redirect.github.com/prometheus/client_golang/issues/1496)
- \[FEATURE] api/v1: Add `WithLimit` parameter to all API methods that supports it. [#&#8203;1544](https://redirect.github.com/prometheus/client_golang/issues/1544)
- \[FEATURE] prometheus: Add support for created timestamps in constant histograms and constant summaries. [#&#8203;1537](https://redirect.github.com/prometheus/client_golang/issues/1537)
- \[FEATURE] process-collectors: Add network usage metrics: `process_network_receive_bytes_total` and `process_network_transmit_bytes_total`. [#&#8203;1555](https://redirect.github.com/prometheus/client_golang/issues/1555)
- \[FEATURE] promlint: Add duplicated metric lint rule. [#&#8203;1472](https://redirect.github.com/prometheus/client_golang/issues/1472)
- \[BUGFIX] promlint: Relax metric type in name linter rule. [#&#8203;1455](https://redirect.github.com/prometheus/client_golang/issues/1455)
- \[BUGFIX] promhttp: Make sure server
  instrumentation wrapping supports new and future extra responseWriter methods. [#&#8203;1480](https://redirect.github.com/prometheus/client_golang/issues/1480)
- \[BUGFIX] testutil: Functions using compareMetricFamilies are now failing if filtered metricNames are not in the input. [#&#8203;1424](https://redirect.github.com/prometheus/client_golang/issues/1424)

<details>
  <summary>All commits</summary>

- feat(prometheus/testutil/promlint/validations): refine lintMetricType… by [@&#8203;foehammer127](https://redirect.github.com/foehammer127) in [#&#8203;1455](https://redirect.github.com/prometheus/client_golang/pull/1455)
- Bump github.com/prometheus/client\_golang from 1.18.0 to 1.19.0 in /examples/middleware by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1457](https://redirect.github.com/prometheus/client_golang/pull/1457)
- Bump github.com/prometheus/client\_model from 0.5.0 to 0.6.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1458](https://redirect.github.com/prometheus/client_golang/pull/1458)
- Bump golang.org/x/sys from 0.16.0 to 0.17.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1459](https://redirect.github.com/prometheus/client_golang/pull/1459)
- Bump github.com/prometheus/client\_golang from 1.18.0 to 1.19.0 in /tutorial/whatsup by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1461](https://redirect.github.com/prometheus/client_golang/pull/1461)
- Merge Release 1.19 back to main by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;1462](https://redirect.github.com/prometheus/client_golang/pull/1462)
- Bump the github-actions group with 2 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1456](https://redirect.github.com/prometheus/client_golang/pull/1456)
- Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1466](https://redirect.github.com/prometheus/client_golang/pull/1466)
- Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 in /examples/middleware by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1467](https://redirect.github.com/prometheus/client_golang/pull/1467)
- Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 in /tutorial/whatsup by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1469](https://redirect.github.com/prometheus/client_golang/pull/1469)
- Add LintDuplicateMetric to promlint by [@&#8203;bboreham](https://redirect.github.com/bboreham) in [#&#8203;1472](https://redirect.github.com/prometheus/client_golang/pull/1472)
- Auto-update Go Collector Metrics for new Go versions by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1476](https://redirect.github.com/prometheus/client_golang/pull/1476)
- Implement Unwrap() for responseWriterDelegator by [@&#8203;igor-drozdov](https://redirect.github.com/igor-drozdov) in [#&#8203;1480](https://redirect.github.com/prometheus/client_golang/pull/1480)
- Bump golang.org/x/sys from 0.17.0 to 0.18.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1485](https://redirect.github.com/prometheus/client_golang/pull/1485)
- Bump github.com/prometheus/procfs from 0.12.0 to 0.13.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1486](https://redirect.github.com/prometheus/client_golang/pull/1486)
- ci: Remove hardcoded supported Go versions from go.yml by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1489](https://redirect.github.com/prometheus/client_golang/pull/1489)
- feat: metrics generation workflow by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1481](https://redirect.github.com/prometheus/client_golang/pull/1481)
- fix: remove redundant go module in middleware example by [@&#8203;majolo](https://redirect.github.com/majolo) in [#&#8203;1492](https://redirect.github.com/prometheus/client_golang/pull/1492)
- chore: Refactor how base metrics are added to Sched metrics by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;1483](https://redirect.github.com/prometheus/client_golang/pull/1483)
- gocollector: Add regex option to allow collection of debug runtime metrics by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;1389](https://redirect.github.com/prometheus/client_golang/pull/1389)
- Bump github.com/prometheus/common from 0.48.0 to 0.52.3 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1498](https://redirect.github.com/prometheus/client_golang/pull/1498)
- chore: fix function name in comment by [@&#8203;oftenoccur](https://redirect.github.com/oftenoccur) in [#&#8203;1497](https://redirect.github.com/prometheus/client_golang/pull/1497)
- build(deps): bump golang.org/x/net from 0.20.0 to 0.23.0 in /tutorial/whatsup by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1501](https://redirect.github.com/prometheus/client_golang/pull/1501)
- build(deps): bump golang.org/x/net from 0.22.0 to 0.23.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1502](https://redirect.github.com/prometheus/client_golang/pull/1502)
- feat(dependency): replace go-spew package  by [@&#8203;dongjiang1989](https://redirect.github.com/dongjiang1989) in [#&#8203;1499](https://redirect.github.com/prometheus/client_golang/pull/1499)
- build(deps): bump github.com/prometheus/common from 0.52.3 to 0.53.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1504](https://redirect.github.com/prometheus/client_golang/pull/1504)
- build(deps): bump github.com/cespare/xxhash/v2 from 2.2.0 to 2.3.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1505](https://redirect.github.com/prometheus/client_golang/pull/1505)
- build(deps): bump google.golang.org/protobuf from 1.33.0 to 1.34.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1506](https://redirect.github.com/prometheus/client_golang/pull/1506)
- build(deps): bump golang.org/x/sys from 0.18.0 to 0.19.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1507](https://redirect.github.com/prometheus/client_golang/pull/1507)
- build(deps): bump github.com/prometheus/client\_model from 0.6.0 to 0.6.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1508](https://redirect.github.com/prometheus/client_golang/pull/1508)
- build(deps): bump github.com/prometheus/common from 0.48.0 to 0.53.0 in /tutorial/whatsup by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1509](https://redirect.github.com/prometheus/client_golang/pull/1509)
- improved code more clean by [@&#8203;lilijreey](https://redirect.github.com/lilijreey) in [#&#8203;1511](https://redirect.github.com/prometheus/client_golang/pull/1511)
- build(deps): bump the github-actions group with 3 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1510](https://redirect.github.com/prometheus/client_golang/pull/1510)
- \[CI]: Add Concurrency Grouping to GitHub Workflows by [@&#8203;Ishani217](https://redirect.github.com/Ishani217) in [#&#8203;1444](https://redirect.github.com/prometheus/client_golang/pull/1444)
- Add CollectAndFormat to testutil, allowing caller to assert as they want to on the exported metric  by [@&#8203;jcass8695](https://redirect.github.com/jcass8695) in [#&#8203;1503](https://redirect.github.com/prometheus/client_golang/pull/1503)
- testutil compareMetricFamilies: make less error-prone by [@&#8203;leonnicolas](https://redirect.github.com/leonnicolas) in [#&#8203;1424](https://redirect.github.com/prometheus/client_golang/pull/1424)
- improved code more clean use time.IsZero() replace t = time.Time{}   by [@&#8203;lilijreey](https://redirect.github.com/lilijreey) in [#&#8203;1515](https://redirect.github.com/prometheus/client_golang/pull/1515)
- add native histogram exemplar support by [@&#8203;fatsheep9146](https://redirect.github.com/fatsheep9146) in [#&#8203;1471](https://redirect.github.com/prometheus/client_golang/pull/1471)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1514](https://redirect.github.com/prometheus/client_golang/pull/1514)
- build(deps): bump golang.org/x/sys from 0.19.0 to 0.20.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1523](https://redirect.github.com/prometheus/client_golang/pull/1523)
- build(deps): bump google.golang.org/protobuf from 1.34.0 to 1.34.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1522](https://redirect.github.com/prometheus/client_golang/pull/1522)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1524](https://redirect.github.com/prometheus/client_golang/pull/1524)
- Add PR template for changelog automation by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1517](https://redirect.github.com/prometheus/client_golang/pull/1517)
- Auto label PRs by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1518](https://redirect.github.com/prometheus/client_golang/pull/1518)
- Fix: Auto label PRs [#&#8203;1518](https://redirect.github.com/prometheus/client_golang/issues/1518) by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1525](https://redirect.github.com/prometheus/client_golang/pull/1525)
- build(deps): bump github.com/prometheus/procfs from 0.13.0 to 0.15.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1527](https://redirect.github.com/prometheus/client_golang/pull/1527)
- ci: Group all changelog-related CI jobs into single one by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;1526](https://redirect.github.com/prometheus/client_golang/pull/1526)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1530](https://redirect.github.com/prometheus/client_golang/pull/1530)
- Remove synchronize trigger from changelog workflow by [@&#8203;SachinSahu431](https://redirect.github.com/SachinSahu431) in [#&#8203;1532](https://redirect.github.com/prometheus/client_golang/pull/1532)
- feat: Support zstd compression by [@&#8203;mrueg](https://redirect.github.com/mrueg) in [#&#8203;1496](https://redirect.github.com/prometheus/client_golang/pull/1496)
- Fix golangci-lint config by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;1536](https://redirect.github.com/prometheus/client_golang/pull/1536)
- build(deps): bump github.com/prometheus/client\_golang from 1.19.0 to 1.19.1 in /tutorial/whatsup by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1529](https://redirect.github.com/prometheus/client_golang/pull/1529)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1531](https://redirect.github.com/prometheus/client_golang/pull/1531)
- Cleanup NOTICE file by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;1541](https://redirect.github.com/prometheus/client_golang/pull/1541)
- Remove inlined upstream code by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;1539](https://redirect.github.com/prometheus/client_golang/pull/1539)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1545](https://redirect.github.com/prometheus/client_golang/pull/1545)
- client: Add Option to provide limit query param for APIs that support it by [@&#8203;abbyssoul](https://redirect.github.com/abbyssoul) in [#&#8203;1544](https://redirect.github.com/prometheus/client_golang/pull/1544)
- Allow creating constant histogram and summary metrics with a created timestamp by [@&#8203;swar8080](https://redirect.github.com/swar8080) in [#&#8203;1537](https://redirect.github.com/prometheus/client_golang/pull/1537)
- Update README.md by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1556](https://redirect.github.com/prometheus/client_golang/pull/1556)
- Temporarily remove required CI job for changelog. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1560](https://redirect.github.com/prometheus/client_golang/pull/1560)
- build(deps): bump github.com/prometheus/common from 0.53.0 to 0.55.0 in /tutorial/whatsup by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1549](https://redirect.github.com/prometheus/client_golang/pull/1549)
- build(deps): bump golang.org/x/sys from 0.20.0 to 0.21.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1552](https://redirect.github.com/prometheus/client_golang/pull/1552)
- build(deps): bump github.com/klauspost/compress from 1.17.8 to 1.17.9 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1553](https://redirect.github.com/prometheus/client_golang/pull/1553)
- fix: Update Go tests by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [#&#8203;1562](https://redirect.github.com/prometheus/client_golang/pull/1562)
- process\_collector: collect received/transmitted bytes by [@&#8203;huwcbjones](https://redirect.github.com/huwcbjones) in [#&#8203;1555](https://redirect.github.com/prometheus/client_golang/pull/1555)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1561](https://redirect.github.com/prometheus/client_golang/pull/1561)
- chore: Remove half-implemented changelog automation by [@&#8203;ArthurSens](https://redirect.github.com/ArthurSens) in [#&#8203;1564](https://redirect.github.com/prometheus/client_golang/pull/1564)
- build(deps): bump the github-actions group across 1 directory with 3 updates by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1565](https://redirect.github.com/prometheus/client_golang/pull/1565)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1563](https://redirect.github.com/prometheus/client_golang/pull/1563)
- build(deps): bump google.golang.org/protobuf from 1.34.1 to 1.34.2 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [#&#8203;1551](https://redirect.github.com/prometheus/client_golang/pull/1551)
- deps: Updated to prometheus/common to 0.55 by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;1566](https://redirect.github.com/prometheus/client_golang/pull/1566)
- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;1567](https://redirect.github.com/prometheus/client_golang/pull/1567)
- tutorials: Renamed tutorial -> tutorials for consistency + fixed tutorial code. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#

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

## Discussion

### Comment by @red-hat-konflux on November 06, 2025 at 12:22 PM UTC

### ℹ Artifact update notice

##### File name: rh_identity_transform/go.mod

In order to perform the update(s) described in the table above, Renovate ran the `go get` command, which resulted in the following additional change(s):


- 3 additional dependencies were updated


Details:


| **Package**                     | **Change**              |
| :------------------------------ | :---------------------- |
| `github.com/klauspost/compress` | `v1.17.11` -> `v1.18.0` |
| `github.com/prometheus/common`  | `v0.48.0` -> `v0.66.1`  |
| `github.com/prometheus/procfs`  | `v0.12.0` -> `v0.16.1`  |

### Comment by @red-hat-konflux on November 11, 2025 at 04:31 PM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`v1.23.2`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/48*
