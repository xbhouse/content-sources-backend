---
type: pull_request
number: 466
title: "Build: Bump the go group with 12 updates"
state: closed
author: dependabot
created: 2023-11-13T04:28:42Z
updated: 2023-11-20T05:03:41Z
closed: 2023-11-20T05:03:38Z
base_branch: main
head_branch: dependabot/go_modules/go-9dc6f3d970
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/466
---

# Pull Request #466: Build: Bump the go group with 12 updates

**Author**: @dependabot
**Created**: November 13, 2023 at 04:28 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9dc6f3d970`

## Description

Bumps the go group with 12 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.6` | `1.0.8` |
| [github.com/go-playground/validator/v10](https://github.com/go-playground/validator) | `10.15.5` | `10.16.0` |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.11.2` | `4.11.3` |
| [github.com/labstack/gommon](https://github.com/labstack/gommon) | `0.4.0` | `0.4.1` |
| [github.com/archdx/zerolog-sentry](https://github.com/archdx/zerolog-sentry) | `1.6.1` | `1.7.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.21.2` | `1.22.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.13.43` | `1.15.2` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.24.2` | `1.27.0` |
| [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) | `2023.10.1697742345` | `2023.10.1698674075` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.4.3` | `5.5.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.2.1` | `9.3.0` |
| [github.com/RedHatInsights/event-schemas-go](https://github.com/RedHatInsights/event-schemas-go) | `1.0.5` | `1.0.6` |

Updates `github.com/content-services/yummy` from 1.0.6 to 1.0.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/yummy/releases">github.com/content-services/yummy's releases</a>.</em></p>
<blockquote>
<h2>v1.0.8</h2>
<h2>What's Changed</h2>
<ul>
<li>handling case of no comps.xml in repo by <a href="https://github.com/xbhouse"><code>@​xbhouse</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/18">#18</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.7...v1.0.8"><code>v1.0.7...v1.0.8</code></a></p>
<h2>v1.0.7</h2>
<p>No release notes provided.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/7a6f14fbf00fb67ccc3da36480f41dcf0dba14e6"><code>7a6f14f</code></a> handling case of no comps.xml in repo (<a href="https://redirect.github.com/content-services/yummy/issues/18">#18</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/28b1d171d807caa031b5dfc456601c8d58b94d49"><code>28b1d17</code></a> Fixes 2442: add support for parsing groups and envs (<a href="https://redirect.github.com/content-services/yummy/issues/16">#16</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.6...v1.0.8">compare view</a></li>
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

Updates `github.com/labstack/echo/v4` from 4.11.2 to 4.11.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.11.3</h2>
<p><strong>Security</strong></p>
<ul>
<li>'c.Attachment' and 'c.Inline' should escape filename in 'Content-Disposition' header to avoid 'Reflect File Download' vulnerability. <a href="https://redirect.github.com/labstack/echo/pull/2541">#2541</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Tests: refactor context tests to be separate functions <a href="https://redirect.github.com/labstack/echo/pull/2540">#2540</a></li>
<li>Proxy middleware: reuse echo request context <a href="https://redirect.github.com/labstack/echo/pull/2537">#2537</a></li>
<li>Mark unmarshallable yaml struct tags as ignored <a href="https://redirect.github.com/labstack/echo/pull/2536">#2536</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.11.3 - 2023-11-07</h2>
<p><strong>Security</strong></p>
<ul>
<li>'c.Attachment' and 'c.Inline' should escape filename in 'Content-Disposition' header to avoid 'Reflect File Download' vulnerability. <a href="https://redirect.github.com/labstack/echo/pull/2541">#2541</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Tests: refactor context tests to be separate functions <a href="https://redirect.github.com/labstack/echo/pull/2540">#2540</a></li>
<li>Proxy middleware: reuse echo request context <a href="https://redirect.github.com/labstack/echo/pull/2537">#2537</a></li>
<li>Mark unmarshallable yaml struct tags as ignored <a href="https://redirect.github.com/labstack/echo/pull/2536">#2536</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/4b26cde851bc7a51e624c04dcc5d37be1ce0c84f"><code>4b26cde</code></a> Changelog for v4.11.3 (<a href="https://redirect.github.com/labstack/echo/issues/2542">#2542</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/14daeb968049b71296a80b91abd3883afd02b4d1"><code>14daeb9</code></a> Security: c.Attachment and c.Inline should escape name in `Content-Dispositio...</li>
<li><a href="https://github.com/labstack/echo/commit/50ebcd8d7c17457489df7bcbbcaa3745c687fd32"><code>50ebcd8</code></a> refactor context tests to be separate functions (<a href="https://redirect.github.com/labstack/echo/issues/2540">#2540</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/c7d6d4373fdfbef5d6f44df0a8ef410c198420ee"><code>c7d6d43</code></a> proxy middleware: reuse echo request context (<a href="https://redirect.github.com/labstack/echo/issues/2537">#2537</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/69a0de84158fd7cad326599d145c2248bcc15a69"><code>69a0de8</code></a> Mark unmarshallable yaml struct tags as ignored (<a href="https://redirect.github.com/labstack/echo/issues/2536">#2536</a>)</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.11.2...v4.11.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/labstack/gommon` from 0.4.0 to 0.4.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/gommon/releases">github.com/labstack/gommon's releases</a>.</em></p>
<blockquote>
<h2>v0.4.1 small updates</h2>
<ul>
<li>
<p>feat(random): use crypto/rand for random string generator <a href="https://redirect.github.com/labstack/gommon/issues/55">#55</a></p>
</li>
<li>
<p>update deps and CI flow <a href="https://redirect.github.com/labstack/gommon/issues/57">#57</a></p>
</li>
<li>
<p>update deps <a href="https://redirect.github.com/labstack/gommon/issues/54">#54</a></p>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/gommon/commit/bf0335abda3637e22e0d0fe5660023f22d8af857"><code>bf0335a</code></a> update deps and CI flow (<a href="https://redirect.github.com/labstack/gommon/issues/57">#57</a>)</li>
<li><a href="https://github.com/labstack/gommon/commit/508eabfa166278eb264b02aa96e64cfdc406e463"><code>508eabf</code></a> feat(random): use crypto/rand for random string generator (<a href="https://redirect.github.com/labstack/gommon/issues/55">#55</a>)</li>
<li><a href="https://github.com/labstack/gommon/commit/d43909e2f99b1d97fbd95deace5a8a004d348135"><code>d43909e</code></a> update deps</li>
<li>See full diff in <a href="https://github.com/labstack/gommon/compare/v0.4.0...v0.4.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/archdx/zerolog-sentry` from 1.6.1 to 1.7.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/9a0b4be1371b58baee77fbd708cf123f5eb3e587"><code>9a0b4be</code></a> Merge pull request <a href="https://redirect.github.com/archdx/zerolog-sentry/issues/13">#13</a> from saleh199/add-new-with-hub</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/221cdc66f9fc9baf89e6e25d288c40be2477d576"><code>221cdc6</code></a> feat: add NewWithHub function to utilize existing Sentry hub</li>
<li>See full diff in <a href="https://github.com/archdx/zerolog-sentry/compare/v1.6.1...v1.7.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.21.2 to 1.22.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5aa6e2b081b5d1b038a0cc5d9198c7a37abf6b4c"><code>5aa6e2b</code></a> Release 2023-11-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7dbcd52c14c8f992cb2fec090d826ab30881155e"><code>7dbcd52</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9b178445e4365543a410efefd6b0262c2ea3eb84"><code>9b17844</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/efa54869c68f9d770c9b83fda176c452906688ce"><code>efa5486</code></a> Allowlist bucket owner header (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2358">#2358</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5118d3bcdd253fbfd02b8e44966ba7a79ecfbf59"><code>5118d3b</code></a> Release 2023-11-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3c195fc3bcd7f20ae84bf0f999dc89e2487eaee"><code>f3c195f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f24b6a0665313f69c97fd3fa62b351358962bda4"><code>f24b6a0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f5a08768ef820ff5efd62a49ba50c61c9ca5dbcb"><code>f5a0876</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3674c3886336f797ad40b18bb1db8c35f4c7ed91"><code>3674c38</code></a> Release 2023-11-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b856b28d54158b2e3596afcd799cfbafddd843ca"><code>b856b28</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.21.2...v1.22.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.13.43 to 1.15.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/386724971857987a5d2a50f506f23df615765ac7"><code>3867249</code></a> Release 2022-03-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2fb960e531eb00693870edd03e816f4189e8f50a"><code>2fb960e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6201345923b416b9de0f9c444f585ed2ef9a53a4"><code>6201345</code></a> Update SDK's smithy-go dependency to v1.11.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa4f5d50dd26eb346a18184f3a82ca9a5603607d"><code>aa4f5d5</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9f1b668782a3e111f99685e0eaf6545430d6669d"><code>9f1b668</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cd1c8a3764935dd86fe4cdf8aed382adc0c1a79e"><code>cd1c8a3</code></a> codegen: Fix EventSteam code generation bug (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1638">#1638</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5399c50f8a2fc8b966d2368bfb1c8aaf546efefb"><code>5399c50</code></a> Release 2022-03-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/57ed4144aa9e38dfd1a8c41bbae9f4b4c2f56dac"><code>57ed414</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7797b2c8bb8e3393455231024242783a40d6a30"><code>e7797b2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a7068e374e813ff24326d09f2903bae0e3d57cb"><code>0a7068e</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.13.43...config/v1.15.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.24.2 to 1.27.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-07-01)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/databasemigrationservice/CHANGELOG.md#v1200-2022-07-01">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Added new features for AWS DMS version 3.4.7 that includes new endpoint settings for S3, OpenSearch, Postgres, SQLServer and Oracle.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/rds/CHANGELOG.md#v1215-2022-07-01">v1.21.5</a>
<ul>
<li><strong>Documentation</strong>: Adds support for additional retention periods to Performance Insights.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/s3/CHANGELOG.md#v1270-2022-07-01">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: Add presign support for HeadBucket, DeleteObject, and DeleteBucket. Fixes <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1076">#1076</a>.</li>
</ul>
</li>
</ul>
<h1>Release (2022-06-30)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/athena</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/athena/CHANGELOG.md#v1160-2022-06-30">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This feature introduces the API support for Athena's parameterized query and BatchGetPreparedStatement API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/customerprofiles</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/customerprofiles/CHANGELOG.md#v1180-2022-06-30">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release adds the optional MinAllowedConfidenceScoreForMerging parameter to the CreateDomain, UpdateDomain, and GetAutoMergingPreview APIs in Customer Profiles. This parameter is used as a threshold to influence the profile auto-merging step of the Identity Resolution process.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/emr</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/emr/CHANGELOG.md#v1200-2022-06-30">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for the ExecutionRoleArn parameter in the AddJobFlowSteps and DescribeStep APIs. Customers can use ExecutionRoleArn to specify the IAM role used for each job they submit using the AddJobFlowSteps API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/glue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/glue/CHANGELOG.md#v1270-2022-06-30">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: This release adds tag as an input of CreateDatabase</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kendra</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/kendra/CHANGELOG.md#v1290-2022-06-30">v1.29.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Kendra now provides a data source connector for alfresco</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mwaa</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/mwaa/CHANGELOG.md#v1130-2022-06-30">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Documentation updates for Amazon Managed Workflows for Apache Airflow.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/pricing</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/pricing/CHANGELOG.md#v1160-2022-06-30">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Documentation update for GetProducts Response.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/wellarchitected</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/wellarchitected/CHANGELOG.md#v1160-2022-06-30">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Added support for UpdateGlobalSettings API. Added status filter to ListWorkloadShares and ListLensShares.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/workmail</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/workmail/CHANGELOG.md#v1160-2022-06-30">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for managing user availability configurations in Amazon WorkMail.</li>
</ul>
</li>
</ul>
<h1>Release (2022-06-29)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.16.6
<ul>
<li><strong>Bug Fix</strong>: Fix aws/signer/v4 to not double sign Content-Length header. Fixes <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1728">#1728</a>. Thanks to <a href="https://github.com/matelang"><code>@​matelang</code></a> for creating the issue and PR.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appstream</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/appstream/CHANGELOG.md#v1170-2022-06-29">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Includes support for StreamingExperienceSettings in CreateStack and UpdateStack APIs</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/elasticloadbalancingv2/CHANGELOG.md#v1187-2022-06-29">v1.18.7</a>
<ul>
<li><strong>Documentation</strong>: This release adds two attributes for ALB. One, helps to preserve the host header and the other helps to modify, preserve, or remove the X-Forwarded-For header in the HTTP request.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/emr</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/emr/CHANGELOG.md#v1190-2022-06-29">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces additional optional parameter &quot;Throughput&quot; to VolumeSpecification to enable user to configure throughput for gp3 ebs volumes.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/medialive</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.27.0/service/medialive/CHANGELOG.md#v1210-2022-06-29">v1.21.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ece16973b6d828c6c0f547b4f6df80d4178eb49"><code>7ece169</code></a> Release 2022-07-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e0cb637c968882528d019af13e4cffbb8de614d8"><code>e0cb637</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56779da847f7fc665e45aac3d52fb66510b5c07a"><code>56779da</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e614a231bfb8beeb7adb07f71096f9abd5e930d7"><code>e614a23</code></a> codegen: add support for addition Amazon S3 presigned URL operations (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1747">#1747</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2f406b86bf03c67378326d302633ee5e9372387"><code>e2f406b</code></a> Release 2022-06-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef6a9bbfea108cc2e37b55fc94728829657dcc9f"><code>ef6a9bb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cbf4363fbf70ecc7bf18f1b464d3db063f82bd37"><code>cbf4363</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3d809f443eef139ce2c97419e576ece5deda9a02"><code>3d809f4</code></a> Release 2022-06-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d4ec10bfb9f8b55b29330aabf79bfb3d46d2ad68"><code>d4ec10b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79fe35a216cafcccb299098d99ab2cb0e494c6f8"><code>79fe35a</code></a> Update SDK's smithy-go dependency to v1.12.0</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.24.2...service/s3/v1.27.0">compare view</a></li>
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

### Comment by @app-sre-bot on November 13, 2023 at 04:28 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on November 20, 2023 at 05:03 AM UTC

Superseded by #476.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/466*
