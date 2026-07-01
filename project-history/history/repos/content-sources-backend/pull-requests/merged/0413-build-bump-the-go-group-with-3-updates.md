---
type: pull_request
number: 413
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2023-10-02T13:58:59Z
updated: 2023-10-02T21:00:30Z
closed: 2023-10-02T21:00:26Z
merged: 2023-10-02T21:00:26Z
base_branch: main
head_branch: dependabot/go_modules/go-d26fdb8f3b
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/413
---

# Pull Request #413: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: October 02, 2023 at 01:58 PM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-d26fdb8f3b`

## Description

Bumps the go group with 3 updates: [github.com/go-playground/validator/v10](https://github.com/go-playground/validator), [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) and [github.com/redis/go-redis/v9](https://github.com/redis/go-redis).

Updates `github.com/go-playground/validator/v10` from 10.15.4 to 10.15.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-playground/validator/releases">github.com/go-playground/validator/v10's releases</a>.</em></p>
<blockquote>
<h2>Release 10.15.5</h2>
<h2>What was fixed?</h2>
<p>Fixed <code>CIDRIPv4</code> validation, ty <a href="https://github.com/martinlehoux"><code>@​martinlehoux</code></a> for the <a href="https://redirect.github.com/go-playground/validator/pull/945">PR</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-playground/validator/commit/94a637ab9fbbb0bc0fe8a278f0352d0b14e2c365"><code>94a637a</code></a> feat(BREAKING): Change CIDRIPv4 validation (<a href="https://redirect.github.com/go-playground/validator/issues/945">#945</a>)</li>
<li>See full diff in <a href="https://github.com/go-playground/validator/compare/v10.15.4...v10.15.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.16.0 to 1.17.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.17.0</h2>
<h2>What's Changed</h2>
<ul>
<li>[CHANGE] Minimum required go version is now 1.19 (we also test client_golang against new 1.21 version). <a href="https://redirect.github.com/prometheus/client_golang/issues/1325">#1325</a></li>
<li>[FEATURE] Add support for Created Timestamps in Counters, Summaries and Historams. <a href="https://redirect.github.com/prometheus/client_golang/issues/1313">#1313</a></li>
<li>[ENHANCEMENT] Enable detection of a native histogram without observations. <a href="https://redirect.github.com/prometheus/client_golang/issues/1314">#1314</a></li>
</ul>
<!-- raw HTML omitted -->
<ul>
<li>Merge v1.16.0 to main by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1293">prometheus/client_golang#1293</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1297">prometheus/client_golang#1297</a></li>
<li>ci: define minimal permissions to GitHub workflows by <a href="https://github.com/diogoteles08"><code>@​diogoteles08</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1295">prometheus/client_golang#1295</a></li>
<li>Do not allocate memory when there's no constraints by <a href="https://github.com/Okhoshi"><code>@​Okhoshi</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1296">prometheus/client_golang#1296</a></li>
<li>Bump golang.org/x/sys from 0.8.0 to 0.9.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1306">prometheus/client_golang#1306</a></li>
<li>Bump google.golang.org/grpc from 1.45.0 to 1.53.0 in /tutorial/whatsup by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1307">prometheus/client_golang#1307</a></li>
<li>histogram: Enable detection of a native histogram without observations by <a href="https://github.com/beorn7"><code>@​beorn7</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1314">prometheus/client_golang#1314</a></li>
<li>Bump github.com/prometheus/procfs from 0.10.1 to 0.11.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1305">prometheus/client_golang#1305</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1302">prometheus/client_golang#1302</a></li>
<li>Fix data-race in metric without <code>code</code> and <code>method</code> but with <code>WithLabelFromCtx</code> by <a href="https://github.com/tigrato"><code>@​tigrato</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1318">prometheus/client_golang#1318</a></li>
<li>Add missing tick &quot;`&quot; in README by <a href="https://github.com/ZiViZiViZ"><code>@​ZiViZiViZ</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1321">prometheus/client_golang#1321</a></li>
<li>Bump golang.org/x/sys from 0.9.0 to 0.10.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1320">prometheus/client_golang#1320</a></li>
<li>Bump github.com/prometheus/procfs from 0.11.0 to 0.11.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1319">prometheus/client_golang#1319</a></li>
<li>docs: trivial grammar fixes to improve readability in promauto Godoc by <a href="https://github.com/sengi"><code>@​sengi</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1322">prometheus/client_golang#1322</a></li>
<li>Add Go 1.21 support by <a href="https://github.com/kakkoyun"><code>@​kakkoyun</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1325">prometheus/client_golang#1325</a></li>
<li>Bump client_model by <a href="https://github.com/ArthurSens"><code>@​ArthurSens</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1323">prometheus/client_golang#1323</a></li>
<li>histogram docs: Fixed minor nit. by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1324">prometheus/client_golang#1324</a></li>
<li>Update building by <a href="https://github.com/SuperQ"><code>@​SuperQ</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1326">prometheus/client_golang#1326</a></li>
<li>Bump golang.org/x/sys from 0.10.0 to 0.11.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1331">prometheus/client_golang#1331</a></li>
<li>Bump github.com/prometheus/client_golang from 1.15.1-0.20230416215738-0963f595c689 to 1.16.0 in /tutorial/whatsup by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1329">prometheus/client_golang#1329</a></li>
<li>Bump github.com/prometheus/client_golang from 1.13.1 to 1.16.0 in /examples/middleware by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1328">prometheus/client_golang#1328</a></li>
<li>Bump github.com/prometheus/common from 0.42.0 to 0.44.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1284">prometheus/client_golang#1284</a></li>
<li>Bump github.com/prometheus/common from 0.42.0 to 0.44.0 in /tutorial/whatsup by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1330">prometheus/client_golang#1330</a></li>
<li>Bump google.golang.org/protobuf from 1.30.0 to 1.31.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1304">prometheus/client_golang#1304</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1332">prometheus/client_golang#1332</a></li>
<li>Synchronize common files from prometheus/prometheus by <a href="https://github.com/prombot"><code>@​prombot</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1338">prometheus/client_golang#1338</a></li>
<li>Cleanup golangci-lint errcheck by <a href="https://github.com/SuperQ"><code>@​SuperQ</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1339">prometheus/client_golang#1339</a></li>
<li>Add go_godebug_non_default_behavior_tlsmaxrsasize_events_total by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1348">prometheus/client_golang#1348</a></li>
<li>Extend Counters, Summaries and Histograms with creation timestamp by <a href="https://github.com/ArthurSens"><code>@​ArthurSens</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1313">prometheus/client_golang#1313</a></li>
<li>Fix typos in comments, tests, and errors by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1346">prometheus/client_golang#1346</a></li>
<li>Deprecated comment should begin with &quot;Deprecated:&quot; by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1347">prometheus/client_golang#1347</a></li>
<li>Add changelog entry for 1.17 by <a href="https://github.com/ArthurSens"><code>@​ArthurSens</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1352">prometheus/client_golang#1352</a></li>
</ul>
<!-- raw HTML omitted -->
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/diogoteles08"><code>@​diogoteles08</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1295">prometheus/client_golang#1295</a></li>
<li><a href="https://github.com/tigrato"><code>@​tigrato</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1318">prometheus/client_golang#1318</a></li>
<li><a href="https://github.com/ZiViZiViZ"><code>@​ZiViZiViZ</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1321">prometheus/client_golang#1321</a></li>
<li><a href="https://github.com/sengi"><code>@​sengi</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1322">prometheus/client_golang#1322</a></li>
<li><a href="https://github.com/ArthurSens"><code>@​ArthurSens</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1323">prometheus/client_golang#1323</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.17.0 / 2023-09-27</h2>
<ul>
<li>[CHANGE] Minimum required go version is now 1.19 (we also test client_golang against new 1.21 version). <a href="https://redirect.github.com/prometheus/client_golang/issues/1325">#1325</a></li>
<li>[FEATURE] Add support for Created Timestamps in Counters, Summaries and Historams. <a href="https://redirect.github.com/prometheus/client_golang/issues/1313">#1313</a></li>
<li>[ENHANCEMENT] Enable detection of a native histogram without observations. <a href="https://redirect.github.com/prometheus/client_golang/issues/1314">#1314</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/fa1408ee351f6aba15c6d0207f7a0021eb3af406"><code>fa1408e</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1352">#1352</a> from prometheus/arthursens/cut-1.17.0</li>
<li><a href="https://github.com/prometheus/client_golang/commit/24a72b819f07e2abdd7aa2ed910a2d343c02f9a3"><code>24a72b8</code></a> Add changelog entry for 1.17</li>
<li><a href="https://github.com/prometheus/client_golang/commit/1bae6c1e6314f6a20be183a7277059630780232a"><code>1bae6c1</code></a> Deprecated comment should begin with &quot;Deprecated:&quot; (<a href="https://redirect.github.com/prometheus/client_golang/issues/1347">#1347</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/bbab8fe770ef961310d1c7ba045d68f3708e6463"><code>bbab8fe</code></a> Fix typos in comments, tests, and errors (<a href="https://redirect.github.com/prometheus/client_golang/issues/1346">#1346</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/df7fa494179f42a8c08cf21d95da9ae09564f907"><code>df7fa49</code></a> Extend Counters, Summaries and Histograms with creation timestamp (<a href="https://redirect.github.com/prometheus/client_golang/issues/1313">#1313</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/74cc26257c40439fda01b450a6def1bd69d77d1b"><code>74cc262</code></a> Add go_godebug_non_default_behavior_tlsmaxrsasize_events_total (<a href="https://redirect.github.com/prometheus/client_golang/issues/1348">#1348</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/d03abf3a31c973a5bc2c2dc698fb41b661a0f0c5"><code>d03abf3</code></a> Cleanup golangci-lint errcheck (<a href="https://redirect.github.com/prometheus/client_golang/issues/1339">#1339</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/ca6ba04f2fd1225668db5d01fa5a65bc2e668898"><code>ca6ba04</code></a> Update common Prometheus files (<a href="https://redirect.github.com/prometheus/client_golang/issues/1338">#1338</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/51d24f86807c24ce866355d0c82258420f571c17"><code>51d24f8</code></a> Update common Prometheus files (<a href="https://redirect.github.com/prometheus/client_golang/issues/1332">#1332</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/c17edf09ed77d45c7b0df50f2d9979e4f132a09e"><code>c17edf0</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1304">#1304</a> from prometheus/dependabot/go_modules/google.golang....</li>
<li>Additional commits viewable in <a href="https://github.com/prometheus/client_golang/compare/v1.16.0...v1.17.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.2.0 to 9.2.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.2.1</h2>
<h2>🧰 Maintenance</h2>
<ul>
<li>chore(deps): bump actions/stale from 3 to 8 (<a href="https://redirect.github.com/redis/go-redis/issues/2732">#2732</a>)</li>
<li>Add stream interface back to <code>Cmdable</code> (<a href="https://redirect.github.com/redis/go-redis/issues/2725">#2725</a>)</li>
<li>Remove redundant nil check in gears (<a href="https://redirect.github.com/redis/go-redis/issues/2728">#2728</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/Juneezee"><code>@​Juneezee</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/gabrielgio"><code>@​gabrielgio</code></a> and <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/f994ff1cd96299a5c8029ae3403af7b17ef06e8a"><code>f994ff1</code></a> Bump version to 9.2.1 (<a href="https://redirect.github.com/redis/go-redis/issues/2735">#2735</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/0b5e1866b1b37cfafd7e5d403446207c0143b6ec"><code>0b5e186</code></a> chore(deps): bump actions/stale from 3 to 8 (<a href="https://redirect.github.com/redis/go-redis/issues/2732">#2732</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c6fe509f4a0c6746544a6320f6d1237e74424a7a"><code>c6fe509</code></a> Add stream interface back to <code>Cmdable</code> (<a href="https://redirect.github.com/redis/go-redis/issues/2725">#2725</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/275af739718caca7af07dde35c234517d70b6b07"><code>275af73</code></a> refactor(gears): remove redundant nil check (<a href="https://redirect.github.com/redis/go-redis/issues/2728">#2728</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.2.0...v9.2.1">compare view</a></li>
</ul>
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
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Discussion

### Comment by @app-sre-bot on October 02, 2023 at 02:00 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 02, 2023 at 02:36 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on October 02, 2023 at 05:17 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/413*
