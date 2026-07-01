---
type: pull_request
number: 1351
title: "Build: Bump the go group across 1 directory with 5 updates"
state: merged
author: dependabot
created: 2026-01-05T04:06:06Z
updated: 2026-01-07T10:30:46Z
closed: 2026-01-07T10:30:44Z
merged: 2026-01-07T10:30:44Z
base_branch: main
head_branch: dependabot/go_modules/go-f6e78264a8
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1351
---

# Pull Request #1351: Build: Bump the go group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: January 05, 2026 at 04:06 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-f6e78264a8`

## Description

Bumps the go group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.14.0` | `4.15.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.94.0` | `1.95.0` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.12.1766203719` | `2025.12.1766446639` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.7.6` | `5.8.0` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.77.0` | `1.78.0` |


Updates `github.com/labstack/echo/v4` from 4.14.0 to 4.15.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.15.0</h2>
<p><strong>Security</strong></p>
<p>WARNING: <strong>If your application relies on cross-origin or same-site (same subdomain) requests do not blindly push this version to production</strong></p>
<p>The CSRF middleware now supports the <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Sec-Fetch-Site"><strong>Sec-Fetch-Site</strong></a> header as a modern, defense-in-depth approach to <a href="https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#fetch-metadata-headers">CSRF
protection</a>, implementing the OWASP-recommended Fetch Metadata API alongside the traditional token-based mechanism.</p>
<p><strong>How it works:</strong></p>
<p>Modern browsers automatically send the <code>Sec-Fetch-Site</code> header with all requests, indicating the relationship
between the request origin and the target. The middleware uses this to make security decisions:</p>
<ul>
<li><strong><code>same-origin</code></strong> or <strong><code>none</code></strong>: Requests are allowed (exact origin match or direct user navigation)</li>
<li><strong><code>same-site</code></strong>: Falls back to token validation (e.g., subdomain to main domain)</li>
<li><strong><code>cross-site</code></strong>: Blocked by default with 403 error for unsafe methods (POST, PUT, DELETE, PATCH)</li>
</ul>
<p>For browsers that don't send this header (older browsers), the middleware seamlessly falls back to
traditional token-based CSRF protection.</p>
<p><strong>New Configuration Options:</strong></p>
<ul>
<li><code>TrustedOrigins []string</code>: Allowlist specific origins for cross-site requests (useful for OAuth callbacks, webhooks)</li>
<li><code>AllowSecFetchSiteFunc func(echo.Context) (bool, error)</code>: Custom logic for same-site/cross-site request validation</li>
</ul>
<p><strong>Example:</strong></p>
<pre lang="go"><code>e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
    // Allow OAuth callbacks from trusted provider
    TrustedOrigins: []string{&quot;https://oauth-provider.com&quot;},
<pre><code>// Custom validation for same-site requests
AllowSecFetchSiteFunc: func(c echo.Context) (bool, error) {
    // Your custom authorization logic here
    return validateCustomAuth(c), nil
    // return true, err  // blocks request with error
    // return true, nil  // allows CSRF request through
    // return false, nil // falls back to legacy token logic
},
</code></pre>
<p>}))<br />
</code></pre></p>
<p>PR: <a href="https://redirect.github.com/labstack/echo/pull/2858">labstack/echo#2858</a></p>
<p><strong>Type-Safe Generic Parameter Binding</strong></p>
<ul>
<li>
<p>Added generic functions for type-safe parameter extraction and context access by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2856">labstack/echo#2856</a></p>
<p>Echo now provides generic functions for extracting path, query, and form parameters with automatic type conversion,
eliminating manual string parsing and type assertions.</p>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.15.0 - 2026-01-01</h2>
<p><strong>Security</strong></p>
<p>NB: <strong>If your application relies on cross-origin or same-site (same subdomain) requests do not blindly push this version to production</strong></p>
<p>The CSRF middleware now supports the <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Sec-Fetch-Site"><strong>Sec-Fetch-Site</strong></a> header as a modern, defense-in-depth approach to <a href="https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#fetch-metadata-headers">CSRF
protection</a>, implementing the OWASP-recommended Fetch Metadata API alongside the traditional token-based mechanism.</p>
<p><strong>How it works:</strong></p>
<p>Modern browsers automatically send the <code>Sec-Fetch-Site</code> header with all requests, indicating the relationship
between the request origin and the target. The middleware uses this to make security decisions:</p>
<ul>
<li><strong><code>same-origin</code></strong> or <strong><code>none</code></strong>: Requests are allowed (exact origin match or direct user navigation)</li>
<li><strong><code>same-site</code></strong>: Falls back to token validation (e.g., subdomain to main domain)</li>
<li><strong><code>cross-site</code></strong>: Blocked by default with 403 error for unsafe methods (POST, PUT, DELETE, PATCH)</li>
</ul>
<p>For browsers that don't send this header (older browsers), the middleware seamlessly falls back to
traditional token-based CSRF protection.</p>
<p><strong>New Configuration Options:</strong></p>
<ul>
<li><code>TrustedOrigins []string</code>: Allowlist specific origins for cross-site requests (useful for OAuth callbacks, webhooks)</li>
<li><code>AllowSecFetchSiteFunc func(echo.Context) (bool, error)</code>: Custom logic for same-site/cross-site request validation</li>
</ul>
<p><strong>Example:</strong></p>
<pre lang="go"><code>e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
    // Allow OAuth callbacks from trusted provider
    TrustedOrigins: []string{&quot;https://oauth-provider.com&quot;},
<pre><code>// Custom validation for same-site requests
AllowSecFetchSiteFunc: func(c echo.Context) (bool, error) {
    // Your custom authorization logic here
    return validateCustomAuth(c), nil
    // return true, err  // blocks request with error
    // return true, nil  // allows CSRF request through
    // return false, nil // falls back to legacy token logic
},
</code></pre>
<p>}))<br />
</code></pre></p>
<p>PR: <a href="https://redirect.github.com/labstack/echo/pull/2858">labstack/echo#2858</a></p>
<p><strong>Type-Safe Generic Parameter Binding</strong></p>
<ul>
<li>
<p>Added generic functions for type-safe parameter extraction and context access by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2856">labstack/echo#2856</a></p>
<p>Echo now provides generic functions for extracting path, query, and form parameters with automatic type conversion,</p>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/482bb46fe5c7eb7c9fd7bec7d3128433dea21bee"><code>482bb46</code></a> v4.15.0 changelog</li>
<li><a href="https://github.com/labstack/echo/commit/d0f9d1e73503f38a82719562a7ff28ae06730d9e"><code>d0f9d1e</code></a> CRSF with Sec-Fetch-Site=same-site falls back to legacy token</li>
<li><a href="https://github.com/labstack/echo/commit/f3fc61848f0d61ded6df3f5cd73ebdecfd1ffb8d"><code>f3fc618</code></a> CRSF with Sec-Fetch-Site checks</li>
<li><a href="https://github.com/labstack/echo/commit/4dcb9b44f0a14663d74bb44644591f0ec4d68af8"><code>4dcb9b4</code></a> licence headers</li>
<li><a href="https://github.com/labstack/echo/commit/cbc0ac1dbc7b1a03feb45b21f52783e5a1d24df0"><code>cbc0ac1</code></a> Add PathParam(Or)/QueryParam(Or)/FormParam(Or) generic functions</li>
<li><a href="https://github.com/labstack/echo/commit/6b14f4ef3f37387827fa1cd514ea32a63425ee0c"><code>6b14f4e</code></a> Add Context.Get generic functions</li>
<li><a href="https://github.com/labstack/echo/commit/321530d2c2d12b4f25b55253d11fc3e0506f9889"><code>321530d</code></a> disable test - returns different error under Windows</li>
<li><a href="https://github.com/labstack/echo/commit/c8abd9f7db5e816161f2171d913c1f65e59ac547"><code>c8abd9f</code></a> disable flaky test</li>
<li><a href="https://github.com/labstack/echo/commit/9fe43f78b8195896a27a8c8a5219ca7eb08fae26"><code>9fe43f7</code></a> fix Rate limiter disallows fractional rates</li>
<li><a href="https://github.com/labstack/echo/commit/1b5122aaed169a882fba920ae7105ee01c54d023"><code>1b5122a</code></a> document things to reduce false positives</li>
<li>Additional commits viewable in <a href="https://github.com/labstack/echo/compare/v4.14.0...v4.15.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.94.0 to 1.95.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c6af97c5aac787ce84a5b3794be2c85e61faae56"><code>c6af97c</code></a> Release 2025-12-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44c79bcf5ea5c91aa2cfacf1266550b840cddf72"><code>44c79bc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cbd895944cc8a09617c1b9100720f582a3250f2d"><code>cbd8959</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8db5742cd5c378d0529a19c693c7cc16fc16f882"><code>8db5742</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/567ea661f59a4b913e2d18078a96e90e488562e9"><code>567ea66</code></a> Release 2025-12-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6bedc004ae386aa60707c65519004b170b6193fa"><code>6bedc00</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4b32467ac353bf36f904362c5e8be520c9150588"><code>4b32467</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0183b9c8f1c84d6553146a29612356b754b9dab0"><code>0183b9c</code></a> Release 2025-12-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3fc0009489b9b9b6c48aecd3ae16f6c988618a4"><code>e3fc000</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/552358d4e6aad86db68adb7514859428829f0817"><code>552358d</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.94.0...service/s3/v1.95.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.12.1766203719 to 2025.12.1766446639
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/a38df6698cbc1d420ba852a4f49b92da34a2bfdc"><code>a38df66</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276a9bdbabb4c27b8e3dae9584...</li>
<li><a href="https://github.com/content-services/zest/commit/22d3252612329db82b35ba22e9e167ecd7f45cbf"><code>22d3252</code></a> Update pulp bindings to e69db48356e528a464be3da896237b39b55a8e40a7d54eb5892a9...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.12.1766203719...release/v2025.12.1766446639">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.6 to 5.8.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.8.0 (December 26, 2025)</h1>
<ul>
<li>Require Go 1.24+</li>
<li>Remove golang.org/x/crypto dependency</li>
<li>Add OptionShouldPing to control ResetSession ping behavior (ilyam8)</li>
<li>Fix: Avoid overflow when MaxConns is set to MaxInt32</li>
<li>Fix: Close batch pipeline after a query error (Anthonin Bonnefoy)</li>
<li>Faster shutdown of pgxpool.Pool background goroutines (Blake Gentry)</li>
<li>Add pgxpool ping timeout (Amirsalar Safaei)</li>
<li>Fix: Rows.FieldDescriptions for empty query</li>
<li>Scan unknown types into *any as string or []byte based on format code</li>
<li>Optimize pgtype.Numeric (Philip Dubé)</li>
<li>Add AfterNetConnect hook to pgconn.Config</li>
<li>Fix: Handle for preparing statements that fail during the Describe phase</li>
<li>Fix overflow in numeric scanning (Ilia Demianenko)</li>
<li>Fix: json/jsonb sql.Scanner source type is []byte</li>
<li>Migrate from math/rand to math/rand/v2 (Mathias Bogaert)</li>
<li>Optimize internal iobufpool (Mathias Bogaert)</li>
<li>Optimize stmtcache invalidation (Mathias Bogaert)</li>
<li>Fix: missing error case in interval parsing (Maxime Soulé)</li>
<li>Fix: invalidate statement/description cache in Exec (James Hartig)</li>
<li>ColumnTypeLength method return the type length for varbit type (DengChan)</li>
<li>Array and Composite codecs handle typed nils</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/fe8740aa0679b67e13d2f1744bce5b61567d584e"><code>fe8740a</code></a> Release v5.8.0</li>
<li><a href="https://github.com/jackc/pgx/commit/e5dde5a51169fac139fb5ff82f5e9ce0155d7f62"><code>e5dde5a</code></a> Skip test on CockroachDB</li>
<li><a href="https://github.com/jackc/pgx/commit/06f2d82cac0f9fe6a6246987412b60b33241ed42"><code>06f2d82</code></a> Remove trailing space</li>
<li><a href="https://github.com/jackc/pgx/commit/2cf78dd906605e0e3fe9772bfc4d52e2ad60becc"><code>2cf78dd</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2448">#2448</a> from DengChan/column_type_lenth_varbit</li>
<li><a href="https://github.com/jackc/pgx/commit/2d1c4ef1aac082ec1d16c0ef2d83dc66586a9acb"><code>2d1c4ef</code></a> Skip tests on CockroachDB</li>
<li><a href="https://github.com/jackc/pgx/commit/1a5fa7fc7f069c51225244f7632696166edf3ef1"><code>1a5fa7f</code></a> Array and Composite codecs handle typed nils</li>
<li><a href="https://github.com/jackc/pgx/commit/5736d0976dddf0421c8b9193489add01810676fb"><code>5736d09</code></a> ColumnTypeLength method return the type length for varbit type.</li>
<li><a href="https://github.com/jackc/pgx/commit/4c1308c14b2eabaaa365e8cd69120b1de8f953e9"><code>4c1308c</code></a> Revert &quot;stdlib matches native pgx scanning support&quot;</li>
<li><a href="https://github.com/jackc/pgx/commit/14ce2b7a62479e02eb19dc34c639d5f799d3e96f"><code>14ce2b7</code></a> Skip test on CockroachDB</li>
<li><a href="https://github.com/jackc/pgx/commit/65b2724ea5dd0dc3f75a9b5d98029dbbeeec5bd9"><code>65b2724</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2443">#2443</a> from jameshartig/x-invalidate-cache-in-exec</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.7.6...v5.8.0">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.77.0 to 1.78.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.78.0</h2>
<h1>Behavior Changes</h1>
<ul>
<li>client: Reject target URLs containing unbracketed colons in the hostname in Go version 1.26+. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8716">#8716</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/neild"><code>@​neild</code></a></li>
</ul>
</li>
</ul>
<h1>New Features</h1>
<ul>
<li>stats/otel: Add backend service label to wrr metrics as part of A89. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8737">#8737</a>)</li>
<li>stats/otel: Add subchannel metrics (without the disconnection reason) to eventually replace the pickfirst metrics. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8738">#8738</a>)</li>
<li>client: Wait for all pending goroutines to complete when closing a graceful switch balancer. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8746">#8746</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/twz123"><code>@​twz123</code></a></li>
</ul>
</li>
</ul>
<h1>Bug Fixes</h1>
<ul>
<li>transport/client : Return status code <code>Unknown</code> on malformed grpc-status. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8735">#8735</a>)</li>
<li>client: Add <code>experimental.AcceptCompressors</code> so callers can restrict the <code>grpc-accept-encoding</code> header advertised for a call. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8718">#8718</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/iblancasa"><code>@​iblancasa</code></a></li>
</ul>
</li>
<li>xds: Fix a bug in <code>StringMatcher</code> where regexes would match incorrectly when ignore_case is set to true. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8723">#8723</a>)</li>
<li>xds/resolver:
<ul>
<li>Drop previous route resources and report an error when no matching virtual host is found.</li>
<li>Only log LDS/RDS configuration errors following a successful update and retain the last valid resource to prevent transient failures. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8711">#8711</a>)</li>
</ul>
</li>
<li>client:
<ul>
<li>Change connectivity state to CONNECTING when creating the name resolver (as part of exiting IDLE).</li>
<li>Change connectivity state to TRANSIENT_FAILURE if name resolver creation fails (as part of exiting IDLE).</li>
<li>Change connectivity state to IDLE after idle timeout expires even when current state is TRANSIENT_FAILURE.</li>
<li>Fix a bug that resulted in <code>OnFinish</code> call option not being invoked for RPCs where stream creation failed. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8710">#8710</a>)</li>
</ul>
</li>
<li>xdsclient: Fix a race in the xdsClient that could lead to resource-not-found errors. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8627">#8627</a>)</li>
</ul>
<h1>Performance Improvements</h1>
<ul>
<li>mem: Round up to nearest 4KiB for pool allocations larger than 1MiB. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8705">#8705</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/cjc25"><code>@​cjc25</code></a></li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/9df039ef2c921978514b600c9d5c6bf25cce54f6"><code>9df039e</code></a> Change version to 1.78.0 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8761">#8761</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/9b990b6355c443ecf9e71f118f7097b62bc3299a"><code>9b990b6</code></a> gracefulswitch: Wait for all goroutines on close (<a href="https://redirect.github.com/grpc/grpc-go/issues/8746">#8746</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/6677d9a9cf1dd8227673253015027de0addeeafb"><code>6677d9a</code></a> xds: Fixing a typo (<a href="https://redirect.github.com/grpc/grpc-go/issues/8760">#8760</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/d35cedde1ee806f3c578aba8c59bec7117ae0bc3"><code>d35cedd</code></a> xds/resolver: pass route's auto_host_rewrite to LB picker (gRFC A81) (<a href="https://redirect.github.com/grpc/grpc-go/issues/8740">#8740</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/d931fdc379623f36d2050487887f5465a18b7912"><code>d931fdc</code></a> client: allow overriding grpc-accept-encoding header (<a href="https://redirect.github.com/grpc/grpc-go/issues/8718">#8718</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/0800ec75223cd0995d599002581eafe2151c6df3"><code>0800ec7</code></a> xds/clusterimpl: update TestChildPolicyChangeOnConfigUpdate to use custom lb ...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/6553ea1a1d99ff4e3a516499330bf47607e7708f"><code>6553ea1</code></a> stats/otel: Add subchannel metrics (A94) (<a href="https://redirect.github.com/grpc/grpc-go/issues/8738">#8738</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/81a00cecc0abe8a7d7140967f96d9cc0729a3aa4"><code>81a00ce</code></a> grpc: Fixing spelling typo (<a href="https://redirect.github.com/grpc/grpc-go/issues/8756">#8756</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/e413838c3b7b8b3e94754cb3704751e49f917358"><code>e413838</code></a> client: Change connectivity state to CONNECTING when creating the name resolv...</li>
<li><a href="https://github.com/grpc/grpc-go/commit/f9d2bdb34edcd95f0ca9e2cfaba692722cb85ee2"><code>f9d2bdb</code></a> stats/otel: Add grpc.lb.backend_service label to wrr metrics (A89) (<a href="https://redirect.github.com/grpc/grpc-go/issues/8737">#8737</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/grpc/grpc-go/compare/v1.77.0...v1.78.0">compare view</a></li>
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

### Review by @TenSt - Approved on January 07, 2026 at 10:30 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1351*
