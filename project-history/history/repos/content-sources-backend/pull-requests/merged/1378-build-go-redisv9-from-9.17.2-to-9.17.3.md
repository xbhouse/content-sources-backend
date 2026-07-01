---
type: pull_request
number: 1378
title: "Build: go-redis/v9 from 9.17.2 to 9.17.3"
state: merged
author: dependabot
created: 2026-01-26T04:40:27Z
updated: 2026-01-26T13:14:18Z
closed: 2026-01-26T13:14:15Z
merged: 2026-01-26T13:14:15Z
base_branch: main
head_branch: dependabot/go_modules/go-88646566e6
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1378
---

# Pull Request #1378: Build: go-redis/v9 from 9.17.2 to 9.17.3

**Author**: @dependabot
**Created**: January 26, 2026 at 04:40 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-88646566e6`

## Description

Bumps the go group with 1 update: [github.com/redis/go-redis/v9](https://github.com/redis/go-redis).

Updates `github.com/redis/go-redis/v9` from 9.17.2 to 9.17.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.17.3</h2>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Connection Pool</strong>: Fixed zombie <code>wantConn</code> elements accumulation in <code>wantConnQueue</code> that could cause resource leaks in high concurrency scenarios with dial failures (<a href="https://redirect.github.com/redis/go-redis/pull/3680">#3680</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>Stream Commands</strong>: Fixed <code>XADD</code> and <code>XTRIM</code> commands to use exact threshold (<code>=</code>) when <code>Approx</code> is false, ensuring precise stream trimming behavior (<a href="https://redirect.github.com/redis/go-redis/pull/3684">#3684</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Connection Pool</strong>: Added <code>ConnMaxLifetimeJitter</code> configuration to distribute connection expiration times and prevent the thundering herd problem when many connections expire simultaneously (<a href="https://redirect.github.com/redis/go-redis/pull/3666">#3666</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>Client Options</strong>: Added <code>DialerRetries</code> and <code>DialerRetryTimeout</code> fields to <code>ClusterOptions</code>, <code>RingOptions</code>, and <code>FailoverOptions</code> to allow configuring connection retry behavior for cluster, ring, and sentinel clients (<a href="https://redirect.github.com/redis/go-redis/pull/3686">#3686</a>) by <a href="https://github.com/naveenchander30"><code>@​naveenchander30</code></a></li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/cyningsun"><code>@​cyningsun</code></a>, <a href="https://github.com/naveenchander30"><code>@​naveenchander30</code></a>, and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<hr />
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.17.2...v9.17.3">https://github.com/redis/go-redis/compare/v9.17.2...v9.17.3</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/v9.17.3/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.17.3 (2026-01-25)</h1>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Connection Pool</strong>: Fixed zombie <code>wantConn</code> elements accumulation in <code>wantConnQueue</code> that could cause resource leaks in high concurrency scenarios with dial failures (<a href="https://redirect.github.com/redis/go-redis/pull/3680">#3680</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>Stream Commands</strong>: Fixed <code>XADD</code> and <code>XTRIM</code> commands to use exact threshold (<code>=</code>) when <code>Approx</code> is false, ensuring precise stream trimming behavior (<a href="https://redirect.github.com/redis/go-redis/pull/3684">#3684</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Connection Pool</strong>: Added <code>ConnMaxLifetimeJitter</code> configuration to distribute connection expiration times and prevent the thundering herd problem when many connections expire simultaneously (<a href="https://redirect.github.com/redis/go-redis/pull/3666">#3666</a>) by <a href="https://github.com/cyningsun"><code>@​cyningsun</code></a></li>
<li><strong>Client Options</strong>: Added <code>DialerRetries</code> and <code>DialerRetryTimeout</code> fields to <code>ClusterOptions</code>, <code>RingOptions</code>, and <code>FailoverOptions</code> to allow configuring connection retry behavior for cluster, ring, and sentinel clients (<a href="https://redirect.github.com/redis/go-redis/pull/3686">#3686</a>) by <a href="https://github.com/naveenchander30"><code>@​naveenchander30</code></a></li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/cyningsun"><code>@​cyningsun</code></a>, <a href="https://github.com/naveenchander30"><code>@​naveenchander30</code></a>, and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<hr />
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.17.2...v9.17.3">https://github.com/redis/go-redis/compare/v9.17.2...v9.17.3</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/0a836fb24c808795dfa561ddfdba613e6b4961ea"><code>0a836fb</code></a> chore(release): 9.17.3 patch with bugfixes (<a href="https://redirect.github.com/redis/go-redis/issues/3688">#3688</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2668e112ee20c81389703df94df04b057cca8b86"><code>2668e11</code></a> feat(options): add DialerRetries and DialerRetryTimeout to ClusterOptions, Ri...</li>
<li><a href="https://github.com/redis/go-redis/commit/981333fb426f5d4fe985d1071b481a40ceb88a45"><code>981333f</code></a> fix(pool): fix wantConnQueue zombie elements and add comprehensive test cover...</li>
<li><a href="https://github.com/redis/go-redis/commit/87be269e743f511d0babfa81f285c96bc8b40868"><code>87be269</code></a> fix(xadd,xtrim): when approx is false, should have <code>=</code> (<a href="https://redirect.github.com/redis/go-redis/issues/3684">#3684</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/5a616396ba090a173d1f920b57d807b24f6c7880"><code>5a61639</code></a> feat(pool): add ConnMaxLifetimeJitter to prevent connection thundering herd (...</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.17.2...v9.17.3">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/redis/go-redis/v9&package-manager=go_modules&previous-version=9.17.2&new-version=9.17.3)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @TenSt on January 26, 2026 at 11:54 AM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on January 26, 2026 at 11:53 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1378*
