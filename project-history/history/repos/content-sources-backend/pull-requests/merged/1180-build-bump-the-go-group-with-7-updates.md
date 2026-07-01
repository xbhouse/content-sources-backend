---
type: pull_request
number: 1180
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2025-08-18T07:00:19Z
updated: 2025-08-19T17:06:04Z
closed: 2025-08-19T17:06:00Z
merged: 2025-08-19T17:06:00Z
base_branch: main
head_branch: dependabot/go_modules/go-d47a6807c4
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1180
---

# Pull Request #1180: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: August 18, 2025 at 07:00 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-d47a6807c4`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.37.2` | `1.38.0` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.30.3` | `1.31.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.3` | `1.18.4` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.55.0` | `1.56.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.86.0` | `1.87.0` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.35.0` | `0.35.1` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.12.0` | `9.12.1` |

Updates `github.com/aws/aws-sdk-go-v2` from 1.37.2 to 1.38.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ab2d66f8e3df0b6850a13986c7d8c2d23cbd822"><code>0ab2d66</code></a> Release 2025-08-11</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae8100862e46c1ec9fc7a8f4a8c745928f571118"><code>ae81008</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6cf56c18f9973dcebe9e1be40fbdb9291c4610ec"><code>6cf56c1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e25292fe9b6bacc04a83e82c17e94cb1a0acdd9"><code>5e25292</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/14e9fb777684530dbf617624f37f33821d72b04a"><code>14e9fb7</code></a> upgrade to smithy v1.61.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fcdf6abd15520856083261043702400afd060f3d"><code>fcdf6ab</code></a> regen</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223029934bf4bd2bacace5c7f6d506e3bb60b62f"><code>2230299</code></a> fix changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/76aa8d73a6c12220e48133ca2f817423bc912190"><code>76aa8d7</code></a> feat: add support for per service options to Config (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3145">#3145</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8afe3272e8a75b70d1b93cc3d27a4fb15ae6ce34"><code>8afe327</code></a> Release 2025-08-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4d6e55d41fd68ed0157bc0622cf8e2a7987b9fc3"><code>4d6e55d</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.37.2...v1.38.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.30.3 to 1.31.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f445866bcc850b67c71e36882488e10f7c782e3"><code>2f44586</code></a> Release 2024-09-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/22d087682834495c3aebc0c0f1aa2db37c4785a6"><code>22d0876</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5454ab9dfb0bc726f1f79b8d807e81fc7124797a"><code>5454ab9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06150d96305d6b6c19db0a2e5d1c1f4fa4a95612"><code>06150d9</code></a> add tracing and metrics support (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2798">#2798</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18f7b40ea83d7ee75092085609a808b68c4d2000"><code>18f7b40</code></a> Release 2024-09-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e91c9c4e48a47fa92d82251dd6ba0f0edd4bff3d"><code>e91c9c4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6df0a09c479b5e8e59204eda1c6a481c311e5629"><code>6df0a09</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98ae6886ccefe00db38a3088e273e524abf3b196"><code>98ae688</code></a> Release 2024-09-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/222928c4e7cc17d406f6c24f2e69b6747e29d43e"><code>222928c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/887c4de750c99c575f37828f1a646cefadfc4aa7"><code>887c4de</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.3...v1.31.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.3 to 1.18.4
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-12-02)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appsync</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/appsync/CHANGELOG.md#v1170-2022-12-02">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Fixes the URI for the evaluatecode endpoint to include the /v1 prefix (ie. &quot;/v1/dataplane-evaluatecode&quot;).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/ecs/CHANGELOG.md#v1201-2022-12-02">v1.20.1</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Amazon ECS</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/fms/CHANGELOG.md#v1210-2022-12-02">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: AWS Firewall Manager now supports Fortigate Cloud Native Firewall as a Service as a third-party policy type.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediaconvert</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/mediaconvert/CHANGELOG.md#v1280-2022-12-02">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: The AWS Elemental MediaConvert SDK has added support for configurable ID3 eMSG box attributes and the ability to signal them with InbandEventStream tags in DASH and CMAF outputs.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/medialive</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/medialive/CHANGELOG.md#v1250-2022-12-02">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: Updates to Event Signaling and Management (ESAM) API and documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/polly</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/polly/CHANGELOG.md#v1210-2022-12-02">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: Add language code for Finnish (fi-FI)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/proton</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/proton/CHANGELOG.md#v1180-2022-12-02">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: CreateEnvironmentAccountConnection RoleArn input is now optional</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/redshiftserverless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/redshiftserverless/CHANGELOG.md#v130-2022-12-02">v1.3.0</a>
<ul>
<li><strong>Feature</strong>: Add Table Level Restore operations for Amazon Redshift Serverless. Add multi-port support for Amazon Redshift Serverless endpoints. Add Tagging support to Snapshots and Recovery Points in Amazon Redshift Serverless.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sns</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/sns/CHANGELOG.md#v1187-2022-12-02">v1.18.7</a>
<ul>
<li><strong>Documentation</strong>: This release adds the message payload-filtering feature to the SNS Subscribe, SetSubscriptionAttributes, and GetSubscriptionAttributes API actions</li>
</ul>
</li>
</ul>
<h1>Release (2022-12-01)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/codecatalyst</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/codecatalyst/CHANGELOG.md#v100-2022-12-01">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: This release adds operations that support customers using the AWS Toolkits and Amazon CodeCatalyst, a unified software development service that helps developers develop, deploy, and maintain applications in the cloud. For more information, see the documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/comprehend</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/comprehend/CHANGELOG.md#v1200-2022-12-01">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Comprehend now supports semi-structured documents (such as PDF files or image files) as inputs for custom analysis using the synchronous APIs (ClassifyDocument and DetectEntities).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/gamelift</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/gamelift/CHANGELOG.md#v1160-2022-12-01">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: GameLift introduces a new feature, GameLift Anywhere. GameLift Anywhere allows you to integrate your own compute resources with GameLift. You can also use GameLift Anywhere to iteratively test your game servers without uploading the build to GameLift for every iteration.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/pipes</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/pipes/CHANGELOG.md#v100-2022-12-01">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: AWS introduces new Amazon EventBridge Pipes which allow you to connect sources (SQS, Kinesis, DDB, Kafka, MQ) to Targets (14+ EventBridge Targets) without any code, with filtering, batching, input transformation, and an optional Enrichment stage (Lambda, StepFunctions, ApiGateway, ApiDestinations)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sfn</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/sfn/CHANGELOG.md#v1160-2022-12-01">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for the AWS Step Functions Map state in Distributed mode. The changes include a new MapRun resource and several new and modified APIs.</li>
</ul>
</li>
</ul>
<h1>Release (2022-11-30)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/accessanalyzer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/accessanalyzer/CHANGELOG.md#v1180-2022-11-30">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for S3 cross account access points. IAM Access Analyzer will now produce public or cross account findings when it detects bucket delegation to external account access points.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/athena</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/athena/CHANGELOG.md#v1200-2022-11-30">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release includes support for using Apache Spark in Amazon Athena.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/dataexchange</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.4/service/dataexchange/CHANGELOG.md#v1170-2022-11-30">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: This release enables data providers to license direct access to data in their Amazon S3 buckets or AWS Lake Formation data lakes through AWS Data Exchange. Subscribers get read-only access to the data and can use it in downstream AWS services, like Amazon Athena, without creating or managing copies.</li>
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
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93c3f1871b862d743e0bd2e2e7180246df3a9212"><code>93c3f18</code></a> Release 2022-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7254028f8bc89095326d9e3657fdbc98b98cca94"><code>7254028</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f43ad83db1b3da1c2ea37857524148c91189cb4c"><code>f43ad83</code></a> Update SDK's smithy-go dependency to v1.13.5</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/77d257ee120e67d45a5de6f0d6478f313a21b92a"><code>77d257e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/779e29ff5a4bcebe1ab7088ab12c4c95ce06f8aa"><code>779e29f</code></a> Release 2022-12-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f64d7d2b0a0033996b32ba9e1b18e5a923452b84"><code>f64d7d2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bc59f75ee4683ca886c3d701b49bb81db2efd4d"><code>9bc59f7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d9c18aa2bdd4c237a4919452f58e29c20ba484cc"><code>d9c18aa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0259b169b753daf77ad332c680a9ad1e3f56753d"><code>0259b16</code></a> Release 2022-11-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ee0277f1abad4856afc13ced2bfb90a43dbd9d34"><code>ee0277f</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.3...config/v1.18.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.55.0 to 1.56.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/077df5deb4f94e0cacc5c64c4538e49b6c711563"><code>077df5d</code></a> Release 2024-06-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3320b138b7ef295ad49afcaa9b04b6dca5e0e8ad"><code>3320b13</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1315201166439e42a027b644eb29de62e5e1ecee"><code>1315201</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8dddc9c41a16c7f622ba149d5f15b78e33ff7f1c"><code>8dddc9c</code></a> add SDK-specific feature tracking (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2682">#2682</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54f11c0ac651fbec8d65abe9cea3740af17c1460"><code>54f11c0</code></a> Release 2024-06-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5c4ad008ac68562397d65931315af7e4f1f5c73"><code>d5c4ad0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7057a6165b255266172c22528a5373da8f03f9e"><code>e7057a6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/374440df88873e8e363b1b25cb427a1042ea9018"><code>374440d</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2490">#2490</a> from aws/feat-aid-endpoints</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3133994bcb30e4ee3342eb900b0559e2919f696e"><code>3133994</code></a> fix changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5ceabb7c9115e06de25742efe5b167297f8a44b8"><code>5ceabb7</code></a> merge from main</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.55.0...service/s3/v1.56.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.86.0 to 1.87.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ab2d66f8e3df0b6850a13986c7d8c2d23cbd822"><code>0ab2d66</code></a> Release 2025-08-11</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae8100862e46c1ec9fc7a8f4a8c745928f571118"><code>ae81008</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6cf56c18f9973dcebe9e1be40fbdb9291c4610ec"><code>6cf56c1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e25292fe9b6bacc04a83e82c17e94cb1a0acdd9"><code>5e25292</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/14e9fb777684530dbf617624f37f33821d72b04a"><code>14e9fb7</code></a> upgrade to smithy v1.61.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fcdf6abd15520856083261043702400afd060f3d"><code>fcdf6ab</code></a> regen</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223029934bf4bd2bacace5c7f6d506e3bb60b62f"><code>2230299</code></a> fix changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/76aa8d73a6c12220e48133ca2f817423bc912190"><code>76aa8d7</code></a> feat: add support for per service options to Config (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3145">#3145</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8afe3272e8a75b70d1b93cc3d27a4fb15ae6ce34"><code>8afe327</code></a> Release 2025-08-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4d6e55d41fd68ed0157bc0622cf8e2a7987b9fc3"><code>4d6e55d</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.86.0...service/s3/v1.87.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.35.0 to 0.35.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.35.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Fix race conditions when accessing the scope during logging operations (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1050">#1050</a>)</li>
<li>Fix nil pointer dereference with malformed URLs when tracing is enabled in <code>fasthttp</code> and <code>fiber</code> integrations (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1055">#1055</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Bump <code>github.com/gofiber/fiber/v2</code> from 2.52.5 to 2.52.9 in <code>/fiber</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1067">#1067</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.35.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Fix race conditions when accessing the scope during logging operations (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1050">#1050</a>)</li>
<li>Fix nil pointer dereference with malformed URLs when tracing is enabled in <code>fasthttp</code> and <code>fiber</code> integrations (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1055">#1055</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Bump <code>github.com/gofiber/fiber/v2</code> from 2.52.5 to 2.52.9 in <code>/fiber</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1067">#1067</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/efaa5db1884130d134beca60749a0b2b23ce3018"><code>efaa5db</code></a> release: 0.35.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/b3adb472fa42077f9bab2e4fc3e15d076484048f"><code>b3adb47</code></a> Prepare 0.35.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1069">#1069</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a100bf57d3d5cfef009b3b5368eb9ffac47361f7"><code>a100bf5</code></a> Fix race conditions on log when accessing scope (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1050">#1050</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/17421c4a35ef82f863499236934c72fc8a93081f"><code>17421c4</code></a> Fix nil pointer dereference with malformed URLs when tracing is enabled (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1055">#1055</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/55af2784371b4b1f944277434f5244ea7603dc2b"><code>55af278</code></a> build(deps): bump github.com/gofiber/fiber/v2 in /fiber (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1067">#1067</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c846e7c35235313291d7c164d6d274e8b6dd61c7"><code>c846e7c</code></a> Merge branch 'release/0.35.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.35.0...v0.35.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.12.0 to 9.12.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.12.1</h2>
<h2>🚀 Highlights</h2>
<p>In the last version (9.12.0) the client introduced bigger write and read buffer sizes. The default value was 512KiB.
However, users reported that this is too big for most use cases and can lead to high memory usage.
In this version the default value is changed to 256KiB. The <code>README.md</code> was updated to reflect the
correct default value and include a note that the default value can be changed.</p>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix(options): Add buffer sizes to failover. Update README (<a href="https://redirect.github.com/redis/go-redis/pull/3468">#3468</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>fix(options): Add buffer sizes to failover. Update README (<a href="https://redirect.github.com/redis/go-redis/pull/3468">#3468</a>)</li>
<li>chore: update &amp; fix otel example (<a href="https://redirect.github.com/redis/go-redis/pull/3466">#3466</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/ndyakov"><code>@​ndyakov</code></a> and <a href="https://github.com/vmihailenco"><code>@​vmihailenco</code></a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.12.1 (2025-08-11)</h1>
<h2>🚀 Highlights</h2>
<p>In the last version (9.12.0) the client introduced bigger write and read buffer sized. The default value we set was 512KiB.
However, users reported that this is too big for most use cases and can lead to high memory usage.
In this version the default value is changed to 256KiB. The <code>README.md</code> was updated to reflect the
correct default value and include a note that the default value can be changed.</p>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix(options): Add buffer sizes to failover. Update README (<a href="https://redirect.github.com/redis/go-redis/pull/3468">#3468</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>fix(options): Add buffer sizes to failover. Update README (<a href="https://redirect.github.com/redis/go-redis/pull/3468">#3468</a>)</li>
<li>chore: update &amp; fix otel example (<a href="https://redirect.github.com/redis/go-redis/pull/3466">#3466</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/ndyakov"><code>@​ndyakov</code></a> and <a href="https://github.com/vmihailenco"><code>@​vmihailenco</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/7b4a537aef9e2670fff6990e81f278021e7c1499"><code>7b4a537</code></a> chore(release): 9.12.1, failover client buffer size fixes (<a href="https://redirect.github.com/redis/go-redis/issues/3469">#3469</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/94cfffa4176956d1745dcee587dc278faaed829c"><code>94cfffa</code></a> fix(options): Add buffer sizes to failover. Update README (<a href="https://redirect.github.com/redis/go-redis/issues/3468">#3468</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2c29dedc2db20e1a39e11a7f21040683cd116611"><code>2c29ded</code></a> chore(otel): upgrade otel example to Uptrace v2 (<a href="https://redirect.github.com/redis/go-redis/issues/3466">#3466</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.12.0...v9.12.1">compare view</a></li>
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

### Comment by @swadeley on August 18, 2025 at 07:41 AM UTC

`NOENT: no such file or directory, chmod '/usr/local/share/.cache/yarn/v6/npm-mkdirp-1.0.4`

### Comment by @rverdile on August 18, 2025 at 03:20 PM UTC

/retest

### Comment by @rverdile on August 18, 2025 at 07:31 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 19, 2025 at 05:05 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1180*
