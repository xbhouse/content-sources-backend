---
type: pull_request
number: 678
title: "Build: Bump the go group across 1 directory with 6 updates"
state: merged
author: dependabot
created: 2024-05-28T16:26:27Z
updated: 2024-05-28T17:27:35Z
closed: 2024-05-28T17:27:31Z
merged: 2024-05-28T17:27:31Z
base_branch: main
head_branch: dependabot/go_modules/go-1d4fcceee7
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/678
---

# Pull Request #678: Build: Bump the go group across 1 directory with 6 updates

**Author**: @dependabot
**Created**: May 28, 2024 at 04:26 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-1d4fcceee7`

## Description

Bumps the go group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.32.0` | `1.33.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.15` | `1.17.16` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.35.4` | `1.35.5` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.5.1715867131` | `2024.5.1716497237` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.27.0` | `0.28.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.5` | `5.6.0` |


Updates `github.com/rs/zerolog` from 1.32.0 to 1.33.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/rs/zerolog/commit/c78e50e2da70f4ae63e1b65222c3acf12e9ba699"><code>c78e50e</code></a> Add fields order (<a href="https://redirect.github.com/rs/zerolog/issues/550">#550</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/8582bed24f021b174e2dfddbd2cfc72a8e39ccfd"><code>8582bed</code></a> fix: use <code>TimestampFunc</code> in busrt sampler (<a href="https://redirect.github.com/rs/zerolog/issues/671">#671</a>) (<a href="https://redirect.github.com/rs/zerolog/issues/672">#672</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/7d9db06a5369272008eee64a52ebf09443caff9e"><code>7d9db06</code></a> Allow setting floating point precision in JSON. (<a href="https://redirect.github.com/rs/zerolog/issues/663">#663</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/e5aa7e36278c0627297fad5d24a13e3f71bd15dd"><code>e5aa7e3</code></a> Revert <a href="https://redirect.github.com/rs/zerolog/issues/662">#662</a></li>
<li><a href="https://github.com/rs/zerolog/commit/0efa414907bc6775424a4987c9e13b63d17bd4c2"><code>0efa414</code></a> Fix panic caused by an extra malformed <code>level</code> field (<a href="https://redirect.github.com/rs/zerolog/issues/665">#665</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/eb081e1fa2325f7aab25a9386e463f67b37b421d"><code>eb081e1</code></a> chore: fix some typos in comments (<a href="https://redirect.github.com/rs/zerolog/issues/667">#667</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/2d899f0cf9919b091561c58233e624ffcc8b8824"><code>2d899f0</code></a> set debug log color (<a href="https://redirect.github.com/rs/zerolog/issues/662">#662</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/74cf37a3965b19e5e433dcd45ccdcfbb13cf239f"><code>74cf37a</code></a> Add EmptyFields method to remove all the fileds from logger (<a href="https://redirect.github.com/rs/zerolog/issues/575">#575</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/e5edd4b8ec0845c561d88e9bcdb3fe57b2acc249"><code>e5edd4b</code></a> Refactor: make code in comment valid and runable (<a href="https://redirect.github.com/rs/zerolog/issues/654">#654</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/582007f21da46f136517edb419ed44931ae72b73"><code>582007f</code></a> Add a time.Location to ConsoleWriter. (This allows UTC timestamps.) (<a href="https://redirect.github.com/rs/zerolog/issues/531">#531</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/rs/zerolog/compare/v1.32.0...v1.33.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.15 to 1.17.16
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8abec4c1a7b9326c05fee376dbd47cb653bbb4e3"><code>8abec4c</code></a> Release 2024-05-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70e7095bf016028d2c47f9f1f9ed7a1d88f8c920"><code>70e7095</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b2a3406d2e37fb520b27f9b0c2df1b1f6ee0ef5"><code>0b2a340</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1eb2d96cc305af9fe8ade05392db9b858373ace"><code>c1eb2d9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4c990d18e139d86883c00c5e82078950d638ef48"><code>4c990d1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c6c1626e4ea9aa69ddc5bbb64750c4a348ca3684"><code>c6c1626</code></a> s3: handle unrecognized values for Expires in responses (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2653">#2653</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8209abb7fa1aeb513228b4d8c1a459aeb6209d4d"><code>8209abb</code></a> Release 2024-05-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/81ad16867997bc40b56a9ced651840674f23d384"><code>81ad168</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5c92ae723530d8f711bbf1b1af03e7331fd42aff"><code>5c92ae7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6eeecd923623136c47f2ec3a429fc1ddb8acb252"><code>6eeecd9</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.15...credentials/v1.17.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.35.4 to 1.35.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f2706c8029255757ad0367000f64a65669a5d740"><code>f2706c8</code></a> Release 2023-12-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1dac0c95969dfffe03b376d8f937a437885bd316"><code>1dac0c9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cbd9216b092dbe1912addd94831de73df0072c4f"><code>cbd9216</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c7357bb31076e6e6d59faa8d1b93746ef051d066"><code>c7357bb</code></a> fix: reinstate presence of retryer when functional opts run but still respect...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c5c34b332ad29a72846f0d4ed228db09ee64ef7f"><code>c5c34b3</code></a> fix: translation of ini service sections into shared config (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2416">#2416</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3c7fbff96e191320f18e8513d20a73937d19675"><code>b3c7fbf</code></a> update express cache key (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2414">#2414</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9b90af4cc4235797576d8d7e3f1e27b782c60656"><code>9b90af4</code></a> fix: add non-vhostable buckets to path when using legacy endpoint resolver (#...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d643a8f0a8ad09075f41a60acba6a208cb36c58"><code>0d643a8</code></a> Release 2023-12-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e38b534ccea2707475381f985bc9236970bab554"><code>e38b534</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f394daf705a925e626f6a5c52aa433ff5504c7f1"><code>f394daf</code></a> Update SDK's smithy-go dependency to v1.19.0</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.35.4...service/ecs/v1.35.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.5.1715867131 to 2024.5.1716497237
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/0145964a299cfef6eee5dacd54c53f3e2e530c7c"><code>0145964</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276db5d53a5027b8e3dae95843...</li>
<li><a href="https://github.com/content-services/zest/commit/9bdbe820fa7708e5648394b38296c1ddaf00a483"><code>9bdbe82</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b735aa23b2ec37d2e4353b86e2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.5.1715867131...release/v2024.5.1716497237">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.27.0 to 0.28.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.28.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.28.0.</p>
<h3>Features</h3>
<ul>
<li>Add a <code>Fiber</code> performance tracing &amp; error reporting integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/795">#795</a>)</li>
<li>Add performance tracing to the <code>Echo</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/722">#722</a>)</li>
<li>Add performance tracing to the <code>FastHTTP</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/723">#732</a>)</li>
<li>Add performance tracing to the <code>Iris</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/809">#809</a>)</li>
<li>Add performance tracing to the <code>Negroni</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/808">#808</a>)</li>
<li>Add <code>FailureIssueThreshold</code> &amp; <code>RecoveryThreshold</code> to <code>MonitorConfig</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/775">#775</a>)</li>
<li>Use <code>errors.Unwrap()</code> to create exception groups (<a href="https://redirect.github.com/getsentry/sentry-go/pull/792">#792</a>)</li>
<li>Add support for matching on strings for <code>ClientOptions.IgnoreErrors</code> &amp; <code>ClientOptions.IgnoreTransactions</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/819">#819</a>)</li>
<li>Add <code>http.request.method</code> attribute for performance span data (<a href="https://redirect.github.com/getsentry/sentry-go/pull/786">#786</a>)</li>
<li>Accept <code>interface{}</code> for span data values (<a href="https://redirect.github.com/getsentry/sentry-go/pull/784">#784</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>Fix missing stack trace for parsing error in <code>logrusentry</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/689">#689</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.28.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.28.0.</p>
<h3>Features</h3>
<ul>
<li>Add a <code>Fiber</code> performance tracing &amp; error reporting integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/795">#795</a>)</li>
<li>Add performance tracing to the <code>Echo</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/722">#722</a>)</li>
<li>Add performance tracing to the <code>FastHTTP</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/723">#732</a>)</li>
<li>Add performance tracing to the <code>Iris</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/809">#809</a>)</li>
<li>Add performance tracing to the <code>Negroni</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/808">#808</a>)</li>
<li>Add <code>FailureIssueThreshold</code> &amp; <code>RecoveryThreshold</code> to <code>MonitorConfig</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/775">#775</a>)</li>
<li>Use <code>errors.Unwrap()</code> to create exception groups (<a href="https://redirect.github.com/getsentry/sentry-go/pull/792">#792</a>)</li>
<li>Add support for matching on strings for <code>ClientOptions.IgnoreErrors</code> &amp; <code>ClientOptions.IgnoreTransactions</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/819">#819</a>)</li>
<li>Add <code>http.request.method</code> attribute for performance span data (<a href="https://redirect.github.com/getsentry/sentry-go/pull/786">#786</a>)</li>
<li>Accept <code>interface{}</code> for span data values (<a href="https://redirect.github.com/getsentry/sentry-go/pull/784">#784</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>Fix missing stack trace for parsing error in <code>logrusentry</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/689">#689</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/670c82527c5f26c0644b23a5b09021c0881beec3"><code>670c825</code></a> release: 0.28.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/29be191c0ed2a51ba65354e19e70e1e270c039f7"><code>29be191</code></a> Prepare 0.28.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/828">#828</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/23c5137c84c74371307ce682f1283f3c44dd921b"><code>23c5137</code></a> Fix go lint issues (<a href="https://redirect.github.com/getsentry/sentry-go/issues/825">#825</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/3192e9fd92d7a90fe3263178492bc6529b959cbf"><code>3192e9f</code></a> add SendEventWithContext method (<a href="https://redirect.github.com/getsentry/sentry-go/issues/824">#824</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9a9f28523f39be5033612ba2dce989af82856ef4"><code>9a9f285</code></a> feat(metrics): add basic metrics type support (<a href="https://redirect.github.com/getsentry/sentry-go/issues/789">#789</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/84427d888594e3555ba66165eeaaeb714a660a9d"><code>84427d8</code></a> match strings for ignoreErrors as well (<a href="https://redirect.github.com/getsentry/sentry-go/issues/819">#819</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7954115fa671ec0f125a52e9df471a6f15476bb0"><code>7954115</code></a> Add<code>FailureIssueThreshold</code> &amp; <code>RecoveryThreshold</code> to MonitorConfig (<a href="https://redirect.github.com/getsentry/sentry-go/issues/775">#775</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bf970d1cb308145380b73eefaffc35c58c1b34b0"><code>bf970d1</code></a> fix linter (<a href="https://redirect.github.com/getsentry/sentry-go/issues/816">#816</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/1492d375054c96fbccf3d810dd3fe06d88ace067"><code>1492d37</code></a> fix failing tests (<a href="https://redirect.github.com/getsentry/sentry-go/issues/815">#815</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/df20ce63bbede6de539d2ad2f15f7daf2b703c65"><code>df20ce6</code></a> normalize http.response.status_code (<a href="https://redirect.github.com/getsentry/sentry-go/issues/811">#811</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.27.0...v0.28.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.5.5 to 5.6.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.6.0 (May 25, 2024)</h1>
<ul>
<li>Add StrictNamedArgs (Tomas Zahradnicek)</li>
<li>Add support for macaddr8 type (Carlos Pérez-Aradros Herce)</li>
<li>Add SeverityUnlocalized field to PgError / Notice</li>
<li>Performance optimization of RowToStructByPos/Name (Zach Olstein)</li>
<li>Allow customizing context canceled behavior for pgconn</li>
<li>Add ScanLocation to pgtype.Timestamp[tz]Codec</li>
<li>Add custom data to pgconn.PgConn</li>
<li>Fix ResultReader.Read() to handle nil values</li>
<li>Do not encode interval microseconds when they are 0 (Carlos Pérez-Aradros Herce)</li>
<li>pgconn.SafeToRetry checks for wrapped errors (tjasko)</li>
<li>Failed connection attempts include all errors</li>
<li>Optimize LargeObject.Read (Mitar)</li>
<li>Add tracing for connection acquire and release from pool (ngavinsir)</li>
<li>Fix encode driver.Valuer not called when nil</li>
<li>Add support for custom JSON marshal and unmarshal (Mitar)</li>
<li>Use Go default keepalive for TCP connections (Hans-Joachim Kliemeck)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/572d7fff326f1befdbf9f36a0c0a2b6661432079"><code>572d7ff</code></a> Release v5.6.0</li>
<li><a href="https://github.com/jackc/pgx/commit/b4911f1da7ee4746365dd38acafd2ae546bca0a7"><code>b4911f1</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2019">#2019</a> from jackc/fix-encode-driver-valuer-on-pointer</li>
<li><a href="https://github.com/jackc/pgx/commit/24c0a5e8ff740778bbed9810629cb83c9ea18bf9"><code>24c0a5e</code></a> remove keepalive and rely on GOLANG default (since go 1.13 default is 15s)</li>
<li><a href="https://github.com/jackc/pgx/commit/9ca9203afbd95ed53bf29eb0d3ac42f20ae5d17f"><code>9ca9203</code></a> Move typed nil handling to Map.Encode from anynil</li>
<li><a href="https://github.com/jackc/pgx/commit/79cab4640f4ba70c4be37ff2071ed4180ba8b63a"><code>79cab46</code></a> Only use anynil inside of pgtype</li>
<li><a href="https://github.com/jackc/pgx/commit/6ea2d248a38eda25a686ebca7ff9ea4c313f513c"><code>6ea2d24</code></a> Remove anynil.NormalizeSlice</li>
<li><a href="https://github.com/jackc/pgx/commit/c1075bfff0752a01d946f40a15e0893f6637c579"><code>c1075bf</code></a> Remove some special casing for QueryExecModeExec</li>
<li><a href="https://github.com/jackc/pgx/commit/cf6074fe5c7c43fc5c73d56761d20f73f2bc378c"><code>cf6074f</code></a> Remove unused anynil.Normalize</li>
<li><a href="https://github.com/jackc/pgx/commit/13beb380f51e3d744e27e7d55a3078ac8753ccf9"><code>13beb38</code></a> Fix encode driver.Valuer on nil-able non-pointers</li>
<li><a href="https://github.com/jackc/pgx/commit/fec45c802ba9f19107aa24f9aea46f2f21577621"><code>fec45c8</code></a> Refactor appendParamsForQueryExecModeExec</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.5...v5.6.0">compare view</a></li>
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

### Comment by @app-sre-bot on May 28, 2024 at 04:27 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on May 28, 2024 at 04:27 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on May 28, 2024 at 05:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/678*
