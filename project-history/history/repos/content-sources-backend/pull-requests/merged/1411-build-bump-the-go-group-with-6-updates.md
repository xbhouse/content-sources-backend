---
type: pull_request
number: 1411
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2026-03-16T04:39:20Z
updated: 2026-03-16T08:49:33Z
closed: 2026-03-16T08:49:31Z
merged: 2026-03-16T08:49:31Z
base_branch: main
head_branch: dependabot/go_modules/go-7b9640a113
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1411
---

# Pull Request #1411: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: March 16, 2026 at 04:39 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-7b9640a113`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.133.0` | `0.134.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.3` | `1.41.4` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.11` | `1.32.12` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.11` | `1.19.12` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.64.0` | `1.64.1` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.96.4` | `1.97.1` |

Updates `github.com/getkin/kin-openapi` from 0.133.0 to 0.134.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.134.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3filter: feat(multipart-ct-decoder): default body decoder on binary schema by <a href="https://github.com/mieltn"><code>@​mieltn</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1097">getkin/kin-openapi#1097</a></li>
<li>openapi3: ensure SchemaErrors get the correct path for allOf schemas by <a href="https://github.com/jamesfcarter"><code>@​jamesfcarter</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1098">getkin/kin-openapi#1098</a></li>
<li>openapi3filter: Fix empty string handling in parameter validation by <a href="https://github.com/utah-KT"><code>@​utah-KT</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1096">getkin/kin-openapi#1096</a></li>
<li>tidy(docs): Update references to oapi-codegen project URL (issue getkin#1094) by <a href="https://github.com/frenchi"><code>@​frenchi</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1095">getkin/kin-openapi#1095</a></li>
<li>openapi2conv: fix allOf inside additionalProperties by <a href="https://github.com/dani-maarouf"><code>@​dani-maarouf</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1103">getkin/kin-openapi#1103</a></li>
<li>openapi3filter: Reject requests with body when non expected (issue 1100) by <a href="https://github.com/CynanX"><code>@​CynanX</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1101">getkin/kin-openapi#1101</a></li>
<li>openapi3: fix for RFC3339 validation: time-offset is a required component by <a href="https://github.com/mpreu"><code>@​mpreu</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1104">getkin/kin-openapi#1104</a></li>
<li>openapi3: add file path to origin location tracking by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1128">getkin/kin-openapi#1128</a></li>
<li>openapi3filter: fix bug where absent optional properties fail validation in form-urlencoded requests by <a href="https://github.com/thierry-f-78"><code>@​thierry-f-78</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1110">getkin/kin-openapi#1110</a></li>
<li>feat: add document-scoped format validators to prevent global state pollution by <a href="https://github.com/the-corp-mark"><code>@​the-corp-mark</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1126">getkin/kin-openapi#1126</a></li>
<li>openapi3: process discriminator mapping values as refs by <a href="https://github.com/RaduPetreTarean"><code>@​RaduPetreTarean</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1108">getkin/kin-openapi#1108</a></li>
<li>openapi3: serialize Extensions when using $ref by <a href="https://github.com/irees"><code>@​irees</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1131">getkin/kin-openapi#1131</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/mieltn"><code>@​mieltn</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1097">getkin/kin-openapi#1097</a></li>
<li><a href="https://github.com/jamesfcarter"><code>@​jamesfcarter</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1098">getkin/kin-openapi#1098</a></li>
<li><a href="https://github.com/utah-KT"><code>@​utah-KT</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1096">getkin/kin-openapi#1096</a></li>
<li><a href="https://github.com/frenchi"><code>@​frenchi</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1095">getkin/kin-openapi#1095</a></li>
<li><a href="https://github.com/dani-maarouf"><code>@​dani-maarouf</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1103">getkin/kin-openapi#1103</a></li>
<li><a href="https://github.com/CynanX"><code>@​CynanX</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1101">getkin/kin-openapi#1101</a></li>
<li><a href="https://github.com/mpreu"><code>@​mpreu</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1104">getkin/kin-openapi#1104</a></li>
<li><a href="https://github.com/thierry-f-78"><code>@​thierry-f-78</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1110">getkin/kin-openapi#1110</a></li>
<li><a href="https://github.com/the-corp-mark"><code>@​the-corp-mark</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1126">getkin/kin-openapi#1126</a></li>
<li><a href="https://github.com/RaduPetreTarean"><code>@​RaduPetreTarean</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1108">getkin/kin-openapi#1108</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.133.0...v0.134.0">https://github.com/getkin/kin-openapi/compare/v0.133.0...v0.134.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/713eff1d1285b8e4934b218ebf4757b3f00452d7"><code>713eff1</code></a> openapi3: serialize Extensions when using $ref (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1131">#1131</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/88234d0958c940e9e35e58536e5e47e57f98d285"><code>88234d0</code></a> openapi3: process discriminator mapping values as refs (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1108">#1108</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/b4a86cedc7677a73b646de9f1d3cc8230c8280e5"><code>b4a86ce</code></a> feat: add document-scoped format validators to prevent global state pollution...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/fd00a26032ccf10396e178330b3aad1f49c3b1a4"><code>fd00a26</code></a> openapi3filter: fix bug where absent optional properties fail validation in f...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/ede8b1fc0c72a7fe2e01abbb26b3a25bdd209d37"><code>ede8b1f</code></a> openapi3: add file path to origin location tracking (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1128">#1128</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/45db2adb0102203579276f24456ce494b59e751a"><code>45db2ad</code></a> Fix RFC3339 validation (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1104">#1104</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/2ac63467aaa45671538aeee38b52e75e9fe914a8"><code>2ac6346</code></a> feat: support rejecting when request body present but not required by specifi...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/6321ee8adad6038eadae32ecd49e58b84ff91603"><code>6321ee8</code></a> openapi2conv: fix allOf inside additionalProperties (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1103">#1103</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/f53e403411003d3552409ee916c213eb11d44d9c"><code>f53e403</code></a> tidy(docs): Update references to oapi-codegen project URL (issue getkin#1094)...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/6a04fdf87012fc4319cfcb7a1239f135668768d1"><code>6a04fdf</code></a> openapi3: Allow usage of empty string (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1096">#1096</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.133.0...v0.134.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.3 to 1.41.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b84293d4af5f8b777c48b73cc39669c5b10f914b"><code>b84293d</code></a> Release 2026-03-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6f286944a7c5e3e0e37513ca0f5fcc9167c86eab"><code>6f28694</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0f243647f9fff51c14a6e792293c000724e366b"><code>f0f2436</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/042a1eae898b51269bcd56c88e37d301f96bfdef"><code>042a1ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3d420736f493c08749a2ed5c7c13a15c5e5ee26"><code>f3d4207</code></a> test sigv4 stream signer (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3347">#3347</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56f2f26e22271bbd833466b6c1a0fbbecd07ff53"><code>56f2f26</code></a> Add polly SynthesizeSpeech presign missing fields serd (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3344">#3344</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a330a45638e61e13a356cfba1e7260698e70603a"><code>a330a45</code></a> add changelog for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3283">#3283</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58b98f6bdb598cb4a2825cfc946c1e6a295303d1"><code>58b98f6</code></a> Remove X-Amz-Security-Token header on redirect to different host (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3283">#3283</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/238eeadab024b378aec5b23c585322e5af479cf5"><code>238eead</code></a> add changelog for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3238">#3238</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65e8aea1eb9777a16c12bb870db28285fbcdc093"><code>65e8aea</code></a> docs: fix typo (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3238">#3238</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.3...v1.41.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.11 to 1.32.12
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b84293d4af5f8b777c48b73cc39669c5b10f914b"><code>b84293d</code></a> Release 2026-03-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6f286944a7c5e3e0e37513ca0f5fcc9167c86eab"><code>6f28694</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0f243647f9fff51c14a6e792293c000724e366b"><code>f0f2436</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/042a1eae898b51269bcd56c88e37d301f96bfdef"><code>042a1ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3d420736f493c08749a2ed5c7c13a15c5e5ee26"><code>f3d4207</code></a> test sigv4 stream signer (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3347">#3347</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56f2f26e22271bbd833466b6c1a0fbbecd07ff53"><code>56f2f26</code></a> Add polly SynthesizeSpeech presign missing fields serd (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3344">#3344</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a330a45638e61e13a356cfba1e7260698e70603a"><code>a330a45</code></a> add changelog for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3283">#3283</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58b98f6bdb598cb4a2825cfc946c1e6a295303d1"><code>58b98f6</code></a> Remove X-Amz-Security-Token header on redirect to different host (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3283">#3283</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/238eeadab024b378aec5b23c585322e5af479cf5"><code>238eead</code></a> add changelog for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3238">#3238</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65e8aea1eb9777a16c12bb870db28285fbcdc093"><code>65e8aea</code></a> docs: fix typo (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3238">#3238</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.11...config/v1.32.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.11 to 1.19.12
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/08f1f0b3e3d3f09b699c84f1f5b56b026fba6e15"><code>08f1f0b</code></a> Release 2022-10-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e1e20e6ca01d3acf5529dbfa059bda3b2ff5393"><code>0e1e20e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/281c268a255720163c78c38c97a92553fabf8f94"><code>281c268</code></a> Update SDK's smithy-go dependency to v1.13.4</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/db7c0a3fd1c72951a0673c13b6602b943285796c"><code>db7c0a3</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1eae80df178a5e3cd03b1cf04a6c7c9648e65e5a"><code>1eae80d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/17628c478a72ed2bc3596c4b7f24a49fa2251107"><code>17628c4</code></a> EC2 IMDS client logging fixes (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1891">#1891</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/765544941191825edd26162f9790bf11f059d426"><code>7655449</code></a> Release 2022-10-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dcae829ecc334f91502afd6d7ae2295861db9885"><code>dcae829</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b82766b858e595943b26924ad1f107cd04363d66"><code>b82766b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1c05fb6452a1f74985ff6deb7a642b9eb441274a"><code>1c05fb6</code></a> Implements IsCredentialsProvider for checking if a provider matches a target ...</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.11...service/sqs/v1.19.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.64.0 to 1.64.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb83a1cde349c71506c3addbdca49d50b7fcb3d2"><code>cb83a1c</code></a> Release 2024-10-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8503c8359d6964807111f3dc1cdefef9e0e8d44d"><code>8503c83</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/317e15b2f7380e9cf1e1ca3574c88d57475d58f9"><code>317e15b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ff516832203c3ac03426b6cab998485b331cbc5"><code>3ff5168</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.64.0...service/s3/v1.64.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.96.4 to 1.97.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b84293d4af5f8b777c48b73cc39669c5b10f914b"><code>b84293d</code></a> Release 2026-03-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6f286944a7c5e3e0e37513ca0f5fcc9167c86eab"><code>6f28694</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0f243647f9fff51c14a6e792293c000724e366b"><code>f0f2436</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/042a1eae898b51269bcd56c88e37d301f96bfdef"><code>042a1ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3d420736f493c08749a2ed5c7c13a15c5e5ee26"><code>f3d4207</code></a> test sigv4 stream signer (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3347">#3347</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56f2f26e22271bbd833466b6c1a0fbbecd07ff53"><code>56f2f26</code></a> Add polly SynthesizeSpeech presign missing fields serd (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3344">#3344</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a330a45638e61e13a356cfba1e7260698e70603a"><code>a330a45</code></a> add changelog for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3283">#3283</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58b98f6bdb598cb4a2825cfc946c1e6a295303d1"><code>58b98f6</code></a> Remove X-Amz-Security-Token header on redirect to different host (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3283">#3283</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/238eeadab024b378aec5b23c585322e5af479cf5"><code>238eead</code></a> add changelog for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3238">#3238</a></li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65e8aea1eb9777a16c12bb870db28285fbcdc093"><code>65e8aea</code></a> docs: fix typo (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3238">#3238</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.96.4...service/s3/v1.97.1">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on March 16, 2026 at 08:32 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1411*
