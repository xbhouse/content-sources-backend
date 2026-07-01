---
type: pull_request
number: 1638
title: "chore(deps): bump the babel group across 1 directory with 5 updates"
state: merged
author: dependabot
created: 2026-05-26T23:58:22Z
updated: 2026-05-28T19:19:54Z
closed: 2026-05-28T19:11:57Z
merged: 2026-05-28T19:11:56Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-5b42412689
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1638
---

# Pull Request #1638: chore(deps): bump the babel group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: May 26, 2026 at 11:58 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-5b42412689`

## Description

Bumps the babel group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser) | `7.28.6` | `7.29.7` |
| [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime) | `7.29.0` | `7.29.7` |
| [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env) | `7.29.5` | `7.29.7` |
| [@babel/preset-react](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react) | `7.28.5` | `7.29.7` |
| [@babel/preset-typescript](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-typescript) | `7.28.5` | `7.29.7` |


Updates `@babel/eslint-parser` from 7.28.6 to 7.29.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases">@​babel/eslint-parser's releases</a>.</em></p>
<blockquote>
<h2>v7.29.7 (2026-05-25)</h2>
<p>Re-release all packages with npm provenance attestations</p>
<h2>v7.29.6 (2026-05-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18014">#18014</a> Catchup source map position in preserveFormat (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18001">#18001</a> [7.x packport]Improve input source map handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17998">#17998</a> Preserve original identifier names from input sourcemaps (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17992">#17992</a>) (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Mateusz Burzyński (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
<h2>v7.29.5 (2026-05-05)</h2>
<h4>:house:  Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li>Update <code>@babel/*</code> dependencies</li>
</ul>
</li>
</ul>
<h2>v7.29.4 (2026-05-05)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-modules-systemjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17974">#17974</a> [7.x backport]fix(systemjs): improve module string name support (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 1</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h2>v7.29.3 (2026-04-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17923">#17923</a> Support flow extends bound (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-proposal-decorators</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17931">#17931</a> fix(decorators): replace super within all removed static elements (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-register</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17915">#17915</a> Fix thread synchronization issues in <code>@babel/register</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-bugfix-safari-rest-destructuring-rhs-array</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17788">#17788</a> Add bugfix plugin for Safari array rest destructuring bug (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/4fba7541180bf5f58256d8e358b544e3831ad090"><code>4fba754</code></a> v7.29.7</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.7/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/plugin-transform-runtime` from 7.29.0 to 7.29.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases">@​babel/plugin-transform-runtime's releases</a>.</em></p>
<blockquote>
<h2>v7.29.7 (2026-05-25)</h2>
<p>Re-release all packages with npm provenance attestations</p>
<h2>v7.29.6 (2026-05-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18014">#18014</a> Catchup source map position in preserveFormat (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18001">#18001</a> [7.x packport]Improve input source map handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17998">#17998</a> Preserve original identifier names from input sourcemaps (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17992">#17992</a>) (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Mateusz Burzyński (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
<h2>v7.29.5 (2026-05-05)</h2>
<h4>:house:  Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li>Update <code>@babel/*</code> dependencies</li>
</ul>
</li>
</ul>
<h2>v7.29.4 (2026-05-05)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-modules-systemjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17974">#17974</a> [7.x backport]fix(systemjs): improve module string name support (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 1</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h2>v7.29.3 (2026-04-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17923">#17923</a> Support flow extends bound (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-proposal-decorators</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17931">#17931</a> fix(decorators): replace super within all removed static elements (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-register</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17915">#17915</a> Fix thread synchronization issues in <code>@babel/register</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-bugfix-safari-rest-destructuring-rhs-array</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17788">#17788</a> Add bugfix plugin for Safari array rest destructuring bug (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/4fba7541180bf5f58256d8e358b544e3831ad090"><code>4fba754</code></a> v7.29.7</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.7/packages/babel-plugin-transform-runtime">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-env` from 7.29.5 to 7.29.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases">@​babel/preset-env's releases</a>.</em></p>
<blockquote>
<h2>v7.29.7 (2026-05-25)</h2>
<p>Re-release all packages with npm provenance attestations</p>
<h2>v7.29.6 (2026-05-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18014">#18014</a> Catchup source map position in preserveFormat (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18001">#18001</a> [7.x packport]Improve input source map handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17998">#17998</a> Preserve original identifier names from input sourcemaps (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17992">#17992</a>) (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Mateusz Burzyński (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/4fba7541180bf5f58256d8e358b544e3831ad090"><code>4fba754</code></a> v7.29.7</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.7/packages/babel-preset-env">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-react` from 7.28.5 to 7.29.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases">@​babel/preset-react's releases</a>.</em></p>
<blockquote>
<h2>v7.29.7 (2026-05-25)</h2>
<p>Re-release all packages with npm provenance attestations</p>
<h2>v7.29.6 (2026-05-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18014">#18014</a> Catchup source map position in preserveFormat (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18001">#18001</a> [7.x packport]Improve input source map handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17998">#17998</a> Preserve original identifier names from input sourcemaps (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/17992">#17992</a>) (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Mateusz Burzyński (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
<h2>v7.29.5 (2026-05-05)</h2>
<h4>:house:  Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li>Update <code>@babel/*</code> dependencies</li>
</ul>
</li>
</ul>
<h2>v7.29.4 (2026-05-05)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-modules-systemjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17974">#17974</a> [7.x backport]fix(systemjs): improve module string name support (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 1</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h2>v7.29.3 (2026-04-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17923">#17923</a> Support flow extends bound (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-proposal-decorators</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17931">#17931</a> fix(decorators): replace super within all removed static elements (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-register</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17915">#17915</a> Fix thread synchronization issues in <code>@babel/register</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-bugfix-safari-rest-destructuring-rhs-array</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17788">#17788</a> Add bugfix plugin for Safari array rest destructuring bug (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/4fba7541180bf5f58256d8e358b544e3831ad090"><code>4fba754</code></a> v7.29.7</li>
<li><a href="https://github.com/babel/babel/commit/f3a22268bdc4fc6748cbc2be718a4d1090bdaf00"><code>f3a2226</code></a> [babel 7] Delete Babel 8 fixtures (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/17729">#17729</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.7/packages/babel-preset-react">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-typescript` from 7.28.5 to 7.29.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases">@​babel/preset-typescript's releases</a>.</em></p>
<blockquote>
<h2>v7.29.7 (2026-05-25)</h2>
<p>Re-release all packages with npm provenance attestations</p>
<h2>v7.29.6 (2026-05-25)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18014">#18014</a> Catchup source map position in preserveFormat (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/18001">#18001</a> [7.x packport]Improve input source map handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>, <code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17998">#17998</a> Preserve original identifier names from input sourcemaps (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-typescript/issues/17992">#17992</a>) (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 3</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Mateusz Burzyński (<a href="https://github.com/Andarist"><code>@​Andarist</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
<h2>v7.29.5 (2026-05-05)</h2>
<h4>:house:  Internal</h4>
<ul>
<li><code>babel-preset-env</code>
<ul>
<li>Update <code>@babel/*</code> dependencies</li>
</ul>
</li>
</ul>
<h2>v7.29.4 (2026-05-05)</h2>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-transform-modules-systemjs</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17974">#17974</a> [7.x backport]fix(systemjs): improve module string name support (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 1</h4>
<ul>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
<h2>v7.29.3 (2026-04-30)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17923">#17923</a> Support flow extends bound (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-proposal-decorators</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17931">#17931</a> fix(decorators): replace super within all removed static elements (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-register</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17915">#17915</a> Fix thread synchronization issues in <code>@babel/register</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-compat-data</code>, <code>babel-plugin-bugfix-safari-rest-destructuring-rhs-array</code>, <code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17788">#17788</a> Add bugfix plugin for Safari array rest destructuring bug (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-parser</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/4fba7541180bf5f58256d8e358b544e3831ad090"><code>4fba754</code></a> v7.29.7</li>
<li><a href="https://github.com/babel/babel/commit/f3a22268bdc4fc6748cbc2be718a4d1090bdaf00"><code>f3a2226</code></a> [babel 7] Delete Babel 8 fixtures (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-typescript/issues/17729">#17729</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.7/packages/babel-preset-typescript">compare view</a></li>
</ul>
</details>
<br />

---

## Discussion

### Comment by @codecov-commenter on May 27, 2026 at 12:01 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1638?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.58%. Comparing base ([`78c91be`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/78c91be9ef123cd63d2da3ad661c9f05f741e8a8?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`7cf331c`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7cf331c246c19e027276737e781e8ceb63423fa6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1638   +/-   ##
=======================================
  Coverage   77.58%   77.58%           
=======================================
  Files         103      103           
  Lines        3266     3266           
  Branches      729      729           
=======================================
  Hits         2534     2534           
  Misses        655      655           
  Partials       77       77           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1638?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @xbhouse on May 28, 2026 at 05:51 PM UTC

@dependabot rebase

---

## Reviews

### Review by @swadeley - Approved on May 28, 2026 at 07:04 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1638*
