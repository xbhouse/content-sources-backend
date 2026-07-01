---
type: pull_request
number: 1142
title: "Build: Bump github.com/getsentry/sentry-go 0.34.0 > 0.34.1 in go group"
state: merged
author: dependabot
created: 2025-07-14T05:01:37Z
updated: 2025-07-14T12:34:20Z
closed: 2025-07-14T12:34:11Z
merged: 2025-07-14T12:34:11Z
base_branch: main
head_branch: dependabot/go_modules/go-b04adf4e6f
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1142
---

# Pull Request #1142: Build: Bump github.com/getsentry/sentry-go 0.34.0 > 0.34.1 in go group

**Author**: @dependabot
**Created**: July 14, 2025 at 05:01 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-b04adf4e6f`

## Description

Bumps the go group with 1 update: [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go).

Updates `github.com/getsentry/sentry-go` from 0.34.0 to 0.34.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.34.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.34.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Allow flush to be used multiple times without issues, particularly for the batch logger (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1051">#1051</a>)</li>
<li>Fix race condition in <code>Scope.GetSpan()</code> method by adding proper mutex locking (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1044">#1044</a>)</li>
<li>Guard transport on <code>Close()</code> to prevent panic when called multiple times (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1044">#1044</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.34.1</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.34.1.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Allow flush to be used multiple times without issues, particularly for the batch logger (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1051">#1051</a>)</li>
<li>Fix race condition in <code>Scope.GetSpan()</code> method by adding proper mutex locking (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1044">#1044</a>)</li>
<li>Guard transport on <code>Close()</code> to prevent panic when called multiple times (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1044">#1044</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/e32fa329b491127c1d4feb30f749655a9d204b0c"><code>e32fa32</code></a> release: 0.34.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/79f9cbe46ccd65f7fd432328e55b13f788f2542e"><code>79f9cbe</code></a> Prepare 0.34.1 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1052">#1052</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/26683315ec9c94890226f320ae159d682ffad54e"><code>2668331</code></a> Allow flush to be used multiple times (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1051">#1051</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/221b2acbb88dfe7a525a8a9896d5d62403a96552"><code>221b2ac</code></a> Fix race condition on GetSpan and guard transport on Close (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1044">#1044</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/d20dadde3bc47b6dc9f232286cc70e50fd114e9a"><code>d20dadd</code></a> Merge branch 'release/0.34.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.34.0...v0.34.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/getsentry/sentry-go&package-manager=go_modules&previous-version=0.34.0&new-version=0.34.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Review by @swadeley - Approved on July 14, 2025 at 12:34 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1142*
