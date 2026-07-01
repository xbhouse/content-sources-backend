---
type: pull_request
number: 1447
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2026-04-06T04:53:53Z
updated: 2026-04-06T13:30:53Z
closed: 2026-04-06T13:30:51Z
merged: 2026-04-06T13:30:51Z
base_branch: main
head_branch: dependabot/go_modules/go-85b6aff00b
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1447
---

# Pull Request #1447: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: April 06, 2026 at 04:53 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-85b6aff00b`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/lib/pq](https://github.com/lib/pq) | `1.12.0` | `1.12.3` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.13` | `1.32.14` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.13` | `1.19.14` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.65.0` | `1.68.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.97.3` | `1.98.0` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.79.3` | `1.80.0` |

Updates `github.com/lib/pq` from 1.12.0 to 1.12.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/releases">github.com/lib/pq's releases</a>.</em></p>
<blockquote>
<h2>v1.12.3</h2>
<ul>
<li>Send datestyle startup parameter, improving compatbility with database engines that use a different default datestyle such as EnterpriseDB (<a href="https://redirect.github.com/lib/pq/issues/1312">#1312</a>).</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1312">#1312</a>: <a href="https://redirect.github.com/lib/pq/pull/1312">lib/pq#1312</a></p>
<h2>v1.12.2</h2>
<ul>
<li>Treat io.ErrUnexpectedEOF as driver.ErrBadConn so database/sql discards the connection. Since v1.12.0 this could result in permanently broken connections, especially with CockroachDB which frequently sends partial messages (<a href="https://redirect.github.com/lib/pq/issues/1299">#1299</a>).</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1299">#1299</a>: <a href="https://redirect.github.com/lib/pq/pull/1299">lib/pq#1299</a></p>
<h2>v1.12.1</h2>
<ul>
<li>
<p>Look for pgpass file in ~/.pgpass instead of ~/.postgresql/pgpass (<a href="https://redirect.github.com/lib/pq/issues/1300">#1300</a>).</p>
</li>
<li>
<p>Don't clear password if directly set on pq.Config (<a href="https://redirect.github.com/lib/pq/issues/1302">#1302</a>).</p>
</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1300">#1300</a>: <a href="https://redirect.github.com/lib/pq/pull/1300">lib/pq#1300</a>
<a href="https://redirect.github.com/lib/pq/issues/1302">#1302</a>: <a href="https://redirect.github.com/lib/pq/pull/1302">lib/pq#1302</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/lib/pq/blob/master/CHANGELOG.md">github.com/lib/pq's changelog</a>.</em></p>
<blockquote>
<h2>v1.12.3 (2026-04-03)</h2>
<ul>
<li>Send datestyle startup parameter, improving compatbility with database engines
that use a different default datestyle such as EnterpriseDB (<a href="https://redirect.github.com/lib/pq/issues/1312">#1312</a>).</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1312">#1312</a>: <a href="https://redirect.github.com/lib/pq/pull/1312">lib/pq#1312</a></p>
<h2>v1.12.2 (2026-04-02)</h2>
<ul>
<li>Treat io.ErrUnexpectedEOF as driver.ErrBadConn so database/sql discards the
connection. Since v1.12.0 this could result in permanently broken connections,
especially with CockroachDB which frequently sends partial messages (<a href="https://redirect.github.com/lib/pq/issues/1299">#1299</a>).</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1299">#1299</a>: <a href="https://redirect.github.com/lib/pq/pull/1299">lib/pq#1299</a></p>
<h2>v1.12.1 (2026-03-30)</h2>
<ul>
<li>
<p>Look for pgpass file in ~/.pgpass instead of ~/.postgresql/pgpass (<a href="https://redirect.github.com/lib/pq/issues/1300">#1300</a>).</p>
</li>
<li>
<p>Don't clear password if directly set on pq.Config (<a href="https://redirect.github.com/lib/pq/issues/1302">#1302</a>).</p>
</li>
</ul>
<p><a href="https://redirect.github.com/lib/pq/issues/1300">#1300</a>: <a href="https://redirect.github.com/lib/pq/pull/1300">lib/pq#1300</a>
<a href="https://redirect.github.com/lib/pq/issues/1302">#1302</a>: <a href="https://redirect.github.com/lib/pq/pull/1302">lib/pq#1302</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lib/pq/commit/1f3e3d92865dd313b4e146968684d7e3836c76e8"><code>1f3e3d9</code></a> Send datestyle as a startup parameter (<a href="https://redirect.github.com/lib/pq/issues/1312">#1312</a>)</li>
<li><a href="https://github.com/lib/pq/commit/32ba56b8f9c09575e3320f0043f4f0bdf0ad2009"><code>32ba56b</code></a> Expand tests for multiple result sets</li>
<li><a href="https://github.com/lib/pq/commit/c2cfac15d5048670f784616c0c3dca56f97f49c0"><code>c2cfac1</code></a> Release v1.12.2</li>
<li><a href="https://github.com/lib/pq/commit/859f10493799ae5b3fc3706bbef2ee48764dc787"><code>859f104</code></a> Test CockroachDB</li>
<li><a href="https://github.com/lib/pq/commit/12e464c3afecfb945fc764001837c137fa764e37"><code>12e464c</code></a> Allow multiple matches and regexps in pqtest.ErrorContains()</li>
<li><a href="https://github.com/lib/pq/commit/6d77ced41719616090c9e7eec2c313a18640bc3f"><code>6d77ced</code></a> Treat io.ErrUnexpectedEOF as driver.ErrBadConn in handleError</li>
<li><a href="https://github.com/lib/pq/commit/71daecbc4522cf9cb6c399e19b910d22356ebb87"><code>71daecb</code></a> Ensure transactions are closed in pqtest</li>
<li><a href="https://github.com/lib/pq/commit/8f448230b50d3c2f796fd20622daaf8ebe3d173c"><code>8f44823</code></a> Set PGAPPNAME for tests</li>
<li><a href="https://github.com/lib/pq/commit/4af2196aa02298c23461f2baf538a0679b66a093"><code>4af2196</code></a> Fix healthcheck</li>
<li><a href="https://github.com/lib/pq/commit/38a54e44b0a91e12314291c9102714e7f503ba98"><code>38a54e4</code></a> Split out testdata/init a bit</li>
<li>Additional commits viewable in <a href="https://github.com/lib/pq/compare/v1.12.0...v1.12.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.13 to 1.32.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d15107640a5073c5ce53dc395210858e316d5e82"><code>d151076</code></a> Release 2026-04-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e33c9a9f67577e2dd5fcd3b15b9a8a2e83683f53"><code>e33c9a9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e06655901c88ef90432f41a7b9c92b33b4537c75"><code>e066559</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a635ee4303d8cac381be8788519a1a8bba5ccf12"><code>a635ee4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9074b3ddace9c195f00993c8ae509bc8977a6c4c"><code>9074b3d</code></a> Release 2026-04-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f6ad4c082d3a1995b9add26af4eb7553b13df252"><code>f6ad4c0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/12a971a61d5226cd72102c66b188ae4968cb9b81"><code>12a971a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8bd8eee331c2fa1553e57d18d5dbbf1dc83820fc"><code>8bd8eee</code></a> chore: add additional text to CONTRIBUTING.md (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3372">#3372</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4deb657ddb9c5d8ac36c335391aafb19034e9de"><code>e4deb65</code></a> Release 2026-03-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1f758f207351302f9f4e1a2aa4c255a75d81446b"><code>1f758f2</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.13...config/v1.32.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.13 to 1.19.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d15107640a5073c5ce53dc395210858e316d5e82"><code>d151076</code></a> Release 2026-04-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e33c9a9f67577e2dd5fcd3b15b9a8a2e83683f53"><code>e33c9a9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e06655901c88ef90432f41a7b9c92b33b4537c75"><code>e066559</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a635ee4303d8cac381be8788519a1a8bba5ccf12"><code>a635ee4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9074b3ddace9c195f00993c8ae509bc8977a6c4c"><code>9074b3d</code></a> Release 2026-04-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f6ad4c082d3a1995b9add26af4eb7553b13df252"><code>f6ad4c0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/12a971a61d5226cd72102c66b188ae4968cb9b81"><code>12a971a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8bd8eee331c2fa1553e57d18d5dbbf1dc83820fc"><code>8bd8eee</code></a> chore: add additional text to CONTRIBUTING.md (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3372">#3372</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4deb657ddb9c5d8ac36c335391aafb19034e9de"><code>e4deb65</code></a> Release 2026-03-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1f758f207351302f9f4e1a2aa4c255a75d81446b"><code>1f758f2</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.13...credentials/v1.19.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.65.0 to 1.68.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18fea5d7df1f6a923b62e7d3d38663ff61814878"><code>18fea5d</code></a> Release 2024-11-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/480ffa8ead88c95db859a612652abadada25dc95"><code>480ffa8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4075928ac9d98d563b31ef28d984f9e2ec2794e3"><code>4075928</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a88650657acbec1d0814bbe89adae439c53842f"><code>2a88650</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c527ee8b854b0a2813c0332b1de89357e2a5621f"><code>c527ee8</code></a> Release 2024-11-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fdd683997b576a3d0fc469b0e9ba66667f704a1d"><code>fdd6839</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6809be00b712a49abef6f172f86360f8a2ece61a"><code>6809be0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f040881463aa7a5995e5484c1bda820c57e55f98"><code>f040881</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/249090ec218d0e6ab266a47a759baab368bb6b1f"><code>249090e</code></a> Release 2024-11-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c158e1685a3f4277a93be5030f469437386c4920"><code>c158e16</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.65.0...service/s3/v1.68.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.97.3 to 1.98.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4deb657ddb9c5d8ac36c335391aafb19034e9de"><code>e4deb65</code></a> Release 2026-03-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1f758f207351302f9f4e1a2aa4c255a75d81446b"><code>1f758f2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba7e432545fa3203d98f593b2aceaba66c02db7a"><code>ba7e432</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/607cb0afad6e6a655a356be54f9fdec5cc558a80"><code>607cb0a</code></a> Release 2026-03-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a44005fd4a2e2b7308c0346da970890757aaeda5"><code>a44005f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dbbd8465f03f1931d85d38af45f18362cbc4c469"><code>dbbd846</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5b5c3f960e98cf024a8b7d6a6a8f0a1452e028fa"><code>5b5c3f9</code></a> Revert &quot;drop service/internal/benchmark (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3368">#3368</a>)&quot; (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3369">#3369</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ca3f9d52a718e4be3a78be0d05e6437ae362e20"><code>7ca3f9d</code></a> drop service/internal/benchmark (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3368">#3368</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/338088bc3ce801d773863578da6a0a835d6012af"><code>338088b</code></a> Release 2026-03-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0e5f3dc09ddb548bccfc11b297229839516bff9"><code>f0e5f3d</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.97.3...service/s3/v1.98.0">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.79.3 to 1.80.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.80.0</h2>
<h1>Behavior Changes</h1>
<ul>
<li>balancer: log a warning if a balancer is registered with uppercase letters, as balancer names should be lowercase. In a future release, balancer names will be treated as case-insensitive; see <a href="https://redirect.github.com/grpc/grpc-go/issues/5288">#5288</a> for details. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8837">#8837</a>)</li>
<li>xds: update resource error handling and re-resolution logic (<a href="https://redirect.github.com/grpc/grpc-go/issues/8907">#8907</a>)
<ul>
<li>Re-resolve all <code>LOGICAL_DNS</code> clusters simultaneously when re-resolution is requested.</li>
<li>Fail all in-flight RPCs immediately upon receipt of listener or route resource errors, instead of allowing them to complete.</li>
</ul>
</li>
</ul>
<h1>Bug Fixes</h1>
<ul>
<li>xds: support the LB policy configured in <code>LOGICAL_DNS</code> cluster resources instead of defaulting to <code>pick_first</code>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8733">#8733</a>)</li>
<li>credentials/tls: perform per-RPC authority validation against the leaf certificate instead of the entire peer certificate chain. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8831">#8831</a>)</li>
<li>xds: enabling A76 ring hash endpoint keys no longer causes EDS resources with invalid proxy metadata to be NACKed when HTTP CONNECT (gRFC A86) is disabled. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8875">#8875</a>)</li>
<li>xds: validate that the sum of endpoint weights in a locality does not exceed the maximum <code>uint32</code> value. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8899">#8899</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/RAVEYUS"><code>@​RAVEYUS</code></a></li>
</ul>
</li>
<li>xds: fix incorrect proto field access in the weighted round robin (WRR) configuration where <code>blackout_period</code> was used instead of <code>weight_expiration_period</code>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8915">#8915</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/gregbarasch"><code>@​gregbarasch</code></a></li>
</ul>
</li>
<li>xds/rbac: handle addresses with ports in IP matchers. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8990">#8990</a>)</li>
</ul>
<h1>New Features</h1>
<ul>
<li>ringhash: enable gRFC A76 (endpoint hash keys and request hash headers) by default. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8922">#8922</a>)</li>
</ul>
<h1>Performance Improvements</h1>
<ul>
<li>credentials/alts: pool write buffers to reduce memory allocations and usage. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8919">#8919</a>)</li>
<li>grpc: enable the use of pooled write buffers for buffering HTTP/2 frame writes by default. This reduces memory usage when connections are idle. Use the <a href="https://pkg.go.dev/google.golang.org/grpc#WithSharedWriteBuffer">WithSharedWriteBuffer</a> dial option or the <a href="https://pkg.go.dev/google.golang.org/grpc#SharedWriteBuffer">SharedWriteBuffer</a> server option to disable this feature. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8957">#8957</a>)</li>
<li>xds/priority: stop caching child LB policies removed from the configuration. This will help reduce memory and cpu usage when localities are constantly switching between priorities. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8997">#8997</a>)</li>
<li>mem: add a faster tiered buffer pool; use the experimental <a href="https://pkg.go.dev/google.golang.org/grpc/mem@master#NewBinaryTieredBufferPool">mem.NewBinaryTieredBufferPool</a> function to create such pools. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8775">#8775</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/397e45edaa68f8763773bbaaf539cf7894169cd2"><code>397e45e</code></a> Change version to 1.80.0 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8948">#8948</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/64ebf0a600005838970e6ba1eb0a9e46e528ed73"><code>64ebf0a</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8997">#8997</a> to v1.80.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/9027">#9027</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/e45ed241865981b6973cdd0dd69571456d570282"><code>e45ed24</code></a> xds/rbac: add additional handling for addresses with ports (<a href="https://redirect.github.com/grpc/grpc-go/issues/8990">#8990</a>) (<a href="https://redirect.github.com/grpc/grpc-go/issues/9022">#9022</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/c78d26e03e129f5cb357b757037fcded2333b74e"><code>c78d26e</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8957">#8957</a> to v1.80.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/9007">#9007</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/bd7cd3c1abbd27fb751275a58886444d52103482"><code>bd7cd3c</code></a> grpc: enforce strict path checking for incoming requests on the server (<a href="https://redirect.github.com/grpc/grpc-go/issues/8987">#8987</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/b6597b3d328c1ed6b003f9a23b942af7148352ca"><code>b6597b3</code></a> xds/clusterimpl: use xdsConfig for updates and remove redundant fields from L...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/1d4fa8a7b772553e82137b059ad4a8f632a1c522"><code>1d4fa8a</code></a> xds: change cdsbalancer to use update from dependency manager (<a href="https://redirect.github.com/grpc/grpc-go/issues/8907">#8907</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/8f47d364511c8eb0517b47e1a39f13a1370c6a10"><code>8f47d36</code></a> attributes: Replace internal map with linked list (<a href="https://redirect.github.com/grpc/grpc-go/issues/8933">#8933</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/22e1ee8085952b4bdadf2928c187d665f6daff99"><code>22e1ee8</code></a> xds: add panic recovery in xdsclient resource unmarshalling. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8895">#8895</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/7136e99ee323c26984174eb3cec85c201fef9946"><code>7136e99</code></a> credentials/alts: Pool write buffers (<a href="https://redirect.github.com/grpc/grpc-go/issues/8919">#8919</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/grpc/grpc-go/compare/v1.79.3...v1.80.0">compare view</a></li>
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

### Review by @TenSt - Approved on April 06, 2026 at 01:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1447*
