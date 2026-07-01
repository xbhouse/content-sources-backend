---
type: pull_request
number: 1234
title: "chore(deps): bump the babel group with 4 updates"
state: closed
author: dependabot
created: 2025-01-08T10:20:21Z
updated: 2025-01-13T18:58:29Z
closed: 2025-01-13T18:58:27Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-b7993e4002
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1234
---

# Pull Request #1234: chore(deps): bump the babel group with 4 updates

**Author**: @dependabot
**Created**: January 08, 2025 at 10:20 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-b7993e4002`

## Description

Bumps the babel group with 4 updates: [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser), [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env), [@babel/preset-react](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react) and [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime).

Updates `@babel/eslint-parser` from 7.25.8 to 7.25.9
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/eslint-parser</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.25.9 (2024-10-22)</h2>
<p>Thanks <a href="https://github.com/victorenator"><code>@​victorenator</code></a> for your first PR!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-template</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16905">#16905</a> fix: Keep type annotations in <code>syntacticPlaceholders</code> mode (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-compilation-targets</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16907">#16907</a> fix: support BROWSERSLIST{,_CONFIG} env (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16884">#16884</a> Analyze <code>ClassAccessorProperty</code> to prevent the <code>no-undef</code> rule (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-transform-fixture-test-runner</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16914">#16914</a> remove test options flaky (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><em>Every package</em>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16917">#16917</a> fix: Accidentally published <code>tsconfig</code> files (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16918">#16918</a> perf: Make <code>VISITOR_KEYS</code> etc. faster to access (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 4</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Viktar Vaŭčkievič (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/eslint-parser</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.25.9 (2024-10-22)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-template</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16905">#16905</a> fix: Keep type annotations in <code>syntacticPlaceholders</code> mode (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-compilation-targets</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16907">#16907</a> fix: support BROWSERSLIST{,_CONFIG} env (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16884">#16884</a> Analyze <code>ClassAccessorProperty</code> to prevent the <code>no-undef</code> rule (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-transform-fixture-test-runner</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16914">#16914</a> remove test options flaky (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16918">#16918</a> perf: Make <code>VISITOR_KEYS</code> etc. faster to access (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/b07957ebb316a1e2fc67454fc7423508bb942e63"><code>b07957e</code></a> v7.25.9</li>
<li><a href="https://github.com/babel/babel/commit/af917594e4df3decdde2ce0b1614d607b27367a5"><code>af91759</code></a> fix: Accidentally publishing useless files (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/16917">#16917</a>)</li>
<li><a href="https://github.com/babel/babel/commit/33d705e8df0ec660c1e8e67154949d2d89269350"><code>33d705e</code></a> Analyze <code>ClassAccessorProperty</code> to prevent the <code>no-undef</code> rule (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/16884">#16884</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.25.9/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-env` from 7.25.8 to 7.26.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-env</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.26.0 (2024-10-25)</h2>
<p>Thanks <a href="https://github.com/timofei-iatsenko"><code>@​timofei-iatsenko</code></a> for your first PR!</p>
<p>You can find the release blog post with some highlights at <a href="https://babeljs.io/blog/2024/10/25/7.26.0">https://babeljs.io/blog/2024/10/25/7.26.0</a>.</p>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-core</code>, <code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-syntax-import-assertions</code>, <code>babel-plugin-syntax-import-attributes</code>, <code>babel-preset-env</code>, <code>babel-standalone</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16850">#16850</a> Enable import attributes parsing by default (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16862">#16862</a> feat: support async plugin's pre/post (<a href="https://github.com/timofei-iatsenko"><code>@​timofei-iatsenko</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-proposal-regexp-modifiers</code>, <code>babel-plugin-transform-regexp-modifiers</code>, <code>babel-preset-env</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16692">#16692</a> Add <code>transform-regexp-modifiers</code> to <code>preset-env</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16849">#16849</a> feat: add <code>startIndex</code> parser option (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-syntax-flow</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16841">#16841</a> Always enable parsing of Flow enums (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-preset-typescript</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16794">#16794</a> Support <code>import()</code> in <code>rewriteImportExtensions</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16708">#16708</a> Add experimental format-preserving mode to <code>@babel/generator</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16928">#16928</a> Workaround Node.js bug for parallel loading of TLA modules (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16926">#16926</a> Fix loading of modules with TLA in Node.js 23 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-proposal-json-modules</code>, <code>babel-plugin-transform-json-modules</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16924">#16924</a> Rename <code>proposal-json-modules</code> to <code>transform-json-modules</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-code-frame</code>, <code>babel-highlight</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16896">#16896</a> Inline <code>@babel/highlight</code> in <code>@babel/code-frame</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16732">#16732</a> Add <code>kind</code> to <code>TSModuleDeclaration</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-helper-module-transforms</code>, <code>babel-plugin-transform-modules-commonjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16882">#16882</a> perf: Improve module transforms (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Dylan Piercey (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li>Timofei Iatsenko (<a href="https://github.com/timofei-iatsenko"><code>@​timofei-iatsenko</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.25.9 (2024-10-22)</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-env</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.26.0 (2024-10-25)</h2>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-core</code>, <code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-syntax-import-assertions</code>, <code>babel-plugin-syntax-import-attributes</code>, <code>babel-preset-env</code>, <code>babel-standalone</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16850">#16850</a> Enable import attributes parsing by default (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16862">#16862</a> feat: support async plugin's pre/post (<a href="https://github.com/timofei-iatsenko"><code>@​timofei-iatsenko</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-proposal-regexp-modifiers</code>, <code>babel-plugin-transform-regexp-modifiers</code>, <code>babel-preset-env</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16692">#16692</a> Add <code>transform-regexp-modifiers</code> to <code>preset-env</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16849">#16849</a> feat: add <code>startIndex</code> parser option (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-syntax-flow</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16841">#16841</a> Always enable parsing of Flow enums (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-preset-typescript</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16794">#16794</a> Support <code>import()</code> in <code>rewriteImportExtensions</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16708">#16708</a> Add experimental format-preserving mode to <code>@babel/generator</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16928">#16928</a> Workaround Node.js bug for parallel loading of TLA modules (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16926">#16926</a> Fix loading of modules with TLA in Node.js 23 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-proposal-json-modules</code>, <code>babel-plugin-transform-json-modules</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16924">#16924</a> Rename <code>proposal-json-modules</code> to <code>transform-json-modules</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-code-frame</code>, <code>babel-highlight</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16896">#16896</a> Inline <code>@babel/highlight</code> in <code>@babel/code-frame</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16732">#16732</a> Add <code>kind</code> to <code>TSModuleDeclaration</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-helper-module-transforms</code>, <code>babel-plugin-transform-modules-commonjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16882">#16882</a> perf: Improve module transforms (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.25.9 (2024-10-22)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-template</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16905">#16905</a> fix: Keep type annotations in <code>syntacticPlaceholders</code> mode (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-compilation-targets</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16907">#16907</a> fix: support BROWSERSLIST{,_CONFIG} env (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16884">#16884</a> Analyze <code>ClassAccessorProperty</code> to prevent the <code>no-undef</code> rule (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-transform-fixture-test-runner</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16914">#16914</a> remove test options flaky (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/63d30381c169780460e01bdb6669c5e01af1dfbe"><code>63d3038</code></a> v7.26.0</li>
<li><a href="https://github.com/babel/babel/commit/64fa46676b5729fcc53fbc71ccd4615d3017fe08"><code>64fa466</code></a> Enable import attributes parsing by default (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/16850">#16850</a>)</li>
<li><a href="https://github.com/babel/babel/commit/816b293a698bce6162b2b3ed62a4bd6e9020eb3a"><code>816b293</code></a> Add <code>transform-regexp-modifiers</code> to <code>preset-env</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/16692">#16692</a>)</li>
<li><a href="https://github.com/babel/babel/commit/b07957ebb316a1e2fc67454fc7423508bb942e63"><code>b07957e</code></a> v7.25.9</li>
<li><a href="https://github.com/babel/babel/commit/af917594e4df3decdde2ce0b1614d607b27367a5"><code>af91759</code></a> fix: Accidentally publishing useless files (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/16917">#16917</a>)</li>
<li><a href="https://github.com/babel/babel/commit/490a330aa82cd0f1e65ca8615f5c58bd97aa0dda"><code>490a330</code></a> Update compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/16909">#16909</a>)</li>
<li><a href="https://github.com/babel/babel/commit/d65873827b00e4a0a3ed8fe59000cebd5d1dd82e"><code>d658738</code></a> fix: support BROWSERSLIST{,_CONFIG} env (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/16907">#16907</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.0/packages/babel-preset-env">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-react` from 7.25.7 to 7.26.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-react</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.26.3 (2024-12-04)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16958">#16958</a> [preserveFormat] force semicolons when invalidating ASI (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-builder-binary-assignment-operator-visitor</code>, <code>babel-plugin-transform-exponentiation-operator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16895">#16895</a> Remove helper-builder-binary-assignment-operator-visitor (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16959">#16959</a> perf: Reduce the use of temporary objects (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16923">#16923</a> perf: Improve scope information collection performance (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16964">#16964</a> perf: Avoid repeated traversal when creating scope (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-modules-commonjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16954">#16954</a> perf: Remove use of <code>simplifyAccess</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 4</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.26.2 (2024-10-30)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16903">#16903</a> fix: Parse placeholder for TS namespace (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16937">#16937</a> fix: Account for offsets when creating new Position instances (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16948">#16948</a> Fix mapping of tokens with generated nodes in between (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 6</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Dylan Piercey (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
<li>fisker Cheung (<a href="https://github.com/fisker"><code>@​fisker</code></a>)</li>
</ul>
<h2>v7.26.1 (2024-10-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16936">#16936</a> fix(parser): offset internal index locations by startIndex (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Dylan Piercey (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-react</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.26.3 (2024-12-04)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16958">#16958</a> [preserveFormat] force semicolons when invalidating ASI (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-builder-binary-assignment-operator-visitor</code>, <code>babel-plugin-transform-exponentiation-operator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16895">#16895</a> Remove helper-builder-binary-assignment-operator-visitor (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16959">#16959</a> perf: Reduce the use of temporary objects (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16923">#16923</a> perf: Improve scope information collection performance (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16964">#16964</a> perf: Avoid repeated traversal when creating scope (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-modules-commonjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16954">#16954</a> perf: Remove use of <code>simplifyAccess</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.2 (2024-10-30)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16903">#16903</a> fix: Parse placeholder for TS namespace (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16937">#16937</a> fix: Account for offsets when creating new Position instances (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16948">#16948</a> Fix mapping of tokens with generated nodes in between (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.1 (2024-10-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16936">#16936</a> fix(parser): offset internal index locations by startIndex (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.0 (2024-10-25)</h2>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-core</code>, <code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-syntax-import-assertions</code>, <code>babel-plugin-syntax-import-attributes</code>, <code>babel-preset-env</code>, <code>babel-standalone</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16850">#16850</a> Enable import attributes parsing by default (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16862">#16862</a> feat: support async plugin's pre/post (<a href="https://github.com/timofei-iatsenko"><code>@​timofei-iatsenko</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-proposal-regexp-modifiers</code>, <code>babel-plugin-transform-regexp-modifiers</code>, <code>babel-preset-env</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16692">#16692</a> Add <code>transform-regexp-modifiers</code> to <code>preset-env</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16849">#16849</a> feat: add <code>startIndex</code> parser option (<a href="https://github.com/DylanPiercey"><code>@​DylanPiercey</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-syntax-flow</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16841">#16841</a> Always enable parsing of Flow enums (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-preset-typescript</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16794">#16794</a> Support <code>import()</code> in <code>rewriteImportExtensions</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16708">#16708</a> Add experimental format-preserving mode to <code>@babel/generator</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/36ca8faeb52245fe268fb33b235f61ace073b2cc"><code>36ca8fa</code></a> v7.26.3</li>
<li><a href="https://github.com/babel/babel/commit/3ff105860880a6248dc8b0536d7ca42f00fc0a0a"><code>3ff1058</code></a> [react] Only default to <code>development</code> when set through env (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/16927">#16927</a>)</li>
<li><a href="https://github.com/babel/babel/commit/79e50b22bc9dda5703ebfa99180c7d9ff16b873d"><code>79e50b2</code></a> test: add node 23 to ci matrix</li>
<li><a href="https://github.com/babel/babel/commit/b07957ebb316a1e2fc67454fc7423508bb942e63"><code>b07957e</code></a> v7.25.9</li>
<li><a href="https://github.com/babel/babel/commit/af917594e4df3decdde2ce0b1614d607b27367a5"><code>af91759</code></a> fix: Accidentally publishing useless files (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/16917">#16917</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.3/packages/babel-preset-react">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/plugin-transform-runtime` from 7.25.7 to 7.25.9
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/plugin-transform-runtime</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.25.9 (2024-10-22)</h2>
<p>Thanks <a href="https://github.com/victorenator"><code>@​victorenator</code></a> for your first PR!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-template</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16905">#16905</a> fix: Keep type annotations in <code>syntacticPlaceholders</code> mode (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-compilation-targets</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16907">#16907</a> fix: support BROWSERSLIST{,_CONFIG} env (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16884">#16884</a> Analyze <code>ClassAccessorProperty</code> to prevent the <code>no-undef</code> rule (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-transform-fixture-test-runner</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16914">#16914</a> remove test options flaky (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><em>Every package</em>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16917">#16917</a> fix: Accidentally published <code>tsconfig</code> files (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16918">#16918</a> perf: Make <code>VISITOR_KEYS</code> etc. faster to access (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 4</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Viktar Vaŭčkievič (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.25.8 (2024-10-10)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16888">#16888</a> Restore public API of <code>resolvePlugin</code>/<code>resolvePreset</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-plugin-proposal-async-do-expressions</code>, <code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>, <code>babel-plugin-proposal-explicit-resource-management</code>, <code>babel-plugin-proposal-export-default-from</code>, <code>babel-plugin-proposal-function-bind</code>, <code>babel-plugin-proposal-function-sent</code>, <code>babel-plugin-proposal-import-defer</code>, <code>babel-plugin-proposal-partial-application</code>, <code>babel-plugin-proposal-throw-expressions</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-plugin-transform-dynamic-import</code>, <code>babel-plugin-transform-export-namespace-from</code>, <code>babel-plugin-transform-json-strings</code>, <code>babel-plugin-transform-logical-assignment-operators</code>, <code>babel-plugin-transform-nullish-coalescing-operator</code>, <code>babel-plugin-transform-numeric-separator</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-catch-binding</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-private-property-in-object</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16824">#16824</a> Inline one-line syntax plugins (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/plugin-transform-runtime</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.25.9 (2024-10-22)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-template</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16905">#16905</a> fix: Keep type annotations in <code>syntacticPlaceholders</code> mode (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-compilation-targets</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16907">#16907</a> fix: support BROWSERSLIST{,_CONFIG} env (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16884">#16884</a> Analyze <code>ClassAccessorProperty</code> to prevent the <code>no-undef</code> rule (<a href="https://github.com/victorenator"><code>@​victorenator</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-helper-transform-fixture-test-runner</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16914">#16914</a> remove test options flaky (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16918">#16918</a> perf: Make <code>VISITOR_KEYS</code> etc. faster to access (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.25.8 (2024-10-10)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16888">#16888</a> Restore public API of <code>resolvePlugin</code>/<code>resolvePreset</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-plugin-proposal-async-do-expressions</code>, <code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>, <code>babel-plugin-proposal-explicit-resource-management</code>, <code>babel-plugin-proposal-export-default-from</code>, <code>babel-plugin-proposal-function-bind</code>, <code>babel-plugin-proposal-function-sent</code>, <code>babel-plugin-proposal-import-defer</code>, <code>babel-plugin-proposal-partial-application</code>, <code>babel-plugin-proposal-throw-expressions</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-plugin-transform-dynamic-import</code>, <code>babel-plugin-transform-export-namespace-from</code>, <code>babel-plugin-transform-json-strings</code>, <code>babel-plugin-transform-logical-assignment-operators</code>, <code>babel-plugin-transform-nullish-coalescing-operator</code>, <code>babel-plugin-transform-numeric-separator</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-catch-binding</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-private-property-in-object</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16824">#16824</a> Inline one-line syntax plugins (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/b07957ebb316a1e2fc67454fc7423508bb942e63"><code>b07957e</code></a> v7.25.9</li>
<li><a href="https://github.com/babel/babel/commit/af917594e4df3decdde2ce0b1614d607b27367a5"><code>af91759</code></a> fix: Accidentally publishing useless files (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/16917">#16917</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.25.9/packages/babel-plugin-transform-runtime">compare view</a></li>
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

### Comment by @codecov-commenter on January 08, 2025 at 11:49 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1234?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.60%. Comparing base [(`b390c17`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b390c17e5ab1495dc6567876b50adc6abb8543c6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`1ae2b6c`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1ae2b6c6e7776cd3c1c7bbe2360a91ee9199fe42?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1234   +/-   ##
=======================================
  Coverage   63.60%   63.60%           
=======================================
  Files         124      124           
  Lines        3267     3267           
  Branches      860      860           
=======================================
  Hits         2078     2078           
  Misses       1189     1189           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1234?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @dependabot on January 13, 2025 at 06:58 PM UTC

Superseded by #1243.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1234*
