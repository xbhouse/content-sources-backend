---
type: pull_request
number: 642
title: "Build: Bump the go group with 4 updates"
state: closed
author: dependabot
created: 2024-04-22T04:23:24Z
updated: 2024-04-22T13:06:24Z
closed: 2024-04-22T13:06:22Z
base_branch: main
head_branch: dependabot/go_modules/go-2cefcd3aae
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/642
---

# Pull Request #642: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: April 22, 2024 at 04:23 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-2cefcd3aae`

## Description

Bumps the go group with 4 updates: [github.com/content-services/tang](https://github.com/content-services/tang), [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate), [github.com/labstack/echo/v4](https://github.com/labstack/echo) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/content-services/tang` from 0.0.4 to 0.0.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/tang/releases">github.com/content-services/tang's releases</a>.</em></p>
<blockquote>
<h2>v0.0.5</h2>
<h2>What's Changed</h2>
<ul>
<li>Add RpmRepositoryVersionErrataList to tangy by <a href="https://github.com/Andrewgdewar"><code>@​Andrewgdewar</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/6">content-services/tang#6</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Andrewgdewar"><code>@​Andrewgdewar</code></a> made their first contribution in <a href="https://redirect.github.com/content-services/tang/pull/6">content-services/tang#6</a></li>
<li><a href="https://github.com/xbhouse"><code>@​xbhouse</code></a> made their first contribution in <a href="https://redirect.github.com/content-services/tang/pull/6">content-services/tang#6</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.4...v0.0.5">https://github.com/content-services/tang/compare/v0.0.4...v0.0.5</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/b9362eb2cebb25d10fe4048d56a5ea85ba55d451"><code>b9362eb</code></a> Fixes 2941: Add RpmRepositoryVersionErrataList to tangy (<a href="https://redirect.github.com/content-services/tang/issues/6">#6</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.4...v0.0.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/golang-migrate/migrate/v4` from 4.17.0 to 4.17.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.17.1</h2>
<h2>Changelog</h2>
<ul>
<li>4bc6777 Add dktesting.Cleanup() method</li>
<li>bead4a9 Added documentation and test for lock strategy</li>
<li>129922a Added support for pgx locking table</li>
<li>a860f0c Bump github.com/dvsekhvalnov/jose2go from 1.5.0 to 1.6.0</li>
<li>d1df97b Bump github.com/jackc/pgx/v4 from 4.18.1 to 4.18.2</li>
<li>a78d1ab Bump github.com/jackc/pgx/v5 from 5.3.1 to 5.5.4</li>
<li>2e0872f Bump google.golang.org/protobuf from 1.31.0 to 1.33.0</li>
<li>1b707a7 Cleanup cassandra images after tests run</li>
<li>49cac86 Cleanup mongodb images after tests run</li>
<li>2884a8e Cleanup postgres images after tests run</li>
<li>b1d02e2 Cleanup sqlserver images after tests run</li>
<li>06614d9 Cleanup yugabytedb images after tests run</li>
<li>e913336 Drop support for Go 1.20 and add support for Go 1.22</li>
<li>f4950c1 Fallback to dktest.DefaultCleanupTimeout if the dktest.Options doesn't have one specified</li>
<li>5aa4670 Fix GoReleaser deprecations</li>
<li>d63a5c2 Only test against YugabyteDB LTS releases</li>
<li>091ad5d Quote locktable from config in queries</li>
<li>1a002d0 Set golangci-lint to 1.54.2 (latest is broken) (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1046">#1046</a>)</li>
<li>f100226 Update dktest from v0.4.0 to v0.4.1 to fix docker vulnerability</li>
<li>ff8a961 Update yugabyte test images</li>
<li>0350a00 [sqlserver] Always access version table with explicit schema</li>
<li>8147693 [sqlserver] Ensure version table in provided schema</li>
<li>7f85f9c chore: fix some typos</li>
<li>9d70a39 chore: fix some typos in comments</li>
<li>94b8fa5 rqlite is spelled with all lowercase</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/0c456c49a6debb0b2deb5a97a10d9e85d157b9d5"><code>0c456c4</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1068">#1068</a> from goodfirm/master</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/f1002267588a5617925665667251bdc24d276e1f"><code>f100226</code></a> Update dktest from v0.4.0 to v0.4.1 to fix docker vulnerability</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/c52377528f14f3b8004e571716fa10c0b79478bd"><code>c523775</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1055">#1055</a> from golang-migrate/dependabot/go_modules/github.com...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/a78d1abfa81a65e9ac972f242410656812756a03"><code>a78d1ab</code></a> Bump github.com/jackc/pgx/v5 from 5.3.1 to 5.5.4</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/837776f0c7880859623f09ff6487a70a303b6cbe"><code>837776f</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1054">#1054</a> from golang-migrate/dependabot/go_modules/google.gol...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/128b6509ccfc527e7bf490b75eff79bd170fd15f"><code>128b650</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1050">#1050</a> from golang-migrate/dependabot/go_modules/github.com...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/2e0872f3e8e61eb2f1af90b80749e134ca4b7188"><code>2e0872f</code></a> Bump google.golang.org/protobuf from 1.31.0 to 1.33.0</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/d1df97b2d1b80b99139de504f44e3efcf6de9555"><code>d1df97b</code></a> Bump github.com/jackc/pgx/v4 from 4.18.1 to 4.18.2</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/1a002d0e9e86a80c67409347498a7b5041238792"><code>1a002d0</code></a> Set golangci-lint to 1.54.2 (latest is broken) (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1046">#1046</a>)</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/2c5df8742b6f395056b42a435f461f6308f39ade"><code>2c5df87</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1072">#1072</a> from dhui/dktesting-cleanup</li>
<li>Additional commits viewable in <a href="https://github.com/golang-migrate/migrate/compare/v4.17.0...v4.17.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/labstack/echo/v4` from 4.11.4 to 4.12.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.12.0 - 2024-04-15</h2>
<p><strong>Security</strong></p>
<ul>
<li>Update golang.org/x/net dep because of <a href="https://pkg.go.dev/vuln/GO-2024-2687">GO-2024-2687</a> by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2625">labstack/echo#2625</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>binder: make binding to Map work better with string destinations by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2554">labstack/echo#2554</a></li>
<li>README.md: add Encore as sponsor by <a href="https://github.com/marcuskohlberg"><code>@​marcuskohlberg</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2579">labstack/echo#2579</a></li>
<li>Reorder paragraphs in README.md by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2581">labstack/echo#2581</a></li>
<li>CI: upgrade actions/checkout to v4 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2584">labstack/echo#2584</a></li>
<li>Remove default charset from 'application/json' Content-Type header by <a href="https://github.com/doortts"><code>@​doortts</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2568">labstack/echo#2568</a></li>
<li>CI: Use Go 1.22 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2588">labstack/echo#2588</a></li>
<li>binder: allow binding to a nil map by <a href="https://github.com/georgmu"><code>@​georgmu</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2574">labstack/echo#2574</a></li>
<li>Add Skipper Unit Test In BasicBasicAuthConfig and Add More Detail Explanation regarding BasicAuthValidator by <a href="https://github.com/RyoKusnadi"><code>@​RyoKusnadi</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2461">labstack/echo#2461</a></li>
<li>fix some typos by <a href="https://github.com/teslaedison"><code>@​teslaedison</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2603">labstack/echo#2603</a></li>
<li>fix: some typos by <a href="https://github.com/pomadev"><code>@​pomadev</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2596">labstack/echo#2596</a></li>
<li>Allow ResponseWriters to unwrap writers when flushing/hijacking by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2595">labstack/echo#2595</a></li>
<li>Add SPDX licence comments to files.  by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2604">labstack/echo#2604</a></li>
<li>Upgrade deps by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2605">labstack/echo#2605</a></li>
<li>Change type definition blocks to single declarations. This helps copy… by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2606">labstack/echo#2606</a></li>
<li>Fix Real IP logic by <a href="https://github.com/cl-bvl"><code>@​cl-bvl</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2550">labstack/echo#2550</a></li>
<li>Default binder can use <code>UnmarshalParams(params []string) error</code> inter… by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2607">labstack/echo#2607</a></li>
<li>Default binder can bind pointer to slice as struct field. For example  <code>*[]string</code> by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2608">labstack/echo#2608</a></li>
<li>Remove maxparam dependence from Context by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2611">labstack/echo#2611</a></li>
<li>When route is registered with empty path it is normalized to <code>/</code>.  by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2616">labstack/echo#2616</a></li>
<li>proxy middleware should use httputil.ReverseProxy for SSE requests by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2624">labstack/echo#2624</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/marcuskohlberg"><code>@​marcuskohlberg</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2579">labstack/echo#2579</a></li>
<li><a href="https://github.com/doortts"><code>@​doortts</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2568">labstack/echo#2568</a></li>
<li><a href="https://github.com/georgmu"><code>@​georgmu</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2574">labstack/echo#2574</a></li>
<li><a href="https://github.com/RyoKusnadi"><code>@​RyoKusnadi</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2461">labstack/echo#2461</a></li>
<li><a href="https://github.com/teslaedison"><code>@​teslaedison</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2603">labstack/echo#2603</a></li>
<li><a href="https://github.com/pomadev"><code>@​pomadev</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2596">labstack/echo#2596</a></li>
<li><a href="https://github.com/cl-bvl"><code>@​cl-bvl</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2550">labstack/echo#2550</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.11.4...v4.12.0">https://github.com/labstack/echo/compare/v4.11.4...v4.12.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.12.0 - 2024-04-15</h2>
<p><strong>Security</strong></p>
<ul>
<li>Update golang.org/x/net dep because of <a href="https://pkg.go.dev/vuln/GO-2024-2687">GO-2024-2687</a> by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2625">labstack/echo#2625</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>binder: make binding to Map work better with string destinations by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2554">labstack/echo#2554</a></li>
<li>README.md: add Encore as sponsor by <a href="https://github.com/marcuskohlberg"><code>@​marcuskohlberg</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2579">labstack/echo#2579</a></li>
<li>Reorder paragraphs in README.md by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2581">labstack/echo#2581</a></li>
<li>CI: upgrade actions/checkout to v4 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2584">labstack/echo#2584</a></li>
<li>Remove default charset from 'application/json' Content-Type header by <a href="https://github.com/doortts"><code>@​doortts</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2568">labstack/echo#2568</a></li>
<li>CI: Use Go 1.22 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2588">labstack/echo#2588</a></li>
<li>binder: allow binding to a nil map by <a href="https://github.com/georgmu"><code>@​georgmu</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2574">labstack/echo#2574</a></li>
<li>Add Skipper Unit Test In BasicBasicAuthConfig and Add More Detail Explanation regarding BasicAuthValidator by <a href="https://github.com/RyoKusnadi"><code>@​RyoKusnadi</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2461">labstack/echo#2461</a></li>
<li>fix some typos by <a href="https://github.com/teslaedison"><code>@​teslaedison</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2603">labstack/echo#2603</a></li>
<li>fix: some typos by <a href="https://github.com/pomadev"><code>@​pomadev</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2596">labstack/echo#2596</a></li>
<li>Allow ResponseWriters to unwrap writers when flushing/hijacking by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2595">labstack/echo#2595</a></li>
<li>Add SPDX licence comments to files.  by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2604">labstack/echo#2604</a></li>
<li>Upgrade deps by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2605">labstack/echo#2605</a></li>
<li>Change type definition blocks to single declarations. This helps copy… by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2606">labstack/echo#2606</a></li>
<li>Fix Real IP logic by <a href="https://github.com/cl-bvl"><code>@​cl-bvl</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2550">labstack/echo#2550</a></li>
<li>Default binder can use <code>UnmarshalParams(params []string) error</code> inter… by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2607">labstack/echo#2607</a></li>
<li>Default binder can bind pointer to slice as struct field. For example  <code>*[]string</code> by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2608">labstack/echo#2608</a></li>
<li>Remove maxparam dependence from Context by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2611">labstack/echo#2611</a></li>
<li>When route is registered with empty path it is normalized to <code>/</code>.  by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2616">labstack/echo#2616</a></li>
<li>proxy middleware should use httputil.ReverseProxy for SSE requests by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2624">labstack/echo#2624</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/88c379ff77278f553a0f3c44d27786b5a450b6e9"><code>88c379f</code></a> Changelog for v4.12.0 (<a href="https://redirect.github.com/labstack/echo/issues/2626">#2626</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/e0f2a02e4c12d0f75a2dba63851ecdc979afba94"><code>e0f2a02</code></a> proxy middleware should use http proxy for SSE requests (<a href="https://redirect.github.com/labstack/echo/issues/2624">#2624</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/306202540516a52f31acf9e24bcdb01c4734e1ef"><code>3062025</code></a> Update golang.org/x/* deps (<a href="https://redirect.github.com/labstack/echo/issues/2625">#2625</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/447c92d842e2f91c07eb032f619bea212beeee60"><code>447c92d</code></a> When route is registered with empty path it is normalized to <code>/</code>. Make sure t...</li>
<li><a href="https://github.com/labstack/echo/commit/d549290448fc65ae075f0b960b25489f6a38cb78"><code>d549290</code></a> Remove maxparam dependence from Context (<a href="https://redirect.github.com/labstack/echo/issues/2611">#2611</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/011acb4732fe4bdca1cb6d8ea9c29735e0b941f7"><code>011acb4</code></a> default binder can bind pointer to slice as struct field. For example `*[]str...</li>
<li><a href="https://github.com/labstack/echo/commit/c57fcb3746c4bfdab1b65363aa9e9edc7b6cab28"><code>c57fcb3</code></a> Default binder can use <code>UnmarshalParams(params []string) error</code> interface to ...</li>
<li><a href="https://github.com/labstack/echo/commit/a3b0ba24d3596de3bc2dde618125ac3acb1c774c"><code>a3b0ba2</code></a> Fix Real IP logic (<a href="https://redirect.github.com/labstack/echo/issues/2550">#2550</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/3598f295f95f316bbeb252b7b332fe34e120815c"><code>3598f29</code></a> Change type definition blocks to single declarations. This helps copy/pasting...</li>
<li><a href="https://github.com/labstack/echo/commit/5f7bedfb86e10bf0024236adfea544d0f5a82689"><code>5f7bedf</code></a> update makefile</li>
<li>Additional commits viewable in <a href="https://github.com/labstack/echo/compare/v4.11.4...v4.12.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.3.1710791249 to 2024.4.1713443852
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/3ba79db50bace1452715252021cf76b3c1e31795"><code>3ba79db</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e9a994d82c87e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/a5baa3e29b18af1fc633cc58e5447a39b0c2d0a9"><code>a5baa3e</code></a> Update pulp bindings to e69db48356e528a464be3da896237b4da4984dfa7d54eb5892a9d...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.3.1710791249...release/v2024.4.1713443852">compare view</a></li>
</ul>
</details>
<br />


Dependabot will resolve any conflicts with this PR as long as you don't alter it yourself. You can also trigger a rebase manually by commenting `@dependabot rebase`.

[//]: # (dependabot-automerge-start)
[//]: # (dependabot-automerge-end)

---

<details>
<summary>Dependabot commands and options</summary>
<br />

You can trigger Dependabot actions by commenting on this PR:
- `@dependabot rebase` will rebase this PR
- `@dependabot recreate` will recreate this PR, overwriting any edits that have been made to it
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Discussion

### Comment by @app-sre-bot on April 22, 2024 at 04:24 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on April 22, 2024 at 08:12 AM UTC

Lint errors related  to content guard

### Comment by @jlsherrill on April 22, 2024 at 12:36 PM UTC

updated here: https://github.com/content-services/content-sources-backend/pull/643

### Comment by @dependabot on April 22, 2024 at 01:06 PM UTC

Looks like these dependencies are no longer updatable, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/642*
