---
type: pull_request
number: 855
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2024-10-21T04:10:01Z
updated: 2024-10-21T13:36:24Z
closed: 2024-10-21T13:36:16Z
merged: 2024-10-21T13:36:16Z
base_branch: main
head_branch: dependabot/go_modules/go-3b11dfd5f0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/855
---

# Pull Request #855: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: October 21, 2024 at 04:10 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-3b11dfd5f0`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/swaggo/swag](https://github.com/swaggo/swag) | `1.16.3` | `1.16.4` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.20.4` | `1.20.5` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.10.1728596503` | `2024.10.1729249374` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.29.0` | `0.29.1` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.6.1` | `9.7.0` |

Updates `github.com/swaggo/swag` from 1.16.3 to 1.16.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/swaggo/swag/releases">github.com/swaggo/swag's releases</a>.</em></p>
<blockquote>
<h2>v1.16.4</h2>
<h2>Changelog</h2>
<ul>
<li>85254b4 Added multiline support for <a href="https://github.com/description"><code>@​description</code></a> attribute for securityDefinitions (<a href="https://redirect.github.com/swaggo/swag/issues/1786">#1786</a>)</li>
<li>7204462 Added suport for parsing comments inside of function bodies (<a href="https://redirect.github.com/swaggo/swag/issues/1824">#1824</a>)</li>
<li>d323b48 Feat: Add support for parenthesis in router patterns (<a href="https://redirect.github.com/swaggo/swag/issues/1859">#1859</a>)</li>
<li>15dae35 Feat: multi-arch docker image (<a href="https://redirect.github.com/swaggo/swag/issues/1756">#1756</a>)</li>
<li>83fe3ca Fix generics used with function scoped types (<a href="https://redirect.github.com/swaggo/swag/issues/1883">#1883</a>)</li>
<li>ff50cd6 Fix global overrides for any/interface ref types (<a href="https://redirect.github.com/swaggo/swag/issues/1835">#1835</a>)</li>
<li>1d730c5 Fix param comment escaping issue (<a href="https://redirect.github.com/swaggo/swag/issues/1890">#1890</a>)</li>
<li>697572a Fixes Issue 1829 (<a href="https://redirect.github.com/swaggo/swag/issues/1830">#1830</a>)</li>
<li>28de14c Flags to parse internal and dependency package (<a href="https://redirect.github.com/swaggo/swag/issues/1894">#1894</a>)</li>
<li>4c2f8dd Handle case of empty GOROOT (<a href="https://redirect.github.com/swaggo/swag/issues/1798">#1798</a>)</li>
<li>87e7d9c Update docker go build version to 1.21 (<a href="https://redirect.github.com/swaggo/swag/issues/1758">#1758</a>)</li>
<li>4fd8a36 Update docs for request and response headers (<a href="https://redirect.github.com/swaggo/swag/issues/1825">#1825</a>)</li>
<li>56fde5c Update operation.go (<a href="https://redirect.github.com/swaggo/swag/issues/1753">#1753</a>)</li>
<li>807dd1f [Issue 1812] fix misalignment in expected.json and api.go messing with parser_test (<a href="https://redirect.github.com/swaggo/swag/issues/1836">#1836</a>)</li>
<li>91624ad add support for &quot;title&quot; tag (<a href="https://redirect.github.com/swaggo/swag/issues/1762">#1762</a>)</li>
<li>f32d4d3 adds support for complex types with function scope (<a href="https://redirect.github.com/swaggo/swag/issues/1813">#1813</a>)</li>
<li>c7f1cd8 adds support for pointer function scoped fields (<a href="https://redirect.github.com/swaggo/swag/issues/1841">#1841</a>)</li>
<li>8a47dcb bump go version (<a href="https://redirect.github.com/swaggo/swag/issues/1797">#1797</a>)</li>
<li>0834357 chore(deps): bump golang.org/x/net from 0.17.0 to 0.23.0 (<a href="https://redirect.github.com/swaggo/swag/issues/1793">#1793</a>)</li>
<li>1bb1445 chore(deps): bump golang.org/x/net in /example/celler (<a href="https://redirect.github.com/swaggo/swag/issues/1794">#1794</a>)</li>
<li>6aa6613 chore(deps): bump golang.org/x/net in /example/go-module-support (<a href="https://redirect.github.com/swaggo/swag/issues/1795">#1795</a>)</li>
<li>d5af957 chore(deps): bump golang.org/x/net in /example/markdown (<a href="https://redirect.github.com/swaggo/swag/issues/1792">#1792</a>)</li>
<li>0368d7d chore(deps): bump golang.org/x/net in /example/object-map-example (<a href="https://redirect.github.com/swaggo/swag/issues/1796">#1796</a>)</li>
<li>b8662de chore(deps): bump google.golang.org/protobuf (<a href="https://redirect.github.com/swaggo/swag/issues/1773">#1773</a>)</li>
<li>4a11e23 chore(deps): bump google.golang.org/protobuf (<a href="https://redirect.github.com/swaggo/swag/issues/1774">#1774</a>)</li>
<li>937c239 chore(deps): bump google.golang.org/protobuf in /example/celler (<a href="https://redirect.github.com/swaggo/swag/issues/1775">#1775</a>)</li>
<li>103ac42 chore: Update ci.yml (<a href="https://redirect.github.com/swaggo/swag/issues/1902">#1902</a>)</li>
<li>90aa46f chore: fix some typos in comments (<a href="https://redirect.github.com/swaggo/swag/issues/1788">#1788</a>)</li>
<li>e55c557 feat: read from stdin, write to stdout (<a href="https://redirect.github.com/swaggo/swag/issues/1831">#1831</a>) (<a href="https://redirect.github.com/swaggo/swag/issues/1832">#1832</a>)</li>
<li>fd2fa83 fix issue: <a href="https://redirect.github.com/swaggo/swag/issues/1780">#1780</a>: filter $GOROOT path (<a href="https://redirect.github.com/swaggo/swag/issues/1827">#1827</a>)</li>
<li>10030b0 fix parse nested structs and aliases (<a href="https://redirect.github.com/swaggo/swag/issues/1866">#1866</a>)</li>
<li>7159b0f fix: failing assert in enums test on 32bit (<a href="https://redirect.github.com/swaggo/swag/issues/1634">#1634</a>)</li>
<li>928264c fix: remove dropped tags from general infos (<a href="https://redirect.github.com/swaggo/swag/issues/1764">#1764</a>)</li>
<li>a74d34c fix：parse all field names declared in a row (<a href="https://redirect.github.com/swaggo/swag/issues/1872">#1872</a>)</li>
<li>0b9e347 new release (<a href="https://redirect.github.com/swaggo/swag/issues/1901">#1901</a>)</li>
<li>a3c6d12 support markdown description for declaration (<a href="https://redirect.github.com/swaggo/swag/issues/1893">#1893</a>)</li>
<li>9069105 update README (<a href="https://redirect.github.com/swaggo/swag/issues/1856">#1856</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/swaggo/swag/commit/0b9e347c196710ea155a147782bf51707a600c2c"><code>0b9e347</code></a> new release (<a href="https://redirect.github.com/swaggo/swag/issues/1901">#1901</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/103ac425bb10efd43610f7a16f17298bfd61e71c"><code>103ac42</code></a> chore: Update ci.yml (<a href="https://redirect.github.com/swaggo/swag/issues/1902">#1902</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/d323b4897592ac355ae3ee11fe140a115472d25d"><code>d323b48</code></a> Feat: Add support for parenthesis in router patterns (<a href="https://redirect.github.com/swaggo/swag/issues/1859">#1859</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/7159b0f7b3f11a516cf10b1b59e3513f955f912a"><code>7159b0f</code></a> fix: failing assert in enums test on 32bit (<a href="https://redirect.github.com/swaggo/swag/issues/1634">#1634</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/28de14c88f4745e163a184d96adbb009e29a52f5"><code>28de14c</code></a> Flags to parse internal and dependency package (<a href="https://redirect.github.com/swaggo/swag/issues/1894">#1894</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/a74d34c4566f59d7d8567f635e3909c7f86d823f"><code>a74d34c</code></a> fix：parse all field names declared in a row (<a href="https://redirect.github.com/swaggo/swag/issues/1872">#1872</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/4fd8a361b8eb343a3e463a8db5101568f0442c1e"><code>4fd8a36</code></a> Update docs for request and response headers (<a href="https://redirect.github.com/swaggo/swag/issues/1825">#1825</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/9069105bebbc90b10915f4ba2b014c1e875f2359"><code>9069105</code></a> update README (<a href="https://redirect.github.com/swaggo/swag/issues/1856">#1856</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/a3c6d12319acd519d0785c6211c4c8885e032e95"><code>a3c6d12</code></a> support markdown description for declaration (<a href="https://redirect.github.com/swaggo/swag/issues/1893">#1893</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/1d730c53a45b03640e1d15c6b0aacea6abbba2ce"><code>1d730c5</code></a> Fix param comment escaping issue (<a href="https://redirect.github.com/swaggo/swag/issues/1890">#1890</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/swaggo/swag/compare/v1.16.3...v1.16.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.20.4 to 1.20.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.20.5 / 2024-10-15</h2>
<p>We decided to revert <a href="https://redirect.github.com/prometheus/client_golang/pull/1424">the <code>testutil</code> change</a> that made our util functions less error-prone, but created a lot of work for our downstream users. Apologies for the pain! This revert should not cause any major breaking change, even if you already did the work--unless you depend on the <a href="https://redirect.github.com/grafana/mimir/pull/9624#issuecomment-2413401565">exact error message</a>.</p>
<p>Going forward, we plan to reinforce our release testing strategy <a href="https://redirect.github.com/prometheus/client_golang/issues/1646">[1]</a>,<a href="https://redirect.github.com/prometheus/client_golang/issues/1648">[2]</a> and deliver an enhanced <a href="https://redirect.github.com/prometheus/client_golang/issues/1639"><code>testutil</code> package/module</a> with more flexible and safer APIs.</p>
<p>Thanks to <a href="https://github.com/dashpole"><code>@​dashpole</code></a> <a href="https://github.com/dgrisonnet"><code>@​dgrisonnet</code></a> <a href="https://github.com/kakkoyun"><code>@​kakkoyun</code></a> <a href="https://github.com/ArthurSens"><code>@​ArthurSens</code></a> <a href="https://github.com/vesari"><code>@​vesari</code></a> <a href="https://github.com/logicalhan"><code>@​logicalhan</code></a> <a href="https://github.com/krajorama"><code>@​krajorama</code></a> <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> who helped in this patch release! 🤗</p>
<h3>Changelog</h3>
<p>[BUGFIX] testutil: Reverted <a href="https://redirect.github.com/prometheus/client_golang/issues/1424">#1424</a>; functions using compareMetricFamilies are (again) only failing if filtered metricNames are in the expected input. <a href="https://redirect.github.com/prometheus/client_golang/issues/1645">#1645</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.20.5 / 2024-10-15</h2>
<ul>
<li>[BUGFIX] testutil: Reverted <a href="https://redirect.github.com/prometheus/client_golang/issues/1424">#1424</a>; functions using compareMetricFamilies are (again) only failing if filtered metricNames are in the expected input.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/48e12a185519fd76b4e514b597483781d9ba4093"><code>48e12a1</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1645">#1645</a> from prometheus/cut-1204-pr1424</li>
<li><a href="https://github.com/prometheus/client_golang/commit/504ad9bf5c6419449d2cacf8cf8855bfdcfcfc18"><code>504ad9b</code></a> Cut 1.20.5; update comments.</li>
<li><a href="https://github.com/prometheus/client_golang/commit/584a7ce3d935e4fdca7b893f5f741d59f3289140"><code>584a7ce</code></a> Revert &quot;testutil compareMetricFamilies: make less error-prone (<a href="https://redirect.github.com/prometheus/client_golang/issues/1424">#1424</a>)&quot;</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.20.4...v1.20.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.10.1728596503 to 2024.10.1729249374
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/3ee09a31080ef9655936b97e7f70eef0d2ed1402"><code>3ee09a3</code></a> Update pulp bindings to 386d86254354e29d3b864523aed47865646d38f67968b5e2da9ab...</li>
<li><a href="https://github.com/content-services/zest/commit/ee12d5cab768abcffe3c7e4f7c69ba2cd154d02f"><code>ee12d5c</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7eaeb63663187e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/3703a96de36004fa6b5266b81b20847c51c8cd11"><code>3703a96</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276459ee8eb027b8e3dae95843...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.10.1728596503...release/v2024.10.1729249374">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.29.0 to 0.29.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.29.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.29.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Correlate errors to the current trace (<a href="https://redirect.github.com/getsentry/sentry-go/pull/886">#886</a>)</li>
<li>Set the trace context when the transaction finishes (<a href="https://redirect.github.com/getsentry/sentry-go/pull/888">#888</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Update the <code>sentrynegroni</code> integration to use the latest (v3.1.1) version of Negroni (<a href="https://redirect.github.com/getsentry/sentry-go/pull/885">#885</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.29.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.29.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Correlate errors to the current trace (<a href="https://redirect.github.com/getsentry/sentry-go/pull/886">#886</a>)</li>
<li>Set the trace context when the transaction finishes (<a href="https://redirect.github.com/getsentry/sentry-go/pull/888">#888</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Update the <code>sentrynegroni</code> integration to use the latest (v3.1.1) version of Negroni (<a href="https://redirect.github.com/getsentry/sentry-go/pull/885">#885</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/4a56cffc01693c179caa16bc435784744bd7f157"><code>4a56cff</code></a> release: 0.29.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/963131ab03d36a5050f202cede027d8162038f4a"><code>963131a</code></a> Prepare 0.29.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/891">#891</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/20684c77f12c8c592f4f63b104f0da96c66699d6"><code>20684c7</code></a> Push a new scope when a span is started (<a href="https://redirect.github.com/getsentry/sentry-go/issues/886">#886</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2935366b03fa3d88cd603a04cc4a0613fb33dba9"><code>2935366</code></a> Fix applying the <code>trace</code> context on events (<a href="https://redirect.github.com/getsentry/sentry-go/issues/888">#888</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/951e36393487a3bd7421e56efddad657ebdfa92d"><code>951e363</code></a> build(deps): bump golangci/golangci-lint-action from 6.1.0 to 6.1.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/890">#890</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e009c2ee19025057d4eca5cf6bb9d1f6077c160e"><code>e009c2e</code></a> build(deps): bump codecov/codecov-action from 4.5.0 to 4.6.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/889">#889</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/0e63d3959730ac153666afbfa440c6822f246e2b"><code>0e63d39</code></a> upgrade negroni to latest version (<a href="https://redirect.github.com/getsentry/sentry-go/issues/885">#885</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a6acd05f8d419d0445475b9018bd59fc0503b322"><code>a6acd05</code></a> Merge branch 'release/0.29.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.29.0...v0.29.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.6.1 to 9.7.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.7.0</h2>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Support Redis search and query capabilities (<a href="https://redirect.github.com/redis/go-redis/issues/2801">#2801</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3098">#3098</a>)</li>
<li>Support indexing and querying empty values (<a href="https://redirect.github.com/redis/go-redis/issues/3053">#3053</a>)</li>
<li>Support for Redis JSON with RESP2 protocol (<a href="https://redirect.github.com/redis/go-redis/issues/3146">#3146</a>)</li>
</ul>
<h2>🛠️ Improvements</h2>
<p>We're glad to announce that we added a search and query support in the current release.</p>
<h2>🧰 Maintenance</h2>
<ul>
<li>Documentation examples (<a href="https://redirect.github.com/redis/go-redis/issues/3102">#3102</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3106">#3106</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3110">#3110</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3111">#3111</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3113">#3113</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3114">#3114</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3115">#3115</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3123">#3123</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3124">#3124</a>)</li>
<li>retract v9.5.3 of redisotel and other extra packages (<a href="https://redirect.github.com/redis/go-redis/issues/3108">#3108</a>)</li>
<li>Add test coverage reporting and Codecov badge (<a href="https://redirect.github.com/redis/go-redis/issues/3055">#3055</a>)</li>
<li>Updated module version that points to retracted package version (<a href="https://redirect.github.com/redis/go-redis/issues/3074">#3074</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/andy-stark-redis"><code>@​andy-stark-redis</code></a>, <a href="https://github.com/ipechorin"><code>@​ipechorin</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a> and <a href="https://github.com/vladvildanov"><code>@​vladvildanov</code></a></p>
<h2>9.7.0-beta.1</h2>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Support Redis search and query capabilities (<a href="https://redirect.github.com/redis/go-redis/issues/2801">#2801</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3098">#3098</a>)</li>
<li>Support indexing and querying empty values (<a href="https://redirect.github.com/redis/go-redis/issues/3053">#3053</a>)</li>
</ul>
<h2>🛠️ Improvements</h2>
<p>We're glad to announce that we added a search and query support in the current release.</p>
<h2>🧰 Maintenance</h2>
<ul>
<li>Documentation examples (<a href="https://redirect.github.com/redis/go-redis/issues/3102">#3102</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3106">#3106</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3110">#3110</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3111">#3111</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3113">#3113</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3114">#3114</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3115">#3115</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3123">#3123</a>, <a href="https://redirect.github.com/redis/go-redis/issues/3124">#3124</a>)</li>
<li>retract v9.5.3 of redisotel and other extra packages (<a href="https://redirect.github.com/redis/go-redis/issues/3108">#3108</a>)</li>
<li>Add test coverage reporting and Codecov badge (<a href="https://redirect.github.com/redis/go-redis/issues/3055">#3055</a>)</li>
<li>Updated module version that points to retracted package version (<a href="https://redirect.github.com/redis/go-redis/issues/3074">#3074</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/andy-stark-redis"><code>@​andy-stark-redis</code></a>, <a href="https://github.com/ipechorin"><code>@​ipechorin</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a> and <a href="https://github.com/vladvildanov"><code>@​vladvildanov</code></a></p>
<h2>9.6.2</h2>
<h1>Changes</h1>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/ed37c33a9037483ad2a6b1042e5eb6df89009a1c"><code>ed37c33</code></a> Updated package version [9.7] (<a href="https://redirect.github.com/redis/go-redis/issues/3159">#3159</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/135f8e3b12d7dce2ed41fd9e428ce90b085f52d0"><code>135f8e3</code></a> Fix field name spellings (<a href="https://redirect.github.com/redis/go-redis/issues/3132">#3132</a>) (<a href="https://redirect.github.com/redis/go-redis/issues/3156">#3156</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/ac2e91d9d9ccabefd9a158d5ab294345d2368707"><code>ac2e91d</code></a> Support Json with Resp 2 (<a href="https://redirect.github.com/redis/go-redis/issues/3146">#3146</a>) (<a href="https://redirect.github.com/redis/go-redis/issues/3155">#3155</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/ec680aec142159539c2594584a5dd31a416a4683"><code>ec680ae</code></a> Remove direct read from TLS underlying conn (<a href="https://redirect.github.com/redis/go-redis/issues/3138">#3138</a>) (<a href="https://redirect.github.com/redis/go-redis/issues/3154">#3154</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/ad131f49b0210b3252648b880921b9f6d2fe452e"><code>ad131f4</code></a> Updated package version (<a href="https://redirect.github.com/redis/go-redis/issues/3134">#3134</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/d9eeed131a32b4aeaf1e1915b6717f8cd2af9124"><code>d9eeed1</code></a> Fix Flaky Test: should handle FTAggregate with Unstable RESP3 Search Module a...</li>
<li><a href="https://github.com/redis/go-redis/commit/e99abe45469cefe078a05e3c656ca452f3cce646"><code>e99abe4</code></a> DOC-4237 added Bloom filter examples (<a href="https://redirect.github.com/redis/go-redis/issues/3115">#3115</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/9e79c9bd394cb62efb0e00954a9014c09f0aeaa6"><code>9e79c9b</code></a> DOC-4228 JSON code examples (<a href="https://redirect.github.com/redis/go-redis/issues/3114">#3114</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/9e3709c4040acfa504bb7803dba709e70e780dc9"><code>9e3709c</code></a> DOC-4234 added bitmap examples (<a href="https://redirect.github.com/redis/go-redis/issues/3124">#3124</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/69b4c0165fc7981144e5cfe3b3d028827f4e82a0"><code>69b4c01</code></a> DOC-4241 added t-digest examples (<a href="https://redirect.github.com/redis/go-redis/issues/3123">#3123</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.6.1...v9.7.0">compare view</a></li>
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

### Comment by @app-sre-bot on October 21, 2024 at 04:10 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on October 21, 2024 at 08:05 AM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on October 21, 2024 at 01:36 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/855*
