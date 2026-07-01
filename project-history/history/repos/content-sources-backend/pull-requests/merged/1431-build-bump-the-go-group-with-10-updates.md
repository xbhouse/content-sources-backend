---
type: pull_request
number: 1431
title: "Build: Bump the go group with 10 updates"
state: merged
author: dependabot
created: 2026-03-30T04:54:55Z
updated: 2026-03-30T07:18:51Z
closed: 2026-03-30T07:18:49Z
merged: 2026-03-30T07:18:49Z
base_branch: main
head_branch: dependabot/go_modules/go-11e28d5c5a
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1431
---

# Pull Request #1431: Build: Bump the go group with 10 updates

**Author**: @dependabot
**Created**: March 30, 2026 at 04:54 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-11e28d5c5a`

## Description

Bumps the go group with 10 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.34.0` | `1.35.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.4` | `1.41.5` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.12` | `1.32.13` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.12` | `1.19.13` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.64.1` | `1.65.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.97.1` | `1.97.3` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.3.1774377330` | `2026.3.1774545586` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.43.0` | `0.44.1` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.6.0` | `1.7.0` |
| [github.com/redhatinsights/platform-go-middlewares/v2](https://github.com/redhatinsights/platform-go-middlewares) | `2.0.0` | `2.1.0` |

Updates `github.com/rs/zerolog` from 1.34.0 to 1.35.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/rs/zerolog/commit/13966551e7cc9f174723bddde692f0f0f9a833aa"><code>1396655</code></a> Bump CI Go matrix minimum from 1.21 to 1.23</li>
<li><a href="https://github.com/rs/zerolog/commit/4b65a2f6f63437b9db07d57733d34d381b8f0696"><code>4b65a2f</code></a> Bump actions/cache from 4 to 5 (<a href="https://redirect.github.com/rs/zerolog/issues/741">#741</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/b83579670fae2aef5d260f1350975e4c0b16864c"><code>b835796</code></a> Bump actions/setup-go from 5 to 6 (<a href="https://redirect.github.com/rs/zerolog/issues/742">#742</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/134caf82aab8ea832731c5b4e5967d970abbc538"><code>134caf8</code></a> Added sanitization of journald keys (<a href="https://redirect.github.com/rs/zerolog/issues/751">#751</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/e133b6a5173130c6ce2349645111829ba80d589e"><code>e133b6a</code></a> Added variadic StrsV, ObjectsV, and StringersV (<a href="https://redirect.github.com/rs/zerolog/issues/752">#752</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/82017d8fff8d7739ba4daf9b1cf096d501ba7732"><code>82017d8</code></a> Bump github.com/coreos/go-systemd/v22 from 22.6.0 to 22.7.0 (<a href="https://redirect.github.com/rs/zerolog/issues/753">#753</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/2f5b8a91be2cea93c2a6792d4ad00f5a06967dae"><code>2f5b8a9</code></a> fix: UpdateContext skips Nop and zero-value loggers (<a href="https://redirect.github.com/rs/zerolog/issues/754">#754</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/d64c9a7138d98305e963fd437246d2ca0f0c4833"><code>d64c9a7</code></a> Add slog.Handler implementation for zerolog (<a href="https://redirect.github.com/rs/zerolog/issues/755">#755</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/a0d61dc2c7439bdea7e1a9920713b1134367be58"><code>a0d61dc</code></a> fix: return dict to Event pool (<a href="https://redirect.github.com/rs/zerolog/issues/749">#749</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/f6fbd330be6e308d528e1f0bd7972c358b41370f"><code>f6fbd33</code></a> Test coverage improvements (<a href="https://redirect.github.com/rs/zerolog/issues/748">#748</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/rs/zerolog/compare/v1.34.0...v1.35.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.4 to 1.41.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90650dd22735ab68f6089ae5c39b6614286ae9ec"><code>90650dd</code></a> Release 2026-03-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dd88818bee7d632a8b9da6e2c78ef92e23c94c62"><code>dd88818</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b662c50138bd393927871b46e84ee3483377f5be"><code>b662c50</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/500a9cb3522a0e71d798d7079ff5856b23c2cac1"><code>500a9cb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6221102f763bd65d7e403fa62c3a1e3d39e24dc6"><code>6221102</code></a> fix stale skew and delayed skew healing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3359">#3359</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a39373433a121800bc68efa743a7486eb07aa3f"><code>0a39373</code></a> fix order of generated event header handlers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3361">#3361</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/098f3898271e2eaaf8a92e38d1d928fb018805a6"><code>098f389</code></a> Only generate resolveAccountID when it's required (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3360">#3360</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6ebab66428e97db0ee252fea042d56b1313cb9f6"><code>6ebab66</code></a> Release 2026-03-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2ec3beebb986a5e74e50d0c105119d84e1e934e"><code>b2ec3be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abc126f6b35bfe2f77e2505f6d04f8ceced971ee"><code>abc126f</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.4...v1.41.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.12 to 1.32.13
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90650dd22735ab68f6089ae5c39b6614286ae9ec"><code>90650dd</code></a> Release 2026-03-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dd88818bee7d632a8b9da6e2c78ef92e23c94c62"><code>dd88818</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b662c50138bd393927871b46e84ee3483377f5be"><code>b662c50</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/500a9cb3522a0e71d798d7079ff5856b23c2cac1"><code>500a9cb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6221102f763bd65d7e403fa62c3a1e3d39e24dc6"><code>6221102</code></a> fix stale skew and delayed skew healing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3359">#3359</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a39373433a121800bc68efa743a7486eb07aa3f"><code>0a39373</code></a> fix order of generated event header handlers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3361">#3361</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/098f3898271e2eaaf8a92e38d1d928fb018805a6"><code>098f389</code></a> Only generate resolveAccountID when it's required (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3360">#3360</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6ebab66428e97db0ee252fea042d56b1313cb9f6"><code>6ebab66</code></a> Release 2026-03-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2ec3beebb986a5e74e50d0c105119d84e1e934e"><code>b2ec3be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abc126f6b35bfe2f77e2505f6d04f8ceced971ee"><code>abc126f</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.12...config/v1.32.13">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.12 to 1.19.13
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90650dd22735ab68f6089ae5c39b6614286ae9ec"><code>90650dd</code></a> Release 2026-03-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dd88818bee7d632a8b9da6e2c78ef92e23c94c62"><code>dd88818</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b662c50138bd393927871b46e84ee3483377f5be"><code>b662c50</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/500a9cb3522a0e71d798d7079ff5856b23c2cac1"><code>500a9cb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6221102f763bd65d7e403fa62c3a1e3d39e24dc6"><code>6221102</code></a> fix stale skew and delayed skew healing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3359">#3359</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a39373433a121800bc68efa743a7486eb07aa3f"><code>0a39373</code></a> fix order of generated event header handlers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3361">#3361</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/098f3898271e2eaaf8a92e38d1d928fb018805a6"><code>098f389</code></a> Only generate resolveAccountID when it's required (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3360">#3360</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6ebab66428e97db0ee252fea042d56b1313cb9f6"><code>6ebab66</code></a> Release 2026-03-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2ec3beebb986a5e74e50d0c105119d84e1e934e"><code>b2ec3be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abc126f6b35bfe2f77e2505f6d04f8ceced971ee"><code>abc126f</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/sqs/v1.19.12...credentials/v1.19.13">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.64.1 to 1.65.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a76a2ae73fe6ae04c8dba07570145eba0582555"><code>7a76a2a</code></a> Release 2024-10-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e35b8bedbb56d7b39d8ccc60cc120a7b61d5fec5"><code>e35b8be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e9587148dadaebdfeda731a68bb30740aedfcdd"><code>6e95871</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d58df32b098fca0d31acca9056f8f56f5c1cdca"><code>0d58df3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/280579adfcd8a5e23658566d3033f40b5e73d447"><code>280579a</code></a> add HTTP client metrics from smithy-go (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2815">#2815</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.64.1...service/s3/v1.65.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.97.1 to 1.97.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90650dd22735ab68f6089ae5c39b6614286ae9ec"><code>90650dd</code></a> Release 2026-03-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dd88818bee7d632a8b9da6e2c78ef92e23c94c62"><code>dd88818</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b662c50138bd393927871b46e84ee3483377f5be"><code>b662c50</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/500a9cb3522a0e71d798d7079ff5856b23c2cac1"><code>500a9cb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6221102f763bd65d7e403fa62c3a1e3d39e24dc6"><code>6221102</code></a> fix stale skew and delayed skew healing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3359">#3359</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a39373433a121800bc68efa743a7486eb07aa3f"><code>0a39373</code></a> fix order of generated event header handlers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3361">#3361</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/098f3898271e2eaaf8a92e38d1d928fb018805a6"><code>098f389</code></a> Only generate resolveAccountID when it's required (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3360">#3360</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6ebab66428e97db0ee252fea042d56b1313cb9f6"><code>6ebab66</code></a> Release 2026-03-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2ec3beebb986a5e74e50d0c105119d84e1e934e"><code>b2ec3be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abc126f6b35bfe2f77e2505f6d04f8ceced971ee"><code>abc126f</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.97.1...service/s3/v1.97.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.3.1774377330 to 2026.3.1774545586
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/1d6132ff12d808c01a5edfea3dd58226d2ad315e"><code>1d6132f</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a972853bd94d9057bd286834ede...</li>
<li><a href="https://github.com/content-services/zest/commit/6c2898714cf47dfdde96ce36e6372d7913211135"><code>6c28987</code></a> update template to not encode params</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.3.1774377330...release/v2026.3.1774545586">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.43.0 to 0.44.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.44.1</h2>
<blockquote>
<p>[!NOTE]<br />
v0.44.0 had to be released as v0.44.1 due to a technical issue.</p>
</blockquote>
<h3>New Features ✨</h3>
<ul>
<li>Add RemoveAttribute api on the scope. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1224">#1224</a></li>
<li>Deprecate <code>Scope.SetExtra</code>, <code>Scope.SetExtras</code>, and <code>Scope.RemoveExtra</code> in favor of <code>Scope.SetAttributes</code> and <code>Scope.RemoveAttribute</code> by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1224">#1224</a>
<ul>
<li>The recommended migration path is to use <code>SetAttributes</code> to attach values to logs and metrics. Note that attributes do not appear on error events; if you only capture errors, use <code>SetTag</code> or <code>SetContext</code> instead.</li>
<li>Before:</li>
</ul>
<pre lang="go"><code>scope.SetExtra(&quot;key.string&quot;, &quot;str&quot;)
scope.SetExtra(&quot;key.int&quot;, 42)
</code></pre>
<ul>
<li>After (for error events) — use tags and contexts:</li>
</ul>
<pre lang="go"><code>scope.SetTag(&quot;key.string&quot;, &quot;str&quot;)
scope.SetContext(&quot;my_data&quot;, sentry.Context{&quot;key.int&quot;: 42})
</code></pre>
<ul>
<li>After (for logs and metrics) — use attributes:</li>
</ul>
<pre lang="go"><code>scope.SetAttributes(
    attribute.String(&quot;key.string&quot;, &quot;str&quot;),
    attribute.Int(&quot;key.int&quot;, 42),
)
</code></pre>
</li>
<li>Add support for homogenous arrays by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1203">#1203</a></li>
<li>Add support for client reports by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1192">#1192</a></li>
<li>Add org id propagation in sentry_baggage by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1210">#1210</a></li>
<li>Add OrgID and StrictTraceContinuation client options. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1210">#1210</a></li>
<li>Add the option to set attributes on the scope by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1208">#1208</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>(serialization) Pre-serialize mutable event fields to prevent race panics by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1214">#1214</a></li>
<li>Use HEROKU_BUILD_COMMIT with HEROKU_SLUG_COMMIT as fallback by <a href="https://github.com/ericapisani"><code>@​ericapisani</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1220">#1220</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Ai</h4>
<ul>
<li>Add AGENTS.md and testing guidelines by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1216">#1216</a></li>
<li>Add dotagents configuration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1211">#1211</a></li>
</ul>
<h4>Deps</h4>
<ul>
<li>Bump github.com/buger/jsonparser from 1.1.1 to 1.1.2 in /zerolog by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1231">#1231</a></li>
<li>Bump github.com/gofiber/fiber/v2 from 2.52.11 to 2.52.12 in /fiber by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1209">#1209</a></li>
</ul>
<h4>Other</h4>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.44.1</h2>
<blockquote>
<p>[!NOTE]<br />
The v0.44.0 is missing due to a technical issue and had to be released again as v0.44.1</p>
</blockquote>
<h3>New Features ✨</h3>
<ul>
<li>Add RemoveAttribute api on the scope. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1224">#1224</a></li>
<li>Deprecate <code>Scope.SetExtra</code>, <code>Scope.SetExtras</code>, and <code>Scope.RemoveExtra</code> in favor of <code>Scope.SetAttributes</code> and <code>Scope.RemoveAttribute</code> by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1224">#1224</a>
<ul>
<li>The recommended migration path is to use <code>SetAttributes</code> to attach values to logs and metrics. Note that attributes do not appear on error events; if you only capture errors, use <code>SetTag</code> or <code>SetContext</code> instead.</li>
<li>Before:</li>
</ul>
<pre lang="go"><code>scope.SetExtra(&quot;key.string&quot;, &quot;str&quot;)
scope.SetExtra(&quot;key.int&quot;, 42)
</code></pre>
<ul>
<li>After (for error events) — use tags and contexts:</li>
</ul>
<pre lang="go"><code>scope.SetTag(&quot;key.string&quot;, &quot;str&quot;)
scope.SetContext(&quot;my_data&quot;, sentry.Context{&quot;key.int&quot;: 42})
</code></pre>
<ul>
<li>After (for logs and metrics) — use attributes:</li>
</ul>
<pre lang="go"><code>scope.SetAttributes(
    attribute.String(&quot;key.string&quot;, &quot;str&quot;),
    attribute.Int(&quot;key.int&quot;, 42),
)
</code></pre>
</li>
<li>Add support for homogenous arrays by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1203">#1203</a></li>
<li>Add support for client reports by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1192">#1192</a></li>
<li>Add org id propagation in sentry_baggage by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1210">#1210</a></li>
<li>Add OrgID and StrictTraceContinuation client options. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1210">#1210</a></li>
<li>Add the option to set attributes on the scope by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1208">#1208</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>(serialization) Pre-serialize mutable event fields to prevent race panics by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1214">#1214</a></li>
<li>Use HEROKU_BUILD_COMMIT with HEROKU_SLUG_COMMIT as fallback by <a href="https://github.com/ericapisani"><code>@​ericapisani</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1220">#1220</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Ai</h4>
<ul>
<li>Add AGENTS.md and testing guidelines by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1216">#1216</a></li>
<li>Add dotagents configuration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1211">#1211</a></li>
</ul>
<h4>Deps</h4>
<ul>
<li>Bump github.com/buger/jsonparser from 1.1.1 to 1.1.2 in /zerolog by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1231">#1231</a></li>
<li>Bump github.com/gofiber/fiber/v2 from 2.52.11 to 2.52.12 in /fiber by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1209">#1209</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/16414f29918c1fb660d4792046f08224dfefc632"><code>16414f2</code></a> release: 0.44.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d26d3ecd4bceb8b900147c2e8faf993c1fdf8720"><code>d26d3ec</code></a> ci: add preReleaseCommand for craft (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1232">#1232</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/65538eb2c2b714f1aec781267612d97e4a4f8d94"><code>65538eb</code></a> build(deps): bump github.com/buger/jsonparser in /zerolog (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1231">#1231</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/91096be62b80b40a2e1d8e62f42115dede34c054"><code>91096be</code></a> chore: pin GitHub Actions to full-length commit SHAs (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1230">#1230</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/df391b0baee60760635a948da32e82e362d7a84e"><code>df391b0</code></a> feat: add RemoveAttribute api on the scope (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1224">#1224</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/340c142cf974aaba7dcb6545101fe125a7d8ad7c"><code>340c142</code></a> fix: TestAsyncTransport_SendEnvelope flakiness (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1226">#1226</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e561a25fe7f0b594a005138bf71db4fd12faacf3"><code>e561a25</code></a> build: Bump getsentry/craft to 2.24.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1225">#1225</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/259b20563d61c7da69d8d83361427e1a8bbd873a"><code>259b205</code></a> feat: add support for homogenous arrays (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1203">#1203</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bf26e9a107879c566124e27d5af51bd1851428d1"><code>bf26e9a</code></a> fix: fix flaky TestAsyncTransport_FlushWithContext (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1222">#1222</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d1edaf4192ea39be6e16ec5145ded927adfed9e9"><code>d1edaf4</code></a> chore(ai): Add AGENTS.md and testing guidelines (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1216">#1216</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.43.0...v0.44.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.6.0 to 1.7.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.7.0</h2>
<ul>
<li>Added CheckForUpdateBulk endpoint</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.6.0...v1.7.0">https://github.com/project-kessel/kessel-sdk-go/compare/v1.6.0...v1.7.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/be8d85c214fcf19788a3a4d3e430925796c19746"><code>be8d85c</code></a> chore: regenerate protobuf code</li>
<li>See full diff in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.6.0...v1.7.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redhatinsights/platform-go-middlewares/v2` from 2.0.0 to 2.1.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redhatinsights/platform-go-middlewares/releases">github.com/redhatinsights/platform-go-middlewares/v2's releases</a>.</em></p>
<blockquote>
<h2>v2.1.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/go-chi/chi/v5 from 5.2.1 to 5.2.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/115">RedHatInsights/platform-go-middlewares#115</a></li>
<li>fix: typo in gobump weekly by <a href="https://github.com/lzap"><code>@​lzap</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/113">RedHatInsights/platform-go-middlewares#113</a></li>
<li>Update to Go 1.21 by <a href="https://github.com/lzap"><code>@​lzap</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/105">RedHatInsights/platform-go-middlewares#105</a></li>
<li>chore: bump dependencies via gobump by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/116">RedHatInsights/platform-go-middlewares#116</a></li>
<li>Bump github.com/onsi/gomega from 1.34.1 to 1.38.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/121">RedHatInsights/platform-go-middlewares#121</a></li>
<li>Bump github.com/go-chi/chi/v5 from 5.2.2 to 5.2.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/124">RedHatInsights/platform-go-middlewares#124</a></li>
<li>Bump actions/setup-go from 5 to 6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/125">RedHatInsights/platform-go-middlewares#125</a></li>
<li>Bump go.uber.org/zap from 1.27.0 to 1.27.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/128">RedHatInsights/platform-go-middlewares#128</a></li>
<li>Bump actions/checkout from 4 to 6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/127">RedHatInsights/platform-go-middlewares#127</a></li>
<li>chore: bump dependencies via gobump by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/141">RedHatInsights/platform-go-middlewares#141</a></li>
<li>feat: add ZapWriteSyncer for zapcore.WriteSyncer compatibility by <a href="https://github.com/charlesmulder"><code>@​charlesmulder</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/142">RedHatInsights/platform-go-middlewares#142</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] made their first contribution in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/116">RedHatInsights/platform-go-middlewares#116</a></li>
<li><a href="https://github.com/charlesmulder"><code>@​charlesmulder</code></a> made their first contribution in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/142">RedHatInsights/platform-go-middlewares#142</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/platform-go-middlewares/compare/v2.0.0...v2.1.0">https://github.com/RedHatInsights/platform-go-middlewares/compare/v2.0.0...v2.1.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/eab6b9f262e51c99285175fa9887a66f89b4b660"><code>eab6b9f</code></a> feat: add ZapWriteSyncer for zapcore.WriteSyncer compatibility (<a href="https://redirect.github.com/redhatinsights/platform-go-middlewares/issues/142">#142</a>)</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/f56dcef5ad63f1a6ca06e5bbbe9b8bd3d91d2c7f"><code>f56dcef</code></a> chore: bump dependencies via gobump</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/2c6264c98d8963f4b520e7a1efc2b943f63a241f"><code>2c6264c</code></a> Bump actions/checkout from 4 to 6</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/a1827e4910d2ecb353d5cb75b8f67792ee99cbb5"><code>a1827e4</code></a> Bump go.uber.org/zap from 1.27.0 to 1.27.1</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/19ca70f21ad60e171d59aa4bcdb2ed2ef91bb59c"><code>19ca70f</code></a> Bump actions/setup-go from 5 to 6</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/97d7bfee00aa8947c49e2ba25a30397aaa927c69"><code>97d7bfe</code></a> Bump github.com/go-chi/chi/v5 from 5.2.2 to 5.2.3</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/f6c3652f3f404bda76d7104503904211b52da2da"><code>f6c3652</code></a> Revert &quot;Bump github.com/onsi/gomega from 1.34.1 to 1.38.1&quot;</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/32c633b92be554e113cfc6ab175b1db451a8af09"><code>32c633b</code></a> Bump github.com/onsi/gomega from 1.34.1 to 1.38.1</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/fbd178d1c30018465d6256aa6f4c6eef6f0107b2"><code>fbd178d</code></a> chore: bump Go dependencies</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/1d43c04a98964546fcd068f0467125bd876d4167"><code>1d43c04</code></a> Update to Go 1.21</li>
<li>Additional commits viewable in <a href="https://github.com/redhatinsights/platform-go-middlewares/compare/v2.0.0...v2.1.0">compare view</a></li>
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

### Review by @swadeley - Approved on March 30, 2026 at 07:18 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1431*
