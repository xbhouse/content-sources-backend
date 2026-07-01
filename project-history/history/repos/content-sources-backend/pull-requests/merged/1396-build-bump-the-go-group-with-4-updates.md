---
type: pull_request
number: 1396
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2026-02-23T04:42:47Z
updated: 2026-02-23T09:57:23Z
closed: 2026-02-23T09:57:21Z
merged: 2026-02-23T09:57:21Z
base_branch: main
head_branch: dependabot/go_modules/go-63caf26562
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1396
---

# Pull Request #1396: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: February 23, 2026 at 04:42 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-63caf26562`

## Description

Bumps the go group with 4 updates: [github.com/labstack/echo/v4](https://github.com/labstack/echo), [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) and [github.com/redis/go-redis/v9](https://github.com/redis/go-redis).

Updates `github.com/labstack/echo/v4` from 4.15.0 to 4.15.1
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
<h2>v5.0.4 - 2026-02-15</h2>
<p><strong>Enhancements</strong></p>
<ul>
<li>Remove unused import 'errors' from README example by <a href="https://github.com/kumapower17"><code>@​kumapower17</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2889">labstack/echo#2889</a></li>
<li>Fix Graceful shutdown: after <code>http.Server.Serve</code> returns we need to wait for graceful shutdown goroutine to finish by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2898">labstack/echo#2898</a></li>
<li>Update location of oapi-codegen in README by <a href="https://github.com/mromaszewicz"><code>@​mromaszewicz</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2896">labstack/echo#2896</a></li>
<li>Add Go 1.26 to CI flow by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2899">labstack/echo#2899</a></li>
<li>Add new function <code>echo.StatusCode</code> by <a href="https://github.com/suwakei"><code>@​suwakei</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2892">labstack/echo#2892</a></li>
<li>CSRF: support older token-based CSRF protection handler that want to render token into template by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2894">labstack/echo#2894</a></li>
<li>Add <code>echo.ResolveResponseStatus</code> function to help middleware/handlers determine HTTP status code and echo.Response by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2900">labstack/echo#2900</a></li>
</ul>
<h2>v5.0.3 - 2026-02-06</h2>
<p><strong>Security</strong></p>
<ul>
<li>Fix directory traversal vulnerability under Windows in Static middleware when default Echo filesystem is used. Reported by <a href="https://github.com/shblue21"><code>@​shblue21</code></a>.</li>
</ul>
<p>This applies to cases when:</p>
<ul>
<li>Windows is used as OS</li>
<li><code>middleware.StaticConfig.Filesystem</code> is <code>nil</code> (default)</li>
<li><code>echo.Filesystem</code> is has not been set explicitly (default)</li>
</ul>
<p>Exposure is restricted to the active process working directory and its subfolders.</p>
<h2>v5.0.2 - 2026-02-02</h2>
<p><strong>Security</strong></p>
<ul>
<li>Fix Static middleware with <code>config.Browse=true</code> lists all files/subfolders from <code>config.Filesystem</code> root and not starting from <code>config.Root</code> in <a href="https://redirect.github.com/labstack/echo/pull/2887">labstack/echo#2887</a></li>
</ul>
<h2>v5.0.1 - 2026-01-28</h2>
<ul>
<li>Panic MW: will now return a custom PanicStackError with stack trace by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2871">labstack/echo#2871</a></li>
<li>Docs: add missing err parameter to DenyHandler example by <a href="https://github.com/cgalibern"><code>@​cgalibern</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2878">labstack/echo#2878</a></li>
<li>improve: improve websocket checks in IsWebSocket() [per RFC 6455] by <a href="https://github.com/raju-mechatronics"><code>@​raju-mechatronics</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2875">labstack/echo#2875</a></li>
<li>fix: Context.Json() should not send status code before serialization is complete by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2877">labstack/echo#2877</a></li>
</ul>
<h2>v5.0.0 - 2026-01-18</h2>
<p>Echo <code>v5</code> is maintenance release with <strong>major breaking changes</strong></p>
<ul>
<li><code>Context</code> is now struct instead of interface and we can add method to it in the future in minor versions.</li>
<li>Adds new <code>Router</code> interface for possible new routing implementations.</li>
<li>Drops old logging interface and uses moderm <code>log/slog</code> instead.</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/6f3a84a50585110dc71c39585e7c97faf8fcaf0a"><code>6f3a84a</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/2905">#2905</a> from aldas/v4_crsf_token_fallback</li>
<li><a href="https://github.com/labstack/echo/commit/24fa4d07ff994074cfb47ca8a6b088e17cbe1711"><code>24fa4d0</code></a> CSRF: support older token-based CSRF protection handler that want to render t...</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.15.0...v4.15.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.7 to 1.32.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/841d59050195096aa618d48d77f0d57523dd3e3a"><code>841d590</code></a> Release 2026-02-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/080daf79d9875ed54bf4d86815beb3355c745e4d"><code>080daf7</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e7ec908729ff5518f51302e6291af09dcc353b9"><code>3e7ec90</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d87017019ee67c5aebbf2daf4d292530089109bc"><code>d870170</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c4012c3ecd1c1e5877cc5232ccb6a2c14305b343"><code>c4012c3</code></a> Release 2026-02-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fa02ccfb5e61483cac9607c04d04f8c3a3a620c3"><code>fa02ccf</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a9fbb34ac8ec20314e152578edb4e68054abb46"><code>5a9fbb3</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/76d06b19066f3c0955a25ae2ee4edf4f9ac9988f"><code>76d06b1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03e49f78e5a25bc8a7af19549ae2f46aa05abeda"><code>03e49f7</code></a> Release 2026-02-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb0ac66bdf2d0d77085982ae8271d95db6f11e0e"><code>eb0ac66</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.7...config/v1.32.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.7 to 1.19.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54aed732316b5162e5c4382a1f2d3891175d0254"><code>54aed73</code></a> Release 2025-02-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/082781faee876f9d612fa7c113b4304a29766b14"><code>082781f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ed185b604684a86547e679154975f1914f97312"><code>3ed185b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03da7378d668622cd880ec741d57e93cc370efa1"><code>03da737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a8ccb619ffbfe00e99a83e99729b948f20be29"><code>c8a8ccb</code></a> Bump go version to 1.22 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3010">#3010</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8b7c7bf6d9a1c63d0c5262724ae8a15a44e366a6"><code>8b7c7bf</code></a> fix missing AccountIDEndpointMode binding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3013">#3013</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90f9d1081a37acaf792ccda5bfb07e2ee7590a9e"><code>90f9d10</code></a> Release 2025-02-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40dd351c61c016749a3f4105cca0c965e7c66d7b"><code>40dd351</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06352dfafe9067da1956229d6925efed328d5ff6"><code>06352df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/833566b553122ebd5bfa1237ee7c905a8db0d687"><code>833566b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/m2/v1.19.7...service/m2/v1.19.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.17.3 to 9.18.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.18.0</h2>
<h3>Redis 8.6 Support</h3>
<p>Added support for Redis 8.6, including new commands and features for streams idempotent production and HOTKEYS.</p>
<h3>Smart Client Handoff (Maintenance Notifications) for Cluster</h3>
<p><strong>note: Pending RS version release</strong></p>
<p>This release introduces comprehensive support for Redis Enterprise Cluster maintenance notifications via SMIGRATING/SMIGRATED push notifications. The client now automatically handles slot migrations by:</p>
<ul>
<li><strong>Relaxing timeouts during migration</strong> (SMIGRATING) to prevent false failures</li>
<li><strong>Triggering lazy cluster state reloads</strong> upon completion (SMIGRATED)</li>
<li>Enabling seamless operations during Redis Enterprise maintenance windows</li>
</ul>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3643">#3643</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h3>OpenTelemetry Native Metrics Support</h3>
<p>Added comprehensive OpenTelemetry metrics support following the <a href="https://opentelemetry.io/docs/specs/semconv/database/database-metrics/">OpenTelemetry Database Client Semantic Conventions</a>. The implementation uses a Bridge Pattern to keep the core library dependency-free while providing optional metrics instrumentation through the new <code>extra/redisotel-native</code> package.</p>
<p><strong>Metric groups include:</strong></p>
<ul>
<li>Command metrics: Operation duration with retry tracking</li>
<li>Connection basic: Connection count and creation time</li>
<li>Resiliency: Errors, handoffs, timeout relaxation</li>
<li>Connection advanced: Wait time and use time</li>
<li>Pubsub metrics: Published and received messages</li>
<li>Stream metrics: Processing duration and maintenance notifications</li>
</ul>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3637">#3637</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
<h2>✨ New Features</h2>
<ul>
<li><strong>HOTKEYS Commands</strong>: Added support for Redis HOTKEYS feature for identifying hot keys based on CPU consumption and network utilization (<a href="https://redirect.github.com/redis/go-redis/pull/3695">#3695</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>Streams Idempotent Production</strong>: Added support for Redis 8.6+ Streams Idempotent Production with <code>ProducerID</code>, <code>IdempotentID</code>, <code>IdempotentAuto</code> in <code>XAddArgs</code> and new <code>XCFGSET</code> command (<a href="https://redirect.github.com/redis/go-redis/pull/3693">#3693</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>NaN Values for TimeSeries</strong>: Added support for NaN (Not a Number) values in Redis time series commands (<a href="https://redirect.github.com/redis/go-redis/pull/3687">#3687</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>DialerRetries Options</strong>: Added <code>DialerRetries</code> and <code>DialerRetryTimeout</code> to <code>ClusterOptions</code>, <code>RingOptions</code>, and <code>FailoverOptions</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3686">#3686</a>) by <a href="https://github.com/naveenchander30"><code>@​naveenchander30</code></a></li>
<li><strong>ConnMaxLifetimeJitter</strong>: Added jitter configuration to distribute connection expiration times and prevent thundering herd (<a href="https://redirect.github.com/redis/go-redis/pull/3666">#3666</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>Digest Helper Functions</strong>: Added <code>DigestString</code> and <code>DigestBytes</code> helper functions for client-side xxh3 hashing compatible with Redis DIGEST command (<a href="https://redirect.github.com/redis/go-redis/pull/3679">#3679</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>SMIGRATED New Format</strong>: Updated SMIGRATED parser to support new format and remember original host:port (<a href="https://redirect.github.com/redis/go-redis/pull/3697">#3697</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Cluster State Reload Interval</strong>: Added cluster state reload interval option for maintenance notifications (<a href="https://redirect.github.com/redis/go-redis/pull/3663">#3663</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>PubSub nil pointer dereference</strong>: Fixed nil pointer dereference in PubSub after <code>WithTimeout()</code> - <code>pubSubPool</code> is now properly cloned (<a href="https://redirect.github.com/redis/go-redis/pull/3710">#3710</a>) by <a href="https://github.com/apps/copilot-swe-agent"><code>@​Copilot</code></a></li>
<li><strong>MaintNotificationsConfig nil check</strong>: Guard against nil <code>MaintNotificationsConfig</code> in <code>initConn</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3707">#3707</a>) by <a href="https://github.com/veeceey"><code>@​veeceey</code></a></li>
<li><strong>wantConnQueue zombie elements</strong>: Fixed zombie <code>wantConn</code> elements accumulation in <code>wantConnQueue</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3680">#3680</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>XADD/XTRIM approx flag</strong>: Fixed XADD and XTRIM to use <code>=</code> when approx is false (<a href="https://redirect.github.com/redis/go-redis/pull/3684">#3684</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Sentinel timeout retry</strong>: When connection to a sentinel times out, attempt to connect to other sentinels (<a href="https://redirect.github.com/redis/go-redis/pull/3654">#3654</a>) by <a href="https://github.com/cxljs"><code>@​cxljs</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.18.0 (2026-02-16)</h1>
<h2>🚀 Highlights</h2>
<h3>Redis 8.6 Support</h3>
<p>Added support for Redis 8.6, including new commands and features for streams idempotent production and HOTKEYS.</p>
<h3>Smart Client Handoff (Maintenance Notifications) for Cluster</h3>
<p>This release introduces comprehensive support for Redis Cluster maintenance notifications via SMIGRATING/SMIGRATED push notifications. The client now automatically handles slot migrations by:</p>
<ul>
<li><strong>Relaxing timeouts during migration</strong> (SMIGRATING) to prevent false failures</li>
<li><strong>Triggering lazy cluster state reloads</strong> upon completion (SMIGRATED)</li>
<li>Enabling seamless operations during Redis Enterprise maintenance windows</li>
</ul>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3643">#3643</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h3>OpenTelemetry Native Metrics Support</h3>
<p>Added comprehensive OpenTelemetry metrics support following the <a href="https://opentelemetry.io/docs/specs/semconv/database/database-metrics/">OpenTelemetry Database Client Semantic Conventions</a>. The implementation uses a Bridge Pattern to keep the core library dependency-free while providing optional metrics instrumentation through the new <code>extra/redisotel-native</code> package.</p>
<p><strong>Metric groups include:</strong></p>
<ul>
<li>Command metrics: Operation duration with retry tracking</li>
<li>Connection basic: Connection count and creation time</li>
<li>Resiliency: Errors, handoffs, timeout relaxation</li>
<li>Connection advanced: Wait time and use time</li>
<li>Pubsub metrics: Published and received messages</li>
<li>Stream metrics: Processing duration and maintenance notifications</li>
</ul>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3637">#3637</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
<h2>✨ New Features</h2>
<ul>
<li><strong>HOTKEYS Commands</strong>: Added support for Redis HOTKEYS feature for identifying hot keys based on CPU consumption and network utilization (<a href="https://redirect.github.com/redis/go-redis/pull/3695">#3695</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>Streams Idempotent Production</strong>: Added support for Redis 8.6+ Streams Idempotent Production with <code>ProducerID</code>, <code>IdempotentID</code>, <code>IdempotentAuto</code> in <code>XAddArgs</code> and new <code>XCFGSET</code> command (<a href="https://redirect.github.com/redis/go-redis/pull/3693">#3693</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>NaN Values for TimeSeries</strong>: Added support for NaN (Not a Number) values in Redis time series commands (<a href="https://redirect.github.com/redis/go-redis/pull/3687">#3687</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>DialerRetries Options</strong>: Added <code>DialerRetries</code> and <code>DialerRetryTimeout</code> to <code>ClusterOptions</code>, <code>RingOptions</code>, and <code>FailoverOptions</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3686">#3686</a>) by <a href="https://github.com/naveenchander30"><code>@​naveenchander30</code></a></li>
<li><strong>ConnMaxLifetimeJitter</strong>: Added jitter configuration to distribute connection expiration times and prevent thundering herd (<a href="https://redirect.github.com/redis/go-redis/pull/3666">#3666</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>Digest Helper Functions</strong>: Added <code>DigestString</code> and <code>DigestBytes</code> helper functions for client-side xxh3 hashing compatible with Redis DIGEST command (<a href="https://redirect.github.com/redis/go-redis/pull/3679">#3679</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>SMIGRATED New Format</strong>: Updated SMIGRATED parser to support new format and remember original host:port (<a href="https://redirect.github.com/redis/go-redis/pull/3697">#3697</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Cluster State Reload Interval</strong>: Added cluster state reload interval option for maintenance notifications (<a href="https://redirect.github.com/redis/go-redis/pull/3663">#3663</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>PubSub nil pointer dereference</strong>: Fixed nil pointer dereference in PubSub after <code>WithTimeout()</code> - <code>pubSubPool</code> is now properly cloned (<a href="https://redirect.github.com/redis/go-redis/pull/3710">#3710</a>) by <a href="https://github.com/apps/copilot-swe-agent"><code>@​Copilot</code></a></li>
<li><strong>MaintNotificationsConfig nil check</strong>: Guard against nil <code>MaintNotificationsConfig</code> in <code>initConn</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3707">#3707</a>) by <a href="https://github.com/veeceey"><code>@​veeceey</code></a></li>
<li><strong>wantConnQueue zombie elements</strong>: Fixed zombie <code>wantConn</code> elements accumulation in <code>wantConnQueue</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3680">#3680</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>XADD/XTRIM approx flag</strong>: Fixed XADD and XTRIM to use <code>=</code> when approx is false (<a href="https://redirect.github.com/redis/go-redis/pull/3684">#3684</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Sentinel timeout retry</strong>: When connection to a sentinel times out, attempt to connect to other sentinels (<a href="https://redirect.github.com/redis/go-redis/pull/3654">#3654</a>) by <a href="https://github.com/cxljs"><code>@​cxljs</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/90faf0692313160b1140f96716763227912aec44"><code>90faf06</code></a> chore(release): update versions in deps (<a href="https://redirect.github.com/redis/go-redis/issues/3712">#3712</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/bf8e8e3b16dc5b4f51b69e9bb9a49b72fb9e3e79"><code>bf8e8e3</code></a> chore(release): v9.18.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3711">#3711</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a881cd4280bf6006c14eda4b9449168b30e7dd02"><code>a881cd4</code></a> fix(clone): nil pointer dereference in PubSub after WithTimeout() (<a href="https://redirect.github.com/redis/go-redis/issues/3710">#3710</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/ee6e9dbf09125366d1270f31969b493325f524f1"><code>ee6e9db</code></a> feat(otel): Add OpenTelemetry Native Metrics Support (<a href="https://redirect.github.com/redis/go-redis/issues/3637">#3637</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b53f2b0c9d1fa4e19680a20de2050aa7c2c5ab8d"><code>b53f2b0</code></a> feat(sch): MaintNotifications for ClusterClient (<a href="https://redirect.github.com/redis/go-redis/issues/3643">#3643</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f25343d2813d9f718a9158175d747173756f02c2"><code>f25343d</code></a> chore(tests): Add comprehensive TLS tests and example (<a href="https://redirect.github.com/redis/go-redis/issues/3681">#3681</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/33ca5cb864ce5567e8c955f66aba2723e89100d3"><code>33ca5cb</code></a> feat(commands): Add support for Redis HOTKEYS commands (<a href="https://redirect.github.com/redis/go-redis/issues/3695">#3695</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/34f4568e4f847465a5219fb51267c41433407c52"><code>34f4568</code></a> fix(conn): guard against nil MaintNotificationsConfig in initConn (<a href="https://redirect.github.com/redis/go-redis/issues/3707">#3707</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2fc030f0d1020e3c7a52011343fcd10b995c22df"><code>2fc030f</code></a> perf(options): perf Fuzz Test Go File (<a href="https://redirect.github.com/redis/go-redis/issues/3692">#3692</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/63ed1fd9cd6084e3bd68903d9f9a72df165693a7"><code>63ed1fd</code></a> Add support for Redis Streams Idempotent Production (<a href="https://redirect.github.com/redis/go-redis/issues/3693">#3693</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.17.3...v9.18.0">compare view</a></li>
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

### Review by @TenSt - Approved on February 23, 2026 at 09:57 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1396*
