---
type: pull_request
number: 408
title: "Bump github.com/redis/go-redis/v9 from 9.2.0 to 9.2.1"
state: closed
author: dependabot
created: 2023-10-02T04:57:31Z
updated: 2023-10-02T13:59:02Z
closed: 2023-10-02T13:58:52Z
base_branch: main
head_branch: dependabot/go_modules/github.com/redis/go-redis/v9-9.2.1
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/408
---

# Pull Request #408: Bump github.com/redis/go-redis/v9 from 9.2.0 to 9.2.1

**Author**: @dependabot
**Created**: October 02, 2023 at 04:57 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/redis/go-redis/v9-9.2.1`

## Description

Bumps [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) from 9.2.0 to 9.2.1.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.2.1</h2>
<h2>🧰 Maintenance</h2>
<ul>
<li>chore(deps): bump actions/stale from 3 to 8 (<a href="https://redirect.github.com/redis/go-redis/issues/2732">#2732</a>)</li>
<li>Add stream interface back to <code>Cmdable</code> (<a href="https://redirect.github.com/redis/go-redis/issues/2725">#2725</a>)</li>
<li>Remove redundant nil check in gears (<a href="https://redirect.github.com/redis/go-redis/issues/2728">#2728</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/Juneezee"><code>@​Juneezee</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/gabrielgio"><code>@​gabrielgio</code></a> and <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/f994ff1cd96299a5c8029ae3403af7b17ef06e8a"><code>f994ff1</code></a> Bump version to 9.2.1 (<a href="https://redirect.github.com/redis/go-redis/issues/2735">#2735</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/0b5e1866b1b37cfafd7e5d403446207c0143b6ec"><code>0b5e186</code></a> chore(deps): bump actions/stale from 3 to 8 (<a href="https://redirect.github.com/redis/go-redis/issues/2732">#2732</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c6fe509f4a0c6746544a6320f6d1237e74424a7a"><code>c6fe509</code></a> Add stream interface back to <code>Cmdable</code> (<a href="https://redirect.github.com/redis/go-redis/issues/2725">#2725</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/275af739718caca7af07dde35c234517d70b6b07"><code>275af73</code></a> refactor(gears): remove redundant nil check (<a href="https://redirect.github.com/redis/go-redis/issues/2728">#2728</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.2.0...v9.2.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/redis/go-redis/v9&package-manager=go_modules&previous-version=9.2.0&new-version=9.2.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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


</details>

---

## Discussion

### Comment by @app-sre-bot on October 02, 2023 at 04:58 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on October 02, 2023 at 01:58 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/408*
