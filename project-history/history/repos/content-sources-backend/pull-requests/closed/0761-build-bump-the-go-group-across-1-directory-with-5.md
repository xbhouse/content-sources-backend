---
type: pull_request
number: 761
title: "Build: Bump the go group across 1 directory with 5 updates"
state: closed
author: dependabot
created: 2024-08-05T04:38:56Z
updated: 2024-08-07T19:51:08Z
closed: 2024-08-07T19:51:06Z
base_branch: main
head_branch: dependabot/go_modules/go-a52da44633
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/761
---

# Pull Request #761: Build: Bump the go group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: August 05, 2024 at 04:38 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-a52da44633`

## Description

Bumps the go group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.126.0` | `0.127.0` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.10` | `1.25.11` |
| [github.com/archdx/zerolog-sentry](https://github.com/archdx/zerolog-sentry) | `1.8.3` | `1.8.4` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.7.1721415179` | `2024.8.1722539431` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.6.0` | `9.6.1` |


Updates `github.com/getkin/kin-openapi` from 0.126.0 to 0.127.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.127.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3: include local reference parts in refPath saved by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/978">getkin/kin-openapi#978</a></li>
<li>fix: update type: file to type: string and format: binary by <a href="https://github.com/reversearrow"><code>@​reversearrow</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/980">getkin/kin-openapi#980</a></li>
<li>test: add a test for <a href="https://redirect.github.com/getkin/kin-openapi/issues/499">#499</a> by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/982">getkin/kin-openapi#982</a></li>
<li>test: add a test for <a href="https://redirect.github.com/getkin/kin-openapi/issues/961">#961</a> by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/981">getkin/kin-openapi#981</a></li>
<li>openapi3gen: generate &quot;nullable: true&quot; for pointers by <a href="https://github.com/endertunc"><code>@​endertunc</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/987">getkin/kin-openapi#987</a></li>
<li>openapi3filter: Remove inconsistency for arrays in queries by <a href="https://github.com/TheSadlig"><code>@​TheSadlig</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/990">getkin/kin-openapi#990</a></li>
<li>openapi3filter: remove double call by <a href="https://github.com/AriehSchneier"><code>@​AriehSchneier</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/993">getkin/kin-openapi#993</a></li>
<li>openapi3filter: Fix default arrays application for query parameters by <a href="https://github.com/TheSadlig"><code>@​TheSadlig</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/992">getkin/kin-openapi#992</a></li>
<li>routers: downgrade gorilla/mux to v1.8.0 for <a href="https://redirect.github.com/getkin/kin-openapi/issues/988">#988</a> by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/996">getkin/kin-openapi#996</a></li>
<li>openapi3: export <code>ComponentRef</code> for usage in <code>RefNameResolver</code> by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/998">getkin/kin-openapi#998</a></li>
<li>Add note on gorillamux by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/999">getkin/kin-openapi#999</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/reversearrow"><code>@​reversearrow</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/980">getkin/kin-openapi#980</a></li>
<li><a href="https://github.com/endertunc"><code>@​endertunc</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/987">getkin/kin-openapi#987</a></li>
<li><a href="https://github.com/TheSadlig"><code>@​TheSadlig</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/990">getkin/kin-openapi#990</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.126.0...v0.127.0">https://github.com/getkin/kin-openapi/compare/v0.126.0...v0.127.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/d7226466cdeb3a02e90c8cd242c631896e5a2d32"><code>d722646</code></a> Add note on gorillamux (<a href="https://redirect.github.com/getkin/kin-openapi/issues/999">#999</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/b5bcd71617fd31686e538427298c0a98751968f3"><code>b5bcd71</code></a> openapi3: export <code>ComponentRef</code> for usage in <code>RefNameResolver</code> (<a href="https://redirect.github.com/getkin/kin-openapi/issues/998">#998</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/f7ebd9c86c592b96e0cb91690772e580469206aa"><code>f7ebd9c</code></a> routers: downgrade gorilla/mux to v1.8.0 for <a href="https://redirect.github.com/getkin/kin-openapi/issues/988">#988</a> (<a href="https://redirect.github.com/getkin/kin-openapi/issues/996">#996</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/cf9684e8c121a9475531e64a786732139a81d6d9"><code>cf9684e</code></a> openapi3filter: Fix default arrays application for query parameters (<a href="https://redirect.github.com/getkin/kin-openapi/issues/992">#992</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/cd0a337347566473f362e3944de8f09a6cb547b0"><code>cd0a337</code></a> openapi3filter: remove double call (<a href="https://redirect.github.com/getkin/kin-openapi/issues/993">#993</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/1a2712f0758d88f28e5ed3a2f4e78975494ff898"><code>1a2712f</code></a> openapi3filter: Remove inconsistency for arrays in queries (<a href="https://redirect.github.com/getkin/kin-openapi/issues/990">#990</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/25ce76510a5f4602e0ac76b2916d587e7315fe1d"><code>25ce765</code></a> fixes <a href="https://redirect.github.com/getkin/kin-openapi/issues/968">#968</a> - openapi3gen: generate &quot;nullable: true&quot; for pointers (<a href="https://redirect.github.com/getkin/kin-openapi/issues/987">#987</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/1a819a1374ef609f37591483a559fd5f2cb81905"><code>1a819a1</code></a> test: add a test for <a href="https://redirect.github.com/getkin/kin-openapi/issues/961">#961</a> (<a href="https://redirect.github.com/getkin/kin-openapi/issues/981">#981</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/8c15898390de6f2074a20bc04aa7d08f63ff9bda"><code>8c15898</code></a> test: add a test for <a href="https://redirect.github.com/getkin/kin-openapi/issues/499">#499</a> (<a href="https://redirect.github.com/getkin/kin-openapi/issues/982">#982</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/a7bc0ff0db3adb4569b327771c5a3b5bffbb9a17"><code>a7bc0ff</code></a> fix: update type: file to type: string and format: binary (<a href="https://redirect.github.com/getkin/kin-openapi/issues/980">#980</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.126.0...v0.127.0">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.25.10 to 1.25.11
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/4a50b36f638c6899089e6e3457425528ce693933"><code>4a50b36</code></a> ci: Add PostgreSQL 14 and 15 to GitHub Actions matrix (<a href="https://redirect.github.com/go-gorm/gorm/issues/7081">#7081</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/11c4331058e5fed7d05707c5a6a6997947509f41"><code>11c4331</code></a> feat: add MapColumns method (<a href="https://redirect.github.com/go-gorm/gorm/issues/6901">#6901</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/8a0af58cc5315f2b8562cac2c06fbdea05cd0e03"><code>8a0af58</code></a> fix map fields with clickhouse driver</li>
<li><a href="https://github.com/go-gorm/gorm/commit/4f6291154b3c0a2e8d8158232709e381a484c5cf"><code>4f62911</code></a> Allow to support other field types</li>
<li><a href="https://github.com/go-gorm/gorm/commit/109f239fae95b310f64236bcc0fbe87bbfb1edfc"><code>109f239</code></a> add DB level propagation for the Unscoped flag (<a href="https://redirect.github.com/go-gorm/gorm/issues/7007">#7007</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/79bf7f92edec8f9be220524ed61ce9ae66ed197b"><code>79bf7f9</code></a> fix CI for sqlserver</li>
<li><a href="https://github.com/go-gorm/gorm/commit/3d09f7947f97598ef8b951cc05d9b5a886e800cc"><code>3d09f79</code></a> only listen local port</li>
<li><a href="https://github.com/go-gorm/gorm/commit/73a988ceb22651e01c968a9ec20ae1709e73c8e6"><code>73a988c</code></a> fix(scan): update Scan function to reset structs to zero values for each scan...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/05167fd5918f8e653f78a635ba590581d02309a3"><code>05167fd</code></a> fix: use reflect.Append when preloading nested associations (<a href="https://redirect.github.com/go-gorm/gorm/issues/7014">#7014</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/78c6dfd712aca47ab1945921cdb37d1185aea7f3"><code>78c6dfd</code></a> Fix association replace non-addressable panic (<a href="https://redirect.github.com/go-gorm/gorm/issues/7012">#7012</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/go-gorm/gorm/compare/v1.25.10...v1.25.11">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/archdx/zerolog-sentry` from 1.8.3 to 1.8.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/5ae2fbd8e0ace288cf39b966b5b81ab948e00298"><code>5ae2fbd</code></a> drop unsafe bytes to str conv</li>
<li>See full diff in <a href="https://github.com/archdx/zerolog-sentry/compare/v1.8.3...v1.8.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.7.1721415179 to 2024.8.1722539431
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/bd43ce2dd878fb5906b102e23903fd32ce7897cf"><code>bd43ce2</code></a> Update pulp bindings to e69db48356e528a464be3da896237b8a2298890a7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/b232aa63c3b6e7c8905879b0bc9fcd710c437b5c"><code>b232aa6</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d46466546057952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/1588a88104d510e1b0061affd6b3b790e8b4ecf2"><code>1588a88</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84745eaa369dc37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/06adeeae8fe6206d653c1fc107c4cd122a6752ed"><code>06adeea</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a98268a9dcb7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/6f39367ce225df214ce894c31ca303803c6c4fb2"><code>6f39367</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/15">#15</a> from jlsherrill/fix</li>
<li><a href="https://github.com/content-services/zest/commit/30af2293c1d64d43283b09b4615dc76cb5832d33"><code>30af229</code></a> Fix template for params in get request</li>
<li><a href="https://github.com/content-services/zest/commit/7669ad3f577ee06df9f945028eb2f6c11afd8016"><code>7669ad3</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e4755ebb4eb2cb7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/bb144f539e7418ccc668c5c2cfd0f5506b6b3fa7"><code>bb144f5</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b735ea54669037d2e4353b86e2...</li>
<li><a href="https://github.com/content-services/zest/commit/073e6e27dc4ec2ed68dc2b9aa442873f33da6c12"><code>073e6e2</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276d9e84683127b8e3dae95843...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.7.1721415179...release/v2024.8.1722539431">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.6.0 to 9.6.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.6.1</h2>
<h1>Changes</h1>
<h2>9.6</h2>
<p>This release contains all new features from version 9.6.</p>
<h3>🚀 New Features</h3>
<ul>
<li>Support Hash-field expiration commands (<a href="https://redirect.github.com/redis/go-redis/issues/2991">#2991</a>)</li>
<li>Support Hash-field expiration commands in Pipeline &amp; Fix HExpire HExpireWithArgs expiration (<a href="https://redirect.github.com/redis/go-redis/issues/3038">#3038</a>)</li>
<li>Support NOVALUES parameter for HSCAN (<a href="https://redirect.github.com/redis/go-redis/issues/2925">#2925</a>)</li>
<li>Added test case for CLIENT KILL with MAXAGE option (<a href="https://redirect.github.com/redis/go-redis/issues/2971">#2971</a>)</li>
<li>Add support for XREAD last entry (<a href="https://redirect.github.com/redis/go-redis/issues/3005">#3005</a>)</li>
<li>Reduce the type assertion of CheckConn (<a href="https://redirect.github.com/redis/go-redis/issues/3066">#3066</a>)</li>
</ul>
<h2>9.6.1</h2>
<p>In addition minor changes were performed to retract version 9.5.3 and 9.5.4 that were released accidentally.</p>
<h3>🧰 Maintenance</h3>
<ul>
<li>Change CI to 7.4.0-RC2 (<a href="https://redirect.github.com/redis/go-redis/issues/3070">#3070</a>)</li>
</ul>
<h3>🎁 Package Distribution</h3>
<ul>
<li>Retract versions 9.5.3 and 9.5.4 (<a href="https://redirect.github.com/redis/go-redis/issues/3069">#3069</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/b1ron"><code>@​b1ron</code></a>, <a href="https://github.com/gerzse"><code>@​gerzse</code></a>, <a href="https://github.com/haines"><code>@​haines</code></a>, <a href="https://github.com/immersedin"><code>@​immersedin</code></a>, <a href="https://github.com/naiqianz"><code>@​naiqianz</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/srikar-jilugu"><code>@​srikar-jilugu</code></a>, <a href="https://github.com/tzongw"><code>@​tzongw</code></a>, <a href="https://github.com/vladvildanov"><code>@​vladvildanov</code></a>, <a href="https://github.com/vmihailenco"><code>@​vmihailenco</code></a> and <a href="https://github.com/monkey92t"><code>@​monkey92t</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/c3610cde42ad05854fa05e312d4b34cf1a809cc2"><code>c3610cd</code></a> cp retracts and bump version to 9.6.1 (<a href="https://redirect.github.com/redis/go-redis/issues/3069">#3069</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b0b94e79e07ac153ac2813266f0ac5d3e5bfb76c"><code>b0b94e7</code></a> Change CI to 7.4.0-RC2 (<a href="https://redirect.github.com/redis/go-redis/issues/3070">#3070</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.6.0...v9.6.1">compare view</a></li>
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

### Comment by @app-sre-bot on August 05, 2024 at 04:39 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on August 07, 2024 at 07:51 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/761*
