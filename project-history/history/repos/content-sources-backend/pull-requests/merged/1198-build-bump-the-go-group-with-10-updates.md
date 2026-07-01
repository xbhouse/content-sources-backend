---
type: pull_request
number: 1198
title: "Build: Bump the go group with 10 updates"
state: merged
author: dependabot
created: 2025-09-15T04:10:51Z
updated: 2025-09-15T08:48:08Z
closed: 2025-09-15T08:48:00Z
merged: 2025-09-15T08:48:00Z
base_branch: main
head_branch: dependabot/go_modules/go-9783b5a6dc
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1198
---

# Pull Request #1198: Build: Bump the go group with 10 updates

**Author**: @dependabot
**Created**: September 15, 2025 at 04:10 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9783b5a6dc`

## Description

Bumps the go group with 10 updates:

| Package | From | To |
| --- | --- | --- |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.30.5` | `1.31.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.38.3` | `1.39.0` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.6` | `1.31.8` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.10` | `1.18.12` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.57.2` | `1.57.4` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.87.3` | `1.88.1` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.35.1` | `0.35.2` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.1.0` | `1.1.1` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.13.0` | `9.14.0` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.75.0` | `1.75.1` |

Updates `gorm.io/gorm` from 1.30.5 to 1.31.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.31.0</h2>
<h2>Changes</h2>
<ul>
<li>Add association operation support to generics Set API and enable conditional bulk association updates <a href="https://github.com/jinzhu"><code>@​jinzhu</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7581">#7581</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/b289563d54de23e742e2a4cba9d77af01f641b0f"><code>b289563</code></a> Remove unnecessary OpCreateValue mode</li>
<li><a href="https://github.com/go-gorm/gorm/commit/e015bd08ebb5cc27aabf1731735324138879a44b"><code>e015bd0</code></a> Add association operation support to generics Set API and enable conditional ...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/faee391f05f3426660598d22f03881f115d3aa39"><code>faee391</code></a> fix: add missing Count method to CreateInterface (closes <a href="https://redirect.github.com/go-gorm/gorm/issues/7580">#7580</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.30.5...v1.31.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.38.3 to 1.39.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/648027edb8aeba036195538174a63cbccaca8c16"><code>648027e</code></a> Release 2025-09-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a3b9b7b1303ebf9fc1837f70253fe086d0ee72a2"><code>a3b9b7b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67dad834dcc6ab1c14a2b9935bd92bc1ff91ddd9"><code>67dad83</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2bfe86a2ca763380b5fb498feab0886d41b96dcc"><code>2bfe86a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59e7410f279040b4c186e149402eebf13c5e06d9"><code>59e7410</code></a> add businessmetrics feature ID for env-based bearer token (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3182">#3182</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1cdc15880e2f5087470b6f813eb0a5cd1451c95f"><code>1cdc158</code></a> Patching override s3expire shape (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3180">#3180</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1745ede8f69372b0e08c4a8bcc83d43e76caeb26"><code>1745ede</code></a> Release 2025-09-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f84de5375a6e97e530819caba27fa0c7bc773cd6"><code>f84de53</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7cdaa318493e10d7ef9dd87eb1130b9806c48c98"><code>7cdaa31</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/498b5c42d5b3fc143fc3449665faa3635f468d09"><code>498b5c4</code></a> remove service/sms (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3177">#3177</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.38.3...v1.39.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.6 to 1.31.8
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70e9f3d91a369c9147a9b09f2e66e270ef5febbb"><code>70e9f3d</code></a> Release 2025-09-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e620ecbfcfd29508d5de745cb341201563159188"><code>e620ecb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b412203666022514559534c7efd5a010521fcf51"><code>b412203</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/89f42ef7608024a6714dd3dbd290ab37302f0bec"><code>89f42ef</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d71b109a7510e2e82f5f4afaa4b272cbdcb2de42"><code>d71b109</code></a> Bump smithy go version to allow unused required parameter in endpoint rule se...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f68827f17283ffb439c64aa951a6dd4852bef8e2"><code>f68827f</code></a> Release 2025-09-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1748bf98381993674dd5be1f6c716cd3e9df80b"><code>d1748bf</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00307c01f6a8843dc17230cd5ea3043e976260f7"><code>00307c0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/424be9309a8c1473c78a45da53a5a6c0f4907362"><code>424be93</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/648027edb8aeba036195538174a63cbccaca8c16"><code>648027e</code></a> Release 2025-09-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.31.6...config/v1.31.8">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.10 to 1.18.12
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3f28b5134e7e9a047147b5773af62c6012c34df6"><code>3f28b51</code></a> Release 2023-02-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e8d17fd21083f2ec6ec6139f89b35c841994932"><code>6e8d17f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/60dbdbb0da35b8b8374e0b4e1a9536f1aed7e458"><code>60dbdbb</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/212910ac25d59e959c0a534bd6264a13fb9ca8c8"><code>212910a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb8cb66b4433cd5f125abc7c6b74de39c19f25fb"><code>eb8cb66</code></a> Upgrade smithy to 1.27.2, correct query empty list serialization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/24db9f5f6e6d387f115ee8d9393ae6d158e8ef0c"><code>24db9f5</code></a> Update processcreds.CredentialProcessResponse visibility to public (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1921">#1921</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bd3003e29f1fdc34859fcd194a61d6655b8ea492"><code>bd3003e</code></a> dependency: upgrade smithy to 1.27.2 and correct query empty list serialization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d94f223e8c3dc3c63d001ef443a25c056d4131e"><code>0d94f22</code></a> Release 2023-02-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2eec85ed13c74cf77315398edf974481fb5f4dd8"><code>2eec85e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ca6e32eedbb1e707ad3675879e17782e93db67e"><code>4ca6e32</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.10...config/v1.18.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.57.2 to 1.57.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/983f19260821af7fce8a653741e61f89dd851c68"><code>983f192</code></a> Release 2025-06-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5c1277d48ccc889ad50f25725530106ecde1699"><code>a5c1277</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a42991177cb12bd8edd1e05a06f4a069b73c4bfb"><code>a429911</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ea1cecfb150d10d7f314dddd54bc838fdcba26e"><code>4ea1cec</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5b11c8d01f258aa7641c77c5dd0f97a2ad3d73c9"><code>5b11c8d</code></a> remove changelog directions for now because of <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3107">#3107</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79f492ceb29aeb14397d3551e310dd2d48f82057"><code>79f492c</code></a> fixup changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4f82369defdab68029247cd3eb74dda480efe412"><code>4f82369</code></a> use UTC() in v4 event stream signing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3105">#3105</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/755839b2eebb246c7eec79b65404aee105196d5b"><code>755839b</code></a> Release 2025-06-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba3d22d775f3c1ef15df4da2d0eb4aa440d99798"><code>ba3d22d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/01587c6c41afce3f3dc36f044ded13e5fcdc6746"><code>01587c6</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.57.2...service/ecs/v1.57.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.87.3 to 1.88.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70e9f3d91a369c9147a9b09f2e66e270ef5febbb"><code>70e9f3d</code></a> Release 2025-09-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e620ecbfcfd29508d5de745cb341201563159188"><code>e620ecb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b412203666022514559534c7efd5a010521fcf51"><code>b412203</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/89f42ef7608024a6714dd3dbd290ab37302f0bec"><code>89f42ef</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d71b109a7510e2e82f5f4afaa4b272cbdcb2de42"><code>d71b109</code></a> Bump smithy go version to allow unused required parameter in endpoint rule se...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f68827f17283ffb439c64aa951a6dd4852bef8e2"><code>f68827f</code></a> Release 2025-09-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1748bf98381993674dd5be1f6c716cd3e9df80b"><code>d1748bf</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00307c01f6a8843dc17230cd5ea3043e976260f7"><code>00307c0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/424be9309a8c1473c78a45da53a5a6c0f4907362"><code>424be93</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/648027edb8aeba036195538174a63cbccaca8c16"><code>648027e</code></a> Release 2025-09-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.87.3...service/s3/v1.88.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.35.1 to 0.35.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.35.2</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.2.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Fix OpenTelemetry spans being created as transactions instead of child spans (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1073">#1073</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Add <code>MockTransport</code> to test clients for improved testing (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1071">#1071</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.35.2</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.2.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Fix OpenTelemetry spans being created as transactions instead of child spans (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1073">#1073</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Add <code>MockTransport</code> to test clients for improved testing (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1071">#1071</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/febe2bad9219772857f8d699b40d77788223dcf6"><code>febe2ba</code></a> release: 0.35.2</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d0c9af5317516624657728e36f868652d71fe9cb"><code>d0c9af5</code></a> Prepare 0.35.2 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1078">#1078</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e9566db189127d363c518331a7579b4aea62cd29"><code>e9566db</code></a> Fix otel span being created as a transaction (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1073">#1073</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/6013529525adfd565819acf46efe753eb6b6ea51"><code>6013529</code></a> Run test pipeline once (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1072">#1072</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e6bbe0fdb4d2f01214b2e4acefea1df22fae1103"><code>e6bbe0f</code></a> Add mockTransport on test clients (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1071">#1071</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/6fa7ff43d5c31247e69fe93d48f86b8ebc9b9296"><code>6fa7ff4</code></a> Merge branch 'release/0.35.1'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.35.1...v0.35.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.1.0 to 1.1.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.1.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Add go client release instructions by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/9">project-kessel/kessel-sdk-go#9</a></li>
<li>chore: regenerate protos by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/10">project-kessel/kessel-sdk-go#10</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.1.0...v1.1.1">https://github.com/project-kessel/kessel-sdk-go/compare/v1.1.0...v1.1.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/b3508074cf83a1bcc083f57167ca845e6b262276"><code>b350807</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/10">#10</a> from lennysgarage/regenerate-protos</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/bd961635e787f1bbd5b2890f2e3d6b3e8dd6adf0"><code>bd96163</code></a> chore: regenerate protos</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/2d81d723d53625a4705bc2f91cb5c535a17fe9bf"><code>2d81d72</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/9">#9</a> from lennysgarage/release-instructions</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/1466f2a66de9b415010b8af1ce95d036489b7886"><code>1466f2a</code></a> go not python</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/abba8b01d4f901496adf4e438e8f07b2e033a4b4"><code>abba8b0</code></a> Add go client release instructions</li>
<li>See full diff in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.1.0...v1.1.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.13.0 to 9.14.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.14.0</h2>
<h2>Highlights</h2>
<ul>
<li>Added batch process method to the pipeline (<a href="https://redirect.github.com/redis/go-redis/pull/3510">#3510</a>)</li>
</ul>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Added batch process method to the pipeline (<a href="https://redirect.github.com/redis/go-redis/pull/3510">#3510</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix: SetErr on Cmd if the command cannot be queued correctly in multi/exec (<a href="https://redirect.github.com/redis/go-redis/pull/3509">#3509</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>Updates release drafter config to exclude dependabot (<a href="https://redirect.github.com/redis/go-redis/pull/3511">#3511</a>)</li>
<li>chore(deps): bump actions/setup-go from 5 to 6 (<a href="https://redirect.github.com/redis/go-redis/pull/3504">#3504</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/elena-kolevksa"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a> and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.14.0 (2025-09-10)</h1>
<h2>Highlights</h2>
<ul>
<li>Added batch process method to the pipeline (<a href="https://redirect.github.com/redis/go-redis/pull/3510">#3510</a>)</li>
</ul>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Added batch process method to the pipeline (<a href="https://redirect.github.com/redis/go-redis/pull/3510">#3510</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix: SetErr on Cmd if the command cannot be queued correctly in multi/exec (<a href="https://redirect.github.com/redis/go-redis/pull/3509">#3509</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>Updates release drafter config to exclude dependabot (<a href="https://redirect.github.com/redis/go-redis/pull/3511">#3511</a>)</li>
<li>chore(deps): bump actions/setup-go from 5 to 6 (<a href="https://redirect.github.com/redis/go-redis/pull/3504">#3504</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/elena-kolevksa"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a> and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/2da6ca07c065db5f24bf47cbf70510c80e3190ba"><code>2da6ca0</code></a> chore(release): Update the rest of the versions (<a href="https://redirect.github.com/redis/go-redis/issues/3513">#3513</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c11a70448132e808ea8e6f33775ace839859dc0d"><code>c11a704</code></a> chore(release): v9.14.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3512">#3512</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8f5469abd04faaaceb1d211dbe15ca3538ed1e5f"><code>8f5469a</code></a> chore(ci): Update release drafter config to exclude dependabot (<a href="https://redirect.github.com/redis/go-redis/issues/3511">#3511</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a264ffb8a4a923043364329cbfdbf7577a64c293"><code>a264ffb</code></a> fix: SetErr on Cmd if the command cannot be queued correctly in multi/exec (#...</li>
<li><a href="https://github.com/redis/go-redis/commit/e0853aba634dd9fb50a55919c2442ffe7d382013"><code>e0853ab</code></a> Added batch process method to the pipeline (<a href="https://redirect.github.com/redis/go-redis/issues/3510">#3510</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/65e1c22065050e7390350482f41728f470fe7994"><code>65e1c22</code></a> chore(deps): bump actions/setup-go from 5 to 6 (<a href="https://redirect.github.com/redis/go-redis/issues/3504">#3504</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.13.0...v9.14.0">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.75.0 to 1.75.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.75.1</h2>
<h1>Bug Fixes</h1>
<ul>
<li>transport: Fix a data race while copying headers for stats handlers in the std lib http2 server transport. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8519">#8519</a>)</li>
<li>xdsclient:
<ul>
<li>Fix a data race caused while reporting load to LRS. (<a href="https://redirect.github.com/grpc/grpc-go/pull/8483">#8483</a>)</li>
<li>Fix regression preventing empty node IDs when creating an LRS client. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8483">#8483</a>)</li>
</ul>
</li>
<li>server: Fix a regression preventing streams from being cancelled or timed out when blocked on flow control. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8528">#8528</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/b4dc263cb692d1951a1842cc877d913d30de0559"><code>b4dc263</code></a> Change version to 1.75.1 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8559">#8559</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/1fffee7cbccdfcf4a973ea3ce4d7e7dd2aa63cfd"><code>1fffee7</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8528">#8528</a> to v1.75.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/8555">#8555</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/a52e42be78d406f40dd5888e999c229fb2d63044"><code>a52e42b</code></a> Cherry pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8483">#8483</a> into v1.75.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/8541">#8541</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/369c9aa6cedbd23e46df8d6ba851e6a2fc3e459a"><code>369c9aa</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8519">#8519</a> to v1.75.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/8530">#8530</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/7269d5fe7026b29be98c5db85e978916b37496dd"><code>7269d5f</code></a> Change version to 1.75.1-dev (<a href="https://redirect.github.com/grpc/grpc-go/issues/8494">#8494</a>)</li>
<li>See full diff in <a href="https://github.com/grpc/grpc-go/compare/v1.75.0...v1.75.1">compare view</a></li>
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

### Comment by @swadeley on September 15, 2025 at 08:44 AM UTC

Only `test_bulk_import` failed

---

## Reviews

### Review by @swadeley - Approved on September 15, 2025 at 08:47 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1198*
