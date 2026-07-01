---
type: pull_request
number: 1319
title: "chore(deps): bump @babel/helpers from 7.26.7 to 7.27.0"
state: closed
author: dependabot
created: 2025-04-24T18:02:03Z
updated: 2026-01-21T12:42:44Z
closed: 2026-01-21T12:42:35Z
base_branch: master
head_branch: dependabot/npm_and_yarn/babel/helpers-7.27.0
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1319
---

# Pull Request #1319: chore(deps): bump @babel/helpers from 7.26.7 to 7.27.0

**Author**: @dependabot
**Created**: April 24, 2025 at 06:02 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/babel/helpers-7.27.0`

## Description

Bumps [@babel/helpers](https://github.com/babel/babel/tree/HEAD/packages/babel-helpers) from 7.26.7 to 7.27.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/releases"><code>@​babel/helpers</code>'s releases</a>.</em></p>
<blockquote>
<h2>v7.27.0 (2025-03-24)</h2>
<p>Thanks <a href="https://github.com/ishchhabra"><code>@​ishchhabra</code></a> and <a href="https://github.com/vovkasm"><code>@​vovkasm</code></a> for your first PRs!</p>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16977">#16977</a> Default <code>importAttributesKeyword</code> to <code>with</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17169">#17169</a> Allow <code>traverseFast</code> to exit early (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17110">#17110</a> Add <code>ImportAttributes</code> to <code>Standardized</code> and move its parser test fixtures (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17100">#17100</a> fix(babel-generator): add named export of generate function (<a href="https://github.com/vovkasm"><code>@​vovkasm</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-template</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17149">#17149</a> Add <code>allowYieldOutsideFunction</code> to parser (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17102">#17102</a> feat: Add <code>upToScope</code> parameter to <code>hasBinding</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17082">#17082</a> Support ESTree AccessorProperty (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17162">#17162</a> feat(babel-types): Add support for BigInt literal conversion in valueToNode (<a href="https://github.com/ishchhabra"><code>@​ishchhabra</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-transform-class-properties</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16816">#16816</a> fix: Class reference in type throws error (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17170">#17170</a> fix: Reset child scopes when <code>scope.crawl()</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-preset-typescript</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17118">#17118</a> Fix: align behaviour to tsc <code>rewriteRelativeImportExtensions</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-cli</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17182">#17182</a> fix: <code>@babel/cli</code> generates duplicate inline source maps (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-named-capturing-groups-regex</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17175">#17175</a> Generate computed proto key (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16870">#16870</a> perf: Improve builders of <code>@babel/types</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-create-regexp-features-plugin</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17176">#17176</a> fix: improve duplicate named groups check (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>Committers: 5</h4>
<ul>
<li>Babel Bot (<a href="https://github.com/babel-bot"><code>@​babel-bot</code></a>)</li>
<li>Huáng Jùnliàng (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
<li>Ish Chhabra (<a href="https://github.com/ishchhabra"><code>@​ishchhabra</code></a>)</li>
<li>Vladimir Timofeev (<a href="https://github.com/vovkasm"><code>@​vovkasm</code></a>)</li>
<li><a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a></li>
</ul>
<h2>v7.26.10 (2025-03-11)</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/babel/babel/blob/main/CHANGELOG.md"><code>@​babel/helpers</code>'s changelog</a>.</em></p>
<blockquote>
<h2>v7.27.0 (2025-03-24)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-generator</code>, <code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16977">#16977</a> Default <code>importAttributesKeyword</code> to <code>with</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:rocket: New Feature</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-traverse</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17169">#17169</a> Allow <code>traverseFast</code> to exit early (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17110">#17110</a> Add <code>ImportAttributes</code> to <code>Standardized</code> and move its parser test fixtures (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-generator</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17100">#17100</a> fix(babel-generator): add named export of generate function (<a href="https://github.com/vovkasm"><code>@​vovkasm</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>, <code>babel-template</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17149">#17149</a> Add <code>allowYieldOutsideFunction</code> to parser (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-typescript</code>, <code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17102">#17102</a> feat: Add <code>upToScope</code> parameter to <code>hasBinding</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17082">#17082</a> Support ESTree AccessorProperty (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17162">#17162</a> feat(babel-types): Add support for BigInt literal conversion in valueToNode (<a href="https://github.com/ishchhabra"><code>@​ishchhabra</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-helper-create-class-features-plugin</code>, <code>babel-plugin-transform-class-properties</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16816">#16816</a> fix: Class reference in type throws error (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-traverse</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17170">#17170</a> fix: Reset child scopes when <code>scope.crawl()</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helpers</code>, <code>babel-preset-typescript</code>, <code>babel-runtime-corejs2</code>, <code>babel-runtime-corejs3</code>, <code>babel-runtime</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17118">#17118</a> Fix: align behaviour to tsc <code>rewriteRelativeImportExtensions</code> (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-cli</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17182">#17182</a> fix: <code>@babel/cli</code> generates duplicate inline source maps (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-plugin-transform-named-capturing-groups-regex</code>, <code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17175">#17175</a> Generate computed proto key (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:running_woman: Performance</h4>
<ul>
<li><code>babel-types</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/16870">#16870</a> perf: Improve builders of <code>@babel/types</code> (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
<li><code>babel-helper-create-regexp-features-plugin</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17176">#17176</a> fix: improve duplicate named groups check (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h2>v7.26.10 (2025-03-11)</h2>
<h4>:eyeglasses: Spec Compliance</h4>
<ul>
<li><code>babel-parser</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17159">#17159</a> Disallow decorator in array pattern (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
</ul>
<h4>:bug: Bug Fix</h4>
<ul>
<li><code>babel-parser</code>, <code>babel-template</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17164">#17164</a> Fix: always initialize ExportDeclaration attributes (<a href="https://github.com/JLHwung"><code>@​JLHwung</code></a>)</li>
</ul>
</li>
<li><code>babel-core</code>
<ul>
<li><a href="https://redirect.github.com/babel/babel/pull/17142">#17142</a> fix: &quot;Map maximum size exceeded&quot; in deepClone (<a href="https://github.com/liuxingbaoyu"><code>@​liuxingbaoyu</code></a>)</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/babel/babel/commit/5c350eab83dd12268add44cce0eeda6c898211e3"><code>5c350ea</code></a> v7.27.0</li>
<li><a href="https://github.com/babel/babel/commit/ca4865a7f43a6a56aec242e23e4a3e318cf0ca92"><code>ca4865a</code></a> Fix: align behaviour to tsc <code>rewriteRelativeImportExtensions</code> (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-helpers/issues/17118">#17118</a>)</li>
<li><a href="https://github.com/babel/babel/commit/e1ce99df422971175249509e7bbc2b327b8f7957"><code>e1ce99d</code></a> v7.26.10</li>
<li><a href="https://github.com/babel/babel/commit/d5952e80c0faa5ec20e35085531b6e572d31dad4"><code>d5952e8</code></a> Fix processing of replacement pattern with named capture groups (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-helpers/issues/17173">#17173</a>)</li>
<li><a href="https://github.com/babel/babel/commit/64bca7b5f308cd52c192a5c821a57f6d1b0475f4"><code>64bca7b</code></a> v7.26.9</li>
<li><a href="https://github.com/babel/babel/commit/4cf5c9e0fbe8899bb9eb0dac7c615406a4fe926d"><code>4cf5c9e</code></a> [babel 8] Use <code>@babel/types</code> for parser's return type (<a href="https://github.com/babel/babel/tree/HEAD/packages/babel-helpers/issues/17117">#17117</a>)</li>
<li>See full diff in <a href="https://github.com/babel/babel/commits/v7.27.0/packages/babel-helpers">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=@babel/helpers&package-manager=npm_and_yarn&previous-version=7.26.7&new-version=7.27.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

### Comment by @codecov-commenter on April 24, 2025 at 06:24 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1319?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 100.00%. Comparing base [(`610b99b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/610b99babb66500e77e79cac7ac5a30dbbe93f6c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`b9cf4b9`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/b9cf4b9284beda54b9c6763700ca249b5e1d1b5e?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff            @@
##            master     #1319   +/-   ##
=========================================
  Coverage   100.00%   100.00%           
=========================================
  Files            1         1           
  Lines            6         6           
  Branches         2         2           
=========================================
  Hits             6         6           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1319/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1319/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1319/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1319?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @leSamo on May 06, 2025 at 07:39 AM UTC

/retest

### Comment by @dependabot on January 21, 2026 at 12:42 PM UTC

OK, I won't notify you again about this release, but will get in touch when a new version is available. If you'd rather skip all updates until the next major or minor version, let me know by commenting `@dependabot ignore this major version` or `@dependabot ignore this minor version`. You can also ignore all major, minor, or patch releases for a dependency by adding an [`ignore` condition](https://docs.github.com/en/code-security/supply-chain-security/configuration-options-for-dependency-updates#ignore) with the desired `update_types` to your config file.

If you change your mind, just re-open this PR and I'll resolve any conflicts on it.

---

## Reviews

### Review by @leSamo - Approved on May 06, 2025 at 07:21 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1319*
