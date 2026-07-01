---
type: pull_request
number: 1504
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2026-05-25T05:50:36Z
updated: 2026-05-26T08:28:16Z
closed: 2026-05-26T08:28:13Z
merged: 2026-05-26T08:28:13Z
base_branch: main
head_branch: dependabot/go_modules/go-e3fe47c5a6
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1504
---

# Pull Request #1504: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: May 25, 2026 at 05:50 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-e3fe47c5a6`

## Description

Bumps the go group with 4 updates: [github.com/IBM/sarama](https://github.com/IBM/sarama), [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) and [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest).

Updates `github.com/IBM/sarama` from 1.48.2 to 1.49.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.49.0 (2026-05-18)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:rotating_light: Breaking Changes</h3>
<ul>
<li>fix(consumer): decouple FetchRequest.MaxBytes from MaxResponseSize by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3538">IBM/sarama#3538</a></li>
</ul>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(consumer): warn on sustained partition retries by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3535">IBM/sarama#3535</a></li>
<li>feat(producer): add Produce v8 request/response support by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3540">IBM/sarama#3540</a></li>
<li>feat(consumer): cap partition consumer retries by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3539">IBM/sarama#3539</a></li>
<li>feat: support FindCoordinator V3 protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3544">IBM/sarama#3544</a></li>
<li>feat: support describe acls v2 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3548">IBM/sarama#3548</a></li>
<li>feat: support create acls v2 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3549">IBM/sarama#3549</a></li>
<li>feat: support delete acls v2 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3550">IBM/sarama#3550</a></li>
<li>feat: support sasl authenticate v2 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3551">IBM/sarama#3551</a></li>
<li>feat: support create partitions v2 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3554">IBM/sarama#3554</a></li>
<li>feat: support join group v7 by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3555">IBM/sarama#3555</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: flexible decoder out-of-bounds panic by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3543">IBM/sarama#3543</a></li>
<li>fix(consumer): size partial-batch retry correctly by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3541">IBM/sarama#3541</a></li>
<li>feat(consumer): add OffsetCommit v8 request/response support by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3545">IBM/sarama#3545</a></li>
<li>fix: decode nullable ACL describe error messages by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3552">IBM/sarama#3552</a></li>
<li>fix(consumer): lease preferred read replicas by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3553">IBM/sarama#3553</a></li>
<li>fix(producer): honour Retry.Backoff in idempotent retryBatch by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3557">IBM/sarama#3557</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: better bounds checking by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3546">IBM/sarama#3546</a></li>
<li>chore: bump deps in ./examples tree by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3558">IBM/sarama#3558</a></li>
<li>docs: add AlterPartitionReassignments example and functional test by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3556">IBM/sarama#3556</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.48.2...v1.49.0">https://github.com/IBM/sarama/compare/v1.48.2...v1.49.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/63f561ad8a227a1e9adf5ea35882aa7135deb360"><code>63f561a</code></a> chore(ci): Update github/codeql-action action to v4.35.5 (<a href="https://redirect.github.com/IBM/sarama/issues/3559">#3559</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4009ac0b864233e22207c978a8567b70cefe4b4d"><code>4009ac0</code></a> fix(producer): honour Retry.Backoff in idempotent retryBatch (<a href="https://redirect.github.com/IBM/sarama/issues/3557">#3557</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/65e434dc7e72ef3d1ba7d2765e0ea236fcd5f947"><code>65e434d</code></a> docs: add AlterPartitionReassignments example and functional test (<a href="https://redirect.github.com/IBM/sarama/issues/3556">#3556</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/5cf9fb80ee2a57817a2adfc9cb14e615b0295e18"><code>5cf9fb8</code></a> fix(consumer): lease preferred read replicas (<a href="https://redirect.github.com/IBM/sarama/issues/3553">#3553</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/833fe3dde707b17ee0ea7b4eda137fe041c68699"><code>833fe3d</code></a> chore: bump deps in ./examples tree (<a href="https://redirect.github.com/IBM/sarama/issues/3558">#3558</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4d592d0f8424d039a28a08f7739952d4f0bbeb2e"><code>4d592d0</code></a> feat: support join group v7 (<a href="https://redirect.github.com/IBM/sarama/issues/3555">#3555</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/bcee29d75752fb966e438da7ff0adf4dea770233"><code>bcee29d</code></a> fix: decode nullable ACL describe error messages (<a href="https://redirect.github.com/IBM/sarama/issues/3552">#3552</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/fb8ed8ddd1ff3e26cff2ddb12dd1e036d0860b50"><code>fb8ed8d</code></a> feat: support create partitions v2 (<a href="https://redirect.github.com/IBM/sarama/issues/3554">#3554</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/d23b63f8803e70ee0239449ce98be7a841834e11"><code>d23b63f</code></a> feat: support sasl authenticate v2 (<a href="https://redirect.github.com/IBM/sarama/issues/3551">#3551</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a31a726b76e7ce6a4ecfcf079bb24f4cb2471f41"><code>a31a726</code></a> feat: support delete acls v2 (<a href="https://redirect.github.com/IBM/sarama/issues/3550">#3550</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.48.2...v1.49.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.17 to 1.32.18
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/db9f4e546dfe2f62a6bc3bf54b9da42ebace6372"><code>db9f4e5</code></a> Release 2026-05-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34e7ddc9400e830a9ae226a7e3c2161e5ece4f19"><code>34e7ddc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f9db036cf7b3b8a1ea5eb67c3d296da4b48b6e2b"><code>f9db036</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae5eae1e3ec46433bd99496bfa6936f8f09a2e72"><code>ae5eae1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/429dbdd2a35d325aabc5757edfc9ebf09c2ad12e"><code>429dbdd</code></a> Feat discover endpoint partition validation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3410">#3410</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ab4f5b60785064ec6346c922604d94b63d9c7299"><code>ab4f5b6</code></a> Release 2026-05-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/757a09909a97a15e5a481d9839b83f15b8fdc4bc"><code>757a099</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/02c8323ee6c99be82dae3a3923616756cb164525"><code>02c8323</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f4ac954c5b3567f7918fbaa845bd05a8b211f54e"><code>f4ac954</code></a> Bump smithy-go version and update imports for evenstream protocoltests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3420">#3420</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d937001e020def8b587dccbe5d803933ce57bfd"><code>6d93700</code></a> Add replace for credentials dependency added on go.mod (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3419">#3419</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.17...config/v1.32.18">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.16 to 1.19.17
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/db9f4e546dfe2f62a6bc3bf54b9da42ebace6372"><code>db9f4e5</code></a> Release 2026-05-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34e7ddc9400e830a9ae226a7e3c2161e5ece4f19"><code>34e7ddc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f9db036cf7b3b8a1ea5eb67c3d296da4b48b6e2b"><code>f9db036</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ae5eae1e3ec46433bd99496bfa6936f8f09a2e72"><code>ae5eae1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/429dbdd2a35d325aabc5757edfc9ebf09c2ad12e"><code>429dbdd</code></a> Feat discover endpoint partition validation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3410">#3410</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ab4f5b60785064ec6346c922604d94b63d9c7299"><code>ab4f5b6</code></a> Release 2026-05-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/757a09909a97a15e5a481d9839b83f15b8fdc4bc"><code>757a099</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/02c8323ee6c99be82dae3a3923616756cb164525"><code>02c8323</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f4ac954c5b3567f7918fbaa845bd05a8b211f54e"><code>f4ac954</code></a> Bump smithy-go version and update imports for evenstream protocoltests (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3420">#3420</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6d937001e020def8b587dccbe5d803933ce57bfd"><code>6d93700</code></a> Add replace for credentials dependency added on go.mod (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3419">#3419</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.16...credentials/v1.19.17">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.5.1778263552 to 2026.5.1778785514
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/bf471c3bb513fe9d84910d56984d47fac72e4220"><code>bf471c3</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a6b69d36b4fb7a8d53bd4983...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.5.1778263552...release/v2026.5.1778785514">compare view</a></li>
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

### Review by @swadeley - Approved on May 26, 2026 at 08:28 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1504*
