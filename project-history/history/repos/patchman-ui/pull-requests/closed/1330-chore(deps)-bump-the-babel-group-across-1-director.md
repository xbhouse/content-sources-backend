---
type: pull_request
number: 1330
title: "chore(deps): bump the babel group across 1 directory with 5 updates"
state: closed
author: dependabot
created: 2025-05-05T19:05:49Z
updated: 2025-06-16T20:18:57Z
closed: 2025-06-16T20:18:56Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-96275f3acc
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1330
---

# Pull Request #1330: chore(deps): bump the babel group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: May 05, 2025 at 07:05 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-96275f3acc`

## Description

Bumps the babel group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser) | `7.26.5` | `7.27.1` |
| [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime) | `7.25.9` | `7.27.1` |
| [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env) | `7.26.7` | `7.27.1` |
| [@babel/preset-react](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react) | `7.26.3` | `7.27.1` |
| [babel-loader](https://github.com/babel/babel-loader) | `9.2.1` | `10.0.0` |


Updates `@babel/eslint-parser` from 7.26.5 to 7.27.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/eslint-parser</code>'s releases</a>.</em></p>
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
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/eslint-parser</code>'s changelog</a>.</em></p>
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
<li><a href="https://github.com/babel/babel/commit/9a40d852e7bcf65cfadfaaacf0b85d66a0b59d6e"><code>9a40d85</code></a> [Babel 8]: Remove record and tuple syntax support (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17242">#17242</a>)</li>
<li><a href="https://github.com/babel/babel/commit/3766c4ddfaaf5429eb8443f42d172e5845a0feef"><code>3766c4d</code></a> Fill optional AST properties when both estree and typescript parser plugin ar...</li>
<li><a href="https://github.com/babel/babel/commit/0de0e2b0a8a4c6fec519bf1882e32ed49424129d"><code>0de0e2b</code></a> Bump typescript-eslint to 8.29.1 (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17232">#17232</a>)</li>
<li><a href="https://github.com/babel/babel/commit/c1614c804c52e370dd48a3aba59bb17eb24dfbf8"><code>c1614c8</code></a> Fill optional AST properties when both estree and typescript parser plugin ar...</li>
<li><a href="https://github.com/babel/babel/commit/7979e3fc56cfba9d62b4dfa4ef71cfc025a8863f"><code>7979e3f</code></a> Fill optional AST properties when both estree and typescript parser plugin ar...</li>
<li><a href="https://github.com/babel/babel/commit/c18e4791b48f519edc3bdd88c9541cc72145e36d"><code>c18e479</code></a> test: add basic typescript-eslint integration tests (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17219">#17219</a>)</li>
<li><a href="https://github.com/babel/babel/commit/317e332e650bc04907bc787ab79f930288a3e71e"><code>317e332</code></a> Enforce node protocol import (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17207">#17207</a>)</li>
<li><a href="https://github.com/babel/babel/commit/fdc0fb59e119ee0b38bced63867a344a5b4bc2f3"><code>fdc0fb5</code></a> [Babel 8] Bump nodejs requirements to <code>^20.19.0 || &gt;= 22.12.0</code> (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17204">#17204</a>)</li>
<li><a href="https://github.com/babel/babel/commit/5c350eab83dd12268add44cce0eeda6c898211e3"><code>5c350ea</code></a> v7.27.0</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel/commits/v7.27.1/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/plugin-transform-runtime` from 7.25.9 to 7.27.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/plugin-transform-runtime</code>'s releases</a>.</em></p>
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
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/plugin-transform-runtime</code>'s changelog</a>.</em></p>
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
<li><a href="https://github.com/babel/babel/commit/317e332e650bc04907bc787ab79f930288a3e71e"><code>317e332</code></a> Enforce node protocol import (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17207">#17207</a>)</li>
<li><a href="https://github.com/babel/babel/commit/fdc0fb59e119ee0b38bced63867a344a5b4bc2f3"><code>fdc0fb5</code></a> [Babel 8] Bump nodejs requirements to <code>^20.19.0 || &gt;= 22.12.0</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17204">#17204</a>)</li>
<li><a href="https://github.com/babel/babel/commit/e1ce99df422971175249509e7bbc2b327b8f7957"><code>e1ce99d</code></a> v7.26.10</li>
<li><a href="https://github.com/babel/babel/commit/ce0e87b12cf50a88bd06d4e39bf4ff205da862a2"><code>ce0e87b</code></a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17138">#17138</a>)</li>
<li><a href="https://github.com/babel/babel/commit/64bca7b5f308cd52c192a5c821a57f6d1b0475f4"><code>64bca7b</code></a> v7.26.9</li>
<li><a href="https://github.com/babel/babel/commit/5315446a9a9348bd9646464bfa7b0e4381647e04"><code>5315446</code></a> [babel 8] Remove babel 7-specific imports (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17111">#17111</a>)</li>
<li><a href="https://github.com/babel/babel/commit/059394159de3af74446be6fba896dd4ce178333c"><code>0593941</code></a> v7.26.8</li>
<li><a href="https://github.com/babel/babel/commit/1137c29a30badc6e51d58ab21983ba2824140e5b"><code>1137c29</code></a> Update dependency babel-plugin-polyfill-corejs3 to ^0.11.0 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17097">#17097</a>)</li>
<li><a href="https://github.com/babel/babel/commit/cd24cc07ef6558b7f6510f9177f6393c91b0549f"><code>cd24cc0</code></a> chore: Update TS 5.7 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17053">#17053</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.27.1/packages/babel-plugin-transform-runtime">compare view</a></li>
</ul>
</details>
<br />

Updates `@babel/preset-env` from 7.26.7 to 7.27.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-env</code>'s releases</a>.</em></p>
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
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-env</code>'s changelog</a>.</em></p>
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
<li><a href="https://github.com/babel/babel/commit/5c5e77c05b52f4d333c4910014efe700022d8cef"><code>5c5e77c</code></a> Tune plugin compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17256">#17256</a>)</li>
<li><a href="https://github.com/babel/babel/commit/aca4f0138c4b163f39fb398e862dc4400ae9c986"><code>aca4f01</code></a> chore: bump compat-data sources (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17253">#17253</a>)</li>
<li><a href="https://github.com/babel/babel/commit/726438499badd1454911073ddba20c8acfaa3781"><code>7264384</code></a> Update firefox bugfix compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17228">#17228</a>)</li>
<li><a href="https://github.com/babel/babel/commit/0f95b748a9a5f90c3b23d4d72299684991049243"><code>0f95b74</code></a> Reduce <code>regeneratorRuntime</code> size (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17213">#17213</a>)</li>
<li><a href="https://github.com/babel/babel/commit/317e332e650bc04907bc787ab79f930288a3e71e"><code>317e332</code></a> Enforce node protocol import (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17207">#17207</a>)</li>
<li><a href="https://github.com/babel/babel/commit/14ef1e972277bc7b80527fc2aa4d4ffc7662c768"><code>14ef1e9</code></a> Babel 8 cleanup (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17211">#17211</a>)</li>
<li><a href="https://github.com/babel/babel/commit/97105cb9f285964065a525c7ab4fd718c57e2a40"><code>97105cb</code></a> Re-convert regeneratorRuntime to helper format (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17205">#17205</a>)</li>
<li><a href="https://github.com/babel/babel/commit/fdc0fb59e119ee0b38bced63867a344a5b4bc2f3"><code>fdc0fb5</code></a> [Babel 8] Bump nodejs requirements to <code>^20.19.0 || &gt;= 22.12.0</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17204">#17204</a>)</li>
<li><a href="https://github.com/babel/babel/commit/624e699be58e49670bc4683cac404922a870dbda"><code>624e699</code></a> [Babel 8] Align esmodules: true behaviour to intersect (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17188">#17188</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel/commits/v7.27.1/packages/babel-preset-env">compare view</a></li>
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
<li><code>babel-cli</code>, <code>babel-compat-data</code>, <code>babel-core</code>, <code>babel-generator</code>, <code>babel-helper-compilation-targets</code>, <code>babel-helper-fixtures</code>, <code>babel-helper-module-imports</code>, <code>babel-helper-module-transforms</code>, <code>babel-helper-plugin-test-runner</code>, <code>babel-helper-transform-fixture-test-runner</code>, <code>babel-helpers</code>, <code>babel-node</code>, <code>babel-parser</code>, <code>babel-plugin-transform-modules-amd</code>, <code>babel-plugin-transform-modules-common...

_Description has been truncated_

---

## Discussion

### Comment by @dependabot on June 16, 2025 at 08:18 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1330*
