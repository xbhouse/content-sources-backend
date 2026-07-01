---
type: pull_request
number: 752
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2024-07-22T04:28:50Z
updated: 2024-07-23T16:32:44Z
closed: 2024-07-22T07:37:58Z
merged: 2024-07-22T07:37:58Z
base_branch: main
head_branch: dependabot/go_modules/go-4444efb690
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/752
---

# Pull Request #752: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: July 22, 2024 at 04:28 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-4444efb690`

## Description

Bumps the go group with 4 updates: [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) and [github.com/redis/go-redis/v9](https://github.com/redis/go-redis).

Updates `github.com/ProtonMail/go-crypto` from 1.1.0-alpha.3-proton to 1.1.0-alpha.5-proton
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>v1.1.0-alpha.5-proton</h2>
<p>This pre-release is v1.1.0-alpha.5 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
<h2>v1.1.0-alpha.4-proton</h2>
<p>This pre-release is v1.1.0-alpha.4 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/5cc763e7d18b851cdea066c107d041a5ef789357"><code>5cc763e</code></a> Fix HMAC generation (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/204">#204</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/5c6f2b4a007b3dfcdfb4d6b01524ae6f4fd9f997"><code>5c6f2b4</code></a> Replace ioutil.ReadAll with io.ReadAll</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/08bd3f7f24ced7fd3cd725c473e8a3e5b19525c0"><code>08bd3f7</code></a> fix(v2): Adapt NewForwardingEntity to refactored NewEntity</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/635733fe337ef25d592744813b6d68005ea4cf19"><code>635733f</code></a> fix(v2): Do not allow encrpytion with a forwarding key</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/6c1edd7a1b34b3813a1b2bfaae06c2b4216d06e6"><code>6c1edd7</code></a> feat: Add symmetric keys to v2</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/736343def1559f058750e4b8180377106908f7f3"><code>736343d</code></a> fix: Address warnings</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/0c847822be897f660382151fe2541b55980b7adc"><code>0c84782</code></a> feat: Add forwarding to v2 api</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/aaf4fba548e61f4a88bed7006550554b0f937a40"><code>aaf4fba</code></a> fix: Address rebase on version 2 issues</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/7c4521c9d34f75c4cd5bf7353e6e8fc6b012414a"><code>7c4521c</code></a> Use fingerprints instead of KeyIDs</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/c38aca0cfeba9f273e90747f4fffe28933eac0f6"><code>c38aca0</code></a> Create a copy of the encrypted key when forwarding</li>
<li>Additional commits viewable in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.3-proton...v1.1.0-alpha.5-proton">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.26 to 1.17.27
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b8a0918056d0ae923f811f525c17066bbb45f050"><code>b8a0918</code></a> Release 2024-07-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03bfb50e05a9a21733c8d890a162d6f9a60b64be"><code>03bfb50</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/616f9a1d3b3ab9e88787ec023b9b6bf26007c048"><code>616f9a1</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/168009d4c428ff78571eaeba9bd03225efa30851"><code>168009d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03768e0d0276b360a6abaa4d30318d4aedc44995"><code>03768e0</code></a> Release 2024-07-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/444505424966e891f064f852c74f6376003fbe1a"><code>4445054</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00acc1e027ff423f22b39bf7678d4615a50c1139"><code>00acc1e</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1b3794c34c13d5bca1dfa19bac440cb09ca4c422"><code>1b3794c</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.26...credentials/v1.17.27">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.7.1720720340 to 2024.7.1721415179
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/0fb9cba19acdab215f979aabf23eac6c625d4bb6"><code>0fb9cba</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84745a8a5ae3037d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/82466d767c13e717752c28dc692f6538e8861457"><code>82466d7</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b7352b3ae88c37d2e4353b86e2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.7.1720720340...release/v2024.7.1721415179">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.5.2 to 9.6.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.6.0</h2>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Support Hash-field expiration commands (<a href="https://redirect.github.com/redis/go-redis/issues/2991">#2991</a>)</li>
<li>Support Hash-field expiration commands in Pipeline &amp; Fix HExpire HExpireWithArgs expiration (<a href="https://redirect.github.com/redis/go-redis/issues/3038">#3038</a>)</li>
<li>Support NOVALUES parameter for HSCAN (<a href="https://redirect.github.com/redis/go-redis/issues/2925">#2925</a>)</li>
<li>Added test case for CLIENT KILL with MAXAGE option (<a href="https://redirect.github.com/redis/go-redis/issues/2971">#2971</a>)</li>
<li>Add support for XREAD last entry (<a href="https://redirect.github.com/redis/go-redis/issues/3005">#3005</a>)</li>
<li>Reduce the type assertion of CheckConn (<a href="https://redirect.github.com/redis/go-redis/issues/3066">#3066</a>)</li>
</ul>
<h2>🛠️ Improvements</h2>
<p>This release includes support for Redis Community Edition (CE) 7.4.0, ensuring compatibility with the latest features and improvements introduced in Redis CE 7.4.0.</p>
<h2>🧰 Maintenance</h2>
<ul>
<li>chore(deps): bump golangci/golangci-lint-action from 4 to 6 (<a href="https://redirect.github.com/redis/go-redis/issues/2993">#2993</a>)</li>
<li>Avoid unnecessary retry delay in cluster client following MOVED and ASK redirection (<a href="https://redirect.github.com/redis/go-redis/issues/3048">#3048</a>)</li>
<li>add test for tls connCheck <a href="https://redirect.github.com/redis/go-redis/issues/3025">#3025</a> (<a href="https://redirect.github.com/redis/go-redis/issues/3047">#3047</a>)</li>
<li>fix node routing in slotClosestNode (<a href="https://redirect.github.com/redis/go-redis/issues/3043">#3043</a>)</li>
<li>Update pubsub.go (<a href="https://redirect.github.com/redis/go-redis/issues/3042">#3042</a>)</li>
<li>Change monitor test to run manually (<a href="https://redirect.github.com/redis/go-redis/issues/3041">#3041</a>)</li>
<li>chore(deps): bump rojopolis/spellcheck-github-actions from 0.36.0 to 0.38.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3028">#3028</a>)</li>
<li>Add <code>(*StatusCmd).Bytes()</code> method (<a href="https://redirect.github.com/redis/go-redis/issues/3030">#3030</a>)</li>
<li>chore(deps): bump golang.org/x/net from 0.20.0 to 0.23.0 in /example/otel (<a href="https://redirect.github.com/redis/go-redis/issues/3000">#3000</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/b1ron"><code>@​b1ron</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/gerzse"><code>@​gerzse</code></a>, <a href="https://github.com/haines"><code>@​haines</code></a>, <a href="https://github.com/immersedin"><code>@​immersedin</code></a>, <a href="https://github.com/naiqianz"><code>@​naiqianz</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/srikar-jilugu"><code>@​srikar-jilugu</code></a>, <a href="https://github.com/tzongw"><code>@​tzongw</code></a>, <a href="https://github.com/vladvildanov"><code>@​vladvildanov</code></a> and <a href="https://github.com/vmihailenco"><code>@​vmihailenco</code></a> <a href="https://github.com/monkey92t"><code>@​monkey92t</code></a></p>
<h2>9.6.0b2</h2>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Add support for XREAD last entry (<a href="https://redirect.github.com/redis/go-redis/issues/3005">#3005</a>)</li>
</ul>
<h2>9.6.0b1</h2>
<h1>Changes</h1>
<h2>🚀 New Features</h2>
<ul>
<li>Support Hash-field expiration commands (<a href="https://redirect.github.com/redis/go-redis/issues/2991">#2991</a>)</li>
<li>Support NOVALUES parameter for HSCAN (<a href="https://redirect.github.com/redis/go-redis/issues/2925">#2925</a>)</li>
<li>Added test case for CLIENT KILL with MAXAGE option (<a href="https://redirect.github.com/redis/go-redis/issues/2971">#2971</a>)</li>
</ul>
<h2>🛠️ Improvements</h2>
<p>This release includes support for Redis Community Edition (CE) 7.4.0, ensuring compatibility with the latest features and improvements introduced in Redis CE 7.4.0.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/44ba2ee08b2d41948a5e499aa4b5378dc6ce7aeb"><code>44ba2ee</code></a> feat: reduce the type assertion of CheckConn (<a href="https://redirect.github.com/redis/go-redis/issues/3066">#3066</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/9cfeb30f627941a769ef25fe0933a0e55dddb8fd"><code>9cfeb30</code></a> Bump version of go-redis to 9.6.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3059">#3059</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/49bfbf2a1e5418d57d95f5b032ecf9bda7fc7922"><code>49bfbf2</code></a> chore(deps): bump golangci/golangci-lint-action from 4 to 6 (<a href="https://redirect.github.com/redis/go-redis/issues/2993">#2993</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2c146dc1042d078bb6ef2ea76f35d7d4a0212357"><code>2c146dc</code></a> Avoid unnecessary retry delay following MOVED and ASK redirection (<a href="https://redirect.github.com/redis/go-redis/issues/3048">#3048</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/233ff45ac9791ac57b7b1dbad864238f2c20a422"><code>233ff45</code></a> add test for tls connCheck <a href="https://redirect.github.com/redis/go-redis/issues/3025">#3025</a> (<a href="https://redirect.github.com/redis/go-redis/issues/3047">#3047</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/5c9c468bd9e899a4eb4914ba47efc501698637d9"><code>5c9c468</code></a> Support Hash-field expiration commands in Pipeline &amp; Fix HExpire HExpireWithA...</li>
<li><a href="https://github.com/redis/go-redis/commit/1a0a68f9cd23e4d69bdd03b86df21e7a1ece9219"><code>1a0a68f</code></a> Support Hash-field expiration for 7.4 CE RC2 (<a href="https://redirect.github.com/redis/go-redis/issues/3040">#3040</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/27dd84c77739364410bf34cb19cd065502e2c71c"><code>27dd84c</code></a> fix node routing in slotClosestNode (<a href="https://redirect.github.com/redis/go-redis/issues/3043">#3043</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/51b6e2c85ff25d282ec097c2d3d0a3e48f2d477e"><code>51b6e2c</code></a> Update pubsub.go (<a href="https://redirect.github.com/redis/go-redis/issues/3042">#3042</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/649d3b6d3121c457225a73ae4af9151245fbeac2"><code>649d3b6</code></a> Change monitor test to run manually (<a href="https://redirect.github.com/redis/go-redis/issues/3041">#3041</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.5.2...v9.6.0">compare view</a></li>
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

### Comment by @app-sre-bot on July 22, 2024 at 04:30 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on July 23, 2024 at 04:32 PM UTC

@dependabot show  gorm.io/gorm ignore conditions

### Comment by @dependabot on July 23, 2024 at 04:32 PM UTC

<details>
<summary>Ignore Conditions</summary>
<br />

|Dependency| Ignore Condition|
|-|-|
| gorm.io/gorm | [< 1.26, > 1.25.5] |
</details>

### Comment by @jlsherrill on July 23, 2024 at 04:32 PM UTC

@dependabot unignore gorm.io/gorm

### Comment by @dependabot on July 23, 2024 at 04:32 PM UTC

OK, I will stop ignoring the gorm.io/gorm dependency.

### Comment by @dependabot on July 23, 2024 at 04:32 PM UTC

Looks like this PR has already been merged. If you've moved back to an older version of these dependencies a new PR will be created next time Dependabot checks your dependencies.

---

## Reviews

### Review by @swadeley - Approved on July 22, 2024 at 07:37 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/752*
