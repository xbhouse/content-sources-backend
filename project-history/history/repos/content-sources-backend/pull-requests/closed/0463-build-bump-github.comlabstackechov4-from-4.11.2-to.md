---
type: pull_request
number: 463
title: "Build: Bump github.com/labstack/echo/v4 from 4.11.2 to 4.11.3"
state: closed
author: dependabot
created: 2023-11-13T04:25:17Z
updated: 2023-11-13T04:28:49Z
closed: 2023-11-13T04:28:47Z
base_branch: main
head_branch: dependabot/go_modules/github.com/labstack/echo/v4-4.11.3
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/463
---

# Pull Request #463: Build: Bump github.com/labstack/echo/v4 from 4.11.2 to 4.11.3

**Author**: @dependabot
**Created**: November 13, 2023 at 04:25 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/labstack/echo/v4-4.11.3`

## Description

Bumps [github.com/labstack/echo/v4](https://github.com/labstack/echo) from 4.11.2 to 4.11.3.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/releases">github.com/labstack/echo/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.11.3</h2>
<p><strong>Security</strong></p>
<ul>
<li>'c.Attachment' and 'c.Inline' should escape filename in 'Content-Disposition' header to avoid 'Reflect File Download' vulnerability. <a href="https://redirect.github.com/labstack/echo/pull/2541">#2541</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Tests: refactor context tests to be separate functions <a href="https://redirect.github.com/labstack/echo/pull/2540">#2540</a></li>
<li>Proxy middleware: reuse echo request context <a href="https://redirect.github.com/labstack/echo/pull/2537">#2537</a></li>
<li>Mark unmarshallable yaml struct tags as ignored <a href="https://redirect.github.com/labstack/echo/pull/2536">#2536</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/master/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.11.3 - 2023-11-07</h2>
<p><strong>Security</strong></p>
<ul>
<li>'c.Attachment' and 'c.Inline' should escape filename in 'Content-Disposition' header to avoid 'Reflect File Download' vulnerability. <a href="https://redirect.github.com/labstack/echo/pull/2541">#2541</a></li>
</ul>
<p><strong>Enhancements</strong></p>
<ul>
<li>Tests: refactor context tests to be separate functions <a href="https://redirect.github.com/labstack/echo/pull/2540">#2540</a></li>
<li>Proxy middleware: reuse echo request context <a href="https://redirect.github.com/labstack/echo/pull/2537">#2537</a></li>
<li>Mark unmarshallable yaml struct tags as ignored <a href="https://redirect.github.com/labstack/echo/pull/2536">#2536</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/4b26cde851bc7a51e624c04dcc5d37be1ce0c84f"><code>4b26cde</code></a> Changelog for v4.11.3 (<a href="https://redirect.github.com/labstack/echo/issues/2542">#2542</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/14daeb968049b71296a80b91abd3883afd02b4d1"><code>14daeb9</code></a> Security: c.Attachment and c.Inline should escape name in `Content-Dispositio...</li>
<li><a href="https://github.com/labstack/echo/commit/50ebcd8d7c17457489df7bcbbcaa3745c687fd32"><code>50ebcd8</code></a> refactor context tests to be separate functions (<a href="https://redirect.github.com/labstack/echo/issues/2540">#2540</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/c7d6d4373fdfbef5d6f44df0a8ef410c198420ee"><code>c7d6d43</code></a> proxy middleware: reuse echo request context (<a href="https://redirect.github.com/labstack/echo/issues/2537">#2537</a>)</li>
<li><a href="https://github.com/labstack/echo/commit/69a0de84158fd7cad326599d145c2248bcc15a69"><code>69a0de8</code></a> Mark unmarshallable yaml struct tags as ignored (<a href="https://redirect.github.com/labstack/echo/issues/2536">#2536</a>)</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.11.2...v4.11.3">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/labstack/echo/v4&package-manager=go_modules&previous-version=4.11.2&new-version=4.11.3)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

## Discussion

### Comment by @app-sre-bot on November 13, 2023 at 04:25 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on November 13, 2023 at 04:28 AM UTC

Superseded by #466.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/463*
