---
type: pull_request
number: 1464
title: "chore(deps): bump js-yaml from 3.14.1 to 3.14.2"
state: closed
author: dependabot
created: 2026-01-09T20:58:17Z
updated: 2026-02-16T07:45:39Z
closed: 2026-02-16T07:45:37Z
base_branch: master
head_branch: dependabot/npm_and_yarn/js-yaml-3.14.2
labels: ["dependencies", "patch"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1464
---

# Pull Request #1464: chore(deps): bump js-yaml from 3.14.1 to 3.14.2

**Author**: @dependabot
**Created**: January 09, 2026 at 08:58 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `patch`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/js-yaml-3.14.2`

## Description

Bumps [js-yaml](https://github.com/nodeca/js-yaml) from 3.14.1 to 3.14.2.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/nodeca/js-yaml/blob/master/CHANGELOG.md">js-yaml's changelog</a>.</em></p>
<blockquote>
<h2>[3.14.2] - 2025-11-15</h2>
<h3>Security</h3>
<ul>
<li>Backported v4.1.1 fix to v3</li>
</ul>
<h2>[4.1.1] - 2025-11-12</h2>
<h3>Security</h3>
<ul>
<li>Fix prototype pollution issue in yaml merge (&lt;&lt;) operator.</li>
</ul>
<h2>[4.1.0] - 2021-04-15</h2>
<h3>Added</h3>
<ul>
<li>Types are now exported as <code>yaml.types.XXX</code>.</li>
<li>Every type now has <code>options</code> property with original arguments kept as they were
(see <code>yaml.types.int.options</code> as an example).</li>
</ul>
<h3>Changed</h3>
<ul>
<li><code>Schema.extend()</code> now keeps old type order in case of conflicts
(e.g. Schema.extend([ a, b, c ]).extend([ b, a, d ]) is now ordered as <code>abcd</code> instead of <code>cbad</code>).</li>
</ul>
<h2>[4.0.0] - 2021-01-03</h2>
<h3>Changed</h3>
<ul>
<li>Check <a href="https://github.com/nodeca/js-yaml/blob/master/migrate_v3_to_v4.md">migration guide</a> to see details for all breaking changes.</li>
<li>Breaking: &quot;unsafe&quot; tags <code>!!js/function</code>, <code>!!js/regexp</code>, <code>!!js/undefined</code> are
moved to <a href="https://github.com/nodeca/js-yaml-js-types">js-yaml-js-types</a> package.</li>
<li>Breaking: removed <code>safe*</code> functions. Use <code>load</code>, <code>loadAll</code>, <code>dump</code>
instead which are all now safe by default.</li>
<li><code>yaml.DEFAULT_SAFE_SCHEMA</code> and <code>yaml.DEFAULT_FULL_SCHEMA</code> are removed, use
<code>yaml.DEFAULT_SCHEMA</code> instead.</li>
<li><code>yaml.Schema.create(schema, tags)</code> is removed, use <code>schema.extend(tags)</code> instead.</li>
<li><code>!!binary</code> now always mapped to <code>Uint8Array</code> on load.</li>
<li>Reduced nesting of <code>/lib</code> folder.</li>
<li>Parse numbers according to YAML 1.2 instead of YAML 1.1 (<code>01234</code> is now decimal,
<code>0o1234</code> is octal, <code>1:23</code> is parsed as string instead of base60).</li>
<li><code>dump()</code> no longer quotes <code>:</code>, <code>[</code>, <code>]</code>, <code>(</code>, <code>)</code> except when necessary, <a href="https://redirect.github.com/nodeca/js-yaml/issues/470">#470</a>, <a href="https://redirect.github.com/nodeca/js-yaml/issues/557">#557</a>.</li>
<li>Line and column in exceptions are now formatted as <code>(X:Y)</code> instead of
<code>at line X, column Y</code> (also present in compact format), <a href="https://redirect.github.com/nodeca/js-yaml/issues/332">#332</a>.</li>
<li>Code snippet created in exceptions now contains multiple lines with line numbers.</li>
<li><code>dump()</code> now serializes <code>undefined</code> as <code>null</code> in collections and removes keys with
<code>undefined</code> in mappings, <a href="https://redirect.github.com/nodeca/js-yaml/issues/571">#571</a>.</li>
<li><code>dump()</code> with <code>skipInvalid=true</code> now serializes invalid items in collections as null.</li>
<li>Custom tags starting with <code>!</code> are now dumped as <code>!tag</code> instead of <code>!&lt;!tag&gt;</code>, <a href="https://redirect.github.com/nodeca/js-yaml/issues/576">#576</a>.</li>
<li>Custom tags starting with <code>tag:yaml.org,2002:</code> are now shorthanded using <code>!!</code>, <a href="https://redirect.github.com/nodeca/js-yaml/issues/258">#258</a>.</li>
</ul>
<h3>Added</h3>
<ul>
<li>Added <code>.mjs</code> (es modules) support.</li>
<li>Added <code>quotingType</code> and <code>forceQuotes</code> options for dumper to configure
string literal style, <a href="https://redirect.github.com/nodeca/js-yaml/issues/290">#290</a>, <a href="https://redirect.github.com/nodeca/js-yaml/issues/529">#529</a>.</li>
<li>Added <code>styles: { '!!null': 'empty' }</code> option for dumper
(serializes <code>{ foo: null }</code> as &quot;<code>foo: </code>&quot;), <a href="https://redirect.github.com/nodeca/js-yaml/issues/570">#570</a>.</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/nodeca/js-yaml/commit/9963d366dfbde0c69722452bcd40b41e7e4160a0"><code>9963d36</code></a> 3.14.2 released</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/10d3c8e70a6888543f5cdb656bb39f73e0ea77c1"><code>10d3c8e</code></a> dist rebuild</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/5278870a17454fe8621dbd8c445c412529525266"><code>5278870</code></a> fix prototype pollution in merge (&lt;&lt;) (<a href="https://redirect.github.com/nodeca/js-yaml/issues/731">#731</a>)</li>
<li>See full diff in <a href="https://github.com/nodeca/js-yaml/compare/3.14.1...3.14.2">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=js-yaml&package-manager=npm_and_yarn&previous-version=3.14.1&new-version=3.14.2)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

You can trigger a rebase of this PR by commenting `@dependabot rebase`.

[//]: # (dependabot-automerge-start)
[//]: # (dependabot-automerge-end)

---

<details>
<summary>Dependabot commands and options</summary>
<br />

You can trigger Dependabot actions by commenting on this PR:
- `@dependabot rebase` will rebase this PR
- `@dependabot recreate` will recreate this PR, overwriting any edits that have been made to it
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.

---

## Discussion

### Comment by @codecov-commenter on February 10, 2026 at 10:22 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1464?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`87ec4bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/87ec4bdb48b1c3c0e5e99802e10cc4341592d4c7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`7ca8346`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7ca83460033c39857fbac1e85bc24ceb16213c40?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 13 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1464   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      899      892    -7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1464?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @swadeley on February 10, 2026 at 10:53 AM UTC

@dependabot rebase

### Comment by @dependabot on February 10, 2026 at 10:54 AM UTC

Looks like this PR has been edited by someone other than Dependabot. That means Dependabot can't rebase it - sorry!

If you're happy for Dependabot to recreate it from scratch, overwriting any edits, you can request `@dependabot recreate`.


### Comment by @swadeley on February 10, 2026 at 10:54 AM UTC

@dependabot recreate

### Comment by @swadeley on February 10, 2026 at 01:02 PM UTC

@dependabot recreate

### Comment by @dependabot on February 16, 2026 at 07:45 AM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

## Reviews

### Review by @swadeley - Approved on February 10, 2026 at 07:52 AM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1464*
