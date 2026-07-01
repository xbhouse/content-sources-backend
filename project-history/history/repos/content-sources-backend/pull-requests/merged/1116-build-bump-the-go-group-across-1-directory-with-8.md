---
type: pull_request
number: 1116
title: "Build: Bump the go group across 1 directory with 8 updates"
state: merged
author: dependabot
created: 2025-05-26T04:09:02Z
updated: 2025-05-27T20:34:24Z
closed: 2025-05-27T20:34:16Z
merged: 2025-05-27T20:34:16Z
base_branch: main
head_branch: dependabot/go_modules/go-9f0bcde136
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1116
---

# Pull Request #1116: Build: Bump the go group across 1 directory with 8 updates

**Author**: @dependabot
**Created**: May 26, 2025 at 04:09 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9f0bcde136`

## Description

Bumps the go group with 8 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.2.0` | `1.3.0` |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.13.3` | `4.13.4` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.26.1` | `1.30.0` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.49.0` | `1.50.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.79.3` | `1.79.4` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.5.1746754376` | `2025.5.1747750822` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.32.0` | `0.33.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.7.4` | `5.7.5` |


Updates `github.com/ProtonMail/go-crypto` from 1.2.0 to 1.3.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.3.0</h2>
<h2>What's Changed</h2>
<ul>
<li>API v2: Tolerate invalid key signatures if one verifies in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/284">ProtonMail/go-crypto#284</a></li>
<li>Enforce acceptable hash functions in clearsign in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/281">ProtonMail/go-crypto#281</a></li>
<li>Allow to set a decompressed message size limit in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/285">ProtonMail/go-crypto#285</a></li>
<li>API v1: Only allow acceptable hashes when writing signatures in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/286">ProtonMail/go-crypto#286</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.2.0...v1.3.0">https://github.com/ProtonMail/go-crypto/compare/v1.2.0...v1.3.0</a></p>
<h2>Release v1.3.0-proton</h2>
<p>This release is v1.3.0 with support for the following non-standardized features:</p>
<ul>
<li>Presistent symmetric keys experimental + latest draft <a href="https://www.ietf.org/archive/id/draft-ietf-openpgp-persistent-symmetric-keys-00.html">draft-ietf-openpgp-persistent-symmetric-keys-00</a></li>
<li>Automatic forwarding <a href="https://www.ietf.org/archive/id/draft-wussler-openpgp-forwarding-00.html">draft-wussler-openpgp-forwarding-00</a></li>
<li>Post-quantum algorithms <a href="https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/">draft-ietf-openpgp-pqc</a> (Updated to draft-ietf-openpgp-pqc-09)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3b22d8539b95b3b7e76a911053023e6ef9ef51d6"><code>3b22d85</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/286">#286</a> from ProtonMail/feat/v1-api-enforce-signature-hashes</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/a9af95cb02243b6a8bb81800e0e45e0d74430a69"><code>a9af95c</code></a> feat(v1): Only allow acceptable hashes when writing signatures</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3c6060347fd1028352e327406d90df96b3c9edb4"><code>3c60603</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/285">#285</a> from ProtonMail/feat/decompression-size-limit</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/d66470076a7f33e0e1ff7b9cbeb6199751e21ddb"><code>d664700</code></a> feat: Allow to set a decompressed size limit</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/e1ea40eb5d76f94f0be9813f0b10591fbc159d75"><code>e1ea40e</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/281">#281</a> from ProtonMail/fix/clearsign-check-hash</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/9cd4c3a20f9c3ae39f7248d76aae585105adc5fb"><code>9cd4c3a</code></a> feat(clearsign): Write complete legacy hash header</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/244eb1cfffe50b52c9167cc15e4170e3ed77e741"><code>244eb1c</code></a> fix(clearsign): Enforce acceptable hash functions in clearsign</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/a994e322317842fe2675f5c2ed8c665137cf2032"><code>a994e32</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/284">#284</a> from ProtonMail/feat/less-restrictive-key-signature</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/8fe0e95d7221c857daf24fb71ab51d320df42722"><code>8fe0e95</code></a> feat(v2): Tolerate invalid key signatures if one passes</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.2.0...v1.3.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/labstack/echo/v4` from 4.13.3 to 4.13.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.13.4</h2>
<h2>What's Changed</h2>
<ul>
<li>chore: fix some typos in comment by <a href="https://github.com/zhuhaicity"><code>@​zhuhaicity</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2735">labstack/echo#2735</a></li>
<li>CI: test with Go 1.24 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2748">labstack/echo#2748</a></li>
<li>Add support for TLS WebSocket proxy by <a href="https://github.com/t-ibayashi-safie"><code>@​t-ibayashi-safie</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2762">labstack/echo#2762</a></li>
</ul>
<p><strong>Security</strong></p>
<ul>
<li>Update dependencies for <a href="https://pkg.go.dev/vuln/GO-2025-3487">GO-2025-3487</a>, <a href="https://pkg.go.dev/vuln/GO-2025-3503">GO-2025-3503</a> and <a href="https://pkg.go.dev/vuln/GO-2025-3595">GO-2025-3595</a> in <a href="https://redirect.github.com/labstack/echo/pull/2780">labstack/echo#2780</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/zhuhaicity"><code>@​zhuhaicity</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2735">labstack/echo#2735</a></li>
<li><a href="https://github.com/t-ibayashi-safie"><code>@​t-ibayashi-safie</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2762">labstack/echo#2762</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.13.3...v4.13.4">https://github.com/labstack/echo/compare/v4.13.3...v4.13.4</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.13.4 - 2025-05-22</h2>
<p><strong>Enhancements</strong></p>
<ul>
<li>chore: fix some typos in comment by <a href="https://github.com/zhuhaicity"><code>@​zhuhaicity</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2735">labstack/echo#2735</a></li>
<li>CI: test with Go 1.24 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2748">labstack/echo#2748</a></li>
<li>Add support for TLS WebSocket proxy by <a href="https://github.com/t-ibayashi-safie"><code>@​t-ibayashi-safie</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2762">labstack/echo#2762</a></li>
</ul>
<p><strong>Security</strong></p>
<ul>
<li>Update dependencies for <a href="https://pkg.go.dev/vuln/GO-2025-3487">GO-2025-3487</a>, <a href="https://pkg.go.dev/vuln/GO-2025-3503">GO-2025-3503</a> and <a href="https://pkg.go.dev/vuln/GO-2025-3595">GO-2025-3595</a> in <a href="https://redirect.github.com/labstack/echo/pull/2780">labstack/echo#2780</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/98ca08e7dd64075b858e758d6693bf9799340756"><code>98ca08e</code></a> Improve changelog for 4.13.4 (<a href="https://redirect.github.com/labstack/echo/issues/2783">#2783</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/f24aaff49bc5ba613222176237927a0409c4cdbb"><code>f24aaff</code></a> Revert &quot;CORS: reject requests with 401 for non-preflight request with not mat...</li>
<li><a href="https://github.com/labstack/echo/commit/9f50a659e90d7e6ff9317d7ea740af4fd55c0f57"><code>9f50a65</code></a> Changelog for 4.13.4 (<a href="https://redirect.github.com/labstack/echo/issues/2781">#2781</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/d735cb6a2ec191f68d26a807fb730180458ec771"><code>d735cb6</code></a> Upgrade dependencies (<a href="https://redirect.github.com/labstack/echo/issues/2780">#2780</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/de44c53a5b16f7dca451f337f7221a1448c92007"><code>de44c53</code></a> Add support for TLS WebSocket proxy (<a href="https://redirect.github.com/labstack/echo/issues/2762">#2762</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/c44f6283f02a2cd945deb9363e0db5ab8e3afab0"><code>c44f628</code></a> CI: test with Go 1.24 (<a href="https://redirect.github.com/labstack/echo/issues/2748">#2748</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/ce0b12ae531b8c6523797af2425c9c0b6d772c4e"><code>ce0b12a</code></a> chore: fix some typos in comment (<a href="https://redirect.github.com/labstack/echo/issues/2735">#2735</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/ee3e1297788e8fc3543489ebc0d4e940be7c6532"><code>ee3e129</code></a> CORS: reject requests with 401 for non-preflight request with not matching or...</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.13.3...v4.13.4">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.26.1 to 1.30.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.30.0</h2>
<h2>Changes</h2>
<ul>
<li>(WIP) Implement Generics API <a href="https://github.com/jinzhu"><code>@​jinzhu</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7424">#7424</a>)</li>
<li>perf(schema): avoid redundant strings.ToLower call <a href="https://github.com/1911860538"><code>@​1911860538</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7464">#7464</a>)</li>
<li>fix: return init dialector error <a href="https://github.com/codingplz"><code>@​codingplz</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7379">#7379</a>)</li>
<li>perf: break early on match failure in ParseConstraint <a href="https://github.com/1911860538"><code>@​1911860538</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7402">#7402</a>)</li>
<li>feat: error message show field name <a href="https://github.com/pi12138"><code>@​pi12138</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7452">#7452</a>)</li>
<li>perf: use strings.IndexByte to replace strings.Index <a href="https://github.com/1911860538"><code>@​1911860538</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7454">#7454</a>)</li>
<li>perf: use strings.Cut to replace strings.SplitN <a href="https://github.com/1911860538"><code>@​1911860538</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7455">#7455</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/c44405a25b0fb15c20265e672b8632b8774793ca"><code>c44405a</code></a> Implement Generics API (<a href="https://redirect.github.com/go-gorm/gorm/issues/7424">#7424</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/751c1d6b45245843405c87d8b4b94e26c26b1ba8"><code>751c1d6</code></a> perf(schema): avoid redundant strings.ToLower call (<a href="https://redirect.github.com/go-gorm/gorm/issues/7464">#7464</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/8e7ab46c1b865386bbb8f5322d64880ee5ce6f40"><code>8e7ab46</code></a> fix: return init dialector error (<a href="https://redirect.github.com/go-gorm/gorm/issues/7379">#7379</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/e3037e4ef0394faf118cf926fd03283995d6da1f"><code>e3037e4</code></a> perf: break early on match failure in ParseConstraint (<a href="https://redirect.github.com/go-gorm/gorm/issues/7402">#7402</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/12043304197c54237d430825e90adc50d63ef383"><code>1204330</code></a> feat: error message show field name (<a href="https://redirect.github.com/go-gorm/gorm/issues/7452">#7452</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/9703eb775f203e916759bfa51af6724b05ff04ba"><code>9703eb7</code></a> perf: use strings.IndexByte to replace strings.Index (<a href="https://redirect.github.com/go-gorm/gorm/issues/7454">#7454</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/1c966e0d25a6cdedbb7d599eae0cd060a0572ddf"><code>1c966e0</code></a> perf: use strings.Cut to replace strings.SplitN (<a href="https://redirect.github.com/go-gorm/gorm/issues/7455">#7455</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.26.1...v1.30.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.49.0 to 1.50.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4334b43c27eeb4c4ecfcba1a87cc08f963fe91a3"><code>4334b43</code></a> Release 2024-02-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9e29187dc08c2e9bf8f66166841ca9c75e0624ab"><code>9e29187</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f672c49fb7b4199becd3ba0a6e9d2b35f215e985"><code>f672c49</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f935782a0d8ee350736d4515cc05001e4ac3d01"><code>7f93578</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/da7bcdacddcf05ee9c897137fdbcdc8caf501698"><code>da7bcda</code></a> feat: add client config passthrough to waiter opts (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2499">#2499</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1bb048272ad54b3cbeb3b6da99f4be8090bea5d2"><code>1bb0482</code></a> Release 2024-02-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cc83a2b05532dfe5e6ccba25881887352134fbbd"><code>cc83a2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f44d9967b75499363425c29a5cf40375e9869a7"><code>7f44d99</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4280ccb2671145220be23f299bc0db4a2c061ff5"><code>4280ccb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a264562983a722b90b7a08e5547de85274006e30"><code>a264562</code></a> fix awsjson error deserialization to not expect string code (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2489">#2489</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.49.0...service/s3/v1.50.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.79.3 to 1.79.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8e8487a51e9eb22a101c49cc61b98ca8990c1322"><code>8e8487a</code></a> Release 2025-05-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7c47834d1a8af069c5be836bf58ba7c15dff0e54"><code>7c47834</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0394af4e72ef0953760e9dc2d93d680f03bd9a1"><code>b0394af</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78e08c4cdb30821be24c3a353e3a794f5ada036d"><code>78e08c4</code></a> Handle checksum for unseekable 0-length s3 request body (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3086">#3086</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6a7e8b52e910b0dbd30c6b26a864a39c19e02294"><code>6a7e8b5</code></a> Release 2025-05-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/422e4b47599a61c3c2ad09948981666d695ac611"><code>422e4b4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff53c6a27faed701ca87d9aa887d45e808652283"><code>ff53c6a</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/892de6fd423d5f9adfb955818042921f606d32ad"><code>892de6f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ec5036806f84653f73348bd7b2e15a21a44fe2ee"><code>ec50368</code></a> Release 2025-05-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/16611cb7c081bfb7221a27fc21a9454f3048a3e1"><code>16611cb</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.79.3...service/s3/v1.79.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.5.1746754376 to 2025.5.1747750822
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/1ce6e3fe606a57374ada493845d8ad15b0f7a3c6"><code>1ce6e3f</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84744b63ab6ef37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/26dc245b8bdde0a9430fce019ad924039edf87af"><code>26dc245</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97943386aab037ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/420967825c87e352f6fff69a41538bb0472be249"><code>4209678</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276e49ae256027b8e3dae95843...</li>
<li><a href="https://github.com/content-services/zest/commit/701618d57983a7d83738619d0a680227caabe4df"><code>701618d</code></a> Update pulp bindings to 386d86254354e29d3b864523aed47886e3586b067968b5e2da9ab...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.5.1746754376...release/v2025.5.1747750822">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.32.0 to 0.33.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.33.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.33.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Rename the internal <code>Logger</code> to <code>DebugLogger</code>. This feature was only used when you set <code>Debug: True</code> in your <code>sentry.Init()</code> call. If you haven't used the Logger directly, no changes are necessary. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1012">#1012</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add support for <a href="https://docs.sentry.io/product/explore/logs/">Structured Logging</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1010">#1010</a>)</p>
<pre lang="go"><code>logger := sentry.NewLogger(ctx)
logger.Info(ctx, &quot;Hello, Logs!&quot;)
</code></pre>
<p>You can learn more about Sentry Logs on our <a href="https://docs.sentry.io/product/explore/logs/">docs</a> and the <a href="https://github.com/getsentry/sentry-go/blob/master/_examples/logs/main.go">examples</a>.</p>
</li>
<li>
<p>Add new attributes APIs, which are currently only exposed on logs. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1007">#1007</a>)</p>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Do not push a new scope on <code>StartSpan</code>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1013">#1013</a>)</li>
<li>Fix an issue where the propagated smapling decision wasn't used. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/995">#995</a>)</li>
<li>[Otel] Prefer <code>httpRoute</code> over <code>httpTarget</code> for span descriptions. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1002">#1002</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Update <code>github.com/stretchr/testify</code> to v1.8.4. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/988">#988</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.33.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.33.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Rename the internal <code>Logger</code> to <code>DebugLogger</code>. This feature was only used when you set <code>Debug: True</code> in your <code>sentry.Init()</code> call. If you haven't used the Logger directly, no changes are necessary. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1012">#1012</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add support for <a href="https://docs.sentry.io/product/explore/logs/">Structured Logging</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1010">#1010</a>)</p>
<pre lang="go"><code>logger := sentry.NewLogger(ctx)
logger.Info(ctx, &quot;Hello, Logs!&quot;)
</code></pre>
<p>You can learn more about Sentry Logs on our <a href="https://docs.sentry.io/product/explore/logs/">docs</a> and the <a href="https://github.com/getsentry/sentry-go/blob/master/_examples/logs/main.go">examples</a>.</p>
</li>
<li>
<p>Add new attributes APIs, which are currently only exposed on logs. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1007">#1007</a>)</p>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Do not push a new scope on <code>StartSpan</code>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1013">#1013</a>)</li>
<li>Fix an issue where the propagated smapling decision wasn't used. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/995">#995</a>)</li>
<li>[Otel] Prefer <code>httpRoute</code> over <code>httpTarget</code> for span descriptions. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1002">#1002</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Update <code>github.com/stretchr/testify</code> to v1.8.4. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/988">#988</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/03a4debf6ca740cd1f5a0fdbef0522066c1ceef9"><code>03a4deb</code></a> release: 0.33.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/34ba180a74ec11f739d35bb35286e070f75f4795"><code>34ba180</code></a> Prepare 0.33.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1016">#1016</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c340089ff5e8a6cd896a0720e3245c2e0609f895"><code>c340089</code></a> Add initial logger implementation (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1010">#1010</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/00a135a57d4261f0c58a4ac821d392811fdcd04a"><code>00a135a</code></a> Do not push a new scope on <code>StartSpan</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1013">#1013</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/64de096a6f33547c9aae4b1b6abed0ae0b0848bb"><code>64de096</code></a> Rename <code>Logger</code> to <code>DebugLogger</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1012">#1012</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/523e1b40a94f6edcb62e6fcd52af678f68c03c6b"><code>523e1b4</code></a> Bump github.com/stretchr/testify on otel submodule (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1011">#1011</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bf20c9f50aba46ec972260075dc28049d757b896"><code>bf20c9f</code></a> Prefer http route over http target for span description (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1002">#1002</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/021ccc1973f29e4444aa3a2d7b2bbca4d5697ea7"><code>021ccc1</code></a> Add attributes API (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1007">#1007</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c7d162cfa490f2b759109187296f7dd8acfcf66e"><code>c7d162c</code></a> Update github.com/stretchr/testify to v1.8.4 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/988">#988</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a309f56a2b567b4976eaac5681177598f6d4b986"><code>a309f56</code></a> build(deps): bump actions/create-github-app-token from 1.11.5 to 1.12.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/987">#987</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.32.0...v0.33.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.4 to 5.7.5
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.5 (May 17, 2025)</h1>
<ul>
<li>Support sslnegotiation connection option (divyam234)</li>
<li>Update golang.org/x/crypto to v0.37.0. This placates security scanners that were unable to see that pgx did not use the behavior affected by <a href="https://pkg.go.dev/vuln/GO-2025-3487">https://pkg.go.dev/vuln/GO-2025-3487</a>.</li>
<li>TraceLog now logs Acquire and Release at the debug level (dave sinclair)</li>
<li>Add support for PGTZ environment variable</li>
<li>Add support for PGOPTIONS environment variable</li>
<li>Unpin memory used by Rows quicker</li>
<li>Remove PlanScan memoization. This resolves a rare issue where scanning could be broken for one type by first scanning another. The problem was in the memoization system and benchmarking revealed that memoization was not providing any meaningful benefit.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/15bca4a4e14e0049777c1245dba4c16300fe4fd0"><code>15bca4a</code></a> Release v5.7.5</li>
<li><a href="https://github.com/jackc/pgx/commit/1d557f9116c5d8fd1c6242dbc4db1e06f44b09e1"><code>1d557f9</code></a> Remove PlanScan memoization</li>
<li><a href="https://github.com/jackc/pgx/commit/de7fe81d78655c58bc41427fef4c32c317b20884"><code>de7fe81</code></a> Use reflect.TypeFor instead of reflect.TypeOf</li>
<li><a href="https://github.com/jackc/pgx/commit/d9eb089bd72b1adc0e5347df30ba74080e5445a8"><code>d9eb089</code></a> Remove unused function</li>
<li><a href="https://github.com/jackc/pgx/commit/6be24eb08d57825e0ac68696b2ac50e4d80dea42"><code>6be24eb</code></a> Fix comment typo</li>
<li><a href="https://github.com/jackc/pgx/commit/07871c0a346cdcabfa0e39996b00557665a3b56c"><code>07871c0</code></a> Zero internal baseRows references to allow GC earlier</li>
<li><a href="https://github.com/jackc/pgx/commit/777e7e5cdf2d349c37e1eef8eedc0e21857e9b95"><code>777e7e5</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2313">#2313</a> from stampy88/tracelog_pool_additions</li>
<li><a href="https://github.com/jackc/pgx/commit/151bd026ec836cbcb2b90c01424300ee19987bb8"><code>151bd02</code></a> Switched to <code>LogLevelDebug</code></li>
<li><a href="https://github.com/jackc/pgx/commit/540fcaa9b908880ed9e82ccd3e560a3232e55a7d"><code>540fcaa</code></a> Add support for PGOPTIONS environment variable</li>
<li><a href="https://github.com/jackc/pgx/commit/3a248e3822b1178c27ad311b4110bb125f7ebb5a"><code>3a248e3</code></a> Add support for PGTZ environment variable</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.7.4...v5.7.5">compare view</a></li>
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

### Comment by @jlsherrill on May 27, 2025 at 05:44 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on May 27, 2025 at 08:34 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1116*
