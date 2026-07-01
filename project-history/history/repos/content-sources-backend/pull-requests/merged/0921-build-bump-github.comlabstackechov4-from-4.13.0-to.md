---
type: pull_request
number: 921
title: "Build: Bump github.com/labstack/echo/v4 from 4.13.0 to 4.13.2"
state: merged
author: dependabot
created: 2024-12-16T04:54:09Z
updated: 2024-12-16T13:57:15Z
closed: 2024-12-16T13:57:11Z
merged: 2024-12-16T13:57:11Z
base_branch: main
head_branch: dependabot/go_modules/go-ca8e3cba94
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/921
---

# Pull Request #921: Build: Bump github.com/labstack/echo/v4 from 4.13.0 to 4.13.2

**Author**: @dependabot
**Created**: December 16, 2024 at 04:54 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-ca8e3cba94`

## Description

Bumps the go group with 1 update: [github.com/labstack/echo/v4](https://github.com/labstack/echo).

Updates `github.com/labstack/echo/v4` from 4.13.0 to 4.13.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.13.2 - update dependencies</h2>
<p><strong>Security</strong></p>
<ul>
<li>Update dependencies (dependabot reports <a href="https://pkg.go.dev/vuln/GO-2024-3321">https://pkg.go.dev/vuln/GO-2024-3321</a> by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2721">labstack/echo#2721</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.13.1...v4.13.2">https://github.com/labstack/echo/compare/v4.13.1...v4.13.2</a></p>
<h2>v4.13.1</h2>
<p><strong>Fixes</strong></p>
<ul>
<li>Fix BindBody ignoring <code>Transfer-Encoding: chunked</code> requests (introduced in <a href="https://redirect.github.com/labstack/echo/pull/2710">#2710</a>) by <a href="https://github.com/178inaba"><code>@​178inaba</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2717">labstack/echo#2717</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/labstack/echo/compare/v4.13.0...v4.13.1">https://github.com/labstack/echo/compare/v4.13.0...v4.13.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.13.2 - 2024-12-12</h2>
<p><strong>Security</strong></p>
<ul>
<li>Update dependencies (dependabot reports <a href="https://pkg.go.dev/vuln/GO-2024-3321">GO-2024-3321</a>) in <a href="https://redirect.github.com/labstack/echo/pull/2721">labstack/echo#2721</a></li>
</ul>
<h2>v4.13.1 - 2024-12-11</h2>
<p><strong>Fixes</strong></p>
<ul>
<li>Fix BindBody ignoring <code>Transfer-Encoding: chunked</code> requests by <a href="https://github.com/178inaba"><code>@​178inaba</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2717">labstack/echo#2717</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/692bc2abb6e83be2efe45f89f142f59dbca3a6d9"><code>692bc2a</code></a> Update dependencies (dependabot reports <a href="https://pkg.go.dev/vuln/GO-2024-3321">https://pkg.go.dev/vuln/GO-2024-3321</a>)...</li>
<li><a href="https://github.com/labstack/echo/commit/fd3f07447eddb53e096b0df3be38bbf79123d605"><code>fd3f074</code></a> Changelog for v4.13.1 (<a href="https://redirect.github.com/labstack/echo/issues/2719">#2719</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/0368ed87f28eb07e7c5cca2dbff5585a019b4fa1"><code>0368ed8</code></a> Add Conditions to Ensure Bind Succeeds with <code>Transfer-Encoding: chunked</code> (<a href="https://redirect.github.com/labstack/echo/issues/2717">#2717</a>)</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.13.0...v4.13.2">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/labstack/echo/v4&package-manager=go_modules&previous-version=4.13.0&new-version=4.13.2)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Review by @jlsherrill - Approved on December 16, 2024 at 01:57 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/921*
