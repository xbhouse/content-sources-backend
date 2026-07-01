---
type: pull_request
number: 564
title: "Build: Bump the go group with 5 updates"
state: closed
author: dependabot
created: 2024-02-06T13:46:03Z
updated: 2024-02-06T13:50:01Z
closed: 2024-02-06T13:50:00Z
base_branch: main
head_branch: dependabot/go_modules/go-9d8d89fcee
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/564
---

# Pull Request #564: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: February 06, 2024 at 01:46 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9d8d89fcee`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.31.0` | `1.32.0` |
| [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) | `1.5.4` | `1.5.6` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.7-0.20240204074919-46816ad31dde` | `1.25.7` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.40.1` | `1.42.1` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.1.1706018763` | `2024.2.1706885661` |

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

Updates `gorm.io/gorm` from 1.25.7-0.20240204074919-46816ad31dde to 1.25.7
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/commits/v1.25.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.40.1 to 1.42.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.42.1 (2023-11-07)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: make fetchInitialOffset use correct protocol by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2705">IBM/sarama#2705</a></li>
<li>fix(config): relax ClientID validation after 1.0.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2706">IBM/sarama#2706</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1">https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1</a></p>
<h2>Version 1.42.0 (2023-11-02)</h2>
<!-- raw HTML omitted -->
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
<!-- raw HTML omitted -->
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
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3">https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
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
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3">https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/d88a48a6a6fd6cf1bd89e76192216378379fe799"><code>d88a48a</code></a> chore: update CHANGELOG.md to v1.42.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2711">#2711</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/385b3b435e85e6e3d1e29299b49330a0e9954f98"><code>385b3b4</code></a> fix(config): relax ClientID validation after 1.0.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2706">#2706</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3364ff02d7e986ba0eb75d1c2dbb9a62a614dd15"><code>3364ff0</code></a> chore(doc): add CODE_OF_CONDUCT.md</li>
<li><a href="https://github.com/IBM/sarama/commit/768496e6e6f2e9de5cfb2d1fca3534a6bb737f92"><code>768496e</code></a> chore(ci): bump actions/dependency-review-action from 3.1.0 to 3.1.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2707">#2707</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/27710af1531a2f9a57dc21d4291441be6df274c6"><code>27710af</code></a> fix: make fetchInitialOffset use correct protocol (<a href="https://redirect.github.com/IBM/sarama/issues/2705">#2705</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a46917f5b87dea431c327b5725e37bd843a2dc7a"><code>a46917f</code></a> chore(ci): bump actions/dependency-review-action from 2.5.1 to 3.1.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2702">#2702</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4168f7ccba85976c3b03fe01d3d693c876c617a2"><code>4168f7c</code></a> chore(ci): bump ossf/scorecard-action from 2.1.2 to 2.3.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2703">#2703</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/7155d51d84883025253f297233f95ba13683de33"><code>7155d51</code></a> chore(ci): add kafka 3.6.0 to FVT and versions</li>
<li><a href="https://github.com/IBM/sarama/commit/e0c3c627b80ae4468f9e1dee60306dc2ce3fb284"><code>e0c3c62</code></a> fix(txmgr): ErrOffsetsLoadInProgress is retriable</li>
<li><a href="https://github.com/IBM/sarama/commit/2e077cf8d86cc06bffc2614b80a699caa868ad2b"><code>2e077cf</code></a> Fix default retention time value in offset commit (<a href="https://redirect.github.com/IBM/sarama/issues/2700">#2700</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.40.1...v1.42.1">compare view</a></li>
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

### Comment by @app-sre-bot on February 06, 2024 at 01:48 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on February 06, 2024 at 01:48 PM UTC

@dependabot  unignore github.com/jackc/pgx/v5

### Comment by @dependabot on February 06, 2024 at 01:48 PM UTC

OK, I will stop ignoring the github.com/jackc/pgx/v5 dependency.

### Comment by @dependabot on February 06, 2024 at 01:49 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/564*
