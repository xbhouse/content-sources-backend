---
type: pull_request
number: 910
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2024-12-02T04:49:45Z
updated: 2024-12-02T17:51:50Z
closed: 2024-12-02T17:50:46Z
merged: 2024-12-02T17:50:46Z
base_branch: main
head_branch: dependabot/go_modules/go-58ffd425a1
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/910
---

# Pull Request #910: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: December 02, 2024 at 04:49 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-58ffd425a1`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.2` | `1.1.3` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.5` | `1.32.6` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.46` | `1.17.47` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.44.0` | `1.45.0` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.11.1732298036` | `2024.11.1732799476` |

Updates `github.com/ProtonMail/go-crypto` from 1.1.2 to 1.1.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>v1.1.3</h2>
<h2>What's Changed</h2>
<ul>
<li>Add argon2 test vector for 32-bit platforms by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/245">ProtonMail/go-crypto#245</a></li>
<li>Validate argon2 params on read by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/246">ProtonMail/go-crypto#246</a></li>
<li>Adapt aead preferences on key generation by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/248">ProtonMail/go-crypto#248</a></li>
<li>Improve AEAD handling by <a href="https://github.com/twiss"><code>@​twiss</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/247">ProtonMail/go-crypto#247</a></li>
<li>Improve error message for decryption with a session key by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/237">ProtonMail/go-crypto#237</a></li>
<li>Flag to allow signing key decryption by <a href="https://github.com/lubux"><code>@​lubux</code></a> in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/251">ProtonMail/go-crypto#251</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.2...v1.1.3">https://github.com/ProtonMail/go-crypto/compare/v1.1.2...v1.1.3</a></p>
<h2>Release v1.1.3-proton.1</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.3 with support for the following non-standardized features:</p>
<ul>
<li>Presistent symmetric keys <a href="https://www.ietf.org/archive/id/draft-ietf-openpgp-persistent-symmetric-keys-00.html">draft-ietf-openpgp-persistent-symmetric-keys-00</a></li>
<li>Automatic forwarding <a href="https://www.ietf.org/archive/id/draft-wussler-openpgp-forwarding-00.html">draft-wussler-openpgp-forwarding-00</a></li>
<li>Post-quantum algorithms <a href="https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/">draft-ietf-openpgp-pqc</a></li>
</ul>
<p>Patches v1.1.3-proton:</p>
<ul>
<li>Marked forwarding key should not be usable in encryption.</li>
</ul>
<h2>Release v1.1.3-proton</h2>
<h2>What's Changed</h2>
<p>This release is v1.1.3 with support for the following non-standardized features:</p>
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
<li><a href="https://github.com/ProtonMail/go-crypto/commit/5521d835096caef67f37fdad5bdc8f276d999747"><code>5521d83</code></a> Flag to allow signing key decryption (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/251">#251</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/5e3e39db239450fc8ab5457cd1047a74b0dbc8d9"><code>5e3e39d</code></a> Improve error message for decryption with a session key (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/237">#237</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/33a08b33597276fdbfcad598d240f8f1d226378b"><code>33a08b3</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/247">#247</a> from ProtonMail/improve-aead</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/9ad55726228a5d5ddd7eaf240f0906cdf8166ef8"><code>9ad5572</code></a> Adapt aead preferences on key generation (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/248">#248</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/531d9f55ff62af70a700931be8de0f22cdd0a830"><code>531d9f5</code></a> Improve documentation</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/1efe4a001d60d4077152d3c98607386025a28d5f"><code>1efe4a0</code></a> Deprecate SerializeEncryptedKey[WithHiddenOption] and SerializeSymmetricKeyEn...</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/ee678440c71d1d90c43f22c6aa5b4a12a2719909"><code>ee67844</code></a> Add SerializeSymmetricKeyEncryptedAEADReuseKey</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/63e3da1cfd65188df420a89a2ce111cb619547e8"><code>63e3da1</code></a> [v2] Use AEAD if all public keys support it</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/d7733dc924c82d4ce2c7bdc1ae59ae83689f360e"><code>d7733dc</code></a> Validate argon2 params on read (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/246">#246</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/c0ca2b83cbbb6844914e9a5f7b9bb29e169acae5"><code>c0ca2b8</code></a> Add argon2 test vector for 32-bit platforms (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/245">#245</a>)</li>
<li>See full diff in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.2...v1.1.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.5 to 1.32.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/776903f3dd7208803912e19d3aa25006a7fbdeee"><code>776903f</code></a> Release 2024-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/170b13cac4658e0909b13468d3959f94c358faf3"><code>170b13c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1a02e760211a1d0762f664973ea4d896376a621"><code>c1a02e7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2835f7bb8d4925acb36a5ab7813cca6ac8977cb7"><code>2835f7b</code></a> Fix user agent to add business metrics at the end instead of prepend them (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2">#2</a>...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4965dd291e68f9345c5eae4dec8a0f63bda436"><code>ba4965d</code></a> Release 2024-11-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55149b05036afa39e3b7c6dd945543faee4069e6"><code>55149b0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fd6bb8b363cce171fba74b72d2bb4142b9306b7c"><code>fd6bb8b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65ab4f88b048175d8fc13fab6f72f31b739a2455"><code>65ab4f8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8939ed049c8b3490f9a3dbd4e4d3d56cf22a27f2"><code>8939ed0</code></a> Release 2024-11-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23cf36b2492a3caf859a6235c6911f1c2fb715ca"><code>23cf36b</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.5...v1.32.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.46 to 1.17.47
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/776903f3dd7208803912e19d3aa25006a7fbdeee"><code>776903f</code></a> Release 2024-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/170b13cac4658e0909b13468d3959f94c358faf3"><code>170b13c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1a02e760211a1d0762f664973ea4d896376a621"><code>c1a02e7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2835f7bb8d4925acb36a5ab7813cca6ac8977cb7"><code>2835f7b</code></a> Fix user agent to add business metrics at the end instead of prepend them (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2">#2</a>...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4965dd291e68f9345c5eae4dec8a0f63bda436"><code>ba4965d</code></a> Release 2024-11-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55149b05036afa39e3b7c6dd945543faee4069e6"><code>55149b0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fd6bb8b363cce171fba74b72d2bb4142b9306b7c"><code>fd6bb8b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65ab4f88b048175d8fc13fab6f72f31b739a2455"><code>65ab4f8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8939ed049c8b3490f9a3dbd4e4d3d56cf22a27f2"><code>8939ed0</code></a> Release 2024-11-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23cf36b2492a3caf859a6235c6911f1c2fb715ca"><code>23cf36b</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.46...credentials/v1.17.47">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.44.0 to 1.45.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/13546abeb3c6bf9f0bfa97855fd8f52a2e26373e"><code>13546ab</code></a> Release 2023-11-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59ce389800703c1509810f19ab77c467531ab07f"><code>59ce389</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8257e839ab34958c84f1671ca89b92cf30360b43"><code>8257e83</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ef931c3c2a1f28e4bc68efe5a0582e3ec32f2ac"><code>8ef931c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6afe5286d243f06aa040d3954b765b38366b88c2"><code>6afe528</code></a> Release 2023-11-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/769b1e08ea9acae6b2488864977a444dccb59e62"><code>769b1e0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b6c76999b24cbe735e03123800eaf7427ab6caf2"><code>b6c7699</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/57f3065c220672b63d1a1af0f424757048805904"><code>57f3065</code></a> breakfix: convert public access block config fields to nilable like s3 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2385">#2385</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.44.0...service/s3/v1.45.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.11.1732298036 to 2024.11.1732799476
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/ee25f1ea7da8997b1d5ea61cb8fedf1067c59286"><code>ee25f1e</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84748b9922def37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/9f0555dd0ceffd6f512d38c776af6213ebe6dd38"><code>9f0555d</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d88eedd54057952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/567dd4e1faf1fde39234ecae418db5f049935680"><code>567dd4e</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7eaabbb4b5c87e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/2335540ffe43e16ffe5d656d6904df2cb8b96ea6"><code>2335540</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a33a32ad80b7a8d53bd49838...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.11.1732298036...release/v2024.11.1732799476">compare view</a></li>
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

### Comment by @swadeley on December 02, 2024 at 01:57 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on December 02, 2024 at 05:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/910*
