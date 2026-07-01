---
type: pull_request
number: 448
title: "Build: Bump gorm.io/driver/postgres from 1.5.3 to 1.5.4"
state: closed
author: dependabot
created: 2023-10-30T04:26:14Z
updated: 2023-10-30T13:00:21Z
closed: 2023-10-30T13:00:16Z
base_branch: main
head_branch: dependabot/go_modules/gorm.io/driver/postgres-1.5.4
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/448
---

# Pull Request #448: Build: Bump gorm.io/driver/postgres from 1.5.3 to 1.5.4

**Author**: @dependabot
**Created**: October 30, 2023 at 04:26 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/gorm.io/driver/postgres-1.5.4`

## Description

[//]: # (dependabot-start)
⚠️  **Dependabot is rebasing this PR** ⚠️ 

Rebasing might not happen immediately, so don't worry if this takes some time.

Note: if you make any changes to this PR yourself, they will take precedence over the rebase.

---

[//]: # (dependabot-end)

Bumps [gorm.io/driver/postgres](https://github.com/go-gorm/postgres) from 1.5.3 to 1.5.4.
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/postgres/commit/68a877c11af52552b7eb5e061a66240a0f6508cb"><code>68a877c</code></a> Update go.mod</li>
<li><a href="https://github.com/go-gorm/postgres/commit/a7214f0c523c2d4f209198fa2cca543fac84d712"><code>a7214f0</code></a> allow ignoring escape (<a href="https://redirect.github.com/go-gorm/postgres/issues/221">#221</a>)</li>
<li><a href="https://github.com/go-gorm/postgres/commit/76303659a1acc223f15afd9dac2c9cedc01bb237"><code>7630365</code></a> fix: change DefaultValueValue to the same as other drivers (<a href="https://redirect.github.com/go-gorm/postgres/issues/222">#222</a>)</li>
<li>See full diff in <a href="https://github.com/go-gorm/postgres/compare/v1.5.3...v1.5.4">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=gorm.io/driver/postgres&package-manager=go_modules&previous-version=1.5.3&new-version=1.5.4)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @app-sre-bot on October 30, 2023 at 04:27 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 30, 2023 at 10:35 AM UTC

[test]

### Comment by @dependabot on October 30, 2023 at 01:00 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/448*
