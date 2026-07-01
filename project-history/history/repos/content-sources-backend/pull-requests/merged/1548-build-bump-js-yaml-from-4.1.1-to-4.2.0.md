---
type: pull_request
number: 1548
title: "Build: Bump js-yaml from 4.1.1 to 4.2.0"
state: merged
author: dependabot
created: 2026-06-22T07:29:01Z
updated: 2026-06-22T08:20:53Z
closed: 2026-06-22T08:20:52Z
merged: 2026-06-22T08:20:51Z
base_branch: main
head_branch: dependabot/npm_and_yarn/js-yaml-4.2.0
labels: ["dependencies", "javascript"]
url: https://github.com/content-services/content-sources-backend/pull/1548
---

# Pull Request #1548: Build: Bump js-yaml from 4.1.1 to 4.2.0

**Author**: @dependabot
**Created**: June 22, 2026 at 07:29 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `javascript`
**Base**: `main` ← **Head**: `dependabot/npm_and_yarn/js-yaml-4.2.0`

## Description

Bumps [js-yaml](https://github.com/nodeca/js-yaml) from 4.1.1 to 4.2.0.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/nodeca/js-yaml/blob/master/CHANGELOG.md">js-yaml's changelog</a>.</em></p>
<blockquote>
<h2>[4.2.0] - 2026-06-01</h2>
<h3>Added</h3>
<ul>
<li>Added <code>docs/safety.md</code> with notes about processing untrusted YAML.</li>
<li>Added <code>maxDepth</code> (100) loader option. Not a problem, but gives a better
exception instead of RangeError on stack overflow.</li>
<li>Added <code>maxMergeSeqLength</code> (20) loader option. Not a problem after <code>merge</code> fix,
but an additional restriction for safety.</li>
<li>Added sourcemaps to <code>dist/</code> builds.</li>
</ul>
<h3>Changed</h3>
<ul>
<li>Stop resolving numbers with underscores as numeric scalars, <a href="https://redirect.github.com/nodeca/js-yaml/issues/627">#627</a>.</li>
<li>Switched dev toolchains to Vite / neostandard.</li>
<li>Updated demo.</li>
<li>Reorganized tests.</li>
<li><code>dist/</code> files are no longer kept in the repository.</li>
</ul>
<h3>Fixed</h3>
<ul>
<li>Fix parsing of properties on the first implicit block mapping key, <a href="https://redirect.github.com/nodeca/js-yaml/issues/62">#62</a>.</li>
<li>Fix trailing whitespace handling when folding flow scalar lines, <a href="https://redirect.github.com/nodeca/js-yaml/issues/307">#307</a>.</li>
<li>Reject top-level block scalars without content indentation, <a href="https://redirect.github.com/nodeca/js-yaml/issues/280">#280</a>.</li>
<li>Ensure numbers survive round-trip, <a href="https://redirect.github.com/nodeca/js-yaml/issues/737">#737</a>.</li>
<li>Fix test coverage for issue <a href="https://redirect.github.com/nodeca/js-yaml/issues/221">#221</a>.</li>
<li>Fix flow scalar trailing whitespace folding, <a href="https://redirect.github.com/nodeca/js-yaml/issues/307">#307</a>.</li>
<li>Fix digits in YAML named tag handles.</li>
</ul>
<h3>Security</h3>
<ul>
<li>Fix potential DoS via quadratic complexity in merge - deduplicate repeated
elements (makes sense for malformed files &gt; 10K).</li>
</ul>
<h2>[3.14.2] - 2025-11-15</h2>
<h3>Security</h3>
<ul>
<li>Backported v4.1.1 fix to v3</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/nodeca/js-yaml/commit/590dbabadd172b099c07654fab2eabec8c7a07b9"><code>590dbab</code></a> 4.2.0 released</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/f944dc5cd132251752499bdb157f33027d362177"><code>f944dc5</code></a> Add package.json funding field</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/f6927192223355b64f2a6e19f3071ccc579ee718"><code>f692719</code></a> Changelog update</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/9971a068fe7fb67eeedec7cf15fd2aba1b71bd79"><code>9971a06</code></a> Fix digits in YAML named tag handles</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/464a5b854b682691d35cc119500f74595d2e513c"><code>464a5b8</code></a> Fix flow scalar trailing whitespace folding, close <a href="https://redirect.github.com/nodeca/js-yaml/issues/307">#307</a></li>
<li><a href="https://github.com/nodeca/js-yaml/commit/1fda4f715d368c74e6e65e44d3201946c9577e54"><code>1fda4f7</code></a> Tests for <a href="https://redirect.github.com/nodeca/js-yaml/issues/567">#567</a>, <a href="https://redirect.github.com/nodeca/js-yaml/issues/565">#565</a></li>
<li><a href="https://github.com/nodeca/js-yaml/commit/031ad079a65da5018cbe6280f3b9306e39699b5a"><code>031ad07</code></a> Stop resolving numbers with underscores as numeric scalars, <a href="https://redirect.github.com/nodeca/js-yaml/issues/627">#627</a></li>
<li><a href="https://github.com/nodeca/js-yaml/commit/e46d223b37110ec2ba5b5b1da672a8f6862832d6"><code>e46d223</code></a> CI config update</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/9023feec80d19f4c39e73d1b25fc7644e93bf628"><code>9023fee</code></a> Add lockfile</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/990e6f4d1c51d03b53da0dc6df3c7fc0f61ad909"><code>990e6f4</code></a> Docs update</li>
<li>Additional commits viewable in <a href="https://github.com/nodeca/js-yaml/compare/4.1.1...4.2.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=js-yaml&package-manager=npm_and_yarn&previous-version=4.1.1&new-version=4.2.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Reviews

### Review by @swadeley - Approved on June 22, 2026 at 07:32 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1548*
