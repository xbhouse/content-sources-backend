---
type: pull_request
number: 522
title: "Build: Bump the go group with 6 updates"
state: closed
author: dependabot
created: 2024-01-08T04:10:10Z
updated: 2024-01-09T15:30:03Z
closed: 2024-01-09T15:30:01Z
base_branch: main
head_branch: dependabot/go_modules/go-0444256019
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/522
---

# Pull Request #522: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: January 08, 2024 at 04:10 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-0444256019`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.9` | `1.0.10` |
| [github.com/redhatinsights/platform-go-middlewares](https://github.com/redhatinsights/platform-go-middlewares) | `0.20.0` | `1.0.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.24.0` | `1.24.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.16.13` | `1.16.14` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.30.1` | `1.30.2` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.3.1` | `9.4.0` |

Updates `github.com/content-services/yummy` from 1.0.9 to 1.0.10
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/yummy/releases">github.com/content-services/yummy's releases</a>.</em></p>
<blockquote>
<h2>v1.0.10</h2>
<h2>What's Changed</h2>
<ul>
<li>handling compressed comps by <a href="https://github.com/xbhouse"><code>@​xbhouse</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/21">#21</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.9...v1.0.10"><code>v1.0.9...v1.0.10</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/f05aed1b577e2af4c32e8fff1abb94cff2dc5604"><code>f05aed1</code></a> handle compressed comps (<a href="https://redirect.github.com/content-services/yummy/issues/21">#21</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.9...v1.0.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redhatinsights/platform-go-middlewares` from 0.20.0 to 1.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redhatinsights/platform-go-middlewares/releases">github.com/redhatinsights/platform-go-middlewares's releases</a>.</em></p>
<blockquote>
<h2>v1.0.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Create README and tag v1.0.0 by <a href="https://github.com/lzap"><code>@​lzap</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/46">RedHatInsights/platform-go-middlewares#46</a></li>
<li>Add dependabot updates by <a href="https://github.com/subpop"><code>@​subpop</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/48">RedHatInsights/platform-go-middlewares#48</a></li>
<li>Bump actions/checkout from 3 to 4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/49">RedHatInsights/platform-go-middlewares#49</a></li>
<li>Bump actions/setup-go from 3 to 5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/50">RedHatInsights/platform-go-middlewares#50</a></li>
<li>Bump github.com/onsi/ginkgo from 1.8.0 to 1.16.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/51">RedHatInsights/platform-go-middlewares#51</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.38.51 to 1.49.13 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/53">RedHatInsights/platform-go-middlewares#53</a></li>
<li>Bump github.com/go-chi/chi from 4.0.2+incompatible to 4.1.2+incompatible by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/52">RedHatInsights/platform-go-middlewares#52</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/dependabot"><code>@​dependabot</code></a> made their first contribution in <a href="https://redirect.github.com/RedHatInsights/platform-go-middlewares/pull/49">RedHatInsights/platform-go-middlewares#49</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/platform-go-middlewares/compare/v0.20.0...v1.0.0">https://github.com/RedHatInsights/platform-go-middlewares/compare/v0.20.0...v1.0.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/1d808e9ce96a9d87b05464aaade1deb453c921eb"><code>1d808e9</code></a> Bump github.com/go-chi/chi from 4.0.2+incompatible to 4.1.2+incompatible</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/997ce1c461d7765eef31317020aafcd39154f663"><code>997ce1c</code></a> Bump github.com/aws/aws-sdk-go from 1.38.51 to 1.49.13</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/f67f4efa8a3399b2d1011c51e71cfa6cc62c1830"><code>f67f4ef</code></a> Bump github.com/onsi/ginkgo from 1.8.0 to 1.16.5</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/4eddedc7f961b2b036cbba0607570fa1291328e5"><code>4eddedc</code></a> Bump actions/setup-go from 3 to 5</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/a9c41d97fd605d4ce07070456ea3d4e56d6d859e"><code>a9c41d9</code></a> Bump actions/checkout from 3 to 4</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/e194b1b674cf9e09d62d8dd7b8a86d90a50a5744"><code>e194b1b</code></a> Add dependabot updates</li>
<li><a href="https://github.com/RedHatInsights/platform-go-middlewares/commit/e3779317d1aa349fdb5c7e1025d0b2bdacfe2ea2"><code>e377931</code></a> Create README and tag v1.0.0</li>
<li>See full diff in <a href="https://github.com/redhatinsights/platform-go-middlewares/compare/v0.20.0...v1.0.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.24.0 to 1.24.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/47dd1b1bcbde244357a82ef00fa6a61a9bfb9b39"><code>47dd1b1</code></a> Release 2024-01-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23145e3e605a93582020facfe13350b4153714e1"><code>23145e3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/00e57bb7feb2d104387103aa4fbcd55dfff3a6a7"><code>00e57bb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/73e1a99f2fa858ca56627779852768a9198ba057"><code>73e1a99</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2441">#2441</a> from RanVaknin/fix-documentation-config</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0f8ad11593c219c52ad7fb12998c75ade39dc7ad"><code>0f8ad11</code></a> Fix SRA auth trailing checksum retry bug (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2438">#2438</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/efbc5aa622a882167129e69a88aa50c730cd1904"><code>efbc5aa</code></a> Release 2024-01-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78357bb87682230e24b15c01e807d7375a9474e4"><code>78357bb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e465ddd60d18e91b34de5917534cfa1542323027"><code>e465ddd</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/384ad3b7ec21eecb3c6c38b69f86fb6342906b11"><code>384ad3b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1126a91e53a20b18bc1db74225a8417bfb63f427"><code>1126a91</code></a> changelog added</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.24.0...v1.24.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.13 to 1.16.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/289bbd880a1aa7f88dc5a8be4a48869861aa949b"><code>289bbd8</code></a> Release 2022-09-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4169b1a526e46fd72ae275925d4f0a09152355bf"><code>4169b1a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e02802d3d2ccda7699a55800f7a8e58fa59cc2b8"><code>e02802d</code></a> Update SDK's smithy-go dependency to v1.13.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9f98fe25508538ae25d0b07ffe2709f5fae6e0d4"><code>9f98fe2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/824a8eb58b66dbcad91432561577be45ccbefadf"><code>824a8eb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/05a6aab6a1c03d3835dfe3285fabc2e275229b70"><code>05a6aab</code></a> Release 2022-09-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a4fbfe93646991e088c05daeb1586dd685973ce"><code>5a4fbfe</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9fde2d8c6bb409930211633e24afe7da8286ef7"><code>b9fde2d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18081400926772d62a7e8476cb156762deaaa5fc"><code>1808140</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.13...v1.16.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.30.1 to 1.30.2
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.30.1...service/s3/v1.30.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.3.1 to 9.4.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.4.0</h2>
<h1>Changes</h1>
<h2>Breaking Changes</h2>
<ul>
<li>Revert <a href="https://redirect.github.com/redis/go-redis/issues/2818">#2818</a> due to it be a breaking change (<a href="https://redirect.github.com/redis/go-redis/issues/2861">#2861</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>Add Suffix support to default client set info (<a href="https://redirect.github.com/redis/go-redis/issues/2852">#2852</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>return raw value instead of function calling in Result() (<a href="https://redirect.github.com/redis/go-redis/issues/2831">#2831</a>)</li>
<li>Add Redis Enterprise tests (<a href="https://redirect.github.com/redis/go-redis/issues/2847">#2847</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/lowang-bh"><code>@​lowang-bh</code></a> and <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/531f06861b26fc01b1b75f777d7155d286b0943c"><code>531f068</code></a> 9.4.0 (<a href="https://redirect.github.com/redis/go-redis/issues/2862">#2862</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/27581fcb437fa7387d1dcfb6a43c5ccf1f4d2273"><code>27581fc</code></a> Change Z Member type to interface (<a href="https://redirect.github.com/redis/go-redis/issues/2861">#2861</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/a32be3d93d60b9b9a3711bb3c248684246168bae"><code>a32be3d</code></a> Add Suffix support to default client set info (<a href="https://redirect.github.com/redis/go-redis/issues/2852">#2852</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/d8e3e95866c094700ffc8f7e2705c2b8fb7a3793"><code>d8e3e95</code></a> return raw value instead of funcation calling in Result() (<a href="https://redirect.github.com/redis/go-redis/issues/2831">#2831</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/00229238c00584052588ee33f448ca31e0cc543d"><code>0022923</code></a> Change Env vars in RE CI (<a href="https://redirect.github.com/redis/go-redis/issues/2856">#2856</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/2e121910233d77b23d76a64e7d37628d1a73c0c1"><code>2e12191</code></a> Change RE build in CI (<a href="https://redirect.github.com/redis/go-redis/issues/2855">#2855</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b76230924354e418e500679c0460706dfcf5f82b"><code>b762309</code></a> Add RE tests (<a href="https://redirect.github.com/redis/go-redis/issues/2847">#2847</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.3.1...v9.4.0">compare view</a></li>
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

### Comment by @app-sre-bot on January 08, 2024 at 04:12 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on January 08, 2024 at 01:26 PM UTC

[test]

### Comment by @jlsherrill on January 09, 2024 at 01:40 PM UTC

[test]

### Comment by @dependabot on January 09, 2024 at 03:29 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/522*
