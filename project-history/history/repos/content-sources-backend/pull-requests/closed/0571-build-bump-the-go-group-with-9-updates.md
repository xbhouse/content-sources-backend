---
type: pull_request
number: 571
title: "Build: Bump the go group with 9 updates"
state: closed
author: dependabot
created: 2024-02-12T04:20:08Z
updated: 2024-02-14T08:29:46Z
closed: 2024-02-14T08:29:44Z
base_branch: main
head_branch: dependabot/go_modules/go-1afcc04e51
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/571
---

# Pull Request #571: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: February 12, 2024 at 04:20 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-1afcc04e51`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.10` | `1.0.11` |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.31.0` | `1.32.0` |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.4` | `1.5.6` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.5` | `1.25.7-0.20240204074919-46816ad31dde` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.40.1` | `1.42.2` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.31.0` | `1.32.0` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.1.1706018763` | `2024.2.1706885661` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.26.0` | `0.27.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.2` | `5.5.3` |

Updates `github.com/content-services/yummy` from 1.0.10 to 1.0.11
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/yummy/releases">github.com/content-services/yummy's releases</a>.</em></p>
<blockquote>
<h2>v1.0.11</h2>
<h2>What's Changed</h2>
<ul>
<li>Fixes 3539: introspection fails if comps file not detected by <a href="https://github.com/xbhouse"><code>@​xbhouse</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/23">#23</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.10...v1.0.11"><code>v1.0.10...v1.0.11</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/cba215cb5240c4e5218f1abba131a2e088faf64d"><code>cba215c</code></a> Fixes 3539: introspection fails if comps file not detected (<a href="https://redirect.github.com/content-services/yummy/issues/23">#23</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/e205d78a7378384a30d6c2622f74df34136e695c"><code>e205d78</code></a> Bump github.com/cloudflare/circl from 1.3.3 to 1.3.7 (<a href="https://redirect.github.com/content-services/yummy/issues/22">#22</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.10...v1.0.11">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/rs/zerolog` from 1.31.0 to 1.32.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/rs/zerolog/commit/147ae653507db96d498366e368f715f4da23bf25"><code>147ae65</code></a> Fix prettylog piping (<a href="https://redirect.github.com/rs/zerolog/issues/644">#644</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/4d78dc5ffa901e1c2c5905300ea701c9275f8c52"><code>4d78dc5</code></a> Add forwarding close methods to several writer implementations (<a href="https://redirect.github.com/rs/zerolog/issues/636">#636</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/c1ab4ed9bfdd5d1fb1b849c5223d8c43e6c189fc"><code>c1ab4ed</code></a> Make Log.Fatal() call Close on the writer before os.Exit(1) (<a href="https://redirect.github.com/rs/zerolog/issues/634">#634</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/417580d1ceea7b20325a9b8db31cd65640a7eaa2"><code>417580d</code></a> add: ContextLogger Interface LogObjectMarshaler (<a href="https://redirect.github.com/rs/zerolog/issues/623">#623</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/602e90aeea83612a96a15d50729b401facc27395"><code>602e90a</code></a> Fixed failing tests  (<a href="https://redirect.github.com/rs/zerolog/issues/626">#626</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/a9ec232b3e0c4cb1f9c7195359ab3f410d75d803"><code>a9ec232</code></a> Add stacktrace to Context (<a href="https://redirect.github.com/rs/zerolog/issues/630">#630</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/3e8ae07aba942a31de38f7e36d1f2ab79d22d82b"><code>3e8ae07</code></a> Refactor: change Hook(h Hook) to Hook(hooks ...Hook) (<a href="https://redirect.github.com/rs/zerolog/issues/629">#629</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/7fa45a4dda359d9a657a2960078097415417ec73"><code>7fa45a4</code></a> fixed documentation for tracing hook (<a href="https://redirect.github.com/rs/zerolog/issues/621">#621</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/93fb5cb2158f94a351c82b4f712017f8c8d758b8"><code>93fb5cb</code></a> Add TriggerLevelWriter. (<a href="https://redirect.github.com/rs/zerolog/issues/602">#602</a>)</li>
<li><a href="https://github.com/rs/zerolog/commit/83e03c75d923022dd5d22fd8baa469efcf0e9840"><code>83e03c7</code></a> stop using deprecated io/ioutils package (<a href="https://redirect.github.com/rs/zerolog/issues/620">#620</a>) (<a href="https://redirect.github.com/rs/zerolog/issues/620">#620</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/rs/zerolog/compare/v1.31.0...v1.32.0">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/driver/postgres` from 1.5.4 to 1.5.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/postgres/commit/b3b67da9ecabf3a6de0a81ef9cab4f26686492a7"><code>b3b67da</code></a> only retract v1.5.5</li>
<li><a href="https://github.com/go-gorm/postgres/commit/d715278f3f57d8f653bec8726cf845f8d608f2fd"><code>d715278</code></a> fix: retract v1.5.5 (<a href="https://redirect.github.com/go-gorm/postgres/issues/249">#249</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/ae7e52706ecbafb57bbba7bedded29e47f1028b6"><code>ae7e527</code></a> update gorm to master latest (<a href="https://redirect.github.com/go-gorm/postgres/issues/242">#242</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/1cfbc8652cd8c56d1cd05c1a3e111318b3dd551d"><code>1cfbc86</code></a> refactor: distinguish between Unique and UniqueIndex (<a href="https://redirect.github.com/go-gorm/postgres/issues/195">#195</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/dc711bd7c69394788e6532ad9c89f5bdcb12b007"><code>dc711bd</code></a> let gorm support  limit and offset binding parameters,change the BindVar of p...</li>
<li><a href="https://github.com/go-gorm/postgres/commit/438e4fdec86ee25732eddd756741bf3c9882a60c"><code>438e4fd</code></a> chore(deps): bump actions/setup-go from 4 to 5 (<a href="https://redirect.github.com/go-gorm/postgres/issues/227">#227</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/e6d2435140dc0305dae1e61780f3a55cdced47c7"><code>e6d2435</code></a> override identifier length for postgres (<a href="https://redirect.github.com/go-gorm/postgres/issues/232">#232</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/postgres/compare/v1.5.4...v1.5.6">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.25.5 to 1.25.7-0.20240204074919-46816ad31dde
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/commits">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.40.1 to 1.42.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.42.2 (2024-02-09)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<p>⚠️ The go.mod directive has been bumped to 1.18 as the minimum version of Go required for the module. This was necessary to continue to receive updates from some of the third party dependencies that Sarama makes use of for compression.</p>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: update go directive to 1.18 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2713">IBM/sarama#2713</a></li>
<li>feat: return KError instead of errors in AlterConfigs and DescribeConfig by <a href="https://github.com/zhuliquan"><code>@​zhuliquan</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2472">IBM/sarama#2472</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: don't waste time for backoff on member id required error by <a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2759">IBM/sarama#2759</a></li>
<li>fix: prevent ConsumerGroup.Close infinitely locking by <a href="https://github.com/maqdev"><code>@​maqdev</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2717">IBM/sarama#2717</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump golang.org/x/net from 0.17.0 to 0.18.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2716">IBM/sarama#2716</a></li>
<li>chore(deps): bump golang.org/x/sync to v0.5.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2718">IBM/sarama#2718</a></li>
<li>chore(deps): bump github.com/pierrec/lz4/v4 from 4.1.18 to 4.1.19 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2739">IBM/sarama#2739</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2748">IBM/sarama#2748</a></li>
<li>chore(deps): bump the golang-org-x group with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2734">IBM/sarama#2734</a></li>
<li>chore(deps): bump the golang-org-x group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2764">IBM/sarama#2764</a></li>
<li>chore(deps): bump github.com/pierrec/lz4/v4 from 4.1.19 to 4.1.21 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2763">IBM/sarama#2763</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 in /examples/exactly_once by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2749">IBM/sarama#2749</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 in /examples/consumergroup by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2750">IBM/sarama#2750</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 in /examples/sasl_scram_client by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2751">IBM/sarama#2751</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 in /examples/interceptors by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2752">IBM/sarama#2752</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 in /examples/http_server by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2753">IBM/sarama#2753</a></li>
<li>chore(deps): bump github.com/eapache/go-resiliency from 1.4.0 to 1.5.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2745">IBM/sarama#2745</a></li>
<li>chore(deps): bump golang.org/x/crypto from 0.15.0 to 0.17.0 in /examples/txn_producer by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2754">IBM/sarama#2754</a></li>
<li>chore(deps): bump go.opentelemetry.io/otel/sdk from 1.19.0 to 1.22.0 in /examples/interceptors by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2767">IBM/sarama#2767</a></li>
<li>chore(deps): bump the golang-org-x group with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2793">IBM/sarama#2793</a></li>
<li>chore(deps): bump go.opentelemetry.io/otel/exporters/stdout/stdoutmetric from 0.42.0 to 1.23.1 in /examples/interceptors by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2792">IBM/sarama#2792</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>fix(examples): housekeeping of code and deps by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2720">IBM/sarama#2720</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>fix(test): retry MockBroker Listen for EADDRINUSE by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2721">IBM/sarama#2721</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/maqdev"><code>@​maqdev</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2717">IBM/sarama#2717</a></li>
<li><a href="https://github.com/zhuliquan"><code>@​zhuliquan</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2472">IBM/sarama#2472</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.1...v1.42.2">https://github.com/IBM/sarama/compare/v1.42.1...v1.42.2</a></p>
<h2>Version 1.42.1 (2023-11-07)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: make fetchInitialOffset use correct protocol by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2705">IBM/sarama#2705</a></li>
<li>fix(config): relax ClientID validation after 1.0.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2706">IBM/sarama#2706</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1">https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
<h2>Version 1.42.1 (2023-11-07)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: make fetchInitialOffset use correct protocol by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2705">IBM/sarama#2705</a></li>
<li>fix(config): relax ClientID validation after 1.0.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2706">IBM/sarama#2706</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1">https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1</a></p>
<h2>Version 1.42.0 (2023-11-02)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>Asynchronously close brokers during a RefreshBrokers by <a href="https://github.com/bmassemin"><code>@​bmassemin</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2693">IBM/sarama#2693</a></li>
<li>Fix data race on Broker.done channel by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2698">IBM/sarama#2698</a></li>
<li>fix: data race in Broker.AsyncProduce by <a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2678">IBM/sarama#2678</a></li>
<li>Fix default retention time value in offset commit by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2700">IBM/sarama#2700</a></li>
<li>fix(txmgr): ErrOffsetsLoadInProgress is retriable by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2701">IBM/sarama#2701</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(ci): improve ossf scorecard result by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2685">IBM/sarama#2685</a></li>
<li>chore(ci): add kafka 3.6.0 to FVT and versions by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2692">IBM/sarama#2692</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>chore(ci): ossf scorecard.yml by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2683">IBM/sarama#2683</a></li>
<li>fix(ci): always run CodeQL on every commit by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2689">IBM/sarama#2689</a></li>
<li>chore(doc): add OpenSSF Scorecard badge by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2691">IBM/sarama#2691</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/bmassemin"><code>@​bmassemin</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2693">IBM/sarama#2693</a></li>
<li><a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2678">IBM/sarama#2678</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.3...v1.42.0">https://github.com/IBM/sarama/compare/v1.41.3...v1.42.0</a></p>
<h2>Version 1.41.3 (2023-10-17)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: pre-compile regex for parsing kafka version by <a href="https://github.com/qshuai"><code>@​qshuai</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2663">IBM/sarama#2663</a></li>
<li>fix(client): ignore empty Metadata responses when refreshing by <a href="https://github.com/HaoSunUber"><code>@​HaoSunUber</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2672">IBM/sarama#2672</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump the golang-org-x group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2661">IBM/sarama#2661</a></li>
<li>chore(deps): bump golang.org/x/net from 0.16.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2671">IBM/sarama#2671</a></li>
</ul>
<h3>:memo: Documentation</h3>
<ul>
<li>fix(docs): correct topic name in rebalancing strategy example by <a href="https://github.com/maksadbek"><code>@​maksadbek</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2657">IBM/sarama#2657</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/maksadbek"><code>@​maksadbek</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2657">IBM/sarama#2657</a></li>
<li><a href="https://github.com/qshuai"><code>@​qshuai</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2663">IBM/sarama#2663</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/5f63a84f47c39bf08a1c276f1f6b5f1d754e9fc3"><code>5f63a84</code></a> chore(deps): bump go.opentelemetry.io/otel/exporters/stdout/stdoutmetric (<a href="https://redirect.github.com/IBM/sarama/issues/2792">#2792</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/5ddd66176fdd0bc346af79d383a901bba4df8125"><code>5ddd661</code></a> chore(deps): bump the golang-org-x group with 1 update (<a href="https://redirect.github.com/IBM/sarama/issues/2793">#2793</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/d0bd279e8a7c9d759258d5092bda2b53dac3571e"><code>d0bd279</code></a> chore(ci): bump docker/bake-action (<a href="https://redirect.github.com/IBM/sarama/issues/2782">#2782</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/25c9c1a880e385781e1a39b49f8e7239e3d5e729"><code>25c9c1a</code></a> chore(ci): bump github/codeql-action from 2.22.9 to 3.24.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2783">#2783</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/f8b2129d5b7892baeb8ca2af9c10376e679639fc"><code>f8b2129</code></a> chore(deps): bump go.opentelemetry.io/otel/sdk in /examples/interceptors (<a href="https://redirect.github.com/IBM/sarama/issues/2767">#2767</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/7d7befd2855168ab4be8061dd8aa9e8c44840304"><code>7d7befd</code></a> chore(deps): bump golang.org/x/crypto in /examples/txn_producer (<a href="https://redirect.github.com/IBM/sarama/issues/2754">#2754</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/d44bd99f33efbbef553be9fd5a3ebacbc609611d"><code>d44bd99</code></a> chore(deps): bump github.com/eapache/go-resiliency from 1.4.0 to 1.5.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2745">#2745</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/b2999b45a3da08480e25e3979fd47fb18e00e587"><code>b2999b4</code></a> fix(test): retry MockBroker Listen for EADDRINUSE (<a href="https://redirect.github.com/IBM/sarama/issues/2721">#2721</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/d382099e8012ee7f271e5d2c0d4a81b41c841c83"><code>d382099</code></a> feat: return KError instead of errors in AlterConfigs and DescribeConfigs (<a href="https://redirect.github.com/IBM/sarama/issues/2">#2</a>...</li>
<li><a href="https://github.com/IBM/sarama/commit/5c3ffeacb37429233c7eae5602e013e0f84d40c2"><code>5c3ffea</code></a> fix: prevent ConsumerGroup.Close infinitely locking (<a href="https://redirect.github.com/IBM/sarama/issues/2717">#2717</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.40.1...v1.42.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.31.0 to 1.32.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-04-19)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/comprehend</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/comprehend/CHANGELOG.md#v1240-2023-04-19">v1.24.0</a>
<ul>
<li><strong>Feature</strong>: This release supports native document models for custom classification, in addition to plain-text models. You train native document models using documents (PDF, Word, images) in their native format.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/ecs/CHANGELOG.md#v1260-2023-04-19">v1.26.0</a>
<ul>
<li><strong>Feature</strong>: This release supports the Account Setting &quot;TagResourceAuthorization&quot; that allows for enhanced Tagging security controls.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ram</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/ram/CHANGELOG.md#v1180-2023-04-19">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for customer managed permissions. Customer managed permissions enable customers to author and manage tailored permissions for resources shared using RAM.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/rds/CHANGELOG.md#v1431-2023-04-19">v1.43.1</a>
<ul>
<li><strong>Documentation</strong>: Adds support for the ImageId parameter of CreateCustomDBEngineVersion to RDS Custom for Oracle</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/s3/CHANGELOG.md#v1320-2023-04-19">v1.32.0</a>
<ul>
<li><strong>Feature</strong>: Provides support for &quot;Snow&quot; Storage class.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/secretsmanager</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/secretsmanager/CHANGELOG.md#v1194-2023-04-19">v1.19.4</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Secrets Manager</li>
</ul>
</li>
</ul>
<h1>Release (2023-04-17)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/appflow/CHANGELOG.md#v1260-2023-04-17">v1.26.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a Client Token parameter to the following AppFlow APIs: Create/Update Connector Profile, Create/Update Flow, Start Flow, Register Connector, Update Connector Registration. The Client Token parameter allows idempotent operations for these APIs.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/drs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/drs/CHANGELOG.md#v1130-2023-04-17">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Changed existing APIs and added new APIs to support using an account-level launch configuration template with AWS Elastic Disaster Recovery.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/dynamodb</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/dynamodb/CHANGELOG.md#v1195-2023-04-17">v1.19.5</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for DynamoDB API</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/emrserverless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/emrserverless/CHANGELOG.md#v170-2023-04-17">v1.7.0</a>
<ul>
<li><strong>Feature</strong>: The GetJobRun API has been updated to include the job's billed resource utilization. This utilization shows the aggregate vCPU, memory and storage that AWS has billed for the job run. The billed resources include a 1-minute minimum usage for workers, plus additional storage over 20 GB per worker.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/internetmonitor</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/internetmonitor/CHANGELOG.md#v120-2023-04-17">v1.2.0</a>
<ul>
<li><strong>Feature</strong>: This release includes a new configurable value, TrafficPercentageToMonitor, which allows users to adjust the amount of traffic monitored by percentage</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotwireless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/iotwireless/CHANGELOG.md#v1270-2023-04-17">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: Supports the new feature of LoRaWAN roaming, allows to configure MaxEirp for LoRaWAN gateway, and allows to configure PingSlotPeriod for LoRaWAN multicast group</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lambda</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/lambda/CHANGELOG.md#v1330-2023-04-17">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: Add Python 3.10 (python3.10) support to AWS Lambda</li>
</ul>
</li>
</ul>
<h1>Release (2023-04-14)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/ecs/CHANGELOG.md#v1251-2023-04-14">v1.25.1</a>
<ul>
<li><strong>Documentation</strong>: This release supports  ephemeral storage for AWS Fargate Windows containers.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/lambda</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/lambda/CHANGELOG.md#v1320-2023-04-14">v1.32.0</a>
<ul>
<li><strong>Feature</strong>: This release adds SnapStart related exceptions to InvokeWithResponseStream API. IAM access related documentation is also added for this API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/migrationhubrefactorspaces/CHANGELOG.md#v198-2023-04-14">v1.9.8</a>
<ul>
<li><strong>Documentation</strong>: Doc only update for Refactor Spaces environments without network bridge feature.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.32.0/service/rds/CHANGELOG.md#v1430-2023-04-14">v1.43.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/86ce13f2993caadbbea28ec36e5c9ec7b5a8496f"><code>86ce13f</code></a> Release 2023-04-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e607ef176a31afd100e8f901ab9e2598319dc97f"><code>e607ef1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0135839133079a5f85e3ac63ba9f94097e0d9dbb"><code>0135839</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3bd75c8364f39e87c751630369f68640c4da695"><code>b3bd75c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3d01a6361729cb0833e4561c11d9dc702081ba2"><code>b3d01a6</code></a> Release 2023-04-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/30a27302cb1df28acd9c2a0e63becc12f0be3ec0"><code>30a2730</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7c3afdaca27d616b207cd1c4555b391a6c5987a8"><code>7c3afda</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/749f0c848ac9589e8fd6a50c14358c5633debca6"><code>749f0c8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b83b3053dda120e4687168a211cf5eca4c071bfe"><code>b83b305</code></a> Release 2023-04-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cdcc36abd06bd254c4b9e677511d0d6d93f307b3"><code>cdcc36a</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.31.0...service/s3/v1.32.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.1.1706018763 to 2024.2.1706885661
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/227d345e927f46afe6b5d10c438a1e4abc916138"><code>227d345</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e475934d9b341b7a98d596b8e32...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.1.1706018763...release/v2024.2.1706885661">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.26.0 to 0.27.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.27.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.27.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li><code>Exception.ThreadId</code> is now typed as <code>uint64</code>. It was wrongly typed as <code>string</code> before. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/770">#770</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Export <code>Event.Attachments</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/771">#771</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.27.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.27.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li><code>Exception.ThreadId</code> is now typed as <code>uint64</code>. It was wrongly typed as <code>string</code> before. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/770">#770</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Export <code>Event.Attachments</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/771">#771</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/ba344b4a9be534eb7a06866f33f9225009331bf9"><code>ba344b4</code></a> release: 0.27.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/49126c6921b6fbc2ee1b09fe3f059f533dc4b7da"><code>49126c6</code></a> Prepare 0.27.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/772">#772</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/30f4edc2a707a40894fae57107faed4ec350697d"><code>30f4edc</code></a> Export <code>Event.Attachments</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/771">#771</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/4185c9f7216fd062374f92a07beb2eb2edaa186f"><code>4185c9f</code></a> Fix Exception.ThreadID type (<a href="https://redirect.github.com/getsentry/sentry-go/issues/770">#770</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/824589bd5be49de6242d5f867a94c7df05510a4e"><code>824589b</code></a> Merge branch 'release/0.26.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.26.0...v0.27.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.5.2 to 5.5.3
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.3 (February 3, 2024)</h1>
<ul>
<li>Fix: prepared statement already exists</li>
<li>Improve CopyFrom auto-conversion of text-ish values</li>
<li>Add ltree type support (Florent Viel)</li>
<li>Make some properties of Batch and QueuedQuery public (Pavlo Golub)</li>
<li>Add AppendRows function (Edoardo Spadolini)</li>
<li>Optimize convert UUID [16]byte to string (Kirill Malikov)</li>
<li>Fix: LargeObject Read and Write of more than ~1GB at a time (Mitar)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/6f8f6ede6c67ce6b755feafe5c44f93577545147"><code>6f8f6ed</code></a> Update changelog for v5.5.3</li>
<li><a href="https://github.com/jackc/pgx/commit/576b6c88f631354d8deaddf9fd30cdb0632adb1a"><code>576b6c8</code></a> Bump actions/setup-go version</li>
<li><a href="https://github.com/jackc/pgx/commit/7caa448ac8858ee0c3fdfc15d11c32cf3bba34f3"><code>7caa448</code></a> Skip test on CockroachDB</li>
<li><a href="https://github.com/jackc/pgx/commit/832b4f97718c2d9d2eb16bbd2fef1d05ede7aab5"><code>832b4f9</code></a> Fix: prepared statement already exists</li>
<li><a href="https://github.com/jackc/pgx/commit/fd4411453fbd592586601b71f05eb06ea1c74906"><code>fd44114</code></a> Improve Conn.LoadType documentation</li>
<li><a href="https://github.com/jackc/pgx/commit/34da2fed9570ec3ff22dfa181323ef9cf702b2ca"><code>34da2fe</code></a> Improve CopyFrom auto-conversion of text-ish values</li>
<li><a href="https://github.com/jackc/pgx/commit/7b5fcac46526c55c6c3ed32812a0002104bae2c3"><code>7b5fcac</code></a> Add timetz and []timetz OID constants</li>
<li><a href="https://github.com/jackc/pgx/commit/0819a17da8863cd8bd6819dfed246d6101f67086"><code>0819a17</code></a> Remove openssl from TLS test setup</li>
<li><a href="https://github.com/jackc/pgx/commit/bf1c1d7848c087c99fb30a67657a5ec3197f7f32"><code>bf1c1d7</code></a> create ltree extension in pg setup for tests</li>
<li><a href="https://github.com/jackc/pgx/commit/0fa533386c98bd9ee1b492114d34f78b236ec2b7"><code>0fa5333</code></a> add ltree pgtype support</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.2...v5.5.3">compare view</a></li>
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

### Comment by @app-sre-bot on February 12, 2024 at 04:21 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on February 14, 2024 at 08:29 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/571*
