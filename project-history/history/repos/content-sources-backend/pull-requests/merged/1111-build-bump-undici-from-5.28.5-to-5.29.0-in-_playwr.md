---
type: pull_request
number: 1111
title: "Build: Bump undici from 5.28.5 to 5.29.0 in /_playwright-tests"
state: merged
author: dependabot
created: 2025-05-15T16:32:59Z
updated: 2025-07-04T13:21:17Z
closed: 2025-07-04T13:21:09Z
merged: 2025-07-04T13:21:09Z
base_branch: main
head_branch: dependabot/npm_and_yarn/_playwright-tests/undici-5.29.0
labels: ["dependencies", "javascript"]
url: https://github.com/content-services/content-sources-backend/pull/1111
---

# Pull Request #1111: Build: Bump undici from 5.28.5 to 5.29.0 in /_playwright-tests

**Author**: @dependabot
**Created**: May 15, 2025 at 04:32 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `javascript`
**Base**: `main` ← **Head**: `dependabot/npm_and_yarn/_playwright-tests/undici-5.29.0`

## Description

Bumps [undici](https://github.com/nodejs/undici) from 5.28.5 to 5.29.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/nodejs/undici/releases">undici's releases</a>.</em></p>
<blockquote>
<h2>v5.29.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Fix tests in v5.x for Node 20 by <a href="https://github.com/mcollina"><code>@​mcollina</code></a> in <a href="https://redirect.github.com/nodejs/undici/pull/4104">nodejs/undici#4104</a></li>
<li>Removed clients with unrecoverable errors from the Pool <a href="https://redirect.github.com/nodejs/undici/pull/4088">nodejs/undici#4088</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/nodejs/undici/compare/v5.28.5...v5.29.0">https://github.com/nodejs/undici/compare/v5.28.5...v5.29.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/nodejs/undici/commit/9528f6853a72a637823e863f1dd12ec47a3fc9e7"><code>9528f68</code></a> Bumped v5.29.0</li>
<li><a href="https://github.com/nodejs/undici/commit/f1d75a4e107866110c48003f00e5d0de52ef2cce"><code>f1d75a4</code></a> increase timeout for redirect test</li>
<li><a href="https://github.com/nodejs/undici/commit/2d31ed61f7ca12ef6d89a323dc236346364ac379"><code>2d31ed6</code></a> remove fuzzing tests</li>
<li><a href="https://github.com/nodejs/undici/commit/6b36d49cb2fa14217baa11b6fd27ee20b661ea4c"><code>6b36d49</code></a> fix redirect test in Node v16</li>
<li><a href="https://github.com/nodejs/undici/commit/648dd8f7ba3654db09a095361a167e3576db8cd0"><code>648dd8f</code></a> more fix for the wpt runner on Windows</li>
<li><a href="https://github.com/nodejs/undici/commit/a0516bae59b6aa8aa8124f9ae5cfed79541d10e2"><code>a0516ba</code></a> don't use internal header state for cookies (<a href="https://redirect.github.com/nodejs/undici/issues/3295">#3295</a>)</li>
<li><a href="https://github.com/nodejs/undici/commit/87ce4af0e58657506cedb2d07a5ba24f964b733f"><code>87ce4af</code></a> fix test/client for node 20</li>
<li><a href="https://github.com/nodejs/undici/commit/c2c8fd55b778267ad8b2e9ee04218c038a5d02af"><code>c2c8fd5</code></a> fix: accept v20 SSL specific error for alpn selection in http/2</li>
<li><a href="https://github.com/nodejs/undici/commit/82200bd10b7073ac235f8fc48d4daa82b350cd4c"><code>82200bd</code></a> [v6.x] fix wpts on windows (<a href="https://redirect.github.com/nodejs/undici/issues/4093">#4093</a>)</li>
<li><a href="https://github.com/nodejs/undici/commit/47546fa68d04eec5b96ab93225c3bc08b77cd94f"><code>47546fa</code></a> test: fix windows wpt (<a href="https://redirect.github.com/nodejs/undici/issues/4050">#4050</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/nodejs/undici/compare/v5.28.5...v5.29.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=undici&package-manager=npm_and_yarn&previous-version=5.28.5&new-version=5.29.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

You can trigger a rebase of this PR by commenting `@dependabot rebase`.

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

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.


---

## Discussion

### Comment by @rverdile on May 28, 2025 at 02:59 PM UTC

@dependabot rebase

### Comment by @swadeley on June 16, 2025 at 12:51 PM UTC

@dependabot rebase

### Comment by @rverdile on July 02, 2025 at 06:37 PM UTC

@dependabot rebase

### Comment by @rverdile on July 02, 2025 at 08:08 PM UTC

@dependabot rebase

### Comment by @swadeley on July 03, 2025 at 08:18 PM UTC

https://github.com/dependabot rebase

### Comment by @dependabot on July 03, 2025 at 08:18 PM UTC

Looks like this PR is already up-to-date with main! If you'd still like to recreate it from scratch, overwriting any edits, you can request `@dependabot recreate`.

### Comment by @swadeley on July 03, 2025 at 08:18 PM UTC

/retest

### Comment by @swadeley on July 03, 2025 at 09:10 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on July 04, 2025 at 01:21 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1111*
