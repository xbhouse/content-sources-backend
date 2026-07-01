---
type: pull_request
number: 459
title: "Build: Bump the go group with 9 updates"
state: closed
author: dependabot
created: 2023-11-06T04:46:53Z
updated: 2023-11-13T04:25:52Z
closed: 2023-11-13T04:25:50Z
base_branch: main
head_branch: dependabot/go_modules/go-de361cfb57
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/459
---

# Pull Request #459: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: November 06, 2023 at 04:46 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-de361cfb57`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.6` | `1.0.7` |
| [github.com/go-playground/validator/v10](https://github.com/go-playground/validator) | `10.15.5` | `10.16.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.21.2` | `1.22.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.13.43` | `1.15.1` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.24.2` | `1.26.0` |
| [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) | `2023.10.1697742345` | `2023.10.1698674075` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.4.3` | `5.5.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.2.1` | `9.3.0` |
| [github.com/RedHatInsights/event-schemas-go](https://github.com/RedHatInsights/event-schemas-go) | `1.0.5` | `1.0.6` |

Updates `github.com/content-services/yummy` from 1.0.6 to 1.0.7
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/28b1d171d807caa031b5dfc456601c8d58b94d49"><code>28b1d17</code></a> Fixes 2442: add support for parsing groups and envs (<a href="https://redirect.github.com/content-services/yummy/issues/16">#16</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.6...v1.0.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/go-playground/validator/v10` from 10.15.5 to 10.16.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-playground/validator/releases">github.com/go-playground/validator/v10's releases</a>.</em></p>
<blockquote>
<h2>Release 10.16.0</h2>
<h2>What's new or fixed?</h2>
<ul>
<li>Added new <code>issn</code> validator. TY <a href="https://github.com/mrcook"><code>@​mrcook</code></a> for the <a href="https://redirect.github.com/go-playground/validator/pull/1166">PR</a></li>
<li>Added improvement/fix to the file url validation. TY <a href="https://github.com/nodivbyzero"><code>@​nodivbyzero</code></a> for the <a href="https://redirect.github.com/go-playground/validator/pull/1171">PR</a></li>
<li>Fix onof tag in examples. TY <a href="https://github.com/gren236"><code>@​gren236</code></a> for the <a href="https://redirect.github.com/go-playground/validator/pull/1184">PR</a></li>
<li>Add <code>fmt.Stringer</code> interface support to <code>uuid</code> validations allowing most UUID validation libraries which implement it to work transparently now. TY <a href="https://github.com/JoshGlazebrook"><code>@​JoshGlazebrook</code></a> for the <a href="https://redirect.github.com/go-playground/validator/pull/1189">PR</a></li>
<li>Added new <code>omitnil</code> similar to <code>omitempty</code> but for pointers. this is mainly for code generation ease. TY <a href="https://github.com/tarampampam"><code>@​tarampampam</code></a> for the <a href="https://redirect.github.com/go-playground/validator/pull/1187">PR</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-playground/validator/commit/84254aeb5a59e615ec0b66ab53b988bc0677f55e"><code>84254ae</code></a> update README</li>
<li><a href="https://github.com/go-playground/validator/commit/124662234ff9aa571e50e1fda6048a551bb78af3"><code>1246622</code></a> add ISSN validator + tests + documentation (<a href="https://redirect.github.com/go-playground/validator/issues/1166">#1166</a>)</li>
<li><a href="https://github.com/go-playground/validator/commit/c8569618d4489374973a0ee650972fcc786ed7cb"><code>c856961</code></a> fix file URL validator (<a href="https://redirect.github.com/go-playground/validator/issues/1171">#1171</a>)</li>
<li><a href="https://github.com/go-playground/validator/commit/aa969096c5644e7487fba0ef9768a6ac712a72fb"><code>aa96909</code></a> Add <code>omitnil</code> modifier (<a href="https://redirect.github.com/go-playground/validator/issues/1187">#1187</a>)</li>
<li><a href="https://github.com/go-playground/validator/commit/b0c733737d69470434d9f2358e6e83f6de60669a"><code>b0c7337</code></a> update fieldMatchesRegexByStringerValOrString for backward compatibility</li>
<li><a href="https://github.com/go-playground/validator/commit/4c1bd614e306e580fff06cd83334f30d8e35db8c"><code>4c1bd61</code></a> Add support for validating against uuid values that are structs which impleme...</li>
<li><a href="https://github.com/go-playground/validator/commit/adda84d6e65cd12b4ea52850e60698490d806a3f"><code>adda84d</code></a> Fix oneof tag in simple example (<a href="https://redirect.github.com/go-playground/validator/issues/1184">#1184</a>)</li>
<li>See full diff in <a href="https://github.com/go-playground/validator/compare/v10.15.5...v10.16.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.21.2 to 1.22.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ee5e3f05637540596cc7aab1359742000a8d533a"><code>ee5e3f0</code></a> Release 2023-11-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b65c226f47aa1f837699664bdc65c3c3e3611765"><code>b65c226</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a194b9b0344774a5af100d11ea2066c5b0cf234"><code>7a194b9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cb924a0007bc681d12f382a604368e0660827ee"><code>0cb924a</code></a> Add support for configured endpoints. (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2328">#2328</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/61039fea9cc9e080c53382850c87685b5406fd68"><code>61039fe</code></a> Release 2023-10-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/797e0560769725635218fc30a2554c1bbaccc01b"><code>797e056</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/822585d3f621a7c5844584d8e471c32f852702aa"><code>822585d</code></a> Update SDK's smithy-go dependency to v1.16.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abf753db747dd256f3ee69712a19d1d3dc681f23"><code>abf753d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/99861c071109ce5ee4f1cb3b72ead2062b3bd86c"><code>99861c0</code></a> lang: bump minimum go version to 1.19 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2338">#2338</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ac0a53ac45acaadc4526fd25b643dc46032b02a"><code>2ac0a53</code></a> Release 2023-10-30</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.21.2...v1.22.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.13.43 to 1.15.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-03-23)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.16.0
<ul>
<li><strong>Feature</strong>: Update CredentialsCache to make use of two new optional CredentialsProvider interfaces to give the cache, per provider, behavior how the cache handles credentials that fail to refresh, and adjusting expires time. See <a href="https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws#CredentialsCache">aws.CredentialsCache</a> for more details.</li>
<li><strong>Feature</strong>: Update <code>ec2rolecreds</code> package's <code>Provider</code> to implememnt support for CredentialsCache new optional caching strategy interfaces, HandleFailRefreshCredentialsCacheStrategy and AdjustExpiresByCredentialsCacheStrategy.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/credentials</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/credentials/CHANGELOG.md#v1110-2022-03-23">v1.11.0</a>
<ul>
<li><strong>Feature</strong>: Update <code>ec2rolecreds</code> package's <code>Provider</code> to implememnt support for CredentialsCache new optional caching strategy interfaces, HandleFailRefreshCredentialsCacheStrategy and AdjustExpiresByCredentialsCacheStrategy.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/auditmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/auditmanager/CHANGELOG.md#v1180-2022-03-23">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release updates 1 API parameter, the SnsArn attribute. The character length and regex pattern for the SnsArn attribute have been updated, which enables you to deselect an SNS topic when using the UpdateSettings operation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ebs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/ebs/CHANGELOG.md#v1150-2022-03-23">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Increased the maximum supported value for the Timeout parameter of the StartSnapshot API from 60 minutes to 4320 minutes.  Changed the HTTP error code for ConflictException from 503 to 409.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticache</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/elasticache/CHANGELOG.md#v1202-2022-03-23">v1.20.2</a>
<ul>
<li><strong>Documentation</strong>: Doc only update for ElastiCache</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/gamesparks</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/gamesparks/CHANGELOG.md#v100-2022-03-23">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: Released the preview of Amazon GameSparks, a fully managed AWS service that provides a multi-service backend for game developers.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/redshift</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/redshift/CHANGELOG.md#v1220-2022-03-23">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new [--encrypted | --no-encrypted] field in restore-from-cluster-snapshot API. Customers can now restore an unencrypted snapshot to a cluster encrypted with AWS Managed Key or their own KMS key.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ssm</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/ssm/CHANGELOG.md#v1230-2022-03-23">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: Update AddTagsToResource, ListTagsForResource, and RemoveTagsFromResource APIs to reflect the support for tagging Automation resources. Includes other minor documentation updates.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/transfer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/transfer/CHANGELOG.md#v1181-2022-03-23">v1.18.1</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for AWS Transfer Family to describe how to remove an associated workflow from a server.</li>
</ul>
</li>
</ul>
<h1>Release (2022-03-22)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/costexplorer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/costexplorer/CHANGELOG.md#v1180-2022-03-22">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Added three new APIs to support tagging and resource-level authorization on Cost Explorer resources: TagResource, UntagResource, ListTagsForResource.  Added optional parameters to CreateCostCategoryDefinition, CreateAnomalySubscription and CreateAnomalyMonitor APIs to support Tag On Create.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/ecs/CHANGELOG.md#v1182-2022-03-22">v1.18.2</a>
<ul>
<li><strong>Documentation</strong>: Documentation only update to address tickets</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lakeformation</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/lakeformation/CHANGELOG.md#v1160-2022-03-22">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: The release fixes the incorrect permissions called out in the documentation - DESCRIBE_TAG, ASSOCIATE_TAG, DELETE_TAG, ALTER_TAG. This trebuchet release fixes the corresponding SDK and documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/location</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/location/CHANGELOG.md#v1160-2022-03-22">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Location Service now includes a MaxResults parameter for GetDevicePositionHistory requests.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/polly</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/polly/CHANGELOG.md#v1140-2022-03-22">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Polly adds new Catalan voice - Arlet. Arlet is available as Neural voice only.</li>
</ul>
</li>
</ul>
<h1>Release (2022-03-21)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/chimesdkmeetings/CHANGELOG.md#v180-2022-03-21">v1.8.0</a>
<ul>
<li><strong>Feature</strong>: Add support for media replication to link multiple WebRTC media sessions together to reach larger and global audiences. Participants connected to a replica session can be granted access to join the primary session and can switch sessions with their existing WebRTC connection</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecr</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/ecr/CHANGELOG.md#v1170-2022-03-21">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: This release includes a fix in the DescribeImageScanFindings paginated output.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediaconnect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.15.1/service/mediaconnect/CHANGELOG.md#v1160-2022-03-21">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for selecting a maintenance window.</li>
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
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5399c50f8a2fc8b966d2368bfb1c8aaf546efefb"><code>5399c50</code></a> Release 2022-03-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/57ed4144aa9e38dfd1a8c41bbae9f4b4c2f56dac"><code>57ed414</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7797b2c8bb8e3393455231024242783a40d6a30"><code>e7797b2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a7068e374e813ff24326d09f2903bae0e3d57cb"><code>0a7068e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cbd1bab25397591da5983f6a598df23a8daa67c6"><code>cbd1bab</code></a> Update IMDS credential provider to handle credentials fail to refresh (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1634">#1634</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f83f630afb5c3f149058ae004c4f75a6acecc85d"><code>f83f630</code></a> Remove unneeded destructive integration tests for Amazon S3 Control API (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1637">#1637</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ccdd0e4fab2786009c8728b18298d7972b66a851"><code>ccdd0e4</code></a> Release 2022-03-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5b16b20bd80356d4a8dca1cb5454e96149cbc626"><code>5b16b20</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d273f3acd96f9b88b3d41cddff6c47b07cb51bb5"><code>d273f3a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d78523a009125fa2a74fcf3201db87ba42fe2322"><code>d78523a</code></a> Release 2022-03-21</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.13.43...config/v1.15.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.24.2 to 1.26.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-03-08)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Feature</strong>: Updated <code>github.com/aws/smithy-go</code> to latest version</li>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplify</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/amplify/CHANGELOG.md#v1110-2022-03-08">v1.11.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/amplifyuibuilder/CHANGELOG.md#v150-2022-03-08">v1.5.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/appflow/CHANGELOG.md#v1140-2022-03-08">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apprunner</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/apprunner/CHANGELOG.md#v1110-2022-03-08">v1.11.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/athena</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/athena/CHANGELOG.md#v1140-2022-03-08">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/braket</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/braket/CHANGELOG.md#v1150-2022-03-08">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/chimesdkmeetings/CHANGELOG.md#v170-2022-03-08">v1.7.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudtrail</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/cloudtrail/CHANGELOG.md#v1150-2022-03-08">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/connect/CHANGELOG.md#v1190-2022-03-08">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/devopsguru</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/devopsguru/CHANGELOG.md#v1160-2022-03-08">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/ec2/CHANGELOG.md#v1310-2022-03-08">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecr</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/ecr/CHANGELOG.md#v1160-2022-03-08">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/ecs/CHANGELOG.md#v1180-2022-03-08">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticache</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/elasticache/CHANGELOG.md#v1200-2022-03-08">v1.20.0</a>
<ul>
<li><strong>Documentation</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/finspacedata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/finspacedata/CHANGELOG.md#v1100-2022-03-08">v1.10.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fis</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/fis/CHANGELOG.md#v1120-2022-03-08">v1.12.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fsx</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/fsx/CHANGELOG.md#v1200-2022-03-08">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/gamelift</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/gamelift/CHANGELOG.md#v1140-2022-03-08">v1.14.0</a>
<ul>
<li><strong>Documentation</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/greengrassv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/greengrassv2/CHANGELOG.md#v1150-2022-03-08">v1.15.0</a>
<ul>
<li><strong>Documentation</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/internal/checksum</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/internal/checksum/CHANGELOG.md#v110-2022-03-08">v1.1.0</a>
<ul>
<li><strong>Feature</strong>:  Updates the SDK's checksum validation logic to require opt-in to output response payload validation. The SDK was always preforming output response payload checksum validation, not respecting the output validation model option. Fixes <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1606">#1606</a></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kafkaconnect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/kafkaconnect/CHANGELOG.md#v180-2022-03-08">v1.8.0</a>
<ul>
<li><strong>Feature</strong>: Updated service client model to latest release.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kendra</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.26.0/service/kendra/CHANGELOG.md#v1220-2022-03-08">v1.22.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4841e655a6d87ca9cc61221b3f434fb9e1880c09"><code>4841e65</code></a> Release 2022-03-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3b155ed5bbc44a03cfdb5545764f1d29a547fa01"><code>3b155ed</code></a> Updated smithy-go dependency version</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bb23000f36f9f1fcca6481cfef2ee6fa49fc5272"><code>bb23000</code></a> Update API Models (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1613">#1613</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d133d67c3da4011f6153b2f9dbbd3cd4fb480c47"><code>d133d67</code></a> service/internal/checksum: Fix handling of require checksum (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1612">#1612</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dfc3d8c71f208b0eb2b1dd23b750c72bdb1cb65b"><code>dfc3d8c</code></a> service/internal/checksum: Fix opt-in to checksum output validation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1607">#1607</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e6432821e1e212895376c492e01cffaa7f39262e"><code>e643282</code></a> aws/retry: Fix adaptive rate limit test (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1601">#1601</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0aa6c439e6e0d94ad5550986067a208c0e0ef3ed"><code>0aa6c43</code></a> Release 2022-02-24.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2195b593da3b68d2cc1be848bb2af7adcf2b5e99"><code>2195b59</code></a> Update SDK's API clients (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1602">#1602</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7653643543a19592b94a3db6d5bdc97621d19864"><code>7653643</code></a> Release 2022-02-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78b771ed3b1de4da39ff01ae73e30ad90e88522c"><code>78b771e</code></a> Update smithy-go dependency version</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.24.2...service/s3/v1.26.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2023` from 2023.10.1697742345 to 2023.10.1698674075
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/561f1a92f17bb124471ad02a5bb752e0252a0fe4"><code>561f1a9</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a8945b6e50b7a8d53bd49838...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2023.10.1697742345...release/v2023.10.1698674075">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.4.3 to 5.5.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.0 (November 4, 2023)</h1>
<ul>
<li>Add CollectExactlyOneRow. (Julien GOTTELAND)</li>
<li>Add OpenDBFromPool to create *database/sql.DB from *pgxpool.Pool. (Lev Zakharov)</li>
<li>Prepare can automatically choose statement name based on sql. This makes it easier to explicitly manage prepared statements.</li>
<li>Statement cache now uses deterministic, stable statement names.</li>
<li>database/sql prepared statement names are deterministically generated.</li>
<li>Fix: SendBatch wasn't respecting context cancellation.</li>
<li>Fix: Timeout error from pipeline is now normalized.</li>
<li>Fix: database/sql encoding json.RawMessage to []byte.</li>
<li>CancelRequest: Wait for the cancel request to be acknowledged by the server. This should improve PgBouncer compatibility. (Anton Levakin)</li>
<li>stdlib: Use Ping instead of CheckConn in ResetSession</li>
<li>Add json.Marshaler and json.Unmarshaler for Float4, Float8 (Kirill Mironov)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/96f5f9cd952eccb28ce575089953c8ba7d99d77b"><code>96f5f9c</code></a> Release v5.5.0</li>
<li><a href="https://github.com/jackc/pgx/commit/d3fb6e00da0c7bc7fce3b22a895a0fb72e39322e"><code>d3fb6e0</code></a> implement json.Marshaler and json.Unmarshaler for Float4, Float8</li>
<li><a href="https://github.com/jackc/pgx/commit/cf6ef75f916648857cf7b46322d6d7af7d372917"><code>cf6ef75</code></a> stdlib: Use Ping instead of CheckConn in ResetSession</li>
<li><a href="https://github.com/jackc/pgx/commit/7a4bb7edb5ff48b44734c8a7c29d3aa24b0ba432"><code>7a4bb7e</code></a> Add link to pgx presentation to README.md</li>
<li><a href="https://github.com/jackc/pgx/commit/6f7400f4282d174f272c9548c21dd9dd75fc4bec"><code>6f7400f</code></a> fix typo in the comment in the pgconn.go</li>
<li><a href="https://github.com/jackc/pgx/commit/304697de36f37f92f8ad7b6eeb5061ea66324d3d"><code>304697d</code></a> CancelRequest: Wait for the cancel request to be acknowledged by the server</li>
<li><a href="https://github.com/jackc/pgx/commit/5d0f904831a84cd6b7f17ef4539d2217ae3f8a66"><code>5d0f904</code></a> update TestConnContextCanceledCancelsRunningQueryOnServer</li>
<li><a href="https://github.com/jackc/pgx/commit/6ca3d8ed4e9c23cdd67f9bb9b90a4184f73ea1b8"><code>6ca3d8e</code></a> Revert &quot;CancelRequest: don't try to read the reply&quot;</li>
<li><a href="https://github.com/jackc/pgx/commit/81ddcfdefb391c7e80b0a9909682cf0c4ee19bf4"><code>81ddcfd</code></a> Fix spurious deadline exceeded error</li>
<li><a href="https://github.com/jackc/pgx/commit/45f807fdb4742e5d09cd1efde84b9ddd56ebcd9d"><code>45f807f</code></a> Special case the underlying type of []byte</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.4.3...v5.5.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.2.1 to 9.3.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.3.0</h2>
<h2>What's new?</h2>
<h3>JSON support</h3>
<p>We are continuing to add support for more <a href="https://redis.io/docs/data-types/">Redis data types</a> in Go-Redis. Today, we are happy to announce support for JSON. <a href="https://redis.io/docs/data-types/json/">JSON is a Redis data structure</a> for storing, querying, and manipulating a single JSON document.</p>
<p>With a JSON key in Redis, you can:</p>
<ul>
<li>Use it as a versatile hierarchical data type</li>
<li>Opt for it as a sophisticated alternative to the traditional hash data structure</li>
<li>Treat it as a singular document in a document-based database</li>
</ul>
<p>The following example demonstrate how to get started with JSON in Go-Redis:</p>
<pre lang="go"><code>var ctx = context.Background()
client := redis.NewClient(&amp;redis.Options{
	Addr: &quot;localhost:6379&quot;,
})
<p>type Bicycle struct {
Brand       string
Model       string
Price       int
}</p>
<p>bicycle := Bicycle{
Brand: &quot;Velorim&quot;,
Model: &quot;Jigger&quot;,
Price: 270,
}</p>
<p>_, err := client.JSONSet(ctx, &quot;bicycle:1&quot;, &quot;$&quot;, bicycle).Result()
if err != nil {
panic(err)
}</p>
<p>res, err := client.JSONGet(ctx, &quot;bicycle:1&quot;, &quot;.Model&quot;).Result()
if err != nil {
panic(err)
}
fmt.Println(&quot;bicycle:1 model is&quot;, res)
</code></pre></p>
<p><a href="https://redis.io/docs/data-types/json/">Learn more about JSON support in Redis</a></p>
<h3>Other notable improvements</h3>
<ul>
<li>Allow using pointers of simple data types as command values (<a href="https://redirect.github.com/redis/go-redis/issues/2745">#2745</a>) (<a href="https://redirect.github.com/redis/go-redis/issues/2753">#2753</a>)</li>
<li>Add InfoMap command (<a href="https://redirect.github.com/redis/go-redis/issues/2665">#2665</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/21bd40a47e56e61c0598ea1bdf8e02e67d1aa651"><code>21bd40a</code></a> Version 9.3.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2774">#2774</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/828fd2d3b83b813739ba56625fb637fe84f9c200"><code>828fd2d</code></a> chore(deps): bump google.golang.org/grpc in /example/otel (<a href="https://redirect.github.com/redis/go-redis/issues/2775">#2775</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/898bd9aa5155d4f0c83b7a7260d51d47affcd2e2"><code>898bd9a</code></a> chore(deps): bump golang.org/x/net in /example/otel (<a href="https://redirect.github.com/redis/go-redis/issues/2776">#2776</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/15682e3227597d12d7f76862b712a75fbb3d02d3"><code>15682e3</code></a> feat: support write the types of pointer of simple data types (<a href="https://redirect.github.com/redis/go-redis/issues/2745">#2745</a>) (<a href="https://redirect.github.com/redis/go-redis/issues/2753">#2753</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/81947daa8d6b2bea7e3a12c210b022f0a88d4006"><code>81947da</code></a> Handle wrapped errors in scripter.Run (<a href="https://redirect.github.com/redis/go-redis/issues/2674">#2674</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/84f46c330142e8ac08d12d9d68aef316cd272f73"><code>84f46c3</code></a> BUG: BFReserveArgs - error_rate &amp; capacity (<a href="https://redirect.github.com/redis/go-redis/issues/2763">#2763</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/343016bf72212bb22c07d3e56393ef78f252c397"><code>343016b</code></a> add InfoMap command (<a href="https://redirect.github.com/redis/go-redis/issues/2665">#2665</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/4408f8cfb2d4b3c49fbb5d17e6dfb85f04ce96c2"><code>4408f8c</code></a> free turn when leave with error (<a href="https://redirect.github.com/redis/go-redis/issues/2658">#2658</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7ebb537c804cf74322fe4910969e8af4f938683b"><code>7ebb537</code></a> remove duplicate declaration (<a href="https://redirect.github.com/redis/go-redis/issues/2773">#2773</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a5fe17472addc919cab2495559586540813e1864"><code>a5fe174</code></a> Option types must propagage missing fields (<a href="https://redirect.github.com/redis/go-redis/issues/2726">#2726</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.2.1...v9.3.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/event-schemas-go` from 1.0.5 to 1.0.6
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/event-schemas-go/releases">github.com/RedHatInsights/event-schemas-go's releases</a>.</em></p>
<blockquote>
<h2>v1.0.6</h2>
<h2><a href="https://github.com/RedHatInsights/event-schemas-go/compare/v1.0.5...v1.0.6">1.0.6</a> (2023-10-31)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>deps:</strong> bump api from <code>2a6f76d</code> to <code>f7bc074</code> (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/b6be569b6f79f45e20ab264d6448f00a895ccd53">b6be569</a>)</li>
<li><strong>deps:</strong> bump api from <code>a57120f</code> to <code>2a6f76d</code> (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/4a5043580ac4b254c827b1b27c8df33123394980">4a50435</a>)</li>
<li><strong>deps:</strong> bump api from <code>bb5a18a</code> to <code>a57120f</code> (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/cc0110834a2270148dc441c0fb703efdde6fef98">cc01108</a>)</li>
<li>Regenerate types (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/0272de44ea953a7f0332278e7138355bcd08b0be">0272de4</a>)</li>
<li>Regenerate types (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/5f0cbffe9eaa8b0425e17f7a36abea920b37c7d4">5f0cbff</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/event-schemas-go/blob/main/CHANGELOG.md">github.com/RedHatInsights/event-schemas-go's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/RedHatInsights/event-schemas-go/compare/v1.0.5...v1.0.6">1.0.6</a> (2023-10-31)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>deps:</strong> bump api from <code>2a6f76d</code> to <code>f7bc074</code> (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/b6be569b6f79f45e20ab264d6448f00a895ccd53">b6be569</a>)</li>
<li><strong>deps:</strong> bump api from <code>a57120f</code> to <code>2a6f76d</code> (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/4a5043580ac4b254c827b1b27c8df33123394980">4a50435</a>)</li>
<li><strong>deps:</strong> bump api from <code>bb5a18a</code> to <code>a57120f</code> (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/cc0110834a2270148dc441c0fb703efdde6fef98">cc01108</a>)</li>
<li>Regenerate types (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/0272de44ea953a7f0332278e7138355bcd08b0be">0272de4</a>)</li>
<li>Regenerate types (<a href="https://github.com/RedHatInsights/event-schemas-go/commit/5f0cbffe9eaa8b0425e17f7a36abea920b37c7d4">5f0cbff</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/43faf7f0db0f028398b5799e740836033448fa91"><code>43faf7f</code></a> chore(main): release 1.0.6 (<a href="https://redirect.github.com/RedHatInsights/event-schemas-go/issues/45">#45</a>)</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/0272de44ea953a7f0332278e7138355bcd08b0be"><code>0272de4</code></a> fix: Regenerate types</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/b6be569b6f79f45e20ab264d6448f00a895ccd53"><code>b6be569</code></a> fix(deps): bump api from <code>2a6f76d</code> to <code>f7bc074</code></li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/80295e52a6e6f825f8c34c4698f8ad2e1d005ea2"><code>80295e5</code></a> chore(deps): bump quicktype from 23.0.75 to 23.0.76</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/e98a89b116869de2d55e32b415e99d4b6224e509"><code>e98a89b</code></a> chore(deps): bump quicktype from 23.0.59 to 23.0.75</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/a5c82a2ecebe93bb75974564c94c3d13432a533a"><code>a5c82a2</code></a> chore(deps): bump quicktype from 23.0.54 to 23.0.59</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/4e52157e1db86d86503552a107e397258a86b931"><code>4e52157</code></a> chore(deps): bump quicktype from 23.0.49 to 23.0.54</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/514b711f4e322604268af6f867f0f314a73c3456"><code>514b711</code></a> chore(deps): bump quicktype from 23.0.48 to 23.0.49</li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/4a5043580ac4b254c827b1b27c8df33123394980"><code>4a50435</code></a> fix(deps): bump api from <code>a57120f</code> to <code>2a6f76d</code></li>
<li><a href="https://github.com/RedHatInsights/event-schemas-go/commit/5f0cbffe9eaa8b0425e17f7a36abea920b37c7d4"><code>5f0cbff</code></a> fix: Regenerate types</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/event-schemas-go/compare/v1.0.5...v1.0.6">compare view</a></li>
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

### Comment by @app-sre-bot on November 06, 2023 at 04:48 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on November 13, 2023 at 04:25 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/459*
