---
type: pull_request
number: 744
title: "Build: Bump the go group across 1 directory with 5 updates"
state: merged
author: dependabot
created: 2024-07-15T04:12:23Z
updated: 2024-07-15T16:01:56Z
closed: 2024-07-15T16:01:48Z
merged: 2024-07-15T16:01:47Z
base_branch: main
head_branch: dependabot/go_modules/go-13856c7b8e
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/744
---

# Pull Request #744: Build: Bump the go group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: July 15, 2024 at 04:12 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-13856c7b8e`

## Description

Bumps the go group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.30.1` | `1.30.3` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.24` | `1.17.26` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.37.1` | `1.37.3` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.7.1720031367` | `2024.7.1720720340` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.5.3` | `9.5.4` |


Updates `github.com/aws/aws-sdk-go-v2` from 1.30.1 to 1.30.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0599931efcf551dc63f84ec669d7fb5cfe14f64c"><code>0599931</code></a> Release 2024-07-10.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ca2a30ec8d823300ec948b784c491904eeba8929"><code>ca2a30e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2dff5f9bddaf1f23ea57e1e04d80f6ce2d54221"><code>e2dff5f</code></a> temporarily re-add jmespath dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2712">#2712</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/383fd26928547348efed0ee1f8914d9e7a1b0287"><code>383fd26</code></a> Release 2024-07-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a055f9d9e17eb7ef2206e8e37ba98d661d13e6a"><code>4a055f9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3457953351e6993f57f5ab8e373af8af70be45b"><code>e345795</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/52a10ac239cc1aab2d070fbb004f6445ca0ddcc0"><code>52a10ac</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/adab0de9e4ae75aee4f894c99fd1e2ce2de0035e"><code>adab0de</code></a> remove unused jmespath dependency from main module (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2707">#2707</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e07cc82b25692dce8f68e0b5bd0d0c5cdbcd279"><code>0e07cc8</code></a> Release 2024-07-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e3583451cd8a91dbca2cc22672c2cfa0c7860cf"><code>5e35834</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.1...v1.30.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.24 to 1.17.26
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0599931efcf551dc63f84ec669d7fb5cfe14f64c"><code>0599931</code></a> Release 2024-07-10.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ca2a30ec8d823300ec948b784c491904eeba8929"><code>ca2a30e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2dff5f9bddaf1f23ea57e1e04d80f6ce2d54221"><code>e2dff5f</code></a> temporarily re-add jmespath dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2712">#2712</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/383fd26928547348efed0ee1f8914d9e7a1b0287"><code>383fd26</code></a> Release 2024-07-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a055f9d9e17eb7ef2206e8e37ba98d661d13e6a"><code>4a055f9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3457953351e6993f57f5ab8e373af8af70be45b"><code>e345795</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/52a10ac239cc1aab2d070fbb004f6445ca0ddcc0"><code>52a10ac</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/adab0de9e4ae75aee4f894c99fd1e2ce2de0035e"><code>adab0de</code></a> remove unused jmespath dependency from main module (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2707">#2707</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e07cc82b25692dce8f68e0b5bd0d0c5cdbcd279"><code>0e07cc8</code></a> Release 2024-07-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e3583451cd8a91dbca2cc22672c2cfa0c7860cf"><code>5e35834</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.24...credentials/v1.17.26">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.37.1 to 1.37.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65023f312c12ce2707836e552ab55bb4f03f3345"><code>65023f3</code></a> Release 2023-08-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0742a92e9b623ddaf650eaaea4abc3889b8a17f2"><code>0742a92</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e3340f027ed652b13be24b63af875fe2d4985f7"><code>0e3340f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4526489235229d00a0735e6e28202556ae023729"><code>4526489</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18635b8d0644a641334d0b8cec6b6e1da2c8217c"><code>18635b8</code></a> Add X-Amz-Server-Side-Encryption-Context header to required signed headers al...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4682b6a4b93c5c23a30a064d5f1c8bc11c0636d4"><code>4682b6a</code></a> Achieve DisableMRAP parity in config loaders (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2239">#2239</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ecab8687887d1131a9880aac17e12e589c04b4a"><code>7ecab86</code></a> Release 2023-08-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1ac4adda4fb31aeca5db9e309a428d62ba85bf90"><code>1ac4add</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4f353cd96758cd06e28ea2bfb47d865c52b972d4"><code>4f353cd</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/53f9e53a3e4404ee656500a4f5bf0da450d4f0ff"><code>53f9e53</code></a> remove checks on push because its redundant (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2240">#2240</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.37.1...service/ssm/v1.37.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.7.1720031367 to 2024.7.1720720340
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/6002d88aad298c4af17362448d452d32c1036841"><code>6002d88</code></a> Update pulp bindings to e69db48356e528a464be3da896237b89399b65ca7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/3d94e6efa44f72e8baf45c693477bf13fdf68054"><code>3d94e6e</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/14">#14</a> from jlsherrill/repo</li>
<li><a href="https://github.com/content-services/zest/commit/333c3ad4b7cdd94c55670659bdf67e05fb88b3dc"><code>333c3ad</code></a> check new image repo for pulp builds</li>
<li><a href="https://github.com/content-services/zest/commit/2ce7aed68478f95e469abb73338c4eb74743c327"><code>2ce7aed</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/13">#13</a> from jlsherrill/pulp_change</li>
<li><a href="https://github.com/content-services/zest/commit/407a076e3a98e00e74d710240048daf196750d57"><code>407a076</code></a> Build: use new pulp image</li>
<li><a href="https://github.com/content-services/zest/commit/3d936277c3d86e3613a5828f168aa1a320d1509f"><code>3d93627</code></a> Update pulp bindings to 386d86254354e29d3b864523aed4782b68a448f67968b5e2da9ab...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.7.1720031367...release/v2024.7.1720720340">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.5.3 to 9.5.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.5.4</h2>
<h1>Changes</h1>
<ul>
<li>add test for tls connCheck <a href="https://redirect.github.com/redis/go-redis/issues/3025">#3025</a> (<a href="https://redirect.github.com/redis/go-redis/issues/3047">#3047</a>)</li>
<li>Support Hash-field expiration commands in Pipeline &amp; Fix HExpire HExpireWithArgs expiration (<a href="https://redirect.github.com/redis/go-redis/issues/3038">#3038</a>)</li>
<li>Support Hash-field expiration for 7.4 CE RC2 (<a href="https://redirect.github.com/redis/go-redis/issues/3040">#3040</a>)</li>
<li>fix node routing in slotClosestNode (<a href="https://redirect.github.com/redis/go-redis/issues/3043">#3043</a>)</li>
<li>Update pubsub.go (<a href="https://redirect.github.com/redis/go-redis/issues/3042">#3042</a>)</li>
<li>Change monitor test to run manually (<a href="https://redirect.github.com/redis/go-redis/issues/3041">#3041</a>)</li>
<li>chore(deps): bump rojopolis/spellcheck-github-actions from 0.36.0 to 0.38.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3028">#3028</a>)</li>
<li>Add <code>(*StatusCmd).Bytes()</code> method (<a href="https://redirect.github.com/redis/go-redis/issues/3030">#3030</a>)</li>
<li>chore(deps): bump golang.org/x/net from 0.20.0 to 0.23.0 in /example/otel (<a href="https://redirect.github.com/redis/go-redis/issues/3000">#3000</a>)</li>
<li>Add support for XREAD last entry (<a href="https://redirect.github.com/redis/go-redis/issues/3005">#3005</a>)</li>
<li>Remove Redis 7.4.0 tests from RE (<a href="https://redirect.github.com/redis/go-redis/issues/3035">#3035</a>)</li>
<li>Support NOVALUES parameter for HSCAN (<a href="https://redirect.github.com/redis/go-redis/issues/2925">#2925</a>)</li>
<li>Added test case for CLIENT KILL with MAXAGE option (<a href="https://redirect.github.com/redis/go-redis/issues/2971">#2971</a>)</li>
<li>Change redis version from 7.2 to 7.4 in makefile (<a href="https://redirect.github.com/redis/go-redis/issues/3034">#3034</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>TimeSeries insertion filters for close samples (<a href="https://redirect.github.com/redis/go-redis/issues/3003">#3003</a>)</li>
<li>RediSearch Support (<a href="https://redirect.github.com/redis/go-redis/issues/2801">#2801</a>)</li>
<li>Support Hash-field expiration commands (<a href="https://redirect.github.com/redis/go-redis/issues/2991">#2991</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/b1ron"><code>@​b1ron</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>, <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot], <a href="https://github.com/gerzse"><code>@​gerzse</code></a>, <a href="https://github.com/haines"><code>@​haines</code></a>, <a href="https://github.com/immersedin"><code>@​immersedin</code></a>, <a href="https://github.com/naiqianz"><code>@​naiqianz</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/srikar-jilugu"><code>@​srikar-jilugu</code></a>, <a href="https://github.com/tzongw"><code>@​tzongw</code></a> and <a href="https://github.com/vladvildanov"><code>@​vladvildanov</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/0858ed24e6d08c1940bcb86ccbf8913c761d9ec0"><code>0858ed2</code></a> add test for tls connCheck <a href="https://redirect.github.com/redis/go-redis/issues/3025">#3025</a> (<a href="https://redirect.github.com/redis/go-redis/issues/3047">#3047</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8a0c59b101805f2e45272100134ac4ad83bd69ee"><code>8a0c59b</code></a> TimeSeries insertion filters for close samples (<a href="https://redirect.github.com/redis/go-redis/issues/3003">#3003</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/75398584cfef898f89f411081f0fe05eb9b7cb68"><code>7539858</code></a> Support Hash-field expiration commands in Pipeline &amp; Fix HExpire HExpireWithA...</li>
<li><a href="https://github.com/redis/go-redis/commit/6a584c1e2bcf31ab3bc3830093d0140fbeb6cf3e"><code>6a584c1</code></a> Support Hash-field expiration for 7.4 CE RC2 (<a href="https://redirect.github.com/redis/go-redis/issues/3040">#3040</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/9c1f4f06423d81f5602b274066d65fbd9fc37703"><code>9c1f4f0</code></a> fix node routing in slotClosestNode (<a href="https://redirect.github.com/redis/go-redis/issues/3043">#3043</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/26e0c49acf9328e0c7e3852d6a68e4cc63ce14d7"><code>26e0c49</code></a> Update pubsub.go (<a href="https://redirect.github.com/redis/go-redis/issues/3042">#3042</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/4cf03dbd3981918b66d4d50c64ffa7042d7c1edf"><code>4cf03db</code></a> Change monitor test to run manually (<a href="https://redirect.github.com/redis/go-redis/issues/3041">#3041</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/daf84a091137c24d80da645784b8e5cdcb02b8e6"><code>daf84a0</code></a> chore(deps): bump rojopolis/spellcheck-github-actions (<a href="https://redirect.github.com/redis/go-redis/issues/3028">#3028</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/244a3e22da43d11315a18cf1dede9b86de95888f"><code>244a3e2</code></a> RediSearch Support (<a href="https://redirect.github.com/redis/go-redis/issues/2801">#2801</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2d7382e8ccabf1f295ec740bd6a6cce384cbb620"><code>2d7382e</code></a> Add <code>(*StatusCmd).Bytes()</code> method (<a href="https://redirect.github.com/redis/go-redis/issues/3030">#3030</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.5.3...v9.5.4">compare view</a></li>
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

### Comment by @app-sre-bot on July 15, 2024 at 04:13 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on July 15, 2024 at 12:31 PM UTC

/test

---

## Reviews

### Review by @jlsherrill - Approved on July 15, 2024 at 04:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/744*
