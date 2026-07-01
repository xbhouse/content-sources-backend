---
type: pull_request
number: 585
title: "Build: Bump the go group with 8 updates"
state: closed
author: dependabot
created: 2024-02-26T05:01:52Z
updated: 2024-02-28T15:09:48Z
closed: 2024-02-28T15:09:46Z
base_branch: main
head_branch: dependabot/go_modules/go-bb6d1fc9a8
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/585
---

# Pull Request #585: Build: Bump the go group with 8 updates

**Author**: @dependabot
**Created**: February 26, 2024 at 05:01 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-bb6d1fc9a8`

## Description

Bumps the go group with 8 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.42.2` | `1.43.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.25.0` | `1.25.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.0` | `1.17.4` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.33.1` | `1.34.1` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.2.1707758687` | `2024.2.1708559921` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.5.0` | `9.5.1` |

Updates `github.com/IBM/sarama` from 1.42.2 to 1.43.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.43.0 (2024-02-22)</h2>
<!-- raw HTML omitted -->
<blockquote>
<p>[!NOTE]<br />
The go.mod directive has been bumped to 1.19 as the minimum version of Go required for the module. This was necessary to continue to receive updates from some of the third party dependencies that Sarama makes use of for compression.</p>
</blockquote>
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: update go directive to 1.19 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2795">IBM/sarama#2795</a></li>
<li>feat: add BuildSpnFunc to GSSAPIConfig for allow custom spn by <a href="https://github.com/fooofei"><code>@​fooofei</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2807">IBM/sarama#2807</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>Use %v formatting words and remove unnecessary newline by <a href="https://github.com/puellanivis"><code>@​puellanivis</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2802">IBM/sarama#2802</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump github.com/klauspost/compress from 1.16.7 to 1.17.6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2784">IBM/sarama#2784</a></li>
<li>chore(deps): bump github.com/eapache/go-resiliency from 1.5.0 to 1.6.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2810">IBM/sarama#2810</a></li>
<li>chore(deps): bump github.com/klauspost/compress from 1.17.6 to 1.17.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2811">IBM/sarama#2811</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(doc): add v1.42.2 to CHANGELOG.md by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2796">IBM/sarama#2796</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/puellanivis"><code>@​puellanivis</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2802">IBM/sarama#2802</a></li>
<li><a href="https://github.com/fooofei"><code>@​fooofei</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2807">IBM/sarama#2807</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.2...v1.43.0">https://github.com/IBM/sarama/compare/v1.42.2...v1.43.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/71f19e408a4daadc0a8ae2377a2a00e7bdacf4dc"><code>71f19e4</code></a> chore(deps): bump github.com/klauspost/compress from 1.17.6 to 1.17.7 (<a href="https://redirect.github.com/IBM/sarama/issues/2811">#2811</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/9d8250f2caa61e77b12d2c887acf621f74a5fea7"><code>9d8250f</code></a> chore(deps): bump github.com/eapache/go-resiliency from 1.5.0 to 1.6.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2810">#2810</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/b89bd604c3536603d5d61b24946b7ed30bb93506"><code>b89bd60</code></a> chore(ci): bump actions/dependency-review-action from 4.0.0 to 4.1.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2809">#2809</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/c850b24c8f0e1348773f00564e9c7081cfbbe505"><code>c850b24</code></a> chore(ci): bump github/codeql-action from 3.24.0 to 3.24.3 (<a href="https://redirect.github.com/IBM/sarama/issues/2808">#2808</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/38f2d151cf92c34799c4d88eddd4b823c361e33b"><code>38f2d15</code></a> feat: add BuildSpnFunc to GSSAPIConfig for allow custom spn (<a href="https://redirect.github.com/IBM/sarama/issues/2807">#2807</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3e385a677e5b0aaacc8ba0a56be18a53550275ca"><code>3e385a6</code></a> chore(ci): bump golangci/golangci-lint-action from 3.7.0 to 4.0.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2800">#2800</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/d282073f6bbc647c140faac2e1f06a6efdfd7c8f"><code>d282073</code></a> chore(deps): bump github.com/klauspost/compress from 1.16.7 to 1.17.6 (<a href="https://redirect.github.com/IBM/sarama/issues/2784">#2784</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/09923287b9e939d1e9942935aed6825d5eb3b489"><code>0992328</code></a> feat: update go directive to 1.19 (<a href="https://redirect.github.com/IBM/sarama/issues/2795">#2795</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/dba621b316b139c3c83412157d672b9f9fccd3dc"><code>dba621b</code></a> chore(doc): add v1.42.2 to CHANGELOG.md (<a href="https://redirect.github.com/IBM/sarama/issues/2796">#2796</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/33633d891de73c74ba05138b83d37233aa83850f"><code>33633d8</code></a> Use %v formatting words and remove unnecessary newline (<a href="https://redirect.github.com/IBM/sarama/issues/2802">#2802</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.42.2...v1.43.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.25.0 to 1.25.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7095341032d4a2b1fe04aa5e96e254321169e114"><code>7095341</code></a> Release 2024-02-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bcf04e6e66831a0fde59d0060af9a1d71158c9d5"><code>bcf04e6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ca190b054e7a07bcde270356ffe0ad06fd722ccc"><code>ca190b0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6397a6448e500f7386d005ad2fc3a822f3eaf380"><code>6397a64</code></a> move common middleware stack ops to service client modules (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2516">#2516</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e9b00e26a17fcdb0b01772ad19fc6f733abac749"><code>e9b00e2</code></a> Release 2024-02-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cfc53e095ae2e97185c6f594d1725320f152304"><code>0cfc53e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/92d6c194cfde05280d8836ef7b56c36fd7fd926d"><code>92d6c19</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18adb3114cc789def28fac1b718657d33443ed5c"><code>18adb31</code></a> add middleware snapshot tests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2513">#2513</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7888f0e10eb7ef5db0b67a8a27bb8eeee984257b"><code>7888f0e</code></a> Release 2024-02-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a9cd1dd63c4532e0afe4e5314750bf08f1ae68b"><code>4a9cd1d</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.25.0...v1.25.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.0 to 1.17.4
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-02-03)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/autoscaling</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/autoscaling/CHANGELOG.md#v1262-2023-02-03">v1.26.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudformation</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/cloudformation/CHANGELOG.md#v1262-2023-02-03">v1.26.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudsearch</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/cloudsearch/CHANGELOG.md#v1141-2023-02-03">v1.14.1</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudwatch</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/cloudwatch/CHANGELOG.md#v1252-2023-02-03">v1.25.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/docdb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/docdb/CHANGELOG.md#v1202-2023-02-03">v1.20.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/ec2/CHANGELOG.md#v1841-2023-02-03">v1.84.1</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticache</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/elasticache/CHANGELOG.md#v1262-2023-02-03">v1.26.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/elasticbeanstalk/CHANGELOG.md#v1151-2023-02-03">v1.15.1</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/elasticloadbalancing/CHANGELOG.md#v1152-2023-02-03">v1.15.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/elasticloadbalancingv2/CHANGELOG.md#v1193-2023-02-03">v1.19.3</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iam</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/iam/CHANGELOG.md#v1192-2023-02-03">v1.19.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/neptune</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/neptune/CHANGELOG.md#v1192-2023-02-03">v1.19.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/proton</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/proton/CHANGELOG.md#v1200-2023-02-03">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Add new GetResourcesSummary API</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/rds/CHANGELOG.md#v1402-2023-02-03">v1.40.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/redshift</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/redshift/CHANGELOG.md#v1272-2023-02-03">v1.27.2</a>
<ul>
<li><strong>Documentation</strong>: Corrects descriptions of the parameters for the API operations RestoreFromClusterSnapshot, RestoreTableFromClusterSnapshot, and CreateCluster.</li>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ses</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/ses/CHANGELOG.md#v1151-2023-02-03">v1.15.1</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sns</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/sns/CHANGELOG.md#v1201-2023-02-03">v1.20.1</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sqs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/sqs/CHANGELOG.md#v1202-2023-02-03">v1.20.2</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sts</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/sts/CHANGELOG.md#v1183-2023-02-03">v1.18.3</a>
<ul>
<li><strong>Dependency Update</strong>: Upgrade smithy to 1.27.2 and correct empty query list serialization.</li>
</ul>
</li>
</ul>
<h1>Release (2023-02-02)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appconfig</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.4/service/appconfig/CHANGELOG.md#v1160-2023-02-02">v1.16.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3f28b5134e7e9a047147b5773af62c6012c34df6"><code>3f28b51</code></a> Release 2023-02-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e8d17fd21083f2ec6ec6139f89b35c841994932"><code>6e8d17f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/60dbdbb0da35b8b8374e0b4e1a9536f1aed7e458"><code>60dbdbb</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/212910ac25d59e959c0a534bd6264a13fb9ca8c8"><code>212910a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb8cb66b4433cd5f125abc7c6b74de39c19f25fb"><code>eb8cb66</code></a> Upgrade smithy to 1.27.2, correct query empty list serialization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/24db9f5f6e6d387f115ee8d9393ae6d158e8ef0c"><code>24db9f5</code></a> Update processcreds.CredentialProcessResponse visibility to public (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1921">#1921</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bd3003e29f1fdc34859fcd194a61d6655b8ea492"><code>bd3003e</code></a> dependency: upgrade smithy to 1.27.2 and correct query empty list serialization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d94f223e8c3dc3c63d001ef443a25c056d4131e"><code>0d94f22</code></a> Release 2023-02-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2eec85ed13c74cf77315398edf974481fb5f4dd8"><code>2eec85e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ca6e32eedbb1e707ad3675879e17782e93db67e"><code>4ca6e32</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.17.0...v1.17.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.33.1 to 1.34.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-06-15)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/auditmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/auditmanager/CHANGELOG.md#v1250-2023-06-15">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces 2 Audit Manager features: CSV exports and new manual evidence options. You can now export your evidence finder results in CSV format. In addition, you can now add manual evidence to a control by entering free-form text or uploading a file from your browser.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/efs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/efs/CHANGELOG.md#v1203-2023-06-15">v1.20.3</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for EFS.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/guardduty</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/guardduty/CHANGELOG.md#v1232-2023-06-15">v1.23.2</a>
<ul>
<li><strong>Documentation</strong>: Updated descriptions for some APIs.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/location</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/location/CHANGELOG.md#v1240-2023-06-15">v1.24.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Location Service adds categories to places, including filtering on those categories in searches. Also, you can now add metadata properties to your geofences.</li>
</ul>
</li>
</ul>
<h1>Release (2023-06-13)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudtrail</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/cloudtrail/CHANGELOG.md#v1270-2023-06-13">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: This feature allows users to view dashboards for CloudTrail Lake event data stores.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/codegurusecurity</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/codegurusecurity/CHANGELOG.md#v100-2023-06-13">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: Initial release of Amazon CodeGuru Security APIs</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/drs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/drs/CHANGELOG.md#v1140-2023-06-13">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Added APIs to support network replication and recovery using AWS Elastic Disaster Recovery.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/ec2/CHANGELOG.md#v11000-2023-06-13">v1.100.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces a new feature, EC2 Instance Connect Endpoint, that enables you to connect to a resource over TCP, without requiring the resource to have a public IPv4 address.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/imagebuilder</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/imagebuilder/CHANGELOG.md#v1235-2023-06-13">v1.23.5</a>
<ul>
<li><strong>Documentation</strong>: Change the Image Builder ImagePipeline dateNextRun field to more accurately describe the data.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lightsail</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/lightsail/CHANGELOG.md#v1270-2023-06-13">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: This release adds pagination for the Get Certificates API operation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/s3/CHANGELOG.md#v1340-2023-06-13">v1.34.0</a>
<ul>
<li><strong>Feature</strong>: Integrate double encryption feature to SDKs.</li>
<li><strong>Bug Fix</strong>: Fix HeadObject to return types.Nound when an object does not exist. Fixes <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2084">2084</a></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/securityhub</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/securityhub/CHANGELOG.md#v1330-2023-06-13">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: Add support for Security Hub Automation Rules</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/simspaceweaver</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/simspaceweaver/CHANGELOG.md#v130-2023-06-13">v1.3.0</a>
<ul>
<li><strong>Feature</strong>: This release fixes using aws-us-gov ARNs in API calls and adds documentation for snapshot APIs.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/verifiedpermissions</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/verifiedpermissions/CHANGELOG.md#v100-2023-06-13">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: GA release of Amazon Verified Permissions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/wafv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/wafv2/CHANGELOG.md#v1350-2023-06-13">v1.35.0</a>
<ul>
<li><strong>Feature</strong>: You can now detect and block fraudulent account creation attempts with the new AWS WAF Fraud Control account creation fraud prevention (ACFP) managed rule group AWSManagedRulesACFPRuleSet.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/wellarchitected</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.34.1/service/wellarchitected/CHANGELOG.md#v1210-2023-06-13">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: AWS Well-Architected now supports Profiles that help customers prioritize which questions to focus on first by providing a list of prioritized questions that are better aligned with their business goals and outcomes.</li>
</ul>
</li>
</ul>
<h1>Release (2023-06-12)</h1>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c92ebc88890d9400122d55e2fc458931e32423ae"><code>c92ebc8</code></a> Release 2023-06-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3488004c6c0a1ab352f4e05d492f6078c1a86dcf"><code>3488004</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ee898185413e0b3f92f599dc991f492d13c2b27"><code>3ee8981</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/09d9f49930ba1d3f5070cc14467a817b8ea0d65d"><code>09d9f49</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/844ff45cdc76182229ad098c95bf3f5ab8c20e9f"><code>844ff45</code></a> Release 2023-06-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cc294064366687ef764f7bb73a7b08d7222a4f28"><code>cc29406</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa907c6969a3d103e6659d9f098f55d13dfde232"><code>aa907c6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f4a5607dec5b0efe12534373d24bc8885b54531"><code>2f4a560</code></a> chore: typo in Retryer interface (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2148">#2148</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba5a08fbe430c7930ad2227785226200adecd20c"><code>ba5a08f</code></a> fix: S3 HeadObject NotFound error never being returned (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2139">#2139</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9fd4e02a39d0870cef2223437ca2d5daff639b17"><code>9fd4e02</code></a> Release 2023-06-12</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.33.1...service/s3/v1.34.1">compare view</a></li>
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

Updates `github.com/content-services/zest/release/v2024` from 2024.2.1707758687 to 2024.2.1708559921
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/c973bda5ab87f83c30b313e0e334fb67ee9b129d"><code>c973bda</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a452768989ab35f27b8e3dae95843...</li>
<li><a href="https://github.com/content-services/zest/commit/64c8aea5e6de57a89b96fe44ad79aff09bf251c3"><code>64c8aea</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e4759de4b65e0b7a98d596b8e32...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.2.1707758687...release/v2024.2.1708559921">compare view</a></li>
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

### Comment by @app-sre-bot on February 26, 2024 at 05:03 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on February 26, 2024 at 02:34 PM UTC

oktotest

### Comment by @jlsherrill on February 26, 2024 at 02:53 PM UTC

[test]

### Comment by @jlsherrill on February 27, 2024 at 04:05 PM UTC

[test]

### Comment by @dependabot on February 28, 2024 at 03:09 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/585*
