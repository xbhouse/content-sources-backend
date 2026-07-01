---
type: pull_request
number: 578
title: "Build: Bump the go group with 11 updates"
state: closed
author: dependabot
created: 2024-02-14T08:35:48Z
updated: 2024-02-19T04:40:20Z
closed: 2024-02-19T04:40:19Z
base_branch: main
head_branch: dependabot/go_modules/go-34d9ab1422
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/578
---

# Pull Request #578: Build: Bump the go group with 11 updates

**Author**: @dependabot
**Created**: February 14, 2024 at 08:35 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-34d9ab1422`

## Description

Bumps the go group with 11 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.31.0` | `1.32.0` |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.4` | `1.5.6` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.5` | `1.25.7-0.20240204074919-46816ad31dde` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.40.1` | `1.42.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.24.1` | `1.25.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.16.16` | `1.17.0` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.31.0` | `1.33.0` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.1.1706018763` | `2024.2.1707758687` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.26.0` | `0.27.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.2` | `5.5.3` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.3` | `1.25.4` |

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

Updates `github.com/aws/aws-sdk-go-v2` from 1.24.1 to 1.25.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4fd9126607b61a0fe22af3a6a31b4e4792a66d83"><code>4fd9126</code></a> Release 2024-02-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d22cecd4af5dc762e59c2bd156e3a774cbc1f2db"><code>d22cecd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b640bf508c997be09b65a3b0d7f2be07ee1e7ac6"><code>b640bf5</code></a> Update SDK's smithy-go dependency to v1.20.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94e885c36103117b48764e0e94d88da729b84047"><code>94e885c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3a6cd78eabbe87d11278117682674044d1e4408"><code>b3a6cd7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f328e6ca9ec4effe5f82d79c2076c59788c05b8"><code>5f328e6</code></a> chore: bump min go to 1.20 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2494">#2494</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3cb2c978aed75796d8db56dab5f496329af8b182"><code>3cb2c97</code></a> ci: fix CI scripts for main branch case (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2491">#2491</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e44db4cf355dbb0324cba2ee49cafed1454ba26"><code>0e44db4</code></a> Release 2024-02-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1d8fb3bda76e7fa0373768c10f90aa2aa5281e4a"><code>1d8fb3b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4d556f1d9af36afc7d5109010e3330975877f13"><code>e4d556f</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.24.1...v1.25.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.16 to 1.17.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-10-21)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.17.0
<ul>
<li><strong>Feature</strong>: Adds <code>aws.IsCredentialsProvider</code> for inspecting <code>CredentialProvider</code> types when needing to determine if the underlying implementation type matches a target type. This resolves an issue where <code>CredentialsCache</code> could mask <code>AnonymousCredentials</code> providers, breaking downstream detection logic.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/cognitoidentityprovider/CHANGELOG.md#v1210-2022-10-21">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new &quot;DeletionProtection&quot; field to the UserPool in Cognito. Application admins can configure this value with either ACTIVE or INACTIVE value. Setting this field to ACTIVE will prevent a user pool from accidental deletion.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/eventbridge</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/eventbridge/CHANGELOG.md#v11616-2022-10-21">v1.16.16</a>
<ul>
<li><strong>Bug Fix</strong>: The SDK client has been updated to utilize the <code>aws.IsCredentialsProvider</code> function for determining if <code>aws.AnonymousCredentials</code> has been configured for the <code>CredentialProvider</code>.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/s3/CHANGELOG.md#v1290-2022-10-21">v1.29.0</a>
<ul>
<li><strong>Feature</strong>: S3 on Outposts launches support for automatic bucket-style alias. You can use the automatic access point alias instead of an access point ARN for any object-level operation in an Outposts bucket.</li>
<li><strong>Bug Fix</strong>: The SDK client has been updated to utilize the <code>aws.IsCredentialsProvider</code> function for determining if <code>aws.AnonymousCredentials</code> has been configured for the <code>CredentialProvider</code>.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/sagemaker/CHANGELOG.md#v1490-2022-10-21">v1.49.0</a>
<ul>
<li><strong>Feature</strong>: CreateInferenceRecommenderjob API now supports passing endpoint details directly, that will help customers to identify the max invocation and max latency they can achieve for their model and the associated endpoint along with getting recommendations on other instances.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sts</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/sts/CHANGELOG.md#v1170-2022-10-21">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Add presign functionality for sts:AssumeRole operation</li>
</ul>
</li>
</ul>
<h1>Release (2022-10-20)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/devopsguru</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/devopsguru/CHANGELOG.md#v1200-2022-10-20">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: This release adds information about the resources DevOps Guru is analyzing.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/globalaccelerator</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/globalaccelerator/CHANGELOG.md#v1150-2022-10-20">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Global Accelerator now supports AddEndpoints and RemoveEndpoints operations for standard endpoint groups.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/resiliencehub</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/resiliencehub/CHANGELOG.md#v170-2022-10-20">v1.7.0</a>
<ul>
<li><strong>Feature</strong>: In this release, we are introducing support for regional optimization for AWS Resilience Hub applications. It also includes a few documentation updates to improve clarity.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rum</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/rum/CHANGELOG.md#v170-2022-10-20">v1.7.0</a>
<ul>
<li><strong>Feature</strong>: CloudWatch RUM now supports Extended CloudWatch Metrics with Additional Dimensions</li>
</ul>
</li>
</ul>
<h1>Release (2022-10-19)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/chimesdkmessaging/CHANGELOG.md#v1116-2022-10-19">v1.11.6</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Chime Messaging SDK</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudtrail</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/cloudtrail/CHANGELOG.md#v1190-2022-10-19">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: This release includes support for exporting CloudTrail Lake query results to an Amazon S3 bucket.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/configservice/CHANGELOG.md#v1270-2022-10-19">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: This release adds resourceType enums for AppConfig, AppSync, DataSync, EC2, EKS, Glue, GuardDuty, SageMaker, ServiceDiscovery, SES, Route53 types.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/connect/CHANGELOG.md#v1330-2022-10-19">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: This release adds API support for managing phone numbers that can be used across multiple AWS regions through telephony traffic distribution.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/managedblockchain</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/managedblockchain/CHANGELOG.md#v1130-2022-10-19">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Adding new Accessor APIs for Amazon Managed Blockchain</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.17.0/service/s3/CHANGELOG.md#v1280-2022-10-19">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: Updates internal logic for constructing API endpoints. We have added rule-based endpoints and internal model parameters.</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/765544941191825edd26162f9790bf11f059d426"><code>7655449</code></a> Release 2022-10-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dcae829ecc334f91502afd6d7ae2295861db9885"><code>dcae829</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b82766b858e595943b26924ad1f107cd04363d66"><code>b82766b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1c05fb6452a1f74985ff6deb7a642b9eb441274a"><code>1c05fb6</code></a> Implements IsCredentialsProvider for checking if a provider matches a target ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0fab39aa32e09221c21383a8e658ef94d240b7e4"><code>0fab39a</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1888">#1888</a> from aws/isvita/issues-1787</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56eb993a7dfc0195e4ca3fca306d4985f6c7c99a"><code>56eb993</code></a> added changelog file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cde8cbcb42326be0c822cc557405619f1cefd92a"><code>cde8cbc</code></a> Release 2022-10-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d7765f947faec45f0b4e22bb8051ca3a1c337d93"><code>d7765f9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9dab7e94b7446829a09a1f7a302251bac37771d"><code>b9dab7e</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93ed3ee62605d8af3021010c5ddbf2141574af00"><code>93ed3ee</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.16...v1.17.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.31.0 to 1.33.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-04-24)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.18.0
<ul>
<li><strong>Feature</strong>: add recursion detection middleware to all SDK requests to avoid recursion invocation in Lambda</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/appflow/CHANGELOG.md#v1270-2023-04-24">v1.27.0</a>
<ul>
<li><strong>Feature</strong>: Increased the max length for RefreshToken and AuthCode from 2048 to 4096.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/codecatalyst</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/codecatalyst/CHANGELOG.md#v125-2023-04-24">v1.2.5</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Amazon CodeCatalyst.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ec2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/ec2/CHANGELOG.md#v1940-2023-04-24">v1.94.0</a>
<ul>
<li><strong>Feature</strong>: API changes to AWS Verified Access related to identity providers' information.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/mediaconvert</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/mediaconvert/CHANGELOG.md#v1360-2023-04-24">v1.36.0</a>
<ul>
<li><strong>Feature</strong>: This release introduces a noise reduction pre-filter, linear interpolation deinterlace mode, video pass-through, updated default job settings, and expanded LC-AAC Stereo audio bitrate ranges.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rekognition</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/rekognition/CHANGELOG.md#v1250-2023-04-24">v1.25.0</a>
<ul>
<li><strong>Feature</strong>: Added new status result to Liveness session status.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/route53</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/route53/CHANGELOG.md#v1280-2023-04-24">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: added paginator for listResourceRecordSets</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/s3</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/s3/CHANGELOG.md#v1330-2023-04-24">v1.33.0</a>
<ul>
<li><strong>Feature</strong>: added custom paginators for listMultipartUploads and ListObjectVersions</li>
</ul>
</li>
</ul>
<h1>Release (2023-04-21)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/connect</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/connect/CHANGELOG.md#v1520-2023-04-21">v1.52.0</a>
<ul>
<li><strong>Feature</strong>: This release adds a new API CreateParticipant. For Amazon Connect Chat, you can use this new API to customize chat flow experiences.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/ecs/CHANGELOG.md#v1261-2023-04-21">v1.26.1</a>
<ul>
<li><strong>Documentation</strong>: Documentation update to address various Amazon ECS tickets.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/fms</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/fms/CHANGELOG.md#v1230-2023-04-21">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: AWS Firewall Manager adds support for multiple administrators. You can now delegate more than one administrator per organization.</li>
</ul>
</li>
</ul>
<h1>Release (2023-04-20)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chime</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/chime/CHANGELOG.md#v1230-2023-04-20">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for Hindi and Thai languages and additional Amazon Transcribe parameters to the StartMeetingTranscription API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/chimesdkmediapipelines/CHANGELOG.md#v140-2023-04-20">v1.4.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for specifying the recording file format in an S3 recording sink configuration.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/chimesdkmeetings/CHANGELOG.md#v1150-2023-04-20">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for Hindi and Thai languages and additional Amazon Transcribe parameters to the StartMeetingTranscription API.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/gamelift</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/gamelift/CHANGELOG.md#v1180-2023-04-20">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Amazon GameLift supports creating Builds for Windows 2016 operating system.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/guardduty</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/guardduty/CHANGELOG.md#v1210-2023-04-20">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for the new Lambda Protection feature.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iot</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/iot/CHANGELOG.md#v1360-2023-04-20">v1.36.0</a>
<ul>
<li><strong>Feature</strong>: Support additional OTA states in GetOTAUpdate API</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.33.0/service/sagemaker/CHANGELOG.md#v1740-2023-04-20">v1.74.0</a>
<ul>
<li><strong>Feature</strong>: Amazon SageMaker Canvas adds ModelRegisterSettings support for CanvasAppSettings.</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/60a0ef0c0a0c3bc3abebf6f269162d1bc0f78cff"><code>60a0ef0</code></a> Release 2023-04-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/08a6461bf0d204cd7ce1ade0575a6c11a17228e6"><code>08a6461</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/31f1bdacb89b9735596a6022c00ee48a7981e308"><code>31f1bda</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c9ad0d0754d1d06c3e9c0cee4d3ace49603a5f83"><code>c9ad0d0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aadb5c6f004afaae30372c043a1cafb8078cd7ad"><code>aadb5c6</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2106">#2106</a> from syall/apigw-exports-nullability-mediaconvert-fix</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44acf0cb2f6a0a93ba51c7d345e68bfea447780f"><code>44acf0c</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2105">#2105</a> from aws/recursion-detection</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6c0e9379bcd832b35fd7ff0b576cf9b113da04e5"><code>6c0e937</code></a> add runtime checks</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6c2cf2a0b335b3f8e2f9fd8636c8a3e3ff34eae1"><code>6c2cf2a</code></a> Fix APIGW exports</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/75061a435a5c35f9347b1dadf2345d2437cb2ee0"><code>75061a4</code></a> capitalize recursion detection trace ID var name</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6f809488a24e385a1a9453dce38f94aa0edc1ee2"><code>6f80948</code></a> modify recursion detection and test</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.31.0...service/s3/v1.33.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.1.1706018763 to 2024.2.1707758687
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/3084bb3c3fff756ef78ccfa21672749b338dfe4f"><code>3084bb3</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27aa23486951b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/67fc5be09a79ba2da94597d7c9289628c605de38"><code>67fc5be</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e475939e4863cb7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/c430716af479cc3b512a6ed006bda7d60e301927"><code>c430716</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/10">#10</a> from jlsherrill/newlines</li>
<li><a href="https://github.com/content-services/zest/commit/e8df4658805746553e66e3c7aaca3be7eee3784b"><code>e8df465</code></a> Use temp file for fixing json</li>
<li><a href="https://github.com/content-services/zest/commit/69932164e3f2ec353e53d1a62eaa077788eb7dcb"><code>6993216</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/9">#9</a> from jlsherrill/newlines</li>
<li><a href="https://github.com/content-services/zest/commit/12eca0e6b2bb3e18e09b750ca1aad47a6117e494"><code>12eca0e</code></a> fix \n strings in doc</li>
<li><a href="https://github.com/content-services/zest/commit/227d345e927f46afe6b5d10c438a1e4abc916138"><code>227d345</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e475934d9b341b7a98d596b8e32...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.1.1706018763...release/v2024.2.1707758687">compare view</a></li>
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

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.3 to 1.25.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>Release v1.25.4</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/aws/aws-sdk-go from 1.49.19 to 1.49.23 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/409">RedHatInsights/insights-operator-utils#409</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.49.23 to 1.49.24 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/410">RedHatInsights/insights-operator-utils#410</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.49.24 to 1.50.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/414">RedHatInsights/insights-operator-utils#414</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.50.2 to 1.50.10 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/424">RedHatInsights/insights-operator-utils#424</a></li>
<li>Bump github.com/rs/zerolog from 1.31.0 to 1.32.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/423">RedHatInsights/insights-operator-utils#423</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.50.10 to 1.50.16 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/430">RedHatInsights/insights-operator-utils#430</a></li>
<li>Handle Kafka Broker addresses as string instead of slice by <a href="https://github.com/epapbak"><code>@​epapbak</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/431">RedHatInsights/insights-operator-utils#431</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.3...v1.25.4">https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.3...v1.25.4</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/1c8925348a35854f2c27d439a3745bf6e6f90588"><code>1c89253</code></a> Handle Kafka Broker addresses as string instead of slice</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/68803c21d92fb48b9aaa08a3257a6a783e12e1cd"><code>68803c2</code></a> Bump github.com/aws/aws-sdk-go from 1.50.10 to 1.50.16 (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/430">#430</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/2f2d614447e21454c711c793947ce26c06832ceb"><code>2f2d614</code></a> Bump github.com/rs/zerolog from 1.31.0 to 1.32.0 (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/423">#423</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/74d5a9bf5654598e7783c2d4eeaf4d2fb2bea025"><code>74d5a9b</code></a> Bump github.com/aws/aws-sdk-go from 1.50.2 to 1.50.10 (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/424">#424</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/f5c28f578f2b96a3a9caafbc538d29db13365636"><code>f5c28f5</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/414">#414</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/132f822796a9b2a137065b31268f4258b2853ac3"><code>132f822</code></a> Bump github.com/aws/aws-sdk-go from 1.49.24 to 1.50.2</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/181c22753b0b721b12acb68be8bedc0c6da4e00a"><code>181c227</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/410">#410</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/171ffc56f2a569f46f48dd5718a1e2dd2a8f3cc4"><code>171ffc5</code></a> Bump github.com/aws/aws-sdk-go from 1.49.23 to 1.49.24</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/badbd1278f90c4032d46dd33b7d7e7f03efd06b9"><code>badbd12</code></a> Bump github.com/aws/aws-sdk-go from 1.49.19 to 1.49.23</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.3...v1.25.4">compare view</a></li>
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

### Comment by @app-sre-bot on February 14, 2024 at 08:36 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on February 19, 2024 at 04:40 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/578*
