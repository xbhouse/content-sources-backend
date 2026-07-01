---
type: pull_request
number: 930
title: "Build: Bump the go group with 7 updates"
state: closed
author: dependabot
created: 2024-12-23T04:54:54Z
updated: 2024-12-30T04:04:53Z
closed: 2024-12-30T04:04:52Z
base_branch: main
head_branch: dependabot/go_modules/go-fff0b05aee
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/930
---

# Pull Request #930: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: December 23, 2024 at 04:54 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-fff0b05aee`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/tang](https://github.com/content-services/tang) | `0.0.10` | `0.0.11` |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.13.2` | `4.13.3` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.6` | `1.32.7` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.47` | `1.17.48` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.45.0` | `1.45.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.12.1733335787` | `2024.12.1734541842` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.7.1` | `5.7.2` |

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

### Comment by @dependabot on December 30, 2024 at 04:04 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/930*
