---
type: pull_request
number: 1318
title: "chore(deps): bump axios from 1.7.9 to 1.8.4"
state: closed
author: dependabot
created: 2025-04-24T18:01:37Z
updated: 2026-01-21T12:42:43Z
closed: 2026-01-21T12:42:34Z
base_branch: master
head_branch: dependabot/npm_and_yarn/axios-1.8.4
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1318
---

# Pull Request #1318: chore(deps): bump axios from 1.7.9 to 1.8.4

**Author**: @dependabot
**Created**: April 24, 2025 at 06:01 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/axios-1.8.4`

## Description

Bumps [axios](https://github.com/axios/axios) from 1.7.9 to 1.8.4.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
<h2>Release v1.8.4</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>buildFullPath:</strong> handle <code>allowAbsoluteUrls: false</code> without <code>baseURL</code> (<a href="https://redirect.github.com/axios/axios/issues/6833">#6833</a>) (<a href="https://github.com/axios/axios/commit/f10c2e0de7fde0051f848609a29c2906d0caa1d9">f10c2e0</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+5/-1 ([#6833](https://github.com/axios/axios/issues/6833) )">Marc Hassan</a></li>
</ul>
<h2>Release v1.8.3</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li>add missing type for allowAbsoluteUrls (<a href="https://redirect.github.com/axios/axios/issues/6818">#6818</a>) (<a href="https://github.com/axios/axios/commit/10fa70ef14fe39558b15a179f0e82f5f5e5d11b2">10fa70e</a>)</li>
<li><strong>xhr/fetch:</strong> pass <code>allowAbsoluteUrls</code> to <code>buildFullPath</code> in <code>xhr</code> and <code>fetch</code> adapters (<a href="https://redirect.github.com/axios/axios/issues/6814">#6814</a>) (<a href="https://github.com/axios/axios/commit/ec159e507bdf08c04ba1a10fe7710094e9e50ec9">ec159e5</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/Electroid" title="+6/-0 ([#6811](https://github.com/axios/axios/issues/6811) )">Ashcon Partovi</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/StefanBRas" title="+4/-0 ([#6818](https://github.com/axios/axios/issues/6818) )">StefanBRas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+2/-2 ([#6814](https://github.com/axios/axios/issues/6814) )">Marc Hassan</a></li>
</ul>
<h2>Release v1.8.2</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>http-adapter:</strong> add allowAbsoluteUrls to path building (<a href="https://redirect.github.com/axios/axios/issues/6810">#6810</a>) (<a href="https://github.com/axios/axios/commit/fb8eec214ce7744b5ca787f2c3b8339b2f54b00f">fb8eec2</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/lexcorp16" title="+1/-1 ([#6810](https://github.com/axios/axios/issues/6810) )">Fasoro-Joseph Alexander</a></li>
</ul>
<h2>Release v1.8.1</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>utils:</strong> move <code>generateString</code> to platform utils to avoid importing crypto module into client builds; (<a href="https://redirect.github.com/axios/axios/issues/6789">#6789</a>) (<a href="https://github.com/axios/axios/commit/36a5a620bec0b181451927f13ac85b9888b86cec">36a5a62</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+51/-47 ([#6789](https://github.com/axios/axios/issues/6789) )">Dmitriy Mozgovoy</a></li>
</ul>
<h2>Release v1.8.0</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>examples:</strong> application crashed when navigating examples in browser (<a href="https://redirect.github.com/axios/axios/issues/5938">#5938</a>) (<a href="https://github.com/axios/axios/commit/1260ded634ec101dd5ed05d3b70f8e8f899dba6c">1260ded</a>)</li>
<li>missing word in SUPPORT_QUESTION.yml (<a href="https://redirect.github.com/axios/axios/issues/6757">#6757</a>) (<a href="https://github.com/axios/axios/commit/1f890b13f2c25a016f3c84ae78efb769f244133e">1f890b1</a>)</li>
<li><strong>utils:</strong> replace getRandomValues with crypto module (<a href="https://redirect.github.com/axios/axios/issues/6788">#6788</a>) (<a href="https://github.com/axios/axios/commit/23a25af0688d1db2c396deb09229d2271cc24f6c">23a25af</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v1.x/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/axios/axios/compare/v1.8.3...v1.8.4">1.8.4</a> (2025-03-19)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>buildFullPath:</strong> handle <code>allowAbsoluteUrls: false</code> without <code>baseURL</code> (<a href="https://redirect.github.com/axios/axios/issues/6833">#6833</a>) (<a href="https://github.com/axios/axios/commit/f10c2e0de7fde0051f848609a29c2906d0caa1d9">f10c2e0</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+5/-1 ([#6833](https://github.com/axios/axios/issues/6833) )">Marc Hassan</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.8.2...v1.8.3">1.8.3</a> (2025-03-10)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>add missing type for allowAbsoluteUrls (<a href="https://redirect.github.com/axios/axios/issues/6818">#6818</a>) (<a href="https://github.com/axios/axios/commit/10fa70ef14fe39558b15a179f0e82f5f5e5d11b2">10fa70e</a>)</li>
<li><strong>xhr/fetch:</strong> pass <code>allowAbsoluteUrls</code> to <code>buildFullPath</code> in <code>xhr</code> and <code>fetch</code> adapters (<a href="https://redirect.github.com/axios/axios/issues/6814">#6814</a>) (<a href="https://github.com/axios/axios/commit/ec159e507bdf08c04ba1a10fe7710094e9e50ec9">ec159e5</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/Electroid" title="+6/-0 ([#6811](https://github.com/axios/axios/issues/6811) )">Ashcon Partovi</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/StefanBRas" title="+4/-0 ([#6818](https://github.com/axios/axios/issues/6818) )">StefanBRas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+2/-2 ([#6814](https://github.com/axios/axios/issues/6814) )">Marc Hassan</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.8.1...v1.8.2">1.8.2</a> (2025-03-07)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>http-adapter:</strong> add allowAbsoluteUrls to path building (<a href="https://redirect.github.com/axios/axios/issues/6810">#6810</a>) (<a href="https://github.com/axios/axios/commit/fb8eec214ce7744b5ca787f2c3b8339b2f54b00f">fb8eec2</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/lexcorp16" title="+1/-1 ([#6810](https://github.com/axios/axios/issues/6810) )">Fasoro-Joseph Alexander</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.8.0...v1.8.1">1.8.1</a> (2025-02-26)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>utils:</strong> move <code>generateString</code> to platform utils to avoid importing crypto module into client builds; (<a href="https://redirect.github.com/axios/axios/issues/6789">#6789</a>) (<a href="https://github.com/axios/axios/commit/36a5a620bec0b181451927f13ac85b9888b86cec">36a5a62</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+51/-47 ([#6789](https://github.com/axios/axios/issues/6789) )">Dmitriy Mozgovoy</a></li>
</ul>
<h1><a href="https://github.com/axios/axios/compare/v1.7.9...v1.8.0">1.8.0</a> (2025-02-25)</h1>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/9f6f97bcfb7c510103dfcb05641ae8e8f5c08bcc"><code>9f6f97b</code></a> chore(release): v1.8.4 (<a href="https://redirect.github.com/axios/axios/issues/6844">#6844</a>)</li>
<li><a href="https://github.com/axios/axios/commit/f10c2e0de7fde0051f848609a29c2906d0caa1d9"><code>f10c2e0</code></a> fix(buildFullPath): handle <code>allowAbsoluteUrls: false</code> without <code>baseURL</code> (<a href="https://redirect.github.com/axios/axios/issues/6833">#6833</a>)</li>
<li><a href="https://github.com/axios/axios/commit/1e6632cd6253c89be7cb1ff757ba91afbd247de1"><code>1e6632c</code></a> chore(deps): bump tj-actions/changed-files in the github-actions group (<a href="https://redirect.github.com/axios/axios/issues/6838">#6838</a>)</li>
<li><a href="https://github.com/axios/axios/commit/39ec206483a89921732bdc8a5be67e350bfc23f0"><code>39ec206</code></a> chore(release): v1.8.3 (<a href="https://redirect.github.com/axios/axios/issues/6819">#6819</a>)</li>
<li><a href="https://github.com/axios/axios/commit/10fa70ef14fe39558b15a179f0e82f5f5e5d11b2"><code>10fa70e</code></a> fix: add missing type for allowAbsoluteUrls (<a href="https://redirect.github.com/axios/axios/issues/6818">#6818</a>)</li>
<li><a href="https://github.com/axios/axios/commit/7821ef9f5be2d62fbc3f01040d9df6f2225eb9d8"><code>7821ef9</code></a> docs: update readme to include bun install (<a href="https://redirect.github.com/axios/axios/issues/6811">#6811</a>)</li>
<li><a href="https://github.com/axios/axios/commit/ec159e507bdf08c04ba1a10fe7710094e9e50ec9"><code>ec159e5</code></a> fix(xhr/fetch): pass <code>allowAbsoluteUrls</code> to <code>buildFullPath</code> in <code>xhr</code> and `fet...</li>
<li><a href="https://github.com/axios/axios/commit/a9f7689b0c4b6d68c7f587c3aa376860da509d94"><code>a9f7689</code></a> chore(release): v1.8.2 (<a href="https://redirect.github.com/axios/axios/issues/6812">#6812</a>)</li>
<li><a href="https://github.com/axios/axios/commit/fb8eec214ce7744b5ca787f2c3b8339b2f54b00f"><code>fb8eec2</code></a> fix(http-adapter): add allowAbsoluteUrls to path building (<a href="https://redirect.github.com/axios/axios/issues/6810">#6810</a>)</li>
<li><a href="https://github.com/axios/axios/commit/98120457559e573024862e2925d56295a965ad7e"><code>9812045</code></a> chore(sponsor): update sponsor block (<a href="https://redirect.github.com/axios/axios/issues/6804">#6804</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v1.7.9...v1.8.4">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=axios&package-manager=npm_and_yarn&previous-version=1.7.9&new-version=1.8.4)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

You can trigger a rebase of this PR by commenting `@dependabot rebase`.

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

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.


---

## Discussion

### Comment by @codecov-commenter on April 24, 2025 at 06:22 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1318?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 100.00%. Comparing base [(`d9468ad`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d9468ad3ce659ea24fd7870e17bb495c71cc28f7?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`a61ff45`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a61ff45c398c4ccfb77332adfd49f3b84033bc12?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##            master     #1318   +/-   ##
=========================================
  Coverage   100.00%   100.00%           
=========================================
  Files            1         1           
  Lines            6         6           
  Branches         2         2           
=========================================
  Hits             6         6           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1318/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1318/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1318/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1318?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @gkarat on May 02, 2025 at 02:41 PM UTC

/retest

### Comment by @dependabot on January 21, 2026 at 12:42 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

## Reviews

### Review by @gkarat - Approved on May 02, 2025 at 02:35 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1318*
