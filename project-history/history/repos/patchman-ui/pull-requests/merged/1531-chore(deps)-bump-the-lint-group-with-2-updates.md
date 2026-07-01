---
type: pull_request
number: 1531
title: "chore(deps): bump the lint group with 2 updates"
state: merged
author: dependabot
created: 2026-03-03T18:43:33Z
updated: 2026-03-04T09:17:59Z
closed: 2026-03-04T09:17:49Z
merged: 2026-03-04T09:17:49Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-da4c6aa409
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1531
---

# Pull Request #1531: chore(deps): bump the lint group with 2 updates

**Author**: @dependabot
**Created**: March 03, 2026 at 06:43 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-da4c6aa409`

## Description

Bumps the lint group with 2 updates: [eslint-plugin-playwright](https://github.com/mskelton/eslint-plugin-playwright) and [stylelint](https://github.com/stylelint/stylelint).

Updates `eslint-plugin-playwright` from 2.7.1 to 2.9.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/mskelton/eslint-plugin-playwright/releases">eslint-plugin-playwright's releases</a>.</em></p>
<blockquote>
<h2>v2.9.0</h2>
<h1><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.8.0...v2.9.0">2.9.0</a> (2026-03-02)</h1>
<h3>Bug Fixes</h3>
<ul>
<li><strong>no-restricted-roles:</strong> Catch all uses, not just on page methods (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/1861fa57fd21b8d2d17cbd96238fdf5970277686">1861fa5</a>)</li>
<li>Support nested locators everywhere (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/0e48186885a8868d3163cdd2d2732c3f227056ab">0e48186</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>no-duplicate-hooks:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fe3ca54178cde017388aa5d844553a9b7a9d1307">fe3ca54</a>)</li>
<li><strong>no-duplicate-slow:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/2f0b67d841dd8f091f7c8ab44a6eadb34255127b">2f0b67d</a>)</li>
<li><strong>prefer-hooks-in-order:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/e8ae16ef219ce943749b194fc972b27a0162e8cb">e8ae16e</a>)</li>
<li><strong>prefer-hooks-on-top:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/5ab929677dea5263b69e2fe44a43cf711b1bbb16">5ab9296</a>)</li>
<li><strong>prefer-locator:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fcab221c092c290c8b2b851b669ea0b1774ec75c">fcab221</a>)</li>
<li><strong>prefer-to-have-count:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fcbf086ae637cf569530856e7690b6da85a5c5fe">fcbf086</a>)</li>
<li><strong>prefer-to-have-length:</strong> Mark as recommended (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/c6c923e6b382c1bee4de89e2cf9781511a37e7a0">c6c923e</a>)</li>
</ul>
<h2>v2.8.0</h2>
<h1><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.7.1...v2.8.0">2.8.0</a> (2026-02-27)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>Add missing test coverage and fix several minor bugs (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/434">#434</a>) (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/e3398ec61da52de205e7c9af2896633357769f74">e3398ec</a>)</li>
<li><strong>missing-playwright-await:</strong> Handle spread elements (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/df3016323819f7bc335fd1841971dccc2ae64f51">df30163</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/430">#430</a></li>
<li><strong>missing-playwright-await:</strong> Support more promise edge cases (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/b4cdcbd010a2b4dfc7ee14ab5bdc655897389f19">b4cdcbd</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>Auto-detect <code>test.extend()</code> fixtures and import aliases (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/432">#432</a>) (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8b22ee7b1f7823d81bafda82e240dd51106726dd">8b22ee7</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/7e16bd565cfccd365a6a8f1f7f6fe29a1c868036"><code>7e16bd5</code></a> docs: Fix readme for some rules</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/c27110b2ba1950fd8cc7bccafaa1fc817a3cdef7"><code>c27110b</code></a> Merge pull request <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/439">#439</a> from mskelton/new-recommended-rules</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/5ab929677dea5263b69e2fe44a43cf711b1bbb16"><code>5ab9296</code></a> feat(prefer-hooks-on-top): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/e8ae16ef219ce943749b194fc972b27a0162e8cb"><code>e8ae16e</code></a> feat(prefer-hooks-in-order): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fcab221c092c290c8b2b851b669ea0b1774ec75c"><code>fcab221</code></a> feat(prefer-locator): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/c6c923e6b382c1bee4de89e2cf9781511a37e7a0"><code>c6c923e</code></a> feat(prefer-to-have-length): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fcbf086ae637cf569530856e7690b6da85a5c5fe"><code>fcbf086</code></a> feat(prefer-to-have-count): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/fe3ca54178cde017388aa5d844553a9b7a9d1307"><code>fe3ca54</code></a> feat(no-duplicate-hooks): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/2f0b67d841dd8f091f7c8ab44a6eadb34255127b"><code>2f0b67d</code></a> feat(no-duplicate-slow): Mark as recommended</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/0e48186885a8868d3163cdd2d2732c3f227056ab"><code>0e48186</code></a> fix: Support nested locators everywhere</li>
<li>Additional commits viewable in <a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.7.1...v2.9.0">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 17.3.0 to 17.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.4.0</h2>
<p>It adds 2 options to the rules and fixes 7 bugs.</p>
<ul>
<li>Added: <code>ignoreAtRules: []</code> to <code>at-rule-no-vendor-prefix</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9096">#9096</a>) (<a href="https://github.com/theacrat"><code>@​theacrat</code></a>).</li>
<li>Added: <code>ignoreMediaFeatureNames: []</code> to <code>media-feature-name-no-vendor-prefix</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9097">#9097</a>) (<a href="https://github.com/theacrat"><code>@​theacrat</code></a>).</li>
<li>Fixed: performance of selector cloning rules (<a href="https://redirect.github.com/stylelint/stylelint/issues/9089">#9089</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>*-empty-line-before</code> performance (<a href="https://redirect.github.com/stylelint/stylelint/issues/9092">#9092</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> performance (<a href="https://redirect.github.com/stylelint/stylelint/issues/9090">#9090</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>no-irregular-whitespace</code> performance (<a href="https://redirect.github.com/stylelint/stylelint/issues/9091">#9091</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false negatives for at-rule descriptors (<a href="https://redirect.github.com/stylelint/stylelint/issues/9109">#9109</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false positives for <code>corner-shape</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9099">#9099</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false positives for double-slashed properties (<a href="https://redirect.github.com/stylelint/stylelint/issues/9099">#9099</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.4.0 - 2026-02-25</h2>
<p>It adds 2 options to the rules and fixes 7 bugs.</p>
<ul>
<li>Added: <code>ignoreAtRules: []</code> to <code>at-rule-no-vendor-prefix</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9096">#9096</a>) (<a href="https://github.com/theacrat"><code>@​theacrat</code></a>).</li>
<li>Added: <code>ignoreMediaFeatureNames: []</code> to <code>media-feature-name-no-vendor-prefix</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9097">#9097</a>) (<a href="https://github.com/theacrat"><code>@​theacrat</code></a>).</li>
<li>Fixed: performance of selector cloning rules (<a href="https://redirect.github.com/stylelint/stylelint/pull/9089">#9089</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>*-empty-line-before</code> performance (<a href="https://redirect.github.com/stylelint/stylelint/pull/9092">#9092</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> performance (<a href="https://redirect.github.com/stylelint/stylelint/pull/9090">#9090</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>no-irregular-whitespace</code> performance (<a href="https://redirect.github.com/stylelint/stylelint/pull/9091">#9091</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false negatives for at-rule descriptors (<a href="https://redirect.github.com/stylelint/stylelint/pull/9109">#9109</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false positives for <code>corner-shape</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9099">#9099</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false positives for double-slashed properties (<a href="https://redirect.github.com/stylelint/stylelint/pull/9099">#9099</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/556b7ad511a411eb9ac6784342e85a8ed6123071"><code>556b7ad</code></a> Release 17.4.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9113">#9113</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/05f64adb508c2f16b9fa057b897e1281b3fea7e2"><code>05f64ad</code></a> Run <code>npm audit</code> for <code>minimatch</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9112">#9112</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/d358479e2617fddacdf585c2293c40593f7007ef"><code>d358479</code></a> Document steps for adding an option and fixing a rule bug (<a href="https://redirect.github.com/stylelint/stylelint/issues/9111">#9111</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/467c5c9f52ca65edc6bbfe9b4ffac4429fd5f2db"><code>467c5c9</code></a> Fix <code>property-no-unknown</code> false negatives for at-rule descriptors (<a href="https://redirect.github.com/stylelint/stylelint/issues/9109">#9109</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/4250b21ed9b6ad88280157f15201df5c2016b7df"><code>4250b21</code></a> Document opening a pull request requests in CONTRIBUTING (<a href="https://redirect.github.com/stylelint/stylelint/issues/9110">#9110</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/effb787e250d9d0fc1107f4437e0cff430d5fada"><code>effb787</code></a> Remove <code>husky</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8984">#8984</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/f38a18c51fbba6d11c4e0577e74afd37f390d31a"><code>f38a18c</code></a> Remove redundant <code>@types/file-entry-cache</code> type dep (<a href="https://redirect.github.com/stylelint/stylelint/issues/9108">#9108</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/83ed11bfc2fc99d8f6f685bfbfb21044215bc2a9"><code>83ed11b</code></a> Add <code>ignoreAtRules: []</code> to <code>at-rule-no-vendor-prefix</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9096">#9096</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/0a8f468dcb1618fbbffa122e5753c031a6d870be"><code>0a8f468</code></a> Add <code>ignoreMediaFeatureNames: []</code> to <code>media-feature-name-no-vendor-prefix</code> (#...</li>
<li><a href="https://github.com/stylelint/stylelint/commit/83cfd657fb788a6994dd56958a919f3d96df769a"><code>83cfd65</code></a> Fix <code>property-no-unknown</code> false positives for <code>corner-shape</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9099">#9099</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/17.3.0...17.4.0">compare view</a></li>
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

### Comment by @codecov-commenter on March 03, 2026 at 07:04 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1531?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 73.67%. Comparing base ([`336412a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/336412abf0e3b35550a1599f977a099658320866?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b0cd3f9`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b0cd3f99df732f60ca764d50697ae74cec07b787?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1531   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1531?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on March 04, 2026 at 09:17 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1531*
