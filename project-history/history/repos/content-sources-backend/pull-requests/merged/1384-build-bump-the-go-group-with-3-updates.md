---
type: pull_request
number: 1384
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2026-02-02T04:38:44Z
updated: 2026-02-02T18:26:03Z
closed: 2026-02-02T18:26:01Z
merged: 2026-02-02T18:26:01Z
base_branch: main
head_branch: dependabot/go_modules/go-4eb5746356
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1384
---

# Pull Request #1384: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: February 02, 2026 at 04:38 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-4eb5746356`

## Description

Bumps the go group with 3 updates: [github.com/lib/pq](https://github.com/lib/pq), [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) and [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go).

Updates `github.com/lib/pq` from 1.10.9 to 1.11.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/releases">github.com/lib/pq's releases</a>.</em></p>
<blockquote>
<h2>v1.11.1</h2>
<p>This fixes two regressions present in the v1.11.0 release:</p>
<ul>
<li>
<p>Fix build on 32bit systems, Windows, and Plan 9 (<a href="https://redirect.github.com/lib/pq/issues/1253">#1253</a>).</p>
</li>
<li>
<p>Named []byte types and pointers to []byte (e.g. <code>*[]byte</code>, <code>json.RawMessage</code>) would be treated as an array instead of bytea (<a href="https://redirect.github.com/lib/pq/issues/1252">#1252</a>).</p>
</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1252">#1252</a>: <a href="https://redirect.github.com/lib/pq/pull/1252">lib/pq#1252</a>
<a href="https://redirect.github.com/lib/pq/issues/1253">#1253</a>: <a href="https://redirect.github.com/lib/pq/pull/1253">lib/pq#1253</a></p>
<h2>v1.11.0</h2>
<p>This version of pq requires Go 1.21 or newer.</p>
<p>pq now supports only maintained PostgreSQL releases, which is PostgreSQL 14 and newer. Previously PostgreSQL 8.4 and newer were supported.</p>
<h3>Features</h3>
<ul>
<li>
<p>The <code>pq.Error.Error()</code> text  includes the position of the error (if reported by PostgreSQL) and SQLSTATE code (<a href="https://redirect.github.com/lib/pq/issues/1219">#1219</a>, <a href="https://redirect.github.com/lib/pq/issues/1224">#1224</a>):</p>
<pre><code>pq: column &quot;columndoesntexist&quot; does not exist at column 8 (42703)
pq: syntax error at or near &quot;)&quot; at position 2:71 (42601)
</code></pre>
</li>
<li>
<p>The <code>pq.Error.ErrorWithDetail()</code> method prints a more detailed multiline message, with the Detail, Hint, and error position (if any) (<a href="https://redirect.github.com/lib/pq/issues/1219">#1219</a>):</p>
<pre><code>ERROR:   syntax error at or near &quot;)&quot; (42601)
CONTEXT: line 12, column 1:
<pre><code> 10 |     name           varchar,
 11 |     version        varchar,
 12 | );
      ^
</code></pre>
<p></code></pre></p>
</li>
<li>
<p>Add <code>Config</code>, <code>NewConfig()</code>, and <code>NewConnectorConfig()</code> to supply connection details in a more structured way (<a href="https://redirect.github.com/lib/pq/issues/1240">#1240</a>).</p>
</li>
<li>
<p>Support <code>hostaddr</code> and <code>$PGHOSTADDR</code> (<a href="https://redirect.github.com/lib/pq/issues/1243">#1243</a>).</p>
</li>
<li>
<p>Support multiple values in <code>host</code>, <code>port</code>, and <code>hostaddr</code>, which are each tried in order, or randomly if <code>load_balance_hosts=random</code> is set (<a href="https://redirect.github.com/lib/pq/issues/1246">#1246</a>).</p>
</li>
<li>
<p>Support <code>target_session_attrs</code> connection parameter (<a href="https://redirect.github.com/lib/pq/issues/1246">#1246</a>).</p>
</li>
<li>
<p>Support [<code>sslnegotiation</code>] to use SSL without negotiation (<a href="https://redirect.github.com/lib/pq/issues/1180">#1180</a>).</p>
</li>
<li>
<p>Allow using a custom <code>tls.Config</code>, for example for encrypted keys (<a href="https://redirect.github.com/lib/pq/issues/1228">#1228</a>).</p>
</li>
<li>
<p>Add <code>PQGO_DEBUG=1</code> print the communication with PostgreSQL to stderr, to aid in debugging, testing, and bug reports (<a href="https://redirect.github.com/lib/pq/issues/1223">#1223</a>).</p>
</li>
<li>
<p>Add support for NamedValueChecker interface (<a href="https://redirect.github.com/lib/pq/issues/1125">#1125</a>, <a href="https://redirect.github.com/lib/pq/issues/1238">#1238</a>).</p>
</li>
</ul>
<h3>Fixes</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/blob/master/CHANGELOG.md">github.com/lib/pq's changelog</a>.</em></p>
<blockquote>
<h2>v1.11.1 (2025-01-29)</h2>
<p>This fixes two regressions present in the v1.11.0 release:</p>
<ul>
<li>
<p>Fix build on 32bit systems, Windows, and Plan 9 (<a href="https://redirect.github.com/lib/pq/issues/1253">#1253</a>).</p>
</li>
<li>
<p>Named []byte types and pointers to []byte (e.g. <code>*[]byte</code>, <code>json.RawMessage</code>)
would be treated as an array instead of bytea (<a href="https://redirect.github.com/lib/pq/issues/1252">#1252</a>).</p>
</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1252">#1252</a>: <a href="https://redirect.github.com/lib/pq/pull/1252">lib/pq#1252</a>
<a href="https://redirect.github.com/lib/pq/issues/1253">#1253</a>: <a href="https://redirect.github.com/lib/pq/pull/1253">lib/pq#1253</a></p>
<h2>v1.11.0 (2025-01-28)</h2>
<p>This version of pq requires Go 1.21 or newer.</p>
<p>pq now supports only maintained PostgreSQL releases, which is PostgreSQL 14 and
newer. Previously PostgreSQL 8.4 and newer were supported.</p>
<h3>Features</h3>
<ul>
<li>
<p>The <code>pq.Error.Error()</code> text  includes the position of the error (if reported
by PostgreSQL) and SQLSTATE code (<a href="https://redirect.github.com/lib/pq/issues/1219">#1219</a>, <a href="https://redirect.github.com/lib/pq/issues/1224">#1224</a>):</p>
<pre><code>pq: column &quot;columndoesntexist&quot; does not exist at column 8 (42703)
pq: syntax error at or near &quot;)&quot; at position 2:71 (42601)
</code></pre>
</li>
<li>
<p>The <code>pq.Error.ErrorWithDetail()</code> method prints a more detailed multiline
message, with the Detail, Hint, and error position (if any) (<a href="https://redirect.github.com/lib/pq/issues/1219">#1219</a>):</p>
<pre><code>ERROR:   syntax error at or near &quot;)&quot; (42601)
CONTEXT: line 12, column 1:
<pre><code> 10 |     name           varchar,
 11 |     version        varchar,
 12 | );
      ^
</code></pre>
<p></code></pre></p>
</li>
<li>
<p>Add <code>Config</code>, <code>NewConfig()</code>, and <code>NewConnectorConfig()</code> to supply connection
details in a more structured way (<a href="https://redirect.github.com/lib/pq/issues/1240">#1240</a>).</p>
</li>
<li>
<p>Support <code>hostaddr</code> and <code>$PGHOSTADDR</code> (<a href="https://redirect.github.com/lib/pq/issues/1243">#1243</a>).</p>
</li>
<li>
<p>Support multiple values in <code>host</code>, <code>port</code>, and <code>hostaddr</code>, which are each
tried in order, or randomly if <code>load_balance_hosts=random</code> is set (<a href="https://redirect.github.com/lib/pq/issues/1246">#1246</a>).</p>
</li>
<li>
<p>Support <code>target_session_attrs</code> connection parameter (<a href="https://redirect.github.com/lib/pq/issues/1246">#1246</a>).</p>
</li>
<li>
<p>Support [<code>sslnegotiation</code>] to use SSL without negotiation (<a href="https://redirect.github.com/lib/pq/issues/1180">#1180</a>).</p>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lib/pq/commit/eec526cee8f50c61b6294a8299e3ae358ab28d7b"><code>eec526c</code></a> Release v1.11.1 (<a href="https://redirect.github.com/lib/pq/issues/1255">#1255</a>)</li>
<li><a href="https://github.com/lib/pq/commit/1928a1d75260e84d8e4422cf91252cb7577e4fcf"><code>1928a1d</code></a> Fix []byte types incorrectly converted to PostgreSQL array (<a href="https://redirect.github.com/lib/pq/issues/1252">#1252</a>)</li>
<li><a href="https://github.com/lib/pq/commit/9e2aa8e7b098fe3d2743826ffb6bd41db13fac28"><code>9e2aa8e</code></a> Run staticcheck on all GOOS/GOARCH combinations</li>
<li><a href="https://github.com/lib/pq/commit/c9320c42e8721adae182adf2ee3eda3708844ac8"><code>c9320c4</code></a> Fix build on Windows and Plan9</li>
<li><a href="https://github.com/lib/pq/commit/28095269d64d6d501a0e1341e3559755b556e39e"><code>2809526</code></a> Fix build on 32bit systems</li>
<li><a href="https://github.com/lib/pq/commit/8e88f7e928ecc429c45050e378b751c232a35981"><code>8e88f7e</code></a> Release 1.11.0</li>
<li><a href="https://github.com/lib/pq/commit/0ad30496f8aa96a983fce5490ba050523d8fc5d1"><code>0ad3049</code></a> Handle pre-protocol errors to prevent memory exhaustion</li>
<li><a href="https://github.com/lib/pq/commit/f1fae2ee3828fe6e103c0ec4dfcb568d906e5cb6"><code>f1fae2e</code></a> Add pqtest.Fake.Close()</li>
<li><a href="https://github.com/lib/pq/commit/3815d03993a59ea6ffd0206a5dae62913c6ce877"><code>3815d03</code></a> Remove assumption that the auth response is AuthenticateOk</li>
<li><a href="https://github.com/lib/pq/commit/589ad43c7e1f33330cf8f461d2ffca19785f1110"><code>589ad43</code></a> Implement load_balance_hosts=random</li>
<li>Additional commits viewable in <a href="https://github.com/lib/pq/compare/v1.10.9...v1.11.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.95.1 to 1.96.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdb98c543b9f2ddcfde6670b97871fb144ec18e9"><code>bdb98c5</code></a> Release 2026-01-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c878d57e69ca5f53552ba84d850857743b967c22"><code>c878d57</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4f5d5034684faf53e349b53f8f67b1bcb47a2b95"><code>4f5d503</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b7bf952165576a32eec513a5247ff571fac0a35b"><code>b7bf952</code></a> Feat release s3 transfer manager v2 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3293">#3293</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0baa1dcc4d6ea45d92a2292a6c51b3ea276d4359"><code>0baa1dc</code></a> Release 2026-01-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/11eae4b993c32b7640271465becb6dbae44230de"><code>11eae4b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/297caa5614123ece8565d8ee5b4f3de091b79fbb"><code>297caa5</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/43d96e4ef276b89340b8323e139b676f8d00fea9"><code>43d96e4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de58dc6cdc4c35ac4687d53cff781a6027a0f52f"><code>de58dc6</code></a> Release 2026-01-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dba39e60706bddbc976de20328b7c15df9fb6640"><code>dba39e6</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.95.1...service/s3/v1.96.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.41.0 to 0.42.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.42.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>refactor Telemetry Processor to use TelemetryItem instead of ItemConvertible by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1180">#1180</a>
<ul>
<li>remove ToEnvelopeItem from single log items</li>
<li>rename TelemetryBuffer to Telemetry Processor to adhere to spec</li>
<li>remove unsed ToEnvelopeItem(dsn) from Event.</li>
</ul>
</li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add metric support by <a href="https://github.com/aldy505"><code>@​aldy505</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1151">#1151</a>
<ul>
<li>support for three metric methods (counter, gauge, distribution)</li>
<li>custom metric units</li>
<li>unexport batchlogger</li>
</ul>
</li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Release</h4>
<ul>
<li>Fix changelog-preview permissions by <a href="https://github.com/BYK"><code>@​BYK</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1181">#1181</a></li>
<li>Switch from action-prepare-release to Craft by <a href="https://github.com/BYK"><code>@​BYK</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1167">#1167</a></li>
</ul>
<h4>Other</h4>
<ul>
<li>(repo) Add Claude Code settings with basic permissions by <a href="https://github.com/philipphofmann"><code>@​philipphofmann</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1175">#1175</a></li>
<li>Update release and changelog-preview workflows by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1177">#1177</a></li>
<li>Bump echo to 4.10.1 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1174">#1174</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.42.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>refactor Telemetry Processor to use TelemetryItem instead of ItemConvertible by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1180">#1180</a>
<ul>
<li>remove ToEnvelopeItem from single log items</li>
<li>rename TelemetryBuffer to Telemetry Processor to adhere to spec</li>
<li>remove unsed ToEnvelopeItem(dsn) from Event.</li>
</ul>
</li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add metric support by <a href="https://github.com/aldy505"><code>@​aldy505</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1151">#1151</a>
<ul>
<li>support for three metric methods (counter, gauge, distribution)</li>
<li>custom metric units</li>
<li>unexport batchlogger</li>
</ul>
</li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Release</h4>
<ul>
<li>Fix changelog-preview permissions by <a href="https://github.com/BYK"><code>@​BYK</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1181">#1181</a></li>
<li>Switch from action-prepare-release to Craft by <a href="https://github.com/BYK"><code>@​BYK</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1167">#1167</a></li>
</ul>
<h4>Other</h4>
<ul>
<li>(repo) Add Claude Code settings with basic permissions by <a href="https://github.com/philipphofmann"><code>@​philipphofmann</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1175">#1175</a></li>
<li>Update release and changelog-preview workflows by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1177">#1177</a></li>
<li>Bump echo to 4.10.1 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1174">#1174</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/21004229d97bbc308e095c53cb84d82449be411a"><code>2100422</code></a> release: 0.42.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/ba17897c5492456000157ebccd4ba11896be3c50"><code>ba17897</code></a> refactor!: update Telemetry Processor logic (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1180">#1180</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/5d2aa89e695dc575dc8e9d4285528478969e3650"><code>5d2aa89</code></a> feat: Add metrics support (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1151">#1151</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e090e4a55b2dce773817967a6adc2c2ef24d6a25"><code>e090e4a</code></a> ci(release): Fix changelog-preview permissions (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1181">#1181</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9d2368e9525b0c5c4326211c9054b9d8f7482618"><code>9d2368e</code></a> ci: update release and changelog-preview workflows (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1177">#1177</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c551ab56f3dda443c07c205e2b0ecc2c75b123cb"><code>c551ab5</code></a> ci(release): Switch from action-prepare-release to Craft (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1167">#1167</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/21b1d0f4b5121b2963e96e4d0abb60a2a6905995"><code>21b1d0f</code></a> chore(repo): Add Claude Code settings with basic permissions (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1175">#1175</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e84cccc9bd44f977b29085846cee283b4015ba6f"><code>e84cccc</code></a> chore: ignore local Claude settings (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1176">#1176</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/b9f4109494f8a828fe1f5eb8e0fbb5f77d042e65"><code>b9f4109</code></a> chore: bump echo to 4.10.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1174">#1174</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9c96788510b14b8d8beac6384b883e8b1abdee80"><code>9c96788</code></a> Merge branch 'release/0.41.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.41.0...v0.42.0">compare view</a></li>
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

### Comment by @rverdile on February 02, 2026 at 03:05 PM UTC

@dependabot rebase

### Comment by @dependabot on February 02, 2026 at 03:05 PM UTC

Looks like this PR is already up-to-date with main! If you'd still like to recreate it from scratch, overwriting any edits, you can request `@dependabot recreate`.

### Comment by @rverdile on February 02, 2026 at 05:08 PM UTC

@dependabot recreate

---

## Reviews

### Review by @swadeley - Approved on February 02, 2026 at 07:25 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1384*
