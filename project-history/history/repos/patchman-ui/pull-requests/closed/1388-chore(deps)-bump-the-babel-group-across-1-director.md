---
type: pull_request
number: 1388
title: "chore(deps): bump the babel group across 1 directory with 3 updates"
state: closed
author: dependabot
created: 2025-09-22T18:19:34Z
updated: 2025-10-27T18:02:32Z
closed: 2025-10-27T18:02:30Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-462493d5f5
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1388
---

# Pull Request #1388: chore(deps): bump the babel group across 1 directory with 3 updates

**Author**: @dependabot
**Created**: September 22, 2025 at 06:19 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-462493d5f5`

## Description

Bumps the babel group with 3 updates in the / directory: [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser), [babel-jest](https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest) and [babel-loader](https://github.com/babel/babel-loader).

Updates `@babel/eslint-parser` from 7.28.0 to 7.28.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/eslint-parser</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.28.4 (2025-09-05)</h2>
<p>Thanks <a href="https://github.com/gwillen"><code>@​gwillen</code></a> and <a href="https://github.com/mrginglymus"><code>@​mrginglymus</code></a> for your first PRs!</p>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-core</code>, <code>babel-helper-check-duplicate-nodes</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17493">#17493</a> Update Jest to v30.1.1 (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17455">#17455</a> chore: Clean up <code>transform-regenerator</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17474">#17474</a> Switch to <code>@​jridgewell/remapping</code> (<a href="https://github.com/mrginglymus"><code>@​mrginglymus</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Bill Collins (<a href="https://github.com/mrginglymus"><code>@​mrginglymus</code></a>)</li>
<li>Glenn Willen (<a href="https://github.com/gwillen"><code>@​gwillen</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.28.3 (2025-08-14)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17443">#17443</a> [static blocks] Do not inject new static fields after static code (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17465">#17465</a> fix(parser/typescript): parse <code>import(&quot;./a&quot;, {with:{},})</code> (<a href="https://github.com/easrng"><code>@​easrng</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17478">#17478</a> fix(parser): stop subscript parsing on async arrow (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17363">#17363</a> Do not save last yield in call in temp var (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:memo: Documentation</h4>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17448">#17448</a> move eslint-{parser,plugin} docs to the website (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17454">#17454</a> Enable type checking for <code>scripts</code> and <code>babel-worker.cjs</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h4>:microscope: Output optimization</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17444">#17444</a> Optimize do expression output (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Jam Balaya (<a href="https://github.com/JamBalaya56562"><code>@​JamBalaya56562</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li>easrng (<a href="https://github.com/easrng"><code>@​easrng</code></a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/eslint-parser</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.28.4 (2025-09-05)</h2>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-core</code>, <code>babel-helper-check-duplicate-nodes</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17493">#17493</a> Update Jest to v30.1.1 (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17455">#17455</a> chore: Clean up <code>transform-regenerator</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17474">#17474</a> Switch to <code>@​jridgewell/remapping</code> (<a href="https://github.com/mrginglymus"><code>@​mrginglymus</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.28.3 (2025-08-14)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17443">#17443</a> [static blocks] Do not inject new static fields after static code (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17465">#17465</a> fix(parser/typescript): parse <code>import(&quot;./a&quot;, {with:{},})</code> (<a href="https://github.com/easrng"><code>@​easrng</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17478">#17478</a> fix(parser): stop subscript parsing on async arrow (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-plugin-transform-regenerator</code>, <code>babel-plugin-transform-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17363">#17363</a> Do not save last yield in call in temp var (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:memo: Documentation</h4>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17448">#17448</a> move eslint-{parser,plugin} docs to the website (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17454">#17454</a> Enable type checking for <code>scripts</code> and <code>babel-worker.cjs</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h4>:microscope: Output optimization</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>, <code>babel-plugin-proposal-do-expressions</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17444">#17444</a> Optimize do expression output (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.28.2 (2025-07-24)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17445">#17445</a> [babel 7] Make <code>operator</code> param in <code>t.tsTypeOperator</code> optional (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-regenerator</code>, <code>babel-preset-env</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17441">#17441</a> fix: <code>regeneratorDefine</code> compatibility with es5 strict mode (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.28.1 (2025-07-12)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17426">#17426</a> fix: <code>regenerator</code> correctly handles <code>throw</code> outside of <code>try</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:memo: Documentation</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17422">#17422</a> Add missing FunctionParameter docs (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
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
<li><a href="https://github.com/babel/babel/commit/35055e392079a65830b7bf5b1d1c1fc4de90a78f"><code>35055e3</code></a> v7.28.4</li>
<li><a href="https://github.com/babel/babel/commit/f04083a70573804935797e5a4d7d8f647d30a59a"><code>f04083a</code></a> [Babel 8] Align TSMappedType AST (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17479">#17479</a>)</li>
<li><a href="https://github.com/babel/babel/commit/432a7ffbff568efb608a5ddd8e87aea39c76bdad"><code>432a7ff</code></a> fix(parser/typescript): parse <code>import(&quot;./a&quot;, {with:{},})</code> (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17465">#17465</a>)</li>
<li><a href="https://github.com/babel/babel/commit/e77e3b02d731da53463c5c80c66858cbb0c6dcfb"><code>e77e3b0</code></a> move eslint-{parser,plugin} docs to the website (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17448">#17448</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.28.4/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<br />

Updates `babel-jest` from 29.7.0 to 30.1.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">babel-jest's releases</a>.</em></p>
<blockquote>
<h2>30.1.2</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Correct snapshot header regexp to work with newline across OSes (<a href="https://redirect.github.com/jestjs/jest/pull/15803">#15803</a>)</li>
</ul>
<h2>30.1.1</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot warning not handling Windows end-of-line sequences (<a href="https://redirect.github.com/jestjs/jest/pull/15800">#15800</a>)</li>
</ul>
<h2>30.1.0</h2>
<h2>Features</h2>
<ul>
<li><code>[jest-leak-detector]</code> Configurable GC aggressiveness regarding to V8 heap snapshot generation (<a href="https://redirect.github.com/jestjs/jest/pull/15793/">#15793</a>)</li>
<li><code>[jest-runtime]</code> Reduce redundant ReferenceError messages</li>
<li><code>[jest-core]</code> Include test modules that failed to load when --onlyFailures is active</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>`[jest-snapshot-utils] Fix deprecated goo.gl snapshot guide link not getting replaced with fully canonical URL (<a href="https://redirect.github.com/jestjs/jest/pull/15787">#15787</a>)</li>
<li><code>[jest-circus]</code> Fix <code>it.concurrent</code> not working with <code>describe.skip</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15765">#15765</a>)</li>
<li><code>[jest-snapshot]</code> Fix mangled inline snapshot updates when used with Prettier 3 and CRLF line endings</li>
<li><code>[jest-runtime]</code> Importing from <code>@jest/globals</code> in more than one file no longer breaks relative paths (<a href="https://redirect.github.com/jestjs/jest/issues/15772">#15772</a>)</li>
</ul>
<h1>Chore</h1>
<ul>
<li><code>[expect]</code> Update docblock for <code>toContain()</code> to display info on substring check (<a href="https://redirect.github.com/jestjs/jest/pull/15789">#15789</a>)</li>
</ul>
<h2>30.0.2</h2>
<h2>What's Changed</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-matcher-utils]</code> Make 'deepCyclicCopyObject' safer by setting descriptors to a null-prototype object (<a href="https://redirect.github.com/jestjs/jest/pull/15689">#15689</a>)</li>
<li><code>[jest-util]</code> Make garbage collection protection property writable (<a href="https://redirect.github.com/jestjs/jest/pull/15689">#15689</a>)</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">https://github.com/jestjs/jest/blob/main/CHANGELOG.md</a></p>
<h2>Jest 30.0.1</h2>
<h2>What's Changed</h2>
<h3>Features</h3>
<ul>
<li><code>[jest-resolver]</code> Implement the <code>defaultAsyncResolver</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15679">#15679</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-resolver]</code> Resolve builtin modules correctly (<a href="https://redirect.github.com/jestjs/jest/pull/15683">#15683</a>)</li>
<li><code>[jest-environment-node, jest-util]</code> Avoid setting globals cleanup protection symbol when feature is off (<a href="https://redirect.github.com/jestjs/jest/pull/15684">#15684</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">babel-jest's changelog</a>.</em></p>
<blockquote>
<h2>30.1.2</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Correct snapshot header regexp to work with newline across OSes (<a href="https://redirect.github.com/jestjs/jest/pull/15803">#15803</a>)</li>
</ul>
<h2>30.1.1</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot warning not handling Windows end-of-line sequences (<a href="https://redirect.github.com/jestjs/jest/pull/15800">#15800</a>)</li>
</ul>
<h2>30.1.0</h2>
<h2>Features</h2>
<ul>
<li><code>[jest-leak-detector]</code> Configurable GC aggressiveness regarding to V8 heap snapshot generation (<a href="https://redirect.github.com/jestjs/jest/pull/15793/">#15793</a>)</li>
<li><code>[jest-runtime]</code> Reduce redundant ReferenceError messages</li>
<li><code>[jest-core]</code> Include test modules that failed to load when --onlyFailures is active</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot guide link not getting replaced with fully canonical URL (<a href="https://redirect.github.com/jestjs/jest/pull/15787">#15787</a>)</li>
<li><code>[jest-circus]</code> Fix <code>it.concurrent</code> not working with <code>describe.skip</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15765">#15765</a>)</li>
<li><code>[jest-snapshot]</code> Fix mangled inline snapshot updates when used with Prettier 3 and CRLF line endings</li>
<li><code>[jest-runtime]</code> Importing from <code>@jest/globals</code> in more than one file no longer breaks relative paths (<a href="https://redirect.github.com/jestjs/jest/issues/15772">#15772</a>)</li>
</ul>
<h1>Chore</h1>
<ul>
<li><code>[expect]</code> Update docblock for <code>toContain()</code> to display info on substring check (<a href="https://redirect.github.com/jestjs/jest/pull/15789">#15789</a>)</li>
</ul>
<h2>30.0.5</h2>
<h3>Features</h3>
<ul>
<li><code>[jest-config]</code> Allow <code>testMatch</code> to take a string value</li>
<li><code>[jest-worker]</code> Let <code>workerIdleMemoryLimit</code> accept 0 to always restart worker child processes</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[expect]</code> Fix <code>bigint</code> error (<a href="https://redirect.github.com/jestjs/jest/pull/15702">#15702</a>)</li>
</ul>
<h2>30.0.4</h2>
<h3>Features</h3>
<ul>
<li><code>[expect]</code> The <code>Inverse</code> type is now exported (<a href="https://redirect.github.com/jestjs/jest/pull/15714">#15714</a>)</li>
<li><code>[expect]</code> feat: support <code>async functions</code> in <code>toBe</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15704">#15704</a>)</li>
</ul>
<h3>Fixes</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jestjs/jest/commit/ebfa31cc9787303e8698a1a029a162a18e8974aa"><code>ebfa31c</code></a> v30.1.2</li>
<li><a href="https://github.com/jestjs/jest/commit/d347c0f3f87f976a1dbd9761d503e45f5ced2a7e"><code>d347c0f</code></a> v30.1.1</li>
<li><a href="https://github.com/jestjs/jest/commit/4d5f41d0885c1d9630c81b4fd47f74ab0615e18f"><code>4d5f41d</code></a> v30.1.0</li>
<li><a href="https://github.com/jestjs/jest/commit/22236cf58b66039f81893537c90dee290bab427f"><code>22236cf</code></a> v30.0.5</li>
<li><a href="https://github.com/jestjs/jest/commit/f4296d2bc85c1405f84ddf613a25d0bc3766b7e5"><code>f4296d2</code></a> v30.0.4</li>
<li><a href="https://github.com/jestjs/jest/commit/393acbfac31f64bb38dff23c89224797caded83c"><code>393acbf</code></a> v30.0.2</li>
<li><a href="https://github.com/jestjs/jest/commit/5ce865b4060189fe74cd486544816c079194a0f7"><code>5ce865b</code></a> v30.0.1</li>
<li><a href="https://github.com/jestjs/jest/commit/469f665c2d3bea4a55a194ceebae88724b7202cd"><code>469f665</code></a> v30.0.0</li>
<li><a href="https://github.com/jestjs/jest/commit/ce14203d9156f830a8e24a6e3e8205f670a72a40"><code>ce14203</code></a> v30.0.0-rc.1</li>
<li><a href="https://github.com/jestjs/jest/commit/a2218e4f794f914884c403ecceb274ada595f2b9"><code>a2218e4</code></a> Stop using <code>import X = require('…')</code>. (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/15659">#15659</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/jestjs/jest/commits/v30.1.2/packages/babel-jest">compare view</a></li>
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

### Comment by @codecov-commenter on September 29, 2025 at 10:14 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1388?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.46%. Comparing base ([`2596af5`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2596af56e545e1fe3f331dd66d72db671b157bd5?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b706400`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b70640010163f1c7f7bee5d170d8e002c7ea4543?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff            @@
##           master    #1388    +/-   ##
========================================
  Coverage   62.46%   62.46%            
========================================
  Files         126      126            
  Lines        3312     3312            
  Branches      857      829    -28     
========================================
  Hits         2069     2069            
- Misses       1122     1243   +121     
+ Partials      121        0   -121     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1388/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1388/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.46% <ø> (?)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1388/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (?)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1388/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.46% <ø> (?)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1388?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @dependabot on October 27, 2025 at 06:02 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1388*
