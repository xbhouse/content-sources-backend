---
type: pull_request
number: 447
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2023-10-30T04:25:19Z
updated: 2023-10-30T12:59:51Z
closed: 2023-10-30T12:59:47Z
merged: 2023-10-30T12:59:47Z
base_branch: main
head_branch: dependabot/go_modules/go-bc7d90a435
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/447
---

# Pull Request #447: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: October 30, 2023 at 04:25 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-bc7d90a435`

## Description

Bumps the go group with 4 updates: [github.com/google/uuid](https://github.com/google/uuid), [gorm.io/driver/postgres](https://github.com/go-gorm/postgres), [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) and [go.uber.org/goleak](https://github.com/uber-go/goleak).

Updates `github.com/google/uuid` from 1.3.1 to 1.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/google/uuid/releases">github.com/google/uuid's releases</a>.</em></p>
<blockquote>
<h2>v1.4.0</h2>
<h2><a href="https://github.com/google/uuid/compare/v1.3.1...v1.4.0">1.4.0</a> (2023-10-26)</h2>
<h3>Features</h3>
<ul>
<li>UUIDs slice type with Strings() convenience method (<a href="https://redirect.github.com/google/uuid/issues/133">#133</a>) (<a href="https://github.com/google/uuid/commit/cd5fbbdd02f3e3467ac18940e07e062be1f864b4">cd5fbbd</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>Clarify that Parse's job is to parse but not necessarily validate strings. (Documents current behavior)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/google/uuid/blob/master/CHANGELOG.md">github.com/google/uuid's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/google/uuid/compare/v1.3.1...v1.4.0">1.4.0</a> (2023-10-26)</h2>
<h3>Features</h3>
<ul>
<li>UUIDs slice type with Strings() convenience method (<a href="https://redirect.github.com/google/uuid/issues/133">#133</a>) (<a href="https://github.com/google/uuid/commit/cd5fbbdd02f3e3467ac18940e07e062be1f864b4">cd5fbbd</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>Clarify that Parse's job is to parse but not necessarily validate strings. (Documents current behavior)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/google/uuid/commit/8de8764e294f072b7a2f1a209e88fdcdb1ebc875"><code>8de8764</code></a> chore(master): release 1.4.0 (<a href="https://redirect.github.com/google/uuid/issues/134">#134</a>)</li>
<li><a href="https://github.com/google/uuid/commit/7c22e97ff7647f3b21c3e0870ab335c3889de467"><code>7c22e97</code></a> Clarify the documentation of Parse to state its job is to parse, not validate...</li>
<li><a href="https://github.com/google/uuid/commit/cd5fbbdd02f3e3467ac18940e07e062be1f864b4"><code>cd5fbbd</code></a> feat: UUIDs slice type with Strings() convenience method (<a href="https://redirect.github.com/google/uuid/issues/133">#133</a>)</li>
<li><a href="https://github.com/google/uuid/commit/47f5b3936c94efb365bdfc62716912ed9e66326f"><code>47f5b39</code></a> docs: fix a typo in CONTRIBUTING.md (<a href="https://redirect.github.com/google/uuid/issues/130">#130</a>)</li>
<li><a href="https://github.com/google/uuid/commit/542ddabd47d7bfa79359b7b4e2af7f975354e35f"><code>542ddab</code></a> chore(tests): add Fuzz tests (<a href="https://redirect.github.com/google/uuid/issues/128">#128</a>)</li>
<li><a href="https://github.com/google/uuid/commit/06716f6a60da5ba158f1d53a8236a534968ff76e"><code>06716f6</code></a> chore(tests): Add json.Unmarshal test with empty value cases (<a href="https://redirect.github.com/google/uuid/issues/116">#116</a>)</li>
<li>See full diff in <a href="https://github.com/google/uuid/compare/v1.3.1...v1.4.0">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/driver/postgres` from 1.5.3 to 1.5.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/postgres/commit/68a877c11af52552b7eb5e061a66240a0f6508cb"><code>68a877c</code></a> Update go.mod</li>
<li><a href="https://github.com/go-gorm/postgres/commit/a7214f0c523c2d4f209198fa2cca543fac84d712"><code>a7214f0</code></a> allow ignoring escape (<a href="https://redirect.github.com/go-gorm/postgres/issues/221">#221</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/76303659a1acc223f15afd9dac2c9cedc01bb237"><code>7630365</code></a> fix: change DefaultValueValue to the same as other drivers (<a href="https://redirect.github.com/go-gorm/postgres/issues/222">#222</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/postgres/compare/v1.5.3...v1.5.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2023` from 2023.9.1695058604 to 2023.10.1697742345
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/cb708943c99b0a8158a6ec394c7974a28e271b33"><code>cb70894</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b9798a248d33037ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/e87329752f918bc50ae1f50154392f9983e3b505"><code>e873297</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/3">#3</a> from jlsherrill/compose</li>
<li><a href="https://github.com/content-services/zest/commit/ee96fb58bff95681a33982762e2f88fbfaec6d48"><code>ee96fb5</code></a> update docker compose for new pulp</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2023.9.1695058604...release/v2023.10.1697742345">compare view</a></li>
</ul>
</details>
<br />

Updates `go.uber.org/goleak` from 1.2.1 to 1.3.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/uber-go/goleak/releases">go.uber.org/goleak's releases</a>.</em></p>
<blockquote>
<h2>v1.3.0</h2>
<h3>Fixed</h3>
<ul>
<li>Built-in ignores now match function names more accurately.
They will no longer ignore stacks because of file names
that look similar to function names. (<a href="https://redirect.github.com/uber-go/goleak/issues/112">#112</a>)</li>
</ul>
<h3>Added</h3>
<ul>
<li>Add an <code>IgnoreAnyFunction</code> option to ignore stack traces
that have the provided function anywhere in the stack. (<a href="https://redirect.github.com/uber-go/goleak/issues/113">#113</a>)</li>
<li>Ignore <code>testing.runFuzzing</code> and <code>testing.runFuzzTests</code> alongside
other already-ignored test functions (<code>testing.RunTests</code>, etc). (<a href="https://redirect.github.com/uber-go/goleak/issues/105">#105</a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>Miscellaneous CI-related fixes. (<a href="https://redirect.github.com/uber-go/goleak/issues/103">#103</a>, <a href="https://redirect.github.com/uber-go/goleak/issues/108">#108</a>, <a href="https://redirect.github.com/uber-go/goleak/issues/114">#114</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/uber-go/goleak/blob/master/CHANGELOG.md">go.uber.org/goleak's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/uber-go/goleak/compare/v1.2.1...v1.3.0">1.3.0</a></h2>
<h3>Fixed</h3>
<ul>
<li>Built-in ignores now match function names more accurately.
They will no longer ignore stacks because of file names
that look similar to function names. (<a href="https://redirect.github.com/uber-go/goleak/issues/112">#112</a>)</li>
</ul>
<h3>Added</h3>
<ul>
<li>Add an <code>IgnoreAnyFunction</code> option to ignore stack traces
that have the provided function anywhere in the stack. (<a href="https://redirect.github.com/uber-go/goleak/issues/113">#113</a>)</li>
<li>Ignore <code>testing.runFuzzing</code> and <code>testing.runFuzzTests</code> alongside
other already-ignored test functions (<code>testing.RunTests</code>, etc). (<a href="https://redirect.github.com/uber-go/goleak/issues/105">#105</a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>Miscellaneous CI-related fixes. (<a href="https://redirect.github.com/uber-go/goleak/issues/103">#103</a>, <a href="https://redirect.github.com/uber-go/goleak/issues/108">#108</a>, <a href="https://redirect.github.com/uber-go/goleak/issues/114">#114</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/uber-go/goleak/commit/31095c657c34bba405a8d480db27989aa5f60b9c"><code>31095c6</code></a> Release v1.3.0 (<a href="https://redirect.github.com/uber-go/goleak/issues/115">#115</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/5643445eeef62283311d3a89541d609ae40bc757"><code>5643445</code></a> Revert &quot;Make maxSleep and maxRetires configurable when building options (<a href="https://redirect.github.com/uber-go/goleak/issues/94">#94</a>)...</li>
<li><a href="https://github.com/uber-go/goleak/commit/7b4998fbaa50c54bbb8516b02e0870175cee73b1"><code>7b4998f</code></a> options: Add IgnoreAnyFunction (<a href="https://redirect.github.com/uber-go/goleak/issues/113">#113</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/66916c2a65fe0b2b9c0a07e3ac8415d111a58551"><code>66916c2</code></a> README: Fix build status badge (<a href="https://redirect.github.com/uber-go/goleak/issues/114">#114</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/ecabcf9fed27e23ccaccd816471dd97d1e6452b0"><code>ecabcf9</code></a> ignores: Don't use strings.Contains (<a href="https://redirect.github.com/uber-go/goleak/issues/112">#112</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/91de685688cf4c932e998dbba47a678c051a7d1d"><code>91de685</code></a> stack: Parse all functions (<a href="https://redirect.github.com/uber-go/goleak/issues/111">#111</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/25cbb67949a29168fe22878d215ad49bce416fb1"><code>25cbb67</code></a> internal/stack: Use control flow for state (<a href="https://redirect.github.com/uber-go/goleak/issues/110">#110</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/f995fdb5a4f35b501e1037fcb553f5fab2711340"><code>f995fdb</code></a> ci: Use golangci-lint for linting (<a href="https://redirect.github.com/uber-go/goleak/issues/108">#108</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/03cb6e633cd60859e24327222d86dae4c36b76f4"><code>03cb6e6</code></a> Document incompatibility of VerifyNone with t.Parallel (<a href="https://redirect.github.com/uber-go/goleak/issues/107">#107</a>)</li>
<li><a href="https://github.com/uber-go/goleak/commit/3e8744fc59d447f73dc0ec0769f5892ab8e2f1ed"><code>3e8744f</code></a> Ignore testing.runFuzzing and testing.runFuzzTests (<a href="https://redirect.github.com/uber-go/goleak/issues/105">#105</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/uber-go/goleak/compare/v1.2.1...v1.3.0">compare view</a></li>
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
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Discussion

### Comment by @app-sre-bot on October 30, 2023 at 04:27 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 30, 2023 at 10:35 AM UTC

[test]

### Comment by @jlsherrill on October 30, 2023 at 12:06 PM UTC

/retest

### Comment by @jlsherrill on October 30, 2023 at 12:59 PM UTC

ci test failed due to slow infra, merging

---

## Reviews

### Review by @jlsherrill - Approved on October 30, 2023 at 12:58 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/447*
