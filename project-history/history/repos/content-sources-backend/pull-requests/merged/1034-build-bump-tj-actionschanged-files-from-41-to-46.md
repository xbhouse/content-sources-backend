---
type: pull_request
number: 1034
title: "Build: Bump tj-actions/changed-files from 41 to 46"
state: merged
author: dependabot
created: 2025-03-18T08:24:52Z
updated: 2025-03-18T16:00:49Z
closed: 2025-03-18T16:00:39Z
merged: 2025-03-18T16:00:39Z
base_branch: main
head_branch: dependabot/github_actions/dot-github/workflows/tj-actions/changed-files-46
labels: ["dependencies", "github_actions"]
url: https://github.com/content-services/content-sources-backend/pull/1034
---

# Pull Request #1034: Build: Bump tj-actions/changed-files from 41 to 46

**Author**: @dependabot
**Created**: March 18, 2025 at 08:24 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `github_actions`
**Base**: `main` ← **Head**: `dependabot/github_actions/dot-github/workflows/tj-actions/changed-files-46`

## Description

Bumps [tj-actions/changed-files](https://github.com/tj-actions/changed-files) from 41 to 46.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/tj-actions/changed-files/releases">tj-actions/changed-files's releases</a>.</em></p>
<blockquote>
<h2>v46</h2>
<blockquote>
<p>[!WARNING]<br />
<strong>Security Alert:</strong> A critical security issue was identified in this action due to a compromised commit.</p>
<p>This commit has been <strong>removed</strong> from all tags and branches, and necessary measures have been implemented to prevent similar issues in the future.</p>
<h4><strong>Action Required:</strong></h4>
<ul>
<li><strong>Review your workflows executed between March 14 and March 15.</strong> If you notice unexpected output under the <code>changed-files</code> section, decode  it using the following command:  <code>echo 'xxx' | base64 -d | base64 -d</code><br />
If the output contains sensitive information (e.g., tokens or secrets), <strong>revoke and rotate those secrets immediately</strong>.</li>
<li><strong>If your workflows reference this commit directly by its SHA</strong>, you must update them immediately to avoid using the compromised version.</li>
<li><strong>If you are using tagged versions</strong> (e.g., <code>v35</code>, <code>v44.5.1</code>), no action is required as these tags have been updated and are now safe to use.</li>
</ul>
<p>Additionally, as a precaution, we recommend rotating any secrets that may have been exposed during this timeframe to ensure the continued security of your workflows.</p>
</blockquote>
<h1>Changes in v46.0.1</h1>
<h2>What's Changed</h2>
<ul>
<li>update: sync-release-version.yml to use signed commits by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2472">tj-actions/changed-files#2472</a></li>
<li>Updated README.md by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2473">tj-actions/changed-files#2473</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/tj-actions/changed-files/compare/v46...v46.0.1">https://github.com/tj-actions/changed-files/compare/v46...v46.0.1</a></p>
<hr />
<h1>Changes in v46.0.0</h1>
<h2>What's Changed</h2>
<ul>
<li>docs: update docs to highlight security issues by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2465">tj-actions/changed-files#2465</a></li>
<li>fix: update github workflow update-readme.yml by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2466">tj-actions/changed-files#2466</a></li>
<li>fix: update permission in update-readme.yml workflow by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2467">tj-actions/changed-files#2467</a></li>
<li>fix: update update-readme.yml to sign-commits by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2468">tj-actions/changed-files#2468</a></li>
<li>Updated README.md by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2469">tj-actions/changed-files#2469</a></li>
<li>update: sync-release-version.yml by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2471">tj-actions/changed-files#2471</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/github-actions"><code>@​github-actions</code></a> made their first contribution in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2469">tj-actions/changed-files#2469</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/tj-actions/changed-files/compare/v45.0.5...v46.0.0">https://github.com/tj-actions/changed-files/compare/v45.0.5...v46.0.0</a></p>
<h2>What's Changed</h2>
<ul>
<li>docs: update docs to highlight security issues by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2465">tj-actions/changed-files#2465</a></li>
<li>fix: update github workflow update-readme.yml by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2466">tj-actions/changed-files#2466</a></li>
<li>fix: update permission in update-readme.yml workflow by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2467">tj-actions/changed-files#2467</a></li>
<li>fix: update update-readme.yml to sign-commits by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2468">tj-actions/changed-files#2468</a></li>
<li>Updated README.md by <a href="https://github.com/github-actions"><code>@​github-actions</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2469">tj-actions/changed-files#2469</a></li>
<li>update: sync-release-version.yml by <a href="https://github.com/jackton1"><code>@​jackton1</code></a> in <a href="https://redirect.github.com/tj-actions/changed-files/pull/2471">tj-actions/changed-files#2471</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/tj-actions/changed-files/compare/v45.0.5...v46.0.0">https://github.com/tj-actions/changed-files/compare/v45.0.5...v46.0.0</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/tj-actions/changed-files/blob/main/HISTORY.md">tj-actions/changed-files's changelog</a>.</em></p>
<blockquote>
<h1>Changelog</h1>
<h1><a href="https://github.com/tj-actions/changed-files/compare/v46.0.0...v46.0.1">46.0.1</a> - (2025-03-16)</h1>
<h2><!-- raw HTML omitted -->🔄 Update</h2>
<ul>
<li>Updated README.md (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2473">#2473</a>)</li>
</ul>
<p>Co-authored-by: github-actions[bot] <!-- raw HTML omitted --> (<a href="https://github.com/tj-actions/changed-files/commit/2f7c5bfce28377bc069a65ba478de0a74aa0ca32">2f7c5bf</a>)  - (github-actions[bot])</p>
<ul>
<li>Sync-release-version.yml to use signed commits (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2472">#2472</a>) (<a href="https://github.com/tj-actions/changed-files/commit/4189ec62c445484531e9ad97157d990be96e88ee">4189ec6</a>)  - (Tonye Jack)</li>
</ul>
<h1><a href="https://github.com/tj-actions/changed-files/compare/v45.0.9...v46.0.0">46.0.0</a> - (2025-03-16)</h1>
<h2><!-- raw HTML omitted -->🐛 Bug Fixes</h2>
<ul>
<li>Update update-readme.yml to sign-commits (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2468">#2468</a>) (<a href="https://github.com/tj-actions/changed-files/commit/0f1ffe61855cb317d5fd66122c14dc0627eab141">0f1ffe6</a>)  - (Tonye Jack)</li>
<li>Update permission in update-readme.yml workflow (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2467">#2467</a>) (<a href="https://github.com/tj-actions/changed-files/commit/ddef03e37c84cfb9ee89fa055b86359aaf949c86">ddef03e</a>)  - (Tonye Jack)</li>
<li>Update github workflow update-readme.yml (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2466">#2466</a>) (<a href="https://github.com/tj-actions/changed-files/commit/9c2df0d54a911c819d7368d7e5ed7c01c0796e0a">9c2df0d</a>)  - (Tonye Jack)</li>
</ul>
<h2><!-- raw HTML omitted -->➖ Remove</h2>
<ul>
<li>Deleted renovate.json (<a href="https://github.com/tj-actions/changed-files/commit/e37e952786556966c1fb6183c5937b3966bab099">e37e952</a>)  - (Tonye Jack)</li>
</ul>
<h2><!-- raw HTML omitted -->🔄 Update</h2>
<ul>
<li>Sync-release-version.yml (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2471">#2471</a>) (<a href="https://github.com/tj-actions/changed-files/commit/4cd184a1dd542b79cca1d4d7938e4154a6520ca7">4cd184a</a>)  - (Tonye Jack)</li>
<li>Updated README.md (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2469">#2469</a>)</li>
</ul>
<p>Co-authored-by: github-actions[bot] <!-- raw HTML omitted --> (<a href="https://github.com/tj-actions/changed-files/commit/5cbf22026d05fbef0c027d1b1f118fe3a1b6e435">5cbf220</a>)  - (github-actions[bot])</p>
<h2><!-- raw HTML omitted -->📚 Documentation</h2>
<ul>
<li>Update docs to highlight security issues (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2465">#2465</a>) (<a href="https://github.com/tj-actions/changed-files/commit/65253327cf47481b4b1b4b9fea78e143a1353147">6525332</a>)  - (Tonye Jack)</li>
</ul>
<h1><a href="https://github.com/tj-actions/changed-files/compare/v45.0.4...v45.0.9">45.0.9</a> - (2025-03-15)</h1>
<h2><!-- raw HTML omitted -->🐛 Bug Fixes</h2>
<ul>
<li><strong>deps:</strong> Update dependency <code>@​octokit/rest</code> to v21.1.1 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2435">#2435</a>) (<a href="https://github.com/tj-actions/changed-files/commit/fb8dcda5fb8954cec37773d2b275a8579c86c781">fb8dcda</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update dependency <code>@​octokit/rest</code> to v21.1.0 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2394">#2394</a>) (<a href="https://github.com/tj-actions/changed-files/commit/7b72c97d739f955f5cadca0d59799d826ae9f6c9">7b72c97</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update dependency yaml to v2.7.0 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2383">#2383</a>) (<a href="https://github.com/tj-actions/changed-files/commit/5f974c28f5044c411f0c9e7becf3f172029cf9cf">5f974c2</a>)  - (renovate[bot])</li>
</ul>
<h2><!-- raw HTML omitted -->⚙️ Miscellaneous Tasks</h2>
<ul>
<li><strong>deps:</strong> Lock file maintenance (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2460">#2460</a>) (<a href="https://github.com/tj-actions/changed-files/commit/9200e69727eb73eb060652b19946b8a2fdfb654b">9200e69</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update dependency <code>@​types/node</code> to v22.13.10 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2459">#2459</a>) (<a href="https://github.com/tj-actions/changed-files/commit/e650cfdae513481a20f538e88d98b39106523006">e650cfd</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update dependency eslint-config-prettier to v10.1.1 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2458">#2458</a>) (<a href="https://github.com/tj-actions/changed-files/commit/82af21f4a05896ca18c950539469bee225c45a89">82af21f</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update dependency eslint-config-prettier to v10.1.0 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2457">#2457</a>) (<a href="https://github.com/tj-actions/changed-files/commit/82fa4a6402582d5c8c9c0e95b7ff7cc88992bbb4">82fa4a6</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update peter-evans/create-pull-request action to v7.0.8 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2455">#2455</a>) (<a href="https://github.com/tj-actions/changed-files/commit/315505acf41d2913b71af48080fb158cd01f79e7">315505a</a>)  - (renovate[bot])</li>
<li><strong>deps:</strong> Update dependency <code>@​types/node</code> to v22.13.9 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2454">#2454</a>) (<a href="https://github.com/tj-actions/changed-files/commit/c8e1cdb9ea135ee549963c167ffaec5e7d4a71cd">c8e1cdb</a>)  - (renovate[bot])</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/tj-actions/changed-files/commit/2f7c5bfce28377bc069a65ba478de0a74aa0ca32"><code>2f7c5bf</code></a> Updated README.md (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2473">#2473</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/4189ec62c445484531e9ad97157d990be96e88ee"><code>4189ec6</code></a> update: sync-release-version.yml to use signed commits (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2472">#2472</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/4cd184a1dd542b79cca1d4d7938e4154a6520ca7"><code>4cd184a</code></a> update: sync-release-version.yml (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2471">#2471</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/5cbf22026d05fbef0c027d1b1f118fe3a1b6e435"><code>5cbf220</code></a> Updated README.md (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2469">#2469</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/0f1ffe61855cb317d5fd66122c14dc0627eab141"><code>0f1ffe6</code></a> fix: update update-readme.yml to sign-commits (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2468">#2468</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/ddef03e37c84cfb9ee89fa055b86359aaf949c86"><code>ddef03e</code></a> fix: update permission in update-readme.yml workflow (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2467">#2467</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/9c2df0d54a911c819d7368d7e5ed7c01c0796e0a"><code>9c2df0d</code></a> fix: update github workflow update-readme.yml (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2466">#2466</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/65253327cf47481b4b1b4b9fea78e143a1353147"><code>6525332</code></a> docs: update docs to highlight security issues (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2465">#2465</a>)</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/e37e952786556966c1fb6183c5937b3966bab099"><code>e37e952</code></a> Deleted renovate.json</li>
<li><a href="https://github.com/tj-actions/changed-files/commit/a284dc1814e3fd07f2e34267fc8f81227ed29fb8"><code>a284dc1</code></a> Upgraded to v45.0.8 (<a href="https://redirect.github.com/tj-actions/changed-files/issues/2462">#2462</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/tj-actions/changed-files/compare/v41...v46">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=tj-actions/changed-files&package-manager=github_actions&previous-version=41&new-version=46)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/content-services/content-sources-backend/network/alerts).

</details>

---

## Reviews

### Review by @rverdile - Approved on March 18, 2025 at 03:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1034*
