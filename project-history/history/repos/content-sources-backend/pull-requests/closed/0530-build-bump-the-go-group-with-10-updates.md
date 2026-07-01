---
type: pull_request
number: 530
title: "Build: Bump the go group with 10 updates"
state: closed
author: dependabot
created: 2024-01-15T04:41:53Z
updated: 2024-01-18T21:15:27Z
closed: 2024-01-18T21:15:25Z
base_branch: main
head_branch: dependabot/go_modules/go-3250ef6a9a
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/530
---

# Pull Request #530: Build: Bump the go group with 10 updates

**Author**: @dependabot
**Created**: January 15, 2024 at 04:41 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-3250ef6a9a`

## Description

Bumps the go group with 10 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.9` | `1.0.10` |
| [github.com/redhatinsights/platform-go-middlewares](https://github.com/redhatinsights/platform-go-middlewares) | `0.20.0` | `1.0.0` |
| [github.com/archdx/zerolog-sentry](https://github.com/archdx/zerolog-sentry) | `1.8.0` | `1.8.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.24.0` | `1.24.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.16.13` | `1.16.14` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.30.1` | `1.31.0` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.25.0` | `0.26.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.1` | `5.5.2` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.3.1` | `9.4.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.24.13` | `1.25.0` |

Updates `github.com/content-services/yummy` from 1.0.9 to 1.0.10
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/yummy/releases">github.com/content-services/yummy's releases</a>.</em></p>
<blockquote>
<h2>v1.0.10</h2>
<h2>What's Changed</h2>
<ul>
<li>handling compressed comps by <a href="https://github.com/xbhouse"><code>@​xbhouse</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/21">#21</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.9...v1.0.10"><code>v1.0.9...v1.0.10</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/f05aed1b577e2af4c32e8fff1abb94cff2dc5604"><code>f05aed1</code></a> handle compressed comps (<a href="https://redirect.github.com/content-services/yummy/issues/21">#21</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.9...v1.0.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redhatinsights/platform-go-middlewares` from 0.20.0 to 1.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redhatinsights/platform-go-middlewares/releases">github.com/redhatinsights/platform-go-middlewares's releases</a>.</em></p>
<blockquote>
<h2>v1.0.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Create README and tag v1.0.0 by <a href="https://github.com/lzap"><code>@​lzap</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/46">RedHatInsights/platform-go-middlewares#46</a></li>
<li>Add dependabot updates by <a href="https://github.com/subpop"><code>@​subpop</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/48">RedHatInsights/platform-go-middlewares#48</a></li>
<li>Bump actions/checkout from 3 to 4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/49">RedHatInsights/platform-go-middlewares#49</a></li>
<li>Bump actions/setup-go from 3 to 5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/50">RedHatInsights/platform-go-middlewares#50</a></li>
<li>Bump github.com/onsi/ginkgo from 1.8.0 to 1.16.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/51">RedHatInsights/platform-go-middlewares#51</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.38.51 to 1.49.13 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/53">RedHatInsights/platform-go-middlewares#53</a></li>
<li>Bump github.com/go-chi/chi from 4.0.2+incompatible to 4.1.2+incompatible by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/52">RedHatInsights/platform-go-middlewares#52</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/dependabot"><code>@​dependabot</code></a> made their first contribution in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/49">RedHatInsights/platform-go-middlewares#49</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/platform-go-middlewares/compare/v0.20.0...v1.0.0">https://github.com/RedHatInsights/platform-go-middlewares/compare/v0.20.0...v1.0.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/1d808e9ce96a9d87b05464aaade1deb453c921eb"><code>1d808e9</code></a> Bump github.com/go-chi/chi from 4.0.2+incompatible to 4.1.2+incompatible</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/997ce1c461d7765eef31317020aafcd39154f663"><code>997ce1c</code></a> Bump github.com/aws/aws-sdk-go from 1.38.51 to 1.49.13</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/f67f4efa8a3399b2d1011c51e71cfa6cc62c1830"><code>f67f4ef</code></a> Bump github.com/onsi/ginkgo from 1.8.0 to 1.16.5</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/4eddedc7f961b2b036cbba0607570fa1291328e5"><code>4eddedc</code></a> Bump actions/setup-go from 3 to 5</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/a9c41d97fd605d4ce07070456ea3d4e56d6d859e"><code>a9c41d9</code></a> Bump actions/checkout from 3 to 4</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/e194b1b674cf9e09d62d8dd7b8a86d90a50a5744"><code>e194b1b</code></a> Add dependabot updates</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/e3779317d1aa349fdb5c7e1025d0b2bdacfe2ea2"><code>e377931</code></a> Create README and tag v1.0.0</li>
<li>See full diff in <a href="https://github.com/redhatinsights/platform-go-middlewares/compare/v0.20.0...v1.0.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/archdx/zerolog-sentry` from 1.8.0 to 1.8.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/a5e7c6e3cfc1dabe917a8ce22c4a7642a0c9bafd"><code>a5e7c6e</code></a> Merge pull request <a href="https://redirect.github.com/archdx/zerolog-sentry/issues/16">#16</a> from ezr-ondrej/breadcrumbs_with_level</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/262337b415b6872a4f36c1dc54acf9d7bb6d9576"><code>262337b</code></a> Merge pull request <a href="https://redirect.github.com/archdx/zerolog-sentry/issues/15">#15</a> from ezr-ondrej/trace_level_panic</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/518c55f95908dd884157ca2b6c9565d4c031ba0b"><code>518c55f</code></a> Add breadcrumbs support for WriteLevel</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/d4254603aef570666ea2619ed4aad72056dbb259"><code>d425460</code></a> Fixes a panic when Trace level log fired</li>
<li>See full diff in <a href="https://github.com/archdx/zerolog-sentry/compare/v1.8.0...v1.8.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.24.0 to 1.24.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/47dd1b1bcbde244357a82ef00fa6a61a9bfb9b39"><code>47dd1b1</code></a> Release 2024-01-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23145e3e605a93582020facfe13350b4153714e1"><code>23145e3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00e57bb7feb2d104387103aa4fbcd55dfff3a6a7"><code>00e57bb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/73e1a99f2fa858ca56627779852768a9198ba057"><code>73e1a99</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2441">#2441</a> from RanVaknin/fix-documentation-config</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0f8ad11593c219c52ad7fb12998c75ade39dc7ad"><code>0f8ad11</code></a> Fix SRA auth trailing checksum retry bug (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2438">#2438</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/efbc5aa622a882167129e69a88aa50c730cd1904"><code>efbc5aa</code></a> Release 2024-01-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78357bb87682230e24b15c01e807d7375a9474e4"><code>78357bb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e465ddd60d18e91b34de5917534cfa1542323027"><code>e465ddd</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/384ad3b7ec21eecb3c6c38b69f86fb6342906b11"><code>384ad3b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1126a91e53a20b18bc1db74225a8417bfb63f427"><code>1126a91</code></a> changelog added</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.24.0...v1.24.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.13 to 1.16.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/289bbd880a1aa7f88dc5a8be4a48869861aa949b"><code>289bbd8</code></a> Release 2022-09-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4169b1a526e46fd72ae275925d4f0a09152355bf"><code>4169b1a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e02802d3d2ccda7699a55800f7a8e58fa59cc2b8"><code>e02802d</code></a> Update SDK's smithy-go dependency to v1.13.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9f98fe25508538ae25d0b07ffe2709f5fae6e0d4"><code>9f98fe2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/824a8eb58b66dbcad91432561577be45ccbefadf"><code>824a8eb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/05a6aab6a1c03d3835dfe3285fabc2e275229b70"><code>05a6aab</code></a> Release 2022-09-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a4fbfe93646991e088c05daeb1586dd685973ce"><code>5a4fbfe</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9fde2d8c6bb409930211633e24afe7da8286ef7"><code>b9fde2d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18081400926772d62a7e8476cb156762deaaa5fc"><code>1808140</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.13...v1.16.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.30.1 to 1.31.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-03-21)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/chimesdkmessaging/CHANGELOG.md#v1130-2023-03-21">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Chime SDK messaging customers can now manage streaming configuration for messaging data for archival and analysis.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cleanrooms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/cleanrooms/CHANGELOG.md#v110-2023-03-21">v1.1.0</a>
<ul>
<li><strong>Feature</strong>: GA Release of AWS Clean Rooms, Added Tagging Functionality</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/ec2/CHANGELOG.md#v1910-2023-03-21">v1.91.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for AWS Network Firewall, AWS PrivateLink, and Gateway Load Balancers to Amazon VPC Reachability Analyzer, and it makes the path destination optional as long as a destination address in the filter at source is provided.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/internal/s3shared</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/internal/s3shared/CHANGELOG.md#v1140-2023-03-21">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: port v1 sdk 100-continue http header customization for s3 PutObject/UploadPart request and enable user config</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotsitewise</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/iotsitewise/CHANGELOG.md#v1280-2023-03-21">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: Provide support for tagging of data streams and enabling tag based authorization for property alias</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mgn</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/mgn/CHANGELOG.md#v1180-2023-03-21">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces the Import and export feature and expansion of the post-launch actions</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/s3/CHANGELOG.md#v1310-2023-03-21">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: port v1 sdk 100-continue http header customization for s3 PutObject/UploadPart request and enable user config</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-20)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationautoscaling</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/applicationautoscaling/CHANGELOG.md#v1190-2023-03-20">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: With this release customers can now tag their Application Auto Scaling registered targets with key-value pairs and manage IAM permissions for all the tagged resources centrally.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/neptune</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/neptune/CHANGELOG.md#v1200-2023-03-20">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release makes following few changes. db-cluster-identifier is now a required parameter of create-db-instance. describe-db-cluster will now return PendingModifiedValues and GlobalClusterIdentifier fields in the response.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3outposts</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/s3outposts/CHANGELOG.md#v1160-2023-03-20">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: S3 On Outposts added support for endpoint status, and a failed endpoint reason, if any</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/workdocs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/workdocs/CHANGELOG.md#v1140-2023-03-20">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new API, SearchResources, which enable users to search through metadata and content of folders, documents, document versions and comments in a WorkDocs site.</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-17)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/billingconductor</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/billingconductor/CHANGELOG.md#v160-2023-03-17">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new filter to ListAccountAssociations API and a new filter to ListBillingGroups API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/configservice/CHANGELOG.md#v1300-2023-03-17">v1.30.0</a>
<ul>
<li><strong>Feature</strong>: This release adds resourceType enums for types released from October 2022 through February 2023.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/service/databasemigrationservice/CHANGELOG.md#v1250-2023-03-17">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: S3 setting to create AWS Glue Data Catalog. Oracle setting to control conversion of timestamp column. Support for Kafka SASL Plain authentication. Setting to map boolean from PostgreSQL to Redshift. SQL Server settings to force lob lookup on inline LOBs and to control access of database logs.</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-16)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/config</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.31.0/config/CHANGELOG.md#v11818-2023-03-16">v1.18.18</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.30.1...service/s3/v1.31.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.25.0 to 0.26.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.26.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.26.0.</p>
<h3>Breaking Changes</h3>
<p>As previously announced, this release removes some methods from the SDK.</p>
<ul>
<li><code>sentry.TransactionName()</code> use <code>sentry.WithTransactionName()</code> instead.</li>
<li><code>sentry.OpName()</code> use <code>sentry.WithOpName()</code> instead.</li>
<li><code>sentry.TransctionSource()</code> use <code>sentry.WithTransactionSource()</code> instead.</li>
<li><code>sentry.SpanSampled()</code> use <code>sentry.WithSpanSampled()</code> instead.</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add <code>WithDescription</code> span option (<a href="https://redirect.github.com/getsentry/sentry-go/pull/751">#751</a>)</p>
<pre lang="go"><code>span := sentry.StartSpan(ctx, &quot;http.client&quot;, WithDescription(&quot;GET /api/users&quot;))
</code></pre>
</li>
<li>
<p>Add support for package name parsing in Go 1.20 and higher (<a href="https://redirect.github.com/getsentry/sentry-go/pull/730">#730</a>)</p>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Apply <code>ClientOptions.SampleRate</code> only to errors &amp; messages (<a href="https://redirect.github.com/getsentry/sentry-go/pull/754">#754</a>)</li>
<li>Check if git is available before executing any git commands (<a href="https://redirect.github.com/getsentry/sentry-go/pull/737">#737</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.26.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.26.0.</p>
<h3>Breaking Changes</h3>
<p>As previously announced, this release removes some methods from the SDK.</p>
<ul>
<li><code>sentry.TransactionName()</code> use <code>sentry.WithTransactionName()</code> instead.</li>
<li><code>sentry.OpName()</code> use <code>sentry.WithOpName()</code> instead.</li>
<li><code>sentry.TransctionSource()</code> use <code>sentry.WithTransactionSource()</code> instead.</li>
<li><code>sentry.SpanSampled()</code> use <code>sentry.WithSpanSampled()</code> instead.</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add <code>WithDescription</code> span option (<a href="https://redirect.github.com/getsentry/sentry-go/pull/751">#751</a>)</p>
<pre lang="go"><code>span := sentry.StartSpan(ctx, &quot;http.client&quot;, WithDescription(&quot;GET /api/users&quot;))
</code></pre>
</li>
<li>
<p>Add support for package name parsing in Go 1.20 and higher (<a href="https://redirect.github.com/getsentry/sentry-go/pull/730">#730</a>)</p>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Apply <code>ClientOptions.SampleRate</code> only to errors &amp; messages (<a href="https://redirect.github.com/getsentry/sentry-go/pull/754">#754</a>)</li>
<li>Check if git is available before executing any git commands (<a href="https://redirect.github.com/getsentry/sentry-go/pull/737">#737</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/81ffb249e3bd090121e61e495982960f89f0e678"><code>81ffb24</code></a> release: 0.26.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/46c2ed3cf1371ee435ea2f55322694afe0b0d3bf"><code>46c2ed3</code></a> Fix revive CI issues (<a href="https://redirect.github.com/getsentry/sentry-go/issues/763">#763</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/32f1b19f9fcfec2158fca065705393e4ccc2409a"><code>32f1b19</code></a> Remove depguard (<a href="https://redirect.github.com/getsentry/sentry-go/issues/762">#762</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d60b74342c3659202c17c10cc127fd01b9f4cf68"><code>d60b743</code></a> Update CI (<a href="https://redirect.github.com/getsentry/sentry-go/issues/761">#761</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7fa0b42c04a087c8a8c7398a7468c0785cf999ac"><code>7fa0b42</code></a> doc: fix typo (<a href="https://redirect.github.com/getsentry/sentry-go/issues/760">#760</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9954a51f45fdea2816f3ddd0aac05f7834a84108"><code>9954a51</code></a> Prepare 0.26.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/759">#759</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d302d89316ac53e40fa0563236e36f932fb4384a"><code>d302d89</code></a> Only apply <code>SampleRate </code> to errors/messages (<a href="https://redirect.github.com/getsentry/sentry-go/issues/754">#754</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/fff66d18d3019c0c75c88edb26f8d79d52d8a126"><code>fff66d1</code></a> feat(tracing): Add WithDescription SpanOption (<a href="https://redirect.github.com/getsentry/sentry-go/issues/751">#751</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/032ecc1861700897888aec5644403a3d0a435007"><code>032ecc1</code></a> Remove Span.isTransaction attribute (<a href="https://redirect.github.com/getsentry/sentry-go/issues/736">#736</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/03c634502f8fb416a744892f1ba3c4d3baba291c"><code>03c6345</code></a> ci: improve flaky tests (<a href="https://redirect.github.com/getsentry/sentry-go/issues/744">#744</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.25.0...v0.26.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.5.1 to 5.5.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.2 (January 13, 2024)</h1>
<ul>
<li>Allow NamedArgs to start with underscore</li>
<li>pgproto3: Maximum message body length support (jeremy.spriet)</li>
<li>Upgrade golang.org/x/crypto to v0.17.0</li>
<li>Add snake_case support to RowToStructByName (Tikhon Fedulov)</li>
<li>Fix: update description cache after exec prepare (James Hartig)</li>
<li>Fix: pipeline checks if it is closed (James Hartig and Ryan Fowler)</li>
<li>Fix: normalize timeout / context errors during TLS startup (Samuel Stauffer)</li>
<li>Add OnPgError for easier centralized error handling (James Hartig)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/b7de418d46878f65f91c377297d98b8b1a9c406d"><code>b7de418</code></a> Release v5.5.2</li>
<li><a href="https://github.com/jackc/pgx/commit/b99e2bb7e0818428092e955cb0ee9cff45504bfd"><code>b99e2bb</code></a> Document max read and write sizes for large objects</li>
<li><a href="https://github.com/jackc/pgx/commit/52f215142261839b465103c9eb2423685a0cdc39"><code>52f2151</code></a> Allow NamedArgs to start with underscore</li>
<li><a href="https://github.com/jackc/pgx/commit/dfb6489612684f664c0176416e5d114d26cbd10c"><code>dfb6489</code></a> fix typo in doc.go</li>
<li><a href="https://github.com/jackc/pgx/commit/9346d48035562654e15933d2477356150b7934ff"><code>9346d48</code></a> fix OpenDBFromPool example</li>
<li><a href="https://github.com/jackc/pgx/commit/1fdd17041a29847878409c3d4c855c58b242d3b8"><code>1fdd170</code></a> feat(pgproto3): expose MaxExpectedBodyLen and ActualBodyLen in ExceededMaxBod...</li>
<li><a href="https://github.com/jackc/pgx/commit/f654d61d792865c428cff645c9c979ec7ab7f531"><code>f654d61</code></a> Make note about possible parse config error message redaction change</li>
<li><a href="https://github.com/jackc/pgx/commit/5d26bbefd81aaeab3fb3d80fb61bce122b099cb2"><code>5d26bbe</code></a> Make pgconn.ConnectError and pgconn.ParseConfigError public</li>
<li><a href="https://github.com/jackc/pgx/commit/44768b5a01bddfcb30a6ff7e4c2b8db313c060b3"><code>44768b5</code></a> fix a typo in config_test.go</li>
<li><a href="https://github.com/jackc/pgx/commit/6f2ce92356f35263258ef7baa4cde86a3e1c6255"><code>6f2ce92</code></a> Upgrade golang.org/x/crypto to v0.17.0</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.1...v5.5.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.3.1 to 9.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.4.0</h2>
<h1>Changes</h1>
<h2>Breaking Changes</h2>
<ul>
<li>Revert <a href="https://redirect.github.com/redis/go-redis/issues/2818">#2818</a> due to it be a breaking change (<a href="https://redirect.github.com/redis/go-redis/issues/2861">#2861</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>Add Suffix support to default client set info (<a href="https://redirect.github.com/redis/go-redis/issues/2852">#2852</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>return raw value instead of function calling in Result() (<a href="https://redirect.github.com/redis/go-redis/issues/2831">#2831</a>)</li>
<li>Add Redis Enterprise tests (<a href="https://redirect.github.com/redis/go-redis/issues/2847">#2847</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/lowang-bh"><code>@​lowang-bh</code></a> and <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/531f06861b26fc01b1b75f777d7155d286b0943c"><code>531f068</code></a> 9.4.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2862">#2862</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/27581fcb437fa7387d1dcfb6a43c5ccf1f4d2273"><code>27581fc</code></a> Change Z Member type to interface (<a href="https://redirect.github.com/redis/go-redis/issues/2861">#2861</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a32be3d93d60b9b9a3711bb3c248684246168bae"><code>a32be3d</code></a> Add Suffix support to default client set info (<a href="https://redirect.github.com/redis/go-redis/issues/2852">#2852</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/d8e3e95866c094700ffc8f7e2705c2b8fb7a3793"><code>d8e3e95</code></a> return raw value instead of funcation calling in Result() (<a href="https://redirect.github.com/redis/go-redis/issues/2831">#2831</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/00229238c00584052588ee33f448ca31e0cc543d"><code>0022923</code></a> Change Env vars in RE CI (<a href="https://redirect.github.com/redis/go-redis/issues/2856">#2856</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2e121910233d77b23d76a64e7d37628d1a73c0c1"><code>2e12191</code></a> Change RE build in CI (<a href="https://redirect.github.com/redis/go-redis/issues/2855">#2855</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b76230924354e418e500679c0460706dfcf5f82b"><code>b762309</code></a> Add RE tests (<a href="https://redirect.github.com/redis/go-redis/issues/2847">#2847</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.3.1...v9.4.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.24.13 to 1.25.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/ee63cd11459da087ac734039ab360f318f4addee"><code>ee63cd1</code></a> Refactor MultiBrokerConfiguration into a slice of BrokerConfiguration objects</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/3dd575a980c60dcd7b7a7dae280d30a24387520e"><code>3dd575a</code></a> support for multiple broker addresses when using clowder config</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/88a487584aea272ab1a7c720eab6bb0fc9c5c4fe"><code>88a4875</code></a> Bump github.com/aws/aws-sdk-go from 1.49.16 to 1.49.17</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/78d35d32df7100cd462a18b6d92af73cb94de577"><code>78d35d3</code></a> Bump github.com/aws/aws-sdk-go from 1.49.15 to 1.49.16 (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/398">#398</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/77b80d2bb46bf37662da70665af95fcd5dc83997"><code>77b80d2</code></a> Bump github.com/redis/go-redis/v9 from 9.3.1 to 9.4.0 (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/399">#399</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/ce05bb315dadbbaf29fd62622f24262c5bacfa73"><code>ce05bb3</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/397">#397</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/6bfc212a13897cd5bf89b6dbebe0e7525e660641"><code>6bfc212</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/396">#396</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/3f4649985d9a7b6d95bbef4c80c2c99ffe74915d"><code>3f46499</code></a> Bump github.com/aws/aws-sdk-go from 1.49.13 to 1.49.15</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/0592d4d903595060ffb2dca16c863d23d05c52dd"><code>0592d4d</code></a> Bump github.com/xdg/scram</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/e37a6f7f5578fd55dbe8b980f4f91914a67ecf6d"><code>e37a6f7</code></a> [CCXDEV-12032] Add common Kafka configuration and clowder-related utilities (...</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.24.13...v1.25.0">compare view</a></li>
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

### Comment by @app-sre-bot on January 15, 2024 at 04:42 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on January 18, 2024 at 09:15 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/530*
