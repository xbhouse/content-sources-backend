---
type: pull_request
number: 696
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2024-06-10T04:14:52Z
updated: 2024-06-10T14:26:17Z
closed: 2024-06-10T14:26:10Z
merged: 2024-06-10T14:26:09Z
base_branch: main
head_branch: dependabot/go_modules/go-30f0a0a5e0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/696
---

# Pull Request #696: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: June 10, 2024 at 04:14 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-30f0a0a5e0`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/redhatinsights/app-common-go](https://github.com/redhatinsights/app-common-go) | `1.6.7` | `1.6.8` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.27.0` | `1.27.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.16` | `1.17.18` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.35.5` | `1.35.7` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.5.1717078140` | `2024.6.1717791813` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.5.2` | `9.5.3` |

Updates `github.com/redhatinsights/app-common-go` from 1.6.7 to 1.6.8
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redhatinsights/app-common-go/releases">github.com/redhatinsights/app-common-go's releases</a>.</em></p>
<blockquote>
<h2>v1.6.8</h2>
<h2>What's Changed</h2>
<ul>
<li>Added in hostname by <a href="https://github.com/psav"><code>@​psav</code></a> in <a href="https://redirect.github.com/RedHatInsights/app-common-go/pull/30">RedHatInsights/app-common-go#30</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/app-common-go/compare/v1.6.7...v1.6.8">https://github.com/RedHatInsights/app-common-go/compare/v1.6.7...v1.6.8</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/app-common-go/commit/d9098d75295940b3a394969f9601bad96524fdc0"><code>d9098d7</code></a> Merge pull request <a href="https://redirect.github.com/redhatinsights/app-common-go/issues/30">#30</a> from RedHatInsights/psav/add_hostname</li>
<li><a href="https://github.com/RedHatInsights/app-common-go/commit/615bab7fc28a0ce60f7265e8dce93ba3dd6c2419"><code>615bab7</code></a> Added in hostname</li>
<li>See full diff in <a href="https://github.com/redhatinsights/app-common-go/compare/v1.6.7...v1.6.8">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.27.0 to 1.27.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d872461174d9d665cc722432f821f98564ff0496"><code>d872461</code></a> Release 2024-06-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dd82156cd7a8e147832dc91f9cc13fe36b653905"><code>dd82156</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/85ebac0de6de947471fe99152ec14ff57531d0be"><code>85ebac0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/490e6cb793af49c88decf4a5442a011731b88aa4"><code>490e6cb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4892b059a0a1fb3c8e506ee82444a6fdbab76df0"><code>4892b05</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2668">#2668</a> from aws/feature-clock-skew</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8bdbe6a414748c1111e662a12a55526cbf17dd84"><code>8bdbe6a</code></a> Merge branch 'main' into feature-clock-skew</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c4f8ba325267231aee0010ec7e17030c8f6d30e4"><code>c4f8ba3</code></a> run goimports</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2474a05c366431ce19c355c05a31806e91224962"><code>2474a05</code></a> update snapshot tests</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9fa716b2497e384328135fff32a75a6fa47396eb"><code>9fa716b</code></a> regen clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d365fb33d3a9b6f7fd3bd0153beacc1a507649a"><code>6d365fb</code></a> address feedback from pr, mostly moving things around</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.27.0...v1.27.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.16 to 1.17.18
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d872461174d9d665cc722432f821f98564ff0496"><code>d872461</code></a> Release 2024-06-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dd82156cd7a8e147832dc91f9cc13fe36b653905"><code>dd82156</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/85ebac0de6de947471fe99152ec14ff57531d0be"><code>85ebac0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/490e6cb793af49c88decf4a5442a011731b88aa4"><code>490e6cb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4892b059a0a1fb3c8e506ee82444a6fdbab76df0"><code>4892b05</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2668">#2668</a> from aws/feature-clock-skew</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8bdbe6a414748c1111e662a12a55526cbf17dd84"><code>8bdbe6a</code></a> Merge branch 'main' into feature-clock-skew</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c4f8ba325267231aee0010ec7e17030c8f6d30e4"><code>c4f8ba3</code></a> run goimports</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2474a05c366431ce19c355c05a31806e91224962"><code>2474a05</code></a> update snapshot tests</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9fa716b2497e384328135fff32a75a6fa47396eb"><code>9fa716b</code></a> regen clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d365fb33d3a9b6f7fd3bd0153beacc1a507649a"><code>6d365fb</code></a> address feedback from pr, mostly moving things around</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.16...credentials/v1.17.18">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.35.5 to 1.35.7
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/390cf19a7507fe64e912123df2e1ad4a41a3f2c4"><code>390cf19</code></a> Release 2023-03-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c37c72ab935cbaa8816613808c1153c9da810583"><code>c37c72a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1e51932fdca008df18fc461095401f4d04f375a"><code>d1e5193</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2506101c1841a1816094280739e3a8d0a845d90b"><code>2506101</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c93b5ccbfae85a4c02349f59e3f05182b8885154"><code>c93b5cc</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2051">#2051</a> from aws/add100ContinueCustomization</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c01aac6bfc14a2af5b66b5b696df247e3d68b162"><code>c01aac6</code></a> Keep one changelog for PR</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3780faa4eee262c3c701f586fca007bf4be17115"><code>3780faa</code></a> Keep one changelog for PR</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b94b5b7b7b969e869f38aa708da5d3274095b9a7"><code>b94b5b7</code></a> Merge remote-tracking branch 'origin/add100ContinueCustomization' into add100...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6174ff2c68e774053928a686c20075c6281774bb"><code>6174ff2</code></a> Change some variable name and use operation shape id to represent operation s...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/83491fca07a1af29fb9f6773aba1fbabaee2f329"><code>83491fc</code></a> add changelog to last commit</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.35.5...service/ssm/v1.35.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.5.1717078140 to 2024.6.1717791813
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/0deeb44e7eb644b11ac174b23a769d46eac584ad"><code>0deeb44</code></a> Update pulp bindings to e69db48356e528a464be3da896237b8399366a0a7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/0a3179c746971eed539ff95a5bf46c1e7d79f68b"><code>0a3179c</code></a> Update pulp bindings to e69db48356e528a464be3da896237b83e64364ca7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/98104da76bd2f10c52dfefc4e8d22608debb7b9c"><code>98104da</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b979a2eb92e3f37ae5486b92ad2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.5.1717078140...release/v2024.6.1717791813">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.5.2 to 9.5.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/f752b9a9d5cc158381c2ffe5b13c531037426b39"><code>f752b9a</code></a> Release/v9.5.3 (<a href="https://redirect.github.com/redis/go-redis/issues/3018">#3018</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.5.2...v9.5.3">compare view</a></li>
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

### Comment by @app-sre-bot on June 10, 2024 at 04:15 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on June 10, 2024 at 08:07 AM UTC

/retest

### Comment by @swadeley on June 10, 2024 at 02:09 PM UTC

[Test]

---

## Reviews

### Review by @swadeley - Approved on June 10, 2024 at 02:25 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/696*
