---
type: pull_request
number: 1408
title: "Build: Bump the go group with 7 updates"
state: merged
author: dependabot
created: 2026-03-09T04:39:49Z
updated: 2026-03-09T10:56:31Z
closed: 2026-03-09T10:56:29Z
merged: 2026-03-09T10:56:29Z
base_branch: main
head_branch: dependabot/go_modules/go-63b0a3c54b
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1408
---

# Pull Request #1408: Build: Bump the go group with 7 updates

**Author**: @dependabot
**Created**: March 09, 2026 at 04:39 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-63b0a3c54b`

## Description

Bumps the go group with 7 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.2` | `1.41.3` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.10` | `1.32.11` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.10` | `1.19.11` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.63.2` | `1.64.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.96.2` | `1.96.4` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.4.0` | `1.5.0` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.79.1` | `1.79.2` |

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.2 to 1.41.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/967c9203011d777f902647cf450b9770ac282af8"><code>967c920</code></a> Release 2026-03-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23daf2be0b1306a205b3c531f5c41f43a78eb0e1"><code>23daf2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae690a7271153ba5f29301306f02ba2da685b85d"><code>ae690a7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/42acfb99f7d25603cabbcce98511dc80a9b50e0d"><code>42acfb9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/790446eef174929d0796cbd608c43cdea0bba443"><code>790446e</code></a> Run new Go tool 'go fix' on non-codegen files (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3337">#3337</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e64a7d0f3ff5248deaf30e95b113355b1db3816e"><code>e64a7d0</code></a> migrate protocol test codegen to smithy-go (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3335">#3335</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5b9ac6e3a465972c1f735ea67ab46e8fd9a6a09"><code>d5b9ac6</code></a> Bump Go version to 1.24 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3334">#3334</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ddb9538ac4bfe2b4bc0a8292188f45afd3c8f35e"><code>ddb9538</code></a> Release 2026-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/deda70f050cf39c85aca091d7b90717889ce1166"><code>deda70f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d5acbdc0cb17f3229165a0586f7d5b4a67a67a2"><code>7d5acbd</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.2...v1.41.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.10 to 1.32.11
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/967c9203011d777f902647cf450b9770ac282af8"><code>967c920</code></a> Release 2026-03-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23daf2be0b1306a205b3c531f5c41f43a78eb0e1"><code>23daf2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae690a7271153ba5f29301306f02ba2da685b85d"><code>ae690a7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/42acfb99f7d25603cabbcce98511dc80a9b50e0d"><code>42acfb9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/790446eef174929d0796cbd608c43cdea0bba443"><code>790446e</code></a> Run new Go tool 'go fix' on non-codegen files (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3337">#3337</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e64a7d0f3ff5248deaf30e95b113355b1db3816e"><code>e64a7d0</code></a> migrate protocol test codegen to smithy-go (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3335">#3335</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5b9ac6e3a465972c1f735ea67ab46e8fd9a6a09"><code>d5b9ac6</code></a> Bump Go version to 1.24 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3334">#3334</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ddb9538ac4bfe2b4bc0a8292188f45afd3c8f35e"><code>ddb9538</code></a> Release 2026-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/deda70f050cf39c85aca091d7b90717889ce1166"><code>deda70f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d5acbdc0cb17f3229165a0586f7d5b4a67a67a2"><code>7d5acbd</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.10...config/v1.32.11">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.10 to 1.19.11
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/967c9203011d777f902647cf450b9770ac282af8"><code>967c920</code></a> Release 2026-03-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23daf2be0b1306a205b3c531f5c41f43a78eb0e1"><code>23daf2b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae690a7271153ba5f29301306f02ba2da685b85d"><code>ae690a7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/42acfb99f7d25603cabbcce98511dc80a9b50e0d"><code>42acfb9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/790446eef174929d0796cbd608c43cdea0bba443"><code>790446e</code></a> Run new Go tool 'go fix' on non-codegen files (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3337">#3337</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e64a7d0f3ff5248deaf30e95b113355b1db3816e"><code>e64a7d0</code></a> migrate protocol test codegen to smithy-go (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3335">#3335</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5b9ac6e3a465972c1f735ea67ab46e8fd9a6a09"><code>d5b9ac6</code></a> Bump Go version to 1.24 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3334">#3334</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ddb9538ac4bfe2b4bc0a8292188f45afd3c8f35e"><code>ddb9538</code></a> Release 2026-02-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/deda70f050cf39c85aca091d7b90717889ce1166"><code>deda70f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d5acbdc0cb17f3229165a0586f7d5b4a67a67a2"><code>7d5acbd</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/iam/v1.19.10...credentials/v1.19.11">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.63.2 to 1.64.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6191b25230d7cf4333b1e3bc2e60400fb56a4591"><code>6191b25</code></a> Release 2024-10-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6b963d71ea3b0e63acbbba67d6f2ad725af67f73"><code>6b963d7</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/71fb8f0b33c0fe73e98456beb34fb3c0b6ec66e1"><code>71fb8f0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b132e7e9f290e38657de2666fa3c7f3d827f6bdf"><code>b132e7e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c5e364e134b6207b81c762ad18f41d09ec531f21"><code>c5e364e</code></a> Release 2024-10-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d4103b74479201ef1934bcee5414c15f0bada5f1"><code>d4103b7</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e237ef3019e8ef46b2871331f68f61b503fad3d1"><code>e237ef3</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/98c6308bfd1e3154f8f35c33eb0e36034a5eecf7"><code>98c6308</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be2f3e03a2d633228d04f560a87e5459655083e5"><code>be2f3e0</code></a> Release 2024-09-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/88e60a128ba0ed1aa4e4c3e87b5be4633cf0c491"><code>88e60a1</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.63.2...service/s3/v1.64.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.96.2 to 1.96.4
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/38de242faf3a1397673276165379524b2d581a72"><code>38de242</code></a> Release 2026-03-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ecc8d0553abec63d7525a7e1a8305cbd0e7f62d3"><code>ecc8d05</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/17f88d86269b84cd3747c35b545cc8b71b836bf1"><code>17f88d8</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32a2e29e85af4a2e545bc9638c66ee8a13715755"><code>32a2e29</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/10568eb7a3c808d4e80897a3890923a892d3450c"><code>10568eb</code></a> read the correct signing name for sigv4a (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3338">#3338</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/20f32535850690d0afffda17c38579f7ec5a91d8"><code>20f3253</code></a> track protocol benchmarks (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3339">#3339</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/921ad894f6f0ef36ffd9651ceaaa6cae330c74f1"><code>921ad89</code></a> fix event stream v2 deser (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3340">#3340</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0e4dafd630f15bc40204da3228b1821853e70e4"><code>f0e4daf</code></a> Release 2026-03-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f94e7b7b8d702c4fe9cf38ef769897682ee25748"><code>f94e7b7</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fb776cb55624781072f2c70e6d633b728a34d164"><code>fb776cb</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.96.2...service/s3/v1.96.4">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.4.0 to 1.5.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.5.0</h2>
<h2>What's Changed</h2>
<ul>
<li>RHCLOUD-43506 - Add CI worklows by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/24">project-kessel/kessel-sdk-go#24</a></li>
<li>Update generated protobuf files by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/25">project-kessel/kessel-sdk-go#25</a></li>
<li>Bump actions/checkout from 5 to 6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/23">project-kessel/kessel-sdk-go#23</a></li>
<li>Bump peter-evans/create-pull-request from 7 to 8 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/27">project-kessel/kessel-sdk-go#27</a></li>
<li>Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.10-20250912141014-52f32327d4b0.1 to 1.36.10-20251209175733-2a1774d88802.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/28">project-kessel/kessel-sdk-go#28</a></li>
<li>Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.10-20251209175733-2a1774d88802.1 to 1.36.11-20251209175733-2a1774d88802.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/30">project-kessel/kessel-sdk-go#30</a></li>
<li>Update generated protobuf files by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/33">project-kessel/kessel-sdk-go#33</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.4.0...v1.5.0">https://github.com/project-kessel/kessel-sdk-go/compare/v1.4.0...v1.5.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/6d043c9d52848d9afc84dee572281b9a8dbc0efc"><code>6d043c9</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/33">#33</a> from project-kessel/buf-generate-update</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/9728eb462e87adaff7ed5ae0f70c3b01bfaee2e6"><code>9728eb4</code></a> Regenerate protobuf files from buf</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/2f9267a404d67cec4789a5c79f6f1260dee65bfe"><code>2f9267a</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/30">#30</a> from project-kessel/dependabot/go_modules/buf.build/ge...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/dfe20b7e7cc9250421784ab3597f8a2f5b29de1d"><code>dfe20b7</code></a> Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/81c616a57b997fc7301dc320a3d12bdda349f5a1"><code>81c616a</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/28">#28</a> from project-kessel/dependabot/go_modules/buf.build/ge...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/7cedf01b5702f15ff0b19bab787ef8eefd2bcacb"><code>7cedf01</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/27">#27</a> from project-kessel/dependabot/github_actions/peter-ev...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/ab1faccde5ad8f8f4c5a8ba9f96e67fb58934d55"><code>ab1facc</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/23">#23</a> from project-kessel/dependabot/github_actions/actions/...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/cfcfe21c3eba3338e102550331d3475c95ae1e8c"><code>cfcfe21</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/25">#25</a> from project-kessel/buf-generate-update</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/25ac7ff33713cb6a8bb6d9a9ea98c82f2e9d72c1"><code>25ac7ff</code></a> Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/236496c04a103a5435cd105fc3df8c1f8ef09d77"><code>236496c</code></a> Bump peter-evans/create-pull-request from 7 to 8</li>
<li>Additional commits viewable in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.4.0...v1.5.0">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.79.1 to 1.79.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.79.2</h2>
<h1>Bug Fixes</h1>
<ul>
<li>stats: Prevent redundant error logging in health/ORCA producers by skipping stats/tracing processing when no stats handler is configured. (<a href="https://redirect.github.com/grpc/grpc-go/pull/8874">grpc/grpc-go#8874</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/8902ab6efea590f5b3861126559eaa26fa9783b2"><code>8902ab6</code></a> Change the version to release 1.79.2 (<a href="https://redirect.github.com/grpc/grpc-go/issues/8947">#8947</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/a9286705aa689bee321ec674323b6896284f3e02"><code>a928670</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/8874">#8874</a> to v1.79.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/8904">#8904</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/06df3638c0bcee88197b1033b3ba83e1eb8bc010"><code>06df363</code></a> Change version to 1.79.2-dev (<a href="https://redirect.github.com/grpc/grpc-go/issues/8903">#8903</a>)</li>
<li>See full diff in <a href="https://github.com/grpc/grpc-go/compare/v1.79.1...v1.79.2">compare view</a></li>
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

## Discussion

### Comment by @TenSt on March 09, 2026 at 10:13 AM UTC

@dependabot recreate

---

## Reviews

### Review by @TenSt - Approved on March 09, 2026 at 10:12 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1408*
