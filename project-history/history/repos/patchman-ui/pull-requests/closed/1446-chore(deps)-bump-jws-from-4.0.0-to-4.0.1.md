---
type: pull_request
number: 1446
title: "chore(deps): bump jws from 4.0.0 to 4.0.1"
state: closed
author: dependabot
created: 2025-12-04T17:01:13Z
updated: 2026-01-21T12:42:49Z
closed: 2026-01-21T12:42:39Z
base_branch: master
head_branch: dependabot/npm_and_yarn/jws-4.0.1
labels: ["dependencies", "patch"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1446
---

# Pull Request #1446: chore(deps): bump jws from 4.0.0 to 4.0.1

**Author**: @dependabot
**Created**: December 04, 2025 at 05:01 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `patch`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/jws-4.0.1`

## Description

Bumps [jws](https://github.com/brianloveswords/node-jws) from 4.0.0 to 4.0.1.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/brianloveswords/node-jws/releases">jws's releases</a>.</em></p>
<blockquote>
<h2>v4.0.1</h2>
<h3>Changed</h3>
<ul>
<li>Fix advisory GHSA-869p-cjfg-cm3x: createSign and createVerify now require
that a non empty secret is provided (via opts.secret, opts.privateKey or opts.key)
when using HMAC algorithms.</li>
<li>Upgrading JWA version to 2.0.1, addressing a compatibility issue for Node &gt;= 25.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/auth0/node-jws/blob/master/CHANGELOG.md">jws's changelog</a>.</em></p>
<blockquote>
<h2>[4.0.1]</h2>
<h3>Changed</h3>
<ul>
<li>Fix advisory GHSA-869p-cjfg-cm3x: createSign and createVerify now require
that a non empty secret is provided (via opts.secret, opts.privateKey or opts.key)
when using HMAC algorithms.</li>
<li>Upgrading JWA version to 2.0.1, adressing a compatibility issue for Node &gt;= 25.</li>
</ul>
<h2>[3.2.3]</h2>
<h3>Changed</h3>
<ul>
<li>Fix advisory GHSA-869p-cjfg-cm3x: createSign and createVerify now require
that a non empty secret is provided (via opts.secret, opts.privateKey or opts.key)
when using HMAC algorithms.</li>
<li>Upgrading JWA version to 1.4.2, adressing a compatibility issue for Node &gt;= 25.</li>
</ul>
<h2>[3.0.0]</h2>
<h3>Changed</h3>
<ul>
<li><strong>BREAKING</strong>: <code>jwt.verify</code> now requires an <code>algorithm</code> parameter, and
<code>jws.createVerify</code> requires an <code>algorithm</code> option. The <code>&quot;alg&quot;</code> field
signature headers is ignored. This mitigates a critical security flaw
in the library which would allow an attacker to generate signatures with
arbitrary contents that would be accepted by <code>jwt.verify</code>. See
<a href="https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/">https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/</a>
for details.</li>
</ul>
<h2><a href="https://github.com/brianloveswords/node-jws/compare/v1.0.1...v2.0.0">2.0.0</a> - 2015-01-30</h2>
<h3>Changed</h3>
<ul>
<li>
<p><strong>BREAKING</strong>: Default payload encoding changed from <code>binary</code> to
<code>utf8</code>. <code>utf8</code> is a is a more sensible default than <code>binary</code> because
many payloads, as far as I can tell, will contain user-facing
strings that could be in any language. (<!-- raw HTML omitted -->[6b6de48]<!-- raw HTML omitted -->)</p>
</li>
<li>
<p>Code reorganization, thanks [<a href="https://github.com/fearphage"><code>@​fearphage</code></a>]! (<!-- raw HTML omitted --><a href="https://github.com/brianloveswords/node-jws/commit/7880050">7880050</a><!-- raw HTML omitted -->)</p>
</li>
</ul>
<h3>Added</h3>
<ul>
<li>Option in all relevant methods for <code>encoding</code>. For those few users
that might be depending on a <code>binary</code> encoding of the messages, this
is for them. (<!-- raw HTML omitted -->[6b6de48]<!-- raw HTML omitted -->)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/auth0/node-jws/commit/34c45b2c04434f925b638de6a061de9339c0ea2e"><code>34c45b2</code></a> Merge commit from fork</li>
<li><a href="https://github.com/auth0/node-jws/commit/49bc39b1f5509a630e0c6849527d8bc66b29ddf5"><code>49bc39b</code></a> version 4.0.1</li>
<li><a href="https://github.com/auth0/node-jws/commit/d42350ccab74db06c95f2279d1674d7d6a1692f4"><code>d42350c</code></a> Enhance tests for HMAC streaming sign and verify</li>
<li><a href="https://github.com/auth0/node-jws/commit/5cb007cf826c70f178c9975d31e949adff75e61b"><code>5cb007c</code></a> Improve secretOrKey initialization in VerifyStream</li>
<li><a href="https://github.com/auth0/node-jws/commit/f9a2e1c8c61ed80d1aa97f03ec32ccb920cf51cb"><code>f9a2e1c</code></a> Improve secret handling in SignStream</li>
<li><a href="https://github.com/auth0/node-jws/commit/b9fb8d30e9c009ade6379f308590f1b0703eefc3"><code>b9fb8d3</code></a> Merge pull request <a href="https://redirect.github.com/brianloveswords/node-jws/issues/102">#102</a> from auth0/SRE-57-Upload-opslevel-yaml</li>
<li><a href="https://github.com/auth0/node-jws/commit/95b75ee56c64d4f8c09c70e9e9662d813bab5685"><code>95b75ee</code></a> Upload OpsLevel YAML</li>
<li><a href="https://github.com/auth0/node-jws/commit/8857ee77623104e5cf9955932165ddf9cea1b72c"><code>8857ee7</code></a> test: remove unused variable (<a href="https://redirect.github.com/brianloveswords/node-jws/issues/96">#96</a>)</li>
<li>See full diff in <a href="https://github.com/brianloveswords/node-jws/compare/v4.0.0...v4.0.1">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~julien.wollscheid">julien.wollscheid</a>, a new releaser for jws since your current version.</p>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=jws&package-manager=npm_and_yarn&previous-version=4.0.0&new-version=4.0.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @dependabot on January 21, 2026 at 12:42 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1446*
