---
type: pull_request
number: 651
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2024-04-29T04:47:04Z
updated: 2024-04-29T07:39:48Z
closed: 2024-04-29T07:39:41Z
merged: 2024-04-29T07:39:41Z
base_branch: main
head_branch: dependabot/go_modules/go-1c95ea5de8
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/651
---

# Pull Request #651: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: April 29, 2024 at 04:47 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-1c95ea5de8`

## Description

Bumps the go group with 3 updates: [github.com/IBM/sarama](https://github.com/IBM/sarama), [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/IBM/sarama` from 1.43.1 to 1.43.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.43.2 (2024-04-25)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>chore(ci): add 32-bit alignment check by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2874">IBM/sarama#2874</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump golang.org/x/net from 0.21.0 to 0.23.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2866">IBM/sarama#2866</a></li>
<li>chore(deps): bump the golang-org-x group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2853">IBM/sarama#2853</a></li>
<li>chore(deps): bump github.com/klauspost/compress from 1.17.7 to 1.17.8 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2857">IBM/sarama#2857</a></li>
<li>chore(deps): bump golang.org/x/net from 0.21.0 to 0.23.0 in /examples/txn_producer in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2865">IBM/sarama#2865</a></li>
<li>chore(deps): bump golang.org/x/net from 0.21.0 to 0.23.0 in /examples/consumergroup in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2867">IBM/sarama#2867</a></li>
<li>chore(deps): bump golang.org/x/net from 0.21.0 to 0.23.0 in /examples/exactly_once in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2868">IBM/sarama#2868</a></li>
<li>chore(deps): bump golang.org/x/net from 0.22.0 to 0.23.0 in /examples/interceptors in the go_modules group by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2869">IBM/sarama#2869</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.43.1...v1.43.2">https://github.com/IBM/sarama/compare/v1.43.1...v1.43.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/0ab2bb77aeca321f41a0953a8c6f52472607a59e"><code>0ab2bb7</code></a> chore(ci): go mod tidy examples tree (<a href="https://redirect.github.com/IBM/sarama/issues/2883">#2883</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a8d8e7a3dbea00655cde0f9575cd029e33c46021"><code>a8d8e7a</code></a> chore(ci): bump actions/upload-artifact from 4.3.1 to 4.3.3 (<a href="https://redirect.github.com/IBM/sarama/issues/2882">#2882</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e4d54f00dae43e0160bf8c5910c70a3f50e6e689"><code>e4d54f0</code></a> chore(ci): re-enable gocache in golangci-lint step</li>
<li><a href="https://github.com/IBM/sarama/commit/18c39e443d2afb3df0d8b3964da114ec87747a71"><code>18c39e4</code></a> chore(ci): bump golangci/golangci-lint-action from 4.0.0 to 5.0.0</li>
<li><a href="https://github.com/IBM/sarama/commit/3d0d4273173435bf58e032dce8ee8882268611ba"><code>3d0d427</code></a> chore(ci): bump docker/setup-buildx-action from 3.2.0 to 3.3.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2880">#2880</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/cf0ca719d5ac8b08267eba54bfd7005df74e0758"><code>cf0ca71</code></a> chore(ci): bump actions/dependency-review-action from 4.1.3 to 4.2.5 (<a href="https://redirect.github.com/IBM/sarama/issues/2879">#2879</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/709e3137c01483ec6769d35cc998c733546dc3c4"><code>709e313</code></a> chore(ci): bump ubi8/ubi-minimal from 8.8 to 8.9 (<a href="https://redirect.github.com/IBM/sarama/issues/2878">#2878</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/681001f4beeee1bcbcda13fc1b1b61b78fc6a2b2"><code>681001f</code></a> fix(fv): use docker's tini/init for FV containers</li>
<li><a href="https://github.com/IBM/sarama/commit/1cb915497044fafac96ec0b3502182c4e02591f3"><code>1cb9154</code></a> fix(ci): fix dependabot docker action</li>
<li><a href="https://github.com/IBM/sarama/commit/532dda654d3e92575fbd0d98a7c3ae93060bcdb6"><code>532dda6</code></a> chore(ci): bump actions/checkout from 4.1.1 to 4.1.4 (<a href="https://redirect.github.com/IBM/sarama/issues/2876">#2876</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.43.1...v1.43.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.4.5-beta.1 to 4.4.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/bc06dd0bd58d4d288c60439272593911acdf94b0"><code>bc06dd0</code></a> Update bindings to release/v4.4.6</li>
<li><a href="https://github.com/content-services/caliri/commit/f45aa3bb027a75e73627fad947e9d4c7d16f2673"><code>f45aa3b</code></a> Merge pull request <a href="https://redirect.github.com/content-services/caliri/issues/1">#1</a> from jlsherrill/run</li>
<li><a href="https://github.com/content-services/caliri/commit/259c6c0c1ca66d3daf1c42ffdcb246bd864f376a"><code>259c6c0</code></a> Allow manual workflow trigger</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.4.5-beta.1...release/v4.4.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.4.1713443852 to 2024.4.1714138895
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/6016bc49b16831d200b964da5fe721579094a6df"><code>6016bc4</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a9724abdee85c57bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/3b862e23439112bc6eb769f462a605cef6d0e251"><code>3b862e2</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73b46a2b43c37d2e4353b86e2...</li>
<li><a href="https://github.com/content-services/zest/commit/bda5057e9a7b915e2be7389233ca90cb665c40dd"><code>bda5057</code></a> Update pulp bindings to e69db48356e528a464be3da896237b4d28edd90a7d54eb5892a9d...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.4.1713443852...release/v2024.4.1714138895">compare view</a></li>
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

### Comment by @app-sre-bot on April 29, 2024 at 04:48 AM UTC

Can one of the admins verify this patch?

---

## Reviews

### Review by @swadeley - Approved on April 29, 2024 at 07:39 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/651*
