---
type: pull_request
number: 1232
title: "Bump github.com/gin-gonic/gin from 1.9.0 to 1.9.1"
state: merged
author: dependabot
created: 2023-06-01T20:36:10Z
updated: 2023-06-02T09:07:11Z
closed: 2023-06-02T09:07:03Z
merged: 2023-06-02T09:07:03Z
base_branch: master
head_branch: dependabot/go_modules/github.com/gin-gonic/gin-1.9.1
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-engine/pull/1232
---

# Pull Request #1232: Bump github.com/gin-gonic/gin from 1.9.0 to 1.9.1

**Author**: @dependabot
**Created**: June 01, 2023 at 08:36 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/go_modules/github.com/gin-gonic/gin-1.9.1`

## Description

Bumps [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) from 1.9.0 to 1.9.1.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/gin-gonic/gin/releases">github.com/gin-gonic/gin's releases</a>.</em></p>
<blockquote>
<h2>v1.9.1</h2>
<h2>Changelog</h2>
<h3>BUG FIXES</h3>
<ul>
<li>fix Request.Context() checks <a href="https://redirect.github.com/gin-gonic/gin/pull/3512">#3512</a></li>
</ul>
<h3>SECURITY</h3>
<ul>
<li>fix lack of escaping of filename in Content-Disposition <a href="https://redirect.github.com/gin-gonic/gin/pull/3556">#3556</a></li>
</ul>
<h3>ENHANCEMENTS</h3>
<ul>
<li>refactor: use bytes.ReplaceAll directly <a href="https://redirect.github.com/gin-gonic/gin/pull/3455">#3455</a></li>
<li>convert strings and slices using the officially recommended way <a href="https://redirect.github.com/gin-gonic/gin/pull/3344">#3344</a></li>
<li>improve render code coverage <a href="https://redirect.github.com/gin-gonic/gin/pull/3525">#3525</a></li>
</ul>
<h3>DOCS</h3>
<ul>
<li>docs: changed documentation link for trusted proxies <a href="https://redirect.github.com/gin-gonic/gin/pull/3575">#3575</a></li>
<li>chore: improve linting, testing, and GitHub Actions setup <a href="https://redirect.github.com/gin-gonic/gin/pull/3583">#3583</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/gin-gonic/gin/blob/master/CHANGELOG.md">github.com/gin-gonic/gin's changelog</a>.</em></p>
<blockquote>
<h2>Gin v1.9.1</h2>
<h3>BUG FIXES</h3>
<ul>
<li>fix Request.Context() checks <a href="https://redirect.github.com/gin-gonic/gin/pull/3512">#3512</a></li>
</ul>
<h3>SECURITY</h3>
<ul>
<li>fix lack of escaping of filename in Content-Disposition <a href="https://redirect.github.com/gin-gonic/gin/pull/3556">#3556</a></li>
</ul>
<h3>ENHANCEMENTS</h3>
<ul>
<li>refactor: use bytes.ReplaceAll directly <a href="https://redirect.github.com/gin-gonic/gin/pull/3455">#3455</a></li>
<li>convert strings and slices using the officially recommended way <a href="https://redirect.github.com/gin-gonic/gin/pull/3344">#3344</a></li>
<li>improve render code coverage <a href="https://redirect.github.com/gin-gonic/gin/pull/3525">#3525</a></li>
</ul>
<h3>DOCS</h3>
<ul>
<li>docs: changed documentation link for trusted proxies <a href="https://redirect.github.com/gin-gonic/gin/pull/3575">#3575</a></li>
<li>chore: improve linting, testing, and GitHub Actions setup <a href="https://redirect.github.com/gin-gonic/gin/pull/3583">#3583</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/gin-gonic/gin/commit/4ea0e648e38a63d6caff14100f5eab5c50912bcd"><code>4ea0e64</code></a> Ready release gin 1.9.1 (by: thinkerou) (<a href="https://redirect.github.com/gin-gonic/gin/issues/3630">#3630</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/bb1fc2e0fe97c63dab1527baab88d01183853b8f"><code>bb1fc2e</code></a> fix Request.Context() checks (<a href="https://redirect.github.com/gin-gonic/gin/issues/3512">#3512</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/2d4bbec941551479b1fdf1e54ece03e6e82a7e72"><code>2d4bbec</code></a> fix lack of escaping of filename in Content-Disposition (<a href="https://redirect.github.com/gin-gonic/gin/issues/3556">#3556</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/9f5ecd4be440f2789db917aa93c1b8afa2276917"><code>9f5ecd4</code></a> chore(deps): bump actions/setup-go from 3 to 4 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3543">#3543</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/20cd6bcfc41148ae4acb01290496f818a61306aa"><code>20cd6bc</code></a> chore(deps): bump github.com/go-playground/validator/v10 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3610">#3610</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/6bdc725c8dfdc203530f1c64c7ea1aaaf4aeaa40"><code>6bdc725</code></a> Fix typos in ISSUE_TEMPLATE.md (<a href="https://redirect.github.com/gin-gonic/gin/issues/3616">#3616</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/1ab268989db62a6dd86264cb20e14160e25a6a6d"><code>1ab2689</code></a> chore(deps): bump golang.org/x/net from 0.9.0 to 0.10.0 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3599">#3599</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/6a0556ed5a67d1d12ae3e7ea2c0121b6c3b894e2"><code>6a0556e</code></a> improve render code coverage (<a href="https://redirect.github.com/gin-gonic/gin/issues/3525">#3525</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/eac2daac64811197970b5d2f6406e4ae6c31cb5e"><code>eac2daa</code></a> chore: update dependencies for various packages and libraries (<a href="https://redirect.github.com/gin-gonic/gin/issues/3585">#3585</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/757a638b7bbdd998375432fb22f693e82d7a7840"><code>757a638</code></a> chore: improve linting, testing, and GitHub Actions setup (<a href="https://redirect.github.com/gin-gonic/gin/issues/3583">#3583</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/gin-gonic/gin/compare/v1.9.0...v1.9.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/gin-gonic/gin&package-manager=go_modules&previous-version=1.9.0&new-version=1.9.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-engine/network/alerts).

</details>

---

## Discussion

### Comment by @jira-linking on June 01, 2023 at 08:36 PM UTC

Commits missing Jira IDs:
2f78521ca1b3a49b8fe13772487a3838a0940207


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1232*
