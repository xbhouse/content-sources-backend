---
type: pull_request
number: 1511
title: "chore(deps): bump the babel group with 6 updates"
state: merged
author: dependabot
created: 2026-02-16T15:32:35Z
updated: 2026-02-18T14:32:33Z
closed: 2026-02-18T14:32:23Z
merged: 2026-02-18T14:32:23Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-547e6ebbe4
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1511
---

# Pull Request #1511: chore(deps): bump the babel group with 6 updates

**Author**: @dependabot
**Created**: February 16, 2026 at 03:32 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-547e6ebbe4`

## Description

Bumps the babel group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [@babel/eslint-parser](https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser) | `7.28.0` | `7.28.6` |
| [@babel/plugin-transform-runtime](https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime) | `7.28.3` | `7.29.0` |
| [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env) | `7.28.3` | `7.29.0` |
| [@babel/preset-react](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react) | `7.27.1` | `7.28.5` |
| [babel-jest](https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest) | `29.7.0` | `30.2.0` |
| [babel-loader](https://github.com/babel/babel-loader) | `9.2.1` | `10.0.0` |

Updates `@babel/eslint-parser` from 7.28.0 to 7.28.6
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/eslint-parser</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.28.6 (2026-01-12)</h2>
<p>Thanks <a href="https://github.com/kadhirash"><code>@​kadhirash</code></a> and <a href="https://github.com/kolvian"><code>@​kolvian</code></a> for your first PRs!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-cli</code>, <code>babel-code-frame</code>, <code>babel-core</code>, <code>babel-helper-check-duplicate-nodes</code>, <code>babel-helper-fixtures</code>, <code>babel-helper-plugin-utils</code>, <code>babel-node</code>, <code>babel-plugin-transform-flow-comments</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-property-mutators</code>, <code>babel-preset-env</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17589">#17589</a> Improve Unicode handling in code-frame tokenizer (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17556">#17556</a> fix: <code>transform-regenerator</code> correctly handles scope (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-react-jsx</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17538">#17538</a> fix: Keep jsx comments (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-core</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17606">#17606</a> Polish(standalone): improve message on invalid preset/plugin (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-plugin-bugfix-v8-static-class-fields-redefine-readonly</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-proposal-import-attributes-to-assertions</code>, <code>babel-plugin-proposal-import-wasm-source</code>, <code>babel-plugin-syntax-async-do-expressions</code>, <code>babel-plugin-syntax-decorators</code>, <code>babel-plugin-syntax-destructuring-private</code>, <code>babel-plugin-syntax-do-expressions</code>, <code>babel-plugin-syntax-explicit-resource-management</code>, <code>babel-plugin-syntax-export-default-from</code>, <code>babel-plugin-syntax-flow</code>, <code>babel-plugin-syntax-function-bind</code>, <code>babel-plugin-syntax-function-sent</code>, <code>babel-plugin-syntax-import-assertions</code>, <code>babel-plugin-syntax-import-attributes</code>, <code>babel-plugin-syntax-import-defer</code>, <code>babel-plugin-syntax-import-source</code>, <code>babel-plugin-syntax-jsx</code>, <code>babel-plugin-syntax-module-blocks</code>, <code>babel-plugin-syntax-optional-chaining-assign</code>, <code>babel-plugin-syntax-partial-application</code>, <code>babel-plugin-syntax-pipeline-operator</code>, <code>babel-plugin-syntax-throw-expressions</code>, <code>babel-plugin-syntax-typescript</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-class-properties</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-plugin-transform-dotall-regex</code>, <code>babel-plugin-transform-duplicate-named-capturing-groups-regex</code>, <code>babel-plugin-transform-explicit-resource-management</code>, <code>babel-plugin-transform-exponentiation-operator</code>, <code>babel-plugin-transform-json-strings</code>, <code>babel-plugin-transform-logical-assignment-operators</code>, <code>babel-plugin-transform-nullish-coalescing-operator</code>, <code>babel-plugin-transform-numeric-separator</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-catch-binding</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-private-methods</code>, <code>babel-plugin-transform-private-property-in-object</code>, <code>babel-plugin-transform-regexp-modifiers</code>, <code>babel-plugin-transform-unicode-property-regex</code>, <code>babel-plugin-transform-unicode-sets-regex</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17580">#17580</a> Allow Babel 8 in compatible Babel 7 plugins (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-plugin-transform-react-jsx</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17555">#17555</a> perf: Use lighter traversal for jsx <code>__source,__self</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 7</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Eliot Pontarelli (<a href="https://github.com/kolvian"><code>@​kolvian</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Kadhirash Sivakumar (<a href="https://github.com/kadhirash"><code>@​kadhirash</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
<li>coderaiser (<a href="https://github.com/coderaiser"><code>@​coderaiser</code></a>)</li>
</ul>
<h2>v7.28.5 (2025-10-23)</h2>
<p>Thank you <a href="https://github.com/CO0Ki3"><code>@​CO0Ki3</code></a>, <a href="https://github.com/Olexandr88"><code>@​Olexandr88</code></a>, and <a href="https://github.com/youthfulhps"><code>@​youthfulhps</code></a> for your first PRs!</p>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17446">#17446</a> Allow <code>Runtime Errors for Function Call Assignment Targets</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-validator-identifier</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17501">#17501</a> fix: update identifier to unicode 17 (<a href="https://github.com/fisker"><code>@​fisker</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17534">#17534</a> Allow mixing private destructuring and rest (<a href="https://github.com/CO0Ki3"><code>@​CO0Ki3</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17521">#17521</a> Improve <code>@babel/parser</code> error typing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17491">#17491</a> fix: improve ts-only declaration parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-discard-binding</code>, <code>babel-plugin-transform-destructuring</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/d7f400889567ae18ef9ac41b024b5120f6060e17"><code>d7f4008</code></a> v7.28.6</li>
<li><a href="https://github.com/babel/babel/commit/beea88c2bddd158a52e3c08739f47cbca970beea"><code>beea88c</code></a> [babel 8] Rename <code>TSImportType.argument</code> to <code>.source</code> (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17610">#17610</a>)</li>
<li><a href="https://github.com/babel/babel/commit/61647ae2397c82c3c71f077b5ab109106a5cac0f"><code>61647ae</code></a> v7.28.5</li>
<li><a href="https://github.com/babel/babel/commit/42cb285b59fc99a8102d69bef6223b75617e9f46"><code>42cb285</code></a> Improve <code>@babel/core</code> types (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17404">#17404</a>)</li>
<li><a href="https://github.com/babel/babel/commit/e7031b725d6a478f4fe14a3cadd0ed2c62a46017"><code>e7031b7</code></a> [Babel 8] Treat <code>allowSuperOutsideMethod</code> as top-level only (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17505">#17505</a>)</li>
<li><a href="https://github.com/babel/babel/commit/35055e392079a65830b7bf5b1d1c1fc4de90a78f"><code>35055e3</code></a> v7.28.4</li>
<li><a href="https://github.com/babel/babel/commit/f04083a70573804935797e5a4d7d8f647d30a59a"><code>f04083a</code></a> [Babel 8] Align TSMappedType AST (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17479">#17479</a>)</li>
<li><a href="https://github.com/babel/babel/commit/432a7ffbff568efb608a5ddd8e87aea39c76bdad"><code>432a7ff</code></a> fix(parser/typescript): parse <code>import(&quot;./a&quot;, {with:{},})</code> (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17465">#17465</a>)</li>
<li><a href="https://github.com/babel/babel/commit/e77e3b02d731da53463c5c80c66858cbb0c6dcfb"><code>e77e3b0</code></a> move eslint-{parser,plugin} docs to the website (<a href="https://github.com/babel/babel/tree/HEAD/eslint/babel-eslint-parser/issues/17448">#17448</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.28.6/eslint/babel-eslint-parser">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for <code>@​babel/eslint-parser</code> since your current version.</p>
</details>
<br />

Updates `@babel/plugin-transform-runtime` from 7.28.3 to 7.29.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/plugin-transform-runtime</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.29.0 (2026-01-31)</h2>
<p>Thanks <a href="https://github.com/simbahax"><code>@​simbahax</code></a> for your first PR!</p>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17750">#17750</a> [7.x backport] Add attributes import declaration builder (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17663">#17663</a> [7.x backport] feat(standalone): export async transform (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17725">#17725</a> [7.x backport] feat: read standalone targets from data-targets (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17765">#17765</a> fix(parser): correctly parse type assertions in <code>extends</code> clause (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17723">#17723</a> [7.x backport] fix(parser): improve super type argument parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17708">#17708</a> fix(traverse): provide a hub when traversing a File or Program and no parentPath is given (<a href="https://github.com/simbahax"><code>@​simbahax</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-block-scoping</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17737">#17737</a> [7.x backport] fix: Rename switch discriminant references when body creates shadowing variable (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-generator</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17642">#17642</a> [Babel 7] Improve generator performance (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 6</h4>
<ul>
<li>David (<a href="https://github.com/simbahax"><code>@​simbahax</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
<li><a href="https://github.com/magic-akari"><code>@​magic-akari</code></a></li>
</ul>
<h2>v7.28.6 (2026-01-12)</h2>
<p>Thanks <a href="https://github.com/kadhirash"><code>@​kadhirash</code></a> and <a href="https://github.com/kolvian"><code>@​kolvian</code></a> for your first PRs!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-cli</code>, <code>babel-code-frame</code>, <code>babel-core</code>, <code>babel-helper-check-duplicate-nodes</code>, <code>babel-helper-fixtures</code>, <code>babel-helper-plugin-utils</code>, <code>babel-node</code>, <code>babel-plugin-transform-flow-comments</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-property-mutators</code>, <code>babel-preset-env</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17589">#17589</a> Improve Unicode handling in code-frame tokenizer (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17556">#17556</a> fix: <code>transform-regenerator</code> correctly handles scope (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-react-jsx</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17538">#17538</a> fix: Keep jsx comments (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-core</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17606">#17606</a> Polish(standalone): improve message on invalid preset/plugin (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-plugin-bugfix-v8-static-class-fields-redefine-readonly</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-proposal-import-attributes-to-assertions</code>, <code>babel-plugin-proposal-import-wasm-source</code>, <code>babel-plugin-syntax-async-do-expressions</code>, <code>babel-plugin-syntax-decorators</code>, <code>babel-plugin-syntax-destructuring-private</code>, <code>babel-plugin-syntax-do-expressions</code>, <code>babel-plugin-syntax-explicit-resource-management</code>, <code>babel-plugin-syntax-export-default-from</code>, <code>babel-plugin-syntax-flow</code>, <code>babel-plugin-syntax-function-bind</code>, <code>babel-plugin-syntax-function-sent</code>, <code>babel-plugin-syntax-import-assertions</code>, <code>babel-plugin-syntax-import-attributes</code>, <code>babel-plugin-syntax-import-defer</code>, <code>babel-plugin-syntax-import-source</code>, <code>babel-plugin-syntax-jsx</code>, <code>babel-plugin-syntax-module-blocks</code>, <code>babel-plugin-syntax-optional-chaining-assign</code>, <code>babel-plugin-syntax-partial-application</code>, <code>babel-plugin-syntax-pipeline-operator</code>, <code>babel-plugin-syntax-throw-expressions</code>, <code>babel-plugin-syntax-typescript</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-class-properties</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-plugin-transform-dotall-regex</code>, <code>babel-plugin-transform-duplicate-named-capturing-groups-regex</code>, <code>babel-plugin-transform-explicit-resource-management</code>, <code>babel-plugin-transform-exponentiation-operator</code>, <code>babel-plugin-transform-json-strings</code>, <code>babel-plugin-transform-logical-assignment-operators</code>, <code>babel-plugin-transform-nullish-coalescing-operator</code>, <code>babel-plugin-transform-numeric-separator</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-catch-binding</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-private-methods</code>, <code>babel-plugin-transform-private-property-in-object</code>, <code>babel-plugin-transform-regexp-modifiers</code>, <code>babel-plugin-transform-unicode-property-regex</code>, <code>babel-plugin-transform-unicode-sets-regex</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17580">#17580</a> Allow Babel 8 in compatible Babel 7 plugins (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
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
<li><a href="https://github.com/babel/babel/commit/aa8394e454337d118ac3d40bfa3ee1a3cb3f3ed2"><code>aa8394e</code></a> v7.29.0</li>
<li><a href="https://github.com/babel/babel/commit/0053db620c05acf0036f593b5aaf4e372daa79d0"><code>0053db6</code></a> Update polyfill packages (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17727">#17727</a>)</li>
<li><a href="https://github.com/babel/babel/commit/f3a22268bdc4fc6748cbc2be718a4d1090bdaf00"><code>f3a2226</code></a> [babel 7] Delete Babel 8 fixtures (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17729">#17729</a>)</li>
<li><a href="https://github.com/babel/babel/commit/6287e94600ea3ff60450696cbce0bed82435d905"><code>6287e94</code></a> Run transform-runtime tests also in Babel 8 (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17615">#17615</a>)</li>
<li><a href="https://github.com/babel/babel/commit/61647ae2397c82c3c71f077b5ab109106a5cac0f"><code>61647ae</code></a> v7.28.5</li>
<li><a href="https://github.com/babel/babel/commit/42cb285b59fc99a8102d69bef6223b75617e9f46"><code>42cb285</code></a> Improve <code>@babel/core</code> types (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17404">#17404</a>)</li>
<li><a href="https://github.com/babel/babel/commit/85cafbe9cd805b31cd7e1d9b6b8ddaf76f09f97e"><code>85cafbe</code></a> Type check runtime scripts (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17522">#17522</a>)</li>
<li><a href="https://github.com/babel/babel/commit/3fc8d05093c1d2a017b973905625b847877f8934"><code>3fc8d05</code></a> [babel 8] Update default <code>@babel/runtime</code> version (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17512">#17512</a>)</li>
<li><a href="https://github.com/babel/babel/commit/1a6890949aa4736d6b6b2102b0e3e576c2261df5"><code>1a68909</code></a> [babel 8] Remove <code>semver</code> dependency from transform-runtime (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-plugin-transform-runtime/issues/17511">#17511</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.0/packages/babel-plugin-transform-runtime">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for <code>@​babel/plugin-transform-runtime</code> since your current version.</p>
</details>
<br />

Updates `@babel/preset-env` from 7.28.3 to 7.29.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-env</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.29.0 (2026-01-31)</h2>
<p>Thanks <a href="https://github.com/simbahax"><code>@​simbahax</code></a> for your first PR!</p>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17750">#17750</a> [7.x backport] Add attributes import declaration builder (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17663">#17663</a> [7.x backport] feat(standalone): export async transform (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17725">#17725</a> [7.x backport] feat: read standalone targets from data-targets (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17765">#17765</a> fix(parser): correctly parse type assertions in <code>extends</code> clause (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17723">#17723</a> [7.x backport] fix(parser): improve super type argument parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17708">#17708</a> fix(traverse): provide a hub when traversing a File or Program and no parentPath is given (<a href="https://github.com/simbahax"><code>@​simbahax</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-block-scoping</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17737">#17737</a> [7.x backport] fix: Rename switch discriminant references when body creates shadowing variable (<a href="https://github.com/magic-akari"><code>@​magic-akari</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-generator</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17642">#17642</a> [Babel 7] Improve generator performance (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 6</h4>
<ul>
<li>David (<a href="https://github.com/simbahax"><code>@​simbahax</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
<li><a href="https://github.com/magic-akari"><code>@​magic-akari</code></a></li>
</ul>
<h2>v7.28.6 (2026-01-12)</h2>
<p>Thanks <a href="https://github.com/kadhirash"><code>@​kadhirash</code></a> and <a href="https://github.com/kolvian"><code>@​kolvian</code></a> for your first PRs!</p>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-cli</code>, <code>babel-code-frame</code>, <code>babel-core</code>, <code>babel-helper-check-duplicate-nodes</code>, <code>babel-helper-fixtures</code>, <code>babel-helper-plugin-utils</code>, <code>babel-node</code>, <code>babel-plugin-transform-flow-comments</code>, <code>babel-plugin-transform-modules-commonjs</code>, <code>babel-plugin-transform-property-mutators</code>, <code>babel-preset-env</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17589">#17589</a> Improve Unicode handling in code-frame tokenizer (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-regenerator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17556">#17556</a> fix: <code>transform-regenerator</code> correctly handles scope (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-react-jsx</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17538">#17538</a> fix: Keep jsx comments (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-core</code>, <code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17606">#17606</a> Polish(standalone): improve message on invalid preset/plugin (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-plugin-bugfix-v8-static-class-fields-redefine-readonly</code>, <code>babel-plugin-proposal-decorators</code>, <code>babel-plugin-proposal-import-attributes-to-assertions</code>, <code>babel-plugin-proposal-import-wasm-source</code>, <code>babel-plugin-syntax-async-do-expressions</code>, <code>babel-plugin-syntax-decorators</code>, <code>babel-plugin-syntax-destructuring-private</code>, <code>babel-plugin-syntax-do-expressions</code>, <code>babel-plugin-syntax-explicit-resource-management</code>, <code>babel-plugin-syntax-export-default-from</code>, <code>babel-plugin-syntax-flow</code>, <code>babel-plugin-syntax-function-bind</code>, <code>babel-plugin-syntax-function-sent</code>, <code>babel-plugin-syntax-import-assertions</code>, <code>babel-plugin-syntax-import-attributes</code>, <code>babel-plugin-syntax-import-defer</code>, <code>babel-plugin-syntax-import-source</code>, <code>babel-plugin-syntax-jsx</code>, <code>babel-plugin-syntax-module-blocks</code>, <code>babel-plugin-syntax-optional-chaining-assign</code>, <code>babel-plugin-syntax-partial-application</code>, <code>babel-plugin-syntax-pipeline-operator</code>, <code>babel-plugin-syntax-throw-expressions</code>, <code>babel-plugin-syntax-typescript</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-plugin-transform-async-to-generator</code>, <code>babel-plugin-transform-class-properties</code>, <code>babel-plugin-transform-class-static-block</code>, <code>babel-plugin-transform-dotall-regex</code>, <code>babel-plugin-transform-duplicate-named-capturing-groups-regex</code>, <code>babel-plugin-transform-explicit-resource-management</code>, <code>babel-plugin-transform-exponentiation-operator</code>, <code>babel-plugin-transform-json-strings</code>, <code>babel-plugin-transform-logical-assignment-operators</code>, <code>babel-plugin-transform-nullish-coalescing-operator</code>, <code>babel-plugin-transform-numeric-separator</code>, <code>babel-plugin-transform-object-rest-spread</code>, <code>babel-plugin-transform-optional-catch-binding</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-plugin-transform-private-methods</code>, <code>babel-plugin-transform-private-property-in-object</code>, <code>babel-plugin-transform-regexp-modifiers</code>, <code>babel-plugin-transform-unicode-property-regex</code>, <code>babel-plugin-transform-unicode-sets-regex</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17580">#17580</a> Allow Babel 8 in compatible Babel 7 plugins (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
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
<li><a href="https://github.com/babel/babel/commit/aa8394e454337d118ac3d40bfa3ee1a3cb3f3ed2"><code>aa8394e</code></a> v7.29.0</li>
<li><a href="https://github.com/babel/babel/commit/0053db620c05acf0036f593b5aaf4e372daa79d0"><code>0053db6</code></a> Update polyfill packages (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17727">#17727</a>)</li>
<li><a href="https://github.com/babel/babel/commit/f3a22268bdc4fc6748cbc2be718a4d1090bdaf00"><code>f3a2226</code></a> [babel 7] Delete Babel 8 fixtures (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17729">#17729</a>)</li>
<li><a href="https://github.com/babel/babel/commit/d7f400889567ae18ef9ac41b024b5120f6060e17"><code>d7f4008</code></a> v7.28.6</li>
<li><a href="https://github.com/babel/babel/commit/99dcba5e71de3bd81ce14077cfa5b6df58e9b177"><code>99dcba5</code></a> chore: enable some ts-eslint rules (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17592">#17592</a>)</li>
<li><a href="https://github.com/babel/babel/commit/c92c4919771105140015167f25f7bacac77c90d9"><code>c92c491</code></a> Improve Unicode handling in code-frame tokenizer (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17589">#17589</a>)</li>
<li><a href="https://github.com/babel/babel/commit/61647ae2397c82c3c71f077b5ab109106a5cac0f"><code>61647ae</code></a> v7.28.5</li>
<li><a href="https://github.com/babel/babel/commit/42cb285b59fc99a8102d69bef6223b75617e9f46"><code>42cb285</code></a> Improve <code>@babel/core</code> types (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17404">#17404</a>)</li>
<li><a href="https://github.com/babel/babel/commit/ae363aed504a91f1ac0b79ad46dbd072658d364a"><code>ae363ae</code></a> chore: Fix typo in variable name (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17535">#17535</a>)</li>
<li><a href="https://github.com/babel/babel/commit/1edfcaa48f5d5c6aaf1345a03deb106cefff89a8"><code>1edfcaa</code></a> Update compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17487">#17487</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.0/packages/babel-preset-env">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for <code>@​babel/preset-env</code> since your current version.</p>
</details>
<br />

Updates `@babel/preset-react` from 7.27.1 to 7.28.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/preset-react</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.28.5 (2025-10-23)</h2>
<p>Thank you <a href="https://github.com/CO0Ki3"><code>@​CO0Ki3</code></a>, <a href="https://github.com/Olexandr88"><code>@​Olexandr88</code></a>, and <a href="https://github.com/youthfulhps"><code>@​youthfulhps</code></a> for your first PRs!</p>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17446">#17446</a> Allow <code>Runtime Errors for Function Call Assignment Targets</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-validator-identifier</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17501">#17501</a> fix: update identifier to unicode 17 (<a href="https://github.com/fisker"><code>@​fisker</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-plugin-proposal-destructuring-private</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17534">#17534</a> Allow mixing private destructuring and rest (<a href="https://github.com/CO0Ki3"><code>@​CO0Ki3</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17521">#17521</a> Improve <code>@babel/parser</code> error typing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17491">#17491</a> fix: improve ts-only declaration parsing (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-proposal-discard-binding</code>, <code>babel-plugin-transform-destructuring</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17519">#17519</a> fix: <code>rest</code> correctly returns plain array (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-helper-member-expression-to-functions</code>, <code>babel-plugin-transform-block-scoping</code>, <code>babel-plugin-transform-optional-chaining</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17503">#17503</a> Fix <code>JSXIdentifier</code> handling in <code>isReferencedIdentifier</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17504">#17504</a> fix: ensure scope.push register in anonymous fn (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:house: Internal</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17494">#17494</a> Type checking babel-types scripts (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17490">#17490</a> Faster finding of locations in <code>buildCodeFrameError</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 8</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Byeongho Yoo (<a href="https://github.com/youthfulhps"><code>@​youthfulhps</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Hyeon Dokko (<a href="https://github.com/CO0Ki3"><code>@​CO0Ki3</code></a>)</li>
<li>Nicolò Ribaudo (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
<li><a href="https://github.com/Olexandr88"><code>@​Olexandr88</code></a></li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
<li>fisker Cheung (<a href="https://github.com/fisker"><code>@​fisker</code></a>)</li>
</ul>
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
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/preset-react</code>'s changelog</a>.</em></p>
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
<p>This file contains the changelog starting from v8.0.0-alpha.0.</p>
<ul>
<li>See <a href="https://github.com/babel/babel/blob/main/.github/CHANGELOG-v7.15.0-v7.28.5.md">CHANGELOG - v7.15.0 to v7.28.5</a> for v7.15.0 to v7.28.5 changes (the last common release between the v8 and v7 release lines was v7.28.5).</li>
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
<h2>v8.0.0-rc.2 (2026-02-15)</h2>
<h4>:boom: Breaking Change</h4>
<ul>
<li>Other
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17766">#17766</a> Remove unused code for old ESLint versions (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-code-frame</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17772">#17772</a> Remove deprecated default export from <code>@babel/code-frame</code> (<a href="https://github.com/fisker"><code>@​fisker</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helpers</code>, <code>babel-plugin-transform-async-generator-functions</code>, <code>babel-runtime-corejs3</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17797">#17797</a> fix: Properly handle <code>await</code> in <code>finally</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17796">#17796</a> Support ESLint 10 (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-preset-env</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17787">#17787</a> Fix: preset-env include/exclude should accept bugfix plugins (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17781">#17781</a> fix: preserve trailing comma in optional call args (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li><a href="https://redirect.github.com/babel/babel/pull/17774">#17774</a> Fix <code>undefined</code> indentation when exactly 64 indents (<a href="https://github.com/YoussefHenna"><code>@​YoussefHenna</code></a>)</li>
</ul>
</li>
<li><code>babel-standalone</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17770">#17770</a> fix: ensure <code>targets.esmodules</code> is validated (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:nail_care: Polish</h4>
<ul>
<li><code>babel-core</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/61647ae2397c82c3c71f077b5ab109106a5cac0f"><code>61647ae</code></a> v7.28.5</li>
<li><a href="https://github.com/babel/babel/commit/42cb285b59fc99a8102d69bef6223b75617e9f46"><code>42cb285</code></a> Improve <code>@babel/core</code> types (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-react/issues/17404">#17404</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.28.5/packages/babel-preset-react">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by [GitHub Actions](<a href="https://www.npmjs.com/~GitHub">https://www.npmjs.com/~GitHub</a> Actions), a new releaser for <code>@​babel/preset-react</code> since your current version.</p>
</details>
<br />

Updates `babel-jest` from 29.7.0 to 30.2.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">babel-jest's releases</a>.</em></p>
<blockquote>
<h2>30.2.0</h2>
<h3>Chore &amp; Maintenance</h3>
<ul>
<li><code>[*]</code> Update example repo for testing React Native projects (<a href="https://redirect.github.com/jestjs/jest/pull/15832">#15832</a>)</li>
<li><code>[*]</code> Update <code>jest-watch-typeahead</code> to v3 (<a href="https://redirect.github.com/jestjs/jest/pull/15830">#15830</a>)</li>
</ul>
<h2>Features</h2>
<ul>
<li><code>[jest-environment-jsdom-abstract]</code> Add support for JSDOM v27 (<a href="https://redirect.github.com/jestjs/jest/pull/15834">#15834</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[babel-jest]</code> Export the <code>TransformerConfig</code> interface (<a href="https://redirect.github.com/jestjs/jest/pull/15820">#15820</a>)</li>
<li><code>[jest-config]</code> Fix <code>jest.config.ts</code> with TS loader specified in docblock pragma (<a href="https://redirect.github.com/jestjs/jest/pull/15839">#15839</a>)</li>
</ul>
<h2>30.1.3</h2>
<h3>Fixes</h3>
<ul>
<li>Fix <code>unstable_mockModule</code> with <code>node:</code> prefixed core modules.</li>
</ul>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">babel-jest's changelog</a>.</em></p>
<blockquote>
<h2>30.2.0</h2>
<h3>Chore &amp; Maintenance</h3>
<ul>
<li><code>[*]</code> Update example repo for testing React Native projects (<a href="https://redirect.github.com/jestjs/jest/pull/15832">#15832</a>)</li>
<li><code>[*]</code> Update <code>jest-watch-typeahead</code> to v3 (<a href="https://redirect.github.com/jestjs/jest/pull/15830">#15830</a>)</li>
</ul>
<h2>Features</h2>
<ul>
<li><code>[jest-environment-jsdom-abstract]</code> Add support for JSDOM v27 (<a href="https://redirect.github.com/jestjs/jest/pull/15834">#15834</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-matcher-utils]</code> Fix infinite recursion with self-referential getters in <code>deepCyclicCopyReplaceable</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15831">#15831</a>)</li>
<li><code>[babel-jest]</code> Export the <code>TransformerConfig</code> interface (<a href="https://redirect.github.com/jestjs/jest/pull/15820">#15820</a>)</li>
<li><code>[jest-config]</code> Fix <code>jest.config.ts</code> with TS loader specified in docblock pragma (<a href="https://redirect.github.com/jestjs/jest/pull/15839">#15839</a>)</li>
</ul>
<h2>30.1.3</h2>
<h3>Fixes</h3>
<ul>
<li>Fix <code>unstable_mockModule</code> with <code>node:</code> prefixed core modules.</li>
</ul>
<h2>30.1.2</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Correct snapshot header regexp to work with newline across OSes (<a href="https://redirect.github.com/jestjs/jest/pull/15803">#15803</a>)</li>
</ul>
<h2>30.1.1</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot warning not handling Windows end-of-line sequences (<a href="https://redirect.github.com/jestjs/jest/pull/15800">#15800</a>)</li>
<li><code>[jest-snapshot-utils]</code> Improve messaging about goo.gl snapshot link change (<a href="https://redirect.github.com/jestjs/jest/pull/15821">#15821</a>)</li>
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
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jestjs/jest/commit/855864e3f9751366455246790be2bf912d4d0dac"><code>855864e</code></a> v30.2.0</li>
<li><a href="https://github.com/jestjs/jest/commit/d2a2491b642a015b9f176418b1175fcf2dbb905f"><code>d2a2491</code></a> Support Babel 8 in plugins and presets (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/15750">#15750</a>)</li>
<li><a href="https://github.com/jestjs/jest/commit/ddf97f7f891965041b76a294b64cae7fd3fcc78d"><code>ddf97f7</code></a> Export TransformerConfig interface from <code>babel-jest</code> (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/15820">#15820</a>)</li>
<li><a href="https://github.com/jestjs/jest/commit/15e3e7cb0bb91c7466f60ce854287ca5b711bd69"><code>15e3e7c</code></a> Update babel-plugin-istanbul to 7.0.1 (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/15813">#15813</a>)</li>
<li><a href="https://github.com/jestjs/jest/commit/ebfa31cc9787303e8698a1a029a162a18e8974aa"><code>ebfa31c</code></a> v30.1.2</li>
<li><a href="https://github.com/jestjs/jest/commit/d347c0f3f87f976a1dbd9761d503e45f5ced2a7e"><code>d347c0f</code></a> v30.1.1</li>
<li><a href="https://github.com/jestjs/jest/commit/4d5f41d0885c1d9630c81b4fd47f74ab0615e18f"><code>4d5f41d</code></a> v30.1.0</li>
<li><a href="https://github.com/jestjs/jest/commit/22236cf58b66039f81893537c90dee290bab427f"><code>22236cf</code></a> v30.0.5</li>
<li><a href="https://github.com/jestjs/jest/commit/f4296d2bc85c1405f84ddf613a25d0bc3766b7e5"><code>f4296d2</code></a> v30.0.4</li>
<li><a href="https://github.com/jestjs/jest/commit/393acbfac31f64bb38dff23c89224797caded83c"><code>393acbf</code></a> v30.0.2</li>
<li>Additional commits viewable in <a href="https://github.com/jestjs/jest/commits/v30.2.0/packages/babel-jest">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Discussion

### Comment by @codecov-commenter on February 16, 2026 at 03:35 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1511?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.50%. Comparing base ([`6fb82f8`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6fb82f8c94950b1c0440f115e45f410a5a957313?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`2ee2800`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/2ee2800b7c5b89d08e735c49a4705abfb9f176fe?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1511      +/-   ##
==========================================
- Coverage   62.53%   62.50%   -0.04%     
==========================================
  Files         127      127              
  Lines        3310     3312       +2     
  Branches      899      892       -7     
==========================================
  Hits         2070     2070              
- Misses       1120     1122       +2     
  Partials      120      120              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1511?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on February 18, 2026 at 02:32 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1511*
