---
type: pull_request
number: 668
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2024-05-13T04:47:56Z
updated: 2024-05-13T07:16:58Z
closed: 2024-05-13T07:16:49Z
merged: 2024-05-13T07:16:49Z
base_branch: main
head_branch: dependabot/go_modules/go-8d7e83ef4c
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/668
---

# Pull Request #668: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: May 13, 2024 at 04:47 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-8d7e83ef4c`

## Description

Bumps the go group with 4 updates: [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/prometheus/client_golang` from 1.19.0 to 1.19.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.19.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Security patches for <code>golang.org/x/sys</code> and <code>google.golang.org/protobuf</code></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/lukasauk"><code>@​lukasauk</code></a> made their first contribution in <a href="https://redirect.github.com/prometheus/client_golang/pull/1494">prometheus/client_golang#1494</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.19.0...v1.19.1">https://github.com/prometheus/client_golang/compare/v1.19.0...v1.19.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>Unreleased</h2>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/6e3f4b1091875216850a486b1c2eb0e5ea852f98"><code>6e3f4b1</code></a> Cut 1.19.1 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1494">#1494</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/cad1bfa2b87c26affca469785dd6b166936f696c"><code>cad1bfa</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1454">#1454</a> from prometheus/small-nits</li>
<li><a href="https://github.com/prometheus/client_golang/commit/0aa8c9f68b59bbb17c8c871eb89d615ad19998ed"><code>0aa8c9f</code></a> Rephrase incompatibility with common v0.48.0</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.19.0...v1.19.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.11 to 1.17.13
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0805b749042a533d0ba958f20cff2788db6b600f"><code>0805b74</code></a> Release 2024-05-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ba0718c37c4d4c7e658a5f3bcf46fcf212632d1"><code>8ba0718</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/566a90120ec3a7157dc661e68c48d32f7fe238fd"><code>566a901</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/765a6d269e76fb984c748bf40afe713bdb55978d"><code>765a6d2</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ac88f621440fe4a4183d70e2f7c144a04958e86"><code>0ac88f6</code></a> do NOT serialize empty lists in ec2query (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2630">#2630</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0457ec5571818a8dc0c3305727b3af668d8b0955"><code>0457ec5</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2638">#2638</a> from aws/feat-remove-honeycode</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a8dc075b402637ec448050d7b0d57ac51eef5a52"><code>a8dc075</code></a> add changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a47278ecebafca591cae06fbceb9aa8192b4a2e"><code>9a47278</code></a> remove honeycode from v2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/343ec354ea5aa20b132ce04f35427ede87a26c9c"><code>343ec35</code></a> drop x/net runtime dependency which was only used for testing (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2637">#2637</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1c71d2f79c8b3d7dfba3afc65a37cdfa3dfb7c1a"><code>1c71d2f</code></a> Release 2024-05-09</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.17.11...credentials/v1.17.13">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.35.1 to 1.35.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3f28b5134e7e9a047147b5773af62c6012c34df6"><code>3f28b51</code></a> Release 2023-02-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e8d17fd21083f2ec6ec6139f89b35c841994932"><code>6e8d17f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/60dbdbb0da35b8b8374e0b4e1a9536f1aed7e458"><code>60dbdbb</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/212910ac25d59e959c0a534bd6264a13fb9ca8c8"><code>212910a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/eb8cb66b4433cd5f125abc7c6b74de39c19f25fb"><code>eb8cb66</code></a> Upgrade smithy to 1.27.2, correct query empty list serialization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/24db9f5f6e6d387f115ee8d9393ae6d158e8ef0c"><code>24db9f5</code></a> Update processcreds.CredentialProcessResponse visibility to public (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1921">#1921</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bd3003e29f1fdc34859fcd194a61d6655b8ea492"><code>bd3003e</code></a> dependency: upgrade smithy to 1.27.2 and correct query empty list serialization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d94f223e8c3dc3c63d001ef443a25c056d4131e"><code>0d94f22</code></a> Release 2023-02-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2eec85ed13c74cf77315398edf974481fb5f4dd8"><code>2eec85e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ca6e32eedbb1e707ad3675879e17782e93db67e"><code>4ca6e32</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.35.1...service/ssm/v1.35.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.4.1714507825 to 2024.5.1715373015
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/6c18acd182e92a8cdbddae252b225d9ee7aa407f"><code>6c18acd</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47596388ab51b7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/939131ee5eaad60393985dffdc8b9d2d961e60ec"><code>939131e</code></a> Update pulp bindings to e69db48356e528a464be3da896237b42655922ca7d54eb5892a9d...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.4.1714507825...release/v2024.5.1715373015">compare view</a></li>
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

### Comment by @app-sre-bot on May 13, 2024 at 04:48 AM UTC

Can one of the admins verify this patch?

---

## Reviews

### Review by @swadeley - Approved on May 13, 2024 at 07:12 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/668*
