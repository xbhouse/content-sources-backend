---
type: pull_request
number: 584
title: "Build: Bump the go group with 9 updates"
state: closed
author: dependabot
created: 2024-02-21T20:48:00Z
updated: 2024-02-26T04:54:32Z
closed: 2024-02-26T04:54:31Z
base_branch: main
head_branch: dependabot/go_modules/go-0b25df848e
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/584
---

# Pull Request #584: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: February 21, 2024 at 08:48 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-0b25df848e`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.4` | `1.5.6` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.5` | `1.25.7-0.20240204074919-46816ad31dde` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.25.0` | `1.25.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.0` | `1.17.2` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.33.1` | `1.33.3` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.2.1707758687` | `2024.2.1708461232` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.5.0` | `9.5.1` |

Updates `gorm.io/driver/postgres` from 1.5.4 to 1.5.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/postgres/commit/b3b67da9ecabf3a6de0a81ef9cab4f26686492a7"><code>b3b67da</code></a> only retract v1.5.5</li>
<li><a href="https://github.com/go-gorm/postgres/commit/d715278f3f57d8f653bec8726cf845f8d608f2fd"><code>d715278</code></a> fix: retract v1.5.5 (<a href="https://redirect.github.com/go-gorm/postgres/issues/249">#249</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/ae7e52706ecbafb57bbba7bedded29e47f1028b6"><code>ae7e527</code></a> update gorm to master latest (<a href="https://redirect.github.com/go-gorm/postgres/issues/242">#242</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/1cfbc8652cd8c56d1cd05c1a3e111318b3dd551d"><code>1cfbc86</code></a> refactor: distinguish between Unique and UniqueIndex (<a href="https://redirect.github.com/go-gorm/postgres/issues/195">#195</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/dc711bd7c69394788e6532ad9c89f5bdcb12b007"><code>dc711bd</code></a> let gorm support  limit and offset binding parameters,change the BindVar of p...</li>
<li><a href="https://github.com/go-gorm/postgres/commit/438e4fdec86ee25732eddd756741bf3c9882a60c"><code>438e4fd</code></a> chore(deps): bump actions/setup-go from 4 to 5 (<a href="https://redirect.github.com/go-gorm/postgres/issues/227">#227</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/e6d2435140dc0305dae1e61780f3a55cdced47c7"><code>e6d2435</code></a> override identifier length for postgres (<a href="https://redirect.github.com/go-gorm/postgres/issues/232">#232</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/postgres/compare/v1.5.4...v1.5.6">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.25.5 to 1.25.7-0.20240204074919-46816ad31dde
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/commits">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.25.0 to 1.25.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7888f0e10eb7ef5db0b67a8a27bb8eeee984257b"><code>7888f0e</code></a> Release 2024-02-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a9cd1dd63c4532e0afe4e5314750bf08f1ae68b"><code>4a9cd1d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bebcd8c54f31abf4fc37a2205de460dae58ac122"><code>bebcd8c</code></a> Update SDK's smithy-go dependency to v1.20.1</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/374e0b974c5f03f15165917cd839a03110c36404"><code>374e0b9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98b0dc8950774e9a1691307dc845a50e066d847c"><code>98b0dc8</code></a> dep: drop dependency on go-cmp in protocol tests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2506">#2506</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de4f646b7271b9698a67f63bff876ad40eb690ca"><code>de4f646</code></a> Release 2024-02-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a31327acbd07ca1d42ca83107993d7bcfa223dd8"><code>a31327a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1333952165a0f8ceea1c7de5d20d06af4925c30"><code>d133395</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b135f991ff3bef2f7c9ce8117dd59e1abae43a0f"><code>b135f99</code></a> fix(2502): zero region should explicitly fail endpoint resolution (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2503">#2503</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7cb105f17b3a12d4b69219ce9957d517dba299d8"><code>7cb105f</code></a> Release 2024-02-19</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.25.0...v1.25.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.0 to 1.17.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-12-02)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appsync</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/appsync/CHANGELOG.md#v1170-2022-12-02">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Fixes the URI for the evaluatecode endpoint to include the /v1 prefix (ie. &quot;/v1/dataplane-evaluatecode&quot;).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/ecs/CHANGELOG.md#v1201-2022-12-02">v1.20.1</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Amazon ECS</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/fms/CHANGELOG.md#v1210-2022-12-02">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: AWS Firewall Manager now supports Fortigate Cloud Native Firewall as a Service as a third-party policy type.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediaconvert</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/mediaconvert/CHANGELOG.md#v1280-2022-12-02">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: The AWS Elemental MediaConvert SDK has added support for configurable ID3 eMSG box attributes and the ability to signal them with InbandEventStream tags in DASH and CMAF outputs.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/medialive</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/medialive/CHANGELOG.md#v1250-2022-12-02">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: Updates to Event Signaling and Management (ESAM) API and documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/polly</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/polly/CHANGELOG.md#v1210-2022-12-02">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: Add language code for Finnish (fi-FI)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/proton</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/proton/CHANGELOG.md#v1180-2022-12-02">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: CreateEnvironmentAccountConnection RoleArn input is now optional</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/redshiftserverless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/redshiftserverless/CHANGELOG.md#v130-2022-12-02">v1.3.0</a>
<ul>
<li><strong>Feature</strong>: Add Table Level Restore operations for Amazon Redshift Serverless. Add multi-port support for Amazon Redshift Serverless endpoints. Add Tagging support to Snapshots and Recovery Points in Amazon Redshift Serverless.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sns</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/sns/CHANGELOG.md#v1187-2022-12-02">v1.18.7</a>
<ul>
<li><strong>Documentation</strong>: This release adds the message payload-filtering feature to the SNS Subscribe, SetSubscriptionAttributes, and GetSubscriptionAttributes API actions</li>
</ul>
</li>
</ul>
<h1>Release (2022-12-01)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/codecatalyst</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/codecatalyst/CHANGELOG.md#v100-2022-12-01">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: This release adds operations that support customers using the AWS Toolkits and Amazon CodeCatalyst, a unified software development service that helps developers develop, deploy, and maintain applications in the cloud. For more information, see the documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/comprehend</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/comprehend/CHANGELOG.md#v1200-2022-12-01">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Comprehend now supports semi-structured documents (such as PDF files or image files) as inputs for custom analysis using the synchronous APIs (ClassifyDocument and DetectEntities).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/gamelift</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/gamelift/CHANGELOG.md#v1160-2022-12-01">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: GameLift introduces a new feature, GameLift Anywhere. GameLift Anywhere allows you to integrate your own compute resources with GameLift. You can also use GameLift Anywhere to iteratively test your game servers without uploading the build to GameLift for every iteration.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/pipes</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/pipes/CHANGELOG.md#v100-2022-12-01">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: AWS introduces new Amazon EventBridge Pipes which allow you to connect sources (SQS, Kinesis, DDB, Kafka, MQ) to Targets (14+ EventBridge Targets) without any code, with filtering, batching, input transformation, and an optional Enrichment stage (Lambda, StepFunctions, ApiGateway, ApiDestinations)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sfn</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/sfn/CHANGELOG.md#v1160-2022-12-01">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for the AWS Step Functions Map state in Distributed mode. The changes include a new MapRun resource and several new and modified APIs.</li>
</ul>
</li>
</ul>
<h1>Release (2022-11-30)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/accessanalyzer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/accessanalyzer/CHANGELOG.md#v1180-2022-11-30">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for S3 cross account access points. IAM Access Analyzer will now produce public or cross account findings when it detects bucket delegation to external account access points.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/athena</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/athena/CHANGELOG.md#v1200-2022-11-30">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release includes support for using Apache Spark in Amazon Athena.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/dataexchange</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.2/service/dataexchange/CHANGELOG.md#v1170-2022-11-30">v1.17.0</a>
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.17.0...v1.17.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.33.1 to 1.33.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1257d4e205318d7dea08f5785d51d8ced05ccde6"><code>1257d4e</code></a> Release 2022-12-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f89f556a125709070bf3614e565233a90d2ee17c"><code>f89f556</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7c6f2cfe5d8e2c7be134e2210a01ebacda945c33"><code>7c6f2cf</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7cd0d9d9b4d7dd972da4aaf8983d6aca8d3f5874"><code>7cd0d9d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6067fb20cd68b5537e2fad6be41c367be5b78006"><code>6067fb2</code></a> Unify home dir logic between shared config and sso (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1960">#1960</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6fad028bc1ecb509e19336122056af6e6ff26abd"><code>6fad028</code></a> Release 2022-12-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cc6589758306fb4c877c9b4d232a2a3512df79b6"><code>cc65897</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e04c471a8ec448ab62e3768e36b2d1f2731058eb"><code>e04c471</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2de26a22ba8b2db393adce4828a90d7d9b079959"><code>2de26a2</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ad009a5d8d87c090c79d66e06d865795f994798"><code>7ad009a</code></a> Release 2022-12-13</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.33.1...service/ssm/v1.33.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2` from 2.15.0 to 2.15.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.15.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump andstor/file-existence-action from 2 to 3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1009">cloudevents/sdk-go#1009</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /test/conformance by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/993">cloudevents/sdk-go#993</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /test/benchmark by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/994">cloudevents/sdk-go#994</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/kafka by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/995">cloudevents/sdk-go#995</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /test/integration by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/996">cloudevents/sdk-go#996</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/kafka_sarama/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/997">cloudevents/sdk-go#997</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/http by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/998">cloudevents/sdk-go#998</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/nats by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/999">cloudevents/sdk-go#999</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/stan by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1004">cloudevents/sdk-go#1004</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/nats_jetstream by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1003">cloudevents/sdk-go#1003</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/nats/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1002">cloudevents/sdk-go#1002</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/nats_jetstream/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1001">cloudevents/sdk-go#1001</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/stan/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1000">cloudevents/sdk-go#1000</a></li>
<li>Propose the <code>confluent-kafka-go</code> binding for Kafka by <a href="https://github.com/yanmxa"><code>@​yanmxa</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1008">cloudevents/sdk-go#1008</a></li>
<li>Sync CESQL tck tests by <a href="https://github.com/Cali0707"><code>@​Cali0707</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1010">cloudevents/sdk-go#1010</a></li>
<li>Fix docstring typos in nats and jetstream protocol by <a href="https://github.com/jafossum"><code>@​jafossum</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1013">cloudevents/sdk-go#1013</a></li>
<li>Bump golangci/golangci-lint-action from 3 to 4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1016">cloudevents/sdk-go#1016</a></li>
<li>Bump the bundler group across 1 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1011">cloudevents/sdk-go#1011</a></li>
<li>Remove vi swp file by <a href="https://github.com/duglin"><code>@​duglin</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1020">cloudevents/sdk-go#1020</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Cali0707"><code>@​Cali0707</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1010">cloudevents/sdk-go#1010</a></li>
<li><a href="https://github.com/jafossum"><code>@​jafossum</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1013">cloudevents/sdk-go#1013</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.1">https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/67e389964131d55d65cd14b4eb32d57a47312695"><code>67e3899</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1020">#1020</a> from duglin/oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f0061e0ad9a7071a091fc4c92c91f504c1462bc5"><code>f0061e0</code></a> oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/4cc6c2d62d63da00f5b00efd4147e7dd5e40ee4c"><code>4cc6c2d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1011">#1011</a> from cloudevents/dependabot/bundler/docs/bundler-sec...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/b6949b02324839338c4bb8fbc7031cd0fba55899"><code>b6949b0</code></a> Bump the bundler group across 1 directories with 1 update</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/df51395f567d75bc5e2ac51a1796ae81522d2e87"><code>df51395</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1016">#1016</a> from cloudevents/dependabot/github_actions/golangci/...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/1af6e066857109cba2b636642ffa10a30facfce3"><code>1af6e06</code></a> Bump golangci/golangci-lint-action from 3 to 4</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/2574a05f7b10376e4d536157f447f1d4f3ead380"><code>2574a05</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1013">#1013</a> from jafossum/fix-nats-typos</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/2458ec7a7a89b645cfacd05ab75f99a1ca2dafb5"><code>2458ec7</code></a> Fix docstring typos in nats and jetstream protocol</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/04ed212498e406a218b44a6067400e76281bd169"><code>04ed212</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1010">#1010</a> from Cali0707/sync-tck-tests</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/06b03210273a6c6adda4879467c7eff2ba38f32f"><code>06b0321</code></a> fix like expression parser</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/v2` from 2.15.0 to 2.15.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.15.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump andstor/file-existence-action from 2 to 3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1009">cloudevents/sdk-go#1009</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /test/conformance by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/993">cloudevents/sdk-go#993</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /test/benchmark by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/994">cloudevents/sdk-go#994</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/kafka by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/995">cloudevents/sdk-go#995</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /test/integration by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/996">cloudevents/sdk-go#996</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/kafka_sarama/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/997">cloudevents/sdk-go#997</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/http by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/998">cloudevents/sdk-go#998</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/nats by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/999">cloudevents/sdk-go#999</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/stan by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1004">cloudevents/sdk-go#1004</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /samples/nats_jetstream by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1003">cloudevents/sdk-go#1003</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/nats/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1002">cloudevents/sdk-go#1002</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/nats_jetstream/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1001">cloudevents/sdk-go#1001</a></li>
<li>Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /protocol/stan/v2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1000">cloudevents/sdk-go#1000</a></li>
<li>Propose the <code>confluent-kafka-go</code> binding for Kafka by <a href="https://github.com/yanmxa"><code>@​yanmxa</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1008">cloudevents/sdk-go#1008</a></li>
<li>Sync CESQL tck tests by <a href="https://github.com/Cali0707"><code>@​Cali0707</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1010">cloudevents/sdk-go#1010</a></li>
<li>Fix docstring typos in nats and jetstream protocol by <a href="https://github.com/jafossum"><code>@​jafossum</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1013">cloudevents/sdk-go#1013</a></li>
<li>Bump golangci/golangci-lint-action from 3 to 4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1016">cloudevents/sdk-go#1016</a></li>
<li>Bump the bundler group across 1 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1011">cloudevents/sdk-go#1011</a></li>
<li>Remove vi swp file by <a href="https://github.com/duglin"><code>@​duglin</code></a> in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1020">cloudevents/sdk-go#1020</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Cali0707"><code>@​Cali0707</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1010">cloudevents/sdk-go#1010</a></li>
<li><a href="https://github.com/jafossum"><code>@​jafossum</code></a> made their first contribution in <a href="https://redirect.github.com/cloudevents/sdk-go/pull/1013">cloudevents/sdk-go#1013</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.1">https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/67e389964131d55d65cd14b4eb32d57a47312695"><code>67e3899</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1020">#1020</a> from duglin/oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f0061e0ad9a7071a091fc4c92c91f504c1462bc5"><code>f0061e0</code></a> oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/4cc6c2d62d63da00f5b00efd4147e7dd5e40ee4c"><code>4cc6c2d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1011">#1011</a> from cloudevents/dependabot/bundler/docs/bundler-sec...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/b6949b02324839338c4bb8fbc7031cd0fba55899"><code>b6949b0</code></a> Bump the bundler group across 1 directories with 1 update</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/df51395f567d75bc5e2ac51a1796ae81522d2e87"><code>df51395</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1016">#1016</a> from cloudevents/dependabot/github_actions/golangci/...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/1af6e066857109cba2b636642ffa10a30facfce3"><code>1af6e06</code></a> Bump golangci/golangci-lint-action from 3 to 4</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/2574a05f7b10376e4d536157f447f1d4f3ead380"><code>2574a05</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1013">#1013</a> from jafossum/fix-nats-typos</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/2458ec7a7a89b645cfacd05ab75f99a1ca2dafb5"><code>2458ec7</code></a> Fix docstring typos in nats and jetstream protocol</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/04ed212498e406a218b44a6067400e76281bd169"><code>04ed212</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1010">#1010</a> from Cali0707/sync-tck-tests</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/06b03210273a6c6adda4879467c7eff2ba38f32f"><code>06b0321</code></a> fix like expression parser</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.2.1707758687 to 2024.2.1708461232
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/64c8aea5e6de57a89b96fe44ad79aff09bf251c3"><code>64c8aea</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e4759de4b65e0b7a98d596b8e32...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.2.1707758687...release/v2024.2.1708461232">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.5.0 to 9.5.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.5.1</h2>
<h1>Changes</h1>
<p>Note: This release fixes the SETINFO issue from 9.5.0. This release restores connections to redis versions that do not have SETINFO. Thank you to our amazing community for their help with this issue</p>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>bug: Fix SETINFO ensuring it is set-and-forget (<a href="https://redirect.github.com/redis/go-redis/issues/2915">#2915</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>docs: README update to highlight how to disable sending client identification (<a href="https://redirect.github.com/redis/go-redis/issues/2913">#2913</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/d43a9fa887d9284ba42fcd46d46e97c56b34e132"><code>d43a9fa</code></a> Bump go-redis version to 9.5.1 (<a href="https://redirect.github.com/redis/go-redis/issues/2917">#2917</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/5da49b1abaef3bc65acae10debdbc72d7f5f32a1"><code>5da49b1</code></a> bug: Fix SETINFO ensuring it is set-and-forget (<a href="https://redirect.github.com/redis/go-redis/issues/2915">#2915</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/99527f0ac1e50d37fd25692b9484cba84c52f611"><code>99527f0</code></a> docs: README update to highlight how to disable sending client identification...</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.5.0...v9.5.1">compare view</a></li>
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

### Comment by @app-sre-bot on February 21, 2024 at 08:48 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on February 21, 2024 at 08:58 PM UTC

@dependabot ignore  https://github.com/go-gorm/gorm minor version
@dependabot ignore https://github.com/go-gorm/postgres  minor version

### Comment by @jlsherrill on February 22, 2024 at 02:29 PM UTC

https://github.com/dependabot ignore https://github.com/go-gorm/gorm minor version

### Comment by @jlsherrill on February 22, 2024 at 02:30 PM UTC

@dependabot ignore https://github.com/go-gorm/gorm minor version

### Comment by @jlsherrill on February 22, 2024 at 03:05 PM UTC

@dependabot ignore gorm.io/gorm minor version
@dependabot ignore gorm.io/driver/postgres minor version

### Comment by @dependabot on February 22, 2024 at 03:05 PM UTC

OK, I won't notify you about version 1.25.x of gorm.io/gorm again, unless you unignore it.

### Comment by @jlsherrill on February 22, 2024 at 03:05 PM UTC

https://github.com/dependabot ignore gorm.io/driver/postgres minor version

### Comment by @dependabot on February 22, 2024 at 03:05 PM UTC

OK, I won't notify you about version 1.5.x of gorm.io/driver/postgres again, unless you unignore it.

### Comment by @dependabot on February 26, 2024 at 04:54 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/584*
