---
type: pull_request
number: 1472
title: "chore(deps): bump tar and @currents/playwright"
state: closed
author: dependabot
created: 2026-01-16T21:24:22Z
updated: 2026-01-21T01:24:06Z
closed: 2026-01-21T01:24:04Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-bf7315278b
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1472
---

# Pull Request #1472: chore(deps): bump tar and @currents/playwright

**Author**: @dependabot
**Created**: January 16, 2026 at 09:24 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-bf7315278b`

## Description

Removes [tar](https://github.com/isaacs/node-tar). It's no longer used after updating ancestor dependency [@currents/playwright](https://github.com/currents-dev/currents-playwright-changelog). These dependencies need to be updated together.

Removes `tar`

Updates `@currents/playwright` from 1.19.1 to 1.20.2
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/currents-dev/currents-playwright-changelog/blob/main/CHANGELOG.md"><code>@​currents/playwright</code>'s changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v1.20.1...v1.20.2">1.20.2</a> (2026-01-13)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>Add additional debug logs to identify orchestration discovery timeouts (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/687">#687</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/853d026c58fdb87d1820ba301c3776e9a6494625">853d026</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v1.20.0...v1.20.1">1.20.1</a> (2025-12-29)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>bump eslint from 9.35.0 to 9.39.2 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/667">#667</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/625010d4c12ba2e40cea93b5fe1ba43a240e63e0">625010d</a>)</li>
<li>bump next from 15.4.8 to 15.4.9 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/665">#665</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/54be510ad801f16453066b9a2d0f96d7e9f8948c">54be510</a>)</li>
<li>bump next from 15.4.9 to 15.4.10 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/672">#672</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/eedc32f5e8f64e8109ec9a47127e97ab61a9dd81">eedc32f</a>)</li>
<li>handle orchestration cancellation without hanging [CSR-3540] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/673">#673</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/ecf81f46ee5cb64f9c82c90bcb48b9abf50f6d5a">ecf81f4</a>)</li>
</ul>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v1.17.0...v1.20.0">1.20.0</a> (2025-12-08)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>bump axios from 1.12.2 to 1.13.2 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/656">#656</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/deaa5c2479d02a59efe6301343e5036095d323d0">deaa5c2</a>)</li>
<li>bump dotenv from 17.2.2 to 17.2.3 (<a href="https://github.com/currents-dev/currents-playwright/commit/7edbe4fa5909ca7b4bb1da250d89bb266a4e35bc">7edbe4f</a>)</li>
<li>Switch from c12 to using jiti directly for config loading (<a href="https://github.com/currents-dev/currents-playwright/commit/a830bc63e690074582cdb943fa138454a8b1f169">a830bc6</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v1.19.1...v1.19.2">1.19.2</a> (2025-12-03)</h2>
<h3>Bug Fixes</h3>
<p>fix: set dotenv import as quiet to remove log
fix: downgrade c12 to reduce config loading timing issues</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/currents-dev/currents-playwright-changelog/commits">compare view</a></li>
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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @dependabot on January 21, 2026 at 01:24 AM UTC

Superseded by #1477.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1472*
