---
type: pull_request
number: 775
title: "chore(deps): bump urijs from 1.19.9 to 1.19.11"
state: closed
author: dependabot
created: 2022-04-13T00:20:36Z
updated: 2022-04-21T09:39:53Z
closed: 2022-04-21T09:39:52Z
base_branch: master
head_branch: dependabot/npm_and_yarn/urijs-1.19.11
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-ui/pull/775
---

# Pull Request #775: chore(deps): bump urijs from 1.19.9 to 1.19.11

**Author**: @dependabot
**Created**: April 13, 2022 at 12:20 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/urijs-1.19.11`

## Description

Bumps [urijs](https://github.com/medialize/URI.js) from 1.19.9 to 1.19.11.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/medialize/URI.js/releases">urijs's releases</a>.</em></p>
<blockquote>
<h2>1.19.11 (April 3rd 2022)</h2>
<ul>
<li><strong>SECURITY</strong> fixing <a href="http://medialize.github.io/URI.js/docs.html#static-parse"><code>URI.parse()</code></a> handle excessive slashes in scheme-relative URLs - disclosed by <a href="https://github.com/zeyu2001">zeyu2001</a> via <a href="https://huntr.dev/">https://huntr.dev/</a></li>
<li><strong>SECURITY</strong> fixing <a href="http://medialize.github.io/URI.js/docs.html#static-parse"><code>URI.parse()</code></a> remove <code>\r</code> (CR), <code>\n</code>, (LF) <code>\t</code> (TAB) - disclosed by <a href="https://github.com/haxatron">haxatron</a> via <a href="https://huntr.dev/">https://huntr.dev/</a></li>
</ul>
<h2>1.19.10 (March 5th 2022)</h2>
<ul>
<li><strong>SECURITY</strong> fixing <a href="http://medialize.github.io/URI.js/docs.html#static-parse"><code>URI.parse()</code></a> handle excessive colons in protocol delimiter - disclosed by <a href="https://github.com/huydoppa">huydoppa</a> via <a href="https://huntr.dev/">https://huntr.dev/</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/medialize/URI.js/blob/gh-pages/CHANGELOG.md">urijs's changelog</a>.</em></p>
<blockquote>
<h3>1.19.11 (April 3rd 2022)</h3>
<ul>
<li><strong>SECURITY</strong> fixing <a href="http://medialize.github.io/URI.js/docs.html#static-parse"><code>URI.parse()</code></a> handle excessive slashes in scheme-relative URLs - disclosed by <a href="https://github.com/zeyu2001">zeyu2001</a> via <a href="https://huntr.dev/">https://huntr.dev/</a></li>
<li><strong>SECURITY</strong> fixing <a href="http://medialize.github.io/URI.js/docs.html#static-parse"><code>URI.parse()</code></a> remove <code>\r</code> (CR), <code>\n</code>, (LF) <code>\t</code> (TAB) - disclosed by <a href="https://github.com/haxatron">haxatron</a> via <a href="https://huntr.dev/">https://huntr.dev/</a></li>
</ul>
<h3>1.19.10 (March 5th 2022)</h3>
<ul>
<li><strong>SECURITY</strong> fixing <a href="http://medialize.github.io/URI.js/docs.html#static-parse"><code>URI.parse()</code></a> handle excessive colons in protocol delimiter - disclosed by <a href="https://github.com/huydoppa">huydoppa</a> via <a href="https://huntr.dev/">https://huntr.dev/</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/medialize/URI.js/commit/b655c1b972111ade9f181b02374305942e68e30a"><code>b655c1b</code></a> chore(build): bumping to version 1.19.11</li>
<li><a href="https://github.com/medialize/URI.js/commit/b0c9796aa1a95a85f40924fb18b1e5da3dc8ffae"><code>b0c9796</code></a> fix(parse): handle CR,LF,TAB</li>
<li><a href="https://github.com/medialize/URI.js/commit/88805fd3da03bd7a5e60947adb49d182011f1277"><code>88805fd</code></a> fix(parse): handle excessive slashes in scheme-relative URLs</li>
<li><a href="https://github.com/medialize/URI.js/commit/926b2aa1099f177f82d0a998da4b43e69fe56ec8"><code>926b2aa</code></a> chore(build): bumping to version 1.19.10</li>
<li><a href="https://github.com/medialize/URI.js/commit/a8166fe02f3af6dc1b2b888dcbb807155aad9509"><code>a8166fe</code></a> fix(parse): handle excessive colons in scheme delimiter</li>
<li>See full diff in <a href="https://github.com/medialize/URI.js/compare/v1.19.9...v1.19.11">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=urijs&package-manager=npm_and_yarn&previous-version=1.19.9&new-version=1.19.11)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @dependabot on April 21, 2022 at 09:39 AM UTC

Looks like urijs is up-to-date now, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/775*
