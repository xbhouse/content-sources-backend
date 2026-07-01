---
type: pull_request
number: 398
title: "Bump github.com/redis/go-redis/v9 from 9.1.0 to 9.2.0"
state: closed
author: dependabot
created: 2023-09-21T16:17:35Z
updated: 2023-09-26T14:23:48Z
closed: 2023-09-26T14:23:46Z
base_branch: main
head_branch: dependabot/go_modules/github.com/redis/go-redis/v9-9.2.0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/398
---

# Pull Request #398: Bump github.com/redis/go-redis/v9 from 9.1.0 to 9.2.0

**Author**: @dependabot
**Created**: September 21, 2023 at 04:17 PM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/redis/go-redis/v9-9.2.0`

## Description

Bumps [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) from 9.1.0 to 9.2.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.2.0</h2>
<h2>What's new?</h2>
<h3>Time series support</h3>
<p>We are happy to announce support for <a href="https://redis.io/docs/data-types/timeseries/">Time Series</a>. This enhancement allows developers to leverage the time series data structure directly within Go-Redis, enabling efficient ingestion, querying, and storage of time-sequential data. Whether you're tracking metrics, logs, or other time-sensitive information, this feature offers a robust solution to manage time series data seamlessly. We encourage developers to explore this new capability and provide feedback for further improvements.
<a href="https://redis.io/docs/data-types/timeseries/">Learn more about Redis Time Series</a></p>
<h3>Better support for Redis 7.2</h3>
<p>Go-Redis now supports <a href="https://redis.io/commands/waitaof/">WAITAOF</a> and <code>CLIENT SETINFO</code> commands introduced in <a href="https://redis.com/blog/introducing-redis-7-2/">Redis 7.2</a>.</p>
<h3>Other notable improvements</h3>
<ul>
<li>Add the ability to set a connection growth limit with MaxActiveConns configuration setting (<a href="https://redirect.github.com/redis/go-redis/issues/2646">#2646</a>)</li>
<li>Add support for multiple values in the bitfield cmd (<a href="https://redirect.github.com/redis/go-redis/issues/2648">#2648</a>)</li>
</ul>
<h2>Breaking Changes</h2>
<ul>
<li>Changing the suffix for probablistic commands acceptings arguments to WithArgs from Args (<a href="https://redirect.github.com/redis/go-redis/issues/2701">#2701</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>Making command structs digestible (<a href="https://redirect.github.com/redis/go-redis/issues/2716">#2716</a>)</li>
<li>change interfaces probabilistic and redis gears from private to public (<a href="https://redirect.github.com/redis/go-redis/issues/2695">#2695</a>)</li>
<li>Adding CONTRIBUTING guidelines (<a href="https://redirect.github.com/redis/go-redis/issues/2718">#2718</a>)</li>
<li>Adding Go 1.21.x for CI coverage (<a href="https://redirect.github.com/redis/go-redis/issues/2697">#2697</a>)</li>
<li>chore(deps): bump actions/checkout from 3 to 4 (<a href="https://redirect.github.com/redis/go-redis/issues/2702">#2702</a>)</li>
<li>chore(deps): bump github.com/bsm/ginkgo/v2 from 2.9.5 to 2.12.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2690">#2690</a>)</li>
<li>chore(deps): bump github.com/bsm/gomega from 1.26.0 to 1.27.10 (<a href="https://redirect.github.com/redis/go-redis/issues/2689">#2689</a>)</li>
<li>Adding stale issues workflow (<a href="https://redirect.github.com/redis/go-redis/issues/2700">#2700</a>)</li>
<li>Updating redis binary for makefile to 7.2.1 (<a href="https://redirect.github.com/redis/go-redis/issues/2693">#2693</a>)</li>
<li>Skip flaky tests (<a href="https://redirect.github.com/redis/go-redis/issues/2699">#2699</a>)</li>
<li>Format code and fix go vet (<a href="https://redirect.github.com/redis/go-redis/issues/2696">#2696</a>)</li>
<li>Use time duration calculation (<a href="https://redirect.github.com/redis/go-redis/issues/2651">#2651</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/NikanV"><code>@​NikanV</code></a>, <a href="https://github.com/SoulPancake"><code>@​SoulPancake</code></a>, <a href="https://github.com/chayim"><code>@​chayim</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/nvorobev"><code>@​nvorobev</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/peczenyj"><code>@​peczenyj</code></a>, <a href="https://github.com/taytzehao"><code>@​taytzehao</code></a> and <a href="https://github.com/wzlove"><code>@​wzlove</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/dac3314bc640e4a4c3c8bca03bae9a4ec67b0de8"><code>dac3314</code></a> Bump version to 9.2.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2722">#2722</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6199a2af2c59d85ab536408a4c8b92f8cba07f2f"><code>6199a2a</code></a> Making command structs digestable (<a href="https://redirect.github.com/redis/go-redis/issues/2716">#2716</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e07f7e62b8197a8f47edeeef230dfbf53ef05aa5"><code>e07f7e6</code></a> chore(deps): bump github.com/bsm/ginkgo/v2 from 2.9.5 to 2.12.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2690">#2690</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7ecd7ac1c7b84f19fe723539d27089686ade1925"><code>7ecd7ac</code></a> chore(deps): bump github.com/bsm/gomega from 1.26.0 to 1.27.10 (<a href="https://redirect.github.com/redis/go-redis/issues/2689">#2689</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e23ea028bd95ae0f8a76a59ea53b7557e0c19fe5"><code>e23ea02</code></a> Added MaxActiveConns (<a href="https://redirect.github.com/redis/go-redis/issues/2646">#2646</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/934c6a3fe0073b3afcca6c293a35486f2971c034"><code>934c6a3</code></a> make public probabilistic and redis gears interfaces (<a href="https://redirect.github.com/redis/go-redis/issues/2695">#2695</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/0637c53f103b127f77bce430a7ac5b566a61aa30"><code>0637c53</code></a> Added the support for WAITAOF which is a new command in redis ver7.2.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2629">#2629</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7acc0cd2546899c179b5952710cba486d3200ee3"><code>7acc0cd</code></a> useTime duration calculation (<a href="https://redirect.github.com/redis/go-redis/issues/2651">#2651</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e8ad794e965a1b4e2ba63fe9a6975a14ebf289b9"><code>e8ad794</code></a> Format code and fix go vet (<a href="https://redirect.github.com/redis/go-redis/issues/2696">#2696</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/1a7d2f4ad4e91b9fac784aa1b94254f838863e3b"><code>1a7d2f4</code></a> upgrade bitfield cmd to <code>add multiple values</code> (<a href="https://redirect.github.com/redis/go-redis/issues/2648">#2648</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.1.0...v9.2.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/redis/go-redis/v9&package-manager=go_modules&previous-version=9.1.0&new-version=9.2.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @app-sre-bot on September 21, 2023 at 04:18 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on September 21, 2023 at 06:33 PM UTC

[test]

### Comment by @dependabot on September 26, 2023 at 02:23 PM UTC

Looks like github.com/redis/go-redis/v9 is up-to-date now, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/398*
