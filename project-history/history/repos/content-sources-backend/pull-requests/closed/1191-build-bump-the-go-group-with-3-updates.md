---
type: pull_request
number: 1191
title: "Build: Bump the go group with 3 updates"
state: closed
author: dependabot
created: 2025-09-08T04:05:32Z
updated: 2025-09-08T14:29:08Z
closed: 2025-09-08T14:29:06Z
base_branch: main
head_branch: dependabot/go_modules/go-9fb4ec1f9e
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1191
---

# Pull Request #1191: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: September 08, 2025 at 04:05 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9fb4ec1f9e`

## Description

Bumps the go group with 3 updates: [gorm.io/gorm](https://github.com/go-gorm/gorm), [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) and [github.com/redis/go-redis/v9](https://github.com/redis/go-redis).

Updates `gorm.io/gorm` from 1.30.2 to 1.30.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.30.3</h2>
<h2>Changes</h2>
<ul>
<li>No changes</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/38404259fe7935212fc14156989e3de381a2d3f7"><code>3840425</code></a> fix(generics): resolve CurrentTable in Raw/Exec</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.30.2...v1.30.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.23.0 to 1.23.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.23.2 - 2025-09-05</h2>
<p>This release is made to upgrade to prometheus/common v0.66.1, which drops the dependencies github.com/grafana/regexp and go.uber.org/atomic and replaces gopkg.in/yaml.v2 with go.yaml.in/yaml/v2 (a drop-in replacement). There are no functional changes.</p>
<!-- raw HTML omitted -->
<ul>
<li>[release-1.23] Upgrade to prometheus/common@v0.66.1 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1869">prometheus/client_golang#1869</a></li>
<li>[release-1.23] Cut v1.23.2 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1870">prometheus/client_golang#1870</a></li>
</ul>
<!-- raw HTML omitted -->
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.23.1...v1.23.2">https://github.com/prometheus/client_golang/compare/v1.23.1...v1.23.2</a></p>
<h2>v1.23.1 - 2025-09-04</h2>
<p>This release is made to be compatible with a backwards incompatible API change in prometheus/common v0.66.0. There are no functional changes.</p>
<!-- raw HTML omitted -->
<ul>
<li>[release-1.23] Upgrade to prometheus/common v0.66 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1866">prometheus/client_golang#1866</a></li>
<li>[release-1.23] Cut v1.23.1 by <a href="https://github.com/aknuds1"><code>@​aknuds1</code></a> in <a href="https://redirect.github.com/prometheus/client_golang/pull/1867">prometheus/client_golang#1867</a></li>
</ul>
<!-- raw HTML omitted -->
<p><strong>Full Changelog</strong>: <a href="https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.1">https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.1</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.23.2 / 2025-09-05</h2>
<p>This release is made to upgrade to prometheus/common v0.66.1, which drops the dependencies github.com/grafana/regexp and go.uber.org/atomic and replaces gopkg.in/yaml.v2 with go.yaml.in/yaml/v2 (a drop-in replacement).
There are no functional changes.</p>
<h2>1.23.1 / 2025-09-04</h2>
<p>This release is made to be compatible with a backwards incompatible API change
in prometheus/common v0.66.0. There are no functional changes.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/8179a560819f2c64ef6ade70e6ae4c73aecaca3c"><code>8179a56</code></a> Cut v1.23.2 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1870">#1870</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/4142b5908bb6c8f5e412b72de5ea4b927d8c219d"><code>4142b59</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1869">#1869</a> from prometheus/arve/upgrade-common</li>
<li><a href="https://github.com/prometheus/client_golang/commit/4ff40f0d918efc0f59701d13622913805c2425b4"><code>4ff40f0</code></a> Cut v1.23.1 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1867">#1867</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/989b0298944e64f88a54ac9c70cd8c8121f10bc9"><code>989b029</code></a> Upgrade to prometheus/common v0.66 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1866">#1866</a>)</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.23.0...v1.23.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.12.1 to 9.13.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.13.0</h2>
<h2>Highlights</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
<li>Fixes on Read and Write buffer sizes and UniversalOptions</li>
</ul>
<h2>Changes</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>fix(test): fix a timing issue in pubsub test (<a href="https://redirect.github.com/redis/go-redis/pull/3498">#3498</a>)</li>
<li>Allow users to enable read-write splitting in failover mode. (<a href="https://redirect.github.com/redis/go-redis/pull/3482">#3482</a>)</li>
<li>Set the read/write buffer size of the sentinel client to 4KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3476">#3476</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>Support subscriptions against cluster slave nodes (<a href="https://redirect.github.com/redis/go-redis/pull/3480">#3480</a>)</li>
<li>Add wait metrics to otel (<a href="https://redirect.github.com/redis/go-redis/pull/3493">#3493</a>)</li>
<li>Clean failing timeout implementation (<a href="https://redirect.github.com/redis/go-redis/pull/3472">#3472</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>Do not assume that all non-IP hosts are loopbacks (<a href="https://redirect.github.com/redis/go-redis/pull/3085">#3085</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>fix(make test): Add default env in makefile (<a href="https://redirect.github.com/redis/go-redis/pull/3491">#3491</a>)</li>
<li>Update the introduction to running tests in README.md (<a href="https://redirect.github.com/redis/go-redis/pull/3495">#3495</a>)</li>
<li>test: Add comprehensive edge case tests for IncrByFloat command (<a href="https://redirect.github.com/redis/go-redis/pull/3477">#3477</a>)</li>
<li>Set the default read/write buffer size of Redis connection to 32KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3483">#3483</a>)</li>
<li>Bumps test image to 8.2.1-pre (<a href="https://redirect.github.com/redis/go-redis/pull/3478">#3478</a>)</li>
<li>fix UniversalOptions miss ReadBufferSize and WriteBufferSize options (<a href="https://redirect.github.com/redis/go-redis/pull/3485">#3485</a>)</li>
<li>chore(deps): bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/redis/go-redis/pull/3484">#3484</a>)</li>
<li>Removes dry run for stale issues policy (<a href="https://redirect.github.com/redis/go-redis/pull/3471">#3471</a>)</li>
<li>Update otel metrics URL (<a href="https://redirect.github.com/redis/go-redis/pull/3474">#3474</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/cybersmeashish"><code>@​cybersmeashish</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/suever"><code>@​suever</code></a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.13.0 (2025-09-03)</h1>
<h2>Highlights</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
<li>Fixes on Read and Write buffer sizes and UniversalOptions</li>
</ul>
<h2>Changes</h2>
<ul>
<li>Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/pull/3496">#3496</a>)</li>
<li>fix(test): fix a timing issue in pubsub test (<a href="https://redirect.github.com/redis/go-redis/pull/3498">#3498</a>)</li>
<li>Allow users to enable read-write splitting in failover mode. (<a href="https://redirect.github.com/redis/go-redis/pull/3482">#3482</a>)</li>
<li>Set the read/write buffer size of the sentinel client to 4KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3476">#3476</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>Support subscriptions against cluster slave nodes (<a href="https://redirect.github.com/redis/go-redis/pull/3480">#3480</a>)</li>
<li>Add wait metrics to otel (<a href="https://redirect.github.com/redis/go-redis/pull/3493">#3493</a>)</li>
<li>Clean failing timeout implementation (<a href="https://redirect.github.com/redis/go-redis/pull/3472">#3472</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>Do not assume that all non-IP hosts are loopbacks (<a href="https://redirect.github.com/redis/go-redis/pull/3085">#3085</a>)</li>
<li>Ensure that JSON.GET returns Nil response (<a href="https://redirect.github.com/redis/go-redis/pull/3470">#3470</a>)</li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>fix(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/pull/3499">#3499</a>)</li>
<li>fix(make test): Add default env in makefile (<a href="https://redirect.github.com/redis/go-redis/pull/3491">#3491</a>)</li>
<li>Update the introduction to running tests in README.md (<a href="https://redirect.github.com/redis/go-redis/pull/3495">#3495</a>)</li>
<li>test: Add comprehensive edge case tests for IncrByFloat command (<a href="https://redirect.github.com/redis/go-redis/pull/3477">#3477</a>)</li>
<li>Set the default read/write buffer size of Redis connection to 32KiB (<a href="https://redirect.github.com/redis/go-redis/pull/3483">#3483</a>)</li>
<li>Bumps test image to 8.2.1-pre (<a href="https://redirect.github.com/redis/go-redis/pull/3478">#3478</a>)</li>
<li>fix UniversalOptions miss ReadBufferSize and WriteBufferSize options (<a href="https://redirect.github.com/redis/go-redis/pull/3485">#3485</a>)</li>
<li>chore(deps): bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/redis/go-redis/pull/3484">#3484</a>)</li>
<li>Removes dry run for stale issues policy (<a href="https://redirect.github.com/redis/go-redis/pull/3471">#3471</a>)</li>
<li>Update otel metrics URL (<a href="https://redirect.github.com/redis/go-redis/pull/3474">#3474</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/cybersmeashish"><code>@​cybersmeashish</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/htemelski-redis"><code>@​htemelski-redis</code></a>, <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/suever"><code>@​suever</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/52bda7a35ac3b6032a563e23329b912cb0a0a589"><code>52bda7a</code></a> chore(release): 9.13.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3500">#3500</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/19fdc488a76e982038f240642ab00a90d8c10d9d"><code>19fdc48</code></a> chore(otel): register wait metrics (<a href="https://redirect.github.com/redis/go-redis/issues/3499">#3499</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/10121e9e1c8800432e788a02fc0c24d740391702"><code>10121e9</code></a> feat(osscluster): Support subscriptions against cluster slave nodes (<a href="https://redirect.github.com/redis/go-redis/issues/3480">#3480</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6f41b600c5f69da9da9b33dafd1e733d2500b37a"><code>6f41b60</code></a> fix(client): Do not assume that all non-IP hosts are loopbacks (<a href="https://redirect.github.com/redis/go-redis/issues/3085">#3085</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/f0058063a9e2b98e1a3c2f810499db6da4d312e6"><code>f005806</code></a> feat(otel): Add wait metrics to otel (<a href="https://redirect.github.com/redis/go-redis/issues/3493">#3493</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fafec3f3ce2b203257c17b3d9c9a51710d6e3e66"><code>fafec3f</code></a> Pipeliner expose queued commands (<a href="https://redirect.github.com/redis/go-redis/issues/3496">#3496</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6b9cbe8c547646508d209de57549077d388ac904"><code>6b9cbe8</code></a> fix(test): fix a timing issue in pubsub test (<a href="https://redirect.github.com/redis/go-redis/issues/3498">#3498</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/e91f6ced09ee3a4e18bf347b35ffb473388fbee2"><code>e91f6ce</code></a> fix(make test): Add default env in makefile (<a href="https://redirect.github.com/redis/go-redis/issues/3491">#3491</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6bc723834038ec30f06e1028d0c7bdffe18df642"><code>6bc7238</code></a> Fix the ReplicaOnly option does not take effect when using NewFailoverCluster...</li>
<li><a href="https://github.com/redis/go-redis/commit/bb94ac7898d17bf747901f5eed0d4015bfd0ea5e"><code>bb94ac7</code></a> chore(readme): Update the introduction to running tests in README.md (<a href="https://redirect.github.com/redis/go-redis/issues/3495">#3495</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.12.1...v9.13.0">compare view</a></li>
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

### Comment by @swadeley on September 08, 2025 at 07:36 AM UTC

```
 go: golang.org/x/text@v0.29.0 requires go >= 1.24.0 (running go 1.23.12)
Error: Process completed with exit code 1.
```

### Comment by @dependabot on September 08, 2025 at 02:29 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1191*
