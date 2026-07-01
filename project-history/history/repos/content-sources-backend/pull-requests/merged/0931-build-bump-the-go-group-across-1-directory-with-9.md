---
type: pull_request
number: 931
title: "Build: Bump the go group across 1 directory with 9 updates"
state: merged
author: dependabot
created: 2024-12-30T04:07:45Z
updated: 2025-01-02T01:41:00Z
closed: 2025-01-02T01:40:53Z
merged: 2025-01-02T01:40:53Z
base_branch: main
head_branch: dependabot/go_modules/go-2eef2f5c35
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/931
---

# Pull Request #931: Build: Bump the go group across 1 directory with 9 updates

**Author**: @dependabot
**Created**: December 30, 2024 at 04:07 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-2eef2f5c35`

## Description

Bumps the go group with 9 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/tang](https://github.com/content-services/tang) | `0.0.10` | `0.0.11` |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.13.2` | `4.13.3` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.43.3` | `1.44.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.6` | `1.32.7` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.47` | `1.17.48` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.45.0` | `1.45.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.12.1733335787` | `2024.12.1734541842` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.7.1` | `5.7.2` |
| [github.com/lzap/cloudwatchwriter2](https://github.com/lzap/cloudwatchwriter2) | `1.2.0` | `1.3.0` |


Updates `github.com/content-services/tang` from 0.0.10 to 0.0.11
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/tang/releases">github.com/content-services/tang's releases</a>.</em></p>
<blockquote>
<h2>v0.0.11</h2>
<h2>What's Changed</h2>
<ul>
<li>Build: Change to new pulp image by <a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/11">content-services/tang#11</a></li>
<li>Update to go 1.23 and update deps by <a href="https://github.com/rverdile"><code>@​rverdile</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/12">content-services/tang#12</a></li>
<li>Fixes 4740: Change base query to use inner joins by <a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/13">content-services/tang#13</a></li>
<li>Support go 1.22 by <a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/15">content-services/tang#15</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.8...v0.0.10">https://github.com/content-services/tang/compare/v0.0.8...v0.0.10</a></p>
<h2>What's Changed</h2>
<ul>
<li>Update action by <a href="https://github.com/Andrewgdewar"><code>@​Andrewgdewar</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/16">content-services/tang#16</a></li>
<li>Fixes 4932: Add list module Streams endpoint by <a href="https://github.com/Andrewgdewar"><code>@​Andrewgdewar</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/14">content-services/tang#14</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.10...v0.0.11">https://github.com/content-services/tang/compare/v0.0.10...v0.0.11</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/6507fd1961e880986615e0791336e71c908369a6"><code>6507fd1</code></a> Fixes 4932: Add list module Streams endpoint (<a href="https://redirect.github.com/content-services/tang/issues/14">#14</a>)</li>
<li><a href="https://github.com/content-services/tang/commit/169c0b1ad27144b00f47c9ce848b6c993df9f0d1"><code>169c0b1</code></a> Update action (<a href="https://redirect.github.com/content-services/tang/issues/16">#16</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.10...v0.0.11">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/labstack/echo/v4` from 4.13.2 to 4.13.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.13.3</h2>
<p><strong>Security</strong></p>
<ul>
<li>Update golang.org/x/net dependency <a href="https://pkg.go.dev/vuln/GO-2024-3333">GO-2024-3333</a> in <a href="https://redirect.github.com/labstack/echo/pull/2722">labstack/echo#2722</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.13.2...v4.13.3">https://github.com/labstack/echo/compare/v4.13.2...v4.13.3</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.13.3 - 2024-12-19</h2>
<p><strong>Security</strong></p>
<ul>
<li>Update golang.org/x/net dependency <a href="https://pkg.go.dev/vuln/GO-2024-3333">GO-2024-3333</a> in <a href="https://redirect.github.com/labstack/echo/pull/2722">labstack/echo#2722</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/45524e39d60d424d8ac175001eed70d6ace92930"><code>45524e3</code></a> Update golang.org/x/net dependency [GO-2024-3333](<a href="https://pkg.go.dev/vuln/GO-">https://pkg.go.dev/vuln/GO-</a>...</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.13.2...v4.13.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.43.3 to 1.44.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.44.0 (2024-12-27)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: update go directive to 1.20 by <a href="https://github.com/mauri870"><code>@​mauri870</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2933">IBM/sarama#2933</a></li>
<li>feat(producer): add retry buffer tuning option to prevent OOM by <a href="https://github.com/wanwenli"><code>@​wanwenli</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3026">IBM/sarama#3026</a></li>
<li>feat(admin): implement leader election api by <a href="https://github.com/chengjoey"><code>@​chengjoey</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3030">IBM/sarama#3030</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: log SASL connection and handshake errors by <a href="https://github.com/pierDipi"><code>@​pierDipi</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2995">IBM/sarama#2995</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump the golang-org-x group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3010">IBM/sarama#3010</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.28.0 to 0.31.0 in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3041">IBM/sarama#3041</a></li>
<li>chore(deps): bump the golang-org-x group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3040">IBM/sarama#3040</a></li>
<li>chore(deps): bump github.com/pierrec/lz4/v4 from 4.1.21 to 4.1.22 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3038">IBM/sarama#3038</a></li>
<li>chore(deps): bump the go_modules group across 2 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3035">IBM/sarama#3035</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.22.0 to 0.31.0 in /examples/consumergroup in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3033">IBM/sarama#3033</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.22.0 to 0.31.0 in /examples/txn_producer in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3034">IBM/sarama#3034</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.22.0 to 0.31.0 in /examples/interceptors in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3032">IBM/sarama#3032</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.22.0 to 0.31.0 in /examples/exactly_once in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3031">IBM/sarama#3031</a></li>
<li>chore(deps): bump github.com/stretchr/testify from 1.9.0 to 1.10.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3020">IBM/sarama#3020</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: add newer kafka versions and bump Go in CI by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2969">IBM/sarama#2969</a></li>
<li>fix(lint): resolve IDENTICAL_BRANCHES issue in broker by <a href="https://github.com/frzifus"><code>@​frzifus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2992">IBM/sarama#2992</a></li>
<li>chore: add version consts for 3.8.1+3.9.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3011">IBM/sarama#3011</a></li>
<li>fix(client): refactor duplicated replica+partition logic by <a href="https://github.com/Trinoooo"><code>@​Trinoooo</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2925">IBM/sarama#2925</a></li>
<li>chore(deps): bump golang.org/x/net to v0.33.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3044">IBM/sarama#3044</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/mauri870"><code>@​mauri870</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2933">IBM/sarama#2933</a></li>
<li><a href="https://github.com/frzifus"><code>@​frzifus</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2992">IBM/sarama#2992</a></li>
<li><a href="https://github.com/pierDipi"><code>@​pierDipi</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2995">IBM/sarama#2995</a></li>
<li><a href="https://github.com/wanwenli"><code>@​wanwenli</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3026">IBM/sarama#3026</a></li>
<li><a href="https://github.com/Trinoooo"><code>@​Trinoooo</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2925">IBM/sarama#2925</a></li>
<li><a href="https://github.com/chengjoey"><code>@​chengjoey</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3030">IBM/sarama#3030</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.43.3...v1.44.0">https://github.com/IBM/sarama/compare/v1.43.3...v1.44.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/1358013ce2dbdca2dacf1964b5c26c3beaa3fe1d"><code>1358013</code></a> feat(admin): implement leader election api</li>
<li><a href="https://github.com/IBM/sarama/commit/88fd71312bfba56f8eb30e8284a48cd20cffca2b"><code>88fd713</code></a> chore(ci): bump actions versions to latest</li>
<li><a href="https://github.com/IBM/sarama/commit/d604ce61a9c04e05116c46d64ce2004e4278cdea"><code>d604ce6</code></a> chore(deps): bump golang.org/x/net to v0.33.0</li>
<li><a href="https://github.com/IBM/sarama/commit/a17194df7d560172eebf4b2bcddbf481f238c949"><code>a17194d</code></a> chore(deps): bump github.com/stretchr/testify from 1.9.0 to 1.10.0</li>
<li><a href="https://github.com/IBM/sarama/commit/0757b96086e311f7485293a6fa195eaf68392a38"><code>0757b96</code></a> fix(client): refactor duplicated partition logic</li>
<li><a href="https://github.com/IBM/sarama/commit/1a2e1d35ba41f794ee68e0ebfd92ab4cd15156eb"><code>1a2e1d3</code></a> fix(client): refactor duplicated replica logic</li>
<li><a href="https://github.com/IBM/sarama/commit/7c75cc4c059f60cf515d8215c8eb50f2de268d47"><code>7c75cc4</code></a> feat(producer): add retry buffer tuning option to prevent OOM (<a href="https://redirect.github.com/IBM/sarama/issues/3026">#3026</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/91728dcd9e4834a83f86205e06a6b7ed626f02f1"><code>91728dc</code></a> chore(ci): bump actions/upload-artifact from 4.3.6 to 4.4.3</li>
<li><a href="https://github.com/IBM/sarama/commit/0f8a89662fa8b96e139e4231f9622f57b5cb58a8"><code>0f8a896</code></a> chore(deps): bump golang.org/x/crypto</li>
<li><a href="https://github.com/IBM/sarama/commit/f7dfef72d41db8ff795dc600e44c1e2ce92f8d40"><code>f7dfef7</code></a> chore(deps): bump golang.org/x/crypto</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.43.3...v1.44.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.6 to 1.32.7
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a964704cb2640ed57a74b9b37a53dcda7b6b7dd"><code>5a96470</code></a> Release 2024-12-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/653aa807b912e104f5e1e84e0510b4dffd76c751"><code>653aa80</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d02b239e896c5791e295c9a30a5281f56a8f7c39"><code>d02b239</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/698d709c21bc7922489aaba8c8207c9d7253c2fe"><code>698d709</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/885de40869f9bcee29ad11d60967aa0f1b571d46"><code>885de40</code></a> Fix improper use of Printf-style functions (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2934">#2934</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/858298a55393392fb161c5bd0ae3b9c5251996bf"><code>858298a</code></a> Release 2024-12-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f58264af808a255782999422056bccb06552dcbd"><code>f58264a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/df31082d87044a000a1524dbb654651f32713e10"><code>df31082</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/346690ed8f5b974ab26532aa93d5fa92a58d3571"><code>346690e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/45154546e09b48505c8798f7e5f3846ee1e0453a"><code>4515454</code></a> Release 2024-12-17</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.6...v1.32.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.47 to 1.17.48
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a964704cb2640ed57a74b9b37a53dcda7b6b7dd"><code>5a96470</code></a> Release 2024-12-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/653aa807b912e104f5e1e84e0510b4dffd76c751"><code>653aa80</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d02b239e896c5791e295c9a30a5281f56a8f7c39"><code>d02b239</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/698d709c21bc7922489aaba8c8207c9d7253c2fe"><code>698d709</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/885de40869f9bcee29ad11d60967aa0f1b571d46"><code>885de40</code></a> Fix improper use of Printf-style functions (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2934">#2934</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/858298a55393392fb161c5bd0ae3b9c5251996bf"><code>858298a</code></a> Release 2024-12-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f58264af808a255782999422056bccb06552dcbd"><code>f58264a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/df31082d87044a000a1524dbb654651f32713e10"><code>df31082</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/346690ed8f5b974ab26532aa93d5fa92a58d3571"><code>346690e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/45154546e09b48505c8798f7e5f3846ee1e0453a"><code>4515454</code></a> Release 2024-12-17</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.47...credentials/v1.17.48">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.0 to 1.45.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2de0027dc478a6ae80e9f2d24d904a425169a23b"><code>2de0027</code></a> Release 2023-11-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0c890c5eaf354ff23feb727ded9f50aaee9f1c4"><code>f0c890c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e032d9ea8d98d366f2467a72834d2cc0ee865edd"><code>e032d9e</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/507661ff1edbc896fbdfe3ea2e4c2e74be3b4e3c"><code>507661f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4128360684a451476e33c0f979921bc46ff63656"><code>4128360</code></a> fix: respect functional option modifications to RetryMaxAttempts (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2390">#2390</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d66ab2c239681700d947dfe859c3804377d0b4a8"><code>d66ab2c</code></a> Release 2023-11-27.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a264f87e1b9454dfa8a8f7226e6c72409fb3359d"><code>a264f87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d50b35dc1b0030e8ff5b1468b996d0779a25269"><code>7d50b35</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.45.0...service/s3/v1.45.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.12.1733335787 to 2024.12.1734541842
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/e72aa8ed88fe78c7b7af4ac526d63905a9556efd"><code>e72aa8e</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7ed23bab84c87e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/11fb0faeb839484e580fc9eb7abb9fac1b6f2556"><code>11fb0fa</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/21">#21</a> from jlsherrill/compose-update</li>
<li><a href="https://github.com/content-services/zest/commit/e1194590b5f0f41eeaaca10f2194dacf7076bfe2"><code>e119459</code></a> update docker compose for pulp</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.12.1733335787...release/v2024.12.1734541842">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.1 to 5.7.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.2 (December 21, 2024)</h1>
<ul>
<li>Fix prepared statement already exists on batch prepare failure</li>
<li>Add commit query to tx options (Lucas Hild)</li>
<li>Fix pgtype.Timestamp json unmarshal (Shean de Montigny-Desautels)</li>
<li>Add message body size limits in frontend and backend (zene)</li>
<li>Add xid8 type</li>
<li>Ensure planning encodes and scans cannot infinitely recurse</li>
<li>Implement pgtype.UUID.String() (Konstantin Grachev)</li>
<li>Switch from ExecParams to Exec in ValidateConnectTargetSessionAttrs functions (Alexander Rumyantsev)</li>
<li>Update golang.org/x/crypto</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/24fbe353ed5c3f53d379b3ae370229a5638a2868"><code>24fbe35</code></a> Create changelog for v5.7.2</li>
<li><a href="https://github.com/jackc/pgx/commit/3a1593b25b9de8b43e37d75ebc6d77589fe02001"><code>3a1593b</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2198">#2198</a> from alexandear/fix-nilness</li>
<li><a href="https://github.com/jackc/pgx/commit/9d851d7c98e255b25beaa69250a89cfe2f34f9ba"><code>9d851d7</code></a> Fix integration benchmarks</li>
<li><a href="https://github.com/jackc/pgx/commit/dacffdc7e22c41863ba80d96e5a2a53bf79c333c"><code>dacffdc</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2196">#2196</a> from alexandear/docs-improve-links</li>
<li><a href="https://github.com/jackc/pgx/commit/bc7c84077055c5af6f1e65d584523590a9a554df"><code>bc7c840</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2195">#2195</a> from LucasHild/master</li>
<li><a href="https://github.com/jackc/pgx/commit/043685147f04bf42465dd77497e91d3226895d52"><code>0436851</code></a> Handle errors in generate_certs</li>
<li><a href="https://github.com/jackc/pgx/commit/25329273dacf0b56fcb799d6dd0811a87f5a48ad"><code>2532927</code></a> Improve links in README</li>
<li><a href="https://github.com/jackc/pgx/commit/ad87d47089d16a2b59998c5eba94354ca6a9dca2"><code>ad87d47</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2194">#2194</a> from alexandear/refactor/pgconn-tests</li>
<li><a href="https://github.com/jackc/pgx/commit/7cf7bc6054ca2a4e2079864ee7bd0e8af75a24d6"><code>7cf7bc6</code></a> Simplify pgconn tests by using T.TempDir</li>
<li><a href="https://github.com/jackc/pgx/commit/3e6c719698b0a153fc96f1258024a1068ff3ffea"><code>3e6c719</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2189">#2189</a> from pankona/update-crypto</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.7.1...v5.7.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/lzap/cloudwatchwriter2` from 1.2.0 to 1.3.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/a109d00e23ef63bf83d60ba37a66b775fc33dc5e"><code>a109d00</code></a> fix: ensure time is monotonic</li>
<li>See full diff in <a href="https://github.com/lzap/cloudwatchwriter2/compare/v1.2.0...v1.3.0">compare view</a></li>
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

## Reviews

### Review by @swadeley - Approved on January 02, 2025 at 01:40 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/931*
