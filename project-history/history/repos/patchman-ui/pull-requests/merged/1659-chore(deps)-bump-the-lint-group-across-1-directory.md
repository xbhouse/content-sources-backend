---
type: pull_request
number: 1659
title: "chore(deps): bump the lint group across 1 directory with 2 updates"
state: merged
author: dependabot
created: 2026-06-11T18:34:21Z
updated: 2026-06-11T19:48:28Z
closed: 2026-06-11T19:48:18Z
merged: 2026-06-11T19:48:18Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-5ca40f8bc7
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1659
---

# Pull Request #1659: chore(deps): bump the lint group across 1 directory with 2 updates

**Author**: @dependabot
**Created**: June 11, 2026 at 06:34 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-5ca40f8bc7`

## Description

Bumps the lint group with 2 updates in the / directory: [stylelint](https://github.com/stylelint/stylelint) and [stylelint-scss](https://github.com/stylelint-scss/stylelint-scss).

Updates `stylelint` from 17.12.0 to 17.13.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.13.0</h2>
<p>It fixes 3 bugs, including a false negative one.</p>
<ul>
<li>Fixed: <code>declaration-block-no-duplicate-properties</code> false negatives for interleaved non-consecutive duplicates with <code>ignore: [&quot;consecutive-duplicates(-*)&quot;]</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9324">#9324</a>) (<a href="https://github.com/sarathfrancis90"><code>@​sarathfrancis90</code></a>).</li>
<li>Fixed: <code>selector-max-type</code> false positives for nested selectors (<a href="https://redirect.github.com/stylelint/stylelint/issues/9319">#9319</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>selector-type-no-unknown</code> false positives for <code>install</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9308">#9308</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.13.0 - 2026-06-06</h2>
<p>It fixes 3 bugs, including a false negative one.</p>
<ul>
<li>Fixed: <code>declaration-block-no-duplicate-properties</code> false negatives for interleaved non-consecutive duplicates with <code>ignore: [&quot;consecutive-duplicates(-*)&quot;]</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9324">#9324</a>) (<a href="https://github.com/sarathfrancis90"><code>@​sarathfrancis90</code></a>).</li>
<li>Fixed: <code>selector-max-type</code> false positives for nested selectors (<a href="https://redirect.github.com/stylelint/stylelint/pull/9319">#9319</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>selector-type-no-unknown</code> false positives for <code>install</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9308">#9308</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/7fcee2b3155adb43daa2078762a15c2d9a24e69b"><code>7fcee2b</code></a> Release 17.13.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9342">#9342</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/3b7287b2eb88474f63f6d17474e87ba836bf1f43"><code>3b7287b</code></a> Refactor to reuse shared utilities (<a href="https://redirect.github.com/stylelint/stylelint/issues/9337">#9337</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/8e889c3394f410916343a5e4c2b372ec514184d0"><code>8e889c3</code></a> Bump lint-staged from 17.0.4 to 17.0.5 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9334">#9334</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/a74aab46df5a94afa0cf4fddc9faf6eeacf34293"><code>a74aab4</code></a> Bump the stylelint-actions group with 5 updates (<a href="https://redirect.github.com/stylelint/stylelint/issues/9333">#9333</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/74c644828b839add46650d67390a0f1094dbcf7a"><code>74c6448</code></a> Fix <code>declaration-block-no-duplicate-properties</code> false negatives for interleav...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/1cd26ac54ed16b92f77487f85b9305ca320f4085"><code>1cd26ac</code></a> Skip changeset verification on fork PRs CI (<a href="https://redirect.github.com/stylelint/stylelint/issues/9331">#9331</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/712b9867f7a1fa631813f577d7941686a10d3a58"><code>712b986</code></a> Fix vulnerable dependencies via <code>npm audit fix</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9328">#9328</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/27196b71dbc255bd612ad26dee40bf276a8123b0"><code>27196b7</code></a> Fix CI badge in README.md (<a href="https://redirect.github.com/stylelint/stylelint/issues/9329">#9329</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/179bba2e5a3264f9e96baf3318afdf113ea4f8dc"><code>179bba2</code></a> Refactor to use <code>@import</code> over <code>@typedef</code> for simple imports (<a href="https://redirect.github.com/stylelint/stylelint/issues/9326">#9326</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/94eab544e488dddb43829353649da8debd6427b5"><code>94eab54</code></a> Document using our PR template (<a href="https://redirect.github.com/stylelint/stylelint/issues/9327">#9327</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/17.12.0...17.13.0">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint-scss` from 7.1.1 to 7.2.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/releases">stylelint-scss's releases</a>.</em></p>
<blockquote>
<h2>7.2.0</h2>
<ul>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> autofix for Sass functions in custom properties (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1287">#1287</a>).</li>
<li>Added: <code>declaration-property-value-no-unknown</code> support for namespaced functions (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1289">#1289</a>).</li>
<li>Updated: <code>at-rule-conditional-no-parentheses</code> improve autofix (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1286">#1286</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v7.1.1...v7.2.0">https://github.com/stylelint-scss/stylelint-scss/compare/v7.1.1...v7.2.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/blob/master/CHANGELOG.md">stylelint-scss's changelog</a>.</em></p>
<blockquote>
<h1>7.2.0</h1>
<ul>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> autofix for Sass functions in custom properties (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1287">#1287</a>).</li>
<li>Added: <code>declaration-property-value-no-unknown</code> support for namespaced functions (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1289">#1289</a>).</li>
<li>Updated: <code>at-rule-conditional-no-parentheses</code> improve autofix (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1286">#1286</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v7.1.1...v7.2.0">https://github.com/stylelint-scss/stylelint-scss/compare/v7.1.1...v7.2.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/e7c92d143a99cac7822bc99632f6064037b7a9a9"><code>e7c92d1</code></a> 7.2.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/ecdc3c37fe99119fa37dc0dbd5fa833aad93b393"><code>ecdc3c3</code></a> Prepare version 7.2.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1301">#1301</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/ddffe930d88339827a6dedcfb88d8d3c41539a5f"><code>ddffe93</code></a> build(deps-dev): bump stylelint from 17.10.0 to 17.12.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1300">#1300</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/6eb763054361ae49bd521237060903ea9b9efa3a"><code>6eb7630</code></a> Update Jest to latest version (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1299">#1299</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/ac4dffb473979c1d520069242e10f9e052b8be83"><code>ac4dffb</code></a> build(deps-dev): bump <code>@​lavamoat/allow-scripts</code> from 5.0.1 to 5.0.2 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1298">#1298</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/ac1e7faca9d5145b99b985d77aec1a1c18bce0aa"><code>ac1e7fa</code></a> build(deps): bump <code>@​csstools/css-calc</code> from 3.2.0 to 3.2.1 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1297">#1297</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/b483811b8af2d577a8950b23250257ad0162fdc9"><code>b483811</code></a> build(deps-dev): bump postcss from 8.5.14 to 8.5.15 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1295">#1295</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/2c8816586fd3b3df8a88e57a10ef5d8684e5c27e"><code>2c88165</code></a> build(deps-dev): bump lint-staged from 16.4.0 to 17.0.5 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1292">#1292</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/f4aea2666ee4840c7cf27245a8379c2f38db0db0"><code>f4aea26</code></a> build(deps): bump <code>@​csstools/css-syntax-patches-for-csstree</code> (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1294">#1294</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/db3c20ed4598659e019f4c52cfecaaddb84ca611"><code>db3c20e</code></a> build(deps-dev): bump eslint from 10.3.0 to 10.4.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1293">#1293</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v7.1.1...v7.2.0">compare view</a></li>
</ul>
</details>
<br />


---

## Reviews

### Review by @swadeley - Approved on June 11, 2026 at 07:24 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1659*
