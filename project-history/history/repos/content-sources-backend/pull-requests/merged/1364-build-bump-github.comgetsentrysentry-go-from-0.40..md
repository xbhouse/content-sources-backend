---
type: pull_request
number: 1364
title: "Build: Bump github.com/getsentry/sentry-go from 0.40.0 to 0.41.0"
state: merged
author: dependabot
created: 2026-01-19T04:58:00Z
updated: 2026-01-19T14:27:48Z
closed: 2026-01-19T14:27:47Z
merged: 2026-01-19T14:27:47Z
base_branch: main
head_branch: dependabot/go_modules/go-f302018b9d
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1364
---

# Pull Request #1364: Build: Bump github.com/getsentry/sentry-go from 0.40.0 to 0.41.0

**Author**: @dependabot
**Created**: January 19, 2026 at 04:58 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-f302018b9d`

## Description

Bumps the go group with 1 update: [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go).

Updates `github.com/getsentry/sentry-go` from 0.40.0 to 0.41.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.41.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.41.0.</p>
<h3>Features</h3>
<ul>
<li>Add HTTP client integration for distributed tracing via <code>sentryhttpclient</code> package (<a href="https://redirect.github.com/getsentry/sentry-go/pull/876">#876</a>)
<ul>
<li>Provides an <code>http.RoundTripper</code> implementation that automatically creates spans for outgoing HTTP requests</li>
<li>Supports trace propagation targets configuration via <code>WithTracePropagationTargets</code> option</li>
<li>Example usage:
<pre lang="go"><code>import sentryhttpclient &quot;github.com/getsentry/sentry-go/httpclient&quot;
<p>roundTripper := sentryhttpclient.NewSentryRoundTripper(nil)<br />
client := &amp;http.Client{<br />
Transport: roundTripper,<br />
}<br />
</code></pre></p>
</li>
</ul>
</li>
<li>Add <code>ClientOptions.PropagateTraceparent</code> option to control W3C <code>traceparent</code> header propagation in outgoing HTTP requests (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1161">#1161</a>)</li>
<li>Add <code>SpanID</code> field to structured logs (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1169">#1169</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.41.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.41.0.</p>
<h3>Features</h3>
<ul>
<li>Add HTTP client integration for distributed tracing via <code>sentryhttpclient</code> package (<a href="https://redirect.github.com/getsentry/sentry-go/pull/876">#876</a>)
<ul>
<li>Provides an <code>http.RoundTripper</code> implementation that automatically creates spans for outgoing HTTP requests</li>
<li>Supports trace propagation targets configuration via <code>WithTracePropagationTargets</code> option</li>
<li>Example usage:
<pre lang="go"><code>import sentryhttpclient &quot;github.com/getsentry/sentry-go/httpclient&quot;
<p>roundTripper := sentryhttpclient.NewSentryRoundTripper(nil)<br />
client := &amp;http.Client{<br />
Transport: roundTripper,<br />
}<br />
</code></pre></p>
</li>
</ul>
</li>
<li>Add <code>ClientOptions.PropagateTraceparent</code> option to control W3C <code>traceparent</code> header propagation in outgoing HTTP requests (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1161">#1161</a>)</li>
<li>Add <code>SpanID</code> field to structured logs (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1169">#1169</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/87e197c2a57d353c1c65395d6c19fc0d22dec084"><code>87e197c</code></a> release: 0.41.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a130ff15f9c6f87ea5551a03fc3e820cbeca9976"><code>a130ff1</code></a> chore: prepare 0.41.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1172">#1172</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f3a5a3fd378106ca6616f91aba75e06286e3738a"><code>f3a5a3f</code></a> feat: refactor log serialization to match docs (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1169">#1169</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/736c662dfb1e8010bc725c63b3d82a7229a7d99e"><code>736c662</code></a> build(deps): bump actions/create-github-app-token from 2.1.4 to 2.2.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1164">#1164</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/30c8b2ebd69b240a0b1a7b55584d129a8480592d"><code>30c8b2e</code></a> build(deps): bump golangci/golangci-lint-action from 8.0.0 to 9.2.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1165">#1165</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/335c5f12f366017ad662c24c15c41503367bf0db"><code>335c5f1</code></a> build(deps): bump codecov/codecov-action from 5.5.1 to 5.5.2 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1163">#1163</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/77f5f7dbbf872092fa5eba95fed31e04963c0f01"><code>77f5f7d</code></a> build(deps): bump actions/checkout from 5 to 6 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1162">#1162</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f5cdcd9b1abd85d2e7230b1e07d3b38f1248950d"><code>f5cdcd9</code></a> build(deps): bump actions/cache from 4 to 5 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1166">#1166</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/8cb62cff2b2022aeb85d2e4f1194ddaac5a6d5df"><code>8cb62cf</code></a> feat: add propagateTraceparent option (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1161">#1161</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d7582e8848ca8fdc8aae4a8c189b83be1c9e2bd1"><code>d7582e8</code></a> feat: add http client integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/876">#876</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.40.0...v0.41.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/getsentry/sentry-go&package-manager=go_modules&previous-version=0.40.0&new-version=0.41.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Review by @rverdile - Approved on January 19, 2026 at 02:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1364*
