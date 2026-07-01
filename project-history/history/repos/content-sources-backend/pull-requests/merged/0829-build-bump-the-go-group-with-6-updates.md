---
type: pull_request
number: 829
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2024-09-23T04:58:34Z
updated: 2024-09-23T12:08:50Z
closed: 2024-09-23T12:08:42Z
merged: 2024-09-23T12:08:42Z
base_branch: main
head_branch: dependabot/go_modules/go-844434d70e
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/829
---

# Pull Request #829: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: September 23, 2024 at 04:58 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-844434d70e`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.20.3` | `1.20.4` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.30.5` | `1.31.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.32` | `1.17.34` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.39.0` | `1.40.0` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.4.15` | `4.4.16` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.9.1726259707` | `2024.9.1726834826` |

Updates `github.com/prometheus/client_golang` from 1.20.3 to 1.20.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.20.4</h2>
<ul>
<li>[BUGFIX] histograms: Fix a possible data race when appending exemplars vs metrics gather. <a href="https://redirect.github.com/prometheus/client_golang/issues/1623">#1623</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>Unreleased</h2>
<ul>
<li>[BUGFIX] histograms: Fix possible data race when appending exemplars vs metrics gather. <a href="https://redirect.github.com/prometheus/client_golang/issues/1623">#1623</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/05fcde9fe4eb93d3fb7b56ebe51acf80536f0583"><code>05fcde9</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1623">#1623</a> from krajorama/data-race-in-histogram-write</li>
<li><a href="https://github.com/prometheus/client_golang/commit/209f4c041ed1764866f44dd053a8d94aa051c610"><code>209f4c0</code></a> Add changelog</li>
<li><a href="https://github.com/prometheus/client_golang/commit/1e398ccb1259d20836e3003885bdd949cb21e635"><code>1e398cc</code></a> native histogram: Fix race between Write and addExemplar</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.20.3...v1.20.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.30.5 to 1.31.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f445866bcc850b67c71e36882488e10f7c782e3"><code>2f44586</code></a> Release 2024-09-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/22d087682834495c3aebc0c0f1aa2db37c4785a6"><code>22d0876</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5454ab9dfb0bc726f1f79b8d807e81fc7124797a"><code>5454ab9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06150d96305d6b6c19db0a2e5d1c1f4fa4a95612"><code>06150d9</code></a> add tracing and metrics support (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2798">#2798</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18f7b40ea83d7ee75092085609a808b68c4d2000"><code>18f7b40</code></a> Release 2024-09-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e91c9c4e48a47fa92d82251dd6ba0f0edd4bff3d"><code>e91c9c4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6df0a09c479b5e8e59204eda1c6a481c311e5629"><code>6df0a09</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98ae6886ccefe00db38a3088e273e524abf3b196"><code>98ae688</code></a> Release 2024-09-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/222928c4e7cc17d406f6c24f2e69b6747e29d43e"><code>222928c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/887c4de750c99c575f37828f1a646cefadfc4aa7"><code>887c4de</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.5...v1.31.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.32 to 1.17.34
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2f445866bcc850b67c71e36882488e10f7c782e3"><code>2f44586</code></a> Release 2024-09-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/22d087682834495c3aebc0c0f1aa2db37c4785a6"><code>22d0876</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5454ab9dfb0bc726f1f79b8d807e81fc7124797a"><code>5454ab9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06150d96305d6b6c19db0a2e5d1c1f4fa4a95612"><code>06150d9</code></a> add tracing and metrics support (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2798">#2798</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18f7b40ea83d7ee75092085609a808b68c4d2000"><code>18f7b40</code></a> Release 2024-09-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e91c9c4e48a47fa92d82251dd6ba0f0edd4bff3d"><code>e91c9c4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6df0a09c479b5e8e59204eda1c6a481c311e5629"><code>6df0a09</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98ae6886ccefe00db38a3088e273e524abf3b196"><code>98ae688</code></a> Release 2024-09-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/222928c4e7cc17d406f6c24f2e69b6747e29d43e"><code>222928c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/887c4de750c99c575f37828f1a646cefadfc4aa7"><code>887c4de</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.32...credentials/v1.17.34">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.39.0 to 1.40.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8c233fb8a848bbf03d3dcca6898761eb9f39ff16"><code>8c233fb</code></a> Release 2023-09-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d302e513827cb6145bd44b8ba6bc258f486409f9"><code>d302e51</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/01077b015910af3d286b7d3012c83b681bcdb35f"><code>01077b0</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/def1175584cccca2dc1c0f0172d099bca64a6db6"><code>def1175</code></a> Release 2023-09-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4bd57bb24e1c7947aca5af700b34bfcb1272e8b"><code>b4bd57b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94187c8551d09c44e29ad0d5874e7bc76eba0ae6"><code>94187c8</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/297cdcf0c1ca3acf438778d2caa5680972dd644b"><code>297cdcf</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1fe8c9e878a2f909704831729451710b918eeae5"><code>1fe8c9e</code></a> Release 2023-09-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48b3ce975d5de63715cc219526d3622ab39f2715"><code>48b3ce9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0f2416d4dc83c8347d84510fe3fcbc41c63952d"><code>b0f2416</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.39.0...service/s3/v1.40.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.4.15 to 4.4.16
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/4068e465aff0928541b8589fce38ba6393126ac2"><code>4068e46</code></a> Update bindings to release/v4.4.16</li>
<li><a href="https://github.com/content-services/caliri/commit/6eb30466cffa6145e1d87a5e3ea470a167e1048e"><code>6eb3046</code></a> Merge pull request <a href="https://redirect.github.com/content-services/caliri/issues/3">#3</a> from rverdile/cron-for-build</li>
<li><a href="https://github.com/content-services/caliri/commit/abf6239e67433541ddfd9ea3362af3c033556bd4"><code>abf6239</code></a> automate build with cron</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.4.15...release/v4.4.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.9.1726259707 to 2024.9.1726834826
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/0d38dd076cfdb948d602c3d7266b4938429074ca"><code>0d38dd0</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b979a4b6e2de137ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/c84a9d94c5a839c447b92c832139c035ca30ce33"><code>c84a9d9</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47556e325661b7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/3f06904666ec04cc2160c4429c4d9cee8a933f88"><code>3f06904</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a9423a2950b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/edd2070e216fdb19138e02e1b9af8207c070fcd6"><code>edd2070</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a8474542e3542c37d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/b4a1eb687416dff550b54039a4edb11d6debe47d"><code>b4a1eb6</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a8474549bea69c37d356e3e82bb4...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.9.1726259707...release/v2024.9.1726834826">compare view</a></li>
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

### Comment by @app-sre-bot on September 23, 2024 at 05:00 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on September 23, 2024 at 08:16 AM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on September 23, 2024 at 12:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/829*
