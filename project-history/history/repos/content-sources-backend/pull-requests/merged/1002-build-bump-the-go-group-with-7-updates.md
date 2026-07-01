---
type: pull_request
number: 1002
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2025-03-03T04:17:32Z
updated: 2025-03-03T08:15:05Z
closed: 2025-03-03T08:14:57Z
merged: 2025-03-03T08:14:57Z
base_branch: main
head_branch: dependabot/go_modules/go-2fa1be97d5
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1002
---

# Pull Request #1002: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: March 03, 2025 at 04:17 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-2fa1be97d5`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.5` | `1.1.6` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.45.0` | `1.45.1` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.2` | `1.36.3` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.7` | `1.29.8` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.60` | `1.17.61` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.45.14` | `1.46.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.77.1` | `1.78.0` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.5 to 1.1.6
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.1.6</h2>
<h2>What's Changed</h2>
<ul>
<li>Fix <code>PublicKey.KeyIdString</code> to return a valid key id by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/269">ProtonMail/go-crypto#269</a></li>
<li>Allow Key Flags override <a href="https://github.com/davrux"><code>@​davrux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/272">ProtonMail/go-crypto#272</a></li>
<li>Only check that message signatures are newer than the key by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/275">ProtonMail/go-crypto#275</a></li>
<li>openpgp/clearsign: just use rand.Reader in tests by <a href="https://github.com/mdosch"><code>@​mdosch</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/276">ProtonMail/go-crypto#276</a></li>
<li>Make Issuer Key ID signature subpacket non-critical by <a href="https://github.com/caarlos0"><code>@​caarlos0</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/266">ProtonMail/go-crypto#266</a></li>
<li>v2 API: Improve error messages for encryption key selection by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/271">ProtonMail/go-crypto#271</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.5...v1.1.6">https://github.com/ProtonMail/go-crypto/compare/v1.1.5...v1.1.6</a></p>
<h2>Release v1.1.6-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.6 with support for the following non-standardized features:</p>
<ul>
<li>Presistent symmetric keys <a href="https://www.ietf.org/archive/id/draft-ietf-openpgp-persistent-symmetric-keys-00.html">draft-ietf-openpgp-persistent-symmetric-keys-00</a></li>
<li>Automatic forwarding <a href="https://www.ietf.org/archive/id/draft-wussler-openpgp-forwarding-00.html">draft-wussler-openpgp-forwarding-00</a></li>
<li>Post-quantum algorithms <a href="https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/">draft-ietf-openpgp-pqc</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/e52eada5c60c4406d02e11195d91d46f0356beda"><code>e52eada</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/271">#271</a> from ProtonMail/feat/improve-errors-key-selection</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/4bf9d904d3586c894d7de3c53b01b05fe740a6da"><code>4bf9d90</code></a> feat(v2): Improve error message for encryption key selection</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/d47bb3817800ff5bfab2fb11f38534701bed32cb"><code>d47bb38</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/266">#266</a> from caarlos0/issuer-key-id</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/756ebbd0d3eff84640b9cbbbd2c9a12a9ca6b279"><code>756ebbd</code></a> Make Issuer Key ID signature subpacket non-critical</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/44ef98c5a687cc0d3b611be2061b32dabcc949dd"><code>44ef98c</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/276">#276</a> from mdosch/fix-random-source-is-broken</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b105e24ae66f11e269a0efb8eb54ee0add82e90b"><code>b105e24</code></a> Merge branch 'main' into fix-random-source-is-broken</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/89b0776c10277d4c5f673f6d808ee4e1fa690b6a"><code>89b0776</code></a> Only check that message signatures are newer than the key</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/2b2dbe9eca392b9fd7e193961df11320f5a5cb80"><code>2b2dbe9</code></a> openpgp/clearsign: just use rand.Reader in tests</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/2732e09cbc206553bd7af8acf73077037bcddb9f"><code>2732e09</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/275">#275</a> from ProtonMail/only-check-msg-sig-newer-than-key</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/8e272e796276ef5ff91e878d680be94faa76ee5b"><code>8e272e7</code></a> Only check that message signatures are newer than the key</li>
<li>Additional commits viewable in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.5...v1.1.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.45.0 to 1.45.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.45.1 (2025-03-02)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(producer): add MaxBufferBytes to limit retry buffer size by <a href="https://github.com/wanwenli"><code>@​wanwenli</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3088">IBM/sarama#3088</a></li>
<li>feat(producer): add sync pool for channel reuse by <a href="https://github.com/kasimtj"><code>@​kasimtj</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3109">IBM/sarama#3109</a></li>
<li>feat: exponential backoff for clients (KIP-580) by <a href="https://github.com/wanwenli"><code>@​wanwenli</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3099">IBM/sarama#3099</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(sasl): add nilguard around token to prevent panic by <a href="https://github.com/hoo47"><code>@​hoo47</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3076">IBM/sarama#3076</a></li>
<li>fix(test): consumer group fetch request messages by <a href="https://github.com/stsmurf"><code>@​stsmurf</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3081">IBM/sarama#3081</a></li>
<li>fix: remove redundant nil check by <a href="https://github.com/knbr13"><code>@​knbr13</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3089">IBM/sarama#3089</a></li>
<li>fix(consumer): add recovery from no leader partitions by <a href="https://github.com/liutao365"><code>@​liutao365</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3101">IBM/sarama#3101</a></li>
<li>produce: set MaxTimestamp by <a href="https://github.com/rockwotj"><code>@​rockwotj</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3108">IBM/sarama#3108</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump go.opentelemetry.io/otel from 1.24.0 to 1.29.0 in /examples/interceptors by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3071">IBM/sarama#3071</a></li>
<li>chore(deps): bump the otel group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3072">IBM/sarama#3072</a></li>
<li>chore(deps): bump the golang-x group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3098">IBM/sarama#3098</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(deps): prevent otel upgrades for now by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3069">IBM/sarama#3069</a></li>
<li>chore: add version constant for kafka 3.7.2 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3073">IBM/sarama#3073</a></li>
<li>chore(ci): fetch kafka 4.0 via tar.gz rather than git by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3079">IBM/sarama#3079</a></li>
<li>fix(ci): tighten up github workflows by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3080">IBM/sarama#3080</a></li>
<li>chore(ci): analyse actions in codeql by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3085">IBM/sarama#3085</a></li>
<li>chore(ci): bump golangci-lint version to v1.63.4 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3090">IBM/sarama#3090</a></li>
<li>feat(ci): add dedicated staticcheck run by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3091">IBM/sarama#3091</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/hoo47"><code>@​hoo47</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3076">IBM/sarama#3076</a></li>
<li><a href="https://github.com/stsmurf"><code>@​stsmurf</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3081">IBM/sarama#3081</a></li>
<li><a href="https://github.com/knbr13"><code>@​knbr13</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3089">IBM/sarama#3089</a></li>
<li><a href="https://github.com/liutao365"><code>@​liutao365</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3101">IBM/sarama#3101</a></li>
<li><a href="https://github.com/rockwotj"><code>@​rockwotj</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3108">IBM/sarama#3108</a></li>
<li><a href="https://github.com/kasimtj"><code>@​kasimtj</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3109">IBM/sarama#3109</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.45.0...v1.45.1">https://github.com/IBM/sarama/compare/v1.45.0...v1.45.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/4001b2edfe5a77943569fb200fec946912672b3c"><code>4001b2e</code></a> chore(ci): bump ubi9/ubi-minimal from <code>b870979</code> to <code>3902bab</code> (<a href="https://redirect.github.com/IBM/sarama/issues/3095">#3095</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/6e71c2f3f1788a24ad77a7e3631bc75ba79de12a"><code>6e71c2f</code></a> chore(deps): bump the golang-x group across 1 directory with 2 updates (<a href="https://redirect.github.com/IBM/sarama/issues/3098">#3098</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/13c92aa8d0f08c371b6a8add4199e9d651f36d44"><code>13c92aa</code></a> feat: exponential backoff for clients (<a href="https://redirect.github.com/IBM/sarama/issues/3099">#3099</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4a028c47c8105c89d0ae7e497771c4e0f54e5e3b"><code>4a028c4</code></a> chore(ci): bump github/codeql-action from 3.28.5 to 3.28.10 (<a href="https://redirect.github.com/IBM/sarama/issues/3104">#3104</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/f2963797deee4e61fa75a99e2e4552c3720884d3"><code>f296379</code></a> chore(ci): bump docker/bake-action from 6.2.0 to 6.4.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3105">#3105</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/c2e0d941cfc1df3a623ae2ec0a1d62618b1c2a91"><code>c2e0d94</code></a> feat(producer): add sync pool for channel reuse (<a href="https://redirect.github.com/IBM/sarama/issues/3109">#3109</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3c678852873cf8d24af6bb02a39e59d0ec04462c"><code>3c67885</code></a> produce: set MaxTimestamp (<a href="https://redirect.github.com/IBM/sarama/issues/3108">#3108</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/60592f6a3bfce97873986f36c0ae1831d0221bf1"><code>60592f6</code></a> fix(consumer): add recovery from no leader partitions (<a href="https://redirect.github.com/IBM/sarama/issues/3101">#3101</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/9ae475a39e964d257e630d75d02541ab5fa70b1c"><code>9ae475a</code></a> Merge pull request <a href="https://redirect.github.com/IBM/sarama/issues/3091">#3091</a> from IBM/dnwe/staticcheck</li>
<li><a href="https://github.com/IBM/sarama/commit/b3aef996e2232c85ee396c441360462865389957"><code>b3aef99</code></a> fix(ci): update ignore directives for staticcheck</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.45.0...v1.45.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.36.2 to 1.36.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c33f3e8d6bd4acd3c33d7abf391c7040305d5f12"><code>c33f3e8</code></a> Release 2025-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e30151fa29807db629931d5acbd0baaa7a4d812"><code>3e30151</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d47f9d487e991a7964e40ae23436808c9c45354"><code>7d47f9d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8c63a09fc01aba1de4b792f78b7fce378801901"><code>c8c63a0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c50ebd4f921acfa7114e34690f20a2634814b23e"><code>c50ebd4</code></a> Reapply &quot;Track credential providers via User-Agent Feature ids&quot; (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3017">#3017</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef40d0c7f18167e851e939d28357a541102a6a94"><code>ef40d0c</code></a> Release 2025-02-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d21bb4570986d62ee45c76a28f30c51b51cad62"><code>2d21bb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8d810fdae9adcf883bc03efd776a115ba8f78635"><code>8d810fd</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0538c527fdf601468f57bdcdfea4cc9c526c20e"><code>a0538c5</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21ec8a2702f18c177c1a2f2d24dda340eef5bed4"><code>21ec8a2</code></a> Release 2025-02-25</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.36.2...v1.36.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.7 to 1.29.8
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c33f3e8d6bd4acd3c33d7abf391c7040305d5f12"><code>c33f3e8</code></a> Release 2025-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e30151fa29807db629931d5acbd0baaa7a4d812"><code>3e30151</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d47f9d487e991a7964e40ae23436808c9c45354"><code>7d47f9d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8c63a09fc01aba1de4b792f78b7fce378801901"><code>c8c63a0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c50ebd4f921acfa7114e34690f20a2634814b23e"><code>c50ebd4</code></a> Reapply &quot;Track credential providers via User-Agent Feature ids&quot; (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3017">#3017</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef40d0c7f18167e851e939d28357a541102a6a94"><code>ef40d0c</code></a> Release 2025-02-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d21bb4570986d62ee45c76a28f30c51b51cad62"><code>2d21bb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8d810fdae9adcf883bc03efd776a115ba8f78635"><code>8d810fd</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0538c527fdf601468f57bdcdfea4cc9c526c20e"><code>a0538c5</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21ec8a2702f18c177c1a2f2d24dda340eef5bed4"><code>21ec8a2</code></a> Release 2025-02-25</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.7...config/v1.29.8">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.60 to 1.17.61
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c33f3e8d6bd4acd3c33d7abf391c7040305d5f12"><code>c33f3e8</code></a> Release 2025-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e30151fa29807db629931d5acbd0baaa7a4d812"><code>3e30151</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d47f9d487e991a7964e40ae23436808c9c45354"><code>7d47f9d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8c63a09fc01aba1de4b792f78b7fce378801901"><code>c8c63a0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c50ebd4f921acfa7114e34690f20a2634814b23e"><code>c50ebd4</code></a> Reapply &quot;Track credential providers via User-Agent Feature ids&quot; (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3017">#3017</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef40d0c7f18167e851e939d28357a541102a6a94"><code>ef40d0c</code></a> Release 2025-02-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d21bb4570986d62ee45c76a28f30c51b51cad62"><code>2d21bb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8d810fdae9adcf883bc03efd776a115ba8f78635"><code>8d810fd</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0538c527fdf601468f57bdcdfea4cc9c526c20e"><code>a0538c5</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21ec8a2702f18c177c1a2f2d24dda340eef5bed4"><code>21ec8a2</code></a> Release 2025-02-25</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.60...credentials/v1.17.61">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.14 to 1.46.0
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/polly/v1.45.14...service/s3/v1.46.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.77.1 to 1.78.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c33f3e8d6bd4acd3c33d7abf391c7040305d5f12"><code>c33f3e8</code></a> Release 2025-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e30151fa29807db629931d5acbd0baaa7a4d812"><code>3e30151</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d47f9d487e991a7964e40ae23436808c9c45354"><code>7d47f9d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8c63a09fc01aba1de4b792f78b7fce378801901"><code>c8c63a0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c50ebd4f921acfa7114e34690f20a2634814b23e"><code>c50ebd4</code></a> Reapply &quot;Track credential providers via User-Agent Feature ids&quot; (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3017">#3017</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef40d0c7f18167e851e939d28357a541102a6a94"><code>ef40d0c</code></a> Release 2025-02-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d21bb4570986d62ee45c76a28f30c51b51cad62"><code>2d21bb4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8d810fdae9adcf883bc03efd776a115ba8f78635"><code>8d810fd</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0538c527fdf601468f57bdcdfea4cc9c526c20e"><code>a0538c5</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21ec8a2702f18c177c1a2f2d24dda340eef5bed4"><code>21ec8a2</code></a> Release 2025-02-25</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.77.1...service/s3/v1.78.0">compare view</a></li>
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

## Reviews

### Review by @swadeley - Approved on March 03, 2025 at 08:14 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1002*
