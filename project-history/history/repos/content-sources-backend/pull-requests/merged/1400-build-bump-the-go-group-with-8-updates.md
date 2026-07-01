---
type: pull_request
number: 1400
title: "Build: Bump the go group with 8 updates"
state: merged
author: dependabot
created: 2026-03-02T04:40:20Z
updated: 2026-03-02T09:50:53Z
closed: 2026-03-02T09:50:51Z
merged: 2026-03-02T09:50:51Z
base_branch: main
head_branch: dependabot/go_modules/go-05dc8a43cf
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1400
---

# Pull Request #1400: Build: Bump the go group with 8 updates

**Author**: @dependabot
**Created**: March 02, 2026 at 04:40 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-05dc8a43cf`

## Description

Bumps the go group with 8 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.3.0` | `1.4.0` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.46.3` | `1.47.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.1` | `1.41.2` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.9` | `1.32.10` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.9` | `1.19.10` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.63.1` | `1.63.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.96.0` | `1.96.2` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.42.0` | `0.43.0` |

Updates `github.com/ProtonMail/go-crypto` from 1.3.0 to 1.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>Release v1.4.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Ignore leading and trailing whitespaces in the armor body in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/288">ProtonMail/go-crypto#288</a></li>
<li>Update key_generation.go, rename variables to avoid shadowing in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/290">ProtonMail/go-crypto#290</a></li>
<li>Add InsecureGenerateNonCriticalKeyFlags option to generate non-critical key flags signature subpackets in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/291">ProtonMail/go-crypto#291</a></li>
<li>Add InsecureGenerateNonCriticalSignatureCreationTime option to generate non-critical signature creation time subpackets in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/292">ProtonMail/go-crypto#292</a></li>
<li>Bump dependencies and min go version to 1.23 in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/294">ProtonMail/go-crypto#294</a></li>
<li>ECDHv4: Error on low-order x25519 public key curve points in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/299">ProtonMail/go-crypto#299</a></li>
<li>Cleartext: Only allow valid hashes in header in <a href="https://redirect.github.com/ProtonMail/go-crypto/pull/298">ProtonMail/go-crypto#298</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/ProtonMail/go-crypto/compare/v1.3.0...v1.4.0">https://github.com/ProtonMail/go-crypto/compare/v1.3.0...v1.4.0</a></p>
<h2>Release v1.4.0-proton</h2>
<p>This release is v1.4.0 with support for the following non-standardized features:</p>
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
<li><a href="https://github.com/ProtonMail/go-crypto/commit/a8cc4f09f6cb247ab2180b45029ddaa736674f87"><code>a8cc4f0</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/298">#298</a> from ProtonMail/feat/cleartext-hash-header</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/57f891b6b4d198fb18adb4877d4192fb96f9f5a0"><code>57f891b</code></a> Merge branch 'main' into feat/cleartext-hash-header</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/da5c190d0ba1061cb21d8d311f6032c8bc43e80d"><code>da5c190</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/299">#299</a> from ProtonMail/fix/ecdh-low-order-curve-points</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/3cc59b0193219e5850b874a52873671c65a0c907"><code>3cc59b0</code></a> Merge branch 'main' into feat/cleartext-hash-header</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b11bd2375b66c0f6b33c355563b59029edd7f117"><code>b11bd23</code></a> fix(ecdh): Do not allow low order public key points</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b6bdd12c063898caa069ec5379fe5080d1bafcd1"><code>b6bdd12</code></a> Merge pull request <a href="https://redirect.github.com/ProtonMail/go-crypto/issues/294">#294</a> from ProtonMail/chore/bump-go-and-circl</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b1ff3d56014e94dfd37402b2c1c25d239bb38405"><code>b1ff3d5</code></a> Bump crypto dependencies and min go version to 1.23</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/cfb2af9d2cff2ca2e2d403f9118a1a9265e86e02"><code>cfb2af9</code></a> fix(cleartext): Check hashes in headers</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/de877883d43979a32c829304ad8119dddd8b0dd9"><code>de87788</code></a> Add InsecureGenerateNonCriticalSignatureCreationTime option to generate non-c...</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/0906643e7d5a122dfda584b78b4367e6b0980220"><code>0906643</code></a> Add InsecureGenerateNonCriticalKeyFlags option to generate non-critical key f...</li>
<li>Additional commits viewable in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.3.0...v1.4.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.46.3 to 1.47.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.47.0 (2026-02-27)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>perf(admin): modernize DescribeCluster RPC handling by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3390">IBM/sarama#3390</a></li>
<li>test: expand Java interop tests to cover all compression codecs by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3423">IBM/sarama#3423</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(client): add nilguards to updateBroker by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3393">IBM/sarama#3393</a></li>
<li>fix(broker): auto-close broken connections by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3412">IBM/sarama#3412</a></li>
<li>fix: set version from IBM/sarama, not main app by <a href="https://github.com/adamdecaf"><code>@​adamdecaf</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3415">IBM/sarama#3415</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: finish up the move to atomic types by <a href="https://github.com/puellanivis"><code>@​puellanivis</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3399">IBM/sarama#3399</a></li>
<li>chore: tear down zk in functional tests by <a href="https://github.com/edoardocomar"><code>@​edoardocomar</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3420">IBM/sarama#3420</a></li>
<li>chore: migrate from eapache/go-xerial-snappy to klauspost/compress/sn… by <a href="https://github.com/edoardocomar"><code>@​edoardocomar</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3421">IBM/sarama#3421</a></li>
<li>fix(test): resolve FVT issues in Kafka v2.x interop tests by <a href="https://github.com/edoardocomar"><code>@​edoardocomar</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3424">IBM/sarama#3424</a></li>
<li>feat: add kafka 4.1.1 constants and use in FVT by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3437">IBM/sarama#3437</a></li>
<li>chore: add Kafka 4.0.1 and replace 4.0.0 in FVT by <a href="https://github.com/edoardocomar"><code>@​edoardocomar</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3439">IBM/sarama#3439</a></li>
<li>ci(lint): unblock Go 1.26 lint and handle gosec noise by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3454">IBM/sarama#3454</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update module golang.org/x/crypto to v0.45.0 [security] → v0.48.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3383">#3383</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3425">#3425</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.42.0 to 0.45.0 across /examples (consumergroup, exactly_once, sasl_scram_client, http_server, txn_producer, interceptors) by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3382">#3382</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3381">#3381</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3380">#3380</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3379">#3379</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3378">#3378</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3377">#3377</a></li>
<li>fix(deps): update module golang.org/x/sync to v0.18.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3385">#3385</a></li>
<li>fix(deps): update module golang.org/x/net to v0.49.0 → v0.51.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] and <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3426">#3426</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3427">#3427</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3453">#3453</a></li>
<li>chore(deps): update golangci/golangci-lint-action action to v9 → v9.2.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3370">#3370</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3400">#3400</a></li>
<li>chore(deps): update dependency golangci/golangci-lint to v2.6.2 → v2.8.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3366">#3366</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3401">#3401</a></li>
<li>chore(deps): update docker/bake-action action to v6.10.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3392">#3392</a></li>
<li>chore(deps): update docker/setup-buildx-action action to v3.12.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3416">#3416</a></li>
<li>chore(deps): bump github.com/klauspost/compress from 1.18.1 to 1.18.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3397">#3397</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3430">#3430</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3442">#3442</a></li>
<li>chore(deps): bump github.com/pierrec/lz4/v4 from 4.1.22 to 4.1.25 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3411">#3411</a>, <a href="https://redirect.github.com/IBM/sarama/issues/3432">#3432</a></li>
<li>chore(deps): bump the golang-x group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3405">#3405</a></li>
<li>chore(deps): bump github.com/xdg-go/scram from 1.1.2 to 1.2.0 in /examples/sasl_scram_client by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3394">#3394</a></li>
<li>chore(deps): update dependency dominikh/go-tools to v2026 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/issues/3446">#3446</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3390">IBM/sarama#3390</a></li>
<li><a href="https://github.com/edoardocomar"><code>@​edoardocomar</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3420">IBM/sarama#3420</a></li>
<li><a href="https://github.com/adamdecaf"><code>@​adamdecaf</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3415">IBM/sarama#3415</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.46.3...v1.47.0">https://github.com/IBM/sarama/compare/v1.46.3...v1.47.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/c6a7ea0fa29599bc584b79dac403ff3723130bbd"><code>c6a7ea0</code></a> chore(deps): bump github.com/klauspost/compress from 1.18.3 to 1.18.4 (<a href="https://redirect.github.com/IBM/sarama/issues/3442">#3442</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/66eabe8c7fe1f5a10be3c321edd0b0707d85cfb9"><code>66eabe8</code></a> chore(deps): update dependency dominikh/go-tools to v2026 (<a href="https://redirect.github.com/IBM/sarama/issues/3446">#3446</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e46c2133b1e3639bf6b2a93474e97a32714016ff"><code>e46c213</code></a> ci(gosec): suppress G117 and G704 false positives</li>
<li><a href="https://github.com/IBM/sarama/commit/7f0e26cb7c78bb028419e49cbf4aa44bd997f4ef"><code>7f0e26c</code></a> ci(workflows): bump golangci-lint to v2.10.1</li>
<li><a href="https://github.com/IBM/sarama/commit/7a33045c6f772d5e20ff37856104a792a09462a9"><code>7a33045</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.7 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/34acfaa6337d6b6b8ef7b1594b767bffcb191a99"><code>34acfaa</code></a> chore(ci): bump the actions group with 2 updates (<a href="https://redirect.github.com/IBM/sarama/issues/3450">#3450</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4df4d820713832ba38cf5f1759e8825929428a55"><code>4df4d82</code></a> chore(ci): bump github/codeql-action from 4.31.10 to 4.32.4 (<a href="https://redirect.github.com/IBM/sarama/issues/3451">#3451</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/d4862cefab60f05e6db48a674e42a1a5edd154ad"><code>d4862ce</code></a> chore(deps): bump golang.org/x/net from 0.50.0 to 0.51.0 in the golang-x grou...</li>
<li><a href="https://github.com/IBM/sarama/commit/c1bf7ec2e57c6509fa2605c456f341748bcaab5a"><code>c1bf7ec</code></a> chore: add ai-assistance guidelines to CONTRIBUTING.md (<a href="https://redirect.github.com/IBM/sarama/issues/3452">#3452</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/0a3b6e34c8956dcae9e2d70c3dbf45c3620af839"><code>0a3b6e3</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.7 Docker dige...</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.46.3...v1.47.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.1 to 1.41.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/238dde72705f3e20fcf3e63a7b8fc98a52692f31"><code>238dde7</code></a> Release 2026-02-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3a957b913ac755abfb4dcbebbe66011017a4074c"><code>3a957b9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2972c7d7e82ea9775baaae52082f11b102df0cb6"><code>2972c7d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa037947b6153d8626180fe493ac7319a4fef1b4"><code>aa03794</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7712be78a0c47df47c37a2d2e423134c8d29746e"><code>7712be7</code></a> Feature add new eventstream implementation to support bedrockruntime#InvokeMo...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/da36a982260c7f190112204be5a81ee031cde183"><code>da36a98</code></a> Release 2026-02-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9d7b59c43a463f90850ac26e9abc7bcdd608f2e4"><code>9d7b59c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e75eb3f3e33cbb3ebb85b4e89fff0e8e7c4700be"><code>e75eb3f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2c987cd0f909ae349935bb7b4df1cf847a16c1d1"><code>2c987cd</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/766b176ae9db6c2afb31cfc1f15678d9a1ae7ae1"><code>766b176</code></a> parameterize generate-dev (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3329">#3329</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.1...v1.41.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.9 to 1.32.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/238dde72705f3e20fcf3e63a7b8fc98a52692f31"><code>238dde7</code></a> Release 2026-02-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3a957b913ac755abfb4dcbebbe66011017a4074c"><code>3a957b9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2972c7d7e82ea9775baaae52082f11b102df0cb6"><code>2972c7d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa037947b6153d8626180fe493ac7319a4fef1b4"><code>aa03794</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7712be78a0c47df47c37a2d2e423134c8d29746e"><code>7712be7</code></a> Feature add new eventstream implementation to support bedrockruntime#InvokeMo...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/da36a982260c7f190112204be5a81ee031cde183"><code>da36a98</code></a> Release 2026-02-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9d7b59c43a463f90850ac26e9abc7bcdd608f2e4"><code>9d7b59c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e75eb3f3e33cbb3ebb85b4e89fff0e8e7c4700be"><code>e75eb3f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2c987cd0f909ae349935bb7b4df1cf847a16c1d1"><code>2c987cd</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/766b176ae9db6c2afb31cfc1f15678d9a1ae7ae1"><code>766b176</code></a> parameterize generate-dev (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3329">#3329</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.9...config/v1.32.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.9 to 1.19.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fcc0f5daa41033b7a696f8cc5f53a9fc8696a274"><code>fcc0f5d</code></a> Release 2023-04-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cd750e0b2188951c525ae7917d47ae9e2d013a1b"><code>cd750e0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1bc2f0514c73be727d3536e829fef18911bb45ae"><code>1bc2f05</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b964f5ca3ccae40ef8d7a56fd1b1ad040764b5fa"><code>b964f5c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fd6901588c6d13a7889787328d0628134afd14cc"><code>fd69015</code></a> fix APIGW exports nullability exceptions</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fae239abb72a392c50e05aa567b2e5cc2b93a10b"><code>fae239a</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2089">#2089</a> from aws/auditAccessibility</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/acf33a2872dbd9d3547842f38ccd38863048c121"><code>acf33a2</code></a> Update aws-sdk-go-v2's comment codegened from Smithy Go's updated document sm...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/27360c189034eb456ebe7458bb8e019bc45d686c"><code>27360c1</code></a> fix APIGW exports nullability exceptions</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/30383d567f67c2a67b2b40a462a8c284c49d1796"><code>30383d5</code></a> Release 2023-04-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/352f89c2d23ec6249a699c732ba5c9ae050f833f"><code>352f89c</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/m2/v1.19.9...service/iam/v1.19.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.63.1 to 1.63.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d0a2d1af31ec2c2c33fcaf31ee9698bc8edcab45"><code>d0a2d1a</code></a> Release 2024-09-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/85c956cba02f9770379584312819cd74506276c4"><code>85c956c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c0e4d46f88f4bd78ca16ae07c3fa6688b9b7851d"><code>c0e4d46</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e3cdfa3b9d35093e9053dc138493132f516d37d"><code>3e3cdfa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e211ac0d542b6529cb7f81c05d230e5a2a6f91d6"><code>e211ac0</code></a> Release 2024-09-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c00259a1f703a31a5fcf7ebd05466a2401d44b14"><code>c00259a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a6f2cebf76c68d3234807cdbdbdb073c8a3b514b"><code>a6f2ceb</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9437d07ba03173c7489d04efbb4cdf36083a010c"><code>9437d07</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.63.1...service/s3/v1.63.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.96.0 to 1.96.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dfcf25b6d15674848f71bd7e5ee8ed22b115c6a2"><code>dfcf25b</code></a> Release 2026-02-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/981d0b47535d2bd4fa4887b9d03113ef1cac6357"><code>981d0b4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/60d901e42f63e035c405d7362eed33679a04a72f"><code>60d901e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51e7bd80bf802c2e62d357fcef56b5f265e1c449"><code>51e7bd8</code></a> Fix: support arbitrary pre-calculated checksum values on s3 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3332">#3332</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b53e5ba605419d859b48e03ecf2e31ecb3032e8"><code>2b53e5b</code></a> Release 2026-02-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ebfd72de7f69e281ded6bf346ae64b4e2d0fab39"><code>ebfd72d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/87640ce3e7fe8ae0c14a348d0265e1795a62dec9"><code>87640ce</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba64e1bad306a5a426f7539c769b0b95c89027d4"><code>ba64e1b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3f1e54552ce88d7fe213d1eaa28f363de5a9c6f"><code>f3f1e54</code></a> Release 2026-02-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c4df0e13c9580d56383a87c99af979e75997489c"><code>c4df0e1</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.96.0...service/s3/v1.96.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.42.0 to 0.43.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.43.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Add support for go 1.26 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1193">#1193</a>
<ul>
<li>bump minimum supported go version to 1.24</li>
</ul>
</li>
<li>change type signature of attributes for Logs and Metrics. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1205">#1205</a>
<ul>
<li>users are not supposed to modify Attributes directly on the Log/Metric itself, but this is still is a breaking change on the type.</li>
</ul>
</li>
<li>Send uint64 overflowing attributes as numbers. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1198">#1198</a>
<ul>
<li>The SDK was converting overflowing uint64 attributes to strings for slog and logrus integrations. To eliminate double types for these attributes, the SDK now sends the overflowing attribute as is, and lets the server handle the overflow appropriately.</li>
<li>It is expected that overflowing unsigned integers would now get dropped, instead of converted to strings.</li>
</ul>
</li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add zap logging integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1184">#1184</a></li>
<li>Log specific message for RequestEntityTooLarge by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1185">#1185</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Improve otel span map cleanup performance by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1200">#1200</a></li>
<li>Ensure correct signal delivery on multi-client setups by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1190">#1190</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<ul>
<li>Bump golang.org/x/crypto to 0.48.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1196">#1196</a></li>
<li>Use go1.24.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1195">#1195</a></li>
<li>Bump github.com/gofiber/fiber/v2 from 2.52.9 to 2.52.11 in /fiber by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1191">#1191</a></li>
<li>Bump getsentry/craft from 2.19.0 to 2.20.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1187">#1187</a></li>
</ul>
<h4>Other</h4>
<ul>
<li>Add omitzero and remove custom serialization by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1197">#1197</a></li>
<li>Rename Telemetry Processor components by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1186">#1186</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.43.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Add support for go 1.26 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1193">#1193</a>
<ul>
<li>bump minimum supported go version to 1.24</li>
</ul>
</li>
<li>change type signature of attributes for Logs and Metrics. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1205">#1205</a>
<ul>
<li>users are not supposed to modify Attributes directly on the Log/Metric itself, but this is still is a breaking change on the type.</li>
</ul>
</li>
<li>Send uint64 overflowing attributes as numbers. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1198">#1198</a>
<ul>
<li>The SDK was converting overflowing uint64 attributes to strings for slog and logrus integrations. To eliminate double types for these attributes, the SDK now sends the overflowing attribute as is, and lets the server handle the overflow appropriately.</li>
<li>It is expected that overflowing unsigned integers would now get dropped, instead of converted to strings.</li>
</ul>
</li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add zap logging integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1184">#1184</a></li>
<li>Log specific message for RequestEntityTooLarge by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1185">#1185</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Improve otel span map cleanup performance by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1200">#1200</a></li>
<li>Ensure correct signal delivery on multi-client setups by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1190">#1190</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<ul>
<li>Bump golang.org/x/crypto to 0.48.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1196">#1196</a></li>
<li>Use go1.24.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1195">#1195</a></li>
<li>Bump github.com/gofiber/fiber/v2 from 2.52.9 to 2.52.11 in /fiber by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1191">#1191</a></li>
<li>Bump getsentry/craft from 2.19.0 to 2.20.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1187">#1187</a></li>
</ul>
<h4>Other</h4>
<ul>
<li>Add omitzero and remove custom serialization by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1197">#1197</a></li>
<li>Rename Telemetry Processor components by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1186">#1186</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/8dbf970375c9e4f6d16026bfd47801400938ddaa"><code>8dbf970</code></a> release: 0.43.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/44caaaebddd7fb0941d4e5ec7138959eb31fbca4"><code>44caaae</code></a> ref!: modify Attributes type for Logs and Metrics (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1205">#1205</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/71a723095f8f006fe8536f37f917fbf4a969186b"><code>71a7230</code></a> fix: improve otel span map cleanup performance (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1200">#1200</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/0fe1ae59c07fca8ab35ad7e030b0243514eb0485"><code>0fe1ae5</code></a> feat: add zap logging integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1184">#1184</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e1e8ceb07d7ae747e384a6b12131a5d2cb033527"><code>e1e8ceb</code></a> feat!: Send uint64 overflowing attributes as numbers for slog and logrus (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1198">#1198</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/0faba363a348eceddedb918de72f6805f39ee68a"><code>0faba36</code></a> chore: add omitzero and remove custom serialization (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1197">#1197</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/648d2d28e7f2c1f1ceafc5cad578c76660a87f00"><code>648d2d2</code></a> build(deps): bump golang.org/x/crypt to 0.48.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1196">#1196</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/18c50b96a3d5af19267cd2796182c2bdb2684172"><code>18c50b9</code></a> fix(deps): use go1.24.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1195">#1195</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2d5637b9baf119468a58e03d0fc450af6e1c03fc"><code>2d5637b</code></a> build(deps)!: add support for go1.26 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1193">#1193</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f057a8dab7d5897213c810007bc19e78a61e2dc3"><code>f057a8d</code></a> fix: ensure correct signal delivery on multi-client setups (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1190">#1190</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.42.0...v0.43.0">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on March 02, 2026 at 09:50 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1400*
