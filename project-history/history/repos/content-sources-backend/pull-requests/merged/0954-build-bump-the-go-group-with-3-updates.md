---
type: pull_request
number: 954
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2025-01-27T04:19:27Z
updated: 2025-01-27T14:13:55Z
closed: 2025-01-27T14:13:50Z
merged: 2025-01-27T14:13:50Z
base_branch: main
head_branch: dependabot/go_modules/go-6dec1ecaaf
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/954
---

# Pull Request #954: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: January 27, 2025 at 04:19 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-6dec1ecaaf`

## Description

Bumps the go group with 3 updates: [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) and [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2).

Updates `github.com/aws/aws-sdk-go-v2` from 1.33.0 to 1.34.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/676a8b1bf0174c8763e19d99b68b988e67e2d398"><code>676a8b1</code></a> Release 2025-01-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1737386a85235b72e9676ed261b72cddb61355df"><code>1737386</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3bc09da29fb3dd079526f7ed141520f69245e445"><code>3bc09da</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb98deef60318ce9a61cda159ebfb0166d88539b"><code>cb98dee</code></a> Fix flex checksum validation cfg (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2981">#2981</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9c764018fe28b27912a0b976614d9e806e3f8268"><code>9c76401</code></a> fix bad changelog type</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ed8a3caa0df9ce36a5b60aebeee201187098d205"><code>ed8a3ca</code></a> Reduce fmt.Sprintf allocations in query encoding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2919">#2919</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5773a9a070873393eb2e7eed37bd647e12e1267"><code>d5773a9</code></a> Add FixUnmarshalIndividualSetValues option to DecoderOptions of dynamodb (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2896">#2896</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58e23dc0311cec940749e34ddfc542dbb00ff7a3"><code>58e23dc</code></a> fix codegen test failing in main</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/640d919419375c1bb9041ffa6dd024b60243a1ed"><code>640d919</code></a> fix broken jmespath waiters in cloudwatch and autoscaling (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2984">#2984</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/613a6cfc607af8470ceec5b7391f9231fa1f98dd"><code>613a6cf</code></a> Optimize/directory traversal (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2970">#2970</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.33.0...v1.34.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.54 to 1.17.55
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/676a8b1bf0174c8763e19d99b68b988e67e2d398"><code>676a8b1</code></a> Release 2025-01-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1737386a85235b72e9676ed261b72cddb61355df"><code>1737386</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3bc09da29fb3dd079526f7ed141520f69245e445"><code>3bc09da</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb98deef60318ce9a61cda159ebfb0166d88539b"><code>cb98dee</code></a> Fix flex checksum validation cfg (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2981">#2981</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9c764018fe28b27912a0b976614d9e806e3f8268"><code>9c76401</code></a> fix bad changelog type</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ed8a3caa0df9ce36a5b60aebeee201187098d205"><code>ed8a3ca</code></a> Reduce fmt.Sprintf allocations in query encoding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2919">#2919</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5773a9a070873393eb2e7eed37bd647e12e1267"><code>d5773a9</code></a> Add FixUnmarshalIndividualSetValues option to DecoderOptions of dynamodb (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2896">#2896</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/58e23dc0311cec940749e34ddfc542dbb00ff7a3"><code>58e23dc</code></a> fix codegen test failing in main</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/640d919419375c1bb9041ffa6dd024b60243a1ed"><code>640d919</code></a> fix broken jmespath waiters in cloudwatch and autoscaling (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2984">#2984</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/613a6cfc607af8470ceec5b7391f9231fa1f98dd"><code>613a6cf</code></a> Optimize/directory traversal (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2970">#2970</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.54...credentials/v1.17.55">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.6 to 1.45.8
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/776903f3dd7208803912e19d3aa25006a7fbdeee"><code>776903f</code></a> Release 2024-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/170b13cac4658e0909b13468d3959f94c358faf3"><code>170b13c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1a02e760211a1d0762f664973ea4d896376a621"><code>c1a02e7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2835f7bb8d4925acb36a5ab7813cca6ac8977cb7"><code>2835f7b</code></a> Fix user agent to add business metrics at the end instead of prepend them (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2">#2</a>...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4965dd291e68f9345c5eae4dec8a0f63bda436"><code>ba4965d</code></a> Release 2024-11-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55149b05036afa39e3b7c6dd945543faee4069e6"><code>55149b0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fd6bb8b363cce171fba74b72d2bb4142b9306b7c"><code>fd6bb8b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65ab4f88b048175d8fc13fab6f72f31b739a2455"><code>65ab4f8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8939ed049c8b3490f9a3dbd4e4d3d56cf22a27f2"><code>8939ed0</code></a> Release 2024-11-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23cf36b2492a3caf859a6235c6911f1c2fb715ca"><code>23cf36b</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/polly/v1.45.6...service/polly/v1.45.8">compare view</a></li>
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

## Reviews

### Review by @jlsherrill - Approved on January 27, 2025 at 02:13 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/954*
