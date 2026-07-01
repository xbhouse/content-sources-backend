---
type: pull_request
number: 784
title: "Build: Bump the go group across 1 directory with 7 updates"
state: merged
author: dependabot
created: 2024-08-26T04:17:44Z
updated: 2024-08-29T12:41:21Z
closed: 2024-08-29T12:41:13Z
merged: 2024-08-29T12:41:13Z
base_branch: main
head_branch: dependabot/go_modules/go-04fdd67acc
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/784
---

# Pull Request #784: Build: Bump the go group across 1 directory with 7 updates

**Author**: @dependabot
**Created**: August 26, 2024 at 04:17 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-04fdd67acc`

## Description

Bumps the go group with 7 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.19.1` | `1.20.2` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.43.2` | `1.43.3` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.30.3` | `1.30.4` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.27` | `1.17.29` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.37.3` | `1.37.5` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.8.1723121474` | `2024.8.1723836162` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.6` | `1.25.9` |


Updates `github.com/prometheus/client_golang` from 1.19.1 to 1.20.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.20.2</h2>
<ul>
<li>[BUGFIX] promhttp: Unset Content-Encoding header when data is uncompressed. <a href="https://redirect.github.com/prometheus/client_golang/issues/1596">#1596</a></li>
</ul>
<h2>v1.20.1</h2>
<p>This release contains the critical fix for the <a href="https://redirect.github.com/prometheus/client_golang/issues/1584">issue</a>. Thanks to <a href="https://github.com/geberl"><code>@​geberl</code></a>, <a href="https://github.com/CubicrootXYZ"><code>@​CubicrootXYZ</code></a>, <a href="https://github.com/zetaab"><code>@​zetaab</code></a> and <a href="https://github.com/timofurrer"><code>@​timofurrer</code></a> for helping us with the investigation!</p>
<ul>
<li>[BUGFIX] process-collector: Fixed unregistered descriptor error when using process collector with PedanticRegistry on Linux machines. <a href="https://redirect.github.com/prometheus/client_golang/issues/1587">#1587</a></li>
</ul>
<h2>v1.20.0</h2>
<p>Thanks everyone for contributions!</p>
<p>:warning: In this release we remove one (broken anyway, given Go runtime changes) metric and add three new (representing GOGC, GOMEMLIMIT and GOMAXPROCS flags) to the default <code>collectors.NewGoCollector()</code> collector. Given its popular usage, expect your binary to expose two additional metric.</p>
<h2>Changes</h2>
<ul>
<li>[CHANGE] :warning: go-collector: Remove <code>go_memstat_lookups_total</code> metric which was always 0; Go runtime stopped sharing pointer lookup statistics. <a href="https://redirect.github.com/prometheus/client_golang/issues/1577">#1577</a></li>
<li>[FEATURE] :warning: go-collector: Add 3 default metrics: <code>go_gc_gogc_percent</code>, <code>go_gc_gomemlimit_bytes</code> and <code>go_sched_gomaxprocs_threads</code> as those are recommended by the Go team. <a href="https://redirect.github.com/prometheus/client_golang/issues/1559">#1559</a></li>
<li>[FEATURE] go-collector: Add more information to all metrics' HELP e.g. the exact <code>runtime/metrics</code> sourcing each metric (if relevant). <a href="https://redirect.github.com/prometheus/client_golang/issues/1568">#1568</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1578">#1578</a></li>
<li>[FEATURE] testutil: Add CollectAndFormat method. <a href="https://redirect.github.com/prometheus/client_golang/issues/1503">#1503</a></li>
<li>[FEATURE] histograms: Add support for exemplars in native histograms. <a href="https://redirect.github.com/prometheus/client_golang/issues/1471">#1471</a></li>
<li>[FEATURE] promhttp: Add experimental support for <code>zstd</code> on scrape, controlled by the request <code>Accept-Encoding</code> header. <a href="https://redirect.github.com/prometheus/client_golang/issues/1496">#1496</a></li>
<li>[FEATURE] api/v1: Add <code>WithLimit</code> parameter to all API methods that supports it. <a href="https://redirect.github.com/prometheus/client_golang/issues/1544">#1544</a></li>
<li>[FEATURE] prometheus: Add support for created timestamps in constant histograms and constant summaries. <a href="https://redirect.github.com/prometheus/client_golang/issues/1537">#1537</a></li>
<li>[FEATURE] process-collectors: Add network usage metrics: <code>process_network_receive_bytes_total</code> and <code>process_network_transmit_bytes_total</code>. <a href="https://redirect.github.com/prometheus/client_golang/issues/1555">#1555</a></li>
<li>[FEATURE] promlint: Add duplicated metric lint rule. <a href="https://redirect.github.com/prometheus/client_golang/issues/1472">#1472</a></li>
<li>[BUGFIX] promlint: Relax metric type in name linter rule. <a href="https://redirect.github.com/prometheus/client_golang/issues/1455">#1455</a></li>
<li>[BUGFIX] promhttp: Make sure server
instrumentation wrapping supports new and future extra responseWriter methods. <a href="https://redirect.github.com/prometheus/client_golang/issues/1480">#1480</a></li>
<li>[BUGFIX] testutil: Functions using compareMetricFamilies are now failing if filtered metricNames are not in the input. <a href="https://redirect.github.com/prometheus/client_golang/issues/1424">#1424</a></li>
</ul>
<!-- raw HTML omitted -->
<ul>
<li>feat(prometheus/testutil/promlint/validations): refine lintMetricType… by <a href="https://github.com/foehammer127"><code>@​foehammer127</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1455">prometheus/client_golang#1455</a></li>
<li>Bump github.com/prometheus/client_golang from 1.18.0 to 1.19.0 in /examples/middleware by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1457">prometheus/client_golang#1457</a></li>
<li>Bump github.com/prometheus/client_model from 0.5.0 to 0.6.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1458">prometheus/client_golang#1458</a></li>
<li>Bump golang.org/x/sys from 0.16.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1459">prometheus/client_golang#1459</a></li>
<li>Bump github.com/prometheus/client_golang from 1.18.0 to 1.19.0 in /tutorial/whatsup by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1461">prometheus/client_golang#1461</a></li>
<li>Merge Release 1.19 back to main by <a href="https://github.com/ArthurSens"><code>@​ArthurSens</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1462">prometheus/client_golang#1462</a></li>
<li>Bump the github-actions group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1456">prometheus/client_golang#1456</a></li>
<li>Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1466">prometheus/client_golang#1466</a></li>
<li>Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 in /examples/middleware by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1467">prometheus/client_golang#1467</a></li>
<li>Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 in /tutorial/whatsup by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1469">prometheus/client_golang#1469</a></li>
<li>Add LintDuplicateMetric to promlint by <a href="https://github.com/bboreham"><code>@​bboreham</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1472">prometheus/client_golang#1472</a></li>
<li>Auto-update Go Collector Metrics for new Go versions by <a href="https://github.com/SachinSahu431"><code>@​SachinSahu431</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1476">prometheus/client_golang#1476</a></li>
<li>Implement Unwrap() for responseWriterDelegator by <a href="https://github.com/igor-drozdov"><code>@​igor-drozdov</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1480">prometheus/client_golang#1480</a></li>
<li>Bump golang.org/x/sys from 0.17.0 to 0.18.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1485">prometheus/client_golang#1485</a></li>
<li>Bump github.com/prometheus/procfs from 0.12.0 to 0.13.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1486">prometheus/client_golang#1486</a></li>
<li>ci: Remove hardcoded supported Go versions from go.yml by <a href="https://github.com/SachinSahu431"><code>@​SachinSahu431</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1489">prometheus/client_golang#1489</a></li>
<li>feat: metrics generation workflow by <a href="https://github.com/SachinSahu431"><code>@​SachinSahu431</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1481">prometheus/client_golang#1481</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.20.2 / 2024-08-23</h2>
<ul>
<li>[BUGFIX] promhttp: Unset Content-Encoding header when data is uncompressed. <a href="https://redirect.github.com/prometheus/client_golang/issues/1596">#1596</a></li>
</ul>
<h2>1.20.1 / 2024-08-20</h2>
<ul>
<li>[BUGFIX] process-collector: Fixed unregistered descriptor error when using process collector with <code>PedanticRegistry</code> on linux machines. <a href="https://redirect.github.com/prometheus/client_golang/issues/1587">#1587</a></li>
</ul>
<h2>1.20.0 / 2024-08-14</h2>
<ul>
<li>[CHANGE] :warning: go-collector: Remove <code>go_memstat_lookups_total</code> metric which was always 0; Go runtime stopped sharing pointer lookup statistics. <a href="https://redirect.github.com/prometheus/client_golang/issues/1577">#1577</a></li>
<li>[FEATURE] :warning: go-collector: Add 3 default metrics: <code>go_gc_gogc_percent</code>, <code>go_gc_gomemlimit_bytes</code> and <code>go_sched_gomaxprocs_threads</code> as those are recommended by the Go team. <a href="https://redirect.github.com/prometheus/client_golang/issues/1559">#1559</a></li>
<li>[FEATURE] go-collector: Add more information to all metrics' HELP e.g. the exact <code>runtime/metrics</code> sourcing each metric (if relevant). <a href="https://redirect.github.com/prometheus/client_golang/issues/1568">#1568</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1578">#1578</a></li>
<li>[FEATURE] testutil: Add CollectAndFormat method. <a href="https://redirect.github.com/prometheus/client_golang/issues/1503">#1503</a></li>
<li>[FEATURE] histograms: Add support for exemplars in native histograms. <a href="https://redirect.github.com/prometheus/client_golang/issues/1471">#1471</a></li>
<li>[FEATURE] promhttp: Add experimental support for <code>zstd</code> on scrape, controlled by the request <code>Accept-Encoding</code> header. <a href="https://redirect.github.com/prometheus/client_golang/issues/1496">#1496</a></li>
<li>[FEATURE] api/v1: Add <code>WithLimit</code> parameter to all API methods that supports it. <a href="https://redirect.github.com/prometheus/client_golang/issues/1544">#1544</a></li>
<li>[FEATURE] prometheus: Add support for created timestamps in constant histograms and constant summaries. <a href="https://redirect.github.com/prometheus/client_golang/issues/1537">#1537</a></li>
<li>[FEATURE] process-collector: Add network usage metrics: <code>process_network_receive_bytes_total</code> and <code>process_network_transmit_bytes_total</code>. <a href="https://redirect.github.com/prometheus/client_golang/issues/1555">#1555</a></li>
<li>[FEATURE] promlint: Add duplicated metric lint rule. <a href="https://redirect.github.com/prometheus/client_golang/issues/1472">#1472</a></li>
<li>[BUGFIX] promlint: Relax metric type in name linter rule. <a href="https://redirect.github.com/prometheus/client_golang/issues/1455">#1455</a></li>
<li>[BUGFIX] promhttp: Make sure server instrumentation wrapping supports new and future extra responseWriter methods. <a href="https://redirect.github.com/prometheus/client_golang/issues/1480">#1480</a></li>
<li>[BUGFIX] testutil: Functions using compareMetricFamilies are now failing if filtered metricNames are not in the input. <a href="https://redirect.github.com/prometheus/client_golang/issues/1424">#1424</a></li>
</ul>
<h2>1.19.0 / 2024-02-27</h2>
<p>The module <code>prometheus/common v0.48.0</code> introduced an incompatibility when used together with client_golang (See <a href="https://redirect.github.com/prometheus/client_golang/pull/1448">prometheus/client_golang#1448</a> for more details). If your project uses client_golang and you want to use <code>prometheus/common v0.48.0</code> or higher, please update client_golang to v1.19.0.</p>
<ul>
<li>[CHANGE] Minimum required go version is now 1.20 (we also test client_golang against new 1.22 version). <a href="https://redirect.github.com/prometheus/client_golang/issues/1445">#1445</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1449">#1449</a></li>
<li>[FEATURE] collectors: Add version collector. <a href="https://redirect.github.com/prometheus/client_golang/issues/1422">#1422</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1427">#1427</a></li>
</ul>
<h2>1.18.0 / 2023-12-22</h2>
<ul>
<li>[FEATURE] promlint: Allow creation of custom metric validations. <a href="https://redirect.github.com/prometheus/client_golang/issues/1311">#1311</a></li>
<li>[FEATURE] Go programs using client_golang can be built in wasip1 OS. <a href="https://redirect.github.com/prometheus/client_golang/issues/1350">#1350</a></li>
<li>[BUGFIX] histograms: Add timer to reset ASAP after bucket limiting has happened. <a href="https://redirect.github.com/prometheus/client_golang/issues/1367">#1367</a></li>
<li>[BUGFIX] testutil: Fix comparison of metrics with empty Help strings. <a href="https://redirect.github.com/prometheus/client_golang/issues/1378">#1378</a></li>
<li>[ENHANCEMENT] Improved performance of <code>MetricVec.WithLabelValues(...)</code>. <a href="https://redirect.github.com/prometheus/client_golang/issues/1360">#1360</a></li>
</ul>
<h2>1.17.0 / 2023-09-27</h2>
<ul>
<li>[CHANGE] Minimum required go version is now 1.19 (we also test client_golang against new 1.21 version). <a href="https://redirect.github.com/prometheus/client_golang/issues/1325">#1325</a></li>
<li>[FEATURE] Add support for Created Timestamps in Counters, Summaries and Historams. <a href="https://redirect.github.com/prometheus/client_golang/issues/1313">#1313</a></li>
<li>[ENHANCEMENT] Enable detection of a native histogram without observations. <a href="https://redirect.github.com/prometheus/client_golang/issues/1314">#1314</a></li>
</ul>
<h2>1.16.0 / 2023-06-15</h2>
<ul>
<li>[BUGFIX] api: Switch to POST for LabelNames, Series, and QueryExemplars. <a href="https://redirect.github.com/prometheus/client_golang/issues/1252">#1252</a></li>
<li>[BUGFIX] api: Fix undefined execution order in return statements. <a href="https://redirect.github.com/prometheus/client_golang/issues/1260">#1260</a></li>
<li>[BUGFIX] native histograms: Fix bug in bucket key calculation. <a href="https://redirect.github.com/prometheus/client_golang/issues/1279">#1279</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/67121dc55a7c1ac353141e360a2b184c5e3d5529"><code>67121dc</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1596">#1596</a> from mrueg/fix-uncompressed-content-header</li>
<li><a href="https://github.com/prometheus/client_golang/commit/187acd4ca1a05e3207990e427cae7057a1e58961"><code>187acd4</code></a> Cut 1.20.2</li>
<li><a href="https://github.com/prometheus/client_golang/commit/f7f8f3a1e2c09da4b88b78b7e04a9252187cf71f"><code>f7f8f3a</code></a> fix: Unset Content-Encoding header when uncompressed</li>
<li><a href="https://github.com/prometheus/client_golang/commit/2254d6c3087d3660f73785ae123c7cc705f2a9c2"><code>2254d6c</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1587">#1587</a> from prometheus/fix-processcollector</li>
<li><a href="https://github.com/prometheus/client_golang/commit/4a15d0584653a36547195042512218cad99888c6"><code>4a15d05</code></a> Cut 1.20.1</li>
<li><a href="https://github.com/prometheus/client_golang/commit/f2dd7b35fd3f28b6a59b64953aab727b5c0c0404"><code>f2dd7b3</code></a> Use pedantic registry in other places too, to double check.</li>
<li><a href="https://github.com/prometheus/client_golang/commit/261fe84cd4f670b8e51bf74f6f6bbb96aa1fce20"><code>261fe84</code></a> bugfix: Pass network metrics to processCollector's Describe() function</li>
<li><a href="https://github.com/prometheus/client_golang/commit/5bf3341b6684a18bfdc019c9b2c154ab6a42c0a1"><code>5bf3341</code></a> Use NewPedanticRegistry in Process' Collector tests</li>
<li><a href="https://github.com/prometheus/client_golang/commit/73b811c54a628c7a7fe43005fb81351f64da9289"><code>73b811c</code></a> Cut 1.20.0 release. (<a href="https://redirect.github.com/prometheus/client_golang/issues/1580">#1580</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/7ce508988aa3bfc91d56afa9f82b7167b36c4590"><code>7ce5089</code></a> gocollector: Attach original runtime/metrics metric name to help. (<a href="https://redirect.github.com/prometheus/client_golang/issues/1578">#1578</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prometheus/client_golang/compare/v1.19.1...v1.20.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.43.2 to 1.43.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.43.3 (2024-08-12)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: declare assignor variable for examples &amp; clean up log format by <a href="https://github.com/kumakichi"><code>@​kumakichi</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2909">IBM/sarama#2909</a></li>
<li>fix(consumer): maintain ordering of offset commit requests by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2947">IBM/sarama#2947</a></li>
<li>fix(producer): treat ErrKafkaStorageError as retriable by <a href="https://github.com/richardartoul"><code>@​richardartoul</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2939">IBM/sarama#2939</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump the golang-org-x group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2956">IBM/sarama#2956</a></li>
<li>chore(deps): bump github.com/eapache/go-resiliency from 1.6.0 to 1.7.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2944">IBM/sarama#2944</a></li>
<li>chore(deps): bump github.com/klauspost/compress from 1.17.8 to 1.17.9 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2926">IBM/sarama#2926</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>fix(ci): correct docker-compose install by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2954">IBM/sarama#2954</a></li>
</ul>
<h3>:memo: Documentation</h3>
<ul>
<li>fix(doc): correct JVM's config name corresponding to MaxWaitTime by <a href="https://github.com/abhipranay"><code>@​abhipranay</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2893">IBM/sarama#2893</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/abhipranay"><code>@​abhipranay</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2893">IBM/sarama#2893</a></li>
<li><a href="https://github.com/kumakichi"><code>@​kumakichi</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2909">IBM/sarama#2909</a></li>
<li><a href="https://github.com/richardartoul"><code>@​richardartoul</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2939">IBM/sarama#2939</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.43.2...v1.43.3">https://github.com/IBM/sarama/compare/v1.43.2...v1.43.3</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/a46c78d8078332141d3c7e9a6c6cc22e14af0f74"><code>a46c78d</code></a> chore(ci): bump ubi8/ubi-minimal from <code>f729a7f</code> to <code>de2a0a2</code> (<a href="https://redirect.github.com/IBM/sarama/issues/2958">#2958</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3aea989f608e0e59178376e874c27ba448cdd9ab"><code>3aea989</code></a> fix(producer): treat ErrKafkaStorageError as retriable (<a href="https://redirect.github.com/IBM/sarama/issues/2939">#2939</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/9ded62996a0d7407dffb1049949e5fe2a489eb79"><code>9ded629</code></a> --- (<a href="https://redirect.github.com/IBM/sarama/issues/2908">#2908</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e40dc2c2e2aa399c0228b687af5f95bef733c3bf"><code>e40dc2c</code></a> chore(ci): bump golangci/golangci-lint-action from 5.0.0 to 6.0.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2896">#2896</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/845ec7c8a8d530248a394df09a2f31803fa3a2c2"><code>845ec7c</code></a> chore(deps): bump github.com/klauspost/compress from 1.17.8 to 1.17.9 (<a href="https://redirect.github.com/IBM/sarama/issues/2926">#2926</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/1466fb161ad6511c22ca8d7f093462ab4c1ade8d"><code>1466fb1</code></a> chore(deps): bump github.com/eapache/go-resiliency from 1.6.0 to 1.7.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2944">#2944</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/73090d640c5e7ad45c334d9e4b248668cb627756"><code>73090d6</code></a> chore(ci): bump github/codeql-action from 3.25.2 to 3.26.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2955">#2955</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/97a9f1d632add9a9b10c03382834b986c0b28c48"><code>97a9f1d</code></a> chore(deps): bump the golang-org-x group across 1 directory with 2 updates (#...</li>
<li><a href="https://github.com/IBM/sarama/commit/c82965870bbeceaa086277abc4ac36a5f4763de4"><code>c829658</code></a> chore(ci): remove <code>version</code> field from compose yml</li>
<li><a href="https://github.com/IBM/sarama/commit/08ded321b5bca081179d27057d294d5ff76087e4"><code>08ded32</code></a> fix(ci): correct docker-compose install</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.43.2...v1.43.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.30.3 to 1.30.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1d44177adda2c58035abe5f2e878c5c5a60db913"><code>1d44177</code></a> Release 2024-08-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bab4916aa8605e843377c9c049b2e758d4b010f8"><code>bab4916</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d4ecf5c0327da6f66fcbb8ec4e181ed8dd35b57"><code>2d4ecf5</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/df4f4e1401905d647c86951dd196ee3f8b235f41"><code>df4f4e1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51ca5b5729bc6a8599fe5b8d9aa193df22290db3"><code>51ca5b5</code></a> bump smithy-go and go version (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2750">#2750</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/39ddd61ec78f6972734348d9fa30a0cfdbf5b486"><code>39ddd61</code></a> Release 2024-08-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ffaf282ebc42c4c01e4d9972a64d275f0cae8cbe"><code>ffaf282</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6beafa89ca37adbf20feafe2d726d83b80cca773"><code>6beafa8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/57f6f2dde26a255c92624621daf6a5262e05f676"><code>57f6f2d</code></a> Release 2024-08-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cdeb8a6a1be147976e46b8dfe8c5699578893d69"><code>cdeb8a6</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.3...v1.30.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.27 to 1.17.29
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1d210dd11afc2ef0efa2e980d6b6c09cac0f5cb"><code>d1d210d</code></a> Release 2024-08-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/afb0408b5f2aff8e2b5e40e8f8abcfb82e989198"><code>afb0408</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/05f44921828cb1d34457ad682ece64a13cb6be85"><code>05f4492</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/600651a884fd9902b62f34bec4fbaab05438d23f"><code>600651a</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ebaae4054cb459d4988401e2f514d4e1515bef92"><code>ebaae40</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/38fe812f7abf976b6b1a5122c0c3bf303502cfc5"><code>38fe812</code></a> Release 2024-08-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b5042114218829dd202359ca689f1cd56357852f"><code>b504211</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/02f304545f970b76f77b46543d71cb9bfcadaa3b"><code>02f3045</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f7d2cf6ed8b48b201c35341d0b6f58cf24caa983"><code>f7d2cf6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4d361264c636ebc3906105bf272b9c0a72673d10"><code>4d36126</code></a> Release 2024-08-20</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.27...credentials/v1.17.29">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.37.3 to 1.37.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4599f78694cabb6853addabc6f92cb197fdb5647"><code>4599f78</code></a> Release 2023-08-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2a439ce0fdda24816a5ea71c083765af67e93599"><code>2a439ce</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/46f31d76b34aefbc66bbfb08d2ef7e22267d4b66"><code>46f31d7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/525486a4cfa9b8364ea5346a526dd1f92d3edc76"><code>525486a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4e3176642b69937575f61f752f7d094fb6f4084"><code>b4e3176</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2189">#2189</a> from aws/feat-presign-polly</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/30fe9b8de4c8538156248ae21b577a3e49170750"><code>30fe9b8</code></a> Modify and Merge polly mod import</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f6c1d4839450e7f99850868570ff0c47850b72ca"><code>f6c1d48</code></a> sync polly presigner from main</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/396f97a998c5fab5bc7ee7bd4297c6d0f42b661e"><code>396f97a</code></a> Release 2023-08-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/08ab45438df135fc66f2eaf80032694ff2486db3"><code>08ab454</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b50999bacbf780bf55d5c2b051c96e5c4467a5f6"><code>b50999b</code></a> Update SDK's smithy-go dependency to v1.14.2</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.37.3...service/ssm/v1.37.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.8.1723121474 to 2024.8.1723836162
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/79557a9f2cd6aa7bb211a2f105b43f3d8510b5ca"><code>79557a9</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a847455298968137d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/5dc380794202772fe28d85282f9cc53e3bd83bd6"><code>5dc3807</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d44946523157952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/7015c8456d69ac59f80474d5bc47f2665e0bc905"><code>7015c84</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b979ab9b3a38c37ae5486b92ad2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.8.1723121474...release/v2024.8.1723836162">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.6 to 1.25.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/c7356847402b656bd7e0387f5e7ac3f382fc158d"><code>c735684</code></a> [CCXDEV-13127] Revert sarama migration (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/573">#573</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/fa0ce65809cba3e65a283f13fa1baf7691a0b11f"><code>fa0ce65</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/572">#572</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/68940852ce7b6f7d793fc3b18c69183d29e61b2b"><code>6894085</code></a> Bump github.com/IBM/sarama from 1.43.2 to 1.43.3</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/8601fc95f372fd70dccf67e5e88839ae0dc748db"><code>8601fc9</code></a> replace github.com/Shopify/sarama with github.com/IBM/sarama (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/571">#571</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/16c41f8fe965b097ed7ade03b22b04f038ec4f63"><code>16c41f8</code></a> reduce error level messages (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/570">#570</a>)</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.6...v1.25.9">compare view</a></li>
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

### Comment by @app-sre-bot on August 26, 2024 at 04:18 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on August 26, 2024 at 01:17 PM UTC

[test]

### Comment by @jlsherrill on August 26, 2024 at 02:12 PM UTC

[test]

### Comment by @swadeley on August 27, 2024 at 01:57 PM UTC

/retest

### Comment by @swadeley on August 29, 2024 at 11:24 AM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on August 29, 2024 at 12:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/784*
