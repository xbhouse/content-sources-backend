---
type: pull_request
number: 1132
title: "Build: Bump the go group across 1 directory with 9 updates"
state: merged
author: dependabot
created: 2025-06-24T12:42:33Z
updated: 2025-06-25T19:03:42Z
closed: 2025-06-25T19:03:34Z
merged: 2025-06-25T19:03:34Z
base_branch: main
head_branch: dependabot/go_modules/go-0e06cc3290
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1132
---

# Pull Request #1132: Build: Bump the go group across 1 directory with 9 updates

**Author**: @dependabot
**Created**: June 24, 2025 at 12:42 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-0e06cc3290`

## Description

Bumps the go group with 8 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.4` | `1.36.5` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.16` | `1.29.17` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.50.2` | `1.51.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.80.2` | `1.81.0` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.6.1749665466` | `2025.6.1750263718` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.33.0` | `0.34.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.10.0` | `9.11.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.14` | `1.25.15` |


Updates `github.com/aws/aws-sdk-go-v2` from 1.36.4 to 1.36.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe7486fa32d029fc295c3146b3348cfd8ea7a9e7"><code>fe7486f</code></a> Release 2025-06-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1b6cafd86b0bbb70f1a42e4553fabcf59df8ec16"><code>1b6cafd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8db11372bfef4048b8991442ca84a8ee60b058d0"><code>8db1137</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3c5cc0eaa2b8346d78db7d7195df4be67a0e43d0"><code>3c5cc0e</code></a> Bump smithy go version (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3115">#3115</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67464865ce54f160aec17af0bae17103f7176c85"><code>6746486</code></a> Release 2025-06-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f536b0d13ffc5515db661a67aad22df63ff52955"><code>f536b0d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b014b35c2e215e681ee57b0ba9ee5c24f3916802"><code>b014b35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5c12a3b4209ebed699b200c3eb7a795669753f43"><code>5c12a3b</code></a> correct smithy go bumping to correct commit in its main branch (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3112">#3112</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b65321fe08ee5c78f469fc58f3c387343319210f"><code>b65321f</code></a> Bump codegen for CBOR empty check (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3111">#3111</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e6c02cdac88a318c65b484823eb8fd7ab8ece7c1"><code>e6c02cd</code></a> Release 2025-06-12</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.36.4...v1.36.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.16 to 1.29.17
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe7486fa32d029fc295c3146b3348cfd8ea7a9e7"><code>fe7486f</code></a> Release 2025-06-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1b6cafd86b0bbb70f1a42e4553fabcf59df8ec16"><code>1b6cafd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8db11372bfef4048b8991442ca84a8ee60b058d0"><code>8db1137</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3c5cc0eaa2b8346d78db7d7195df4be67a0e43d0"><code>3c5cc0e</code></a> Bump smithy go version (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3115">#3115</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67464865ce54f160aec17af0bae17103f7176c85"><code>6746486</code></a> Release 2025-06-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f536b0d13ffc5515db661a67aad22df63ff52955"><code>f536b0d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b014b35c2e215e681ee57b0ba9ee5c24f3916802"><code>b014b35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5c12a3b4209ebed699b200c3eb7a795669753f43"><code>5c12a3b</code></a> correct smithy go bumping to correct commit in its main branch (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3112">#3112</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b65321fe08ee5c78f469fc58f3c387343319210f"><code>b65321f</code></a> Bump codegen for CBOR empty check (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3111">#3111</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e6c02cdac88a318c65b484823eb8fd7ab8ece7c1"><code>e6c02cd</code></a> Release 2025-06-12</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.16...config/v1.29.17">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.69 to 1.17.70
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe7486fa32d029fc295c3146b3348cfd8ea7a9e7"><code>fe7486f</code></a> Release 2025-06-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1b6cafd86b0bbb70f1a42e4553fabcf59df8ec16"><code>1b6cafd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8db11372bfef4048b8991442ca84a8ee60b058d0"><code>8db1137</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3c5cc0eaa2b8346d78db7d7195df4be67a0e43d0"><code>3c5cc0e</code></a> Bump smithy go version (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3115">#3115</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67464865ce54f160aec17af0bae17103f7176c85"><code>6746486</code></a> Release 2025-06-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f536b0d13ffc5515db661a67aad22df63ff52955"><code>f536b0d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b014b35c2e215e681ee57b0ba9ee5c24f3916802"><code>b014b35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5c12a3b4209ebed699b200c3eb7a795669753f43"><code>5c12a3b</code></a> correct smithy go bumping to correct commit in its main branch (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3112">#3112</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b65321fe08ee5c78f469fc58f3c387343319210f"><code>b65321f</code></a> Bump codegen for CBOR empty check (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3111">#3111</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e6c02cdac88a318c65b484823eb8fd7ab8ece7c1"><code>e6c02cd</code></a> Release 2025-06-12</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.69...credentials/v1.17.70">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.50.2 to 1.51.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e9b00e26a17fcdb0b01772ad19fc6f733abac749"><code>e9b00e2</code></a> Release 2024-02-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cfc53e095ae2e97185c6f594d1725320f152304"><code>0cfc53e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/92d6c194cfde05280d8836ef7b56c36fd7fd926d"><code>92d6c19</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18adb3114cc789def28fac1b718657d33443ed5c"><code>18adb31</code></a> add middleware snapshot tests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2513">#2513</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7888f0e10eb7ef5db0b67a8a27bb8eeee984257b"><code>7888f0e</code></a> Release 2024-02-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a9cd1dd63c4532e0afe4e5314750bf08f1ae68b"><code>4a9cd1d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bebcd8c54f31abf4fc37a2205de460dae58ac122"><code>bebcd8c</code></a> Update SDK's smithy-go dependency to v1.20.1</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/374e0b974c5f03f15165917cd839a03110c36404"><code>374e0b9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98b0dc8950774e9a1691307dc845a50e066d847c"><code>98b0dc8</code></a> dep: drop dependency on go-cmp in protocol tests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2506">#2506</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.50.2...service/s3/v1.51.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.80.2 to 1.81.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a99eaedffb9b9a74793f62e7f45e2978c5f8c9f7"><code>a99eaed</code></a> Release 2025-06-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50ffb0cc9c8580688446d3181c5b7ec553b47532"><code>50ffb0c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7df1e36feff81981f15575380c6dd756adeffd4f"><code>7df1e36</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b6f660e2b597a2b765bc0878e1d7ed2ffd5b0333"><code>b6f660e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe7486fa32d029fc295c3146b3348cfd8ea7a9e7"><code>fe7486f</code></a> Release 2025-06-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1b6cafd86b0bbb70f1a42e4553fabcf59df8ec16"><code>1b6cafd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8db11372bfef4048b8991442ca84a8ee60b058d0"><code>8db1137</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3c5cc0eaa2b8346d78db7d7195df4be67a0e43d0"><code>3c5cc0e</code></a> Bump smithy go version (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3115">#3115</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67464865ce54f160aec17af0bae17103f7176c85"><code>6746486</code></a> Release 2025-06-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f536b0d13ffc5515db661a67aad22df63ff52955"><code>f536b0d</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.80.2...service/s3/v1.81.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.6.1749665466 to 2025.6.1750263718
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/2cf2b6c33964f0f81f9a94906a438a72ccbd387f"><code>2cf2b6c</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47548b5a236b0b7a98d596b8e3...</li>
<li><a href="https://github.com/content-services/zest/commit/047f06cb1fcf596a5786e4c0e54c91ee95c46a45"><code>047f06c</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d593b9d3ea157952eab83695...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.6.1749665466...release/v2025.6.1750263718">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.33.0 to 0.34.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.34.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.34.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Logrus structured logging support replaces the <code>sentrylogrus.Hook</code> signature from a <code>*Hook</code> to an interface.</li>
</ul>
<pre lang="go"><code>var hook *sentrylogrus.Hook
hook = sentrylogrus.New(
    // ... your setup
)
<p>// should change the definition to<br />
var hook sentrylogrus.Hook<br />
hook = sentrylogrus.New(<br />
// ... your setup<br />
)<br />
</code></pre></p>
<h3>Features</h3>
<ul>
<li>Structured logging support for <a href="https://pkg.go.dev/log/slog">slog</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1033">#1033</a>)</li>
</ul>
<pre lang="go"><code>ctx := context.Background()
handler := sentryslog.Option{
    EventLevel: []slog.Level{slog.LevelError, sentryslog.LevelFatal}, // Only Error and Fatal as events
    LogLevel:   []slog.Level{slog.LevelWarn, slog.LevelInfo},         // Only Warn and Info as logs
}.NewSentryHandler(ctx)
logger := slog.New(handler)
logger.Info(&quot;hello&quot;))
</code></pre>
<ul>
<li>Structured logging support for <a href="https://github.com/sirupsen/logrus">logrus</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1036">#1036</a>)</li>
</ul>
<pre lang="go"><code>logHook, _ := sentrylogrus.NewLogHook(
    []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}, 
    sentry.ClientOptions{
        Dsn: &quot;your-dsn&quot;,
        EnableLogs: true, // Required for log entries    
    })
defer logHook.Flush(5 * time.Secod)
logrus.RegisterExitHandler(func() {
    logHook.Flush(5 * time.Second)
})
<p>logger := logrus.New()<br />
logger.AddHook(logHook)<br />
logger.Infof(&quot;hello&quot;)<br />
&lt;/tr&gt;&lt;/table&gt;<br />
</code></pre></p>
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.34.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.34.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>Logrus structured logging support replaces the <code>sentrylogrus.Hook</code> signature from a <code>*Hook</code> to an interface.</li>
</ul>
<pre lang="go"><code>var hook *sentrylogrus.Hook
hook = sentrylogrus.New(
    // ... your setup
)
<p>// should change the definition to<br />
var hook sentrylogrus.Hook<br />
hook = sentrylogrus.New(<br />
// ... your setup<br />
)<br />
</code></pre></p>
<h3>Features</h3>
<ul>
<li>Structured logging support for <a href="https://pkg.go.dev/log/slog">slog</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1033">#1033</a>)</li>
</ul>
<pre lang="go"><code>ctx := context.Background()
handler := sentryslog.Option{
    EventLevel: []slog.Level{slog.LevelError, sentryslog.LevelFatal}, // Only Error and Fatal as events
    LogLevel:   []slog.Level{slog.LevelWarn, slog.LevelInfo},         // Only Warn and Info as logs
}.NewSentryHandler(ctx)
logger := slog.New(handler)
logger.Info(&quot;hello&quot;))
</code></pre>
<ul>
<li>Structured logging support for <a href="https://github.com/sirupsen/logrus">logrus</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1036">#1036</a>)</li>
</ul>
<pre lang="go"><code>logHook, _ := sentrylogrus.NewLogHook(
    []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}, 
    sentry.ClientOptions{
        Dsn: &quot;your-dsn&quot;,
        EnableLogs: true, // Required for log entries    
    })
defer logHook.Flush(5 * time.Secod)
logrus.RegisterExitHandler(func() {
    logHook.Flush(5 * time.Second)
})
<p>logger := logrus.New()<br />
logger.AddHook(logHook)<br />
&lt;/tr&gt;&lt;/table&gt;<br />
</code></pre></p>
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/4d9fea959a8e45ec25a6df36df2c0d17b5992799"><code>4d9fea9</code></a> release: 0.34.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/3c922c49634b573346da3f1e996ecaf6a63816be"><code>3c922c4</code></a> Prepare 0.34.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1040">#1040</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/349dd613822f8d4ef8c43878c35bd48c1599f931"><code>349dd61</code></a> Remove unneeded handling on logrus data (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1041">#1041</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/5bc6aade52f4b58055737b38464910a44e5305b3"><code>5bc6aad</code></a> Structured logging support for slog (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1033">#1033</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f43c181c99e9ad63e03f5b99078d13a277ea57a5"><code>f43c181</code></a> Structured logging support for logrus (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1036">#1036</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f28f4987b4f0e5465500149c8c7bb93611334ab5"><code>f28f498</code></a> Add slog fingerprint handler (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1039">#1039</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/28379526fd1bf44c5829d4c98579bf6820fcac7f"><code>2837952</code></a> Implement flush with context (<a href="https://redirect.github.com/getsentry/sentry-go/issues/935">#935</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e75b1200b85950f4a82528feadb11ba0f340548f"><code>e75b120</code></a> Fix remove sentry.origin for sentry logger (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1038">#1038</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/69f7e3f6fa9360747935d0040c70d248c96e09d7"><code>69f7e3f</code></a> chore: Rename bulider.go to builder.go (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1037">#1037</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/4c908b750b8f71f837214858e7087698da43b712"><code>4c908b7</code></a> Fix use correct Level signature for logs (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1034">#1034</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.33.0...v0.34.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.10.0 to 9.11.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.11.0</h2>
<h2>🚀 Highlights</h2>
<p>Fixes <code>TxPipeline</code> to work correctly in cluster scenarios, allowing execution of commands
only in the same slot for a given transaction.</p>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Set cluster slot for <code>scan</code> commands, rather than random (<a href="https://redirect.github.com/redis/go-redis/pull/2623">#2623</a>)</li>
<li>Add CredentialsProvider field to UniversalOptions (<a href="https://redirect.github.com/redis/go-redis/pull/2927">#2927</a>)</li>
<li>feat(redisotel): add WithCallerEnabled option (<a href="https://redirect.github.com/redis/go-redis/pull/3415">#3415</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix(txpipeline): keyless commands should take the slot of the keyed (<a href="https://redirect.github.com/redis/go-redis/pull/3411">#3411</a>)</li>
<li>fix(loading): cache the loaded flag for slave nodes (<a href="https://redirect.github.com/redis/go-redis/pull/3410">#3410</a>)</li>
<li>fix(txpipeline): should return error on multi/exec on multiple slots (<a href="https://redirect.github.com/redis/go-redis/pull/3408">#3408</a>)</li>
<li>fix: check if the shard exists to avoid returning nil (<a href="https://redirect.github.com/redis/go-redis/pull/3396">#3396</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>feat: optimize connection pool waitTurn (<a href="https://redirect.github.com/redis/go-redis/pull/3412">#3412</a>)</li>
<li>chore(ci): update CI redis builds (<a href="https://redirect.github.com/redis/go-redis/pull/3407">#3407</a>)</li>
<li>chore: remove a redundant method from <code>Ring</code>, <code>Client</code> and <code>ClusterClient</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3401">#3401</a>)</li>
<li>test: refactor TestBasicCredentials using table-driven tests (<a href="https://redirect.github.com/redis/go-redis/pull/3406">#3406</a>)</li>
<li>perf: reduce unnecessary memory allocation operations (<a href="https://redirect.github.com/redis/go-redis/pull/3399">#3399</a>)</li>
<li>fix: insert entry during iterating over a map (<a href="https://redirect.github.com/redis/go-redis/pull/3398">#3398</a>)</li>
<li>DOC-5229 probabilistic data type examples (<a href="https://redirect.github.com/redis/go-redis/pull/3413">#3413</a>)</li>
<li>chore(deps): bump rojopolis/spellcheck-github-actions from 0.49.0 to 0.51.0 (<a href="https://redirect.github.com/redis/go-redis/pull/3414">#3414</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/andy-stark-redis"><code>@​andy-stark-redis</code></a>, <a href="https://github.com/boekkooi-impossiblecloud"><code>@​boekkooi-impossiblecloud</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/dcherubini"><code>@​dcherubini</code></a>, <a href="https://github.com/iamamirsalehi"><code>@​iamamirsalehi</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/pete-woods"><code>@​pete-woods</code></a>, <a href="https://github.com/twz915"><code>@​twz915</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/0decfdc6ed1710b0abf5b41741bd02223c877c92"><code>0decfdc</code></a> chore(release): v9.11.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3416">#3416</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/4ac591c7c4ae10e703f5e08bbe93462fd5371678"><code>4ac591c</code></a> Set correct cluster slot for scan commands, similarly to Java's Jedis client ...</li>
<li><a href="https://github.com/redis/go-redis/commit/0383d08a351adae039a6a2648d0ea1d5ee97afd3"><code>0383d08</code></a> feat(client): Add CredentialsProvider field to UniversalOptions (<a href="https://redirect.github.com/redis/go-redis/issues/2927">#2927</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fa475cbc99f7fd36a2500839a8e7aae7fcd4e565"><code>fa475cb</code></a> feat(redisotel): add WithCallerEnabled option (<a href="https://redirect.github.com/redis/go-redis/issues/3415">#3415</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/05f42e23271eed66e2682d9cbe542bcc6b084f17"><code>05f42e2</code></a> fix(txpipeline): keyless commands should take the slot of the keyed (<a href="https://redirect.github.com/redis/go-redis/issues/3411">#3411</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/884f9970c025c346c7b223f62aae45e903fb0b56"><code>884f997</code></a> DOC-5229 probabilistic data type examples (<a href="https://redirect.github.com/redis/go-redis/issues/3413">#3413</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e642856ed3329077633c28d47cfb4ac66ca3be0b"><code>e642856</code></a> chore(deps): bump rojopolis/spellcheck-github-actions (<a href="https://redirect.github.com/redis/go-redis/issues/3414">#3414</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7d97cc1c594658d91b04d4d641e992452345a0de"><code>7d97cc1</code></a> feat: optimize connection pool waitTurn (<a href="https://redirect.github.com/redis/go-redis/issues/3412">#3412</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f4358acffcb236879a7467b3b13b6ce5d211a536"><code>f4358ac</code></a> [CAE-1046] fix(loading): cache the loaded flag for slave nodes (<a href="https://redirect.github.com/redis/go-redis/issues/3410">#3410</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/4c635cc563b6fd5c3a90f293f2efeb9ccc84953f"><code>4c635cc</code></a> fix(txpipeline): should return error on multi/exec on multiple slots [CAE-102...</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.10.0...v9.11.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.14 to 1.25.15
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>v1.25.15</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/redis/go-redis/v9 from 9.9.0 to 9.10.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/618">RedHatInsights/insights-operator-utils#618</a></li>
<li>[CCXDEV-15205] Change provider of sarama package by <a href="https://github.com/JiriPapousek"><code>@​JiriPapousek</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/619">RedHatInsights/insights-operator-utils#619</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.14...v1.25.15">https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.14...v1.25.15</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4b516342fed9c0c5b3459e10bd06ddb7d5af88ff"><code>4b51634</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/619">#619</a> from JiriPapousek/ibm-sarama</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/27bd93152a4fd39ebdfc0c39851fcae98ac391e6"><code>27bd931</code></a> Change provider of sarama package</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/1d5469fe2ba61db8850cf933844752102bdbec6f"><code>1d5469f</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/618">#618</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/92f6f835b28218e458264cfeec69f2b52b870d4f"><code>92f6f83</code></a> Bump github.com/redis/go-redis/v9 from 9.9.0 to 9.10.0</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.14...v1.25.15">compare view</a></li>
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

### Comment by @rverdile on June 25, 2025 at 02:06 PM UTC

@dependabot rebase

---

## Reviews

### Review by @jlsherrill - Approved on June 25, 2025 at 07:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1132*
