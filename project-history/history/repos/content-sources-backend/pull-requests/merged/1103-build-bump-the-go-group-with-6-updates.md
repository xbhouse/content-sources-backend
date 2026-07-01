---
type: pull_request
number: 1103
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2025-05-05T04:36:14Z
updated: 2025-05-06T10:35:26Z
closed: 2025-05-06T10:35:18Z
merged: 2025-05-06T10:35:18Z
base_branch: main
head_branch: dependabot/go_modules/go-ac897e1bca
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1103
---

# Pull Request #1103: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: May 05, 2025 at 04:36 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-ac897e1bca`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.131.0` | `0.132.0` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.47.3` | `1.48.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.79.2` | `1.79.3` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.4.1745594022` | `2025.5.1746105577` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.7.3` | `9.8.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.12` | `1.25.13` |

Updates `github.com/getkin/kin-openapi` from 0.131.0 to 0.132.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.132.0</h2>
<h2>What's Changed</h2>
<ul>
<li>style: Use fmt.Sprint without formating by <a href="https://github.com/gaiaz-iusipov"><code>@​gaiaz-iusipov</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1061">getkin/kin-openapi#1061</a></li>
<li>openapi3filter: don't consume request body in <code>AuthenticatorFunc</code> by <a href="https://github.com/jamietanna"><code>@​jamietanna</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1064">getkin/kin-openapi#1064</a></li>
<li>openapi2conv: fix for refs on items within additional properties by <a href="https://github.com/roberth1988"><code>@​roberth1988</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1067">getkin/kin-openapi#1067</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/gaiaz-iusipov"><code>@​gaiaz-iusipov</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1061">getkin/kin-openapi#1061</a></li>
<li><a href="https://github.com/roberth1988"><code>@​roberth1988</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1067">getkin/kin-openapi#1067</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.131.0...v0.132.0">https://github.com/getkin/kin-openapi/compare/v0.131.0...v0.132.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/72b5d5177d5993bbbc8c863e36ef036850558658"><code>72b5d51</code></a> openapi2conv: fix for refs on items within additional properties (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1067">#1067</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/75e2cc5ad8729cf119ad0274ea0678e2e360c731"><code>75e2cc5</code></a> openapi3filter: don't consume request body in <code>AuthenticatorFunc</code> (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1064">#1064</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/3cf8f7dad86555f6ad46c5972cc1bfb517afb148"><code>3cf8f7d</code></a> style: Use fmt.Sprint without formating (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1061">#1061</a>)</li>
<li>See full diff in <a href="https://github.com/getkin/kin-openapi/compare/v0.131.0...v0.132.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.47.3 to 1.48.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ba37053fff9055b216a303ab591f6fa5e4c80c1"><code>4ba3705</code></a> Release 2024-01-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c49b1531db5ba4a8e2a42d5fd31c8b10641adca6"><code>c49b153</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/80508f5111049f7f13295423d8fa8dcb3d26bd90"><code>80508f5</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/686a64c6d4ae0bb6e8139e4761034a0cc25caf47"><code>686a64c</code></a> codegen: track upstream sigv4a trait changes (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2442">#2442</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/47dd1b1bcbde244357a82ef00fa6a61a9bfb9b39"><code>47dd1b1</code></a> Release 2024-01-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23145e3e605a93582020facfe13350b4153714e1"><code>23145e3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00e57bb7feb2d104387103aa4fbcd55dfff3a6a7"><code>00e57bb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/73e1a99f2fa858ca56627779852768a9198ba057"><code>73e1a99</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2441">#2441</a> from RanVaknin/fix-documentation-config</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0f8ad11593c219c52ad7fb12998c75ade39dc7ad"><code>0f8ad11</code></a> Fix SRA auth trailing checksum retry bug (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2438">#2438</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/efbc5aa622a882167129e69a88aa50c730cd1904"><code>efbc5aa</code></a> Release 2024-01-03</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.47.3...service/s3/v1.48.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.2 to 1.79.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cfef7b8bb84e6f11f5294901ebbb45eccc702b96"><code>cfef7b8</code></a> Release 2025-04-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6de0aa4647f8e70309b94987a6a7d9dbe41a8be7"><code>6de0aa4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c2224b26d5dcf3cf6117b90649bd17f5d7561e3c"><code>c2224b2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6f897c6ad66a638bf769a2791e19699b08e7fbc3"><code>6f897c6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/977ee1d009c7bffd8b456d12c276b859caf7f4da"><code>977ee1d</code></a> don't perform checksum validation in non-200 responses (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3075">#3075</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b06fd664465de19f26e46ada588801c1c3013f78"><code>b06fd66</code></a> Release 2025-04-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/989866c3da168aeb7f89d2f8ec6804c573c5789a"><code>989866c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a3a3dfc32c5aa5df53568222252db40d894263c"><code>5a3a3df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa132cce708ed0b059b7cbd23962ef755b8b080d"><code>aa132cc</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3aace4178c90b4b6acd3e6485e04e752c7c30a85"><code>3aace41</code></a> Transfer Manager v2 downloader concurrency fix and version control (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3058">#3058</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.2...service/s3/v1.79.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.4.1745594022 to 2025.5.1746105577
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/2d53147dd6d9c82f1abbc899a704e207dd525867"><code>2d53147</code></a> Update pulp bindings to 386d86254354e29d3b864523aed47884be82d9067968b5e2da9ab...</li>
<li><a href="https://github.com/content-services/zest/commit/ea97913dd5c6be997f584546d5fe87c234022042"><code>ea97913</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47569b455431b7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/a5374381f4930c4ee799f66569bd02e00bc4e28d"><code>a537438</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276e2d48333127b8e3dae95843...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.4.1745594022...release/v2025.5.1746105577">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.7.3 to 9.8.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>v9.8.0</h2>
<h1>9.8.0 (2025-04-30)</h1>
<h2>🚀 Highlights</h2>
<ul>
<li><strong>Redis 8 Support</strong>: Full compatibility with Redis 8.0, including testing and CI integration</li>
<li><strong>Enhanced Hash Operations</strong>: Added support for new hash commands (<code>HGETDEL</code>, <code>HGETEX</code>, <code>HSETEX</code>) and <code>HSTRLEN</code> command</li>
<li><strong>Search Improvements</strong>: Enabled Search DIALECT 2 by default and added <code>CountOnly</code> argument for <code>FT.Search</code></li>
</ul>
<h2>✨ New Features</h2>
<ul>
<li>Added support for new hash commands: <code>HGETDEL</code>, <code>HGETEX</code>, <code>HSETEX</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3305">#3305</a>)</li>
<li>Added <code>HSTRLEN</code> command for hash operations (<a href="https://redirect.github.com/redis/go-redis/pull/2843">#2843</a>)</li>
<li>Added <code>Do</code> method for raw query by single connection from <code>pool.Conn()</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3182">#3182</a>)</li>
<li>Prevent false-positive marshaling by treating zero time.Time as empty in isEmptyValue (<a href="https://redirect.github.com/redis/go-redis/pull/3273">#3273</a>)</li>
<li>Added FailoverClusterClient support for Universal client (<a href="https://redirect.github.com/redis/go-redis/pull/2794">#2794</a>)</li>
<li>Added support for cluster mode with <code>IsClusterMode</code> config parameter (<a href="https://redirect.github.com/redis/go-redis/pull/3255">#3255</a>)</li>
<li>Added client name support in <code>HELLO</code> RESP handshake (<a href="https://redirect.github.com/redis/go-redis/pull/3294">#3294</a>)</li>
<li><strong>Enabled Search DIALECT 2 by default</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3213">#3213</a>)</li>
<li>Added read-only option for failover configurations (<a href="https://redirect.github.com/redis/go-redis/pull/3281">#3281</a>)</li>
<li>Added <code>CountOnly</code> argument for <code>FT.Search</code> to use <code>LIMIT 0 0</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3338">#3338</a>)</li>
<li>Added <code>DB</code> option support in <code>NewFailoverClusterClient</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3342">#3342</a>)</li>
<li>Added <code>nil</code> check for the options when creating a client (<a href="https://redirect.github.com/redis/go-redis/pull/3363">#3363</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>Fixed <code>PubSub</code> concurrency safety issues (<a href="https://redirect.github.com/redis/go-redis/pull/3360">#3360</a>)</li>
<li>Fixed panic caused when argument is <code>nil</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3353">#3353</a>)</li>
<li>Improved error handling when fetching master node from sentinels (<a href="https://redirect.github.com/redis/go-redis/pull/3349">#3349</a>)</li>
<li>Fixed connection pool timeout issues and increased retries (<a href="https://redirect.github.com/redis/go-redis/pull/3298">#3298</a>)</li>
<li>Fixed context cancellation error leading to connection spikes on Primary instances (<a href="https://redirect.github.com/redis/go-redis/pull/3190">#3190</a>)</li>
<li>Fixed RedisCluster client to consider <code>MASTERDOWN</code> a retriable error (<a href="https://redirect.github.com/redis/go-redis/pull/3164">#3164</a>)</li>
<li>Fixed tracing to show complete commands instead of truncated versions (<a href="https://redirect.github.com/redis/go-redis/pull/3290">#3290</a>)</li>
<li>Fixed OpenTelemetry instrumentation to prevent multiple span reporting (<a href="https://redirect.github.com/redis/go-redis/pull/3168">#3168</a>)</li>
<li>Fixed <code>FT.Search</code> Limit argument and added <code>CountOnly</code> argument for limit 0 0 (<a href="https://redirect.github.com/redis/go-redis/pull/3338">#3338</a>)</li>
<li>Fixed missing command in interface (<a href="https://redirect.github.com/redis/go-redis/pull/3344">#3344</a>)</li>
<li>Fixed slot calculation for <code>COUNTKEYSINSLOT</code> command (<a href="https://redirect.github.com/redis/go-redis/pull/3327">#3327</a>)</li>
<li>Updated PubSub implementation with correct context (<a href="https://redirect.github.com/redis/go-redis/pull/3329">#3329</a>)</li>
</ul>
<h2>📚 Documentation</h2>
<ul>
<li>Added hash search examples (<a href="https://redirect.github.com/redis/go-redis/pull/3357">#3357</a>)</li>
<li>Fixed documentation comments (<a href="https://redirect.github.com/redis/go-redis/pull/3351">#3351</a>)</li>
<li>Added <code>CountOnly</code> search example (<a href="https://redirect.github.com/redis/go-redis/pull/3345">#3345</a>)</li>
<li>Added examples for list commands: <code>LLEN</code>, <code>LPOP</code>, <code>LPUSH</code>, <code>LRANGE</code>, <code>RPOP</code>, <code>RPUSH</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3234">#3234</a>)</li>
<li>Added <code>SADD</code> and <code>SMEMBERS</code> command examples (<a href="https://redirect.github.com/redis/go-redis/pull/3242">#3242</a>)</li>
<li>Updated <code>README.md</code> to use Redis Discord guild (<a href="https://redirect.github.com/redis/go-redis/pull/3331">#3331</a>)</li>
<li>Updated <code>HExpire</code> command documentation (<a href="https://redirect.github.com/redis/go-redis/pull/3355">#3355</a>)</li>
<li>Featured OpenTelemetry instrumentation more prominently (<a href="https://redirect.github.com/redis/go-redis/pull/3316">#3316</a>)</li>
<li>Updated <code>README.md</code> with additional information (<a href="https://github.com/redis/go-redis/commit/310ce55">#310ce55</a>)</li>
</ul>
<h2>⚡ Performance and Reliability</h2>
<ul>
<li>Bound connection pool background dials to configured dial timeout (<a href="https://redirect.github.com/redis/go-redis/pull/3089">#3089</a>)</li>
<li>Ensured context isn't exhausted via concurrent query (<a href="https://redirect.github.com/redis/go-redis/pull/3334">#3334</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/fba6dececdf138af0b34928e28d514f9205ed7d2"><code>fba6dec</code></a> Merge branch 'master' into v9.8</li>
<li><a href="https://github.com/redis/go-redis/commit/46ede21fdbe2ab00aff180ab14f2a490996c81ba"><code>46ede21</code></a> chore(release): Update version to v9.8.0</li>
<li><a href="https://github.com/redis/go-redis/commit/22992116f37f38c0b34ad2a6a5f19ff244c3f8ad"><code>2299211</code></a> feat(options): panic when options are nil (<a href="https://redirect.github.com/redis/go-redis/issues/3363">#3363</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8f582356b31a8bdc9bfad7327abb4e279eddbdb5"><code>8f58235</code></a> chore(ci): Use redis 8 rc2 image. (<a href="https://redirect.github.com/redis/go-redis/issues/3361">#3361</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/09dc3510a2c7c38a6be81f8fcf22ed8ed7410e40"><code>09dc351</code></a> migrate golangci-lint config to v2 format (<a href="https://redirect.github.com/redis/go-redis/issues/3354">#3354</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/cb2cfb000d5321b7f49204175dfa24f2ab69d18e"><code>cb2cfb0</code></a> fix: <code>PubSub</code> isn't concurrency-safe (<a href="https://redirect.github.com/redis/go-redis/issues/3360">#3360</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/46d4b20ee247198ad0386e93286773c755e3632a"><code>46d4b20</code></a> feat: func isEmptyValue support time.Time (<a href="https://redirect.github.com/redis/go-redis/issues/3273">#3273</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f9b0e70f4c0e34210d8429592fc165922d91e706"><code>f9b0e70</code></a> update HExpire command documentation (<a href="https://redirect.github.com/redis/go-redis/issues/3355">#3355</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b28606cf2599bdebce2b2f534445e605f4d76e1e"><code>b28606c</code></a> Update README.md, use redis discord guild (<a href="https://redirect.github.com/redis/go-redis/issues/3331">#3331</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/adb479820c600565622a8c7d739d7b38ce86e5b1"><code>adb4798</code></a> fix: Fix panic caused when arg is nil (<a href="https://redirect.github.com/redis/go-redis/issues/3353">#3353</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.7.3...v9.8.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.12 to 1.25.13
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/0b34ee10c9f25f4a2b3525c5a94522094793ab07"><code>0b34ee1</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/611">#611</a> from JiriPapousek/CCXDEV-14922</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/3cb8d01ad5e12d3b3f354fedabdca7016f19894d"><code>3cb8d01</code></a> Replace ubuntu-20.04 with latest in all GH runners</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/3e64016e04413a13405e0d93ed75c6fd93ce80ab"><code>3e64016</code></a> Add username arg to createRedisClient function</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/39ebf74bb33b44d722d217cf6f92621189608934"><code>39ebf74</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/607">#607</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/0cf06181e30f0cd7739822f356e13d9b4381d1a3"><code>0cf0618</code></a> Bump github.com/prometheus/client_model from 0.6.1 to 0.6.2</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/c5a3df211d9638ab370a6b35154d31e7d29a1866"><code>c5a3df2</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/606">#606</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/682f661f3efe2ab988de4969ed54a5fff9357c2a"><code>682f661</code></a> Bump github.com/getsentry/sentry-go from 0.31.1 to 0.32.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/37c827a3ef49b4697cdf01c391d51f1b635fdaff"><code>37c827a</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/605">#605</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/23efd7b32af049d66386e7b9bd0563fa583bd699"><code>23efd7b</code></a> Bump github.com/prometheus/client_golang from 1.21.1 to 1.22.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/5c5235ee1f73e9da883386cc317d5311a648c7cb"><code>5c5235e</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/604">#604</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.12...v1.25.13">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


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

## Reviews

### Review by @swadeley - Approved on May 06, 2025 at 10:35 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1103*
