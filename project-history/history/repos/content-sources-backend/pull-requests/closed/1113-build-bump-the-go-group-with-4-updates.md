---
type: pull_request
number: 1113
title: "Build: Bump the go group with 4 updates"
state: closed
author: dependabot
created: 2025-05-19T04:19:09Z
updated: 2025-05-26T04:06:50Z
closed: 2025-05-26T04:06:48Z
base_branch: main
head_branch: dependabot/go_modules/go-16df04b9f6
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1113
---

# Pull Request #1113: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: May 19, 2025 at 04:19 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-16df04b9f6`

## Description

Bumps the go group with 4 updates: [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2), [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest), [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) and [github.com/jackc/pgx/v5](https://github.com/jackc/pgx).

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.49.0 to 1.50.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4334b43c27eeb4c4ecfcba1a87cc08f963fe91a3"><code>4334b43</code></a> Release 2024-02-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9e29187dc08c2e9bf8f66166841ca9c75e0624ab"><code>9e29187</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f672c49fb7b4199becd3ba0a6e9d2b35f215e985"><code>f672c49</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f935782a0d8ee350736d4515cc05001e4ac3d01"><code>7f93578</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/da7bcdacddcf05ee9c897137fdbcdc8caf501698"><code>da7bcda</code></a> feat: add client config passthrough to waiter opts (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2499">#2499</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1bb048272ad54b3cbeb3b6da99f4be8090bea5d2"><code>1bb0482</code></a> Release 2024-02-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cc83a2b05532dfe5e6ccba25881887352134fbbd"><code>cc83a2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f44d9967b75499363425c29a5cf40375e9869a7"><code>7f44d99</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4280ccb2671145220be23f299bc0db4a2c061ff5"><code>4280ccb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a264562983a722b90b7a08e5547de85274006e30"><code>a264562</code></a> fix awsjson error deserialization to not expect string code (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2489">#2489</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.49.0...service/s3/v1.50.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.5.1746754376 to 2025.5.1747405236
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/26dc245b8bdde0a9430fce019ad924039edf87af"><code>26dc245</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97943386aab037ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/420967825c87e352f6fff69a41538bb0472be249"><code>4209678</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276e49ae256027b8e3dae95843...</li>
<li><a href="https://github.com/content-services/zest/commit/701618d57983a7d83738619d0a680227caabe4df"><code>701618d</code></a> Update pulp bindings to 386d86254354e29d3b864523aed47886e3586b067968b5e2da9ab...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.5.1746754376...release/v2025.5.1747405236">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.32.0 to 0.33.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.33.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.33.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Rename the internal <code>Logger</code> to <code>DebugLogger</code>. This feature was only used when you set <code>Debug: True</code> in your <code>sentry.Init()</code> call. If you haven't used the Logger directly, no changes are necessary. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1012">#1012</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add support for <a href="https://docs.sentry.io/product/explore/logs/">Structured Logging</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1010">#1010</a>)</p>
<pre lang="go"><code>logger := sentry.NewLogger(ctx)
logger.Info(ctx, &quot;Hello, Logs!&quot;)
</code></pre>
<p>You can learn more about Sentry Logs on our <a href="https://docs.sentry.io/product/explore/logs/">docs</a> and the <a href="https://github.com/getsentry/sentry-go/blob/master/_examples/logs/main.go">examples</a>.</p>
</li>
<li>
<p>Add new attributes APIs, which are currently only exposed on logs. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1007">#1007</a>)</p>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Do not push a new scope on <code>StartSpan</code>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1013">#1013</a>)</li>
<li>Fix an issue where the propagated smapling decision wasn't used. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/995">#995</a>)</li>
<li>[Otel] Prefer <code>httpRoute</code> over <code>httpTarget</code> for span descriptions. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1002">#1002</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Update <code>github.com/stretchr/testify</code> to v1.8.4. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/988">#988</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.33.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.33.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Rename the internal <code>Logger</code> to <code>DebugLogger</code>. This feature was only used when you set <code>Debug: True</code> in your <code>sentry.Init()</code> call. If you haven't used the Logger directly, no changes are necessary. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1012">#1012</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add support for <a href="https://docs.sentry.io/product/explore/logs/">Structured Logging</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1010">#1010</a>)</p>
<pre lang="go"><code>logger := sentry.NewLogger(ctx)
logger.Info(ctx, &quot;Hello, Logs!&quot;)
</code></pre>
<p>You can learn more about Sentry Logs on our <a href="https://docs.sentry.io/product/explore/logs/">docs</a> and the <a href="https://github.com/getsentry/sentry-go/blob/master/_examples/logs/main.go">examples</a>.</p>
</li>
<li>
<p>Add new attributes APIs, which are currently only exposed on logs. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1007">#1007</a>)</p>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Do not push a new scope on <code>StartSpan</code>. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1013">#1013</a>)</li>
<li>Fix an issue where the propagated smapling decision wasn't used. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/995">#995</a>)</li>
<li>[Otel] Prefer <code>httpRoute</code> over <code>httpTarget</code> for span descriptions. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1002">#1002</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Update <code>github.com/stretchr/testify</code> to v1.8.4. (<a href="https://redirect.github.com/getsentry/sentry-go/issues/988">#988</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/03a4debf6ca740cd1f5a0fdbef0522066c1ceef9"><code>03a4deb</code></a> release: 0.33.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/34ba180a74ec11f739d35bb35286e070f75f4795"><code>34ba180</code></a> Prepare 0.33.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1016">#1016</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c340089ff5e8a6cd896a0720e3245c2e0609f895"><code>c340089</code></a> Add initial logger implementation (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1010">#1010</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/00a135a57d4261f0c58a4ac821d392811fdcd04a"><code>00a135a</code></a> Do not push a new scope on <code>StartSpan</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1013">#1013</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/64de096a6f33547c9aae4b1b6abed0ae0b0848bb"><code>64de096</code></a> Rename <code>Logger</code> to <code>DebugLogger</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1012">#1012</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/523e1b40a94f6edcb62e6fcd52af678f68c03c6b"><code>523e1b4</code></a> Bump github.com/stretchr/testify on otel submodule (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1011">#1011</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bf20c9f50aba46ec972260075dc28049d757b896"><code>bf20c9f</code></a> Prefer http route over http target for span description (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1002">#1002</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/021ccc1973f29e4444aa3a2d7b2bbca4d5697ea7"><code>021ccc1</code></a> Add attributes API (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1007">#1007</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c7d162cfa490f2b759109187296f7dd8acfcf66e"><code>c7d162c</code></a> Update github.com/stretchr/testify to v1.8.4 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/988">#988</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a309f56a2b567b4976eaac5681177598f6d4b986"><code>a309f56</code></a> build(deps): bump actions/create-github-app-token from 1.11.5 to 1.12.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/987">#987</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.32.0...v0.33.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.4 to 5.7.5
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.5 (May 17, 2025)</h1>
<ul>
<li>Support sslnegotiation connection option (divyam234)</li>
<li>Update golang.org/x/crypto to v0.37.0. This placates security scanners that were unable to see that pgx did not use the behavior affected by <a href="https://pkg.go.dev/vuln/GO-2025-3487">https://pkg.go.dev/vuln/GO-2025-3487</a>.</li>
<li>TraceLog now logs Acquire and Release at the debug level (dave sinclair)</li>
<li>Add support for PGTZ environment variable</li>
<li>Add support for PGOPTIONS environment variable</li>
<li>Unpin memory used by Rows quicker</li>
<li>Remove PlanScan memoization. This resolves a rare issue where scanning could be broken for one type by first scanning another. The problem was in the memoization system and benchmarking revealed that memoization was not providing any meaningful benefit.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/15bca4a4e14e0049777c1245dba4c16300fe4fd0"><code>15bca4a</code></a> Release v5.7.5</li>
<li><a href="https://github.com/jackc/pgx/commit/1d557f9116c5d8fd1c6242dbc4db1e06f44b09e1"><code>1d557f9</code></a> Remove PlanScan memoization</li>
<li><a href="https://github.com/jackc/pgx/commit/de7fe81d78655c58bc41427fef4c32c317b20884"><code>de7fe81</code></a> Use reflect.TypeFor instead of reflect.TypeOf</li>
<li><a href="https://github.com/jackc/pgx/commit/d9eb089bd72b1adc0e5347df30ba74080e5445a8"><code>d9eb089</code></a> Remove unused function</li>
<li><a href="https://github.com/jackc/pgx/commit/6be24eb08d57825e0ac68696b2ac50e4d80dea42"><code>6be24eb</code></a> Fix comment typo</li>
<li><a href="https://github.com/jackc/pgx/commit/07871c0a346cdcabfa0e39996b00557665a3b56c"><code>07871c0</code></a> Zero internal baseRows references to allow GC earlier</li>
<li><a href="https://github.com/jackc/pgx/commit/777e7e5cdf2d349c37e1eef8eedc0e21857e9b95"><code>777e7e5</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2313">#2313</a> from stampy88/tracelog_pool_additions</li>
<li><a href="https://github.com/jackc/pgx/commit/151bd026ec836cbcb2b90c01424300ee19987bb8"><code>151bd02</code></a> Switched to <code>LogLevelDebug</code></li>
<li><a href="https://github.com/jackc/pgx/commit/540fcaa9b908880ed9e82ccd3e560a3232e55a7d"><code>540fcaa</code></a> Add support for PGOPTIONS environment variable</li>
<li><a href="https://github.com/jackc/pgx/commit/3a248e3822b1178c27ad311b4110bb125f7ebb5a"><code>3a248e3</code></a> Add support for PGTZ environment variable</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.7.4...v5.7.5">compare view</a></li>
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

### Comment by @dependabot on May 26, 2025 at 04:06 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1113*
