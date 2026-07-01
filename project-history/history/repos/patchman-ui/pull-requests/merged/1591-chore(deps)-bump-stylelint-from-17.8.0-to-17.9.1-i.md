---
type: pull_request
number: 1591
title: "chore(deps): bump stylelint from 17.8.0 to 17.9.1 in the lint group"
state: merged
author: dependabot
created: 2026-04-28T18:44:59Z
updated: 2026-04-29T15:02:51Z
closed: 2026-04-29T15:02:39Z
merged: 2026-04-29T15:02:39Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-c226a97001
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1591
---

# Pull Request #1591: chore(deps): bump stylelint from 17.8.0 to 17.9.1 in the lint group

**Author**: @dependabot
**Created**: April 28, 2026 at 06:44 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-c226a97001`

## Description

Bumps the lint group with 1 update: [stylelint](https://github.com/stylelint/stylelint).

Updates `stylelint` from 17.8.0 to 17.9.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>17.9.1</h2>
<p>It fixes 4 bugs. We also documented the <code>messageArgs</code> each rule provides to the <code>message</code> configuration property.</p>
<ul>
<li>Fixed: <code>ConfigurationError</code> regression for custom syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/issues/9245">#9245</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: MD5 hash algorithm to SHA256 for caching (<a href="https://redirect.github.com/stylelint/stylelint/issues/9241">#9241</a>) (<a href="https://github.com/rkdfx"><code>@​rkdfx</code></a>).</li>
<li>Fixed: <code>property-no-deprecated</code> autofix for <code>page-break-*: always</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9214">#9214</a>) (<a href="https://github.com/rkdfx"><code>@​rkdfx</code></a>).</li>
<li>Fixed: <code>selector-no-deprecated</code> false positives for <code>::part()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9227">#9227</a>) (<a href="https://github.com/SaekiTominaga"><code>@​SaekiTominaga</code></a>).</li>
</ul>
<h2>17.9.0</h2>
<p>It adds 3 new features. Adding the <code>referenceFiles</code> property to your configuration object makes the <code>no-unknown-animations</code>, <code>no-unknown-custom-media</code> and <code>no-unknown-custom-properties</code> rules more useful.</p>
<ul>
<li>Added: experimental <code>referenceFiles</code> to configuration object (<a href="https://redirect.github.com/stylelint/stylelint/issues/9179">#9179</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: experimental <code>abortSignal</code> option to Node.js API for cancellation support (<a href="https://redirect.github.com/stylelint/stylelint/issues/9213">#9213</a>) (<a href="https://github.com/adalinesimonian"><code>@​adalinesimonian</code></a>).</li>
<li>Added: <code>maxWarnings</code> to configuration object (<a href="https://redirect.github.com/stylelint/stylelint/issues/9181">#9181</a>) (<a href="https://github.com/mrginglymus"><code>@​mrginglymus</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>17.9.1 - 2026-04-27</h2>
<p>It fixes 4 bugs. We also documented the <code>messageArgs</code> each rule provides to the <code>message</code> configuration property.</p>
<ul>
<li>Fixed: <code>ConfigurationError</code> regression for custom syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/pull/9245">#9245</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: MD5 hash algorithm to SHA256 for caching (<a href="https://redirect.github.com/stylelint/stylelint/pull/9241">#9241</a>) (<a href="https://github.com/rkdfx"><code>@​rkdfx</code></a>).</li>
<li>Fixed: <code>property-no-deprecated</code> autofix for <code>page-break-*: always</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9214">#9214</a>) (<a href="https://github.com/rkdfx"><code>@​rkdfx</code></a>).</li>
<li>Fixed: <code>selector-no-deprecated</code> false positives for <code>::part()</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/9227">#9227</a>) (<a href="https://github.com/SaekiTominaga"><code>@​SaekiTominaga</code></a>).</li>
</ul>
<h2>17.9.0 - 2026-04-23</h2>
<p>It adds 3 new features. Adding the <code>referenceFiles</code> property to your configuration object makes the <code>no-unknown-animations</code>, <code>no-unknown-custom-media</code> and <code>no-unknown-custom-properties</code> rules more useful.</p>
<ul>
<li>Added: experimental <code>referenceFiles</code> to configuration object (<a href="https://redirect.github.com/stylelint/stylelint/pull/9179">#9179</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: experimental <code>abortSignal</code> option to Node.js API for cancellation support (<a href="https://redirect.github.com/stylelint/stylelint/pull/9213">#9213</a>) (<a href="https://github.com/adalinesimonian"><code>@​adalinesimonian</code></a>).</li>
<li>Added: <code>maxWarnings</code> to configuration object (<a href="https://redirect.github.com/stylelint/stylelint/pull/9181">#9181</a>) (<a href="https://github.com/mrginglymus"><code>@​mrginglymus</code></a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/53c881784669bc068d492f49bb96761f4015e9d1"><code>53c8817</code></a> Release 17.9.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9248">#9248</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/fa750547b3154ff233387bf47050215979e7ac45"><code>fa75054</code></a> Fix <code>property-no-deprecated</code> autofix for <code>page-break-*: always</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9214">#9214</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/02b039e9a1150f8abb4749a14a25b869ba34f7cd"><code>02b039e</code></a> Fix <code>selector-no-deprecated</code> false positives for <code>::part()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/9227">#9227</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/af22bece4dfa23b58244d8dce02c772a52179789"><code>af22bec</code></a> Document message arguments of rules (<a href="https://redirect.github.com/stylelint/stylelint/issues/9226">#9226</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/20f6e9de1035cf830ca07e3fa6f60c62012bf11b"><code>20f6e9d</code></a> Fix <code>ConfigurationError</code> regression for custom syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/issues/9245">#9245</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/21a57e89a4a52f8a7ecb8b099940f88c738ce18d"><code>21a57e8</code></a> Fix MD5 hash algorithm to SHA256 for caching (<a href="https://redirect.github.com/stylelint/stylelint/issues/9241">#9241</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/cee404b4519dfe5e82151323733adf86a08ddd87"><code>cee404b</code></a> Release 17.9.0 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9242">#9242</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/b0af5aed125a88c9442e1821d763b97c2431ec80"><code>b0af5ae</code></a> Bump prettier from 3.8.1 to 3.8.3 (<a href="https://redirect.github.com/stylelint/stylelint/issues/9240">#9240</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/e2c2c43d87ee2515b9e5830f25d547fcfe526d45"><code>e2c2c43</code></a> Bump eslint-plugin-jest from 29.15.1 to 29.15.2 in the eslint group (<a href="https://redirect.github.com/stylelint/stylelint/issues/9239">#9239</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/68d008ec824b3597b2b4c28b85cdbaf30150b38e"><code>68d008e</code></a> Bump <code>@​csstools/css-syntax-patches-for-csstree</code> from 1.1.2 to 1.1.3 in the csst...</li>
<li>Additional commits viewable in <a href="https://github.com/stylelint/stylelint/compare/17.8.0...17.9.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=stylelint&package-manager=npm_and_yarn&previous-version=17.8.0&new-version=17.9.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @codecov-commenter on April 28, 2026 at 06:48 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1591?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`46c9cfe`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/46c9cfe36b6f5cde2136fa986a7f7a092abe4688?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`dec7dd1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/dec7dd10aec2f408c71cc5711dca2ecc595bc287?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1591   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1591/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1591/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `75.88% <ø> (?)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1591/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `75.88% <ø> (?)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1591?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on April 29, 2026 at 01:20 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1591*
