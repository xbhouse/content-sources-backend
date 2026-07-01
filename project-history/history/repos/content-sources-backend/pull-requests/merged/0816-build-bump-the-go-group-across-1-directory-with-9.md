---
type: pull_request
number: 816
title: "Build: Bump the go group across 1 directory with 9 updates"
state: merged
author: dependabot
created: 2024-09-16T04:10:22Z
updated: 2024-09-17T11:38:27Z
closed: 2024-09-17T11:38:23Z
merged: 2024-09-17T11:38:23Z
base_branch: main
head_branch: dependabot/go_modules/go-02adbe50e6
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/816
---

# Pull Request #816: Build: Bump the go group across 1 directory with 9 updates

**Author**: @dependabot
**Created**: September 16, 2024 at 04:10 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-02adbe50e6`

## Description

Bumps the go group with 9 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) | `4.17.1` | `4.18.1` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.25.11` | `1.25.12` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.20.2` | `1.20.3` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.30.4` | `1.30.5` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.30` | `1.17.32` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.37.5` | `1.39.0` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.8.1723836162` | `2024.9.1726259707` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.28.1` | `0.29.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.6.0` | `5.7.1` |


Updates `github.com/golang-migrate/migrate/v4` from 4.17.1 to 4.18.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.18.1</h2>
<h2>Changes</h2>
<ul>
<li>Update dktest from v0.4.2 to v0.4.3</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/golang-migrate/migrate/compare/v4.18.0...v4.18.1">https://github.com/golang-migrate/migrate/compare/v4.18.0...v4.18.1</a></p>
<h2>v4.18.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump golang.org/x/net from 0.21.0 to 0.23.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1074">golang-migrate/migrate#1074</a></li>
<li>Make MySQL SetVersion compatible with sql_safe_update by <a href="https://github.com/maxmati"><code>@​maxmati</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1070">golang-migrate/migrate#1070</a></li>
<li><a href="https://redirect.github.com/golang-migrate/migrate/issues/1104">#1104</a> Fix golanglint-ci lint config &amp; lint errors by <a href="https://github.com/Shion1305"><code>@​Shion1305</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1105">golang-migrate/migrate#1105</a></li>
<li>🧑‍💻 improve error message for invalid source, database, resolves: <a href="https://redirect.github.com/golang-migrate/migrate/issues/1102">#1102</a> by <a href="https://github.com/Shion1305"><code>@​Shion1305</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1103">golang-migrate/migrate#1103</a></li>
<li>Upgrade go-sqlite3 to v1.14.22 by <a href="https://github.com/gjabell"><code>@​gjabell</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1098">golang-migrate/migrate#1098</a></li>
<li>Fix redshift tests by <a href="https://github.com/dhui"><code>@​dhui</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1151">golang-migrate/migrate#1151</a></li>
<li>chore: remove deprecated <code>rand.Seed()</code> in testing.docker by <a href="https://github.com/joschi"><code>@​joschi</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1149">golang-migrate/migrate#1149</a></li>
<li>build: use Go 1.23 to build the project by <a href="https://github.com/joschi"><code>@​joschi</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1133">golang-migrate/migrate#1133</a></li>
<li>fix(tests): fix Docker imports by <a href="https://github.com/joschi"><code>@​joschi</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1153">golang-migrate/migrate#1153</a></li>
<li>test(postgres): run tests with PostgreSQL 16 by <a href="https://github.com/joschi"><code>@​joschi</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1126">golang-migrate/migrate#1126</a></li>
<li>test(mysql): run tests with MySQL 8.0, 8.4, and 9.0 by <a href="https://github.com/joschi"><code>@​joschi</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1127">golang-migrate/migrate#1127</a></li>
<li>Support .deb package for Noble Numbat (24.04) by <a href="https://github.com/muzammilar"><code>@​muzammilar</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1119">golang-migrate/migrate#1119</a></li>
<li>Bump google.golang.org/grpc from 1.64.0 to 1.64.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1152">golang-migrate/migrate#1152</a></li>
<li>CORRECTION DONE near m.Step(2) if you want to explicitly set the number of migrations to run by <a href="https://github.com/adityassharma-ss"><code>@​adityassharma-ss</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1101">golang-migrate/migrate#1101</a></li>
</ul>
<p>A special thanks to <a href="https://github.com/joschi"><code>@​joschi</code></a> fixing tests and builds!</p>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/maxmati"><code>@​maxmati</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1070">golang-migrate/migrate#1070</a></li>
<li><a href="https://github.com/Shion1305"><code>@​Shion1305</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1105">golang-migrate/migrate#1105</a></li>
<li><a href="https://github.com/gjabell"><code>@​gjabell</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1098">golang-migrate/migrate#1098</a></li>
<li><a href="https://github.com/joschi"><code>@​joschi</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1149">golang-migrate/migrate#1149</a></li>
<li><a href="https://github.com/muzammilar"><code>@​muzammilar</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1119">golang-migrate/migrate#1119</a></li>
<li><a href="https://github.com/adityassharma-ss"><code>@​adityassharma-ss</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1101">golang-migrate/migrate#1101</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/golang-migrate/migrate/compare/v4.17.1...v4.17.2">https://github.com/golang-migrate/migrate/compare/v4.17.1...v4.17.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/555501f7cdbf52a34bf66d49eb37eb72b22b2a85"><code>555501f</code></a> Update dktest from v0.4.2 to v0.4.3</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/d5eb594426c00e26024d5824ee2db92b23de140a"><code>d5eb594</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1101">#1101</a> from adityassharma-ss/patch-1</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/67c71f9a9d602a600873762f49e643d3975d89e7"><code>67c71f9</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1152">#1152</a> from golang-migrate/dependabot/go_modules/google.gol...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/1e02b42a7a32cf14f3c0068eb32eb8591f70c97b"><code>1e02b42</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1119">#1119</a> from muzammilar/add-noble-numbat</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/5b6f623dec0bc0c804fbcea5f2b21367f40e62ce"><code>5b6f623</code></a> Bump google.golang.org/grpc from 1.64.0 to 1.64.1</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/ddd7fa61dad2e5ecf6785eb1b03ead5bc1af83c9"><code>ddd7fa6</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1127">#1127</a> from joschi/tests-mysql-9</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/7e8f6be051309f99a8ea84ee7b2c69da0623cedf"><code>7e8f6be</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1126">#1126</a> from joschi/tests-postgres-16</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/65a3bd56cca7a6ad6072e4593fbda57f79eac97c"><code>65a3bd5</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1153">#1153</a> from joschi/fix-docker-imports</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/e40e64c4acea524918036dcf47d2600c818ff84d"><code>e40e64c</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1133">#1133</a> from joschi/go-1.23</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/6ceb5a946734e7c80636462c5c3830175c8edb3c"><code>6ceb5a9</code></a> fix(tests): fix Docker imports</li>
<li>Additional commits viewable in <a href="https://github.com/golang-migrate/migrate/compare/v4.17.1...v4.18.1">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.25.11 to 1.25.12
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/0daaf1747cfa4e4850376ad50a7834fb78b0cc0e"><code>0daaf17</code></a> fix: AfterQuery using safer right trim while clearing from clause's join adde...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/0dbfda5d7eb8f97ce524b8b692aa9fbee7b4e693"><code>0dbfda5</code></a> fix memory leaks in PrepareStatementDB (<a href="https://redirect.github.com/go-gorm/gorm/issues/7142">#7142</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.25.11...v1.25.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.20.2 to 1.20.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.20.3</h2>
<ul>
<li>[BUGFIX] histograms: Fix possible data race when appending exemplars. <a href="https://redirect.github.com/prometheus/client_golang/issues/1608">#1608</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/v1.20.3/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.20.3 / 2024-09-05</h2>
<ul>
<li>[BUGFIX] histograms: Fix possible data race when appending exemplars. <a href="https://redirect.github.com/prometheus/client_golang/issues/1608">#1608</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/ef2f87ea986252194ea960187b20b409180044dd"><code>ef2f87e</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1620">#1620</a> from prometheus/arthursens/prepare-1.20.3</li>
<li><a href="https://github.com/prometheus/client_golang/commit/937ac63d3d2dda83847f4ca842d62edabce4e743"><code>937ac63</code></a> Add changelog entry for 1.20.3</li>
<li><a href="https://github.com/prometheus/client_golang/commit/6e9914db5af255f5def17d54a7ca9c531771f4ca"><code>6e9914d</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1608">#1608</a> from krajorama/index-out-of-range-native-histogram-e...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/d6b8c8925bd16626cf168e642eb70724b17a0d61"><code>d6b8c89</code></a> Update comments with more explanations</li>
<li><a href="https://github.com/prometheus/client_golang/commit/504566f07c680f68743c3a5d239dede48538c7ec"><code>504566f</code></a> Use simplified solution from <a href="https://redirect.github.com/prometheus/client_golang/issues/1609">#1609</a> for the data race</li>
<li><a href="https://github.com/prometheus/client_golang/commit/dc8e9a4d8a4c7c64d5ae2c9d29a91bb1407d549b"><code>dc8e9a4</code></a> fix: native histogram: Simplify and fix addExemplar</li>
<li><a href="https://github.com/prometheus/client_golang/commit/dc819ceb1b0f906f1ab124f7492693970733a54d"><code>dc819ce</code></a> Use a trivial solution to <a href="https://redirect.github.com/prometheus/client_golang/issues/1605">#1605</a></li>
<li><a href="https://github.com/prometheus/client_golang/commit/e061dfae88c0dd63ff477a153096a1ba28f69f7e"><code>e061dfa</code></a> native histogram: use exemplars in concurrency test</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.20.2...v1.20.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.30.4 to 1.30.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2b751d1ba71f59175a41f9cae5f159f1044360f"><code>a2b751d</code></a> Release 2024-09-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e22c2495bd38232c640776ef3c1a84fb3145a8b9"><code>e22c249</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff0cf6fbfa7c7a12388606d5568135f2beae6599"><code>ff0cf6f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3120376762853b3098fda7e9217fb39fe1bf5105"><code>3120376</code></a> refactoring of buildQuery to accept a list of maintained headers to l… (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2773">#2773</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ed838eab2a963cb16301501c8b8c3e29dac4c20"><code>4ed838e</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2768">#2768</a> from bhavya2109sharma/presignedurl-requestpayer-change</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d4bd42fdc82c2954f429bd9eeac3096f5590eaac"><code>d4bd42f</code></a> Merge branch 'main' into presignedurl-requestpayer-change</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0353706229a89749afa93324432ece53da37822b"><code>0353706</code></a> Added Changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/97e2d3fec04bc8bfd69eaf22ce476b12f8673810"><code>97e2d3f</code></a> Release 2024-08-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4cca52b4a81df57e91b1e5d0a65fd6df89606d02"><code>4cca52b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a5146734eb1e6d59acba18b5967c78dc7a5e42"><code>c8a5146</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.4...v1.30.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.30 to 1.17.32
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f1d71c59a1499981873f70db8c92d07242779670"><code>f1d71c5</code></a> Release 2024-09-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4813e114f3c2ce8db07624f656073440adc1140"><code>e4813e1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e8bb90d8c0d03824fc8bf8656bd2de6346e48f8"><code>0e8bb90</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6a5875b13d01e8270c1cbdfae0c63ee479386917"><code>6a5875b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f7030de42b2a92ad4f66b153ae4d88ad28a6b2fe"><code>f7030de</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2b751d1ba71f59175a41f9cae5f159f1044360f"><code>a2b751d</code></a> Release 2024-09-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e22c2495bd38232c640776ef3c1a84fb3145a8b9"><code>e22c249</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff0cf6fbfa7c7a12388606d5568135f2beae6599"><code>ff0cf6f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3120376762853b3098fda7e9217fb39fe1bf5105"><code>3120376</code></a> refactoring of buildQuery to accept a list of maintained headers to l… (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2773">#2773</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ed838eab2a963cb16301501c8b8c3e29dac4c20"><code>4ed838e</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2768">#2768</a> from bhavya2109sharma/presignedurl-requestpayer-change</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.30...credentials/v1.17.32">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.37.5 to 1.39.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/807c921d812ada1f36e8e016035a375ee20b6379"><code>807c921</code></a> Release 2023-09-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/39b991de3a1abe48d06850b17381a53ab8c345e8"><code>39b991d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b1d6dafaff6a5fe77514b0329508ed467404f282"><code>b1d6daf</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5771d2244ce2f85d536e4ee5c9f014b9b7cfff4a"><code>5771d22</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7143703367f9e82e2add61e58fcf19dfe3d7cb43"><code>7143703</code></a> Release 2023-09-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1803c873fb010d3991d36927eafc5a8683bc9ea4"><code>1803c87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/228b865841787ce518d849ac7584225b65b22518"><code>228b865</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f416d6651c5eb1850565963502dca48f946ec1e0"><code>f416d66</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0b396e56da1e19abc0855af33cdfbf3c7531882"><code>f0b396e</code></a> Release 2023-09-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/53896255622f253e3273ffe2bada16ab94a199d0"><code>5389625</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.37.5...service/s3/v1.39.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.8.1723836162 to 2024.9.1726259707
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/8081eef6769be5ffbc3f947404c68c083107cb70"><code>8081eef</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e4755222d62bcb7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/c834534462e1820818e54088e62dc725ab30c1da"><code>c834534</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4782d2b5539067968b5e2da9ab...</li>
<li><a href="https://github.com/content-services/zest/commit/57dfc37593857ba396674cf7439f5aafe56b4ebb"><code>57dfc37</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7ebd696bd5f87e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/689da9b325b777d7b3feabb6e0b796462eca4412"><code>689da9b</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/19">#19</a> from jlsherrill/template_fix</li>
<li><a href="https://github.com/content-services/zest/commit/38e4c26c6b19bfd9564e6530ea614e9c580c2f1f"><code>38e4c26</code></a> update template for style and upgrade to go 1.23</li>
<li><a href="https://github.com/content-services/zest/commit/e93b773dac373affa5c40e23227a11c0fc2cb383"><code>e93b773</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/18">#18</a> from jlsherrill/revert123</li>
<li><a href="https://github.com/content-services/zest/commit/bc7548258fbd708238df765659ce840b364e65d5"><code>bc75482</code></a> Revert &quot;Upgrade to go 1.23&quot;</li>
<li><a href="https://github.com/content-services/zest/commit/c4401a5b54f5877592dda2ce5a771e8aca7f2142"><code>c4401a5</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/17">#17</a> from jlsherrill/123</li>
<li><a href="https://github.com/content-services/zest/commit/5c638497e6c7364e540893db465d7fa403d55eb3"><code>5c63849</code></a> Upgrade to go 1.23</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.8.1723836162...release/v2024.9.1726259707">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.28.1 to 0.29.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.29.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.29.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Remove the <code>sentrymartini</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/861">#861</a>)</li>
<li>The <code>WrapResponseWriter</code> has been moved from the <code>sentryhttp</code> package to the <code>internal/httputils</code> package. If you've imported it previosuly, you'll need to copy the implementation in your project. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/871">#871</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add new convenience methods to continue a trace and propagate tracing headers for error-only use cases. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/862">#862</a>)</p>
<p>If you are not using one of our integrations, you can manually continue an incoming trace by using <code>sentry.ContinueTrace()</code> by providing the <code>sentry-trace</code> and <code>baggage</code> header received from a downstream SDK.</p>
<pre lang="go"><code>hub := sentry.CurrentHub()
sentry.ContinueTrace(hub, r.Header.Get(sentry.SentryTraceHeader), r.Header.Get(sentry.SentryBaggageHeader)),
</code></pre>
<p>You can use <code>hub.GetTraceparent()</code> and <code>hub.GetBaggage()</code> to fetch the necessary header values for outgoing HTTP requests.</p>
<pre lang="go"><code>hub := sentry.GetHubFromContext(ctx)
req, _ := http.NewRequest(&quot;GET&quot;, &quot;http://localhost:3000&quot;, nil)
req.Header.Add(sentry.SentryTraceHeader, hub.GetTraceparent())
req.Header.Add(sentry.SentryBaggageHeader, hub.GetBaggage())
</code></pre>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Initialize <code>HTTPTransport.limit</code> if <code>nil</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/844">#844</a>)</li>
<li>Fix <code>sentry.StartTransaction()</code> returning a transaction with an outdated context on existing transactions (<a href="https://redirect.github.com/getsentry/sentry-go/pull/854">#854</a>)</li>
<li>Treat <code>Proxy-Authorization</code> as a sensitive header (<a href="https://redirect.github.com/getsentry/sentry-go/pull/859">#859</a>)</li>
<li>Add support for the <code>http.Hijacker</code> interface to the <code>sentrynegroni</code> package (<a href="https://redirect.github.com/getsentry/sentry-go/pull/871">#871</a>)</li>
<li>Go version &gt;= 1.23: Use value from <code>http.Request.Pattern</code> for HTTP transaction names when using <code>sentryhttp</code> &amp; <code>sentrynegroni</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/875">#875</a>)</li>
<li>Go version &gt;= 1.21: Fix closure functions name grouping (<a href="https://redirect.github.com/getsentry/sentry-go/pull/877">#877</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Collect <code>span</code> origins (<a href="https://redirect.github.com/getsentry/sentry-go/pull/849">#849</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.29.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.29.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Remove the <code>sentrymartini</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/861">#861</a>)</li>
<li>The <code>WrapResponseWriter</code> has been moved from the <code>sentryhttp</code> package to the <code>internal/httputils</code> package. If you've imported it previosuly, you'll need to copy the implementation in your project. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/871">#871</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add new convenience methods to continue a trace and propagate tracing headers for error-only use cases. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/862">#862</a>)</p>
<p>If you are not using one of our integrations, you can manually continue an incoming trace by using <code>sentry.ContinueTrace()</code> by providing the <code>sentry-trace</code> and <code>baggage</code> header received from a downstream SDK.</p>
<pre lang="go"><code>hub := sentry.CurrentHub()
sentry.ContinueTrace(hub, r.Header.Get(sentry.SentryTraceHeader), r.Header.Get(sentry.SentryBaggageHeader)),
</code></pre>
<p>You can use <code>hub.GetTraceparent()</code> and <code>hub.GetBaggage()</code> to fetch the necessary header values for outgoing HTTP requests.</p>
<pre lang="go"><code>hub := sentry.GetHubFromContext(ctx)
req, _ := http.NewRequest(&quot;GET&quot;, &quot;http://localhost:3000&quot;, nil)
req.Header.Add(sentry.SentryTraceHeader, hub.GetTraceparent())
req.Header.Add(sentry.SentryBaggageHeader, hub.GetBaggage())
</code></pre>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Initialize <code>HTTPTransport.limit</code> if <code>nil</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/844">#844</a>)</li>
<li>Fix <code>sentry.StartTransaction()</code> returning a transaction with an outdated context on existing transactions (<a href="https://redirect.github.com/getsentry/sentry-go/pull/854">#854</a>)</li>
<li>Treat <code>Proxy-Authorization</code> as a sensitive header (<a href="https://redirect.github.com/getsentry/sentry-go/pull/859">#859</a>)</li>
<li>Add support for the <code>http.Hijacker</code> interface to the <code>sentrynegroni</code> package (<a href="https://redirect.github.com/getsentry/sentry-go/pull/871">#871</a>)</li>
<li>Go version &gt;= 1.23: Use value from <code>http.Request.Pattern</code> for HTTP transaction names when using <code>sentryhttp</code> &amp; <code>sentrynegroni</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/875">#875</a>)</li>
<li>Go version &gt;= 1.21: Fix closure functions name grouping (<a href="https://redirect.github.com/getsentry/sentry-go/pull/877">#877</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Collect <code>span</code> origins (<a href="https://redirect.github.com/getsentry/sentry-go/pull/849">#849</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/1c2b1d275d5aaeeff18bd2c610c5612572aacc1c"><code>1c2b1d2</code></a> release: 0.29.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/17096d27f90544dbff0454cdde996eb2837cbe38"><code>17096d2</code></a> Prepare 0.29.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/883">#883</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/61d84f83c7bfa1447bb4fe14d94a518e0b4ca40f"><code>61d84f8</code></a> Revert &quot;Add ability to skip frames (<a href="https://redirect.github.com/getsentry/sentry-go/issues/852">#852</a>)&quot; (<a href="https://redirect.github.com/getsentry/sentry-go/issues/882">#882</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/871a3667e7d3339f59c6da5d9e9f221f601a7624"><code>871a366</code></a> Move ResponseWriter wrapper to internal (<a href="https://redirect.github.com/getsentry/sentry-go/issues/880">#880</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/24ffea568912a2e729f80c388be5400c67f698de"><code>24ffea5</code></a> Tracing without performance improvements (<a href="https://redirect.github.com/getsentry/sentry-go/issues/862">#862</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7085a1daa3e2ac8e3deaaa1377a48947118bc12a"><code>7085a1d</code></a> fix(http): use route name from 'r.Pattern' instead (<a href="https://redirect.github.com/getsentry/sentry-go/issues/875">#875</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e5d46d55523ec18a984c580b2c3b258ed86bd172"><code>e5d46d5</code></a> Implement hijacker interface for Negroni integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/871">#871</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bff876d40bcdd92fa2beea917e66c7c25a5af1d1"><code>bff876d</code></a> Function grouping prefix fix (<a href="https://redirect.github.com/getsentry/sentry-go/issues/877">#877</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/72f3bff9fc92731899074717f2e734cd95d440d9"><code>72f3bff</code></a> ci: add go1.23 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/872">#872</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/11d3c795ddc3d0a28ddb26755584ba4a7d314e7e"><code>11d3c79</code></a> build(deps): bump golangci/golangci-lint-action from 6.0.1 to 6.1.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/867">#867</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.28.1...v0.29.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.6.0 to 5.7.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.1 (September 10, 2024)</h1>
<ul>
<li>Fix data race in tracelog.TraceLog</li>
<li>Update puddle to v2.2.2. This removes the import of nanotime via linkname.</li>
<li>Update golang.org/x/crypto and golang.org/x/text</li>
</ul>
<h1>5.7.0 (September 7, 2024)</h1>
<ul>
<li>Add support for sslrootcert=system (Yann Soubeyrand)</li>
<li>Add LoadTypes to load multiple types in a single SQL query (Nick Farrell)</li>
<li>Add XMLCodec supports encoding + scanning XML column type like json (nickcruess-soda)</li>
<li>Add MultiTrace (Stepan Rabotkin)</li>
<li>Add TraceLogConfig with customizable TimeKey (stringintech)</li>
<li>pgx.ErrNoRows wraps sql.ErrNoRows to aid in database/sql compatibility with native pgx functions (merlin)</li>
<li>Support scanning binary formatted uint32 into string / TextScanner (jennifersp)</li>
<li>Fix interval encoding to allow 0s and avoid extra spaces (Carlos Pérez-Aradros Herce)</li>
<li>Update pgservicefile - fixes panic when parsing invalid file</li>
<li>Better error message when reading past end of batch</li>
<li>Don't print url when url.Parse returns an error (Kevin Biju)</li>
<li>Fix snake case name normalization collision in RowToStructByName with db tag (nolandseigler)</li>
<li>Fix: Scan and encode types with underlying types of arrays</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/672c4a3a24849b1f34857817e6ed76f6581bbe90"><code>672c4a3</code></a> Release v5.7.1</li>
<li><a href="https://github.com/jackc/pgx/commit/f8a5a5c9e35993bb3311ed95b6660fab0af5f07e"><code>f8a5a5c</code></a> Update golang.org/x/crypto and golang.org/x/text</li>
<li><a href="https://github.com/jackc/pgx/commit/ab36c2c0dd2034ca4a8e610d3ec9d524240b4802"><code>ab36c2c</code></a> Upgrade puddle to v2.2.2</li>
<li><a href="https://github.com/jackc/pgx/commit/ce66b1dae4d1e9fc68d576c029463562e2324d63"><code>ce66b1d</code></a> Fix data race with TraceLog.Config initialization</li>
<li><a href="https://github.com/jackc/pgx/commit/d1205a6dbccfcf0e53cc3c39741b4a0f60f209ae"><code>d1205a6</code></a> Release v5.7.0</li>
<li><a href="https://github.com/jackc/pgx/commit/97d20ccfadaa55c571f74677899f1b89f75c8e39"><code>97d20cc</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2115">#2115</a> from ninedraft/sql-err-no-rows</li>
<li><a href="https://github.com/jackc/pgx/commit/e9bd382c519ca618090d684a120a7e006900f777"><code>e9bd382</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2114">#2114</a> from jennifersp/master</li>
<li><a href="https://github.com/jackc/pgx/commit/603f2337d66b2fe44263e4fe11823d452ed87c0d"><code>603f233</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2113">#2113</a> from mateuszkowalke/master</li>
<li><a href="https://github.com/jackc/pgx/commit/035bbbe0cbd53928c4a67f79576dc29fb10021fb"><code>035bbbe</code></a> Use sql.ErrNoRows as value for pgx.ErrNoRows</li>
<li><a href="https://github.com/jackc/pgx/commit/73bbced270612f0bc25ad2843534d60ff5933ee1"><code>73bbced</code></a> add byte length check to uint32</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.6.0...v5.7.1">compare view</a></li>
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

### Comment by @app-sre-bot on September 16, 2024 at 04:12 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on September 16, 2024 at 02:23 PM UTC

[Test]

### Comment by @swadeley on September 16, 2024 at 02:25 PM UTC

[test]

### Comment by @jlsherrill on September 16, 2024 at 05:57 PM UTC

/retest

### Comment by @swadeley on September 16, 2024 at 07:14 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on September 17, 2024 at 11:38 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/816*
