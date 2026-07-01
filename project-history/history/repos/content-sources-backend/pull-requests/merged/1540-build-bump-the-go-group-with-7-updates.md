---
type: pull_request
number: 1540
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2026-06-15T04:42:57Z
updated: 2026-06-15T08:45:25Z
closed: 2026-06-15T08:45:23Z
merged: 2026-06-15T08:45:23Z
base_branch: main
head_branch: dependabot/go_modules/go-3e8188967f
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1540
---

# Pull Request #1540: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: June 15, 2026 at 04:42 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-3e8188967f`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.50.1` | `1.50.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.11` | `1.42.0` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.22` | `1.32.24` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.21` | `1.19.23` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.74.4` | `1.75.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.103.1` | `1.103.3` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.5.1779396769` | `2026.6.1780674492` |

Updates `github.com/IBM/sarama` from 1.50.1 to 1.50.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.50.2 (2026-06-05)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(consumer): add support for SyncGroupRequest/Response v5 (KIP-559) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3591">IBM/sarama#3591</a></li>
<li>feat(txn): add protocol support for TxnOffsetCommit v3 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3592">IBM/sarama#3592</a></li>
<li>feat(txn): support consumer group metadata in TxnOffsetCommit v3 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3593">IBM/sarama#3593</a></li>
<li>feat(admin): add protocol support for DeleteRecords v2 (KIP-482) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3594">IBM/sarama#3594</a></li>
<li>feat(protocol): add support for DescribeConfigs v3 and v4 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3596">IBM/sarama#3596</a></li>
<li>feat(admin): add DescribeConfigs for multiple resources by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3600">IBM/sarama#3600</a></li>
<li>feat(consumer): option to cap decompressed batch size by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3604">IBM/sarama#3604</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(admin): retry ACL and SCRAM ops on stale controller by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3598">IBM/sarama#3598</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>fix(deps): update module github.com/pierrec/lz4/v4 to v4.1.27 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3597">IBM/sarama#3597</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.50.1...v1.50.2">https://github.com/IBM/sarama/compare/v1.50.1...v1.50.2</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
<h2>Version 1.50.2 (2026-06-05)</h2>
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(consumer): add support for SyncGroupRequest/Response v5 (KIP-559) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3591">IBM/sarama#3591</a></li>
<li>feat(txn): add protocol support for TxnOffsetCommit v3 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3592">IBM/sarama#3592</a></li>
<li>feat(txn): support consumer group metadata in TxnOffsetCommit v3 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3593">IBM/sarama#3593</a></li>
<li>feat(admin): add protocol support for DeleteRecords v2 (KIP-482) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3594">IBM/sarama#3594</a></li>
<li>feat(protocol): add support for DescribeConfigs v3 and v4 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3596">IBM/sarama#3596</a></li>
<li>feat(admin): add DescribeConfigs for multiple resources by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3600">IBM/sarama#3600</a></li>
<li>feat(consumer): option to cap decompressed batch size by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3604">IBM/sarama#3604</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(admin): retry ACL and SCRAM ops on stale controller by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3598">IBM/sarama#3598</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>fix(deps): update module github.com/pierrec/lz4/v4 to v4.1.27 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3597">IBM/sarama#3597</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.50.1...v1.50.2">https://github.com/IBM/sarama/compare/v1.50.1...v1.50.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/959f157424135d7313749b9eb201a395a9c70c4c"><code>959f157</code></a> feat(consumer): option to cap decompressed batch size (<a href="https://redirect.github.com/IBM/sarama/issues/3604">#3604</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/f6d1c191dcd6114a9b94a8c2fbcf4ac0bbf8c279"><code>f6d1c19</code></a> feat(admin): add DescribeConfigs for multiple resources (<a href="https://redirect.github.com/IBM/sarama/issues/3600">#3600</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3b4f72dcc7e1060dc87bfd10ff7694d6b9792a7b"><code>3b4f72d</code></a> chore(ci): Update docker dependencies (<a href="https://redirect.github.com/IBM/sarama/issues/3603">#3603</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/870f951379fc549f3b1f62342a179750a376f891"><code>870f951</code></a> chore(ci): bump actions/checkout from 6.0.2 to 6.0.3 in the actions group (<a href="https://redirect.github.com/IBM/sarama/issues/3">#3</a>...</li>
<li><a href="https://github.com/IBM/sarama/commit/5cb411e57ecbdb8680cce25ad0a22bacb5ba21b8"><code>5cb411e</code></a> chore(ci): bump github/codeql-action from 4.36.0 to 4.36.1 (<a href="https://redirect.github.com/IBM/sarama/issues/3602">#3602</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/23ebd069883405a125e8c427ecb2106af58f50a3"><code>23ebd06</code></a> feat(protocol): add support for DescribeConfigs v3 and v4 (<a href="https://redirect.github.com/IBM/sarama/issues/3596">#3596</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/45b4840aa51af8207a07292512e9afca92f45206"><code>45b4840</code></a> fix(admin): retry ACL and SCRAM ops on stale controller (<a href="https://redirect.github.com/IBM/sarama/issues/3598">#3598</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3248a774b4446b0a6ded6efb8f9ed039b929c120"><code>3248a77</code></a> feat(admin): add protocol support for DeleteRecords v2 (KIP-482) (<a href="https://redirect.github.com/IBM/sarama/issues/3594">#3594</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/810caff536720ed2950c072b4a323ec6ac6ab684"><code>810caff</code></a> feat(txn): support consumer group metadata in TxnOffsetCommit v3 (<a href="https://redirect.github.com/IBM/sarama/issues/3593">#3593</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/94ffb2bafecaa766444717942d8c20063782a122"><code>94ffb2b</code></a> fix(deps): update module github.com/pierrec/lz4/v4 to v4.1.27 (<a href="https://redirect.github.com/IBM/sarama/issues/3597">#3597</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.50.1...v1.50.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.11 to 1.42.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a3190f79c9e3f7dce5d602be375c07ecd8973cc"><code>9a3190f</code></a> Release 2026-06-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b20dd5bb52bbb193b036c3e81525e7a06f3f819b"><code>b20dd5b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/75a45eaa192d32d7e1247a666f9572a61443b48a"><code>75a45ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e736f55ba6d11d83f9b8d858bbd9ab6c1a883ee2"><code>e736f55</code></a> Add preview of changes for standard retry mode behind flag (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3400">#3400</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba08dc9776ad30b7e5299dd98f3f0360d39656d4"><code>ba08dc9</code></a> Release 2026-06-05.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a67e21ffef3d8b3c9c84e9ce7d90314b646f395"><code>9a67e21</code></a> Revert schema serde (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3442">#3442</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51692f80a1919052106b080468a7d231649d10a4"><code>51692f8</code></a> s3/transfermanager: avoid double-closing concurrentReader channel after read ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f696d5bbafd406a16429239844c6e70627aee935"><code>f696d5b</code></a> Release 2026-06-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7efb8fd6bd7b33d2242586a63a32eca93f774f89"><code>7efb8fd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a420c53c9599f9267285650a61009933de8c9bb"><code>1a420c5</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.11...v1.42.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.22 to 1.32.24
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a3190f79c9e3f7dce5d602be375c07ecd8973cc"><code>9a3190f</code></a> Release 2026-06-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b20dd5bb52bbb193b036c3e81525e7a06f3f819b"><code>b20dd5b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/75a45eaa192d32d7e1247a666f9572a61443b48a"><code>75a45ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e736f55ba6d11d83f9b8d858bbd9ab6c1a883ee2"><code>e736f55</code></a> Add preview of changes for standard retry mode behind flag (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3400">#3400</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba08dc9776ad30b7e5299dd98f3f0360d39656d4"><code>ba08dc9</code></a> Release 2026-06-05.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a67e21ffef3d8b3c9c84e9ce7d90314b646f395"><code>9a67e21</code></a> Revert schema serde (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3442">#3442</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51692f80a1919052106b080468a7d231649d10a4"><code>51692f8</code></a> s3/transfermanager: avoid double-closing concurrentReader channel after read ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f696d5bbafd406a16429239844c6e70627aee935"><code>f696d5b</code></a> Release 2026-06-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7efb8fd6bd7b33d2242586a63a32eca93f774f89"><code>7efb8fd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a420c53c9599f9267285650a61009933de8c9bb"><code>1a420c5</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.22...config/v1.32.24">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.21 to 1.19.23
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a3190f79c9e3f7dce5d602be375c07ecd8973cc"><code>9a3190f</code></a> Release 2026-06-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b20dd5bb52bbb193b036c3e81525e7a06f3f819b"><code>b20dd5b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/75a45eaa192d32d7e1247a666f9572a61443b48a"><code>75a45ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e736f55ba6d11d83f9b8d858bbd9ab6c1a883ee2"><code>e736f55</code></a> Add preview of changes for standard retry mode behind flag (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3400">#3400</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba08dc9776ad30b7e5299dd98f3f0360d39656d4"><code>ba08dc9</code></a> Release 2026-06-05.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a67e21ffef3d8b3c9c84e9ce7d90314b646f395"><code>9a67e21</code></a> Revert schema serde (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3442">#3442</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51692f80a1919052106b080468a7d231649d10a4"><code>51692f8</code></a> s3/transfermanager: avoid double-closing concurrentReader channel after read ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f696d5bbafd406a16429239844c6e70627aee935"><code>f696d5b</code></a> Release 2026-06-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7efb8fd6bd7b33d2242586a63a32eca93f774f89"><code>7efb8fd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a420c53c9599f9267285650a61009933de8c9bb"><code>1a420c5</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.21...credentials/v1.19.23">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.74.4 to 1.75.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2e9697d8ebe330a7435716c2f31b1cea4dff3c0"><code>e2e9697</code></a> Release 2025-01-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6576a0939a79d5f31eef10164750faedd78a45d4"><code>6576a09</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f762573ab5d9286d9751d49091f6a240c12c0742"><code>f762573</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c94df29ecd457e8ec40931fd2fe54d8da2f93ce2"><code>c94df29</code></a> add transfer manager doc header (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2990">#2990</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/880543ce2034570eb3b93c4811289c3b0e55600f"><code>880543c</code></a> revert the revert on the transfer manager beta (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2993">#2993</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8da49e527e317a77ef0f1b2f52b4dc72a4fbd720"><code>8da49e5</code></a> switch to code-generated waiters for remaining services (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2994">#2994</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c7c68659ce67e5b7e18f31bc66068cec9e3d790d"><code>c7c6865</code></a> Release 2025-01-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70f736c5dc0b8652c5fe5c387b2165c3b9beddb1"><code>70f736c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/28731c2bdef3c2555a95632396b6d4936e58099d"><code>28731c2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3505e4b255c327a1fa38f870612c327b93302dc0"><code>3505e4b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/eks/v1.74.4...service/s3/v1.75.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.103.1 to 1.103.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a3190f79c9e3f7dce5d602be375c07ecd8973cc"><code>9a3190f</code></a> Release 2026-06-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b20dd5bb52bbb193b036c3e81525e7a06f3f819b"><code>b20dd5b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/75a45eaa192d32d7e1247a666f9572a61443b48a"><code>75a45ea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e736f55ba6d11d83f9b8d858bbd9ab6c1a883ee2"><code>e736f55</code></a> Add preview of changes for standard retry mode behind flag (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3400">#3400</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba08dc9776ad30b7e5299dd98f3f0360d39656d4"><code>ba08dc9</code></a> Release 2026-06-05.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a67e21ffef3d8b3c9c84e9ce7d90314b646f395"><code>9a67e21</code></a> Revert schema serde (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3442">#3442</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51692f80a1919052106b080468a7d231649d10a4"><code>51692f8</code></a> s3/transfermanager: avoid double-closing concurrentReader channel after read ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f696d5bbafd406a16429239844c6e70627aee935"><code>f696d5b</code></a> Release 2026-06-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7efb8fd6bd7b33d2242586a63a32eca93f774f89"><code>7efb8fd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a420c53c9599f9267285650a61009933de8c9bb"><code>1a420c5</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.103.1...service/s3/v1.103.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.5.1779396769 to 2026.6.1780674492
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/5f86b8ff2d02ef44e4eaf06d5ba09b34c767959a"><code>5f86b8f</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a9728d35b2454157bd286834ede...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.5.1779396769...release/v2026.6.1780674492">compare view</a></li>
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

### Review by @swadeley - Approved on June 15, 2026 at 08:45 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1540*
