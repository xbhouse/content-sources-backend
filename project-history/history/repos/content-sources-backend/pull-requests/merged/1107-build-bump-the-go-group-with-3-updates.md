---
type: pull_request
number: 1107
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2025-05-12T04:30:24Z
updated: 2025-05-12T06:59:22Z
closed: 2025-05-12T06:59:16Z
merged: 2025-05-12T06:59:16Z
base_branch: main
head_branch: dependabot/go_modules/go-e27edc86e1
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1107
---

# Pull Request #1107: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: May 12, 2025 at 04:30 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-e27edc86e1`

## Description

Bumps the go group with 3 updates: [gorm.io/gorm](https://github.com/go-gorm/gorm), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest).

Updates `gorm.io/gorm` from 1.26.0 to 1.26.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.26.1</h2>
<h2>Changes</h2>
<ul>
<li>fix: int type variable defaultMaxSize overflows in 32-bit environment <a href="https://github.com/iTanken"><code>@​iTanken</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7439">#7439</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/e5b867e785039199800f0549c60fe688dc0304d2"><code>e5b867e</code></a> remove unnecessary session-level configuration for prepared statements</li>
<li><a href="https://github.com/go-gorm/gorm/commit/8c4e8e2d2a63ef019048bd988a2016948605920b"><code>8c4e8e2</code></a> fix: int type variable defaultMaxSize overflows in 32-bit environment (<a href="https://redirect.github.com/go-gorm/gorm/issues/7439">#7439</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.26.0...v1.26.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.48.0 to 1.49.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4fd9126607b61a0fe22af3a6a31b4e4792a66d83"><code>4fd9126</code></a> Release 2024-02-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d22cecd4af5dc762e59c2bd156e3a774cbc1f2db"><code>d22cecd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b640bf508c997be09b65a3b0d7f2be07ee1e7ac6"><code>b640bf5</code></a> Update SDK's smithy-go dependency to v1.20.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94e885c36103117b48764e0e94d88da729b84047"><code>94e885c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b3a6cd78eabbe87d11278117682674044d1e4408"><code>b3a6cd7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f328e6ca9ec4effe5f82d79c2076c59788c05b8"><code>5f328e6</code></a> chore: bump min go to 1.20 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2494">#2494</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3cb2c978aed75796d8db56dab5f496329af8b182"><code>3cb2c97</code></a> ci: fix CI scripts for main branch case (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2491">#2491</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e44db4cf355dbb0324cba2ee49cafed1454ba26"><code>0e44db4</code></a> Release 2024-02-12</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1d8fb3bda76e7fa0373768c10f90aa2aa5281e4a"><code>1d8fb3b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4d556f1d9af36afc7d5109010e3330975877f13"><code>e4d556f</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.48.0...service/s3/v1.49.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.5.1746105577 to 2025.5.1746754376
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/7a6f987d9904ae076f9f2277a1cbfe9cd33d8184"><code>7a6f987</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7dd422683df57952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/c4c12a1a88fe4b10b295b1dea3c06c0b8ea6f390"><code>c4c12a1</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a8474456d4de4c37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/59728141060c630185badd5ed40147ad4fd82312"><code>5972814</code></a> Update pulp bindings to e69db48356e528a464be3da896237b683e4a2b0a7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/9252864c22461cc4e6fd8e9d930c5f58854bff18"><code>9252864</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e89aa6e38f87e8639db952b3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.5.1746105577...release/v2025.5.1746754376">compare view</a></li>
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

### Review by @swadeley - Approved on May 12, 2025 at 06:59 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1107*
