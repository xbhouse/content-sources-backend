---
type: pull_request
number: 1448
title: "chore(deps): bump the lint group across 1 directory with 7 updates"
state: closed
author: dependabot
created: 2025-12-08T08:37:12Z
updated: 2026-01-20T18:44:38Z
closed: 2026-01-20T18:44:36Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-961510d8c5
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1448
---

# Pull Request #1448: chore(deps): bump the lint group across 1 directory with 7 updates

**Author**: @dependabot
**Created**: December 08, 2025 at 08:37 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-961510d8c5`

## Description

Bumps the lint group with 7 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [eslint](https://github.com/eslint/eslint) | `9.37.0` | `9.39.1` |
| [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier) | `9.1.2` | `10.1.8` |
| [eslint-plugin-playwright](https://github.com/playwright-community/eslint-plugin-playwright) | `2.3.0` | `2.4.0` |
| [eslint-plugin-unused-imports](https://github.com/sweepline/eslint-plugin-unused-imports) | `4.2.0` | `4.3.0` |
| [stylelint](https://github.com/stylelint/stylelint) | `16.23.1` | `16.26.1` |
| [stylelint-config-recommended-scss](https://github.com/stylelint-scss/stylelint-config-recommended-scss) | `14.1.0` | `16.0.2` |
| [stylelint-scss](https://github.com/stylelint-scss/stylelint-scss) | `6.12.1` | `6.13.0` |


Updates `eslint` from 9.37.0 to 9.39.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v9.39.1</h2>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/650753ee3976784343ceb40170619dab1aa9fe0d"><code>650753e</code></a> fix: Only pass node to JS lang visitor methods (<a href="https://redirect.github.com/eslint/eslint/issues/20283">#20283</a>) (Nicholas C. Zakas)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/51b51f4f1ce82ef63264c4e45d9ef579bcd73f8e"><code>51b51f4</code></a> docs: add a section on when to use extends vs cascading (<a href="https://redirect.github.com/eslint/eslint/issues/20268">#20268</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/b44d42699dcd1729b7ecb50ca70e4c1c17f551f1"><code>b44d426</code></a> docs: Update README (GitHub Actions Bot)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/92db329211c8da5ce8340a4d4c05ce9c12845381"><code>92db329</code></a> chore: update <code>@eslint/js</code> version to 9.39.1 (<a href="https://redirect.github.com/eslint/eslint/issues/20284">#20284</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/c7ebefc9eaf99b76b30b0d3cf9960807a47367c4"><code>c7ebefc</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/61778f6ca33c0f63962a91d6a75a4fa5db9f47d2"><code>61778f6</code></a> chore: update eslint-config-eslint dependency <code>@​eslint/js</code> to ^9.39.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20275">#20275</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/d9ca2fcd9ad63331bfd329a69534e1ff04f231e8"><code>d9ca2fc</code></a> ci: Add rangeStrategy to eslint group in renovate config (<a href="https://redirect.github.com/eslint/eslint/issues/20266">#20266</a>) (唯然)</li>
<li><a href="https://github.com/eslint/eslint/commit/009e5076ff5a4bd845f55e17676e3bb88f47c280"><code>009e507</code></a> test: fix version tests for ESLint v10 (<a href="https://redirect.github.com/eslint/eslint/issues/20274">#20274</a>) (Milos Djermanovic)</li>
</ul>
<h2>v9.39.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/cc57d87a3f119e9d39c55e044e526ae067fa31ce"><code>cc57d87</code></a> feat: update error loc to key in <code>no-dupe-class-members</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20259">#20259</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/126552fcf35da3ddcefa527db06dabc54c04041c"><code>126552f</code></a> feat: update error location in <code>for-direction</code> and <code>no-dupe-args</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20258">#20258</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/167d0970d3802a66910e9820f31dcd717fab0b2a"><code>167d097</code></a> feat: update <code>complexity</code> rule to highlight only static block header (<a href="https://redirect.github.com/eslint/eslint/issues/20245">#20245</a>) (jaymarvelz)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/15f5c7c168d0698683943f51dd617f14a5e6815c"><code>15f5c7c</code></a> fix: forward traversal <code>step.args</code> to visitors (<a href="https://redirect.github.com/eslint/eslint/issues/20253">#20253</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/5a1a534e877f7c4c992885867f923df307c3929d"><code>5a1a534</code></a> fix: allow JSDoc comments in object-shorthand rule (<a href="https://redirect.github.com/eslint/eslint/issues/20167">#20167</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/e86b813eb660f1a5adc8e143a70d9b683cd12362"><code>e86b813</code></a> fix: Use more types from <code>@​eslint/core</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20257">#20257</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/927272d1f0d5683b029b729d368a96527f283323"><code>927272d</code></a> fix: correct <code>Scope</code> typings (<a href="https://redirect.github.com/eslint/eslint/issues/20198">#20198</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/37f76d9c539bb6fc816fedb7be4486b71a58620a"><code>37f76d9</code></a> fix: use <code>AST.Program</code> type for Program node (<a href="https://redirect.github.com/eslint/eslint/issues/20244">#20244</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/ae07f0b3334ebd22ae2e7b09bca5973b96aa9768"><code>ae07f0b</code></a> fix: unify timing report for concurrent linting (<a href="https://redirect.github.com/eslint/eslint/issues/20188">#20188</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/b165d471be6062f4475b972155b02654a974a0e9"><code>b165d47</code></a> fix: correct <code>Rule</code> typings (<a href="https://redirect.github.com/eslint/eslint/issues/20199">#20199</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/fb97cda70d87286a7dbd2457f578ef578d6905e8"><code>fb97cda</code></a> fix: improve error message for missing fix function in suggestions (<a href="https://redirect.github.com/eslint/eslint/issues/20218">#20218</a>) (jaymarvelz)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/d3e81e30ee6be5a21151b7a17ef10a714b6059c0"><code>d3e81e3</code></a> docs: Always recommend to include a files property (<a href="https://redirect.github.com/eslint/eslint/issues/20158">#20158</a>) (Percy Ma)</li>
<li><a href="https://github.com/eslint/eslint/commit/0f0385f1404dcadaba4812120b1ad02334dbd66a"><code>0f0385f</code></a> docs: use consistent naming recommendation (<a href="https://redirect.github.com/eslint/eslint/issues/20250">#20250</a>) (Alex M. Spieslechner)</li>
<li><a href="https://github.com/eslint/eslint/commit/a3b145609ac649fac837c8c0515cbb2a9321ca40"><code>a3b1456</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/cf5f2dd58dd98084a21da04fe7b9054b9478d552"><code>cf5f2dd</code></a> docs: fix correct tag of <code>no-useless-constructor</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20255">#20255</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/10b995c8e5473de8d66d3cd99d816e046f35e3ec"><code>10b995c</code></a> docs: add TS options and examples for <code>nofunc</code> in <code>no-use-before-define</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20249">#20249</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/2584187e4a305ea7a98e1a5bd4dca2a60ad132f8"><code>2584187</code></a> docs: remove repetitive word in comment (<a href="https://redirect.github.com/eslint/eslint/issues/20242">#20242</a>) (reddaisyy)</li>
<li><a href="https://github.com/eslint/eslint/commit/637216bd4f2aae7c928ad04a4e40eecffb50c9e5"><code>637216b</code></a> docs: update CLI flags migration instructions (<a href="https://redirect.github.com/eslint/eslint/issues/20238">#20238</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/e7cda3bdf1bdd664e6033503a3315ad81736b200"><code>e7cda3b</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/7b9446f7cc2054aa2cdf8e6225f4ac15a03671a8"><code>7b9446f</code></a> docs: handle empty flags sections on the feature flags page (<a href="https://redirect.github.com/eslint/eslint/issues/20222">#20222</a>) (sethamus)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/dfe3c1b2034228765c48c8a445554223767dd16d"><code>dfe3c1b</code></a> chore: update <code>@eslint/js</code> version to 9.39.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20270">#20270</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/2375a6de8263393c129d41cac1b407b40111a73c"><code>2375a6d</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/a1f4e52d67c94bef61edd1607dcd130047c1baf0"><code>a1f4e52</code></a> chore: update <code>@eslint</code> dependencies (<a href="https://redirect.github.com/eslint/eslint/issues/20265">#20265</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/c7d32298482752eeac9fb46378d4f1ea095f3836"><code>c7d3229</code></a> chore: update dependency <code>@​eslint/core</code> to ^0.17.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20256">#20256</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/27549bc774c7c2dc5c569070a3e87c62f602bf7d"><code>27549bc</code></a> chore: update fuzz testing to not error if code sample minimizer fails (<a href="https://redirect.github.com/eslint/eslint/issues/20252">#20252</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/a1370ee40e9d8e0e41843f3278cd745fc1ad543f"><code>a1370ee</code></a> ci: bump actions/setup-node from 5 to 6 (<a href="https://redirect.github.com/eslint/eslint/issues/20230">#20230</a>) (dependabot[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/9e7fad4a1867709060686d03e0ec1d0d69671cfb"><code>9e7fad4</code></a> chore: add script to auto-generate eslint:recommended configuration (<a href="https://redirect.github.com/eslint/eslint/issues/20208">#20208</a>) (唯然)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/e2772811a8595d161870835ff04822b25a2cdf45"><code>e277281</code></a> 9.39.1</li>
<li><a href="https://github.com/eslint/eslint/commit/4cdf397b30b2b749865ea0fcf4d30eb8ba458896"><code>4cdf397</code></a> Build: changelog update for 9.39.1</li>
<li><a href="https://github.com/eslint/eslint/commit/92db329211c8da5ce8340a4d4c05ce9c12845381"><code>92db329</code></a> chore: update <code>@eslint/js</code> version to 9.39.1 (<a href="https://redirect.github.com/eslint/eslint/issues/20284">#20284</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/c7ebefc9eaf99b76b30b0d3cf9960807a47367c4"><code>c7ebefc</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/650753ee3976784343ceb40170619dab1aa9fe0d"><code>650753e</code></a> fix: Only pass node to JS lang visitor methods (<a href="https://redirect.github.com/eslint/eslint/issues/20283">#20283</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/51b51f4f1ce82ef63264c4e45d9ef579bcd73f8e"><code>51b51f4</code></a> docs: add a section on when to use extends vs cascading (<a href="https://redirect.github.com/eslint/eslint/issues/20268">#20268</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/61778f6ca33c0f63962a91d6a75a4fa5db9f47d2"><code>61778f6</code></a> chore: update eslint-config-eslint dependency <code>@​eslint/js</code> to ^9.39.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20275">#20275</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/d9ca2fcd9ad63331bfd329a69534e1ff04f231e8"><code>d9ca2fc</code></a> ci: Add rangeStrategy to eslint group in renovate config (<a href="https://redirect.github.com/eslint/eslint/issues/20266">#20266</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/009e5076ff5a4bd845f55e17676e3bb88f47c280"><code>009e507</code></a> test: fix version tests for ESLint v10 (<a href="https://redirect.github.com/eslint/eslint/issues/20274">#20274</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/b44d42699dcd1729b7ecb50ca70e4c1c17f551f1"><code>b44d426</code></a> docs: Update README</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v9.37.0...v9.39.1">compare view</a></li>
</ul>
</details>
<br />

Updates `eslint-config-prettier` from 9.1.2 to 10.1.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/releases">eslint-config-prettier's releases</a>.</em></p>
<blockquote>
<h2>v10.1.8</h2>
<p>republish latest version</p>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.5...v10.1.8">https://github.com/prettier/eslint-config-prettier/compare/v10.1.5...v10.1.8</a></p>
<h2>v10.1.5</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/332">#332</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4"><code>60fef02</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - chore: add <code>funding</code> field into <code>package.json</code></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.4...v10.1.5">https://github.com/prettier/eslint-config-prettier/compare/v10.1.4...v10.1.5</a></p>
<h2>v10.1.4</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/328">#328</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850"><code>94b4799</code></a> Thanks <a href="https://github.com/silvenon"><code>@​silvenon</code></a>! - fix(cli): do not crash on no rules configured</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.3...v10.1.4">https://github.com/prettier/eslint-config-prettier/compare/v10.1.3...v10.1.4</a></p>
<h2>v10.1.3</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/325">#325</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df"><code>4e95a1d</code></a> Thanks <a href="https://github.com/pilikan"><code>@​pilikan</code></a>! - fix: this package is <code>commonjs</code>, align its types correctly</li>
</ul>
<h3>New Contributors</h3>
<ul>
<li><a href="https://github.com/pilikan"><code>@​pilikan</code></a> made their first contribution in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/325">prettier/eslint-config-prettier#325</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.2...v10.1.3">https://github.com/prettier/eslint-config-prettier/compare/v10.1.2...v10.1.3</a></p>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/blob/main/CHANGELOG.md">eslint-config-prettier's changelog</a>.</em></p>
<blockquote>
<h1>eslint-config-prettier</h1>
<h2>10.1.5</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/332">#332</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4"><code>60fef02</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - chore: add <code>funding</code> field into <code>package.json</code></li>
</ul>
<h2>10.1.4</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/328">#328</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850"><code>94b4799</code></a> Thanks <a href="https://github.com/silvenon"><code>@​silvenon</code></a>! - fix(cli): do not crash on no rules configured</li>
</ul>
<h2>10.1.3</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/325">#325</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df"><code>4e95a1d</code></a> Thanks <a href="https://github.com/pilikan"><code>@​pilikan</code></a>! - fix: this package is <code>commonjs</code>, align its types correctly</li>
</ul>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/prettier/eslint-config-prettier/commits/v10.1.8">compare view</a></li>
</ul>
</details>
<br />

Updates `eslint-plugin-playwright` from 2.3.0 to 2.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/playwright-community/eslint-plugin-playwright/releases">eslint-plugin-playwright's releases</a>.</em></p>
<blockquote>
<h2>v2.4.0</h2>
<h1><a href="https://github.com/playwright-community/eslint-plugin-playwright/compare/v2.3.0...v2.4.0">2.4.0</a> (2025-11-30)</h1>
<h3>Bug Fixes</h3>
<ul>
<li><strong>missing-playwright-await:</strong> prevent infinite recursion in checkValidity (<a href="https://github.com/playwright-community/eslint-plugin-playwright/commit/9ce346ddde659050714bbe770363e3cbe1361c9c">9ce346d</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>expect-expect:</strong> Support regex patterns (<a href="https://redirect.github.com/playwright-community/eslint-plugin-playwright/issues/390">#390</a>) (<a href="https://github.com/playwright-community/eslint-plugin-playwright/commit/fdd025339b68173cb5aec57f83c8bc9792388be1">fdd0253</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/ca5601dc112543fff4b66eee542491bb4ddeccf9"><code>ca5601d</code></a> chore: Fix formatting</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/323f954debbded3048c3dbded065639a9a250f32"><code>323f954</code></a> chore: Fix formatting</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/9ce346ddde659050714bbe770363e3cbe1361c9c"><code>9ce346d</code></a> fix(missing-playwright-await): prevent infinite recursion in checkValidity</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/0d3c9573899913830d6095892547bd7ad03bf09e"><code>0d3c957</code></a> chore: Swap utils for dedent</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/b434289fe42836fb66879ec2633813788d9fbd04"><code>b434289</code></a> Fix standalone expect in fixture (<a href="https://redirect.github.com/playwright-community/eslint-plugin-playwright/issues/403">#403</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fdd025339b68173cb5aec57f83c8bc9792388be1"><code>fdd0253</code></a> feat(expect-expect): Support regex patterns (<a href="https://redirect.github.com/playwright-community/eslint-plugin-playwright/issues/390">#390</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/ed7c44d4d9c68cc73b530bfa0beb44a77d9a19e9"><code>ed7c44d</code></a> docs: Use ESLint <code>defineConfig</code> and <code>extends</code> (<a href="https://redirect.github.com/playwright-community/eslint-plugin-playwright/issues/402">#402</a>)</li>
<li>See full diff in <a href="https://github.com/playwright-community/eslint-plugin-playwright/compare/v2.3.0...v2.4.0">compare view</a></li>
</ul>
</details>
<br />

Updates `eslint-plugin-unused-imports` from 4.2.0 to 4.3.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/sweepline/eslint-plugin-unused-imports/commit/6e02be791dbbe3c65476aa354dedb7ebfa94d7a8"><code>6e02be7</code></a> chore: release v4.3.0</li>
<li><a href="https://github.com/sweepline/eslint-plugin-unused-imports/commit/44a71812c1531a0f158c9f310036d5eba63a08ef"><code>44a7181</code></a> fix: preserve imports used in JSDoc <a href="https://github.com/link"><code>@​link</code></a> tags (<a href="https://redirect.github.com/sweepline/eslint-plugin-unused-imports/issues/111">#111</a>)</li>
<li><a href="https://github.com/sweepline/eslint-plugin-unused-imports/commit/275d96e58a67aea38dd23c875e5f67634a4a5653"><code>275d96e</code></a> chore: update deps</li>
<li>See full diff in <a href="https://github.com/sweepline/eslint-plugin-unused-imports/compare/v4.2.0...v4.3.0">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 16.23.1 to 16.26.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>16.26.1</h2>
<p>It fixes numerous false positive bugs, including many in the <code>declaration-property-value-no-unknown</code> rule for the latest CSS specifications.</p>
<ul>
<li>Fixed: <code>*-no-unknown</code> false positives for latest specs by integrating <code>@csstools/css-syntax-patches-for-csstree</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8850">#8850</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>at-rule-no-unknown</code> false positives for <code>@function</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8851">#8851</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>attr()</code>, <code>if()</code> and custom functions (<a href="https://redirect.github.com/stylelint/stylelint/issues/8853">#8853</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>function-url-quotes</code> false positives when URLs require quoting (<a href="https://redirect.github.com/stylelint/stylelint/issues/8804">#8804</a>) (<a href="https://github.com/taearls"><code>@​taearls</code></a>).</li>
<li>Fixed: <code>selector-pseudo-element-no-unknown</code> false positives for <code>::scroll-button()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8856">#8856</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
<h2>16.26.0</h2>
<p>It adds 1 feature and fixes 2 bugs.</p>
<ul>
<li>Added: support for <code>customSyntax</code> with function export (<a href="https://redirect.github.com/stylelint/stylelint/issues/8834">#8834</a>) (<a href="https://github.com/silverwind"><code>@​silverwind</code></a>).</li>
<li>Fixed: <code>custom-property-no-missing-var-function</code> false positives for style query in <code>if()</code> function (<a href="https://redirect.github.com/stylelint/stylelint/issues/8813">#8813</a>) (<a href="https://github.com/sajdakabir"><code>@​sajdakabir</code></a>).</li>
<li>Fixed: <code>media-feature-range-notation</code> false positives for multiple queries and <code>except: exact-value</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8832">#8832</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
</ul>
<h2>16.25.0</h2>
<p>It adds 3 new features, including experimental support for bulk suppressions. It's also our first <a href="https://github.blog/changelog/2025-08-26-releases-now-support-immutability-in-public-preview/">immutable release</a>, with the package published to npm using <a href="https://github.blog/changelog/2025-07-31-npm-trusted-publishing-with-oidc-is-generally-available/">trusted publishing</a> and our dependencies updated on a <a href="https://github.blog/changelog/2025-07-01-dependabot-supports-configuration-of-a-minimum-package-age/">cool down</a> for improved supply chain security.</p>
<ul>
<li>Added: support for bulk suppressions (<a href="https://redirect.github.com/stylelint/stylelint/issues/8564">#8564</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignoreAtRules: []</code> to <code>no-invalid-position-declaration</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8781">#8781</a>) (<a href="https://github.com/jrmlt"><code>@​jrmlt</code></a>).</li>
<li>Added: rule name to custom messages (<a href="https://redirect.github.com/stylelint/stylelint/issues/8774">#8774</a>) (<a href="https://github.com/jhae-de"><code>@​jhae-de</code></a>).</li>
</ul>
<h2>16.24.0</h2>
<p>It adds 1 new rule, adds 1 option to a rule and fixes 2 bugs.</p>
<ul>
<li>Added: <code>rule-nesting-at-rule-required-list</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8680">#8680</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>ignoreAtRules: []</code> to <code>nesting-selector-no-missing-scoping-root</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8743">#8743</a>) (<a href="https://github.com/karlhorky"><code>@​karlhorky</code></a>).</li>
<li>Fixed: <code>function-no-unknown</code> false positives for <code>contrast-color()</code> and <code>sibling-*()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8729">#8729</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for <code>:heading</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8749">#8749</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>16.26.1 - 2025-11-28</h2>
<p>It fixes numerous false positive bugs, including many in the <code>declaration-property-value-no-unknown</code> rule for the latest CSS specifications.</p>
<ul>
<li>Fixed: <code>*-no-unknown</code> false positives for latest specs by integrating <code>@csstools/css-syntax-patches-for-csstree</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8850">#8850</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>at-rule-no-unknown</code> false positives for <code>@function</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8851">#8851</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>attr()</code>, <code>if()</code> and custom functions (<a href="https://redirect.github.com/stylelint/stylelint/pull/8853">#8853</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>function-url-quotes</code> false positives when URLs require quoting (<a href="https://redirect.github.com/stylelint/stylelint/pull/8804">#8804</a>) (<a href="https://github.com/taearls"><code>@​taearls</code></a>).</li>
<li>Fixed: <code>selector-pseudo-element-no-unknown</code> false positives for <code>::scroll-button()</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8856">#8856</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
<h2>16.26.0 - 2025-11-21</h2>
<p>It adds 1 feature and fixes 2 bugs.</p>
<ul>
<li>Added: support for <code>customSyntax</code> with function export (<a href="https://redirect.github.com/stylelint/stylelint/pull/8834">#8834</a>) (<a href="https://github.com/silverwind"><code>@​silverwind</code></a>).</li>
<li>Fixed: <code>custom-property-no-missing-var-function</code> false positives for style query in <code>if()</code> function (<a href="https://redirect.github.com/stylelint/stylelint/pull/8813">#8813</a>) (<a href="https://github.com/sajdakabir"><code>@​sajdakabir</code></a>).</li>
<li>Fixed: <code>media-feature-range-notation</code> false positives for multiple queries and <code>except: exact-value</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8832">#8832</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
</ul>
<h2>16.25.0 - 2025-10-03</h2>
<p>It adds 3 new features, including experimental support for bulk suppressions. It's also our first <a href="https://github.blog/changelog/2025-08-26-releases-now-support-immutability-in-public-preview/">immutable release</a>, with the package published to npm using <a href="https://github.blog/changelog/2025-07-31-npm-trusted-publishing-with-oidc-is-generally-available/">trusted publishing</a> and our dependencies updated on a <a href="https://github.blog/changelog/2025-07-01-dependabot-supports-configuration-of-a-minimum-package-age/">cool down</a> for improved supply chain security.</p>
<ul>
<li>Added: support for bulk suppressions (<a href="https://redirect.github.com/stylelint/stylelint/pull/8564">#8564</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignoreAtRules: []</code> to <code>no-invalid-position-declaration</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8781">#8781</a>) (<a href="https://github.com/jrmlt"><code>@​jrmlt</code></a>).</li>
<li>Added: rule name to custom messages (<a href="https://redirect.github.com/stylelint/stylelint/pull/8774">#8774</a>) (<a href="https://github.com/jhae-de"><code>@​jhae-de</code></a>).</li>
</ul>
<h2>16.24.0 - 2025-09-07</h2>
<p>It adds 1 new rule, adds 1 option to a rule and fixes 2 bugs.</p>
<ul>
<li>Added: <code>rule-nesting-at-rule-required-list</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8680">#8680</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>ignoreAtRules: []</code> to <code>nesting-selector-no-missing-scoping-root</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8743">#8743</a>) (<a href="https://github.com/karlhorky"><code>@​karlhorky</code></a>).</li>
<li>Fixed: <code>function-no-unknown</code> false positives for <code>contrast-color()</code> and <code>sibling-*()</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8729">#8729</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for <code>:heading</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8749">#8749</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/b96814344b7d1088e3459c44dcafebfbdabff412"><code>b968143</code></a> Release 16.26.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8857">#8857</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/2b24b9cd5030b4ef6726d575ea71d34005dd9929"><code>2b24b9c</code></a> Fix <code>selector-pseudo-element-no-unknown</code> false positives for `::scroll-button...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/f152564f037047a4f1a40c812fba77dde05d0062"><code>f152564</code></a> Fix <code>*-no-unknown</code> false positives for latest specs by integrating `@csstools...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/431cb53c0a181eaacc3b208a71c0e765c14faedf"><code>431cb53</code></a> Fix <code>at-rule-no-unknown</code> false positives for <code>@function</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8851">#8851</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/119097ea694cca6bf477ac534fd02c39c8b37c8e"><code>119097e</code></a> Fix <code>declaration-property-value-no-unknown</code> false positives for <code>attr()</code> and ...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/4b9c68be0763a87df187a7fc9de00bced940d916"><code>4b9c68b</code></a> Fix <code>function-url-quotes</code> false positives when URLs require quoting (<a href="https://redirect.github.com/stylelint/stylelint/issues/8804">#8804</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/8cc4ced2e8938785aa29559609984df8c4d83431"><code>8cc4ced</code></a> Bump rollup from 4.52.5 to 4.53.2 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8848">#8848</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/4383feb6dfacb57fc334ab6441ba32e7ea4e3008"><code>4383feb</code></a> Bump file-entry-cache from 11.1.0 to 11.1.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8846">#8846</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/a8a7560c49f78ce1baaa1fd182c03685c12c7b37"><code>a8a7560</code></a> Bump the eslint group with 2 updates (<a href="https://redirect.github.com/stylelint/stylelint/issues/8845">#8845</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/947ad33c1562b03e54b440693db69c5fbb4b39fb"><code>947ad33</code></a> Fix <code>patch-package</code> warning about mismatched <code>@types/css-tree</code> version (<a href="https://redirect.github.com/stylelint/stylelint/issues/8844">#8844</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/16.23.1...16.26.1">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for stylelint since your current version.</p>
</details>
<br />

Updates `stylelint-config-recommended-scss` from 14.1.0 to 16.0.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/releases">stylelint-config-recommended-scss's releases</a>.</em></p>
<blockquote>
<h2>16.0.2</h2>
<ul>
<li>Removed: <code>stylelint</code> less than <code>16.24.0</code> from peer dependencies. Fixes incorrect minimum peer dependency version.</li>
</ul>
<h2>16.0.1</h2>
<ul>
<li>Fixed: false positive for <code>nesting-selector-no-missing-scoping-root</code>.</li>
</ul>
<h2>16.0.0</h2>
<ul>
<li>Changed: updated to <a href="https://github.com/stylelint/stylelint-config-recommended/releases/tag/17.0.0"><code>stylelint-config-recommended@17.0.0</code></a>.</li>
<li>Changed: updated to <a href="https://github.com/stylelint-scss/stylelint-scss/releases/tag/v6.12.1"><code>stylelint-scss@6.12.1</code></a>.</li>
<li>Removed: <code>stylelint</code> less than <code>16.23.1</code> from peer dependencies.</li>
</ul>
<h2>15.0.1</h2>
<ul>
<li>Fixed: change minimum supported Node.js version to <code>20</code>.</li>
</ul>
<h2>15.0.0</h2>
<ul>
<li>Changed: updated to <a href="https://github.com/stylelint/stylelint-config-recommended/releases/tag/16.0.0"><code>stylelint-config-recommended@16.0.0</code></a>.</li>
<li>Changed: updated to <a href="https://github.com/stylelint-scss/stylelint-scss/releases/tag/v6.12.0"><code>stylelint-scss@6.12.0</code></a>.</li>
<li>Removed: <code>stylelint</code> less than <code>16.16.0</code> from peer dependencies.</li>
<li>Removed: Node.js less than <code>22</code> support.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/blob/master/CHANGELOG.md">stylelint-config-recommended-scss's changelog</a>.</em></p>
<blockquote>
<h1>16.0.2</h1>
<ul>
<li>Removed: <code>stylelint</code> less than <code>16.24.0</code> from peer dependencies. Fixes incorrect minimum peer dependency version.</li>
</ul>
<h1>16.0.1</h1>
<ul>
<li>Fixed: false positive for <code>nesting-selector-no-missing-scoping-root</code>.</li>
</ul>
<h1>16.0.0</h1>
<ul>
<li>Changed: updated to <a href="https://github.com/stylelint/stylelint-config-recommended/releases/tag/17.0.0"><code>stylelint-config-recommended@17.0.0</code></a>.</li>
<li>Changed: updated to <a href="https://github.com/stylelint-scss/stylelint-scss/releases/tag/v6.12.1"><code>stylelint-scss@6.12.1</code></a>.</li>
<li>Removed: <code>stylelint</code> less than <code>16.23.1</code> from peer dependencies.</li>
</ul>
<h1>15.0.1</h1>
<ul>
<li>Fixed: change minimum supported Node.js version to <code>20</code>.</li>
</ul>
<h1>15.0.0</h1>
<ul>
<li>Changed: updated to <a href="https://github.com/stylelint/stylelint-config-recommended/releases/tag/16.0.0"><code>stylelint-config-recommended@16.0.0</code></a>.</li>
<li>Changed: updated to <a href="https://github.com/stylelint-scss/stylelint-scss/releases/tag/v6.12.0"><code>stylelint-scss@6.12.0</code></a>.</li>
<li>Removed: <code>stylelint</code> less than <code>16.16.0</code> from peer dependencies.</li>
<li>Removed: Node.js less than <code>22</code> support.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/7f3fe7130faed10a0aa57428520b5d2c0963d273"><code>7f3fe71</code></a> 16.0.2</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/2b8ff0c3b4285038b0ba3f63a666dd3611230acf"><code>2b8ff0c</code></a> Prepare 16.0.2</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/2e8b4b64b7f54ed72b2a9a12635466657153864e"><code>2e8b4b6</code></a> Bump eslint from 9.35.0 to 9.36.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/371">#371</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/94d6fd331be808dbe5de70e70e5b13c2ed39703a"><code>94d6fd3</code></a> Bump stylelint peer dep to 16.24.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/370">#370</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/a9edfb274b12b75e1319388bea3ca73e6df56907"><code>a9edfb2</code></a> 16.0.1</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/02dc994e6ddc8f69e2bf68578b031f58ee1f802b"><code>02dc994</code></a> Prepare 16.0.1</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/82329b174b110bc02ce0cc0fa24dcd5ce9e6ed69"><code>82329b1</code></a> Revert stylelint peer dep version bump (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/368">#368</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/02e99ca3a19a1bf81747b6bdd525273f54a04e74"><code>02e99ca</code></a> Bump eslint from 9.32.0 to 9.35.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/367">#367</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/0c168c85de633d2a84197bf0128ed99a477caa8e"><code>0c168c8</code></a> Bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/363">#363</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/c554a0e637101ed37403499b84b0a0cce3ee7e8c"><code>c554a0e</code></a> Fix false positive for nesting-selector-no-missing-scoping-root (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/365">#365</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/compare/v14.1.0...v16.0.2">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint-scss` from 6.12.1 to 6.13.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/releases">stylelint-scss's releases</a>.</em></p>
<blockquote>
<h2>6.13.0</h2>
<ul>
<li>Added: <code>at-mixin-argumentless-call-parentheses</code> handle mixin calls with content block arguments (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1185">#1185</a>).</li>
<li>Added: <code>at-function-pattern</code>, <code>at-mixin-pattern</code>, <code>dollar-variable-pattern</code>, <code>percent-placeholder-pattern</code> add support for arguments in custom messages (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1187">#1187</a>).</li>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> check for CSS custom properties, add autofix, rule documentation improvements (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1195">#1195</a>).</li>
<li>Fixed: <code>dollar-variable-colon-space-after</code> prevent TypeError for dynamically created nodes (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1159">#1159</a>).</li>
<li>Fixed: <code>load-partial-extension</code> add missing link to docs (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1202">#1202</a>).</li>
<li>Fixed: migrate rules to use autofix callback instead of deprecated <code>context.fix</code> (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1206">#1206</a>).</li>
<li>Updated: <code>stylelint</code> peer dependency version to <code>^16.8.2</code> (required by autofix callback) (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1206">#1206</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.12.1...v6.13.0">https://github.com/stylelint-scss/stylelint-scss/compare/v6.12.1...v6.13.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/blob/master/CHANGELOG.md">stylelint-scss's changelog</a>.</em></p>
<blockquote>
<h1>6.13.0</h1>
<ul>
<li>Added: <code>at-mixin-argumentless-call-parentheses</code> handle mixin calls with content block arguments (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1185">#1185</a>).</li>
<li>Added: <code>at-function-pattern</code>, <code>at-mixin-pattern</code>, <code>dollar-variable-pattern</code>, <code>percent-placeholder-pattern</code> add support for arguments in custom messages (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1187">#1187</a>).</li>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> check for CSS custom properties, add autofix, rule documentation improvements (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1195">#1195</a>).</li>
<li>Fixed: <code>dollar-variable-colon-space-after</code> prevent TypeError for dynamically created nodes (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1159">#1159</a>).</li>
<li>Fixed: <code>load-partial-extension</code> add missing link to docs (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1202">#1202</a>).</li>
<li>Fixed: migrate rules to use autofix callback instead of deprecated <code>context.fix</code> (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1206">#1206</a>).</li>
<li>Updated: <code>stylelint</code> peer dependency version to <code>^16.8.2</code> (required by autofix callback) (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1206">#1206</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.12.1...v6.13.0">https://github.com/stylelint-scss/stylelint-scss/compare/v6.12.1...v6.13.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/86037262a5368edebcf1d06056f1c7f08f6a9460"><code>8603726</code></a> 6.13.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/6ecc2852a8422bbaab0057c7eef215dac7eb0249"><code>6ecc285</code></a> Prepare version 6.13.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1210">#1210</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/9e019659cc46eff0552815674ef27dd0d70f3eff"><code>9e01965</code></a> Update contributors list (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1209">#1209</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/db73b8941d806b5fc89b9ed74a1c5f42fa9c8c5c"><code>db73b89</code></a> Update Stylelint to 16.26.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1208">#1208</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/c1da86f0181ffc1f0fea0a0433bb57447037aa3a"><code>c1da86f</code></a> build(deps-dev): bump prettier from 3.7.3 to 3.7.4 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1207">#1207</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/4de4085e450103c0cce37a89a79d9246d30031e2"><code>4de4085</code></a> Migrate rules to use autofix callback (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1206">#1206</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/2d3022d6272770d1cffac721e5a616b33e1be292"><code>2d3022d</code></a> Bump Node version to 24 in CI (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1205">#1205</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/b9cb7843a56a5cf1c6259d0504a69b6dca10af12"><code>b9cb784</code></a> build(deps): bump mdn-data from 2.24.0 to 2.25.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1204">#1204</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/a8634ada61c3a83bc99b5bf6f2ea1269246f905d"><code>a8634ad</code></a> build(deps-dev): bump js-yaml from 3.14.1 to 3.14.2 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1203">#1203</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-scss/commit/3784a783ea78fcbee69c3d5378de3acd0ce91c31"><code>3784a78</code></a> build(deps): bump postcss-selector-parser from 7.1.0 to 7.1.1 (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1197">#1197</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.12.1...v6.13.0">compare view</a></li>
</ul>
</details>
<br />


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
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.

---

## Discussion

### Comment by @dependabot on January 20, 2026 at 06:44 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1448*
