---
type: pull_request
number: 500
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2023-12-11T04:21:29Z
updated: 2023-12-11T14:22:58Z
closed: 2023-12-11T14:22:49Z
merged: 2023-12-11T14:22:49Z
base_branch: main
head_branch: dependabot/go_modules/go-8878dab762
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/500
---

# Pull Request #500: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: December 11, 2023 at 04:21 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-8878dab762`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/spf13/viper](https://github.com/spf13/viper) | `1.17.0` | `1.18.1` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.23.5` | `1.24.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.16.9` | `1.16.12` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.29.2` | `1.29.5` |
| [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) | `2023.11.1701177874` | `2023.12.1701890198` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.5.0` | `5.5.1` |

Updates `github.com/spf13/viper` from 1.17.0 to 1.18.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/spf13/viper/releases">github.com/spf13/viper's releases</a>.</em></p>
<blockquote>
<h2>v1.18.1</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Merge missing struct keys inside UnmarshalExact by <a href="https://github.com/krakowski"><code>@​krakowski</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1704">spf13/viper#1704</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/spf13/viper/compare/v1.18.0...v1.18.1">https://github.com/spf13/viper/compare/v1.18.0...v1.18.1</a></p>
<h2>v1.18.0</h2>
<h2>Major changes</h2>
<p>Highlighting some of the changes for better visibility.</p>
<p>Please share your feedback in the Discussion forum. Thanks! ❤️</p>
<h3><code>AutomaticEnv</code> works with <code>Unmarshal</code></h3>
<p>Previously,  environment variables that weren't bound manually or had no defaults could not be mapped by <code>Unmarshal</code>. (The problem is explained in details in this issue: <a href="https://redirect.github.com/spf13/viper/issues/761">#761</a>)</p>
<p><a href="https://redirect.github.com/spf13/viper/issues/1429">#1429</a> introduced a solution that solves that issue.</p>
<h2>What's Changed</h2>
<h3>Enhancements 🚀</h3>
<ul>
<li>chore: rename files according to enabled build tags by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1642">spf13/viper#1642</a></li>
<li>test: replace ifs with asserts to simplify tests by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1656">spf13/viper#1656</a></li>
<li>ci: enable test shuffle and fix tests by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1643">spf13/viper#1643</a></li>
<li>fix: gocritic lint issues by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1696">spf13/viper#1696</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Implement viper.BindStruct for automatic unmarshalling from environment variables by <a href="https://github.com/krakowski"><code>@​krakowski</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1429">spf13/viper#1429</a></li>
<li>fix isPathShadowedInFlatMap type cast bug by <a href="https://github.com/linuxsong"><code>@​linuxsong</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1585">spf13/viper#1585</a></li>
</ul>
<h3>Dependency Updates ⬆️</h3>
<ul>
<li>build(deps): bump github/codeql-action from 2.21.9 to 2.22.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1661">spf13/viper#1661</a></li>
<li>build(deps): bump golang.org/x/net from 0.15.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1659">spf13/viper#1659</a></li>
<li>build(deps): bump actions/checkout from 4.1.0 to 4.1.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1663">spf13/viper#1663</a></li>
<li>build(deps): bump actions/github-script from 6.4.1 to 7.0.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1686">spf13/viper#1686</a></li>
<li>build(deps): bump github/codeql-action from 2.22.3 to 2.22.8 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1688">spf13/viper#1688</a></li>
<li>build(deps): bump github.com/spf13/afero from 1.10.0 to 1.11.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1692">spf13/viper#1692</a></li>
<li>build(deps): bump actions/dependency-review-action from 3.1.0 to 3.1.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1690">spf13/viper#1690</a></li>
<li>build(deps): bump cachix/install-nix-action from 23 to 24 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1689">spf13/viper#1689</a></li>
<li>build(deps): bump github.com/nats-io/nkeys from 0.4.5 to 0.4.6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1672">spf13/viper#1672</a></li>
<li>build(deps): bump github.com/spf13/cast from 1.5.1 to 1.6.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1691">spf13/viper#1691</a></li>
<li>build(deps): bump github.com/fsnotify/fsnotify from 1.6.0 to 1.7.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1668">spf13/viper#1668</a></li>
<li>chore: update dependencies by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1694">spf13/viper#1694</a></li>
<li>chore: update crypt by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1701">spf13/viper#1701</a></li>
</ul>
<h3>Other Changes</h3>
<ul>
<li>Add info about multiple hosts for remote config by <a href="https://github.com/KaymeKaydex"><code>@​KaymeKaydex</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1684">spf13/viper#1684</a></li>
<li>refactor: drop fsonitfy wrapper by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1693">spf13/viper#1693</a></li>
<li>Note Get* behavior on parse failure by <a href="https://github.com/scop"><code>@​scop</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1687">spf13/viper#1687</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/spf13/viper/commit/fb6eb1e8e9fbaa424b113248af35f048c5d86dea"><code>fb6eb1e</code></a> fix: merge missing struct keys inside UnmarshalExact</li>
<li><a href="https://github.com/spf13/viper/commit/f5fcb4a104f76357c270c3944c26a2e2c8a06fdf"><code>f5fcb4a</code></a> chore: update crypt</li>
<li><a href="https://github.com/spf13/viper/commit/f7363633d1c17fd98cd70783fabc133c41d3f40e"><code>f736363</code></a> fix isPathShadowedInFlatMap type cast bug (<a href="https://redirect.github.com/spf13/viper/issues/1585">#1585</a>)</li>
<li><a href="https://github.com/spf13/viper/commit/36a38682ba4fe1bb823952e52c461c297bb1767e"><code>36a3868</code></a> Review changes</li>
<li><a href="https://github.com/spf13/viper/commit/f0c4ccd6cd7527b198aecc6d50b3ffd6d06e36d1"><code>f0c4ccd</code></a> fix: gocritic lint issues</li>
<li><a href="https://github.com/spf13/viper/commit/3a23b80b1120cd3e6aaf21d91784fce81b8054df"><code>3a23b80</code></a> ci: enable test shuffle; fix tests</li>
<li><a href="https://github.com/spf13/viper/commit/73dfb94c57ad48bcdf3d40b7f47d69c2962d800d"><code>73dfb94</code></a> feat: make Unmarshal work with AutomaticEnv</li>
<li><a href="https://github.com/spf13/viper/commit/6ea31ae4ca569f048d53a06e91574f1ebe4b0225"><code>6ea31ae</code></a> refactor: move all settings code to a getter</li>
<li><a href="https://github.com/spf13/viper/commit/c4dcd31f68e5d77ce447c0091dd1ca6d7e169807"><code>c4dcd31</code></a> fix: godot lint issues</li>
<li><a href="https://github.com/spf13/viper/commit/4c9b2a26ae1de051cce83f4aea2527dfa8c00c53"><code>4c9b2a2</code></a> Note Get* behavior on parse failure</li>
<li>Additional commits viewable in <a href="https://github.com/spf13/viper/compare/v1.17.0...v1.18.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.23.5 to 1.24.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d643a8f0a8ad09075f41a60acba6a208cb36c58"><code>0d643a8</code></a> Release 2023-12-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e38b534ccea2707475381f985bc9236970bab554"><code>e38b534</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f394daf705a925e626f6a5c52aa433ff5504c7f1"><code>f394daf</code></a> Update SDK's smithy-go dependency to v1.19.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4fb654cee93178383ad85934adc3e16cb906baf7"><code>4fb654c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8594ca09ef8095c3287128f364527b1fea3bc49c"><code>8594ca0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e33a34c60f1001aa0a85cb9eb682d1b3153bed43"><code>e33a34c</code></a> fix codegen ci (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2418">#2418</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/57a201c7bb214737f78cf871fbeafaa929bf1d8d"><code>57a201c</code></a> regenerate partitions (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2415">#2415</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be06f0299409272364d03e76196939db023ef807"><code>be06f02</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2344">#2344</a> from aws/feat-request-compression-2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70bf4107034085068fb6847667657fbb78367be4"><code>70bf410</code></a> Change config loading logic</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8c02c46d4857cbf2c1eb607484e3b94458fe50eb"><code>8c02c46</code></a> Release 2023-12-06</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.23.5...v1.24.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.9 to 1.16.12
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1091d0a12dd9d5894ad3961214a6bea5def3062"><code>d1091d0</code></a> Release 2022-08-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a1140b1b7e70a06b9b49493f095853fc891d3599"><code>a1140b1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8f3045a981b27bde412ecd3143bbea128b34285e"><code>8f3045a</code></a> Update SDK's smithy-go dependency to v1.13.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/462d0469260d76219599fb16d025d9db50ea6c56"><code>462d046</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7aba71b74a65085bc9d1ffa60b893c90011b792d"><code>7aba71b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1609fe847a5af31ae8a4a306b779e7c13bd45bb6"><code>1609fe8</code></a> credentials/ssocreds: Add SSOTokenProvider for Bearer Token auth (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1818">#1818</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8e755b49d93958105eefdd8bbc3f2500e2c98738"><code>8e755b4</code></a> Release 2022-08-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d33e0c036c685d4e7d1305216a474e541047aea5"><code>d33e0c0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/129270893ae61b7eef49470a157344f23f722453"><code>1292708</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d6d53e3e9422ffcfe80d842a351d33bb8dfde70"><code>7d6d53e</code></a> add GitHub code owner for auto assigned reviewers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1817">#1817</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.9...v1.16.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.29.2 to 1.29.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93c3f1871b862d743e0bd2e2e7180246df3a9212"><code>93c3f18</code></a> Release 2022-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7254028f8bc89095326d9e3657fdbc98b98cca94"><code>7254028</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f43ad83db1b3da1c2ea37857524148c91189cb4c"><code>f43ad83</code></a> Update SDK's smithy-go dependency to v1.13.5</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/77d257ee120e67d45a5de6f0d6478f313a21b92a"><code>77d257e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/779e29ff5a4bcebe1ab7088ab12c4c95ce06f8aa"><code>779e29f</code></a> Release 2022-12-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f64d7d2b0a0033996b32ba9e1b18e5a923452b84"><code>f64d7d2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bc59f75ee4683ca886c3d701b49bb81db2efd4d"><code>9bc59f7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d9c18aa2bdd4c237a4919452f58e29c20ba484cc"><code>d9c18aa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0259b169b753daf77ad332c680a9ad1e3f56753d"><code>0259b16</code></a> Release 2022-11-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ee0277f1abad4856afc13ced2bfb90a43dbd9d34"><code>ee0277f</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.29.2...service/s3/v1.29.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2023` from 2023.11.1701177874 to 2023.12.1701890198
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/24b7e22030912744f4f6ac418d412e45a9576fcd"><code>24b7e22</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97984db6aa5037ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/6d00048b6585ba82ae1480a93e765653db739912"><code>6d00048</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a8474e4282edb037d356e3e82bb4...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2023.11.1701177874...release/v2023.12.1701890198">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.5.0 to 5.5.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.5.1 (December 9, 2023)</h1>
<ul>
<li>Add CopyFromFunc helper function. (robford)</li>
<li>Add PgConn.Deallocate method that uses PostgreSQL protocol Close message.</li>
<li>pgx uses new PgConn.Deallocate method. This allows deallocating statements to work in a failed transaction. This fixes a case where the prepared statement map could become invalid.</li>
<li>Fix: Prefer driver.Valuer over json.Marshaler for json fields. (Jacopo)</li>
<li>Fix: simple protocol SQL sanitizer previously panicked if an invalid $0 placeholder was used. This now returns an error instead. (maksymnevajdev)</li>
<li>Add pgtype.Numeric.ScanScientific (Eshton Robateau)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/ba05097642a802eb5a76343f27a9a3515ed86498"><code>ba05097</code></a> Release v5.5.1</li>
<li><a href="https://github.com/jackc/pgx/commit/384fe7775c9d3b981959f51cd41c47da5aff8016"><code>384fe77</code></a> Batch.Queue: document always uses the conn's DefaultQueryExecMode</li>
<li><a href="https://github.com/jackc/pgx/commit/20bf953a17fc7cf6619536a87bd285b9273dbcf9"><code>20bf953</code></a> pull out changes into new public function</li>
<li><a href="https://github.com/jackc/pgx/commit/12582a0fd451ab0f3d6fde2cea213a9cdd3622ba"><code>12582a0</code></a> bitsize largest option is 64</li>
<li><a href="https://github.com/jackc/pgx/commit/905f2526672dcceb32f9c9e1e7ded3b10760d0c2"><code>905f252</code></a> uncomment tests</li>
<li><a href="https://github.com/jackc/pgx/commit/9927e14bbf44db9a9eb7da723276427851a8e41e"><code>9927e14</code></a> remove dead line</li>
<li><a href="https://github.com/jackc/pgx/commit/95b2f85e60740c62f95a8dc7cfedc7562cb629aa"><code>95b2f85</code></a> support scientific notation big floats</li>
<li><a href="https://github.com/jackc/pgx/commit/913e4c848720ca998441e80d2224e8a084c1a442"><code>913e4c8</code></a> Update changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/31321c2017fc68ce86aaf05218ae3e6e44426591"><code>31321c2</code></a> Add race detector to bug report template</li>
<li><a href="https://github.com/jackc/pgx/commit/319c3172f2d74a6e1e9d2efb1d94ac6f45e635b9"><code>319c317</code></a> fix panic in prepared sql</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.5.0...v5.5.1">compare view</a></li>
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

### Comment by @app-sre-bot on December 11, 2023 at 04:22 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on December 11, 2023 at 12:56 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on December 11, 2023 at 02:22 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/500*
