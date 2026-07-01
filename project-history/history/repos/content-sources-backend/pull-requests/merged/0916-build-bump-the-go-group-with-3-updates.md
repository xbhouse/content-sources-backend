---
type: pull_request
number: 916
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2024-12-09T04:36:03Z
updated: 2024-12-12T14:02:32Z
closed: 2024-12-12T14:02:24Z
merged: 2024-12-12T14:02:24Z
base_branch: main
head_branch: dependabot/go_modules/go-645ad18d53
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/916
---

# Pull Request #916: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: December 09, 2024 at 04:36 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-645ad18d53`

## Description

Bumps the go group with 3 updates: [github.com/labstack/echo/v4](https://github.com/labstack/echo), [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) and [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go).

Updates `github.com/labstack/echo/v4` from 4.12.0 to 4.13.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>JWT Middleware Removed</h2>
<h2>BREAKING CHANGE: JWT Middleware Removed from Core</h2>
<p>The JWT middleware has been <strong>removed from Echo core</strong> due to another security vulnerability, <a href="https://nvd.nist.gov/vuln/detail/CVE-2024-51744">CVE-2024-51744</a>. For more details, refer to issue <a href="https://redirect.github.com/labstack/echo/issues/2699">#2699</a>. A drop-in replacement is available in the <a href="https://github.com/labstack/echo-jwt">labstack/echo-jwt</a> repository or see <a href="https://github.com/labstack/echo-jwt/discussions/29">alternative implementation</a></p>
<p><strong>Important</strong>: Direct assignments like <code>token := c.Get(&quot;user&quot;).(*jwt.Token)</code> will now cause a panic due to an invalid cast. Update your code accordingly. Replace the current imports from <code>&quot;github.com/golang-jwt/jwt&quot;</code> in your handlers to the new middleware version using <code>&quot;github.com/golang-jwt/jwt/v5&quot;</code>.</p>
<p>Background:</p>
<p>The version of <code>golang-jwt/jwt</code> (v3.2.2) previously used in Echo core has been in an unmaintained state for some time. This is not the first vulnerability affecting this library; earlier issues were addressed in [PR <a href="https://redirect.github.com/labstack/echo/issues/1946">#1946</a>](<a href="https://redirect.github.com/labstack/echo/pull/1946">labstack/echo#1946</a>).
JWT middleware was marked as deprecated in Echo core as of <a href="https://github.com/labstack/echo/releases/tag/v4.10.0">v4.10.0</a> on 2022-12-27. If you did not notice that, consider leveraging tools like <a href="https://staticcheck.dev/">Staticcheck</a> to catch such deprecations earlier in you dev/CI flow.  For bonus points - check out <a href="https://github.com/securego/gosec">gosec</a>.</p>
<p>We sincerely apologize for any inconvenience caused by this change. While we strive to maintain backward compatibility within Echo core, recurring security issues with third-party dependencies have forced this decision.</p>
<p><strong>Enhancements</strong></p>
<ul>
<li>remove jwt middleware by <a href="https://github.com/stevenwhitehead"><code>@​stevenwhitehead</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2701">labstack/echo#2701</a></li>
<li>optimization: struct alignment by <a href="https://github.com/behnambm"><code>@​behnambm</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2636">labstack/echo#2636</a></li>
<li>bind: Maintain backwards compatibility for map[string]interface{} binding by <a href="https://github.com/thesaltree"><code>@​thesaltree</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2656">labstack/echo#2656</a></li>
<li>Add Go 1.23 to CI by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2675">labstack/echo#2675</a></li>
<li>improve <code>MultipartForm</code> test by <a href="https://github.com/martinyonatann"><code>@​martinyonatann</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2682">labstack/echo#2682</a></li>
<li><code>bind</code> : add support of multipart multi files by <a href="https://github.com/martinyonatann"><code>@​martinyonatann</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2684">labstack/echo#2684</a></li>
<li>Add TemplateRenderer struct to ease creating renderers for <code>html/template</code> and <code>text/template</code> packages. by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2690">labstack/echo#2690</a></li>
<li>Refactor TestBasicAuth to utilize table-driven test format by <a href="https://github.com/ErikOlson"><code>@​ErikOlson</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2688">labstack/echo#2688</a></li>
<li>Remove broken header by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2705">labstack/echo#2705</a></li>
<li>fix(bind body): content-length can be -1 by <a href="https://github.com/phamvinhdat"><code>@​phamvinhdat</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2710">labstack/echo#2710</a></li>
<li>CORS middleware should compile allowOrigin regexp at creation by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2709">labstack/echo#2709</a></li>
<li>Shorten Github issue template and add test example by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2711">labstack/echo#2711</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/behnambm"><code>@​behnambm</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2636">labstack/echo#2636</a></li>
<li><a href="https://github.com/thesaltree"><code>@​thesaltree</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2656">labstack/echo#2656</a></li>
<li><a href="https://github.com/martinyonatann"><code>@​martinyonatann</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2682">labstack/echo#2682</a></li>
<li><a href="https://github.com/ErikOlson"><code>@​ErikOlson</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2688">labstack/echo#2688</a></li>
<li><a href="https://github.com/phamvinhdat"><code>@​phamvinhdat</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2710">labstack/echo#2710</a></li>
<li><a href="https://github.com/stevenwhitehead"><code>@​stevenwhitehead</code></a> made their first contribution in <a href="https://redirect.github.com/labstack/echo/pull/2701">labstack/echo#2701</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.12.0...v4.13.0">https://github.com/labstack/echo/compare/v4.12.0...v4.13.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.13.0 - 2024-12-04</h2>
<p><strong>BREAKING CHANGE</strong> JWT Middleware Removed from Core use <a href="https://github.com/labstack/echo-jwt">labstack/echo-jwt</a> instead</p>
<p>The JWT middleware has been <strong>removed from Echo core</strong> due to another security vulnerability, <a href="https://nvd.nist.gov/vuln/detail/CVE-2024-51744">CVE-2024-51744</a>. For more details, refer to issue <a href="https://redirect.github.com/labstack/echo/issues/2699">#2699</a>. A drop-in replacement is available in the <a href="https://github.com/labstack/echo-jwt">labstack/echo-jwt</a> repository.</p>
<p><strong>Important</strong>: Direct assignments like <code>token := c.Get(&quot;user&quot;).(*jwt.Token)</code> will now cause a panic due to an invalid cast. Update your code accordingly. Replace the current imports from <code>&quot;github.com/golang-jwt/jwt&quot;</code> in your handlers to the new middleware version using <code>&quot;github.com/golang-jwt/jwt/v5&quot;</code>.</p>
<p>Background:</p>
<p>The version of <code>golang-jwt/jwt</code> (v3.2.2) previously used in Echo core has been in an unmaintained state for some time. This is not the first vulnerability affecting this library; earlier issues were addressed in [PR <a href="https://redirect.github.com/labstack/echo/issues/1946">#1946</a>](<a href="https://redirect.github.com/labstack/echo/pull/1946">labstack/echo#1946</a>).
JWT middleware was marked as deprecated in Echo core as of <a href="https://github.com/labstack/echo/releases/tag/v4.10.0">v4.10.0</a> on 2022-12-27. If you did not notice that, consider leveraging tools like <a href="https://staticcheck.dev/">Staticcheck</a> to catch such deprecations earlier in you dev/CI flow.  For bonus points - check out <a href="https://github.com/securego/gosec">gosec</a>.</p>
<p>We sincerely apologize for any inconvenience caused by this change. While we strive to maintain backward compatibility within Echo core, recurring security issues with third-party dependencies have forced this decision.</p>
<p><strong>Enhancements</strong></p>
<ul>
<li>remove jwt middleware by <a href="https://github.com/stevenwhitehead"><code>@​stevenwhitehead</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2701">labstack/echo#2701</a></li>
<li>optimization: struct alignment by <a href="https://github.com/behnambm"><code>@​behnambm</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2636">labstack/echo#2636</a></li>
<li>bind: Maintain backwards compatibility for map[string]interface{} binding by <a href="https://github.com/thesaltree"><code>@​thesaltree</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2656">labstack/echo#2656</a></li>
<li>Add Go 1.23 to CI by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2675">labstack/echo#2675</a></li>
<li>improve <code>MultipartForm</code> test by <a href="https://github.com/martinyonatann"><code>@​martinyonatann</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2682">labstack/echo#2682</a></li>
<li><code>bind</code> : add support of multipart multi files by <a href="https://github.com/martinyonatann"><code>@​martinyonatann</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2684">labstack/echo#2684</a></li>
<li>Add TemplateRenderer struct to ease creating renderers for <code>html/template</code> and <code>text/template</code> packages. by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2690">labstack/echo#2690</a></li>
<li>Refactor TestBasicAuth to utilize table-driven test format by <a href="https://github.com/ErikOlson"><code>@​ErikOlson</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2688">labstack/echo#2688</a></li>
<li>Remove broken header by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2705">labstack/echo#2705</a></li>
<li>fix(bind body): content-length can be -1 by <a href="https://github.com/phamvinhdat"><code>@​phamvinhdat</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2710">labstack/echo#2710</a></li>
<li>CORS middleware should compile allowOrigin regexp at creation by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2709">labstack/echo#2709</a></li>
<li>Shorten Github issue template and add test example by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2711">labstack/echo#2711</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/3b017855b4d331002e2b8b28e903679b875ae3e9"><code>3b01785</code></a> Changelog for 4.13.0 (<a href="https://redirect.github.com/labstack/echo/issues/2712">#2712</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/fe2627778114fc774a1b10920e1cd55fdd97cf00"><code>fe26277</code></a> remove jwt middleware</li>
<li><a href="https://github.com/labstack/echo/commit/9e73691837f52c7fdf4898cbe5bf1d157387bdb0"><code>9e73691</code></a> Shorten Github issue template and add test example</li>
<li><a href="https://github.com/labstack/echo/commit/118c1632f274a400fd4ba168c7a789950b3c35a2"><code>118c163</code></a> CORS middleware should compile allowOrigin regexp at creation.</li>
<li><a href="https://github.com/labstack/echo/commit/a973e3bc431e157c190c9e5e0764fb22fedf4407"><code>a973e3b</code></a> add unit-test</li>
<li><a href="https://github.com/labstack/echo/commit/c4410fe0b88ee0d395e23ba8b0fdfcfbce18a7ca"><code>c4410fe</code></a> fix(bind body): content-length can be -1</li>
<li><a href="https://github.com/labstack/echo/commit/5d98929328ad3ea28c7d6d0e1e8f2bee43c0704f"><code>5d98929</code></a> Remove broken header</li>
<li><a href="https://github.com/labstack/echo/commit/5a0b4dd8063575995cbcb746a0fb31266a0de3db"><code>5a0b4dd</code></a> clean up field assignments with default values to make test configuration mor...</li>
<li><a href="https://github.com/labstack/echo/commit/03c0236fb344fa1f201f9ef6fa7020f00e0bd035"><code>03c0236</code></a> refactor basic_auth_test to utilize table driven tests</li>
<li><a href="https://github.com/labstack/echo/commit/822d11a465aa640634789dd39babb14220f887eb"><code>822d11a</code></a> Add TemplateRenderer struct to ease creating renderers for <code>html/template</code> an...</li>
<li>Additional commits viewable in <a href="https://github.com/labstack/echo/compare/v4.12.0...v4.13.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.11.1732799476 to 2024.12.1733335787
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/a419b23580b82e40b802bc47645532e6737d20bc"><code>a419b23</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a4527643d56d6e127b8e3dae95843...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.11.1732799476...release/v2024.12.1733335787">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.29.1 to 0.30.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.30.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.30.0.</p>
<h3>Features</h3>
<ul>
<li>Add <code>sentryzerolog</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/857">#857</a>)</li>
<li>Add <code>sentryslog</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/865">#865</a>)</li>
<li>Always set Mechanism Type to generic (<a href="https://redirect.github.com/getsentry/sentry-go/pull/897">#896</a>)</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Prevent panic in <code>fasthttp</code> and <code>fiber</code> integration in case a malformed URL has to be parsed (<a href="https://redirect.github.com/getsentry/sentry-go/pull/912">#912</a>)</li>
</ul>
<h3>Misc</h3>
<p>Drop support for Go 1.18, 1.19 and 1.20. The currently supported Go versions are the last 3 stable releases: 1.23, 1.22 and 1.21.</p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.30.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.30.0.</p>
<h3>Features</h3>
<ul>
<li>Add <code>sentryzerolog</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/857">#857</a>)</li>
<li>Add <code>sentryslog</code> integration (<a href="https://redirect.github.com/getsentry/sentry-go/pull/865">#865</a>)</li>
<li>Always set Mechanism Type to generic (<a href="https://redirect.github.com/getsentry/sentry-go/pull/897">#896</a>)</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Prevent panic in <code>fasthttp</code> and <code>fiber</code> integration in case a malformed URL has to be parsed (<a href="https://redirect.github.com/getsentry/sentry-go/pull/912">#912</a>)</li>
</ul>
<h3>Misc</h3>
<p>Drop support for Go 1.18, 1.19 and 1.20. The currently supported Go versions are the last 3 stable releases: 1.23, 1.22 and 1.21.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/c93781f745eb4d49c1976f64c6e12672acb0f41e"><code>c93781f</code></a> Run go mod tidy</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/3099ce162c0160326767fa25309ace99d223b86a"><code>3099ce1</code></a> Update slog and zerlog go.mod</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/1db5420dd17d47db35564cdf312ac4029d3b1783"><code>1db5420</code></a> release: 0.30.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/4ee64441e9a8a225faa5da0204e93eb9c3b843cf"><code>4ee6444</code></a> Prepare 0.30.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/908">#908</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/fd9725b0a09c8319f5d76126040fd0111a419748"><code>fd9725b</code></a> Prevent panic in case url.Parsing fails in fasthttp and fiber (<a href="https://redirect.github.com/getsentry/sentry-go/issues/912">#912</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9ed1fa94ccecb20f0084e0d9bde8fb0814c5b401"><code>9ed1fa9</code></a> fix: remove dsn from slog example (<a href="https://redirect.github.com/getsentry/sentry-go/issues/910">#910</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/fc433de848aa381ecd201257659362aef23ca69b"><code>fc433de</code></a> fix: use correct Go import for slog package (<a href="https://redirect.github.com/getsentry/sentry-go/issues/909">#909</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/10bed82eb3c0dab93f602e9a95c0573b8183ceff"><code>10bed82</code></a> add readme for slog integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/905">#905</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/297ccafbf8fb749656d42e1682bb4c016c10cc92"><code>297ccaf</code></a> build(deps): bump codecov/codecov-action from 5.0.2 to 5.0.7 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/906">#906</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9b8b7bea7a90bf023c627550e813a176ffb53f53"><code>9b8b7be</code></a> add readme for zerolog integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/904">#904</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.29.1...v0.30.0">compare view</a></li>
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

### Review by @jlsherrill - Approved on December 12, 2024 at 02:02 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/916*
