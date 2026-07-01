---
type: pull_request
number: 514
title: "Build: Bump the go group with 8 updates"
state: closed
author: dependabot
created: 2023-12-25T04:51:24Z
updated: 2024-01-01T04:38:22Z
closed: 2024-01-01T04:38:21Z
base_branch: main
head_branch: dependabot/go_modules/go-2317ad8384
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/514
---

# Pull Request #514: Build: Bump the go group with 8 updates

**Author**: @dependabot
**Created**: December 25, 2023 at 04:51 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-2317ad8384`

## Description

Bumps the go group with 8 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.8` | `1.0.9` |
| [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) | `4.16.2` | `4.17.0` |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.11.3` | `4.11.4` |
| [github.com/archdx/zerolog-sentry](https://github.com/archdx/zerolog-sentry) | `1.7.0` | `1.8.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.16.12` | `1.16.13` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.30.0` | `1.30.1` |
| [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) | `2023.12.1702603647` | `2023.12.1703173361` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.3.0` | `9.3.1` |

Updates `github.com/content-services/yummy` from 1.0.8 to 1.0.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/f6daea7657a2acf97e6eb8410bf6fd0c4fa18aab"><code>f6daea7</code></a> Bump golang.org/x/crypto from 0.13.0 to 0.17.0 (<a href="https://redirect.github.com/content-services/yummy/issues/20">#20</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/9835a459c0174dfdaaf322d99ed868d95c1de06e"><code>9835a45</code></a> Add golangci config file and remove CI tests (<a href="https://redirect.github.com/content-services/yummy/issues/19">#19</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.8...v1.0.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/golang-migrate/migrate/v4` from 4.16.2 to 4.17.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.17.0</h2>
<h2>Changelog</h2>
<ul>
<li>cf03803 Add rqlite 8.0.0 to tested database versions</li>
<li>12968a7 Add syntax highlighting to Postgres example</li>
<li>50112e7 Add to clickhouse README.md database creation</li>
<li>5ded96d Bump golang.org/x/crypto from 0.14.0 to 0.17.0</li>
<li>c3ebd52 Bump google.golang.org/grpc from 1.55.0 to 1.56.3</li>
<li>5026488 Clean up require directive grouping</li>
<li>3b02b18 Correct a spelling mistake</li>
<li>cd17c5a Drop support for Go 1.19 and add support for Go 1.21</li>
<li>839421e Leverage quoteIdentifier from pgx</li>
<li>bad30b5 Mention migradaptor</li>
<li>fb22436 Merge remote-tracking branch 'origin/master' into upgrade-spanner</li>
<li>bfedabb Merge remote-tracking branch 'upstream/master'</li>
<li>92dec35 Move supported go version to standard place</li>
<li>4078ef8 New release prep</li>
<li>9fe7383 Quote in drop as well</li>
<li>691f687 Reformat ScyllaDB/Cassandra docs</li>
<li>90a3ac4 Remove cluster adaptation for tables to pass tests</li>
<li>64755d0 Update README.md</li>
<li>f2c4b52 Update aws-sdk-go from v1.44.301 to v1.49.6</li>
<li>876a13d Update aws-sdk-go to adress vulerabilitiy</li>
<li>b567287 Update from alpine 3.18 to 3.19</li>
<li>f2e0b33 Update lib/pq to fix cert permissions issues</li>
<li>208ac53 Update spanner to fix security issue See also: <a href="https://redirect.github.com/golang-migrate/migrate/pull/952">golang-migrate/migrate#952</a></li>
<li>72957b6 Updated version of spanner to support sequences and generate uuid</li>
<li>7d03609 add 8.11 and 8.12 versions and remove debug logging</li>
<li>7a72550 add tests for scylladb. add scylladb to docs</li>
<li>90273fe clickhouse: Quote db name in ensureVersionTable</li>
<li>5163ac7 feature: add rqlite support</li>
<li>ee8a8e5 fix: typo</li>
<li>f8afa5a small changes to retry failed by timeout CI</li>
<li>669437c update rqlite 8 container version to 8.0.6</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/cd17c5a808d1889d17d73a68355c18d1dea6c49d"><code>cd17c5a</code></a> Drop support for Go 1.19 and add support for Go 1.21</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/f2c4b52679390693211279658779107d41d0e872"><code>f2c4b52</code></a> Update aws-sdk-go from v1.44.301 to v1.49.6</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/c7c50118a0139d665c27ae77877f4f727a44df7f"><code>c7c5011</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/956">#956</a> from Kenai/master</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/9c551d43982bf4679ba13bb27138a2072c557716"><code>9c551d4</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/875">#875</a> from no-name16/clickhouse_create_database_migrations</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/ab24e7674d1ec618d5b5f891dd6e278e53c1af7f"><code>ab24e76</code></a> Merge branch 'master' into clickhouse_create_database_migrations</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/e8edcdce660b71a30735b21e928a18063236075e"><code>e8edcdc</code></a> Merge branch 'master' into master</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/691f687eb895737efe55c9f11b7cef8fc4e87d1e"><code>691f687</code></a> Reformat ScyllaDB/Cassandra docs</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/eb64ffa1ad92c145b4e03dc8abe417899daa2aed"><code>eb64ffa</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/947">#947</a> from mkorolyov/scylladb_support</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/a94396c3a811ffc101a27b7a913de836ee01f65b"><code>a94396c</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/857">#857</a> from luca-nardelli/master</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/0ba6fc3166450652df00a266d32f4c83c51750c9"><code>0ba6fc3</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1002">#1002</a> from jsabbatini-upguard/upgrade-spanner</li>
<li>Additional commits viewable in <a href="https://github.com/golang-migrate/migrate/compare/v4.16.2...v4.17.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/labstack/echo/v4` from 4.11.3 to 4.11.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.11.4 upgrade dependencies</h2>
<p><strong>Security</strong></p>
<ul>
<li>Upgrade golang.org/x/crypto to v0.17.0 to fix vulnerability <a href="https://pkg.go.dev/vuln/GO-2023-2402">issue</a> <a href="https://redirect.github.com/labstack/echo/pull/2562">#2562</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Update deps and mark Go version to 1.18 as this is what golang.org/x/* use <a href="https://redirect.github.com/labstack/echo/pull/2563">#2563</a></li>
<li>Request logger: add example for Slog <a href="https://pkg.go.dev/log/slog">https://pkg.go.dev/log/slog</a> <a href="https://redirect.github.com/labstack/echo/pull/2543">#2543</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.11.4 - 2023-12-20</h2>
<p><strong>Security</strong></p>
<ul>
<li>Upgrade golang.org/x/crypto to v0.17.0 to fix vulnerability <a href="https://pkg.go.dev/vuln/GO-2023-2402">issue</a> <a href="https://redirect.github.com/labstack/echo/pull/2562">#2562</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Update deps and mark Go version to 1.18 as this is what golang.org/x/* use <a href="https://redirect.github.com/labstack/echo/pull/2563">#2563</a></li>
<li>Request logger: add example for Slog <a href="https://pkg.go.dev/log/slog">https://pkg.go.dev/log/slog</a> <a href="https://redirect.github.com/labstack/echo/pull/2543">#2543</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/226e4f076a69de85b71cf059d8a3c0fa8feafcaf"><code>226e4f0</code></a> Changelog for v4.11.4 (<a href="https://redirect.github.com/labstack/echo/issues/2564">#2564</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/209c6a199af0d6443f640528351064ba31b5f864"><code>209c6a1</code></a> Update deps and mark Go version to 1.18 as this is what golang.org/x/* use. (...</li>
<li><a href="https://github.com/labstack/echo/commit/287a82c228efce23fac50e84d37e8690896bf5a5"><code>287a82c</code></a> Upgrade golang.org/x/crypto to v0.17.0 to fix vulnerability issue (<a href="https://redirect.github.com/labstack/echo/issues/2562">#2562</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/584cb85a6b749846ac26a8cd151244ab281f2abc"><code>584cb85</code></a> request logger: add example for Slog <a href="https://pkg.go.dev/log/slog">https://pkg.go.dev/log/slog</a> (<a href="https://redirect.github.com/labstack/echo/issues/2543">#2543</a>)</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.11.3...v4.11.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/archdx/zerolog-sentry` from 1.7.0 to 1.8.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/a41ff05a5b597c45a941e06b8f0d6dca50ac403f"><code>a41ff05</code></a> Merge pull request <a href="https://redirect.github.com/archdx/zerolog-sentry/issues/14">#14</a> from etkecc/master</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/3f53baa3c58d64eca2db6317a479a17fb27d8184"><code>3f53baa</code></a> add breadcrumbs support</li>
<li>See full diff in <a href="https://github.com/archdx/zerolog-sentry/compare/v1.7.0...v1.8.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.12 to 1.16.13
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d47c319e3287ca645b14f72ab80e0dbd42f1bfbf"><code>d47c319</code></a> Release 2022-08-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d0933debf440808caec8914d686a7da6cbc7a08f"><code>d0933de</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/988edc87e10b6ec43ceb9d9ccbbb1426792fe18e"><code>988edc8</code></a> Update SDK's smithy-go dependency to v1.13.1</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d6e010073bfe049472ec2f8389bff50d8411257a"><code>d6e0100</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f8a5aa28593b79d5f43c2cd6f19ba5a2e65ee7d4"><code>f8a5aa2</code></a> Release 2022-08-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8072564c2dcb9c303d132b992f60581bd99da3a1"><code>8072564</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9f60590930ad361560059f3dee4ce5f6770964f0"><code>9f60590</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8874a3eb1498c42fae4abe421e52205dac8c9ebc"><code>8874a3e</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.12...v1.16.13">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.30.0 to 1.30.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f95afa96016c325b62fdfe1af320bf30bee5f779"><code>f95afa9</code></a> Release 2023-01-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/77f1f40a3cac8e545acc6f9ecdbc931a81ef57ec"><code>77f1f40</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/298fc3272d4f7b46a27a528af38d0104d025f8d3"><code>298fc32</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bfcdd49c999e4e0911126ac793c1b6e3a038e8de"><code>bfcdd49</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00925a868bfc43dafa526259389ce44799c1b029"><code>00925a8</code></a> Release 2023-01-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/346bddbedc2b580f0181058be355c7bce5587bf1"><code>346bddb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7facad1bf8549631a5bafa6056f54b2473907544"><code>7facad1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/247b038d92b5cd4c466c887a3d5653fa36b6cbe5"><code>247b038</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0691f2ce89932a87343c737b69d3f56f54e39f5"><code>b0691f2</code></a> Release 2023-01-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a847e5fbd346b1944e70d2c6d5e0f96c553e5c74"><code>a847e5f</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.30.0...service/s3/v1.30.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2023` from 2023.12.1702603647 to 2023.12.1703173361
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/f4f5f8eeba01310fcce99710186a74a3513cc761"><code>f4f5f8e</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d29633bd5c57952eab83695d...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2023.12.1702603647...release/v2023.12.1703173361">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.3.0 to 9.3.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.3.1</h2>
<h2>Updates and Enhancements</h2>
<ul>
<li>Documentation and examples have been updated (<a href="https://redirect.github.com/redis/go-redis/issues/2806">#2806</a>).</li>
<li>Redis values can now be scanned into pointer fields (<a href="https://redirect.github.com/redis/go-redis/issues/2787">#2787</a>).</li>
<li>The URL format error in the Documentation has been corrected (<a href="https://redirect.github.com/redis/go-redis/issues/2789">#2789</a>).</li>
<li>Cmder annotation has been added (<a href="https://redirect.github.com/redis/go-redis/issues/2816">#2816</a>).</li>
<li>The Z member type has been changed to string (<a href="https://redirect.github.com/redis/go-redis/issues/2818">#2818</a>).</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>BITFIELD_RO Command has been introduced (<a href="https://redirect.github.com/redis/go-redis/issues/2820">#2820</a>).</li>
<li>Monitor Command is now supported (<a href="https://redirect.github.com/redis/go-redis/issues/2830">#2830</a>).</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>&quot;COMMAND&quot; command is now executed only when readonly (<a href="https://redirect.github.com/redis/go-redis/issues/2815">#2815</a>).</li>
</ul>
<h2>🧰 Maintenance and Dependency Updates</h2>
<ul>
<li>The usage of TSMadd ktvSlices has been clarified in the docstring (<a href="https://redirect.github.com/redis/go-redis/issues/2827">#2827</a>).</li>
<li>Dependencies have been updated:
<ul>
<li>rojopolis/spellcheck-github-actions from 0.34.0 to 0.35.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2807">#2807</a>)</li>
<li>actions/stale from 8 to 9 (<a href="https://redirect.github.com/redis/go-redis/issues/2828">#2828</a>)</li>
<li>actions/setup-go from 4 to 5 (<a href="https://redirect.github.com/redis/go-redis/issues/2829">#2829</a>)</li>
</ul>
</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/RyoMiyashita"><code>@​RyoMiyashita</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/lzakharov"><code>@​lzakharov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/rfyiamcool"><code>@​rfyiamcool</code></a>, <a href="https://github.com/rouzier"><code>@​rouzier</code></a>, <a href="https://github.com/splundid"><code>@​splundid</code></a> and <a href="https://github.com/x1nchen"><code>@​x1nchen</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/73c879df57f61fe7d231f9248f3c448ace37bd77"><code>73c879d</code></a> 9.3.1 (<a href="https://redirect.github.com/redis/go-redis/issues/2835">#2835</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/5665b0f7ea26298a8c7ff45323348e7b04b8d491"><code>5665b0f</code></a> Clarify TSMadd ktvSlices usage in docstring (<a href="https://redirect.github.com/redis/go-redis/issues/2827">#2827</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/9073e4056de44fe547b4a84840f868165f2f414a"><code>9073e40</code></a> Update docs and examples (<a href="https://redirect.github.com/redis/go-redis/issues/2806">#2806</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c82876433698ae24c7a2504a340f535e758af252"><code>c828764</code></a> Allow scanning redis values into pointer fields (<a href="https://redirect.github.com/redis/go-redis/issues/2787">#2787</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/716906addaa475a699df1d26869db99128fb1e80"><code>716906a</code></a> docs: correct the format error of the url in 'Connecting via a redis url' sec...</li>
<li><a href="https://github.com/redis/go-redis/commit/86c68be278a8541f6419513940ffaefa45341e78"><code>86c68be</code></a> Execute &quot;COMMAND&quot; command only when readonly  (<a href="https://redirect.github.com/redis/go-redis/issues/2815">#2815</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8c695488a247283d92e4d03c1774d9e2c4583244"><code>8c69548</code></a> fix: add Cmder annotation (<a href="https://redirect.github.com/redis/go-redis/issues/2816">#2816</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e535124055c53d281390d3ec5957e437a885430d"><code>e535124</code></a> Change Z member to string (<a href="https://redirect.github.com/redis/go-redis/issues/2818">#2818</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/1d784578dfea5d9a33a7b0fc0d989cc929c98f2b"><code>1d78457</code></a> Add BITFIELD_RO Command (<a href="https://redirect.github.com/redis/go-redis/issues/2820">#2820</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/277e8b7d9f9c5c5dc5690e0e11a029278bb25297"><code>277e8b7</code></a> Support Monitor Command (<a href="https://redirect.github.com/redis/go-redis/issues/2830">#2830</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.3.0...v9.3.1">compare view</a></li>
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

### Comment by @app-sre-bot on December 25, 2023 at 04:52 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on January 01, 2024 at 04:38 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/514*
