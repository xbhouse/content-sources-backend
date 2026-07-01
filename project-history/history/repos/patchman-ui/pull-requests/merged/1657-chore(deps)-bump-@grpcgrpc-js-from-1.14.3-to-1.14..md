---
type: pull_request
number: 1657
title: "chore(deps): bump @grpc/grpc-js from 1.14.3 to 1.14.4"
state: merged
author: dependabot
created: 2026-06-11T14:14:47Z
updated: 2026-06-11T18:56:11Z
closed: 2026-06-11T18:56:03Z
merged: 2026-06-11T18:56:03Z
base_branch: master
head_branch: dependabot/npm_and_yarn/grpc/grpc-js-1.14.4
labels: ["dependencies", "patch"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1657
---

# Pull Request #1657: chore(deps): bump @grpc/grpc-js from 1.14.3 to 1.14.4

**Author**: @dependabot
**Created**: June 11, 2026 at 02:14 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `patch`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/grpc/grpc-js-1.14.4`

## Description

Bumps [@grpc/grpc-js](https://github.com/grpc/grpc-node) from 1.14.3 to 1.14.4.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-node/releases">@​grpc/grpc-js's releases</a>.</em></p>
<blockquote>
<h2><code>@​grpc/grpc-js</code> 1.14.4</h2>
<ul>
<li>Fix a bug that could cause servers to crash when handling malformed requests (<a href="https://github.com/grpc/grpc-node/security/advisories/GHSA-5375-pq7m-f5r2">advisory GHSA-5375-pq7m-f5r2</a>)</li>
<li>Fix a bug that could cause clients and servers to crash when handling malformed compressed messages (<a href="https://github.com/grpc/grpc-node/security/advisories/GHSA-99f4-grh7-6pcq">advisory GHSA-99f4-grh7-6pcq</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-node/commit/a380735ba9b0351214f2faa578350a559dd486ff"><code>a380735</code></a> Merge pull request <a href="https://redirect.github.com/grpc/grpc-node/issues/3052">#3052</a> from murgatroid99/grpc-js_1.14.4</li>
<li><a href="https://github.com/grpc/grpc-node/commit/5b8d37b03d91122ec0b9bc5e27dd26ffa7448337"><code>5b8d37b</code></a> Merge commit from fork</li>
<li><a href="https://github.com/grpc/grpc-node/commit/6a97456cc88d2b74e1527b356de98bf8ee8d7a40"><code>6a97456</code></a> Merge commit from fork</li>
<li><a href="https://github.com/grpc/grpc-node/commit/e5e0b1d3ff14fa7c5eeef10b309d694bc3ff7e96"><code>e5e0b1d</code></a> grpc-js: Bump version to 1.14.4</li>
<li><a href="https://github.com/grpc/grpc-node/commit/5029a2668164d1ba6de6ed4dcf6d35d5c4ff6cf4"><code>5029a26</code></a> Make compression error a static string</li>
<li><a href="https://github.com/grpc/grpc-node/commit/2fe55fd76a8bb59eaab5f39e3552b5f84985a163"><code>2fe55fd</code></a> Fix crashes when receiving malformed compressed data</li>
<li><a href="https://github.com/grpc/grpc-node/commit/234f9172b2ff35e586ca7d4e788557aad5985668"><code>234f917</code></a> Fix server crash when handling invalid requests</li>
<li><a href="https://github.com/grpc/grpc-node/commit/acef8d4adfa091188e9dd572cedf4d87b0f69b21"><code>acef8d4</code></a> Merge pull request <a href="https://redirect.github.com/grpc/grpc-node/issues/3043">#3043</a> from murgatroid99/rbac_types_change_fix_1.14</li>
<li><a href="https://github.com/grpc/grpc-node/commit/4f3c58fda2136eb0038a39d54804acb06a8419ea"><code>4f3c58f</code></a> grpc-js-xds: Update RBAC code to handle Node type change, pin <code>@​types/node</code></li>
<li>See full diff in <a href="https://github.com/grpc/grpc-node/compare/@grpc/grpc-js@1.14.3...@grpc/grpc-js@1.14.4">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on June 11, 2026 at 06:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1657?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.27%. Comparing base ([`49d0fb1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/49d0fb1e3781c62e9029b9593a6cc546021cf53b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`da9c3e6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/da9c3e69eea11828d10e01b91931f53ab7fd07ac?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1657   +/-   ##
=======================================
  Coverage   77.27%   77.27%           
=======================================
  Files         103      103           
  Lines        3287     3287           
  Branches      740      735    -5     
=======================================
  Hits         2540     2540           
  Misses        668      668           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1657?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on June 11, 2026 at 06:49 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1657*
