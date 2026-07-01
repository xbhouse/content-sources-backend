---
type: pull_request
number: 1565
title: "chore(deps): bump handlebars from 4.7.8 to 4.7.9"
state: merged
author: dependabot
created: 2026-03-30T13:45:02Z
updated: 2026-04-01T12:41:05Z
closed: 2026-04-01T12:40:56Z
merged: 2026-04-01T12:40:56Z
base_branch: master
head_branch: dependabot/npm_and_yarn/handlebars-4.7.9
labels: ["dependencies", "patch"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1565
---

# Pull Request #1565: chore(deps): bump handlebars from 4.7.8 to 4.7.9

**Author**: @dependabot
**Created**: March 30, 2026 at 01:45 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `patch`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/handlebars-4.7.9`

## Description

Bumps [handlebars](https://github.com/handlebars-lang/handlebars.js) from 4.7.8 to 4.7.9.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/handlebars-lang/handlebars.js/releases">handlebars's releases</a>.</em></p>
<blockquote>
<h2>v4.7.9</h2>
<ul>
<li>fix: enable shell mode for spawn to resolve Windows EINVAL issue - e0137c2</li>
<li>fix type &quot;RuntimeOptions&quot; also accepting string partials - eab1d14</li>
<li>feat(types): set <code>hash</code> to be a <code>Record&lt;string, any&gt;</code> - de4414d</li>
<li>fix non-contiguous program indices - 4512766</li>
<li>refactor: rename i to startPartIndex - e497a35</li>
<li>security: fix security issues - 68d8df5
<ul>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-2w6w-674q-4c4q">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-2w6w-674q-4c4q</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-3mfm-83xf-c92r">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-3mfm-83xf-c92r</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-xhpv-hc6g-r9c6">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-xhpv-hc6g-r9c6</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-xjpj-3mr7-gcpf">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-xjpj-3mr7-gcpf</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-9cx6-37pm-9jff">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-9cx6-37pm-9jff</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-2qvq-rjwj-gvw9">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-2qvq-rjwj-gvw9</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-7rx3-28cr-v5wh">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-7rx3-28cr-v5wh</a></li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-442j-39wm-28r2">https://github.com/handlebars-lang/handlebars.js/security/advisories/GHSA-442j-39wm-28r2</a></li>
</ul>
</li>
</ul>
<p><a href="https://github.com/handlebars-lang/handlebars.js/compare/v4.7.8...v4.7.9">Commits</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/handlebars-lang/handlebars.js/blob/v4.7.9/release-notes.md">handlebars's changelog</a>.</em></p>
<blockquote>
<h2>v4.7.9 - March 26th, 2026</h2>
<ul>
<li>fix: enable shell mode for spawn to resolve Windows EINVAL issue - e0137c2</li>
<li>fix type &quot;RuntimeOptions&quot; also accepting string partials - eab1d14</li>
<li>feat(types): set <code>hash</code> to be a <code>Record&lt;string, any&gt;</code> - de4414d</li>
<li>fix non-contiguous program indices - 4512766</li>
<li>refactor: rename i to startPartIndex - e497a35</li>
<li>security: fix security issues - 68d8df5</li>
</ul>
<p><a href="https://github.com/handlebars-lang/handlebars.js/compare/v4.7.8...v4.7.9">Commits</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/dce542c9a660048d31f0981ac8a45c08b919bddb"><code>dce542c</code></a> v4.7.9</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/8a41389ba5b2624b6f43a5463d8e2533b843a562"><code>8a41389</code></a> Update release notes</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/68d8df5a88e0a26fe9e6084c5c6aaebe67b07da2"><code>68d8df5</code></a> Fix security issues</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/b2a083136b11e1da9f0f47a11f749a9830a49328"><code>b2a0831</code></a> Fix browser tests</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/9f98c1629834abf8de5a127caff8a2eab03d2c12"><code>9f98c16</code></a> Fix release script</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/45443b4290475dfb7cec32a85d344f12ab345eb9"><code>45443b4</code></a> Revert &quot;Improve partial indenting performance&quot;</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/8841a5f6d35096aee95d68e1e49636a4cb5c661e"><code>8841a5f</code></a> Fix CI errors with linting</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/e0137c26f2202593bca7cc25184e733e87d54709"><code>e0137c2</code></a> fix: enable shell mode for spawn to resolve Windows EINVAL issue</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/e914d6037ffb0dd371f7e4823cdb019732ae66d7"><code>e914d60</code></a> Improve rendering performance</li>
<li><a href="https://github.com/handlebars-lang/handlebars.js/commit/7de4b41c344a5d702edca93d1841b59642fa32bd"><code>7de4b41</code></a> Upgrade GitHub Actions checkout and setup-node on 4.x branch</li>
<li>Additional commits viewable in <a href="https://github.com/handlebars-lang/handlebars.js/compare/v4.7.8...v4.7.9">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on March 30, 2026 at 01:48 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1565?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`a08d0d6`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/a08d0d69deb64ea8523582130fbabff1dde6be51?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`478b64a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/478b64a705ebcc4bbb693ac2c7dcd0981c485be6?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1565   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      686      686           
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1565?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @dominikvagner - Approved on April 01, 2026 at 12:40 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1565*
