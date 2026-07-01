---
type: pull_request
number: 1471
title: "Build: Bump the go group with 9 updates"
state: merged
author: dependabot
created: 2026-04-20T05:00:13Z
updated: 2026-04-20T08:31:17Z
closed: 2026-04-20T08:31:15Z
merged: 2026-04-20T08:31:15Z
base_branch: main
head_branch: dependabot/go_modules/go-398403d4b6
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1471
---

# Pull Request #1471: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: April 20, 2026 at 05:00 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-398403d4b6`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/labstack/gommon](https://github.com/labstack/gommon) | `0.4.2` | `0.5.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.5` | `1.41.6` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.14` | `1.32.16` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.14` | `1.19.15` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.68.0` | `1.69.1` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.99.0` | `1.99.1` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.4.1775140412` | `2026.4.1775836578` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.45.0` | `0.45.1` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.9.1` | `5.9.2` |

Updates `github.com/labstack/gommon` from 0.4.2 to 0.5.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/gommon/releases">github.com/labstack/gommon's releases</a>.</em></p>
<blockquote>
<h2>v0.5.0</h2>
<h2>Highlights</h2>
<ul>
<li><strong><code>email</code>: SMTPS / implicit TLS on port 465.</strong> <code>smtp.SendMail</code> only speaks plain + STARTTLS, so Resend/SendGrid/etc. on <code>:465</code> hang on the handshake. Detect port 465 and dial TLS directly. Added <code>Email.TLSConfig</code> (custom root pool / ServerName; always cloned per send) and <code>Email.DialTimeout</code> (scoped to the TCP/TLS connect phase).</li>
<li><strong><code>email</code>: no silent cleartext downgrade.</strong> Drive <code>Hello()</code> explicitly so a failed EHLO can't be swallowed and mis-read as &quot;STARTTLS not advertised&quot;.</li>
<li><strong><code>log</code>: silence 14 <code>go vet</code> printf warnings.</strong> Split the internal <code>log()</code> method; public signatures unchanged. <code>TestCallerFile</code> guards the <code>runtime.Caller</code> skip.</li>
<li><strong><code>random</code>: fix <code>sync.Pool</code> copy in <code>New()</code>.</strong> Construct the pool directly on the struct — <code>sync.Pool</code> must not be copied after first use.</li>
</ul>
<h2>Toolchain (breaking)</h2>
<ul>
<li>Go directive bumped <code>1.18</code> → <code>1.23.0</code> to align with <code>labstack/echo</code>. Consumers on Go &lt;1.23 should stay on v0.4.2.</li>
<li>CI matrix: <code>1.23 / 1.24 / 1.25 / 1.26</code> × <code>ubuntu / macos / windows</code>.</li>
<li>Deps refreshed: <code>testify 1.8.4 → 1.11.1</code>, <code>go-colorable 0.1.13 → 0.1.14</code>, <code>go-isatty 0.0.20 → 0.0.21</code>, <code>x/sys 0.15.0 → 0.29.0</code> (highest that still supports Go 1.23).</li>
</ul>
<h2>Non-breaking code changes</h2>
<ul>
<li><code>bytes/bytes_test.go</code>: replaced <code>Parse(\&quot;8EiB\&quot;)</code> assertions with <code>Parse(\&quot;7EiB\&quot;)</code> — 2^63 overflowed <code>int64</code> and relied on implementation-defined float-to-int behavior.</li>
</ul>
<h2>Full diff</h2>
<p><a href="https://redirect.github.com/labstack/gommon/pull/62">labstack/gommon#62</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/gommon/commit/2659cdaeb998f92ea22bbeb46892eb0e5d79fda3"><code>2659cda</code></a> fix: SMTPS support + align toolchain with echo (<a href="https://redirect.github.com/labstack/gommon/issues/62">#62</a>)</li>
<li>See full diff in <a href="https://github.com/labstack/gommon/compare/v0.4.2...v0.5.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.5 to 1.41.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bc9c51d733fd437a60a8710531477ca0347dcc4"><code>9bc9c51</code></a> Release 2026-04-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b41455a8d88032b9b76b0bafbe23de0a61a6f29"><code>2b41455</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2327a1b15b5f16414e715325bc9579082fde0767"><code>2327a1b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9225797aa8ce8661d61f4628973aface4e01c922"><code>9225797</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7b6f675adb365c7b069aaf2a8e81c4b8f8a5d08e"><code>7b6f675</code></a> Bump smithy-go to 1.25 (again) (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3390">#3390</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8b6a97220901cccfd8a06f6f25aef65a967384c"><code>c8b6a97</code></a> Release 2026-04-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a05c2ba948715d305aff4519154abe2c4799246"><code>4a05c2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6cd4cc9201382be2638c8b4951b5d92de73c1161"><code>6cd4cc9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3153bed494fb9e426640b30c7b91be0510b20362"><code>3153bed</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe3cef6cc115a2915f5410c739dc73d51553d709"><code>fe3cef6</code></a> Make AccountIDEndpointRouting aware of bdd endpoints (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3387">#3387</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.5...v1.41.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.14 to 1.32.16
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bc9c51d733fd437a60a8710531477ca0347dcc4"><code>9bc9c51</code></a> Release 2026-04-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b41455a8d88032b9b76b0bafbe23de0a61a6f29"><code>2b41455</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2327a1b15b5f16414e715325bc9579082fde0767"><code>2327a1b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9225797aa8ce8661d61f4628973aface4e01c922"><code>9225797</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7b6f675adb365c7b069aaf2a8e81c4b8f8a5d08e"><code>7b6f675</code></a> Bump smithy-go to 1.25 (again) (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3390">#3390</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8b6a97220901cccfd8a06f6f25aef65a967384c"><code>c8b6a97</code></a> Release 2026-04-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a05c2ba948715d305aff4519154abe2c4799246"><code>4a05c2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6cd4cc9201382be2638c8b4951b5d92de73c1161"><code>6cd4cc9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3153bed494fb9e426640b30c7b91be0510b20362"><code>3153bed</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe3cef6cc115a2915f5410c739dc73d51553d709"><code>fe3cef6</code></a> Make AccountIDEndpointRouting aware of bdd endpoints (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3387">#3387</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.14...config/v1.32.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.14 to 1.19.15
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bc9c51d733fd437a60a8710531477ca0347dcc4"><code>9bc9c51</code></a> Release 2026-04-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b41455a8d88032b9b76b0bafbe23de0a61a6f29"><code>2b41455</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2327a1b15b5f16414e715325bc9579082fde0767"><code>2327a1b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9225797aa8ce8661d61f4628973aface4e01c922"><code>9225797</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7b6f675adb365c7b069aaf2a8e81c4b8f8a5d08e"><code>7b6f675</code></a> Bump smithy-go to 1.25 (again) (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3390">#3390</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8b6a97220901cccfd8a06f6f25aef65a967384c"><code>c8b6a97</code></a> Release 2026-04-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a05c2ba948715d305aff4519154abe2c4799246"><code>4a05c2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6cd4cc9201382be2638c8b4951b5d92de73c1161"><code>6cd4cc9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3153bed494fb9e426640b30c7b91be0510b20362"><code>3153bed</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe3cef6cc115a2915f5410c739dc73d51553d709"><code>fe3cef6</code></a> Make AccountIDEndpointRouting aware of bdd endpoints (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3387">#3387</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.14...credentials/v1.19.15">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.68.0 to 1.69.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34ec0a43cbf69475be3bbbb6329c49687cc018df"><code>34ec0a4</code></a> Release 2025-11-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/540441aa03ecb2b3b2b2f1eab1b71661b0ae2481"><code>540441a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9869a2b39a8edb77fc98e0fc70d9bb059c6534a5"><code>9869a2b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a6b38c167b9ce07461c4f0af81be087055267121"><code>a6b38c1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3953a0d499bf4a96a11e3377af3e5291831eed6c"><code>3953a0d</code></a> add explicit message deser to all s3 errors (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3237">#3237</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7eec69683d04dbab33314fb6de43f8a43684163"><code>e7eec69</code></a> Fix panic during auth scheme resolution due to region validation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3235">#3235</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d68b3a05c22b3bde751da6bb70e6fe01fd02407f"><code>d68b3a0</code></a> Release 2025-11-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/292a19869df57271d51b382018591d71f09f72d3"><code>292a198</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dfeabc64ff80e2ee65951eb84d616072c8cd4b60"><code>dfeabc6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a7a1be5d2c14e9270e927654b63272fdfbee1aa6"><code>a7a1be5</code></a> Release 2025-11-21</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.68.0...service/ecs/v1.69.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.99.0 to 1.99.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bc9c51d733fd437a60a8710531477ca0347dcc4"><code>9bc9c51</code></a> Release 2026-04-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b41455a8d88032b9b76b0bafbe23de0a61a6f29"><code>2b41455</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2327a1b15b5f16414e715325bc9579082fde0767"><code>2327a1b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9225797aa8ce8661d61f4628973aface4e01c922"><code>9225797</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7b6f675adb365c7b069aaf2a8e81c4b8f8a5d08e"><code>7b6f675</code></a> Bump smithy-go to 1.25 (again) (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3390">#3390</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8b6a97220901cccfd8a06f6f25aef65a967384c"><code>c8b6a97</code></a> Release 2026-04-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a05c2ba948715d305aff4519154abe2c4799246"><code>4a05c2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6cd4cc9201382be2638c8b4951b5d92de73c1161"><code>6cd4cc9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3153bed494fb9e426640b30c7b91be0510b20362"><code>3153bed</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe3cef6cc115a2915f5410c739dc73d51553d709"><code>fe3cef6</code></a> Make AccountIDEndpointRouting aware of bdd endpoints (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3387">#3387</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.99.0...service/s3/v1.99.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.4.1775140412 to 2026.4.1775836578
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/8dfc8951a4b098dfce2425410d5334671829e8c3"><code>8dfc895</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73d5b434b64037d2e4353b86e...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.4.1775140412...release/v2026.4.1775836578">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.45.0 to 0.45.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.45.1</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Add missing TracesSampler fields for SamplingContext by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1259">#1259</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.45.1</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Add missing TracesSampler fields for SamplingContext by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1259">#1259</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/17971f980de78683f2505ad824635879644787fa"><code>17971f9</code></a> release: 0.45.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/284b2f67bf832b3cc967d62efd6cac5c3ca5b742"><code>284b2f6</code></a> fix: add missing TracesSampler fields for SamplingContext (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1259">#1259</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e1ee192a5cc5ebafbc775a35143524e80961e7d4"><code>e1ee192</code></a> Merge branch 'release/0.45.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.45.0...v0.45.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.9.1 to 5.9.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.9.2 (April 18, 2026)</h1>
<p>Fix SQL Injection via placeholder confusion with dollar quoted string literals (GHSA-j88v-2chj-qfwx)</p>
<p>SQL injection can occur when:</p>
<ol>
<li>The non-default simple protocol is used.</li>
<li>A dollar quoted string literal is used in the SQL query.</li>
<li>That query contains text that would be would be interpreted outside as a placeholder outside of a string literal.</li>
<li>The value of that placeholder is controllable by the attacker.</li>
</ol>
<p>e.g.</p>
<pre lang="go"><code>attackValue := `$tag$; drop table canary; --`
_, err = tx.Exec(ctx, `select $tag$ $1 $tag$, $1`, pgx.QueryExecModeSimpleProtocol, attackValue)
</code></pre>
<p>This is unlikely to occur outside of a contrived scenario.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/0aeabbcf11d859229c1f0b20e710d3596c76bf27"><code>0aeabbc</code></a> Release v5.9.2</li>
<li><a href="https://github.com/jackc/pgx/commit/60644f84918a8af66d14a4b0d865d4edafd955da"><code>60644f8</code></a> Fix SQL sanitizer bugs with dollar-quoted strings and placeholder overflow</li>
<li><a href="https://github.com/jackc/pgx/commit/a5680bc945aa7c6ebac2778d859ee7b4ba86db60"><code>a5680bc</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2531">#2531</a> from dolmen-go/godoc-add-links</li>
<li><a href="https://github.com/jackc/pgx/commit/e34e4524007062710c6a4fb9c8655b75a486b5cd"><code>e34e452</code></a> doc: Add godoc links</li>
<li><a href="https://github.com/jackc/pgx/commit/08c9bb1f0d8fa6cc10ed8c713e68b1baa64dfe2c"><code>08c9bb1</code></a> Fix Stringer types encoded as text instead of numeric value in composite fields</li>
<li><a href="https://github.com/jackc/pgx/commit/96b4dbdfd0458cb425bf8454d292a23978872cc8"><code>96b4dbd</code></a> Remove unstable test</li>
<li><a href="https://github.com/jackc/pgx/commit/acf88e0065682e8948696d26fa6438669c4cabee"><code>acf88e0</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2526">#2526</a> from abrightwell/abrightwell-min-proto</li>
<li><a href="https://github.com/jackc/pgx/commit/2f81f1fc03bef99593e92c64ad9cac954c00e8e6"><code>2f81f1f</code></a> Update <code>max_protocol_version</code> and <code>min_protocol_version</code> defaults</li>
<li>See full diff in <a href="https://github.com/jackc/pgx/compare/v5.9.1...v5.9.2">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on April 20, 2026 at 07:20 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1471*
