---
type: pull_request
number: 847
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2024-10-14T04:10:26Z
updated: 2024-10-14T13:29:27Z
closed: 2024-10-14T13:29:20Z
merged: 2024-10-14T13:29:20Z
base_branch: main
head_branch: dependabot/go_modules/go-41cb0903be
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/847
---

# Pull Request #847: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: October 14, 2024 at 04:10 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-41cb0903be`

## Description

Bumps the go group with 4 updates: [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.0 to 1.32.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cbb5aa17f9078cb45dc0e82d3e1d0abee3744a9"><code>0cbb5aa</code></a> Release 2024-10-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54c1dd6c74185b0c7df78159ec4d5b2c27e9e280"><code>54c1dd6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2cde144eedda9f509141751c3011ca64a6b6528e"><code>2cde144</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67fbd35762ef8694839df209714d2ec2c33d3df9"><code>67fbd35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa04330cb6978ccb6a7bb3e198b3fb21abbd6333"><code>aa04330</code></a> Allow non-nil but empty headers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2826">#2826</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a4e5bb42c08ff5a4e0e601a7461c8466565e44e"><code>5a4e5bb</code></a> add feature tracking for cbor protocol (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2821">#2821</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/183987cda0c2487a1b25c8e9cbf8dba510046c73"><code>183987c</code></a> add annotations to deprecated services and introduce codegen integration for ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b737dc9eb14847cd97d3b30ad6a1394efd73245b"><code>b737dc9</code></a> Release 2024-10-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7279a51bbcd597f4aa7aeeb599c017d3d1679fb6"><code>7279a51</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a1b1f5a17c687371cc53c5dfbb2bf5ff467ff51a"><code>a1b1f5a</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.0...v1.32.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.39 to 1.17.41
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cbb5aa17f9078cb45dc0e82d3e1d0abee3744a9"><code>0cbb5aa</code></a> Release 2024-10-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54c1dd6c74185b0c7df78159ec4d5b2c27e9e280"><code>54c1dd6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2cde144eedda9f509141751c3011ca64a6b6528e"><code>2cde144</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67fbd35762ef8694839df209714d2ec2c33d3df9"><code>67fbd35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa04330cb6978ccb6a7bb3e198b3fb21abbd6333"><code>aa04330</code></a> Allow non-nil but empty headers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2826">#2826</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a4e5bb42c08ff5a4e0e601a7461c8466565e44e"><code>5a4e5bb</code></a> add feature tracking for cbor protocol (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2821">#2821</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/183987cda0c2487a1b25c8e9cbf8dba510046c73"><code>183987c</code></a> add annotations to deprecated services and introduce codegen integration for ...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b737dc9eb14847cd97d3b30ad6a1394efd73245b"><code>b737dc9</code></a> Release 2024-10-07</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7279a51bbcd597f4aa7aeeb599c017d3d1679fb6"><code>7279a51</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a1b1f5a17c687371cc53c5dfbb2bf5ff467ff51a"><code>a1b1f5a</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.39...credentials/v1.17.41">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.41.0 to 1.41.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de4f646b7271b9698a67f63bff876ad40eb690ca"><code>de4f646</code></a> Release 2024-02-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a31327acbd07ca1d42ca83107993d7bcfa223dd8"><code>a31327a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d1333952165a0f8ceea1c7de5d20d06af4925c30"><code>d133395</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b135f991ff3bef2f7c9ce8117dd59e1abae43a0f"><code>b135f99</code></a> fix(2502): zero region should explicitly fail endpoint resolution (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2503">#2503</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7cb105f17b3a12d4b69219ce9957d517dba299d8"><code>7cb105f</code></a> Release 2024-02-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7f9c4d2eee136647971352eb9d06b196cc9abaf0"><code>7f9c4d2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a60d4b2fc37bab61dc219d9e210369c4bf4e445e"><code>a60d4b2</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be0accd6c6511b727d38cbe98489753e481d68b3"><code>be0accd</code></a> fix: panic due to potentially uncomparable wrapped fn in express_resolve (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2500">#2500</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4334b43c27eeb4c4ecfcba1a87cc08f963fe91a3"><code>4334b43</code></a> Release 2024-02-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9e29187dc08c2e9bf8f66166841ca9c75e0624ab"><code>9e29187</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.41.0...service/fsx/v1.41.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.10.1727972141 to 2024.10.1728596503
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/c6861f618f94d99bf4192c64adba9391c0cca1b3"><code>c6861f6</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d85eb38da157952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/049c3e1a6c99c8347826714bfa25df16cf9db566"><code>049c3e1</code></a> Update pulp bindings to e69db48356e528a464be3da896237bd3a2e69e1a7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/3ace198185ee6fb58f3a415d291f4197b2e9ca7a"><code>3ace198</code></a> Update pulp bindings to 386d86254354e29d3b864523aed47863a3db8d167968b5e2da9ab...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.10.1727972141...release/v2024.10.1728596503">compare view</a></li>
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

### Comment by @app-sre-bot on October 14, 2024 at 04:12 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on October 14, 2024 at 08:23 AM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on October 14, 2024 at 01:29 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/847*
