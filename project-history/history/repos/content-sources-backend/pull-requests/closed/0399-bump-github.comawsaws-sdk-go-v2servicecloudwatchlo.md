---
type: pull_request
number: 399
title: "Bump github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs from 1.23.5 to 1.24.0"
state: closed
author: dependabot
created: 2023-09-21T16:19:06Z
updated: 2023-09-26T14:23:46Z
closed: 2023-09-26T14:23:44Z
base_branch: main
head_branch: dependabot/go_modules/github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs-1.24.0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/399
---

# Pull Request #399: Bump github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs from 1.23.5 to 1.24.0

**Author**: @dependabot
**Created**: September 21, 2023 at 04:19 PM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs-1.24.0`

## Description

Bumps [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) from 1.23.5 to 1.24.0.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-01-14)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Feature</strong>: Updated <code>github.com/aws/smithy-go</code> to latest version</li>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.13.0
<ul>
<li><strong>Bug Fix</strong>: Updates the Retry middleware to release the retry token, on subsequent attempts. This fixes <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1413">#1413</a>, and is based on PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1424">#1424</a></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/feature/dynamodb/attributevalue/CHANGELOG.md#v160-2022-01-14">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: Adds new MarshalWithOptions and UnmarshalWithOptions helpers allowing Encoding and Decoding options to be specified when serializing AttributeValues. Addresses issue: <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1494">aws/aws-sdk-go-v2#1494</a></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/feature/dynamodbstreams/attributevalue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/feature/dynamodbstreams/attributevalue/CHANGELOG.md#v160-2022-01-14">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: Adds new MarshalWithOptions and UnmarshalWithOptions helpers allowing Encoding and Decoding options to be specified when serializing AttributeValues. Addresses issue: <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1494">aws/aws-sdk-go-v2#1494</a></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appsync</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/appsync/CHANGELOG.md#v1120-2022-01-14">v1.12.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/autoscalingplans</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/autoscalingplans/CHANGELOG.md#v1100-2022-01-14">v1.10.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/computeoptimizer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/computeoptimizer/CHANGELOG.md#v1150-2022-01-14">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/costexplorer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/costexplorer/CHANGELOG.md#v1150-2022-01-14">v1.15.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/databasemigrationservice/CHANGELOG.md#v1160-2022-01-14">v1.16.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databrew</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/databrew/CHANGELOG.md#v1160-2022-01-14">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/ec2/CHANGELOG.md#v1280-2022-01-14">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticache</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/elasticache/CHANGELOG.md#v1180-2022-01-14">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticsearchservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/elasticsearchservice/CHANGELOG.md#v1130-2022-01-14">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/finspacedata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/finspacedata/CHANGELOG.md#v180-2022-01-14">v1.8.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/fms/CHANGELOG.md#v1130-2022-01-14">v1.13.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/glue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/glue/CHANGELOG.md#v1190-2022-01-14">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/honeycode</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/honeycode/CHANGELOG.md#v190-2022-01-14">v1.9.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/identitystore</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/identitystore/CHANGELOG.md#v1120-2022-01-14">v1.12.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ioteventsdata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/ioteventsdata/CHANGELOG.md#v190-2022-01-14">v1.9.0</a>
<ul>
<li><strong>Documentation</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotwireless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/iotwireless/CHANGELOG.md#v1160-2022-01-14">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kendra</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/kendra/CHANGELOG.md#v1200-2022-01-14">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lexmodelsv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/lexmodelsv2/CHANGELOG.md#v1170-2022-01-14">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Updated API models</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lexruntimev2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.24.0/service/lexruntimev2/CHANGELOG.md#v1120-2022-01-14">v1.12.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e10c0d2c8db721bd0e3b16070f3a74f6bc7171de"><code>e10c0d2</code></a> Release 2022-01-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35ab85c1330d3a1f21d0329ba0e1950dfc29e6d1"><code>35ab85c</code></a> Update smithy-go dependency version</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4741932878bc94de3d7c8028e726bb0f9e6c2fcb"><code>4741932</code></a> Fix Retry middleware not releasing retry token (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1560">#1560</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4d9a886e7fd465519a26a770c4b007d6fefae70"><code>e4d9a88</code></a> Update API Models (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1559">#1559</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1b62ab3b6ed72021b2fd081664db46911baa241"><code>c1b62ab</code></a> SDK DefaultsMode Configuration Implementation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1553">#1553</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/07e5d3e7b6286c00376154db25a3e750df5bd0ff"><code>07e5d3e</code></a> i1494 passing opts to attributevalue marshalling (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1495">#1495</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/576b4151f1db87570d680a61c737d816d79f3ab7"><code>576b415</code></a> codegen: Adds smithy-cli dependency to prevent auto use of smithy-cli 1.16 (#...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ec62266ee884956826f40ea3e70428802dd5b9c"><code>3ec6226</code></a> Release 2022-01-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50d7f6a7d5a3d0470e53f1f58357c88aaa391d84"><code>50d7f6a</code></a> Update smithy-go dependency version</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d25393b3864c79522946cf7c00dc30c78429a34"><code>2d25393</code></a> config: Add WithCredentialCacheOptions for LoadDefaultConfig's LoadOptions (#...</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.23.5...service/s3/v1.24.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs&package-manager=go_modules&previous-version=1.23.5&new-version=1.24.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)


</details>

---

## Discussion

### Comment by @app-sre-bot on September 21, 2023 at 04:20 PM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on September 26, 2023 at 02:23 PM UTC

Looks like github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs is up-to-date now, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/399*
