---
type: pull_request
number: 1546
title: "Build: Bump the go group with 11 updates"
state: merged
author: dependabot
created: 2026-06-22T04:38:43Z
updated: 2026-06-22T07:28:44Z
closed: 2026-06-22T07:28:42Z
merged: 2026-06-22T07:28:42Z
base_branch: main
head_branch: dependabot/go_modules/go-0fb3c01f0e
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1546
---

# Pull Request #1546: Build: Bump the go group with 11 updates

**Author**: @dependabot
**Created**: June 22, 2026 at 04:38 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-0fb3c01f0e`

## Description

Bumps the go group with 11 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.15.2` | `4.15.4` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.50.2` | `1.50.3` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.24` | `1.32.25` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.23` | `1.19.24` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.75.2` | `1.76.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.103.3` | `1.104.0` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.6.1780674492` | `2026.6.1781299727` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.46.2` | `0.47.0` |
| [github.com/getsentry/sentry-go/zerolog](https://github.com/getsentry/sentry-go) | `0.46.2` | `0.47.0` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.9.1` | `1.10.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.20.0` | `9.20.1` |

Updates `github.com/labstack/echo/v4` from 4.15.2 to 4.15.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.15.4</h2>
<p><strong>Security</strong></p>
<p>Fixes <a href="https://github.com/labstack/echo/security/advisories/GHSA-vfp3-v2gw-7wfq">GHSA-vfp3-v2gw-7wfq</a>: an encoded path separator (<code>%2F</code> or <code>%5C</code>) in a static file URL could bypass route-level middleware (e.g. authentication on a sibling route) and disclose static files. Both <code>StaticDirectoryHandler</code> (used by <code>Static</code>/<code>StaticFS</code>) and the <code>Static</code> middleware are affected. Backport of the v5 fix (<a href="https://redirect.github.com/labstack/echo/pull/3016">#3016</a>, released in v5.2.1). Thanks to <a href="https://github.com/a-tt-om"><code>@​a-tt-om</code></a> and <a href="https://github.com/oran-gugu"><code>@​oran-gugu</code></a> for reporting.</p>
<hr />
<p>Make serving static file releated methods  and middleware not unescape path by default - so how the way Router interprets paths and Static methods/middleware is consistent.</p>
<p>Given following situation:</p>
<pre lang="go"><code>// 0.
// given folder structure:
// private.txt
// public/
// public/index.html
// public/text.txt
// public/admin/private.txt
<p>// 1. share <code>public/</code> folder contents from the server root. This folder actually contains subfolder <code>admin</code> which<br />
// contents we want to forbid from downloading<br />
e.Static(&quot;/&quot;, &quot;public&quot;)</p>
<p>// 2. naively assume that everything under /admin folder is now forbidden<br />
e.GET(&quot;/admin/*&quot;, func(c *Context) error {<br />
return ErrForbidden<br />
})<br />
</code></pre></p>
<p>Then requests to <code>/admin%2fprivate.txt</code> would not be matched to <code>GET /admin/*</code> route (routing does not look unescaped path) and static file serving will use unescaped path to serve the file.</p>
<p>Note: this way of &quot;guarding&quot; subfolders will never work for for paths like <code>/assets/../admin%2fprivate.txt</code> which will <code>path.Clean(&quot;/assets/../admin%2fprivate.txt&quot;)</code> to <code>/admin/private.txt</code> and are servable if static file serving is configured to unescape paths.</p>
<p>If you want to guard routes - use middlewares on <code>Static*</code> methods and before <code>Static</code> middleware.</p>
<p><strong>Breaking change / migration:</strong> If you serve files whose names contain URL-encoded characters (e.g., <code>/hello%20world.txt</code> → <code>hello world.txt</code>), you must now opt in:</p>
<pre lang="go"><code>	e := echo.New()
	e.EnablePathUnescapingStaticFiles = true  // &lt;-- enable old behavior
	e.Static(&quot;/&quot;, &quot;public&quot;)
</code></pre>
<p>for static middleware</p>
<pre lang="go"><code>	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		EnablePathUnescaping: true, // &lt;-- enable old behavior
	}))
</code></pre>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/v4.15.4/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.15.4 - 2026-06-15</h2>
<p><strong>Security</strong></p>
<p>Fixes <a href="https://github.com/labstack/echo/security/advisories/GHSA-vfp3-v2gw-7wfq">GHSA-vfp3-v2gw-7wfq</a></p>
<p>Make serving static file releated methods  and middleware not unescape path by default - so how the way Router interprets paths and Static methods/middleware is consistent.</p>
<p>Given following situation:</p>
<pre lang="go"><code>// 0.
// given folder structure:
// private.txt
// public/
// public/index.html
// public/text.txt
// public/admin/private.txt
<p>// 1. share <code>public/</code> folder contents from the server root. This folder actually contains subfolder <code>admin</code> which<br />
// contents we want to forbid from downloading<br />
e.Static(&quot;/&quot;, &quot;public&quot;)</p>
<p>// 2. naively assume that everything under /admin folder is now forbidden<br />
e.GET(&quot;/admin/*&quot;, func(c *Context) error {<br />
return ErrForbidden<br />
})<br />
</code></pre></p>
<p>Then requests to <code>/admin%2fprivate.txt</code> would not be matched to <code>GET /admin/*</code> route (routing does not look unescaped path) and static file serving will use unescaped path to serve the file.</p>
<p>Note: this way of &quot;guarding&quot; subfolders will never work for for paths like <code>/assets/../admin%2fprivate.txt</code> which will <code>path.Clean(&quot;/assets/../admin%2fprivate.txt&quot;)</code> to <code>/admin/private.txt</code> and are servable if static file serving is configured to unescape paths.</p>
<p>If you want to guard routes - use middlewares on <code>Static*</code> methods and before <code>Static</code> middleware.</p>
<p><strong>Breaking change / migration:</strong> If you serve files whose names contain URL-encoded characters (e.g., <code>/hello%20world.txt</code> → <code>hello world.txt</code>), you must now opt in:</p>
<pre lang="go"><code>	e := echo.New()
	e.EnablePathUnescapingStaticFiles = true  // &lt;-- enable old behavior
	e.Static(&quot;/&quot;, &quot;public&quot;)
</code></pre>
<p>for static middleware</p>
<pre lang="go"><code>	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		EnablePathUnescaping: true, // &lt;-- enable old behavior
	}))
</code></pre>
<h2>v4.15.3 - 2026-06-14</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/ec79b584025d300acc9cb0a2b127a7209c902fd2"><code>ec79b58</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/3020">#3020</a> from aldas/v4_v4-15-4_changelog</li>
<li><a href="https://github.com/labstack/echo/commit/2714c07b793e5b3cd8b961cf6bb47ee59fafc368"><code>2714c07</code></a> Changelog for v4.15.4 - security fix</li>
<li><a href="https://github.com/labstack/echo/commit/13f0ed18cd5e08607bbed96b78cd477ad92380e2"><code>13f0ed1</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/3019">#3019</a> from aldas/v4_backport_3016</li>
<li><a href="https://github.com/labstack/echo/commit/d16a4ecf059440d2486809c9f06909547e8928ed"><code>d16a4ec</code></a> backport PR 3016 from v4</li>
<li><a href="https://github.com/labstack/echo/commit/8f167b9d45ace4f78ebd20efba9185389097f742"><code>8f167b9</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/3018">#3018</a> from aldas/v4_remove_v5_dep</li>
<li><a href="https://github.com/labstack/echo/commit/9afa4bae5e86c0846722e696aff1ba4abe8feed0"><code>9afa4ba</code></a> remove dependency on labstack/echo v5 introduced in go.mod and go.sum</li>
<li><a href="https://github.com/labstack/echo/commit/1e05f6351a8c9b7efc6b2d8d0c853c7559d70a43"><code>1e05f63</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/3017">#3017</a> from aldas/v4_ci_updates</li>
<li><a href="https://github.com/labstack/echo/commit/11a3cc46b94acee9a15be53f445aa358a07a2e7d"><code>11a3cc4</code></a> Update dependencies and add ignore for linting</li>
<li><a href="https://github.com/labstack/echo/commit/26bd016499257c16e824f29e8da075d77a3717ff"><code>26bd016</code></a> Update CI action versions</li>
<li><a href="https://github.com/labstack/echo/commit/aa52f6a5c78f791d89fe8979ffeb03bae1addd83"><code>aa52f6a</code></a> ci: run workflows on the v4 branch, not just master (<a href="https://redirect.github.com/labstack/echo/issues/3013">#3013</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/labstack/echo/compare/v4.15.2...v4.15.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.50.2 to 1.50.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.50.3 (2026-06-15)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: support Fetch v12 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3609">IBM/sarama#3609</a></li>
<li>feat: support CreateTopics v6 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3622">IBM/sarama#3622</a></li>
<li>feat: support DeleteTopics v5 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3624">IBM/sarama#3624</a></li>
<li>feat: support AddPartitionsToTxn v3 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3625">IBM/sarama#3625</a></li>
<li>feat: support Produce v9 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3629">IBM/sarama#3629</a></li>
<li>feat: support CreatePartitions v3 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3631">IBM/sarama#3631</a></li>
<li>feat: expose broker features from ApiVersions v3 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3633">IBM/sarama#3633</a></li>
<li>feat: support UpdateFeatures  by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3632">IBM/sarama#3632</a></li>
<li>feat: add OnAssignmentBalanceStrategy, the onAssignment half of the assignor contract by <a href="https://github.com/lizthegrey"><code>@​lizthegrey</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3627">IBM/sarama#3627</a></li>
<li>feat: support ApiVersions V4 and discover upgradable features in test by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3634">IBM/sarama#3634</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: avoid makeslice &quot;len out of range&quot; panics on invalid response by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3612">IBM/sarama#3612</a></li>
<li>fix: avoid makeslice panic in request decode too by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3613">IBM/sarama#3613</a></li>
<li>fix(test): wait for leaders after topic create by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3623">IBM/sarama#3623</a></li>
<li>fix(consumer): don't panic on requeue after unref by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3630">IBM/sarama#3630</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update golang-x to v0.53.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3618">IBM/sarama#3618</a></li>
<li>fix(deps): update module golang.org/x/sys to v0.46.0 - autoclosed by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3617">IBM/sarama#3617</a></li>
<li>fix(deps): update module golang.org/x/sync to v0.21.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3616">IBM/sarama#3616</a></li>
<li>fix(deps): update module golang.org/x/net to v0.56.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3626">IBM/sarama#3626</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: update default Kafka version to 2.6 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3608">IBM/sarama#3608</a></li>
<li>chore: more protocol version placeholders by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3610">IBM/sarama#3610</a></li>
<li>chore: extend fuzz testing to more types by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3614">IBM/sarama#3614</a></li>
<li>chore: clean up Int32 array encode/decode by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3619">IBM/sarama#3619</a></li>
<li>chore: apply Go's modernize fixes by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3621">IBM/sarama#3621</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.50.2...v1.50.3">https://github.com/IBM/sarama/compare/v1.50.2...v1.50.3</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/ff2eaba697d9974880b614b6043994ed3616d055"><code>ff2eaba</code></a> feat: support ApiVersions V4</li>
<li><a href="https://github.com/IBM/sarama/commit/0483979b44d3a4b74a02987389173c8ed8fe7532"><code>0483979</code></a> test: discover upgradable features via ApiVersions</li>
<li><a href="https://github.com/IBM/sarama/commit/d39f8ea185eda0df1c1ebca49e4eeafdf6f50240"><code>d39f8ea</code></a> feat: add OnAssignmentBalanceStrategy, the onAssignment half of the assignor ...</li>
<li><a href="https://github.com/IBM/sarama/commit/498dceb139ed9f421785473a039dd38f6d373f37"><code>498dceb</code></a> feat(admin): add ClusterAdmin UpdateFeatures</li>
<li><a href="https://github.com/IBM/sarama/commit/5f485029f0a2a076564f6c783acd62ed8074d6d3"><code>5f48502</code></a> feat: support UpdateFeatures v0</li>
<li><a href="https://github.com/IBM/sarama/commit/36c1c46cb0831ceeee080f579bc5769157afb458"><code>36c1c46</code></a> feat: expose broker features from ApiVersions v3 (<a href="https://redirect.github.com/IBM/sarama/issues/3633">#3633</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/b7d610681e9c016237f8104c6c2c817bc47df684"><code>b7d6106</code></a> feat: support CreatePartitions v3 (<a href="https://redirect.github.com/IBM/sarama/issues/3631">#3631</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/fdbaec72a260bd1c099ef5c80ded22953411cbe8"><code>fdbaec7</code></a> feat: support Produce v9 (<a href="https://redirect.github.com/IBM/sarama/issues/3629">#3629</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/277248f81decdd805092a1e260031b3d6e663c6f"><code>277248f</code></a> fix(consumer): don't panic on requeue after unref (<a href="https://redirect.github.com/IBM/sarama/issues/3630">#3630</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/ce73e81cd1a20e26362d7cb061f66c5237798191"><code>ce73e81</code></a> fix(deps): update module golang.org/x/net to v0.56.0</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.50.2...v1.50.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.24 to 1.32.25
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0016334f314179e3bb1d63ac7a5dbcb2ee7b3ee1"><code>0016334</code></a> Release 2026-06-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/396a18287497ccfa2cff73a10ae0dd946a167352"><code>396a182</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3eb85338c6dfc5721c5933253b2d66e5f8e96830"><code>3eb8533</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7bbea79fac21abb6cc7b037466074575e2138f34"><code>7bbea79</code></a> Release 2026-06-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1fefa6c28b6345014a2da360a6e47fc5bd5b7e72"><code>1fefa6c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2cf49ff138b824c9a5963c0f8f34429c86a199a"><code>a2cf49f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d87be6894e8451a757cc33d65bca734f51e256a9"><code>d87be68</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.24...config/v1.32.25">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.23 to 1.19.24
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0016334f314179e3bb1d63ac7a5dbcb2ee7b3ee1"><code>0016334</code></a> Release 2026-06-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/396a18287497ccfa2cff73a10ae0dd946a167352"><code>396a182</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3eb85338c6dfc5721c5933253b2d66e5f8e96830"><code>3eb8533</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7bbea79fac21abb6cc7b037466074575e2138f34"><code>7bbea79</code></a> Release 2026-06-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1fefa6c28b6345014a2da360a6e47fc5bd5b7e72"><code>1fefa6c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2cf49ff138b824c9a5963c0f8f34429c86a199a"><code>a2cf49f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d87be6894e8451a757cc33d65bca734f51e256a9"><code>d87be68</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.23...credentials/v1.19.24">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.75.2 to 1.76.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/061fd1b6d9940c6b37974e6aa64a5f9dc4ee1dc4"><code>061fd1b</code></a> Release 2025-02-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/92f835caa600f219455691ac912d89ce0aff0beb"><code>92f835c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56564d338547c28d596e453c30bee493f7e45845"><code>56564d3</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54678cf84cea516b77aae31bd1756283a24789e4"><code>54678cf</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1244c281833381d10002f3b03d6e67d6e1b686ad"><code>1244c28</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d4c5594ade58a82f1886cb43dcdd979619b5d6d3"><code>d4c5594</code></a> Fix for <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2383">#2383</a> (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3004">#3004</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78fa10aa9eaaa0851b0006145382ec0a0f4304c5"><code>78fa10a</code></a> Release 2025-02-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d65d975c21ababb3796264383b9b3f576257545"><code>2d65d97</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/64e842d4303abb468d7e07bcf1d92522c1081bae"><code>64e842d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f255e31c2ce564ae829a7f5a8caaec3887bbc0d1"><code>f255e31</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.75.2...service/s3/v1.76.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.103.3 to 1.104.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/99943aa792387847a5ee723320a98af18e5c3271"><code>99943aa</code></a> Release 2026-06-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/85941352ba17c55c1b715211b60bc39c8b3f94b6"><code>8594135</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/14151e5a04b45402df5df389db1490197e305ae9"><code>14151e5</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f43e77cd1b3bb74336f67e0e273dcf3fe183ff5c"><code>f43e77c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9fcc53abdf577092eb228458d78d0d2af8329c17"><code>9fcc53a</code></a> Release 2026-06-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65154c157b52b67a6029315c162f36d70ed9a99f"><code>65154c1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b86c7b982b4d6caec7b6a466468d2112f4fad1fd"><code>b86c7b9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/750e42a6493c92739684c3b99b99dba175bf6260"><code>750e42a</code></a> Release 2026-06-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f64ce2ceba5c9b7811df5f08f6a0ec0c5d23edaf"><code>f64ce2c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04e975554b5f9d665bb76157ec00bf7002763a09"><code>04e9755</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.103.3...service/s3/v1.104.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.6.1780674492 to 2026.6.1781299727
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/89d178ad381807a498a36cc9429d1f1433a4c4a7"><code>89d178a</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d5da53e55e057952eab83695...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.6.1780674492...release/v2026.6.1781299727">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.46.2 to 0.47.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.47.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Fix <code>transaction_info</code> source getting set incorrectly across HTTP middleware integrations (http, fasthttp, fiber). Users should now expect traces to properly get grouped with their parameterized path. Transactions in affected integrations may regroup after upgrading. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1325">#1325</a></li>
<li>remove deprecated<code>otel.NewSentrySpanProcessor</code>. Users should now use the <code>sentryotlp.NewTraceExporter</code> instead by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1307">#1307</a>
<pre lang="go"><code>// Before
sentry.Init(sentry.ClientOptions{Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0})
<p>tp := sdktrace.NewTracerProvider(
sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
)
otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
otel.SetTracerProvider(tp)</p>
<p>// After:
sentry.Init(sentry.ClientOptions{
Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0,
Integrations: func(i []sentry.Integration) []sentry.Integration {
return append(i, sentryotel.NewOtelIntegration())
},
})</p>
<p>exporter, _ := sentryotlp.NewTraceExporter(ctx, dsn)
tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
otel.SetTracerProvider(tp)
</code></pre></p>
</li>
<li>Enable logs by default to skip double allow behavior. Enabling logs now happens once when setting up either <code>sentry.NewLogger</code> or any supported integration. Also the EnableLogs flag changes to DisableLogs for a global override switch by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1306">#1306</a></li>
<li>Remove the <code>ContextifyFrames</code> integration. The recommended way to add source context is <a href="https://docs.sentry.io/integrations/source-code-mgmt/source-context/">SCM</a> by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1302">#1302</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add fiber v3 integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Bump fasthttp from 1.51.0 to 1.71.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Add sentrysql SQL tracing integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1305">#1305</a>
<ul>
<li>Supports multiple integration paths depending on how your app opens database connections: <code>sentrysql.Open(...)</code>, <code>sentrysql.OpenDB(...)</code>, and wrapped drivers/connectors for custom setups.</li>
<li>Database metadata is not inferred in every setup. If the database name is not discoverable automatically, pass <code>sentrysql.WithDatabaseName(...)</code> so spans are populated correctly.</li>
<li>Example:</li>
</ul>
<pre lang="go"><code> // Simple driver-based setup
 db, err := sentrysql.Open(&quot;sqlite&quot;, &quot;:memory:&quot;,
     sentrysql.WithDatabaseSystem(sentrysql.SystemSQLite),
     sentrysql.WithDatabaseName(&quot;main&quot;),
 )
</code></pre>
</li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.47.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Fix <code>transaction_info</code> source getting set incorrectly across HTTP middleware integrations (http, fasthttp, fiber). Users should now expect traces to properly get grouped with their parameterized path. Transactions in affected integrations may regroup after upgrading. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1325">#1325</a></li>
<li>remove deprecated<code>otel.NewSentrySpanProcessor</code>. Users should now use the <code>sentryotlp.NewTraceExporter</code> instead by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1307">#1307</a>
<pre lang="go"><code>// Before
sentry.Init(sentry.ClientOptions{Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0})
<p>tp := sdktrace.NewTracerProvider(
sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
)
otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
otel.SetTracerProvider(tp)</p>
<p>// After:
sentry.Init(sentry.ClientOptions{
Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0,
Integrations: func(i []sentry.Integration) []sentry.Integration {
return append(i, sentryotel.NewOtelIntegration())
},
})</p>
<p>exporter, _ := sentryotlp.NewTraceExporter(ctx, dsn)
tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
otel.SetTracerProvider(tp)
</code></pre></p>
</li>
<li>Enable logs by default to skip double allow behavior. Enabling logs now happens once when setting up either <code>sentry.NewLogger</code> or any supported integration. Also the EnableLogs flag changes to DisableLogs for a global override switch by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1306">#1306</a></li>
<li>Remove the <code>ContextifyFrames</code> integration. The recommended way to add source context is <a href="https://docs.sentry.io/integrations/source-code-mgmt/source-context/">SCM</a> by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1302">#1302</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add fiber v3 integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Bump fasthttp from 1.51.0 to 1.71.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Add sentrysql SQL tracing integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1305">#1305</a>
<ul>
<li>Supports multiple integration paths depending on how your app opens database connections: <code>sentrysql.Open(...)</code>, <code>sentrysql.OpenDB(...)</code>, and wrapped drivers/connectors for custom setups.</li>
<li>Database metadata is not inferred in every setup. If the database name is not discoverable automatically, pass <code>sentrysql.WithDatabaseName(...)</code> so spans are populated correctly.</li>
<li>Example:</li>
</ul>
<pre lang="go"><code> // Simple driver-based setup
 db, err := sentrysql.Open(&quot;sqlite&quot;, &quot;:memory:&quot;,
     sentrysql.WithDatabaseSystem(sentrysql.SystemSQLite),
     sentrysql.WithDatabaseName(&quot;main&quot;),
 )
</code></pre>
</li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/9b7a5624089638c2ba32feb0f19ada8d43a6cc45"><code>9b7a562</code></a> release: 0.47.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/320597ca56231b04d8cbecd7fb1638022a5f0f31"><code>320597c</code></a> chore: update bump-version script to also bump crosstest (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1327">#1327</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/02a836d7aa0d9714e7b043165aa8f576e051278c"><code>02a836d</code></a> build(deps): sync go.work (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1326">#1326</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/dc18868e4665622e8754b14693dfe74a03a0353d"><code>dc18868</code></a> feat: add fiberv3 integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1324">#1324</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f970fb569e96394f1bee44e24cdbcdb8c03f3b5e"><code>f970fb5</code></a> feat: add sql transaction instrumentation (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1305">#1305</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/1010c034417791bc4e8863f8d350684e64557a7b"><code>1010c03</code></a> fix!: transaction source for integrations (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1325">#1325</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/97b9ddf7adee2af74ce3b129d4eef204de7cb307"><code>97b9ddf</code></a> build(deps): bump getsentry/github-workflows from 71588ddf95134f804e82c5970a8...</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9cce79eaf28fec76f526628878f0819330bd85c0"><code>9cce79e</code></a> feat(sql): add lexer and obfuscator (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1299">#1299</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/063a5730fdb192c7fe3330f67a7e433841d8aeb9"><code>063a573</code></a> feat(sql): add span instrumentation (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1286">#1286</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/34d7db0afca784b8a73bc1b336ffa04163733c0d"><code>34d7db0</code></a> chore(otel): remove unused semconv helpers (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1321">#1321</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.46.2...v0.47.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go/zerolog` from 0.46.2 to 0.47.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go/zerolog's releases</a>.</em></p>
<blockquote>
<h2>0.47.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Fix <code>transaction_info</code> source getting set incorrectly across HTTP middleware integrations (http, fasthttp, fiber). Users should now expect traces to properly get grouped with their parameterized path. Transactions in affected integrations may regroup after upgrading. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1325">#1325</a></li>
<li>remove deprecated<code>otel.NewSentrySpanProcessor</code>. Users should now use the <code>sentryotlp.NewTraceExporter</code> instead by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1307">#1307</a>
<pre lang="go"><code>// Before
sentry.Init(sentry.ClientOptions{Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0})
<p>tp := sdktrace.NewTracerProvider(
sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
)
otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
otel.SetTracerProvider(tp)</p>
<p>// After:
sentry.Init(sentry.ClientOptions{
Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0,
Integrations: func(i []sentry.Integration) []sentry.Integration {
return append(i, sentryotel.NewOtelIntegration())
},
})</p>
<p>exporter, _ := sentryotlp.NewTraceExporter(ctx, dsn)
tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
otel.SetTracerProvider(tp)
</code></pre></p>
</li>
<li>Enable logs by default to skip double allow behavior. Enabling logs now happens once when setting up either <code>sentry.NewLogger</code> or any supported integration. Also the EnableLogs flag changes to DisableLogs for a global override switch by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1306">#1306</a></li>
<li>Remove the <code>ContextifyFrames</code> integration. The recommended way to add source context is <a href="https://docs.sentry.io/integrations/source-code-mgmt/source-context/">SCM</a> by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1302">#1302</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add fiber v3 integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Bump fasthttp from 1.51.0 to 1.71.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Add sentrysql SQL tracing integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1305">#1305</a>
<ul>
<li>Supports multiple integration paths depending on how your app opens database connections: <code>sentrysql.Open(...)</code>, <code>sentrysql.OpenDB(...)</code>, and wrapped drivers/connectors for custom setups.</li>
<li>Database metadata is not inferred in every setup. If the database name is not discoverable automatically, pass <code>sentrysql.WithDatabaseName(...)</code> so spans are populated correctly.</li>
<li>Example:</li>
</ul>
<pre lang="go"><code> // Simple driver-based setup
 db, err := sentrysql.Open(&quot;sqlite&quot;, &quot;:memory:&quot;,
     sentrysql.WithDatabaseSystem(sentrysql.SystemSQLite),
     sentrysql.WithDatabaseName(&quot;main&quot;),
 )
</code></pre>
</li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go/zerolog's changelog</a>.</em></p>
<blockquote>
<h2>0.47.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Fix <code>transaction_info</code> source getting set incorrectly across HTTP middleware integrations (http, fasthttp, fiber). Users should now expect traces to properly get grouped with their parameterized path. Transactions in affected integrations may regroup after upgrading. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1325">#1325</a></li>
<li>remove deprecated<code>otel.NewSentrySpanProcessor</code>. Users should now use the <code>sentryotlp.NewTraceExporter</code> instead by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1307">#1307</a>
<pre lang="go"><code>// Before
sentry.Init(sentry.ClientOptions{Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0})
<p>tp := sdktrace.NewTracerProvider(
sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
)
otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
otel.SetTracerProvider(tp)</p>
<p>// After:
sentry.Init(sentry.ClientOptions{
Dsn: dsn, EnableTracing: true, TracesSampleRate: 1.0,
Integrations: func(i []sentry.Integration) []sentry.Integration {
return append(i, sentryotel.NewOtelIntegration())
},
})</p>
<p>exporter, _ := sentryotlp.NewTraceExporter(ctx, dsn)
tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
otel.SetTracerProvider(tp)
</code></pre></p>
</li>
<li>Enable logs by default to skip double allow behavior. Enabling logs now happens once when setting up either <code>sentry.NewLogger</code> or any supported integration. Also the EnableLogs flag changes to DisableLogs for a global override switch by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1306">#1306</a></li>
<li>Remove the <code>ContextifyFrames</code> integration. The recommended way to add source context is <a href="https://docs.sentry.io/integrations/source-code-mgmt/source-context/">SCM</a> by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1302">#1302</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add fiber v3 integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Bump fasthttp from 1.51.0 to 1.71.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1324">#1324</a></li>
<li>Add sentrysql SQL tracing integration by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1305">#1305</a>
<ul>
<li>Supports multiple integration paths depending on how your app opens database connections: <code>sentrysql.Open(...)</code>, <code>sentrysql.OpenDB(...)</code>, and wrapped drivers/connectors for custom setups.</li>
<li>Database metadata is not inferred in every setup. If the database name is not discoverable automatically, pass <code>sentrysql.WithDatabaseName(...)</code> so spans are populated correctly.</li>
<li>Example:</li>
</ul>
<pre lang="go"><code> // Simple driver-based setup
 db, err := sentrysql.Open(&quot;sqlite&quot;, &quot;:memory:&quot;,
     sentrysql.WithDatabaseSystem(sentrysql.SystemSQLite),
     sentrysql.WithDatabaseName(&quot;main&quot;),
 )
</code></pre>
</li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/9b7a5624089638c2ba32feb0f19ada8d43a6cc45"><code>9b7a562</code></a> release: 0.47.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/320597ca56231b04d8cbecd7fb1638022a5f0f31"><code>320597c</code></a> chore: update bump-version script to also bump crosstest (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1327">#1327</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/02a836d7aa0d9714e7b043165aa8f576e051278c"><code>02a836d</code></a> build(deps): sync go.work (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1326">#1326</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/dc18868e4665622e8754b14693dfe74a03a0353d"><code>dc18868</code></a> feat: add fiberv3 integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1324">#1324</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f970fb569e96394f1bee44e24cdbcdb8c03f3b5e"><code>f970fb5</code></a> feat: add sql transaction instrumentation (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1305">#1305</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/1010c034417791bc4e8863f8d350684e64557a7b"><code>1010c03</code></a> fix!: transaction source for integrations (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1325">#1325</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/97b9ddf7adee2af74ce3b129d4eef204de7cb307"><code>97b9ddf</code></a> build(deps): bump getsentry/github-workflows from 71588ddf95134f804e82c5970a8...</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9cce79eaf28fec76f526628878f0819330bd85c0"><code>9cce79e</code></a> feat(sql): add lexer and obfuscator (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1299">#1299</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/063a5730fdb192c7fe3330f67a7e433841d8aeb9"><code>063a573</code></a> feat(sql): add span instrumentation (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1286">#1286</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/34d7db0afca784b8a73bc1b336ffa04163733c0d"><code>34d7db0</code></a> chore(otel): remove unused semconv helpers (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1321">#1321</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.46.2...v0.47.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.9.1 to 1.10.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.10.0</h2>
<h2>What's Changed</h2>
<ul>
<li>RHCLOUD-47069 - Update documentation around how list_workspaces handles pagination by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/55">project-kessel/kessel-sdk-go#55</a></li>
<li>RHCLOUD-48468 - Add consistency token support to ListWorkspaces() by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/62">project-kessel/kessel-sdk-go#62</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.1...v1.10.0">https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.1...v1.10.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/8622ce56988ccffe829e4ca829037c33ff0c0197"><code>8622ce5</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/62">#62</a> from lennysgarage/RHCLOUD-48468</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/a06717dd7a9f560495aa298978e0f3bae319f60d"><code>a06717d</code></a> Add listWorkspacesOption param to listWorkspaces</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/7a102b49a29444661d2dea4c343ed62d1828f9de"><code>7a102b4</code></a> update godoc to reflect new # of params</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/e5cece2b0e1a91fc17b409b17dd656b2b40c67ef"><code>e5cece2</code></a> Needed to merge in latest changes &amp; update func</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/e4f9f3df975cca6cfcb66230dd395a81849ac576"><code>e4f9f3d</code></a> Merge branch 'main' into RHCLOUD-48468</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/f6640dc492553a92600ce85828a6e34fa039efc4"><code>f6640dc</code></a> Update example so lint runs properly</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/763c42dcf92da25fa88fc1ec69e563796f8eea4c"><code>763c42d</code></a> Revert &quot;why is this not working&quot;</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/ca3fcc0c76acecb17dbb997a14e489bc89aaffeb"><code>ca3fcc0</code></a> why is this not working</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/a43fad9907ec70b9174559b722cfd0f7fb8c8f3f"><code>a43fad9</code></a> Update func definition</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/caaa71dec3ba0c338bf5ae316da11e2755d69053"><code>caaa71d</code></a> lint all example files</li>
<li>Additional commits viewable in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.1...v1.10.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.20.0 to 9.20.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.20.1</h2>
<p>This is a patch release containing bug fixes only. There are no new features or breaking changes; upgrading from 9.20.0 is a drop-in replacement.</p>
<h2>🚀 Highlights</h2>
<h3>RESP3 pub/sub message loss fixed</h3>
<p><code>PeekPushNotificationName</code> previously inspected only the bytes already buffered by <code>bufio</code>, so when a push frame header straddled a buffer fill boundary it could return a <strong>truncated</strong> notification name (e.g. <code>&quot;messa&quot;</code> instead of <code>&quot;message&quot;</code>). The push processor then mis-routed the frame and <code>ReadReply</code> silently dropped it, causing intermittent RESP3 pub/sub message loss. The peek now grows its window (36 bytes → up to 4 KiB) and reads more from the connection until the header is complete, cleanly separating incomplete prefixes from corrupt frames (including overflow-safe bulk-length handling). Fixes <a href="https://redirect.github.com/redis/go-redis/issues/3839">#3839</a>.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3842">#3842</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>RESP3 push peeking</strong>: <code>PeekPushNotificationName</code> no longer returns a truncated notification name when a push frame header spans a buffer boundary, preventing silent RESP3 pub/sub message loss (fixes <a href="https://redirect.github.com/redis/go-redis/issues/3839">#3839</a>) (<a href="https://redirect.github.com/redis/go-redis/pull/3842">#3842</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong><code>FT.HYBRID</code> vector params</strong>: Vector data is now always sent via <code>PARAMS</code> with auto-generated param names (<code>__vector_param_N</code>, with collision avoidance) when <code>VectorParamName</code> is omitted, since Redis no longer accepts inline vector blobs; the <code>FTHybridOptions.Params</code> map is no longer mutated, so the same options struct can be reused across calls (<a href="https://redirect.github.com/redis/go-redis/pull/3844">#3844</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong><code>CLUSTER SHARDS</code> forward compatibility</strong>: Unknown shard- and node-level attributes in the <code>CLUSTER SHARDS</code> reply are now skipped via <code>DiscardNext()</code> instead of erroring, so clients keep working when the server introduces new fields (<a href="https://redirect.github.com/redis/go-redis/pull/3843">#3843</a>) by <a href="https://github.com/madolson"><code>@​madolson</code></a></li>
<li><strong>PubSub double reconnect</strong>: <code>PubSub.releaseConn</code> no longer reconnects twice when a connection is both unusable (or pending handoff) and reports a bad-connection error, avoiding a wasted connection establish-then-close cycle (<a href="https://redirect.github.com/redis/go-redis/pull/3833">#3833</a>) by <a href="https://github.com/cxljs"><code>@​cxljs</code></a></li>
</ul>
<h2>👥 Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/madolson"><code>@​madolson</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<hr />
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.20.0...v9.20.1">https://github.com/redis/go-redis/compare/v9.20.0...v9.20.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.20.1 (2026-06-11)</h1>
<p>This is a patch release containing bug fixes only. There are no new features or breaking changes; upgrading from 9.20.0 is a drop-in replacement.</p>
<h2>🚀 Highlights</h2>
<h3>RESP3 pub/sub message loss fixed</h3>
<p><code>PeekPushNotificationName</code> previously inspected only the bytes already buffered by <code>bufio</code>, so when a push frame header straddled a buffer fill boundary it could return a <strong>truncated</strong> notification name (e.g. <code>&quot;messa&quot;</code> instead of <code>&quot;message&quot;</code>). The push processor then mis-routed the frame and <code>ReadReply</code> silently dropped it, causing intermittent RESP3 pub/sub message loss. The peek now grows its window (36 bytes → up to 4 KiB) and reads more from the connection until the header is complete, cleanly separating incomplete prefixes from corrupt frames (including overflow-safe bulk-length handling). Fixes <a href="https://redirect.github.com/redis/go-redis/issues/3839">#3839</a>.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3842">#3842</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>RESP3 push peeking</strong>: <code>PeekPushNotificationName</code> no longer returns a truncated notification name when a push frame header spans a buffer boundary, preventing silent RESP3 pub/sub message loss (fixes <a href="https://redirect.github.com/redis/go-redis/issues/3839">#3839</a>) (<a href="https://redirect.github.com/redis/go-redis/pull/3842">#3842</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong><code>FT.HYBRID</code> vector params</strong>: Vector data is now always sent via <code>PARAMS</code> with auto-generated param names (<code>__vector_param_N</code>, with collision avoidance) when <code>VectorParamName</code> is omitted, since Redis no longer accepts inline vector blobs; the <code>FTHybridOptions.Params</code> map is no longer mutated, so the same options struct can be reused across calls (<a href="https://redirect.github.com/redis/go-redis/pull/3844">#3844</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong><code>CLUSTER SHARDS</code> forward compatibility</strong>: Unknown shard- and node-level attributes in the <code>CLUSTER SHARDS</code> reply are now skipped via <code>DiscardNext()</code> instead of erroring, so clients keep working when the server introduces new fields (<a href="https://redirect.github.com/redis/go-redis/pull/3843">#3843</a>) by <a href="https://github.com/madolson"><code>@​madolson</code></a></li>
<li><strong>PubSub double reconnect</strong>: <code>PubSub.releaseConn</code> no longer reconnects twice when a connection is both unusable (or pending handoff) and reports a bad-connection error, avoiding a wasted connection establish-then-close cycle (<a href="https://redirect.github.com/redis/go-redis/pull/3833">#3833</a>) by <a href="https://github.com/cxljs"><code>@​cxljs</code></a></li>
</ul>
<h2>👥 Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/madolson"><code>@​madolson</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<hr />
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.20.0...v9.20.1">https://github.com/redis/go-redis/compare/v9.20.0...v9.20.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/a13416bf1c0556075b843dd0a9cb6db3b3cf2789"><code>a13416b</code></a> chore(release): 9.20.1 (<a href="https://redirect.github.com/redis/go-redis/issues/3847">#3847</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/10dc44f424d58b9c2b4df29e529ca92a2f8fa986"><code>10dc44f</code></a> fix(push): fix peeking when push name is truncated (<a href="https://redirect.github.com/redis/go-redis/issues/3842">#3842</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e1a2d68b27e70ca8f3ff5eec61d0dd2b74554670"><code>e1a2d68</code></a> fix(ft.hybrid): Always generate vector param names if they are not provided b...</li>
<li><a href="https://github.com/redis/go-redis/commit/a4b234f4a4f2551d1179181bbc136db132bf5912"><code>a4b234f</code></a> chore(deps): bump codecov/codecov-action from 6 to 7 (<a href="https://redirect.github.com/redis/go-redis/issues/3845">#3845</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/974e717dbf3008497e54aec34636e34452e6872b"><code>974e717</code></a> fix(command): ignore unknown fields in CLUSTER SHARDS response (<a href="https://redirect.github.com/redis/go-redis/issues/3843">#3843</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/65d6abde0d812234d26763930b5e84ff05963c96"><code>65d6abd</code></a> fix(pubsub): prevent double reconnect in releaseConn (<a href="https://redirect.github.com/redis/go-redis/issues/3833">#3833</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.20.0...v9.20.1">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on June 22, 2026 at 07:28 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1546*
