---
type: pull_request
number: 1083
title: "Build: Bump golang.org/x/net from 0.35.0 to 0.38.0"
state: closed
author: dependabot
created: 2025-04-16T23:12:54Z
updated: 2025-04-24T19:29:04Z
closed: 2025-04-24T19:28:55Z
base_branch: main
head_branch: dependabot/go_modules/golang.org/x/net-0.38.0
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1083
---

# Pull Request #1083: Build: Bump golang.org/x/net from 0.35.0 to 0.38.0

**Author**: @dependabot
**Created**: April 16, 2025 at 11:12 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/golang.org/x/net-0.38.0`

## Description

Bumps [golang.org/x/net](https://github.com/golang/net) from 0.35.0 to 0.38.0.
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang/net/commit/e1fcd82abba34df74614020343be8eb1fe85f0d9"><code>e1fcd82</code></a> html: properly handle trailing solidus in unquoted attribute value in foreign...</li>
<li><a href="https://github.com/golang/net/commit/ebed060e8f30f20235f74808c22125fd86b15edd"><code>ebed060</code></a> internal/http3: fix build of tests with GOEXPERIMENT=nosynctest</li>
<li><a href="https://github.com/golang/net/commit/1f1fa29e0a46fffe18c43a9da8daa5a0b180dfa9"><code>1f1fa29</code></a> publicsuffix: regenerate table</li>
<li><a href="https://github.com/golang/net/commit/12150816f701c912a32a376754ab28dd3878833a"><code>1215081</code></a> http2: improve error when server sends HTTP/1</li>
<li><a href="https://github.com/golang/net/commit/312450e473eae9f9e6173ad895c80bc5ea2f79ad"><code>312450e</code></a> html: ensure &lt;search&gt; tag closes &lt;p&gt; and update tests</li>
<li><a href="https://github.com/golang/net/commit/09731f9bf919b00b344c763894cd1920b3d96d90"><code>09731f9</code></a> http2: improve handling of lost PING in Server</li>
<li><a href="https://github.com/golang/net/commit/55989e24b972a90ab99308fdc7ea1fb58a96fef1"><code>55989e2</code></a> http2/h2c: use ResponseController for hijacking connections</li>
<li><a href="https://github.com/golang/net/commit/2914f46773171f4fa13e276df1135bafef677801"><code>2914f46</code></a> websocket: re-recommend gorilla/websocket</li>
<li><a href="https://github.com/golang/net/commit/99b3ae0643f9a2f9d820fcbba5f9e4c83b23bd48"><code>99b3ae0</code></a> go.mod: update golang.org/x dependencies</li>
<li><a href="https://github.com/golang/net/commit/85d1d54551b68719346cb9fec24b911da4e452a1"><code>85d1d54</code></a> go.mod: update golang.org/x dependencies</li>
<li>Additional commits viewable in <a href="https://github.com/golang/net/compare/v0.35.0...v0.38.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=golang.org/x/net&package-manager=go_modules&previous-version=0.35.0&new-version=0.38.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

## Discussion

### Comment by @rverdile on April 24, 2025 at 07:28 PM UTC

closing because required upgrade to 1.23

### Comment by @dependabot on April 24, 2025 at 07:28 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1083*
