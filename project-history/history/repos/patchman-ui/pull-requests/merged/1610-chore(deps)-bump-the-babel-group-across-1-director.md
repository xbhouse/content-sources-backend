---
type: pull_request
number: 1610
title: "chore(deps): bump the babel group across 1 directory with 2 updates"
state: merged
author: dependabot
created: 2026-05-08T14:27:38Z
updated: 2026-05-08T14:58:08Z
closed: 2026-05-08T14:57:58Z
merged: 2026-05-08T14:57:58Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-f6842ae598
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1610
---

# Pull Request #1610: chore(deps): bump the babel group across 1 directory with 2 updates

**Author**: @dependabot
**Created**: May 08, 2026 at 02:27 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-f6842ae598`

## Description

Bumps the babel group with 2 updates in the / directory: [@babel/preset-env](https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env) and [babel-jest](https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest).

Updates `@babel/preset-env` from 7.29.2 to 7.29.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases">@​babel/preset-env's releases</a>.</em></p>
<blockquote>
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
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17782">#17782</a> Improve trailing comma comment handling (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:memo: Documentation</h4>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17847">#17847</a> Replace npmjs.com links with npmx.dev (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-helper-import-to-platform-api</code>, <code>babel-plugin-proposal-import-wasm-source</code>, <code>babel-plugin-transform-json-modules</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17818">#17818</a> Load async Wasm and JSON imports in parallel (<a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a>)</li>
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
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/3cd910d838332b988ed83bdd2ddc22e849e7ea5d"><code>3cd910d</code></a> v7.29.5</li>
<li><a href="https://github.com/babel/babel/commit/3d399f8c8c1e5308bb25e11947d90a111399ac0d"><code>3d399f8</code></a> [7.x backport]docs(preset-env): update CONTRIBUTING.md (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17976">#17976</a>)</li>
<li><a href="https://github.com/babel/babel/commit/183db7bc040a68057489f8981d02962345a322ed"><code>183db7b</code></a> v7.29.3</li>
<li><a href="https://github.com/babel/babel/commit/268f246f21e51b2204ba6dc5349055504cc7420d"><code>268f246</code></a> Add bugfix plugin for Safari array rest destructuring bug (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17788">#17788</a>)</li>
<li><a href="https://github.com/babel/babel/commit/f8524d80799e136313e55da0468777a57d1bf6b6"><code>f8524d8</code></a> Update compat data (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-preset-env/issues/17686">#17686</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.29.5/packages/babel-preset-env">compare view</a></li>
</ul>
</details>
<br />

Updates `babel-jest` from 30.3.0 to 30.4.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">babel-jest's releases</a>.</em></p>
<blockquote>
<h2>v30.4.1</h2>
<h2>Features</h2>
<ul>
<li><code>[jest-config, jest-core, jest-runner, jest-schemas, jest-types]</code> Allow custom runner configuration options via tuple format <code>['runner-path', {options}]</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16141">#16141</a>)</li>
</ul>
<h2>Fixes</h2>
<ul>
<li><code>[jest-runtime]</code> Align CJS-from-ESM default export with Node: <code>module.exports</code> is always the ESM default, <code>__esModule</code> unwrapping is no longer applied (<a href="https://redirect.github.com/jestjs/jest/pull/16143">#16143</a>)</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/jestjs/jest/compare/v30.4.0...v30.4.1">https://github.com/jestjs/jest/compare/v30.4.0...v30.4.1</a></p>
<h2>v30.4.0</h2>
<p>Big release! 😀</p>
<p>Main feature is a rewrite of our custom runtime in preparation for stabilisation of native support of ESM. As part of that work <code>require(esm)</code> module is now supported on Node 24.9+ (still requires <code>--experimental-vm-modules</code> like before).</p>
<p>In addition we now support fake timers for the recently released <code>Temporal</code> API in Node v26.</p>
<p>React 19 is also supported properly in <code>pretty-format</code>, meaning snapshots of React components now work like they should.</p>
<p>Due to all the changes, there might be regressions that snuck in. Please report them!</p>
<p>Full list of changes below</p>
<h2>Features</h2>
<ul>
<li><code>[babel-jest]</code> Support collecting coverage from <code>.mts</code>, <code>.cts</code> (and other) files (<a href="https://redirect.github.com/jestjs/jest/pull/15994">#15994</a>)</li>
<li><code>[jest-circus, jest-cli, jest-config, jest-core, jest-jasmine2, jest-types]</code> Add <code>--collect-tests</code> flag to discover and list tests without executing them (<a href="https://redirect.github.com/jestjs/jest/pull/16006">#16006</a>)</li>
<li><code>[jest-config, jest-runner, jest-worker]</code> Add <code>workerGracefulExitTimeout</code> config option to control how long workers are given to exit before being force-killed (<a href="https://redirect.github.com/jestjs/jest/pull/15984">#15984</a>)</li>
<li><code>[jest-config]</code> Add support for <code>jest.config.mts</code> as a valid configuration file (<a href="https://redirect.github.com/jestjs/jest/pull/16005">#16005</a>)</li>
<li><code>[jest-config, jest-core, jest-reporters, jest-runner]</code> <code>verbose</code> and <code>silent</code> can now be set per-project; the project-level value overrides the global value for that project's tests (<a href="https://redirect.github.com/jestjs/jest/pull/16133">#16133</a>)</li>
<li><code>[@jest/fake-timers]</code> Accept <code>Temporal.Duration</code> in <code>jest.advanceTimersByTime()</code> and <code>jest.advanceTimersByTimeAsync()</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16128">#16128</a>)</li>
<li><code>[@jest/fake-timers]</code> Accept <code>Temporal.Instant</code> and <code>Temporal.ZonedDateTime</code> in <code>jest.setSystemTime()</code> and <code>useFakeTimers({now})</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16128">#16128</a>)</li>
<li><code>[@jest/fake-timers]</code> Support faking <code>Temporal.Now.*</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16131">#16131</a>)</li>
<li><code>[jest-mock]</code> Add <code>clearMocksOnScope(scope)</code> on <code>ModuleMocker</code> for clearing every mock function exposed on a scope object (<a href="https://redirect.github.com/jestjs/jest/pull/16088">#16088</a>)</li>
<li><code>[jest-resolve]</code> Add <code>canResolveSync()</code> on <code>Resolver</code> so callers can detect when a user-configured resolver only exports an <code>async</code> hook (<a href="https://redirect.github.com/jestjs/jest/pull/16064">#16064</a>)</li>
<li><code>[jest-runtime]</code> Use synchronous <code>evaluate()</code> for ES modules without top-level <code>await</code> on Node versions that support it (v24.9+), and prefer the synchronous transform path when a sync transformer is configured (<a href="https://redirect.github.com/jestjs/jest/pull/16062">#16062</a>)</li>
<li><code>[jest-runtime]</code> Support <code>require()</code> of ES modules on Node v24.9+ (<a href="https://redirect.github.com/jestjs/jest/pull/16074">#16074</a>)</li>
<li><code>[jest-runtime]</code> Validate TC39 import attributes (<code>with { type: 'json' }</code>) on ESM imports (<a href="https://redirect.github.com/jestjs/jest/pull/16127">#16127</a>)</li>
<li><code>[@jest/transform]</code> Add <code>canTransformSync(filename)</code> on <code>ScriptTransformer</code> so callers can pick the sync vs async transform path (<a href="https://redirect.github.com/jestjs/jest/pull/16062">#16062</a>)</li>
<li><code>[jest-util]</code> Add <code>isError</code> helper (<a href="https://redirect.github.com/jestjs/jest/pull/16076">#16076</a>)</li>
<li><code>[pretty-format]</code> Support React 19 (<a href="https://redirect.github.com/jestjs/jest/pull/16123">#16123</a>)</li>
</ul>
<h2>Fixes</h2>
<ul>
<li><code>[expect-utils]</code> Fix <code>toStrictEqual</code> failing on <code>structuredClone</code> results due to cross-realm constructor mismatch (<a href="https://redirect.github.com/jestjs/jest/pull/15959">#15959</a>)</li>
<li><code>[@jest/expect-utils]</code> Prevent <code>toMatchObject</code>/subset matching from throwing when encountering exotic iterables (<a href="https://redirect.github.com/jestjs/jest/pull/15952">#15952</a>)</li>
<li><code>[fake-timers]</code> Convert <code>Date</code> to milliseconds before passing to <code>@sinonjs/fake-timers</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16029">#16029</a>)</li>
<li><code>[jest]</code> Export <code>GlobalConfig</code> and <code>ProjectConfig</code> TypeScript types (<a href="https://redirect.github.com/jestjs/jest/pull/16132">#16132</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">babel-jest's changelog</a>.</em></p>
<blockquote>
<h2>30.4.1</h2>
<h3>Features</h3>
<ul>
<li><code>[jest-config, jest-core, jest-runner, jest-schemas, jest-types]</code> Allow custom runner configuration options via tuple format <code>['runner-path', {options}]</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16141">#16141</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-runtime]</code> Align CJS-from-ESM default export with Node: <code>module.exports</code> is always the ESM default, <code>__esModule</code> unwrapping is no longer applied (<a href="https://redirect.github.com/jestjs/jest/pull/16143">#16143</a>)</li>
</ul>
<h2>30.4.0</h2>
<h3>Features</h3>
<ul>
<li><code>[babel-jest]</code> Support collecting coverage from <code>.mts</code>, <code>.cts</code> (and other) files (<a href="https://redirect.github.com/jestjs/jest/pull/15994">#15994</a>)</li>
<li><code>[jest-circus, jest-cli, jest-config, jest-core, jest-jasmine2, jest-types]</code> Add <code>--collect-tests</code> flag to discover and list tests without executing them (<a href="https://redirect.github.com/jestjs/jest/pull/16006">#16006</a>)</li>
<li><code>[jest-config, jest-runner, jest-worker]</code> Add <code>workerGracefulExitTimeout</code> config option to control how long workers are given to exit before being force-killed (<a href="https://redirect.github.com/jestjs/jest/pull/15984">#15984</a>)</li>
<li><code>[jest-config]</code> Add support for <code>jest.config.mts</code> as a valid configuration file (<a href="https://redirect.github.com/jestjs/jest/pull/16005">#16005</a>)</li>
<li><code>[jest-config, jest-core, jest-reporters, jest-runner]</code> <code>verbose</code> and <code>silent</code> can now be set per-project; the project-level value overrides the global value for that project's tests (<a href="https://redirect.github.com/jestjs/jest/pull/16133">#16133</a>)</li>
<li><code>[@jest/fake-timers]</code> Accept <code>Temporal.Duration</code> in <code>jest.advanceTimersByTime()</code> and <code>jest.advanceTimersByTimeAsync()</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16128">#16128</a>)</li>
<li><code>[@jest/fake-timers]</code> Accept <code>Temporal.Instant</code> and <code>Temporal.ZonedDateTime</code> in <code>jest.setSystemTime()</code> and <code>useFakeTimers({now})</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16128">#16128</a>)</li>
<li><code>[@jest/fake-timers]</code> Support faking <code>Temporal.Now.*</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16131">#16131</a>)</li>
<li><code>[jest-mock]</code> Add <code>clearMocksOnScope(scope)</code> on <code>ModuleMocker</code> for clearing every mock function exposed on a scope object (<a href="https://redirect.github.com/jestjs/jest/pull/16088">#16088</a>)</li>
<li><code>[jest-resolve]</code> Add <code>canResolveSync()</code> on <code>Resolver</code> so callers can detect when a user-configured resolver only exports an <code>async</code> hook (<a href="https://redirect.github.com/jestjs/jest/pull/16064">#16064</a>)</li>
<li><code>[jest-runtime]</code> Use synchronous <code>evaluate()</code> for ES modules without top-level <code>await</code> on Node versions that support it (v24.9+), and prefer the synchronous transform path when a sync transformer is configured (<a href="https://redirect.github.com/jestjs/jest/pull/16062">#16062</a>)</li>
<li><code>[jest-runtime]</code> Support <code>require()</code> of ES modules on Node v24.9+ (<a href="https://redirect.github.com/jestjs/jest/pull/16074">#16074</a>)</li>
<li><code>[jest-runtime]</code> Validate TC39 import attributes (<code>with { type: 'json' }</code>) on ESM imports (<a href="https://redirect.github.com/jestjs/jest/pull/16127">#16127</a>)</li>
<li><code>[@jest/transform]</code> Add <code>canTransformSync(filename)</code> on <code>ScriptTransformer</code> so callers can pick the sync vs async transform path (<a href="https://redirect.github.com/jestjs/jest/pull/16062">#16062</a>)</li>
<li><code>[jest-util]</code> Add <code>isError</code> helper (<a href="https://redirect.github.com/jestjs/jest/pull/16076">#16076</a>)</li>
<li><code>[pretty-format]</code> Support React 19 (<a href="https://redirect.github.com/jestjs/jest/pull/16123">#16123</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[expect-utils]</code> Fix <code>toStrictEqual</code> failing on <code>structuredClone</code> results due to cross-realm constructor mismatch (<a href="https://redirect.github.com/jestjs/jest/pull/15959">#15959</a>)</li>
<li><code>[@jest/expect-utils]</code> Prevent <code>toMatchObject</code>/subset matching from throwing when encountering exotic iterables (<a href="https://redirect.github.com/jestjs/jest/pull/15952">#15952</a>)</li>
<li><code>[fake-timers]</code> Convert <code>Date</code> to milliseconds before passing to <code>@sinonjs/fake-timers</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16029">#16029</a>)</li>
<li><code>[jest]</code> Export <code>GlobalConfig</code> and <code>ProjectConfig</code> TypeScript types (<a href="https://redirect.github.com/jestjs/jest/pull/16132">#16132</a>)</li>
<li><code>[jest-circus]</code> Prevent crash when <code>asyncError</code> is undefined for non-Error throws (<a href="https://redirect.github.com/jestjs/jest/pull/16003">#16003</a>)</li>
<li><code>[jest-circus, jest-jasmine2]</code> Include <code>Error.cause</code> in JSON <code>failureMessages</code> output (<a href="https://redirect.github.com/jestjs/jest/pull/15967">#15967</a>)</li>
<li><code>[jest-config]</code> Fix preset path resolution on Windows when the preset uses subpath <code>exports</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15961">#15961</a>)</li>
<li><code>[jest-config]</code> Allow <code>collectCoverage</code> and <code>coverageProvider</code> in project config without a validation warning (<a href="https://redirect.github.com/jestjs/jest/pull/16132">#16132</a>)</li>
<li><code>[jest-config]</code> Project config validator now emits &quot;is not supported in an individual project configuration&quot; instead of &quot;probably a typing mistake&quot; for known global-only options (<a href="https://redirect.github.com/jestjs/jest/pull/16132">#16132</a>)</li>
<li><code>[jest-environment-node]</code> Fix <code>--localstorage-file</code> warning on Node 25+ (<a href="https://redirect.github.com/jestjs/jest/pull/16086">#16086</a>)</li>
<li><code>[jest-reporters]</code> Apply global coverage threshold to unmatched pattern files in addition to glob/path thresholds (<a href="https://redirect.github.com/jestjs/jest/pull/16137">#16137</a>)</li>
<li><code>[jest-reporters, jest-runner, jest-runtime, jest-transform]</code> Fix coverage report not showing correct code coverage when using <code>projects</code> config option (<a href="https://redirect.github.com/jestjs/jest/pull/16140">#16140</a>)</li>
<li><code>[jest-runtime]</code> Resolve <code>expect</code> and <code>@jest/expect</code> from the internal module registry so test-file imports share the same <code>JestAssertionError</code> as the global <code>expect</code> (<a href="https://redirect.github.com/jestjs/jest/pull/16130">#16130</a>)</li>
<li><code>[jest-runtime]</code> Improve CJS-from-ESM interop: <code>__esModule</code>/Babel default unwrap, broader named-export coverage, and shared CJS singleton across importers (<a href="https://redirect.github.com/jestjs/jest/pull/16050">#16050</a>)</li>
<li><code>[jest-runtime]</code> Load <code>.js</code> files with ESM syntax but no <code>&quot;type&quot;:&quot;module&quot;</code> marker as native ESM (<a href="https://redirect.github.com/jestjs/jest/pull/16050">#16050</a>)</li>
<li><code>[jest-runtime]</code> Extend the <code>.js</code>-with-ESM-syntax fallback to <code>require()</code> on Node v24.9+ - falls back to <code>require(esm)</code> when the CJS parser rejects ESM syntax (<a href="https://redirect.github.com/jestjs/jest/pull/16078">#16078</a>)</li>
<li><code>[jest-runtime]</code> Fix deadlocks and double-evaluation in concurrent ESM and wasm imports (<a href="https://redirect.github.com/jestjs/jest/pull/16050">#16050</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jestjs/jest/commit/b3b4a09ed3005369dacc7466d1d2122797283785"><code>b3b4a09</code></a> v30.4.1</li>
<li><a href="https://github.com/jestjs/jest/commit/5cbb21e0b3037edb42e503ec1a1ce80efad40c20"><code>5cbb21e</code></a> v30.4.0</li>
<li><a href="https://github.com/jestjs/jest/commit/6886816ab4f2a14a2d53b3020ed02afb2cf0ac9c"><code>6886816</code></a> chore: enable node protocol in imports (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/16077">#16077</a>)</li>
<li><a href="https://github.com/jestjs/jest/commit/d6d44e8d794d8bceb6993c00fcedc95964fba961"><code>d6d44e8</code></a> feat(babel-jest): support collecting coverage from .mts, .cts (and other) fil...</li>
<li><a href="https://github.com/jestjs/jest/commit/a3ab83991bd52b0dc13368d9b2010fe63be59fae"><code>a3ab839</code></a> chore(deps): update babel monorepo to v8.0.0-rc.4 (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/15976">#15976</a>)</li>
<li><a href="https://github.com/jestjs/jest/commit/487f4e86abd0ccbbd8b24111b5699979d40a2731"><code>487f4e8</code></a> chore: remove unused ignore comment (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/16069">#16069</a>)</li>
<li>See full diff in <a href="https://github.com/jestjs/jest/commits/v30.4.1/packages/babel-jest">compare view</a></li>
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

### Comment by @codecov-commenter on May 08, 2026 at 02:31 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1610?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.34%. Comparing base ([`fa92f68`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/fa92f68a74cabf54ab8926d5d7f38f4a00bf217e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`0421fc6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/0421fc67bd2962030dd7f045663db157db153f0d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1610   +/-   ##
=======================================
  Coverage   76.34%   76.34%           
=======================================
  Files         103      103           
  Lines        3187     3187           
  Branches      693      698    +5     
=======================================
  Hits         2433     2433           
  Misses        675      675           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1610?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on May 08, 2026 at 02:45 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1610*
