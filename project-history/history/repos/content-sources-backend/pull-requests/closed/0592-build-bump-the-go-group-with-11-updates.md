---
type: pull_request
number: 592
title: "Build: Bump the go group with 11 updates"
state: closed
author: dependabot
created: 2024-03-04T04:47:04Z
updated: 2024-03-06T13:35:42Z
closed: 2024-03-06T13:35:40Z
base_branch: main
head_branch: dependabot/go_modules/go-c92752c6bc
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/592
---

# Pull Request #592: Build: Bump the go group with 11 updates

**Author**: @dependabot
**Created**: March 04, 2024 at 04:47 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-c92752c6bc`

## Description

Bumps the go group with 11 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.0-alpha.0-proton` | `1.1.0-alpha.1` |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | `1.8.4` | `1.9.0` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.18.0` | `1.19.0` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.42.2` | `1.43.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.25.0` | `1.25.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.0` | `1.17.4` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.33.1` | `1.34.2` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.15.0` | `2.15.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.2.1707758687` | `2024.2.1708559921` |
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

Updates `github.com/stretchr/testify` from 1.8.4 to 1.9.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stretchr/testify/releases">github.com/stretchr/testify's releases</a>.</em></p>
<blockquote>
<h2>v1.9.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Fix Go modules version by <a href="https://github.com/SuperQ"><code>@​SuperQ</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1394">stretchr/testify#1394</a></li>
<li>Document that require is not safe to call in created goroutines by <a href="https://github.com/programmer04"><code>@​programmer04</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1392">stretchr/testify#1392</a></li>
<li>Remove myself from MAINTAINERS.md by <a href="https://github.com/mvdkleijn"><code>@​mvdkleijn</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1367">stretchr/testify#1367</a></li>
<li>Correct spelling/grammar by <a href="https://github.com/echarrod"><code>@​echarrod</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1389">stretchr/testify#1389</a></li>
<li>docs: Update URLs in README by <a href="https://github.com/davidjb"><code>@​davidjb</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1349">stretchr/testify#1349</a></li>
<li>Update mockery link to Github Pages in README by <a href="https://github.com/LandonTClipp"><code>@​LandonTClipp</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1346">stretchr/testify#1346</a></li>
<li>docs: Fix typos in tests and comments by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1410">stretchr/testify#1410</a></li>
<li>CI: tests from go1.17 by <a href="https://github.com/SuperQ"><code>@​SuperQ</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1409">stretchr/testify#1409</a></li>
<li>Fix adding ? when no values passed by <a href="https://github.com/lesichkovm"><code>@​lesichkovm</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1320">stretchr/testify#1320</a></li>
<li>codegen: use standard header for generated files by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1406">stretchr/testify#1406</a></li>
<li>mock: AssertExpectations log reason only on failure by <a href="https://github.com/hikyaru-suzuki"><code>@​hikyaru-suzuki</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1360">stretchr/testify#1360</a></li>
<li>assert: fix flaky TestNeverTrue by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1417">stretchr/testify#1417</a></li>
<li>README: fix typos &quot;set up&quot; vs &quot;setup&quot; by <a href="https://github.com/ossan-dev"><code>@​ossan-dev</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1428">stretchr/testify#1428</a></li>
<li>mock: move regexp compilation outside of <code>Called</code> by <a href="https://github.com/aud10slave"><code>@​aud10slave</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/631">stretchr/testify#631</a></li>
<li>assert: refactor internal func getLen() by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1445">stretchr/testify#1445</a></li>
<li>mock: deprecate type AnythingOfTypeArgument (<a href="https://redirect.github.com/stretchr/testify/issues/1434">#1434</a>) by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1441">stretchr/testify#1441</a></li>
<li>Remove no longer needed assert.canConvert by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1470">stretchr/testify#1470</a></li>
<li>assert: ObjectsAreEqual: use time.Equal for time.Time types by <a href="https://github.com/tscales"><code>@​tscales</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1464">stretchr/testify#1464</a></li>
<li>Bump actions/checkout from 3 to 4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1466">stretchr/testify#1466</a></li>
<li>Bump actions/setup-go from 3.2.0 to 4.1.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1451">stretchr/testify#1451</a></li>
<li>fix: make EventuallyWithT concurrency safe by <a href="https://github.com/czeslavo"><code>@​czeslavo</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1395">stretchr/testify#1395</a></li>
<li>assert: fix httpCode and HTTPBody occur panic when http.Handler read Body by <a href="https://github.com/hidu"><code>@​hidu</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1484">stretchr/testify#1484</a></li>
<li>assert.EqualExportedValues: fix handling of arrays by <a href="https://github.com/zrbecker"><code>@​zrbecker</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1473">stretchr/testify#1473</a></li>
<li>.github: use latest Go versions by <a href="https://github.com/kevinburkesegment"><code>@​kevinburkesegment</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1489">stretchr/testify#1489</a></li>
<li>assert: Deprecate EqualExportedValues by <a href="https://github.com/HaraldNordgren"><code>@​HaraldNordgren</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1488">stretchr/testify#1488</a></li>
<li>suite: refactor test assertions by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1474">stretchr/testify#1474</a></li>
<li>suite: fix SetupSubTest and TearDownSubTest execution order by <a href="https://github.com/linusbarth"><code>@​linusbarth</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1471">stretchr/testify#1471</a></li>
<li>docs: Fix deprecation comments for http package by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1335">stretchr/testify#1335</a></li>
<li>Add map support doc comments to Subset and NotSubset by <a href="https://github.com/jedevc"><code>@​jedevc</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1306">stretchr/testify#1306</a></li>
<li>TestErrorIs/TestNotErrorIs: check error message contents by <a href="https://github.com/craig65535"><code>@​craig65535</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1435">stretchr/testify#1435</a></li>
<li>suite: fix subtest names (fix <a href="https://redirect.github.com/stretchr/testify/issues/1501">#1501</a>) by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1504">stretchr/testify#1504</a></li>
<li>assert: improve unsafe.Pointer tests by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1505">stretchr/testify#1505</a></li>
<li>assert: simplify isNil implementation by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1506">stretchr/testify#1506</a></li>
<li>assert.InEpsilonSlice: fix expected/actual order and other improvements by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1483">stretchr/testify#1483</a></li>
<li>Fix dependency cycle with objx <a href="https://redirect.github.com/stretchr/testify/issues/1292">#1292</a> by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1453">stretchr/testify#1453</a></li>
<li>mock: refactor TestIsArgsEqual by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1444">stretchr/testify#1444</a></li>
<li>mock: optimize argument matching checks by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1416">stretchr/testify#1416</a></li>
<li>assert: fix TestEventuallyTimeout by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1412">stretchr/testify#1412</a></li>
<li>CI: add go 1.21 in GitHub Actions by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1450">stretchr/testify#1450</a></li>
<li>suite: fix recoverAndFailOnPanic to report test failure at the right location by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1502">stretchr/testify#1502</a></li>
<li>Update maintainers by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1533">stretchr/testify#1533</a></li>
<li>assert: Fix EqualValues to handle overflow/underflow by <a href="https://github.com/arjunmahishi"><code>@​arjunmahishi</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1531">stretchr/testify#1531</a></li>
<li>assert: better formatting for Len() error by <a href="https://github.com/kevinburkesegment"><code>@​kevinburkesegment</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1485">stretchr/testify#1485</a></li>
<li>Ensure AssertExpectations does not fail in skipped tests by <a href="https://github.com/ianrose14"><code>@​ianrose14</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1331">stretchr/testify#1331</a></li>
<li>suite: fix deadlock in suite.Require()/Assert() by <a href="https://github.com/arjunmahishi"><code>@​arjunmahishi</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1535">stretchr/testify#1535</a></li>
<li>Revert &quot;assert: ObjectsAreEqual: use time.Equal for time.Time type&quot; by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1537">stretchr/testify#1537</a></li>
<li>[chore] Add issue templates by <a href="https://github.com/arjunmahishi"><code>@​arjunmahishi</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1538">stretchr/testify#1538</a></li>
<li>Update the build status badge by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1540">stretchr/testify#1540</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stretchr/testify/commit/bb548d0473d4e1c9b7bbfd6602c7bf12f7a84dd2"><code>bb548d0</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1552">#1552</a> from stretchr/dependabot/go_modules/github.com/stret...</li>
<li><a href="https://github.com/stretchr/testify/commit/814075f391adffd2bf2b5110a74c51827ba132c4"><code>814075f</code></a> build(deps): bump github.com/stretchr/objx from 0.5.1 to 0.5.2</li>
<li><a href="https://github.com/stretchr/testify/commit/e0456122451b1839c8d58d32df6364e4d0f0a709"><code>e045612</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1339">#1339</a> from bogdandrutu/uintptr</li>
<li><a href="https://github.com/stretchr/testify/commit/5b6926d686d412518f50e888b9ae9b938355e011"><code>5b6926d</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1385">#1385</a> from hslatman/not-implements</li>
<li><a href="https://github.com/stretchr/testify/commit/9f97d67703eff02136d487e6c907e76fdea31a8b"><code>9f97d67</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1550">#1550</a> from stretchr/release-notes</li>
<li><a href="https://github.com/stretchr/testify/commit/bcb0d3fe49ff300fb78288cc144bc61a881f58ec"><code>bcb0d3f</code></a> Include the auto-release notes in releases</li>
<li><a href="https://github.com/stretchr/testify/commit/fb770f8238261aa22f8e0c56f18168ccb90f4a09"><code>fb770f8</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1247">#1247</a> from ccoVeille/typos</li>
<li><a href="https://github.com/stretchr/testify/commit/85d8bb6eea715dcbbb68f7c87b50e1956e20f892"><code>85d8bb6</code></a> fix typos in comments, tests and github templates</li>
<li><a href="https://github.com/stretchr/testify/commit/e2741fa4e9bf2fdfe3ed48d976a7eeebe76c5009"><code>e2741fa</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1548">#1548</a> from arjunmahishi/msgAndArgs</li>
<li><a href="https://github.com/stretchr/testify/commit/6e59f20c0d3883d2bdc589a9e48374ea30601851"><code>6e59f20</code></a> http_assertions: assert that the msgAndArgs actually works in tests</li>
<li>Additional commits viewable in <a href="https://github.com/stretchr/testify/compare/v1.8.4...v1.9.0">compare view</a></li>
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

### Comment by @app-sre-bot on March 04, 2024 at 04:48 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on March 04, 2024 at 01:58 PM UTC

[test]

### Comment by @dependabot on March 06, 2024 at 01:35 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/592*
