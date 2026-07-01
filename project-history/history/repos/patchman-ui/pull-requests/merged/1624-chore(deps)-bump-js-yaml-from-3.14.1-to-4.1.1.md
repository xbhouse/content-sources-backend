---
type: pull_request
number: 1624
title: "chore(deps): bump js-yaml from 3.14.1 to 4.1.1"
state: merged
author: dependabot
created: 2026-05-14T19:30:36Z
updated: 2026-05-20T17:20:23Z
closed: 2026-05-20T17:20:13Z
merged: 2026-05-20T17:20:13Z
base_branch: master
head_branch: dependabot/npm_and_yarn/js-yaml-4.1.1
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1624
---

# Pull Request #1624: chore(deps): bump js-yaml from 3.14.1 to 4.1.1

**Author**: @dependabot
**Created**: May 14, 2026 at 07:30 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/js-yaml-4.1.1`

## Description

Bumps [js-yaml](https://github.com/nodeca/js-yaml) from 3.14.1 to 4.1.1.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/nodeca/js-yaml/blob/master/CHANGELOG.md">js-yaml's changelog</a>.</em></p>
<blockquote>
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
<li>Added <code>replacer</code> option (similar to option in JSON.stringify), <a href="https://redirect.github.com/nodeca/js-yaml/issues/339">#339</a>.</li>
<li>Custom <code>Tag</code> can now handle all tags or multiple tags with the same prefix, <a href="https://redirect.github.com/nodeca/js-yaml/issues/385">#385</a>.</li>
</ul>
<h3>Fixed</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/nodeca/js-yaml/commit/cc482e775913e6625137572a3712d2826170e53a"><code>cc482e7</code></a> 4.1.1 released</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/50968b862e75866ef90e626572fe0b2f97b55f9f"><code>50968b8</code></a> dist rebuild</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/d092d866031751cb27c12d93f3e2470ad74d678b"><code>d092d86</code></a> lint fix</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/383665ff4248ec2192d1274e934462bb30426879"><code>383665f</code></a> fix prototype pollution in merge (&lt;&lt;)</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/0d3ca7a27b03a6c974790a30a89e456007d62976"><code>0d3ca7a</code></a> README.md: HTTP =&gt; HTTPS (<a href="https://redirect.github.com/nodeca/js-yaml/issues/678">#678</a>)</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/49baadd52af887d2991e2c39a6639baa56d6c71b"><code>49baadd</code></a> doc: 'empty' style option for !!null</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/ba3460eb9d3e4478edcbc29edabe17c2157fc9ce"><code>ba3460e</code></a> Fix demo link (<a href="https://redirect.github.com/nodeca/js-yaml/issues/618">#618</a>)</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/2cef47bebf60da141b78b085f3dea3b5733dcc12"><code>2cef47b</code></a> 4.1.0 released</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/810b149ce2d475109722474d91118f0671b15e20"><code>810b149</code></a> dist rebuild</li>
<li><a href="https://github.com/nodeca/js-yaml/commit/2b5620ed8f03ba0df319fe7710f6d7fd44811742"><code>2b5620e</code></a> Export built-in types, type override now preserves order</li>
<li>Additional commits viewable in <a href="https://github.com/nodeca/js-yaml/compare/3.14.1...4.1.1">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on May 14, 2026 at 07:33 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1624?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 77.58%. Comparing base ([`e587878`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e5878784a0c2ac34e38957069d52a4c6e39c58bf?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`b715b8f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b715b8fad96c9376b2cf8f96ec8dd6b9c3f07b01?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1624   +/-   ##
=======================================
  Coverage   77.58%   77.58%           
=======================================
  Files         103      103           
  Lines        3266     3266           
  Branches      734      729    -5     
=======================================
  Hits         2534     2534           
  Misses        655      655           
  Partials       77       77           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1624?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Commented on May 15, 2026 at 08:27 AM UTC

ACK

### Review by @swadeley - Approved on May 20, 2026 at 05:20 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1624*
