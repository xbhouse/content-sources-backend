---
type: pull_request
number: 1164
title: "Build: Bump the go group with 8 updates"
state: merged
author: dependabot
created: 2025-08-04T06:12:46Z
updated: 2025-08-04T14:26:12Z
closed: 2025-08-04T14:26:06Z
merged: 2025-08-04T14:26:06Z
base_branch: main
head_branch: dependabot/go_modules/go-96ea7dd8bc
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1164
---

# Pull Request #1164: Build: Bump the go group with 8 updates

**Author**: @dependabot
**Created**: August 04, 2025 at 06:12 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-96ea7dd8bc`

## Description

Bumps the go group with 8 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/swaggo/swag](https://github.com/swaggo/swag) | `1.16.5` | `1.16.6` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.22.0` | `1.23.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.6` | `1.37.1` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.18` | `1.30.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.71` | `1.18.2` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.53.1` | `1.54.1` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.84.1` | `1.85.1` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.34.1` | `0.35.0` |

Updates `github.com/swaggo/swag` from 1.16.5 to 1.16.6
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/swaggo/swag/releases">github.com/swaggo/swag's releases</a>.</em></p>
<blockquote>
<h2>v1.16.6</h2>
<h2>What's Changed</h2>
<ul>
<li>fix: prevent nil pointer dereference in getFuncDoc when parsing depen… by <a href="https://github.com/gatorjuice"><code>@​gatorjuice</code></a> in <a href="https://redirect.github.com/swaggo/swag/pull/2044">swaggo/swag#2044</a></li>
<li>fix: router with tilde <a href="https://redirect.github.com/swaggo/swag/issues/2004">#2004</a> by <a href="https://github.com/subwiz"><code>@​subwiz</code></a> in <a href="https://redirect.github.com/swaggo/swag/pull/2005">swaggo/swag#2005</a></li>
<li>Feature: allow enum ordered const name override (2nd PR for this) by <a href="https://github.com/drewsilcock"><code>@​drewsilcock</code></a> in <a href="https://redirect.github.com/swaggo/swag/pull/2046">swaggo/swag#2046</a></li>
<li>Use the structs name without the <a href="https://github.com/name"><code>@​name</code></a> comment by <a href="https://github.com/skast96"><code>@​skast96</code></a> in <a href="https://redirect.github.com/swaggo/swag/pull/2043">swaggo/swag#2043</a></li>
<li>feat: allow description line continuation by <a href="https://github.com/berk-karaal"><code>@​berk-karaal</code></a> in <a href="https://redirect.github.com/swaggo/swag/pull/2048">swaggo/swag#2048</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/gatorjuice"><code>@​gatorjuice</code></a> made their first contribution in <a href="https://redirect.github.com/swaggo/swag/pull/2044">swaggo/swag#2044</a></li>
<li><a href="https://github.com/subwiz"><code>@​subwiz</code></a> made their first contribution in <a href="https://redirect.github.com/swaggo/swag/pull/2005">swaggo/swag#2005</a></li>
<li><a href="https://github.com/drewsilcock"><code>@​drewsilcock</code></a> made their first contribution in <a href="https://redirect.github.com/swaggo/swag/pull/2046">swaggo/swag#2046</a></li>
<li><a href="https://github.com/skast96"><code>@​skast96</code></a> made their first contribution in <a href="https://redirect.github.com/swaggo/swag/pull/2043">swaggo/swag#2043</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/swaggo/swag/compare/v1.16.5...v1.16.6">https://github.com/swaggo/swag/compare/v1.16.5...v1.16.6</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/swaggo/swag/commit/f676981e12b892b4a2c5f50cd3dbc96d469ca358"><code>f676981</code></a> feat: allow description line continuation (<a href="https://redirect.github.com/swaggo/swag/issues/2048">#2048</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/252fecd4cbc33868a56d3f3f1092cd2a561073ef"><code>252fecd</code></a> Use the structs name without the <a href="https://github.com/name"><code>@​name</code></a> comment (<a href="https://redirect.github.com/swaggo/swag/issues/2043">#2043</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/b0c5cc990594386ba6189a21313e4e8e68695249"><code>b0c5cc9</code></a> Feature: allow enum ordered const name override (2nd PR for this) (<a href="https://redirect.github.com/swaggo/swag/issues/2046">#2046</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/0f3bf86377c7e1c5bbf280380419e9ded717d83a"><code>0f3bf86</code></a> fix: router with tilde <a href="https://redirect.github.com/swaggo/swag/issues/2004">#2004</a> (<a href="https://redirect.github.com/swaggo/swag/issues/2005">#2005</a>)</li>
<li><a href="https://github.com/swaggo/swag/commit/d7cec5730e4e9ff1c05e98b780c0092b0ad35cbe"><code>d7cec57</code></a> fix: prevent nil pointer dereference in getFuncDoc when parsing dependencies ...</li>
<li>See full diff in <a href="https://github.com/swaggo/swag/compare/v1.16.5...v1.16.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.22.0 to 1.23.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.23.0 - 2025-07-30</h2>
<ul>
<li>[CHANGE] Minimum required Go version is now 1.23, only the two latest Go versions are supported from now on. <a href="https://redirect.github.com/prometheus/client_golang/issues/1812">#1812</a></li>
<li>[FEATURE] Add WrapCollectorWith and WrapCollectorWithPrefix <a href="https://redirect.github.com/prometheus/client_golang/issues/1766">#1766</a></li>
<li>[FEATURE] Add exemplars for native histograms <a href="https://redirect.github.com/prometheus/client_golang/issues/1686">#1686</a></li>
<li>[ENHANCEMENT] exp/api: Bubble up status code from writeResponse <a href="https://redirect.github.com/prometheus/client_golang/issues/1823">#1823</a></li>
<li>[ENHANCEMENT] collector/go: Update runtime metrics for Go v1.23 and v1.24 <a href="https://redirect.github.com/prometheus/client_golang/issues/1833">#1833</a></li>
<li>[BUGFIX] exp/api: client prompt return on context cancellation <a href="https://redirect.github.com/prometheus/client_golang/issues/1729">#1729</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.23.0 / 2025-07-30</h2>
<ul>
<li>[CHANGE] Minimum required Go version is now 1.23, only the two latest Go versions are supported from now on. <a href="https://redirect.github.com/prometheus/client_golang/issues/1812">#1812</a></li>
<li>[FEATURE] Add WrapCollectorWith and WrapCollectorWithPrefix <a href="https://redirect.github.com/prometheus/client_golang/issues/1766">#1766</a></li>
<li>[FEATURE] Add exemplars for native histograms <a href="https://redirect.github.com/prometheus/client_golang/issues/1686">#1686</a></li>
<li>[ENHANCEMENT] exp/api: Bubble up status code from writeResponse <a href="https://redirect.github.com/prometheus/client_golang/issues/1823">#1823</a></li>
<li>[ENHANCEMENT] collector/go: Update runtime metrics for Go v1.23 and v1.24 <a href="https://redirect.github.com/prometheus/client_golang/issues/1833">#1833</a></li>
<li>[BUGFIX] exp/api: client prompt return on context cancellation <a href="https://redirect.github.com/prometheus/client_golang/issues/1729">#1729</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/e4b2208dd8cb6d1425f00250db842ec3c1e8749e"><code>e4b2208</code></a> Cut v1.23.0 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1848">#1848</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/d9492afd3a6f2e9782a7fc10363281bfd5b743bb"><code>d9492af</code></a> cut v1.23.0-rc.1 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1842">#1842</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/aeae8a0b4f54a8fa720d19b88638a2d048596f82"><code>aeae8a0</code></a> Cut v1.23.0-rc.0 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1837">#1837</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/b157309b723f0b8588ea604bb78dbbba196803f2"><code>b157309</code></a> Update common Prometheus files (<a href="https://redirect.github.com/prometheus/client_golang/issues/1832">#1832</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/a704e287f467b79744c30af996b7d710d4c6900d"><code>a704e28</code></a> build(deps): bump the github-actions group with 3 updates (<a href="https://redirect.github.com/prometheus/client_golang/issues/1826">#1826</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/c7743110ad2c599de3d8c23682d978a12f9f36d1"><code>c774311</code></a> Fix errNotImplemented reference (<a href="https://redirect.github.com/prometheus/client_golang/issues/1835">#1835</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/db4db7bb0065a76c75b7df6f61d2cf183ecfc473"><code>db4db7b</code></a> Update runtime metrics for Go v1.23 and v1.24 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1833">#1833</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/99d380ebfe865ae2c973c80184bc97ac0d98f083"><code>99d380e</code></a> Update common Prometheus files (<a href="https://redirect.github.com/prometheus/client_golang/issues/1831">#1831</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/f3ef320dcde30f31188c577ad71e6480e91bf464"><code>f3ef320</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1828">#1828</a> from prometheus/dependabot/go_modules/exp/github.com...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/520c91ae841ff7264b5379fe85e6215fc62734a6"><code>520c91a</code></a> build(deps): bump github.com/prometheus/common in /exp</li>
<li>Additional commits viewable in <a href="https://github.com/prometheus/client_golang/compare/v1.22.0...v1.23.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.36.6 to 1.37.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67148dbe156c9519ccc35f7e46f1df7367c67b81"><code>67148db</code></a> Release 2025-07-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b21029eeb2699fcba879fd3639464fbb887c0178"><code>b21029e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6c2dde6522c75df1cfeb807e6f7823f644298d78"><code>6c2dde6</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7b710ac3e2a805dc0c2efa8bead2abb66bfe6169"><code>7b710ac</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/179df1ca8e873d6374f4739e6bab67f3ed62a3c0"><code>179df1c</code></a> update CONTRIBUTING.md</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/456e51a2895e2ce6c712cebacebb10ad5a8f359e"><code>456e51a</code></a> fix changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ce8fece3fe625af2d8db3d673fcc8265af50485"><code>4ce8fec</code></a> Join errors when retry token is unavailable (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3121">#3121</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1eea3dbe07dffcef2fdba6f8e2206b32b91d8a83"><code>1eea3db</code></a> Release 2025-07-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b63d556f10e3a6a6a48458eeaf5b288feef79c58"><code>b63d556</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1258955cdd238d60ed14ed5ce44cedd6e514f41c"><code>1258955</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.36.6...v1.37.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.18 to 1.30.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/383fd26928547348efed0ee1f8914d9e7a1b0287"><code>383fd26</code></a> Release 2024-07-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a055f9d9e17eb7ef2206e8e37ba98d661d13e6a"><code>4a055f9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3457953351e6993f57f5ab8e373af8af70be45b"><code>e345795</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/52a10ac239cc1aab2d070fbb004f6445ca0ddcc0"><code>52a10ac</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/adab0de9e4ae75aee4f894c99fd1e2ce2de0035e"><code>adab0de</code></a> remove unused jmespath dependency from main module (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2707">#2707</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e07cc82b25692dce8f68e0b5bd0d0c5cdbcd279"><code>0e07cc8</code></a> Release 2024-07-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e3583451cd8a91dbca2cc22672c2cfa0c7860cf"><code>5e35834</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2a28a1eec2a1b6f12f813e4900430a9cc79b6c2"><code>a2a28a1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f5489732257180b7c19456abb050214c2f5cac26"><code>f548973</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7aad565a65dab9b37da7b49c1c79a699c49d85e"><code>e7aad56</code></a> Release 2024-07-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.18...v1.30.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.71 to 1.18.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-11-17)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplify</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/amplify/CHANGELOG.md#v1120-2022-11-17">v1.12.0</a>
<ul>
<li><strong>Feature</strong>: Adds a new value (WEB_COMPUTE) to the Platform enum that allows customers to create Amplify Apps with Server-Side Rendering support.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/appflow/CHANGELOG.md#v1190-2022-11-17">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: AppFlow simplifies the preparation and cataloging of SaaS data into the AWS Glue Data Catalog where your data can be discovered and accessed by AWS analytics and ML services. AppFlow now also supports data field partitioning and file size optimization to improve query performance and reduce cost.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appsync</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/appsync/CHANGELOG.md#v1160-2022-11-17">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces the APPSYNC_JS runtime, and adds support for JavaScript in AppSync functions and AppSync pipeline resolvers.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/databasemigrationservice/CHANGELOG.md#v1220-2022-11-17">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for Internet Protocol Version 6 (IPv6) on DMS Replication Instances</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/ec2/CHANGELOG.md#v1710-2022-11-17">v1.71.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new optional parameter &quot;privateIpAddress&quot; for the CreateNatGateway API. PrivateIPAddress will allow customers to select a custom Private IPv4 address instead of having it be auto-assigned.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/elasticloadbalancingv2/CHANGELOG.md#v11825-2022-11-17">v1.18.25</a>
<ul>
<li><strong>Documentation</strong>: Provides new target group attributes to turn on/off cross zone load balancing and configure target group health for Network Load Balancers and Application Load Balancers. Provides improvements to health check configuration for Network Load Balancers.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/emrserverless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/emrserverless/CHANGELOG.md#v140-2022-11-17">v1.4.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for AWS Graviton2 based applications. You can now select CPU architecture when creating new applications or updating existing ones.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ivschat</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/ivschat/CHANGELOG.md#v110-2022-11-17">v1.1.0</a>
<ul>
<li><strong>Feature</strong>: Adds LoggingConfiguration APIs for IVS Chat - a feature that allows customers to store and record sent messages in a chat room to S3 buckets, CloudWatch logs, or Kinesis firehose.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lambda</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/lambda/CHANGELOG.md#v1250-2022-11-17">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: Add Node 18 (nodejs18.x) support to AWS Lambda.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/personalize</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/personalize/CHANGELOG.md#v1220-2022-11-17">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: This release provides support for creation and use of metric attributions in AWS Personalize</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/polly</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/polly/CHANGELOG.md#v1200-2022-11-17">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Add two new neural voices - Ola (pl-PL) and Hala (ar-AE).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rum</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/rum/CHANGELOG.md#v180-2022-11-17">v1.8.0</a>
<ul>
<li><strong>Feature</strong>: CloudWatch RUM now supports custom events. To use custom events, create an app monitor or update an app monitor with CustomEvent Status as ENABLED.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3control</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/s3control/CHANGELOG.md#v1260-2022-11-17">v1.26.0</a>
<ul>
<li><strong>Feature</strong>: Added 34 new S3 Storage Lens metrics to support additional customer use cases.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/secretsmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/secretsmanager/CHANGELOG.md#v1167-2022-11-17">v1.16.7</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Secrets Manager.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/securityhub</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/securityhub/CHANGELOG.md#v1240-2022-11-17">v1.24.0</a>
<ul>
<li><strong>Feature</strong>: Added SourceLayerArn and SourceLayerHash field for security findings.  Updated AwsLambdaFunction Resource detail</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/servicecatalogappregistry/CHANGELOG.md#v1150-2022-11-17">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for tagged resource associations, which allows you to associate a group of resources with a defined resource tag key and value to the application.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sts</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/sts/CHANGELOG.md#v1174-2022-11-17">v1.17.4</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for AWS Security Token Service.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/textract</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/textract/CHANGELOG.md#v1180-2022-11-17">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for specifying and extracting information from documents using the Signatures feature within Analyze Document API</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/workspaces</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.2/service/workspaces/CHANGELOG.md#v1270-2022-11-17">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: The release introduces CreateStandbyWorkspaces, an API that allows you to create standby WorkSpaces associated with a primary WorkSpace in another Region. DescribeWorkspaces now includes related WorkSpaces properties. DescribeWorkspaceBundles and CreateWorkspaceBundle now return more bundle details.</li>
</ul>
</li>
</ul>
<h1>Release (2022-11-16)</h1>
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
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c38daedb0a1e630091b485a6bf4bfa70e0d3c8d7"><code>c38daed</code></a> Release 2022-11-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/81664391263eb5e42e76689f77528dede681daae"><code>8166439</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a8dc621f42cac4f12f468c1a09b223d641fa0ed"><code>9a8dc62</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb64cae577b688a904bd2079e2cd739fc4210072"><code>cb64cae</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/163bba00b3e3a344d5aa834393fd1644450d1828"><code>163bba0</code></a> Release 2022-11-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51ff054109df5135b16fe0ec3b3a1d46ac0cd2ce"><code>51ff054</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/76bd75134bd8ae9e13aa580391da1f3ba65897a9"><code>76bd751</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00c5ca09c5884ea7181866009102544cdb3b2e90"><code>00c5ca0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b04fed445a9ee5e7910ec954566343b0e40d75b"><code>0b04fed</code></a> Release 2022-11-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f9c367a48d5b40c64882852bbce44d83583fcf43"><code>f9c367a</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.71...config/v1.18.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.53.1 to 1.54.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c964dbd88bc9424349a9225b044ec2322e07f4cb"><code>c964dbd</code></a> Release 2024-05-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ce843953a9cb4c9d2dff75dd5a5c967655c4c8ef"><code>ce84395</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8de91199530fabf1702a2cefab3cee6e065879bc"><code>8de9119</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/05fcf66f17486b8ce9b7372c19910ad68a85a730"><code>05fcf66</code></a> internal: true up internal metrics collection for post-SRA middleware (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2642">#2642</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/114842762c6ffd78ebce1850e31a387992d10acf"><code>1148427</code></a> reformat signer/v4 package doc (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2640">#2640</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b238d3fff478bcd2bbcc17bb36cb10757a09e5a9"><code>b238d3f</code></a> Release 2024-05-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/85f8268a3440c70939ef041a3c44d366916a894a"><code>85f8268</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ebaee4cee8793c276e5933f09591869bed154dda"><code>ebaee4c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/593b9667bb973de276cc9369c2e4046e3f811334"><code>593b966</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bfb02f0947afa4f8fc15430eb4bc290e7ce55f5f"><code>bfb02f0</code></a> Merge customizations for S3</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.53.1...service/s3/v1.54.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.84.1 to 1.85.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67148dbe156c9519ccc35f7e46f1df7367c67b81"><code>67148db</code></a> Release 2025-07-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b21029eeb2699fcba879fd3639464fbb887c0178"><code>b21029e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6c2dde6522c75df1cfeb807e6f7823f644298d78"><code>6c2dde6</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7b710ac3e2a805dc0c2efa8bead2abb66bfe6169"><code>7b710ac</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/179df1ca8e873d6374f4739e6bab67f3ed62a3c0"><code>179df1c</code></a> update CONTRIBUTING.md</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/456e51a2895e2ce6c712cebacebb10ad5a8f359e"><code>456e51a</code></a> fix changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ce8fece3fe625af2d8db3d673fcc8265af50485"><code>4ce8fec</code></a> Join errors when retry token is unavailable (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3121">#3121</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1eea3dbe07dffcef2fdba6f8e2206b32b91d8a83"><code>1eea3db</code></a> Release 2025-07-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b63d556f10e3a6a6a48458eeaf5b288feef79c58"><code>b63d556</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1258955cdd238d60ed14ed5ce44cedd6e514f41c"><code>1258955</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.84.1...service/s3/v1.85.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.34.1 to 0.35.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.35.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Changes to the logging API (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1046">#1046</a>)</li>
</ul>
<p>The logging API now supports a fluent interface for structured logging with attributes:</p>
<pre lang="go"><code>// usage before
logger := sentry.NewLogger(ctx)
// attributes weren't being set permanently
logger.SetAttributes(
    attribute.String(&quot;version&quot;, &quot;1.0.0&quot;),
)
logger.Infof(ctx, &quot;Message with parameters %d and %d&quot;, 1, 2)
<p>// new behavior<br />
ctx := context.Background()<br />
logger := sentry.NewLogger(ctx)</p>
<p>// Set permanent attributes on the logger<br />
logger.SetAttributes(<br />
attribute.String(&quot;version&quot;, &quot;1.0.0&quot;),<br />
)</p>
<p>// Chain attributes on individual log entries<br />
logger.Info().<br />
String(&quot;key.string&quot;, &quot;value&quot;).<br />
Int(&quot;key.int&quot;, 42).<br />
Bool(&quot;key.bool&quot;, true).<br />
Emitf(&quot;Message with parameters %d and %d&quot;, 1, 2)<br />
</code></pre></p>
<h3>Bug Fixes</h3>
<ul>
<li>Correctly serialize <code>FailureIssueThreshold</code> and <code>RecoveryThreshold</code> onto check-in payloads (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1060">#1060</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.35.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Changes to the logging API (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1046">#1046</a>)</li>
</ul>
<p>The logging API now supports a fluent interface for structured logging with attributes:</p>
<pre lang="go"><code>// usage before
logger := sentry.NewLogger(ctx)
// attributes weren't being set permanently
logger.SetAttributes(
    attribute.String(&quot;version&quot;, &quot;1.0.0&quot;),
)
logger.Infof(ctx, &quot;Message with parameters %d and %d&quot;, 1, 2)
<p>// new behavior<br />
ctx := context.Background()<br />
logger := sentry.NewLogger(ctx)</p>
<p>// Set permanent attributes on the logger<br />
logger.SetAttributes(<br />
attribute.String(&quot;version&quot;, &quot;1.0.0&quot;),<br />
)</p>
<p>// Chain attributes on individual log entries<br />
logger.Info().<br />
String(&quot;key.string&quot;, &quot;value&quot;).<br />
Int(&quot;key.int&quot;, 42).<br />
Bool(&quot;key.bool&quot;, true).<br />
Emitf(&quot;Message with parameters %d and %d&quot;, 1, 2)<br />
</code></pre></p>
<h3>Bug Fixes</h3>
<ul>
<li>Correctly serialize <code>FailureIssueThreshold</code> and <code>RecoveryThreshold</code> onto check-in payloads (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1060">#1060</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/a3e3405def544188c1fb81f2d55a2537438b7a8e"><code>a3e3405</code></a> release: 0.35.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/31fcbbd13fd212fa4f06f3e09f00fcc00a8e2bb5"><code>31fcbbd</code></a> Prepare 0.35.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1065">#1065</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2a2b6405c5609da11bd32beac675c6f10b7e85d9"><code>2a2b640</code></a> Add .cursor rules (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1053">#1053</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/274de7b70a1ba3e5c6fbaf8c813c56d150100d9b"><code>274de7b</code></a> Improve logging API (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1046">#1046</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7b8885f9c013a3807225f5ce74527c2d89607490"><code>7b8885f</code></a> fix: include FailureIssueThreshold and RecoveryThreshold in MonitorConfig JSO...</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7b375b01bf57c349b4d0430d0574cd5012121b67"><code>7b375b0</code></a> Merge branch 'release/0.34.1'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.34.1...v0.35.0">compare view</a></li>
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

### Comment by @swadeley on August 04, 2025 at 07:26 AM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on August 04, 2025 at 02:25 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1164*
