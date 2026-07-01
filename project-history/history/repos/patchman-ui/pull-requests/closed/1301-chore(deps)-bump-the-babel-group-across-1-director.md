---
type: pull_request
number: 1301
title: "chore(deps): bump the babel group across 1 directory with 4 updates"
state: closed
author: dependabot
created: 2025-03-03T18:44:50Z
updated: 2025-05-05T18:43:59Z
closed: 2025-05-05T18:43:57Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-b933073bd2
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1301
---

# Pull Request #1301: chore(deps): bump the babel group across 1 directory with 4 updates

**Author**: @dependabot
**Created**: March 03, 2025 at 06:44 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-b933073bd2`

## Description

Bumps the babel group with 4 updates in the / directory: [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser), [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime), [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env) and [babel-loader](https://github.com/babel/babel-loader).

Updates `@babel/eslint-parser` from 7.26.5 to 7.26.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/eslint-parser</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.26.8 (2025-02-08)</h2>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17097">#17097</a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0</li>
</ul>
</li>
</ul>
<h2>v7.26.7 (2025-01-24)</h2>
<p>Thanks <a href="https://github.com/branchseer"><code>@​branchseer</code></a> and <a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a> for your first PRs!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17086">#17086</a> Make &quot;object without properties&quot; helpers ES6-compatible (<a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typeof-symbol</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17085">#17085</a> fix: Correctly handle <code>typeof</code> in arrow functions (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17079">#17079</a> Respect <code>ranges</code> option in estree method value (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17052">#17052</a> Do not try to parse .ts configs as JSON if natively supported (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17050">#17050</a> fix: correctly resolve references to non-constant enum members (<a href="https://github.com/branchseer"><code>@​branchseer</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17025">#17025</a> fix: Remove type-only <code>import x = y.z</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 6</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li>Tony Quetano (<a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a>)</li>
<li><a href="https://github.com/branchseer"><code>@​branchseer</code></a></li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.26.6 (2025-01-13)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-nullish-coalescing-operator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17061">#17061</a> fix: Chaining nullish coalescing operators output size regression (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 1</h4>
<ul>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/eslint-parser</code>'s changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
<blockquote>
<p><strong>Tags:</strong></p>
<ul>
<li>:boom: [Breaking Change]</li>
<li>:eyeglasses: [Spec Compliance]</li>
<li>:rocket: [New Feature]</li>
<li>:bug: [Bug Fix]</li>
<li>:memo: [Documentation]</li>
<li>:house: [Internal]</li>
<li>:nail_care: [Polish]</li>
</ul>
</blockquote>
<p><em>Note: Gaps between patch versions are faulty, broken or test releases.</em></p>
<p>This file contains the changelog starting from v7.15.0.</p>
<ul>
<li>See <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-v7.0.0-v7.14.9.md">CHANGELOG - v7.0.0 to v7.14.9</a> for v7.0.0 to v7.14.9 changes.</li>
<li>See <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-v7-prereleases.md">CHANGELOG - v7 prereleases</a> for v7.0.0-alpha.1 to v7.0.0-rc.4 changes.</li>
<li>See <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-v4.md">CHANGELOG - v4</a>, <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-v5.md">CHANGELOG - v5</a>, and <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-v6.md">CHANGELOG - v6</a> for v4.x-v6.x changes.</li>
<li>See <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-6to5.md">CHANGELOG - 6to5</a> for the pre-4.0.0 version changelog.</li>
<li>See <a href="https://github.com/babel/babel/blob/main/packages/babel-parser/CHANGELOG.md">Babylon's CHANGELOG</a> for the Babylon pre-7.0.0-beta.29 version changelog.</li>
<li>See <a href="https://github.com/babel/babel-eslint/releases"><code>babel-eslint</code>'s releases</a> for the changelog before <code>@babel/eslint-parser</code> 7.8.0.</li>
<li>See <a href="https://github.com/babel/eslint-plugin-babel/releases"><code>eslint-plugin-babel</code>'s releases</a> for the changelog before <code>@babel/eslint-plugin</code> 7.8.0.</li>
</ul>
<!-- raw HTML omitted -->
<!-- raw HTML omitted -->
<h2>v7.26.9 (2025-02-14)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17103">#17103</a> fix: Definition for <code>TSPropertySignature.kind</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17062">#17062</a> Print TypeScript optional/definite in ClassPrivateProperty (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17130">#17130</a> Use <code>.ts</code> files with explicit reexports to solve name conflicts (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17127">#17127</a> Do not depend on <code>@types/gensync</code> in Babel 7 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.7 (2025-01-24)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17086">#17086</a> Make &quot;object without properties&quot; helpers ES6-compatible (<a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typeof-symbol</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17085">#17085</a> fix: Correctly handle <code>typeof</code> in arrow functions (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17079">#17079</a> Respect <code>ranges</code> option in estree method value (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/059394159de3af74446be6fba896dd4ce178333c"><code>0593941</code></a> v7.26.8</li>
<li><a href="https://github.com/babel/babel/commit/e02b0ffd871ab07a47947956f39698346b200b1f"><code>e02b0ff</code></a> [Babel 8] Create TSTemplateLiteralType (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17066">#17066</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.8/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/plugin-transform-runtime` from 7.25.9 to 7.26.9
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/plugin-transform-runtime</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.26.9 (2025-02-14)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17103">#17103</a> fix: Definition for <code>TSPropertySignature.kind</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17062">#17062</a> Print TypeScript optional/definite in ClassPrivateProperty (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17130">#17130</a> Use <code>.ts</code> files with explicit reexports to solve name conflicts (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17127">#17127</a> Do not depend on <code>@types/gensync</code> in Babel 7 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Jamie Kyle (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.26.8 (2025-02-08)</h2>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17097">#17097</a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0</li>
</ul>
</li>
</ul>
<h2>v7.26.7 (2025-01-24)</h2>
<p>Thanks <a href="https://github.com/branchseer"><code>@​branchseer</code></a> and <a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a> for your first PRs!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17086">#17086</a> Make &quot;object without properties&quot; helpers ES6-compatible (<a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typeof-symbol</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17085">#17085</a> fix: Correctly handle <code>typeof</code> in arrow functions (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17079">#17079</a> Respect <code>ranges</code> option in estree method value (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17052">#17052</a> Do not try to parse .ts configs as JSON if natively supported (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17050">#17050</a> fix: correctly resolve references to non-constant enum members (<a href="https://github.com/branchseer"><code>@​branchseer</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17025">#17025</a> fix: Remove type-only <code>import x = y.z</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 6</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li>Tony Quetano (<a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/plugin-transform-runtime</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.26.9 (2025-02-14)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17103">#17103</a> fix: Definition for <code>TSPropertySignature.kind</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17062">#17062</a> Print TypeScript optional/definite in ClassPrivateProperty (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17130">#17130</a> Use <code>.ts</code> files with explicit reexports to solve name conflicts (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17127">#17127</a> Do not depend on <code>@types/gensync</code> in Babel 7 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.7 (2025-01-24)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17086">#17086</a> Make &quot;object without properties&quot; helpers ES6-compatible (<a href="https://github.com/tquetano-netflix"><code>@​tquetano-netflix</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typeof-symbol</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17085">#17085</a> fix: Correctly handle <code>typeof</code> in arrow functions (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17079">#17079</a> Respect <code>ranges</code> option in estree method value (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17052">#17052</a> Do not try to parse .ts configs as JSON if natively supported (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17050">#17050</a> fix: correctly resolve references to non-constant enum members (<a href="https://github.com/branchseer"><code>@​branchseer</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17025">#17025</a> fix: Remove type-only <code>import x = y.z</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.6 (2025-01-13)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-nullish-coalescing-operator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17061">#17061</a> fix: Chaining nullish coalescing operators output size regression (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.5 (2025-01-10)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17011">#17011</a> Allow the dynamic <code>import.defer()</code> form of <code>import defer</code> (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-block-scoped-functions</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17024">#17024</a> chore: Avoid calling <code>isInStrictMode</code> in Babel 7 (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17026">#17026</a> fix: Correctly generate exported const enums in namespace (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17045">#17045</a> [estree] Unify method type parameters handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17013">#17013</a> fix: Correctly set position for <code>@(a.b)()</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/16996">#16996</a> [estree] Adjust the start loc of class methods with type params (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>, <code>babel-plugin-transform-flow-strip-types</code>, <code>babel-types</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/64bca7b5f308cd52c192a5c821a57f6d1b0475f4"><code>64bca7b</code></a> v7.26.9</li>
<li><a href="https://github.com/babel/babel/commit/5315446a9a9348bd9646464bfa7b0e4381647e04"><code>5315446</code></a> [babel 8] Remove babel 7-specific imports (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17111">#17111</a>)</li>
<li><a href="https://github.com/babel/babel/commit/059394159de3af74446be6fba896dd4ce178333c"><code>0593941</code></a> v7.26.8</li>
<li><a href="https://github.com/babel/babel/commit/1137c29a30badc6e51d58ab21983ba2824140e5b"><code>1137c29</code></a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17097">#17097</a>)</li>
<li><a href="https://github.com/babel/babel/commit/cd24cc07ef6558b7f6510f9177f6393c91b0549f"><code>cd24cc0</code></a> chore: Update TS 5.7 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17053">#17053</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.9/packages/babel-plugin-transform-runtime">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-env` from 7.26.7 to 7.26.9
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-env</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.26.9 (2025-02-14)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17103">#17103</a> fix: Definition for <code>TSPropertySignature.kind</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17062">#17062</a> Print TypeScript optional/definite in ClassPrivateProperty (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17130">#17130</a> Use <code>.ts</code> files with explicit reexports to solve name conflicts (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17127">#17127</a> Do not depend on <code>@types/gensync</code> in Babel 7 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Jamie Kyle (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.26.8 (2025-02-08)</h2>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17097">#17097</a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-env</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.26.9 (2025-02-14)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17103">#17103</a> fix: Definition for <code>TSPropertySignature.kind</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17062">#17062</a> Print TypeScript optional/definite in ClassPrivateProperty (<a href="https://github.com/jamiebuilds-signal"><code>@​jamiebuilds-signal</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17130">#17130</a> Use <code>.ts</code> files with explicit reexports to solve name conflicts (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17127">#17127</a> Do not depend on <code>@types/gensync</code> in Babel 7 (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/64bca7b5f308cd52c192a5c821a57f6d1b0475f4"><code>64bca7b</code></a> v7.26.9</li>
<li><a href="https://github.com/babel/babel/commit/5315446a9a9348bd9646464bfa7b0e4381647e04"><code>5315446</code></a> [babel 8] Remove babel 7-specific imports (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17111">#17111</a>)</li>
<li><a href="https://github.com/babel/babel/commit/059394159de3af74446be6fba896dd4ce178333c"><code>0593941</code></a> v7.26.8</li>
<li><a href="https://github.com/babel/babel/commit/1137c29a30badc6e51d58ab21983ba2824140e5b"><code>1137c29</code></a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17097">#17097</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.9/packages/babel-preset-env">compare view</a></li>
</ul>
</details>
<br />

Updates `babel-loader` from 9.2.1 to 10.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel-loader/releases">babel-loader's releases</a>.</em></p>
<blockquote>
<h2>v10.0.0</h2>
<h2>What's Changed</h2>
<h3>Breaking Changes</h3>
<ul>
<li>bump node requirement to <code>^18.20.0 || ^20.10.0 || &gt;=22.0.0</code> and webpack requirement to <code>&gt;= 5.61.0</code> by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1026">babel/babel-loader#1026</a></li>
<li>breaking: use output.hashFunction as loader cache hasher by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1027">babel/babel-loader#1027</a></li>
</ul>
<h3>New Features</h3>
<ul>
<li>Add babel-loader logger by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1034">babel/babel-loader#1034</a></li>
<li>Support cache with external dependencies by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1033">babel/babel-loader#1033</a></li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>[Bugfix] Ensure stability of filename cache-keys by <a href="https://github.com/stefanpenner"><code>@​stefanpenner</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/909">babel/babel-loader#909</a></li>
</ul>
<h3>Docs</h3>
<ul>
<li>docs: clarify that <code>cacheIdentifier</code> is computed from the merged options by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1000">babel/babel-loader#1000</a></li>
<li>Create SECURITY.md by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1032">babel/babel-loader#1032</a></li>
<li>Add babel-loader v10 readme by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1046">babel/babel-loader#1046</a></li>
<li>add readme section for loggingDebug support by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1038">babel/babel-loader#1038</a></li>
<li>Update readme and repo templates by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1041">babel/babel-loader#1041</a></li>
</ul>
<h3>Dependencies</h3>
<ul>
<li>migrate to c8 by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/996">babel/babel-loader#996</a></li>
<li>Bump word-wrap from 1.2.3 to 1.2.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/998">babel/babel-loader#998</a></li>
<li>Bump dev dependencies by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1001">babel/babel-loader#1001</a></li>
<li>Bump braces from 3.0.2 to 3.0.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1020">babel/babel-loader#1020</a></li>
<li>Update deps by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1025">babel/babel-loader#1025</a></li>
<li>refactor: replace <code>find-cache-dir</code> by <code>find-up</code> by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1031">babel/babel-loader#1031</a></li>
<li>Bump webpack from 5.93.0 to 5.94.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1035">babel/babel-loader#1035</a></li>
<li>chore: update dev deps by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1036">babel/babel-loader#1036</a></li>
<li>Bump cross-spawn from 7.0.3 to 7.0.6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1049">babel/babel-loader#1049</a></li>
</ul>
<h3>Internal</h3>
<ul>
<li>Remove caller option check by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1007">babel/babel-loader#1007</a></li>
<li>Update tests by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1003">babel/babel-loader#1003</a></li>
<li>Update metadata test by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1024">babel/babel-loader#1024</a></li>
<li>Migrate to node test runner by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1028">babel/babel-loader#1028</a></li>
<li>chore: use default eslint rules by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1029">babel/babel-loader#1029</a></li>
<li>refactor: use webpack builtin schema util by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1030">babel/babel-loader#1030</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/stefanpenner"><code>@​stefanpenner</code></a> made their first contribution in <a href="https://redirect.github.com/babel/babel-loader/pull/909">babel/babel-loader#909</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/babel/babel-loader/compare/v9.1.3...v10.0.0">https://github.com/babel/babel-loader/compare/v9.1.3...v10.0.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel-loader/commit/10456d3f196a9be611d42cececec6a624cbec2e3"><code>10456d3</code></a> 10.0.0</li>
<li><a href="https://github.com/babel/babel-loader/commit/5a223cf039aee8e5e32c2fce4a4bc5ca20bbd901"><code>5a223cf</code></a> Add babel-loader v10 readme (<a href="https://redirect.github.com/babel/babel-loader/issues/1046">#1046</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/8f8866700fb4bf7b4fe0fde28fa56932d3760733"><code>8f88667</code></a> Bump cross-spawn from 7.0.3 to 7.0.6 (<a href="https://redirect.github.com/babel/babel-loader/issues/1049">#1049</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/f765949ecec8ba7bfdc06c0447bc92d8e0f34422"><code>f765949</code></a> Update readme and repo templates (<a href="https://redirect.github.com/babel/babel-loader/issues/1041">#1041</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/b582028d51b94ae97713aae5ce6f8cb81c7ed4fb"><code>b582028</code></a> add readme section for loggingDebug support (<a href="https://redirect.github.com/babel/babel-loader/issues/1038">#1038</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/a0c450dde1e4af045e743ae2dd3e990dabede4f7"><code>a0c450d</code></a> feat: add babel-loader debug logger (<a href="https://redirect.github.com/babel/babel-loader/issues/1034">#1034</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/d4181b8e20c2421f5f5e44fca14e944866b47a52"><code>d4181b8</code></a> Support cache with external dependencies (<a href="https://redirect.github.com/babel/babel-loader/issues/1033">#1033</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/7fcb53393677f2d2744a428663a95c7e89ab78b0"><code>7fcb533</code></a> chore: update dev deps (<a href="https://redirect.github.com/babel/babel-loader/issues/1036">#1036</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/c2a90e5de5e70c9ce2382dd08bccbb9586059fda"><code>c2a90e5</code></a> Bump webpack from 5.93.0 to 5.94.0 (<a href="https://redirect.github.com/babel/babel-loader/issues/1035">#1035</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/70a3710dbefb75c92831ce557199438fb66874ae"><code>70a3710</code></a> refactor: replace find-cache-dir by find-up (<a href="https://redirect.github.com/babel/babel-loader/issues/1031">#1031</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel-loader/compare/v9.2.1...v10.0.0">compare view</a></li>
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

### Comment by @dependabot on May 05, 2025 at 06:43 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1301*
