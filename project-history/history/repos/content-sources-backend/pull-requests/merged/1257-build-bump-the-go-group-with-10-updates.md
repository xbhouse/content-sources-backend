---
type: pull_request
number: 1257
title: "Build: Bump the go group with 10 updates"
state: merged
author: dependabot
created: 2025-10-27T04:17:24Z
updated: 2025-10-27T09:42:02Z
closed: 2025-10-27T09:41:54Z
merged: 2025-10-27T09:41:54Z
base_branch: main
head_branch: dependabot/go_modules/go-7bc6add7e8
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1257
---

# Pull Request #1257: Build: Bump the go group with 10 updates

**Author**: @dependabot
**Created**: October 27, 2025 at 04:17 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-7bc6add7e8`

## Description

Bumps the go group with 10 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/redhatinsights/app-common-go](https://github.com/redhatinsights/app-common-go) | `1.6.8` | `1.6.9` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.46.2` | `1.46.3` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.39.3` | `1.39.4` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.13` | `1.31.15` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.17` | `1.18.19` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.58.3` | `1.58.5` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.88.5` | `1.88.7` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.10.1760721429` | `2025.10.1761380994` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.36.0` | `0.36.1` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.14.1` | `9.16.0` |

Updates `github.com/redhatinsights/app-common-go` from 1.6.8 to 1.6.9
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redhatinsights/app-common-go/releases">github.com/redhatinsights/app-common-go's releases</a>.</em></p>
<blockquote>
<h2>v1.6.9</h2>
<h2>What's Changed</h2>
<ul>
<li>Update for latest schema by <a href="https://github.com/bsquizz"><code>@​bsquizz</code></a> in <a href="https://redirect.github.com/RedHatInsights/app-common-go/pull/34">RedHatInsights/app-common-go#34</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/bsquizz"><code>@​bsquizz</code></a> made their first contribution in <a href="https://redirect.github.com/RedHatInsights/app-common-go/pull/34">RedHatInsights/app-common-go#34</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/app-common-go/compare/v1.6.8...v1.6.9">https://github.com/RedHatInsights/app-common-go/compare/v1.6.8...v1.6.9</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/app-common-go/commit/7bc20ad0862eee65cc37f1f782c93e44c981c9d4"><code>7bc20ad</code></a> Merge pull request <a href="https://redirect.github.com/redhatinsights/app-common-go/issues/34">#34</a> from bsquizz/update</li>
<li><a href="https://github.com/RedHatInsights/app-common-go/commit/0bb2d00f0c06cfb7d95c39955b62f942e7458d58"><code>0bb2d00</code></a> Update for latest schema</li>
<li>See full diff in <a href="https://github.com/redhatinsights/app-common-go/compare/v1.6.8...v1.6.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.46.2 to 1.46.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.46.3 (2025-10-26)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: wrap KError into error returned by IncrementalAlterConfig by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3352">IBM/sarama#3352</a></li>
<li>fix: assign sequence when flushing retry buffers by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3362">IBM/sarama#3362</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update dependency dominikh/go-tools to v2025 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3351">IBM/sarama#3351</a></li>
<li>chore(deps): update dependency vearutop/teststat to v0.1.27 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3350">IBM/sarama#3350</a></li>
<li>fix(deps): update module github.com/klauspost/compress to v1.18.1 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3355">IBM/sarama#3355</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(ci): extract tool versions and add renovate customManagers by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3346">IBM/sarama#3346</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.46.2...v1.46.3">https://github.com/IBM/sarama/compare/v1.46.2...v1.46.3</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/0447c9abdd8b66f6ee099320bbc21e88fb44e68c"><code>0447c9a</code></a> chore(ci): bump actions/setup-node from 5.0.0 to 6.0.0 in the actions group (...</li>
<li><a href="https://github.com/IBM/sarama/commit/7744b25fcebed358795363d9e19e335cdbd5c056"><code>7744b25</code></a> chore(ci): bump github/codeql-action from 4.30.8 to 4.30.9 (<a href="https://redirect.github.com/IBM/sarama/issues/3358">#3358</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/30f1c1d733ed4c20eeb66648e010747b761dd290"><code>30f1c1d</code></a> fix(deps): update module github.com/klauspost/compress to v1.18.1 (<a href="https://redirect.github.com/IBM/sarama/issues/3355">#3355</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/9a643a1aac64fd193a627448a5d745c4ec0bb034"><code>9a643a1</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.6 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/35ecd4887f2e6c9727192919592cf65fd8f10dc4"><code>35ecd48</code></a> fix: assign sequence when flushing retry buffers (<a href="https://redirect.github.com/IBM/sarama/issues/3362">#3362</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/0347c607be37372b5342160204216c1f086c19b1"><code>0347c60</code></a> fix: wrap KError into error returned by IncrementalAlterConfig (<a href="https://redirect.github.com/IBM/sarama/issues/3352">#3352</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e566998c576e37a3f4abfa5660b52dd0e841de8c"><code>e566998</code></a> chore(deps): update dependency vearutop/teststat to v0.1.27 (<a href="https://redirect.github.com/IBM/sarama/issues/3350">#3350</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4c38778a1851340c85df1db0cd946c3ca5616392"><code>4c38778</code></a> chore(deps): update dependency dominikh/go-tools to v2025 (<a href="https://redirect.github.com/IBM/sarama/issues/3351">#3351</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/9fc1146d97993ce2cab081175efb0d586f13625f"><code>9fc1146</code></a> chore(ci): bump actions/dependency-review-action from 4.8.0 to 4.8.1 in the a...</li>
<li><a href="https://github.com/IBM/sarama/commit/ce734e311d5ea7d24eed599e17314f4f57f0649b"><code>ce734e3</code></a> chore(ci): bump github/codeql-action from 3.30.6 to 4.30.8 (<a href="https://redirect.github.com/IBM/sarama/issues/3349">#3349</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.46.2...v1.46.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.39.3 to 1.39.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4bd7f5481eebe1c422fa85d1956f7ea34d93cf76"><code>4bd7f54</code></a> Release 2025-10-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55fb47cb07949ca70312a359272b10ff29f520df"><code>55fb47c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf727c0b40333a2e643a89909f189ffba0c212b9"><code>bf727c0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ca943fb071778b06e5e7f64f5ddf896f8579b6b"><code>0ca943f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3908bc4e960b7782da16a710fb8a747632af25a5"><code>3908bc4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0c74d28b606e71c1ee23d7ee17ec4949001cf56"><code>a0c74d2</code></a> Release 2025-10-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb611540d594538970ddd91b7802f5152cca8d2f"><code>eb61154</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e20d3e4b6065b1a8cdd97d00342cdc8d8c02561b"><code>e20d3e4</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9acd5faf2ca161219176c3cd8402837c98ed82ce"><code>9acd5fa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a91cc6c72206a1cef64caff0c70fa1bd13bd4543"><code>a91cc6c</code></a> Speed up unit tests by removing duplicate work (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3216">#3216</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.39.3...v1.39.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.13 to 1.31.15
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4bd7f5481eebe1c422fa85d1956f7ea34d93cf76"><code>4bd7f54</code></a> Release 2025-10-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55fb47cb07949ca70312a359272b10ff29f520df"><code>55fb47c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf727c0b40333a2e643a89909f189ffba0c212b9"><code>bf727c0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ca943fb071778b06e5e7f64f5ddf896f8579b6b"><code>0ca943f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3908bc4e960b7782da16a710fb8a747632af25a5"><code>3908bc4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0c74d28b606e71c1ee23d7ee17ec4949001cf56"><code>a0c74d2</code></a> Release 2025-10-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb611540d594538970ddd91b7802f5152cca8d2f"><code>eb61154</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e20d3e4b6065b1a8cdd97d00342cdc8d8c02561b"><code>e20d3e4</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9acd5faf2ca161219176c3cd8402837c98ed82ce"><code>9acd5fa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a91cc6c72206a1cef64caff0c70fa1bd13bd4543"><code>a91cc6c</code></a> Speed up unit tests by removing duplicate work (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3216">#3216</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.31.13...config/v1.31.15">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.17 to 1.18.19
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/390cf19a7507fe64e912123df2e1ad4a41a3f2c4"><code>390cf19</code></a> Release 2023-03-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c37c72ab935cbaa8816613808c1153c9da810583"><code>c37c72a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1e51932fdca008df18fc461095401f4d04f375a"><code>d1e5193</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2506101c1841a1816094280739e3a8d0a845d90b"><code>2506101</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c93b5ccbfae85a4c02349f59e3f05182b8885154"><code>c93b5cc</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2051">#2051</a> from aws/add100ContinueCustomization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c01aac6bfc14a2af5b66b5b696df247e3d68b162"><code>c01aac6</code></a> Keep one changelog for PR</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3780faa4eee262c3c701f586fca007bf4be17115"><code>3780faa</code></a> Keep one changelog for PR</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b94b5b7b7b969e869f38aa708da5d3274095b9a7"><code>b94b5b7</code></a> Merge remote-tracking branch 'origin/add100ContinueCustomization' into add100...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6174ff2c68e774053928a686c20075c6281774bb"><code>6174ff2</code></a> Change some variable name and use operation shape id to represent operation s...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/83491fca07a1af29fb9f6773aba1fbabaee2f329"><code>83491fc</code></a> add changelog to last commit</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.17...config/v1.18.19">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.58.3 to 1.58.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9b0c6553b80f99603b4f8356b88f5baf1328deb"><code>b9b0c65</code></a> Release 2025-10-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2bc8a0ec6f430876fc7de4432ea9cc89c9568f8"><code>e2bc8a0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8691ee380a96c49351e4b5ab8a70bc5d4d100724"><code>8691ee3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51e8a3fe032fc566d31b389f492ab58475a98398"><code>51e8a3f</code></a> bump to go1.23 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3211">#3211</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad2d36cba7c5772b4e8e4caf96939dc41b95c65c"><code>ad2d36c</code></a> Release 2025-10-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19a35d639f969ee328553e632e8cf8b83d324106"><code>19a35d6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35cb02fd50fb125601b9c3b33feb72f3a2bcaa56"><code>35cb02f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f673a1b0a80e666c0128ec606ff053dace9771f1"><code>f673a1b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48421fd812d8592a4eb2b32d11ae07e228969012"><code>48421fd</code></a> Release 2025-10-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fedcba778c21b451a91b4e4bcdd5d6c1554c6a5a"><code>fedcba7</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.58.3...service/route53/v1.58.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.88.5 to 1.88.7
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4bd7f5481eebe1c422fa85d1956f7ea34d93cf76"><code>4bd7f54</code></a> Release 2025-10-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55fb47cb07949ca70312a359272b10ff29f520df"><code>55fb47c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf727c0b40333a2e643a89909f189ffba0c212b9"><code>bf727c0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ca943fb071778b06e5e7f64f5ddf896f8579b6b"><code>0ca943f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3908bc4e960b7782da16a710fb8a747632af25a5"><code>3908bc4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0c74d28b606e71c1ee23d7ee17ec4949001cf56"><code>a0c74d2</code></a> Release 2025-10-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb611540d594538970ddd91b7802f5152cca8d2f"><code>eb61154</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e20d3e4b6065b1a8cdd97d00342cdc8d8c02561b"><code>e20d3e4</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9acd5faf2ca161219176c3cd8402837c98ed82ce"><code>9acd5fa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a91cc6c72206a1cef64caff0c70fa1bd13bd4543"><code>a91cc6c</code></a> Speed up unit tests by removing duplicate work (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3216">#3216</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.88.5...service/s3/v1.88.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.10.1760721429 to 2025.10.1761380994
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/fbda8bf092148781602f61a67ef93283de3827e7"><code>fbda8bf</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e4e9b4a2bb187e8639db952b...</li>
<li><a href="https://github.com/content-services/zest/commit/b9e246c2a3611ad40a7405fc50aed1769dd65366"><code>b9e246c</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a6eae4d525fb7a8d53bd4983...</li>
<li><a href="https://github.com/content-services/zest/commit/b6567fe4fa432eef1d746b3ccc29a2fd5717585a"><code>b6567fe</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b979258438b29c37ae5486b92ad...</li>
<li><a href="https://github.com/content-services/zest/commit/ab85b00296c82074f31ff4930fac4f7af60edcb8"><code>ab85b00</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e4e66338e8c87e8639db952b...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.10.1760721429...release/v2025.10.1761380994">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.36.0 to 0.36.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.36.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.36.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Prevent panic when converting error chains containing non-comparable error types by using a safe fallback for visited detection in exception conversion (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1113">#1113</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.36.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.36.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Prevent panic when converting error chains containing non-comparable error types by using a safe fallback for visited detection in exception conversion (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1113">#1113</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/1bdc7aaa3cfecfecff65be01358edc15c85439a2"><code>1bdc7aa</code></a> release: 0.36.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/fb7dacf7b69dde7d53a6742953add79b6a7c6d98"><code>fb7dacf</code></a> Prepare 0.36.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1114">#1114</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/11f2790668758ef92d9ad67496b1983acbb16bcb"><code>11f2790</code></a> fix: add fallback on non hashable errors (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1113">#1113</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2cf1dc0c1f43fad7a1b2ddba600b0d96c1681f21"><code>2cf1dc0</code></a> Merge branch 'release/0.36.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.36.0...v0.36.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.14.1 to 9.16.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.16.0</h2>
<h2>🚀 Highlights</h2>
<h3>Maintenance Notifications Support</h3>
<p>This release introduces comprehensive support for Redis maintenance notifications, enabling applications to handle server maintenance events gracefully. The new <code>maintnotifications</code> package provides:</p>
<ul>
<li><strong>RESP3 Push Notifications</strong>: Full support for Redis RESP3 protocol push notifications</li>
<li><strong>Connection Handoff</strong>: Automatic connection migration during server maintenance with configurable retry policies and circuit breakers</li>
<li><strong>Graceful Degradation</strong>: Configurable timeout relaxation during maintenance windows to prevent false failures</li>
<li><strong>Event-Driven Architecture</strong>: Background workers with on-demand scaling for efficient handoff processing</li>
</ul>
<p>For detailed usage examples and configuration options, see the <a href="https://github.com/redis/go-redis/tree/master/maintnotifications">maintenance notifications documentation</a>.</p>
<h2>✨ New Features</h2>
<ul>
<li><strong>Trace Filtering</strong>: Add support for filtering traces for specific commands, including pipeline operations and dial operations (<a href="https://redirect.github.com/redis/go-redis/pull/3519">#3519</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3550">#3550</a>)
<ul>
<li>New <code>TraceCmdFilter</code> option to selectively trace commands</li>
<li>Reduces overhead by excluding high-frequency or low-value commands from traces</li>
</ul>
</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Pipeline Error Handling</strong>: Fix issue where pipeline repeatedly sets the same error (<a href="https://redirect.github.com/redis/go-redis/pull/3525">#3525</a>)</li>
<li><strong>Connection Pool</strong>: Ensure re-authentication does not interfere with connection handoff operations (<a href="https://redirect.github.com/redis/go-redis/pull/3547">#3547</a>)</li>
</ul>
<h2>🔧 Improvements</h2>
<ul>
<li><strong>Hash Commands</strong>: Update hash command implementations (<a href="https://redirect.github.com/redis/go-redis/pull/3523">#3523</a>)</li>
<li><strong>OpenTelemetry</strong>: Use <code>metric.WithAttributeSet</code> to avoid unnecessary attribute copying in redisotel (<a href="https://redirect.github.com/redis/go-redis/pull/3552">#3552</a>)</li>
</ul>
<h2>📚 Documentation</h2>
<ul>
<li><strong>Cluster Client</strong>: Add explanation for why <code>MaxRetries</code> is disabled for <code>ClusterClient</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3551">#3551</a>)</li>
</ul>
<h2>🧪 Testing &amp; Infrastructure</h2>
<ul>
<li><strong>E2E Testing</strong>: Upgrade E2E testing framework with improved reliability and coverage (<a href="https://redirect.github.com/redis/go-redis/pull/3541">#3541</a>)</li>
<li><strong>Release Process</strong>: Improved resiliency of the release process (<a href="https://redirect.github.com/redis/go-redis/pull/3530">#3530</a>)</li>
</ul>
<h2>📦 Dependencies</h2>
<ul>
<li>Bump <code>rojopolis/spellcheck-github-actions</code> from 0.51.0 to 0.52.0 (<a href="https://redirect.github.com/redis/go-redis/pull/3520">#3520</a>)</li>
<li>Bump <code>github/codeql-action</code> from 3 to 4 (<a href="https://redirect.github.com/redis/go-redis/pull/3544">#3544</a>)</li>
</ul>
<h2>👥 Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/Sovietaced"><code>@​Sovietaced</code></a>, <a href="https://github.com/Udhayarajan"><code>@​Udhayarajan</code></a>, <a href="https://github.com/boekkooi-impossiblecloud"><code>@​boekkooi-impossiblecloud</code></a>, <a href="https://github.com/Pika-Gopher"><code>@​Pika-Gopher</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/huiyifyj"><code>@​huiyifyj</code></a>, <a href="https://github.com/omid-h70"><code>@​omid-h70</code></a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.16.0 (2025-10-23)</h1>
<h2>🚀 Highlights</h2>
<h3>Maintenance Notifications Support</h3>
<p>This release introduces comprehensive support for Redis maintenance notifications, enabling applications to handle server maintenance events gracefully. The new <code>maintnotifications</code> package provides:</p>
<ul>
<li><strong>RESP3 Push Notifications</strong>: Full support for Redis RESP3 protocol push notifications</li>
<li><strong>Connection Handoff</strong>: Automatic connection migration during server maintenance with configurable retry policies and circuit breakers</li>
<li><strong>Graceful Degradation</strong>: Configurable timeout relaxation during maintenance windows to prevent false failures</li>
<li><strong>Event-Driven Architecture</strong>: Background workers with on-demand scaling for efficient handoff processing</li>
<li><strong>Production-Ready</strong>: Comprehensive E2E testing framework and monitoring capabilities</li>
</ul>
<p>For detailed usage examples and configuration options, see the <a href="https://github.com/redis/go-redis/blob/master/maintnotifications/README.md">maintenance notifications documentation</a>.</p>
<h2>✨ New Features</h2>
<ul>
<li><strong>Trace Filtering</strong>: Add support for filtering traces for specific commands, including pipeline operations and dial operations (<a href="https://redirect.github.com/redis/go-redis/pull/3519">#3519</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3550">#3550</a>)
<ul>
<li>New <code>TraceCmdFilter</code> option to selectively trace commands</li>
<li>Reduces overhead by excluding high-frequency or low-value commands from traces</li>
</ul>
</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Pipeline Error Handling</strong>: Fix issue where pipeline repeatedly sets the same error (<a href="https://redirect.github.com/redis/go-redis/pull/3525">#3525</a>)</li>
<li><strong>Connection Pool</strong>: Ensure re-authentication does not interfere with connection handoff operations (<a href="https://redirect.github.com/redis/go-redis/pull/3547">#3547</a>)</li>
</ul>
<h2>🔧 Improvements</h2>
<ul>
<li><strong>Hash Commands</strong>: Update hash command implementations (<a href="https://redirect.github.com/redis/go-redis/pull/3523">#3523</a>)</li>
<li><strong>OpenTelemetry</strong>: Use <code>metric.WithAttributeSet</code> to avoid unnecessary attribute copying in redisotel (<a href="https://redirect.github.com/redis/go-redis/pull/3552">#3552</a>)</li>
</ul>
<h2>📚 Documentation</h2>
<ul>
<li><strong>Cluster Client</strong>: Add explanation for why <code>MaxRetries</code> is disabled for <code>ClusterClient</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3551">#3551</a>)</li>
</ul>
<h2>🧪 Testing &amp; Infrastructure</h2>
<ul>
<li><strong>E2E Testing</strong>: Upgrade E2E testing framework with improved reliability and coverage (<a href="https://redirect.github.com/redis/go-redis/pull/3541">#3541</a>)</li>
<li><strong>Release Process</strong>: Improved resiliency of the release process (<a href="https://redirect.github.com/redis/go-redis/pull/3530">#3530</a>)</li>
</ul>
<h2>📦 Dependencies</h2>
<ul>
<li>Bump <code>rojopolis/spellcheck-github-actions</code> from 0.51.0 to 0.52.0 (<a href="https://redirect.github.com/redis/go-redis/pull/3520">#3520</a>)</li>
<li>Bump <code>github/codeql-action</code> from 3 to 4 (<a href="https://redirect.github.com/redis/go-redis/pull/3544">#3544</a>)</li>
</ul>
<h2>👥 Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/f1956565eea26d3d0936985188b19009323acf7f"><code>f195656</code></a> chore(release): 9.16.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3557">#3557</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/70dfa383fee810afec312472691e4abd4d8be128"><code>70dfa38</code></a> feat(otel): add trace filter for process pipeline and dial operation (<a href="https://redirect.github.com/redis/go-redis/issues/3550">#3550</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a15e76394c80a8053d2790058d439fc5dd59c112"><code>a15e763</code></a> fix(pool): Pool ReAuth should not interfere with handoff (<a href="https://redirect.github.com/redis/go-redis/issues/3547">#3547</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/14a88145407e75dbff07471afe7fead8ea26a0cf"><code>14a8814</code></a> chore(docs): explain why MaxRetries is disabled for ClusterClient (<a href="https://redirect.github.com/redis/go-redis/issues/3551">#3551</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8ca21d2600f6cab552d22055cc9fd067f2d49d68"><code>8ca21d2</code></a> chore(redisotel): use metric.WithAttributeSet to avoid copy (<a href="https://redirect.github.com/redis/go-redis/issues/3552">#3552</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7aa4a606671d4b0ac3c311c42d4630931a9607e3"><code>7aa4a60</code></a> update gomods to align them with the latest beta (<a href="https://redirect.github.com/redis/go-redis/issues/3539">#3539</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/1e6ee067401605073600014ff66cb554541ed330"><code>1e6ee06</code></a> test(e2e): testing framework upgrade (<a href="https://redirect.github.com/redis/go-redis/issues/3541">#3541</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f7eed76fbcd1340d20981073276e81ca284ae189"><code>f7eed76</code></a> Add support for filtering traces for certain commands (<a href="https://redirect.github.com/redis/go-redis/issues/3519">#3519</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/3d68c7e42f549b4b584135442d102c00a27ef88d"><code>3d68c7e</code></a> chore(deps): bump github/codeql-action from 3 to 4 (<a href="https://redirect.github.com/redis/go-redis/issues/3544">#3544</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/3ad9f9cb2334227d5e59f5b7fc8e1612396756d2"><code>3ad9f9c</code></a> fix: add missing error variable for non-unix build constraints (<a href="https://redirect.github.com/redis/go-redis/issues/3538">#3538</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.14.1...v9.16.0">compare view</a></li>
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

### Review by @swadeley - Approved on October 27, 2025 at 09:41 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1257*
