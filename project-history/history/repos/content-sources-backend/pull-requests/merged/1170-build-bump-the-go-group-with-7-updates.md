---
type: pull_request
number: 1170
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2025-08-11T07:12:08Z
updated: 2025-08-14T19:40:04Z
closed: 2025-08-14T19:39:56Z
merged: 2025-08-14T19:39:56Z
base_branch: main
head_branch: dependabot/go_modules/go-04d5aa339f
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1170
---

# Pull Request #1170: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: August 11, 2025 at 07:12 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-04d5aa339f`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.37.1` | `1.37.2` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.30.2` | `1.30.3` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.2` | `1.18.3` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.54.1` | `1.55.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.85.1` | `1.86.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.11.0` | `9.12.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.26.0` | `1.27.0` |

Updates `github.com/aws/aws-sdk-go-v2` from 1.37.1 to 1.37.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/711cf64d8ec7c903a234c1b84f63e94014de8622"><code>711cf64</code></a> Release 2025-08-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/699889ec855eabb003038d08739bf30f62aa6beb"><code>699889e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a94577c6f8df79c704c641a5309ebbbfa34b65a"><code>5a94577</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/240e984f9408bd5f6c6fbe0075958539e5c6cdb3"><code>240e984</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e08461090ccba679456c05264e2c04bf228138e"><code>2e08461</code></a> support configurable auth scheme preference (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3153">#3153</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5462033e242599cd94c01b98f6d179a0d95ddcbf"><code>5462033</code></a> Release 2025-08-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a84b1c67e1f56dc279bc9224259ae571d03575fe"><code>a84b1c6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3325d43c3e22a41bf8fff831ce10889e584540c9"><code>3325d43</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ddc4289d0a33d61898b6ab3a2e3953ad7495746"><code>2ddc428</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb0a91b264d3c916f1fa4e4d0c113e72e7330866"><code>eb0a91b</code></a> Release 2025-07-31</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.37.1...v1.37.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.30.2 to 1.30.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0599931efcf551dc63f84ec669d7fb5cfe14f64c"><code>0599931</code></a> Release 2024-07-10.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ca2a30ec8d823300ec948b784c491904eeba8929"><code>ca2a30e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2dff5f9bddaf1f23ea57e1e04d80f6ce2d54221"><code>e2dff5f</code></a> temporarily re-add jmespath dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2712">#2712</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.2...v1.30.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.2 to 1.18.3
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-11-22)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/appflow/CHANGELOG.md#v1210-2022-11-22">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: Adding support for Amazon AppFlow to transfer the data to Amazon Redshift databases through Amazon Redshift Data API service. This feature will support the Redshift destination connector on both public and private accessible Amazon Redshift Clusters and Amazon Redshift Serverless.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/kinesisanalyticsv2/CHANGELOG.md#v1150-2022-11-22">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Support for Apache Flink 1.15 in Kinesis Data Analytics.</li>
</ul>
</li>
</ul>
<h1>Release (2022-11-21)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/route53</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/route53/CHANGELOG.md#v1250-2022-11-21">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Route 53 now supports the Asia Pacific (Hyderabad) Region (ap-south-2) for latency records, geoproximity records, and private DNS for Amazon VPCs in that region.</li>
</ul>
</li>
</ul>
<h1>Release (2022-11-18.2)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/ssmsap</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/ssmsap/CHANGELOG.md#v101-2022-11-182">v1.0.1</a>
<ul>
<li><strong>Bug Fix</strong>: Removes old model file for ssm sap and uses the new model file to regenerate client</li>
</ul>
</li>
</ul>
<h1>Release (2022-11-18)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/appflow/CHANGELOG.md#v1200-2022-11-18">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: AppFlow provides a new API called UpdateConnectorRegistration to update a custom connector that customers have previously registered. With this API, customers no longer need to unregister and then register a connector to make an update.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/auditmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/auditmanager/CHANGELOG.md#v1210-2022-11-18">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces a new feature for Audit Manager: Evidence finder. You can now use evidence finder to quickly query your evidence, and add the matching evidence results to an assessment report.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkvoice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/chimesdkvoice/CHANGELOG.md#v100-2022-11-18">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: Amazon Chime Voice Connector, Voice Connector Group and PSTN Audio Service APIs are now available in the Amazon Chime SDK Voice namespace. See <a href="https://docs.aws.amazon.com/chime-sdk/latest/dg/sdk-available-regions.html">https://docs.aws.amazon.com/chime-sdk/latest/dg/sdk-available-regions.html</a> for more details.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudfront</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/cloudfront/CHANGELOG.md#v1210-2022-11-18">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: CloudFront API support for staging distributions and associated traffic management policies.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/connect/CHANGELOG.md#v1380-2022-11-18">v1.38.0</a>
<ul>
<li><strong>Feature</strong>: Added AllowedAccessControlTags and TagRestrictedResource for Tag Based Access Control on Amazon Connect Webpage</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/dynamodb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/dynamodb/CHANGELOG.md#v1176-2022-11-18">v1.17.6</a>
<ul>
<li><strong>Documentation</strong>: Updated minor fixes for DynamoDB documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/dynamodbstreams</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/dynamodbstreams/CHANGELOG.md#v11325-2022-11-18">v1.13.25</a>
<ul>
<li><strong>Documentation</strong>: Updated minor fixes for DynamoDB documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/ec2/CHANGELOG.md#v1720-2022-11-18">v1.72.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for copying an Amazon Machine Image's tags when copying an AMI.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/glue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/glue/CHANGELOG.md#v1350-2022-11-18">v1.35.0</a>
<ul>
<li><strong>Feature</strong>: AWSGlue Crawler - Adding support for Table and Column level Comments with database level datatypes for JDBC based crawler.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotroborunner</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.3/service/iotroborunner/CHANGELOG.md#v100-2022-11-18">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
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
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a32d707aff708a14e7e39927930a1fb2a3bd6f7"><code>7a32d70</code></a> Release 2022-11-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0242bd53d7d653d5af25cb046b4e5cdd9085388d"><code>0242bd5</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/350cc3eb4e1b8bf3e3bd731deb5b43b092d3ac54"><code>350cc3e</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a227ddc7773d739b0cc40d576303d13415c49fbc"><code>a227ddc</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bec5c47e6039349d12ba39e33261630f90dfdd13"><code>bec5c47</code></a> Release 2022-11-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0438d8f026b481eb3217b2d0012b28832ec529d0"><code>0438d8f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1ec853e23add4a10d313e23b08848ace880c1129"><code>1ec853e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b30f43183cc4c83d8ab6bde0c28256999a442c4b"><code>b30f431</code></a> Release 2022-11-18.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4fadedce965266b4132b5f9040ede78897e40edd"><code>4fadedc</code></a> Adding changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/14ebfbd6f287d16eb1afdad9e1248bd108e70d56"><code>14ebfbd</code></a> Remove old model files for ssm map.</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.2...config/v1.18.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.54.1 to 1.55.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54e19df9a118d5510d2d68e743b536dba4ff5d14"><code>54e19df</code></a> Release 2024-06-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0402b496189430bd02ce2dfaa4f172e57f967547"><code>0402b49</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a185f3b0eaa99efa5fb702c6e100de8fa4a06b15"><code>a185f3b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fb6840b93c8a9ec1fffb601a42bab10fedf381b0"><code>fb6840b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/882127d3c9fdd1caff1cdc82551a108b0b00bbd0"><code>882127d</code></a> add codec option to use encoding.Text/Binary(Un)Marshaler when present (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2666">#2666</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5b30240c25ce7b91bce93e01ce88471aa8f41c20"><code>5b30240</code></a> add s3 protocol tests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2667">#2667</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4b0b7efe9ee93c5c64939f7dfd75c821484fa36c"><code>4b0b7ef</code></a> Release 2024-06-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de2612e95a9fb7b54f056dcdcf176cf9edc8f57b"><code>de2612e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cea73de490eeb6e873c3c67fd0ac71064cbd8fbd"><code>cea73de</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a17e5c43549a82f931bfbce884224ade6c5bfcd1"><code>a17e5c4</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.54.1...service/s3/v1.55.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.85.1 to 1.86.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/711cf64d8ec7c903a234c1b84f63e94014de8622"><code>711cf64</code></a> Release 2025-08-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/699889ec855eabb003038d08739bf30f62aa6beb"><code>699889e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a94577c6f8df79c704c641a5309ebbbfa34b65a"><code>5a94577</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/240e984f9408bd5f6c6fbe0075958539e5c6cdb3"><code>240e984</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e08461090ccba679456c05264e2c04bf228138e"><code>2e08461</code></a> support configurable auth scheme preference (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3153">#3153</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5462033e242599cd94c01b98f6d179a0d95ddcbf"><code>5462033</code></a> Release 2025-08-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a84b1c67e1f56dc279bc9224259ae571d03575fe"><code>a84b1c6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3325d43c3e22a41bf8fff831ce10889e584540c9"><code>3325d43</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ddc4289d0a33d61898b6ab3a2e3953ad7495746"><code>2ddc428</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb0a91b264d3c916f1fa4e4d0c113e72e7330866"><code>eb0a91b</code></a> Release 2025-07-31</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.85.1...service/s3/v1.86.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.11.0 to 9.12.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.12.0</h2>
<h2>🚀 Highlights</h2>
<ul>
<li>This release includes support for <a href="https://redis.io/docs/latest/operate/oss_and_stack/stack-with-enterprise/release-notes/redisce/redisos-8.2-release-notes/">Redis 8.2</a>.</li>
<li>Introduces an experimental Query Builders for <code>FTSearch</code>, <code>FTAggregate</code> and other search commands.</li>
<li>Adds support for <code>EPSILON</code> option in <code>FT.VSIM</code>.</li>
<li>Includes bug fixes and improvements contributed by the community related to ring and <a href="https://github.com/redis/go-redis/tree/master/extra/redisotel">redisotel</a>.</li>
</ul>
<h2>Changes</h2>
<ul>
<li>Improve stale issue workflow (<a href="https://redirect.github.com/redis/go-redis/pull/3458">#3458</a>)</li>
<li>chore(ci): Add 8.2 rc2 pre build for CI (<a href="https://redirect.github.com/redis/go-redis/pull/3459">#3459</a>)</li>
<li>Added new stream commands (<a href="https://redirect.github.com/redis/go-redis/pull/3450">#3450</a>)</li>
<li>feat: Add &quot;skip_verify&quot; to Sentinel (<a href="https://redirect.github.com/redis/go-redis/pull/3428">#3428</a>)</li>
<li>fix: <code>errors.Join</code> requires Go 1.20 or later (<a href="https://redirect.github.com/redis/go-redis/pull/3442">#3442</a>)</li>
<li>DOC-4344 document quickstart examples (<a href="https://redirect.github.com/redis/go-redis/pull/3426">#3426</a>)</li>
<li>feat(bitop): add support for the new bitop operations (<a href="https://redirect.github.com/redis/go-redis/pull/3409">#3409</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>feat: recover addIdleConn may occur panic (<a href="https://redirect.github.com/redis/go-redis/pull/2445">#2445</a>)</li>
<li>feat(ring): specify custom health check func via HeartbeatFn option (<a href="https://redirect.github.com/redis/go-redis/pull/2940">#2940</a>)</li>
<li>Add Query Builder for RediSearch commands (<a href="https://redirect.github.com/redis/go-redis/pull/3436">#3436</a>)</li>
<li>add configurable buffer sizes for Redis connections (<a href="https://redirect.github.com/redis/go-redis/pull/3453">#3453</a>)</li>
<li>Add VAMANA vector type to RediSearch (<a href="https://redirect.github.com/redis/go-redis/pull/3449">#3449</a>)</li>
<li>VSIM add <code>EPSILON</code> option (<a href="https://redirect.github.com/redis/go-redis/pull/3454">#3454</a>)</li>
<li>Add closing support to otel metrics instrumentation (<a href="https://redirect.github.com/redis/go-redis/pull/3444">#3444</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix(redisotel): fix buggy append in reportPoolStats (<a href="https://redirect.github.com/redis/go-redis/pull/3122">#3122</a>)</li>
<li>fix(search): return results even if doc is empty (<a href="https://redirect.github.com/redis/go-redis/pull/3457">#3457</a>)</li>
<li>[ISSUE-3402]: Ring.Pipelined return dial timeout error (<a href="https://redirect.github.com/redis/go-redis/pull/3403">#3403</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>Merges stale issues jobs into one job with two steps (<a href="https://redirect.github.com/redis/go-redis/pull/3463">#3463</a>)</li>
<li>improve code readability (<a href="https://redirect.github.com/redis/go-redis/pull/3446">#3446</a>)</li>
<li>chore(release): 9.12.0-beta.1 (<a href="https://redirect.github.com/redis/go-redis/pull/3460">#3460</a>)</li>
<li>DOC-5472 time series doc examples (<a href="https://redirect.github.com/redis/go-redis/pull/3443">#3443</a>)</li>
<li>Add VAMANA compression algorithm tests (<a href="https://redirect.github.com/redis/go-redis/pull/3461">#3461</a>)</li>
<li>bumped redis 8.2 version used in the CI/CD (<a href="https://redirect.github.com/redis/go-redis/pull/3451">#3451</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/andy-stark-redis"><code>@​andy-stark-redis</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/jouir"><code>@​jouir</code></a>, <a href="https://github.com/monkey92t"><code>@​monkey92t</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/rokn"><code>@​rokn</code></a>, <a href="https://github.com/smnvdev"><code>@​smnvdev</code></a>, <a href="https://github.com/strobil"><code>@​strobil</code></a> and <a href="https://github.com/wzy9607"><code>@​wzy9607</code></a></p>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a> made their first contribution in <a href="https://redirect.github.com/redis/go-redis/pull/3409">#3409</a></li>
<li><a href="https://github.com/smnvdev"><code>@​smnvdev</code></a> made their first contribution in <a href="https://redirect.github.com/redis/go-redis/pull/3403">#3403</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.12.0 (2025-08-05)</h1>
<h2>🚀 Highlights</h2>
<ul>
<li>This release includes support for <a href="https://redis.io/docs/latest/operate/oss_and_stack/stack-with-enterprise/release-notes/redisce/redisos-8.2-release-notes/">Redis 8.2</a>.</li>
<li>Introduces an experimental Query Builders for <code>FTSearch</code>, <code>FTAggregate</code> and other search commands.</li>
<li>Adds support for <code>EPSILON</code> option in <code>FT.VSIM</code>.</li>
<li>Includes bug fixes and improvements contributed by the community related to ring and <a href="https://github.com/redis/go-redis/tree/master/extra/redisotel">redisotel</a>.</li>
</ul>
<h2>Changes</h2>
<ul>
<li>Improve stale issue workflow (<a href="https://redirect.github.com/redis/go-redis/pull/3458">#3458</a>)</li>
<li>chore(ci): Add 8.2 rc2 pre build for CI (<a href="https://redirect.github.com/redis/go-redis/pull/3459">#3459</a>)</li>
<li>Added new stream commands (<a href="https://redirect.github.com/redis/go-redis/pull/3450">#3450</a>)</li>
<li>feat: Add &quot;skip_verify&quot; to Sentinel (<a href="https://redirect.github.com/redis/go-redis/pull/3428">#3428</a>)</li>
<li>fix: <code>errors.Join</code> requires Go 1.20 or later (<a href="https://redirect.github.com/redis/go-redis/pull/3442">#3442</a>)</li>
<li>DOC-4344 document quickstart examples (<a href="https://redirect.github.com/redis/go-redis/pull/3426">#3426</a>)</li>
<li>feat(bitop): add support for the new bitop operations (<a href="https://redirect.github.com/redis/go-redis/pull/3409">#3409</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>feat: recover addIdleConn may occur panic (<a href="https://redirect.github.com/redis/go-redis/pull/2445">#2445</a>)</li>
<li>feat(ring): specify custom health check func via HeartbeatFn option (<a href="https://redirect.github.com/redis/go-redis/pull/2940">#2940</a>)</li>
<li>Add Query Builder for RediSearch commands (<a href="https://redirect.github.com/redis/go-redis/pull/3436">#3436</a>)</li>
<li>add configurable buffer sizes for Redis connections (<a href="https://redirect.github.com/redis/go-redis/pull/3453">#3453</a>)</li>
<li>Add VAMANA vector type to RediSearch (<a href="https://redirect.github.com/redis/go-redis/pull/3449">#3449</a>)</li>
<li>VSIM add <code>EPSILON</code> option (<a href="https://redirect.github.com/redis/go-redis/pull/3454">#3454</a>)</li>
<li>Add closing support to otel metrics instrumentation (<a href="https://redirect.github.com/redis/go-redis/pull/3444">#3444</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix(redisotel): fix buggy append in reportPoolStats (<a href="https://redirect.github.com/redis/go-redis/pull/3122">#3122</a>)</li>
<li>fix(search): return results even if doc is empty (<a href="https://redirect.github.com/redis/go-redis/pull/3457">#3457</a>)</li>
<li>[ISSUE-3402]: Ring.Pipelined return dial timeout error (<a href="https://redirect.github.com/redis/go-redis/pull/3403">#3403</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>Merges stale issues jobs into one job with two steps (<a href="https://redirect.github.com/redis/go-redis/pull/3463">#3463</a>)</li>
<li>improve code readability (<a href="https://redirect.github.com/redis/go-redis/pull/3446">#3446</a>)</li>
<li>chore(release): 9.12.0-beta.1 (<a href="https://redirect.github.com/redis/go-redis/pull/3460">#3460</a>)</li>
<li>DOC-5472 time series doc examples (<a href="https://redirect.github.com/redis/go-redis/pull/3443">#3443</a>)</li>
<li>Add VAMANA compression algorithm tests (<a href="https://redirect.github.com/redis/go-redis/pull/3461">#3461</a>)</li>
<li>bumped redis 8.2 version used in the CI/CD (<a href="https://redirect.github.com/redis/go-redis/pull/3451">#3451</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/andy-stark-redis"><code>@​andy-stark-redis</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/jouir"><code>@​jouir</code></a>, <a href="https://github.com/monkey92t"><code>@​monkey92t</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/rokn"><code>@​rokn</code></a>, <a href="https://github.com/smnvdev"><code>@​smnvdev</code></a>, <a href="https://github.com/strobil"><code>@​strobil</code></a> and <a href="https://github.com/wzy9607"><code>@​wzy9607</code></a></p>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a> made their first contribution in <a href="https://redirect.github.com/redis/go-redis/pull/3409">#3409</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/b7838dc4e72b25009e6de064ac2f20495944af39"><code>b7838dc</code></a> chore(release): 9.12.0 / redis 8.2 (<a href="https://redirect.github.com/redis/go-redis/issues/3464">#3464</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6a48d3fec17274e59a9f5d401558693c1f6c5fb9"><code>6a48d3f</code></a> feat: recover addIdleConn may occur panic (<a href="https://redirect.github.com/redis/go-redis/issues/2445">#2445</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/4767d9dfaf5b3485973f0210b1ba03f7431ad533"><code>4767d9d</code></a> fix(redisotel): fix buggy append in reportPoolStats (<a href="https://redirect.github.com/redis/go-redis/issues/3122">#3122</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7158a8dad4284f19ce9e531811275d37410d0623"><code>7158a8d</code></a> feat(ring): specify custom health check func via HeartbeatFn option (<a href="https://redirect.github.com/redis/go-redis/issues/2940">#2940</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8d15d03d4e0de4ed108ba9eef9508bc74cf0f512"><code>8d15d03</code></a> chore(github): merges into one job with two steps (<a href="https://redirect.github.com/redis/go-redis/issues/3463">#3463</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/375fa5d083c9b146fbe709b7ea439e4696845bee"><code>375fa5d</code></a> chore(doc): improve code readability (<a href="https://redirect.github.com/redis/go-redis/issues/3446">#3446</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f006e941d98b849e783cafc73cd14ae674f6eb8b"><code>f006e94</code></a> chore(release): 9.12.0-beta.1 (<a href="https://redirect.github.com/redis/go-redis/issues/3460">#3460</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f93bfa1f36d2da7e4559edfafc46455ede970ac6"><code>f93bfa1</code></a> feat(search): Add Query Builder for RediSearch commands (<a href="https://redirect.github.com/redis/go-redis/issues/3436">#3436</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/421c8a48b4373f58297d0ecae9f201ff94cf8f0d"><code>421c8a4</code></a> chore(doc): DOC-5472 time series doc examples (<a href="https://redirect.github.com/redis/go-redis/issues/3443">#3443</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2ff9249846f71017662353445c6409f5d57bc495"><code>2ff9249</code></a> chore(tests): Add VAMANA compression algorithm tests (<a href="https://redirect.github.com/redis/go-redis/issues/3461">#3461</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.11.0...v9.12.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.26.0 to 1.27.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>Release v1.27.0</h2>
<p>This release introduces a breaking change when migrating <code>aws-sdk-go</code> to <code>v2</code> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/630">RedHatInsights/insights-operator-utils#630</a>.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/ff448a03a2056af6502ed2b9dbe5cb8c39adf768"><code>ff448a0</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/625">#625</a> from joselsegura/cleanup</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/9d139d967e1b684f1491f8f8fab504494e0809b3"><code>9d139d9</code></a> Merge branch 'master' into cleanup</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/f102e88597efe37acbbc1a841260f3f04c6c7853"><code>f102e88</code></a> [CCXDEV-15370] Migrate aws-sdk-go to v2 (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/630">#630</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/9365fdd3980ee1fd489aa161a3a4ef151ca0e540"><code>9365fdd</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/629">#629</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/28797ad5e0a0b23cf3b9d4c48fce6f7708b8610c"><code>28797ad</code></a> Bump github.com/redis/go-redis/v9 from 9.11.0 to 9.12.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/9015af84a371e532d790454291bdd245324da05d"><code>9015af8</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/628">#628</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/b2b9451dc468dd43aeef11667bab144a5c93cafc"><code>b2b9451</code></a> Bump github.com/prometheus/client_golang from 1.22.0 to 1.23.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/aad4dcab7b16b33e62a712c66e033ea9a178073f"><code>aad4dca</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/627">#627</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/03263852ad6d02c1d6c9efe399089a2290786721"><code>0326385</code></a> Bump github.com/getsentry/sentry-go from 0.34.1 to 0.35.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/3461f0a00a51ee3ec0916f289b08ece4b100eecb"><code>3461f0a</code></a> Cleanup and update actions versions</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.26.0...v1.27.0">compare view</a></li>
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

### Comment by @swadeley on August 11, 2025 at 11:00 AM UTC

/retest

### Comment by @jlsherrill on August 11, 2025 at 01:29 PM UTC

/retest

### Comment by @rverdile on August 11, 2025 at 07:45 PM UTC

/retest

### Comment by @swadeley on August 12, 2025 at 07:54 AM UTC

/retest

### Comment by @jlsherrill on August 14, 2025 at 06:01 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 14, 2025 at 07:39 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1170*
