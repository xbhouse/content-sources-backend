---
type: pull_request
number: 1331
title: "chore(deps): bump the lint group across 1 directory with 4 updates"
state: closed
author: dependabot
created: 2025-05-06T07:37:59Z
updated: 2025-05-07T13:49:27Z
closed: 2025-05-07T13:49:25Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-72fc9c27bc
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1331
---

# Pull Request #1331: chore(deps): bump the lint group across 1 directory with 4 updates

**Author**: @dependabot
**Created**: May 06, 2025 at 07:37 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-72fc9c27bc`

## Description

Bumps the lint group with 4 updates in the / directory: [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier), [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react), [stylelint](https://github.com/stylelint/stylelint) and [stylelint-scss](https://github.com/stylelint-scss/stylelint-scss).

Updates `eslint-config-prettier` from 7.2.0 to 10.1.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/releases">eslint-config-prettier's releases</a>.</em></p>
<blockquote>
<h2>v10.1.2</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> Thanks <a href="https://github.com/Fdawgs"><code>@​Fdawgs</code></a>! - chore(package): add homepage for some 3rd-party registry - see <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> for more details</li>
</ul>
<h2>v10.1.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/309">#309</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - fix: separate the <code>/flat</code> entry for compatibility</p>
<p>For flat config users, the previous <code>&quot;eslint-config-prettier&quot;</code> entry still works, but <code>&quot;eslint-config-prettier/flat&quot;</code> adds a new <code>name</code> property for <a href="https://eslint.org/blog/2024/04/eslint-config-inspector/">config-inspector</a>, we just can't add it for the default entry for compatibility.</p>
<p>See also <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/308">prettier/eslint-config-prettier#308</a></p>
<pre lang="ts"><code>// before
import eslintConfigPrettier from &quot;eslint-config-prettier&quot;;
<p>// after<br />
import eslintConfigPrettier from &quot;eslint-config-prettier/flat&quot;;<br />
</code></pre></p>
</li>
</ul>
<h2>v10.1.0</h2>
<h3>Minor Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/306">#306</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/56e2e3466391d0fdfc200e42130309c687aaab53"><code>56e2e34</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - feat: migrate to exports field</li>
</ul>
<h2>v10.0.3</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/294">#294</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/8dbbd6d70b8a56cdfa4ea4e185d3699d5729b38e"><code>8dbbd6d</code></a> Thanks <a href="https://github.com/FloEdelmann"><code>@​FloEdelmann</code></a>! - feat: add name to config</p>
</li>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/280">#280</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/cba57377e4c86d20d17042d6999eabc754fddc03"><code>cba5737</code></a> Thanks <a href="https://github.com/zanminkian"><code>@​zanminkian</code></a>! - feat: add declaration file</p>
</li>
</ul>
<h3>New Contributors</h3>
<ul>
<li><a href="https://github.com/zanminkian"><code>@​zanminkian</code></a> made their first contribution in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/280">prettier/eslint-config-prettier#280</a></li>
<li><a href="https://github.com/FloEdelmann"><code>@​FloEdelmann</code></a> made their first contribution in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/294">prettier/eslint-config-prettier#294</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.0.2...v10.0.3">https://github.com/prettier/eslint-config-prettier/compare/v10.0.2...v10.0.3</a></p>
<h2>v10.0.2</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/299">#299</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/e750edc530c816e0b3ffabfab1f4e46532bccbfe"><code>e750edc</code></a> Thanks <a href="https://github.com/Fdawgs"><code>@​Fdawgs</code></a>! - chore(package): explicitly declare js module type</li>
</ul>
<h2>v10.0.1</h2>
<h1>eslint-config-prettier</h1>
<h2>10.0.1</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/blob/main/CHANGELOG.md">eslint-config-prettier's changelog</a>.</em></p>
<blockquote>
<h2>10.1.2</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> Thanks <a href="https://github.com/Fdawgs"><code>@​Fdawgs</code></a>! - chore(package): add homepage for some 3rd-party registry - see <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> for more details</li>
</ul>
<h2>10.1.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/309">#309</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - fix: separate the <code>/flat</code> entry for compatibility</p>
<p>For flat config users, the previous <code>&quot;eslint-config-prettier&quot;</code> entry still works, but <code>&quot;eslint-config-prettier/flat&quot;</code> adds a new <code>name</code> property for <a href="https://eslint.org/blog/2024/04/eslint-config-inspector/">config-inspector</a>, we just can't add it for the default entry for compatibility.</p>
<p>See also <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/308">prettier/eslint-config-prettier#308</a></p>
<pre lang="ts"><code>// before
import eslintConfigPrettier from &quot;eslint-config-prettier&quot;;
<p>// after<br />
import eslintConfigPrettier from &quot;eslint-config-prettier/flat&quot;;<br />
</code></pre></p>
</li>
</ul>
<h2>10.1.0</h2>
<h3>Minor Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/306">#306</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/56e2e3466391d0fdfc200e42130309c687aaab53"><code>56e2e34</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - feat: migrate to exports field</li>
</ul>
<h2>10.0.3</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/294">#294</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/8dbbd6d70b8a56cdfa4ea4e185d3699d5729b38e"><code>8dbbd6d</code></a> Thanks <a href="https://github.com/FloEdelmann"><code>@​FloEdelmann</code></a>! - feat: add name to config</p>
</li>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/280">#280</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/cba57377e4c86d20d17042d6999eabc754fddc03"><code>cba5737</code></a> Thanks <a href="https://github.com/zanminkian"><code>@​zanminkian</code></a>! - feat: add declaration file</p>
</li>
</ul>
<h2>10.0.2</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/299">#299</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/e750edc530c816e0b3ffabfab1f4e46532bccbfe"><code>e750edc</code></a> Thanks <a href="https://github.com/Fdawgs"><code>@​Fdawgs</code></a>! - chore(package): explicitly declare js module type</li>
</ul>
<h2>10.0.0</h2>
<h3>Major Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/272">#272</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/5be64bef68c3a9bf7202f591f54ffec02572e46b"><code>5be64be</code></a> Thanks <a href="https://github.com/abrahamguo"><code>@​abrahamguo</code></a>! - add support for <a href="https://github.com/stylistic"><code>@​stylistic</code></a> formatting rules</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/8911369cbc66f1f859e19751eaefdea687129de5"><code>8911369</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/322">#322</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> chore(package): add homepage url (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/321">#321</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4ae04c0dea72dd7d950bb575a8d87a90ab5ea787"><code>4ae04c0</code></a> chore(deps): update yarn to v4.8.1 (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/320">#320</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/7499c2e0a39dca3e64a8c03b3e6c68341ed829e8"><code>7499c2e</code></a> chore: ignore <code>eslint-find-rules</code> for ESLint 8 compatibility</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/241c6b8977859ab896bdf7efb431f1a145095480"><code>241c6b8</code></a> chore:  housekeeping, upgrade all (dev) dependencies (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/319">#319</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/9156ab4cf022bb66a2141462a483520ef5b24a65"><code>9156ab4</code></a> chore: add renovate preset</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/f12309bbca9fb051b53fcece9a8491a1222235c8"><code>f12309b</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/310">#310</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/79cf6796bf9de35eb8965333145ca6e989045580"><code>79cf679</code></a> chore: use flat entry for flat config verification</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> fix: separate the <code>/flat</code> entry for compatibility (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/309">#309</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/2c6f80e67f391db7650e5ba32c7a5562d2029664"><code>2c6f80e</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/307">#307</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prettier/eslint-config-prettier/compare/v7.2.0...v10.1.2">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~jounqin">jounqin</a>, a new releaser for eslint-config-prettier since your current version.</p>
</details>
<br />

Updates `eslint-plugin-react` from 7.37.4 to 7.37.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jsx-eslint/eslint-plugin-react/releases">eslint-plugin-react's releases</a>.</em></p>
<blockquote>
<h2>v7.37.5</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-unknown-property</code>]: allow shadow root attrs on <code>\&lt;template&gt;</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
<li>[<code>prop-types</code>]: support <code>ComponentPropsWithRef</code> from a namespace import (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>[] <a href="https://github.com/corydeppen"><code>@​corydeppen</code></a>)</li>
<li>[<code>jsx-no-constructed-context-values</code>]: detect constructed context values in React 19 <code>&lt;Context&gt;</code> usage (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>[] <a href="https://github.com/TildaDares"><code>@​TildaDares</code></a>)</li>
<li>[<code>no-unknown-property</code>]: allow <code>transform-origin</code> on <code>rect</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[Docs] [<code>button-has-type</code>]: clean up phrasing (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>[] <a href="https://github.com/hamirmahal"><code>@​hamirmahal</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3651">jsx-eslint/eslint-plugin-react#3651</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3909">jsx-eslint/eslint-plugin-react#3909</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3910">jsx-eslint/eslint-plugin-react#3910</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">jsx-eslint/eslint-plugin-react#3912</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">jsx-eslint/eslint-plugin-react#3914</a>
[<code>button-has-type</code>]: docs/rules/button-has-type.md
[<code>jsx-no-constructed-context-values</code>]: docs/rules/jsx-no-constructed-context-values.md
[<code>no-unknown-property</code>]: docs/rules/no-unknown-property.md
[<code>prop-types</code>]: docs/rules/prop-types.md</p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jsx-eslint/eslint-plugin-react/blob/master/CHANGELOG.md">eslint-plugin-react's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.4...v7.37.5">7.37.5</a> - 2025.04.03</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-unknown-property</code>]: allow shadow root attrs on <code>\&lt;template&gt;</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
<li>[<code>prop-types</code>]: support <code>ComponentPropsWithRef</code> from a namespace import (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>[] <a href="https://github.com/corydeppen"><code>@​corydeppen</code></a>)</li>
<li>[<code>jsx-no-constructed-context-values</code>]: detect constructed context values in React 19 <code>&lt;Context&gt;</code> usage (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>[] <a href="https://github.com/TildaDares"><code>@​TildaDares</code></a>)</li>
<li>[<code>no-unknown-property</code>]: allow <code>transform-origin</code> on <code>rect</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[Docs] [<code>button-has-type</code>]: clean up phrasing (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>[] <a href="https://github.com/hamirmahal"><code>@​hamirmahal</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">jsx-eslint/eslint-plugin-react#3914</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">jsx-eslint/eslint-plugin-react#3912</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3910">jsx-eslint/eslint-plugin-react#3910</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3909">jsx-eslint/eslint-plugin-react#3909</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3651">jsx-eslint/eslint-plugin-react#3651</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/2c98b83c451a4297edf1787d9a616e50687e27e8"><code>2c98b83</code></a> Update CHANGELOG and bump version</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/2f64deadac51b42fc1a8660fad026ac4c68b92f3"><code>2f64dea</code></a> [Fix] <code>no-unknown-property</code>: allow <code>transform-origin</code> on <code>rect</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/2428618b5a7334b96b7b7eb9629212d07b6fd510"><code>2428618</code></a> [Fix] <code>jsx-no-constructed-context-values</code>: detect constructed context values ...</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/60b731621c98b8d3f6c8c5339a50dc54bf3fd068"><code>60b7316</code></a> [Tests] <code>prop-types</code>: use proper spacing/semis, button type</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/90a00b9318374b402114a4136c6f118b48d9346e"><code>90a00b9</code></a> [Fix] <code>prop-types</code>: support <code>ComponentPropsWithRef</code> from a namespace import</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/3fd9b9223e3f4fc6b34eb6f3ab734a7e2c73743d"><code>3fd9b92</code></a> [Fix] <code>no-unknown-property</code>: allow shadow root attrs on <code>\&lt;template&gt;</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/addad4687b710c022f868ea17f6cabfaaddd8b44"><code>addad46</code></a> [Deps] update <code>object.entries</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/98a31f8e76a4d8aa52caeeb55940f35682b18b2f"><code>98a31f8</code></a> [Dev Deps] update <code>@babel/core</code>, <code>@babel/eslint-parser</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/7eb6ca9144333c828f24abdc98154a45aec46d54"><code>7eb6ca9</code></a> [Docs] <code>button-has-type</code>: clean up phrasing</li>
<li>See full diff in <a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.4...v7.37.5">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 16.14.1 to 16.19.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>16.19.1</h2>
<ul>
<li>Fixed: <code>no-empty-source</code> false positives for non-standard syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8548">#8548</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
</ul>
<h2>16.19.0</h2>
<p>It adds 2 options to 2 rules and fixes 3 bugs.</p>
<ul>
<li>Added: <code>exceptWithoutPropertyFallback: []</code> to <code>function-allowed-list</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8488">#8488</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignore: [&quot;four-into-three-edge-values&quot;]</code> to <code>shorthand-property-no-redundant-values</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8527">#8527</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>compact</code> formatter with pnpm to newline the exit code (<a href="https://redirect.github.com/stylelint/stylelint/issues/8534">#8534</a>) (<a href="https://github.com/konomae"><code>@​konomae</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> range and message for invalid syntax within known functions (<a href="https://redirect.github.com/stylelint/stylelint/issues/8528">#8528</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>no-empty-source</code> false positives for <code>--report-needless-disables</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8536">#8536</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
</ul>
<h2>16.18.0</h2>
<p>It adds 2 new rules and fixes 2 bugs. We've turned on these rules, and the <code>syntax-string-no-invalid</code> and <code>layer-name-pattern</code> ones from recent releases, in our <a href="https://www.npmjs.com/package/stylelint-config-standard">standard config</a>.</p>
<ul>
<li>Added: <code>color-function-alias-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8499">#8499</a>) (<a href="https://github.com/EduardAkhmetshin"><code>@​EduardAkhmetshin</code></a>).</li>
<li>Added: <code>container-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8498">#8498</a>) (<a href="https://github.com/nate10j"><code>@​nate10j</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>math</code> of <code>font-size</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8495">#8495</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
<li>Fixed: <code>font-family-no-missing-generic-family-keyword</code> false positives for <code>math</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8489">#8489</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
</ul>
<h2>16.17.0</h2>
<p>It adds 1 new rule, support for <code>languageOptions</code> to 2 rules, 1 option to a rule, the <code>--compute-edit-info</code> CLI flag (along with support for <code>EditInfo</code> in 3 rules), and fixes 1 bug. <code>EditInfo</code> is useful for automated fixing tools and editor integrations.</p>
<ul>
<li>Added: <code>layer-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8474">#8474</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>--compute-edit-info</code> CLI flag (<a href="https://redirect.github.com/stylelint/stylelint/issues/8473">#8473</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignorePreludeOfAtRules: []</code> to <code>length-zero-no-unit</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8472">#8472</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8475">#8475</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>property-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8476">#8476</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>declaration-block-no-redundant-longhand-properties</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8482">#8482</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>function-url-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8483">#8483</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-attribute-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8484">#8484</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Fixed: <code>custom-property-pattern</code> false negatives for <code>@property</code> preludes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8468">#8468</a>) (<a href="https://github.com/rohitgs28"><code>@​rohitgs28</code></a>).</li>
</ul>
<h2>16.16.0</h2>
<p>It adds support for computing <code>EditInfo</code> to 22 more rules and reverts a change that added <code>context.lexer</code> to our public API in the previous release.</p>
<ul>
<li>Added: <code>at-rule-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8425">#8425</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-deprecated</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8426">#8426</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8427">#8427</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>color-function-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8437">#8437</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8443">#8443</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-property-value-keyword-no-deprecated</code> support for computing <code>EditInfo</code>. (<a href="https://redirect.github.com/stylelint/stylelint/issues/8439">#8439</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>font-family-name-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8419">#8419</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>font-weight-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8420">#8420</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>function-calc-no-unspaced-operator</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8440">#8440</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>function-name-case</code> support for support for computing <code>EditInfo</code>.&quot; (<a href="https://redirect.github.com/stylelint/stylelint/issues/8442">#8442</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>hue-degree-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8444">#8444</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>import-notation</code> support for computing <code>EditInfo</code>. (<a href="https://redirect.github.com/stylelint/stylelint/issues/8445">#8445</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>keyframe-selector-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8457">#8457</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>length-zero-no-unit</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8459">#8459</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>16.19.1 - 2025-04-25</h2>
<ul>
<li>Fixed: <code>no-empty-source</code> false positives for non-standard syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/pull/8548">#8548</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
</ul>
<h2>16.19.0 - 2025-04-23</h2>
<p>It adds 2 options to 2 rules and fixes 3 bugs.</p>
<ul>
<li>Added: <code>exceptWithoutPropertyFallback: []</code> to <code>function-allowed-list</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8488">#8488</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignore: [&quot;four-into-three-edge-values&quot;]</code> to <code>shorthand-property-no-redundant-values</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8527">#8527</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>compact</code> formatter with pnpm to newline the exit code (<a href="https://redirect.github.com/stylelint/stylelint/pull/8534">#8534</a>) (<a href="https://github.com/konomae"><code>@​konomae</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> range and message for invalid syntax within known functions (<a href="https://redirect.github.com/stylelint/stylelint/pull/8528">#8528</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>no-empty-source</code> false positives for <code>--report-needless-disables</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8536">#8536</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
</ul>
<h2>16.18.0 - 2025-04-06</h2>
<p>It adds 2 new rules and fixes 2 bugs. We've turned on these rules, and the <code>syntax-string-no-invalid</code> and <code>layer-name-pattern</code> ones from recent releases, in our <a href="https://www.npmjs.com/package/stylelint-config-standard">standard config</a>.</p>
<ul>
<li>Added: <code>color-function-alias-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8499">#8499</a>) (<a href="https://github.com/EduardAkhmetshin"><code>@​EduardAkhmetshin</code></a>).</li>
<li>Added: <code>container-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8498">#8498</a>) (<a href="https://github.com/nate10j"><code>@​nate10j</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>math</code> of <code>font-size</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8495">#8495</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
<li>Fixed: <code>font-family-no-missing-generic-family-keyword</code> false positives for <code>math</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8489">#8489</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
</ul>
<h2>16.17.0 - 2025-03-26</h2>
<p>It adds 1 new rule, support for <code>languageOptions</code> to 2 rules, 1 option to a rule, the <code>--compute-edit-info</code> CLI flag (along with support for <code>EditInfo</code> in 3 rules), and fixes 1 bug. <code>EditInfo</code> is useful for automated fixing tools and editor integrations.</p>
<ul>
<li>Added: <code>layer-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8474">#8474</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>--compute-edit-info</code> CLI flag (<a href="https://redirect.github.com/stylelint/stylelint/pull/8473">#8473</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignorePreludeOfAtRules: []</code> to <code>length-zero-no-unit</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8472">#8472</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8475">#8475</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>property-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8476">#8476</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>declaration-block-no-redundant-longhand-properties</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8482">#8482</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>function-url-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8483">#8483</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-attribute-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8484">#8484</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Fixed: <code>custom-property-pattern</code> false negatives for <code>@property</code> preludes (<a href="https://redirect.github.com/stylelint/stylelint/pull/8468">#8468</a>) (<a href="https://github.com/rohitgs28"><code>@​rohitgs28</code></a>).</li>
</ul>
<h2>16.16.0 - 2025-03-14</h2>
<p>It adds support for computing <code>EditInfo</code> to 22 more rules and reverts a change that added <code>context.lexer</code> to our public API in the previous release.</p>
<ul>
<li>Added: <code>at-rule-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8425">#8425</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-deprecated</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8426">#8426</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8427">#8427</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>color-function-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8437">#8437</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8443">#8443</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-property-value-keyword-no-deprecated</code> support for computing <code>EditInfo</code>. (<a href="https://redirect.github.com/stylelint/stylelint/pull/8439">#8439</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>font-family-name-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8419">#8419</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>font-weight-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8420">#8420</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>function-calc-no-unspaced-operator</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8440">#8440</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/25968c6066d0a747b6522252c7ea8c35b36883ea"><code>25968c6</code></a> 16.19.1</li>
<li><a href="https://github.com/stylelint/stylelint/commit/015042b5d20e51818ccf17acfe342db6afb94fd5"><code>015042b</code></a> Prepare 16.19.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8551">#8551</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/d4c45d33a4888b0e9c6db87d0091ea00fd4b017d"><code>d4c45d3</code></a> Refactor to allow node for <code>isConfigurationComment()</code> private utility (<a href="https://redirect.github.com/stylelint/stylelint/issues/8549">#8549</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/013c0c82addbe8026c72cdbd1098d00b4741ee26"><code>013c0c8</code></a> Fix <code>no-empty-source</code> false positives for non-standard syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8548">#8548</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/4c63b3939faa4240ca20087e148eca7f2bd8d7a2"><code>4c63b39</code></a> 16.19.0</li>
<li><a href="https://github.com/stylelint/stylelint/commit/f7fb8e71d52d2bb30d506b25896281e28aefbbee"><code>f7fb8e7</code></a> Prepare 16.19.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8523">#8523</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/3c0c86c94ed2f9a30c398628d0d2424a4f10e593"><code>3c0c86c</code></a> Fix <code>compact</code> formatter with pnpm to newline the exit code (<a href="https://redirect.github.com/stylelint/stylelint/issues/8534">#8534</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/660f81844dbbfb98fbdccbc52bc5721b031087ea"><code>660f818</code></a> Bump eslint from 9.24.0 to 9.25.0 in the eslint group (<a href="https://redirect.github.com/stylelint/stylelint/issues/8541">#8541</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/ec3ab9cd783b06dda502f6b5c31696d2a6a451fb"><code>ec3ab9c</code></a> Bump <code>@​changesets/cli</code> from 2.28.1 to 2.29.2 in the changesets group (<a href="https://redirect.github.com/stylelint/stylelint/issues/8540">#8540</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/a512b36677ff11b6dbf1808ef31fe8cc7b0c97f8"><code>a512b36</code></a> Bump codecov/codecov-action from 5.4.0 to 5.4.2 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8542">#8542</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/16.14.1...16.19.1">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint-scss` from 6.11.0 to 6.12.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/releases">stylelint-scss's releases</a>.</em></p>
<blockquote>
<h2>6.12.0</h2>
<ul>
<li>Added: <code>double-slash-comment-whitespace-inside</code> add autofix and fix incorrect error locations (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1146">#1146</a>).</li>
<li>Added: <code>declaration-property-value-no-unknown</code> add support for nested properties with shorthand values (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1142">#1142</a>).</li>
<li>Fixed: <code>operator-no-unspaced</code> don't check Tailwind Directives (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1147">#1147</a>).</li>
<li>Fixed: <code>at-use-no-redundant-alias</code> fix false negative for single quotes (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1130">#1130</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.1...v6.12.0">https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.1...v6.12.0</a></p>
<h2>6.11.1</h2>
<ul>
<li>Fixed: <code>no-duplicate-load-rules</code> fix false positive when using <code>@use</code> and <code>@forward</code> for the same stylesheet (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1124">#1124</a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> add support for Sass' <code>rgba()</code> function with a hex value (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1123">#1123</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.0...v6.11.1">https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.0...v6.11.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/blob/master/CHANGELOG.md">stylelint-scss's changelog</a>.</em></p>
<blockquote>
<h1>6.12.0</h1>
<ul>
<li>Added: <code>double-slash-comment-whitespace-inside</code> add autofix and fix incorrect error locations (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1146">#1146</a>).</li>
<li>Added: <code>declaration-property-value-no-unknown</code> add support for nested properties with shorthand values (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1142">#1142</a>).</li>
<li>Fixed: <code>operator-no-unspaced</code> don't check Tailwind Directives (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1147">#1147</a>).</li>
<li>Fixed: <code>at-use-no-redundant-alias</code> fix false negative for single quotes (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1130">#1130</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.1...v6.12.0">https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.1...v6.12.0</a></p>
<h1>6.11.1</h1>
<ul>
<li>Fixed: <code>no-duplicate-load-rules</code> fix false positive when using <code>@use</code> and <code>@forward</code> for the same stylesheet (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1124">#1124</a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> add support for Sass' <code>rgba()</code> function with a hex value (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1123">#1123</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.0...v6.11.1">https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.0...v6.11.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/d1deb4fe964a1ad483851e544405d1f35b96f7c2"><code>d1deb4f</code></a> 6.12.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/2ceffc2ae8db808d710a1e337d3c3500d11d82a5"><code>2ceffc2</code></a> Prepare version 6.12.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/9ec5456f36a62c4a821170f0ba25e0c832efb6de"><code>9ec5456</code></a> double-slash-comment-whitespace-inside: add autofix to README</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/d586476a77a45c59dec9ed1572223a2cc41794c8"><code>d586476</code></a> Update contributors list</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/e4b8fd4d0dbc1f1abc5732ba322131c4a6813430"><code>e4b8fd4</code></a> build(deps): bump known-css-properties from 0.35.0 to 0.36.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1148">#1148</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/474af2303a2c0ad042d233d3f849f3f6243622e5"><code>474af23</code></a> build(deps-dev): bump jest-preset-stylelint from 7.2.0 to 7.3.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1138">#1138</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/a68c5ab8f364a3ebd3b191737ed5bd05b0493b74"><code>a68c5ab</code></a> operator-no-unspaced: don't check Tailwind Directives (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1147">#1147</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/762387a1fd4f73c9808ba153cfdd7d7337122ad9"><code>762387a</code></a> Improve double-slash-comment-whitespace-inside (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1146">#1146</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/9d9f088275e977989d4cfac9dedae67b6a7c67d1"><code>9d9f088</code></a> build(deps-dev): bump stylelint from 16.16.0 to 16.19.1 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1143">#1143</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/bdfa665dd70a21ab77608df4d1bc2c47defee339"><code>bdfa665</code></a> declaration-property-value-no-unknown support for nested properties with shor...</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.11.0...v6.12.0">compare view</a></li>
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

### Comment by @codecov-commenter on May 06, 2025 at 07:44 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1331?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 100.00%. Comparing base [(`610b99b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/610b99babb66500e77e79cac7ac5a30dbbe93f6c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`b943424`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b943424e3f55018aeea1cc9b489ef0d3c103a22f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##            master     #1331   +/-   ##
=========================================
  Coverage   100.00%   100.00%           
=========================================
  Files            1         1           
  Lines            6         6           
  Branches         2         2           
=========================================
  Hits             6         6           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1331/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1331/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1331/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1331?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @dependabot on May 07, 2025 at 01:49 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1331*
