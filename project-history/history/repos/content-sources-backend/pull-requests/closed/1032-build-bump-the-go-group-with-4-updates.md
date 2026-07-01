---
type: pull_request
number: 1032
title: "Build: Bump the go group with 4 updates"
state: closed
author: dependabot
created: 2025-03-17T04:53:38Z
updated: 2025-03-18T15:58:22Z
closed: 2025-03-18T15:58:20Z
base_branch: main
head_branch: dependabot/go_modules/go-bbec00a15f
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1032
---

# Pull Request #1032: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: March 17, 2025 at 04:53 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-bbec00a15f`

## Description

Bumps the go group with 4 updates: [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi), [github.com/spf13/viper](https://github.com/spf13/viper), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2).

Updates `github.com/getkin/kin-openapi` from 0.129.0 to 0.130.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.130.0</h2>
<h2>What's Changed</h2>
<ul>
<li>feat(openapi3gen): Customize json.RawMessage by <a href="https://github.com/kyleconroy"><code>@​kyleconroy</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1050">getkin/kin-openapi#1050</a></li>
<li>openapi3gen: Fix issue with separate component generated for time.Time by <a href="https://github.com/d1vbyz3r0"><code>@​d1vbyz3r0</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1052">getkin/kin-openapi#1052</a></li>
<li>openapi3filter: Remove redundant ExcludeResponseBody check by <a href="https://github.com/tatsumack"><code>@​tatsumack</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1056">getkin/kin-openapi#1056</a></li>
<li>openapi3: use <strong>origin</strong> to minimize collisions by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1057">getkin/kin-openapi#1057</a></li>
<li>openapi3: delete origin keys only when IncludeOrigin=true by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1055">getkin/kin-openapi#1055</a></li>
<li>openapi3filter: apply default values of an array in a query param with exploded = false by <a href="https://github.com/nhochstr"><code>@​nhochstr</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1054">getkin/kin-openapi#1054</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/kyleconroy"><code>@​kyleconroy</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1050">getkin/kin-openapi#1050</a></li>
<li><a href="https://github.com/d1vbyz3r0"><code>@​d1vbyz3r0</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1052">getkin/kin-openapi#1052</a></li>
<li><a href="https://github.com/tatsumack"><code>@​tatsumack</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1056">getkin/kin-openapi#1056</a></li>
<li><a href="https://github.com/nhochstr"><code>@​nhochstr</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1054">getkin/kin-openapi#1054</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.129.0...v0.130.0">https://github.com/getkin/kin-openapi/compare/v0.129.0...v0.130.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/6da871e0e170b7637eb568c265c08bc2b5d6e7a3"><code>6da871e</code></a> openapi3filter: apply default values of an array in a query param with explod...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/a34baf04836f1c5ee3b315a30150f9a0100301cd"><code>a34baf0</code></a> openapi3: delete origin keys only when IncludeOrigin=true (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1055">#1055</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/2d3e67aa2127f007ee9ca0931d48132a7ec111b1"><code>2d3e67a</code></a> use <strong>origin</strong> to minimize collisions (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1057">#1057</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/e3d68dc02b03b9dbe6bab607142a2b7181b6f2e2"><code>e3d68dc</code></a> Remove redundant ExcludeResponseBody check in ValidateResponse (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1056">#1056</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/050a930912661dd1e1a955421d60d38d2f6dd5a3"><code>050a930</code></a> openapi3gen: Fix issue with separate component generated for time.Time (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1052">#1052</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/72fb819c6aa64a40d1c90dd37c33598840db1221"><code>72fb819</code></a> feat(openapi3gen): Customize json.RawMessage (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1050">#1050</a>)</li>
<li>See full diff in <a href="https://github.com/getkin/kin-openapi/compare/v0.129.0...v0.130.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/spf13/viper` from 1.19.0 to 1.20.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/spf13/viper/releases">github.com/spf13/viper's releases</a>.</em></p>
<blockquote>
<h2>v1.20.0</h2>
<!-- raw HTML omitted -->
<blockquote>
<p>[!WARNING]
This release includes a few minor breaking changes. Read the <a href="https://github.com/spf13/viper/blob/master/UPGRADE.md#v120x">upgrade guide</a> for details.</p>
</blockquote>
<h2>What's Changed</h2>
<h3>Exciting New Features 🎉</h3>
<ul>
<li>New encoding layer by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1869">spf13/viper#1869</a></li>
</ul>
<h3>Enhancements 🚀</h3>
<ul>
<li>Drop Go 1.20 support by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1846">spf13/viper#1846</a></li>
<li>Drop slog shim by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1848">spf13/viper#1848</a></li>
<li>Replace file searching API with a finder by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1849">spf13/viper#1849</a></li>
<li>Finder feature flag by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1852">spf13/viper#1852</a></li>
<li>Allow setting options on the global Viper instance by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1856">spf13/viper#1856</a></li>
<li>Add experimental flag for bind struct by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1854">spf13/viper#1854</a></li>
<li>Make the remote package a separate module by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1860">spf13/viper#1860</a></li>
<li>Add decoder hook option by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1872">spf13/viper#1872</a></li>
<li>Encoder improvements by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1885">spf13/viper#1885</a></li>
<li>Get uint8 by <a href="https://github.com/martinconic"><code>@​martinconic</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1894">spf13/viper#1894</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Fix missing config type when reading from a buffer by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1857">spf13/viper#1857</a></li>
<li>fix: do not allow setting dependencies to nil values by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1871">spf13/viper#1871</a></li>
<li>feat: copy keydelim from parent chart in viper.Sub() by <a href="https://github.com/obs-gh-alexlew"><code>@​obs-gh-alexlew</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1887">spf13/viper#1887</a></li>
</ul>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Drop encoding formats: HCL, Java properties, INI by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1870">spf13/viper#1870</a></li>
</ul>
<h3>Dependency Updates ⬆️</h3>
<ul>
<li>chore: update mapstructure by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1723">spf13/viper#1723</a></li>
<li>chore: update crypt by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1834">spf13/viper#1834</a></li>
<li>build(deps): bump github/codeql-action from 3.25.7 to 3.25.8 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1853">spf13/viper#1853</a></li>
<li>Revert to go-difflib and go-spew releases by <a href="https://github.com/skitt"><code>@​skitt</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1861">spf13/viper#1861</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.3.2 to 4.3.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1862">spf13/viper#1862</a></li>
<li>build(deps): bump github/codeql-action from 3.25.8 to 3.25.10 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1865">spf13/viper#1865</a></li>
<li>build(deps): bump actions/checkout from 4.1.6 to 4.1.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1864">spf13/viper#1864</a></li>
<li>chore: update crypt by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1866">spf13/viper#1866</a></li>
<li>build(deps): bump github/codeql-action from 3.25.10 to 3.25.11 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1876">spf13/viper#1876</a></li>
<li>build(deps): bump google.golang.org/grpc from 1.64.0 to 1.64.1 in /remote by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1878">spf13/viper#1878</a></li>
<li>build(deps): bump actions/setup-go from 5.0.1 to 5.0.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1879">spf13/viper#1879</a></li>
<li>build(deps): bump actions/dependency-review-action from 4.3.3 to 4.3.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1881">spf13/viper#1881</a></li>
<li>build(deps): bump github/codeql-action from 3.25.11 to 3.25.12 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1880">spf13/viper#1880</a></li>
<li>build(deps): bump github/codeql-action from 3.25.12 to 3.25.13 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1883">spf13/viper#1883</a></li>
<li>chore(deps): update crypt by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1884">spf13/viper#1884</a></li>
<li>chore: update dependencies by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1888">spf13/viper#1888</a></li>
<li>build(deps): bump github.com/go-viper/mapstructure/v2 from 2.0.0 to 2.1.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1901">spf13/viper#1901</a></li>
<li>build(deps): bump github.com/spf13/cast from 1.6.0 to 1.7.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1899">spf13/viper#1899</a></li>
<li>build(deps): bump github/codeql-action from 3.25.13 to 3.26.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1897">spf13/viper#1897</a></li>
<li>build(deps): bump golangci/golangci-lint-action from 6.0.1 to 6.1.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1893">spf13/viper#1893</a></li>
<li>build(deps): bump github/codeql-action from 3.26.0 to 3.26.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1903">spf13/viper#1903</a></li>
<li>build(deps): bump github/codeql-action from 3.26.2 to 3.26.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1905">spf13/viper#1905</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/spf13/viper/commit/c038295114ebabab9d57440f76025cfa73d7df3e"><code>c038295</code></a> docs: add update instructions for 1.20</li>
<li><a href="https://github.com/spf13/viper/commit/9c07e0f0633cae75a540fc606c6ebb3bece61826"><code>9c07e0f</code></a> build: disable unused linters</li>
<li><a href="https://github.com/spf13/viper/commit/48112d6a050e8a0af0a0aa2ae0d29834f603154b"><code>48112d6</code></a> ci: add Go 1.24 to the test matrix</li>
<li><a href="https://github.com/spf13/viper/commit/66e3e2877d2aa0b521a39f28f3340181e5603741"><code>66e3e28</code></a> build(deps): bump github.com/spf13/pflag from 1.0.5 to 1.0.6</li>
<li><a href="https://github.com/spf13/viper/commit/17b96ac0d54073bdbd05364a52195e6277c86b16"><code>17b96ac</code></a> New Logo</li>
<li><a href="https://github.com/spf13/viper/commit/8b223a45cef1badfd02317591a316095fb15a5d2"><code>8b223a4</code></a> build(deps): bump github.com/spf13/cast from 1.7.0 to 1.7.1</li>
<li><a href="https://github.com/spf13/viper/commit/91fd3639d7b10935d3e1308d350a7f727834561e"><code>91fd363</code></a> chore: update afero</li>
<li><a href="https://github.com/spf13/viper/commit/e75c48f185d3a75e447bd67f80fa153a88a9c958"><code>e75c48f</code></a> Fix issues reported by testifylint</li>
<li><a href="https://github.com/spf13/viper/commit/a5ea569d2e3cc9f6425530c9c731a0d84962a14d"><code>a5ea569</code></a> build(deps): bump github/codeql-action from 3.27.7 to 3.27.9</li>
<li><a href="https://github.com/spf13/viper/commit/54f2089833b65fa556a510957197de2609059147"><code>54f2089</code></a> build(deps): bump golang.org/x/crypto from 0.27.0 to 0.31.0 in /remote</li>
<li>Additional commits viewable in <a href="https://github.com/spf13/viper/compare/v1.19.0...v1.20.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.46.1 to 1.47.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a88ff68b5524ce9e4f816c8c313267e6d321ce8c"><code>a88ff68</code></a> Release 2023-11-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/023df359dd66f7b3970ddf2874c69d4b78b99126"><code>023df35</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/41f99ac19a3c3fc2e7e177b8f2632d319bb590e4"><code>41f99ac</code></a> Update SDK's smithy-go dependency to v1.18.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/72478fb87eae54a9d5ea84b336daad6795aebe01"><code>72478fb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4c4b852716f661378136f0e4d7e966d180057e26"><code>4c4b852</code></a> feat: add Options() getter to service clients (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2398">#2398</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6b86039a2b6015432b243b342b532f2db7a9c709"><code>6b86039</code></a> Update golang.org/x/net dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2391">#2391</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a955dd146a98d4b218893390a5c72512e0477c7"><code>0a955dd</code></a> Release 2023-11-28.3</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94a398007cfac2c60a62c203faa363a4ecaa4abf"><code>94a3980</code></a> fix: correct wiring of disable s3express auth toggle (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2394">#2394</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4aeeb0d7a4293f0b31c2e0be83e65da6f7fd4ae2"><code>4aeeb0d</code></a> Release 2023-11-28.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e09e153704d4da6ae2bb0ae3875058950d31206b"><code>e09e153</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.46.1...service/s3/v1.47.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.78.1 to 1.78.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d110cbf3144451a2aef547a785810eb3aa99bca"><code>6d110cb</code></a> Release 2025-03-11</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0a916fa998a5b9ad1ea3a985ad7f9b8aa31fcf37"><code>0a916fa</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5945692ff2787ac50fbce991fa3c475ea6c9ff35"><code>5945692</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e856ab0148dca14f452b39f687463f1cae3f1d9f"><code>e856ab0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7aec2edeb6109faed8be0462c11cef81e98bee53"><code>7aec2ed</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3034">#3034</a> from aws/feat-skip-nilbody-checksumvalidation</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b468d874d1653f5196cbce40e50a30ec38d36644"><code>b468d87</code></a> add check for non-exist object</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a56c483c323a5855016932fb4ff5ab1f7a62dac4"><code>a56c483</code></a> add changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59e14e3e748b51dc6ebab67c216a08c3b65c3f89"><code>59e14e3</code></a> only log checksum validation skipped when object is fetched</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d807b4686eefd0d6af69078f4b03861101fd3bcb"><code>d807b46</code></a> Release 2025-03-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/41c3e99e2afc15e58c5f6abe571ab393acb8fab4"><code>41c3e99</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.78.1...service/s3/v1.78.2">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


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

### Comment by @dependabot on March 18, 2025 at 03:58 PM UTC

Looks like these dependencies are no longer updatable, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1032*
