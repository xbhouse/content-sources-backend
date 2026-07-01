---
type: pull_request
number: 1534
title: "chore(deps): bump jest-environment-jsdom and fix UI tests"
state: merged
author: dependabot
created: 2026-03-05T05:48:31Z
updated: 2026-03-17T12:12:29Z
closed: 2026-03-17T12:12:20Z
merged: 2026-03-17T12:12:20Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-8f3f190bed
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1534
---

# Pull Request #1534: chore(deps): bump jest-environment-jsdom and fix UI tests

**Author**: @dependabot
**Created**: March 05, 2026 at 05:48 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-8f3f190bed`

## Description

ThisPR: 

- Fixes wording in tests for satelite-managed and version locked systems.
- Removes [@tootallnate/once](https://github.com/TooTallNate/once). It's no longer used after updating ancestor dependency [jest-environment-jsdom](https://github.com/jestjs/jest/tree/HEAD/packages/jest-environment-jsdom). These dependencies need to be updated together.
 - Updates `jest-environment-jsdom` from 29.7.0 to 30.2.0

<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/releases">jest-environment-jsdom's releases</a>.</em></p>
<blockquote>
<h2>30.2.0</h2>
<h3>Chore &amp; Maintenance</h3>
<ul>
<li><code>[*]</code> Update example repo for testing React Native projects (<a href="https://redirect.github.com/jestjs/jest/pull/15832">#15832</a>)</li>
<li><code>[*]</code> Update <code>jest-watch-typeahead</code> to v3 (<a href="https://redirect.github.com/jestjs/jest/pull/15830">#15830</a>)</li>
</ul>
<h2>Features</h2>
<ul>
<li><code>[jest-environment-jsdom-abstract]</code> Add support for JSDOM v27 (<a href="https://redirect.github.com/jestjs/jest/pull/15834">#15834</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[babel-jest]</code> Export the <code>TransformerConfig</code> interface (<a href="https://redirect.github.com/jestjs/jest/pull/15820">#15820</a>)</li>
<li><code>[jest-config]</code> Fix <code>jest.config.ts</code> with TS loader specified in docblock pragma (<a href="https://redirect.github.com/jestjs/jest/pull/15839">#15839</a>)</li>
</ul>
<h2>30.1.3</h2>
<h3>Fixes</h3>
<ul>
<li>Fix <code>unstable_mockModule</code> with <code>node:</code> prefixed core modules.</li>
</ul>
<h2>30.1.2</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Correct snapshot header regexp to work with newline across OSes (<a href="https://redirect.github.com/jestjs/jest/pull/15803">#15803</a>)</li>
</ul>
<h2>30.1.1</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot warning not handling Windows end-of-line sequences (<a href="https://redirect.github.com/jestjs/jest/pull/15800">#15800</a>)</li>
</ul>
<h2>30.1.0</h2>
<h2>Features</h2>
<ul>
<li><code>[jest-leak-detector]</code> Configurable GC aggressiveness regarding to V8 heap snapshot generation (<a href="https://redirect.github.com/jestjs/jest/pull/15793/">#15793</a>)</li>
<li><code>[jest-runtime]</code> Reduce redundant ReferenceError messages</li>
<li><code>[jest-core]</code> Include test modules that failed to load when --onlyFailures is active</li>
</ul>
<h3>Fixes</h3>
<ul>
<li>`[jest-snapshot-utils] Fix deprecated goo.gl snapshot guide link not getting replaced with fully canonical URL (<a href="https://redirect.github.com/jestjs/jest/pull/15787">#15787</a>)</li>
<li><code>[jest-circus]</code> Fix <code>it.concurrent</code> not working with <code>describe.skip</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15765">#15765</a>)</li>
<li><code>[jest-snapshot]</code> Fix mangled inline snapshot updates when used with Prettier 3 and CRLF line endings</li>
<li><code>[jest-runtime]</code> Importing from <code>@jest/globals</code> in more than one file no longer breaks relative paths (<a href="https://redirect.github.com/jestjs/jest/issues/15772">#15772</a>)</li>
</ul>
<h1>Chore</h1>
<ul>
<li><code>[expect]</code> Update docblock for <code>toContain()</code> to display info on substring check (<a href="https://redirect.github.com/jestjs/jest/pull/15789">#15789</a>)</li>
</ul>
<h2>30.0.2</h2>
<h2>What's Changed</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jestjs/jest/blob/main/CHANGELOG.md">jest-environment-jsdom's changelog</a>.</em></p>
<blockquote>
<h2>30.2.0</h2>
<h3>Chore &amp; Maintenance</h3>
<ul>
<li><code>[*]</code> Update example repo for testing React Native projects (<a href="https://redirect.github.com/jestjs/jest/pull/15832">#15832</a>)</li>
<li><code>[*]</code> Update <code>jest-watch-typeahead</code> to v3 (<a href="https://redirect.github.com/jestjs/jest/pull/15830">#15830</a>)</li>
</ul>
<h2>Features</h2>
<ul>
<li><code>[jest-environment-jsdom-abstract]</code> Add support for JSDOM v27 (<a href="https://redirect.github.com/jestjs/jest/pull/15834">#15834</a>)</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-matcher-utils]</code> Fix infinite recursion with self-referential getters in <code>deepCyclicCopyReplaceable</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15831">#15831</a>)</li>
<li><code>[babel-jest]</code> Export the <code>TransformerConfig</code> interface (<a href="https://redirect.github.com/jestjs/jest/pull/15820">#15820</a>)</li>
<li><code>[jest-config]</code> Fix <code>jest.config.ts</code> with TS loader specified in docblock pragma (<a href="https://redirect.github.com/jestjs/jest/pull/15839">#15839</a>)</li>
</ul>
<h2>30.1.3</h2>
<h3>Fixes</h3>
<ul>
<li>Fix <code>unstable_mockModule</code> with <code>node:</code> prefixed core modules.</li>
</ul>
<h2>30.1.2</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Correct snapshot header regexp to work with newline across OSes (<a href="https://redirect.github.com/jestjs/jest/pull/15803">#15803</a>)</li>
</ul>
<h2>30.1.1</h2>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot warning not handling Windows end-of-line sequences (<a href="https://redirect.github.com/jestjs/jest/pull/15800">#15800</a>)</li>
<li><code>[jest-snapshot-utils]</code> Improve messaging about goo.gl snapshot link change (<a href="https://redirect.github.com/jestjs/jest/pull/15821">#15821</a>)</li>
</ul>
<h2>30.1.0</h2>
<h2>Features</h2>
<ul>
<li><code>[jest-leak-detector]</code> Configurable GC aggressiveness regarding to V8 heap snapshot generation (<a href="https://redirect.github.com/jestjs/jest/pull/15793/">#15793</a>)</li>
<li><code>[jest-runtime]</code> Reduce redundant ReferenceError messages</li>
<li><code>[jest-core]</code> Include test modules that failed to load when --onlyFailures is active</li>
</ul>
<h3>Fixes</h3>
<ul>
<li><code>[jest-snapshot-utils]</code> Fix deprecated goo.gl snapshot guide link not getting replaced with fully canonical URL (<a href="https://redirect.github.com/jestjs/jest/pull/15787">#15787</a>)</li>
<li><code>[jest-circus]</code> Fix <code>it.concurrent</code> not working with <code>describe.skip</code> (<a href="https://redirect.github.com/jestjs/jest/pull/15765">#15765</a>)</li>
<li><code>[jest-snapshot]</code> Fix mangled inline snapshot updates when used with Prettier 3 and CRLF line endings</li>
<li><code>[jest-runtime]</code> Importing from <code>@jest/globals</code> in more than one file no longer breaks relative paths (<a href="https://redirect.github.com/jestjs/jest/issues/15772">#15772</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jestjs/jest/commit/855864e3f9751366455246790be2bf912d4d0dac"><code>855864e</code></a> v30.2.0</li>
<li><a href="https://github.com/jestjs/jest/commit/ebfa31cc9787303e8698a1a029a162a18e8974aa"><code>ebfa31c</code></a> v30.1.2</li>
<li><a href="https://github.com/jestjs/jest/commit/d347c0f3f87f976a1dbd9761d503e45f5ced2a7e"><code>d347c0f</code></a> v30.1.1</li>
<li><a href="https://github.com/jestjs/jest/commit/4d5f41d0885c1d9630c81b4fd47f74ab0615e18f"><code>4d5f41d</code></a> v30.1.0</li>
<li><a href="https://github.com/jestjs/jest/commit/22236cf58b66039f81893537c90dee290bab427f"><code>22236cf</code></a> v30.0.5</li>
<li><a href="https://github.com/jestjs/jest/commit/f4296d2bc85c1405f84ddf613a25d0bc3766b7e5"><code>f4296d2</code></a> v30.0.4</li>
<li><a href="https://github.com/jestjs/jest/commit/393acbfac31f64bb38dff23c89224797caded83c"><code>393acbf</code></a> v30.0.2</li>
<li><a href="https://github.com/jestjs/jest/commit/5ce865b4060189fe74cd486544816c079194a0f7"><code>5ce865b</code></a> v30.0.1</li>
<li><a href="https://github.com/jestjs/jest/commit/469f665c2d3bea4a55a194ceebae88724b7202cd"><code>469f665</code></a> v30.0.0</li>
<li><a href="https://github.com/jestjs/jest/commit/ce14203d9156f830a8e24a6e3e8205f670a72a40"><code>ce14203</code></a> v30.0.0-rc.1</li>
<li>Additional commits viewable in <a href="https://github.com/jestjs/jest/commits/v30.2.0/packages/jest-environment-jsdom">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @swadeley on March 09, 2026 at 10:55 AM UTC

@dependabot recreate

### Comment by @TenSt on March 12, 2026 at 11:19 AM UTC

@dependabot rebase

### Comment by @TenSt on March 12, 2026 at 02:01 PM UTC

@dependabot rebase

### Comment by @codecov-commenter on March 13, 2026 at 11:50 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1534?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 72.34%. Comparing base ([`ccc7f0e`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ccc7f0e96fe9300cec6ca6644b09384af2b02530?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`607789f`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/607789fdc594fd18f92d223535069364a4043871?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1534      +/-   ##
==========================================
+ Coverage   72.30%   72.34%   +0.03%     
==========================================
  Files          99       99              
  Lines        2405     2408       +3     
  Branches      677      678       +1     
==========================================
+ Hits         1739     1742       +3     
  Misses        586      586              
  Partials       80       80              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1534?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @TenSt on March 13, 2026 at 03:19 PM UTC

@dependabot rebase

### Comment by @dependabot on March 13, 2026 at 03:19 PM UTC

Looks like this PR has been edited by someone other than Dependabot. That means Dependabot can't rebase it - sorry!

If you're happy for Dependabot to recreate it from scratch, overwriting any edits, you can request `@dependabot recreate`.


---

## Reviews

### Review by @TenSt - Dismissed on March 12, 2026 at 11:19 AM UTC

### Review by @TenSt - Dismissed on March 13, 2026 at 11:46 AM UTC

### Review by @swadeley - Approved on March 16, 2026 at 05:42 PM UTC

ACK

### Review by @Dugowitch - Approved on March 17, 2026 at 11:31 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1534*
