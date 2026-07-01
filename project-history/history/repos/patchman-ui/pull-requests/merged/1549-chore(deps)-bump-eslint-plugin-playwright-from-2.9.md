---
type: pull_request
number: 1549
title: "chore(deps): bump eslint-plugin-playwright from 2.9.0 to 2.10.0 in the lint group"
state: merged
author: dependabot
created: 2026-03-17T18:47:27Z
updated: 2026-03-18T13:01:38Z
closed: 2026-03-18T13:01:30Z
merged: 2026-03-18T13:01:30Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-994f3c3974
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1549
---

# Pull Request #1549: chore(deps): bump eslint-plugin-playwright from 2.9.0 to 2.10.0 in the lint group

**Author**: @dependabot
**Created**: March 17, 2026 at 06:47 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-994f3c3974`

## Description

Bumps the lint group with 1 update: [eslint-plugin-playwright](https://github.com/mskelton/eslint-plugin-playwright).

Updates `eslint-plugin-playwright` from 2.9.0 to 2.10.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/mskelton/eslint-plugin-playwright/releases">eslint-plugin-playwright's releases</a>.</em></p>
<blockquote>
<h2>v2.10.0</h2>
<h1><a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.9.0...v2.10.0">2.10.0</a> (2026-03-14)</h1>
<h3>Bug Fixes</h3>
<ul>
<li><strong>missing-playwright-await:</strong> Fix false positive with <code>expect().resolves</code> (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/352e15e0e28cda5c7f7fbcd5bd6d01cf634aea3e">352e15e</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/448">#448</a></li>
<li>Support additional promise methods (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8646e62527202cf11da6c00afc7f7e376d00773f">8646e62</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/444">#444</a></li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>missing-playwright-await:</strong> Add <code>includePageLocatorMethods</code> flag for checking more missing awaits (<a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/438">#438</a>) (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/41921f8509bfa90ccef91d86ed874408b60a7abb">41921f8</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/159">#159</a></li>
<li><strong>no-skipped-test:</strong> Support <code>disallowFixme</code> to optionally disable <code>.fixme()</code> annotations (<a href="https://github.com/mskelton/eslint-plugin-playwright/commit/6b42fdb5cf74c6a98b7544e2931bd157cda88e51">6b42fdb</a>), closes <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/446">#446</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/41921f8509bfa90ccef91d86ed874408b60a7abb"><code>41921f8</code></a> feat(missing-playwright-await): Add <code>includePageLocatorMethods</code> flag for chec...</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/2a05e2d6831fc89b2d262f7338c7041b50f83eca"><code>2a05e2d</code></a> chore: Remove name from unit tests</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/352e15e0e28cda5c7f7fbcd5bd6d01cf634aea3e"><code>352e15e</code></a> fix(missing-playwright-await): Fix false positive with <code>expect().resolves</code></li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/6b42fdb5cf74c6a98b7544e2931bd157cda88e51"><code>6b42fdb</code></a> feat(no-skipped-test): Support <code>disallowFixme</code> to optionally disable `.fixme(...</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/080c3f87faa08d2f2c48cc00435f81c961e92e12"><code>080c3f8</code></a> Merge pull request <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/447">#447</a> from mskelton/dependabot/npm_and_yarn/tar-7.5.11</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/21d944621070b9ab83a73f6db914c4f4b2ad7afb"><code>21d9446</code></a> chore(deps): Bump tar from 7.5.10 to 7.5.11</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/8646e62527202cf11da6c00afc7f7e376d00773f"><code>8646e62</code></a> fix: Support additional promise methods</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/306d21e8c32bb26357475842aa3942c120b7f523"><code>306d21e</code></a> Merge pull request <a href="https://redirect.github.com/mskelton/eslint-plugin-playwright/issues/442">#442</a> from mskelton/dependabot/npm_and_yarn/tar-7.5.10</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/981d42e32e392bb1579eabd361a6ff539406586b"><code>981d42e</code></a> chore(deps): Bump tar from 7.5.9 to 7.5.10</li>
<li><a href="https://github.com/mskelton/eslint-plugin-playwright/commit/44b78a9b74aeb90936ed83355e8dc833045749a3"><code>44b78a9</code></a> chore: More tests</li>
<li>See full diff in <a href="https://github.com/mskelton/eslint-plugin-playwright/compare/v2.9.0...v2.10.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=eslint-plugin-playwright&package-manager=npm_and_yarn&previous-version=2.9.0&new-version=2.10.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

## Discussion

### Comment by @codecov-commenter on March 17, 2026 at 06:50 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1549?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 72.34%. Comparing base ([`ccc7f0e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ccc7f0e96fe9300cec6ca6644b09384af2b02530?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`42c00bd`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/42c00bd949af6c0ec98b9e7274b61506613a538a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1549      +/-   ##
==========================================
+ Coverage   72.30%   72.34%   +0.03%     
==========================================
  Files          99       99              
  Lines        2405     2408       +3     
  Branches      677      683       +6     
==========================================
+ Hits         1739     1742       +3     
  Misses        586      586              
  Partials       80       80              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1549?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on March 18, 2026 at 01:01 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1549*
