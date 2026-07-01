---
type: pull_request
number: 620
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2024-04-01T04:41:41Z
updated: 2024-04-01T13:23:23Z
closed: 2024-04-01T13:23:16Z
merged: 2024-04-01T13:23:16Z
base_branch: main
head_branch: dependabot/go_modules/go-f5a5362599
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/620
---

# Pull Request #620: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: April 01, 2024 at 04:41 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-f5a5362599`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.43.0` | `1.43.1` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.25.3` | `1.26.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.7` | `1.17.10` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.34.3` | `1.35.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.3.1710446233` | `2024.3.1710791249` |

Updates `github.com/IBM/sarama` from 1.43.0 to 1.43.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.43.1 (2024-03-27)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: message.max.bytes should default to 1048576 not 1 MB by <a href="https://github.com/puellanivis"><code>@​puellanivis</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2804">IBM/sarama#2804</a></li>
<li>fix: add locking around broker throttle timer to prevent race condition by <a href="https://github.com/chengsha"><code>@​chengsha</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2826">IBM/sarama#2826</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump go.opentelemetry.io/otel/sdk from 1.23.1 to 1.24.0 in /examples/interceptors by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2816">IBM/sarama#2816</a></li>
<li>chore(deps): bump the golang-org-x group with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2825">IBM/sarama#2825</a></li>
<li>chore(deps): bump github.com/stretchr/testify from 1.8.4 to 1.9.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2822">IBM/sarama#2822</a></li>
<li>chore(deps): bump go.opentelemetry.io/otel/exporters/stdout/stdoutmetric from 1.23.1 to 1.24.0 in /examples/interceptors by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2815">IBM/sarama#2815</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/chengsha"><code>@​chengsha</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2826">IBM/sarama#2826</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.43.0...v1.43.1">https://github.com/IBM/sarama/compare/v1.43.0...v1.43.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/4ad35041300e1c15ba58b630745ac8eb05f30c10"><code>4ad3504</code></a> chore(ci): bump docker/setup-buildx-action from 3.0.0 to 3.2.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2838">#2838</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/837fab40795e8dfa834638110871e87a98684959"><code>837fab4</code></a> chore(ci): bump docker/bake-action (<a href="https://redirect.github.com/IBM/sarama/issues/2837">#2837</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/f21c5125746f9d10fd731dfdff54a494098626d1"><code>f21c512</code></a> chore(deps): bump go.opentelemetry.io/otel/exporters/stdout/stdoutmetric (<a href="https://redirect.github.com/IBM/sarama/issues/2815">#2815</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/392187dc6d40c1368c49a93ac0839df1c2573b7a"><code>392187d</code></a> chore(deps): bump github.com/stretchr/testify from 1.8.4 to 1.9.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2822">#2822</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/c432bc2a3daf92654cf7b9d1753d623f1cbef95b"><code>c432bc2</code></a> chore(deps): bump the golang-org-x group with 1 update (<a href="https://redirect.github.com/IBM/sarama/issues/2825">#2825</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e3681a94a82cb99ec890a65844337d0f674001af"><code>e3681a9</code></a> chore(ci): bump actions/dependency-review-action from 4.1.1 to 4.1.3 (<a href="https://redirect.github.com/IBM/sarama/issues/2818">#2818</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/0b515edf39e040706db569feb65b4ebd7ffb1aa3"><code>0b515ed</code></a> chore(deps): bump go.opentelemetry.io/otel/sdk in /examples/interceptors (<a href="https://redirect.github.com/IBM/sarama/issues/2816">#2816</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/ae04724deaeb4fec2a0ba0a20093c40ab8d18a0f"><code>ae04724</code></a> chore(ci): bump github/codeql-action from 3.24.3 to 3.24.6 (<a href="https://redirect.github.com/IBM/sarama/issues/2824">#2824</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a8b3b3d59f7f4e9200fbffc19da3b1eb9dfa4ca0"><code>a8b3b3d</code></a> fix: add locking around broker throttle timer to prevent race condition (<a href="https://redirect.github.com/IBM/sarama/issues/2826">#2826</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/fd84c2b0f0185100dbaec28ca4074289b35cc1b1"><code>fd84c2b</code></a> chore(ci): disable checkout credentials persist (<a href="https://redirect.github.com/IBM/sarama/issues/2812">#2812</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.43.0...v1.43.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.25.3 to 1.26.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0fde27cdffe0657695258e5d5220f7487117e71d"><code>0fde27c</code></a> Release 2024-03-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/57e0d95a95f10e285c38e27401de7da9b7bad559"><code>57e0d95</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e114db5c702d1103267da632ce985e3ce30b17fe"><code>e114db5</code></a> Update SDK's smithy-go dependency to v1.20.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f456f0784da02a3b21c24140a5426be035670352"><code>f456f07</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/96b431ae996732e64848cfa9c384db7ed76e84ef"><code>96b431a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6a694c7d3df3177fe4c4bbfff7e2e7a7733619bb"><code>6a694c7</code></a> dep: upgrade to smithy 1.47.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2587">#2587</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/973665bdd90b166296bf6aa4746fe7d4987dc5cb"><code>973665b</code></a> Release 2024-03-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8b24e40df3cc8d0b78361242cf8e9cc97abf823e"><code>8b24e40</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8788e04046ab19ca6c89a4d95401850455082662"><code>8788e04</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04803960abe0146af57291fa0f97fe347f667998"><code>0480396</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.25.3...v1.26.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.7 to 1.17.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/08f1f0b3e3d3f09b699c84f1f5b56b026fba6e15"><code>08f1f0b</code></a> Release 2022-10-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e1e20e6ca01d3acf5529dbfa059bda3b2ff5393"><code>0e1e20e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/281c268a255720163c78c38c97a92553fabf8f94"><code>281c268</code></a> Update SDK's smithy-go dependency to v1.13.4</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/db7c0a3fd1c72951a0673c13b6602b943285796c"><code>db7c0a3</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1eae80df178a5e3cd03b1cf04a6c7c9648e65e5a"><code>1eae80d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/17628c478a72ed2bc3596c4b7f24a49fa2251107"><code>17628c4</code></a> EC2 IMDS client logging fixes (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1891">#1891</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/765544941191825edd26162f9790bf11f059d426"><code>7655449</code></a> Release 2022-10-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dcae829ecc334f91502afd6d7ae2295861db9885"><code>dcae829</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b82766b858e595943b26924ad1f107cd04363d66"><code>b82766b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1c05fb6452a1f74985ff6deb7a642b9eb441274a"><code>1c05fb6</code></a> Implements IsCredentialsProvider for checking if a provider matches a target ...</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.17.7...config/v1.17.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.34.3 to 1.35.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-04-14)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/appflow/CHANGELOG.md#v1150-2022-04-14">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Enables users to pass custom token URL parameters for Oauth2 authentication during create connector profile</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appstream</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/appstream/CHANGELOG.md#v1160-2022-04-14">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Includes updates for create and update fleet APIs to manage the session scripts locations for Elastic fleets.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/batch</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/batch/CHANGELOG.md#v1180-2022-04-14">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Enables configuration updates for compute environments with BEST_FIT_PROGRESSIVE and SPOT_CAPACITY_OPTIMIZED allocation strategies.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudwatch</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/cloudwatch/CHANGELOG.md#v1181-2022-04-14">v1.18.1</a>
<ul>
<li><strong>Documentation</strong>: Updates documentation for additional statistics in CloudWatch Metric Streams.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/ec2/CHANGELOG.md#v1351-2022-04-14">v1.35.1</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Amazon EC2.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/glue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/glue/CHANGELOG.md#v1230-2022-04-14">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: Auto Scaling for Glue version 3.0 and later jobs to dynamically scale compute resources. This SDK change provides customers with the auto-scaled DPU usage</li>
</ul>
</li>
</ul>
<h1>Release (2022-04-13)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudwatch</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/cloudwatch/CHANGELOG.md#v1180-2022-04-13">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for additional statistics in CloudWatch Metric Streams.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fsx</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/fsx/CHANGELOG.md#v1230-2022-04-13">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for deploying FSx for ONTAP file systems in a single Availability Zone.</li>
</ul>
</li>
</ul>
<h1>Release (2022-04-12)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/devopsguru</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/devopsguru/CHANGELOG.md#v1170-2022-04-12">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: This release adds new APIs DeleteInsight to deletes the insight along with the associated anomalies, events and recommendations.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/ec2/CHANGELOG.md#v1350-2022-04-12">v1.35.0</a>
<ul>
<li><strong>Feature</strong>: X2idn and X2iedn instances are powered by 3rd generation Intel Xeon Scalable processors with an all-core turbo frequency up to 3.5 GHzAmazon EC2. C6a instances are powered by 3rd generation AMD EPYC processors.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/efs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/efs/CHANGELOG.md#v1170-2022-04-12">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Amazon EFS adds support for a ThrottlingException when using the CreateAccessPoint API if the account is nearing the AccessPoint limit(120).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iottwinmaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/iottwinmaker/CHANGELOG.md#v160-2022-04-12">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: This release adds the following new features: 1) ListEntities API now supports search using ExternalId. 2) BatchPutPropertyValue and GetPropertyValueHistory API now allows users to represent time in sub-second level precisions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kinesis</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/kinesis/CHANGELOG.md#v1154-2022-04-12">v1.15.4</a>
<ul>
<li><strong>Bug Fix</strong>: Fixes an issue that caused the unexported constructor function names for EventStream types to be swapped for the event reader and writer respectivly.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lexruntimev2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/lexruntimev2/CHANGELOG.md#v1144-2022-04-12">v1.14.4</a>
<ul>
<li><strong>Bug Fix</strong>: Fixes an issue that caused the unexported constructor function names for EventStream types to be swapped for the event reader and writer respectivly.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/s3/CHANGELOG.md#v1265-2022-04-12">v1.26.5</a>
<ul>
<li><strong>Bug Fix</strong>: Fixes an issue that caused the unexported constructor function names for EventStream types to be swapped for the event reader and writer respectivly.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/transcribestreaming</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ec2/v1.35.1/service/transcribestreaming/CHANGELOG.md#v164-2022-04-12">v1.6.4</a>
<ul>
<li><strong>Bug Fix</strong>: Fixes an issue that caused the unexported constructor function names for EventStream types to be swapped for the event reader and writer respectivly.</li>
</ul>
</li>
</ul>
<h1>Release (2022-04-11)</h1>
<h2>Module Highlights</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/481157636401f71a4dbf9f9014cfecea4e096599"><code>4811576</code></a> Release 2022-04-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c190249cc55e8153f6214652fe6d5114082d8fa0"><code>c190249</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a6a8baca33d7eaa974ed91030f137c7acf75848"><code>0a6a8ba</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06813a8c7c3f770b5d801e2a17354d37071bb343"><code>06813a8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/26414eed14ebdd4c4b47a215b6a5c001421570ca"><code>26414ee</code></a> local-mod-replace.sh: fix shebang (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1664">#1664</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35519bd702e38107c016588e4def765cea76f2e1"><code>35519bd</code></a> Release 2022-04-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d85fefc654014de1cd59a8e093be4ccf62d013f1"><code>d85fefc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fb245b6ceafc2032d397838aa9a97a16ed6c1c94"><code>fb245b6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59e437fd84461ec38aa8fb1c993b0f9bb4479eda"><code>59e437f</code></a> Release 2022-04-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e16bb72762ddfc7312bb1a9c505c91799bebf6de"><code>e16bb72</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/emr/v1.34.3...service/ec2/v1.35.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.3.1710446233 to 2024.3.1710791249
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/4320b8a352dff2d6cb1c0a6b98694cac943df796"><code>4320b8a</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e99624e58c87e8639db952b3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.3.1710446233...release/v2024.3.1710791249">compare view</a></li>
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

### Comment by @app-sre-bot on April 01, 2024 at 04:42 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on April 01, 2024 at 12:53 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on April 01, 2024 at 01:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/620*
