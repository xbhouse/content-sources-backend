---
type: pull_request
number: 1539
title: "chore(deps): bump the babel group with 2 updates"
state: merged
author: dependabot
created: 2026-03-10T18:43:37Z
updated: 2026-03-13T14:53:45Z
closed: 2026-03-13T14:53:35Z
merged: 2026-03-13T14:53:35Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel-a783533071
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1539
---

# Pull Request #1539: chore(deps): bump the babel group with 2 updates

**Author**: @dependabot
**Created**: March 10, 2026 at 06:43 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel-a783533071`

## Description

Bumps the babel group with 2 updates: [babel-jest](https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest) and [babel-loader](https://github.com/babel/babel-loader).

Updates `babel-jest` from 30.2.0 to 30.3.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">babel-jest's releases</a>.</em></p>
<blockquote>
<h2>v30.3.0</h2>
<h3>Features</h3>
<ul>
<li><code>[jest-config]</code> Add <code>defineConfig</code> and <code>mergeConfig</code> helpers for type-safe Jest config (<a href="https://redirect.github.com/jestjs/jest/pull/15844">#15844</a>)</li>
<li><code>[jest-fake-timers]</code> Add <code>setTimerTickMode</code> to configure how timers advance</li>
<li><code>[*]</code> Reduce token usage when run through LLMs (<a href="https://github.com/jestjs/jest/commit/3f17932061c0203999451e5852664093de876709"><code>3f17932</code></a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-config]</code> Keep CLI coverage output when using <code>--json</code> with <code>--outputFile</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15918">#15918</a>)</li>
<li><code>[jest-mock]</code> Use <code>Symbol</code> from test environment (<a href="https://redirect.github.com/jestjs/jest/pull/15858">#15858</a>)</li>
<li><code>[jest-reporters]</code> Fix issue where console output not displayed for GHA reporter even with <code>silent: false</code> option (<a href="https://redirect.github.com/jestjs/jest/pull/15864">#15864</a>)</li>
<li><code>[jest-runtime]</code> Fix issue where user cannot utilize dynamic import despite specifying <code>--experimental-vm-modules</code> Node option (<a href="https://redirect.github.com/jestjs/jest/pull/15842">#15842</a>)</li>
<li><code>[jest-test-sequencer]</code> Fix issue where failed tests due to compilation errors not getting re-executed even with <code>--onlyFailures</code> CLI option (<a href="https://redirect.github.com/jestjs/jest/pull/15851">#15851</a>)</li>
<li><code>[jest-util]</code> Make sure <code>process.features.require_module</code> is <code>false</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15867">#15867</a>)</li>
</ul>
<h3>Chore &amp; Maintenance</h3>
<ul>
<li><code>[*]</code> Replace remaining micromatch uses with picomatch</li>
<li><code>[deps]</code> Update to sinon/fake-timers v15</li>
<li><code>[docs]</code> Update V30 migration guide to notify users on <code>jest.mock()</code> work with case-sensitive path (<a href="https://redirect.github.com/jestjs/jest/pull/15849">#15849</a>)</li>
<li>Updated Twitter icon to match the latest brand guidelines (<a href="https://redirect.github.com/jestjs/jest/pull/15869">#15869</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">babel-jest's changelog</a>.</em></p>
<blockquote>
<h2>30.3.0</h2>
<h3>Features</h3>
<ul>
<li><code>[jest-config]</code> Add <code>defineConfig</code> and <code>mergeConfig</code> helpers for type-safe Jest config (<a href="https://redirect.github.com/jestjs/jest/pull/15844">#15844</a>)</li>
<li><code>[jest-fake-timers]</code> Add <code>setTimerTickMode</code> to configure how timers advance</li>
<li><code>[*]</code> Reduce token usage when run through LLMs (<a href="https://github.com/jestjs/jest/commit/3f17932061c0203999451e5852664093de876709"><code>3f17932</code></a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-config]</code> Keep CLI coverage output when using <code>--json</code> with <code>--outputFile</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15918">#15918</a>)</li>
<li><code>[jest-mock]</code> Use <code>Symbol</code> from test environment (<a href="https://redirect.github.com/jestjs/jest/pull/15858">#15858</a>)</li>
<li><code>[jest-reporters]</code> Fix issue where console output not displayed for GHA reporter even with <code>silent: false</code> option (<a href="https://redirect.github.com/jestjs/jest/pull/15864">#15864</a>)</li>
<li><code>[jest-runtime]</code> Fix issue where user cannot utilize dynamic import despite specifying <code>--experimental-vm-modules</code> Node option (<a href="https://redirect.github.com/jestjs/jest/pull/15842">#15842</a>)</li>
<li><code>[jest-test-sequencer]</code> Fix issue where failed tests due to compilation errors not getting re-executed even with <code>--onlyFailures</code> CLI option (<a href="https://redirect.github.com/jestjs/jest/pull/15851">#15851</a>)</li>
<li><code>[jest-util]</code> Make sure <code>process.features.require_module</code> is <code>false</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15867">#15867</a>)</li>
</ul>
<h3>Chore &amp; Maintenance</h3>
<ul>
<li><code>[*]</code> Replace remaining micromatch uses with picomatch</li>
<li><code>[deps]</code> Update to sinon/fake-timers v15</li>
<li><code>[docs]</code> Update V30 migration guide to notify users on <code>jest.mock()</code> work with case-sensitive path (<a href="https://redirect.github.com/jestjs/jest/pull/15849">#15849</a>)</li>
<li>Updated Twitter icon to match the latest brand guidelines (<a href="https://redirect.github.com/jestjs/jest/pull/15869">#15869</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jestjs/jest/commit/efb59c2e81083f8dc941f20d6d20a3af2dc8d068"><code>efb59c2</code></a> v30.3.0</li>
<li><a href="https://github.com/jestjs/jest/commit/486ae3d6d3ea4da3fccf5d39097f4cebcf132d4f"><code>486ae3d</code></a> chore: update docusaurus (<a href="https://github.com/jestjs/jest/tree/HEAD/packages/babel-jest/issues/15860">#15860</a>)</li>
<li>See full diff in <a href="https://github.com/jestjs/jest/commits/v30.3.0/packages/babel-jest">compare view</a></li>
</ul>
</details>
<br />

Updates `babel-loader` from 10.0.0 to 10.1.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel-loader/releases">babel-loader's releases</a>.</em></p>
<blockquote>
<h2>v10.1.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Revert <a href="https://redirect.github.com/babel/babel-loader/issues/1055">#1055</a> (&quot;use <code>module.findPackageJSON</code> API&quot;) by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1066">babel/babel-loader#1066</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/babel/babel-loader/compare/v10.1.0...v10.1.1">https://github.com/babel/babel-loader/compare/v10.1.0...v10.1.1</a></p>
<h2>v10.1.0</h2>
<h2>What's Changed</h2>
<ul>
<li>refactor: use <code>module.findPackageJSON</code> API by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1055">babel/babel-loader#1055</a></li>
<li>Enable type checking and support Babel 8 by <a href="https://github.com/JLHwung"><code>@​JLHwung</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1056">babel/babel-loader#1056</a></li>
<li>Bump js-yaml from 4.1.0 to 4.1.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/babel/babel-loader/pull/1059">babel/babel-loader#1059</a></li>
<li>fix: mark webpack as optional peer dependency by <a href="https://github.com/chenjiahan"><code>@​chenjiahan</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1061">babel/babel-loader#1061</a></li>
<li>Bump webpack from 5.101.0 to 5.104.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/babel/babel-loader/pull/1062">babel/babel-loader#1062</a></li>
<li>Bump glob from 10.4.5 to 10.5.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/babel/babel-loader/pull/1060">babel/babel-loader#1060</a></li>
<li>Bump minimatch from 3.1.2 to 3.1.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/babel/babel-loader/pull/1063">babel/babel-loader#1063</a></li>
<li>Pin Node.js on CI by <a href="https://github.com/nicolo-ribaudo"><code>@​nicolo-ribaudo</code></a> in <a href="https://redirect.github.com/babel/babel-loader/pull/1064">babel/babel-loader#1064</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/chenjiahan"><code>@​chenjiahan</code></a> made their first contribution in <a href="https://redirect.github.com/babel/babel-loader/pull/1061">babel/babel-loader#1061</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/babel/babel-loader/compare/v10.0.0...v10.1.0">https://github.com/babel/babel-loader/compare/v10.0.0...v10.1.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel-loader/commit/da602105664458752dca3578856ee8d0d6ac80e6"><code>da60210</code></a> 10.1.1</li>
<li><a href="https://github.com/babel/babel-loader/commit/a0a2617e10b39f35b8d1e2893a87f4ee4fe7ebdc"><code>a0a2617</code></a> Revert <a href="https://redirect.github.com/babel/babel-loader/issues/1055">#1055</a> (&quot;use <code>module.findPackageJSON</code> API&quot;) (<a href="https://redirect.github.com/babel/babel-loader/issues/1066">#1066</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/de09ee1426c781987674690be02aca2f2ea61efc"><code>de09ee1</code></a> 10.1.0</li>
<li><a href="https://github.com/babel/babel-loader/commit/e34c360a7b86740fca62158eb1bef89b9fef4507"><code>e34c360</code></a> Pin Node.js on CI (<a href="https://redirect.github.com/babel/babel-loader/issues/1064">#1064</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/3c1e1805511592e7e9f9fe7f60de5439370c4740"><code>3c1e180</code></a> Bump minimatch from 3.1.2 to 3.1.5 (<a href="https://redirect.github.com/babel/babel-loader/issues/1063">#1063</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/e0d4add38856fceeca1a633bb49927e4334999a7"><code>e0d4add</code></a> Bump glob from 10.4.5 to 10.5.0 (<a href="https://redirect.github.com/babel/babel-loader/issues/1060">#1060</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/77e2a66869cf84ddb6444d9b7b9951beb44b68b2"><code>77e2a66</code></a> Bump webpack from 5.101.0 to 5.104.1 (<a href="https://redirect.github.com/babel/babel-loader/issues/1062">#1062</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/faa5dbb4134e4d0c2659ed9bc55cc2e53b82d7e6"><code>faa5dbb</code></a> fix: mark webpack as optional peer dependency (<a href="https://redirect.github.com/babel/babel-loader/issues/1061">#1061</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/146dad2461ca5ba12fd202b33f6aa39be8218295"><code>146dad2</code></a> Bump js-yaml from 4.1.0 to 4.1.1 (<a href="https://redirect.github.com/babel/babel-loader/issues/1059">#1059</a>)</li>
<li><a href="https://github.com/babel/babel-loader/commit/2479ed223262f9ce45f9f7a9b8363a8666d9b41f"><code>2479ed2</code></a> Enable type checking and support Babel 8 (<a href="https://redirect.github.com/babel/babel-loader/issues/1056">#1056</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/babel/babel-loader/compare/v10.0.0...v10.1.1">compare view</a></li>
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

### Comment by @codecov-commenter on March 10, 2026 at 07:03 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1539?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 73.67%. Comparing base ([`f4a7c82`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/f4a7c8284aae975cae5809e77ca8ff8c849f2e50?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`4fd6ef2`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/4fd6ef2cd4941a2995d4ef4e85d753c2a71dab20?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1539   +/-   ##
=======================================
  Coverage   73.67%   73.67%           
=======================================
  Files          97       97           
  Lines        2359     2359           
  Branches      666      671    +5     
=======================================
  Hits         1738     1738           
  Misses        551      551           
  Partials       70       70           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1539?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @TenSt on March 12, 2026 at 11:17 AM UTC

@dependabot rebase

### Comment by @TenSt on March 13, 2026 at 02:46 PM UTC

@dependabot recreate

---

## Reviews

### Review by @TenSt - Approved on March 12, 2026 at 11:17 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1539*
