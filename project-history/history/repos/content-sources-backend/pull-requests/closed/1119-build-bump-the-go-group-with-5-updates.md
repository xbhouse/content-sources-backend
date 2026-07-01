---
type: pull_request
number: 1119
title: "Build: Bump the go group with 5 updates"
state: closed
author: dependabot
created: 2025-06-02T04:51:36Z
updated: 2025-06-03T15:09:58Z
closed: 2025-06-03T15:09:57Z
base_branch: main
head_branch: dependabot/go_modules/go-37ac329d77
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1119
---

# Pull Request #1119: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: June 02, 2025 at 04:51 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-37ac329d77`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.11` | `1.6.0` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.45.1` | `1.45.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.79.4` | `1.80.0` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.5.1747750822` | `2025.5.1748552548` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.8.0` | `9.9.0` |

Updates `gorm.io/driver/postgres` from 1.5.11 to 1.6.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/postgres/commit/1ffb5a75de20776ce5d266d62436e23ec0cfe60e"><code>1ffb5a7</code></a> feat: support dynamic timestamp ScanLocation via DSN timezone (<a href="https://redirect.github.com/go-gorm/postgres/issues/311">#311</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/8385c4afd131d516784f1154d9005601f145588a"><code>8385c4a</code></a> feat(migrator): Add missing postgres type aliases (<a href="https://redirect.github.com/go-gorm/postgres/issues/312">#312</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/84224ecab0a0e50ec9673bacae45e42122ff4b7b"><code>84224ec</code></a> fix: fix migrator recognizing <code>nextval('&quot;table_id_seq&quot;'::regclass)</code> (<a href="https://redirect.github.com/go-gorm/postgres/issues/304">#304</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/04418c920c1f6de324ed333a5f33bf6b894b80bf"><code>04418c9</code></a> PG index options handle CONCURRENT option separately (<a href="https://redirect.github.com/go-gorm/postgres/issues/297">#297</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/9723eacc5e6f3f3e3043925add4f1a51019ae7bc"><code>9723eac</code></a> Add varchar limit detection as per PostgreSQL limits (<a href="https://redirect.github.com/go-gorm/postgres/issues/302">#302</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/dcfe4ea8646a372ad171028441a3100f597feab8"><code>dcfe4ea</code></a> Bump golang.org/x/crypto to v0.31.0 to address CVE-2024-45337 (<a href="https://redirect.github.com/go-gorm/postgres/issues/301">#301</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/c1721250b2600eaf175be7f56f58567745b1ee84"><code>c172125</code></a> fix concurrent issue (<a href="https://redirect.github.com/go-gorm/postgres/issues/295">#295</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/postgres/compare/v1.5.11...v1.6.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.45.1 to 1.45.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.45.2 (2025-05-28)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(decoder): use configurable limit for max number of records in a record batch by <a href="https://github.com/rmb938"><code>@​rmb938</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3120">IBM/sarama#3120</a></li>
<li>fix: ensure mock SyncProducer's SendMessage returns msg.Partition instead of 0 by <a href="https://github.com/magiusdarrigo"><code>@​magiusdarrigo</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3122">IBM/sarama#3122</a></li>
<li>fix: send null instead of empty string when describing default client quotas by <a href="https://github.com/petedannemann"><code>@​petedannemann</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3128">IBM/sarama#3128</a></li>
<li>fix: improve getMetricName performance by <a href="https://github.com/boekkooi-impossiblecloud"><code>@​boekkooi-impossiblecloud</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3156">IBM/sarama#3156</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump github.com/klauspost/compress from 1.17.11 to 1.18.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3103">IBM/sarama#3103</a></li>
<li>chore(deps): bump the golang-x group across 6 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3114">IBM/sarama#3114</a></li>
<li>chore(deps): bump the golang-x group across 7 directories with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3121">IBM/sarama#3121</a></li>
<li>chore(deps): bump the go_modules group across 7 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3148">IBM/sarama#3148</a></li>
<li>chore(deps): bump the go_modules group across 7 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3157">IBM/sarama#3157</a></li>
<li>chore(deps): bump golang.org/x/sync from 0.12.0 to 0.14.0 in the golang-x group across 1 directory by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3161">IBM/sarama#3161</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>chore: bump minimum Go version to 1.23.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3113">IBM/sarama#3113</a></li>
<li>fix(ci): bump golangci-lint to v2 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3160">IBM/sarama#3160</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/rmb938"><code>@​rmb938</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3120">IBM/sarama#3120</a></li>
<li><a href="https://github.com/magiusdarrigo"><code>@​magiusdarrigo</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3122">IBM/sarama#3122</a></li>
<li><a href="https://github.com/petedannemann"><code>@​petedannemann</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3128">IBM/sarama#3128</a></li>
<li><a href="https://github.com/renovate"><code>@​renovate</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3155">IBM/sarama#3155</a></li>
<li><a href="https://github.com/boekkooi-impossiblecloud"><code>@​boekkooi-impossiblecloud</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3156">IBM/sarama#3156</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.45.1...v1.45.2">https://github.com/IBM/sarama/compare/v1.45.1...v1.45.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/d4acbec36f8e76dc731d12b108a84ed0cd3b1e3b"><code>d4acbec</code></a> chore(deps): update ghcr.io/shopify/toxiproxy docker tag to v2.12.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3171">#3171</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/876967c38f756f36211a359e4f7add3626ddee36"><code>876967c</code></a> fix: improve getMetricName performance (<a href="https://redirect.github.com/IBM/sarama/issues/3156">#3156</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/c5f1af503ee85a72a84ac55319fa07a93e79ab35"><code>c5f1af5</code></a> chore(ci): bump github/codeql-action from 3.28.11 to 3.28.18 (<a href="https://redirect.github.com/IBM/sarama/issues/3167">#3167</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/29f5988dcb552956437c52caef1c9880e3f6e85e"><code>29f5988</code></a> chore: Configure Renovate (<a href="https://redirect.github.com/IBM/sarama/issues/3155">#3155</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/eb2bba0204450e638c211091ba503a950c76666b"><code>eb2bba0</code></a> chore(deps): bump golang.org/x/sync from 0.12.0 to 0.14.0 in the golang-x gro...</li>
<li><a href="https://github.com/IBM/sarama/commit/67d982f98d2989a410c42b61d4ff5f4637176f3a"><code>67d982f</code></a> chore(ci): bump the actions group with 2 updates (<a href="https://redirect.github.com/IBM/sarama/issues/3162">#3162</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/29febcd188e75ff4d97c56cb6338fe5616e1ff03"><code>29febcd</code></a> chore(ci): bump docker/bake-action from 6.4.0 to 6.7.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3163">#3163</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/de58f97fd9fcec48b33d862843857439bddf67fd"><code>de58f97</code></a> chore(deps): bump the go_modules group across 7 directories with 1 update (<a href="https://redirect.github.com/IBM/sarama/issues/3">#3</a>...</li>
<li><a href="https://github.com/IBM/sarama/commit/b10aa9fc0c6a54d738d0e7a9ae2a687b32e061d8"><code>b10aa9f</code></a> chore(ci): bump ubi9/ubi-minimal from 9.5 to 9.6 (<a href="https://redirect.github.com/IBM/sarama/issues/3159">#3159</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4d2b166e5c2770f6b23f83eab587d7801878d429"><code>4d2b166</code></a> fix(ci): bump golangci-lint to v2 (<a href="https://redirect.github.com/IBM/sarama/issues/3160">#3160</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.45.1...v1.45.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.4 to 1.80.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c69bd7e10c45c4625623d29f4c7b3889188029e1"><code>c69bd7e</code></a> Release 2025-05-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21fdc96693cdd766dd5ade889d6c7fcfc49997fe"><code>21fdc96</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a4ab8271698624cd04a8b46b4cc2ff8753cf7eb"><code>2a4ab82</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f44412972b960fc96e05eab8e231cfa837991dbe"><code>f444129</code></a> Release 2025-05-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ba713d71fae0f9b80d5d7aebf7b4ad11cdc95b9"><code>4ba713d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98462780675eb7b6cfe4f64eda8912a547dc5165"><code>9846278</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f1b00177a5db154161c2a8f675170b6d7000e1a1"><code>f1b0017</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223d488213a4836f2f13ad84105bc4157e6667f2"><code>223d488</code></a> Remove smoke tests on sunset service (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3098">#3098</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1cfa85a959253a88f176cb2737f6da4f0aed1d5f"><code>1cfa85a</code></a> Release 2025-05-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d078e8ae27f14ef703b4d55c202a7dd0a4eacad3"><code>d078e8a</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.4...service/s3/v1.80.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.5.1747750822 to 2025.5.1748552548
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/9f20ab31c47cb76a5ff8d2a0d2968d4c0917de01"><code>9f20ab3</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7ddd45b9b9f57952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/39323bb33a0ebc01b3bf1d87dfe6b0a58456d760"><code>39323bb</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e8828e449087e8639db952b3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.5.1747750822...release/v2025.5.1748552548">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.8.0 to 9.9.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.9.0</h2>
<h2>🚀 Highlights</h2>
<ul>
<li><strong>Token-based Authentication</strong>: Added <code>StreamingCredentialsProvider</code> for dynamic credential updates (<strong>experimental</strong>)
<ul>
<li>Can be used with <a href="https://github.com/redis/go-redis-entraid">go-redis-entraid</a> for Azure AD authentication</li>
</ul>
</li>
<li><strong>Connection Statistics</strong>: Added connection waiting statistics for better monitoring</li>
<li><strong>Failover Improvements</strong>: Added <code>ParseFailoverURL</code> for easier failover configuration</li>
<li><strong>Ring Client Enhancements</strong>: Added shard access methods for better Pub/Sub management</li>
</ul>
<h2>✨ New Features</h2>
<ul>
<li>Added <code>StreamingCredentialsProvider</code> for token-based authentication (<a href="https://redirect.github.com/redis/go-redis/pull/3320">#3320</a>)
<ul>
<li>Supports dynamic credential updates</li>
<li>Includes connection close hooks</li>
<li>Note: Currently marked as experimental</li>
</ul>
</li>
<li>Added <code>ParseFailoverURL</code> for parsing failover URLs (<a href="https://redirect.github.com/redis/go-redis/pull/3362">#3362</a>)</li>
<li>Added connection waiting statistics (<a href="https://redirect.github.com/redis/go-redis/pull/2804">#2804</a>)</li>
<li>Added new utility functions:
<ul>
<li><code>ParseFloat</code> and <code>MustParseFloat</code> in public utils package (<a href="https://redirect.github.com/redis/go-redis/pull/3371">#3371</a>)</li>
<li>Unit tests for <code>Atoi</code>, <code>ParseInt</code>, <code>ParseUint</code>, and <code>ParseFloat</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3377">#3377</a>)</li>
</ul>
</li>
<li>Added Ring client shard access methods:
<ul>
<li><code>GetShardClients()</code> to retrieve all active shard clients</li>
<li><code>GetShardClientForKey(key string)</code> to get the shard client for a specific key (<a href="https://redirect.github.com/redis/go-redis/pull/3388">#3388</a>)</li>
</ul>
</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>Fixed routing reads to loading slave nodes (<a href="https://redirect.github.com/redis/go-redis/pull/3370">#3370</a>)</li>
<li>Added support for nil lag in XINFO GROUPS (<a href="https://redirect.github.com/redis/go-redis/pull/3369">#3369</a>)</li>
<li>Fixed pool acquisition timeout issues (<a href="https://redirect.github.com/redis/go-redis/pull/3381">#3381</a>)</li>
<li>Optimized unnecessary copy operations (<a href="https://redirect.github.com/redis/go-redis/pull/3376">#3376</a>)</li>
</ul>
<h2>📚 Documentation</h2>
<ul>
<li>Updated documentation for XINFO GROUPS with nil lag support (<a href="https://redirect.github.com/redis/go-redis/pull/3369">#3369</a>)</li>
<li>Added package-level comments for new features</li>
</ul>
<h2>⚡ Performance and Reliability</h2>
<ul>
<li>Optimized <code>ReplaceSpaces</code> function (<a href="https://redirect.github.com/redis/go-redis/pull/3383">#3383</a>)</li>
<li>Set default value for <code>Options.Protocol</code> in <code>init()</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3387">#3387</a>)</li>
<li>Exported pool errors for public consumption (<a href="https://redirect.github.com/redis/go-redis/pull/3380">#3380</a>)</li>
</ul>
<h2>🔧 Dependencies and Infrastructure</h2>
<ul>
<li>Updated Redis CI to version 8.0.1 (<a href="https://redirect.github.com/redis/go-redis/pull/3372">#3372</a>)</li>
<li>Updated spellcheck GitHub Actions (<a href="https://redirect.github.com/redis/go-redis/pull/3389">#3389</a>)</li>
<li>Removed unused parameters (<a href="https://redirect.github.com/redis/go-redis/pull/3382">#3382</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3384">#3384</a>)</li>
</ul>
<h2>🧪 Testing</h2>
<ul>
<li>Added unit tests for pool acquisition timeout (<a href="https://redirect.github.com/redis/go-redis/pull/3381">#3381</a>)</li>
<li>Added unit tests for utility functions (<a href="https://redirect.github.com/redis/go-redis/pull/3377">#3377</a>)</li>
</ul>
<h2>👥 Contributors</h2>
<p>We would like to thank all the contributors who made this release possible:</p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/c935f9614893b42df6552550d70298f381bc3766"><code>c935f96</code></a> release(go-redis): v9.9.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3390">#3390</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/cb1968cad611420b844632393a0bd2881ece90ad"><code>cb1968c</code></a> feat(ring): add GetShardClients and GetShardClientForKey methods to Ring for ...</li>
<li><a href="https://github.com/redis/go-redis/commit/86d418f94077de29be613b9a2c4a324799f96d81"><code>86d418f</code></a> feat: Introducing StreamingCredentialsProvider for token based authentication...</li>
<li><a href="https://github.com/redis/go-redis/commit/28a3c974092a565fdf6af534fcba6da723ed8058"><code>28a3c97</code></a> chore: set the default value for the <code>options.protocol</code> in the <code>init()</code> of `o...</li>
<li><a href="https://github.com/redis/go-redis/commit/66b61c432cea98c2dda9567326ad3200b83845d3"><code>66b61c4</code></a> chore(deps): bump rojopolis/spellcheck-github-actions (<a href="https://redirect.github.com/redis/go-redis/issues/3389">#3389</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/d7ba255394435f444ec274b49c8464eac1cb3568"><code>d7ba255</code></a> fix: prevent routing reads to loading slave nodes (<a href="https://redirect.github.com/redis/go-redis/issues/3370">#3370</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b67455e0998b8767a47c1efd5828144f6709f85c"><code>b67455e</code></a> xinfo-groups: support nil lag in XINFO GROUPS (<a href="https://redirect.github.com/redis/go-redis/issues/3369">#3369</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/03c2c0b088521c67ca9ee8cfad73cfd35bedfb3f"><code>03c2c0b</code></a> chore: remove unused param (<a href="https://redirect.github.com/redis/go-redis/issues/3384">#3384</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c149644da7053ec196579401ff8ce184bf7ec013"><code>c149644</code></a> Unit test for pool acquisition timeout (<a href="https://redirect.github.com/redis/go-redis/issues/3381">#3381</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/bc70b52b42b1347dc78cfe851b9fda4c472589e9"><code>bc70b52</code></a> perf: avoid unnecessary copy operation (<a href="https://redirect.github.com/redis/go-redis/issues/3376">#3376</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.8.0...v9.9.0">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| gorm.io/driver/postgres | [< 1.6, > 1.5.4] |
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

## Discussion

### Comment by @jlsherrill on June 03, 2025 at 02:08 AM UTC

/retest

### Comment by @dependabot on June 03, 2025 at 03:09 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1119*
