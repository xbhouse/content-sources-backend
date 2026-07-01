---
type: pull_request
number: 1493
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2026-05-11T05:12:32Z
updated: 2026-05-11T07:35:27Z
closed: 2026-05-11T07:35:25Z
merged: 2026-05-11T07:35:25Z
base_branch: main
head_branch: dependabot/go_modules/go-3c4b4a53cb
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1493
---

# Pull Request #1493: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: May 11, 2026 at 05:12 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-3c4b4a53cb`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.48.0` | `1.48.1` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.72.0` | `1.73.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.100.1` | `1.101.0` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.4.1777255680` | `2026.4.1777574211` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.46.1` | `0.46.2` |
| [github.com/getsentry/sentry-go/zerolog](https://github.com/getsentry/sentry-go) | `0.46.1` | `0.46.2` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.80.0` | `1.81.0` |

Updates `github.com/IBM/sarama` from 1.48.0 to 1.48.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.48.1 (2026-05-10)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>perf: cache topic batch-size metric lookup by <a href="https://github.com/huynhanx03"><code>@​huynhanx03</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3498">IBM/sarama#3498</a></li>
<li>fix: stabilise TestFuncTxnProduceAndCommitOffset flakes by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3517">IBM/sarama#3517</a></li>
<li>test: relax producer batch metrics assertions by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3523">IBM/sarama#3523</a></li>
<li>fix: prevent race during partition consumer close by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3524">IBM/sarama#3524</a></li>
<li>fix: return leaderless errors in metadata refresh by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3525">IBM/sarama#3525</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update dependency golangci/golangci-lint to v2.12.1 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3509">IBM/sarama#3509</a></li>
<li>chore(deps): bump github.com/klauspost/compress from 1.18.5 to 1.18.6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3508">IBM/sarama#3508</a></li>
<li>chore(deps): bump golang.org/x/sys from 0.43.0 to 0.44.0 in the golang-x group across 1 directory by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3520">IBM/sarama#3520</a></li>
<li>chore(deps): update module golang.org/x/crypto to v0.51.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3521">IBM/sarama#3521</a></li>
<li>fix(deps): update module golang.org/x/net to v0.54.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3526">IBM/sarama#3526</a></li>
<li>chore(deps): update dependency golangci/golangci-lint to v2.12.2 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3515">IBM/sarama#3515</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: add testifylint and fix lint warnings by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3522">IBM/sarama#3522</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/huynhanx03"><code>@​huynhanx03</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3498">IBM/sarama#3498</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.48.0...v1.48.1">https://github.com/IBM/sarama/compare/v1.48.0...v1.48.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/21ec56a94abf6338c64709235a46f4903038ea1d"><code>21ec56a</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.7 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/f671778d9f35db3e235383751dc5b9ec2cf261a2"><code>f671778</code></a> chore(deps): update dependency golangci/golangci-lint to v2.12.2 (<a href="https://redirect.github.com/IBM/sarama/issues/3515">#3515</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3ec1ad6adf178c317935acaaa23d869366306e4a"><code>3ec1ad6</code></a> chore(ci): Update github/codeql-action action to v4.35.4 (<a href="https://redirect.github.com/IBM/sarama/issues/3510">#3510</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/b485cf30fdd900fddbc36d4745001b47c940d179"><code>b485cf3</code></a> fix(deps): update module golang.org/x/net to v0.54.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3526">#3526</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/1fe27c9d8001e28ec9761aaef898d53d72d63161"><code>1fe27c9</code></a> fix: return leaderless errors in metadata refresh (<a href="https://redirect.github.com/IBM/sarama/issues/3525">#3525</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/fa2e3f38a897954e0e0e82176b4b9e95d2b77e5d"><code>fa2e3f3</code></a> fix: prevent race during partition consumer close (<a href="https://redirect.github.com/IBM/sarama/issues/3524">#3524</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/67f8fcee59c47a9d989e45a0d95fb17abeac0f69"><code>67f8fce</code></a> chore: add testifylint and fix lint warnings</li>
<li><a href="https://github.com/IBM/sarama/commit/fcf0655f026814c92a071a3ffd162393268740fb"><code>fcf0655</code></a> chore: move FVT-only variables to fvt files</li>
<li><a href="https://github.com/IBM/sarama/commit/55049cceb15e3ba3302f9f08652d593c171ea093"><code>55049cc</code></a> chore(deps): update module golang.org/x/crypto to v0.51.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3521">#3521</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a471f29f5691624ce48189563d3cbb915f984324"><code>a471f29</code></a> test: relax producer batch metrics assertions (<a href="https://redirect.github.com/IBM/sarama/issues/3523">#3523</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.48.0...v1.48.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.72.0 to 1.73.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8fac7d48b4707a8fdd9cb23b34b928fc83e38777"><code>8fac7d4</code></a> Release 2025-01-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9eff8acbf1cc8bb294171b969a8e2803bf235a07"><code>9eff8ac</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f47aa72637a4492512213811a314dda0b9c9a189"><code>f47aa72</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5934dd36b8ecbc42a321216e8da721e1258ae8a8"><code>5934dd3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6636822440828c3eebaacfe9a182c9eb47895236"><code>6636822</code></a> feat: flexible checksum updates (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2808">#2808</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ffbb7cbe6bfeebb50a773d45df57d3fd126cafb"><code>4ffbb7c</code></a> Release 2025-01-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f70ff490c3474a59ca29dc2661ec5e4f13ec2f7"><code>2f70ff4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/434159f5c088eb7b484fea9179d6c5e6302e4610"><code>434159f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a80fac9aee1efaac1d16f4005ccc93fed22c62c3"><code>a80fac9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44e779abca3ec6dfdc9415ced6af0aa684149bd1"><code>44e779a</code></a> fix several waiter issues (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2953">#2953</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.72.0...service/s3/v1.73.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.100.1 to 1.101.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/84ddd08980744ac0f3bacfe7d2796c861671accc"><code>84ddd08</code></a> Release 2026-05-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23645b402c11575a32d1af93ec237f2f121dd285"><code>23645b4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/848eb597bd63cb770bcd11ee230f9dca68ce1ffe"><code>848eb59</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1d7b13420ed32f0809a5a4a35c630d19bfd0b6d4"><code>1d7b134</code></a> Release 2026-05-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5bbbc97fcab2086757060561ccd1f34101ba7b57"><code>5bbbc97</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8dbb93618325675f55eb0c3eb0c7a766806dadbc"><code>8dbb936</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/369e6498c716c7e02ca7aef318ef311edd5efcbf"><code>369e649</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dc2d13fa6f1db25f1c6d804567e1ecfcdff4f040"><code>dc2d13f</code></a> Release 2026-05-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/da4bcffa913dc4ba46e1ce10a6268bf075547a8d"><code>da4bcff</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a8b1180254cba3e23aa6baece26783395e884d81"><code>a8b1180</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.100.1...service/s3/v1.101.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.4.1777255680 to 2026.4.1777574211
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/d8977e1544115bab2a038825d657de043cb47953"><code>d8977e1</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97923884ee38f37ae5486b92ad...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.4.1777255680...release/v2026.4.1777574211">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.46.1 to 0.46.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.46.2</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Add attachments to new event path by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1295">#1295</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.46.2</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Add attachments to new event path by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1295">#1295</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/1d2598e7580f52f201f06ce6b5d819c11a977f4c"><code>1d2598e</code></a> release: 0.46.2</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/57175c67c4665610f5112a1beecc96178d0bd28f"><code>57175c6</code></a> fix: flaky attachment test (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1296">#1296</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/8d2146849fa2c7fcc2e679367ef9c06959f65e43"><code>8d21468</code></a> fix: add attachments to new event path (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1295">#1295</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e4bcedde0a0f2aa1b8999a6ba72e6c5b174d74a0"><code>e4bcedd</code></a> Merge branch 'release/0.46.1'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.46.1...v0.46.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go/zerolog` from 0.46.1 to 0.46.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go/zerolog's releases</a>.</em></p>
<blockquote>
<h2>0.46.2</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Add attachments to new event path by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1295">#1295</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go/zerolog's changelog</a>.</em></p>
<blockquote>
<h2>0.46.2</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Add attachments to new event path by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1295">#1295</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/1d2598e7580f52f201f06ce6b5d819c11a977f4c"><code>1d2598e</code></a> release: 0.46.2</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/57175c67c4665610f5112a1beecc96178d0bd28f"><code>57175c6</code></a> fix: flaky attachment test (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1296">#1296</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/8d2146849fa2c7fcc2e679367ef9c06959f65e43"><code>8d21468</code></a> fix: add attachments to new event path (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1295">#1295</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e4bcedde0a0f2aa1b8999a6ba72e6c5b174d74a0"><code>e4bcedd</code></a> Merge branch 'release/0.46.1'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.46.1...v0.46.2">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.80.0 to 1.81.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.81.0</h2>
<h1>Behavior Changes</h1>
<ul>
<li>balancer/rls: Switch gauge metrics to asynchronous emission (once per collection cycle) to reduce telemetry noise and align with other gRPC language implementations. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8808">#8808</a>)</li>
</ul>
<h1>Dependencies</h1>
<ul>
<li>Minimum supported Go version is now 1.25. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8969">#8969</a>)</li>
</ul>
<h1>Bug Fixes</h1>
<ul>
<li>xds: Use the leaf cluster's security config for the TLS handshake instead of the aggregate cluster's config. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8956">#8956</a>)</li>
<li>transport: Send a <code>RST_STREAM</code> when receiving an <code>END_STREAM</code> when the stream is not already half-closed. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8832">#8832</a>)</li>
<li>xds: Fix ADS resource name validation to prevent a panic. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8970">#8970</a>)</li>
</ul>
<h1>New Features</h1>
<ul>
<li>grpc/stats: Add support for custom labels in per-call metrics (<a href="https://github.com/grpc/proposal/blob/master/A108-otel-custom-per-call-label.md">gRFC A108</a>). (<a href="https://redirect.github.com/grpc/grpc-go/issues/9008">#9008</a>)</li>
<li>xds: Add support for Server Name Indication (SNI) and SAN validation (<a href="https://github.com/grpc/proposal/blob/master/A101-SNI-setting-and-SNI-SAN-validation.md">gRFC A101</a>). Disabled by default. To enable, set <code>GRPC_EXPERIMENTAL_XDS_SNI=true</code> environment variable. (<a href="https://redirect.github.com/grpc/grpc-go/issues/9016">#9016</a>)</li>
<li>xds: Add support to control which fields get propagated from ORCA backend metric reports to LRS load reports (<a href="https://github.com/grpc/proposal/blob/master/A85-lrs-custom-metrics-changes.md">gRFC A85</a>). Disabled by default. To enable, set <code>GRPC_EXPERIMENTAL_XDS_ORCA_LRS_PROPAGATION=true</code>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/9005">#9005</a>)</li>
<li>xds: Add metrics to track xDS client connectivity and cached resource state (<a href="https://github.com/grpc/proposal/blob/master/A78-grpc-metrics-wrr-pf-xds.md">gRFC A78</a>). (<a href="https://redirect.github.com/grpc/grpc-go/issues/8807">#8807</a>)</li>
<li>stats/otel: Enhance <code>grpc.subchannel.disconnections</code> metric by adding disconnection reason to the <code>grpc.disconnect_error</code> label (<a href="https://github.com/grpc/proposal/blob/master/A94-subchannel-otel-metrics.md">gRFC A94</a>). This provides granular insights into why subchannels are closing. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8973">#8973</a>)</li>
<li>mem: Add <code>mem.Buffer.Slice()</code> API to slice the buffer like a slice. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8977">#8977</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/ash2k"><code>@​ash2k</code></a></li>
</ul>
</li>
</ul>
<h1>Performance Improvements</h1>
<ul>
<li>alts: Pool read buffers to lower memory utilization when sockets are unreadable. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8964">#8964</a>)</li>
<li>transport: Pool HTTP/2 framer read buffers to reduce idle memory consumption. Currently limited to Linux for ALTS and non-encrypted transports (TCP, Unix). To disable, set <code>GRPC_GO_EXPERIMENTAL_HTTP_FRAMER_READ_BUFFER_POOLING=false</code> and report any issues. (<a href="https://redirect.github.com/grpc/grpc-go/issues/9032">#9032</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/cb18228317ff523e63d931b4058b0329585b7dcd"><code>cb18228</code></a> Change version to 1.81.0 (<a href="https://redirect.github.com/grpc/grpc-go/issues/9062">#9062</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/96748f973e20bbfcafa19a8bdffc85ad5da138d1"><code>96748f9</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/9105">#9105</a> to 1.81.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/9106">#9106</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/91832222f0144f76527b630ca55cfea6e1aa015a"><code>9183222</code></a> Cherry pick <a href="https://redirect.github.com/grpc/grpc-go/issues/9055">#9055</a>, <a href="https://redirect.github.com/grpc/grpc-go/issues/9032">#9032</a> to v1.81.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/9095">#9095</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/5cba6da4211f3b130238c792937f5921741b616a"><code>5cba6da</code></a> Revert &quot;deps: update dependencies for all modules (<a href="https://redirect.github.com/grpc/grpc-go/issues/9065">#9065</a>)&quot; (<a href="https://redirect.github.com/grpc/grpc-go/issues/9067">#9067</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/af8a9364aa7523ab24d214e9ef13e6ad64d5c5f9"><code>af8a936</code></a> deps: update dependencies for all modules (<a href="https://redirect.github.com/grpc/grpc-go/issues/9065">#9065</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/cdc60dfaaadde45e16aa3c28237c0e655a722c1a"><code>cdc60df</code></a> transport: optimize heap allocations in ready reader and update syscall conne...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/208d053e3204c806ba9e6205c26aa064c8b42852"><code>208d053</code></a> xds/resolver: pass complete XDSConfig in RPC context for HTTP filters (gRFC A...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/50fe1cc7fd78b78ae638ed90ea78514c934167ac"><code>50fe1cc</code></a> test: Fix flaky test <code>TestServerStreaming_ClientCallRecvMsgTwice</code> in `end2end...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/d574bad188f25ba03d41a506e6f2ef93837ad10b"><code>d574bad</code></a> build(deps): bump go.opentelemetry.io/otel/sdk from 1.42.0 to 1.43.0 (<a href="https://redirect.github.com/grpc/grpc-go/issues/9050">#9050</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/b8bf4d0488a351c563d63797ffba321585d6bb24"><code>b8bf4d0</code></a> build(deps): bump go.opentelemetry.io/otel/sdk from 1.42.0 to 1.43.0 in /inte...</li>
<li>Additional commits viewable in <a href="https://github.com/grpc/grpc-go/compare/v1.80.0...v1.81.0">compare view</a></li>
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

### Review by @swadeley - Approved on May 11, 2026 at 07:35 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1493*
