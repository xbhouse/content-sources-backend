---
type: pull_request
number: 1082
title: "Build: Bump golang.org/x/crypto from 0.33.0 to 0.35.0"
state: closed
author: dependabot
created: 2025-04-14T17:01:29Z
updated: 2025-04-16T17:23:00Z
closed: 2025-04-16T17:22:50Z
base_branch: main
head_branch: dependabot/go_modules/golang.org/x/crypto-0.35.0
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1082
---

# Pull Request #1082: Build: Bump golang.org/x/crypto from 0.33.0 to 0.35.0

**Author**: @dependabot
**Created**: April 14, 2025 at 05:01 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/golang.org/x/crypto-0.35.0`

## Description

Bumps [golang.org/x/crypto](https://github.com/golang/crypto) from 0.33.0 to 0.35.0.
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang/crypto/commit/7292932d45d55c7199324ab0027cc86e8198aa22"><code>7292932</code></a> ssh: limit the size of the internal packet queue while waiting for KEX</li>
<li><a href="https://github.com/golang/crypto/commit/f66f74b0a406b5f6909183531ace593857f1646c"><code>f66f74b</code></a> acme/autocert: check host policy before probing the cache</li>
<li><a href="https://github.com/golang/crypto/commit/b0784b7bfbe0b2c9a59afc1248ed3cb4b6652e85"><code>b0784b7</code></a> x509roots/fallback: drop obsolete build constraint</li>
<li><a href="https://github.com/golang/crypto/commit/911360c8a4f464342b9fe7c23632be57fca87b20"><code>911360c</code></a> all: bump golang.org/x/crypto dependencies of asm generators</li>
<li><a href="https://github.com/golang/crypto/commit/89ff08d67c4d79f9ac619aaf1f7388888798651f"><code>89ff08d</code></a> all: upgrade go directive to at least 1.23.0 [generated]</li>
<li><a href="https://github.com/golang/crypto/commit/e47973b1c1089f6c67ab89261f7aa067b3d611d2"><code>e47973b</code></a> all: update certs for go1.24</li>
<li>See full diff in <a href="https://github.com/golang/crypto/compare/v0.33.0...v0.35.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=golang.org/x/crypto&package-manager=go_modules&previous-version=0.33.0&new-version=0.35.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @rverdile on April 16, 2025 at 05:22 PM UTC

closing because this would require us to upgrade to go 1.23

### Comment by @dependabot on April 16, 2025 at 05:22 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1082*
