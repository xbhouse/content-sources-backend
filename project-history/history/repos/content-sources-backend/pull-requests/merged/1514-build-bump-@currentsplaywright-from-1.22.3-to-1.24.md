---
type: pull_request
number: 1514
title: "Build: Bump @currents/playwright from 1.22.3 to 1.24.0"
state: merged
author: dependabot
created: 2026-06-04T07:52:03Z
updated: 2026-06-04T09:27:59Z
closed: 2026-06-04T09:27:57Z
merged: 2026-06-04T09:27:57Z
base_branch: main
head_branch: dependabot/npm_and_yarn/currents/playwright-1.24.0
labels: ["dependencies", "javascript"]
url: https://github.com/content-services/content-sources-backend/pull/1514
---

# Pull Request #1514: Build: Bump @currents/playwright from 1.22.3 to 1.24.0

**Author**: @dependabot
**Created**: June 04, 2026 at 07:52 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `javascript`
**Base**: `main` ← **Head**: `dependabot/npm_and_yarn/currents/playwright-1.24.0`

## Description

Bumps [@currents/playwright](https://github.com/currents-dev/currents-playwright-changelog) from 1.22.3 to 1.24.0.
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/currents-dev/currents-playwright-changelog/blob/main/CHANGELOG.md">@​currents/playwright's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.4...v2.2.0">2.2.0</a> (2026-06-02)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>bump jiti from 2.6.1 to 2.7.0 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/863">#863</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/7c074b241f6982a57f146ab85ae1d730bd5e1987">7c074b2</a>)</li>
<li>bump tmp from 0.2.6 to 0.2.7 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/865">#865</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/293915d1a47a526ffe06ac95dc341eb59afe8be7">293915d</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>use unidici and drop axios [ENG-568] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/836">#836</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/a3d7f72df8ebd7e851687cce6f72b668b0867a26">a3d7f72</a>), closes <a href="https://redirect.github.com/currents-dev/currents-playwright/issues/813">#813</a></li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.3...v2.1.4">2.1.4</a> (2026-05-29)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>restore git-preferred commit merge and clean CI branch names [ENG-563] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/859">#859</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/a6610719e7c0e807a4677add101284b5ec6e029b">a661071</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.2...v2.1.3">2.1.3</a> (2026-05-29)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>adopt CURRENTS_CI_URL as ciUrl config option [ENG-561] (<a href="https://github.com/currents-dev/currents-playwright/commit/ef9326b3dd3728834d3bb86422ee915afd45aa6a">ef9326b</a>)</li>
<li>bump tmp from 0.2.5 to 0.2.6 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/853">#853</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/73a765d42c7599c251289457753d5689075405b6">73a765d</a>)</li>
<li>bump tmp from 0.2.5 to 0.2.7 in /examples (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/852">#852</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/8beb700ca00b087bdee1882842965363373765d1">8beb700</a>)</li>
<li>bump tmp from 0.2.5 to 0.2.7 in /examples/imported-tests (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/855">#855</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/ba99ef2cb6b415177e76f708fd128f16dc7d1910">ba99ef2</a>)</li>
<li>update dependencies in package-lock.json and package.json (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/858">#858</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/52955430f2c09b0f59d82ac81907f1f300370bd7">5295543</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.1...v2.1.2">2.1.2</a> (2026-05-26)</h2>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.0...v2.1.1">2.1.1</a> (2026-05-22)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>include the fully parallel flag into the accepted flags for pwc-p run. (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/840">#840</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/34615d4da4151d33c27e2efda0dfb3d5a9ec42a8">34615d4</a>)</li>
</ul>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v2.0.0...v2.1.0">2.1.0</a> (2026-05-22)</h1>
<h3>Features</h3>
<ul>
<li>add --pwc-environment option and discovery environment field [CSR-3613] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/833">#833</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/3164c335e3bb82db1794a65a8223442b5dae4718">3164c33</a>)</li>
<li>azure ci data [CSR-4184] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/816">#816</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/d452c276f53b21fc7b920ae199d10a27e96a4fc4">d452c27</a>)</li>
</ul>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v1.23.2...v2.0.0">2.0.0</a> (2026-05-19)</h1>
<h3>Breaking changes</h3>
<p>Breaking changes for orchestration (<code>pwc-p</code>) only.
See the upgrade guide at <a href="https://docs.currents.dev/resources/reporters/currents-playwright/migration-to-playwright-1.60">https://docs.currents.dev/resources/reporters/currents-playwright/migration-to-playwright-1.60</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/currents-dev/currents-playwright-changelog/commits">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=@currents/playwright&package-manager=npm_and_yarn&previous-version=1.22.3&new-version=1.24.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)


</details>

---

## Reviews

### Review by @TenSt - Approved on June 04, 2026 at 09:27 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1514*
