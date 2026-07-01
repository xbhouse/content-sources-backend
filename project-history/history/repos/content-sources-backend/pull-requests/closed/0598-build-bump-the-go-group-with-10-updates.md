---
type: pull_request
number: 598
title: "Build: Bump the go group with 10 updates"
state: closed
author: dependabot
created: 2024-03-06T13:57:35Z
updated: 2024-03-11T04:08:02Z
closed: 2024-03-11T04:08:00Z
base_branch: main
head_branch: dependabot/go_modules/go-a033e139d6
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/598
---

# Pull Request #598: Build: Bump the go group with 10 updates

**Author**: @dependabot
**Created**: March 06, 2024 at 01:57 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-a033e139d6`

## Description

Bumps the go group with 10 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.0-alpha.0-proton` | `1.1.0-alpha.1` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.18.0` | `1.19.0` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.42.2` | `1.43.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.25.0` | `1.25.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.0` | `1.17.6` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.33.1` | `1.34.2` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.3` | `5.5.4` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.5.0` | `9.5.1` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.0-alpha.0-proton to 1.1.0-alpha.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>v1.1.0-alpha.1</h2>
<h1>What's Changed</h1>
<p>Removes the <code>openpgp.VerifyDetachedSignatureAndSaltedHash</code> function and the <code>packet.SaltedHashSpecifier</code> as they are no longer required. They were introduced for verifying the headers in cleartext messages. However, in the latest crypto-refresh specification, cleartext message headers were dropped.</p>
<p>Full Changelog: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.0...v1.1.0-alpha.1">v1.1.0-alpha.0...v1.1.0-alpha.1</a></p>
<h2>v1.1.0-alpha.1-proton</h2>
<p>This pre-release is v1.1.0-alpha.1 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/9d2beb2b7bc7194bbaab3492b36b0538c18b40ef"><code>9d2beb2</code></a> Remove VerifyDetachedSignatureAndSaltedHash and SaltedHashSpecifier (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/196">#196</a>)</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.0...v1.1.0-alpha.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.18.0 to 1.19.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.19.0</h2>
<h2>What's Changed</h2>
<p>The module <code>prometheus/common v0.48.0</code> introduced an incompatibility when used together with client_golang (See <a href="https://redirect.github.com/prometheus/client_golang/pull/1448">prometheus/client_golang#1448</a> for more details). If your project uses client_golang and you want to use <code>prometheus/common v0.48.0</code> or higher, please update client_golang to v1.19.0.</p>
<ul>
<li>[CHANGE] Minimum required go version is now 1.20 (we also test client_golang against new 1.22 version). <a href="https://redirect.github.com/prometheus/client_golang/issues/1445">#1445</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1449">#1449</a></li>
<li>[FEATURE] collectors: Add version collector. <a href="https://redirect.github.com/prometheus/client_golang/issues/1422">#1422</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1427">#1427</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/michurin"><code>@​michurin</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1423">prometheus/client_golang#1423</a></li>
<li><a href="https://github.com/kavu"><code>@​kavu</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1445">prometheus/client_golang#1445</a></li>
<li><a href="https://github.com/ywwg"><code>@​ywwg</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1448">prometheus/client_golang#1448</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.18.0...v1.19.0">https://github.com/prometheus/client_golang/compare/v1.18.0...v1.19.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.19.0 / 2023-02-27</h2>
<p>The module <code>prometheus/common v0.48.0</code> introduced an incompatibility when used together with client_golang (See <a href="https://redirect.github.com/prometheus/client_golang/pull/1448">prometheus/client_golang#1448</a> for more details). If your project uses client_golang and you want to use <code>prometheus/common v0.48.0</code> or higher, please update client_golang to v1.19.0.</p>
<ul>
<li>[CHANGE] Minimum required go version is now 1.20 (we also test client_golang against new 1.22 version). <a href="https://redirect.github.com/prometheus/client_golang/issues/1445">#1445</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1449">#1449</a></li>
<li>[FEATURE] collectors: Add version collector. <a href="https://redirect.github.com/prometheus/client_golang/issues/1422">#1422</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1427">#1427</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/77d4003c72f054ac435df1223deac17b1f8858ea"><code>77d4003</code></a> Add 1.19.0 changelog (<a href="https://redirect.github.com/prometheus/client_golang/issues/1451">#1451</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/14259fa70cfb69f1262f69fdfe58ed5e6318d441"><code>14259fa</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1448">#1448</a> from ywwg/owilliams/content-negotiation</li>
<li><a href="https://github.com/prometheus/client_golang/commit/6d039205b8decc22868f43b0bd0da01b376a36aa"><code>6d03920</code></a> deps: bump prometheus/common version</li>
<li><a href="https://github.com/prometheus/client_golang/commit/353395b3b67b2bee0a219950bf5570779d74a392"><code>353395b</code></a> Remove support for go 1.19 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1449">#1449</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/9dd5d2a64af1e9bd0cbff0516ded6e51d25209bf"><code>9dd5d2a</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1445">#1445</a> from kavu/add_go122_metrics_test</li>
<li><a href="https://github.com/prometheus/client_golang/commit/c906a5e91a4604bd55fb8e26a54b5ba64a81c678"><code>c906a5e</code></a> Add support for Go 1.22</li>
<li><a href="https://github.com/prometheus/client_golang/commit/7ac90362b02729a65109b33d172bafb65d7dab50"><code>7ac9036</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1440">#1440</a> from prometheus/dependabot/github_actions/github-act...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/8c7e30ff0dae76cb87061a37bbccc6c8789196fa"><code>8c7e30f</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1441">#1441</a> from prometheus/dependabot/go_modules/tutorial/whats...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/08769f8257358282749a5180c9dc845f6e065640"><code>08769f8</code></a> Bump github.com/prometheus/common in /tutorial/whatsup</li>
<li><a href="https://github.com/prometheus/client_golang/commit/83d5940383d55377072d10a70316a7e24c5bb47c"><code>83d5940</code></a> Bump the github-actions group with 2 updates</li>
<li>Additional commits viewable in <a href="https://github.com/prometheus/client_golang/compare/v1.18.0...v1.19.0">compare view</a></li>
</ul>
</details>
<br />

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

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.0 to 1.17.6
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-03-10)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/ivschat</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/ivschat/CHANGELOG.md#v140-2023-03-10">v1.4.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new exception returned when calling AWS IVS chat UpdateLoggingConfiguration. Now UpdateLoggingConfiguration can return ConflictException when invalid updates are made in sequence to Logging Configurations.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/secretsmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/secretsmanager/CHANGELOG.md#v1190-2023-03-10">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: The type definitions of SecretString and SecretBinary now have a minimum length of 1 in the model to match the exception thrown when you pass in empty values.</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-09)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/codeartifact</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/codeartifact/CHANGELOG.md#v1170-2023-03-09">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces the generic package format, a mechanism for storing arbitrary binary assets. It also adds a new API, PublishPackageVersion, to allow for publishing generic packages.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/connect/CHANGELOG.md#v1490-2023-03-09">v1.49.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new API, GetMetricDataV2, which returns metric data for Amazon Connect.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/evidently</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/evidently/CHANGELOG.md#v1110-2023-03-09">v1.11.0</a>
<ul>
<li><strong>Feature</strong>: Updated entity override documentation</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/networkmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/networkmanager/CHANGELOG.md#v1175-2023-03-09">v1.17.5</a>
<ul>
<li><strong>Documentation</strong>: This update provides example usage for TransitGatewayRouteTableArn.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/quicksight</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/quicksight/CHANGELOG.md#v1330-2023-03-09">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: This release has two changes: add state persistence feature for embedded dashboard and console in GenerateEmbedUrlForRegisteredUser API; add properties for hidden collapsed row dimensions in PivotTableOptions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/redshiftdata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/redshiftdata/CHANGELOG.md#v1190-2023-03-09">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: Added support for Redshift Serverless workgroup-arn wherever the WorkgroupName parameter is available.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/sagemaker/CHANGELOG.md#v1710-2023-03-09">v1.71.0</a>
<ul>
<li><strong>Feature</strong>: Amazon SageMaker Inference now allows SSM access to customer's model container by setting the &quot;EnableSSMAccess&quot; parameter for a ProductionVariant in CreateEndpointConfig API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/servicediscovery</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/servicediscovery/CHANGELOG.md#v1200-2023-03-09">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Updated all AWS Cloud Map APIs to provide consistent throttling exception (RequestLimitExceeded)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sesv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/sesv2/CHANGELOG.md#v1170-2023-03-09">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces a new recommendation in Virtual Deliverability Manager Advisor, which detects missing or misconfigured Brand Indicator for Message Identification (BIMI) DNS records for customer sending identities.</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-08)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/athena</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/athena/CHANGELOG.md#v1230-2023-03-08">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: A new field SubstatementType is added to GetQueryExecution API, so customers have an error free way to detect the query type and interpret the result.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/dynamodb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/dynamodb/CHANGELOG.md#v1190-2023-03-08">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: Adds deletion protection support to DynamoDB tables. Tables with deletion protection enabled cannot be deleted. Deletion protection is disabled by default, can be enabled via the CreateTable or UpdateTable APIs, and is visible in TableDescription. This setting is not replicated for Global Tables.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/ec2/CHANGELOG.md#v1890-2023-03-08">v1.89.0</a>
<ul>
<li><strong>Feature</strong>: Introducing Amazon EC2 C7g, M7g and R7g instances, powered by the latest generation AWS Graviton3 processors and deliver up to 25% better performance over Graviton2-based instances.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lakeformation</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/lakeformation/CHANGELOG.md#v1200-2023-03-08">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release adds two new API support &quot;GetDataCellsFiler&quot; and &quot;UpdateDataCellsFilter&quot;, and also updates the corresponding documentation.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediapackage</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/mediapackage/CHANGELOG.md#v1210-2023-03-08">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: This release provides the date and time live resources were created.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediapackagevod</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.6/service/mediapackagevod/CHANGELOG.md#v1220-2023-03-08">v1.22.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18bff5e44f4c2c140feaf54b2dd92140630c7e50"><code>18bff5e</code></a> Release 2023-03-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cca7003bfc6acfa24313a1fa3ae5e4aa0a36f4bc"><code>cca7003</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/87bde6fa5748b6f78f1402052c55056f1f3d8ddf"><code>87bde6f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/361bd48e01611b81750b838ec44cabfa7f3c02ad"><code>361bd48</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/61d474d5d66f498c07f08846c308081c0cab2536"><code>61d474d</code></a> Fix staticcheck linter warnings (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2044">#2044</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59a58ae78b0f41b541d431664e98b0711e0af4ed"><code>59a58ae</code></a> Release 2023-03-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa1d8ae433652241f5ea9996f14e385f59a0c7e8"><code>aa1d8ae</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3441f5fbc73bf974fa4af4ad579bfb4a09752f8d"><code>3441f5f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/826cc366c86fba977ed6343ae22247ac722be0d8"><code>826cc36</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5799416de4970748df37d556db57a88e050d99c9"><code>5799416</code></a> Release 2023-03-08</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.17.0...v1.17.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.33.1 to 1.34.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4aeeb0d7a4293f0b31c2e0be83e65da6f7fd4ae2"><code>4aeeb0d</code></a> Release 2023-11-28.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e09e153704d4da6ae2bb0ae3875058950d31206b"><code>e09e153</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8293e2ca285a7333a9931a0a3b37b2bcf7cdc05d"><code>8293e2c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/713fb0f31a188015915d785556ebf16ddb08085a"><code>713fb0f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/830202d722c904c7e3da40e8dde7b9338d08752c"><code>830202d</code></a> Merge customizations for service s3</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2de0027dc478a6ae80e9f2d24d904a425169a23b"><code>2de0027</code></a> Release 2023-11-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0c890c5eaf354ff23feb727ded9f50aaee9f1c4"><code>f0c890c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e032d9ea8d98d366f2467a72834d2cc0ee865edd"><code>e032d9e</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/507661ff1edbc896fbdfe3ea2e4c2e74be3b4e3c"><code>507661f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4128360684a451476e33c0f979921bc46ff63656"><code>4128360</code></a> fix: respect functional option modifications to RetryMaxAttempts (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2390">#2390</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.33.1...service/ecs/v1.34.2">compare view</a></li>
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

Updates `github.com/jackc/pgx/v5` from 5.5.3 to 5.5.4
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.4 (March 4, 2024)</h1>
<p>Fix CVE-2024-27304</p>
<p>SQL injection can occur if an attacker can cause a single query or bind message to exceed 4 GB in size. An integer
overflow in the calculated message size can cause the one large message to be sent as multiple messages under the
attacker's control.</p>
<p>Thanks to Paul Gerste for reporting this issue.</p>
<ul>
<li>Fix behavior of CollectRows to return empty slice if Rows are empty (Felix)</li>
<li>Fix simple protocol encoding of json.RawMessage</li>
<li>Fix *Pipeline.getResults should close pipeline on error</li>
<li>Fix panic in TryFindUnderlyingTypeScanPlan (David Kurman)</li>
<li>Fix deallocation of invalidated cached statements in a transaction</li>
<li>Handle invalid sslkey file</li>
<li>Fix scan float4 into sql.Scanner</li>
<li>Fix pgtype.Bits not making copy of data from read buffer. This would cause the data to be corrupted by future reads.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/da6f2c98f2664b215b40b1606551fdfcc7f3ea5c"><code>da6f2c9</code></a> Update changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/c543134753a0c5d22881c12404025724cb05ffd8"><code>c543134</code></a> SQL sanitizer wraps arguments in parentheses</li>
<li><a href="https://github.com/jackc/pgx/commit/20344dfae83e672eff73a03324398543a1d87f43"><code>20344df</code></a> Check for overflow on uint16 sizes in pgproto3</li>
<li><a href="https://github.com/jackc/pgx/commit/adbb38f298c76e283ffc7c7a3f571036fea47fd4"><code>adbb38f</code></a> Do not allow protocol messages larger than ~1GB</li>
<li><a href="https://github.com/jackc/pgx/commit/c1b0a01ca75ac9eb3a7dbc1396f583ab5dbf9557"><code>c1b0a01</code></a> Fix behavior of CollectRows to return empty slice if Rows are empty</li>
<li><a href="https://github.com/jackc/pgx/commit/88dfc22ae4aa031783cda90841d5358edd85ff2c"><code>88dfc22</code></a> Fix simple protocol encoding of json.RawMessage</li>
<li><a href="https://github.com/jackc/pgx/commit/2e84dccaf57b4fa803149884f2149dfa9e923593"><code>2e84dcc</code></a> *Pipeline.getResults should close pipeline on error</li>
<li><a href="https://github.com/jackc/pgx/commit/d149d3fe5c50d1d98bd6265d3c928519ba4b3f4b"><code>d149d3f</code></a> Fix panic in TryFindUnderlyingTypeScanPlan</li>
<li><a href="https://github.com/jackc/pgx/commit/046f497efb4e92caa9575a0e9c351e4906af14c6"><code>046f497</code></a> deallocateInvalidatedCachedStatements now runs in transactions</li>
<li><a href="https://github.com/jackc/pgx/commit/8896bd697781ed4aee392daa90b90cde142319fe"><code>8896bd6</code></a> Handle invalid sslkey file</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.3...v5.5.4">compare view</a></li>
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

### Comment by @app-sre-bot on March 06, 2024 at 01:58 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on March 06, 2024 at 04:39 PM UTC

[test]

### Comment by @jlsherrill on March 07, 2024 at 05:46 PM UTC

[test]

### Comment by @dependabot on March 11, 2024 at 04:07 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/598*
