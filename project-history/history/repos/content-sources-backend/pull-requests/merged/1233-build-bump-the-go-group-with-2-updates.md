---
type: pull_request
number: 1233
title: "Build: Bump the go group with 2 updates"
state: merged
author: dependabot
created: 2025-10-13T04:05:33Z
updated: 2025-10-14T15:47:44Z
closed: 2025-10-14T15:47:39Z
merged: 2025-10-14T15:47:39Z
base_branch: main
head_branch: dependabot/go_modules/go-0c1bddc6cb
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1233
---

# Pull Request #1233: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: October 13, 2025 at 04:05 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-0c1bddc6cb`

## Description

Bumps the go group with 2 updates: [github.com/IBM/sarama](https://github.com/IBM/sarama) and [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest).

Updates `github.com/IBM/sarama` from 1.46.1 to 1.46.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.46.2 (2025-10-10)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<p>A big focus on improving our support for newer protocol versions in this release, particularly supporting a wider range of flexible versions</p>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>chore: support V5 ListOffsets by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3308">IBM/sarama#3308</a></li>
<li>feat: support DeleteGroups V2 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3320">IBM/sarama#3320</a></li>
<li>feat: support DeleteTopics V4 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3321">IBM/sarama#3321</a></li>
<li>feat: support CreateTopics V5 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3322">IBM/sarama#3322</a></li>
<li>feat: support IncrementalAlterConfigs V1 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3319">IBM/sarama#3319</a></li>
<li>feat: support DescribeGroups V5 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3331">IBM/sarama#3331</a></li>
<li>feat: support SyncGroup V4 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3332">IBM/sarama#3332</a></li>
<li>feat: support LeaveGroup V4 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3334">IBM/sarama#3334</a></li>
<li>feat: support Heartbeat V4 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3335">IBM/sarama#3335</a></li>
<li>feat: support JoinGroup V6 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3339">IBM/sarama#3339</a></li>
<li>feat: support DescribeClientQuotas V1 protocol by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3342">IBM/sarama#3342</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: update map rather than create a new map by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3302">IBM/sarama#3302</a></li>
<li>fix: metadata_response valid version range by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3304">IBM/sarama#3304</a></li>
<li>fix: add V4 as valid CreateTopicsResponse by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3305">IBM/sarama#3305</a></li>
<li>fix: correct requiredVersion for DescribeLogDirsResponse by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3306">IBM/sarama#3306</a></li>
<li>fix: extend TestAllocateBodyProtocolVersions for more testing by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3307">IBM/sarama#3307</a></li>
<li>fix: non-flexible ElectLeadersRequest V0/V1 encode/decode by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3312">IBM/sarama#3312</a></li>
<li>fix: make alterPartitionReassignmentsBlock consistent by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3313">IBM/sarama#3313</a></li>
<li>fix: correct decodeRequest bytesRead return value by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3314">IBM/sarama#3314</a></li>
<li>fix: decoder issues by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3327">IBM/sarama#3327</a></li>
<li>fix: improve KIP-511 behaviour on older Kafka clusters by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3328">IBM/sarama#3328</a></li>
<li>fix: return correct error when encoding by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3333">IBM/sarama#3333</a></li>
<li>fix: correct ApiVersionsResponse handling of ErrUnsupportedVersion by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3337">IBM/sarama#3337</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update ossf/scorecard-action action to v2.4.3 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3318">IBM/sarama#3318</a></li>
<li>fix(deps): update module golang.org/x/net to v0.46.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3343">IBM/sarama#3343</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: remove redundant insufficient data checks by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3300">IBM/sarama#3300</a></li>
<li>refactor: use struct rather than map with one entry by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3301">IBM/sarama#3301</a></li>
<li>chore(ci): adopt gotestsum and re-run flakes by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3311">IBM/sarama#3311</a></li>
<li>refactor: Flexible encoding/decoding refactoring by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3317">IBM/sarama#3317</a></li>
<li>chore(fvt): refactor docker-compose and support KRaft by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3323">IBM/sarama#3323</a></li>
<li>fix(fvt): simplify retry using testify's EventuallyWithT by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3324">IBM/sarama#3324</a></li>
<li>chore: add 3.9.1 and 4.1.0 version constants and FVT by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3325">IBM/sarama#3325</a></li>
<li>refactor: get/put for KError by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3326">IBM/sarama#3326</a></li>
<li>refactor: get/put for throttle time ms time.Duration by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3330">IBM/sarama#3330</a></li>
<li>chore(fvt): improve testFuncConsumerGroupMember by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3329">IBM/sarama#3329</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>fix(fvt): check err before usage by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3338">IBM/sarama#3338</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/4c387a20ac093c8f5417d2b437a21eca421f2ce3"><code>4c387a2</code></a> chore(ci): bump actions/stale from 10.0.0 to 10.1.0 in the actions group (<a href="https://redirect.github.com/IBM/sarama/issues/3340">#3340</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/6bf033d48c25d1ccf5ba17993e57d661184ebdcc"><code>6bf033d</code></a> chore(ci): bump github/codeql-action from 3.30.5 to 3.30.6 (<a href="https://redirect.github.com/IBM/sarama/issues/3341">#3341</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/7efec904e6fba1f4bd1076fe733ac50ef3aed792"><code>7efec90</code></a> feat: support DescribeClientQuotas V1 protocol (<a href="https://redirect.github.com/IBM/sarama/issues/3342">#3342</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/fb2bc3a3c66710ae9888b6c11c7e886c0fe0e4be"><code>fb2bc3a</code></a> fix(deps): update module golang.org/x/net to v0.46.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3343">#3343</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a47d24bbc05085c896f0431e16404bca847b62c1"><code>a47d24b</code></a> feat: support JoinGroup V6 protocol</li>
<li><a href="https://github.com/IBM/sarama/commit/d2368dca4ff947b5012d9d9840095e37f9b81014"><code>d2368dc</code></a> fix: string length -1 should be allowed</li>
<li><a href="https://github.com/IBM/sarama/commit/85f7d7b0cf3e3d4224df99c6b11f276c8fc49fd5"><code>85f7d7b</code></a> refactor: fix scope of err check</li>
<li><a href="https://github.com/IBM/sarama/commit/c00816590d9d13447c500e206e7dfc7750d8cd3c"><code>c008165</code></a> fix: correct ApiVersionsResponse handling of ErrUnsupportedVersion (<a href="https://redirect.github.com/IBM/sarama/issues/3337">#3337</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/bde9591d21dfc7b8ea7831b23098e9a41035149d"><code>bde9591</code></a> fix(fvt): use separate client for producer in testProducingMessages</li>
<li><a href="https://github.com/IBM/sarama/commit/447a313f9f900fee94cd3a118ef9f0d598dafe3e"><code>447a313</code></a> fix(fvt): check err before usage</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.46.1...v1.46.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.10.1759886200 to 2025.10.1760102143
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/e8643ce7762fb58761de2de2bdb3e255ffbe7dfb"><code>e8643ce</code></a> Update pulp bindings to e69db48356e528a464be3da896237b3e5ee8ab9fa7d54eb5892a9...</li>
<li><a href="https://github.com/content-services/zest/commit/96eb8ec9c72b531d07c8fb69fd60c36253bd6822"><code>96eb8ec</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73d668ab954137d2e4353b86e...</li>
<li><a href="https://github.com/content-services/zest/commit/475a369e7c4cb4047f1c8d9ab39ad93d98feb5a0"><code>475a369</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a847432254dea3f37d356e3e82bb...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.10.1759886200...release/v2025.10.1760102143">compare view</a></li>
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

### Comment by @swadeley on October 13, 2025 at 08:53 AM UTC

/retest

### Comment by @rverdile on October 14, 2025 at 01:15 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on October 14, 2025 at 03:47 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1233*
