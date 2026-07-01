---
type: pull_request
number: 1476
title: "chore(deps): bump the lint group across 1 directory with 8 updates"
state: closed
author: dependabot
created: 2026-01-20T18:44:55Z
updated: 2026-02-16T08:22:00Z
closed: 2026-02-16T04:02:19Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-0b1c443f14
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1476
---

# Pull Request #1476: chore(deps): bump the lint group across 1 directory with 8 updates

**Author**: @dependabot
**Created**: January 20, 2026 at 06:44 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-0b1c443f14`

## Description

Bumps the lint group with 7 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [eslint](https://github.com/eslint/eslint) | `9.37.0` | `9.39.2` |
| [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier) | `9.1.2` | `10.1.8` |
| [eslint-plugin-playwright](https://github.com/mskelton/eslint-plugin-playwright) | `2.3.0` | `2.5.0` |
| [eslint-plugin-prettier](https://github.com/prettier/eslint-plugin-prettier) | `5.5.4` | `5.5.5` |
| [eslint-plugin-unused-imports](https://github.com/sweepline/eslint-plugin-unused-imports) | `4.2.0` | `4.3.0` |
| [stylelint](https://github.com/stylelint/stylelint) | `16.23.1` | `17.0.0` |
| [stylelint-config-recommended-scss](https://github.com/stylelint-scss/stylelint-config-recommended-scss) | `14.1.0` | `17.0.0` |


Updates `eslint` from 9.37.0 to 9.39.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v9.39.2</h2>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/57058331946568164449c5caabe2cf206e4fb5d9"><code>5705833</code></a> fix: warn when <code>eslint-env</code> configuration comments are found (<a href="https://redirect.github.com/eslint/eslint/issues/20381">#20381</a>) (sethamus)</li>
</ul>
<h2>Build Related</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/506f1549a64aa65bdddc75c71cb62f0ab94b5a23"><code>506f154</code></a> build: add .scss files entry to knip (<a href="https://redirect.github.com/eslint/eslint/issues/20391">#20391</a>) (Milos Djermanovic)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/7ca0af7f9f89dd4a01736dae01931c45d528171b"><code>7ca0af7</code></a> chore: upgrade to <code>@eslint/js@9.39.2</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20394">#20394</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/c43ce24ff0ce073ec4ad691cd5a50171dfe6cf1e"><code>c43ce24</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/4c9858e47bb9146cf20f546a562bc58a9ee3dae1"><code>4c9858e</code></a> ci: add <code>v9.x-dev</code> branch (<a href="https://redirect.github.com/eslint/eslint/issues/20382">#20382</a>) (Milos Djermanovic)</li>
</ul>
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
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/9278324aa0023d223874825b0d4b6ac75783096a"><code>9278324</code></a> 9.39.2</li>
<li><a href="https://github.com/eslint/eslint/commit/542266ad3c58b47066d4b8ae61d419b423acee8f"><code>542266a</code></a> Build: changelog update for 9.39.2</li>
<li><a href="https://github.com/eslint/eslint/commit/7ca0af7f9f89dd4a01736dae01931c45d528171b"><code>7ca0af7</code></a> chore: upgrade to <code>@eslint/js@9.39.2</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20394">#20394</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/c43ce24ff0ce073ec4ad691cd5a50171dfe6cf1e"><code>c43ce24</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/57058331946568164449c5caabe2cf206e4fb5d9"><code>5705833</code></a> fix: warn when <code>eslint-env</code> configuration comments are found (<a href="https://redirect.github.com/eslint/eslint/issues/20381">#20381</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/506f1549a64aa65bdddc75c71cb62f0ab94b5a23"><code>506f154</code></a> build: add .scss files entry to knip (<a href="https://redirect.github.com/eslint/eslint/issues/20391">#20391</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/4c9858e47bb9146cf20f546a562bc58a9ee3dae1"><code>4c9858e</code></a> ci: add <code>v9.x-dev</code> branch (<a href="https://redirect.github.com/eslint/eslint/issues/20382">#20382</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/e2772811a8595d161870835ff04822b25a2cdf45"><code>e277281</code></a> 9.39.1</li>
<li><a href="https://github.com/eslint/eslint/commit/4cdf397b30b2b749865ea0fcf4d30eb8ba458896"><code>4cdf397</code></a> Build: changelog update for 9.39.1</li>
<li><a href="https://github.com/eslint/eslint/commit/92db329211c8da5ce8340a4d4c05ce9c12845381"><code>92db329</code></a> chore: update <code>@eslint/js</code> version to 9.39.1 (<a href="https://redirect.github.com/eslint/eslint/issues/20284">#20284</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v9.37.0...v9.39.2">compare view</a></li>
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

Updates `eslint-plugin-playwright` from 2.3.0 to 2.5.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/mskelton/eslint-plugin-playwright/releases">eslint-plugin-playwright's releases</a>.</em></p>
<blockquote>
<h2>v2.5.0</h2>
<h1><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.4.1...v2.5.0">2.5.0</a> (2026-01-12)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>Fix lint (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/4e8461dc5b73018d706782c729fbcce67347b7f6">4e8461d</a>)</li>
<li>Fix TypeScript types (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/af7d8702121f58eb5974b81a514470542162823c">af7d870</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>Add <code>enforce-consistent-spacing-between-blocks</code> rule (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/411">#411</a>) (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/a9b78d5d0c7a7c051d9bee85a584ca483dd22777">a9b78d5</a>)</li>
<li>Add no-restricted-locators rule (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/a65200b1773b49ccafbd9a9b8a81e4e9f700bd67">a65200b</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/407">#407</a></li>
<li><strong>prefer-web-first-assertions:</strong> Support <code>allInnerTexts()</code> and <code>allTextContents()</code> (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/36917a86fb7e4ef49837e7657a8363d55f06e461">36917a8</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/362">#362</a></li>
</ul>
<h2>v2.4.1</h2>
<h2><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.4.0...v2.4.1">2.4.1</a> (2026-01-09)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>no-conditional-in-test:</strong> allow nullish coalescing operator (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/2c25b4fe3b7d487a8cc06c43a1bf91fea3f3c7ea">2c25b4f</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/406">#406</a></li>
</ul>
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
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/36917a86fb7e4ef49837e7657a8363d55f06e461"><code>36917a8</code></a> feat(prefer-web-first-assertions): Support <code>allInnerTexts()</code> and `allTextCont...</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/5f4d2f290db48758ff9d285d20dedd3a498ec9c3"><code>5f4d2f2</code></a> chore: Format with oxfmt</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/9b6671e2ac7b8ab2b2f93d0305d7583c86b1c078"><code>9b6671e</code></a> test(consistent-spacing-between-blocks): Add tests for same line</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/461b7d04a71f4772d3d9202b2300705947985907"><code>461b7d0</code></a> chore: Update scheduled release</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/a9b78d5d0c7a7c051d9bee85a584ca483dd22777"><code>a9b78d5</code></a> feat: Add <code>enforce-consistent-spacing-between-blocks</code> rule (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/411">#411</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/4e8461dc5b73018d706782c729fbcce67347b7f6"><code>4e8461d</code></a> fix: Fix lint</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/af7d8702121f58eb5974b81a514470542162823c"><code>af7d870</code></a> fix: Fix TypeScript types</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/66f13f44dea553594569497e1085eb4178af348a"><code>66f13f4</code></a> Fix typo</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/a15487719d49593741d554a53ef6d005a5eafee5"><code>a154877</code></a> chore: Update URLs</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/a65200b1773b49ccafbd9a9b8a81e4e9f700bd67"><code>a65200b</code></a> feat: Add no-restricted-locators rule</li>
<li>Additional commits viewable in <a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.3.0...v2.5.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for eslint-plugin-playwright since your current version.</p>
</details>
<br />

Updates `eslint-plugin-prettier` from 5.5.4 to 5.5.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-plugin-prettier/releases">eslint-plugin-prettier's releases</a>.</em></p>
<blockquote>
<h2>v5.5.5</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-plugin-prettier/pull/772">#772</a> <a href="https://github.com/prettier/eslint-plugin-prettier/commit/7264ed0a6cf47fc36befed32f459e7d875f5992c"><code>7264ed0</code></a> Thanks <a href="https://github.com/BPScott"><code>@​BPScott</code></a>! - Bump prettier-linter-helpers dependency to v1.0.1</p>
</li>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-plugin-prettier/pull/776">#776</a> <a href="https://github.com/prettier/eslint-plugin-prettier/commit/77651a33cd16fd4c50b7346971990b900a42408b"><code>77651a3</code></a> Thanks <a href="https://github.com/aswils"><code>@​aswils</code></a>! - fix: bump synckit for yarn PnP ESM issue</p>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-plugin-prettier/blob/main/CHANGELOG.md">eslint-plugin-prettier's changelog</a>.</em></p>
<blockquote>
<h2>5.5.5</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-plugin-prettier/pull/772">#772</a> <a href="https://github.com/prettier/eslint-plugin-prettier/commit/7264ed0a6cf47fc36befed32f459e7d875f5992c"><code>7264ed0</code></a> Thanks <a href="https://github.com/BPScott"><code>@​BPScott</code></a>! - Bump prettier-linter-helpers dependency to v1.0.1</p>
</li>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-plugin-prettier/pull/776">#776</a> <a href="https://github.com/prettier/eslint-plugin-prettier/commit/77651a33cd16fd4c50b7346971990b900a42408b"><code>77651a3</code></a> Thanks <a href="https://github.com/aswils"><code>@​aswils</code></a>! - fix: bump synckit for yarn PnP ESM issue</p>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prettier/eslint-plugin-prettier/commit/e2c154a7214d4548dad225a56ee1e333d6389b66"><code>e2c154a</code></a> chore: release eslint-plugin-prettier (<a href="https://redirect.github.com/prettier/eslint-plugin-prettier/issues/773">#773</a>)</li>
<li><a href="https://github.com/prettier/eslint-plugin-prettier/commit/6795c1abf6dc9949da8681b05ec31d323794d00c"><code>6795c1a</code></a> build(deps): Bump the actions group across 1 directory with 2 updates (<a href="https://redirect.github.com/prettier/eslint-plugin-prettier/issues/774">#774</a>)</li>
<li><a href="https://github.com/prettier/eslint-plugin-prettier/commit/77651a33cd16fd4c50b7346971990b900a42408b"><code>77651a3</code></a> fix: bump synckit for yarn PnP ESM issue (<a href="https://redirect.github.com/prettier/eslint-plugin-prettier/issues/776">#776</a>)</li>
<li><a href="https://github.com/prettier/eslint-plugin-prettier/commit/7264ed0a6cf47fc36befed32f459e7d875f5992c"><code>7264ed0</code></a> chore: bump prettier-linter-helpers to v1.0.1 (<a href="https://redirect.github.com/prettier/eslint-plugin-prettier/issues/772">#772</a>)</li>
<li><a href="https://github.com/prettier/eslint-plugin-prettier/commit/e11a5b7e71f41b3238da944ba1ee84f7f518a4f4"><code>e11a5b7</code></a> build(deps): Bump the actions group across 1 directory with 3 updates (<a href="https://redirect.github.com/prettier/eslint-plugin-prettier/issues/769">#769</a>)</li>
<li><a href="https://github.com/prettier/eslint-plugin-prettier/commit/befda88381335cd5491d2aaa16b67350ba3cc602"><code>befda88</code></a> ci: enable trusted publishing (<a href="https://redirect.github.com/prettier/eslint-plugin-prettier/issues/757">#757</a>)</li>
<li>See full diff in <a href="https://github.com/prettier/eslint-plugin-prettier/compare/v5.5.4...v5.5.5">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for eslint-plugin-prettier since your current version.</p>
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

Updates `stylelint` from 16.23.1 to 17.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.0.0</h2>
<p>It contains 14 breaking changes, which we've detailed in the <a href="https://github.com/stylelint/stylelint/blob/main/docs/migration-guide/to-17.md">migrating to <code>17.0.0</code> guide</a>. Additionally, it adds 3 options to the rules and fixes 9 bugs. We've also released compatible versions of our <a href="https://www.npmjs.com/package/stylelint-config-standard">shared config</a>, <a href="https://marketplace.visualstudio.com/items?itemName=stylelint.vscode-stylelint">Visual Studio Code extension</a>, <a href="https://www.npmjs.com/package/stylelint-test-rule-node">Node.js Rule Tester</a> and <a href="https://www.npmjs.com/package/jest-preset-stylelint">Jest preset</a>.</p>
<ul>
<li>Removed: CommonJS Node.js API (<a href="https://redirect.github.com/stylelint/stylelint/issues/8859">#8859</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: <code>output</code> property in the Node.js API returned resolved object (<a href="https://redirect.github.com/stylelint/stylelint/issues/8878">#8878</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: support for Node.js less than 20.19.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8867">#8867</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: GitHub formatter (<a href="https://redirect.github.com/stylelint/stylelint/issues/8888">#8888</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: <code>resolveNestedSelectors</code> option from <code>selector-class-pattern</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8931">#8931</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: <code>checkContextFunctionalPseudoClasses</code> option from <code>selector-max-id</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: default <code>fix</code> mode to <code>strict</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8889">#8889</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>report</code> to be consistent and predictable in how it handles the provided position arguments (<a href="https://redirect.github.com/stylelint/stylelint/issues/8217">#8217</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Changed: <code>selector-max-*</code> syntax rules for standard CSS nesting and modern functional pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>*-specificity</code> semantic rules for standard CSS nesting (<a href="https://redirect.github.com/stylelint/stylelint/issues/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>no-duplicate-selectors</code> and <code>selector-no-qualifying-type</code> for standard CSS nesting (<a href="https://redirect.github.com/stylelint/stylelint/issues/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>*-list</code> rules to have consistent behaviour for vendor prefixes and case (<a href="https://redirect.github.com/stylelint/stylelint/issues/8912">#8912</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>*-no-vendor-prefix</code> rules to have consistent behaviour for their <code>ignore*: []</code> secondary options (<a href="https://redirect.github.com/stylelint/stylelint/issues/8924">#8924</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>declaration-property-max-values</code> rule to have consistent behaviour for vendor prefixes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8926">#8926</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>except: [&quot;after-block&quot;]</code> to <code>custom-property-empty-line-before</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8921">#8921</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Added: <code>except: [&quot;after-block&quot;]</code> to <code>declaration-empty-line-before</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8910">#8910</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Added: <code>ignoreSelectors: []</code> to <code>no-duplicate-selectors</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8883">#8883</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: Windows drive letter casing inconsistencies when matching patterns against file paths (<a href="https://redirect.github.com/stylelint/stylelint/issues/8941">#8941</a>) (<a href="https://github.com/adalinesimonian"><code>@​adalinesimonian</code></a>).</li>
<li>Fixed: CLI help to include TypeScript config files (<a href="https://redirect.github.com/stylelint/stylelint/issues/8908">#8908</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>at-rule-descriptor-no-unknown</code> false positives for declarations within feature-value-blocks (<a href="https://redirect.github.com/stylelint/stylelint/issues/8868">#8868</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>declaration-block-no-redundant-longhand-properties</code> false negatives for short and long combinations (<a href="https://redirect.github.com/stylelint/stylelint/issues/8892">#8892</a>) (<a href="https://github.com/nathannewyen"><code>@​nathannewyen</code></a>).</li>
<li>Fixed: <code>media-feature-name-no-unknown</code> false positives for namespaced dollar variables and range context queries (<a href="https://redirect.github.com/stylelint/stylelint/issues/8890">#8890</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>nesting-selector-no-missing-scoping-root</code> false positives for CSS-in-JS (<a href="https://redirect.github.com/stylelint/stylelint/issues/8905">#8905</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>no-invalid-position-declaration</code> false negatives for embedded blocks (<a href="https://redirect.github.com/stylelint/stylelint/issues/8907">#8907</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>selector-no-qualifying-type</code> false negatives for <code>:is/where()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8940">#8940</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>selector-type-no-unknown</code> false positives for MathML 4 tags (<a href="https://redirect.github.com/stylelint/stylelint/issues/8874">#8874</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
</ul>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.0.0 - 2026-01-15</h2>
<p>It contains 14 breaking changes, which we've detailed in the <a href="https://github.com/stylelint/stylelint/blob/main/docs/migration-guide/to-17.md">migrating to <code>17.0.0</code> guide</a>. Additionally, it adds 3 options to the rules and fixes 9 bugs. We've also released compatible versions of our <a href="https://www.npmjs.com/package/stylelint-config-standard">shared config</a>, <a href="https://marketplace.visualstudio.com/items?itemName=stylelint.vscode-stylelint">Visual Studio Code extension</a>, <a href="https://www.npmjs.com/package/stylelint-test-rule-node">Node.js Rule Tester</a> and <a href="https://www.npmjs.com/package/jest-preset-stylelint">Jest preset</a>.</p>
<ul>
<li>Removed: CommonJS Node.js API (<a href="https://redirect.github.com/stylelint/stylelint/pull/8859">#8859</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: <code>output</code> property in the Node.js API returned resolved object (<a href="https://redirect.github.com/stylelint/stylelint/pull/8878">#8878</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: support for Node.js less than 20.19.0 (<a href="https://redirect.github.com/stylelint/stylelint/pull/8867">#8867</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: GitHub formatter (<a href="https://redirect.github.com/stylelint/stylelint/pull/8888">#8888</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: <code>resolveNestedSelectors</code> option from <code>selector-class-pattern</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8931">#8931</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Removed: <code>checkContextFunctionalPseudoClasses</code> option from <code>selector-max-id</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: default <code>fix</code> mode to <code>strict</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8889">#8889</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>report</code> to be consistent and predictable in how it handles the provided position arguments (<a href="https://redirect.github.com/stylelint/stylelint/pull/8217">#8217</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Changed: <code>selector-max-*</code> syntax rules for standard CSS nesting and modern functional pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/pull/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>*-specificity</code> semantic rules for standard CSS nesting (<a href="https://redirect.github.com/stylelint/stylelint/pull/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>no-duplicate-selectors</code> and <code>selector-no-qualifying-type</code> for standard CSS nesting (<a href="https://redirect.github.com/stylelint/stylelint/pull/8913">#8913</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>*-list</code> rules to have consistent behaviour for vendor prefixes and case (<a href="https://redirect.github.com/stylelint/stylelint/pull/8912">#8912</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>*-no-vendor-prefix</code> rules to have consistent behaviour for their <code>ignore*: []</code> secondary options (<a href="https://redirect.github.com/stylelint/stylelint/pull/8924">#8924</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Changed: <code>declaration-property-max-values</code> rule to have consistent behaviour for vendor prefixes (<a href="https://redirect.github.com/stylelint/stylelint/pull/8926">#8926</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>except: [&quot;after-block&quot;]</code> to <code>custom-property-empty-line-before</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8921">#8921</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Added: <code>except: [&quot;after-block&quot;]</code> to <code>declaration-empty-line-before</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8910">#8910</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Added: <code>ignoreSelectors: []</code> to <code>no-duplicate-selectors</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8883">#8883</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: Windows drive letter casing inconsistencies when matching patterns against file paths (<a href="https://redirect.github.com/stylelint/stylelint/pull/8941">#8941</a>) (<a href="https://github.com/adalinesimonian"><code>@​adalinesimonian</code></a>).</li>
<li>Fixed: CLI help to include TypeScript config files (<a href="https://redirect.github.com/stylelint/stylelint/pull/8908">#8908</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>at-rule-descriptor-no-unknown</code> false positives for declarations within feature-value-blocks (<a href="https://redirect.github.com/stylelint/stylelint/pull/8868">#8868</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>declaration-block-no-redundant-longhand-properties</code> false negatives for short and long combinations (<a href="https://redirect.github.com/stylelint/stylelint/pull/8892">#8892</a>) (<a href="https://github.com/nathannewyen"><code>@​nathannewyen</code></a>).</li>
<li>Fixed: <code>media-feature-name-no-unknown</code> false positives for namespaced dollar variables and range context queries (<a href="https://redirect.github.com/stylelint/stylelint/pull/8890">#8890</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>nesting-selector-no-missing-scoping-root</code> false positives for CSS-in-JS (<a href="https://redirect.github.com/stylelint/stylelint/pull/8905">#8905</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>no-invalid-position-declaration</code> false negatives for embedded blocks (<a href="https://redirect.github.com/stylelint/stylelint/pull/8907">#8907</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>selector-no-qualifying-type</code> false negatives for <code>:is/where()</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8940">#8940</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>selector-type-no-unknown</code> false positives for MathML 4 tags (<a href="https://redirect.github.com/stylelint/stylelint/pull/8874">#8874</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
</ul>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/a3e96bd1bd679fc0c70a87955a9e7f633c40b0ca"><code>a3e96bd</code></a> Release 17.0.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8969">#8969</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/4587f265ae5bdf04b702a38c42e92bdd3b73ea98"><code>4587f26</code></a> Revert &quot;Add v17 branch to CI workflow&quot; (<a href="https://redirect.github.com/stylelint/stylelint/issues/8952">#8952</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/0eb7972faeac3d87e4cb4ca511a693e374c31d3c"><code>0eb7972</code></a> Revert &quot;Change <code>target-branch</code> to <code>v17</code> in Dependabot config&quot; (<a href="https://redirect.github.com/stylelint/stylelint/issues/8951">#8951</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/44362b1ecfad69f3a3b2a27f7715ce432dd19801"><code>44362b1</code></a> Merge pull request <a href="https://redirect.github.com/stylelint/stylelint/issues/8870">#8870</a> from stylelint/v17</li>
<li><a href="https://github.com/stylelint/stylelint/commit/71d330357f72c7c828e26caebb4fca49b1356f63"><code>71d3303</code></a> Revert &quot;chore: bump version to 17.0.0&quot; (<a href="https://redirect.github.com/stylelint/stylelint/issues/8917">#8917</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/54716094881bd1b686ca238570599a9f50f3ecd2"><code>5471609</code></a> Use special <code>'module.exports'</code> export (<a href="https://redirect.github.com/stylelint/stylelint/issues/8967">#8967</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/022924e467d711b6d245786ed14032796dd0138b"><code>022924e</code></a> Update <code>@csstools/*</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8968">#8968</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/62fb7c886a639911d521e55cdc2a57b9dbc358fd"><code>62fb7c8</code></a> Document extension in bug template (<a href="https://redirect.github.com/stylelint/stylelint/issues/8965">#8965</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/252c1485c334c8f7ae83c71c5701c2061d695a82"><code>252c148</code></a> Bump the eslint group with 2 updates (<a href="https://redirect.github.com/stylelint/stylelint/issues/8963">#8963</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/2d165c6e4df72d5e4da5567dad520b7313e04dc4"><code>2d165c6</code></a> Bump stylelint/.github/.github/workflows/call-release.yml from 0.5.0 to 0.5.1...</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/16.23.1...17.0.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for stylelint since your current version.</p>
</details>
<br />

Updates `stylelint-config-recommended-scss` from 14.1.0 to 17.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/releases">stylelint-config-recommended-scss's releases</a>.</em></p>
<blockquote>
<h2>17.0.0</h2>
<ul>
<li>Removed: <code>stylelint</code> less than <code>17.0.0</code> from peer dependencies.</li>
<li>Changed: updated to <a href="https://github.com/stylelint/stylelint-config-recommended/releases/tag/18.0.0"><code>stylelint-config-recommended@18.0.0</code></a>.</li>
<li>Changed: updated to <a href="https://github.com/stylelint-scss/stylelint-scss/releases/tag/v7.0.0"><code>stylelint-scss@7.0.0</code></a>.</li>
<li>Changed: module type to ESM.</li>
<li>Fixed: disabled <code>no-descending-specificity</code> rule because of false positives due to standard nesting.</li>
<li>Fixed: disabled <code>no-duplicate-selectors</code> rule because of false positives due to standard nesting.</li>
</ul>
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
<h1>17.0.0</h1>
<ul>
<li>Removed: <code>stylelint</code> less than <code>17.0.0</code> from peer dependencies.</li>
<li>Changed: updated to <a href="https://github.com/stylelint/stylelint-config-recommended/releases/tag/18.0.0"><code>stylelint-config-recommended@18.0.0</code></a>.</li>
<li>Changed: updated to <a href="https://github.com/stylelint-scss/stylelint-scss/releases/tag/v7.0.0"><code>stylelint-scss@7.0.0</code></a>.</li>
<li>Changed: module type to ESM.</li>
<li>Fixed: disabled <code>no-descending-specificity</code> rule because of false positives due to standard nesting.</li>
<li>Fixed: disabled <code>no-duplicate-selectors</code> rule because of false positives due to standard nesting.</li>
</ul>
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
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/15eeb90c7ed77da7668a1b30a5bac6d021deb494"><code>15eeb90</code></a> 17.0.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/47b4ad379b95f14d909dd79c9cf406fa035f3eb6"><code>47b4ad3</code></a> Prepare 17.0.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/383">#383</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/7f3fe7130faed10a0aa57428520b5d2c0963d273"><code>7f3fe71</code></a> 16.0.2</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/2b8ff0c3b4285038b0ba3f63a666dd3611230acf"><code>2b8ff0c</code></a> Prepare 16.0.2</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/2e8b4b64b7f54ed72b2a9a12635466657153864e"><code>2e8b4b6</code></a> Bump eslint from 9.35.0 to 9.36.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/371">#371</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/94d6fd331be808dbe5de70e70e5b13c2ed39703a"><code>94d6fd3</code></a> Bump stylelint peer dep to 16.24.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/370">#370</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/a9edfb274b12b75e1319388bea3ca73e6df56907"><code>a9edfb2</code></a> 16.0.1</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/02dc994e6ddc8f69e2bf68578b031f58ee1f802b"><code>02dc994</code></a> Prepare 16.0.1</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/82329b174b110bc02ce0cc0fa24dcd5ce9e6ed69"><code>82329b1</code></a> Revert stylelint peer dep version bump (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/368">#368</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/02e99ca3a19a1bf81747b6bdd525273f54a04e74"><code>02e99ca</code></a> Bump eslint from 9.32.0 to 9.35.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/367">#367</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/compare/v14.1.0...v17.0.0">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint-scss` from 6.12.1 to 6.14.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-scss/releases">stylelint-scss's releases</a>.</em></p>
<blockquote>
<h2>6.14.0</h2>
<ul>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> report namespaced variables in custom properties (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1225">#1225</a>).</li>
<li>Added: <code>function-disallowed-list</code> check disallowed functions inside <code>@return</code> expressions (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1214">#1214</a>).</li>
<li>Fixed: <code>dollar-variable-no-missing-interpolation</code> fix false positive for variables already inside interpolation (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1213">#1213</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.13.0...v6.14.0">https://github.com/stylelint-scss/stylelint-scss/compare/v6.13.0...v6.14.0</a></p>
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
<h1>6.14.0</h1>
<ul>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> report namespaced variables in custom properties (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1225">#1225</a>).</li>
<li>Added: <code>function-disallowed-list</code> check disallowed functions inside <code>@return</code> expressions (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1214">#1214</a>).</li>
<li>Fixed: <code>dollar-variable-no-missing-interpolation</code> fix false positive for variables already inside interpolation (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1213">#1213</a>).</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stylelint-scss/stylelint-scss/compare/v6.13.0...v6.14.0">https://github.com/stylelint-scss/stylelint-scss/compare/v6.13.0...v6.14.0</a></p>
<h1>6.13.0</h1>
<ul>
<li>Added: <code>at-mixin-argumentless-call-parentheses</code> handle mixin calls with content block arguments (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1185">#1185</a>).</li>
<li>Added: <code>at-function-pattern</code>, <code>at-mixin-pattern</code>, <code>dollar-variable-pattern</code>, <code>percent-placeholder-pattern</code> add support for arguments in custom messages (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1187">#1187</a>).</li>
<li>Added: <code>dollar-variable-no-missing-interpolation</code> check for CSS custom properties, add autofix, rule documentation improvements (<a href="https://redirect.github.com/stylelint-scss/stylelint-scss/issues/1195">#1195</a>).</li>
<li>Fixed: <code>dollar-variable-colon-space-after</code> prevent TypeError for dynamically created nodes (<a href="https://redirect.github.com/stylelint-sc...

_Description has been truncated_

---

## Discussion

### Comment by @swadeley on February 10, 2026 at 10:52 AM UTC

@dependabot rebase

### Comment by @dependabot on February 10, 2026 at 10:53 AM UTC

Looks like this PR has been edited by someone other than Dependabot. That means Dependabot can't rebase it - sorry!

If you're happy for Dependabot to recreate it from scratch, overwriting any edits, you can request `@dependabot recreate`.


### Comment by @swadeley on February 10, 2026 at 10:55 AM UTC

@dependabot recreate

### Comment by @swadeley on February 13, 2026 at 06:26 AM UTC

@dependabot recreate

### Comment by @swadeley on February 13, 2026 at 10:45 AM UTC

npm error Found: eslint@10.0.0

npm error Could not resolve dependency:
npm error peer eslint@"^7.5.0 || ^8.0.0 || ^9.0.0" from @babel/eslint-parser@7.28.0

npm error Conflicting peer dependency: eslint@9.39.2

### Comment by @swadeley on February 16, 2026 at 03:24 AM UTC

@dependabot recreate

### Comment by @dependabot on February 16, 2026 at 04:02 AM UTC

This pull request was built based on a group rule. Closing it will not ignore any of these versions in future pull requests.

To ignore these dependencies, configure [ignore rules](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#ignore) in dependabot.yml

---

## Reviews

### Review by @swadeley - Approved on February 10, 2026 at 10:40 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1476*
