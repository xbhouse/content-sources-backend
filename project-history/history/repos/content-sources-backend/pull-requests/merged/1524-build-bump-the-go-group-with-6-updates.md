---
type: pull_request
number: 1524
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2026-06-08T04:39:17Z
updated: 2026-06-08T07:59:47Z
closed: 2026-06-08T07:59:45Z
merged: 2026-06-08T07:59:45Z
base_branch: main
head_branch: dependabot/go_modules/go-273a1fa6ee
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1524
---

# Pull Request #1524: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: June 08, 2026 at 04:39 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-273a1fa6ee`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.41.9` | `1.41.10` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.20` | `1.32.21` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.19` | `1.19.20` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.74.2` | `1.74.3` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.102.2` | `1.103.0` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.9.2` | `5.10.0` |

Updates `github.com/aws/aws-sdk-go-v2` from 1.41.9 to 1.41.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4d02c534cf7f5c782e10bc3f66e40c081a7bb3a"><code>b4d02c5</code></a> Release 2026-06-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48e375b484e0ae917f61904db440ef17382093d1"><code>48e375b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b8a4fc132e65d8310e2de67e0600eec53e2f4b26"><code>b8a4fc1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e8627b4cc01977004c41ff0f42670a44d500982d"><code>e8627b4</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3430">#3430</a> from aws/fix-remove-ioutil</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4e258a368c5b421d28bae4a18e718b1b4c7dd16a"><code>4e258a3</code></a> chore: update changelog description per review</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e1176dfd921faaf4b6c7c3042afa7ad67d599ab8"><code>e1176df</code></a> chore: add changelog entry</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a157f152df84e8398ca0d9a68b9c406dd0e662e3"><code>a157f15</code></a> chore: regenerate SDK with new smithy-go</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d4a893d796beefa696c6878e4284cb9b4c1a792"><code>0d4a893</code></a> chore: bump SMITHY_GO_CODEGEN_VERSION for ioutil cleanup</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/858d954f8c28c01cfc73beee86ffb9b5867578e6"><code>858d954</code></a> fix: remove deprecated io/ioutil from codegen templates</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35a3c50c9feb3bef0de9412a161cee19b27e00d6"><code>35a3c50</code></a> Release 2026-06-01</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.41.9...v1.41.10">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.20 to 1.32.21
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4d02c534cf7f5c782e10bc3f66e40c081a7bb3a"><code>b4d02c5</code></a> Release 2026-06-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48e375b484e0ae917f61904db440ef17382093d1"><code>48e375b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b8a4fc132e65d8310e2de67e0600eec53e2f4b26"><code>b8a4fc1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e8627b4cc01977004c41ff0f42670a44d500982d"><code>e8627b4</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3430">#3430</a> from aws/fix-remove-ioutil</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4e258a368c5b421d28bae4a18e718b1b4c7dd16a"><code>4e258a3</code></a> chore: update changelog description per review</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e1176dfd921faaf4b6c7c3042afa7ad67d599ab8"><code>e1176df</code></a> chore: add changelog entry</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a157f152df84e8398ca0d9a68b9c406dd0e662e3"><code>a157f15</code></a> chore: regenerate SDK with new smithy-go</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d4a893d796beefa696c6878e4284cb9b4c1a792"><code>0d4a893</code></a> chore: bump SMITHY_GO_CODEGEN_VERSION for ioutil cleanup</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/858d954f8c28c01cfc73beee86ffb9b5867578e6"><code>858d954</code></a> fix: remove deprecated io/ioutil from codegen templates</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35a3c50c9feb3bef0de9412a161cee19b27e00d6"><code>35a3c50</code></a> Release 2026-06-01</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.32.20...config/v1.32.21">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.19 to 1.19.20
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4d02c534cf7f5c782e10bc3f66e40c081a7bb3a"><code>b4d02c5</code></a> Release 2026-06-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48e375b484e0ae917f61904db440ef17382093d1"><code>48e375b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b8a4fc132e65d8310e2de67e0600eec53e2f4b26"><code>b8a4fc1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e8627b4cc01977004c41ff0f42670a44d500982d"><code>e8627b4</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3430">#3430</a> from aws/fix-remove-ioutil</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4e258a368c5b421d28bae4a18e718b1b4c7dd16a"><code>4e258a3</code></a> chore: update changelog description per review</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e1176dfd921faaf4b6c7c3042afa7ad67d599ab8"><code>e1176df</code></a> chore: add changelog entry</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a157f152df84e8398ca0d9a68b9c406dd0e662e3"><code>a157f15</code></a> chore: regenerate SDK with new smithy-go</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d4a893d796beefa696c6878e4284cb9b4c1a792"><code>0d4a893</code></a> chore: bump SMITHY_GO_CODEGEN_VERSION for ioutil cleanup</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/858d954f8c28c01cfc73beee86ffb9b5867578e6"><code>858d954</code></a> fix: remove deprecated io/ioutil from codegen templates</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35a3c50c9feb3bef0de9412a161cee19b27e00d6"><code>35a3c50</code></a> Release 2026-06-01</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.19.19...credentials/v1.19.20">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.74.2 to 1.74.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9b0c6553b80f99603b4f8356b88f5baf1328deb"><code>b9b0c65</code></a> Release 2025-10-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2bc8a0ec6f430876fc7de4432ea9cc89c9568f8"><code>e2bc8a0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8691ee380a96c49351e4b5ab8a70bc5d4d100724"><code>8691ee3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51e8a3fe032fc566d31b389f492ab58475a98398"><code>51e8a3f</code></a> bump to go1.23 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3211">#3211</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad2d36cba7c5772b4e8e4caf96939dc41b95c65c"><code>ad2d36c</code></a> Release 2025-10-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19a35d639f969ee328553e632e8cf8b83d324106"><code>19a35d6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35cb02fd50fb125601b9c3b33feb72f3a2bcaa56"><code>35cb02f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f673a1b0a80e666c0128ec606ff053dace9771f1"><code>f673a1b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48421fd812d8592a4eb2b32d11ae07e228969012"><code>48421fd</code></a> Release 2025-10-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fedcba778c21b451a91b4e4bcdd5d6c1554c6a5a"><code>fedcba7</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/eks/v1.74.2...service/eks/v1.74.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.102.2 to 1.103.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4d02c534cf7f5c782e10bc3f66e40c081a7bb3a"><code>b4d02c5</code></a> Release 2026-06-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48e375b484e0ae917f61904db440ef17382093d1"><code>48e375b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b8a4fc132e65d8310e2de67e0600eec53e2f4b26"><code>b8a4fc1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e8627b4cc01977004c41ff0f42670a44d500982d"><code>e8627b4</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3430">#3430</a> from aws/fix-remove-ioutil</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4e258a368c5b421d28bae4a18e718b1b4c7dd16a"><code>4e258a3</code></a> chore: update changelog description per review</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e1176dfd921faaf4b6c7c3042afa7ad67d599ab8"><code>e1176df</code></a> chore: add changelog entry</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a157f152df84e8398ca0d9a68b9c406dd0e662e3"><code>a157f15</code></a> chore: regenerate SDK with new smithy-go</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0d4a893d796beefa696c6878e4284cb9b4c1a792"><code>0d4a893</code></a> chore: bump SMITHY_GO_CODEGEN_VERSION for ioutil cleanup</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/858d954f8c28c01cfc73beee86ffb9b5867578e6"><code>858d954</code></a> fix: remove deprecated io/ioutil from codegen templates</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35a3c50c9feb3bef0de9412a161cee19b27e00d6"><code>35a3c50</code></a> Release 2026-06-01</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.102.2...service/s3/v1.103.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.9.2 to 5.10.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.10.0 (June 3, 2026)</h1>
<p>This release includes a significant amount of hardening against malicious or compromised PostgreSQL servers,
contributed by Sean Chittenden at CrowdStrike, Inc. This work bounds binary decoders against attacker-controlled
message sizes, caps server-supplied SCRAM iteration counts, adds <code>require_auth</code> to restrict which authentication
methods a server may use (mitigating downgrade attacks under <code>sslmode=prefer</code>), and ensures cancellation requests are
sent over TLS when the original connection used TLS.</p>
<h2>Features</h2>
<ul>
<li>Add <code>require_auth</code> to restrict accepted server authentication methods (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Add <code>ParseConfigOptions.ConnStringAllowedKeys</code> to restrict allowed connection string keys (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Add <code>StructArgs</code> and <code>StrictStructArgs</code> for <code>@</code>-named queries (Tubelight30)</li>
<li>Add <code>ErrConnClosed</code> sentinel error and unwrap it from <code>connLockError</code> (Charlie Tonneslan)</li>
<li>pgxpool: check if connection is expired before acquire (arthurdotwork)</li>
</ul>
<h2>Security Hardening</h2>
<ul>
<li>Encrypt <code>CancelRequest</code> connection when the primary connection used TLS (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Cap server-supplied SCRAM iteration count (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Default Frontend max message body length to ~1 GiB (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Bound hstore binary decode against malicious server input (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Bound array binary decode element length against remaining message bytes (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Bound array element count against remaining message bytes (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Bound range, multirange, and tsvector binary decoders (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Document secure connection configuration (Sean Chittenden at CrowdStrike, Inc.)</li>
<li>Fix panic on malformed geometric text; return an error instead (MaIII)</li>
</ul>
<h2>Fixes</h2>
<ul>
<li>Fix scanning <code>&quot;char&quot;</code> (OID 18) into <code>*string</code> in binary format (luongs3)</li>
<li>Fix handling of typed-nil <code>driver.Valuer</code> in array and composite codecs (Donncha Fahy)</li>
<li>Fix <code>CopyData.Data</code> hex decoding in <code>UnmarshalJSON</code> (Charlie Tonneslan)</li>
<li>Fix data race when context is cancelled during connect</li>
<li>Fix <code>parseKeywordValueSettings</code> rejecting trailing whitespace (alliasgher)</li>
<li>pgconn: preserve full error chain in <code>normalizeTimeoutError</code> (Charlie Tonneslan)</li>
<li>pgconn: use a fresh context for the fallback connection in <code>connectPreferred</code> (Charlie Tonneslan)</li>
<li>pgxpool: fix <code>MaxLifetimeDestroyCount</code> and ping order for acquire-time expiry check</li>
<li>Add missing error check of <code>rows.Err</code> to load types (Jen Altavilla)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/7293fb11125be0373a92f716683f2d494f6fd4b0"><code>7293fb1</code></a> Update changelog for v5.10.0</li>
<li><a href="https://github.com/jackc/pgx/commit/1ade2852841d4ee55677207200f4ffdbc217ce69"><code>1ade285</code></a> pgconn: document secure connection configuration</li>
<li><a href="https://github.com/jackc/pgx/commit/b4d6d4d1be7f381bb81d12ebfecae6b10f5c7562"><code>b4d6d4d</code></a> pgtype: bound range, multirange, and tsvector binary decoders</li>
<li><a href="https://github.com/jackc/pgx/commit/0639b37f8f4fff31dbe73297087e69b3ccc3bf2b"><code>0639b37</code></a> pgconn: add ParseConfigOptions.ConnStringAllowedKeys</li>
<li><a href="https://github.com/jackc/pgx/commit/b28e65b0c3e0cd45c09e7c9ce36e5e29caa6dbe9"><code>b28e65b</code></a> pgtype: bound array element count against remaining message bytes</li>
<li><a href="https://github.com/jackc/pgx/commit/cd1f389d37d775bc8cb11c60363946f928c02c98"><code>cd1f389</code></a> pgtype: bound array binary decode element length against remaining bytes</li>
<li><a href="https://github.com/jackc/pgx/commit/ff27b5bbea012020d1fd8b9bdd56284a88783ef1"><code>ff27b5b</code></a> pgtype: bound hstore binary decode against malicious server input</li>
<li><a href="https://github.com/jackc/pgx/commit/a6002e12a8a393844b48c29d105e7542e7b3a251"><code>a6002e1</code></a> pgproto3: default Frontend max message body length to ~1 GiB</li>
<li><a href="https://github.com/jackc/pgx/commit/44f61732ecdfd08081a1a2ff7227f1e975f0b71e"><code>44f6173</code></a> pgconn: cap server-supplied SCRAM iteration count</li>
<li><a href="https://github.com/jackc/pgx/commit/1a976f7bb91216ea7f8369cb7abe78ce34dc244f"><code>1a976f7</code></a> pgconn: add require_auth to restrict accepted server auth methods</li>
<li>Additional commits viewable in <a href="https://github.com/jackc/pgx/compare/v5.9.2...v5.10.0">compare view</a></li>
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

### Review by @swadeley - Approved on June 08, 2026 at 07:59 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1524*
