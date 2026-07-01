---
type: pull_request
number: 1344
title: "chore(deps): bump the babel group across 1 directory with 6 updates"
state: closed
author: dependabot
created: 2025-06-16T21:12:42Z
updated: 2025-09-22T18:02:22Z
closed: 2025-09-22T18:02:20Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-5888c9a8f1
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1344
---

# Pull Request #1344: chore(deps): bump the babel group across 1 directory with 6 updates

**Author**: @dependabot
**Created**: June 16, 2025 at 09:12 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-5888c9a8f1`

## Description

Bumps the babel group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser) | `7.26.5` | `7.27.5` |
| [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime) | `7.25.9` | `7.27.4` |
| [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env) | `7.26.7` | `7.27.2` |
| [@babel/preset-react](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react) | `7.26.3` | `7.27.1` |
| [babel-jest](https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest) | `29.7.0` | `30.0.0` |
| [babel-loader](https://github.com/babel/babel-loader) | `9.2.1` | `10.0.0` |


Updates `@babel/eslint-parser` from 7.26.5 to 7.27.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/eslint-parser</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.27.5 (2025-06-03)</h2>
<p>Thanks <a href="https://github.com/NullVoxPopuli"><code>@​NullVoxPopuli</code></a> for your first PR!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17359">#17359</a> fix: Unexpected infinite loop with <code>regenerator</code> for <code>try</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17349">#17349</a> Map ESLint's <code>sourceType: commonjs</code> to script (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17333">#17333</a> Improve using declaration errors (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 4</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://github.com/NullVoxPopuli"><code>@​NullVoxPopuli</code></a></li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.27.4 (2025-05-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17323">#17323</a> Disallow using in bare case statement (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17311">#17311</a> Improve parseExpression error messages (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:microscope: Output optimization</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17287">#17287</a> Reduce <code>regenerator</code> size more (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17334">#17334</a> Use shorter method names for regenerator context (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17268">#17268</a> Reduce <code>regenerator</code> helper size (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17238">#17238</a> Split <code>regeneratorRuntime</code> into multiple helpers (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
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
<h2>v7.27.3 (2025-05-27)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17324">#17324</a> Improve multiline comments handling in yield/await expression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/eslint-parser</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.27.5 (2025-06-03)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17359">#17359</a> fix: Unexpected infinite loop with <code>regenerator</code> for <code>try</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17349">#17349</a> Map ESLint's <code>sourceType: commonjs</code> to script (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17333">#17333</a> Improve using declaration errors (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.27.4 (2025-05-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17323">#17323</a> Disallow using in bare case statement (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17311">#17311</a> Improve parseExpression error messages (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:microscope: Output optimization</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17287">#17287</a> Reduce <code>regenerator</code> size more (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17334">#17334</a> Use shorter method names for regenerator context (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17268">#17268</a> Reduce <code>regenerator</code> helper size (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17238">#17238</a> Split <code>regeneratorRuntime</code> into multiple helpers (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.27.3 (2025-05-27)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17324">#17324</a> Improve multiline comments handling in yield/await expression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17328">#17328</a> Correctly set <code>.displayName</code> on <code>GeneratorFunction</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17319">#17319</a> fix: handle shadowed binding in <code>for using of</code> body (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17317">#17317</a> fix: support named evaluation for using declaration (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-decorators</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17321">#17321</a> fix(converter): Remove <code>abstract</code> modifiers in class declaration to expression conversion (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-module-transforms</code>, <code>babel-plugin-proposal-explicit-resource-management</code>, <code>babel-plugin-transform-modules-amd</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-modules-umd</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17257">#17257</a> Preserve class id when transforming using declarations with exported class (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17312">#17312</a> fix(parser): properly handle optional markers in generator class methods (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17307">#17307</a> fix(parser): Terminate modifier parsing at newline (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17308">#17308</a> Improve import phase parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.27.2 (2025-05-06)</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/746ae8e7960b060e475485ad8c7725f6c59a2c54"><code>746ae8e</code></a> v7.27.5</li>
<li><a href="https://github.com/babel/babel/commit/76f3c8588c921c6d4d9ae4855bad3929f169cb64"><code>76f3c85</code></a> fix: support async Babel plugin in eslint-parser (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17348">#17348</a>)</li>
<li><a href="https://github.com/babel/babel/commit/0c51d8f213c98991a127552bb02e76c96e9a1f1b"><code>0c51d8f</code></a> Map ESLint's <code>sourceType: commonjs</code> to script (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17349">#17349</a>)</li>
<li><a href="https://github.com/babel/babel/commit/2d0c76e6fa8fd2331f6cef9eba25f73dab6357bb"><code>2d0c76e</code></a> Improve integrations of using declaration with other transforms (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17330">#17330</a>)</li>
<li><a href="https://github.com/babel/babel/commit/eebd3a06021c13d335b5b0bd79734df3abbea678"><code>eebd3a0</code></a> v7.27.1</li>
<li><a href="https://github.com/babel/babel/commit/9a40d852e7bcf65cfadfaaacf0b85d66a0b59d6e"><code>9a40d85</code></a> [Babel 8]: Remove record and tuple syntax support (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17242">#17242</a>)</li>
<li><a href="https://github.com/babel/babel/commit/3766c4ddfaaf5429eb8443f42d172e5845a0feef"><code>3766c4d</code></a> Fill optional AST properties when both estree and typescript parser plugin ar...</li>
<li><a href="https://github.com/babel/babel/commit/0de0e2b0a8a4c6fec519bf1882e32ed49424129d"><code>0de0e2b</code></a> Bump typescript-eslint to 8.29.1 (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17232">#17232</a>)</li>
<li><a href="https://github.com/babel/babel/commit/c1614c804c52e370dd48a3aba59bb17eb24dfbf8"><code>c1614c8</code></a> Fill optional AST properties when both estree and typescript parser plugin ar...</li>
<li><a href="https://github.com/babel/babel/commit/7979e3fc56cfba9d62b4dfa4ef71cfc025a8863f"><code>7979e3f</code></a> Fill optional AST properties when both estree and typescript parser plugin ar...</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel/commits/v7.27.5/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/plugin-transform-runtime` from 7.25.9 to 7.27.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/plugin-transform-runtime</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.27.4 (2025-05-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17323">#17323</a> Disallow using in bare case statement (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17311">#17311</a> Improve parseExpression error messages (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:microscope: Output optimization</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17287">#17287</a> Reduce <code>regenerator</code> size more (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17334">#17334</a> Use shorter method names for regenerator context (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17268">#17268</a> Reduce <code>regenerator</code> helper size (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17238">#17238</a> Split <code>regeneratorRuntime</code> into multiple helpers (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
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
<h2>v7.27.3 (2025-05-27)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17324">#17324</a> Improve multiline comments handling in yield/await expression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17328">#17328</a> Correctly set <code>.displayName</code> on <code>GeneratorFunction</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17319">#17319</a> fix: handle shadowed binding in <code>for using of</code> body (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17317">#17317</a> fix: support named evaluation for using declaration (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-decorators</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17321">#17321</a> fix(converter): Remove <code>abstract</code> modifiers in class declaration to expression conversion (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-module-transforms</code>, <code>babel-plugin-proposal-explicit-resource-management</code>, <code>babel-plugin-transform-modules-amd</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-modules-umd</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17257">#17257</a> Preserve class id when transforming using declarations with exported class (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17312">#17312</a> fix(parser): properly handle optional markers in generator class methods (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17307">#17307</a> fix(parser): Terminate modifier parsing at newline (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17308">#17308</a> Improve import phase parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 7</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li>Vik R (<a href="https://github.com/vikr01"><code>@​vikr01</code></a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/plugin-transform-runtime</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.27.4 (2025-05-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17323">#17323</a> Disallow using in bare case statement (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17311">#17311</a> Improve parseExpression error messages (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:microscope: Output optimization</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17287">#17287</a> Reduce <code>regenerator</code> size more (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17334">#17334</a> Use shorter method names for regenerator context (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17268">#17268</a> Reduce <code>regenerator</code> helper size (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-classes</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17238">#17238</a> Split <code>regeneratorRuntime</code> into multiple helpers (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.27.3 (2025-05-27)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17324">#17324</a> Improve multiline comments handling in yield/await expression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17328">#17328</a> Correctly set <code>.displayName</code> on <code>GeneratorFunction</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-explicit-resource-management</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17319">#17319</a> fix: handle shadowed binding in <code>for using of</code> body (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17317">#17317</a> fix: support named evaluation for using declaration (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-decorators</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17321">#17321</a> fix(converter): Remove <code>abstract</code> modifiers in class declaration to expression conversion (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-module-transforms</code>, <code>babel-plugin-proposal-explicit-resource-management</code>, <code>babel-plugin-transform-modules-amd</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-modules-umd</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17257">#17257</a> Preserve class id when transforming using declarations with exported class (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17312">#17312</a> fix(parser): properly handle optional markers in generator class methods (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17307">#17307</a> fix(parser): Terminate modifier parsing at newline (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17308">#17308</a> Improve import phase parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.27.2 (2025-05-06)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17289">#17289</a> fix: <code>@babel/parser/bin/index.js</code> contains <code>node:</code> protocol require (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17291">#17291</a> fix: Private class method not found when TS and estree (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-object-rest-spread</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17281">#17281</a> Fix: improve object rest handling in array pattern (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-modules-commonjs</code>, <code>babel-template</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17284">#17284</a> fix(babel-template): Properly handle empty string replacements (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-cli</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/7d069309fdfcedda2928a043f6f7c98135c1242a"><code>7d06930</code></a> v7.27.4</li>
<li><a href="https://github.com/babel/babel/commit/5b9468d9bf1ab4f427241673e9f03593da115a69"><code>5b9468d</code></a> Reduce <code>regenerator</code> size more (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17287">#17287</a>)</li>
<li><a href="https://github.com/babel/babel/commit/d23a1bd4c1f46762eb34888a1a30fa6f77e2f7cf"><code>d23a1bd</code></a> Use shorter method names for regenerator context (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17334">#17334</a>)</li>
<li><a href="https://github.com/babel/babel/commit/fe32019663f91c2e9aa1fbd4fe96bdc2ae27bf3b"><code>fe32019</code></a> Reduce <code>regenerator</code> helper size (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17268">#17268</a>)</li>
<li><a href="https://github.com/babel/babel/commit/a0690e39ea63cdcc3d9282ece739e6677c83ad6e"><code>a0690e3</code></a> Split <code>regeneratorRuntime</code> into multiple helpers (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17238">#17238</a>)</li>
<li><a href="https://github.com/babel/babel/commit/da5e371efabf6c0baab1ec2c888da189e1b610ad"><code>da5e371</code></a> v7.27.3</li>
<li><a href="https://github.com/babel/babel/commit/4b76a5f2c1ae0b8ea308c1687f86cc8cdbaffc88"><code>4b76a5f</code></a> Enable <code>dot-notation</code> rule (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17301">#17301</a>)</li>
<li><a href="https://github.com/babel/babel/commit/eebd3a06021c13d335b5b0bd79734df3abbea678"><code>eebd3a0</code></a> v7.27.1</li>
<li><a href="https://github.com/babel/babel/commit/317e332e650bc04907bc787ab79f930288a3e71e"><code>317e332</code></a> Enforce node protocol import (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17207">#17207</a>)</li>
<li><a href="https://github.com/babel/babel/commit/fdc0fb59e119ee0b38bced63867a344a5b4bc2f3"><code>fdc0fb5</code></a> [Babel 8] Bump nodejs requirements to <code>^20.19.0 || &gt;= 22.12.0</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17204">#17204</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel/commits/v7.27.4/packages/babel-plugin-transform-runtime">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-env` from 7.26.7 to 7.27.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-env</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.27.2 (2025-05-06)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17289">#17289</a> fix: <code>@babel/parser/bin/index.js</code> contains <code>node:</code> protocol require (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17291">#17291</a> fix: Private class method not found when TS and estree (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-object-rest-spread</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17281">#17281</a> Fix: improve object rest handling in array pattern (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-modules-commonjs</code>, <code>babel-template</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17284">#17284</a> fix(babel-template): Properly handle empty string replacements (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-cli</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17285">#17285</a> Enable Node compile cache for <code>@babel/cli</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
<li><a href="https://github.com/magic-akari"><code>@​magic-akari</code></a></li>
</ul>
<h2>v7.27.1 (2025-04-30)</h2>
<p>Thanks <a href="https://github.com/kermanx"><code>@​kermanx</code></a> and <a href="https://github.com/woaitsAryan"><code>@​woaitsAryan</code></a> for your first PRs!</p>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17254">#17254</a> Allow <code>using of</code> as lexical declaration within for (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17230">#17230</a> Disallow get/set in TSPropertySignature (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17193">#17193</a> Stricter TSImportType options parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17137">#17137</a> fix: do expressions should allow early exit (<a href="https://github.com/kermanx"><code>@​kermanx</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-wrap-function</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17251">#17251</a> Fix: propagate argument evaluation errors through async promise chain (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-remap-async-to-generator</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17231">#17231</a> fix apply()/call() annotated as pure (<a href="https://github.com/Lacsw"><code>@​Lacsw</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-fixtures</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17233">#17233</a> Create ChainExpression within TSInstantiationExpression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17226">#17226</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 2) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17224">#17224</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 1) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17080">#17080</a> Fix start of TSParameterProperty (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17228">#17228</a> Update firefox bugfix compat data (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-env</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.27.2 (2025-05-06)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17289">#17289</a> fix: <code>@babel/parser/bin/index.js</code> contains <code>node:</code> protocol require (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17291">#17291</a> fix: Private class method not found when TS and estree (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-object-rest-spread</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17281">#17281</a> Fix: improve object rest handling in array pattern (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-modules-commonjs</code>, <code>babel-template</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17284">#17284</a> fix(babel-template): Properly handle empty string replacements (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-cli</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17285">#17285</a> Enable Node compile cache for <code>@babel/cli</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.27.1 (2025-04-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17254">#17254</a> Allow <code>using of</code> as lexical declaration within for (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17230">#17230</a> Disallow get/set in TSPropertySignature (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17193">#17193</a> Stricter TSImportType options parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17137">#17137</a> fix: do expressions should allow early exit (<a href="https://github.com/kermanx"><code>@​kermanx</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-wrap-function</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17251">#17251</a> Fix: propagate argument evaluation errors through async promise chain (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-remap-async-to-generator</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17231">#17231</a> fix apply()/call() annotated as pure (<a href="https://github.com/Lacsw"><code>@​Lacsw</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-fixtures</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17233">#17233</a> Create ChainExpression within TSInstantiationExpression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17226">#17226</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 2) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17224">#17224</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 1) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17080">#17080</a> Fix start of TSParameterProperty (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17228">#17228</a> Update firefox bugfix compat data (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17156">#17156</a> fix: Objects and arrays with multiple references should not be evaluated (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17216">#17216</a> Fix: support const type parameter in generator (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-bugfix-v8-spread-parameters-in-optional-chaining</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-transform-arrow-functions</code>, <code>babel-plugin-transform-class-properties</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-parameters</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17221">#17221</a> Reduce generated names size for the 10th-11th (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/a8080cd196b381eee2b992956728e4e936f29b3f"><code>a8080cd</code></a> v7.27.2</li>
<li><a href="https://github.com/babel/babel/commit/b9c6efbac3d9f7b158933e1ac89a4630b2d4682a"><code>b9c6efb</code></a> Use <code>.d.ts</code> in import when importing a <code>.d.ts</code> file (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17288">#17288</a>)</li>
<li><a href="https://github.com/babel/babel/commit/1e0bdac593260b5f3e9ab192d1378102c7d4858f"><code>1e0bdac</code></a> [Babel 8] Directly export the JSON files from <code>@babel/compat-data</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17267">#17267</a>)</li>
<li><a href="https://github.com/babel/babel/commit/eebd3a06021c13d335b5b0bd79734df3abbea678"><code>eebd3a0</code></a> v7.27.1</li>
<li><a href="https://github.com/babel/babel/commit/5c5e77c05b52f4d333c4910014efe700022d8cef"><code>5c5e77c</code></a> Tune plugin compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17256">#17256</a>)</li>
<li><a href="https://github.com/babel/babel/commit/aca4f0138c4b163f39fb398e862dc4400ae9c986"><code>aca4f01</code></a> chore: bump compat-data sources (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17253">#17253</a>)</li>
<li><a href="https://github.com/babel/babel/commit/726438499badd1454911073ddba20c8acfaa3781"><code>7264384</code></a> Update firefox bugfix compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17228">#17228</a>)</li>
<li><a href="https://github.com/babel/babel/commit/0f95b748a9a5f90c3b23d4d72299684991049243"><code>0f95b74</code></a> Reduce <code>regeneratorRuntime</code> size (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17213">#17213</a>)</li>
<li><a href="https://github.com/babel/babel/commit/317e332e650bc04907bc787ab79f930288a3e71e"><code>317e332</code></a> Enforce node protocol import (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17207">#17207</a>)</li>
<li><a href="https://github.com/babel/babel/commit/14ef1e972277bc7b80527fc2aa4d4ffc7662c768"><code>14ef1e9</code></a> Babel 8 cleanup (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17211">#17211</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel/commits/v7.27.2/packages/babel-preset-env">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-react` from 7.26.3 to 7.27.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-react</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.27.1 (2025-04-30)</h2>
<p>Thanks <a href="https://github.com/kermanx"><code>@​kermanx</code></a> and <a href="https://github.com/woaitsAryan"><code>@​woaitsAryan</code></a> for your first PRs!</p>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17254">#17254</a> Allow <code>using of</code> as lexical declaration within for (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17230">#17230</a> Disallow get/set in TSPropertySignature (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17193">#17193</a> Stricter TSImportType options parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17137">#17137</a> fix: do expressions should allow early exit (<a href="https://github.com/kermanx"><code>@​kermanx</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-wrap-function</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17251">#17251</a> Fix: propagate argument evaluation errors through async promise chain (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-remap-async-to-generator</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17231">#17231</a> fix apply()/call() annotated as pure (<a href="https://github.com/Lacsw"><code>@​Lacsw</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-fixtures</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17233">#17233</a> Create ChainExpression within TSInstantiationExpression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17226">#17226</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 2) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17224">#17224</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 1) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17080">#17080</a> Fix start of TSParameterProperty (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17228">#17228</a> Update firefox bugfix compat data (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17156">#17156</a> fix: Objects and arrays with multiple references should not be evaluated (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17216">#17216</a> Fix: support const type parameter in generator (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-bugfix-v8-spread-parameters-in-optional-chaining</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-transform-arrow-functions</code>, <code>babel-plugin-transform-class-properties</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-parameters</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17221">#17221</a> Reduce generated names size for the 10th-11th (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17263">#17263</a> Remove unused <code>regenerator-runtime</code> dep in <code>@babel/runtime</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17256">#17256</a> Tune plugin compat data (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17236">#17236</a> migrate babel-compat-data build script to mjs (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-register</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16844">#16844</a> Migrate <code>@babel/register</code> to cts (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17205">#17205</a> Inline regenerator in the relevant packages (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><em>All packages</em>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17207">#17207</a> Enforce node protocol import (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-react</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.27.1 (2025-04-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17254">#17254</a> Allow <code>using of</code> as lexical declaration within for (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17230">#17230</a> Disallow get/set in TSPropertySignature (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17193">#17193</a> Stricter TSImportType options parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17137">#17137</a> fix: do expressions should allow early exit (<a href="https://github.com/kermanx"><code>@​kermanx</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-wrap-function</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17251">#17251</a> Fix: propagate argument evaluation errors through async promise chain (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-remap-async-to-generator</code>, <code>babel-plugin-transform-async-to-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17231">#17231</a> fix apply()/call() annotated as pure (<a href="https://github.com/Lacsw"><code>@​Lacsw</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-fixtures</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17233">#17233</a> Create ChainExpression within TSInstantiationExpression (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17226">#17226</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 2) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17224">#17224</a> Fill optional AST properties when both estree and typescript parser plugin are enabled (Part 1) (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17080">#17080</a> Fix start of TSParameterProperty (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17228">#17228</a> Update firefox bugfix compat data (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17156">#17156</a> fix: Objects and arrays with multiple references should not be evaluated (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17216">#17216</a> Fix: support const type parameter in generator (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-bugfix-v8-spread-parameters-in-optional-chaining</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-transform-arrow-functions</code>, <code>babel-plugin-transform-class-properties</code>, <code>babel-plugin-transform-destructuring</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-parameters</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17221">#17221</a> Reduce generated names size for the 10th-11th (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17263">#17263</a> Remove unused <code>regenerator-runtime</code> dep in <code>@babel/runtime</code> (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17256">#17256</a> Tune plugin compat data (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17236">#17236</a> migrate babel-compat-data build script to mjs (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17232">#17232</a> Bump typescript-eslint to 8.29.1 (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17219">#17219</a> test: add basic typescript-eslint integration tests (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17205">#17205</a> Inline regenerator in the relevant packages (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-register</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16844">#16844</a> Migrate <code>@babel/register</code> to cts (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-cli</code>, <code>babel-compat-data</code>, <code>babel-core</code>, <code>babel-generator</code>, <code>babel-helper-compilation-targets</code>, <code>babel-helper-fixtures</code>, <code>babel-helper-module-imports</code>, <code>babel-helper-module-transforms</code>, <code>babel-helper-plugin-test-runner</code>, <code>babel-helper-transform-fixture-test-runner</code>, <code>babel-helpers</code>, <code>babel-node</code>, <code>babel-parser</code>, <code>babel-plugin-transform-modules-amd</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-modules-systemjs</code>, <code>babel-plugin-transform-modules-umd</code>, <code>babel-plugin-transform-react-display-name</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>, <code>babel-plugin-transform-typeof-symbol</code>, <code>babel-plugin-transform-typescript</code>, <code>babel-preset-env</code>, <code>babel-register</code>, <code>babel-standalone</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17207">#17207</a> Enforce node protocol import (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-regenerator</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/eebd3a06021c13d335b5b0bd79734df3abbea678"><code>eebd3a0</code></a> v7.27.1</li>
<li><a href="https://github.com/babel/babel/commit/fdc0fb59e119ee0b38bced63867a344a5b4bc2f3"><code>fdc0fb5</code></a> [Babel 8] Bump nodejs requirements to <code>^20.19.0 || &gt;= 22.12.0</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/17204">#17204</a>)</li>
<li><a href="https://github.com/babel/babel/commit/cd24cc07ef6558b7f6510f9177f6393c91b0549f"><code>cd24cc0</code></a> chore: Update TS 5.7 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/17053">#17053</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.27.1/packages/babel-preset-react">compare view</a></li>
</ul>
</details>
<br />

Updates `babel-jest` from 29.7.0 to 30.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">babel-jest's releases</a>.</em></p>
<blockquote>
<h2>Jest 30</h2>
<p>Today we are happy to announce the release of Jest 30. This release features a substantial number of changes, fixes, and improvements. While it is one of the largest major releases of Jest ever, we admit that three years for a major release is too long. In the future, we are aiming to make more frequent major releases to keep Jest great for the next decade.</p>
<p>If you want to skip all the news and just get going, run <code>npm install jest@^30.0.0</code> and follow the migration guide: <a href="https://jestjs.io/docs/upgrading-to-jest30">Upgrading from Jest 29 to 30</a>.</p>
<p><a href="https://jestjs.io/blog/2025/06/04/jest-30">Read the full blog post</a></p>
<h3>Features</h3>
<ul>
<li><code>[*]</code> Renamed <code>globalsCleanupMode</code> to <code>globalsCleanup</code> and <code>--waitNextEventLoopTurnForUnhandledRejectionEvents</code> to <code>--waitForUnhandledRejections</code></li>
<li><code>[expect]</code> Add <code>ArrayOf</code> asymmetric matcher for validating array elements. (<a href="https://redirect.github.com/jestjs/jest/pull/15567">#15567</a>)</li>
<li><code>[babel-jest]</code> Add option <code>excludeJestPreset</code> to allow opting out of <code>babel-preset-jest</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15164">#15164</a>)</li>
<li><code>[expect]</code> Revert <a href="https://redirect.github.com/jestjs/jest/pull/15038">#15038</a> to fix <code>expect(fn).toHaveBeenCalledWith(expect.objectContaining(...))</code> when there are multiple calls (<a href="https://redirect.github.com/jestjs/jest/pull/15508">#15508</a>)</li>
<li><code>[jest-circus, jest-cli, jest-config]</code> Add <code>waitNextEventLoopTurnForUnhandledRejectionEvents</code> flag to minimise performance impact of correct detection of unhandled promise rejections introduced in <a href="https://redirect.github.com/jestjs/jest/pull/14315">#14315</a> (<a href="https://redirect.github.com/jestjs/jest/pull/14681">#14681</a>)</li>
<li><code>[jest-circus]</code> Add a <code>waitBeforeRetry</code> option to <code>jest.retryTimes</code> (<a href="https://redirect.github.com/jestjs/jest/pull/14738">#14738</a>)</li>...

_Description has been truncated_

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.

---

## Discussion

### Comment by @dependabot on September 15, 2025 at 06:11 PM UTC

Dependabot tried to update this pull request, but something went wrong. We're looking into it, but in the meantime you can retry the update by commenting `@dependabot rebase`.

### Comment by @dependabot on September 22, 2025 at 06:02 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1344*
