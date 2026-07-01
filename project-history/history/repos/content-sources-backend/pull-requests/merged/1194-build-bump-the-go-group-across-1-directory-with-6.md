---
type: pull_request
number: 1194
title: "Build: Bump the go group across 1 directory with 6 updates"
state: merged
author: dependabot
created: 2025-09-08T17:37:56Z
updated: 2025-09-08T18:49:18Z
closed: 2025-09-08T18:49:12Z
merged: 2025-09-08T18:49:11Z
base_branch: main
head_branch: dependabot/go_modules/go-bc8510c53d
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1194
---

# Pull Request #1194: Build: Bump the go group across 1 directory with 6 updates

**Author**: @dependabot
**Created**: September 08, 2025 at 05:37 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-bc8510c53d`

## Description

Bumps the go group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/spf13/viper](https://github.com/spf13/viper) | `1.20.1` | `1.21.0` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.30.2` | `1.30.5` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.23.0` | `1.23.2` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.7.5` | `5.7.6` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.12.1` | `9.13.0` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.74.2` | `1.75.0` |


Updates `github.com/spf13/viper` from 1.20.1 to 1.21.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/spf13/viper/releases">github.com/spf13/viper's releases</a>.</em></p>
<blockquote>
<h2>v1.21.0</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>Enhancements 🚀</h3>
<ul>
<li>Add support for flags pflag.BoolSlice, pflag.UintSlice and pflag.Float64Slice by <a href="https://github.com/nmvalera"><code>@​nmvalera</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2015">spf13/viper#2015</a></li>
<li>feat: use maintained yaml library by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2040">spf13/viper#2040</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>fix(config): get config type from v.configType or config file ext by <a href="https://github.com/GuillaumeBAECHLER"><code>@​GuillaumeBAECHLER</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2003">spf13/viper#2003</a></li>
<li>fix: config type check when loading any config by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2007">spf13/viper#2007</a></li>
</ul>
<h3>Dependency Updates ⬆️</h3>
<ul>
<li>Update dependencies by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1993">spf13/viper#1993</a></li>
<li>build(deps): bump github.com/spf13/cast from 1.7.1 to 1.8.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2017">spf13/viper#2017</a></li>
<li>build(deps): bump github.com/pelletier/go-toml/v2 from 2.2.3 to 2.2.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2013">spf13/viper#2013</a></li>
<li>build(deps): bump github.com/sagikazarmark/locafero from 0.8.0 to 0.9.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2008">spf13/viper#2008</a></li>
<li>build(deps): bump golang.org/x/net from 0.37.0 to 0.38.0 in /remote by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2016">spf13/viper#2016</a></li>
<li>build(deps): bump github.com/spf13/cast from 1.8.0 to 1.9.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2020">spf13/viper#2020</a></li>
<li>build(deps): bump github.com/go-viper/mapstructure/v2 from 2.2.1 to 2.3.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2028">spf13/viper#2028</a></li>
<li>build(deps): bump github.com/go-viper/mapstructure/v2 from 2.3.0 to 2.4.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2035">spf13/viper#2035</a></li>
<li>build(deps): bump github.com/spf13/pflag from 1.0.6 to 1.0.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2036">spf13/viper#2036</a></li>
<li>build(deps): bump github.com/fsnotify/fsnotify from 1.8.0 to 1.9.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2012">spf13/viper#2012</a></li>
<li>build(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2052">spf13/viper#2052</a></li>
<li>build(deps): bump github.com/go-viper/mapstructure/v2 from 2.3.0 to 2.4.0 in /remote by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2048">spf13/viper#2048</a></li>
<li>build(deps): bump github.com/spf13/pflag from 1.0.7 to 1.0.10 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/spf13/viper/pull/2056">spf13/viper#2056</a></li>
<li>chore: update dependencies by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2057">spf13/viper#2057</a></li>
</ul>
<h3>Other Changes</h3>
<ul>
<li>Update update guide with <code>mapstructure</code> package replacement. by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2004">spf13/viper#2004</a></li>
<li>refactor: use the built-in max/min to simplify the code by <a href="https://github.com/yingshanghuangqiao"><code>@​yingshanghuangqiao</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2029">spf13/viper#2029</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/GuillaumeBAECHLER"><code>@​GuillaumeBAECHLER</code></a> made their first contribution in <a href="https://redirect.github.com/spf13/viper/pull/2003">spf13/viper#2003</a></li>
<li><a href="https://github.com/aldas"><code>@​aldas</code></a> made their first contribution in <a href="https://redirect.github.com/spf13/viper/pull/2004">spf13/viper#2004</a></li>
<li><a href="https://github.com/nmvalera"><code>@​nmvalera</code></a> made their first contribution in <a href="https://redirect.github.com/spf13/viper/pull/2015">spf13/viper#2015</a></li>
<li><a href="https://github.com/yingshanghuangqiao"><code>@​yingshanghuangqiao</code></a> made their first contribution in <a href="https://redirect.github.com/spf13/viper/pull/2029">spf13/viper#2029</a></li>
<li><a href="https://github.com/ccoVeille"><code>@​ccoVeille</code></a> made their first contribution in <a href="https://redirect.github.com/spf13/viper/pull/2046">spf13/viper#2046</a></li>
<li><a href="https://github.com/spacez320"><code>@​spacez320</code></a> made their first contribution in <a href="https://redirect.github.com/spf13/viper/pull/2050">spf13/viper#2050</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/spf13/viper/compare/v1.20.0...v1.21.0">https://github.com/spf13/viper/compare/v1.20.0...v1.21.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/spf13/viper/commit/394040caccbdf5821fa6839386a35f0fb1b1ee9e"><code>394040c</code></a> ci: build on go 1.25</li>
<li><a href="https://github.com/spf13/viper/commit/812f548100ce43c96cf2ef6914fa3aea3c9885e2"><code>812f548</code></a> chore: update dependencies</li>
<li><a href="https://github.com/spf13/viper/commit/d5271efd81369475d8b49e85ce35cfe83359a991"><code>d5271ef</code></a> ci: update stale workflow</li>
<li><a href="https://github.com/spf13/viper/commit/dff303b19f6859ea69d0f48a3343765c25ce5ad2"><code>dff303b</code></a> feat: add a stale issue scheduled action</li>
<li><a href="https://github.com/spf13/viper/commit/12879766ad2fccd4875e8dc5ba96ce5f28357842"><code>1287976</code></a> build(deps): bump github.com/spf13/pflag from 1.0.7 to 1.0.10</li>
<li><a href="https://github.com/spf13/viper/commit/38932cd79521272eb1634d81eb1c200fa1803008"><code>38932cd</code></a> build(deps): bump github.com/go-viper/mapstructure/v2 in /remote</li>
<li><a href="https://github.com/spf13/viper/commit/6d014bec7784acab13d7d7eb2adcb6814bac2f89"><code>6d014be</code></a> build(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1</li>
<li><a href="https://github.com/spf13/viper/commit/b74c7ee1e5e999f16bc2b904608815a59241b316"><code>b74c7ee</code></a> build(deps): bump github.com/fsnotify/fsnotify from 1.8.0 to 1.9.0</li>
<li><a href="https://github.com/spf13/viper/commit/acd05e1543b0c36b36580b6bd29c53a9b62acd98"><code>acd05e1</code></a> fix: linting issues</li>
<li><a href="https://github.com/spf13/viper/commit/ae5a8e23e22a67a6dcc2025cef0b5a2ca7615456"><code>ae5a8e2</code></a> ci: upgrade golangci-lint</li>
<li>Additional commits viewable in <a href="https://github.com/spf13/viper/compare/v1.20.1...v1.21.0">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.30.2 to 1.30.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.30.5</h2>
<h2>Changes</h2>
<ul>
<li>No changes</li>
</ul>
<h2>Release v1.30.4</h2>
<h2>Changes</h2>
<ul>
<li>Add Set-based Create and Update support to Generics API <a href="https://github.com/jinzhu"><code>@​jinzhu</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7578">#7578</a>)</li>
<li>fix: build failure on Go versions below 1.21 <a href="https://github.com/iTanken"><code>@​iTanken</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7572">#7572</a>)</li>
<li>fix slogLogger to support ParameterizedQueries Config <a href="https://github.com/EricWvi"><code>@​EricWvi</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7574">#7574</a>)</li>
</ul>
<h2>Release v1.30.3</h2>
<h2>Changes</h2>
<ul>
<li>No changes</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/688e8ea00a232bd661c08d3d3ba22750c3b3d95e"><code>688e8ea</code></a> Set accepts Assigner for Generics API</li>
<li><a href="https://github.com/go-gorm/gorm/commit/1901911908ffe2ecafb4d7435f3859f3fcad2e73"><code>1901911</code></a> Add Set-based Create and Update support to Generics API (<a href="https://redirect.github.com/go-gorm/gorm/issues/7578">#7578</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/cb6574367217a4d82c7c855d5372d06618a5de19"><code>cb65743</code></a> fix: build failure on Go versions below 1.21, add build constraint for slog.g...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/4087ac7d9d191a2ce3f83451cbb82c19510ceff9"><code>4087ac7</code></a> fix slogLogger to support ParameterizedQueries Config (<a href="https://redirect.github.com/go-gorm/gorm/issues/7574">#7574</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/38404259fe7935212fc14156989e3de381a2d3f7"><code>3840425</code></a> fix(generics): resolve CurrentTable in Raw/Exec</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.30.2...v1.30.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.23.0 to 1.23.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.23.2 - 2025-09-05</h2>
<p>This release is made to upgrade to prometheus/common v0.66.1, which drops the dependencies github.com/grafana/regexp and go.uber.org/atomic and replaces gopkg.in/yaml.v2 with go.yaml.in/yaml/v2 (a drop-in replacement). There are no functional changes.</p>
<!-- raw HTML omitted -->
<ul>
<li>[release-1.23] Upgrade to prometheus/common@v0.66.1 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1869">prometheus/client_golang#1869</a></li>
<li>[release-1.23] Cut v1.23.2 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1870">prometheus/client_golang#1870</a></li>
</ul>
<!-- raw HTML omitted -->
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.23.1...v1.23.2">https://github.com/prometheus/client_golang/compare/v1.23.1...v1.23.2</a></p>
<h2>v1.23.1 - 2025-09-04</h2>
<p>This release is made to be compatible with a backwards incompatible API change in prometheus/common v0.66.0. There are no functional changes.</p>
<!-- raw HTML omitted -->
<ul>
<li>[release-1.23] Upgrade to prometheus/common v0.66 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1866">prometheus/client_golang#1866</a></li>
<li>[release-1.23] Cut v1.23.1 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1867">prometheus/client_golang#1867</a></li>
</ul>
<!-- raw HTML omitted -->
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.1">https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.23.2 / 2025-09-05</h2>
<p>This release is made to upgrade to prometheus/common v0.66.1, which drops the dependencies github.com/grafana/regexp and go.uber.org/atomic and replaces gopkg.in/yaml.v2 with go.yaml.in/yaml/v2 (a drop-in replacement).
There are no functional changes.</p>
<h2>1.23.1 / 2025-09-04</h2>
<p>This release is made to be compatible with a backwards incompatible API change
in prometheus/common v0.66.0. There are no functional changes.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/8179a560819f2c64ef6ade70e6ae4c73aecaca3c"><code>8179a56</code></a> Cut v1.23.2 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1870">#1870</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/4142b5908bb6c8f5e412b72de5ea4b927d8c219d"><code>4142b59</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1869">#1869</a> from prometheus/arve/upgrade-common</li>
<li><a href="https://github.com/prometheus/client_golang/commit/4ff40f0d918efc0f59701d13622913805c2425b4"><code>4ff40f0</code></a> Cut v1.23.1 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1867">#1867</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/989b0298944e64f88a54ac9c70cd8c8121f10bc9"><code>989b029</code></a> Upgrade to prometheus/common v0.66 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1866">#1866</a>)</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.5 to 5.7.6
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.6 (September 8, 2025)</h1>
<ul>
<li>Use ParseConfigError in pgx.ParseConfig and pgxpool.ParseConfig (Yurasov Ilia)</li>
<li>Add PrepareConn hook to pgxpool (Jonathan Hall)</li>
<li>Reduce allocations in QueryContext (Dominique Lefevre)</li>
<li>Add MarshalJSON and UnmarshalJSON for pgtype.Uint32 (Panos Koutsovasilis)</li>
<li>Configure ping behavior on pgxpool with ShouldPing (Christian Kiely)</li>
<li>zeronull int types implement Int64Valuer and Int64Scanner (Li Zeghong)</li>
<li>Fix panic when receiving terminate connection message during CopyFrom (Michal Drausowski)</li>
<li>Fix statement cache not being invalidated on error during batch (Muhammadali Nazarov)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/a2fca037434a0a7096b095d4ed87cdffb03b626e"><code>a2fca03</code></a> Release v5.7.6</li>
<li><a href="https://github.com/jackc/pgx/commit/95fc31294f1d0eeabb0a8dafa36b875eac0df19c"><code>95fc312</code></a> Add link to github.com/KoNekoD/pgx-colon-query-rewriter</li>
<li><a href="https://github.com/jackc/pgx/commit/5534fa9a0206ef351e366ca985ed8e3af4e40f97"><code>5534fa9</code></a> Improve Rows docs</li>
<li><a href="https://github.com/jackc/pgx/commit/a295d68811e6946607f152986d1df961ae824382"><code>a295d68</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2368">#2368</a> from zeghong/zeronull-int-valuer</li>
<li><a href="https://github.com/jackc/pgx/commit/03f32c06bd5efb0fe32305824eb7850df56f0727"><code>03f32c0</code></a> Merge branch 'master' into zeronull-int-valuer</li>
<li><a href="https://github.com/jackc/pgx/commit/82fbe49fecdc4b0a5fa703b1e7fa7c6f80641739"><code>82fbe49</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2372">#2372</a> from WGH-/improve-batch-doc</li>
<li><a href="https://github.com/jackc/pgx/commit/594d9d65dc2e81d44b06ec83ee0e0b9be4f821d2"><code>594d9d6</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2367">#2367</a> from zeghong/zeronull-int-scanner</li>
<li><a href="https://github.com/jackc/pgx/commit/5a18241971b4ea2ae13df0abcf738d6620b6ff55"><code>5a18241</code></a> Merge branch 'master' into zeronull-int-scanner</li>
<li><a href="https://github.com/jackc/pgx/commit/cc34da5884b8f41dd9fb234ff6f90d563c108e5c"><code>cc34da5</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2370">#2370</a> from Saurabh2402/improvement/setup-linters</li>
<li><a href="https://github.com/jackc/pgx/commit/dd81f81e2fc0bf41c6c80d299720897538faf6fa"><code>dd81f81</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2369">#2369</a> from zeghong/go-doc-links</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.7.5...v5.7.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.12.1 to 9.13.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.13.0</h2>
<h2>Highlights</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
<li>Fixes on Read and Write buffer sizes and UniversalOptions</li>
</ul>
<h2>Changes</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>fix(test): fix a timing issue in pubsub test (<a href="https://redirect.github.com/redis/go-redis/pull/3498">#3498</a>)</li>
<li>Allow users to enable read-write splitting in failover mode. (<a href="https://redirect.github.com/redis/go-redis/pull/3482">#3482</a>)</li>
<li>Set the read/write buffer size of the sentinel client to 4KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3476">#3476</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>Support subscriptions against cluster slave nodes (<a href="https://redirect.github.com/redis/go-redis/pull/3480">#3480</a>)</li>
<li>Add wait metrics to otel (<a href="https://redirect.github.com/redis/go-redis/pull/3493">#3493</a>)</li>
<li>Clean failing timeout implementation (<a href="https://redirect.github.com/redis/go-redis/pull/3472">#3472</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>Do not assume that all non-IP hosts are loopbacks (<a href="https://redirect.github.com/redis/go-redis/pull/3085">#3085</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>fix(make test): Add default env in makefile (<a href="https://redirect.github.com/redis/go-redis/pull/3491">#3491</a>)</li>
<li>Update the introduction to running tests in README.md (<a href="https://redirect.github.com/redis/go-redis/pull/3495">#3495</a>)</li>
<li>test: Add comprehensive edge case tests for IncrByFloat command (<a href="https://redirect.github.com/redis/go-redis/pull/3477">#3477</a>)</li>
<li>Set the default read/write buffer size of Redis connection to 32KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3483">#3483</a>)</li>
<li>Bumps test image to 8.2.1-pre (<a href="https://redirect.github.com/redis/go-redis/pull/3478">#3478</a>)</li>
<li>fix UniversalOptions miss ReadBufferSize and WriteBufferSize options (<a href="https://redirect.github.com/redis/go-redis/pull/3485">#3485</a>)</li>
<li>chore(deps): bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/redis/go-redis/pull/3484">#3484</a>)</li>
<li>Removes dry run for stale issues policy (<a href="https://redirect.github.com/redis/go-redis/pull/3471">#3471</a>)</li>
<li>Update otel metrics URL (<a href="https://redirect.github.com/redis/go-redis/pull/3474">#3474</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/cybersmeashish"><code>@​cybersmeashish</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/suever"><code>@​suever</code></a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.13.0 (2025-09-03)</h1>
<h2>Highlights</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
<li>Fixes on Read and Write buffer sizes and UniversalOptions</li>
</ul>
<h2>Changes</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>fix(test): fix a timing issue in pubsub test (<a href="https://redirect.github.com/redis/go-redis/pull/3498">#3498</a>)</li>
<li>Allow users to enable read-write splitting in failover mode. (<a href="https://redirect.github.com/redis/go-redis/pull/3482">#3482</a>)</li>
<li>Set the read/write buffer size of the sentinel client to 4KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3476">#3476</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>Support subscriptions against cluster slave nodes (<a href="https://redirect.github.com/redis/go-redis/pull/3480">#3480</a>)</li>
<li>Add wait metrics to otel (<a href="https://redirect.github.com/redis/go-redis/pull/3493">#3493</a>)</li>
<li>Clean failing timeout implementation (<a href="https://redirect.github.com/redis/go-redis/pull/3472">#3472</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>Do not assume that all non-IP hosts are loopbacks (<a href="https://redirect.github.com/redis/go-redis/pull/3085">#3085</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>fix(make test): Add default env in makefile (<a href="https://redirect.github.com/redis/go-redis/pull/3491">#3491</a>)</li>
<li>Update the introduction to running tests in README.md (<a href="https://redirect.github.com/redis/go-redis/pull/3495">#3495</a>)</li>
<li>test: Add comprehensive edge case tests for IncrByFloat command (<a href="https://redirect.github.com/redis/go-redis/pull/3477">#3477</a>)</li>
<li>Set the default read/write buffer size of Redis connection to 32KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3483">#3483</a>)</li>
<li>Bumps test image to 8.2.1-pre (<a href="https://redirect.github.com/redis/go-redis/pull/3478">#3478</a>)</li>
<li>fix UniversalOptions miss ReadBufferSize and WriteBufferSize options (<a href="https://redirect.github.com/redis/go-redis/pull/3485">#3485</a>)</li>
<li>chore(deps): bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/redis/go-redis/pull/3484">#3484</a>)</li>
<li>Removes dry run for stale issues policy (<a href="https://redirect.github.com/redis/go-redis/pull/3471">#3471</a>)</li>
<li>Update otel metrics URL (<a href="https://redirect.github.com/redis/go-redis/pull/3474">#3474</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/cybersmeashish"><code>@​cybersmeashish</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/suever"><code>@​suever</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/52bda7a35ac3b6032a563e23329b912cb0a0a589"><code>52bda7a</code></a> chore(release): 9.13.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3500">#3500</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/19fdc488a76e982038f240642ab00a90d8c10d9d"><code>19fdc48</code></a> chore(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/issues/3499">#3499</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/10121e9e1c8800432e788a02fc0c24d740391702"><code>10121e9</code></a> feat(osscluster): Support subscriptions against cluster slave nodes (<a href="https://redirect.github.com/redis/go-redis/issues/3480">#3480</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6f41b600c5f69da9da9b33dafd1e733d2500b37a"><code>6f41b60</code></a> fix(client): Do not assume that all non-IP hosts are loopbacks (<a href="https://redirect.github.com/redis/go-redis/issues/3085">#3085</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f0058063a9e2b98e1a3c2f810499db6da4d312e6"><code>f005806</code></a> feat(otel): Add wait metrics to otel (<a href="https://redirect.github.com/redis/go-redis/issues/3493">#3493</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fafec3f3ce2b203257c17b3d9c9a51710d6e3e66"><code>fafec3f</code></a> Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/issues/3496">#3496</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6b9cbe8c547646508d209de57549077d388ac904"><code>6b9cbe8</code></a> fix(test): fix a timing issue in pubsub test (<a href="https://redirect.github.com/redis/go-redis/issues/3498">#3498</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e91f6ced09ee3a4e18bf347b35ffb473388fbee2"><code>e91f6ce</code></a> fix(make test): Add default env in makefile (<a href="https://redirect.github.com/redis/go-redis/issues/3491">#3491</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6bc723834038ec30f06e1028d0c7bdffe18df642"><code>6bc7238</code></a> Fix the ReplicaOnly option does not take effect when using NewFailoverCluster...</li>
<li><a href="https://github.com/redis/go-redis/commit/bb94ac7898d17bf747901f5eed0d4015bfd0ea5e"><code>bb94ac7</code></a> chore(readme): Update the introduction to running tests in README.md (<a href="https://redirect.github.com/redis/go-redis/issues/3495">#3495</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.12.1...v9.13.0">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.74.2 to 1.75.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.75.0</h2>
<h1>Behavior Changes</h1>
<ul>
<li>xds: Remove support for GRPC_EXPERIMENTAL_XDS_FALLBACK environment variable. Fallback support can no longer be disabled. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8482">#8482</a>)</li>
<li>stats: Introduce <code>DelayedPickComplete</code> event, a type alias of <code>PickerUpdated</code>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8465">#8465</a>)
<ul>
<li>This (combined) event will now be emitted only once per call, when a transport is successfully selected for the attempt.</li>
<li>OpenTelemetry metrics will no longer have multiple &quot;Delayed LB pick complete&quot; events in Go, matching other gRPC languages.</li>
<li>A future release will delete the <code>PickerUpdated</code> symbol.</li>
</ul>
</li>
<li>credentials: Properly apply <code>grpc.WithAuthority</code> as the highest-priority option for setting authority, above the setting in the credentials themselves. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8488">#8488</a>)
<ul>
<li>Now that this <code>WithAuthority</code> is available, the credentials should not be used to override the authority.</li>
</ul>
</li>
<li>round_robin: Randomize the order in which addresses are connected to in order to spread out initial RPC load between clients. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8438">#8438</a>)</li>
<li>server: Return status code INTERNAL when a client sends more than one request in unary and server streaming RPC. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8385">#8385</a>)
<ul>
<li>This is a behavior change but also a bug fix to bring gRPC-Go in line with the gRPC spec.</li>
</ul>
</li>
</ul>
<h1>New Features</h1>
<ul>
<li>dns: Add an environment variable (<code>GRPC_ENABLE_TXT_SERVICE_CONFIG</code>) to provide a way to disable TXT lookups in the DNS resolver (by setting it to <code>false</code>).  By default, TXT lookups are enabled, as they were previously. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8377">#8377</a>)</li>
</ul>
<h1>Bug Fixes</h1>
<ul>
<li>xds: Fix regression preventing empty node IDs in xDS bootstrap configuration. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8476">#8476</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/davinci26"><code>@​davinci26</code></a></li>
</ul>
</li>
<li>xds: Fix possible panic when certain invalid resources are encountered. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8412">#8412</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/wooffie"><code>@​wooffie</code></a></li>
</ul>
</li>
<li>xdsclient: Fix a rare panic caused by processing a response from a closed server. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8389">#8389</a>)</li>
<li>stats: Fix metric unit formatting by enclosing non-standard units like <code>call</code> and <code>endpoint</code> in curly braces to comply with UCUM and gRPC OpenTelemetry guidelines. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8481">#8481</a>)</li>
<li>xds: Fix possible panic when clusters are removed from the xds configuration. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8428">#8428</a>)</li>
<li>xdsclient: Fix a race causing &quot;resource doesn not exist&quot; when rapidly subscribing and unsubscribing to the same resource. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8369">#8369</a>)</li>
<li>client: When determining the authority, properly percent-encode (if needed, which is unlikely) when the target string omits the hostname and only specifies a port (<code>grpc.NewClient(&quot;:&lt;port-number-or-name&gt;&quot;)</code>). (<a href="https://redirect.github.com/grpc/grpc-go/issues/8488">#8488</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/b9788ef265596eda98a4391079c70c3992ed47cb"><code>b9788ef</code></a> Change version to 1.75.0 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8493">#8493</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/2bd74b28f5a7d464de4ed6927aef4b69abc0d3af"><code>2bd74b2</code></a> credentials: fix behavior of grpc.WithAuthority and credential handshake prec...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/9fa3267859958a7fa0141a8180102850f3d5842c"><code>9fa3267</code></a> xds: remove xds client fallback environment variable (<a href="https://redirect.github.com/grpc/grpc-go/issues/8482">#8482</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/62ec29fd9b3f9ea3cea6dc08a31e837aa92678b7"><code>62ec29f</code></a> grpc: Fix cardinality violations in non-client streaming RPCs. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8385">#8385</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/85240a5b02defe7b653ccba66866b4370c982b6a"><code>85240a5</code></a> stats: change non-standard units to annotations (<a href="https://redirect.github.com/grpc/grpc-go/issues/8481">#8481</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/ac13172781fae5593fd97ce07c3019c4044a71cd"><code>ac13172</code></a> update deps (<a href="https://redirect.github.com/grpc/grpc-go/issues/8478">#8478</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/0a895bc971a89c68c00070f792a28cc533846780"><code>0a895bc</code></a> examples/opentelemetry:  use experimental metrics in example (<a href="https://redirect.github.com/grpc/grpc-go/issues/8441">#8441</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/8b61e8f7b8fe9b0a4217336f6a4a31731338c3b2"><code>8b61e8f</code></a> xdsclient: do not process updates from closed server channels (<a href="https://redirect.github.com/grpc/grpc-go/issues/8389">#8389</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/7238ab1822875fdc2864e06fb224236dc7cbf3bf"><code>7238ab1</code></a> Allow empty nodeID (<a href="https://redirect.github.com/grpc/grpc-go/issues/8476">#8476</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/9186ebd774370e3b3232d1b202914ff8fc2c56d6"><code>9186ebd</code></a> cleanup: use slices.Equal to simplify code (<a href="https://redirect.github.com/grpc/grpc-go/issues/8472">#8472</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/grpc/grpc-go/compare/v1.74.2...v1.75.0">compare view</a></li>
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

## Reviews

### Review by @swadeley - Approved on September 08, 2025 at 06:49 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1194*
