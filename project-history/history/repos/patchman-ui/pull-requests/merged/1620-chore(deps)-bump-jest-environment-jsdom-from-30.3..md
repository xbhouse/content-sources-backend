---
type: pull_request
number: 1620
title: "chore(deps): bump jest-environment-jsdom from 30.3.0 to 30.4.1 in the jest group across 1 directory"
state: merged
author: dependabot
created: 2026-05-12T22:52:53Z
updated: 2026-05-13T11:38:03Z
closed: 2026-05-13T11:37:54Z
merged: 2026-05-13T11:37:54Z
base_branch: master
head_branch: dependabot/npm_and_yarn/jest-a3fe29250a
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1620
---

# Pull Request #1620: chore(deps): bump jest-environment-jsdom from 30.3.0 to 30.4.1 in the jest group across 1 directory

**Author**: @dependabot
**Created**: May 12, 2026 at 10:52 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/jest-a3fe29250a`

## Description

Bumps the jest group with 1 update in the / directory: [jest-environment-jsdom](https://github.com/jestjs/jest/tree/HEAD/packages/jest-environment-jsdom).

Updates `jest-environment-jsdom` from 30.3.0 to 30.4.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">jest-environment-jsdom's releases</a>.</em></p>
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
<p><em>Sourced from <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">jest-environment-jsdom's changelog</a>.</em></p>
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
<li>See full diff in <a href="https://github.com/jestjs/jest/commits/v30.4.1/packages/jest-environment-jsdom">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on May 12, 2026 at 10:55 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1620?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.34%. Comparing base ([`b17a6d0`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b17a6d0ab3b1eaf14441074252101b3dcfb7d104?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`308696e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/308696ee84a544ef34d07d568e7fe1ea5681ce11?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1620   +/-   ##
=======================================
  Coverage   76.34%   76.34%           
=======================================
  Files         103      103           
  Lines        3187     3187           
  Branches      693      693           
=======================================
  Hits         2433     2433           
  Misses        675      675           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1620?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on May 13, 2026 at 11:05 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1620*
