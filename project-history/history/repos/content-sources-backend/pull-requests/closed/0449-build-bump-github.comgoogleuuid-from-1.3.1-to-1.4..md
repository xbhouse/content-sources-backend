---
type: pull_request
number: 449
title: "Build: Bump github.com/google/uuid from 1.3.1 to 1.4.0"
state: closed
author: dependabot
created: 2023-10-30T04:26:23Z
updated: 2023-10-30T13:00:26Z
closed: 2023-10-30T13:00:19Z
base_branch: main
head_branch: dependabot/go_modules/github.com/google/uuid-1.4.0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/449
---

# Pull Request #449: Build: Bump github.com/google/uuid from 1.3.1 to 1.4.0

**Author**: @dependabot
**Created**: October 30, 2023 at 04:26 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/github.com/google/uuid-1.4.0`

## Description

[//]: # (dependabot-start)
⚠️  **Dependabot is rebasing this PR** ⚠️ 

Rebasing might not happen immediately, so don't worry if this takes some time.

Note: if you make any changes to this PR yourself, they will take precedence over the rebase.

---

[//]: # (dependabot-end)

Bumps [github.com/google/uuid](https://github.com/google/uuid) from 1.3.1 to 1.4.0.
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


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/google/uuid&package-manager=go_modules&previous-version=1.3.1&new-version=1.4.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @dependabot on October 30, 2023 at 01:00 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/449*
