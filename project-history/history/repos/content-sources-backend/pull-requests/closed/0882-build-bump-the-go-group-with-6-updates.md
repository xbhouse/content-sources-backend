---
type: pull_request
number: 882
title: "Build: Bump the go group with 6 updates"
state: closed
author: dependabot
created: 2024-11-11T04:58:22Z
updated: 2024-11-12T16:41:58Z
closed: 2024-11-12T16:41:56Z
base_branch: main
head_branch: dependabot/go_modules/go-f46b4d18e4
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/882
---

# Pull Request #882: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: November 11, 2024 at 04:58 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-f46b4d18e4`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.0-beta.0-proton` | `1.1.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.3` | `1.32.4` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.42` | `1.17.44` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.43.0` | `1.43.2` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.4.17` | `4.4.18` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.10.1730398214` | `2024.11.1731092116` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.0-beta.0-proton to 1.1.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.1.2</h2>
<h2>What's Changed</h2>
<ul>
<li>Cleartext Framework: Exclude the line-ending separator when decoding plaintext by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/242">ProtonMail/go-crypto#242</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.1...v1.1.2">https://github.com/ProtonMail/go-crypto/compare/v1.1.1...v1.1.2</a></p>
<h2>Release v1.1.2-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.2 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
<h2>Release v1.1.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Fix <code>clearsign.Encode</code> backwards compatibility  <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/239">#239</a> by <a href="https://github.com/mdosch"><code>@​mdosch</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/240">ProtonMail/go-crypto#240</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0...v1.1.1">https://github.com/ProtonMail/go-crypto/compare/v1.1.0...v1.1.1</a></p>
<h2>Release v1.1.1-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.1 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
<h2>Release v1.1.0</h2>
<h2>What's Changed</h2>
<p>This release adds full support for the new version of the OpenPGP standard, <a href="https://www.rfc-editor.org/rfc/rfc9580.html">RFC 9580</a>. In addition, the release introduces an improved non-backwards compatible v2 API. The API in the <code>openpgp</code> package remains fully backwards compatible while the new v2 API is located in a separate <code>v2</code> package in <code>openpgp</code>.</p>
<p>For the full changes since <code>v1.0.0</code>, see the previous release notes. For the full changelog, see <a href="https://github.com/ProtonMail/go-crypto/compare/v1.0.0...v1.1.0">https://github.com/ProtonMail/go-crypto/compare/v1.0.0...v1.1.0</a>.</p>
<p>Changes since <code>v1.1.0-beta.0</code>:</p>
<ul>
<li>Replace expiring curve448 integration test vector by <a href="https://github.com/lubux"><code>@​lubux</code></a></li>
<li>Validate input key size in SEIPDv2 decryption by <a href="https://github.com/lubux"><code>@​lubux</code></a></li>
</ul>
<p><strong>Changelog since <code>v1.1.0-beta.0</code></strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-beta.0...v1.1.0">https://github.com/ProtonMail/go-crypto/compare/v1.1.0-beta.0...v1.1.0</a>.</p>
<h2>Release v1.1.0-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.0 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/2d2c7894dbb6ff089c9702e1553111340c3e683d"><code>2d2c789</code></a> feat(cleartext): Do not include line ending separator in plaintext (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/242">#242</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/f8b3f214bfc9c08e7b4af9d1fd406573f5f28d9f"><code>f8b3f21</code></a> Remove cleartext Encode header argument <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/239">#239</a> (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/240">#240</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b97cc3cc403885d5e5f635f2fbe7870c543202a0"><code>b97cc3c</code></a> feat: Validate input key size in SEIPDv2 decryption (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/236">#236</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/20ab0e40584bc971f516c1a45983f82c48239f9a"><code>20ab0e4</code></a> Replace expiring curve448 integration test vector (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/235">#235</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/ef6d67282052194d12bbac3fea88a8124fddd347"><code>ef6d672</code></a> Fix SEIPDv2 packet parsing error messages (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/215">#215</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/4333135edeb9bdcfe35016a6d0b2ae4471be906d"><code>4333135</code></a> Randomize v4 and v5 signatures via custom notation (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/209">#209</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/196c8f5f31c4b8b81c33606338c09ad1d9700940"><code>196c8f5</code></a> Allow parsing empty Key Flags subpackets (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/214">#214</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/c25e074d11e5f1813d0e04ed9c984214d0535a4c"><code>c25e074</code></a> Put support for reading v5 packets behind a feature flag (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/212">#212</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3df78a3ea124d4f5ed9ec02d14e78731282dc15a"><code>3df78a3</code></a> Restrict S2K types (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/213">#213</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3272cd7ba797dbd62275bec987082862e3409bf5"><code>3272cd7</code></a> ECDH with a v6 key must use the full fingerprint (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/211">#211</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-beta.0-proton...v1.1.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.3 to 1.32.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70eb57ac775f782db6856c73f1ca22eae8e48ac2"><code>70eb57a</code></a> Release 2024-11-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cd2c6b1af3d2fd8b114804a604a6b621fbed92c3"><code>cd2c6b1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b2a737b179ea3b7b699584ff6c391619b8f6de3"><code>2b2a737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8c9892f0427c4bf13dac6f944b250fe6512eb6cf"><code>8c9892f</code></a> bump smithy-go codegen to latest (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2890">#2890</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/82897becacb0c658ac8eec09c7e2929f9bfdd0fb"><code>82897be</code></a> fix potential for user-agent lang value mismatch in tests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2888">#2888</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/061540b5a73ead172928c9c5aebc48c011bf4ee1"><code>061540b</code></a> Cloudfront - add expire time in signed cookie. (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2862">#2862</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa3bd1f75d6849d7301f20107485cb54f39dd387"><code>aa3bd1f</code></a> fix makefile to not spam releases for feature/dynamodb/attributevalue (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2885">#2885</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb96051d5b9f85e74e8371384b3d8e8e0d7370cf"><code>eb96051</code></a> Release 2024-11-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f2d000cf7b48c47a1f262676b59b608328af13a"><code>7f2d000</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8542f2f22d9a1d827b6730f8eade2388c41abdaa"><code>8542f2f</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.3...v1.32.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.42 to 1.17.44
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f70834c690f3fecf04ef006f0ea7022884b367f"><code>2f70834</code></a> Release 2024-11-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef9a3c319fb525f680f6d6bdffdfe4153eaa3135"><code>ef9a3c3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c054fe7021957875ecc9707b0cfd49512251d1e9"><code>c054fe7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b68675cf5962947f8001f96ee8cd3b1d66aacc46"><code>b68675c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1d989f31b0248f4d800d6b07bdab20bde40b59c5"><code>1d989f3</code></a> send opt-in query-compatible header where applicable (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2891">#2891</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5d0eb2386ab466c40e4760960f7cf160cdac557b"><code>5d0eb23</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2889">#2889</a> from aws/customize-identity-store-exception-message</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70eb57ac775f782db6856c73f1ca22eae8e48ac2"><code>70eb57a</code></a> Release 2024-11-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cd2c6b1af3d2fd8b114804a604a6b621fbed92c3"><code>cd2c6b1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b2a737b179ea3b7b699584ff6c391619b8f6de3"><code>2b2a737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8c9892f0427c4bf13dac6f944b250fe6512eb6cf"><code>8c9892f</code></a> bump smithy-go codegen to latest (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2890">#2890</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.42...credentials/v1.17.44">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.43.0 to 1.43.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2de0027dc478a6ae80e9f2d24d904a425169a23b"><code>2de0027</code></a> Release 2023-11-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0c890c5eaf354ff23feb727ded9f50aaee9f1c4"><code>f0c890c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e032d9ea8d98d366f2467a72834d2cc0ee865edd"><code>e032d9e</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/507661ff1edbc896fbdfe3ea2e4c2e74be3b4e3c"><code>507661f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4128360684a451476e33c0f979921bc46ff63656"><code>4128360</code></a> fix: respect functional option modifications to RetryMaxAttempts (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2390">#2390</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d66ab2c239681700d947dfe859c3804377d0b4a8"><code>d66ab2c</code></a> Release 2023-11-27.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a264f87e1b9454dfa8a8f7226e6c72409fb3359d"><code>a264f87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d50b35dc1b0030e8ff5b1468b996d0779a25269"><code>7d50b35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/13546abeb3c6bf9f0bfa97855fd8f52a2e26373e"><code>13546ab</code></a> Release 2023-11-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59ce389800703c1509810f19ab77c467531ab07f"><code>59ce389</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.43.0...service/ssm/v1.43.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.4.17 to 4.4.18
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/7d599aa0af6af4c9eb86cdf6faa2f9799b0b5eb6"><code>7d599aa</code></a> Update bindings to release/v4.4.18</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.4.17...release/v4.4.18">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.10.1730398214 to 2024.11.1731092116
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/4b126cb0fe102f59cca1f9a0a10cf547df629ec1"><code>4b126cb</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84748e454ee8f37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/1f4aaf66c247eee82c4c0b0ba2606fe7b5154fc0"><code>1f4aaf6</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4786962d4d2167968b5e2da9ab...</li>
<li><a href="https://github.com/content-services/zest/commit/1d95a5d286160faba2fd068f893b5194be6cb9c9"><code>1d95a5d</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4786944b943f67968b5e2da9ab...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.10.1730398214...release/v2024.11.1731092116">compare view</a></li>
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

### Comment by @app-sre-bot on November 11, 2024 at 05:00 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on November 12, 2024 at 04:37 PM UTC

@dependabot recreate

### Comment by @dependabot on November 12, 2024 at 04:41 PM UTC

Superseded by #886.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/882*
