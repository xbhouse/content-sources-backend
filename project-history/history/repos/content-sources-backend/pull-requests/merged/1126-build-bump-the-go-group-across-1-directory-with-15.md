---
type: pull_request
number: 1126
title: "Build: Bump the go group across 1 directory with 15 updates"
state: merged
author: dependabot
created: 2025-06-16T04:57:24Z
updated: 2025-06-16T12:50:36Z
closed: 2025-06-16T12:50:29Z
merged: 2025-06-16T12:50:29Z
base_branch: main
head_branch: dependabot/go_modules/go-28b9aa6692
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1126
---

# Pull Request #1126: Build: Bump the go group across 1 directory with 15 updates

**Author**: @dependabot
**Created**: June 16, 2025 at 04:57 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-28b9aa6692`

## Description

Bumps the go group with 12 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.11` | `1.6.0` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.45.1` | `1.45.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.3` | `1.36.4` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.14` | `1.29.16` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.50.0` | `1.50.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.79.4` | `1.80.2` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.15.2` | `2.16.1` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.5.1747750822` | `2025.6.1749665466` |
| [github.com/redhatinsights/platform-go-middlewares/v2](https://github.com/redhatinsights/platform-go-middlewares) | `2.0.0-beta.2` | `2.0.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.8.0` | `9.10.0` |
| [github.com/urfave/cli/v2](https://github.com/urfave/cli) | `2.27.6` | `2.27.7` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.13` | `1.25.14` |


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

Updates `github.com/aws/aws-sdk-go-v2` from 1.36.3 to 1.36.4
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.36.3...v1.36.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.14 to 1.29.16
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.14...config/v1.29.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.67 to 1.17.69
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.67...credentials/v1.17.69">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.50.0 to 1.50.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de4f646b7271b9698a67f63bff876ad40eb690ca"><code>de4f646</code></a> Release 2024-02-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a31327acbd07ca1d42ca83107993d7bcfa223dd8"><code>a31327a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1333952165a0f8ceea1c7de5d20d06af4925c30"><code>d133395</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b135f991ff3bef2f7c9ce8117dd59e1abae43a0f"><code>b135f99</code></a> fix(2502): zero region should explicitly fail endpoint resolution (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2503">#2503</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7cb105f17b3a12d4b69219ce9957d517dba299d8"><code>7cb105f</code></a> Release 2024-02-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f9c4d2eee136647971352eb9d06b196cc9abaf0"><code>7f9c4d2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a60d4b2fc37bab61dc219d9e210369c4bf4e445e"><code>a60d4b2</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be0accd6c6511b727d38cbe98489753e481d68b3"><code>be0accd</code></a> fix: panic due to potentially uncomparable wrapped fn in express_resolve (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2500">#2500</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.50.0...service/s3/v1.50.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.4 to 1.80.2
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.4...service/s3/v1.80.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2` from 2.15.2 to 2.16.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.16.1</h2>
<h1>CloudEvents SDK Go v2.16.1</h1>
<h2>🐛 Bug Fixes and Improvements</h2>
<ul>
<li>
<p><strong>⚡ NATS JetStream Enhancement</strong>: Made send subject optional via context by <a href="https://github.com/kmpm"><code>@​kmpm</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1143">cloudevents/sdk-go#1143</a></p>
<ul>
<li>Added WithSubject function to override the default subject when sending messages</li>
<li>Added comprehensive tests and updated samples</li>
<li>Non-breaking enhancement that adds flexibility for NATS users</li>
</ul>
</li>
<li>
<p><strong>📝 CloudEvents JSON Handling Fixes</strong> by <a href="https://github.com/alank-ps"><code>@​alank-ps</code></a>:</p>
<ul>
<li><strong>WriteJson Fix</strong> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1162">cloudevents/sdk-go#1162</a>: Fixed WriteJson to properly handle data as JSON when dataContentType is <code>application/cloudevents+json</code> or batch</li>
<li><strong>ConsumeData Fix</strong> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1164">cloudevents/sdk-go#1164</a>: Fixed consumeData functions to properly recognize structured mode JSON content types</li>
<li>Improves compatibility with the CloudEvents specification</li>
</ul>
</li>
<li>
<p><strong>🔧 CI/Test Improvements</strong>: Fix failing CI tests by <a href="https://github.com/embano1"><code>@​embano1</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1156">cloudevents/sdk-go#1156</a></p>
</li>
</ul>
<h2>🔄 Maintenance and Dependency Updates</h2>
<ul>
<li><strong>🛠️ Dependency Management Overhaul</strong> by <a href="https://github.com/embano1"><code>@​embano1</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1145">cloudevents/sdk-go#1145</a>
<ul>
<li>Added script (<code>hack/update-deps.sh</code>) to update Go dependencies across all modules</li>
<li>Replaced Dependabot with custom script for better dependency management</li>
<li>Removed stale and broken OpenTelemetry samples</li>
</ul>
</li>
</ul>
<p><strong>📦 Key Dependency Updates:</strong></p>
<ul>
<li><strong>github.com/google/go-cmp</strong>: v0.6.0 → v0.7.0</li>
<li><strong>golang.org/x/sync</strong>: v0.12.0 → v0.13.0</li>
<li><strong>github.com/nats-io/nats.go</strong>: v1.37.0 → v1.41.2</li>
<li><strong>github.com/IBM/sarama</strong>: v1.40.1 → v1.45.1</li>
<li><strong>github.com/docker/docker</strong>: v20.10.17 → v27.1.1</li>
<li><strong>go.opentelemetry.io/otel</strong>: v1.18.0 → v1.35.0</li>
<li><strong>🐹 Go version</strong>: Updated from 1.22 to 1.23.0 (toolchain 1.23.8)</li>
</ul>
<h2>🚨 Breaking Changes</h2>
<p>None. All updates are either backward-compatible improvements, bug fixes, or internal refactors.</p>
<h2>👥 New Contributors</h2>
<ul>
<li><a href="https://github.com/github-actions"><code>@​github-actions</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1147">cloudevents/sdk-go#1147</a></li>
<li><a href="https://github.com/kmpm"><code>@​kmpm</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1143">cloudevents/sdk-go#1143</a></li>
<li><a href="https://github.com/alank-ps"><code>@​alank-ps</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1162">cloudevents/sdk-go#1162</a></li>
</ul>
<h2>📋 What's Changed</h2>
<ul>
<li>Dependency Management Improvements by <a href="https://github.com/embano1"><code>@​embano1</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1145">cloudevents/sdk-go#1145</a></li>
<li>chore(deps): Bump the go_modules group across 2 directories with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1146">cloudevents/sdk-go#1146</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1147">cloudevents/sdk-go#1147</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1148">cloudevents/sdk-go#1148</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1149">cloudevents/sdk-go#1149</a></li>
<li>chore(deps-dev): Bump nokogiri from 1.18.4 to 1.18.8 in /docs in the bundler group across 1 directory by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1150">cloudevents/sdk-go#1150</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1151">cloudevents/sdk-go#1151</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1152">cloudevents/sdk-go#1152</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/65b45e42ad3c1be4d50b2a02aaf7313f35d3997c"><code>65b45e4</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1164">#1164</a> from alank-ps/main</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/bb544e06ca8021c6a000d89a6123001444898ad1"><code>bb544e0</code></a> fix: simplify isJSON</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/1bf32a1d22798613db3fddc64427d4c3da60d45c"><code>1bf32a1</code></a> fix: extract IsJson() check to content_type</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/8ad2c064178603736005452336421a5c0bfa336c"><code>8ad2c06</code></a> fix: consumeData(and ...AsBytes) should consider structued mode JSON content ...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/506a8fecb319156ef913a61f55ae27d63c532c12"><code>506a8fe</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1165">#1165</a> from cloudevents/automated-dependency-updates</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/5347cb097a90468a993a9e06bccbc5c58374d0b4"><code>5347cb0</code></a> chore: update dependencies</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/a50d97a11f88b611e9cda2727a0f3e617cbcf06e"><code>a50d97a</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1162">#1162</a> from alank-ps/main</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/753ba72db7084a88384d7c1cdd7c047d875afefd"><code>753ba72</code></a> fix: WriteJson should write data as a JSON value when dataContentType=applica...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/0d35f37e1c151b774feba277f1747d7b5d1a5ebb"><code>0d35f37</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1163">#1163</a> from cloudevents/automated-dependency-updates</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/4546fc768170e4cb0a751a4ed823e564450414cb"><code>4546fc7</code></a> chore: update dependencies</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.2...v2.16.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/v2` from 2.16.0 to 2.16.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.16.1</h2>
<h1>CloudEvents SDK Go v2.16.1</h1>
<h2>🐛 Bug Fixes and Improvements</h2>
<ul>
<li>
<p><strong>⚡ NATS JetStream Enhancement</strong>: Made send subject optional via context by <a href="https://github.com/kmpm"><code>@​kmpm</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1143">cloudevents/sdk-go#1143</a></p>
<ul>
<li>Added WithSubject function to override the default subject when sending messages</li>
<li>Added comprehensive tests and updated samples</li>
<li>Non-breaking enhancement that adds flexibility for NATS users</li>
</ul>
</li>
<li>
<p><strong>📝 CloudEvents JSON Handling Fixes</strong> by <a href="https://github.com/alank-ps"><code>@​alank-ps</code></a>:</p>
<ul>
<li><strong>WriteJson Fix</strong> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1162">cloudevents/sdk-go#1162</a>: Fixed WriteJson to properly handle data as JSON when dataContentType is <code>application/cloudevents+json</code> or batch</li>
<li><strong>ConsumeData Fix</strong> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1164">cloudevents/sdk-go#1164</a>: Fixed consumeData functions to properly recognize structured mode JSON content types</li>
<li>Improves compatibility with the CloudEvents specification</li>
</ul>
</li>
<li>
<p><strong>🔧 CI/Test Improvements</strong>: Fix failing CI tests by <a href="https://github.com/embano1"><code>@​embano1</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1156">cloudevents/sdk-go#1156</a></p>
</li>
</ul>
<h2>🔄 Maintenance and Dependency Updates</h2>
<ul>
<li><strong>🛠️ Dependency Management Overhaul</strong> by <a href="https://github.com/embano1"><code>@​embano1</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1145">cloudevents/sdk-go#1145</a>
<ul>
<li>Added script (<code>hack/update-deps.sh</code>) to update Go dependencies across all modules</li>
<li>Replaced Dependabot with custom script for better dependency management</li>
<li>Removed stale and broken OpenTelemetry samples</li>
</ul>
</li>
</ul>
<p><strong>📦 Key Dependency Updates:</strong></p>
<ul>
<li><strong>github.com/google/go-cmp</strong>: v0.6.0 → v0.7.0</li>
<li><strong>golang.org/x/sync</strong>: v0.12.0 → v0.13.0</li>
<li><strong>github.com/nats-io/nats.go</strong>: v1.37.0 → v1.41.2</li>
<li><strong>github.com/IBM/sarama</strong>: v1.40.1 → v1.45.1</li>
<li><strong>github.com/docker/docker</strong>: v20.10.17 → v27.1.1</li>
<li><strong>go.opentelemetry.io/otel</strong>: v1.18.0 → v1.35.0</li>
<li><strong>🐹 Go version</strong>: Updated from 1.22 to 1.23.0 (toolchain 1.23.8)</li>
</ul>
<h2>🚨 Breaking Changes</h2>
<p>None. All updates are either backward-compatible improvements, bug fixes, or internal refactors.</p>
<h2>👥 New Contributors</h2>
<ul>
<li><a href="https://github.com/github-actions"><code>@​github-actions</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1147">cloudevents/sdk-go#1147</a></li>
<li><a href="https://github.com/kmpm"><code>@​kmpm</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1143">cloudevents/sdk-go#1143</a></li>
<li><a href="https://github.com/alank-ps"><code>@​alank-ps</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1162">cloudevents/sdk-go#1162</a></li>
</ul>
<h2>📋 What's Changed</h2>
<ul>
<li>Dependency Management Improvements by <a href="https://github.com/embano1"><code>@​embano1</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1145">cloudevents/sdk-go#1145</a></li>
<li>chore(deps): Bump the go_modules group across 2 directories with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1146">cloudevents/sdk-go#1146</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1147">cloudevents/sdk-go#1147</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1148">cloudevents/sdk-go#1148</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1149">cloudevents/sdk-go#1149</a></li>
<li>chore(deps-dev): Bump nokogiri from 1.18.4 to 1.18.8 in /docs in the bundler group across 1 directory by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1150">cloudevents/sdk-go#1150</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1151">cloudevents/sdk-go#1151</a></li>
<li>chore: update dependencies by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1152">cloudevents/sdk-go#1152</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/65b45e42ad3c1be4d50b2a02aaf7313f35d3997c"><code>65b45e4</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1164">#1164</a> from alank-ps/main</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/bb544e06ca8021c6a000d89a6123001444898ad1"><code>bb544e0</code></a> fix: simplify isJSON</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/1bf32a1d22798613db3fddc64427d4c3da60d45c"><code>1bf32a1</code></a> fix: extract IsJson() check to content_type</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/8ad2c064178603736005452336421a5c0bfa336c"><code>8ad2c06</code></a> fix: consumeData(and ...AsBytes) should consider structued mode JSON content ...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/506a8fecb319156ef913a61f55ae27d63c532c12"><code>506a8fe</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1165">#1165</a> from cloudevents/automated-dependency-updates</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/5347cb097a90468a993a9e06bccbc5c58374d0b4"><code>5347cb0</code></a> chore: update dependencies</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/a50d97a11f88b611e9cda2727a0f3e617cbcf06e"><code>a50d97a</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1162">#1162</a> from alank-ps/main</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/753ba72db7084a88384d7c1cdd7c047d875afefd"><code>753ba72</code></a> fix: WriteJson should write data as a JSON value when dataContentType=applica...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/0d35f37e1c151b774feba277f1747d7b5d1a5ebb"><code>0d35f37</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1163">#1163</a> from cloudevents/automated-dependency-updates</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/4546fc768170e4cb0a751a4ed823e564450414cb"><code>4546fc7</code></a> chore: update dependencies</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.16.0...v2.16.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.5.1747750822 to 2025.6.1749665466
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/cc184b2d6f7f8bcba0fa47d70814e540d9ac33e5"><code>cc184b2</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a656a3e4e40b7a8d53bd4983...</li>
<li><a href="https://github.com/content-services/zest/commit/6f80ee46117706c9dc2762a0f7b0aa4ef4a0bd0c"><code>6f80ee4</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47548436243d0b7a98d596b8e3...</li>
<li><a href="https://github.com/content-services/zest/commit/abef0dd48c45c02a9c49f2381cb5ad4c455580d5"><code>abef0dd</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e424e98a9bc87e8639db952b...</li>
<li><a href="https://github.com/content-services/zest/commit/bd7a96dd8c5dbb444dc601c912fec12359a3985b"><code>bd7a96d</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276abb3d9b5e127b8e3dae9584...</li>
<li><a href="https://github.com/content-services/zest/commit/8a76545f1ac4d15e7328713e464988ec7a31c3a4"><code>8a76545</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b9792eed98d23f37ae5486b92ad...</li>
<li><a href="https://github.com/content-services/zest/commit/d915defb47ce929dda39c4e8212c8f682ab51069"><code>d915def</code></a> Update pulp bindings to e69db48356e528a464be3da896237b3bb5b38490a7d54eb5892a9...</li>
<li><a href="https://github.com/content-services/zest/commit/375d28f999edeb4c807e15a2d064f4db21e1499b"><code>375d28f</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276abbb2e989f27b8e3dae9584...</li>
<li><a href="https://github.com/content-services/zest/commit/9f20ab31c47cb76a5ff8d2a0d2968d4c0917de01"><code>9f20ab3</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7ddd45b9b9f57952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/39323bb33a0ebc01b3bf1d87dfe6b0a58456d760"><code>39323bb</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e8828e449087e8639db952b3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.5.1747750822...release/v2025.6.1749665466">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redhatinsights/platform-go-middlewares/v2` from 2.0.0-beta.2 to 2.0.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/e8b2f99d29f6ac5a08287df33860a4b9717c7798"><code>e8b2f99</code></a> Stable release v2</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/399ba78709f1c64225b9eb1f9930836e29bb877e"><code>399ba78</code></a> Add missing <code>user_id</code> field to ServiceAccount identity</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/b08e86968acc07cc933eb6b72345915f04877e03"><code>b08e869</code></a> chore: add gobump weekly</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/36c515bb681c408def588e3c85657a38a36cadda"><code>36c515b</code></a> chore: update to Go 1.23+</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/2d456cc74a294daa16460db6fc7aafb10720a8d4"><code>2d456cc</code></a> chore(deps): bump golang.org/x/net to 0.33.0</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/e3ddb9635cb5fd4c7ee9d0cf92075f4c0fd05380"><code>e3ddb96</code></a> Bump github.com/aws/aws-sdk-go from 1.55.3 to 1.55.5</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/bf8970f36989112e91dbc227ae7e8c5f995175db"><code>bf8970f</code></a> Bump github.com/onsi/gomega from 1.33.1 to 1.34.1</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/dfe493f69f2e809cbd0bb3089b46cc48ea7498f3"><code>dfe493f</code></a> Bump github.com/aws/aws-sdk-go from 1.54.20 to 1.55.3</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/dbdff99a6bb469a65df34046085d57c468e0eafe"><code>dbdff99</code></a> Bump github.com/aws/aws-sdk-go from 1.54.19 to 1.54.20</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/44143f6300d86c94a78cdfc5ad5bbaf66b60d53d"><code>44143f6</code></a> Bump github.com/aws/aws-sdk-go from 1.54.15 to 1.54.19</li>
<li>Additional commits viewable in <a href="https://github.com/redhatinsights/platform-go-middlewares/compare/v2.0.0-beta.2...v2.0.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.8.0 to 9.10.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.10.0</h2>
<p>Experimental support for  <a href="https://redis.io/docs/latest/develop/data-types/vector-sets/">vector sets</a>!</p>
<h2>🚀 Highlights</h2>
<p><code>go-redis</code> now supports <a href="https://redis.io/docs/latest/develop/data-types/vector-sets/">vector sets</a>. This data type is marked as &quot;in preview&quot; in Redis and its support in <code>go-redis</code> is marked as experimental. You can find examples in the documentation and in the <code>doctests</code> folder.</p>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>feat: support vectorset (<a href="https://redirect.github.com/redis/go-redis/pull/3375">#3375</a>) <a href="https://github.com/fukua95"><code>@​fukua95</code></a></li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>Add the missing NewFloatSliceResult for testing (<a href="https://redirect.github.com/redis/go-redis/pull/3393">#3393</a>)</li>
<li>DOC-5078 vector set examples (<a href="https://redirect.github.com/redis/go-redis/pull/3394">#3394</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/AndBobsYourUncle"><code>@​AndBobsYourUncle</code></a>, <a href="https://github.com/andy-stark-redis"><code>@​andy-stark-redis</code></a>, <a href="https://github.com/fukua95"><code>@​fukua95</code></a> and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
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
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/75e8370a6f08f55b5337524040de9c2b95687582"><code>75e8370</code></a> chore(release): Update release notes and versions for v9.10.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3395">#3395</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/98b224cf7d7fc9ac86feb8afb5ee212b44e209c4"><code>98b224c</code></a> DOC-5078 vector set examples (<a href="https://redirect.github.com/redis/go-redis/issues/3394">#3394</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/4e22885ca14e81e221039ce8955a5364172b638c"><code>4e22885</code></a> feat: support vectorset (<a href="https://redirect.github.com/redis/go-redis/issues/3375">#3375</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c609828c9b0cdeaf1d5ee5cce12778a3bbf5f829"><code>c609828</code></a> chore(tests): add the missing NewFloatSliceResult for testing (<a href="https://redirect.github.com/redis/go-redis/issues/3393">#3393</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a7772e47c0969fd6bedf6d8fd8e674c6e44a6edc"><code>a7772e4</code></a> Update RELEASE-NOTES.md (<a href="https://redirect.github.com/redis/go-redis/issues/3391">#3391</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c935f9614893b42df6552550d70298f381bc3766"><code>c935f96</code></a> release(go-redis): v9.9.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3390">#3390</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/cb1968cad611420b844632393a0bd2881ece90ad"><code>cb1968c</code></a> feat(ring): add GetShardClients and GetShardClientForKey methods to Ring for ...</li>
<li><a href="https://github.com/redis/go-redis/commit/86d418f94077de29be613b9a2c4a324799f96d81"><code>86d418f</code></a> feat: Introducing StreamingCredentialsProvider for token based authentication...</li>
<li><a href="https://github.com/redis/go-redis/commit/28a3c974092a565fdf6af534fcba6da723ed8058"><code>28a3c97</code></a> chore: set the default value for the <code>options.protocol</code> in the <code>init()</code> of `o...</li>
<li><a href="https://github.com/redis/go-redis/commit/66b61c432cea98c2dda9567326ad3200b83845d3"><code>66b61c4</code></a> chore(deps): bump rojopolis/spellcheck-github-actions (<a href="https://redirect.github.com/redis/go-redis/issues/3389">#3389</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.8.0...v9.10.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/urfave/cli/v2` from 2.27.6 to 2.27.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/urfave/cli/releases">github.com/urfave/cli/v2's releases</a>.</em></p>
<blockquote>
<h2>v2.27.7</h2>
<h2>What's Changed</h2>
<ul>
<li>Update dependencies in v2 series by <a href="https://github.com/meatballhat"><code>@​meatballhat</code></a> in <a href="https://redirect.github.com/urfave/cli/pull/2159">urfave/cli#2159</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/urfave/cli/compare/v2.27.6...v2.27.7">https://github.com/urfave/cli/compare/v2.27.6...v2.27.7</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/urfave/cli/commit/19b951ab78929023a9a670722b26ffb1d67c733a"><code>19b951a</code></a> Merge pull request <a href="https://redirect.github.com/urfave/cli/issues/2159">#2159</a> from urfave/v2-deps-up</li>
<li><a href="https://github.com/urfave/cli/commit/dd134b57770c1c3565ae1f6a23dd487cb129dae2"><code>dd134b5</code></a> Update dependencies in v2 series</li>
<li><a href="https://github.com/urfave/cli/commit/347cd02a7b6384bbd9ec0348dad1f8fac17229c3"><code>347cd02</code></a> Merge pull request <a href="https://redirect.github.com/urfave/cli/issues/2157">#2157</a> from urfave/v2-not-dependabot</li>
<li><a href="https://github.com/urfave/cli/commit/0acf2e41ca6fd9bd5091851f2dbf7d26c49d1f30"><code>0acf2e4</code></a> Dependabot does not work like this (v2)</li>
<li><a href="https://github.com/urfave/cli/commit/c7bc0a984e0b13ebdf7377ddf3c4b55ff2c749f8"><code>c7bc0a9</code></a> Merge pull request <a href="https://redirect.github.com/urfave/cli/issues/2154">#2154</a> from urfave/v2-dependabot-maybe</li>
<li><a href="https://github.com/urfave/cli/commit/6ec03684f1d853a9b51de4d5ac54a62747c5864d"><code>6ec0368</code></a> Is this file needed on each release branch? (v2)</li>
<li>See full diff in <a href="https://github.com/urfave/cli/compare/v2.27.6...v2.27.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.13 to 1.25.14
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>v1.25.14</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/getkin/kin-openapi from 0.131.0 to 0.132.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/612">RedHatInsights/insights-operator-utils#612</a></li>
<li>Bump github.com/redis/go-redis/v9 from 9.7.3 to 9.8.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/613">RedHatInsights/insights-operator-utils#613</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.55.6 to 1.55.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/610">RedHatInsights/insights-operator-utils#610</a></li>
<li>Bump github.com/getsentry/sentry-go from 0.32.0 to 0.33.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/614">RedHatInsights/insights-operator-utils#614</a></li>
<li>Golangci lint by <a href="https://github.com/ikerreyes"><code>@​ikerreyes</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/615">RedHatInsights/insights-operator-utils#615</a></li>
<li>Bump golang.org/x/net from 0.33.0 to 0.38.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/609">RedHatInsights/insights-operator-utils#609</a></li>
<li>Bump github.com/redis/go-redis/v9 from 9.8.0 to 9.9.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/616">RedHatInsights/insights-operator-utils#616</a></li>
<li>modify cloudwatch logger to use HOSTNAME as log stream name in case it's not provided by <a href="https://github.com/Bee-lee"><code>@​Bee-lee</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/617">RedHatInsights/insights-operator-utils#617</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/ikerreyes"><code>@​ikerreyes</code></a> made their first contribution in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/615">RedHatInsights/insights-operator-utils#615</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.13...v1.25.14">https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.13...v1.25.14</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/2e029d158d8dc217ba6c5211882bcfce483e8452"><code>2e029d1</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/617">#617</a> from Bee-lee/dbily-/UJ1yK3O</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/9f1295950f30f24fb422474a80f628d94ae2f02b"><code>9f12959</code></a> use os.Hostname() instead of the env var, as it might not be populated yet</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/ee952a21ebc3f6c0b28c46510fe03261a7512f81"><code>ee952a2</code></a> try using os.Hostname instead of the env var (github runner doesn't have a hn)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4331497789496f0333072fd2259b66e16b457c53"><code>4331497</code></a> add test case to cloudwatch logging to check new log stream name behavior</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/15314cf5ef195fa0d567cf771aeb43fe5ea6e4a0"><code>15314cf</code></a> modify cloudwatch logger to use HOSTNAME as log stream name in case it's not ...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/9e6d2b8ec234f181015fa6ad72c2eb3260ba4a1a"><code>9e6d2b8</code></a> Merge branch 'master' of github.com:RedHatInsights/insights-operator-utils</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/8f4df589cc4f0b4cc2e5e4eccf14648449a6c76e"><code>8f4df58</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/616">#616</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/fb06a90ac9e516cbdfd771878722b4405e3cae35"><code>fb06a90</code></a> Bump github.com/redis/go-redis/v9 from 9.8.0 to 9.9.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/85eba70a4f4c608747fb81ae189eeeb26189e9d4"><code>85eba70</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/609">#609</a> from RedHatInsights/dependabot/go_modules/golang.org/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/8ef5e6c9891741e910da528ac44be7431c0c23a4"><code>8ef5e6c</code></a> Bump golang.org/x/net from 0.33.0 to 0.38.0</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.13...v1.25.14">compare view</a></li>
</ul>
</details>
<br />

Updates `golang.org/x/text` from 0.25.0 to 0.26.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang/text/commit/80721808805f9d846d907c85d73ca6b5b6ecb870"><code>8072180</code></a> go.mod: update golang.org/x dependencies</li>
<li><a href="https://github.com/golang/text/commit/6cacac16ce2f94d21bf864011f5e0d2181a1a8bf"><code>6cacac1</code></a> go.mod: update tagx:ignore'd golang.org/x dependencies</li>
<li>See full diff in <a href="https://github.com/golang/text/compare/v0.25.0...v0.26.0">compare view</a></li>
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

## Reviews

### Review by @swadeley - Approved on June 16, 2025 at 12:50 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1126*
