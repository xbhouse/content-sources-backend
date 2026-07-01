---
type: pull_request
number: 1586
title: "chore(deps): bump the lint group with 2 updates"
state: merged
author: dependabot
created: 2026-04-21T18:45:05Z
updated: 2026-04-22T13:16:44Z
closed: 2026-04-22T13:16:35Z
merged: 2026-04-22T13:16:35Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-77d3d348a9
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1586
---

# Pull Request #1586: chore(deps): bump the lint group with 2 updates

**Author**: @dependabot
**Created**: April 21, 2026 at 06:45 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-77d3d348a9`

## Description

Bumps the lint group with 2 updates: [eslint-plugin-playwright](https://github.com/mskelton/eslint-plugin-playwright) and [stylelint](https://github.com/stylelint/stylelint).

Updates `eslint-plugin-playwright` from 2.10.1 to 2.10.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/mskelton/eslint-plugin-playwright/releases">eslint-plugin-playwright's releases</a>.</em></p>
<blockquote>
<h2>v2.10.2</h2>
<h2><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.10.1...v2.10.2">2.10.2</a> (2026-04-20)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>missing-playwright-await:</strong> Fix false positive when re-assigning awaited variable (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8cca0ac362d9ddbce899195f1433f8d853efc3d0">8cca0ac</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/456">#456</a></li>
<li><strong>no-duplicate-hooks:</strong> handle anonymous describe blocks in forEach loops (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8b4ec601a0f801dc2a8701d66f12e28102ffc934">8b4ec60</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/459">#459</a></li>
<li><strong>valid-test-tags:</strong> Support template literal strings (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/d98a05cb51150bee283109e041e8e458f6d7bc5f">d98a05c</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/460">#460</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/613db7a7f646a307ad966136f6234bf2098500b2"><code>613db7a</code></a> chore: Fix type errors</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8cca0ac362d9ddbce899195f1433f8d853efc3d0"><code>8cca0ac</code></a> fix(missing-playwright-await): Fix false positive when re-assigning awaited v...</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8b4ec601a0f801dc2a8701d66f12e28102ffc934"><code>8b4ec60</code></a> fix(no-duplicate-hooks): handle anonymous describe blocks in forEach loops</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/d98a05cb51150bee283109e041e8e458f6d7bc5f"><code>d98a05c</code></a> fix(valid-test-tags): Support template literal strings</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/1158eda71fb67f4668f606820688d401052d35da"><code>1158eda</code></a> chore(deps): Bump flatted from 3.3.3 to 3.4.2 (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/452">#452</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/6e66967a6e25834233baa4aa74a54618a25cc820"><code>6e66967</code></a> chore(deps): Bump lodash-es from 4.17.23 to 4.18.1 (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/457">#457</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/ab4e713d477e6e7eaf5c8fc76ff2d4ae9038c9d6"><code>ab4e713</code></a> chore(deps): Bump vite from 7.3.1 to 7.3.2 (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/458">#458</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/47cc83a476f4eaf665ec18254af229f32c503ff8"><code>47cc83a</code></a> chore(deps): Bump handlebars from 4.7.8 to 4.7.9 (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/455">#455</a>)</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/b224504473dfed2bfee2024318b898a46b3c7089"><code>b224504</code></a> chore(deps): Bump picomatch from 2.3.1 to 2.3.2 (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/454">#454</a>)</li>
<li>See full diff in <a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.10.1...v2.10.2">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 17.7.0 to 17.8.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.8.0</h2>
<p>It adds 3 new rules and 1 configuration property.</p>
<ul>
<li>Added: <code>languageOptions.directionality</code> configuration property (<a href="https://redirect.github.com/stylelint/stylelint/issues/8687">#8687</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>property-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8687">#8687</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>relative-selector-nesting-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8730">#8730</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>selector-no-deprecated</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8694">#8694</a>) (<a href="https://github.com/immitsu"><code>@​immitsu</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.8.0 - 2026-04-15</h2>
<p>It adds 3 new rules and 1 configuration property.</p>
<ul>
<li>Added: <code>languageOptions.directionality</code> configuration property (<a href="https://redirect.github.com/stylelint/stylelint/pull/8687">#8687</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>property-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8687">#8687</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>relative-selector-nesting-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8730">#8730</a>) (<a href="https://github.com/sw1tch3roo"><code>@​sw1tch3roo</code></a>).</li>
<li>Added: <code>selector-no-deprecated</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8694">#8694</a>) (<a href="https://github.com/immitsu"><code>@​immitsu</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/b329c6fb0ee1c8fd530358ab6b7a11a96aa2c616"><code>b329c6f</code></a> Release 17.8.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9217">#9217</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/e9a4877f6af134af12f3f79d36c66dddd5a54ea4"><code>e9a4877</code></a> Bump postcss from 8.5.8 to 8.5.9 in the postcss group (<a href="https://redirect.github.com/stylelint/stylelint/issues/9216">#9216</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/0dc0d02659448ec2eff65b5b8e0464baa1350c2d"><code>0dc0d02</code></a> Bump eslint from 10.1.0 to 10.2.0 in the eslint group (<a href="https://redirect.github.com/stylelint/stylelint/issues/9215">#9215</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/4978eaf6f9b38905f845b5fc66fbe1f48cc3f29b"><code>4978eaf</code></a> Document <code>selector-no-deprecated</code> not yet in config (<a href="https://redirect.github.com/stylelint/stylelint/issues/9211">#9211</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/9ca941ed5e14d1b918587531ef0172533c829771"><code>9ca941e</code></a> Add <code>relative-selector-nesting-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8730">#8730</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/06807ce45638f6a1107e20f01f02ec0392783063"><code>06807ce</code></a> Add <code>selector-no-deprecated</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8694">#8694</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/16a6090344b085e816f47a30265f732d4f7d5178"><code>16a6090</code></a> Add <code>property-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8687">#8687</a>)</li>
<li>See full diff in <a href="https://github.com/stylelint/stylelint/compare/17.7.0...17.8.0">compare view</a></li>
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

### Comment by @codecov-commenter on April 21, 2026 at 06:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1586?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`b269490`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b269490a14623262fea77f090bef9096b88f8d54?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`95760c3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/95760c3a5d3f0d90e8668129ebd130c094782b8b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1586   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      685      686    +1     
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1586?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on April 22, 2026 at 01:16 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1586*
