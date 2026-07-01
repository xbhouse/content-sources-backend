---
type: pull_request
number: 1502
title: "chore(deps): bump qs and express"
state: merged
author: dependabot
created: 2026-02-14T03:53:17Z
updated: 2026-02-16T08:38:34Z
closed: 2026-02-16T08:38:25Z
merged: 2026-02-16T08:38:25Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-a07fd7252a
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1502
---

# Pull Request #1502: chore(deps): bump qs and express

**Author**: @dependabot
**Created**: February 14, 2026 at 03:53 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-a07fd7252a`

## Description

Bumps [qs](https://github.com/ljharb/qs) and [express](https://github.com/expressjs/express). These dependencies needed to be updated together.
Updates `qs` from 6.14.0 to 6.14.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/ljharb/qs/blob/main/CHANGELOG.md">qs's changelog</a>.</em></p>
<blockquote>
<h2><strong>6.14.2</strong></h2>
<ul>
<li>[Fix] <code>parse</code>: mark overflow objects for indexed notation exceeding <code>arrayLimit</code> (<a href="https://redirect.github.com/ljharb/qs/issues/546">#546</a>)</li>
<li>[Fix] <code>arrayLimit</code> means max count, not max index, in <code>combine</code>/<code>merge</code>/<code>parseArrayValue</code></li>
<li>[Fix] <code>parse</code>: throw on <code>arrayLimit</code> exceeded with indexed notation when <code>throwOnLimitExceeded</code> is true (<a href="https://redirect.github.com/ljharb/qs/issues/529">#529</a>)</li>
<li>[Fix] <code>parse</code>: enforce <code>arrayLimit</code> on <code>comma</code>-parsed values</li>
<li>[Fix] <code>parse</code>: fix error message to reflect arrayLimit as max index; remove extraneous comments (<a href="https://redirect.github.com/ljharb/qs/issues/545">#545</a>)</li>
<li>[Robustness] avoid <code>.push</code>, use <code>void</code></li>
<li>[readme] document that <code>addQueryPrefix</code> does not add <code>?</code> to empty output (<a href="https://redirect.github.com/ljharb/qs/issues/418">#418</a>)</li>
<li>[readme] clarify <code>parseArrays</code> and <code>arrayLimit</code> documentation (<a href="https://redirect.github.com/ljharb/qs/issues/543">#543</a>)</li>
<li>[readme] replace runkit CI badge with shields.io check-runs badge</li>
<li>[meta] fix changelog typo (<code>arrayLength</code> → <code>arrayLimit</code>)</li>
<li>[actions] fix rebase workflow permissions</li>
</ul>
<h2><strong>6.14.1</strong></h2>
<ul>
<li>[Fix] ensure <code>arrayLimit</code> applies to <code>[]</code> notation as well</li>
<li>[Fix] <code>parse</code>: when a custom decoder returns <code>null</code> for a key, ignore that key</li>
<li>[Refactor] <code>parse</code>: extract key segment splitting helper</li>
<li>[meta] add threat model</li>
<li>[actions] add workflow permissions</li>
<li>[Tests] <code>stringify</code>: increase coverage</li>
<li>[Dev Deps] update <code>eslint</code>, <code>@ljharb/eslint-config</code>, <code>npmignore</code>, <code>es-value-fixtures</code>, <code>for-each</code>, <code>object-inspect</code></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ljharb/qs/commit/bdcf0c7f82387c18ac8fabfccd2f440645cef47b"><code>bdcf0c7</code></a> v6.14.2</li>
<li><a href="https://github.com/ljharb/qs/commit/294db90c812ddbe7d7a35d5687c505fd21a2d6a2"><code>294db90</code></a> [readme] document that <code>addQueryPrefix</code> does not add <code>?</code> to empty output</li>
<li><a href="https://github.com/ljharb/qs/commit/5c308e5516c270a78caa6f278465914090f91ec6"><code>5c308e5</code></a> [readme] clarify <code>parseArrays</code> and <code>arrayLimit</code> documentation</li>
<li><a href="https://github.com/ljharb/qs/commit/6addf8cf738d529c54d91f6f3ffb6c1be91bbfdc"><code>6addf8c</code></a> [Fix] <code>parse</code>: mark overflow objects for indexed notation exceeding <code>arrayLimit</code></li>
<li><a href="https://github.com/ljharb/qs/commit/cfc108f662326d6ab540f3545ef0b832baf83cdf"><code>cfc108f</code></a> [Fix] <code>arrayLimit</code> means max count, not max index, in <code>combine</code>/<code>merge</code>/`pars...</li>
<li><a href="https://github.com/ljharb/qs/commit/febb64442a80e49200211fa38d3c96b58024ac77"><code>febb644</code></a> [Fix] <code>parse</code>: throw on <code>arrayLimit</code> exceeded with indexed notation when `thr...</li>
<li><a href="https://github.com/ljharb/qs/commit/f6a7abff1f13d644db9b05fe4f2c98ada6bf8482"><code>f6a7abf</code></a> [Fix] <code>parse</code>: enforce <code>arrayLimit</code> on <code>comma</code>-parsed values</li>
<li><a href="https://github.com/ljharb/qs/commit/fbc5206c25b4d1851cea683f02c10756c521d15a"><code>fbc5206</code></a> [Fix] <code>parse</code>: fix error message to reflect arrayLimit as max index; remove e...</li>
<li><a href="https://github.com/ljharb/qs/commit/1b9a8b4e78c6aff4c22fa559107227f02fd0216a"><code>1b9a8b4</code></a> [actions] fix rebase workflow permissions</li>
<li><a href="https://github.com/ljharb/qs/commit/2a35775614e0fb46ac8a3060201a32a7c23a7fda"><code>2a35775</code></a> [meta] fix changelog typo (<code>arrayLength</code> → <code>arrayLimit</code>)</li>
<li>Additional commits viewable in <a href="https://github.com/ljharb/qs/compare/v6.14.0...v6.14.2">compare view</a></li>
</ul>
</details>
<br />

Updates `express` from 4.21.2 to 4.22.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/expressjs/express/releases">express's releases</a>.</em></p>
<blockquote>
<h2>v4.22.1</h2>
<h2>What's Changed</h2>
<blockquote>
<p>[!IMPORTANT]<br />
The prior release (4.22.0) included an erroneous breaking change related to the extended query parser. There is no actual security vulnerability associated with this behavior (CVE-2024-51999 has been rejected). The change has been fully reverted in this release.</p>
</blockquote>
<ul>
<li>Release: 4.22.1 by <a href="https://github.com/UlisesGascon"><code>@​UlisesGascon</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6934">expressjs/express#6934</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/expressjs/express/compare/4.22.0...v4.22.1">https://github.com/expressjs/express/compare/4.22.0...v4.22.1</a></p>
<h2>4.22.0</h2>
<h2>Important: Security</h2>
<ul>
<li>Security fix for <a href="https://www.cve.org/CVERecord?id=CVE-2024-51999">CVE-2024-51999</a> (<a href="https://github.com/expressjs/express/security/advisories/GHSA-pj86-cfqh-vqx6">GHSA-pj86-cfqh-vqx6</a>)</li>
</ul>
<h2>What's Changed</h2>
<ul>
<li>Refactor: improve readability by <a href="https://github.com/sazk07"><code>@​sazk07</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6190">expressjs/express#6190</a></li>
<li>ci: add support for Node.js@23.0 by <a href="https://github.com/UlisesGascon"><code>@​UlisesGascon</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6080">expressjs/express#6080</a></li>
<li>Method functions with no path should error by <a href="https://github.com/wesleytodd"><code>@​wesleytodd</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/5957">expressjs/express#5957</a></li>
<li>ci: updated github actions ci workflow by <a href="https://github.com/Phillip9587"><code>@​Phillip9587</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6323">expressjs/express#6323</a></li>
<li>ci: reorder <code>npm i</code> steps to fix ci for older node versions by <a href="https://github.com/Phillip9587"><code>@​Phillip9587</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6336">expressjs/express#6336</a></li>
<li>Backport: ci: add node.js 24 to test matrix by <a href="https://github.com/Phillip9587"><code>@​Phillip9587</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6506">expressjs/express#6506</a></li>
<li>chore(4.x): wider range for query test skip by <a href="https://github.com/jonchurch"><code>@​jonchurch</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6513">expressjs/express#6513</a></li>
<li>use tilde notation for certain dependencies by <a href="https://github.com/UlisesGascon"><code>@​UlisesGascon</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6905">expressjs/express#6905</a></li>
<li>deps: qs@6.14.0 by <a href="https://github.com/UlisesGascon"><code>@​UlisesGascon</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6909">expressjs/express#6909</a></li>
<li>deps: use tilde notation for <code>qs</code> by <a href="https://github.com/Phillip9587"><code>@​Phillip9587</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6919">expressjs/express#6919</a></li>
<li>Release: 4.22.0 by <a href="https://github.com/UlisesGascon"><code>@​UlisesGascon</code></a> in <a href="https://redirect.github.com/expressjs/express/pull/6921">expressjs/express#6921</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/expressjs/express/compare/4.21.2...4.22.0">https://github.com/expressjs/express/compare/4.21.2...4.22.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/expressjs/express/blob/v4.22.1/History.md">express's changelog</a>.</em></p>
<blockquote>
<h1>4.22.1 / 2025-12-01</h1>
<ul>
<li>Revert security fix for <a href="https://www.cve.org/CVERecord?id=CVE-2024-51999">CVE-2024-51999</a> (<a href="https://github.com/expressjs/express/security/advisories/GHSA-pj86-cfqh-vqx6">GHSA-pj86-cfqh-vqx6</a>)</li>
</ul>
<h1>4.22.0 / 2025-12-01</h1>
<ul>
<li>Security fix for <a href="https://www.cve.org/CVERecord?id=CVE-2024-51999">CVE-2024-51999</a> (<a href="https://github.com/expressjs/express/security/advisories/GHSA-pj86-cfqh-vqx6">GHSA-pj86-cfqh-vqx6</a>)</li>
<li>deps: use tilde notation for dependencies</li>
<li>deps: qs@6.14.0</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/expressjs/express/commit/12fae14531a78f19a2caaa5d4f58d9b01eaf3194"><code>12fae14</code></a> 4.22.1</li>
<li><a href="https://github.com/expressjs/express/commit/5ddf311af32e772a77fd48b6266ce2f1ba330e1a"><code>5ddf311</code></a> Revert &quot;sec: security patch for CVE-2024-51999&quot;</li>
<li><a href="https://github.com/expressjs/express/commit/49744abd1120484fe64d7bde1cd3197c32523b6e"><code>49744ab</code></a> 4.22.0 (<a href="https://redirect.github.com/expressjs/express/issues/6921">#6921</a>)</li>
<li><a href="https://github.com/expressjs/express/commit/6e97452f600a3b01719fbc5517d833c7646b0bb7"><code>6e97452</code></a> sec: security patch for CVE-2024-51999</li>
<li><a href="https://github.com/expressjs/express/commit/6a23d34d652b9e69a4486d2a2a0dea54b9685fa5"><code>6a23d34</code></a> deps: use tilde notation for <code>qs</code> (<a href="https://redirect.github.com/expressjs/express/issues/6919">#6919</a>)</li>
<li><a href="https://github.com/expressjs/express/commit/8c12cdf93b89a4628b59179e3cc0722fc517d6b3"><code>8c12cdf</code></a> deps: qs@6.14.0 (<a href="https://redirect.github.com/expressjs/express/issues/6909">#6909</a>)</li>
<li><a href="https://github.com/expressjs/express/commit/7fea74fcf02764580f38f2a7f1932dfa54cddd90"><code>7fea74f</code></a> deps: use tilde notation for certain dependencies (<a href="https://redirect.github.com/expressjs/express/issues/6905">#6905</a>)</li>
<li><a href="https://github.com/expressjs/express/commit/dac7a0475a99e9dfc57b3b8e6d5bdf52813f1944"><code>dac7a04</code></a> chore: wider range for query test skip (<a href="https://redirect.github.com/expressjs/express/issues/6513">#6513</a>)</li>
<li><a href="https://github.com/expressjs/express/commit/997919b48879bbd53171c3b4e5dd1b04ad139241"><code>997919b</code></a> ci: add node.js 24 to test matrix (<a href="https://redirect.github.com/expressjs/express/issues/6506">#6506</a>)</li>
<li><a href="https://github.com/expressjs/express/commit/36fb59c6c7d9dfca0b08dfeafb5b6e4a249234a1"><code>36fb59c</code></a> fix(ci): reorder <code>npm i</code> steps to fix ci for older node versions (<a href="https://redirect.github.com/expressjs/express/issues/6336">#6336</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/expressjs/express/compare/4.21.2...v4.22.1">compare view</a></li>
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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @codecov-commenter on February 16, 2026 at 08:05 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1502?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`cedc0ca`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/cedc0ca36801773c2b26c53c05cb886027336837?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b4f6d70`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b4f6d7016a4beefbeba6fcbb8b472d084e545aaf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1502   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      892           
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1502?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on February 16, 2026 at 08:38 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1502*
