---
type: pull_request
number: 1283
title: "chore(deps): bump the babel group with 3 updates"
state: closed
author: dependabot
created: 2025-02-10T19:05:00Z
updated: 2025-02-14T13:48:22Z
closed: 2025-02-14T13:48:21Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-41178f5fee
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1283
---

# Pull Request #1283: chore(deps): bump the babel group with 3 updates

**Author**: @dependabot
**Created**: February 10, 2025 at 07:05 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-41178f5fee`

## Description

Bumps the babel group with 3 updates: [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser), [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime) and [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env).

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
</blockquote>
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

Updates `@babel/plugin-transform-runtime` from 7.25.9 to 7.26.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/plugin-transform-runtime</code>'s releases</a>.</em></p>
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
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/plugin-transform-runtime</code>'s changelog</a>.</em></p>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/059394159de3af74446be6fba896dd4ce178333c"><code>0593941</code></a> v7.26.8</li>
<li><a href="https://github.com/babel/babel/commit/1137c29a30badc6e51d58ab21983ba2824140e5b"><code>1137c29</code></a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17097">#17097</a>)</li>
<li><a href="https://github.com/babel/babel/commit/cd24cc07ef6558b7f6510f9177f6393c91b0549f"><code>cd24cc0</code></a> chore: Update TS 5.7 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17053">#17053</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.8/packages/babel-plugin-transform-runtime">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-env` from 7.26.7 to 7.26.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-env</code>'s releases</a>.</em></p>
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
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-env</code>'s changelog</a>.</em></p>
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
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/059394159de3af74446be6fba896dd4ce178333c"><code>0593941</code></a> v7.26.8</li>
<li><a href="https://github.com/babel/babel/commit/1137c29a30badc6e51d58ab21983ba2824140e5b"><code>1137c29</code></a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17097">#17097</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.26.8/packages/babel-preset-env">compare view</a></li>
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

### Comment by @codecov-commenter on February 10, 2025 at 07:11 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1283?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 63.65%. Comparing base [(`85dba6b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/85dba6b5a340e3d30fedb9c1ced657c52cccb546?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5b2cd1b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/5b2cd1b5e8630b5cbf1feedcb741b735144263a7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1283   +/-   ##
=======================================
  Coverage   63.65%   63.65%           
=======================================
  Files         124      124           
  Lines        3282     3282           
  Branches      868      868           
=======================================
  Hits         2089     2089           
  Misses       1193     1193           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1283?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @dependabot on February 14, 2025 at 01:48 PM UTC

Superseded by #1287.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1283*
