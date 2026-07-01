---
type: pull_request
number: 771
title: "Build: Bump the go group with 2 updates"
state: merged
author: dependabot
created: 2024-08-12T04:50:38Z
updated: 2024-08-12T15:26:15Z
closed: 2024-08-12T15:26:07Z
merged: 2024-08-12T15:26:07Z
base_branch: main
head_branch: dependabot/go_modules/go-0545c85e51
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/771
---

# Pull Request #771: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: August 12, 2024 at 04:50 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-0545c85e51`

## Description

Bumps the go group with 2 updates: [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) and [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils).

Updates `github.com/content-services/zest/release/v2024` from 2024.8.1722539431 to 2024.8.1723121474
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/b646509e0d6939a7198546a8ac57943e4fc00328"><code>b646509</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84745de39864037d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/e94d7a05cb53c6423ed487ae52052e5717f4af7f"><code>e94d7a0</code></a> Merge pull request <a href="https://redirect.github.com/content-services/zest/issues/16">#16</a> from jlsherrill/fix_Deploy</li>
<li><a href="https://github.com/content-services/zest/commit/2e1c84bd3d12ef4aa0e4596d1ff4b4622315c9df"><code>2e1c84b</code></a> Fix compose up in gh action</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.8.1722539431...release/v2024.8.1723121474">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.5 to 1.25.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/a8198f5370b3a1a0163f9031ce158a2a5fa7bfef"><code>a8198f5</code></a> use Warn for situations that are not actual errors (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/569">#569</a>)</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.5...v1.25.6">compare view</a></li>
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

### Comment by @app-sre-bot on August 12, 2024 at 04:51 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on August 12, 2024 at 07:19 AM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on August 12, 2024 at 03:26 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/771*
