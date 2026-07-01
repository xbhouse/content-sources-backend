---
type: pull_request
number: 1456
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2026-04-13T04:52:40Z
updated: 2026-04-13T10:18:26Z
closed: 2026-04-13T10:18:22Z
merged: 2026-04-13T10:18:22Z
base_branch: main
head_branch: dependabot/go_modules/go-7801202b4c
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1456
---

# Pull Request #1456: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: April 13, 2026 at 04:52 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-7801202b4c`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.134.0` | `0.135.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.98.0` | `1.99.0` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.5.0` | `4.7.5` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.3.1774545586` | `2026.4.1775140412` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.44.1` | `0.45.0` |

Updates `github.com/getkin/kin-openapi` from 0.134.0 to 0.135.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.135.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3: strip <strong>origin</strong> from Encodings and ServerVariables maps by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1132">getkin/kin-openapi#1132</a></li>
<li>fix: update yaml3 to prevent panic on empty mapping node in sequence by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1133">getkin/kin-openapi#1133</a></li>
<li>openapi3: strip <strong>origin</strong> from extension values to prevent spurious diffs by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1137">getkin/kin-openapi#1137</a></li>
<li>openapi3: strip <strong>origin</strong> from slices in example values by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1138">getkin/kin-openapi#1138</a></li>
<li>fix: bump yaml and yaml3 to v0.0.4 by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1136">getkin/kin-openapi#1136</a></li>
<li>openapi3: OriginTree approach for origin tracking — separate pass, no inline stripping by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1142">getkin/kin-openapi#1142</a></li>
<li>README: drop go-openapi/spec3 by <a href="https://github.com/zonescape"><code>@​zonescape</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1143">getkin/kin-openapi#1143</a></li>
<li>fix: bump yaml3+yaml to v0.0.9 to fix -root schema origin by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1144">getkin/kin-openapi#1144</a></li>
<li>openapi3: call ReadFromURIFunc before checking IsExternalRefsAllowed by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1146">getkin/kin-openapi#1146</a></li>
<li>fix: use location.String() instead of location.Path for origin file tracking by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1145">getkin/kin-openapi#1145</a></li>
<li>refactor: Replace sort usage with slices package by <a href="https://github.com/jedevc"><code>@​jedevc</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1147">getkin/kin-openapi#1147</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/zonescape"><code>@​zonescape</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1143">getkin/kin-openapi#1143</a></li>
<li><a href="https://github.com/jedevc"><code>@​jedevc</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1147">getkin/kin-openapi#1147</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.134.0...v0.135.0">https://github.com/getkin/kin-openapi/compare/v0.134.0...v0.135.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/9686b66c8cf9b6b242c7d80e44f43a606ca970db"><code>9686b66</code></a> refactor: Replace sort usage with slices package (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1147">#1147</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/a54fb71d4377aa9fe8cc08baf3384fc03e6f3729"><code>a54fb71</code></a> fix: use location.String() instead of location.Path for origin file tracking ...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/0a8216859759a861a2c8ef97fe0c6c90b080acdb"><code>0a82168</code></a> openapi3: call ReadFromURIFunc before checking IsExternalRefsAllowed (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1146">#1146</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/202eaa693448b1736c8dc1b4dfa19a38462e7fb8"><code>202eaa6</code></a> fix: bump yaml3+yaml to v0.0.9 to fix -root schema origin (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1144">#1144</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/6833fc11aa9dbe34e277a172739a1d10e5b65ffc"><code>6833fc1</code></a> README: drop go-openapi/spec3 (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1143">#1143</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/634db2c366fe5f97c225693d4f44a900eb7a6c34"><code>634db2c</code></a> openapi3: OriginTree approach for origin tracking — separate pass, no inline ...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/51a969a1286564d693d21a930b304acbea9840b6"><code>51a969a</code></a> fix: bump yaml and yaml3 to v0.0.4 (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1136">#1136</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/c132e59b59e6c17fef190c02aa0d0d211fb684bb"><code>c132e59</code></a> openapi3: strip <strong>origin</strong> from slices in example values (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1138">#1138</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/12ee8b094c0ee2d12185a80b3bbead173111be34"><code>12ee8b0</code></a> openapi3: strip <strong>origin</strong> from extension values to prevent spurious diffs (#...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/452c31b3264e68667a564cea2109ccb034b4cedb"><code>452c31b</code></a> fix: update yaml3 to prevent panic on empty mapping node in sequence (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1133">#1133</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.134.0...v0.135.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.98.0 to 1.99.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/159f4d23fd3150aa3b14c231c434f46f42b72c2a"><code>159f4d2</code></a> Release 2026-04-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50223dc5168c2eb6316a108f23a7fd0803afba76"><code>50223dc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/af25337605d0de4dc7f4141c6b1627ec31254415"><code>af25337</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/277dc017b3b8ea944865ffb901eea3bf82121f5d"><code>277dc01</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/22bc35ca5a2af77140a47680f80873b4d61b6551"><code>22bc35c</code></a> Release 2026-04-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a35fedd65d0cdad42281199a7659146234893a1"><code>9a35fed</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/867b085ecb8579affb591d0e98bbe362376ce911"><code>867b085</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/61d367c87617dbe79f6cb8fe8809d09ce949d2fa"><code>61d367c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8560a07e8b7ed78b0962b6df18050500a32705e7"><code>8560a07</code></a> Release 2026-04-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5aafdff76b34debb6940077fc4377ed80d24805"><code>a5aafdf</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.98.0...service/s3/v1.99.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.5.0 to 4.7.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/593d940f0fa70532c12a3730aac274d335005f28"><code>593d940</code></a> Update candlepin bindings to release/v4.7.5</li>
<li><a href="https://github.com/content-services/caliri/commit/83d6d68f043d58bbb98590695a307b8f5b3c2a59"><code>83d6d68</code></a> HMS-10377: update candlepin workflows</li>
<li><a href="https://github.com/content-services/caliri/commit/c9019307a801807774d629c8f015f80304110ba9"><code>c901930</code></a> Update candlepin bindings to release/v4.7.5</li>
<li><a href="https://github.com/content-services/caliri/commit/d99fd0be77057a0a65edcba260ec0e8c19232a85"><code>d99fd0b</code></a> HMS-10182: add codeowners</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.5.0...release/v4.7.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.3.1774545586 to 2026.4.1775140412
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/10c7ce5872712acd17c0b6bd59b6f2f0e43115a7"><code>10c7ce5</code></a> Update pulp bindings to 386d86254354e29d3b864523aed47832b5d5e8dc67968b5e2da9a...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.3.1774545586...release/v2026.4.1775140412">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.44.1 to 0.45.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.45.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Add support for Echo v5 by <a href="https://github.com/Scorfly"><code>@​Scorfly</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1183">#1183</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add OTLP trace exporter via new otel/otlp sub-module by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1229">#1229</a>
<ul>
<li>sentryotlp.NewTraceExporter sends OTel spans directly to Sentry's OTLP endpoint.</li>
<li>sentryotel.NewOtelIntegration links Sentry errors, logs, and metrics to the active OTel trace. Works with both direct-to-Sentry and collector-based setups.</li>
<li>NewSentrySpanProcessor, NewSentryPropagator, and SentrySpanMap are deprecated and will be removed in 0.47.0. To Migrate use <code>sentryotlp.NewTraceExporter</code> instead:</li>
</ul>
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
<li>Add IsSensitiveHeader helper to easily distinguish which headers to scrub for PII. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1239">#1239</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>(ci) Update validate-pr action to remove draft enforcement by <a href="https://github.com/stephanie-anderson"><code>@​stephanie-anderson</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1237">#1237</a></li>
<li>(fiber) Use UserContext for transaction to enable OTel trace linking by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1252">#1252</a></li>
<li>Race condition when getting envelope identifier by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1250">#1250</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<ul>
<li>Bump OpenTelemetry SDK to 1.40.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1243">#1243</a></li>
<li>Bump changelog-preview.yml from 2.24.1 to 2.25.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1247">#1247</a></li>
<li>Bump getsentry/craft from 2.24.1 to 2.25.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1248">#1248</a></li>
<li>Bump codecov/codecov-action from 5.5.2 to 6.0.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1245">#1245</a></li>
<li>Bump actions/create-github-app-token from 2.2.1 to 3.0.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1246">#1246</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.45.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Add support for Echo v5 by <a href="https://github.com/Scorfly"><code>@​Scorfly</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1183">#1183</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add OTLP trace exporter via new otel/otlp sub-module by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1229">#1229</a>
<ul>
<li>sentryotlp.NewTraceExporter sends OTel spans directly to Sentry's OTLP endpoint.</li>
<li>sentryotel.NewOtelIntegration links Sentry errors, logs, and metrics to the active OTel trace. Works with both direct-to-Sentry and collector-based setups.</li>
<li>NewSentrySpanProcessor, NewSentryPropagator, and SentrySpanMap are deprecated and will be removed in 0.47.0. To Migrate use <code>sentryotlp.NewTraceExporter</code> instead:</li>
</ul>
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
<li>Add IsSensitiveHeader helper to easily distinguish which headers to scrub for PII. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1239">#1239</a></li>
</ul>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>(ci) Update validate-pr action to remove draft enforcement by <a href="https://github.com/stephanie-anderson"><code>@​stephanie-anderson</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1237">#1237</a></li>
<li>(fiber) Use UserContext for transaction to enable OTel trace linking by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1252">#1252</a></li>
<li>Race condition when getting envelope identifier by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1250">#1250</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<ul>
<li>Bump OpenTelemetry SDK to 1.40.0 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1243">#1243</a></li>
<li>Bump changelog-preview.yml from 2.24.1 to 2.25.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1247">#1247</a></li>
<li>Bump getsentry/craft from 2.24.1 to 2.25.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1248">#1248</a></li>
<li>Bump codecov/codecov-action from 5.5.2 to 6.0.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1245">#1245</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/0f536b12ac10da5ddcccd189a7c245e0b0a0bafd"><code>0f536b1</code></a> release: 0.45.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c8ccbf348f4f0b25686c67d1f902ab5d7c4ae35a"><code>c8ccbf3</code></a> feat: add OTLP integration support (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1229">#1229</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e9ba5cbccfb86dffa0c7544274ac9ca21383e249"><code>e9ba5cb</code></a> fix(fiber): Use UserContext for transaction to enable OTel trace linking (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1252">#1252</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/18546ff0ee3496671c7812502fb5e4b6a28be7b4"><code>18546ff</code></a> test: add serialization panic reproduction (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1249">#1249</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e507be6f0cfb1654fafefa09090c79c2a1e21c5a"><code>e507be6</code></a> fix: race when getting envelope identifier (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1250">#1250</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2cec51a0c3289fc045c8f34c34f6dff2ca361b58"><code>2cec51a</code></a> feat!: add support for Echo v5 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1183">#1183</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/6ba06389e351fc94be995c2c59382575c88f1b82"><code>6ba0638</code></a> build(otel): Bump OpenTelemetry SDK to 1.40.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1243">#1243</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9c4eb663f085adcbeac2913caaf4e8fe52302552"><code>9c4eb66</code></a> build(deps): bump changelog-preview.yml from 2.24.1 to 2.25.2 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1247">#1247</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/b0a3a3080dddad5d5b546dfe287dc737e27a6036"><code>b0a3a30</code></a> build(deps): bump getsentry/craft from 2.24.1 to 2.25.2 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1248">#1248</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2f542592abc8f78eb22773e5e3b4d1f04c99ed1b"><code>2f54259</code></a> build(deps): bump codecov/codecov-action from 5.5.2 to 6.0.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1245">#1245</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.44.1...v0.45.0">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @TenSt - Approved on April 13, 2026 at 10:18 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1456*
