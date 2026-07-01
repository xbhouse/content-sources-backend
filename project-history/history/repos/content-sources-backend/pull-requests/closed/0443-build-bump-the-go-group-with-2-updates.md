---
type: pull_request
number: 443
title: "Build: Bump the go group with 2 updates"
state: closed
author: dependabot
created: 2023-10-25T12:56:40Z
updated: 2023-10-30T04:25:10Z
closed: 2023-10-30T04:25:08Z
base_branch: main
head_branch: dependabot/go_modules/go-7d209ce8da
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/443
---

# Pull Request #443: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: October 25, 2023 at 12:56 PM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-7d209ce8da`

## Description

Bumps the go group with 2 updates: [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) and [go.uber.org/goleak](https://github.com/uber-go/goleak).

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

### Comment by @app-sre-bot on October 25, 2023 at 12:57 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 25, 2023 at 03:39 PM UTC

[test]

### Comment by @jlsherrill on October 26, 2023 at 07:39 PM UTC

/retest

### Comment by @jlsherrill on October 27, 2023 at 05:58 PM UTC

/retest

### Comment by @bsquizz on October 27, 2023 at 07:00 PM UTC

/retest

### Comment by @dependabot on October 30, 2023 at 04:25 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/443*
