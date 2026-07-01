---
type: pull_request
number: 937
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2025-01-13T05:03:25Z
updated: 2025-01-14T17:29:34Z
closed: 2025-01-14T17:29:27Z
merged: 2025-01-14T17:29:26Z
base_branch: main
head_branch: dependabot/go_modules/go-e8b749b528
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/937
---

# Pull Request #937: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: January 13, 2025 at 05:03 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-e8b749b528`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.3` | `1.1.4` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.44.0` | `1.45.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.7` | `1.32.8` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.48` | `1.17.51` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.45.1` | `1.45.3` |
| [github.com/lzap/cloudwatchwriter2](https://github.com/lzap/cloudwatchwriter2) | `1.3.0` | `1.4.2` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.3 to 1.1.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.1.4</h2>
<h2>What's Changed</h2>
<ul>
<li>Emit armor headers in sorted order by <a href="https://github.com/andrewgdotcom"><code>@​andrewgdotcom</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/255">ProtonMail/go-crypto#255</a></li>
<li>Reduce memory usage when AEAD en/decrypting large messages by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/259">ProtonMail/go-crypto#259</a></li>
<li>Update artifact actions to v4 by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/260">ProtonMail/go-crypto#260</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.3...v1.1.4">https://github.com/ProtonMail/go-crypto/compare/v1.1.3...v1.1.4</a></p>
<h2>Release v1.1.4-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.4 with support for the following non-standardized features:</p>
<ul>
<li>Presistent symmetric keys <a href="https://www.ietf.org/archive/id/draft-ietf-openpgp-persistent-symmetric-keys-00.html">draft-ietf-openpgp-persistent-symmetric-keys-00</a></li>
<li>Automatic forwarding <a href="https://www.ietf.org/archive/id/draft-wussler-openpgp-forwarding-00.html">draft-wussler-openpgp-forwarding-00</a></li>
<li>Post-quantum algorithms <a href="https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/">draft-ietf-openpgp-pqc</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3de0301376b64916c60452ffa511647d0f2bc918"><code>3de0301</code></a> Update artifact actions to v4 (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/260">#260</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/be3aef043fdbdabd8239edf24afe70922011212d"><code>be3aef0</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/259">#259</a> from ProtonMail/less-memory-large-msgs</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/1fd5ec842fd37133206b23cba68eda6ae4b905e3"><code>1fd5ec8</code></a> Add tests for reusing buffer in OCB en/decryption</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/df3ee02d77623299d88df3d8a29486b548300d72"><code>df3ee02</code></a> Buffer decrypted bytes more efficiently</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/04cfaf2d9d69903d32db0b4012ad102ebff4a511"><code>04cfaf2</code></a> Reuse plaintext slice for ciphertext when encrypting</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/fee782456821f3ab5f1de998d3e959c7e2cbae98"><code>fee7824</code></a> Reuse ciphertext slice for plaintext when decrypting</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/6fa7f918eaaa1ff22e1b16dfeed2041ec117fe2a"><code>6fa7f91</code></a> Preallocate the chunk size rather than buffering</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/add07bdad9c7162c9cd0898ff615669f95076354"><code>add07bd</code></a> Don't allocate the nonce for each chunk</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b01f065b251063bcab64bdab933e998ea4315ee1"><code>b01f065</code></a> Emit armor headers in deterministically sorted order (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/255">#255</a>)</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.3...v1.1.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.44.0 to 1.45.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.45.0 (2025-01-07)</h2>
<!-- raw HTML omitted -->
<blockquote>
<p>[!NOTE]<br />
The go.mod directive has been bumped to 1.21 as the minimum version of Go required for the module. This was necessary to continue to receive updates from some of the third party dependencies that Sarama makes use of for compression.</p>
</blockquote>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(admin): add retries for GroupCoordinator errors by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3053">IBM/sarama#3053</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump github.com/klauspost/compress from 1.17.9 to 1.17.11 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2999">IBM/sarama#2999</a></li>
<li>chore(deps): bump golang.org/x/net from 0.33.0 to 0.34.0 in the golang-org-x group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3054">IBM/sarama#3054</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: bump minimum go to 1.21 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3048">IBM/sarama#3048</a></li>
<li>chore(test): tag all unittests as !integration by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3047">IBM/sarama#3047</a></li>
<li>chore(test): include kafka 4.0.0 in FV testing by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3045">IBM/sarama#3045</a></li>
<li>fix(ci): restore the Kafka 4.0.0 FV by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3055">IBM/sarama#3055</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.44.0...v1.45.0">https://github.com/IBM/sarama/compare/v1.44.0...v1.45.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/061e437d7d1ef7dba9a8420ac968d87cd5912ccf"><code>061e437</code></a> chore(deps): bump golang.org/x/net in the golang-org-x group (<a href="https://redirect.github.com/IBM/sarama/issues/3054">#3054</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/73852040824158ecc6b9046ae1d708c05a8e94a0"><code>7385204</code></a> Merge pull request <a href="https://redirect.github.com/IBM/sarama/issues/3053">#3053</a> from IBM/dnwe/admin</li>
<li><a href="https://github.com/IBM/sarama/commit/ecbf73af3ad330008903d024fa5bf40958bb9793"><code>ecbf73a</code></a> fix(admin): use named return on retryOnError func</li>
<li><a href="https://github.com/IBM/sarama/commit/060bb3f8da230bb60f00c5de51bc21ce8d7bfa91"><code>060bb3f</code></a> fix(admin): add retries for GroupCoordinator errors</li>
<li><a href="https://github.com/IBM/sarama/commit/80fea25df76a7a73b811db456cf053d7e005d0e5"><code>80fea25</code></a> fix: correct error message for COORDINATOR_NOT_AVAILABLE</li>
<li><a href="https://github.com/IBM/sarama/commit/75bdfbc65ba567010a3962b3a09133013365e524"><code>75bdfbc</code></a> chore(ci): use container_name rather than hostname</li>
<li><a href="https://github.com/IBM/sarama/commit/5ef8084a5bf79e5533e821d59a9b58997ff1e7a5"><code>5ef8084</code></a> fix(ci): patch zookeeper client to 3.5.9 on older Kafka</li>
<li><a href="https://github.com/IBM/sarama/commit/1030b6174d7d592ffdeb701039dc3c7ffa6b7879"><code>1030b61</code></a> chore(ci): add cache cleanup job for closed PRs (<a href="https://redirect.github.com/IBM/sarama/issues/3056">#3056</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/85a9c9fdc6f44a5927469244a3938db56a7f1e08"><code>85a9c9f</code></a> Merge pull request <a href="https://redirect.github.com/IBM/sarama/issues/3055">#3055</a> from IBM/dnwe/fix-fv</li>
<li><a href="https://github.com/IBM/sarama/commit/40117dec6571cfbc7c590ebae60617c1ebcb8b7c"><code>40117de</code></a> fix(ci): pin Kafka 4.0.0 version in build</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.44.0...v1.45.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.7 to 1.32.8
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/31c2f3f89b98bd55ccb8833812d087baa9764e45"><code>31c2f3f</code></a> Release 2025-01-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ed70e6b14012a65186a1175d16bc60a665803d6e"><code>ed70e6b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5aef5b0eac63548e528afefe0431f9acd3b3a22d"><code>5aef5b0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e21e3f6b7c4adaa3db93457d03939c34b369ad8"><code>6e21e3f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90178248e923ae9c90df9b592e6392878c07a7f4"><code>9017824</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ebb7c024c114fe872f65f8a1d58ba49f0cc1a376"><code>ebb7c02</code></a> retry net.ErrClosed by default (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2949">#2949</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19d2a2833744bbc81aeca1896461c2d81e718d4e"><code>19d2a28</code></a> Release 2025-01-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e153a5993473dd666481855453b06dde48d66a7a"><code>e153a59</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/349cb942a7ae0cc5c42b764f39ac441115ccf7cf"><code>349cb94</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/740de30f11f73c77ef70aa722c0cf74fc876133d"><code>740de30</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.7...v1.32.8">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.48 to 1.17.51
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a7d202453f40928e19cc5485138f1683ee66156"><code>7a7d202</code></a> Release 2025-01-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fdaa7e26bf24a7ae4c6f9e66c3a582ed479d9be5"><code>fdaa7e2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/52aba54c34cfd8835897d6eb5e91630ace8656d8"><code>52aba54</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7263a7a06c599e38a282c6be548d0de354e9ad3b"><code>7263a7a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/31c2f3f89b98bd55ccb8833812d087baa9764e45"><code>31c2f3f</code></a> Release 2025-01-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ed70e6b14012a65186a1175d16bc60a665803d6e"><code>ed70e6b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5aef5b0eac63548e528afefe0431f9acd3b3a22d"><code>5aef5b0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e21e3f6b7c4adaa3db93457d03939c34b369ad8"><code>6e21e3f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90178248e923ae9c90df9b592e6392878c07a7f4"><code>9017824</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ebb7c024c114fe872f65f8a1d58ba49f0cc1a376"><code>ebb7c02</code></a> retry net.ErrClosed by default (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2949">#2949</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.48...credentials/v1.17.51">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.1 to 1.45.3
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-06-23)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/devopsguru</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/devopsguru/CHANGELOG.md#v1240-2023-06-23">v1.24.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for encryption via customer managed keys.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fsx</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/fsx/CHANGELOG.md#v1293-2023-06-23">v1.29.3</a>
<ul>
<li><strong>Documentation</strong>: Update to Amazon FSx documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/rds/CHANGELOG.md#v1453-2023-06-23">v1.45.3</a>
<ul>
<li><strong>Documentation</strong>: Documentation improvements for create, describe, and modify DB clusters and DB instances.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/verifiedpermissions</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/verifiedpermissions/CHANGELOG.md#v102-2023-06-23">v1.0.2</a>
<ul>
<li><strong>Documentation</strong>: Added improved descriptions and new code samples to SDK documentation.</li>
</ul>
</li>
</ul>
<h1>Release (2023-06-22)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkidentity</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/chimesdkidentity/CHANGELOG.md#v1120-2023-06-22">v1.12.0</a>
<ul>
<li><strong>Feature</strong>: AppInstanceBots can be configured to be invoked or not using the Target or the CHIME.mentions attribute for ChannelMessages</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/chimesdkmessaging/CHANGELOG.md#v1160-2023-06-22">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: ChannelMessages can be made visible to sender and intended recipient rather than all channel members with the target attribute. For example, a user can send messages to a bot and receive messages back in a group channel without other members seeing them.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kendra</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/kendra/CHANGELOG.md#v1410-2023-06-22">v1.41.0</a>
<ul>
<li><strong>Feature</strong>: Introducing Amazon Kendra Retrieve API that can be used to retrieve relevant passages or text excerpts given an input query.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sfn</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/sfn/CHANGELOG.md#v1180-2023-06-22">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for Versions and Aliases. Adds 8 operations: PublishStateMachineVersion, DeleteStateMachineVersion, ListStateMachineVersions, CreateStateMachineAlias, DescribeStateMachineAlias, UpdateStateMachineAlias, DeleteStateMachineAlias, ListStateMachineAliases</li>
</ul>
</li>
</ul>
<h1>Release (2023-06-21)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/dynamodb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/dynamodb/CHANGELOG.md#v11911-2023-06-21">v1.19.11</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for DynamoDB</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/emr</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/emr/CHANGELOG.md#v1270-2023-06-21">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces a new Amazon EMR EPI called ListSupportedInstanceTypes that returns a list of all instance types supported by a given EMR release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/inspector2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/inspector2/CHANGELOG.md#v1150-2023-06-21">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for Software Bill of Materials (SBOM) export and the general availability of code scanning for AWS Lambda functions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediaconvert</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/mediaconvert/CHANGELOG.md#v1380-2023-06-21">v1.38.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces the bandwidth reduction filter for the HEVC encoder, increases the limits of outputs per job, and updates support for the Nagra SDK to version 1.14.7.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mq</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/mq/CHANGELOG.md#v1150-2023-06-21">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: The Cross Region Disaster Recovery feature allows to replicate a brokers state from one region to another in order to provide customers with multi-region resiliency in the event of a regional outage.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/sagemaker/CHANGELOG.md#v1860-2023-06-21">v1.86.0</a>
<ul>
<li><strong>Feature</strong>: This release provides support in SageMaker for output files in training jobs to be uploaded without compression and enable customer to deploy uncompressed model from S3 to real-time inference Endpoints. In addition, ml.trn1n.32xlarge is added to supported instance type list in training job.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/transfer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/transfer/CHANGELOG.md#v1300-2023-06-21">v1.30.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new parameter StructuredLogDestinations to CreateServer, UpdateServer APIs.</li>
</ul>
</li>
</ul>
<h1>Release (2023-06-20)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.45.3/service/appflow/CHANGELOG.md#v1310-2023-06-20">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: This release adds new API to reset connector metadata cache</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ea71314ea4cf35a4ad53570483bb2fce29a7573d"><code>ea71314</code></a> Release 2023-06-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e9e10db1d1a59634753183ee55a6cdbc4ae31a64"><code>e9e10db</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51d1ac4c31cc6e450e95d4433504320bbe1e8bdd"><code>51d1ac4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/461e68fd42c15c814605fd831753e66383c9bcb5"><code>461e68f</code></a> Release 2023-06-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2103b226f7de1b5cd09bf3bff5ae3546738e48c4"><code>2103b22</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00ec9be203db56ec8586a711daa658c324f643a4"><code>00ec9be</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3104287aa0a6bc2f69d16faf25f7eb7a81ac83ff"><code>3104287</code></a> Release 2023-06-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e05231b21824b5edd4b9a0c0ea35711b7564319"><code>2e05231</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8d53d313b63099dd812cf3d4bfe44346a0e248cd"><code>8d53d31</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cea2987a76c5119b5e01a6dbc9a33aa8461d5828"><code>cea2987</code></a> Release 2023-06-20</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.45.1...service/rds/v1.45.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/lzap/cloudwatchwriter2` from 1.3.0 to 1.4.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/9c90ada9597e34d67f9c646780cb1f1d79c9f894"><code>9c90ada</code></a> fix: do not panic on close when client failed to initialize</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/2618aaf39a62c38df917d88bde177520b2d139ef"><code>2618aaf</code></a> feat: add Handler helper to avoid AWS dependencies</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/a971b1a1fa02aadd1548e2013fbc3f8ae0e0b902"><code>a971b1a</code></a> chore: update intriduction in README</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/ce355b5e4058237e602e76306277c64b07f9af5c"><code>ce355b5</code></a> fix: update examples in the README</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/65cbce8310410dd85a81632117dabf29d33474e0"><code>65cbce8</code></a> chore: update README with new example</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/cdd4c3ef77c1120546151367c01b455d4e19cad4"><code>cdd4c3e</code></a> feat: rewrite the queue to use Go channel</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/89d4a7c19aa9b2e2806381b2e21ec047c7ccb08a"><code>89d4a7c</code></a> fix: use context.Background()</li>
<li>See full diff in <a href="https://github.com/lzap/cloudwatchwriter2/compare/v1.3.0...v1.4.2">compare view</a></li>
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

### Comment by @jlsherrill on January 13, 2025 at 05:15 PM UTC

@dependabot rebase

---

## Reviews

### Review by @jlsherrill - Approved on January 14, 2025 at 05:29 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/937*
