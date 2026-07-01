---
type: pull_request
number: 1184
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2025-08-25T08:00:09Z
updated: 2025-08-25T15:39:33Z
closed: 2025-08-25T15:39:25Z
merged: 2025-08-25T15:39:25Z
base_branch: main
head_branch: dependabot/go_modules/go-9ae1fa6ba2
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1184
---

# Pull Request #1184: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: August 25, 2025 at 08:00 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9ae1fa6ba2`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | `1.10.0` | `1.11.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.38.0` | `1.38.1` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.0` | `1.31.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.4` | `1.18.6` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.56.0` | `1.56.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.87.0` | `1.87.1` |

Updates `github.com/stretchr/testify` from 1.10.0 to 1.11.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stretchr/testify/releases">github.com/stretchr/testify's releases</a>.</em></p>
<blockquote>
<h2>v1.11.0</h2>
<h2>What's Changed</h2>
<h3>Functional Changes</h3>
<p>v1.11.0 Includes a number of performance improvements.</p>
<ul>
<li>Call stack perf change for CallerInfo by <a href="https://github.com/mikeauclair"><code>@​mikeauclair</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1614">stretchr/testify#1614</a></li>
<li>Lazily render mock diff output on successful match by <a href="https://github.com/mikeauclair"><code>@​mikeauclair</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1615">stretchr/testify#1615</a></li>
<li>assert: check early in Eventually, EventuallyWithT, and Never by <a href="https://github.com/cszczepaniak"><code>@​cszczepaniak</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1427">stretchr/testify#1427</a></li>
<li>assert: add IsNotType by <a href="https://github.com/bartventer"><code>@​bartventer</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1730">stretchr/testify#1730</a></li>
<li>assert.JSONEq: shortcut if same strings by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1754">stretchr/testify#1754</a></li>
<li>assert.YAMLEq: shortcut if same strings by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1755">stretchr/testify#1755</a></li>
<li>assert: faster and simpler isEmpty using reflect.Value.IsZero by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1761">stretchr/testify#1761</a></li>
<li>suite: faster methods filtering (internal refactor) by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1758">stretchr/testify#1758</a></li>
</ul>
<h3>Fixes</h3>
<ul>
<li>assert.ErrorAs: log target type by <a href="https://github.com/craig65535"><code>@​craig65535</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1345">stretchr/testify#1345</a></li>
<li>Fix failure message formatting for Positive and Negative asserts in <a href="https://redirect.github.com/stretchr/testify/pull/1062">stretchr/testify#1062</a></li>
<li>Improve ErrorIs message when error is nil but an error was expected by <a href="https://github.com/tsioftas"><code>@​tsioftas</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1681">stretchr/testify#1681</a></li>
<li>fix Subset/NotSubset when calling with mixed input types by <a href="https://github.com/siliconbrain"><code>@​siliconbrain</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1729">stretchr/testify#1729</a></li>
<li>Improve ErrorAs failure message when error is nil by <a href="https://github.com/ccoVeille"><code>@​ccoVeille</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1734">stretchr/testify#1734</a></li>
<li>mock.AssertNumberOfCalls: improve error msg by <a href="https://github.com/3scalation"><code>@​3scalation</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1743">stretchr/testify#1743</a></li>
</ul>
<h3>Documentation, Build &amp; CI</h3>
<ul>
<li>docs: Fix typo in README by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1688">stretchr/testify#1688</a></li>
<li>Replace deprecated io/ioutil with io and os by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1684">stretchr/testify#1684</a></li>
<li>Document consequences of calling t.FailNow() by <a href="https://github.com/greg0ire"><code>@​greg0ire</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1710">stretchr/testify#1710</a></li>
<li>chore: update docs for Unset <a href="https://redirect.github.com/stretchr/testify/issues/1621">#1621</a> by <a href="https://github.com/techfg"><code>@​techfg</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1709">stretchr/testify#1709</a></li>
<li>README: apply gofmt to examples by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1687">stretchr/testify#1687</a></li>
<li>refactor: use %q and %T to simplify fmt.Sprintf by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1674">stretchr/testify#1674</a></li>
<li>Propose Christophe Colombier (ccoVeille) as approver by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1716">stretchr/testify#1716</a></li>
<li>Update documentation for the Error function in assert or require package by <a href="https://github.com/architagr"><code>@​architagr</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1675">stretchr/testify#1675</a></li>
<li>assert: remove deprecated build constraints by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1671">stretchr/testify#1671</a></li>
<li>assert: apply gofumpt to internal test suite by <a href="https://github.com/ccoVeille"><code>@​ccoVeille</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1739">stretchr/testify#1739</a></li>
<li>CI: fix shebang in .ci.*.sh scripts by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1746">stretchr/testify#1746</a></li>
<li>assert,require: enable parallel testing on (almost) all top tests by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1747">stretchr/testify#1747</a></li>
<li>suite.Passed: add one more status test report by <a href="https://github.com/Ararsa-Derese"><code>@​Ararsa-Derese</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1706">stretchr/testify#1706</a></li>
<li>Add Helper() method in internal mocks and assert.CollectT by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1423">stretchr/testify#1423</a></li>
<li>assert.Same/NotSame: improve usage of Sprintf by <a href="https://github.com/ccoVeille"><code>@​ccoVeille</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1742">stretchr/testify#1742</a></li>
<li>mock: enable parallel testing on internal testsuite by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1756">stretchr/testify#1756</a></li>
<li>suite: cleanup use of 'testing' internals at runtime by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1751">stretchr/testify#1751</a></li>
<li>assert: check test failure message for Empty and NotEmpty  by <a href="https://github.com/ccoVeille"><code>@​ccoVeille</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1745">stretchr/testify#1745</a></li>
<li>deps: fix dependency cycle with objx (again) by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1567">stretchr/testify#1567</a></li>
<li>assert.Empty: comprehensive doc of &quot;Empty&quot;-ness rules by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1753">stretchr/testify#1753</a></li>
<li>doc: improve godoc of top level 'testify' package by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1760">stretchr/testify#1760</a></li>
<li>assert.ErrorAs: simplify retrieving the type name by <a href="https://github.com/ccoVeille"><code>@​ccoVeille</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1740">stretchr/testify#1740</a></li>
<li>assert.EqualValues: improve test coverage to 100% by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1763">stretchr/testify#1763</a></li>
<li>suite.Run: simplify running of Setup/TeardownSuite by <a href="https://github.com/renzoarreaza"><code>@​renzoarreaza</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1769">stretchr/testify#1769</a></li>
<li>assert.CallerInfo: micro optimization by using LastIndexByte by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1767">stretchr/testify#1767</a></li>
<li>assert.CallerInfo: micro cleanup by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1768">stretchr/testify#1768</a></li>
<li>assert: refactor Test<em>FileExists and Test</em>DirExists tests to enable parallel testing by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1766">stretchr/testify#1766</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stretchr/testify/commit/b7801fbf5cd58d201296d5d0e132d1849966dbd4"><code>b7801fb</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1778">#1778</a> from stretchr/dependabot/github_actions/actions/chec...</li>
<li><a href="https://github.com/stretchr/testify/commit/69831f3b08c40d56a09d0be93e9d5ae034f1590b"><code>69831f3</code></a> build(deps): bump actions/checkout from 4 to 5</li>
<li><a href="https://github.com/stretchr/testify/commit/a53be35c3b0cfcd5189cffcfd75df60ea581104c"><code>a53be35</code></a> Improve captureTestingT helper</li>
<li><a href="https://github.com/stretchr/testify/commit/aafb604176db7e1f2c9810bc90d644291d057687"><code>aafb604</code></a> mock: improve formatting of error message</li>
<li><a href="https://github.com/stretchr/testify/commit/7218e0390acd2aea3edb18574110ec2753c0aeef"><code>7218e03</code></a> improve error msg</li>
<li><a href="https://github.com/stretchr/testify/commit/929a2126c2702df436312656a0304580b526c6e9"><code>929a212</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1758">#1758</a> from stretchr/dolmen/suite-faster-method-filtering</li>
<li><a href="https://github.com/stretchr/testify/commit/bc7459ec38128532ff32f23cfab4ea0b725210f2"><code>bc7459e</code></a> suite: faster filtering of methods (-testify.m)</li>
<li><a href="https://github.com/stretchr/testify/commit/7d37b5c962954410bcd7a71ff3a77c79514056d1"><code>7d37b5c</code></a> suite: refactor methodFilter</li>
<li><a href="https://github.com/stretchr/testify/commit/c58bc90e5c2a1d1bd5d99e8b4708023ec5a97d46"><code>c58bc90</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1764">#1764</a> from stretchr/dolmen/suite-refactor-stats-for-readab...</li>
<li><a href="https://github.com/stretchr/testify/commit/87101a6e4a5859cee372b6ded7821787b3190cb7"><code>87101a6</code></a> suite.Run: refactor handling of stats</li>
<li>Additional commits viewable in <a href="https://github.com/stretchr/testify/compare/v1.10.0...v1.11.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.38.0 to 1.38.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4ee4da236306b260326a7a913f61cb19355110"><code>ba4ee4d</code></a> Release 2025-08-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40e3d871f1507d7b7a10b101dd65c5c85ec183c2"><code>40e3d87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2be01902dbbdbdec11e3fe7a9ca56aa45c9edcd"><code>b2be019</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dece4e78c3752a54cc2393bf375672ca7b66b260"><code>dece4e7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/294af1979f20160f82f273fd00790466bc8f7daa"><code>294af19</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0df860a876d097b792f61fd35caea13c86247d46"><code>0df860a</code></a> changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/df2bca243bed6101bdee10478def6146a7f7e647"><code>df2bca2</code></a> feature(s3/manager): add option to control default checksums (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3151">#3151</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/315de9ca18b06a3bc807313c9f79b56e2956009a"><code>315de9c</code></a> Release 2025-08-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a9d79d3c8d2dcf70265875f2ed6a8af678454d5"><code>1a9d79d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79594120103fbf7a5aa836f8c640b9c255453835"><code>7959412</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.38.0...v1.38.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.0 to 1.31.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4ee4da236306b260326a7a913f61cb19355110"><code>ba4ee4d</code></a> Release 2025-08-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40e3d871f1507d7b7a10b101dd65c5c85ec183c2"><code>40e3d87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2be01902dbbdbdec11e3fe7a9ca56aa45c9edcd"><code>b2be019</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dece4e78c3752a54cc2393bf375672ca7b66b260"><code>dece4e7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/294af1979f20160f82f273fd00790466bc8f7daa"><code>294af19</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0df860a876d097b792f61fd35caea13c86247d46"><code>0df860a</code></a> changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/df2bca243bed6101bdee10478def6146a7f7e647"><code>df2bca2</code></a> feature(s3/manager): add option to control default checksums (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3151">#3151</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/315de9ca18b06a3bc807313c9f79b56e2956009a"><code>315de9c</code></a> Release 2025-08-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a9d79d3c8d2dcf70265875f2ed6a8af678454d5"><code>1a9d79d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79594120103fbf7a5aa836f8c640b9c255453835"><code>7959412</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.31.0...config/v1.31.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.4 to 1.18.6
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-12-19)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/athena</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/athena/CHANGELOG.md#v1210-2022-12-19">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: Add missed InvalidRequestException in GetCalculationExecutionCode,StopCalculationExecution APIs. Correct required parameters (Payload and Type) in UpdateNotebook API. Change Notebook size from 15 Mb to 10 Mb.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ecs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/ecs/CHANGELOG.md#v1220-2022-12-19">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: This release adds support for alarm-based rollbacks in ECS, a new feature that allows customers to add automated safeguards for Amazon ECS service rolling updates.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kinesisvideo</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/kinesisvideo/CHANGELOG.md#v1140-2022-12-19">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Amazon Kinesis Video Streams offers capabilities to stream video and audio in real-time via WebRTC to the cloud for storage, playback, and analytical processing. Customers can use our enhanced WebRTC SDK and cloud APIs to enable real-time streaming, as well as media ingestion to the cloud.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/kinesisvideowebrtcstorage/CHANGELOG.md#v100-2022-12-19">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
<li><strong>Feature</strong>: Amazon Kinesis Video Streams offers capabilities to stream video and audio in real-time via WebRTC to the cloud for storage, playback, and analytical processing. Customers can use our enhanced WebRTC SDK and cloud APIs to enable real-time streaming, as well as media ingestion to the cloud.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/rds</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/rds/CHANGELOG.md#v1360-2022-12-19">v1.36.0</a>
<ul>
<li><strong>Feature</strong>: Add support for --enable-customer-owned-ip to RDS create-db-instance-read-replica API for RDS on Outposts.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/sagemaker/CHANGELOG.md#v1590-2022-12-19">v1.59.0</a>
<ul>
<li><strong>Feature</strong>: AWS Sagemaker - Sagemaker Images now supports Aliases as secondary identifiers for ImageVersions. SageMaker Images now supports additional metadata for ImageVersions for better images management.</li>
</ul>
</li>
</ul>
<h1>Release (2022-12-16)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/appflow/CHANGELOG.md#v1220-2022-12-16">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: This release updates the ListConnectorEntities API action so that it returns paginated responses that customers can retrieve with next tokens.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/cloudfront</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/cloudfront/CHANGELOG.md#v1222-2022-12-16">v1.22.2</a>
<ul>
<li><strong>Documentation</strong>: Updated documentation for CloudFront</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/datasync</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/datasync/CHANGELOG.md#v1200-2022-12-16">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: AWS DataSync now supports the use of tags with task executions. With this new feature, you can apply tags each time you execute a task, giving you greater control and management over your task executions.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/efs</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/efs/CHANGELOG.md#v1183-2022-12-16">v1.18.3</a>
<ul>
<li><strong>Documentation</strong>: General documentation updates for EFS.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/guardduty</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/guardduty/CHANGELOG.md#v1166-2022-12-16">v1.16.6</a>
<ul>
<li><strong>Documentation</strong>: This release provides the valid characters for the Description and Name field.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotfleetwise</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/iotfleetwise/CHANGELOG.md#v120-2022-12-16">v1.2.0</a>
<ul>
<li><strong>Feature</strong>: Updated error handling for empty resource names in &quot;UpdateSignalCatalog&quot; and &quot;GetModelManifest&quot; operations.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sagemaker</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/service/sagemaker/CHANGELOG.md#v1580-2022-12-16">v1.58.0</a>
<ul>
<li><strong>Feature</strong>: AWS sagemaker - Features: This release adds support for random seed, it's an integer value used to initialize a pseudo-random number generator. Setting a random seed will allow the hyperparameter tuning search strategies to produce more consistent configurations for the same tuning job.</li>
</ul>
</li>
</ul>
<h1>Release (2022-12-15)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.17.3
<ul>
<li><strong>Bug Fix</strong>: Unify logic between shared config and in finding home directory</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/config</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/config/CHANGELOG.md#v1185-2022-12-15">v1.18.5</a>
<ul>
<li><strong>Bug Fix</strong>: Unify logic between shared config and in finding home directory</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/credentials</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/config/v1.18.6/credentials/CHANGELOG.md#v1135-2022-12-15">v1.13.5</a>
<ul>
<li><strong>Bug Fix</strong>: Unify logic between shared config and in finding home directory</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9149b86ed65b3060c4cd185f28ecbcae50e4a39c"><code>9149b86</code></a> Release 2022-12-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a0950979e4fec599726bf5c3e39e1564065e88b9"><code>a095097</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0fdfb08f435dade979fb8190d6cc1f74bd25503"><code>f0fdfb0</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0147dd3e7d655cee5979bd689f1613b4ab7a3d2c"><code>0147dd3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d7b7ceff6365a5725bbc391e04b91f9e5b8eb41c"><code>d7b7cef</code></a> Release 2022-12-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/337b1f42dff5688f937474692dcbdf1381401299"><code>337b1f4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8dc89283777b5c52cca2a01668e58212655e1bf"><code>c8dc892</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7e4aeab5643ff4a28e5b215a0cad215992aad3d"><code>e7e4aea</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1257d4e205318d7dea08f5785d51d8ced05ccde6"><code>1257d4e</code></a> Release 2022-12-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f89f556a125709070bf3614e565233a90d2ee17c"><code>f89f556</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.4...config/v1.18.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.56.0 to 1.56.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b06fd664465de19f26e46ada588801c1c3013f78"><code>b06fd66</code></a> Release 2025-04-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/989866c3da168aeb7f89d2f8ec6804c573c5789a"><code>989866c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a3a3dfc32c5aa5df53568222252db40d894263c"><code>5a3a3df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa132cce708ed0b059b7cbd23962ef755b8b080d"><code>aa132cc</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3aace4178c90b4b6acd3e6485e04e752c7c30a85"><code>3aace41</code></a> Transfer Manager v2 downloader concurrency fix and version control (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3058">#3058</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2e6c6f380f3eb2c998e000ca53c62c193a3c2ed"><code>a2e6c6f</code></a> Release 2025-04-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6c2c1890e46ca5b6a32a8bba7ec66abef9947c19"><code>6c2c189</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9ec7771051f2b623cb9a8fe162266b54c2243632"><code>9ec7771</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3d384f6bb6b1e27c70999c4d6d0a7579e18336af"><code>3d384f6</code></a> Release 2025-04-23</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b1a942dae33bc0324e541bb1bd20697fa557536c"><code>b1a942d</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.56.0...service/ecs/v1.56.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.87.0 to 1.87.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4ee4da236306b260326a7a913f61cb19355110"><code>ba4ee4d</code></a> Release 2025-08-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40e3d871f1507d7b7a10b101dd65c5c85ec183c2"><code>40e3d87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b2be01902dbbdbdec11e3fe7a9ca56aa45c9edcd"><code>b2be019</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dece4e78c3752a54cc2393bf375672ca7b66b260"><code>dece4e7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/294af1979f20160f82f273fd00790466bc8f7daa"><code>294af19</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0df860a876d097b792f61fd35caea13c86247d46"><code>0df860a</code></a> changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/df2bca243bed6101bdee10478def6146a7f7e647"><code>df2bca2</code></a> feature(s3/manager): add option to control default checksums (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3151">#3151</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/315de9ca18b06a3bc807313c9f79b56e2956009a"><code>315de9c</code></a> Release 2025-08-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1a9d79d3c8d2dcf70265875f2ed6a8af678454d5"><code>1a9d79d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79594120103fbf7a5aa836f8c640b9c255453835"><code>7959412</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.87.0...service/s3/v1.87.1">compare view</a></li>
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

### Review by @rverdile - Approved on August 25, 2025 at 03:39 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1184*
