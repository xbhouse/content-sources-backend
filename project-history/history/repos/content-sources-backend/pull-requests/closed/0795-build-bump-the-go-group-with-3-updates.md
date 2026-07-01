---
type: pull_request
number: 795
title: "Build: Bump the go group with 3 updates"
state: closed
author: dependabot
created: 2024-09-02T04:26:40Z
updated: 2024-09-06T14:06:44Z
closed: 2024-09-06T14:06:43Z
base_branch: main
head_branch: dependabot/go_modules/go-5883b7de2e
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/795
---

# Pull Request #795: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: September 02, 2024 at 04:26 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-5883b7de2e`

## Description

Bumps the go group with 3 updates: [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils).

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.29 to 1.17.30
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a6e48aca89f21791c465cb8cf1136d72a28da31b"><code>a6e48ac</code></a> Release 2024-08-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d62062ba2ce184f3976938531cfa74914250327"><code>7d62062</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdf2372ddbac52292105633af192bbb8213bc57d"><code>bdf2372</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/84ca95e16adf482b8a80069d5ffa85814c7f61a2"><code>84ca95e</code></a> omitempty for NULL attribute values from custom marshalers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2739">#2739</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d7a7f5a021d5f64882fc1e219bd12725d9b75d41"><code>d7a7f5a</code></a> save sso cache token expiresAt in UTC (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2709">#2709</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/87cea8b18abee02e2e7343827fcb88b1c80c2ba2"><code>87cea8b</code></a> Release 2024-08-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93f8d22fd394e268d7eff4bac86937a815c784d1"><code>93f8d22</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bd44016700290f46dc6aa2ce6085f53a31794f06"><code>bd44016</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/097b04b54d639fdcff96d16870517a1d5e87569c"><code>097b04b</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2761">#2761</a> from aws/fix-dont-fail-test-if-credentials-set</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c10d54316d559ac37a2e0bb7e63f3ec985a5270e"><code>c10d543</code></a> Don't fail credentials unit tests if credentials are found on a file</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.29...credentials/v1.17.30">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.37.5 to 1.38.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-07-31)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Feature</strong>: Adds support for smithy-modeled endpoint resolution. A new rules-based endpoint resolution will be added to the SDK which will supercede and deprecate existing endpoint resolution. Specifically, EndpointResolver will be deprecated while BaseEndpoint and EndpointResolverV2 will take its place. For more information, please see the Endpoints section in our Developer Guide.</li>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/amplifyuibuilder/CHANGELOG.md#v1120-2023-07-31">v1.12.0</a>
<ul>
<li><strong>Feature</strong>: Amplify Studio releases GraphQL support for codegen job action.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/autoscaling</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/autoscaling/CHANGELOG.md#v1300-2023-07-31">v1.30.0</a>
<ul>
<li><strong>Feature</strong>: You can now configure an instance refresh to set its status to 'failed' when it detects that a specified CloudWatch alarm has gone into the ALARM state. You can also choose to roll back the instance refresh automatically when the alarm threshold is met.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cleanrooms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/cleanrooms/CHANGELOG.md#v130-2023-07-31">v1.3.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces custom SQL queries - an expanded set of SQL you can run. This release adds analysis templates, a new resource for storing pre-defined custom SQL queries ahead of time. This release also adds the Custom analysis rule, which lets you approve analysis templates for querying.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/codestarconnections</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/codestarconnections/CHANGELOG.md#v1150-2023-07-31">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: New integration with the Gitlab provider type.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/drs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/drs/CHANGELOG.md#v1150-2023-07-31">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Add support for in-aws right sizing</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/inspector2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/inspector2/CHANGELOG.md#v1160-2023-07-31">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This release adds 1 new API: BatchGetFindingDetails to retrieve enhanced vulnerability intelligence details for findings.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lookoutequipment</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/lookoutequipment/CHANGELOG.md#v1180-2023-07-31">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release includes new import resource, model versioning and resource policy features.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/omics</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/omics/CHANGELOG.md#v160-2023-07-31">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: Add CreationType filter for ListReadSets</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/rds/CHANGELOG.md#v1490-2023-07-31">v1.49.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for Aurora MySQL local write forwarding, which allows for forwarding of write operations from reader DB instances to the writer DB instance.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/route53</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/route53/CHANGELOG.md#v1290-2023-07-31">v1.29.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Route 53 now supports the Israel (Tel Aviv) Region (il-central-1) for latency records, geoproximity records, and private DNS for Amazon VPCs in that region.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/scheduler</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/scheduler/CHANGELOG.md#v120-2023-07-31">v1.2.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces automatic deletion of schedules in EventBridge Scheduler. If configured, EventBridge Scheduler automatically deletes a schedule after the schedule has completed its last invocation.</li>
</ul>
</li>
</ul>
<h1>Release (2023-07-28.2)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationinsights</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/applicationinsights/CHANGELOG.md#v1180-2023-07-282">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release enable customer to add/remove/update more than one workload for a component</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudformation</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/cloudformation/CHANGELOG.md#v1330-2023-07-282">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: This SDK release is for the feature launch of AWS CloudFormation RetainExceptOnCreate. It adds a new parameter retainExceptOnCreate in the following APIs: CreateStack, UpdateStack, RollbackStack, ExecuteChangeSet.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudfront</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/cloudfront/CHANGELOG.md#v1270-2023-07-282">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: Add a new JavaScript runtime version for CloudFront Functions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/connect/CHANGELOG.md#v1620-2023-07-282">v1.62.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for new number types.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kafka</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/kafka/CHANGELOG.md#v1210-2023-07-282">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: Amazon MSK has introduced new versions of ListClusterOperations and DescribeClusterOperation APIs. These v2 APIs provide information and insights into the ongoing operations of both MSK Provisioned and MSK Serverless clusters.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/pinpoint</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.38.0/service/pinpoint/CHANGELOG.md#v1210-2023-07-282">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: Added support for sending push notifications using the FCM v1 API with json credentials. Amazon Pinpoint customers can now deliver messages to Android devices using both FCM v1 API and the legacy FCM/GCM API</li>
</ul>
</li>
</ul>
<h1>Release (2023-07-28)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/787a81828a3812407a6d90036f07449d77a0f070"><code>787a818</code></a> Release 2023-07-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/953c8521385929d33d632132e6a6aea7463df9c7"><code>953c852</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9ad4f50ad39ef47c7a247a3dc8edd242dd7d4cd1"><code>9ad4f50</code></a> Update SDK's smithy-go dependency to v1.14.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bd77e8d520982bd06debfa82fa74826b7fd7c0f2"><code>bd77e8d</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a0433f2fa1b9ed7acd6a839fa4ff55070a0f3c9"><code>7a0433f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cdfb4c4297f5f0eabc4fd8b1523fcdefe17f9d38"><code>cdfb4c4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5cff0c441acb600e46ad75044570cedb40ea5bf4"><code>5cff0c4</code></a> fix ep20 refs in GH workflows (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2218">#2218</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/26bf48f3e2a8463e1c771b2d2fd4c387215424e7"><code>26bf48f</code></a> Add a new smithy-modeled rules based endpoint resolution. (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2214">#2214</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/28c87ccdf0153d7a8fadec14a0b3aa98c02d4cfc"><code>28c87cc</code></a> Release 2023-07-28.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4f4efd21b0463c0d167db88f1124997369ae5986"><code>4f4efd2</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.37.5...service/s3/v1.38.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.9 to 1.25.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/737dd9f5137f7ad570a9643eeddb94cd69aae360"><code>737dd9f</code></a> [CCXDEV-14052][CCXDEV-14098] Avoid duplicates when using Sentry/Glitchtip (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/577">#577</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4eb4ca0819aed30979b1cf334d0bec6a7bf56e86"><code>4eb4ca0</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/576">#576</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/87c044b561c57cef1a0365df7f2cb3ca3e6fa0c1"><code>87c044b</code></a> Bump github.com/prometheus/client_golang from 1.20.1 to 1.20.2</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/ea3fa7a427c07b900d8e1cad302ed00154c4d5cf"><code>ea3fa7a</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/575">#575</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/71d59cb4bd199ab3a9729e3a33435559781e54d5"><code>71d59cb</code></a> Bump github.com/prometheus/client_golang from 1.20.0 to 1.20.1</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/06947ff285eb5d5e6e470f48490f0b7e08e9d65c"><code>06947ff</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/574">#574</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/aeedde16e251144fec30ce3e583be18969d13af0"><code>aeedde1</code></a> Bump github.com/prometheus/client_golang from 1.19.1 to 1.20.0</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.9...v1.25.10">compare view</a></li>
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

### Comment by @app-sre-bot on September 02, 2024 at 04:27 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on September 06, 2024 at 02:06 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/795*
