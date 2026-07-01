---
type: pull_request
number: 867
title: "Build: Bump the go group with 2 updates"
state: closed
author: dependabot
created: 2024-10-28T04:15:07Z
updated: 2024-10-30T15:16:24Z
closed: 2024-10-30T15:16:22Z
base_branch: main
head_branch: dependabot/go_modules/go-7a06a70af7
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/867
---

# Pull Request #867: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: October 28, 2024 at 04:15 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-7a06a70af7`

## Description

Bumps the go group with 2 updates: [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.41.2 to 1.42.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ee5e3f05637540596cc7aab1359742000a8d533a"><code>ee5e3f0</code></a> Release 2023-11-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b65c226f47aa1f837699664bdc65c3c3e3611765"><code>b65c226</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7a194b9b0344774a5af100d11ea2066c5b0cf234"><code>7a194b9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cb924a0007bc681d12f382a604368e0660827ee"><code>0cb924a</code></a> Add support for configured endpoints. (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2328">#2328</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/61039fea9cc9e080c53382850c87685b5406fd68"><code>61039fe</code></a> Release 2023-10-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/797e0560769725635218fc30a2554c1bbaccc01b"><code>797e056</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/822585d3f621a7c5844584d8e471c32f852702aa"><code>822585d</code></a> Update SDK's smithy-go dependency to v1.16.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/abf753db747dd256f3ee69712a19d1d3dc681f23"><code>abf753d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/99861c071109ce5ee4f1cb3b72ead2062b3bd86c"><code>99861c0</code></a> lang: bump minimum go version to 1.19 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2338">#2338</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ac0a53ac45acaadc4526fd25b643dc46032b02a"><code>2ac0a53</code></a> Release 2023-10-30</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/fsx/v1.41.2...service/s3/v1.42.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.10.1729249374 to 2024.10.1729885990
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/c495fcc7fabe307995e45cc6b920e8fe4940be8b"><code>c495fcc</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b979363d6d52137ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/eca0571eb6e9b242ef5af0d59bd42565b7ad27dd"><code>eca0571</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97936bd49d6037ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/759f989555170f18f57d4949f66a5e44a518ecc6"><code>759f989</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7ea35b59b6087e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/fc424c98ef4f2425854308c30bbbbbf57cae9bb7"><code>fc424c9</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7ea332b2a8087e8639db952b3...</li>
<li><a href="https://github.com/content-services/zest/commit/d4f6587e3d97bb16e44c6ce8df2af2e213eec000"><code>d4f6587</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e475a3824b4e1b7a98d596b8e32...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.10.1729249374...release/v2024.10.1729885990">compare view</a></li>
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

### Comment by @app-sre-bot on October 28, 2024 at 04:15 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on October 28, 2024 at 08:59 AM UTC

`make: *** [mk/go-rules.mk:49: test-unit] Error 1`

### Comment by @swadeley on October 28, 2024 at 09:15 AM UTC

[test]

### Comment by @rverdile on October 30, 2024 at 03:15 PM UTC

@dependabot rebase

### Comment by @dependabot on October 30, 2024 at 03:16 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/867*
