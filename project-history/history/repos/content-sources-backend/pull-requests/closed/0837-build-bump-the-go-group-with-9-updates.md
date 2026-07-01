---
type: pull_request
number: 837
title: "Build: Bump the go group with 9 updates"
state: closed
author: dependabot
created: 2024-10-07T04:48:37Z
updated: 2024-10-07T11:06:35Z
closed: 2024-10-07T11:06:34Z
base_branch: main
head_branch: dependabot/go_modules/go-9b831c1718
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/837
---

# Pull Request #837: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: October 07, 2024 at 04:48 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9b831c1718`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.0-alpha.5-proton` | `1.1.0-beta.0-proton` |
| [github.com/content-services/tang](https://github.com/content-services/tang) | `0.0.8` | `0.0.9` |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.127.1-0.20240830113739-c606b5546b12` | `0.128.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.31.0` | `1.32.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.37` | `1.17.39` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.40.3` | `1.41.0` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.9.1727266727` | `2024.10.1727972141` |
| [golang.org/x/exp](https://github.com/golang/exp) | `0.0.0-20240823005443-9b4947da3948` | `0.0.0-20240909161429-701f63a606c0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.10` | `1.25.11` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.0-alpha.5-proton to 1.1.0-beta.0-proton
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>v1.1.0-beta.0-proton</h2>
<p>This pre-release is v1.1.0-beta.0 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
<h2>v1.1.0-beta.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Allow Salted S2K for high-entropy passphrases by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/219">ProtonMail/go-crypto#219</a></li>
<li>fix data race in getEd25519Sk by <a href="https://github.com/DmitriyMV"><code>@​DmitriyMV</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/220">ProtonMail/go-crypto#220</a></li>
<li>Add back crypto.Signer support for ECDSA signing keys. by <a href="https://github.com/doryiii"><code>@​doryiii</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/227">ProtonMail/go-crypto#227</a></li>
<li>Add support for keyserver preferences and preferred keyserver (closes <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/206">#206</a>) by <a href="https://github.com/andrewgdotcom"><code>@​andrewgdotcom</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/232">ProtonMail/go-crypto#232</a></li>
<li>Fix ECDH using v6 keys by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/233">ProtonMail/go-crypto#233</a></li>
<li>No v6 ECC keys with legacy OIDs by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/234">ProtonMail/go-crypto#234</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.5...v1.1.0-beta.0">https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.5...v1.1.0-beta.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b04e354c4bbb9562ccf0056ca1f00a2d1510c1df"><code>b04e354</code></a> Fix HMAC generation (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/204">#204</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/85bc845effd3bbf13982728d0cbbf3eabbc3470b"><code>85bc845</code></a> Replace ioutil.ReadAll with io.ReadAll</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/5f49c73febb104213fb2b84906d1c7d0d8a2f547"><code>5f49c73</code></a> fix(v2): Adapt NewForwardingEntity to refactored NewEntity</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/9aa010ac90364ddc6b5a2a8fa4b63d39af7064b2"><code>9aa010a</code></a> fix(v2): Do not allow encrpytion with a forwarding key</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/91c2e9ea977b85f9df249d0e41a528eb16d8da8f"><code>91c2e9e</code></a> feat: Add symmetric keys to v2</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3f6d02a4c11f43864892b7b28d372fea30ab71eb"><code>3f6d02a</code></a> fix: Address warnings</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/da0a0cfd6492b29d28a34aa8b277eb2d66c4148c"><code>da0a0cf</code></a> feat: Add forwarding to v2 api</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/0e6a359585d71e74007029d5d7ada829fa3de583"><code>0e6a359</code></a> fix: Address rebase on version 2 issues</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/c602a742e26cc091b51766ad1f2d482b5c077953"><code>c602a74</code></a> Use fingerprints instead of KeyIDs</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/348b81dfe730128505567eb246605e3725e5c56d"><code>348b81d</code></a> Create a copy of the encrypted key when forwarding</li>
<li>Additional commits viewable in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.5-proton...v1.1.0-beta.0-proton">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/tang` from 0.0.8 to 0.0.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/f7e88f5874e826238c153fbb0b8898cc467fbb18"><code>f7e88f5</code></a> Fixes 4740: Change base query to use inner joins (<a href="https://redirect.github.com/content-services/tang/issues/13">#13</a>)</li>
<li><a href="https://github.com/content-services/tang/commit/42f95bd2b7c85a00f0fbccc455e2427e88e9dfc6"><code>42f95bd</code></a> Update to go 1.23 and update deps (<a href="https://redirect.github.com/content-services/tang/issues/12">#12</a>)</li>
<li><a href="https://github.com/content-services/tang/commit/5086af3296452c7412331ef576e2cf0c170f17a1"><code>5086af3</code></a> Build: Change to new pulp image (<a href="https://redirect.github.com/content-services/tang/issues/11">#11</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.8...v0.0.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getkin/kin-openapi` from 0.127.1-0.20240830113739-c606b5546b12 to 0.128.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.128.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3filter: Fix default value for array in for query param by <a href="https://github.com/Tommy-42"><code>@​Tommy-42</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1000">getkin/kin-openapi#1000</a></li>
<li>Add github.com/pb33f/libopenapi by <a href="https://github.com/Jille"><code>@​Jille</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1004">getkin/kin-openapi#1004</a></li>
<li>Introduce an option to override the regex implementation by <a href="https://github.com/alexbakker"><code>@​alexbakker</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1006">getkin/kin-openapi#1006</a></li>
<li>make form required field order deterministic by <a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1008">getkin/kin-openapi#1008</a></li>
<li>openapi2: fix un/marshalling discriminator field by <a href="https://github.com/reversearrow"><code>@​reversearrow</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1011">getkin/kin-openapi#1011</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Tommy-42"><code>@​Tommy-42</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1000">getkin/kin-openapi#1000</a></li>
<li><a href="https://github.com/Jille"><code>@​Jille</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1004">getkin/kin-openapi#1004</a></li>
<li><a href="https://github.com/alexbakker"><code>@​alexbakker</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1006">getkin/kin-openapi#1006</a></li>
<li><a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1008">getkin/kin-openapi#1008</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.127.0...v0.128.0">https://github.com/getkin/kin-openapi/compare/v0.127.0...v0.128.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/getkin/kin-openapi/commits/v0.128.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.31.0 to 1.32.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a76a2ae73fe6ae04c8dba07570145eba0582555"><code>7a76a2a</code></a> Release 2024-10-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e35b8bedbb56d7b39d8ccc60cc120a7b61d5fec5"><code>e35b8be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e9587148dadaebdfeda731a68bb30740aedfcdd"><code>6e95871</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d58df32b098fca0d31acca9056f8f56f5c1cdca"><code>0d58df3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/280579adfcd8a5e23658566d3033f40b5e73d447"><code>280579a</code></a> add HTTP client metrics from smithy-go (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2815">#2815</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb83a1cde349c71506c3addbdca49d50b7fcb3d2"><code>cb83a1c</code></a> Release 2024-10-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8503c8359d6964807111f3dc1cdefef9e0e8d44d"><code>8503c83</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/317e15b2f7380e9cf1e1ca3574c88d57475d58f9"><code>317e15b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ff516832203c3ac03426b6cab998485b331cbc5"><code>3ff5168</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6191b25230d7cf4333b1e3bc2e60400fb56a4591"><code>6191b25</code></a> Release 2024-10-02</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.31.0...v1.32.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.37 to 1.17.39
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a76a2ae73fe6ae04c8dba07570145eba0582555"><code>7a76a2a</code></a> Release 2024-10-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e35b8bedbb56d7b39d8ccc60cc120a7b61d5fec5"><code>e35b8be</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e9587148dadaebdfeda731a68bb30740aedfcdd"><code>6e95871</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d58df32b098fca0d31acca9056f8f56f5c1cdca"><code>0d58df3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/280579adfcd8a5e23658566d3033f40b5e73d447"><code>280579a</code></a> add HTTP client metrics from smithy-go (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2815">#2815</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb83a1cde349c71506c3addbdca49d50b7fcb3d2"><code>cb83a1c</code></a> Release 2024-10-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8503c8359d6964807111f3dc1cdefef9e0e8d44d"><code>8503c83</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/317e15b2f7380e9cf1e1ca3574c88d57475d58f9"><code>317e15b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ff516832203c3ac03426b6cab998485b331cbc5"><code>3ff5168</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6191b25230d7cf4333b1e3bc2e60400fb56a4591"><code>6191b25</code></a> Release 2024-10-02</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.37...credentials/v1.17.39">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.40.3 to 1.41.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/61039fea9cc9e080c53382850c87685b5406fd68"><code>61039fe</code></a> Release 2023-10-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/797e0560769725635218fc30a2554c1bbaccc01b"><code>797e056</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/822585d3f621a7c5844584d8e471c32f852702aa"><code>822585d</code></a> Update SDK's smithy-go dependency to v1.16.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abf753db747dd256f3ee69712a19d1d3dc681f23"><code>abf753d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/99861c071109ce5ee4f1cb3b72ead2062b3bd86c"><code>99861c0</code></a> lang: bump minimum go version to 1.19 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2338">#2338</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ac0a53ac45acaadc4526fd25b643dc46032b02a"><code>2ac0a53</code></a> Release 2023-10-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c10aa0ad45a155d7a6a9968894aed0d8e1cb4e81"><code>c10aa0a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9c456c10923952d6bd1d7d59ded3d70588e1ff36"><code>9c456c1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3cb5dc1d777c4e28cd360728c45e8b5aa2a7b2b0"><code>3cb5dc1</code></a> Release 2023-10-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9b3ad7b1e6ce72730896fe7c9d165543ff158ed3"><code>9b3ad7b</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/rds/v1.40.3...service/s3/v1.41.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.9.1727266727 to 2024.10.1727972141
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/163fac44231272f873ccc7b5af52ec814b399584"><code>163fac4</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a847483938ee6c37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/bb2f32f149d32492712e28602ae3701d3a27e10b"><code>bb2f32f</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b9793e328deec37ae5486b92ad2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.9.1727266727...release/v2024.10.1727972141">compare view</a></li>
</ul>
</details>
<br />

Updates `golang.org/x/exp` from 0.0.0-20240823005443-9b4947da3948 to 0.0.0-20240909161429-701f63a606c0
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/golang/exp/commits">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.10 to 1.25.11
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/f96e3072ba52813abc11c9faa550b2d21eaf6727"><code>f96e307</code></a> Fix inefassign</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/f90ea358aa6ace72497f8d12edd5ca04cf80626b"><code>f90ea35</code></a> Fix gosec and detected issues</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/b0c4de038dc0b127a4b1f42e062c1fda8b599931"><code>b0c4de0</code></a> Add utility function to convert uint64 to uint32 securely</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/45b7622937b85ed39879037c7b4f1c1bdbd86fa4"><code>45b7622</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/578">#578</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/7a2cee440118ae2b0d7181a9de75431b95e0a2b3"><code>7a2cee4</code></a> Bump github.com/getsentry/sentry-go from 0.24.1 to 0.28.1</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.10...v1.25.11">compare view</a></li>
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

### Comment by @app-sre-bot on October 07, 2024 at 04:49 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on October 07, 2024 at 08:01 AM UTC

`panic: failed to resolve "definitions" in fragment in URI: "#/definitions/api.Feature": struct field "definitions" not found`

### Comment by @jlsherrill on October 07, 2024 at 11:05 AM UTC

@dependabot ignore github.com/getkin/kin-openapi minor version

### Comment by @dependabot on October 07, 2024 at 11:05 AM UTC

OK, I won't notify you about version 0.128.x of github.com/getkin/kin-openapi again, unless you unignore it.

### Comment by @dependabot on October 07, 2024 at 11:06 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/837*
