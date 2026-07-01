---
type: pull_request
number: 1486
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2026-05-04T05:06:13Z
updated: 2026-05-05T07:53:09Z
closed: 2026-05-05T07:53:07Z
merged: 2026-05-05T07:53:07Z
base_branch: main
head_branch: dependabot/go_modules/go-36ff66aa15
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1486
---

# Pull Request #1486: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: May 04, 2026 at 05:06 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-36ff66aa15`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/labstack/echo/v4](https://github.com/labstack/echo) | `4.15.1` | `4.15.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.6` | `1.41.7` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.16` | `1.32.17` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.15` | `1.19.16` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.71.0` | `1.72.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.100.0` | `1.100.1` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.8.0` | `1.9.0` |

Updates `github.com/labstack/echo/v4` from 4.15.1 to 4.15.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/labstack/echo/blob/v4.15.2/CHANGELOG.md">github.com/labstack/echo/v4's changelog</a>.</em></p>
<blockquote>
<h2>v4.15.2 - 2026-05-01</h2>
<p><strong>Security</strong></p>
<ul>
<li><code>Context.Scheme()</code> should validate values taken from header by <a href="https://github.com/aldas"><code>@​aldas</code></a> in <a href="https://redirect.github.com/labstack/echo/pull/2962">labstack/echo#2962</a></li>
</ul>
<p>Thanks to <a href="https://github.com/shblue21"><code>@​shblue21</code></a> for reporting this <a href="https://redirect.github.com/labstack/echo/issues/2952">issue</a>.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/labstack/echo/commit/25685e6f5f7b0683105d1386db46ae48eb3de028"><code>25685e6</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/2963">#2963</a> from aldas/v4_changelog_4_15_2</li>
<li><a href="https://github.com/labstack/echo/commit/f9d76893c671df4da9792fc5c122eba01d43c63b"><code>f9d7689</code></a> Changelog for v4.15.2</li>
<li><a href="https://github.com/labstack/echo/commit/37fff28f72264196ad28761316fc7d96fd1c502d"><code>37fff28</code></a> Merge pull request <a href="https://redirect.github.com/labstack/echo/issues/2962">#2962</a> from aldas/v4_valid_proto</li>
<li><a href="https://github.com/labstack/echo/commit/ca4f38a474302aabee93c66bdd0052359882e181"><code>ca4f38a</code></a> Context.Scheme should validate values taken from header</li>
<li><a href="https://github.com/labstack/echo/commit/2e527a70a73b3a93b8346384f607cdc502d3d200"><code>2e527a7</code></a> Update CI, update deps</li>
<li>See full diff in <a href="https://github.com/labstack/echo/compare/v4.15.1...v4.15.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.6 to 1.41.7
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2223642aeafb3b46ae924667ee47d31a1cf5a9d5"><code>2223642</code></a> Release 2026-04-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04c7e460121eafa42577be08289bd0da0de091b2"><code>04c7e46</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f5692571a7afce76e1573da3fbf2180a2c297cc"><code>5f56925</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aac6d2b8fefd47203d3b4bd5f229ca275272ec62"><code>aac6d2b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdaead7d2074e479fc56da77491fe65737413664"><code>bdaead7</code></a> upgrade to smithy-go v1.25.1 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3399">#3399</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/008e12cf64f41c37faeb705f6ce4b77471756f9f"><code>008e12c</code></a> Release 2026-04-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef109d91f0b772aaebc472633edf13ec0fd907ce"><code>ef109d9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6411e6379e509956413b866481ee2362a7cdfc68"><code>6411e63</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e5bf970956db306bf7b5b1bdd8ca71e8cca766b5"><code>e5bf970</code></a> Release 2026-04-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdbb88c9400c80ed2487f7b498a54c6206b5c62a"><code>bdbb88c</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.6...v1.41.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.16 to 1.32.17
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2223642aeafb3b46ae924667ee47d31a1cf5a9d5"><code>2223642</code></a> Release 2026-04-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04c7e460121eafa42577be08289bd0da0de091b2"><code>04c7e46</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f5692571a7afce76e1573da3fbf2180a2c297cc"><code>5f56925</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aac6d2b8fefd47203d3b4bd5f229ca275272ec62"><code>aac6d2b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdaead7d2074e479fc56da77491fe65737413664"><code>bdaead7</code></a> upgrade to smithy-go v1.25.1 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3399">#3399</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/008e12cf64f41c37faeb705f6ce4b77471756f9f"><code>008e12c</code></a> Release 2026-04-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef109d91f0b772aaebc472633edf13ec0fd907ce"><code>ef109d9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6411e6379e509956413b866481ee2362a7cdfc68"><code>6411e63</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e5bf970956db306bf7b5b1bdd8ca71e8cca766b5"><code>e5bf970</code></a> Release 2026-04-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdbb88c9400c80ed2487f7b498a54c6206b5c62a"><code>bdbb88c</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.16...config/v1.32.17">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.15 to 1.19.16
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2223642aeafb3b46ae924667ee47d31a1cf5a9d5"><code>2223642</code></a> Release 2026-04-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04c7e460121eafa42577be08289bd0da0de091b2"><code>04c7e46</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f5692571a7afce76e1573da3fbf2180a2c297cc"><code>5f56925</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aac6d2b8fefd47203d3b4bd5f229ca275272ec62"><code>aac6d2b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdaead7d2074e479fc56da77491fe65737413664"><code>bdaead7</code></a> upgrade to smithy-go v1.25.1 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3399">#3399</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/008e12cf64f41c37faeb705f6ce4b77471756f9f"><code>008e12c</code></a> Release 2026-04-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef109d91f0b772aaebc472633edf13ec0fd907ce"><code>ef109d9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6411e6379e509956413b866481ee2362a7cdfc68"><code>6411e63</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e5bf970956db306bf7b5b1bdd8ca71e8cca766b5"><code>e5bf970</code></a> Release 2026-04-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdbb88c9400c80ed2487f7b498a54c6206b5c62a"><code>bdbb88c</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.15...credentials/v1.19.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.71.0 to 1.72.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9dbe38348e563799aa08abef259c993cfe96f6be"><code>9dbe383</code></a> Release 2025-01-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2fb7c6498bb0940f5f6c3fea87761e96e8a8664a"><code>2fb7c64</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4dd675d9c5cd08e3c4ef1a372c798e5250a8c465"><code>4dd675d</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/11494d46e62b57bd301523aa154560471d56fd74"><code>11494d4</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/85a08fc79360c258e0a80854561e7f3a9633352e"><code>85a08fc</code></a> Release 2025-01-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4babd6cd1d573d15926407bfde15b12da7d2a9fa"><code>4babd6c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8290236384b8646696e2c91d6196b9c9bed9f202"><code>8290236</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/75545e96b1d4118003aa413f355577ed6ae892e0"><code>75545e9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a655009938d14cdf343047a7fd9ebd04894e08b5"><code>a655009</code></a> Release 2024-12-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a7b93d256cdc638582e1e20625cdfcfc38ef5260"><code>a7b93d2</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.71.0...service/s3/v1.72.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.100.0 to 1.100.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2223642aeafb3b46ae924667ee47d31a1cf5a9d5"><code>2223642</code></a> Release 2026-04-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04c7e460121eafa42577be08289bd0da0de091b2"><code>04c7e46</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5f5692571a7afce76e1573da3fbf2180a2c297cc"><code>5f56925</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aac6d2b8fefd47203d3b4bd5f229ca275272ec62"><code>aac6d2b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdaead7d2074e479fc56da77491fe65737413664"><code>bdaead7</code></a> upgrade to smithy-go v1.25.1 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3399">#3399</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/008e12cf64f41c37faeb705f6ce4b77471756f9f"><code>008e12c</code></a> Release 2026-04-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ef109d91f0b772aaebc472633edf13ec0fd907ce"><code>ef109d9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6411e6379e509956413b866481ee2362a7cdfc68"><code>6411e63</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e5bf970956db306bf7b5b1bdd8ca71e8cca766b5"><code>e5bf970</code></a> Release 2026-04-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdbb88c9400c80ed2487f7b498a54c6206b5c62a"><code>bdbb88c</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.100.0...service/s3/v1.100.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.8.0 to 1.9.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.9.0</h2>
<h2>What's Changed</h2>
<ul>
<li>RHCLOUD-46783 - Add AI agent readiness documentation by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/53">project-kessel/kessel-sdk-go#53</a></li>
<li>RHCLOUD-46705 - Implement principal from RH identity helpers by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/50">project-kessel/kessel-sdk-go#50</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.8.0...v1.9.0">https://github.com/project-kessel/kessel-sdk-go/compare/v1.8.0...v1.9.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/e86d9a0262a7793d21d295287963860b7e9f5217"><code>e86d9a0</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/50">#50</a> from lennysgarage/RHCLOUD-46705</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/ccabf492876f95744c3c1b539073a95c3618c492"><code>ccabf49</code></a> Update makefile lint to loop through all example directories and files</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/495a290fbc623be0dbe908f8d5216564ab52d272"><code>495a290</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/53">#53</a> from lennysgarage/agent-readiness</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/1669fda408961f8cee50b5cadd5b5d93eb003d32"><code>1669fda</code></a> Add AI agent readiness documentation</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/f46038aca299e8cc050f310103a284687f2e6056"><code>f46038a</code></a> Match python sdk, pass through domain</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/f39d6040b2fde822ce955d551fbdaac733b57fa2"><code>f39d604</code></a> Address coderabbit feedback</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/2688ba4dc2adf087a276142a0c42545a3cefd8d7"><code>2688ba4</code></a> Implement principal from RH identity helpers</li>
<li>See full diff in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.8.0...v1.9.0">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on May 05, 2026 at 07:52 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1486*
