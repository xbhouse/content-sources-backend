---
type: pull_request
number: 1010
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2025-03-10T04:59:14Z
updated: 2025-03-10T14:10:00Z
closed: 2025-03-10T14:09:51Z
merged: 2025-03-10T14:09:51Z
base_branch: main
head_branch: dependabot/go_modules/go-c1bdcdd262
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1010
---

# Pull Request #1010: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: March 10, 2025 at 04:59 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-c1bdcdd262`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.21.0` | `1.21.1` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.8` | `1.29.9` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.61` | `1.17.62` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.46.0` | `1.46.1` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.78.0` | `1.78.1` |

Updates `github.com/prometheus/client_golang` from 1.21.0 to 1.21.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.21.1 / 2025-03-04</h2>
<p>This release addresses a performance regression introduced in <a href="https://redirect.github.com/prometheus/client_golang/issues/1661">#1661</a> -- thanks to all who <a href="https://redirect.github.com/prometheus/client_golang/issues/1748">reported this quickly</a>:
<a href="https://github.com/chlunde"><code>@​chlunde</code></a>, <a href="https://github.com/dethi"><code>@​dethi</code></a>, <a href="https://github.com/aaronbee"><code>@​aaronbee</code></a> <a href="https://github.com/tsuna"><code>@​tsuna</code></a> <a href="https://github.com/kakkoyun"><code>@​kakkoyun</code></a>  💪🏽. This patch release also fixes the iOS build.</p>
<p>We will be hardening the release process even further (<a href="https://redirect.github.com/prometheus/client_golang/issues/1759">#1759</a>, <a href="https://redirect.github.com/prometheus/client_golang/issues/1761">#1761</a>) to prevent this in future, sorry for the inconvenience!</p>
<p>The high concurrency optimization is planned to be eventually reintroduced, however in a much safer manner, potentially in a separate API.</p>
<ul>
<li>[BUGFIX] prometheus: Revert of <code>Inc</code>, <code>Add</code> and <code>Observe</code> cumulative metric CAS optimizations (<a href="https://redirect.github.com/prometheus/client_golang/issues/1661">#1661</a>), causing regressions on low concurrency cases <a href="https://redirect.github.com/prometheus/client_golang/issues/1757">#1757</a></li>
<li>[BUGFIX] prometheus: Fix GOOS=ios build, broken due to process_collector_* wrong build tags. <a href="https://redirect.github.com/prometheus/client_golang/issues/1758">#1758</a></li>
</ul>
<!-- raw HTML omitted -->
<ul>
<li>Revert &quot;exponential backoff for CAS operations on floats&quot; and cut 1.21.1 by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1757">prometheus/client_golang#1757</a></li>
<li>Fix ios build for 1.21.1 by <a href="https://github.com/bwplotka"><code>@​bwplotka</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1758">prometheus/client_golang#1758</a></li>
</ul>
<!-- raw HTML omitted -->
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.21.0...v1.21.1">https://github.com/prometheus/client_golang/compare/v1.21.0...v1.21.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.21.1 / 2025-03-04</h2>
<ul>
<li>[BUGFIX] prometheus: Revert of <code>Inc</code>, <code>Add</code> and <code>Observe</code> cumulative metric CAS optimizations (<a href="https://redirect.github.com/prometheus/client_golang/issues/1661">#1661</a>), causing regressions on low contention cases.</li>
<li>[BUGFIX] prometheus: Fix GOOS=ios build, broken due to process_collector_* wrong build tags.</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/8a42da3e4bfdc7ea408fcd886064dffc4daffb22"><code>8a42da3</code></a> Fix ios build. (<a href="https://redirect.github.com/prometheus/client_golang/issues/1758">#1758</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/40c62f7a974e27ba11164f7138d60fd7a3b0c326"><code>40c62f7</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1757">#1757</a> from prometheus/revert-121cas</li>
<li><a href="https://github.com/prometheus/client_golang/commit/689f5900164029a73c9a2be6068e1db615107c1f"><code>689f590</code></a> Cut 1.21.1</li>
<li><a href="https://github.com/prometheus/client_golang/commit/9e567a7b4fe821af4fad882552ec06713fb8d291"><code>9e567a7</code></a> Revert &quot;Add: exponential backoff for CAS operations on floats (<a href="https://redirect.github.com/prometheus/client_golang/issues/1661">#1661</a>)&quot;</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.21.0...v1.21.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.8 to 1.29.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bd153c05150d0c3f9d45f5b9353c34df9bab5b5"><code>9bd153c</code></a> Release 2025-03-04.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b967446708e597c3c22c658d8399d11accc9e1f1"><code>b967446</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/188c52a1f53d95e7f99ce0edc73059ce5d0b803c"><code>188c52a</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/defaf5c13afed6bb77e2ee4fafea2bfcd09cea88"><code>defaf5c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8f2dd23c2630bd2ee667a403cf8ea8611c5ce406"><code>8f2dd23</code></a> add test that verifies SRA order of operations (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3025">#3025</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3d547b010751de5a834a749ffdaddff4bcb0d39c"><code>3d547b0</code></a> Release 2025-03-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/554a1497dc16d385e62514b10bc8adacfbd6a8d0"><code>554a149</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3abb221732b42d61c95f5306f528cc2182d37c7a"><code>3abb221</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/95ae39d57f8d986f691754737311102335039848"><code>95ae39d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c62ec38df6cc8e66ecf014365d8bd15c2599bc6d"><code>c62ec38</code></a> create kitchen sink test service (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3023">#3023</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.8...config/v1.29.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.61 to 1.17.62
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bd153c05150d0c3f9d45f5b9353c34df9bab5b5"><code>9bd153c</code></a> Release 2025-03-04.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b967446708e597c3c22c658d8399d11accc9e1f1"><code>b967446</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/188c52a1f53d95e7f99ce0edc73059ce5d0b803c"><code>188c52a</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/defaf5c13afed6bb77e2ee4fafea2bfcd09cea88"><code>defaf5c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8f2dd23c2630bd2ee667a403cf8ea8611c5ce406"><code>8f2dd23</code></a> add test that verifies SRA order of operations (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3025">#3025</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3d547b010751de5a834a749ffdaddff4bcb0d39c"><code>3d547b0</code></a> Release 2025-03-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/554a1497dc16d385e62514b10bc8adacfbd6a8d0"><code>554a149</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3abb221732b42d61c95f5306f528cc2182d37c7a"><code>3abb221</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/95ae39d57f8d986f691754737311102335039848"><code>95ae39d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c62ec38df6cc8e66ecf014365d8bd15c2599bc6d"><code>c62ec38</code></a> create kitchen sink test service (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3023">#3023</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.61...credentials/v1.17.62">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.46.0 to 1.46.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1bb048272ad54b3cbeb3b6da99f4be8090bea5d2"><code>1bb0482</code></a> Release 2024-02-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cc83a2b05532dfe5e6ccba25881887352134fbbd"><code>cc83a2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f44d9967b75499363425c29a5cf40375e9869a7"><code>7f44d99</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4280ccb2671145220be23f299bc0db4a2c061ff5"><code>4280ccb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a264562983a722b90b7a08e5547de85274006e30"><code>a264562</code></a> fix awsjson error deserialization to not expect string code (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2489">#2489</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6ae62c2d6a0db1ecd6312d93a9fdb6ef62d16d33"><code>6ae62c2</code></a> Release 2024-02-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4b7d2951f4bd8ff2f14f58a744131d85fcc67d85"><code>4b7d295</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4f8f9046ab41897266165b3a82323977dd5b6239"><code>4f8f904</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6ef996488fca22310518bbb8121574a93561e34e"><code>6ef9964</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/141feaf30cd2f44bfdf448d14dffb2aca5f303ed"><code>141feaf</code></a> Update README.md (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2498">#2498</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.46.0...service/ssm/v1.46.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.78.0 to 1.78.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9bd153c05150d0c3f9d45f5b9353c34df9bab5b5"><code>9bd153c</code></a> Release 2025-03-04.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b967446708e597c3c22c658d8399d11accc9e1f1"><code>b967446</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/188c52a1f53d95e7f99ce0edc73059ce5d0b803c"><code>188c52a</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/defaf5c13afed6bb77e2ee4fafea2bfcd09cea88"><code>defaf5c</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8f2dd23c2630bd2ee667a403cf8ea8611c5ce406"><code>8f2dd23</code></a> add test that verifies SRA order of operations (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3025">#3025</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3d547b010751de5a834a749ffdaddff4bcb0d39c"><code>3d547b0</code></a> Release 2025-03-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/554a1497dc16d385e62514b10bc8adacfbd6a8d0"><code>554a149</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3abb221732b42d61c95f5306f528cc2182d37c7a"><code>3abb221</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/95ae39d57f8d986f691754737311102335039848"><code>95ae39d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c62ec38df6cc8e66ecf014365d8bd15c2599bc6d"><code>c62ec38</code></a> create kitchen sink test service (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3023">#3023</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.78.0...service/s3/v1.78.1">compare view</a></li>
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

### Review by @jlsherrill - Approved on March 10, 2025 at 02:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1010*
