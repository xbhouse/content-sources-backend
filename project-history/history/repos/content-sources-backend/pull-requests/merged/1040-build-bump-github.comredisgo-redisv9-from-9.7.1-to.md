---
type: pull_request
number: 1040
title: "Build: Bump github.com/redis/go-redis/v9 from 9.7.1 to 9.7.3"
state: merged
author: dependabot
created: 2025-03-20T18:54:42Z
updated: 2025-03-20T20:06:20Z
closed: 2025-03-20T20:06:12Z
merged: 2025-03-20T20:06:12Z
base_branch: main
head_branch: dependabot/go_modules/github.com/redis/go-redis/v9-9.7.3
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1040
---

# Pull Request #1040: Build: Bump github.com/redis/go-redis/v9 from 9.7.1 to 9.7.3

**Author**: @dependabot
**Created**: March 20, 2025 at 06:54 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/redis/go-redis/v9-9.7.3`

## Description

Bumps [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) from 9.7.1 to 9.7.3.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>v9.7.3</h2>
<h2>What's Changed</h2>
<ul>
<li>fix: handle network error on SETINFO (<a href="https://redirect.github.com/redis/go-redis/issues/3295">#3295</a>) (<a href="https://github.com/redis/go-redis/security/advisories/GHSA-92cp-5422-2mw7">CVE-2025-29923</a>)</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.7.1...v9.7.3">https://github.com/redis/go-redis/compare/v9.7.1...v9.7.3</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/a29d91d9ca72a9c26708fba7dc9872f339a3549e"><code>a29d91d</code></a> release 9.7.3, retract 9.7.2 (<a href="https://redirect.github.com/redis/go-redis/issues/3314">#3314</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/ce3034c7b3ee8eeac5b6ac63a33e51db3602cf34"><code>ce3034c</code></a> bump version to 9.7.2</li>
<li><a href="https://github.com/redis/go-redis/commit/0af2b32f9369d81e900e32907b8c1afb1e5d502d"><code>0af2b32</code></a> fix: handle network error on SETINFO (<a href="https://redirect.github.com/redis/go-redis/issues/3295">#3295</a>) (CVE-2025-29923)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.7.1...v9.7.3">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/redis/go-redis/v9&package-manager=go_modules&previous-version=9.7.1&new-version=9.7.3)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Reviews

### Review by @jlsherrill - Approved on March 20, 2025 at 08:06 PM UTC

ack, playwright error is unrelated

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1040*
