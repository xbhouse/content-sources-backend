---
type: pull_request
number: 748
title: "SPM-1058: Bump github.com/gin-gonic/gin from 1.6.3 to 1.7.0"
state: merged
author: dependabot
created: 2021-07-29T11:15:32Z
updated: 2021-07-29T11:26:06Z
closed: 2021-07-29T11:25:29Z
merged: 2021-07-29T11:25:29Z
base_branch: master
head_branch: dependabot/go_modules/github.com/gin-gonic/gin-1.7.0
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-engine/pull/748
---

# Pull Request #748: SPM-1058: Bump github.com/gin-gonic/gin from 1.6.3 to 1.7.0

**Author**: @dependabot
**Created**: July 29, 2021 at 11:15 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/go_modules/github.com/gin-gonic/gin-1.7.0`

## Description

Bumps [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) from 1.6.3 to 1.7.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/gin-gonic/gin/releases">github.com/gin-gonic/gin's releases</a>.</em></p>
<blockquote>
<h2>Release v1.7.0</h2>
<h3>BUGFIXES</h3>
<ul>
<li>fix compile error from <a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2572">#2572</a> (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2600">#2600</a>)</li>
<li>fix: print headers without Authorization header on broken pipe (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2528">#2528</a>)</li>
<li>fix(tree): reassign fullpath when register new node (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2366">#2366</a>)</li>
</ul>
<h3>ENHANCEMENTS</h3>
<ul>
<li>Support params and exact routes without creating conflicts (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2663">#2663</a>)</li>
<li>chore: improve render string performance (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2365">#2365</a>)</li>
<li>Sync route tree to httprouter latest code (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2368">#2368</a>)</li>
<li>chore: rename getQueryCache/getFormCache to initQueryCache/initFormCa (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2375">#2375</a>)</li>
<li>chore(performance): improve countParams (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2378">#2378</a>)</li>
<li>Remove some functions that have the same effect as the bytes package (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2387">#2387</a>)</li>
<li>update:SetMode function (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2321">#2321</a>)</li>
<li>remove a unused type SecureJSONPrefix (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2391">#2391</a>)</li>
<li>Add a redirect sample for POST method (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2389">#2389</a>)</li>
<li>Add CustomRecovery builtin middleware (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2322">#2322</a>)</li>
<li>binding: avoid 2038 problem on 32-bit architectures (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2450">#2450</a>)</li>
<li>Prevent panic in Context.GetQuery() when there is no Request (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2412">#2412</a>)</li>
<li>Add GetUint and GetUint64 method on gin.context (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2487">#2487</a>)</li>
<li>update content-disposition header to MIME-style (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2512">#2512</a>)</li>
<li>reduce allocs and improve the render <code>WriteString</code> (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2508">#2508</a>)</li>
<li>implement &quot;.Unwrap() error&quot; on Error type (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2525">#2525</a>) (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2526">#2526</a>)</li>
<li>Allow bind with a map[string]string (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2484">#2484</a>)</li>
<li>chore: update tree (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2371">#2371</a>)</li>
<li>Support binding for slice/array obj [Rewrite] (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2302">#2302</a>)</li>
<li>basic auth: fix timing oracle (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2609">#2609</a>)</li>
<li>Add mixed param and non-param paths (port of httprouter<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/329">#329</a>) (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2663">#2663</a>)</li>
<li>feat(engine): add trustedproxies and remoteIP (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2632">#2632</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/gin-gonic/gin/blob/master/CHANGELOG.md">github.com/gin-gonic/gin's changelog</a>.</em></p>
<blockquote>
<h2>Gin v1.7.0</h2>
<h3>BUGFIXES</h3>
<ul>
<li>fix compile error from <a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2572">#2572</a> (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2600">#2600</a>)</li>
<li>fix: print headers without Authorization header on broken pipe (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2528">#2528</a>)</li>
<li>fix(tree): reassign fullpath when register new node (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2366">#2366</a>)</li>
</ul>
<h3>ENHANCEMENTS</h3>
<ul>
<li>Support params and exact routes without creating conflicts (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2663">#2663</a>)</li>
<li>chore: improve render string performance (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2365">#2365</a>)</li>
<li>Sync route tree to httprouter latest code (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2368">#2368</a>)</li>
<li>chore: rename getQueryCache/getFormCache to initQueryCache/initFormCa (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2375">#2375</a>)</li>
<li>chore(performance): improve countParams (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2378">#2378</a>)</li>
<li>Remove some functions that have the same effect as the bytes package (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2387">#2387</a>)</li>
<li>update:SetMode function (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2321">#2321</a>)</li>
<li>remove a unused type SecureJSONPrefix (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2391">#2391</a>)</li>
<li>Add a redirect sample for POST method (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2389">#2389</a>)</li>
<li>Add CustomRecovery builtin middleware (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2322">#2322</a>)</li>
<li>binding: avoid 2038 problem on 32-bit architectures (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2450">#2450</a>)</li>
<li>Prevent panic in Context.GetQuery() when there is no Request (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2412">#2412</a>)</li>
<li>Add GetUint and GetUint64 method on gin.context (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2487">#2487</a>)</li>
<li>update content-disposition header to MIME-style (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2512">#2512</a>)</li>
<li>reduce allocs and improve the render <code>WriteString</code> (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2508">#2508</a>)</li>
<li>implement &quot;.Unwrap() error&quot; on Error type (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2525">#2525</a>) (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2526">#2526</a>)</li>
<li>Allow bind with a map[string]string (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2484">#2484</a>)</li>
<li>chore: update tree (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2371">#2371</a>)</li>
<li>Support binding for slice/array obj [Rewrite] (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2302">#2302</a>)</li>
<li>basic auth: fix timing oracle (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2609">#2609</a>)</li>
<li>Add mixed param and non-param paths (port of httprouter<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/329">#329</a>) (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2663">#2663</a>)</li>
<li>feat(engine): add trustedproxies and remoteIP (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/pull/2632">#2632</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/gin-gonic/gin/commit/d496f64540b6707602de50ab57aeea8ff4080b74"><code>d496f64</code></a> bump to v1.7.0 version (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2672">#2672</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/bfc8ca285eb46dad60e037d57c545cd260636711"><code>bfc8ca2</code></a> feat(engine): add trustedproxies and remoteIP (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2632">#2632</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/f3de8132c5d955784deeadb9bcf5752e9fdf0d8c"><code>f3de813</code></a> Add mixed param and non-param paths (port of httprouter#329) (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2663">#2663</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/a331dc6a31473b7208c57ec32e14bfcec3062dbb"><code>a331dc6</code></a> chore: remove duplicate test 'assert.Equal' (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2617">#2617</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/ed6f85c478ba00e5168be1f29ffcdc9a983568b8"><code>ed6f85c</code></a> build: convert to go:build directives (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2664">#2664</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/1bdf86b722026fd650fddfef7fe9bd8342b51b7a"><code>1bdf86b</code></a> Remove the tedious named return value (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2620">#2620</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/e899771392ecf35de8ce10a030ed8fed2207e9cb"><code>e899771</code></a> chore: Deleted spaces (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2622">#2622</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/b01605bb5b43dbf33781970af5ad6633e5549fd1"><code>b01605b</code></a> basic auth: fix timing oracle (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2609">#2609</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/46ddd4259cac975be1eb11b4f1192264f582db16"><code>46ddd42</code></a> Fixes to the graceful shutdown example (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2552">#2552</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/f4bc259de33c561fd3b0ae3e7aaa849c1d251c0b"><code>f4bc259</code></a> fix error gin support min Go version (<a href="https://github-redirect.dependabot.com/gin-gonic/gin/issues/2584">#2584</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/gin-gonic/gin/compare/v1.6.3...v1.7.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/gin-gonic/gin&package-manager=go_modules&previous-version=1.6.3&new-version=1.7.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot use these labels` will set the current labels as the default for future PRs for this repo and language
- `@dependabot use these reviewers` will set the current reviewers as the default for future PRs for this repo and language
- `@dependabot use these assignees` will set the current assignees as the default for future PRs for this repo and language
- `@dependabot use this milestone` will set the current milestone as the default for future PRs for this repo and language

You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-engine/network/alerts).

</details>

---

## Discussion

### Comment by @app-sre-bot on July 29, 2021 at 11:15 AM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on July 29, 2021 at 11:21 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#748](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (a72e7fa) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/11202b0411496bbea0d366aa1ab0942828e33ef4?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (11202b0) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #748   +/-   ##
=======================================
  Coverage   57.03%   57.03%           
=======================================
  Files          81       81           
  Lines        3833     3833           
=======================================
  Hits         2186     2186           
  Misses       1329     1329           
  Partials      318      318           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.03% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [11202b0...a72e7fa](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/748?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/748*
