---
type: pull_request
number: 487
title: "Build: Bump the go group with 2 updates"
state: merged
author: dependabot
created: 2023-11-27T04:38:36Z
updated: 2023-11-27T14:12:49Z
closed: 2023-11-27T14:12:38Z
merged: 2023-11-27T14:12:38Z
base_branch: main
head_branch: dependabot/go_modules/go-a8cd83a2f6
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/487
---

# Pull Request #487: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: November 27, 2023 at 04:38 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-a8cd83a2f6`

## Description

Bumps the go group with 2 updates: [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) and [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2).

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.3 to 1.16.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f14e2a8eba44a6affbdff479aa60775bad7ca0d"><code>5f14e2a</code></a> Release 2022-05-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ffe56d085d32ada3e9feb9bc7394d1debb22947"><code>2ffe56d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/659ec3e5e595af7a6a433958bf55a5f95f4a0980"><code>659ec3e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3442469932326cf3cd56ba7014c410339b8b251"><code>b344246</code></a> Fix iotsecuretunneling and mobile API clients to use correct signing name (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1">#1</a>...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40a800b6768bfb08ea9fce4d0e0786b24404f33d"><code>40a800b</code></a> aws/retry: Fixup documentation typos(<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1698">#1698</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2fc2c1dfcd1c5828f12971b1965a5448f754975e"><code>2fc2c1d</code></a> fix feature/dynamodb/expression/condition.go Or documentation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1702">#1702</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a34da05fbb2c27569c66592c7d489f55acf7bd94"><code>a34da05</code></a> internal/ini: Remove unused and stale fuzzing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1710">#1710</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6f190f6ff1fc106f0f5e4191d4ab152cb415f66f"><code>6f190f6</code></a> Release 2022-05-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/30ac919af49bdd73947b06486ae2f8c41786bb8e"><code>30ac919</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1870c456331c4f9a5193df31149c684ee1f2469"><code>c1870c4</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.3...v1.16.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.27.2 to 1.28.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-10-19)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/chimesdkmessaging/CHANGELOG.md#v1116-2022-10-19">v1.11.6</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Chime Messaging SDK</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudtrail</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/cloudtrail/CHANGELOG.md#v1190-2022-10-19">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: This release includes support for exporting CloudTrail Lake query results to an Amazon S3 bucket.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/configservice/CHANGELOG.md#v1270-2022-10-19">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: This release adds resourceType enums for AppConfig, AppSync, DataSync, EC2, EKS, Glue, GuardDuty, SageMaker, ServiceDiscovery, SES, Route53 types.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/connect/CHANGELOG.md#v1330-2022-10-19">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: This release adds API support for managing phone numbers that can be used across multiple AWS regions through telephony traffic distribution.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/managedblockchain</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/managedblockchain/CHANGELOG.md#v1130-2022-10-19">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Adding new Accessor APIs for Amazon Managed Blockchain</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/s3/CHANGELOG.md#v1280-2022-10-19">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: Updates internal logic for constructing API endpoints. We have added rule-based endpoints and internal model parameters.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/supportapp</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/supportapp/CHANGELOG.md#v110-2022-10-19">v1.1.0</a>
<ul>
<li><strong>Feature</strong>: This release adds the RegisterSlackWorkspaceForOrganization API. You can use the API to register a Slack workspace for an AWS account that is part of an organization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/workspacesweb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/workspacesweb/CHANGELOG.md#v170-2022-10-19">v1.7.0</a>
<ul>
<li><strong>Feature</strong>: WorkSpaces Web now supports user access logging for recording session start, stop, and URL navigation.</li>
</ul>
</li>
</ul>
<h1>Release (2022-10-18)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/frauddetector</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/frauddetector/CHANGELOG.md#v12010-2022-10-18">v1.20.10</a>
<ul>
<li><strong>Documentation</strong>: Documentation Updates for Amazon Fraud Detector</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/sagemaker/CHANGELOG.md#v1480-2022-10-18">v1.48.0</a>
<ul>
<li><strong>Feature</strong>: This change allows customers to enable data capturing while running a batch transform job, and configure monitoring schedule to monitoring the captured data.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/servicediscovery</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/servicediscovery/CHANGELOG.md#v1180-2022-10-18">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Updated the ListNamespaces API to support the NAME and HTTP_NAME filters, and the BEGINS_WITH filter condition.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sesv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/sesv2/CHANGELOG.md#v1140-2022-10-18">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: This release allows subscribers to enable Dedicated IPs (managed) to send email via a fully managed dedicated IP experience. It also adds identities' VerificationStatus in the response of GetEmailIdentity and ListEmailIdentities APIs, and ImportJobs counts in the response of ListImportJobs API.</li>
</ul>
</li>
</ul>
<h1>Release (2022-10-17)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/greengrass</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/greengrass/CHANGELOG.md#v1140-2022-10-17">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: This change allows customers to specify FunctionRuntimeOverride in FunctionDefinitionVersion. This configuration can be used if the runtime on the device is different from the AWS Lambda runtime specified for that function.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/sagemaker/CHANGELOG.md#v1470-2022-10-17">v1.47.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for C7g, C6g, C6gd, C6gn, M6g, M6gd, R6g, and R6gn Graviton instance types in Amazon SageMaker Inference.</li>
</ul>
</li>
</ul>
<h1>Release (2022-10-14)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediaconvert</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.28.0/service/mediaconvert/CHANGELOG.md#v1260-2022-10-14">v1.26.0</a>
<ul>
<li><strong>Feature</strong>: MediaConvert now supports specifying the minimum percentage of the HRD buffer available at the end of each encoded video segment.</li>
</ul>
</li>
</ul>
<h1>Release (2022-10-13)</h1>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/17455fe071684b0c2ffcaea44aad9f20275e8b2f"><code>17455fe</code></a> Release 2022-10-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7cbfe8f471b88a4a2163c245f1e32478521f183e"><code>7cbfe8f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c06105c0b77a6d9d60a38148d7884221280ab341"><code>c06105c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2c1a13cb1e261c32579da4945ca195ac47acb896"><code>2c1a13c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32d9151b2ead0b88818e08bde9c0d5cc09dc0143"><code>32d9151</code></a> Release 2022-10-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0bf72ee7a3ed8ab6154f11cc688cde1ecd7e4df6"><code>0bf72ee</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a95f77ab06163ae95d97581408f72241bc8cf10"><code>5a95f77</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f630721a15604ad9d34f08302aefd34703ae6085"><code>f630721</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b45d0b6af38be20d852523fc6ed00c7fe959dc06"><code>b45d0b6</code></a> Release 2022-10-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c604be9d8956387d3ca26d68b9af37c1bf5cf642"><code>c604be9</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.27.2...service/s3/v1.28.0">compare view</a></li>
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

### Comment by @app-sre-bot on November 27, 2023 at 04:39 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on November 27, 2023 at 01:25 PM UTC

/test

### Comment by @jlsherrill on November 27, 2023 at 01:27 PM UTC

oktotest

### Comment by @jlsherrill on November 27, 2023 at 01:32 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on November 27, 2023 at 02:12 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/487*
