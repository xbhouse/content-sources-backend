---
type: pull_request
number: 1563
title: "chore(deps): bump serialize-javascript and terser-webpack-plugin"
state: merged
author: dependabot
created: 2026-03-29T04:28:29Z
updated: 2026-04-02T13:38:02Z
closed: 2026-04-02T13:37:53Z
merged: 2026-04-02T13:37:53Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-0d13b2d87f
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1563
---

# Pull Request #1563: chore(deps): bump serialize-javascript and terser-webpack-plugin

**Author**: @dependabot
**Created**: March 29, 2026 at 04:28 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-0d13b2d87f`

## Description

Removes [serialize-javascript](https://github.com/yahoo/serialize-javascript). It's no longer used after updating ancestor dependency [terser-webpack-plugin](https://github.com/webpack/terser-webpack-plugin). These dependencies need to be updated together.

Removes `serialize-javascript`

Updates `terser-webpack-plugin` from 5.3.16 to 5.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/webpack/terser-webpack-plugin/releases">terser-webpack-plugin's releases</a>.</em></p>
<blockquote>
<h2>v5.4.0</h2>
<h2><a href="https://github.com/webpack/terser-webpack-plugin/compare/v5.3.17...v5.4.0">5.4.0</a> (2026-03-10)</h2>
<h3>Features</h3>
<ul>
<li>added ability to minimize <code>JSON</code> using <code>jsonMinify</code> (<a href="https://redirect.github.com/webpack/terser-webpack-plugin/issues/657">#657</a>) (<a href="https://github.com/webpack/terser-webpack-plugin/commit/29ac915e58ec2ff81346d936ee2cf63ac6e300b8">29ac915</a>)</li>
</ul>
<h2>v5.3.17</h2>
<h3><a href="https://github.com/webpack/terser-webpack-plugin/compare/v5.3.16...v5.3.17">5.3.17</a> (2026-03-03)</h3>
<h3>Bug Fixes</h3>
<ul>
<li>update <code>serialize-javascript</code> (<a href="https://github.com/webpack/terser-webpack-plugin/commit/37c490c326ffe3416b50028a91bfba5661e1344e">37c490c</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/webpack/terser-webpack-plugin/blob/main/CHANGELOG.md">terser-webpack-plugin's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/webpack/terser-webpack-plugin/compare/v5.3.17...v5.4.0">5.4.0</a> (2026-03-10)</h2>
<h3>Features</h3>
<ul>
<li>added ability to minimizer <code>JSON</code> using <code>jsonMinify</code> (<a href="https://redirect.github.com/webpack/terser-webpack-plugin/issues/657">#657</a>) (<a href="https://github.com/webpack/terser-webpack-plugin/commit/29ac915e58ec2ff81346d936ee2cf63ac6e300b8">29ac915</a>)</li>
</ul>
<h3><a href="https://github.com/webpack/terser-webpack-plugin/compare/v5.3.16...v5.3.17">5.3.17</a> (2026-03-03)</h3>
<h3>Bug Fixes</h3>
<ul>
<li>update <code>serialize-javascript</code> (<a href="https://github.com/webpack/terser-webpack-plugin/commit/37c490c326ffe3416b50028a91bfba5661e1344e">37c490c</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/543da6e282d654d804c8d9d86acdd5fde23d80bf"><code>543da6e</code></a> chore(release): 5.4.0</li>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/29ac915e58ec2ff81346d936ee2cf63ac6e300b8"><code>29ac915</code></a> feat: added ability to minimizer <code>JSON</code> using <code>jsonMinify</code> (<a href="https://redirect.github.com/webpack/terser-webpack-plugin/issues/657">#657</a>)</li>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/e505deecb9230357b77532b88eda9392a682ba9e"><code>e505dee</code></a> fix: align with code</li>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/6f911ffba16ccd2bcfa99cd0bd7e956cab53d147"><code>6f911ff</code></a> chore(release): 5.3.17</li>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/37c490c326ffe3416b50028a91bfba5661e1344e"><code>37c490c</code></a> fix: update <code>serialize-javascript</code></li>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/207764f3ccbab62130b30a97d2df2780faf1bccd"><code>207764f</code></a> chore: deps update (<a href="https://redirect.github.com/webpack/terser-webpack-plugin/issues/652">#652</a>)</li>
<li><a href="https://github.com/webpack/terser-webpack-plugin/commit/a85ab47e575f8ce5e750dad714f2a98382de1774"><code>a85ab47</code></a> chore(deps-dev): bump ajv from 6.12.6 to 6.14.0 (<a href="https://redirect.github.com/webpack/terser-webpack-plugin/issues/648">#648</a>)</li>
<li>See full diff in <a href="https://github.com/webpack/terser-webpack-plugin/compare/v5.3.16...v5.4.0">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on March 29, 2026 at 04:30 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1563?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`a08d0d6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a08d0d69deb64ea8523582130fbabff1dde6be51?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`a55dc3f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a55dc3f25d768324d4e11f2d688e2cf14c4f68bb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1563   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      686      686           
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1563?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on April 02, 2026 at 01:32 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1563*
