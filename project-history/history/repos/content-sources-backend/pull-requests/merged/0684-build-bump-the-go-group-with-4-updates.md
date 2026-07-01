---
type: pull_request
number: 684
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2024-06-03T04:42:31Z
updated: 2024-06-03T10:20:02Z
closed: 2024-06-03T10:19:54Z
merged: 2024-06-03T10:19:54Z
base_branch: main
head_branch: dependabot/go_modules/go-58601118aa
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/684
---

# Pull Request #684: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: June 03, 2024 at 04:42 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-58601118aa`

## Description

Bumps the go group with 4 updates: [github.com/content-services/tang](https://github.com/content-services/tang), [github.com/spf13/viper](https://github.com/spf13/viper), [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) and [github.com/redis/go-redis/v9](https://github.com/redis/go-redis).

Updates `github.com/content-services/tang` from 0.0.7 to 0.0.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/tang/releases">github.com/content-services/tang's releases</a>.</em></p>
<blockquote>
<h2>v0.0.8</h2>
<h2>What's Changed</h2>
<ul>
<li>Refs 3656: list cves in template errata by <a href="https://github.com/xbhouse"><code>@​xbhouse</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/10">content-services/tang#10</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.7...v0.0.8">https://github.com/content-services/tang/compare/v0.0.7...v0.0.8</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/4840e6494f1500f3f467940678925736040eb7d6"><code>4840e64</code></a> Refs 3656: list cves in template errata (<a href="https://redirect.github.com/content-services/tang/issues/10">#10</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.7...v0.0.8">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/spf13/viper` from 1.18.2 to 1.19.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/spf13/viper/releases">github.com/spf13/viper's releases</a>.</em></p>
<blockquote>
<h2>v1.19.0</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>fix!: hide struct binding behind a feature flag by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1720">spf13/viper#1720</a></li>
</ul>
<h3>Dependency Updates ⬆️</h3>
<ul>
<li>build(deps): bump github/codeql-action from 2.22.8 to 2.22.9 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1705">spf13/viper#1705</a></li>
<li>build(deps): bump actions/setup-go from 4.1.0 to 5.0.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1703">spf13/viper#1703</a></li>
<li>build(deps): bump github/codeql-action from 2.22.9 to 3.22.11 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1713">spf13/viper#1713</a></li>
<li>build(deps): bump github.com/pelletier/go-toml/v2 from 2.1.0 to 2.1.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1711">spf13/viper#1711</a></li>
<li>build(deps): bump golang.org/x/crypto from 0.16.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1722">spf13/viper#1722</a></li>
<li>build(deps): bump github/codeql-action from 3.22.11 to 3.23.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1734">spf13/viper#1734</a></li>
<li>build(deps): bump actions/dependency-review-action from 3.1.4 to 3.1.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1731">spf13/viper#1731</a></li>
<li>build(deps): bump mheap/github-action-required-labels from 5.1.0 to 5.2.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1743">spf13/viper#1743</a></li>
<li>build(deps): bump github/codeql-action from 3.23.0 to 3.23.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1742">spf13/viper#1742</a></li>
<li>build(deps): bump actions/dependency-review-action from 3.1.5 to 4.0.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1739">spf13/viper#1739</a></li>
<li>build(deps): bump cachix/install-nix-action from 24 to 25 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1737">spf13/viper#1737</a></li>
<li>build(deps): bump github/codeql-action from 3.23.2 to 3.24.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1751">spf13/viper#1751</a></li>
<li>build(deps): bump github/codeql-action from 3.24.0 to 3.24.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1760">spf13/viper#1760</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.0.0 to 4.1.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1761">spf13/viper#1761</a></li>
<li>build(deps): bump golangci/golangci-lint-action from 3.7.0 to 4.0.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1757">spf13/viper#1757</a></li>
<li>build(deps): bump mheap/github-action-required-labels from 5.2.0 to 5.3.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1759">spf13/viper#1759</a></li>
<li>build(deps): bump github/codeql-action from 3.24.1 to 3.24.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1763">spf13/viper#1763</a></li>
<li>build(deps): bump github.com/sagikazarmark/crypt from 0.17.0 to 0.18.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1774">spf13/viper#1774</a></li>
<li>build(deps): bump github/codeql-action from 3.24.3 to 3.24.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1770">spf13/viper#1770</a></li>
<li>build(deps): bump github.com/stretchr/testify from 1.8.4 to 1.9.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1776">spf13/viper#1776</a></li>
<li>build(deps): bump github/codeql-action from 3.24.5 to 3.24.6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1775">spf13/viper#1775</a></li>
<li>build(deps): bump cachix/install-nix-action from 25 to 26 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1778">spf13/viper#1778</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.1.0 to 4.1.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1767">spf13/viper#1767</a></li>
<li>build(deps): bump github/codeql-action from 3.24.6 to 3.24.9 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1790">spf13/viper#1790</a></li>
<li>build(deps): bump mheap/github-action-required-labels from 5.3.0 to 5.4.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1789">spf13/viper#1789</a></li>
<li>build(deps): bump actions/checkout from 4.1.1 to 4.1.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1780">spf13/viper#1780</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.1.3 to 4.2.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1793">spf13/viper#1793</a></li>
<li>chore: upgrade crypt by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1794">spf13/viper#1794</a></li>
<li>build(deps): bump github.com/pelletier/go-toml/v2 from 2.1.1 to 2.2.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1788">spf13/viper#1788</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.2.4 to 4.2.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1796">spf13/viper#1796</a></li>
<li>build(deps): bump github.com/pelletier/go-toml/v2 from 2.2.0 to 2.2.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1804">spf13/viper#1804</a></li>
<li>build(deps): bump github/codeql-action from 3.24.9 to 3.25.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1806">spf13/viper#1806</a></li>
<li>build(deps): bump golang.org/x/net from 0.22.0 to 0.23.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1807">spf13/viper#1807</a></li>
<li>build(deps): bump actions/checkout from 4.1.2 to 4.1.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1808">spf13/viper#1808</a></li>
<li>build(deps): bump actions/checkout from 4.1.3 to 4.1.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1813">spf13/viper#1813</a></li>
<li>build(deps): bump github/codeql-action from 3.25.1 to 3.25.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1811">spf13/viper#1811</a></li>
<li>build(deps): bump mheap/github-action-required-labels from 5.4.0 to 5.4.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1817">spf13/viper#1817</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.2.5 to 4.3.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1821">spf13/viper#1821</a></li>
<li>build(deps): bump github.com/pelletier/go-toml/v2 from 2.2.1 to 2.2.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1822">spf13/viper#1822</a></li>
<li>build(deps): bump actions/setup-go from 5.0.0 to 5.0.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1824">spf13/viper#1824</a></li>
<li>build(deps): bump github/codeql-action from 3.25.2 to 3.25.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1828">spf13/viper#1828</a></li>
<li>build(deps): bump golangci/golangci-lint-action from 4.0.0 to 6.0.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1829">spf13/viper#1829</a></li>
<li>build(deps): bump github/codeql-action from 3.25.4 to 3.25.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1844">spf13/viper#1844</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/spf13/viper/commit/b9733f03ad014259d08f405c13e3d7f469fa1a8e"><code>b9733f0</code></a> build(deps): bump actions/checkout from 4.1.4 to 4.1.6</li>
<li><a href="https://github.com/spf13/viper/commit/6ecc5c810f98ebe1984fb55ce5ad5c6afc3fc44e"><code>6ecc5c8</code></a> build(deps): bump cachix/install-nix-action from 26 to 27</li>
<li><a href="https://github.com/spf13/viper/commit/248c6fdd03eafa8a524f75af5b5c059af0c952bd"><code>248c6fd</code></a> build(deps): bump github/codeql-action from 3.25.4 to 3.25.7</li>
<li><a href="https://github.com/spf13/viper/commit/abea773f16452659b7b47770862dfb29b61e82d1"><code>abea773</code></a> Update references to bketelsen/crypt</li>
<li><a href="https://github.com/spf13/viper/commit/f17acb4fd40fcd0e22d09217b6680a6297092a33"><code>f17acb4</code></a> build(deps): bump golangci/golangci-lint-action from 4.0.0 to 6.0.1</li>
<li><a href="https://github.com/spf13/viper/commit/8e285a5880531e2ed4933648e388353df1db5106"><code>8e285a5</code></a> build(deps): bump github/codeql-action from 3.25.2 to 3.25.4</li>
<li><a href="https://github.com/spf13/viper/commit/40176207a58a5f2b0b95e72da8b6ebf2df02416e"><code>4017620</code></a> build(deps): bump actions/setup-go from 5.0.0 to 5.0.1</li>
<li><a href="https://github.com/spf13/viper/commit/b67e814385886e962dfe9c304a69e42b8880d3a8"><code>b67e814</code></a> build(deps): bump github.com/pelletier/go-toml/v2 from 2.2.1 to 2.2.2</li>
<li><a href="https://github.com/spf13/viper/commit/4a182c767b9521adfd8981bea4015457105f608e"><code>4a182c7</code></a> build(deps): bump actions/dependency-review-action from 4.2.5 to 4.3.2</li>
<li><a href="https://github.com/spf13/viper/commit/45a0e1214a55ce1fcdc6e66526177e2dc83f2bef"><code>45a0e12</code></a> build(deps): bump mheap/github-action-required-labels</li>
<li>Additional commits viewable in <a href="https://github.com/spf13/viper/compare/v1.18.2...v1.19.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.5.1716497237 to 2024.5.1717078140
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/b106cbd486d60055b82449fd9ef020135d8b870c"><code>b106cbd</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a959bea52fb7a8d53bd49838...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.5.1716497237...release/v2024.5.1717078140">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.5.1 to 9.5.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.5.2</h2>
<h1>Changes</h1>
<ul>
<li>fix: fix <a href="https://redirect.github.com/redis/go-redis/issues/2681">#2681</a> (<a href="https://redirect.github.com/redis/go-redis/issues/2998">#2998</a>)</li>
<li>Remove skipping span creation by checking parent spans (<a href="https://redirect.github.com/redis/go-redis/issues/2980">#2980</a>)</li>
<li>Handle IPv6 in isMovedError (<a href="https://redirect.github.com/redis/go-redis/issues/2981">#2981</a>)</li>
<li>Fix XGroup first pos key (<a href="https://redirect.github.com/redis/go-redis/issues/2983">#2983</a>)</li>
<li>Adding BitfieldRo in BitMapCmdable interface (<a href="https://redirect.github.com/redis/go-redis/issues/2962">#2962</a>)</li>
<li>Optimize docs useless imports and typo (<a href="https://redirect.github.com/redis/go-redis/issues/2970">#2970</a>)</li>
<li>chore: fix some comments (<a href="https://redirect.github.com/redis/go-redis/issues/2967">#2967</a>)</li>
<li>Fix for issues <a href="https://redirect.github.com/redis/go-redis/issues/2959">#2959</a> and <a href="https://redirect.github.com/redis/go-redis/issues/2960">#2960</a> (<a href="https://redirect.github.com/redis/go-redis/issues/2961">#2961</a>)</li>
<li>fix: <a href="https://redirect.github.com/redis/go-redis/issues/2956">#2956</a> (<a href="https://redirect.github.com/redis/go-redis/issues/2957">#2957</a>)</li>
<li>fix misuses of a vs an (<a href="https://redirect.github.com/redis/go-redis/issues/2936">#2936</a>)</li>
<li>add server address and port span attributes to redis otel trace instrumentation (<a href="https://redirect.github.com/redis/go-redis/issues/2826">#2826</a>)</li>
<li>chore(deps): bump google.golang.org/protobuf from 1.32.0 to 1.33.0 in /example/otel (<a href="https://redirect.github.com/redis/go-redis/issues/2944">#2944</a>)</li>
<li>Remove secrets from Redis Enterprise CI (<a href="https://redirect.github.com/redis/go-redis/issues/2938">#2938</a>)</li>
<li>Fix monitor on go 1.19 (<a href="https://redirect.github.com/redis/go-redis/issues/2908">#2908</a>)</li>
<li>chore(deps): bump google.golang.org/protobuf from 1.28.1 to 1.33.0 in /extra/redisprometheus (<a href="https://redirect.github.com/redis/go-redis/issues/2942">#2942</a>)</li>
<li>Change RE image to public RE image (<a href="https://redirect.github.com/redis/go-redis/issues/2935">#2935</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/XSAM"><code>@​XSAM</code></a>, <a href="https://github.com/akash14darshan"><code>@​akash14darshan</code></a>, <a href="https://github.com/daviddzxy"><code>@​daviddzxy</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/esara"><code>@​esara</code></a>, <a href="https://github.com/hakusai22"><code>@​hakusai22</code></a>, <a href="https://github.com/hishope"><code>@​hishope</code></a>, <a href="https://github.com/kindknow"><code>@​kindknow</code></a>, <a href="https://github.com/monkey92t"><code>@​monkey92t</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/singular-seal"><code>@​singular-seal</code></a> and deferdeter</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/2d8fa02ac2b0cda80410674a4abad00ea35217e8"><code>2d8fa02</code></a> fix: fix <a href="https://redirect.github.com/redis/go-redis/issues/2681">#2681</a> (<a href="https://redirect.github.com/redis/go-redis/issues/2998">#2998</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/0f0a28464c3db3ff13f80c44db43dbb4ad6ac2c7"><code>0f0a284</code></a> Remove skipping span creation by checking parent spans (<a href="https://redirect.github.com/redis/go-redis/issues/2980">#2980</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b64d9deef330a51cfbd3e0425b6c26b27c1ee370"><code>b64d9de</code></a> Handle IPv6 in isMovedError (<a href="https://redirect.github.com/redis/go-redis/issues/2981">#2981</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fa9edecebc15475ced01a983d7f6b18db46840b9"><code>fa9edec</code></a> Fix XGroup first pos key (<a href="https://redirect.github.com/redis/go-redis/issues/2983">#2983</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/90c7a414ac819ab9e15cc01cd0bfa6f357236585"><code>90c7a41</code></a> Adding BitfieldRo in BitMapCmdable interface (<a href="https://redirect.github.com/redis/go-redis/issues/2962">#2962</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/9f1522a91e666ee40f807945e49396b09e4f87c2"><code>9f1522a</code></a> Fix typo in comment (<a href="https://redirect.github.com/redis/go-redis/issues/2972">#2972</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f7f34feddf26644661276e90ebd87e748d70a050"><code>f7f34fe</code></a> Optimize docs useless imports and typo (<a href="https://redirect.github.com/redis/go-redis/issues/2970">#2970</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6960bcc08db24d46c6e3b76a49f7087347f51eb6"><code>6960bcc</code></a> chore: fix some comments (<a href="https://redirect.github.com/redis/go-redis/issues/2967">#2967</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f5496d14ddb029523b44ef97bd5e662569728037"><code>f5496d1</code></a> Fix for issues <a href="https://redirect.github.com/redis/go-redis/issues/2959">#2959</a> and <a href="https://redirect.github.com/redis/go-redis/issues/2960">#2960</a> (<a href="https://redirect.github.com/redis/go-redis/issues/2961">#2961</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6833d2f8e18bf050d3d34a3a2bd9881a4de1525c"><code>6833d2f</code></a> fix: <a href="https://redirect.github.com/redis/go-redis/issues/2956">#2956</a> (<a href="https://redirect.github.com/redis/go-redis/issues/2957">#2957</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.5.1...v9.5.2">compare view</a></li>
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

### Comment by @app-sre-bot on June 03, 2024 at 04:43 AM UTC

Can one of the admins verify this patch?

---

## Reviews

### Review by @swadeley - Approved on June 03, 2024 at 10:19 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/684*
