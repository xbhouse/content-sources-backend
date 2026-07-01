---
type: pull_request
number: 1469
title: "chore(deps): bump @remix-run/router and react-router-dom"
state: closed
author: dependabot
created: 2026-01-15T08:31:58Z
updated: 2026-02-10T10:33:53Z
closed: 2026-02-10T10:33:51Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-5c7f151553
labels: ["dependencies", "patch"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1469
---

# Pull Request #1469: chore(deps): bump @remix-run/router and react-router-dom

**Author**: @dependabot
**Created**: January 15, 2026 at 08:31 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `patch`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-5c7f151553`

## Description

Bumps [@remix-run/router](https://github.com/remix-run/react-router/tree/HEAD/packages/router) to 1.23.2 and updates ancestor dependency [react-router-dom](https://github.com/remix-run/react-router/tree/HEAD/packages/react-router-dom). These dependencies need to be updated together.

Updates `@remix-run/router` from 1.23.0 to 1.23.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/remix-run/react-router/blob/@remix-run/router@1.23.2/packages/router/CHANGELOG.md"><code>@​remix-run/router</code>'s changelog</a>.</em></p>
<blockquote>
<h2>1.23.2</h2>
<h3>Patch Changes</h3>
<ul>
<li>Validate redirect locations (<a href="https://redirect.github.com/remix-run/react-router/pull/14707">#14707</a>)</li>
</ul>
<h2>1.23.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>Normalize double-slashes in <code>resolvePath</code> (<a href="https://redirect.github.com/remix-run/react-router/pull/14537">#14537</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/remix-run/react-router/commit/c662ca366a414bf42624dd6cd20a7c414b2602e3"><code>c662ca3</code></a> chore: Update version for release (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/router/issues/14713">#14713</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/98ad6912daec8df0d911f786f18006048efd7ade"><code>98ad691</code></a> chore: Update version for release (pre-v6) (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/router/issues/14710">#14710</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/2fbb84c83dae3695a0435beb0c3c0b467b7c2da2"><code>2fbb84c</code></a> Validate redirect locations (v6) (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/router/issues/14707">#14707</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/26b5d4581fb2829dc7eaeaad413de4735173a6eb"><code>26b5d45</code></a> chore: Update version for release (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/router/issues/14541">#14541</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/919f8a86b95f0c8956e3820743503d5609f572cd"><code>919f8a8</code></a> chore: Update version for release (pre-v6) (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/router/issues/14540">#14540</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/69bf7056f1540918a0324e0af215978cdaf5d486"><code>69bf705</code></a> Normalize double slashes in resolvePath (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/router/issues/14537">#14537</a>)</li>
<li>See full diff in <a href="https://github.com/remix-run/react-router/commits/@remix-run/router@1.23.2/packages/router">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for <code>@​remix-run/router</code> since your current version.</p>
</details>
<br />

Updates `react-router-dom` from 6.30.1 to 6.30.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/remix-run/react-router/releases">react-router-dom's releases</a>.</em></p>
<blockquote>
<h2>react-router-dom-v5-compat@6.4.0-pre.15</h2>
<h3>Patch Changes</h3>
<ul>
<li>Updated dependencies
<ul>
<li>react-router@6.4.0-pre.15</li>
<li>react-router-dom@6.4.0-pre.15</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/remix-run/react-router/blob/main/CHANGELOG.md">react-router-dom's changelog</a>.</em></p>
<blockquote>
<h2>v6.30.3</h2>
<p>Date: 2026-01-07</p>
<h3>Security Notice</h3>
<p>This release addresses 1 security vulnerability:</p>
<ul>
<li><a href="https://github.com/remix-run/react-router/security/advisories/GHSA-2w69-qvjg-hvjx">XSS via Open Redirects</a></li>
</ul>
<h3>Patch Changes</h3>
<ul>
<li>Validate redirect locations (<a href="https://redirect.github.com/remix-run/react-router/pull/14707">#14707</a>)</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/remix-run/react-router/compare/react-router@6.30.2...react-router@6.30.3"><code>v6.30.2...v6.30.3</code></a></p>
<h2>v6.30.2</h2>
<p>Date: 2025-11-13</p>
<h3>Security Notice</h3>
<p>This release addresses 1 security vulnerability:</p>
<ul>
<li><a href="https://github.com/remix-run/react-router/security/advisories/GHSA-9jcx-v3wj-wh4m">Unexpected external redirect via untrusted paths</a></li>
</ul>
<h3>Patch Changes</h3>
<ul>
<li>Normalize double-slashes in <code>resolvePath</code> (<a href="https://redirect.github.com/remix-run/react-router/pull/14537">#14537</a>)</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/remix-run/react-router/compare/react-router@6.30.1...react-router@6.30.2"><code>v6.30.1...v6.30.2</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/remix-run/react-router/commit/c662ca366a414bf42624dd6cd20a7c414b2602e3"><code>c662ca3</code></a> chore: Update version for release (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/react-router-dom/issues/14713">#14713</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/98ad6912daec8df0d911f786f18006048efd7ade"><code>98ad691</code></a> chore: Update version for release (pre-v6) (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/react-router-dom/issues/14710">#14710</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/26b5d4581fb2829dc7eaeaad413de4735173a6eb"><code>26b5d45</code></a> chore: Update version for release (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/react-router-dom/issues/14541">#14541</a>)</li>
<li><a href="https://github.com/remix-run/react-router/commit/919f8a86b95f0c8956e3820743503d5609f572cd"><code>919f8a8</code></a> chore: Update version for release (pre-v6) (<a href="https://github.com/remix-run/react-router/tree/HEAD/packages/react-router-dom/issues/14540">#14540</a>)</li>
<li>See full diff in <a href="https://github.com/remix-run/react-router/commits/react-router-dom@6.30.3/packages/react-router-dom">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for react-router-dom since your current version.</p>
</details>
<br />


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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @dependabot on February 10, 2026 at 10:33 AM UTC

Looks like these dependencies are up-to-date now, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1469*
