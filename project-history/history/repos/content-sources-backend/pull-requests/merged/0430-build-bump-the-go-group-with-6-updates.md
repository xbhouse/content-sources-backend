---
type: pull_request
number: 430
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2023-10-16T04:18:24Z
updated: 2023-10-16T12:00:11Z
closed: 2023-10-16T12:00:02Z
merged: 2023-10-16T12:00:02Z
base_branch: main
head_branch: dependabot/go_modules/go-30bae33b02
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/430
---

# Pull Request #430: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: October 16, 2023 at 04:18 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-30bae33b02`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.11.1` | `4.11.2` |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.2` | `1.5.3` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.4` | `1.25.5` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.21.1` | `1.21.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.13.42` | `1.13.43` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.24.1` | `1.24.2` |

Updates `github.com/labstack/echo/v4` from 4.11.1 to 4.11.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.11.2</h2>
<p><strong>Security</strong></p>
<ul>
<li>Bump golang.org/x/net to prevent CVE-2023-39325 / CVE-2023-44487 HTTP/2 Rapid Reset Attack <a href="https://redirect.github.com/labstack/echo/pull/2527">#2527</a></li>
<li>fix(sec): randomString bias introduced by <a href="https://redirect.github.com/labstack/echo/issues/2490">#2490</a> <a href="https://redirect.github.com/labstack/echo/pull/2492">#2492</a></li>
<li>CSRF/RequestID mw: switch math/random usage to crypto/random <a href="https://redirect.github.com/labstack/echo/pull/2490">#2490</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Delete unused context in body_limit.go <a href="https://redirect.github.com/labstack/echo/pull/2483">#2483</a></li>
<li>Use Go 1.21 in CI <a href="https://redirect.github.com/labstack/echo/pull/2505">#2505</a></li>
<li>Fix some typos <a href="https://redirect.github.com/labstack/echo/pull/2511">#2511</a></li>
<li>Allow CORS middleware to send Access-Control-Max-Age: 0 <a href="https://redirect.github.com/labstack/echo/pull/2518">#2518</a></li>
<li>Bump dependancies <a href="https://redirect.github.com/labstack/echo/pull/2522">#2522</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.11.2 - 2023-10-11</h2>
<p><strong>Security</strong></p>
<ul>
<li>Bump golang.org/x/net to prevent CVE-2023-39325 / CVE-2023-44487 HTTP/2 Rapid Reset Attack <a href="https://redirect.github.com/labstack/echo/pull/2527">#2527</a></li>
<li>fix(sec): randomString bias introduced by <a href="https://redirect.github.com/labstack/echo/issues/2490">#2490</a> <a href="https://redirect.github.com/labstack/echo/pull/2492">#2492</a></li>
<li>CSRF/RequestID mw: switch math/random usage to crypto/random <a href="https://redirect.github.com/labstack/echo/pull/2490">#2490</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Delete unused context in body_limit.go <a href="https://redirect.github.com/labstack/echo/pull/2483">#2483</a></li>
<li>Use Go 1.21 in CI <a href="https://redirect.github.com/labstack/echo/pull/2505">#2505</a></li>
<li>Fix some typos <a href="https://redirect.github.com/labstack/echo/pull/2511">#2511</a></li>
<li>Allow CORS middleware to send Access-Control-Max-Age: 0 <a href="https://redirect.github.com/labstack/echo/pull/2518">#2518</a></li>
<li>Bump dependancies <a href="https://redirect.github.com/labstack/echo/pull/2522">#2522</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/98a523756d875bc13475bcb6237f09e771cbe321"><code>98a5237</code></a> Changelog for v4.11.2 (<a href="https://redirect.github.com/labstack/echo/issues/2529">#2529</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/89ae0e5f2ca6d01665255fd2e479ba98ab5ff4c8"><code>89ae0e5</code></a> Bump dependancies (<a href="https://redirect.github.com/labstack/echo/issues/2522">#2522</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/5780908c7cb110a8c4d56a62e32dc5cbc030a5ab"><code>5780908</code></a> Fix CVE-2023-39325 / CVE-2023-44487 (<a href="https://redirect.github.com/labstack/echo/issues/2527">#2527</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/4bc3e475e3137b6402933eec5e6fde641e0d2320"><code>4bc3e47</code></a> cors middleware: allow sending <code>Access-Control-Max-Age: 0</code> value with config....</li>
<li><a href="https://github.com/labstack/echo/commit/3950c444b726c1de9131d4dee4c9ae708768f26c"><code>3950c44</code></a> fix some typos (<a href="https://redirect.github.com/labstack/echo/issues/2511">#2511</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/77d5ae6a9173d89c49e008607d08df7ba41336f0"><code>77d5ae6</code></a> Use Go 1.21 in CI (<a href="https://redirect.github.com/labstack/echo/issues/2505">#2505</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/e6b96f8873fed46e71e0d34cddb81c533167f954"><code>e6b96f8</code></a> docs: add comments to util.go <code>randomString</code> (<a href="https://redirect.github.com/labstack/echo/issues/2494">#2494</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/b3ec8e0fdd9d904aa5b1b95479da20c4961a59eb"><code>b3ec8e0</code></a> fix(sec): <code>randomString</code> bias (<a href="https://redirect.github.com/labstack/echo/issues/2492">#2492</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/626f13e33830665e08d9d40e333dd13d9de8e672"><code>626f13e</code></a> CSRF/RequestID mw: switch math/random usage to crypto/random</li>
<li><a href="https://github.com/labstack/echo/commit/3f8ae15b57624dcd04bac482e454c9b665476d9f"><code>3f8ae15</code></a> delete unused context in body_limit.go (<a href="https://redirect.github.com/labstack/echo/issues/2483">#2483</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/labstack/echo/compare/v4.11.1...v4.11.2">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/driver/postgres` from 1.5.2 to 1.5.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/postgres/commit/89bd876286916a0c60329b9a8d87b13b8cd116d3"><code>89bd876</code></a> fix: default value with typecast (<a href="https://redirect.github.com/go-gorm/postgres/issues/211">#211</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/b3c309ac0b774b0bb0afea142dc3c6f71c5076fd"><code>b3c309a</code></a> Feature sqlstate support (<a href="https://redirect.github.com/go-gorm/postgres/issues/215">#215</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/f538e03e43252c5eadcac288c1c7370e524f98ae"><code>f538e03</code></a> Sort columns when create comments (<a href="https://redirect.github.com/go-gorm/postgres/issues/217">#217</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/f437ffe7520d60bb73290477bbb864909c744a75"><code>f437ffe</code></a> Fix query to join constraint_column_usage also on table_name. (<a href="https://redirect.github.com/go-gorm/postgres/issues/218">#218</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/be039123c3590962ca833c297802b6af8067c33f"><code>be03912</code></a> chore(deps): bump actions/checkout from 3 to 4 (<a href="https://redirect.github.com/go-gorm/postgres/issues/212">#212</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/948a202c1b04d4648b0864a96d3c1df5ea604cf7"><code>948a202</code></a> chore(deps): bump github.com/jackc/pgx/v5 from 5.3.1 to 5.4.3 (<a href="https://redirect.github.com/go-gorm/postgres/issues/204">#204</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/b04900e2f976f07387c61bc446c7c0825da56b0f"><code>b04900e</code></a> fix text default value produces &quot;'::text'&quot; as value (<a href="https://redirect.github.com/go-gorm/postgres/issues/210">#210</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/1eba30725087e149f180ee8f2815fd5d8b94eb1e"><code>1eba307</code></a> add from as valid clause for update (<a href="https://redirect.github.com/go-gorm/postgres/issues/213">#213</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/7ba909e839c99062711d10be7f1a03f19f8d9694"><code>7ba909e</code></a> Fix get default value from database, close <a href="https://github.com/go-gorm/gorm/is">https://github.com/go-gorm/gorm/is</a>...</li>
<li><a href="https://github.com/go-gorm/postgres/commit/cac4aec6dd625c073af6f7a996411043d61e7db3"><code>cac4aec</code></a> Add timestamp with timezone alias (<a href="https://redirect.github.com/go-gorm/postgres/issues/200">#200</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/go-gorm/postgres/compare/v1.5.2...v1.5.3">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.25.4 to 1.25.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/6bef318891b98263f3568c13093b5860245d2c52"><code>6bef318</code></a> add support for returning in sqlserver (<a href="https://redirect.github.com/go-gorm/gorm/issues/6585">#6585</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/1b240810106fd68f84cfe73bcacaf91a8e4ce1dd"><code>1b24081</code></a> chore(deps): bump actions/checkout from 3 to 4 (<a href="https://redirect.github.com/go-gorm/gorm/issues/6586">#6586</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/8c18714462de07fa3392b99eda089f2f9e3b6042"><code>8c18714</code></a> Don't call MethodByName with a variable arg (<a href="https://redirect.github.com/go-gorm/gorm/issues/6602">#6602</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/12ba285a52fb25c3422e16226666ba791f376c0b"><code>12ba285</code></a> *datatypes.JSON in model causes panic on tx.Statement.Changed (<a href="https://redirect.github.com/go-gorm/gorm/issues/6611">#6611</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/9d8a5bb208f5616638cbaad878a12d5ac73970d3"><code>9d8a5bb</code></a> feat: reuse name (<a href="https://redirect.github.com/go-gorm/gorm/issues/6626">#6626</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/2095d42b4c15de8d0cdaf64fd75e306bec40d9c4"><code>2095d42</code></a> fix: sqlite dialector cannot apply <code>PRIMARY KEY AUTOINCREMENT</code> type (<a href="https://redirect.github.com/go-gorm/gorm/issues/6624">#6624</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/e57e5d8884d801caa4ce0307bcd081f7e889e514"><code>e57e5d8</code></a> Update go.mod</li>
<li><a href="https://github.com/go-gorm/gorm/commit/653732e1c33858f5743a34f9fbfe66428d041760"><code>653732e</code></a> Update go testing versions</li>
<li><a href="https://github.com/go-gorm/gorm/commit/ac07543962994da4c6994ba3907417d7835a2619"><code>ac07543</code></a> Fixed error message when dialector fails to initialize (<a href="https://redirect.github.com/go-gorm/gorm/issues/6509">#6509</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.25.4...v1.25.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.21.1 to 1.21.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3b97d2a02cd4e27c40224f05aa1a7deba24abe2"><code>e3b97d2</code></a> Release 2023-10-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/863010ddb23c242c2a5d49d9f40094a6a49b5525"><code>863010d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6946ef8b9149fe75ac1b427ca2c7f57cdcb64549"><code>6946ef8</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d93ded4536184d38a664b4b75dadd36cbd79878"><code>6d93ded</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bebc232e7f65b02d0b519d11e73cf925c38e716f"><code>bebc232</code></a> fix: fail to load config if configured profile doesn't exist (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2309">#2309</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5de46742b7fb1b72d93d344ee81568800a707267"><code>5de4674</code></a> fix DNS timeout error not retried (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2300">#2300</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.21.1...v1.21.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.13.42 to 1.13.43
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3b97d2a02cd4e27c40224f05aa1a7deba24abe2"><code>e3b97d2</code></a> Release 2023-10-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/863010ddb23c242c2a5d49d9f40094a6a49b5525"><code>863010d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6946ef8b9149fe75ac1b427ca2c7f57cdcb64549"><code>6946ef8</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d93ded4536184d38a664b4b75dadd36cbd79878"><code>6d93ded</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bebc232e7f65b02d0b519d11e73cf925c38e716f"><code>bebc232</code></a> fix: fail to load config if configured profile doesn't exist (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2309">#2309</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5de46742b7fb1b72d93d344ee81568800a707267"><code>5de4674</code></a> fix DNS timeout error not retried (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2300">#2300</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.13.42...credentials/v1.13.43">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.24.1 to 1.24.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-03-21)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/chimesdkmessaging/CHANGELOG.md#v1130-2023-03-21">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Chime SDK messaging customers can now manage streaming configuration for messaging data for archival and analysis.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cleanrooms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/cleanrooms/CHANGELOG.md#v110-2023-03-21">v1.1.0</a>
<ul>
<li><strong>Feature</strong>: GA Release of AWS Clean Rooms, Added Tagging Functionality</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/ec2/CHANGELOG.md#v1910-2023-03-21">v1.91.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for AWS Network Firewall, AWS PrivateLink, and Gateway Load Balancers to Amazon VPC Reachability Analyzer, and it makes the path destination optional as long as a destination address in the filter at source is provided.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/internal/s3shared</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/internal/s3shared/CHANGELOG.md#v1140-2023-03-21">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: port v1 sdk 100-continue http header customization for s3 PutObject/UploadPart request and enable user config</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotsitewise</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/iotsitewise/CHANGELOG.md#v1280-2023-03-21">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: Provide support for tagging of data streams and enabling tag based authorization for property alias</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mgn</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/mgn/CHANGELOG.md#v1180-2023-03-21">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces the Import and export feature and expansion of the post-launch actions</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/s3/CHANGELOG.md#v1310-2023-03-21">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: port v1 sdk 100-continue http header customization for s3 PutObject/UploadPart request and enable user config</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-20)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationautoscaling</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/applicationautoscaling/CHANGELOG.md#v1190-2023-03-20">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: With this release customers can now tag their Application Auto Scaling registered targets with key-value pairs and manage IAM permissions for all the tagged resources centrally.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/neptune</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/neptune/CHANGELOG.md#v1200-2023-03-20">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release makes following few changes. db-cluster-identifier is now a required parameter of create-db-instance. describe-db-cluster will now return PendingModifiedValues and GlobalClusterIdentifier fields in the response.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3outposts</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/s3outposts/CHANGELOG.md#v1160-2023-03-20">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: S3 On Outposts added support for endpoint status, and a failed endpoint reason, if any</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/workdocs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/workdocs/CHANGELOG.md#v1140-2023-03-20">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new API, SearchResources, which enable users to search through metadata and content of folders, documents, document versions and comments in a WorkDocs site.</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-17)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/billingconductor</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/billingconductor/CHANGELOG.md#v160-2023-03-17">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new filter to ListAccountAssociations API and a new filter to ListBillingGroups API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/configservice/CHANGELOG.md#v1300-2023-03-17">v1.30.0</a>
<ul>
<li><strong>Feature</strong>: This release adds resourceType enums for types released from October 2022 through February 2023.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/service/databasemigrationservice/CHANGELOG.md#v1250-2023-03-17">v1.25.0</a>
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
<li><code>github.com/aws/aws-sdk-go-v2/config</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/ecs/v1.24.2/config/CHANGELOG.md#v11818-2023-03-16">v1.18.18</a></li>
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.24.1...service/ecs/v1.24.2">compare view</a></li>
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

### Comment by @app-sre-bot on October 16, 2023 at 04:19 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 16, 2023 at 11:18 AM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on October 16, 2023 at 11:59 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/430*
