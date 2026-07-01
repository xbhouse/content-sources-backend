---
type: pull_request
number: 1336
title: "Build: Bump the go group with 7 updates"
state: closed
author: dependabot
created: 2025-12-15T04:09:53Z
updated: 2025-12-18T20:37:18Z
closed: 2025-12-18T20:37:16Z
base_branch: main
head_branch: dependabot/go_modules/go-f88383a8ed
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1336
---

# Pull Request #1336: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: December 15, 2025 at 04:09 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-f88383a8ed`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.13.4` | `4.14.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.40.0` | `1.41.0` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.2` | `1.32.5` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.2` | `1.19.5` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.61.1` | `1.62.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.92.1` | `1.93.2` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.11.1764189739` | `2025.12.1765576454` |

Updates `github.com/labstack/echo/v4` from 4.13.4 to 4.14.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.14.0</h2>
<p><code>middleware.Logger()</code> has been deprecated. For request logging, use <code>middleware.RequestLogger()</code> or
<code>middleware.RequestLoggerWithConfig()</code>.</p>
<p><code>middleware.RequestLogger()</code> replaces <code>middleware.Logger()</code>, offering comparable configuration while relying on the
Go standard library’s new <code>slog</code> logger.</p>
<p>The previous default output format was JSON. The new default follows the standard <code>slog</code> logger settings.
To continue emitting request logs in JSON, configure <code>slog</code> accordingly:</p>
<pre lang="go"><code>slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
e.Use(middleware.RequestLogger())
</code></pre>
<p>If you are developing anything more substantial than a demo, use <code>middleware.RequestLoggerWithConfig()</code></p>
<p><strong>Security</strong></p>
<ul>
<li>Logger middleware json string escaping and deprecation by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2849">labstack/echo#2849</a></li>
</ul>
<h2>What's Changed</h2>
<ul>
<li>Update deps  by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2807">labstack/echo#2807</a></li>
<li>refactor to use reflect.TypeFor by <a href="https://github.com/cuiweixie"><code>@​cuiweixie</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2812">labstack/echo#2812</a></li>
<li>Use Go 1.25 in CI by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2810">labstack/echo#2810</a></li>
<li>Modernize context.go by replacing interface{} with any by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2822">labstack/echo#2822</a></li>
<li>Fix typo in SetParamValues comment by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2828">labstack/echo#2828</a></li>
<li>Fix typo in ContextTimeout middleware comment by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2827">labstack/echo#2827</a></li>
<li>Improve BasicAuth middleware: use strings.Cut and RFC compliance by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2825">labstack/echo#2825</a></li>
<li>Fix duplicate plus operator in router backtracking logic by <a href="https://github.com/yuya-morimoto"><code>@​yuya-morimoto</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2832">labstack/echo#2832</a></li>
<li>Replace custom private IP range check with built-in net.IP.IsPrivate by <a href="https://github.com/kumapower17"><code>@​kumapower17</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2835">labstack/echo#2835</a></li>
<li>Ensure proxy connection is closed in proxyRaw function(<a href="https://redirect.github.com/labstack/echo/issues/2837">#2837</a>) by <a href="https://github.com/kumapower17"><code>@​kumapower17</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2838">labstack/echo#2838</a></li>
<li>Update deps by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2843">labstack/echo#2843</a></li>
<li>Logger middleware json string escaping and deprecation by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2849">labstack/echo#2849</a></li>
<li>Update golang.org/x/* deps by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2850">labstack/echo#2850</a></li>
<li>Changelog for 4.14.0 by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2851">labstack/echo#2851</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/cuiweixie"><code>@​cuiweixie</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2812">labstack/echo#2812</a></li>
<li><a href="https://github.com/yuya-morimoto"><code>@​yuya-morimoto</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2832">labstack/echo#2832</a></li>
<li><a href="https://github.com/kumapower17"><code>@​kumapower17</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2835">labstack/echo#2835</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.13.4...v4.14.0">https://github.com/labstack/echo/compare/v4.13.4...v4.14.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.14.0 - 2025-12-11</h2>
<p><code>middleware.Logger</code> has been deprecated. For request logging, use <code>middleware.RequestLogger</code> or
<code>middleware.RequestLoggerWithConfig</code>.</p>
<p><code>middleware.RequestLogger</code> replaces <code>middleware.Logger</code>, offering comparable configuration while relying on the
Go standard library’s new <code>slog</code> logger.</p>
<p>The previous default output format was JSON. The new default follows the standard <code>slog</code> logger settings.
To continue emitting request logs in JSON, configure <code>slog</code> accordingly:</p>
<pre lang="go"><code>slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
e.Use(middleware.RequestLogger())
</code></pre>
<p><strong>Security</strong></p>
<ul>
<li>Logger middleware json string escaping and deprecation by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2849">labstack/echo#2849</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Update deps  by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2807">labstack/echo#2807</a></li>
<li>refactor to use reflect.TypeFor by <a href="https://github.com/cuiweixie"><code>@​cuiweixie</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2812">labstack/echo#2812</a></li>
<li>Use Go 1.25 in CI by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2810">labstack/echo#2810</a></li>
<li>Modernize context.go by replacing interface{} with any by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2822">labstack/echo#2822</a></li>
<li>Fix typo in SetParamValues comment by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2828">labstack/echo#2828</a></li>
<li>Fix typo in ContextTimeout middleware comment by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2827">labstack/echo#2827</a></li>
<li>Improve BasicAuth middleware: use strings.Cut and RFC compliance by <a href="https://github.com/vishr"><code>@​vishr</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2825">labstack/echo#2825</a></li>
<li>Fix duplicate plus operator in router backtracking logic by <a href="https://github.com/yuya-morimoto"><code>@​yuya-morimoto</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2832">labstack/echo#2832</a></li>
<li>Replace custom private IP range check with built-in net.IP.IsPrivate by <a href="https://github.com/kumapower17"><code>@​kumapower17</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2835">labstack/echo#2835</a></li>
<li>Ensure proxy connection is closed in proxyRaw function(<a href="https://redirect.github.com/labstack/echo/issues/2837">#2837</a>) by <a href="https://github.com/kumapower17"><code>@​kumapower17</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2838">labstack/echo#2838</a></li>
<li>Update deps by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2843">labstack/echo#2843</a></li>
<li>Update golang.org/x/* deps by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2850">labstack/echo#2850</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/6392cb459842d2c1747902ec2a1809c1387df5d8"><code>6392cb4</code></a> Changelog for 4.14.0</li>
<li><a href="https://github.com/labstack/echo/commit/c9bd2cd8e32d07c2d445ff07300338bf5a28362f"><code>c9bd2cd</code></a> Update golang.org/x/* deps (<a href="https://redirect.github.com/labstack/echo/issues/2850">#2850</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/c12cb08a8679d45dbefe31ab604e60c9ebb64c3c"><code>c12cb08</code></a> Logger middleware json string escaping and deprecation (<a href="https://redirect.github.com/labstack/echo/issues/2849">#2849</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/612967a9fec11b112a16c7b62efc2344eae308e8"><code>612967a</code></a> Update deps</li>
<li><a href="https://github.com/labstack/echo/commit/53b692c4d4de6306d5306d2c81b8335e71ddc016"><code>53b692c</code></a> Ensure proxy connection is closed in proxyRaw function</li>
<li><a href="https://github.com/labstack/echo/commit/e644ff8f7bb01c694cacec3ad22a7471609ea106"><code>e644ff8</code></a> Replace custom private IP range check with built-in net.IP.IsPrivate() method</li>
<li><a href="https://github.com/labstack/echo/commit/40e2e8faf95226d541ecce3ca27def3cb9c7f592"><code>40e2e8f</code></a> Fix typo &quot;+&quot;</li>
<li><a href="https://github.com/labstack/echo/commit/55cb3b625d1228827fa35a3cfc4dd15b3a3a406b"><code>55cb3b6</code></a> Optimize realm quoting to happen once during middleware creation</li>
<li><a href="https://github.com/labstack/echo/commit/dbd583fa4d9e1b327f06cfe9abc97b044b273289"><code>dbd583f</code></a> Add comprehensive tests for realm quoting behavior</li>
<li><a href="https://github.com/labstack/echo/commit/432a2adf46fe74403b7df3a81aca7583da02c054"><code>432a2ad</code></a> Improve BasicAuth middleware: use strings.Cut and RFC compliance</li>
<li>Additional commits viewable in <a href="https://github.com/labstack/echo/compare/v4.13.4...v4.14.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.40.0 to 1.41.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2c5c1de983d67074b607e5e410102ab2821a06e0"><code>2c5c1de</code></a> Release 2025-12-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a968e75f31355127d6d21b5c1eeeef56c7293528"><code>a968e75</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b110f673443a879368b09088ddb3a41b42289810"><code>b110f67</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/82760c3f13f663592956106b7fc61e4f1fc1ce31"><code>82760c3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9e038d3c6c3dff3ed4d7da6c9ca8781bce3f4f69"><code>9e038d3</code></a> [merge 12/8/25] default MaxConnsPerHost to 2048 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3245">#3245</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55b9a85b5962a5529d4383caabb7c7c8cf2e34df"><code>55b9a85</code></a> Release 2025-12-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb36f84a8640782421ce92ac04946ee270cbae98"><code>cb36f84</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/081bac07f1af24e5f7cd1028d120ef41cd1c2e42"><code>081bac0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/25f26c231f47889249f6bf9b6b11ed817097abdf"><code>25f26c2</code></a> Release 2025-12-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aadf9d36831914dd70472d25ba23bf9160304d67"><code>aadf9d3</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.40.0...v1.41.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.2 to 1.32.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d125de3792b20980da07dd1424afa90285d50c90"><code>d125de3</code></a> Release 2024-11-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fec51f3fff1b0f1a2cc913f4f7874977a3d47d0d"><code>fec51f3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fba59970453e163b476913a363a126412641b5fb"><code>fba5997</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b8e5c842fc161b15aa5e13c8a0caf1357b5e9c7"><code>0b8e5c8</code></a> Bump smithy-go dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2902">#2902</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50ba45ce162cd2458c65f5da799fd907ad826561"><code>50ba45c</code></a> Release 2024-11-15.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/444bdffccd6dce18f60ae626b74c087641c3d052"><code>444bdff</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55ab381b20235964ab0c670a29d096821e158e90"><code>55ab381</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94c083768b80bbd372c0e9feb45f02511442b896"><code>94c0837</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2398a7903ca141692f98da65b8537a4a53e9707e"><code>2398a79</code></a> Remove elastictranscoder service's integration test (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2901">#2901</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93e0f294f6c692e39adf1b8ec2c8681ba9ee5f01"><code>93e0f29</code></a> Release 2024-11-15</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.2...v1.32.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.2 to 1.19.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/676a8b1bf0174c8763e19d99b68b988e67e2d398"><code>676a8b1</code></a> Release 2025-01-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1737386a85235b72e9676ed261b72cddb61355df"><code>1737386</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3bc09da29fb3dd079526f7ed141520f69245e445"><code>3bc09da</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb98deef60318ce9a61cda159ebfb0166d88539b"><code>cb98dee</code></a> Fix flex checksum validation cfg (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2981">#2981</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9c764018fe28b27912a0b976614d9e806e3f8268"><code>9c76401</code></a> fix bad changelog type</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ed8a3caa0df9ce36a5b60aebeee201187098d205"><code>ed8a3ca</code></a> Reduce fmt.Sprintf allocations in query encoding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2919">#2919</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5773a9a070873393eb2e7eed37bd647e12e1267"><code>d5773a9</code></a> Add FixUnmarshalIndividualSetValues option to DecoderOptions of dynamodb (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2896">#2896</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58e23dc0311cec940749e34ddfc542dbb00ff7a3"><code>58e23dc</code></a> fix codegen test failing in main</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/640d919419375c1bb9041ffa6dd024b60243a1ed"><code>640d919</code></a> fix broken jmespath waiters in cloudwatch and autoscaling (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2984">#2984</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/613a6cfc607af8470ceec5b7391f9231fa1f98dd"><code>613a6cf</code></a> Optimize/directory traversal (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2970">#2970</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/m2/v1.19.2...service/m2/v1.19.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.61.1 to 1.62.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4bd7f5481eebe1c422fa85d1956f7ea34d93cf76"><code>4bd7f54</code></a> Release 2025-10-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55fb47cb07949ca70312a359272b10ff29f520df"><code>55fb47c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bf727c0b40333a2e643a89909f189ffba0c212b9"><code>bf727c0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ca943fb071778b06e5e7f64f5ddf896f8579b6b"><code>0ca943f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3908bc4e960b7782da16a710fb8a747632af25a5"><code>3908bc4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0c74d28b606e71c1ee23d7ee17ec4949001cf56"><code>a0c74d2</code></a> Release 2025-10-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb611540d594538970ddd91b7802f5152cca8d2f"><code>eb61154</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e20d3e4b6065b1a8cdd97d00342cdc8d8c02561b"><code>e20d3e4</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9acd5faf2ca161219176c3cd8402837c98ed82ce"><code>9acd5fa</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a91cc6c72206a1cef64caff0c70fa1bd13bd4543"><code>a91cc6c</code></a> Speed up unit tests by removing duplicate work (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3216">#3216</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.61.1...service/fsx/v1.62.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.92.1 to 1.93.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8878ab2c86f54b5b35ae7c6b41b47f7c2ab1a128"><code>8878ab2</code></a> Release 2025-12-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3bb7070f414a0d3b0bdf3203af5a88d8c62141cb"><code>3bb7070</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67298ad1e66d5d6c9fc284a45b2add8f6a2a76f1"><code>67298ad</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b57f9c7a9e62e3aea3c32499eec7b6b64bcfd102"><code>b57f9c7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18eff735fe4bb2d88a431b2821d54ed83a6bbe63"><code>18eff73</code></a> upgrade to smithy v1.64.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3246">#3246</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2c5c1de983d67074b607e5e410102ab2821a06e0"><code>2c5c1de</code></a> Release 2025-12-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a968e75f31355127d6d21b5c1eeeef56c7293528"><code>a968e75</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b110f673443a879368b09088ddb3a41b42289810"><code>b110f67</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/82760c3f13f663592956106b7fc61e4f1fc1ce31"><code>82760c3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9e038d3c6c3dff3ed4d7da6c9ca8781bce3f4f69"><code>9e038d3</code></a> [merge 12/8/25] default MaxConnsPerHost to 2048 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3245">#3245</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.92.1...service/s3/v1.93.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.11.1764189739 to 2025.12.1765576454
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/f80c1674ad496963af9cb0b170f2a7121acc700b"><code>f80c167</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4783b666848af67968b5e2da9a...</li>
<li><a href="https://github.com/content-services/zest/commit/1c50f2a702f4e46a8822fc37bfbd150c4b02b43d"><code>1c50f2a</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4783b656d6b8f67968b5e2da9a...</li>
<li><a href="https://github.com/content-services/zest/commit/54de229dfaca56ddd2fcc0b0434a6e54cb0d9047"><code>54de229</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73d39ae994e137d2e4353b86e...</li>
<li><a href="https://github.com/content-services/zest/commit/6c8ff7901be4b31b14360251db21722d959ffe48"><code>6c8ff79</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a9728652e339a157bd286834ede...</li>
<li><a href="https://github.com/content-services/zest/commit/9edc85adc6399e36f30aa918b279a46ebd79db30"><code>9edc85a</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d5b4b459eec57952eab83695...</li>
<li><a href="https://github.com/content-services/zest/commit/05b3aff04b0a07ed9d50c4c2b746afde50048f55"><code>05b3aff</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73d35ae6888037d2e4353b86e...</li>
<li><a href="https://github.com/content-services/zest/commit/03b84cce6d7ea170f1da64b28dc68eaf31ac449a"><code>03b84cc</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e439583384c87e8639db952b...</li>
<li><a href="https://github.com/content-services/zest/commit/08aa5f84ad014baf0eeb1b0c6a6a32501a3c7783"><code>08aa5f8</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e475439452e9ecb7a98d596b8e3...</li>
<li><a href="https://github.com/content-services/zest/commit/933b856c2ba5878cde20be56a8e0ced81ea23f7c"><code>933b856</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73d3e4aa344c37d2e4353b86e...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.11.1764189739...release/v2025.12.1765576454">compare view</a></li>
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

### Comment by @dependabot on December 18, 2025 at 08:37 PM UTC

Looks like these dependencies are no longer updatable, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1336*
