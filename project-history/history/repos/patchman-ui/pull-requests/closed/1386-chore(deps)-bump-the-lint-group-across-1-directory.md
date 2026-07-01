---
type: pull_request
number: 1386
title: "chore(deps): bump the lint group across 1 directory with 5 updates"
state: closed
author: dependabot
created: 2025-09-22T18:16:50Z
updated: 2025-11-10T18:11:18Z
closed: 2025-11-10T18:11:16Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-777ea333af
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1386
---

# Pull Request #1386: chore(deps): bump the lint group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: September 22, 2025 at 06:16 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-777ea333af`

## Description

Bumps the lint group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [eslint](https://github.com/eslint/eslint) | `7.32.0` | `9.36.0` |
| [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier) | `7.2.0` | `10.1.8` |
| [eslint-plugin-cypress](https://github.com/cypress-io/eslint-plugin-cypress) | `2.15.2` | `5.1.1` |
| [stylelint](https://github.com/stylelint/stylelint) | `16.23.1` | `16.24.0` |
| [stylelint-config-recommended-scss](https://github.com/stylelint-scss/stylelint-config-recommended-scss) | `14.1.0` | `16.0.1` |


Updates `eslint` from 7.32.0 to 9.36.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v9.36.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/47afcf668df65eac68d7b04145d037037010a076"><code>47afcf6</code></a> feat: correct <code>preserve-caught-error</code> edge cases (<a href="https://redirect.github.com/eslint/eslint/issues/20109">#20109</a>) (Francesco Trotta)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/75b74d865d3b8e7fa3bcf5ad29f4bf6d18d1310e"><code>75b74d8</code></a> fix: add missing rule option types (<a href="https://redirect.github.com/eslint/eslint/issues/20127">#20127</a>) (ntnyq)</li>
<li><a href="https://github.com/eslint/eslint/commit/1c0d85049e3f30a8809340c1abc881c63b7812ff"><code>1c0d850</code></a> fix: update <code>eslint-all.js</code> to use <code>Object.freeze</code> for <code>rules</code> object (<a href="https://redirect.github.com/eslint/eslint/issues/20116">#20116</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/7d61b7fadc9c5c6f2b131e37e8a3cffa5aae8ee6"><code>7d61b7f</code></a> fix: add missing scope types to <code>Scope.type</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20110">#20110</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/7a670c301b58609017ce8cfda99ee81f95de3898"><code>7a670c3</code></a> fix: correct rule option typings in <code>rules.d.ts</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20084">#20084</a>) (Pixel998)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/b73ab12acd3e87f8d8173cda03499f6cd1f26db6"><code>b73ab12</code></a> docs: update examples to use <code>defineConfig</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20131">#20131</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/31d93926990fba536846ec727d7a2625fc844649"><code>31d9392</code></a> docs: fix typos (<a href="https://redirect.github.com/eslint/eslint/issues/20118">#20118</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/c7f861b3f8c1ac961b4cd4f22483798f3324c62b"><code>c7f861b</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/6b0c08b106aa66f2e9fa484282f0eb63c64a1215"><code>6b0c08b</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/91f97c50468fbdc089c91e99c2ea0fe821911df2"><code>91f97c5</code></a> docs: Update README (GitHub Actions Bot)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/12411e8d450ed26a5f7cca6a78ec05323c9323e8"><code>12411e8</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.36.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20139">#20139</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/488cba6b391b97b2cfc74bbb46fdeacb1361949e"><code>488cba6</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/bac82a2a9c80a3f69087852758d7737aea371f09"><code>bac82a2</code></a> ci: simplify renovate configuration (<a href="https://redirect.github.com/eslint/eslint/issues/19907">#19907</a>) (唯然)</li>
<li><a href="https://github.com/eslint/eslint/commit/c00bb37d62c1bcc0a37f094371be9c40064009f1"><code>c00bb37</code></a> ci: bump actions/labeler from 5 to 6 (<a href="https://redirect.github.com/eslint/eslint/issues/20090">#20090</a>) (dependabot[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/fee751dc8aeab54547af4538332ea5c069ef28b6"><code>fee751d</code></a> refactor: use <code>defaultOptions</code> in rules (<a href="https://redirect.github.com/eslint/eslint/issues/20121">#20121</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/1ace67d9f7903adc3d3f09868aa05b673e7d3f3b"><code>1ace67d</code></a> chore: update example to use <code>defineConfig</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20111">#20111</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/4821963bf765532069c49e9da9ecbe9485b073fc"><code>4821963</code></a> test: add missing loc information to error objects in rule tests (<a href="https://redirect.github.com/eslint/eslint/issues/20112">#20112</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/b42c42e7cd3ac9ee1b5a15f16ff25b325d0482e4"><code>b42c42e</code></a> chore: disallow use of deprecated <code>type</code> property in core rule tests (<a href="https://redirect.github.com/eslint/eslint/issues/20094">#20094</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/7bb498d720dcd054cc042ca4b60b138d8485f07c"><code>7bb498d</code></a> test: remove deprecated <code>type</code> property from core rule tests (<a href="https://redirect.github.com/eslint/eslint/issues/20093">#20093</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/e10cf2ab42fe5b481d980dc652f7504414747733"><code>e10cf2a</code></a> ci: bump actions/setup-node from 4 to 5 (<a href="https://redirect.github.com/eslint/eslint/issues/20089">#20089</a>) (dependabot[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/5cb0ce48ef6cfbbe6d09131c33a53f9d66fe9bd4"><code>5cb0ce4</code></a> refactor: use <code>meta.defaultOptions</code> in <code>preserve-caught-error</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20080">#20080</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/f9f7cb578dced3c14f635e17c75aa6744d291f4d"><code>f9f7cb5</code></a> chore: package.json update for eslint-config-eslint release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/81764b298065a328038cd067bc8fedef97e57500"><code>81764b2</code></a> chore: update <code>eslint</code> peer dependency in <code>eslint-config-eslint</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20079">#20079</a>) (Milos Djermanovic)</li>
</ul>
<h2>v9.35.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/42761fa7c872fb9e14c144b692af6967b3662082"><code>42761fa</code></a> feat: implement suggestions for no-empty-function (<a href="https://redirect.github.com/eslint/eslint/issues/20057">#20057</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/102f44442ac9bf1fcd4ba6ab9fae43ce09199df6"><code>102f444</code></a> feat: implement suggestions for no-empty-static-block (<a href="https://redirect.github.com/eslint/eslint/issues/20056">#20056</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/e51ffff737ca245b3a1d115cb11e1c99737249a3"><code>e51ffff</code></a> feat: add <code>preserve-caught-error</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19913">#19913</a>) (Amnish Singh Arora)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/10e7ae23e30ea0834d9fdeb3a2a1db8103c36cd2"><code>10e7ae2</code></a> fix: update uncloneable options error message (<a href="https://redirect.github.com/eslint/eslint/issues/20059">#20059</a>) (soda-sorcery)</li>
<li><a href="https://github.com/eslint/eslint/commit/bfa46013e7ea9a522c02f72250fa07160f96a6b8"><code>bfa4601</code></a> fix: ignore empty switch statements with comments in no-empty rule (<a href="https://redirect.github.com/eslint/eslint/issues/20045">#20045</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/dfd11deb24fc733faa5db751a2f615eb04e48b15"><code>dfd11de</code></a> fix: add <code>before</code> and <code>after</code> to test case types (<a href="https://redirect.github.com/eslint/eslint/issues/20049">#20049</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/dabbe95c39671c5fa272da012ee1432aa088650f"><code>dabbe95</code></a> fix: correct types for <code>no-restricted-imports</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/20034">#20034</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/ea789c7dd234c1a6be499a4644dd0f5c97615972"><code>ea789c7</code></a> fix: no-loss-of-precision false positive with uppercase exponent (<a href="https://redirect.github.com/eslint/eslint/issues/20032">#20032</a>) (sethamus)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/d265515642f65246bcd45c17979f67c2afb12f95"><code>d265515</code></a> docs: improve phrasing - &quot;if&quot; → &quot;even if&quot; from getting-started section (<a href="https://redirect.github.com/eslint/eslint/issues/20074">#20074</a>) (jjangga0214)</li>
<li><a href="https://github.com/eslint/eslint/commit/a355a0e5b2e6a47cda099b31dc7d112cfb5c4315"><code>a355a0e</code></a> docs: invert comparison logic for example in <code>no-var</code> doc page (<a href="https://redirect.github.com/eslint/eslint/issues/20064">#20064</a>) (OTonGitHub)</li>
<li><a href="https://github.com/eslint/eslint/commit/5082fc206de6946d9d4c20e57301f78839b3b9f2"><code>5082fc2</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/99cfd7e056e1703941c9eb8ca1ae7fdb1987ba9d"><code>99cfd7e</code></a> docs: add missing &quot;the&quot; in rule deprecation docs (<a href="https://redirect.github.com/eslint/eslint/issues/20050">#20050</a>) (Josh Goldberg ✨)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/blob/main/CHANGELOG.md">eslint's changelog</a>.</em></p>
<blockquote>
<p>v9.36.0 - September 19, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/12411e8d450ed26a5f7cca6a78ec05323c9323e8"><code>12411e8</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.36.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20139">#20139</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/488cba6b391b97b2cfc74bbb46fdeacb1361949e"><code>488cba6</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/b73ab12acd3e87f8d8173cda03499f6cd1f26db6"><code>b73ab12</code></a> docs: update examples to use <code>defineConfig</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20131">#20131</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/47afcf668df65eac68d7b04145d037037010a076"><code>47afcf6</code></a> feat: correct <code>preserve-caught-error</code> edge cases (<a href="https://redirect.github.com/eslint/eslint/issues/20109">#20109</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/75b74d865d3b8e7fa3bcf5ad29f4bf6d18d1310e"><code>75b74d8</code></a> fix: add missing rule option types (<a href="https://redirect.github.com/eslint/eslint/issues/20127">#20127</a>) (ntnyq)</li>
<li><a href="https://github.com/eslint/eslint/commit/bac82a2a9c80a3f69087852758d7737aea371f09"><code>bac82a2</code></a> ci: simplify renovate configuration (<a href="https://redirect.github.com/eslint/eslint/issues/19907">#19907</a>) (唯然)</li>
<li><a href="https://github.com/eslint/eslint/commit/1c0d85049e3f30a8809340c1abc881c63b7812ff"><code>1c0d850</code></a> fix: update <code>eslint-all.js</code> to use <code>Object.freeze</code> for <code>rules</code> object (<a href="https://redirect.github.com/eslint/eslint/issues/20116">#20116</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/c00bb37d62c1bcc0a37f094371be9c40064009f1"><code>c00bb37</code></a> ci: bump actions/labeler from 5 to 6 (<a href="https://redirect.github.com/eslint/eslint/issues/20090">#20090</a>) (dependabot[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/fee751dc8aeab54547af4538332ea5c069ef28b6"><code>fee751d</code></a> refactor: use <code>defaultOptions</code> in rules (<a href="https://redirect.github.com/eslint/eslint/issues/20121">#20121</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/31d93926990fba536846ec727d7a2625fc844649"><code>31d9392</code></a> docs: fix typos (<a href="https://redirect.github.com/eslint/eslint/issues/20118">#20118</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/7d61b7fadc9c5c6f2b131e37e8a3cffa5aae8ee6"><code>7d61b7f</code></a> fix: add missing scope types to <code>Scope.type</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20110">#20110</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/1ace67d9f7903adc3d3f09868aa05b673e7d3f3b"><code>1ace67d</code></a> chore: update example to use <code>defineConfig</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20111">#20111</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/4821963bf765532069c49e9da9ecbe9485b073fc"><code>4821963</code></a> test: add missing loc information to error objects in rule tests (<a href="https://redirect.github.com/eslint/eslint/issues/20112">#20112</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/c7f861b3f8c1ac961b4cd4f22483798f3324c62b"><code>c7f861b</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/6b0c08b106aa66f2e9fa484282f0eb63c64a1215"><code>6b0c08b</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/7a670c301b58609017ce8cfda99ee81f95de3898"><code>7a670c3</code></a> fix: correct rule option typings in <code>rules.d.ts</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20084">#20084</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/b42c42e7cd3ac9ee1b5a15f16ff25b325d0482e4"><code>b42c42e</code></a> chore: disallow use of deprecated <code>type</code> property in core rule tests (<a href="https://redirect.github.com/eslint/eslint/issues/20094">#20094</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/7bb498d720dcd054cc042ca4b60b138d8485f07c"><code>7bb498d</code></a> test: remove deprecated <code>type</code> property from core rule tests (<a href="https://redirect.github.com/eslint/eslint/issues/20093">#20093</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/91f97c50468fbdc089c91e99c2ea0fe821911df2"><code>91f97c5</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/e10cf2ab42fe5b481d980dc652f7504414747733"><code>e10cf2a</code></a> ci: bump actions/setup-node from 4 to 5 (<a href="https://redirect.github.com/eslint/eslint/issues/20089">#20089</a>) (dependabot[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/5cb0ce48ef6cfbbe6d09131c33a53f9d66fe9bd4"><code>5cb0ce4</code></a> refactor: use <code>meta.defaultOptions</code> in <code>preserve-caught-error</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20080">#20080</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/f9f7cb578dced3c14f635e17c75aa6744d291f4d"><code>f9f7cb5</code></a> chore: package.json update for eslint-config-eslint release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/81764b298065a328038cd067bc8fedef97e57500"><code>81764b2</code></a> chore: update <code>eslint</code> peer dependency in <code>eslint-config-eslint</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20079">#20079</a>) (Milos Djermanovic)</li>
</ul>
<p>v9.35.0 - September 5, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/da87f2fe792cab5b69b62bf5c15e69ab4f433087"><code>da87f2f</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.35.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20077">#20077</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/af2a0870fdc646091d027516601888923e5bc202"><code>af2a087</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/d265515642f65246bcd45c17979f67c2afb12f95"><code>d265515</code></a> docs: improve phrasing - &quot;if&quot; → &quot;even if&quot; from getting-started section (<a href="https://redirect.github.com/eslint/eslint/issues/20074">#20074</a>) (jjangga0214)</li>
<li><a href="https://github.com/eslint/eslint/commit/70557649e3111c55d8cddf678b6c4079aa6f0ccc"><code>7055764</code></a> test: remove <code>tests/lib/eslint/eslint.config.js</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20065">#20065</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/10e7ae23e30ea0834d9fdeb3a2a1db8103c36cd2"><code>10e7ae2</code></a> fix: update uncloneable options error message (<a href="https://redirect.github.com/eslint/eslint/issues/20059">#20059</a>) (soda-sorcery)</li>
<li><a href="https://github.com/eslint/eslint/commit/42761fa7c872fb9e14c144b692af6967b3662082"><code>42761fa</code></a> feat: implement suggestions for no-empty-function (<a href="https://redirect.github.com/eslint/eslint/issues/20057">#20057</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/102f44442ac9bf1fcd4ba6ab9fae43ce09199df6"><code>102f444</code></a> feat: implement suggestions for no-empty-static-block (<a href="https://redirect.github.com/eslint/eslint/issues/20056">#20056</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/84ffb9680b15e45bfd8c8a5db4731576ddd16fc4"><code>84ffb96</code></a> chore: update <code>@eslint-community/eslint-utils</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20069">#20069</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/a355a0e5b2e6a47cda099b31dc7d112cfb5c4315"><code>a355a0e</code></a> docs: invert comparison logic for example in <code>no-var</code> doc page (<a href="https://redirect.github.com/eslint/eslint/issues/20064">#20064</a>) (OTonGitHub)</li>
<li><a href="https://github.com/eslint/eslint/commit/e51ffff737ca245b3a1d115cb11e1c99737249a3"><code>e51ffff</code></a> feat: add <code>preserve-caught-error</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19913">#19913</a>) (Amnish Singh Arora)</li>
<li><a href="https://github.com/eslint/eslint/commit/5082fc206de6946d9d4c20e57301f78839b3b9f2"><code>5082fc2</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/d5ef9397150cc178e1f9891c3ff49ac4871ec786"><code>d5ef939</code></a> refactor: remove deprecated <code>context.parserOptions</code> usage across rules (<a href="https://redirect.github.com/eslint/eslint/issues/20060">#20060</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/1b3881d7e859bec9589e39888656c33c914a8302"><code>1b3881d</code></a> chore: remove redundant word (<a href="https://redirect.github.com/eslint/eslint/issues/20058">#20058</a>) (pxwanglu)</li>
<li><a href="https://github.com/eslint/eslint/commit/99cfd7e056e1703941c9eb8ca1ae7fdb1987ba9d"><code>99cfd7e</code></a> docs: add missing &quot;the&quot; in rule deprecation docs (<a href="https://redirect.github.com/eslint/eslint/issues/20050">#20050</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/bfa46013e7ea9a522c02f72250fa07160f96a6b8"><code>bfa4601</code></a> fix: ignore empty switch statements with comments in no-empty rule (<a href="https://redirect.github.com/eslint/eslint/issues/20045">#20045</a>) (jaymarvelz)</li>
<li><a href="https://github.com/eslint/eslint/commit/dfd11deb24fc733faa5db751a2f615eb04e48b15"><code>dfd11de</code></a> fix: add <code>before</code> and <code>after</code> to test case types (<a href="https://redirect.github.com/eslint/eslint/issues/20049">#20049</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/6ad8973e5d3c94b8e100b7266f55f8eb0757eb00"><code>6ad8973</code></a> docs: update <code>--no-ignore</code> and <code>--ignore-pattern</code> documentation (<a href="https://redirect.github.com/eslint/eslint/issues/20036">#20036</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/dabbe95c39671c5fa272da012ee1432aa088650f"><code>dabbe95</code></a> fix: correct types for <code>no-restricted-imports</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/20034">#20034</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/8033b195299a1eaa4a0ed6553d9e034a457bb577"><code>8033b19</code></a> docs: add documentation for <code>--no-config-lookup</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20033">#20033</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/ea789c7dd234c1a6be499a4644dd0f5c97615972"><code>ea789c7</code></a> fix: no-loss-of-precision false positive with uppercase exponent (<a href="https://redirect.github.com/eslint/eslint/issues/20032">#20032</a>) (sethamus)</li>
</ul>
<p>v9.34.0 - August 22, 2025</p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/b4857e54e54b5dba96d156cd8d8b4d42dc5a3bf4"><code>b4857e5</code></a> 9.36.0</li>
<li><a href="https://github.com/eslint/eslint/commit/5878a4243f623b46b476eb81043d06244eae5877"><code>5878a42</code></a> Build: changelog update for 9.36.0</li>
<li><a href="https://github.com/eslint/eslint/commit/12411e8d450ed26a5f7cca6a78ec05323c9323e8"><code>12411e8</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.36.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20139">#20139</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/488cba6b391b97b2cfc74bbb46fdeacb1361949e"><code>488cba6</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/b73ab12acd3e87f8d8173cda03499f6cd1f26db6"><code>b73ab12</code></a> docs: update examples to use <code>defineConfig</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20131">#20131</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/47afcf668df65eac68d7b04145d037037010a076"><code>47afcf6</code></a> feat: correct <code>preserve-caught-error</code> edge cases (<a href="https://redirect.github.com/eslint/eslint/issues/20109">#20109</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/75b74d865d3b8e7fa3bcf5ad29f4bf6d18d1310e"><code>75b74d8</code></a> fix: add missing rule option types (<a href="https://redirect.github.com/eslint/eslint/issues/20127">#20127</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/bac82a2a9c80a3f69087852758d7737aea371f09"><code>bac82a2</code></a> ci: simplify renovate configuration (<a href="https://redirect.github.com/eslint/eslint/issues/19907">#19907</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/1c0d85049e3f30a8809340c1abc881c63b7812ff"><code>1c0d850</code></a> fix: update <code>eslint-all.js</code> to use <code>Object.freeze</code> for <code>rules</code> object (<a href="https://redirect.github.com/eslint/eslint/issues/20116">#20116</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/c00bb37d62c1bcc0a37f094371be9c40064009f1"><code>c00bb37</code></a> ci: bump actions/labeler from 5 to 6 (<a href="https://redirect.github.com/eslint/eslint/issues/20090">#20090</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v7.32.0...v9.36.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~eslintbot">eslintbot</a>, a new releaser for eslint since your current version.</p>
</details>
<br />

Updates `eslint-config-prettier` from 7.2.0 to 10.1.8
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
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/9b0b0a47ec28a7a83cf65e8436a8776910379385"><code>9b0b0a4</code></a> fix: release a new latest version</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4c9489339d37bf96d31e0596e64bb8d4cb4308ef"><code>4c94893</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/333">#333</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4"><code>60fef02</code></a> chore: add <code>funding</code> field into <code>package.json</code> (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/332">#332</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/f55501ffe9be65fc9a8ec7d788459fd3a9cb6095"><code>f55501f</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/329">#329</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/50a8a22b0468e3469b7a177e6c81e843bd5cb73e"><code>50a8a22</code></a> chore(deps): update all dependencies (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/330">#330</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850"><code>94b4799</code></a> fix(cli): do not crash on no rules configured (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/328">#328</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/cdc4a5c7e39e7f2d5760c60ea39cecb028fb34dc"><code>cdc4a5c</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/326">#326</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df"><code>4e95a1d</code></a> fix: this package is <code>commonjs</code>, align its types correctly (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/325">#325</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/8911369cbc66f1f859e19751eaefdea687129de5"><code>8911369</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/322">#322</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> chore(package): add homepage url (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/321">#321</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prettier/eslint-config-prettier/compare/v7.2.0...v10.1.8">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~jounqin">jounqin</a>, a new releaser for eslint-config-prettier since your current version.</p>
</details>
<br />

Updates `eslint-plugin-cypress` from 2.15.2 to 5.1.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cypress-io/eslint-plugin-cypress/releases">eslint-plugin-cypress's releases</a>.</em></p>
<blockquote>
<h2>v5.1.1</h2>
<h2><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v5.1.0...v5.1.1">5.1.1</a> (2025-08-14)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>dont throw warning on variable data selector (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/272">#272</a>) (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/0508b75a745bb4f5cc4d64cfbb9206543d407ac7">0508b75</a>)</li>
</ul>
<h2>v5.1.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v5.0.1...v5.1.0">5.1.0</a> (2025-06-03)</h1>
<h3>Features</h3>
<ul>
<li>publish only core files to npm package (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/263">#263</a>) (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/e622a58bf7882ad00cd928dc7e83b2660b33483a">e622a58</a>)</li>
</ul>
<h2>v5.0.1</h2>
<h2><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v5.0.0...v5.0.1">5.0.1</a> (2025-05-30)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>stylistic:</strong> apply javascript formatting conventions (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/262">#262</a>) (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/bbd33fdc74303c608892e071c15627d0cb56ebdb">bbd33fd</a>)</li>
</ul>
<h2>v5.0.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.3.0...v5.0.0">5.0.0</a> (2025-05-29)</h1>
<h3>Breaking Changes</h3>
<ul>
<li>Support for non-flat legacy configurations is removed. The default configuration <code>eslint-plugin-cypress</code> now resolves to the flat configuration instead of to the legacy configuration</li>
</ul>
<h3>Deprecations</h3>
<ul>
<li><code>eslint-plugin-cypress/flat</code> is deprecated and should be replaced by <code>eslint-plugin-cypress</code></li>
</ul>
<h3>Other</h3>
<ul>
<li>remove legacy eslintrc configuration (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/260">#260</a>) (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/618bd21acdb2288e7ab8fc03c7abc0423270a3b8">618bd21</a>)</li>
<li>globals dependency is updated to 16.2.0</li>
<li>semantic-release dependency is update to 24.2.5</li>
</ul>
<h2>v4.3.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.2.1...v4.3.0">4.3.0</a> (2025-04-22)</h1>
<h3>Features</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/0508b75a745bb4f5cc4d64cfbb9206543d407ac7"><code>0508b75</code></a> fix: dont throw warning on variable data selector (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/272">#272</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/3fcfb93d8e950de8a8bd1c722e767946e01a0dcb"><code>3fcfb93</code></a> docs: migrate <a href="https://docs.cypress.io/guides">https://docs.cypress.io/guides</a> links (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/271">#271</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/22b6122daba733a2cb47d77c892b38b2844369e1"><code>22b6122</code></a> test: extend linting to customized test-project eslint config files (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/267">#267</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/b974e989af930ce97246ab2167a4a2f04d529af5"><code>b974e98</code></a> docs: harmonize markdown docs with prettier linting (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/266">#266</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/8c9d1a0bae48e5b66eb9438e71383582695205f5"><code>8c9d1a0</code></a> ci: add prettier linting (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/264">#264</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/e622a58bf7882ad00cd928dc7e83b2660b33483a"><code>e622a58</code></a> feat: publish only core files to npm package (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/263">#263</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/bbd33fdc74303c608892e071c15627d0cb56ebdb"><code>bbd33fd</code></a> fix(stylistic): apply javascript formatting conventions (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/262">#262</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/2412c29423c1a4113710c56a0eff97a50a3f04b4"><code>2412c29</code></a> test: migrate from jest to vitest 3.1.4 (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/261">#261</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/618bd21acdb2288e7ab8fc03c7abc0423270a3b8"><code>618bd21</code></a> feat: remove legacy eslintrc configuration (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/260">#260</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1a4b2cb9a8f31b76ddd66017403b2340ef1edd4b"><code>1a4b2cb</code></a> chore(deps): update eslint to 9.27.0 (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/259">#259</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v2.15.2...v5.1.1">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 16.23.1 to 16.24.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
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
<li><a href="https://github.com/stylelint/stylelint/commit/34ec717b7f167f66a9015f85206e9c6214fdbc0a"><code>34ec717</code></a> 16.24.0</li>
<li><a href="https://github.com/stylelint/stylelint/commit/a55744455046a0626a6a2f07a0a41bfd582ebc57"><code>a557444</code></a> Document changelog summary (<a href="https://redirect.github.com/stylelint/stylelint/issues/8751">#8751</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/6f37f764a3cefac956e08249b426125694ede54f"><code>6f37f76</code></a> Document using latest version of create script (<a href="https://redirect.github.com/stylelint/stylelint/issues/8750">#8750</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/6567514d74d9182d5b3d93122deff1f36cb28fee"><code>6567514</code></a> Prepare 16.24.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8725">#8725</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/57b454298ccb5e7d604d812c1ab64b8772e179ed"><code>57b4542</code></a> Document getting started with alternative package managers (<a href="https://redirect.github.com/stylelint/stylelint/issues/8726">#8726</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/55377c7452bca7e678e6f5788032a7ea0b7cd90f"><code>55377c7</code></a> Fix <code>selector-pseudo-class-no-unknown</code> false positives for <code>:heading</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8749">#8749</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/cfbe7720b53be99a1e75cbaa01fade21351e43d3"><code>cfbe772</code></a> Add <code>ignoreAtRules: []</code> to <code>nesting-selector-no-missing-scoping-root</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8743">#8743</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/5d642e9c9b6b8c7a453903350f112f991cdd9430"><code>5d642e9</code></a> Bump the jest group with 2 updates (<a href="https://redirect.github.com/stylelint/stylelint/issues/8742">#8742</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/e4eac5632c77c17aa72f251eacf3dff461c50c92"><code>e4eac56</code></a> Bump rollup from 4.48.0 to 4.50.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8741">#8741</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/e98952d2436f55319ab4a7d93d85c5310610e54b"><code>e98952d</code></a> Bump rollup from 4.46.2 to 4.48.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8738">#8738</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/16.23.1...16.24.0">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint-config-recommended-scss` from 14.1.0 to 16.0.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/releases">stylelint-config-recommended-scss's releases</a>.</em></p>
<blockquote>
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
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/a9edfb274b12b75e1319388bea3ca73e6df56907"><code>a9edfb2</code></a> 16.0.1</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/02dc994e6ddc8f69e2bf68578b031f58ee1f802b"><code>02dc994</code></a> Prepare 16.0.1</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/82329b174b110bc02ce0cc0fa24dcd5ce9e6ed69"><code>82329b1</code></a> Revert stylelint peer dep version bump (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/368">#368</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/02e99ca3a19a1bf81747b6bdd525273f54a04e74"><code>02e99ca</code></a> Bump eslint from 9.32.0 to 9.35.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/367">#367</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/0c168c85de633d2a84197bf0128ed99a477caa8e"><code>0c168c8</code></a> Bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/363">#363</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/c554a0e637101ed37403499b84b0a0cce3ee7e8c"><code>c554a0e</code></a> Fix false positive for nesting-selector-no-missing-scoping-root (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/365">#365</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/c55487214c69b5b899b38076b964e14c3c9bcfb1"><code>c554872</code></a> 16.0.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/829454dbb7d9d4f913cc93386e8609d4e8d926d0"><code>829454d</code></a> Prepare 16.0.0</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/16ccca803cd16f135c97161060e59f18509c0e14"><code>16ccca8</code></a> Bump eslint-config-stylelint from 24.0.0 to 25.0.0 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/353">#353</a>)</li>
<li><a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/commit/095b126b4bdf9ab8c35cfc3b97587da7e32fc853"><code>095b126</code></a> Bump npm-run-all2 from 8.0.1 to 8.0.4 (<a href="https://redirect.github.com/stylelint-scss/stylelint-config-recommended-scss/issues/351">#351</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint-scss/stylelint-config-recommended-scss/compare/v14.1.0...v16.0.1">compare view</a></li>
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

### Comment by @dependabot on November 10, 2025 at 06:11 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1386*
