---
type: pull_request
number: 1274
title: "Build: Bump the go group with 8 updates"
state: merged
author: dependabot
created: 2025-11-03T04:13:04Z
updated: 2025-11-03T08:25:42Z
closed: 2025-11-03T08:25:33Z
merged: 2025-11-03T08:25:33Z
base_branch: main
head_branch: dependabot/go_modules/go-529b441a98
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1274
---

# Pull Request #1274: Build: Bump the go group with 8 updates

**Author**: @dependabot
**Created**: November 03, 2025 at 04:13 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-529b441a98`

## Description

Bumps the go group with 8 updates:

| Package | From | To |
| --- | --- | --- |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.31.0` | `1.31.1` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.39.4` | `1.39.5` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.15` | `1.31.16` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.19` | `1.18.20` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.58.5` | `1.58.6` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.88.7` | `1.89.1` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.10.1761380994` | `2025.10.1761938327` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.36.1` | `0.36.2` |

Updates `gorm.io/gorm` from 1.31.0 to 1.31.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.31.1</h2>
<h2>Changes</h2>
<ul>
<li>Add Namer-based column lookup to Schema.LookUpField <a href="https://github.com/cmmoran"><code>@​cmmoran</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7619">#7619</a>)</li>
<li>fix: Allow escaped double quotes in struct tag parser <a href="https://github.com/kankankanp"><code>@​kankankanp</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7631">#7631</a>)</li>
<li>Fix slog logger caller frame detection to output correct source file <a href="https://github.com/ifooth"><code>@​ifooth</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7610">#7610</a>)</li>
<li>chore(docs): edited the badge test <a href="https://github.com/Olexandr88"><code>@​Olexandr88</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7635">#7635</a>)</li>
<li>Fix AutoMigrate default value comparison for string fields (issue <a href="https://redirect.github.com/go-gorm/gorm/issues/7590">#7590</a>) <a href="https://github.com/nowindexman"><code>@​nowindexman</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7591">#7591</a>)</li>
<li>fix(UnixSecondSerializer.Value): Avoid panic when handling unsigned integer values <a href="https://github.com/dushaoshuai"><code>@​dushaoshuai</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7608">#7608</a>)</li>
<li>chore: fix some comments <a href="https://github.com/wyrapeseed"><code>@​wyrapeseed</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7615">#7615</a>)</li>
<li>Rename IsValidDBNameChar to IsInvalidDBNameChar <a href="https://github.com/mengxunQAQ"><code>@​mengxunQAQ</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7582">#7582</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/eabca1fa13dcf0acdd7d467300c4bd0414b1e0a2"><code>eabca1f</code></a> Allow Select/Omit for Generics Create, close <a href="https://redirect.github.com/go-gorm/gorm/issues/7638">#7638</a>, <a href="https://redirect.github.com/go-gorm/gorm/issues/7633">#7633</a></li>
<li><a href="https://github.com/go-gorm/gorm/commit/a57abbe12651e3d7a25ddae48ec8b37792053a0d"><code>a57abbe</code></a> Add Namer-based column lookup to Schema.LookUpField (<a href="https://redirect.github.com/go-gorm/gorm/issues/7619">#7619</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/5eaf05a6c9c4d64af3b54085bc4603a250fa5801"><code>5eaf05a</code></a> fix: Allow escaped double quotes in struct tag parser (<a href="https://redirect.github.com/go-gorm/gorm/issues/7631">#7631</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/2c3d109af3057f0d0ee64104011f12ac7bdab9bc"><code>2c3d109</code></a> Fix slog logger caller frame detection to output correct source file (<a href="https://redirect.github.com/go-gorm/gorm/issues/7610">#7610</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/4808ff510c279685eeed189740abd69b8a6dd447"><code>4808ff5</code></a> Update README.md (<a href="https://redirect.github.com/go-gorm/gorm/issues/7635">#7635</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/141388f28b58e6d9c0422cbb5bf4a5f22067610d"><code>141388f</code></a> Fix AutoMigrate default value comparison for string fields (issue <a href="https://redirect.github.com/go-gorm/gorm/issues/7590">#7590</a>) (<a href="https://redirect.github.com/go-gorm/gorm/issues/7591">#7591</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/d9372f551000bef6535f5679c38255ca0be327bc"><code>d9372f5</code></a> fix(UnixSecondSerializer.Value): Avoid panic when handling unsigned integer v...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/d8cdb399566b65270d5a3b651eca0195694f35d5"><code>d8cdb39</code></a> chore: fix some comment (<a href="https://redirect.github.com/go-gorm/gorm/issues/7615">#7615</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/b88148363a954f69fa680b152dfd96a94ffea1e1"><code>b881483</code></a> Rename IsValidDBNameChar to IsInvalidDBNameChar (<a href="https://redirect.github.com/go-gorm/gorm/issues/7582">#7582</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.31.0...v1.31.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.39.4 to 1.39.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d8ed081e9bcc22e6c5eb63fb6bbacfa38d7bcce3"><code>d8ed081</code></a> Release 2025-10-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2c9cb47051246e43a65f78696b8189aaa48c871"><code>a2c9cb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7af054b46fade69938d2682163c4abbbf126b9c7"><code>7af054b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e5ed49bd27c797319223fe6454411c6dc3f62cc"><code>2e5ed49</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3a3b6c778eed688dc91940a6d9160e7445a218f"><code>f3a3b6c</code></a> remove arbitrary response read timeout in kinesis GetRecords (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3221">#3221</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56af15521102c665ba33e5512cd66068c7c32506"><code>56af155</code></a> Release 2025-10-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b23832ac9b9505202125cf5c448c72cd333c819b"><code>b23832a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/33ea965f3eb459c22ba41a8e5da55115da5686bc"><code>33ea965</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f70c7889fa7e09fba31a68c6dbef1e178bbb9964"><code>f70c788</code></a> Release 2025-10-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f8aa1f373d60ee159a2aac412311edad2eaad752"><code>f8aa1f3</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.39.4...v1.39.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.15 to 1.31.16
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d8ed081e9bcc22e6c5eb63fb6bbacfa38d7bcce3"><code>d8ed081</code></a> Release 2025-10-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2c9cb47051246e43a65f78696b8189aaa48c871"><code>a2c9cb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7af054b46fade69938d2682163c4abbbf126b9c7"><code>7af054b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e5ed49bd27c797319223fe6454411c6dc3f62cc"><code>2e5ed49</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3a3b6c778eed688dc91940a6d9160e7445a218f"><code>f3a3b6c</code></a> remove arbitrary response read timeout in kinesis GetRecords (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3221">#3221</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56af15521102c665ba33e5512cd66068c7c32506"><code>56af155</code></a> Release 2025-10-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b23832ac9b9505202125cf5c448c72cd333c819b"><code>b23832a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/33ea965f3eb459c22ba41a8e5da55115da5686bc"><code>33ea965</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f70c7889fa7e09fba31a68c6dbef1e178bbb9964"><code>f70c788</code></a> Release 2025-10-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f8aa1f373d60ee159a2aac412311edad2eaad752"><code>f8aa1f3</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.31.15...config/v1.31.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.19 to 1.18.20
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-04-07)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/dlm</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/dlm/CHANGELOG.md#v1150-2023-04-07">v1.15.0</a>
<ul>
<li><strong>Announcement</strong>: This release includes breaking changes for the timestamp trait on the data lifecycle management client.</li>
<li><strong>Feature</strong>: Updated timestamp format for GetLifecyclePolicy API</li>
<li><strong>Bug Fix</strong>: Correct timestamp type for data lifecycle manager.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/docdb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/docdb/CHANGELOG.md#v1210-2023-04-07">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new parameter 'DBClusterParameterGroupName' to 'RestoreDBClusterFromSnapshot' API to associate the name of the DB cluster parameter group while performing restore.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fsx</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/fsx/CHANGELOG.md#v1288-2023-04-07">v1.28.8</a>
<ul>
<li><strong>Documentation</strong>: Amazon FSx for Lustre now supports creating data repository associations on Persistent_1 and Scratch_2 file systems.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lambda</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/lambda/CHANGELOG.md#v1310-2023-04-07">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new Lambda InvokeWithResponseStream API to support streaming Lambda function responses. The release also adds a new InvokeMode parameter to Function Url APIs to control whether the response will be streamed or buffered.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/quicksight</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/quicksight/CHANGELOG.md#v1340-2023-04-07">v1.34.0</a>
<ul>
<li><strong>Feature</strong>: This release has two changes: adding the OR condition to tag-based RLS rules in CreateDataSet and UpdateDataSet; adding RefreshSchedule and Incremental RefreshProperties operations for users to programmatically configure SPICE dataset ingestions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/redshiftdata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/redshiftdata/CHANGELOG.md#v1193-2023-04-07">v1.19.3</a>
<ul>
<li><strong>Documentation</strong>: Update documentation of API descriptions as needed in support of temporary credentials with IAM identity.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/servicecatalog</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/servicecatalog/CHANGELOG.md#v1181-2023-04-07">v1.18.1</a>
<ul>
<li><strong>Documentation</strong>: Updates description for property</li>
</ul>
</li>
</ul>
<h1>Release (2023-04-06)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudformation</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/cloudformation/CHANGELOG.md#v1270-2023-04-06">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: Including UPDATE_COMPLETE as a failed status for DeleteStack waiter.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/greengrassv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/greengrassv2/CHANGELOG.md#v1220-2023-04-06">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: Add support for SUCCEEDED value in coreDeviceExecutionStatus field. Documentation updates for Greengrass V2.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/proton</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/proton/CHANGELOG.md#v1210-2023-04-06">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for the AWS Proton service sync feature. Service sync enables managing an AWS Proton service (creating and updating instances) and all of it's corresponding service instances from a Git repository.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/rds/CHANGELOG.md#v1421-2023-04-06">v1.42.1</a>
<ul>
<li><strong>Documentation</strong>: Adds and updates the SDK examples</li>
</ul>
</li>
</ul>
<h1>Release (2023-04-05)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/configservice/CHANGELOG.md#v1310-2023-04-05">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: This release adds resourceType enums for types released in March 2023.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/ecs/CHANGELOG.md#v1243-2023-04-05">v1.24.3</a>
<ul>
<li><strong>Documentation</strong>: This is a document only updated to add information about Amazon Elastic Inference (EI).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/identitystore</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/identitystore/CHANGELOG.md#v1167-2023-04-05">v1.16.7</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Identity Store CLI command reference.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ivsrealtime</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/ivsrealtime/CHANGELOG.md#v110-2023-04-05">v1.1.0</a>
<ul>
<li><strong>Feature</strong>: Fix ParticipantToken ExpirationTime format</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/networkfirewall</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/networkfirewall/CHANGELOG.md#v1260-2023-04-05">v1.26.0</a>
<ul>
<li><strong>Feature</strong>: AWS Network Firewall now supports IPv6-only subnets.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/servicecatalog</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.20/service/servicecatalog/CHANGELOG.md#v1180-2023-04-05">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: removed incorrect product type value</li>
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
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/30383d567f67c2a67b2b40a462a8c284c49d1796"><code>30383d5</code></a> Release 2023-04-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/352f89c2d23ec6249a699c732ba5c9ae050f833f"><code>352f89c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50429391777bf4f0d6229e7adfb11c4e59f0f99c"><code>5042939</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d40a16e950d99852243ba56e04536237483b92fa"><code>d40a16e</code></a> NXDOMAIN errors should not be retried (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2083">#2083</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/439f88c4e09f137340366d509e296924e297d978"><code>439f88c</code></a> Add announcement for next release for dlm</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4036a9c3b835693a3f6bc4d9f006bdadb995f26"><code>e4036a9</code></a> Release 2023-04-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1c455e29efcc1bf6681db35faf3455bc3171da81"><code>1c455e2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/adb09d335d690060cf1fcb2a5390543268b3f1db"><code>adb09d3</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3d4ed442f89b68918b77b9322ac876b3a9880660"><code>3d4ed44</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/296e00587eb442c06d03e095daa61d3e9b1c80ed"><code>296e005</code></a> Release 2023-04-05</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.19...config/v1.18.20">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.58.5 to 1.58.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d8ed081e9bcc22e6c5eb63fb6bbacfa38d7bcce3"><code>d8ed081</code></a> Release 2025-10-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2c9cb47051246e43a65f78696b8189aaa48c871"><code>a2c9cb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7af054b46fade69938d2682163c4abbbf126b9c7"><code>7af054b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e5ed49bd27c797319223fe6454411c6dc3f62cc"><code>2e5ed49</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3a3b6c778eed688dc91940a6d9160e7445a218f"><code>f3a3b6c</code></a> remove arbitrary response read timeout in kinesis GetRecords (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3221">#3221</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56af15521102c665ba33e5512cd66068c7c32506"><code>56af155</code></a> Release 2025-10-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b23832ac9b9505202125cf5c448c72cd333c819b"><code>b23832a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/33ea965f3eb459c22ba41a8e5da55115da5686bc"><code>33ea965</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f70c7889fa7e09fba31a68c6dbef1e178bbb9964"><code>f70c788</code></a> Release 2025-10-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f8aa1f373d60ee159a2aac412311edad2eaad752"><code>f8aa1f3</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/route53/v1.58.5...service/cloudwatchlogs/v1.58.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.88.7 to 1.89.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d8ed081e9bcc22e6c5eb63fb6bbacfa38d7bcce3"><code>d8ed081</code></a> Release 2025-10-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2c9cb47051246e43a65f78696b8189aaa48c871"><code>a2c9cb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7af054b46fade69938d2682163c4abbbf126b9c7"><code>7af054b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e5ed49bd27c797319223fe6454411c6dc3f62cc"><code>2e5ed49</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3a3b6c778eed688dc91940a6d9160e7445a218f"><code>f3a3b6c</code></a> remove arbitrary response read timeout in kinesis GetRecords (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3221">#3221</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56af15521102c665ba33e5512cd66068c7c32506"><code>56af155</code></a> Release 2025-10-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b23832ac9b9505202125cf5c448c72cd333c819b"><code>b23832a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/33ea965f3eb459c22ba41a8e5da55115da5686bc"><code>33ea965</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f70c7889fa7e09fba31a68c6dbef1e178bbb9964"><code>f70c788</code></a> Release 2025-10-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f8aa1f373d60ee159a2aac412311edad2eaad752"><code>f8aa1f3</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.88.7...service/s3/v1.89.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.10.1761380994 to 2025.10.1761938327
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/894cbacad89f9922cc7aa0ede2e60d78f9a82bbd"><code>894cbac</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e4ea4b33d5f87e8639db952b...</li>
<li><a href="https://github.com/content-services/zest/commit/fde6d8fc46247fa0621125963f0347164753dffa"><code>fde6d8f</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97925a9e246b137ae5486b92ad...</li>
<li><a href="https://github.com/content-services/zest/commit/231db75d3a36d52502784dbabef8051ebe020bbc"><code>231db75</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d5343ae629057952eab83695...</li>
<li><a href="https://github.com/content-services/zest/commit/64d9ffedd48b5cd933b02dc8c4002e5a42e733a8"><code>64d9ffe</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4783548429b5167968b5e2da9a...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.10.1761380994...release/v2025.10.1761938327">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.36.1 to 0.36.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.36.2</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.36.2.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Fix context propagation for logs to ensure logger instances correctly inherit span and hub information from their creation context (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1118">#1118</a>)
<ul>
<li>Logs now properly propagate trace context from the logger's original context, even when emitted in a different context</li>
<li>The logger will first check the emission context, then fall back to its creation context, and finally to the current hub</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.36.2</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.36.2.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Fix context propagation for logs to ensure logger instances correctly inherit span and hub information from their creation context (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1118">#1118</a>)
<ul>
<li>Logs now properly propagate trace context from the logger's original context, even when emitted in a different context</li>
<li>The logger will first check the emission context, then fall back to its creation context, and finally to the current hub</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/2aecd6ed8f56c837a4d4e6ad6cef8069ede88fec"><code>2aecd6e</code></a> release: 0.36.2</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a30c2b8d1ae64601a0f8652bd0b1d02e23ab7b7d"><code>a30c2b8</code></a> Prepare 0.36.2 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1119">#1119</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/97e65bfdd61eeae5034afba1f772d7501c6c2c67"><code>97e65bf</code></a> fix: add correct context propagation for logs (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1118">#1118</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c982e6f7a63b36c976bc94cff84db597172eb3c2"><code>c982e6f</code></a> feat: add new envelope transport (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1094">#1094</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/10448bc1d8a15179042d12b135d9977b186b919d"><code>10448bc</code></a> feat: add ring buffers (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1093">#1093</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/839f710a0bcebc4d72f6d74f591d959a574c39be"><code>839f710</code></a> Merge branch 'release/0.36.1'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.36.1...v0.36.2">compare view</a></li>
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

### Review by @swadeley - Approved on November 03, 2025 at 08:25 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1274*
