---
type: pull_request
number: 726
title: "Build: Bump the go group across 1 directory with 6 updates"
state: closed
author: dependabot
created: 2024-07-01T04:33:20Z
updated: 2024-07-01T13:55:56Z
closed: 2024-07-01T13:55:55Z
base_branch: main
head_branch: dependabot/go_modules/go-6790a275f9
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/726
---

# Pull Request #726: Build: Bump the go group across 1 directory with 6 updates

**Author**: @dependabot
**Created**: July 01, 2024 at 04:33 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-6790a275f9`

## Description

Bumps the go group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto) | `1.1.0-alpha.2-proton` | `1.1.0-alpha.3-proton` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.27.2` | `1.30.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.18` | `1.17.23` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.35.7` | `1.37.1` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.4.6` | `4.4.9` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.6.1718394381` | `2024.6.1719602258` |


Updates `github.com/ProtonMail/go-crypto` from 1.1.0-alpha.2-proton to 1.1.0-alpha.3-proton
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/ProtonMail/go-crypto/releases">github.com/ProtonMail/go-crypto's releases</a>.</em></p>
<blockquote>
<h2>v1.1.0-alpha.3-proton</h2>
<p>This pre-release is v1.1.0-alpha.3 with support for symmetric keys and automatic forwarding, both of which are not standardized yet.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/dfc52108cb1d8cc96044e7c1ccdc0265e60ba5fe"><code>dfc5210</code></a> Fix HMAC generation (<a href="https://redirect.github.com/ProtonMail/go-crypto/issues/204">#204</a>)</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/a0d3abb281a0021163b3380a7a05643f1a634cfc"><code>a0d3abb</code></a> Replace ioutil.ReadAll with io.ReadAll</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/12d80eef64055dd10f98108f1baa8319f41c6735"><code>12d80ee</code></a> fix(v2): Adapt NewForwardingEntity to refactored NewEntity</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/00db3bea995635782ec4b97ff0d40d494244b1ec"><code>00db3be</code></a> fix(v2): Do not allow encrpytion with a forwarding key</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/49c1020fc9703a6ae7535f51ce4e28bf40faea61"><code>49c1020</code></a> feat: Add symmetric keys to v2</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/bce9b31864fde41c9410067bd7556b6e0768727b"><code>bce9b31</code></a> fix: Address warnings</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/e202bef191f1b5f90fc94abee1367234163946bd"><code>e202bef</code></a> feat: Add forwarding to v2 api</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/2a86d284854416df95abcdf6ff92e86d530ddeef"><code>2a86d28</code></a> fix: Address rebase on version 2 issues</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/a425074c01c1ff26779d7e745ff0f709e53cc96f"><code>a425074</code></a> Use fingerprints instead of KeyIDs</li>
<li><a href="https://github.com/ProtonMail/go-crypto/commit/b55bb97a315ada47aa46dcbcca6ebaa5b2e8c965"><code>b55bb97</code></a> Create a copy of the encrypted key when forwarding</li>
<li>Additional commits viewable in <a href="https://github.com/ProtonMail/go-crypto/compare/v1.1.0-alpha.2-proton...v1.1.0-alpha.3-proton">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.27.2 to 1.30.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4509a600408280c8dcdbc6825ba750cf1628423d"><code>4509a60</code></a> Release 2024-06-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0c61504d74dd81214542aae8a68993166935fa2a"><code>0c61504</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ca59bf2c0fc11f80a59f7f0e5c7f4d7805444e3"><code>7ca59bf</code></a> Update SDK's smithy-go dependency to v1.20.3</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe1af5ba9947870b221f0c17b92e4f6d48e318c9"><code>fe1af5b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b76404f69b8b6b2e398f9fd3f4759e36e3407353"><code>b76404f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ca412df63bb420ae20bf4ae198b8ea57e2aacf6"><code>8ca412d</code></a> Release 2024-06-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6505a148d0d4e3a624fa5fb10c7fc4a226c9257f"><code>6505a14</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/66e5439ffbb921f031bae700e78ecea6b48b3517"><code>66e5439</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/590c9dbbfd52c9f93a2997bcca309c362d7f90e1"><code>590c9db</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7c25c211744bdcff47a7203a7a894b1241f9da50"><code>7c25c21</code></a> Release 2024-06-26</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.27.2...v1.30.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.18 to 1.17.23
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4509a600408280c8dcdbc6825ba750cf1628423d"><code>4509a60</code></a> Release 2024-06-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0c61504d74dd81214542aae8a68993166935fa2a"><code>0c61504</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ca59bf2c0fc11f80a59f7f0e5c7f4d7805444e3"><code>7ca59bf</code></a> Update SDK's smithy-go dependency to v1.20.3</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe1af5ba9947870b221f0c17b92e4f6d48e318c9"><code>fe1af5b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b76404f69b8b6b2e398f9fd3f4759e36e3407353"><code>b76404f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ca412df63bb420ae20bf4ae198b8ea57e2aacf6"><code>8ca412d</code></a> Release 2024-06-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6505a148d0d4e3a624fa5fb10c7fc4a226c9257f"><code>6505a14</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/66e5439ffbb921f031bae700e78ecea6b48b3517"><code>66e5439</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/590c9dbbfd52c9f93a2997bcca309c362d7f90e1"><code>590c9db</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7c25c211744bdcff47a7203a7a894b1241f9da50"><code>7c25c21</code></a> Release 2024-06-26</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.18...credentials/v1.17.23">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.35.7 to 1.37.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5d5009095c5f8bf340673926d7dac305b2dcec3"><code>a5d5009</code></a> Release 2023-07-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff82299a6ebbc5e5f0a9eddff7776c072f0cb4e8"><code>ff82299</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fc5a86f8fdc4e9af70b87362a6a55c6c6f91826f"><code>fc5a86f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/62ffb94b9079e4d5e154d7c7de1e89971d88b395"><code>62ffb94</code></a> fix v4.go PresignHTTP doc (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2212">#2212</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b4d8d26ccac83a0c8e15b01990a475364ac03e52"><code>b4d8d26</code></a> Release 2023-07-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad521060000fcf57abf457d6d9be1cc0290c0a4f"><code>ad52106</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f335bcf78da9b8d4b4208e5b632441991b2c79c5"><code>f335bcf</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d5a9934e2bb72aa010b157965263b7f89698a5a4"><code>d5a9934</code></a> Release 2023-07-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/247060d3224d694b43e53010f1e6d706d0ed3581"><code>247060d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/61540ee88b45c53f89b9bb99db5c8fc079900bf5"><code>61540ee</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ivs/v1.35.7...service/s3/v1.37.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.4.6 to 4.4.9
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/c962defd33a0901c0b76d5d9e2716d231823fb1f"><code>c962def</code></a> Update bindings to release/v4.4.9</li>
<li><a href="https://github.com/content-services/caliri/commit/81f37b5b99ade4b9b8914e8fa0343c587c070b03"><code>81f37b5</code></a> Merge pull request <a href="https://redirect.github.com/content-services/caliri/issues/2">#2</a> from jlsherrill/122</li>
<li><a href="https://github.com/content-services/caliri/commit/f84bf3ad01ed059ee43146cb836301eba675f3ba"><code>f84bf3a</code></a> Upgrade to go 1.22</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.4.6...release/v4.4.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.6.1718394381 to 2024.6.1719602258
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/feea2f51ca649338a6126e50d771eee9aea851d6"><code>feea2f5</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7eb32ae5ed187e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/aaddb32cbcedec1665e7f232098f20408ce9c04f"><code>aaddb32</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a9725bd58d2a157bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/5d6811ad2354b76dbbcdfb9f98e0ea97aac7a5fe"><code>5d6811a</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a9725ba6da8df57bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/9ef630f50cb2d92e3633f9ce505aac02ef00416b"><code>9ef630f</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d438b6d64c57952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/5b2ea5cefbb228478cc6743a20a5bd92d5dc50ad"><code>5b2ea5c</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a9e88d9ab1b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/3d181b2f66e643d6e54b5cbe9016ae77002f5658"><code>3d181b2</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d43332969f57952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/af7d5895c518871aa667b213ae7ab673a9c97ac6"><code>af7d589</code></a> Update pulp bindings to 386d86254354e29d3b864523aed478253b3a64067968b5e2da9ab...</li>
<li><a href="https://github.com/content-services/zest/commit/97e8656e1bad1cc956411364121b1c6be938455f"><code>97e8656</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/12">#12</a> from jlsherrill/122</li>
<li><a href="https://github.com/content-services/zest/commit/f35fee1f17dd9fb35391a6bc3cbbb8f04d378ee7"><code>f35fee1</code></a> use go 1.22</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.6.1718394381...release/v2024.6.1719602258">compare view</a></li>
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

### Comment by @app-sre-bot on July 01, 2024 at 04:34 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on July 01, 2024 at 07:29 AM UTC

lint go typecheck fail

### Comment by @dependabot on July 01, 2024 at 01:55 PM UTC

Looks like these dependencies are no longer updatable, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/726*
