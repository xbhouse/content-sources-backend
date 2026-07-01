---
type: pull_request
number: 1189
title: "Build: Bump the go group with 10 updates"
state: merged
author: dependabot
created: 2025-09-01T07:50:36Z
updated: 2025-09-01T08:40:09Z
closed: 2025-09-01T08:40:02Z
merged: 2025-09-01T08:40:01Z
base_branch: main
head_branch: dependabot/go_modules/go-928076e91e
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1189
---

# Pull Request #1189: Build: Bump the go group with 10 updates

**Author**: @dependabot
**Created**: September 01, 2025 at 07:50 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-928076e91e`

## Description

Bumps the go group with 10 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.132.0` | `0.133.0` |
| [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) | `4.18.3` | `4.19.0` |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | `1.11.0` | `1.11.1` |
| [gorm.io/gorm](https://github.com/go-gorm/gorm) | `1.30.1` | `1.30.2` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.45.2` | `1.46.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.38.1` | `1.38.3` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.2` | `1.31.6` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.6` | `1.18.10` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.56.2` | `1.57.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.87.1` | `1.87.3` |

Updates `github.com/getkin/kin-openapi` from 0.132.0 to 0.133.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.133.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3: resolve Snyk security warning with path traversal by <a href="https://github.com/seborama"><code>@​seborama</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1066">getkin/kin-openapi#1066</a></li>
<li>openapi3: replace bigfloat with decimal128 to fix rounding errors during validation by <a href="https://github.com/Revolyssup"><code>@​Revolyssup</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1068">getkin/kin-openapi#1068</a></li>
<li>openapi2conv: Preserve externalDocs on operations during conversion by <a href="https://github.com/hwustrack"><code>@​hwustrack</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1070">getkin/kin-openapi#1070</a></li>
<li>openapi3: fix ineffectual caching of compiled regexps by <a href="https://github.com/philpearl"><code>@​philpearl</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1076">getkin/kin-openapi#1076</a></li>
<li>openapi3: use Ptr instead of BoolPtr,Float64Ptr,Int64Ptr,Uint64Ptr by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1033">getkin/kin-openapi#1033</a></li>
<li>openapi3: resolve refs in parameter examples by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1086">getkin/kin-openapi#1086</a></li>
<li>openapifilter: Add support for RFC 7396 application/merge-patch+json by <a href="https://github.com/byted"><code>@​byted</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1084">getkin/kin-openapi#1084</a></li>
<li>openapi3filter: use FileBodyDecoder if the format is specified as binary by <a href="https://github.com/dbarrosop"><code>@​dbarrosop</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1088">getkin/kin-openapi#1088</a></li>
<li>openapi3: preserve all validation errors for allOf by <a href="https://github.com/alexbakker"><code>@​alexbakker</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1087">getkin/kin-openapi#1087</a></li>
<li>openapi3filter: support primitive parsing for individual text like parts in multipart/form-data by <a href="https://github.com/nmeheus"><code>@​nmeheus</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1090">getkin/kin-openapi#1090</a></li>
<li>Some coding style fixes and cleaning up by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1093">getkin/kin-openapi#1093</a></li>
<li>openapi2conv: preserve x-fields when converting from v2 to v3 by <a href="https://github.com/saltbo"><code>@​saltbo</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1092">getkin/kin-openapi#1092</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/seborama"><code>@​seborama</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1066">getkin/kin-openapi#1066</a></li>
<li><a href="https://github.com/Revolyssup"><code>@​Revolyssup</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1068">getkin/kin-openapi#1068</a></li>
<li><a href="https://github.com/hwustrack"><code>@​hwustrack</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1070">getkin/kin-openapi#1070</a></li>
<li><a href="https://github.com/philpearl"><code>@​philpearl</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1076">getkin/kin-openapi#1076</a></li>
<li><a href="https://github.com/byted"><code>@​byted</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1084">getkin/kin-openapi#1084</a></li>
<li><a href="https://github.com/dbarrosop"><code>@​dbarrosop</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1088">getkin/kin-openapi#1088</a></li>
<li><a href="https://github.com/nmeheus"><code>@​nmeheus</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1090">getkin/kin-openapi#1090</a></li>
<li><a href="https://github.com/saltbo"><code>@​saltbo</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1092">getkin/kin-openapi#1092</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.132.0...v0.133.0">https://github.com/getkin/kin-openapi/compare/v0.132.0...v0.133.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/2baea3d16906f92e241304527137592a8251afc9"><code>2baea3d</code></a> openapi2conv: preserve x-fields when converting from v2 to v3 (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1092">#1092</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/59b018c1b3242bf93683ddd2fd5372e18270d389"><code>59b018c</code></a> Some coding style fixes and cleaning up (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1093">#1093</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/a33c5f8c6582fc7094f62182bd7c8f62c6bfbcf7"><code>a33c5f8</code></a> openapi3filter: support primitive parsing for individual text like parts in m...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/e00a34088955fdb170656bdd9a54f8a72a77e9c7"><code>e00a340</code></a> openapi3: preserve all validation errors for allOf (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1087">#1087</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/6acf92bcc474992a1328321bb6e875568d513f5f"><code>6acf92b</code></a> openapi3filter: use FileBodyDecoder if the format is specified as binary (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1088">#1088</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/4358c4aa2ae7843b63d0fb264b68995be4e1b180"><code>4358c4a</code></a> feat: add support for RFC 7396 application/merge-patch+json (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1084">#1084</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/d72d75b95b6d35cbd729e9bbd8d03872e0c2df07"><code>d72d75b</code></a> openapi3: resolve refs in parameter examples (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1086">#1086</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/2de45f70d4afe6446074321af3f6d542e91f7f0d"><code>2de45f7</code></a> openapi3: use Ptr instead of BoolPtr,Float64Ptr,Int64Ptr,Uint64Ptr (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1033">#1033</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/783206b0fad6dc1de7c3252362f74fe4ffd00a10"><code>783206b</code></a> openapi3: fix ineffectual caching of compiled regexps (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1076">#1076</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/90fb6412a4b6efaf42cc479a53c9f048d4d30144"><code>90fb641</code></a> openapi2conv: Set externalDocs for Operations (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1070">#1070</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.132.0...v0.133.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/golang-migrate/migrate/v4` from 4.18.3 to 4.19.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.19.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Fixed sqlserver not actually getting a lock if lock is already set by <a href="https://github.com/urbim"><code>@​urbim</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1186">golang-migrate/migrate#1186</a></li>
<li>Bump golang.org/x/oauth2 from 0.18.0 to 0.27.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1299">golang-migrate/migrate#1299</a></li>
<li>Update apt-key to gpg by <a href="https://github.com/sandhilt"><code>@​sandhilt</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1277">golang-migrate/migrate#1277</a></li>
<li>Update dktest to v0.4.6 for docker vuln fix by <a href="https://github.com/dhui"><code>@​dhui</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1309">golang-migrate/migrate#1309</a></li>
<li>refactor: Remove go.uber.org/atomic in favor of std sync/atomic by <a href="https://github.com/romshark"><code>@​romshark</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1303">golang-migrate/migrate#1303</a></li>
<li>Ensure bufferWriter is always closed in Migration.Buffer and propagate close errors by <a href="https://github.com/ckantcs"><code>@​ckantcs</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1308">golang-migrate/migrate#1308</a></li>
<li>Add support for Go 1.25 and drop support for 1.23 by <a href="https://github.com/dhui"><code>@​dhui</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1310">golang-migrate/migrate#1310</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/urbim"><code>@​urbim</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1186">golang-migrate/migrate#1186</a></li>
<li><a href="https://github.com/sandhilt"><code>@​sandhilt</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1277">golang-migrate/migrate#1277</a></li>
<li><a href="https://github.com/romshark"><code>@​romshark</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1303">golang-migrate/migrate#1303</a></li>
<li><a href="https://github.com/ckantcs"><code>@​ckantcs</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1308">golang-migrate/migrate#1308</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/golang-migrate/migrate/compare/v4.18.3...v4.19.0">https://github.com/golang-migrate/migrate/compare/v4.18.3...v4.19.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/8b9c5f77128ef93d65a082208a2009a3911fe6d4"><code>8b9c5f7</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1310">#1310</a> from dhui/update_go</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/b4ec9bccb14cd7b0eb0510ac9d3d01d4be79324f"><code>b4ec9bc</code></a> Add support for Go 1.25 and drop support for 1.23</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/ed4bdd4614e991ca8fba606b38a45f6409a9deb6"><code>ed4bdd4</code></a> Ensure bufferWriter is always closed in Migration.Buffer and propagate close ...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/8945e853c4c84a92bfa42572ac4cfd7874e23108"><code>8945e85</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1303">#1303</a> from romshark/master</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/7108d806dd50c1510510239a161887757f240122"><code>7108d80</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1309">#1309</a> from dhui/dktest_v0.4.6</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/682016f04c2c0f8faa1d122fd22a62563876f71d"><code>682016f</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1277">#1277</a> from sandhilt/doc/change-apt-key-to-gpg</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/f3e6b5a737de1b7c79aaf81168395d6f2eb1fbb0"><code>f3e6b5a</code></a> Replace usage of deprecated docker types</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/0a17402aa2359c8c5cd5a68a0e08a41c4460337f"><code>0a17402</code></a> Update dktest to v0.4.6 for docker vuln fix</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/5eee0c8030ca227bfd27a425988f9c7c4948e90a"><code>5eee0c8</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1299">#1299</a> from golang-migrate/dependabot/go_modules/golang.org...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/642a24d61bb0a870a15020adceeee3e85d5151c3"><code>642a24d</code></a> Bump golang.org/x/oauth2 from 0.18.0 to 0.27.0</li>
<li>Additional commits viewable in <a href="https://github.com/golang-migrate/migrate/compare/v4.18.3...v4.19.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/stretchr/testify` from 1.11.0 to 1.11.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stretchr/testify/releases">github.com/stretchr/testify's releases</a>.</em></p>
<blockquote>
<h2>v1.11.1</h2>
<p>This release fixes <a href="https://redirect.github.com/stretchr/testify/issues/1785">#1785</a> introduced in v1.11.0 where expected argument values implementing the stringer interface (<code>String() string</code>) with a method which mutates their value, when passed to mock.Mock.On (<code>m.On(&quot;Method&quot;, &lt;expected&gt;).Return()</code>) or actual argument values passed to mock.Mock.Called may no longer match one another where they previously did match. The behaviour prior to v1.11.0 where the stringer is always called is restored. Future testify releases may not call the stringer method at all in this case.</p>
<h2>What's Changed</h2>
<ul>
<li>Backport <a href="https://redirect.github.com/stretchr/testify/issues/1786">#1786</a> to release/1.11: mock: revert to pre-v1.11.0 argument matching behavior for mutating stringers by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1788">stretchr/testify#1788</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/stretchr/testify/compare/v1.11.0...v1.11.1">https://github.com/stretchr/testify/compare/v1.11.0...v1.11.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stretchr/testify/commit/2a57335dc9cd6833daa820bc94d9b40c26a7917d"><code>2a57335</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1788">#1788</a> from brackendawson/1785-backport-1.11</li>
<li><a href="https://github.com/stretchr/testify/commit/af8c91234f184009f57ef29027b39ca89cb00100"><code>af8c912</code></a> Backport <a href="https://redirect.github.com/stretchr/testify/issues/1786">#1786</a> to release/1.11</li>
<li>See full diff in <a href="https://github.com/stretchr/testify/compare/v1.11.0...v1.11.1">compare view</a></li>
</ul>
</details>
<br />

Updates `gorm.io/gorm` from 1.30.1 to 1.30.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.30.2</h2>
<h2>Changes</h2>
<ul>
<li>avoid copying structures with embedded mutexs <a href="https://github.com/drakkan"><code>@​drakkan</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7571">#7571</a>)</li>
<li>Add DefaultContextTimeout option <a href="https://github.com/jinzhu"><code>@​jinzhu</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7567">#7567</a>)</li>
<li>Update the docker compose file <a href="https://github.com/moseszane168"><code>@​moseszane168</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7524">#7524</a>)</li>
<li>fix: returning all columns with &quot;on conflict do update&quot;  <a href="https://github.com/phongphan"><code>@​phongphan</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7534">#7534</a>)</li>
<li>feat(slog-integration) <a href="https://github.com/rezamokaram"><code>@​rezamokaram</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7537">#7537</a>)</li>
<li>fix data race in some case <a href="https://github.com/go-gorm/gorm/commit/725aa5b5ff4c0687b06d9a01096b8e4cf96b6c9e">https://github.com/go-gorm/gorm/commit/725aa5b5ff4c0687b06d9a01096b8e4cf96b6c9e</a></li>
<li>performance improvement</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/cace4a6eb85715c8fd648f76728ec554ecafb8e5"><code>cace4a6</code></a> avoid copying structures with embedded mutexs (<a href="https://redirect.github.com/go-gorm/gorm/issues/7571">#7571</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/7ceb0d9e63caebb23098a1886059d2a7dca0bb82"><code>7ceb0d9</code></a> Add DefaultContextTimeout option (<a href="https://redirect.github.com/go-gorm/gorm/issues/7567">#7567</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/4e34a6d21b63e9a9b701a70be9759e5539bf26e9"><code>4e34a6d</code></a> Add tests for sub model</li>
<li><a href="https://github.com/go-gorm/gorm/commit/67de7a8af82adeb86399846986d4c213d09f31b4"><code>67de7a8</code></a> performance improve for schema</li>
<li><a href="https://github.com/go-gorm/gorm/commit/725aa5b5ff4c0687b06d9a01096b8e4cf96b6c9e"><code>725aa5b</code></a> Fix data race, close <a href="https://redirect.github.com/go-gorm/gorm/issues/7287">#7287</a> <a href="https://redirect.github.com/go-gorm/gorm/issues/7110">#7110</a> <a href="https://redirect.github.com/go-gorm/gorm/issues/7539">#7539</a> <a href="https://redirect.github.com/go-gorm/gorm/issues/7108">#7108</a></li>
<li><a href="https://github.com/go-gorm/gorm/commit/6cc2c01268c24fdf7b30cfe4b1a8fb5a10922c65"><code>6cc2c01</code></a> Update the docker compose file (<a href="https://redirect.github.com/go-gorm/gorm/issues/7524">#7524</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/eb90a02a07a950340d9412dfd3a77a74714aaa81"><code>eb90a02</code></a> fix: returning all columns with &quot;on conflict do update&quot; must considered as Sc...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/22d5239dec782d09f54bc11acbd7c99b85f770d2"><code>22d5239</code></a> feat(slog-integration) (<a href="https://redirect.github.com/go-gorm/gorm/issues/7537">#7537</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/gorm/compare/v1.30.1...v1.30.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.45.2 to 1.46.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.46.0 (2025-08-25)</h2>
<blockquote>
<p>[!NOTE]<br />
This release contains significant changes. Notably Sarama will now use the ApiVersionRequest response from each broker to aid in selecting the protocol version to use. The existing <code>Version</code> field in sarama.Config will continue to provide a &quot;pinning&quot; mechanism, but can safely be set to a maximum or higher value than the remote cluster and sarama will sensibly pick compatible versions. There is also a performance improvement relating to MetadataRequests whereby Sarama will avoid having more than a single request to each broker in-flight at any given time. These new (optimal) behaviour is on by default can be opt-ed out via the <code>Metadata.SingleFlight</code> field in Config.</p>
</blockquote>
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(protocol): negotiate API versions by <a href="https://github.com/trapped"><code>@​trapped</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3209">IBM/sarama#3209</a></li>
<li>feat: option to group metadata refreshes so only one is in-flight at a time by <a href="https://github.com/cupcicm"><code>@​cupcicm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3225">IBM/sarama#3225</a></li>
<li>feat: use singleflight metadata by default by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3231">IBM/sarama#3231</a></li>
<li>feat(protocol): support CreateTopicRequest V4 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3238">IBM/sarama#3238</a></li>
<li>feat: always send ApiVersionsRequest and fallback to v0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3234">IBM/sarama#3234</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix(consumer): stuck on the batch with zero records length by <a href="https://github.com/sterligov"><code>@​sterligov</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3221">IBM/sarama#3221</a></li>
<li>fix: sync response header version to clamped request header by <a href="https://github.com/trapped"><code>@​trapped</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3223">IBM/sarama#3223</a></li>
<li>fix(decoder): handle null arrays correctly by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3144">IBM/sarama#3144</a></li>
<li>fix: hardcode lz4 writer blocksize to 64kb by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3258">IBM/sarama#3258</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump the golang-x group across 1 directory with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3185">IBM/sarama#3185</a></li>
<li>chore(deps): bump the golang-x group across 7 directories with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3219">IBM/sarama#3219</a></li>
<li>fix(deps): update module golang.org/x/net to v0.43.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3244">IBM/sarama#3244</a></li>
<li>chore(deps): bump the golang-x group across 6 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3262">IBM/sarama#3262</a></li>
<li>chore(deps): update github/codeql-action action to v3.29.9 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3242">IBM/sarama#3242</a></li>
<li>fix(deps): update github.com/rcrowley/go-metrics digest to 65e299d by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3164">IBM/sarama#3164</a></li>
<li>fix(deps): update module github.com/stretchr/testify to v1.11.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3268">IBM/sarama#3268</a></li>
<li>chore(deps): update docker/bake-action action to v6.9.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3264">IBM/sarama#3264</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(lint): enable copyloopvar by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3214">IBM/sarama#3214</a></li>
<li>chore: fix inconsistent function name in comment by <a href="https://github.com/stellrust"><code>@​stellrust</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3227">IBM/sarama#3227</a></li>
<li>chore(style): refactor compress.go for readability by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3260">IBM/sarama#3260</a></li>
<li>chore: replace unnecessary go-multierror dependency by <a href="https://github.com/bestbug456"><code>@​bestbug456</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3243">IBM/sarama#3243</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/ibm-mend-app"><code>@​ibm-mend-app</code></a>[bot] made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3201">IBM/sarama#3201</a></li>
<li><a href="https://github.com/alexandear"><code>@​alexandear</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3214">IBM/sarama#3214</a></li>
<li><a href="https://github.com/trapped"><code>@​trapped</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3209">IBM/sarama#3209</a></li>
<li><a href="https://github.com/cupcicm"><code>@​cupcicm</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3225">IBM/sarama#3225</a></li>
<li><a href="https://github.com/sterligov"><code>@​sterligov</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3221">IBM/sarama#3221</a></li>
<li><a href="https://github.com/stellrust"><code>@​stellrust</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3227">IBM/sarama#3227</a></li>
<li><a href="https://github.com/bestbug456"><code>@​bestbug456</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3243">IBM/sarama#3243</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.45.2...v1.46.0">https://github.com/IBM/sarama/compare/v1.45.2...v1.46.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/107930bf232af9551df00c37f1e1c588bd097063"><code>107930b</code></a> chore(deps): update docker/bake-action action to v6.9.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3264">#3264</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/f41bb44042889e68d91ea2cb05eb3fdf720f3b6f"><code>f41bb44</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.6 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/088f8c5e884242f8f9d63f59899ee1da89d001dd"><code>088f8c5</code></a> chore: replace unnecessary go-multierror dependency (<a href="https://redirect.github.com/IBM/sarama/issues/3243">#3243</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/dd2886a22b885c63d96b1dcbf8951d2e4c31a82e"><code>dd2886a</code></a> fix(deps): update module github.com/stretchr/testify to v1.11.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3268">#3268</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/c9bdaaa4b393707f388ba15448c5888a1b721782"><code>c9bdaaa</code></a> fix(ci): bump apidiff pin to 1.25.0 compatible version</li>
<li><a href="https://github.com/IBM/sarama/commit/c76c51e169dc0f9d8e6fdcce82f1edf9b1e32ea2"><code>c76c51e</code></a> fix(deps): update github.com/rcrowley/go-metrics digest to 65e299d</li>
<li><a href="https://github.com/IBM/sarama/commit/0c99801095f7e60a9453435e1ee7d3617781be7b"><code>0c99801</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.6 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/f6e84d3c242189fcdd5d805ecc4e122a86bc7d71"><code>f6e84d3</code></a> chore(ci): do not pin docker-compose digests</li>
<li><a href="https://github.com/IBM/sarama/commit/2203569e0f82136f300a7689e158866e7e02968a"><code>2203569</code></a> chore(ci): bump ubi9/ubi-minimal from <code>aaf57d0</code> to <code>295f920</code></li>
<li><a href="https://github.com/IBM/sarama/commit/ba2978c79a886b9116e92cb3b14c0c0481c5ad3b"><code>ba2978c</code></a> chore(ci): fix renovate config</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.45.2...v1.46.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.38.1 to 1.38.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/640b80872b2c49ad5227a4738e15e4189be4863c"><code>640b808</code></a> Release 2025-08-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/657a6e8806831b48766e05f1d9c48b0055b19ebc"><code>657a6e8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34a91bb5d2f6f196769403d5bde31d2725fd28a9"><code>34a91bb</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/09f5a9cb4b1bbda70e3b22599419a610337aeef2"><code>09f5a9c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fff96ca569dd096eb8575d825e466dd55ed3d6f6"><code>fff96ca</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d46f566c13121695d4b22c7964ed7e828d6ae68b"><code>d46f566</code></a> deprecate service/sms (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3176">#3176</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de16d4dc63c5182058b073ebd509688f42c95a15"><code>de16d4d</code></a> Release 2025-08-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d6b83a81af56f06d0ed9edb93aad79ed6313cd51"><code>d6b83a8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0f9cb7f784d4a0c67b9a3383bb1c469b3ebaf07"><code>b0f9cb7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/822549143c11023d454d8af9e6d75bf8cc834a54"><code>8225491</code></a> Release 2025-08-27</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.38.1...v1.38.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.2 to 1.31.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/640b80872b2c49ad5227a4738e15e4189be4863c"><code>640b808</code></a> Release 2025-08-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/657a6e8806831b48766e05f1d9c48b0055b19ebc"><code>657a6e8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34a91bb5d2f6f196769403d5bde31d2725fd28a9"><code>34a91bb</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/09f5a9cb4b1bbda70e3b22599419a610337aeef2"><code>09f5a9c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fff96ca569dd096eb8575d825e466dd55ed3d6f6"><code>fff96ca</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d46f566c13121695d4b22c7964ed7e828d6ae68b"><code>d46f566</code></a> deprecate service/sms (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3176">#3176</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de16d4dc63c5182058b073ebd509688f42c95a15"><code>de16d4d</code></a> Release 2025-08-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d6b83a81af56f06d0ed9edb93aad79ed6313cd51"><code>d6b83a8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0f9cb7f784d4a0c67b9a3383bb1c469b3ebaf07"><code>b0f9cb7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/822549143c11023d454d8af9e6d75bf8cc834a54"><code>8225491</code></a> Release 2025-08-27</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.31.2...config/v1.31.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.6 to 1.18.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b422c03a22416e47fbf567ee4e909329801476a"><code>2b422c0</code></a> Release 2023-01-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9fbc0e483128ded9336cdd440f9f8b98f7037277"><code>9fbc0e4</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/31a07a919d4231cbd3fc9173ff1af2e2a2647bfd"><code>31a07a9</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7d42dcb5b097260736cc9a86d172eb4e3154118d"><code>7d42dcb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e90d54e05b01ddcf201be5cf0ff792bdbbd6113"><code>2e90d54</code></a> fix misleading timestamp value from tests and lint issues (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1994">#1994</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1f6df4a4a0ecb57390da0576e95e47a39493f38a"><code>1f6df4a</code></a> Release 2023-01-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4e52ef7c835ce9bedba44ee0aa355e40225ae96"><code>e4e52ef</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b300bd388f2fed127770fc1c58c4ce12a6662c18"><code>b300bd3</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8157d89e6a517036e190843d80b18d02eff9f0ce"><code>8157d89</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f95afa96016c325b62fdfe1af320bf30bee5f779"><code>f95afa9</code></a> Release 2023-01-23</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.6...config/v1.18.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.56.2 to 1.57.2
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.56.2...service/ssm/v1.57.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.87.1 to 1.87.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/640b80872b2c49ad5227a4738e15e4189be4863c"><code>640b808</code></a> Release 2025-08-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/657a6e8806831b48766e05f1d9c48b0055b19ebc"><code>657a6e8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34a91bb5d2f6f196769403d5bde31d2725fd28a9"><code>34a91bb</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/09f5a9cb4b1bbda70e3b22599419a610337aeef2"><code>09f5a9c</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fff96ca569dd096eb8575d825e466dd55ed3d6f6"><code>fff96ca</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d46f566c13121695d4b22c7964ed7e828d6ae68b"><code>d46f566</code></a> deprecate service/sms (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3176">#3176</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de16d4dc63c5182058b073ebd509688f42c95a15"><code>de16d4d</code></a> Release 2025-08-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d6b83a81af56f06d0ed9edb93aad79ed6313cd51"><code>d6b83a8</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0f9cb7f784d4a0c67b9a3383bb1c469b3ebaf07"><code>b0f9cb7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/822549143c11023d454d8af9e6d75bf8cc834a54"><code>8225491</code></a> Release 2025-08-27</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.87.1...service/s3/v1.87.3">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


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

### Review by @swadeley - Approved on September 01, 2025 at 08:06 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1189*
