---
type: pull_request
number: 603
title: "Build: Bump the go group with 11 updates"
state: merged
author: dependabot
created: 2024-03-13T18:44:53Z
updated: 2024-03-14T14:27:15Z
closed: 2024-03-14T14:27:11Z
merged: 2024-03-14T14:27:11Z
base_branch: main
head_branch: dependabot/go_modules/go-d5b7957ecc
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/603
---

# Pull Request #603: Build: Bump the go group with 11 updates

**Author**: @dependabot
**Created**: March 13, 2024 at 06:44 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-d5b7957ecc`

## Description

Bumps the go group with 11 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.0-alpha.0-proton` | `1.1.0-alpha.1` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.18.0` | `1.19.0` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.42.2` | `1.43.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.25.0` | `1.25.3` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.0` | `1.17.7` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.33.1` | `1.34.3` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.2` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.2` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.3.1709674773` | `2024.3.1710276616` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.3` | `5.5.5` |
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

Updates `github.com/aws/aws-sdk-go-v2` from 1.25.0 to 1.25.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be4380a86275f2ede9d0bba8e7537282cb299edc"><code>be4380a</code></a> Release 2024-03-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d811bc75a06e05b4fccd639d8853d7cf9e14d4f3"><code>d811bc7</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3e8224265f89d71568c710f25c715331d57efd1"><code>b3e8224</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50115544c73020c47a488909b56b547b68626d12"><code>5011554</code></a> regen for copying ddbstreams attribute value after dropping go-cmp</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6c816bc84eb006f8c9bee2c8dbdfab0be56fc69c"><code>6c816bc</code></a> dep: drop go-cmp dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2538">#2538</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/66663e7fd870ff2047adbe6b5b91280ad2bc0fc9"><code>66663e7</code></a> Release 2024-03-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9c613340b1888ba7b81b43421d327b47fc55edad"><code>9c61334</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50fd65b9f584729eed30bbf2ee312d2f0bdd04b0"><code>50fd65b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bd15f869d9963824d1d20e9f9c2f5ed67dab6c64"><code>bd15f86</code></a> Release 2024-03-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3bfebf4fe0c1494f0f0df6f9afbfe5496a8d0e39"><code>3bfebf4</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.25.0...v1.25.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.0 to 1.17.7
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-03-21)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/chimesdkmessaging/CHANGELOG.md#v1130-2023-03-21">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Chime SDK messaging customers can now manage streaming configuration for messaging data for archival and analysis.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cleanrooms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/cleanrooms/CHANGELOG.md#v110-2023-03-21">v1.1.0</a>
<ul>
<li><strong>Feature</strong>: GA Release of AWS Clean Rooms, Added Tagging Functionality</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/ec2/CHANGELOG.md#v1910-2023-03-21">v1.91.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for AWS Network Firewall, AWS PrivateLink, and Gateway Load Balancers to Amazon VPC Reachability Analyzer, and it makes the path destination optional as long as a destination address in the filter at source is provided.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/internal/s3shared</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/internal/s3shared/CHANGELOG.md#v1140-2023-03-21">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: port v1 sdk 100-continue http header customization for s3 PutObject/UploadPart request and enable user config</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotsitewise</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/iotsitewise/CHANGELOG.md#v1280-2023-03-21">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: Provide support for tagging of data streams and enabling tag based authorization for property alias</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mgn</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/mgn/CHANGELOG.md#v1180-2023-03-21">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces the Import and export feature and expansion of the post-launch actions</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/s3/CHANGELOG.md#v1310-2023-03-21">v1.31.0</a>
<ul>
<li><strong>Feature</strong>: port v1 sdk 100-continue http header customization for s3 PutObject/UploadPart request and enable user config</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-20)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationautoscaling</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/applicationautoscaling/CHANGELOG.md#v1190-2023-03-20">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: With this release customers can now tag their Application Auto Scaling registered targets with key-value pairs and manage IAM permissions for all the tagged resources centrally.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/neptune</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/neptune/CHANGELOG.md#v1200-2023-03-20">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release makes following few changes. db-cluster-identifier is now a required parameter of create-db-instance. describe-db-cluster will now return PendingModifiedValues and GlobalClusterIdentifier fields in the response.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3outposts</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/s3outposts/CHANGELOG.md#v1160-2023-03-20">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: S3 On Outposts added support for endpoint status, and a failed endpoint reason, if any</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/workdocs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/workdocs/CHANGELOG.md#v1140-2023-03-20">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new API, SearchResources, which enable users to search through metadata and content of folders, documents, document versions and comments in a WorkDocs site.</li>
</ul>
</li>
</ul>
<h1>Release (2023-03-17)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/billingconductor</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/billingconductor/CHANGELOG.md#v160-2023-03-17">v1.6.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new filter to ListAccountAssociations API and a new filter to ListBillingGroups API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/configservice/CHANGELOG.md#v1300-2023-03-17">v1.30.0</a>
<ul>
<li><strong>Feature</strong>: This release adds resourceType enums for types released from October 2022 through February 2023.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/service/databasemigrationservice/CHANGELOG.md#v1250-2023-03-17">v1.25.0</a>
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
<li><code>github.com/aws/aws-sdk-go-v2/config</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.7/config/CHANGELOG.md#v11818-2023-03-16">v1.18.18</a></li>
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.17.0...v1.17.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.33.1 to 1.34.3
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.33.1...service/emr/v1.34.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2` from 2.15.0 to 2.15.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.15.2</h2>
<h2>What's Changed</h2>
<ul>
<li>Patch for a potential security issue. See <a href="https://github.com/cloudevents/sdk-go/blob/HEAD/TBD">CVE-2024-28110</a>.</li>
<li>Note: this could be a breaking change for people if they purposely change golang's HTTP <code>DefaultClient</code>, or change the CloudEvents <code>Client</code> returned from <code>NewClient</code>, and expect those changes to be visible on other HTTP flows using those Clients. E.g. auth</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.1...v2.15.2">https://github.com/cloudevents/sdk-go/compare/v2.15.1...v2.15.2</a></p>
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
<li><a href="https://github.com/cloudevents/sdk-go/commit/de2f28370b0d2a0f64f92c0c6139fa4b8a7c3851"><code>de2f283</code></a> Merge pull request from GHSA-5pf6-2qwx-pxm2</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/c5f8d9dd259c4197305ece455ad501795c786f7f"><code>c5f8d9d</code></a> Update v2/protocol/http/protocol.go</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/c17d949c7001b7907c2a643771e673c2af317a60"><code>c17d949</code></a> Avoid modifying the DefaultClient's Transport</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/67e389964131d55d65cd14b4eb32d57a47312695"><code>67e3899</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1020">#1020</a> from duglin/oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f0061e0ad9a7071a091fc4c92c91f504c1462bc5"><code>f0061e0</code></a> oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/4cc6c2d62d63da00f5b00efd4147e7dd5e40ee4c"><code>4cc6c2d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1011">#1011</a> from cloudevents/dependabot/bundler/docs/bundler-sec...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/b6949b02324839338c4bb8fbc7031cd0fba55899"><code>b6949b0</code></a> Bump the bundler group across 1 directories with 1 update</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/df51395f567d75bc5e2ac51a1796ae81522d2e87"><code>df51395</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1016">#1016</a> from cloudevents/dependabot/github_actions/golangci/...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/1af6e066857109cba2b636642ffa10a30facfce3"><code>1af6e06</code></a> Bump golangci/golangci-lint-action from 3 to 4</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/2574a05f7b10376e4d536157f447f1d4f3ead380"><code>2574a05</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1013">#1013</a> from jafossum/fix-nats-typos</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/v2` from 2.15.0 to 2.15.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.15.2</h2>
<h2>What's Changed</h2>
<ul>
<li>Patch for a potential security issue. See <a href="https://github.com/cloudevents/sdk-go/blob/HEAD/TBD">CVE-2024-28110</a>.</li>
<li>Note: this could be a breaking change for people if they purposely change golang's HTTP <code>DefaultClient</code>, or change the CloudEvents <code>Client</code> returned from <code>NewClient</code>, and expect those changes to be visible on other HTTP flows using those Clients. E.g. auth</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.1...v2.15.2">https://github.com/cloudevents/sdk-go/compare/v2.15.1...v2.15.2</a></p>
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
<li><a href="https://github.com/cloudevents/sdk-go/commit/de2f28370b0d2a0f64f92c0c6139fa4b8a7c3851"><code>de2f283</code></a> Merge pull request from GHSA-5pf6-2qwx-pxm2</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/c5f8d9dd259c4197305ece455ad501795c786f7f"><code>c5f8d9d</code></a> Update v2/protocol/http/protocol.go</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/c17d949c7001b7907c2a643771e673c2af317a60"><code>c17d949</code></a> Avoid modifying the DefaultClient's Transport</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/67e389964131d55d65cd14b4eb32d57a47312695"><code>67e3899</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1020">#1020</a> from duglin/oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f0061e0ad9a7071a091fc4c92c91f504c1462bc5"><code>f0061e0</code></a> oops</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/4cc6c2d62d63da00f5b00efd4147e7dd5e40ee4c"><code>4cc6c2d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1011">#1011</a> from cloudevents/dependabot/bundler/docs/bundler-sec...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/b6949b02324839338c4bb8fbc7031cd0fba55899"><code>b6949b0</code></a> Bump the bundler group across 1 directories with 1 update</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/df51395f567d75bc5e2ac51a1796ae81522d2e87"><code>df51395</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1016">#1016</a> from cloudevents/dependabot/github_actions/golangci/...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/1af6e066857109cba2b636642ffa10a30facfce3"><code>1af6e06</code></a> Bump golangci/golangci-lint-action from 3 to 4</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/2574a05f7b10376e4d536157f447f1d4f3ead380"><code>2574a05</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1013">#1013</a> from jafossum/fix-nats-typos</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.0...v2.15.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.3.1709674773 to 2024.3.1710276616
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/0a8e7366a027b5ac48e89cb30d8885f7c572ffec"><code>0a8e736</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27aaa595e8dfb7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/3bdbe8a52bd1877e7dc4aecea23f5c6a91f02d5a"><code>3bdbe8a</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a9724e5dbd34f57bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/e9365492fe0339246b58ee93f0b222bcb8aa503d"><code>e936549</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27aa89565b91b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/e6adc2e244308bff36dd558457482fe7c8ef978a"><code>e6adc2e</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73beb94ad5f37d2e4353b86e2...</li>
<li><a href="https://github.com/content-services/zest/commit/0e66ecf734ca1db58dded3b15d6e67d27776921a"><code>0e66ecf</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/11">#11</a> from jlsherrill/unknown</li>
<li><a href="https://github.com/content-services/zest/commit/75d6014a406b116126b7382b1e9a3e3c4d2ceff8"><code>75d6014</code></a> ignore unknown properties</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.3.1709674773...release/v2024.3.1710276616">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.5.3 to 5.5.5
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.5 (March 9, 2024)</h1>
<p>Use spaces instead of parentheses for SQL sanitization.</p>
<p>This still solves the problem of negative numbers creating a line comment, but this avoids breaking edge cases such as
<code>set foo to $1</code> where the substitution is taking place in a location where an arbitrary expression is not allowed.</p>
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
<li><a href="https://github.com/jackc/pgx/commit/78a0a2bf410b28c08359fc8c7550c1200589521c"><code>78a0a2b</code></a> Fix spelling in changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/a17f064492d5e560304aefde2784ef9253f1d0ce"><code>a17f064</code></a> Update changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/49b6aad319f125cd3016b1c00db015d6ca8772db"><code>49b6aad</code></a> Use spaces instead of parentheses for SQL sanitization</li>
<li><a href="https://github.com/jackc/pgx/commit/0cc4c14e620d484d555f2955fb39cfecea89aaa3"><code>0cc4c14</code></a> Add test to validate CollectRows for empty Rows</li>
<li><a href="https://github.com/jackc/pgx/commit/da6f2c98f2664b215b40b1606551fdfcc7f3ea5c"><code>da6f2c9</code></a> Update changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/c543134753a0c5d22881c12404025724cb05ffd8"><code>c543134</code></a> SQL sanitizer wraps arguments in parentheses</li>
<li><a href="https://github.com/jackc/pgx/commit/20344dfae83e672eff73a03324398543a1d87f43"><code>20344df</code></a> Check for overflow on uint16 sizes in pgproto3</li>
<li><a href="https://github.com/jackc/pgx/commit/adbb38f298c76e283ffc7c7a3f571036fea47fd4"><code>adbb38f</code></a> Do not allow protocol messages larger than ~1GB</li>
<li><a href="https://github.com/jackc/pgx/commit/c1b0a01ca75ac9eb3a7dbc1396f583ab5dbf9557"><code>c1b0a01</code></a> Fix behavior of CollectRows to return empty slice if Rows are empty</li>
<li><a href="https://github.com/jackc/pgx/commit/88dfc22ae4aa031783cda90841d5358edd85ff2c"><code>88dfc22</code></a> Fix simple protocol encoding of json.RawMessage</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.3...v5.5.5">compare view</a></li>
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

### Comment by @app-sre-bot on March 13, 2024 at 06:45 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on March 13, 2024 at 06:49 PM UTC

[test]

### Comment by @swadeley on March 13, 2024 at 09:00 PM UTC

/retest

### Comment by @swadeley on March 14, 2024 at 08:55 AM UTC

/retest

### Comment by @swadeley on March 14, 2024 at 02:10 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on March 14, 2024 at 02:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/603*
