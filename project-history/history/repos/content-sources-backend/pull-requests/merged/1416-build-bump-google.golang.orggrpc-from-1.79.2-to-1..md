---
type: pull_request
number: 1416
title: "Build: Bump google.golang.org/grpc from 1.79.2 to 1.79.3"
state: merged
author: dependabot
created: 2026-03-19T02:24:07Z
updated: 2026-03-19T08:23:26Z
closed: 2026-03-19T08:23:23Z
merged: 2026-03-19T08:23:23Z
base_branch: main
head_branch: dependabot/go_modules/google.golang.org/grpc-1.79.3
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1416
---

# Pull Request #1416: Build: Bump google.golang.org/grpc from 1.79.2 to 1.79.3

**Author**: @dependabot
**Created**: March 19, 2026 at 02:24 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/google.golang.org/grpc-1.79.3`

## Description

Bumps [google.golang.org/grpc](https://github.com/grpc/grpc-go) from 1.79.2 to 1.79.3.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.79.3</h2>
<h1>Security</h1>
<ul>
<li>server: fix an authorization bypass where malformed :path headers (missing the leading slash) could bypass path-based restricted &quot;deny&quot; rules in interceptors like <code>grpc/authz</code>. Any request with a non-canonical path is now immediately rejected with an <code>Unimplemented</code> error. (<a href="https://redirect.github.com/grpc/grpc-go/issues/8981">#8981</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/dda86dbd9cecb8b35b58c73d507d81d67761205f"><code>dda86db</code></a> Change version to 1.79.3 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8983">#8983</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/72186f163e75a065c39e6f7df9b6dea07fbdeff5"><code>72186f1</code></a> grpc: enforce strict path checking for incoming requests on the server (<a href="https://redirect.github.com/grpc/grpc-go/issues/8981">#8981</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/97ca3522b239edf6813e2b1106924e9d55e89d43"><code>97ca352</code></a> Changing version to 1.79.3-dev (<a href="https://redirect.github.com/grpc/grpc-go/issues/8954">#8954</a>)</li>
<li>See full diff in <a href="https://github.com/grpc/grpc-go/compare/v1.79.2...v1.79.3">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=google.golang.org/grpc&package-manager=go_modules&previous-version=1.79.2&new-version=1.79.3)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Reviews

### Review by @TenSt - Approved on March 19, 2026 at 08:23 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1416*
