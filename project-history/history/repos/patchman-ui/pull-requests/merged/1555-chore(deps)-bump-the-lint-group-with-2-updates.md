---
type: pull_request
number: 1555
title: "chore(deps): bump the lint group with 2 updates"
state: merged
author: dependabot
created: 2026-03-24T18:45:02Z
updated: 2026-03-26T13:12:40Z
closed: 2026-03-26T13:12:31Z
merged: 2026-03-26T13:12:31Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-a1bc1c395e
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1555
---

# Pull Request #1555: chore(deps): bump the lint group with 2 updates

**Author**: @dependabot
**Created**: March 24, 2026 at 06:45 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-a1bc1c395e`

## Description

Bumps the lint group with 2 updates: [eslint-plugin-playwright](https://github.com/mskelton/eslint-plugin-playwright) and [stylelint](https://github.com/stylelint/stylelint).

Updates `eslint-plugin-playwright` from 2.10.0 to 2.10.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/mskelton/eslint-plugin-playwright/releases">eslint-plugin-playwright's releases</a>.</em></p>
<blockquote>
<h2>v2.10.1</h2>
<h2><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.10.0...v2.10.1">2.10.1</a> (2026-03-18)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>missing-playwright-await:</strong> Don't flag Array.fill as missing await (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/cff96403204e3cac83faf2d1768e3ded1378302d">cff9640</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/450">#450</a></li>
<li>Narrow page detection to prefer false positives (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/10238e173e42725a369db5ee7fb162b1ee99d790">10238e1</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/cff96403204e3cac83faf2d1768e3ded1378302d"><code>cff9640</code></a> fix(missing-playwright-await): Don't flag Array.fill as missing await</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/10238e173e42725a369db5ee7fb162b1ee99d790"><code>10238e1</code></a> fix: Narrow page detection to prefer false positives</li>
<li>See full diff in <a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.10.0...v2.10.1">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 17.4.0 to 17.5.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.5.0</h2>
<p>It deprecates two rule options, adds 1 rule option and fixes 7 bugs. We've also released <a href="https://github.com/stylelint/vscode-stylelint/releases/tag/2.1.0"><code>2.1.0</code></a> of <a href="https://marketplace.visualstudio.com/items?itemName=stylelint.vscode-stylelint">our VS Code extension</a>, which adds 8 new requested features, and our <a href="https://www.npmjs.com/package/@stylelint/language-server">first release</a> of the Stylelint Language Server.</p>
<ul>
<li>Deprecated: <code>*syntax</code> options from <code>declaration-property-value-no-unknown</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9102">#9102</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Added: <code>ignoreMediaFeatureNameValues: {}</code> to <code>media-feature-name-value-no-unknown</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8976">#8976</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>node_modules</code> not ignored when using <code>codeFilename</code> in Node.js API (<a href="https://redirect.github.com/stylelint/stylelint/issues/9130">#9130</a>) (<a href="https://github.com/adalinesimonian"><code>@​adalinesimonian</code></a>).</li>
<li>Fixed: <code>Error TS7016</code> for imported <code>css-tree</code> types (<a href="https://redirect.github.com/stylelint/stylelint/issues/9133">#9133</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Fixed: <code>declaration-property-value-keyword-no-deprecated</code> false positives for function arguments (<a href="https://redirect.github.com/stylelint/stylelint/issues/9116">#9116</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>calc-size()</code> containing <code>size</code> keyword (<a href="https://redirect.github.com/stylelint/stylelint/issues/9105">#9105</a>) (<a href="https://github.com/hriztam"><code>@​hriztam</code></a>).</li>
<li>Fixed: <code>no-descending-specificity</code> &amp; <code>no-duplicate-selectors</code> false negatives for equivalent compound selectors (<a href="https://redirect.github.com/stylelint/stylelint/issues/8977">#8977</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>no-invalid-position-declaration</code> false positives for <code>@mixin</code> and <code>@scope</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9120">#9120</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false negatives for <code>types</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9117">#9117</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.5.0 - 2026-03-19</h2>
<p>It deprecates two rule options, adds 1 rule option and fixes 7 bugs. We've also released <a href="https://github.com/stylelint/vscode-stylelint/releases/tag/2.1.0"><code>2.1.0</code></a> of <a href="https://marketplace.visualstudio.com/items?itemName=stylelint.vscode-stylelint">our VS Code extension</a>, which adds 8 new requested features, and our <a href="https://www.npmjs.com/package/@stylelint/language-server">first release</a> of the Stylelint Language Server.</p>
<ul>
<li>Deprecated: <code>*syntax</code> options from <code>declaration-property-value-no-unknown</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9102">#9102</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Added: <code>ignoreMediaFeatureNameValues: {}</code> to <code>media-feature-name-value-no-unknown</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8976">#8976</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>node_modules</code> not ignored when using <code>codeFilename</code> in Node.js API (<a href="https://redirect.github.com/stylelint/stylelint/pull/9130">#9130</a>) (<a href="https://github.com/adalinesimonian"><code>@​adalinesimonian</code></a>).</li>
<li>Fixed: <code>Error TS7016</code> for imported <code>css-tree</code> types (<a href="https://redirect.github.com/stylelint/stylelint/pull/9133">#9133</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Fixed: <code>declaration-property-value-keyword-no-deprecated</code> false positives for function arguments (<a href="https://redirect.github.com/stylelint/stylelint/pull/9116">#9116</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>calc-size()</code> containing <code>size</code> keyword (<a href="https://redirect.github.com/stylelint/stylelint/pull/9105">#9105</a>) (<a href="https://github.com/hriztam"><code>@​hriztam</code></a>).</li>
<li>Fixed: <code>no-descending-specificity</code> &amp; <code>no-duplicate-selectors</code> false negatives for equivalent compound selectors (<a href="https://redirect.github.com/stylelint/stylelint/pull/8977">#8977</a>) (<a href="https://github.com/kovsu"><code>@​kovsu</code></a>).</li>
<li>Fixed: <code>no-invalid-position-declaration</code> false positives for <code>@mixin</code> and <code>@scope</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9120">#9120</a>) (<a href="https://github.com/ragini-pandey"><code>@​ragini-pandey</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false negatives for <code>types</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9117">#9117</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/8d0198af1bfb6a144d8fce9a094754e41c148237"><code>8d0198a</code></a> Release 17.5.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9160">#9160</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/3ba50ac3e6e915774a587ff119c0b609b506705e"><code>3ba50ac</code></a> Add <code>ignoreMediaFeatureNameValues: {}</code> to `media-feature-name-value-no-unknow...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/77dfd01d01d38711e35e8dead4440df179698808"><code>77dfd01</code></a> Fix <code>no-invalid-position-declaration</code> false positives for <code>@mixin</code> and `@scop...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/7cde2d0ccf5547a28a8f7a6be56e505441a6f437"><code>7cde2d0</code></a> Fix <code>declaration-property-value-keyword-no-deprecated</code> false positives for fu...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/7b10658abb77a2ab309adca2f077062d80467f3a"><code>7b10658</code></a> Fix <code>declaration-property-value-no-unknown</code> false positives for <code>calc-size()</code>...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/e5b9ec7ad4437eae9dd9b2a7677b785dd25c51a3"><code>e5b9ec7</code></a> Deprecate <code>*syntax</code> options from <code>declaration-property-value-no-unknown</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9102">#9102</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/5bd2d21e8a6b47f529314284d162e6dcb37ef681"><code>5bd2d21</code></a> Bump css-tree from 3.1.0 to 3.2.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9158">#9158</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/995caf4ea7c7fd0e352cdec026a3edd9af6b665f"><code>995caf4</code></a> Bump eslint from 10.0.2 to 10.0.3 in the eslint group (<a href="https://redirect.github.com/stylelint/stylelint/issues/9156">#9156</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/f5f6dddc6b1b4e11dfc701dc40ee4b496252e848"><code>f5f6ddd</code></a> Bump lint-staged from 16.2.7 to 16.3.3 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9159">#9159</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/68948180cd615f91baf2bc19d5999be814102f2b"><code>6894818</code></a> Bump the jest group with 2 updates (<a href="https://redirect.github.com/stylelint/stylelint/issues/9157">#9157</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/17.4.0...17.5.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Install script changes</summary>
<p>This version modifies <code>prepare</code> script that runs during installation. Review the package contents before updating.</p>
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

### Comment by @codecov-commenter on March 24, 2026 at 06:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1555?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 72.34%. Comparing base ([`53f6193`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/53f619329d3b2cadabdf24e792274d8e48bbe4fa?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`1b3404e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1b3404e5fbea64825839b9a4e8ec6f73a41a05f6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 5 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1555   +/-   ##
=======================================
  Coverage   72.34%   72.34%           
=======================================
  Files          99       99           
  Lines        2408     2408           
  Branches      683      678    -5     
=======================================
  Hits         1742     1742           
  Misses        586      586           
  Partials       80       80           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1555?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on March 26, 2026 at 01:12 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1555*
