---
type: pull_request
number: 1312
title: "chore(deps): bump the lint group across 1 directory with 6 updates"
state: closed
author: dependabot
created: 2025-04-07T19:15:34Z
updated: 2025-05-06T07:37:50Z
closed: 2025-05-06T07:37:48Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-ac20b27692
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1312
---

# Pull Request #1312: chore(deps): bump the lint group across 1 directory with 6 updates

**Author**: @dependabot
**Created**: April 07, 2025 at 07:15 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-ac20b27692`

## Description

Bumps the lint group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [eslint](https://github.com/eslint/eslint) | `7.32.0` | `9.24.0` |
| [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier) | `7.2.0` | `10.1.1` |
| [eslint-plugin-cypress](https://github.com/cypress-io/eslint-plugin-cypress) | `2.15.2` | `4.2.0` |
| [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) | `7.37.4` | `7.37.5` |
| [stylelint](https://github.com/stylelint/stylelint) | `16.14.1` | `16.18.0` |
| [stylelint-scss](https://github.com/stylelint-scss/stylelint-scss) | `6.11.0` | `6.11.1` |


Updates `eslint` from 7.32.0 to 9.24.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v9.24.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/556c25bbadd640ba9465ca6ec152f10959591666"><code>556c25b</code></a> feat: support loading TS config files using <code>--experimental-strip-types</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19401">#19401</a>) (Arya Emami)</li>
<li><a href="https://github.com/eslint/eslint/commit/72650acdb715fc25c675dc6368877b0e3f8d3885"><code>72650ac</code></a> feat: support TS syntax in <code>init-declarations</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19540">#19540</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/03fb0bca2be41597fcea7c0e84456bbaf2e5acca"><code>03fb0bc</code></a> feat: normalize patterns to handle &quot;./&quot; prefix in files and ignores (<a href="https://redirect.github.com/eslint/eslint/issues/19568">#19568</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/071dcd3a8e34aeeb52d0b9c23c2c4a1e58b45858"><code>071dcd3</code></a> feat: support TS syntax in <code>no-dupe-class-members</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19558">#19558</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/cd72bcc0c8d81fbf47ff3c8fe05ae48e1d862246"><code>cd72bcc</code></a> feat: Introduce a way to suppress violations (<a href="https://redirect.github.com/eslint/eslint/issues/19159">#19159</a>) (Iacovos Constantinou)</li>
<li><a href="https://github.com/eslint/eslint/commit/2a81578ac179b1eeb1484fddee31913ed99042a2"><code>2a81578</code></a> feat: support TS syntax in <code>no-loss-of-precision</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19560">#19560</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/30ae4ed093d19e9950d09c2ab57f43d3564e31c9"><code>30ae4ed</code></a> feat: add new options to class-methods-use-this (<a href="https://redirect.github.com/eslint/eslint/issues/19527">#19527</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/b79ade6c1e0765457637f7ddaa52a39eed3aad38"><code>b79ade6</code></a> feat: support TypeScript syntax in <code>no-array-constructor</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19493">#19493</a>) (Tanuj Kanti)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/b23d1c5f0297c5e2e9a4ff70533f3c0bdbfc34b8"><code>b23d1c5</code></a> fix: deduplicate variable names in no-loop-func error messages (<a href="https://redirect.github.com/eslint/eslint/issues/19595">#19595</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/fb8cdb842edcc035969e14b7b7e3ee372304f2d7"><code>fb8cdb8</code></a> fix: use <code>any[]</code> type for <code>context.options</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19584">#19584</a>) (Francesco Trotta)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/f8578206cc9b9fcd03dc5311f8a2d96b7b3359a5"><code>f857820</code></a> docs: update documentation for <code>--experimental-strip-types</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19594">#19594</a>) (Nikolas Schröter)</li>
<li><a href="https://github.com/eslint/eslint/commit/803e4af48e7fc3a2051e8c384f30fe4a318520e3"><code>803e4af</code></a> docs: simplify gitignore path handling in includeIgnoreFile section (<a href="https://redirect.github.com/eslint/eslint/issues/19596">#19596</a>) (Thomas Broyer)</li>
<li><a href="https://github.com/eslint/eslint/commit/6d979ccc183454e616bc82a598db5402e9d63dcf"><code>6d979cc</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/82177e4108d6b3e63ece6072d923c0a2c08907bf"><code>82177e4</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/e849dc01286cde5b6e2f0e04bf36928710633715"><code>e849dc0</code></a> docs: replace existing var with const (<a href="https://redirect.github.com/eslint/eslint/issues/19578">#19578</a>) (Sweta Tanwar)</li>
<li><a href="https://github.com/eslint/eslint/commit/0c65c628022ff3ce40598c1a6ce95728e7eda317"><code>0c65c62</code></a> docs: don't pass filename when linting rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19571">#19571</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/6be36c99432ecdc72e33b6fb3293cf28f66ab78d"><code>6be36c9</code></a> docs: Update custom-rules.md code example of fixer (<a href="https://redirect.github.com/eslint/eslint/issues/19555">#19555</a>) (Yifan Pan)</li>
</ul>
<h2>Build Related</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/366e3694afd85ab6605adf4fee4dfa1316be8b74"><code>366e369</code></a> build: re-enable Prettier formatting for <code>package.json</code> files (<a href="https://redirect.github.com/eslint/eslint/issues/19569">#19569</a>) (Francesco Trotta)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/ef6742091d49fc1809ad933f1daeff7124f57e93"><code>ef67420</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.24.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19602">#19602</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/4946847bb675ee26c3a52bfe3bca63a0dc5e5c61"><code>4946847</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/a995acbe32471ce8c20cbf9f48b4f3e1d8bc2229"><code>a995acb</code></a> chore: correct 'flter'/'filter' typo in package script (<a href="https://redirect.github.com/eslint/eslint/issues/19587">#19587</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/b9a5efa937046f2d3f97e6caabb67a4bc182c983"><code>b9a5efa</code></a> test: skip symlink test on Windows (<a href="https://redirect.github.com/eslint/eslint/issues/19503">#19503</a>) (fisker Cheung)</li>
<li><a href="https://github.com/eslint/eslint/commit/46eea6d1cbed41d020cb76841ebba30710b0afd0"><code>46eea6d</code></a> chore: remove <code>Rule</code> &amp; <code>FormatterFunction</code> from <code>shared/types.js</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19556">#19556</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/bdcc91d5b61ad1b3e55044767362548c906f5462"><code>bdcc91d</code></a> chore: modify .editorconfig to keep parity with prettier config (<a href="https://redirect.github.com/eslint/eslint/issues/19577">#19577</a>) (Sweta Tanwar)</li>
<li><a href="https://github.com/eslint/eslint/commit/7790d8305a8cef7cc95c331205d59d6b3c2b4e2e"><code>7790d83</code></a> chore: fix some typos in comment (<a href="https://redirect.github.com/eslint/eslint/issues/19576">#19576</a>) (todaymoon)</li>
<li><a href="https://github.com/eslint/eslint/commit/76064a632438533bbb90e253ec72d172e948d200"><code>76064a6</code></a> test: ignore <code>package-lock.json</code> for <code>eslint-webpack-plugin</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19572">#19572</a>) (Francesco Trotta)</li>
</ul>
<h2>v9.23.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/557a0d23755f8af4f2aaab751805c7ba6496fc21"><code>557a0d2</code></a> feat: support TypeScript syntax in no-useless-constructor (<a href="https://redirect.github.com/eslint/eslint/issues/19535">#19535</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/83202412a1ceefd3eba4b97cc9dbe99ab70d59a2"><code>8320241</code></a> feat: support TypeScript syntax in <code>default-param-last</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19431">#19431</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/833c4a301d4f7d21583d520d20d8a6724171733f"><code>833c4a3</code></a> feat: defineConfig() supports &quot;flat/&quot; config prefix (<a href="https://redirect.github.com/eslint/eslint/issues/19533">#19533</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/4a0df16f1ba7bed02d15c561119623199ea2ace0"><code>4a0df16</code></a> feat: circular autofix/conflicting rules detection (<a href="https://redirect.github.com/eslint/eslint/issues/19514">#19514</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/be56a685bf1aadbf59d99d43e71c00802bc9ba27"><code>be56a68</code></a> feat: support TypeScript syntax in <code>class-methods-use-this</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19498">#19498</a>) (Josh Goldberg ✨)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/0e20aa72fec53b16a21c42ac9e82969efa8f94d2"><code>0e20aa7</code></a> fix: move deprecated <code>RuleContext</code> methods to subtype (<a href="https://redirect.github.com/eslint/eslint/issues/19531">#19531</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/cc3bd00795708c4d7c06a6103983245cc9d9845b"><code>cc3bd00</code></a> fix: reporting variable used in catch block in <code>no-useless-assignment</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19423">#19423</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/d46ff832195aa841224a21086afda9d98be45ad6"><code>d46ff83</code></a> fix: <code>no-dupe-keys</code> false positive with proto setter (<a href="https://redirect.github.com/eslint/eslint/issues/19508">#19508</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/e7327736b92686e02721461ac9ccf6e65e0badac"><code>e732773</code></a> fix: navigation of search results on pressing Enter (<a href="https://redirect.github.com/eslint/eslint/issues/19502">#19502</a>) (Tanuj Kanti)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/blob/main/CHANGELOG.md">eslint's changelog</a>.</em></p>
<blockquote>
<p>v9.24.0 - April 4, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/ef6742091d49fc1809ad933f1daeff7124f57e93"><code>ef67420</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.24.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19602">#19602</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/4946847bb675ee26c3a52bfe3bca63a0dc5e5c61"><code>4946847</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/f8578206cc9b9fcd03dc5311f8a2d96b7b3359a5"><code>f857820</code></a> docs: update documentation for <code>--experimental-strip-types</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19594">#19594</a>) (Nikolas Schröter)</li>
<li><a href="https://github.com/eslint/eslint/commit/803e4af48e7fc3a2051e8c384f30fe4a318520e3"><code>803e4af</code></a> docs: simplify gitignore path handling in includeIgnoreFile section (<a href="https://redirect.github.com/eslint/eslint/issues/19596">#19596</a>) (Thomas Broyer)</li>
<li><a href="https://github.com/eslint/eslint/commit/6d979ccc183454e616bc82a598db5402e9d63dcf"><code>6d979cc</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/b23d1c5f0297c5e2e9a4ff70533f3c0bdbfc34b8"><code>b23d1c5</code></a> fix: deduplicate variable names in no-loop-func error messages (<a href="https://redirect.github.com/eslint/eslint/issues/19595">#19595</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/556c25bbadd640ba9465ca6ec152f10959591666"><code>556c25b</code></a> feat: support loading TS config files using <code>--experimental-strip-types</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19401">#19401</a>) (Arya Emami)</li>
<li><a href="https://github.com/eslint/eslint/commit/82177e4108d6b3e63ece6072d923c0a2c08907bf"><code>82177e4</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/a995acbe32471ce8c20cbf9f48b4f3e1d8bc2229"><code>a995acb</code></a> chore: correct 'flter'/'filter' typo in package script (<a href="https://redirect.github.com/eslint/eslint/issues/19587">#19587</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/72650acdb715fc25c675dc6368877b0e3f8d3885"><code>72650ac</code></a> feat: support TS syntax in <code>init-declarations</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19540">#19540</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/03fb0bca2be41597fcea7c0e84456bbaf2e5acca"><code>03fb0bc</code></a> feat: normalize patterns to handle &quot;./&quot; prefix in files and ignores (<a href="https://redirect.github.com/eslint/eslint/issues/19568">#19568</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/b9a5efa937046f2d3f97e6caabb67a4bc182c983"><code>b9a5efa</code></a> test: skip symlink test on Windows (<a href="https://redirect.github.com/eslint/eslint/issues/19503">#19503</a>) (fisker Cheung)</li>
<li><a href="https://github.com/eslint/eslint/commit/46eea6d1cbed41d020cb76841ebba30710b0afd0"><code>46eea6d</code></a> chore: remove <code>Rule</code> &amp; <code>FormatterFunction</code> from <code>shared/types.js</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19556">#19556</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/fb8cdb842edcc035969e14b7b7e3ee372304f2d7"><code>fb8cdb8</code></a> fix: use <code>any[]</code> type for <code>context.options</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19584">#19584</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/071dcd3a8e34aeeb52d0b9c23c2c4a1e58b45858"><code>071dcd3</code></a> feat: support TS syntax in <code>no-dupe-class-members</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19558">#19558</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/e849dc01286cde5b6e2f0e04bf36928710633715"><code>e849dc0</code></a> docs: replace existing var with const (<a href="https://redirect.github.com/eslint/eslint/issues/19578">#19578</a>) (Sweta Tanwar)</li>
<li><a href="https://github.com/eslint/eslint/commit/bdcc91d5b61ad1b3e55044767362548c906f5462"><code>bdcc91d</code></a> chore: modify .editorconfig to keep parity with prettier config (<a href="https://redirect.github.com/eslint/eslint/issues/19577">#19577</a>) (Sweta Tanwar)</li>
<li><a href="https://github.com/eslint/eslint/commit/7790d8305a8cef7cc95c331205d59d6b3c2b4e2e"><code>7790d83</code></a> chore: fix some typos in comment (<a href="https://redirect.github.com/eslint/eslint/issues/19576">#19576</a>) (todaymoon)</li>
<li><a href="https://github.com/eslint/eslint/commit/cd72bcc0c8d81fbf47ff3c8fe05ae48e1d862246"><code>cd72bcc</code></a> feat: Introduce a way to suppress violations (<a href="https://redirect.github.com/eslint/eslint/issues/19159">#19159</a>) (Iacovos Constantinou)</li>
<li><a href="https://github.com/eslint/eslint/commit/2a81578ac179b1eeb1484fddee31913ed99042a2"><code>2a81578</code></a> feat: support TS syntax in <code>no-loss-of-precision</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19560">#19560</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/366e3694afd85ab6605adf4fee4dfa1316be8b74"><code>366e369</code></a> build: re-enable Prettier formatting for <code>package.json</code> files (<a href="https://redirect.github.com/eslint/eslint/issues/19569">#19569</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/30ae4ed093d19e9950d09c2ab57f43d3564e31c9"><code>30ae4ed</code></a> feat: add new options to class-methods-use-this (<a href="https://redirect.github.com/eslint/eslint/issues/19527">#19527</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/b79ade6c1e0765457637f7ddaa52a39eed3aad38"><code>b79ade6</code></a> feat: support TypeScript syntax in <code>no-array-constructor</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19493">#19493</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/0c65c628022ff3ce40598c1a6ce95728e7eda317"><code>0c65c62</code></a> docs: don't pass filename when linting rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19571">#19571</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/76064a632438533bbb90e253ec72d172e948d200"><code>76064a6</code></a> test: ignore <code>package-lock.json</code> for <code>eslint-webpack-plugin</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19572">#19572</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/6be36c99432ecdc72e33b6fb3293cf28f66ab78d"><code>6be36c9</code></a> docs: Update custom-rules.md code example of fixer (<a href="https://redirect.github.com/eslint/eslint/issues/19555">#19555</a>) (Yifan Pan)</li>
</ul>
<p>v9.23.0 - March 21, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/0ac8ea45350fa5819694a3775641e94b1da3282b"><code>0ac8ea4</code></a> chore: update dependencies for v9.23.0 release (<a href="https://redirect.github.com/eslint/eslint/issues/19554">#19554</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/20591c49ff27435b1555111a929a6966febc249f"><code>20591c4</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/901344f9441c746dfa82261a0d00ff6ef35bcdf1"><code>901344f</code></a> chore: update dependency <code>@​eslint/json</code> to ^0.11.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19552">#19552</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/557a0d23755f8af4f2aaab751805c7ba6496fc21"><code>557a0d2</code></a> feat: support TypeScript syntax in no-useless-constructor (<a href="https://redirect.github.com/eslint/eslint/issues/19535">#19535</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/2357edd09beca1c3f70c92df23f2f99b9ebc7a70"><code>2357edd</code></a> build: exclude autogenerated files from Prettier formatting (<a href="https://redirect.github.com/eslint/eslint/issues/19548">#19548</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/5405939efcfe6a038a7c89354eae9c39c8ff21e3"><code>5405939</code></a> docs: show red underlines in TypeScript examples in rules docs (<a href="https://redirect.github.com/eslint/eslint/issues/19547">#19547</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/48b53d6e79945b4f5f66aa2073c2d51ff7896c7c"><code>48b53d6</code></a> docs: replace var with const in examples (<a href="https://redirect.github.com/eslint/eslint/issues/19539">#19539</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/0e20aa72fec53b16a21c42ac9e82969efa8f94d2"><code>0e20aa7</code></a> fix: move deprecated <code>RuleContext</code> methods to subtype (<a href="https://redirect.github.com/eslint/eslint/issues/19531">#19531</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/5228383e3e5c77c7dd07fc9d17b9a57c2ee5bb48"><code>5228383</code></a> chore: fix update-readme formatting (<a href="https://redirect.github.com/eslint/eslint/issues/19544">#19544</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/c39d7db7142ebdb8174da00358b80094eaad39c1"><code>c39d7db</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/a4f87604f4d8d53cb2efbd19aa067606dd1c409e"><code>a4f8760</code></a> docs: revert accidental changes (<a href="https://redirect.github.com/eslint/eslint/issues/19542">#19542</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/5439525925dc26b387cc6cebf0b01f42464b4ab0"><code>5439525</code></a> chore: format JSON files in Trunk (<a href="https://redirect.github.com/eslint/eslint/issues/19541">#19541</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/75adc99eab2878e58fc88f0d4b1b6f9091455914"><code>75adc99</code></a> chore: enabled Prettier in Trunk (<a href="https://redirect.github.com/eslint/eslint/issues/19354">#19354</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/239516856fbf61828f5ac2c8b45e245103c41c04"><code>2395168</code></a> chore: added .git-blame-ignore-revs for Prettier via trunk fmt (<a href="https://redirect.github.com/eslint/eslint/issues/19538">#19538</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/129882d2fdb4e7f597ed78eeadd86377f3d6b078"><code>129882d</code></a> chore: formatted files with Prettier via trunk fmt (<a href="https://redirect.github.com/eslint/eslint/issues/19355">#19355</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/1738dbc36ce556745c230d3592e7f1aa673a1430"><code>1738dbc</code></a> chore: temporarily disable prettier in trunk (<a href="https://redirect.github.com/eslint/eslint/issues/19537">#19537</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/83202412a1ceefd3eba4b97cc9dbe99ab70d59a2"><code>8320241</code></a> feat: support TypeScript syntax in <code>default-param-last</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19431">#19431</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/280128f73def56479e32e7d40879fff05b7f44a2"><code>280128f</code></a> docs: add copy button (<a href="https://redirect.github.com/eslint/eslint/issues/19512">#19512</a>) (xbinaryx)</li>
<li><a href="https://github.com/eslint/eslint/commit/833c4a301d4f7d21583d520d20d8a6724171733f"><code>833c4a3</code></a> feat: defineConfig() supports &quot;flat/&quot; config prefix (<a href="https://redirect.github.com/eslint/eslint/issues/19533">#19533</a>) (Nicholas C. Zakas)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/d49f5b7333e9a46aabdb0cff267a1d36cdbde598"><code>d49f5b7</code></a> 9.24.0</li>
<li><a href="https://github.com/eslint/eslint/commit/9b6ed8a57aa7ec558d6a8a0aca263053fc878eb1"><code>9b6ed8a</code></a> Build: changelog update for 9.24.0</li>
<li><a href="https://github.com/eslint/eslint/commit/ef6742091d49fc1809ad933f1daeff7124f57e93"><code>ef67420</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.24.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19602">#19602</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/4946847bb675ee26c3a52bfe3bca63a0dc5e5c61"><code>4946847</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/f8578206cc9b9fcd03dc5311f8a2d96b7b3359a5"><code>f857820</code></a> docs: update documentation for <code>--experimental-strip-types</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19594">#19594</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/803e4af48e7fc3a2051e8c384f30fe4a318520e3"><code>803e4af</code></a> docs: simplify gitignore path handling in includeIgnoreFile section (<a href="https://redirect.github.com/eslint/eslint/issues/19596">#19596</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/6d979ccc183454e616bc82a598db5402e9d63dcf"><code>6d979cc</code></a> docs: Update README</li>
<li><a href="https://github.com/eslint/eslint/commit/b23d1c5f0297c5e2e9a4ff70533f3c0bdbfc34b8"><code>b23d1c5</code></a> fix: deduplicate variable names in no-loop-func error messages (<a href="https://redirect.github.com/eslint/eslint/issues/19595">#19595</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/556c25bbadd640ba9465ca6ec152f10959591666"><code>556c25b</code></a> feat: support loading TS config files using <code>--experimental-strip-types</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19">#19</a>...</li>
<li><a href="https://github.com/eslint/eslint/commit/82177e4108d6b3e63ece6072d923c0a2c08907bf"><code>82177e4</code></a> docs: Update README</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v7.32.0...v9.24.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~eslintbot">eslintbot</a>, a new releaser for eslint since your current version.</p>
</details>
<br />

Updates `eslint-config-prettier` from 7.2.0 to 10.1.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/releases">eslint-config-prettier's releases</a>.</em></p>
<blockquote>
<h2>v10.1.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/309">#309</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - fix: separate the <code>/flat</code> entry for compatibility</p>
<p>For flat config users, the previous <code>&quot;eslint-config-prettier&quot;</code> entry still works, but <code>&quot;eslint-config-prettier/flat&quot;</code> adds a new <code>name</code> property for <a href="https://eslint.org/blog/2024/04/eslint-config-inspector/">config-inspector</a>, we just can't add it for the default entry for compatibility.</p>
<p>See also <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/308">prettier/eslint-config-prettier#308</a></p>
<pre lang="ts"><code>// before
import eslintConfigPrettier from &quot;eslint-config-prettier&quot;;
<p>// after
import eslintConfigPrettier from &quot;eslint-config-prettier/flat&quot;;
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
<h2>What's Changed</h2>
<ul>
<li>chore: migrate to changeset for automatically releasing by <a href="https://github.com/JounQin"><code>@​JounQin</code></a> in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/278">prettier/eslint-config-prettier#278</a></li>
<li>add support for <code>@stylistic/eslint-plugin</code> by <a href="https://github.com/abrahamguo"><code>@​abrahamguo</code></a> in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/272">prettier/eslint-config-prettier#272</a></li>
</ul>
<h2>New Contributors</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/blob/main/CHANGELOG.md">eslint-config-prettier's changelog</a>.</em></p>
<blockquote>
<h2>10.1.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/309">#309</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - fix: separate the <code>/flat</code> entry for compatibility</p>
<p>For flat config users, the previous <code>&quot;eslint-config-prettier&quot;</code> entry still works, but <code>&quot;eslint-config-prettier/flat&quot;</code> adds a new <code>name</code> property for <a href="https://eslint.org/blog/2024/04/eslint-config-inspector/">config-inspector</a>, we just can't add it for the default entry for compatibility.</p>
<p>See also <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/308">prettier/eslint-config-prettier#308</a></p>
<pre lang="ts"><code>// before
import eslintConfigPrettier from &quot;eslint-config-prettier&quot;;
<p>// after
import eslintConfigPrettier from &quot;eslint-config-prettier/flat&quot;;
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
<h2>Versions before 10.0.0</h2>
<h3>Version 9.1.0 (2023-12-02)</h3>
<ul>
<li>Added: [unicorn/template-indent], (as a [special rule][unicorn/template-indent-special]). Thanks to Gürgün Dayıoğlu (<a href="https://github.com/gurgunday"><code>@​gurgunday</code></a>)!</li>
<li>Changed: All the [formatting rules that were deprecated in ESLint 8.53.0][deprecated-8.53.0] are now excluded if you set the <code>ESLINT_CONFIG_PRETTIER_NO_DEPRECATED</code> environment variable.</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/f12309bbca9fb051b53fcece9a8491a1222235c8"><code>f12309b</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/310">#310</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/79cf6796bf9de35eb8965333145ca6e989045580"><code>79cf679</code></a> chore: use flat entry for flat config verification</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> fix: separate the <code>/flat</code> entry for compatibility (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/309">#309</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/2c6f80e67f391db7650e5ba32c7a5562d2029664"><code>2c6f80e</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/307">#307</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/56e2e3466391d0fdfc200e42130309c687aaab53"><code>56e2e34</code></a> feat: migrate to exports field (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/306">#306</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/c27e7855ee4c12f99d28c7d98c5883c4a04b3267"><code>c27e785</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/304">#304</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/c5d78532b73951313188ab0feb83224e78a9c6e7"><code>c5d7853</code></a> chore: add missing <code>@​stylistic/eslint-plugin-plus</code></li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/1837a4d1e5bc4ca26e81b1d2f6df96801e01795f"><code>1837a4d</code></a> ci: testing, migrate to yarn v4 (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/305">#305</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/8dbbd6d70b8a56cdfa4ea4e185d3699d5729b38e"><code>8dbbd6d</code></a> feat: add name to config (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/294">#294</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/cba57377e4c86d20d17042d6999eabc754fddc03"><code>cba5737</code></a> feat: add declaration file (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/280">#280</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prettier/eslint-config-prettier/compare/v7.2.0...v10.1.1">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~jounqin">jounqin</a>, a new releaser for eslint-config-prettier since your current version.</p>
</details>
<br />

Updates `eslint-plugin-cypress` from 2.15.2 to 4.2.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cypress-io/eslint-plugin-cypress/releases">eslint-plugin-cypress's releases</a>.</em></p>
<blockquote>
<h2>v4.2.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.1.0...v4.2.0">4.2.0</a> (2025-03-07)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>address comments in PR (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/64eaba0a572e15fd7025d6b426c38887c3ceb53c">64eaba0</a>)</li>
<li>doc title (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1fbedaca91e35715648091a3387f46b67bb56157">1fbedac</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>add rule disallow usage of cypress-xpath (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1ae902a3907c984820fbda2010e8c078d00fe503">1ae902a</a>)</li>
</ul>
<h2>v4.1.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.0.0...v4.1.0">4.1.0</a> (2024-10-30)</h1>
<h3>Features</h3>
<ul>
<li><strong>docs:</strong> publish updated readme (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/2bc8d5a82208f7da463f250573d493e6e6c287c4">2bc8d5a</a>)</li>
</ul>
<h2>v4.0.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v3.6.0...v4.0.0">4.0.0</a> (2024-10-11)</h1>
<h3>Features</h3>
<ul>
<li>minimum version eslint v9 (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/b50181ad06d423dd61e611a0e052a074758bfa8f">b50181a</a>)</li>
</ul>
<h3>BREAKING CHANGES</h3>
<ul>
<li>Support ESLint v9 and above only</li>
</ul>
<p>Support for ESlint v7 &amp; v8 is removed
languageOptions ecmaVersion: 2019 and sourceType: module are removed from eslint-plugin-cypress/flat config
globals is updated to 15.11.0
eslint-plugin-n is updated to 17.11.1</p>
<h2>v3.6.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v3.5.0...v3.6.0">3.6.0</a> (2024-10-11)</h1>
<h3>Features</h3>
<ul>
<li>publish eslint v8 deprecation (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/44a772272bbb7eb8a5631e65d7add4f31275d5bb">44a7722</a>)</li>
</ul>
<h2>v3.5.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v3.4.0...v3.5.0">3.5.0</a> (2024-08-12)</h1>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/e04d43c5547b1c4ef1e5c8e75120917d684700bb"><code>e04d43c</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/247">#247</a> from duranbe/no-xpath-rule</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/873c7a0ce64ef4ac332eabfc9d91965a56f69584"><code>873c7a0</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/1">#1</a> from MikeMcC399/pr/247-fix</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/9b70ddccf586b2a6d4fed5b002d3ba43b9f95d83"><code>9b70ddc</code></a> fix pr/247</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1fbedaca91e35715648091a3387f46b67bb56157"><code>1fbedac</code></a> fix: doc title</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/b75ca6ca4a810ff7a555ca1c838bb1d6ae58e741"><code>b75ca6c</code></a> Merge branch 'master' into no-xpath-rule</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/64eaba0a572e15fd7025d6b426c38887c3ceb53c"><code>64eaba0</code></a> fix: address comments in PR</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/9f2b56530e3181045ad55bfe6607964ffba778a0"><code>9f2b565</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/248">#248</a> from MikeMcC399/remove/eslint-plugin-n</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/60a7df3681bf96916d8878ef94804202096af7c9"><code>60a7df3</code></a> chore(deps): remove eslint-plugin-n</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1ae902a3907c984820fbda2010e8c078d00fe503"><code>1ae902a</code></a> feat: add rule disallow usage of cypress-xpath</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/19842d3a9f6d750761484a8013424a6c968da2b1"><code>19842d3</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/245">#245</a> from MikeMcC399/update/node-version</li>
<li>Additional commits viewable in <a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v2.15.2...v4.2.0">compare view</a></li>
</ul>
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

Updates `stylelint` from 16.14.1 to 16.18.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
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
<li>Added: <code>lightness-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8458">#8458</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>media-feature-name-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8456">#8456</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>media-feature-range-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8455">#8455</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>property-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8461">#8461</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>rule-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8460">#8460</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-no-vendor-prefix</code> support for <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8462">#8462</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-not-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8463">#8463</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-pseudo-element-colon-notation</code> support for <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8464">#8464</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-type-case</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8467">#8467</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>shorthand-property-no-redundant-values</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8466">#8466</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>value-keyword-case</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8469">#8469</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>value-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8470">#8470</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
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
<li>Added: <code>at-rule-no-vendor-prefix</code> support for c...

_Description has been truncated_

---

## Discussion

### Comment by @dependabot on May 06, 2025 at 07:37 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1312*
