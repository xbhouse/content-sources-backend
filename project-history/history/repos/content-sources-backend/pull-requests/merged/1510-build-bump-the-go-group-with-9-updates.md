---
type: pull_request
number: 1510
title: "Build: Bump the go group with 9 updates"
state: merged
author: dependabot
created: 2026-06-01T10:49:30Z
updated: 2026-06-02T12:40:35Z
closed: 2026-06-02T12:40:33Z
merged: 2026-06-02T12:40:33Z
base_branch: main
head_branch: dependabot/go_modules/go-2ca7481d03
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1510
---

# Pull Request #1510: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: June 01, 2026 at 10:49 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-2ca7481d03`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.49.0` | `1.50.1` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.7` | `1.41.9` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.18` | `1.32.20` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.17` | `1.19.19` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.74.0` | `1.74.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.101.0` | `1.102.2` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.5.1778785514` | `2026.5.1779396769` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.9.0` | `1.9.1` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.19.0` | `9.20.0` |

Updates `github.com/IBM/sarama` from 1.49.0 to 1.50.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.50.1 (2026-05-27)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: correct requiredVersion for V8 JoinGroup and add protocol version placeholders by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3585">IBM/sarama#3585</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.50.0...v1.50.1">https://github.com/IBM/sarama/compare/v1.50.0...v1.50.1</a></p>
<h2>Version 1.50.0 (2026-05-27)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: add Java-compatible murmur2 partitioner by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3567">IBM/sarama#3567</a></li>
<li>feat(protocol): support LeaveGroupRequest/Response v5 (KIP-800) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3576">IBM/sarama#3576</a></li>
<li>feat(protocol): support JoinGroupRequest/Response v8 (KIP-800) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3577">IBM/sarama#3577</a></li>
<li>feat(consumer_group): set cancellation cause on session context by <a href="https://github.com/prakhar7651"><code>@​prakhar7651</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3575">IBM/sarama#3575</a></li>
<li>feat(consumer_group): send KIP-800 reason on JoinGroup and LeaveGroup by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3584">IBM/sarama#3584</a></li>
<li>feat: add Kafka 4.3.0 version placeholder by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3587">IBM/sarama#3587</a></li>
<li>feat(admin): support OffsetFetchRequest v8 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3565">IBM/sarama#3565</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(client): don't log ErrNoTopicsToUpdateMetadata every tick by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3566">IBM/sarama#3566</a></li>
<li>fix(broker): snapshot fetch meters before deferred Mark by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3563">IBM/sarama#3563</a></li>
<li>fix: prevent len out of range panic on 32bit architectures by <a href="https://github.com/gibmat"><code>@​gibmat</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3579">IBM/sarama#3579</a></li>
<li>fix(offset): retry fetchInitialOffset on top-level coordinator errors by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3574">IBM/sarama#3574</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update docker/bake-action action to v7.2.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3568">IBM/sarama#3568</a></li>
<li>fix(deps): update module golang.org/x/sys to v0.45.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3570">IBM/sarama#3570</a></li>
<li>chore(deps): update golangci/golangci-lint-action action to v9.2.1 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3571">IBM/sarama#3571</a></li>
<li>chore(deps): update module golang.org/x/crypto to v0.52.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3572">IBM/sarama#3572</a></li>
<li>fix(deps): update module golang.org/x/net to v0.55.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3573">IBM/sarama#3573</a></li>
<li>chore(deps): update docker/setup-buildx-action action to v4.1.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3578">IBM/sarama#3578</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>refactor: replace eapache/queue with generic ring buffer by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3560">IBM/sarama#3560</a></li>
<li>test(fvt): use a per-message timeout in follower failover test by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3562">IBM/sarama#3562</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/gibmat"><code>@​gibmat</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3579">IBM/sarama#3579</a></li>
<li><a href="https://github.com/prakhar7651"><code>@​prakhar7651</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3575">IBM/sarama#3575</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.49.0...v1.50.0">https://github.com/IBM/sarama/compare/v1.49.0...v1.50.0</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
<h2>Version 1.50.0 (2026-05-27)</h2>
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: add Java-compatible murmur2 partitioner by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3567">IBM/sarama#3567</a></li>
<li>feat(protocol): support LeaveGroupRequest/Response v5 (KIP-800) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3576">IBM/sarama#3576</a></li>
<li>feat(protocol): support JoinGroupRequest/Response v8 (KIP-800) by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3577">IBM/sarama#3577</a></li>
<li>feat(consumer_group): set cancellation cause on session context by <a href="https://github.com/prakhar7651"><code>@​prakhar7651</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3575">IBM/sarama#3575</a></li>
<li>feat(consumer_group): send KIP-800 reason on JoinGroup and LeaveGroup by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3584">IBM/sarama#3584</a></li>
<li>feat: add Kafka 4.3.0 version placeholder by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3587">IBM/sarama#3587</a></li>
<li>feat(admin): support OffsetFetchRequest v8 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3565">IBM/sarama#3565</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(client): don't log ErrNoTopicsToUpdateMetadata every tick by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3566">IBM/sarama#3566</a></li>
<li>fix(broker): snapshot fetch meters before deferred Mark by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3563">IBM/sarama#3563</a></li>
<li>fix: prevent len out of range panic on 32bit architectures by <a href="https://github.com/gibmat"><code>@​gibmat</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3579">IBM/sarama#3579</a></li>
<li>fix(offset): retry fetchInitialOffset on top-level coordinator errors by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3574">IBM/sarama#3574</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update docker/bake-action action to v7.2.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3568">IBM/sarama#3568</a></li>
<li>fix(deps): update module golang.org/x/sys to v0.45.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3570">IBM/sarama#3570</a></li>
<li>chore(deps): update golangci/golangci-lint-action action to v9.2.1 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3571">IBM/sarama#3571</a></li>
<li>chore(deps): update module golang.org/x/crypto to v0.52.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3572">IBM/sarama#3572</a></li>
<li>fix(deps): update module golang.org/x/net to v0.55.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3573">IBM/sarama#3573</a></li>
<li>chore(deps): update docker/setup-buildx-action action to v4.1.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3578">IBM/sarama#3578</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>refactor: replace eapache/queue with generic ring buffer by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3560">IBM/sarama#3560</a></li>
<li>test(fvt): use a per-message timeout in follower failover test by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3562">IBM/sarama#3562</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/gibmat"><code>@​gibmat</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3579">IBM/sarama#3579</a></li>
<li><a href="https://github.com/prakhar7651"><code>@​prakhar7651</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3575">IBM/sarama#3575</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.49.0...v1.50.0">https://github.com/IBM/sarama/compare/v1.49.0...v1.50.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/bb409c06a0779b2de72360804d9d6c4a8c6b78f7"><code>bb409c0</code></a> chore(doc): update CHANGELOG.md with releases v1.43.0 through v1.50.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3589">#3589</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e8c917d90b9d87fb90b2220a8908d4cfb8d5e396"><code>e8c917d</code></a> fix: correct requiredVersion for V8 JoinGroup and add protocol version placeh...</li>
<li><a href="https://github.com/IBM/sarama/commit/33f660245fb7edca16ecdb414aa5298ef27aa1cf"><code>33f6602</code></a> feat(admin): support OffsetFetchRequest v8 (<a href="https://redirect.github.com/IBM/sarama/issues/3565">#3565</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/2fa22aadbb94f60b3b0e6bfe5c1fb12a36f7c166"><code>2fa22aa</code></a> chore(fvt): bump default docker Kafka version to 4.2.0</li>
<li><a href="https://github.com/IBM/sarama/commit/8577cce86019bebf4f325d6b114571d46ef4df63"><code>8577cce</code></a> chore(ci): add Kafka 4.2.0 to FVT matrix and refine version spread</li>
<li><a href="https://github.com/IBM/sarama/commit/be9392377d17d8c7475eb5811b53315fdcf4d3a6"><code>be93923</code></a> feat: add Kafka 4.3.0 version placeholder</li>
<li><a href="https://github.com/IBM/sarama/commit/822e6d40c7652cb541276b9b4a40c8fabe2b7e07"><code>822e6d4</code></a> fix(offset): retry fetchInitialOffset on top-level coordinator errors (<a href="https://redirect.github.com/IBM/sarama/issues/3574">#3574</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/c515627455abe6fc15cce517aef3b771a079eaa7"><code>c515627</code></a> feat(consumer_group): send KIP-800 reason on JoinGroup and LeaveGroup (<a href="https://redirect.github.com/IBM/sarama/issues/3584">#3584</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/92428ce137a3165eb5fbe9a2f16ab813ede0a297"><code>92428ce</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.8 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/f5cd855f9f062e977450e30aa33d99f0a22bd244"><code>f5cd855</code></a> feat(consumer_group): set cancellation cause on session context (<a href="https://redirect.github.com/IBM/sarama/issues/3575">#3575</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.49.0...v1.50.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.7 to 1.41.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5841d3ae2cfd6e6113ca61b71d69131b84932f4c"><code>5841d3a</code></a> Release 2026-05-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/16ac80d79c282366cae312281b3df925af4e9bf1"><code>16ac80d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/167926f8aca7228f2dd1bed73707505875aafef4"><code>167926f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0fce13e18c6bff397ad77fac4cde4ab3f3b93e0"><code>a0fce13</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223c0057211950899e0117dc027cc299a1dac664"><code>223c005</code></a> update to smithy-go v1.26.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3426">#3426</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74c501189a40c9b937432a1b2a4cacffc851ea76"><code>74c5011</code></a> Release 2026-05-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d82651329a86064a9026f6219cff72921fa74da"><code>7d82651</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79c63d9289784de4914143b7bff67157aa6a2a90"><code>79c63d9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b15b3b873ad5c294d0c010fb1cc56ecb583d1618"><code>b15b3b8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/090e46936630944917cfd6a0990ea3fd6391475b"><code>090e469</code></a> Feat tmv2 parity (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3424">#3424</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.7...v1.41.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.18 to 1.32.20
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5841d3ae2cfd6e6113ca61b71d69131b84932f4c"><code>5841d3a</code></a> Release 2026-05-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/16ac80d79c282366cae312281b3df925af4e9bf1"><code>16ac80d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/167926f8aca7228f2dd1bed73707505875aafef4"><code>167926f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0fce13e18c6bff397ad77fac4cde4ab3f3b93e0"><code>a0fce13</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223c0057211950899e0117dc027cc299a1dac664"><code>223c005</code></a> update to smithy-go v1.26.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3426">#3426</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74c501189a40c9b937432a1b2a4cacffc851ea76"><code>74c5011</code></a> Release 2026-05-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d82651329a86064a9026f6219cff72921fa74da"><code>7d82651</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79c63d9289784de4914143b7bff67157aa6a2a90"><code>79c63d9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b15b3b873ad5c294d0c010fb1cc56ecb583d1618"><code>b15b3b8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/090e46936630944917cfd6a0990ea3fd6391475b"><code>090e469</code></a> Feat tmv2 parity (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3424">#3424</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.18...config/v1.32.20">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.17 to 1.19.19
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5841d3ae2cfd6e6113ca61b71d69131b84932f4c"><code>5841d3a</code></a> Release 2026-05-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/16ac80d79c282366cae312281b3df925af4e9bf1"><code>16ac80d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/167926f8aca7228f2dd1bed73707505875aafef4"><code>167926f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0fce13e18c6bff397ad77fac4cde4ab3f3b93e0"><code>a0fce13</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223c0057211950899e0117dc027cc299a1dac664"><code>223c005</code></a> update to smithy-go v1.26.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3426">#3426</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74c501189a40c9b937432a1b2a4cacffc851ea76"><code>74c5011</code></a> Release 2026-05-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d82651329a86064a9026f6219cff72921fa74da"><code>7d82651</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79c63d9289784de4914143b7bff67157aa6a2a90"><code>79c63d9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b15b3b873ad5c294d0c010fb1cc56ecb583d1618"><code>b15b3b8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/090e46936630944917cfd6a0990ea3fd6391475b"><code>090e469</code></a> Feat tmv2 parity (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3424">#3424</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.17...credentials/v1.19.19">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.74.0 to 1.74.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67db6904b816b95073883b7ad378384c4839b28c"><code>67db690</code></a> Release 2025-09-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32ee1b5d75fc303c0626a6f5e769f4e08cc491a8"><code>32ee1b5</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b431223309a815cffc048072556aa651ee1455f"><code>0b43122</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44786d920f3627b73a99e81c7b6399dbfcf7ab42"><code>44786d9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c98edb73809256823906d7e307ecf3c9abc16700"><code>c98edb7</code></a> update internal endpts comment that was wrong (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3194">#3194</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/88da3c8c5569dece0e99802dab638faa047a0db0"><code>88da3c8</code></a> Release 2025-09-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74a74fc179f8bbd879383cc75fa29a1937266dcc"><code>74a74fc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e6f7ae6139ca69044bb706664b4dbdc31227a32"><code>5e6f7ae</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e722ab42ff6bc6bb810c2937b8e1b41937e17c3"><code>0e722ab</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/41a7d004b9ff794f6007d30168afc825031f2c61"><code>41a7d00</code></a> Release 2025-09-24</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.74.0...service/eks/v1.74.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.101.0 to 1.102.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5841d3ae2cfd6e6113ca61b71d69131b84932f4c"><code>5841d3a</code></a> Release 2026-05-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/16ac80d79c282366cae312281b3df925af4e9bf1"><code>16ac80d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/167926f8aca7228f2dd1bed73707505875aafef4"><code>167926f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0fce13e18c6bff397ad77fac4cde4ab3f3b93e0"><code>a0fce13</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/223c0057211950899e0117dc027cc299a1dac664"><code>223c005</code></a> update to smithy-go v1.26.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3426">#3426</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74c501189a40c9b937432a1b2a4cacffc851ea76"><code>74c5011</code></a> Release 2026-05-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d82651329a86064a9026f6219cff72921fa74da"><code>7d82651</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79c63d9289784de4914143b7bff67157aa6a2a90"><code>79c63d9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b15b3b873ad5c294d0c010fb1cc56ecb583d1618"><code>b15b3b8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/090e46936630944917cfd6a0990ea3fd6391475b"><code>090e469</code></a> Feat tmv2 parity (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3424">#3424</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.101.0...service/s3/v1.102.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.5.1778785514 to 2026.5.1779396769
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/5b4d6c7456438610394191bd6a3787155a839acb"><code>5b4d6c7</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47542e49b395fb7a98d596b8e3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.5.1778785514...release/v2026.5.1779396769">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.9.0 to 1.9.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.9.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.11-20260209202127-80ab13bee0bf.1 to 1.36.11-20260415201107-50325440f8f2.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/52">project-kessel/kessel-sdk-go#52</a></li>
<li>Bump github.com/zitadel/oidc/v3 from 3.45.5 to 3.47.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/56">project-kessel/kessel-sdk-go#56</a></li>
<li>Bump google.golang.org/grpc from 1.79.2 to 1.81.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/60">project-kessel/kessel-sdk-go#60</a></li>
<li>RHCLOUD-47739 - Update locking mechanism for concurrent token refreshes by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/61">project-kessel/kessel-sdk-go#61</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.0...v1.9.1">https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.0...v1.9.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/755a5e529f6e1f10c6d6ab0ae9aef3972731aaa6"><code>755a5e5</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/61">#61</a> from lennysgarage/RHCLOUD-47739</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/d56d80d02d5a109b39c8df1ede5024fe746d0ccb"><code>d56d80d</code></a> Update locking mechanism for concurrent token refreshes</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/0499176cf16b8a6c70a0c19ab03371af4187f896"><code>0499176</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/60">#60</a> from project-kessel/dependabot/go_modules/google.golan...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/f5097bd9bb7094e841e8640fcdcde369f360838d"><code>f5097bd</code></a> Bump google.golang.org/grpc from 1.79.2 to 1.81.1</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/ea9736461d235b700e5d6aced68e8c6eb1914a87"><code>ea97364</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/56">#56</a> from project-kessel/dependabot/go_modules/github.com/z...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/e59a535a44cc6c8b68d3f7216c3337ccc1c2f451"><code>e59a535</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/52">#52</a> from project-kessel/dependabot/go_modules/buf.build/ge...</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/eb64335dab89f9a7682e9ee7319764abff4a5047"><code>eb64335</code></a> Bump github.com/zitadel/oidc/v3 from 3.45.5 to 3.47.5</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/203f83eede9eea12cfbb5916cfa20f3cc771d49d"><code>203f83e</code></a> Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go</li>
<li>See full diff in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.9.0...v1.9.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.19.0 to 9.20.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.20.0</h2>
<h2>🚀 Highlights</h2>
<h3>Redis 8.8 Support</h3>
<p>This release adds support for <strong>Redis 8.8</strong>. The README's supported-versions list now includes Redis 8.8 alongside 8.0/8.2/8.4, and CI exercises the <code>8.8</code> client-libs-test image across the full suite (Makefile, build workflow, doctests, run-tests action, and docker-compose).</p>
<p>Coverage for the new commands that ship in the 8.x line, rounded out in this release:</p>
<ul>
<li><strong><code>AR*</code> array data type</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3813">#3813</a>) — new array data structure, exposed via the <code>ArrayCmdable</code> interface (see the experimental-features highlight below).</li>
<li><strong><code>INCREX</code></strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3816">#3816</a>) — atomic increment with expiration in a single round-trip.</li>
<li><strong><code>XNACK</code></strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3790">#3790</a>) — explicit negative-acknowledge of pending stream entries.</li>
<li><strong><code>XAUTOCLAIM</code> PEL deletes</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3798">#3798</a>) — <code>XAUTOCLAIM</code>/<code>XAUTOCLAIMJUSTID</code> now return the list of deleted message IDs from the pending entries list.</li>
<li><strong><code>TS.RANGE</code> multiple aggregators</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3791">#3791</a>) — <code>TS.RANGE</code>/<code>TS.REVRANGE</code>/<code>TS.MRANGE</code>/<code>TS.MREVRANGE</code> accept multiple aggregators in a single call.</li>
<li><strong><code>Z(UNION|INTER|DIFF)</code> <code>COUNT</code> aggregator</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3802">#3802</a>) — <code>COUNT</code> reducer for sorted-set set operations.</li>
<li><strong><code>JSON.SET FPHA</code></strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3797">#3797</a>) — new <code>FPHA</code> argument that specifies the floating-point type for homogeneous FP arrays.</li>
</ul>
<p>CI image bump (<a href="https://redirect.github.com/redis/go-redis/pull/3814">#3814</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>. Command coverage contributions by <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/Khukharr"><code>@​Khukharr</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, and <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>.</p>
<h3>Stable RESP3 for RediSearch (<code>UnstableResp3</code> deprecated)</h3>
<p><code>FT.SEARCH</code>, <code>FT.AGGREGATE</code>, <code>FT.INFO</code>, <code>FT.SPELLCHECK</code>, and <code>FT.SYNDUMP</code> now parse RESP3 (map) responses into the same typed result objects as RESP2 — <code>Val()</code> and <code>Result()</code> work uniformly on both protocols, no flag required. Previously, RESP3 search responses required <code>UnstableResp3: true</code> and were returned as opaque maps accessible only via <code>RawResult()</code> / <code>RawVal()</code>.</p>
<p>As a result, the <code>UnstableResp3</code> option is now a <strong>no-op</strong> across every options struct (<code>Options</code>, <code>ClusterOptions</code>, <code>UniversalOptions</code>, <code>FailoverOptions</code>, <code>RingOptions</code>) and has been marked <code>// Deprecated:</code>. The field is retained for backwards compatibility — existing code that sets <code>UnstableResp3: true</code> will continue to compile and behave identically — but it will be removed in a future release and new code should not set it. <code>RawResult()</code> / <code>RawVal()</code> continue to work for callers that prefer the raw RESP payload.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3741">#3741</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h3>Experimental Array Data Structure Commands</h3>
<p>Adds an experimental <code>ArrayCmdable</code> interface with the <code>AR*</code> command family (<code>ARSet</code>, <code>ARGet</code>, <code>ARGetRange</code>, <code>ARMSet</code>, <code>ARMGet</code>, <code>ARDel</code>, <code>ARDelRange</code>, <code>ARScan</code>, <code>ARSeek</code>, <code>ARNext</code>, <code>ARLastItems</code>, <code>ARGrep</code>, <code>ARGrepWithValues</code>, <code>ARInfo</code>/<code>ARInfoFull</code>, and typed reducers <code>AROpSum</code>/<code>AROpMin</code>/<code>AROpMax</code>/<code>AROpAnd</code>/<code>AROpOr</code>/<code>AROpXor</code>/<code>AROpMatch</code>/<code>AROpUsed</code>) for working with Redis 8.8's new array data type. <strong>API is experimental and may change in a future release.</strong></p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3813">#3813</a>) by <a href="https://github.com/cxljs"><code>@​cxljs</code></a></p>
<h2>✨ New Features</h2>
<ul>
<li><strong>RESP3 search parser</strong>: First-class RESP3 parsing for <code>FT.SEARCH</code>/<code>FT.AGGREGATE</code>/<code>FT.INFO</code>/<code>FT.SPELLCHECK</code>/<code>FT.SYNDUMP</code> responses with backwards compatibility for RESP2 (<a href="https://redirect.github.com/redis/go-redis/pull/3741">#3741</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>INCREX</strong>: New <code>INCREX</code> command support — atomic increment with expiration (<a href="https://redirect.github.com/redis/go-redis/pull/3816">#3816</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>XNACK</strong>: Client support for the <code>XNACK</code> stream command for explicitly negative-acknowledging pending entries (<a href="https://redirect.github.com/redis/go-redis/pull/3790">#3790</a>) by <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a></li>
<li><strong>TS range multiple aggregators</strong>: <code>TS.RANGE</code>/<code>TS.REVRANGE</code>/<code>TS.MRANGE</code>/<code>TS.MREVRANGE</code> now accept multiple aggregators in a single call (<a href="https://redirect.github.com/redis/go-redis/pull/3791">#3791</a>) by <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a></li>
<li><strong><code>XAutoClaim</code> deleted IDs</strong>: <code>XAUTOCLAIM</code>/<code>XAUTOCLAIMJUSTID</code> now return the list of deleted message IDs from the PEL (<a href="https://redirect.github.com/redis/go-redis/pull/3798">#3798</a>) by <a href="https://github.com/Khukharr"><code>@​Khukharr</code></a></li>
<li><strong><code>JSON.SET FPHA</code></strong>: <code>JSON.SET</code> accepts a new <code>FPHA</code> argument that specifies the floating-point type for homogeneous floating-point arrays (<a href="https://redirect.github.com/redis/go-redis/pull/3797">#3797</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Sorted-set union/intersection COUNT</strong>: <code>ZUNION</code>/<code>ZINTER</code>/<code>ZDIFF</code> aggregator now supports <code>COUNT</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3802">#3802</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong><code>FT.HYBRID</code> vector validation</strong>: Validates hybrid-search vector input types and adds proper typed vector parameters (<a href="https://redirect.github.com/redis/go-redis/pull/3756">#3756</a>) by <a href="https://github.com/DengY11"><code>@​DengY11</code></a></li>
<li><strong>Cluster pool wait stats</strong>: <code>ClusterClient.PoolStats()</code> now accumulates <code>WaitCount</code> and <code>WaitDurationNs</code> across all node pools (previously always zero) (<a href="https://redirect.github.com/redis/go-redis/pull/3809">#3809</a>) by <a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a></li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>TLS-only Cluster PubSub</strong>: <code>CLUSTER SLOTS</code> port-0 entries now fall back to the origin endpoint's port, fixing <code>dial tcp &lt;ip&gt;:0: connection refused</code> on TLS-only clusters started with <code>--port 0 --tls-port &lt;port&gt;</code> (fixes <a href="https://redirect.github.com/redis/go-redis/issues/3726">#3726</a>) (<a href="https://redirect.github.com/redis/go-redis/pull/3828">#3828</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Sharded PubSub reconnect routing</strong>: <code>PubSub.conn()</code> now passes both regular (<code>c.channels</code>) and sharded (<code>c.schannels</code>) channels into the per-PubSub <code>newConn</code> closure. Previously, <code>ClusterClient.SSubscribe</code>-only PubSubs reconnected to a random node (because the routing closure saw an empty channel list), the <code>SSUBSCRIBE</code> was sent to the wrong shard, and the resulting <code>MOVED</code> reply was silently dropped (<a href="https://redirect.github.com/redis/go-redis/pull/3829">#3829</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>ClusterClient <code>Watch</code> retry</strong>: User errors returned from a <code>Watch</code> callback are no longer subjected to cluster-retry classification; transient cluster errors still retry, but a callback returning e.g. <code>net.ErrClosed</code> short-circuits immediately (<a href="https://redirect.github.com/redis/go-redis/pull/3821">#3821</a>) by <a href="https://github.com/obiyang"><code>@​obiyang</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.20.0 (2026-05-28)</h1>
<h2>🚀 Highlights</h2>
<h3>Redis 8.8 Support</h3>
<p>This release adds support for <strong>Redis 8.8</strong>. The README's supported-versions list now includes Redis 8.8 alongside 8.0/8.2/8.4, and CI exercises the <code>8.8-rc1</code> client-libs-test image across the full suite (Makefile, build workflow, doctests, run-tests action, and docker-compose).</p>
<p>Coverage for the new commands that ship in the 8.x line, rounded out in this release:</p>
<ul>
<li><strong><code>AR*</code> array data type</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3813">#3813</a>) — new array data structure, exposed via the <code>ArrayCmdable</code> interface (see the experimental-features highlight below).</li>
<li><strong><code>INCREX</code></strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3816">#3816</a>) — atomic increment with expiration in a single round-trip.</li>
<li><strong><code>XNACK</code></strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3790">#3790</a>) — explicit negative-acknowledge of pending stream entries.</li>
<li><strong><code>XAUTOCLAIM</code> PEL deletes</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3798">#3798</a>) — <code>XAUTOCLAIM</code>/<code>XAUTOCLAIMJUSTID</code> now return the list of deleted message IDs from the pending entries list.</li>
<li><strong><code>TS.RANGE</code> multiple aggregators</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3791">#3791</a>) — <code>TS.RANGE</code>/<code>TS.REVRANGE</code>/<code>TS.MRANGE</code>/<code>TS.MREVRANGE</code> accept multiple aggregators in a single call.</li>
<li><strong><code>Z(UNION|INTER|DIFF)</code> <code>COUNT</code> aggregator</strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3802">#3802</a>) — <code>COUNT</code> reducer for sorted-set set operations.</li>
<li><strong><code>JSON.SET FPHA</code></strong> (<a href="https://redirect.github.com/redis/go-redis/pull/3797">#3797</a>) — new <code>FPHA</code> argument that specifies the floating-point type for homogeneous FP arrays.</li>
</ul>
<p>CI image bump (<a href="https://redirect.github.com/redis/go-redis/pull/3814">#3814</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>. Command coverage contributions by <a href="https://github.com/cxljs"><code>@​cxljs</code></a>, <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a>, <a href="https://github.com/Khukharr"><code>@​Khukharr</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a>, and <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>.</p>
<h3>Stable RESP3 for RediSearch (<code>UnstableResp3</code> deprecated)</h3>
<p><code>FT.SEARCH</code>, <code>FT.AGGREGATE</code>, <code>FT.INFO</code>, <code>FT.SPELLCHECK</code>, and <code>FT.SYNDUMP</code> now parse RESP3 (map) responses into the same typed result objects as RESP2 — <code>Val()</code> and <code>Result()</code> work uniformly on both protocols, no flag required. Previously, RESP3 search responses required <code>UnstableResp3: true</code> and were returned as opaque maps accessible only via <code>RawResult()</code> / <code>RawVal()</code>.</p>
<p>As a result, the <code>UnstableResp3</code> option is now a <strong>no-op</strong> across every options struct (<code>Options</code>, <code>ClusterOptions</code>, <code>UniversalOptions</code>, <code>FailoverOptions</code>, <code>RingOptions</code>) and has been marked <code>// Deprecated:</code>. The field is retained for backwards compatibility — existing code that sets <code>UnstableResp3: true</code> will continue to compile and behave identically — but it will be removed in a future release and new code should not set it. <code>RawResult()</code> / <code>RawVal()</code> continue to work for callers that prefer the raw RESP payload.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3741">#3741</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h3>Experimental Array Data Structure Commands</h3>
<p>Adds an experimental <code>ArrayCmdable</code> interface with the <code>AR*</code> command family (<code>ARSet</code>, <code>ARGet</code>, <code>ARGetRange</code>, <code>ARMSet</code>, <code>ARMGet</code>, <code>ARDel</code>, <code>ARDelRange</code>, <code>ARScan</code>, <code>ARSeek</code>, <code>ARNext</code>, <code>ARLastItems</code>, <code>ARGrep</code>, <code>ARGrepWithValues</code>, <code>ARInfo</code>/<code>ARInfoFull</code>, and typed reducers <code>AROpSum</code>/<code>AROpMin</code>/<code>AROpMax</code>/<code>AROpAnd</code>/<code>AROpOr</code>/<code>AROpXor</code>/<code>AROpMatch</code>/<code>AROpUsed</code>) for working with Redis 8.8's new array data type. <strong>API is experimental and may change in a future release.</strong></p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3813">#3813</a>) by <a href="https://github.com/cxljs"><code>@​cxljs</code></a></p>
<h2>✨ New Features</h2>
<ul>
<li><strong>RESP3 search parser</strong>: First-class RESP3 parsing for <code>FT.SEARCH</code>/<code>FT.AGGREGATE</code>/<code>FT.INFO</code>/<code>FT.SPELLCHECK</code>/<code>FT.SYNDUMP</code> responses with backwards compatibility for RESP2 (<a href="https://redirect.github.com/redis/go-redis/pull/3741">#3741</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>INCREX</strong>: New <code>INCREX</code> command support — atomic increment with expiration (<a href="https://redirect.github.com/redis/go-redis/pull/3816">#3816</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>XNACK</strong>: Client support for the <code>XNACK</code> stream command for explicitly negative-acknowledging pending entries (<a href="https://redirect.github.com/redis/go-redis/pull/3790">#3790</a>) by <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a></li>
<li><strong>TS range multiple aggregators</strong>: <code>TS.RANGE</code>/<code>TS.REVRANGE</code>/<code>TS.MRANGE</code>/<code>TS.MREVRANGE</code> now accept multiple aggregators in a single call (<a href="https://redirect.github.com/redis/go-redis/pull/3791">#3791</a>) by <a href="https://github.com/elena-kolevska"><code>@​elena-kolevska</code></a></li>
<li><strong><code>XAutoClaim</code> deleted IDs</strong>: <code>XAUTOCLAIM</code>/<code>XAUTOCLAIMJUSTID</code> now return the list of deleted message IDs from the PEL (<a href="https://redirect.github.com/redis/go-redis/pull/3798">#3798</a>) by <a href="https://github.com/Khukharr"><code>@​Khukharr</code></a></li>
<li><strong><code>JSON.SET FPHA</code></strong>: <code>JSON.SET</code> accepts a new <code>FPHA</code> argument that specifies the floating-point type for homogeneous floating-point arrays (<a href="https://redirect.github.com/redis/go-redis/pull/3797">#3797</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Sorted-set union/intersection COUNT</strong>: <code>ZUNION</code>/<code>ZINTER</code>/<code>ZDIFF</code> aggregator now supports <code>COUNT</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3802">#3802</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong><code>FT.HYBRID</code> vector validation</strong>: Validates hybrid-search vector input types and adds proper typed vector parameters (<a href="https://redirect.github.com/redis/go-redis/pull/3756">#3756</a>) by <a href="https://github.com/DengY11"><code>@​DengY11</code></a></li>
<li><strong>Cluster pool wait stats</strong>: <code>ClusterClient.PoolStats()</code> now accumulates <code>WaitCount</code> and <code>WaitDurationNs</code> across all node pools (previously always zero) (<a href="https://redirect.github.com/redis/go-redis/pull/3809">#3809</a>) by <a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a></li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>TLS-only Cluster PubSub</strong>: <code>CLUSTER SLOTS</code> port-0 entries now fall back to the origin endpoint's port, fixing <code>dial tcp &lt;ip&gt;:0: connection refused</code> on TLS-only clusters started with <code>--port 0 --tls-port &lt;port&gt;</code> (fixes <a href="https://redirect.github.com/redis/go-redis/issues/3726">#3726</a>) (<a href="https://redirect.github.com/redis/go-redis/pull/3828">#3828</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Sharded PubSub reconnect routing</strong>: <code>PubSub.conn()</code> now passes both regular (<code>c.channels</code>) and sharded (<code>c.schannels</code>) channels into the per-PubSub <code>newConn</code> closure. Previously, <code>ClusterClient.SSubscribe</code>-only PubSubs reconnected to a random node (because the routing closure saw an empty channel list), the <code>SSUBSCRIBE</code> was sent to the wrong shard, and the resulting <code>MOVED</code> reply was silently dropped (<a href="https://redirect.github.com/redis/go-redis/pull/3829">#3829</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/7d05dd3b7ce12a7b8c7923f73da0fede3bfa7c03"><code>7d05dd3</code></a> chore(release): v9.20.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3832">#3832</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/97568827124c58f1338d78b9535edc0eb9435453"><code>9756882</code></a> fix(test): make waitForSentinelClusterStable robust to disconnected r… (<a href="https://redirect.github.com/redis/go-redis/issues/3830">#3830</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/875ce21e4862a7913edb5308b16869141ce5e833"><code>875ce21</code></a> fix(sentinel): do not close sentinel when replica list is empty (<a href="https://redirect.github.com/redis/go-redis/issues/3795">#3795</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8a027f21ba2686511e1a6a159f0c6888f11d48e7"><code>8a027f2</code></a> chore(ci): add govulncheck workflow (<a href="https://redirect.github.com/redis/go-redis/issues/3779">#3779</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/d8407df30af7eba2c24280791fed4b8bd3e62c97"><code>d8407df</code></a> fix(pubsub): include shard channels in newConn routing list (<a href="https://redirect.github.com/redis/go-redis/issues/3829">#3829</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6af9bdc84ba25303f2b9dd9949055ee6224c6d96"><code>6af9bdc</code></a> fix(cluster): fall back to origin port when CLUSTER SLOTS reports port 0 (<a href="https://redirect.github.com/redis/go-redis/issues/3828">#3828</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fa5aa8c4d26a0983f16ebc537696150d80bfc8f0"><code>fa5aa8c</code></a> chore(doc): Update README and CI image. (<a href="https://redirect.github.com/redis/go-redis/issues/3822">#3822</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fdcc6f9221e6d447b1c1a50d946c74d8db51cda3"><code>fdcc6f9</code></a> refactor(keyPos): Enhance key position retrieval with CommandInfo caching (<a href="https://redirect.github.com/redis/go-redis/issues/3">#3</a>...</li>
<li><a href="https://github.com/redis/go-redis/commit/68a8bc1e8457d796d35cae34435529cd6e56a762"><code>68a8bc1</code></a> fix(sentinel): close non-winning sentinel clients in MasterAddr concurrent pr...</li>
<li><a href="https://github.com/redis/go-redis/commit/00bf6d3edc1f50cce7a81d6a11a580cb060fce56"><code>00bf6d3</code></a> fix: avoid retrying ClusterClient Watch callback errors (<a href="https://redirect.github.com/redis/go-redis/issues/3821">#3821</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.19.0...v9.20.0">compare view</a></li>
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

### Comment by @rverdile on June 01, 2026 at 05:18 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on June 02, 2026 at 12:40 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1510*
