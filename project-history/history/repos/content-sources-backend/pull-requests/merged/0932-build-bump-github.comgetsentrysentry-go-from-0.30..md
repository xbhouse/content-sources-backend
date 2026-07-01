---
type: pull_request
number: 932
title: "Build: Bump github.com/getsentry/sentry-go from 0.30.0 to 0.31.1 in the go group"
state: merged
author: dependabot
created: 2025-01-06T04:33:42Z
updated: 2025-01-06T14:14:18Z
closed: 2025-01-06T14:14:13Z
merged: 2025-01-06T14:14:13Z
base_branch: main
head_branch: dependabot/go_modules/go-b9f841e723
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/932
---

# Pull Request #932: Build: Bump github.com/getsentry/sentry-go from 0.30.0 to 0.31.1 in the go group

**Author**: @dependabot
**Created**: January 06, 2025 at 04:33 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-b9f841e723`

## Description

Bumps the go group with 1 update: [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go).

Updates `github.com/getsentry/sentry-go` from 0.30.0 to 0.31.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.31.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.31.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Correct wrong module name for <code>sentry-go/logrus</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/950">#950</a>)</li>
</ul>
<h2>0.31.0</h2>
<h3>Breaking Changes</h3>
<ul>
<li>
<p>Remove support for metrics. Read more about the end of the Metrics beta <a href="https://sentry.zendesk.com/hc/en-us/articles/26369339769883-Metrics-Beta-Ended-on-October-7th">here</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/914">#914</a>)</p>
</li>
<li>
<p>Remove support for profiling. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/915">#915</a>)</p>
</li>
<li>
<p>Remove <code>Segment</code> field from the <code>User</code> struct. This field is no longer used in the Sentry product. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/928">#928</a>)</p>
</li>
<li>
<p>Every integration is now a separate module, reducing the binary size and number of dependencies. Once you update <code>sentry-go</code> to latest version, you'll need to <code>go get</code> the integration you want to use. For example, if you want to use the <code>echo</code> integration, you'll need to run <code>go get github.com/getsentry/sentry-go/echo</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/919">#919</a>).</p>
</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add the ability to override <code>hub</code> in <code>context</code> for integrations that use custom context. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/931">#931</a>)</p>
</li>
<li>
<p>Add <code>HubProvider</code> Hook for <code>sentrylogrus</code>, enabling dynamic Sentry hub allocation for each log entry or goroutine. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/936">#936</a>)</p>
</li>
</ul>
<p>This change enhances compatibility with Sentry's recommendation of using separate hubs per goroutine. To ensure a separate Sentry hub for each goroutine, configure the <code>HubProvider</code> like this:</p>
<pre lang="go"><code>hook, err := sentrylogrus.New(nil, sentry.ClientOptions{})
if err != nil {
    log.Fatalf(&quot;Failed to initialize Sentry hook: %v&quot;, err)
}
<p>// Set a custom HubProvider to generate a new hub for each goroutine or log entry<br />
hook.SetHubProvider(func() *sentry.Hub {<br />
client, _ := sentry.NewClient(sentry.ClientOptions{})<br />
return sentry.NewHub(client, sentry.NewScope())<br />
})</p>
<p>logrus.AddHook(hook)<br />
</code></pre></p>
<h3>Bug Fixes</h3>
<ul>
<li>Add support for closing worker goroutines started by the <code>HTTPTranport</code> to prevent goroutine leaks. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/894">#894</a>)</li>
</ul>
<pre lang="go"><code>client, _ := sentry.NewClient()
defer client.Close()
</code></pre>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.31.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.31.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Correct wrong module name for <code>sentry-go/logrus</code> (<a href="https://redirect.github.com/getsentry/sentry-go/pull/950">#950</a>)</li>
</ul>
<h2>0.31.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.31.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>
<p>Remove support for metrics. Read more about the end of the Metrics beta <a href="https://sentry.zendesk.com/hc/en-us/articles/26369339769883-Metrics-Beta-Ended-on-October-7th">here</a>. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/914">#914</a>)</p>
</li>
<li>
<p>Remove support for profiling. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/915">#915</a>)</p>
</li>
<li>
<p>Remove <code>Segment</code> field from the <code>User</code> struct. This field is no longer used in the Sentry product. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/928">#928</a>)</p>
</li>
<li>
<p>Every integration is now a separate module, reducing the binary size and number of dependencies. Once you update <code>sentry-go</code> to latest version, you'll need to <code>go get</code> the integration you want to use. For example, if you want to use the <code>echo</code> integration, you'll need to run <code>go get github.com/getsentry/sentry-go/echo</code> (<a href="https://redirect.github.com/getsentry/sentry-go/blob/master/redirect.github.com/getsentry/sentry-go/pull/919">#919</a>).</p>
</li>
</ul>
<h3>Features</h3>
<p>Add the ability to override <code>hub</code> in <code>context</code> for integrations that use custom context. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/931">#931</a>)</p>
<ul>
<li>Add <code>HubProvider</code> Hook for <code>sentrylogrus</code>, enabling dynamic Sentry hub allocation for each log entry or goroutine. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/936">#936</a>)</li>
</ul>
<p>This change enhances compatibility with Sentry's recommendation of using separate hubs per goroutine. To ensure a separate Sentry hub for each goroutine, configure the <code>HubProvider</code> like this:</p>
<pre lang="go"><code>hook, err := sentrylogrus.New(nil, sentry.ClientOptions{})
if err != nil {
    log.Fatalf(&quot;Failed to initialize Sentry hook: %v&quot;, err)
}
<p>// Set a custom HubProvider to generate a new hub for each goroutine or log entry<br />
hook.SetHubProvider(func() *sentry.Hub {<br />
client, _ := sentry.NewClient(sentry.ClientOptions{})<br />
return sentry.NewHub(client, sentry.NewScope())<br />
})</p>
<p>logrus.AddHook(hook)<br />
</code></pre></p>
<h3>Bug Fixes</h3>
<ul>
<li>Add support for closing worker goroutines started by the <code>HTTPTranport</code> to prevent goroutine leaks. (<a href="https://redirect.github.com/getsentry/sentry-go/pull/894">#894</a>)</li>
</ul>
<pre lang="go"><code>&lt;/tr&gt;&lt;/table&gt; 
</code></pre>
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/454d469ef94c6b081a560a14fc1fa24d1f12036d"><code>454d469</code></a> Fix lint CI</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/82b1d7a14a6ebdfb624db1d0f632209e2e986bb1"><code>82b1d7a</code></a> release: 0.31.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/08c2b0f0790cb40b8324a4fb9f946af40a50c23c"><code>08c2b0f</code></a> Prepare 0.31.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/951">#951</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/e0b4d38b61cf062effdb2777585c7e40589a0cc7"><code>e0b4d38</code></a> Rename <code>sentry-go/logrus</code> module (<a href="https://redirect.github.com/getsentry/sentry-go/issues/950">#950</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/96cc6a85e202506e2e318d24a4ca828cc1762740"><code>96cc6a8</code></a> Merge branch 'release/0.31.0'</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/17ceb6c8e551c7d06f5c7b72f8ad03935a96ca9e"><code>17ceb6c</code></a> Bump golangci-lint</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9f9f465f0ef7b31747e2770f8d776e82428f1bd0"><code>9f9f465</code></a> release: 0.31.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/934eca557f6d48efc44540de0aba307ecbb0fc17"><code>934eca5</code></a> Prepare 0.31.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/948">#948</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/066c69e96174b1b181cf6da965657efed7ecc936"><code>066c69e</code></a> Add HubProvider functionality to the logrus integration (<a href="https://redirect.github.com/getsentry/sentry-go/issues/936">#936</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/486c78c4ed06346275816527ae69cbed404cd850"><code>486c78c</code></a> Bump golang.org/x/net to 0.33.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/947">#947</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.30.0...v0.31.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/getsentry/sentry-go&package-manager=go_modules&previous-version=0.30.0&new-version=0.31.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Review by @jlsherrill - Approved on January 06, 2025 at 02:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/932*
