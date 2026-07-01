---
type: pull_request
number: 1602
title: "chore(deps): bump stylelint from 17.9.1 to 17.11.0 in the lint group across 1 directory"
state: merged
author: dependabot
created: 2026-05-05T18:43:45Z
updated: 2026-05-08T14:50:38Z
closed: 2026-05-08T14:50:29Z
merged: 2026-05-08T14:50:29Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-f3aa669150
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1602
---

# Pull Request #1602: chore(deps): bump stylelint from 17.9.1 to 17.11.0 in the lint group across 1 directory

**Author**: @dependabot
**Created**: May 05, 2026 at 06:43 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-f3aa669150`

## Description

Bumps the lint group with 1 update in the / directory: [stylelint](https://github.com/stylelint/stylelint).

Updates `stylelint` from 17.9.1 to 17.11.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.11.0</h2>
<p>It adds 2 features, including a <code>loader</code> property to <code>referenceFiles: {}</code> for when the order of appearance in the reference styles matters.</p>
<ul>
<li>Added: <code>loader</code> to experimental <code>referenceFiles: {}</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9251">#9251</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Added: <code>autofixed</code> to the result object (<a href="https://redirect.github.com/stylelint/stylelint/issues/8771">#8771</a>) (<a href="https://github.com/Rob"><code>@​Rob</code></a>--W).</li>
</ul>
<h2>17.10.0</h2>
<p>It adds 3 rules and fixes 4 bugs. You can use the <code>*-layout-mappings</code> rules to enforce logical or physical properties, units and keywords.</p>
<ul>
<li>Added: <code>selector-no-invalid</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/9232">#9232</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>unit-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/9229">#9229</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>value-keyword-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/9233">#9233</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: inconsistent error messages when module is not found (<a href="https://redirect.github.com/stylelint/stylelint/issues/9260">#9260</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
<li>Fixed: <code>property-layout-mappings</code> false negatives for property names in declaration values (<a href="https://redirect.github.com/stylelint/stylelint/issues/9222">#9222</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-layout-mappings</code> false positives for <code>@page</code> properties (<a href="https://redirect.github.com/stylelint/stylelint/issues/9223">#9223</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for nested <code>webkit-scrollbar</code> part (<a href="https://redirect.github.com/stylelint/stylelint/issues/9259">#9259</a>) (<a href="https://github.com/rkdfx"><code>@​rkdfx</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.11.0 - 2026-05-05</h2>
<p>It adds 2 features, including a <code>loader</code> property to <code>referenceFiles: {}</code> for when the order of appearance in the reference styles matters.</p>
<ul>
<li>Added: <code>loader</code> to experimental <code>referenceFiles: {}</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9251">#9251</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Added: <code>autofixed</code> to the result object (<a href="https://redirect.github.com/stylelint/stylelint/pull/8771">#8771</a>) (<a href="https://github.com/Rob--W"><code>@​Rob--W</code></a>).</li>
</ul>
<h2>17.10.0 - 2026-05-03</h2>
<p>It adds 3 rules and fixes 4 bugs. You can use the <code>*-layout-mappings</code> rules to enforce logical or physical properties, units and keywords.</p>
<ul>
<li>Added: <code>selector-no-invalid</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/9232">#9232</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>unit-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/9229">#9229</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>value-keyword-layout-mappings</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/9233">#9233</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: inconsistent error messages when module is not found (<a href="https://redirect.github.com/stylelint/stylelint/pull/9260">#9260</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
<li>Fixed: <code>property-layout-mappings</code> false negatives for property names in declaration values (<a href="https://redirect.github.com/stylelint/stylelint/pull/9222">#9222</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-layout-mappings</code> false positives for <code>@page</code> properties (<a href="https://redirect.github.com/stylelint/stylelint/pull/9223">#9223</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for nested <code>webkit-scrollbar</code> part (<a href="https://redirect.github.com/stylelint/stylelint/pull/9259">#9259</a>) (<a href="https://github.com/rkdfx"><code>@​rkdfx</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/e176acfb995aa7d8b6841cc514976187333c55e8"><code>e176acf</code></a> Release 17.11.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9271">#9271</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/769e791a0389c34983862609da94006c12ccd4a3"><code>769e791</code></a> Require owner review for <code>.npmrc</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9270">#9270</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/c0c362c375a9ef6b105fbb25ed01e86a214c733e"><code>c0c362c</code></a> Harden <code>npm install</code> security (<a href="https://redirect.github.com/stylelint/stylelint/issues/9269">#9269</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/823fdfca02c14a2fa3a0eb03afbd2f48cb25ac34"><code>823fdfc</code></a> Add <code>autofixed</code> to the result object (<a href="https://redirect.github.com/stylelint/stylelint/issues/8771">#8771</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/c19e627e47b9527d47d18c2c8c9b57d460ae2a54"><code>c19e627</code></a> Bump vulnerable dependencies via <code>npm audit fix</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9267">#9267</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/2ce046f96dd1ae70b7d7936cfc0b3dafd29dd784"><code>2ce046f</code></a> Bump string-width from 8.2.0 to 8.2.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9266">#9266</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/a895cddf1ec7d2dd281b6365d739f5492967af56"><code>a895cdd</code></a> Bump the stylelint-actions group with 4 updates (<a href="https://redirect.github.com/stylelint/stylelint/issues/9265">#9265</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/1757a823b3acd9090f77a75866e66c8e9700c77a"><code>1757a82</code></a> Configure Dependabot to group Stylelint actions (<a href="https://redirect.github.com/stylelint/stylelint/issues/9264">#9264</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/34918a424b6dd9be75b335b9e3e1fb19591d26da"><code>34918a4</code></a> Add Dependency Review to CI (<a href="https://redirect.github.com/stylelint/stylelint/issues/9263">#9263</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/70c1f2db5a73c4b00fecdaa1c5de2113f1d25667"><code>70c1f2d</code></a> Add <code>loader</code> to experimental <code>referenceFiles: {}</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9251">#9251</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/17.9.1...17.11.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Install script changes</summary>
<p>This version modifies <code>prepare</code> script that runs during installation. Review the package contents before updating.</p>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on May 05, 2026 at 06:47 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1602?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 76.34%. Comparing base ([`010c257`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/010c25737f959f80cd5dad0250f3e36046e69cd1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`713cdee`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/713cdee314332cc11673d0d75e9bda83642e0ee0?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1602   +/-   ##
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

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1602?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on May 08, 2026 at 09:31 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1602*
