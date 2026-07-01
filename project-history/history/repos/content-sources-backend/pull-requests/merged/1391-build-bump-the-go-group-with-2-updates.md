---
type: pull_request
number: 1391
title: "Build: Bump the go group with 2 updates"
state: merged
author: dependabot
created: 2026-02-16T04:36:01Z
updated: 2026-02-16T05:04:44Z
closed: 2026-02-16T05:04:42Z
merged: 2026-02-16T05:04:42Z
base_branch: main
head_branch: dependabot/go_modules/go-2970eee9eb
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1391
---

# Pull Request #1391: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: February 16, 2026 at 04:36 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-2970eee9eb`

## Description

Bumps the go group with 2 updates: [github.com/lib/pq](https://github.com/lib/pq) and [google.golang.org/grpc](https://github.com/grpc/grpc-go).

Updates `github.com/lib/pq` from 1.11.1 to 1.11.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/releases">github.com/lib/pq's releases</a>.</em></p>
<blockquote>
<h2>v1.11.2</h2>
<p>This fixes two regressions:</p>
<ul>
<li>
<p>Don't send startup parameters if there is no value, improving compatibility with Supavisor (<a href="https://redirect.github.com/lib/pq/issues/1260">#1260</a>).</p>
</li>
<li>
<p>Don't send <code>dbname</code> as a startup parameter if <code>database=[..]</code> is used in the connection string. It's recommended to use dbname=, as database= is not a libpq option, and only worked by accident previously. (<a href="https://redirect.github.com/lib/pq/issues/1261">#1261</a>)</p>
</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1260">#1260</a>: <a href="https://redirect.github.com/lib/pq/pull/1260">lib/pq#1260</a>
<a href="https://redirect.github.com/lib/pq/issues/1261">#1261</a>: <a href="https://redirect.github.com/lib/pq/pull/1261">lib/pq#1261</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/blob/master/CHANGELOG.md">github.com/lib/pq's changelog</a>.</em></p>
<blockquote>
<h2>v1.11.2 (2025-02-10)</h2>
<p>This fixes two regressions:</p>
<ul>
<li>
<p>Don't send startup parameters if there is no value, improving compatibility
with Supavisor (<a href="https://redirect.github.com/lib/pq/issues/1260">#1260</a>).</p>
</li>
<li>
<p>Don't send <code>dbname</code> as a startup parameter if <code>database=[..]</code> is used in the
connection string. It's recommended to use dbname=, as database= is not a
libpq option, and only worked by accident previously. (<a href="https://redirect.github.com/lib/pq/issues/1261">#1261</a>)</p>
</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1260">#1260</a>: <a href="https://redirect.github.com/lib/pq/pull/1260">lib/pq#1260</a>
<a href="https://redirect.github.com/lib/pq/issues/1261">#1261</a>: <a href="https://redirect.github.com/lib/pq/pull/1261">lib/pq#1261</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lib/pq/commit/141280560b25659ae3547e631408e5ffa4e127db"><code>1412805</code></a> Don't send empty startup parameters</li>
<li><a href="https://github.com/lib/pq/commit/0c529db1d8376262a44f20886d4f585d3c1b64df"><code>0c529db</code></a> Don't send dbname= as a startup parameter when database= is used</li>
<li>See full diff in <a href="https://github.com/lib/pq/compare/v1.11.1...v1.11.2">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.78.0 to 1.79.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.79.1</h2>
<h1>Bug Fixes</h1>
<ul>
<li>grpc: Remove the <code>-dev</code> suffix from the User-Agent header. (<a href="https://redirect.github.com/grpc/grpc-go/pull/8902">grpc/grpc-go#8902</a>)</li>
</ul>
<h2>Release 1.79.0</h2>
<h1>API Changes</h1>
<ul>
<li>mem: Add experimental API <code>SetDefaultBufferPool</code> to change the default buffer pool. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8806">#8806</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/vanja-p"><code>@​vanja-p</code></a></li>
</ul>
</li>
<li>experimental/stats: Update <code>MetricsRecorder</code> to require embedding the new <code>UnimplementedMetricsRecorder</code> (a no-op struct) in all implementations for forward compatibility. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8780">#8780</a>)</li>
</ul>
<h1>Behavior Changes</h1>
<ul>
<li>balancer/weightedtarget: Remove handling of <code>Addresses</code> and only handle <code>Endpoints</code> in resolver updates. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8841">#8841</a>)</li>
</ul>
<h1>New Features</h1>
<ul>
<li>experimental/stats: Add support for asynchronous gauge metrics through the new <code>AsyncMetricReporter</code> and <code>RegisterAsyncReporter</code> APIs. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8780">#8780</a>)</li>
<li>pickfirst: Add support for weighted random shuffling of endpoints, as described in <a href="https://redirect.github.com/grpc/proposal/pull/535">gRFC A113</a>.
<ul>
<li>This is enabled by default, and can be turned off using the environment variable <code>GRPC_EXPERIMENTAL_PF_WEIGHTED_SHUFFLING</code>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8864">#8864</a>)</li>
</ul>
</li>
<li>xds: Implement <code>:authority</code> rewriting, as specified in <a href="https://github.com/grpc/proposal/blob/master/A81-xds-authority-rewriting.md">gRFC A81</a>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8779">#8779</a>)</li>
<li>balancer/randomsubsetting: Implement the <code>random_subsetting</code> LB policy, as specified in <a href="https://github.com/grpc/proposal/blob/master/A68-random-subsetting.md">gRFC A68</a>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8650">#8650</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/marek-szews"><code>@​marek-szews</code></a></li>
</ul>
</li>
</ul>
<h1>Bug Fixes</h1>
<ul>
<li>credentials/tls: Fix a bug where the port was not stripped from the authority override before validation. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8726">#8726</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/Atul1710"><code>@​Atul1710</code></a></li>
</ul>
</li>
<li>xds/priority: Fix a bug causing delayed failover to lower-priority clusters when a higher-priority cluster is stuck in <code>CONNECTING</code> state. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8813">#8813</a>)</li>
<li>health: Fix a bug where health checks failed for clients using legacy compression options (<code>WithDecompressor</code> or <code>RPCDecompressor</code>). (<a href="https://redirect.github.com/grpc/grpc-go/issues/8765">#8765</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/sanki92"><code>@​sanki92</code></a></li>
</ul>
</li>
<li>transport: Fix an issue where the HTTP/2 server could skip header size checks when terminating a stream early. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8769">#8769</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/joybestourous"><code>@​joybestourous</code></a></li>
</ul>
</li>
<li>server: Propagate status detail headers, if available, when terminating a stream during request header processing. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8754">#8754</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/joybestourous"><code>@​joybestourous</code></a></li>
</ul>
</li>
</ul>
<h1>Performance Improvements</h1>
<ul>
<li>credentials/alts: Optimize read buffer alignment to reduce copies. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8791">#8791</a>)</li>
<li>mem: Optimize pooling and creation of <code>buffer</code> objects.  (<a href="https://redirect.github.com/grpc/grpc-go/issues/8784">#8784</a>)</li>
<li>transport: Reduce slice re-allocations by reserving slice capacity. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8797">#8797</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/782f2de44f597af18a120527e7682a6670d84289"><code>782f2de</code></a> Change version to 1.79.1 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8902">#8902</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/850eccbb2257bd2de6ac28ee88a7172ab6175629"><code>850eccb</code></a> Change version to 1.79.1-dev (<a href="https://redirect.github.com/grpc/grpc-go/issues/8851">#8851</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/765ff056b6890f6c8341894df4e9668e9bfc18ef"><code>765ff05</code></a> Change version to 1.79.0 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8850">#8850</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/68804be0e78ed0365bb5a576dedc12e2168ed63e"><code>68804be</code></a> Cherry pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8864">#8864</a> to v1.79.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/8896">#8896</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/0381eb650acdae8e423473e64eef07693fe36305"><code>0381eb6</code></a> xds: Support <code>:authority</code> header rewriting for LOGICAL_DNS clusters (<a href="https://redirect.github.com/grpc/grpc-go/issues/8822">#8822</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/90f571db95a0ec223ec45187f7399a06ccdc10cf"><code>90f571d</code></a> xds: remove references to ResolverState.Addresses (<a href="https://redirect.github.com/grpc/grpc-go/issues/8841">#8841</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/679565f9ae655079807f5ab10e07f41acd2af943"><code>679565f</code></a> xds: remove <code>HashKey</code> field from <code>xdsresource.Endpoint</code> struct (<a href="https://redirect.github.com/grpc/grpc-go/issues/8844">#8844</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/bb2073d1e5551b900763979e08e1c11a47a8f150"><code>bb2073d</code></a> mem: Allow overriding the default buffer pool. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8806">#8806</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/bd4444a0a2fdd66245f9e0f0d140aafb5b49044c"><code>bd4444a</code></a> Fix flaky <code>TestServer_RedundantUpdateSuppression</code>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8839">#8839</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/623b3f000b3625aa4a1413f90add1ea367db17c2"><code>623b3f0</code></a> test: add regression test for RecvMsg() error shadowing <a href="https://redirect.github.com/grpc/grpc-go/issues/7510">#7510</a> (<a href="https://redirect.github.com/grpc/grpc-go/issues/8820">#8820</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/grpc/grpc-go/compare/v1.78.0...v1.79.1">compare view</a></li>
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

### Review by @swadeley - Approved on February 16, 2026 at 05:04 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1391*
