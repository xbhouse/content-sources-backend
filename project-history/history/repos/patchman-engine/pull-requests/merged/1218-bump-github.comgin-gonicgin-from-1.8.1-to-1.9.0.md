---
type: pull_request
number: 1218
title: "Bump github.com/gin-gonic/gin from 1.8.1 to 1.9.0"
state: merged
author: dependabot
created: 2023-05-05T02:44:48Z
updated: 2023-05-18T09:33:55Z
closed: 2023-05-18T09:33:46Z
merged: 2023-05-18T09:33:46Z
base_branch: master
head_branch: dependabot/go_modules/github.com/gin-gonic/gin-1.9.0
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-engine/pull/1218
---

# Pull Request #1218: Bump github.com/gin-gonic/gin from 1.8.1 to 1.9.0

**Author**: @dependabot
**Created**: May 05, 2023 at 02:44 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/go_modules/github.com/gin-gonic/gin-1.9.0`

## Description

Bumps [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) from 1.8.1 to 1.9.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/gin-gonic/gin/releases">github.com/gin-gonic/gin's releases</a>.</em></p>
<blockquote>
<h2>v1.9.0</h2>
<h2>Changelog</h2>
<h3>BREAK CHANGES</h3>
<ul>
<li>Stop useless panicking in context and render <a href="https://redirect.github.com/gin-gonic/gin/pull/2150">#2150</a></li>
</ul>
<h3>BUG FIXES</h3>
<ul>
<li>fix(router): tree bug where loop index is not decremented. <a href="https://redirect.github.com/gin-gonic/gin/pull/3460">#3460</a></li>
<li>fix(context): panic on NegotiateFormat - index out of range <a href="https://redirect.github.com/gin-gonic/gin/pull/3397">#3397</a></li>
<li>Add escape logic for header <a href="https://redirect.github.com/gin-gonic/gin/pull/3500">#3500</a> and <a href="https://redirect.github.com/gin-gonic/gin/pull/3503">#3503</a></li>
</ul>
<h3>SECURITY</h3>
<ul>
<li>Fix the GO-2022-0969 and GO-2022-0288 vulnerabilities <a href="https://redirect.github.com/gin-gonic/gin/pull/3333">#3333</a></li>
<li>fix(security): vulnerability GO-2023-1571 <a href="https://redirect.github.com/gin-gonic/gin/pull/3505">#3505</a></li>
</ul>
<h3>ENHANCEMENTS</h3>
<ul>
<li>feat: add sonic json support <a href="https://redirect.github.com/gin-gonic/gin/pull/3184">#3184</a></li>
<li>chore(file): Creates a directory named path <a href="https://redirect.github.com/gin-gonic/gin/pull/3316">#3316</a></li>
<li>fix: modify interface check way <a href="https://redirect.github.com/gin-gonic/gin/pull/3327">#3327</a></li>
<li>remove deprecated of package io/ioutil <a href="https://redirect.github.com/gin-gonic/gin/pull/3395">#3395</a></li>
<li>refactor: avoid calling strings.ToLower twice <a href="https://redirect.github.com/gin-gonic/gin/pull/3433">#3343</a></li>
<li>console logger HTTP status code bug fixed <a href="https://redirect.github.com/gin-gonic/gin/pull/3453">#3453</a></li>
<li>chore(yaml): upgrade dependency to v3 version <a href="https://redirect.github.com/gin-gonic/gin/pull/3456">#3456</a></li>
<li>chore(router): match method added to routergroup for multiple HTTP methods supporting <a href="https://redirect.github.com/gin-gonic/gin/pull/3464">#3464</a></li>
<li>chore(http): add support for go1.20 http.rwUnwrapper to gin.responseWriter <a href="https://redirect.github.com/gin-gonic/gin/pull/3489">#3489</a></li>
</ul>
<h3>DOCS</h3>
<ul>
<li>docs: update markdown format <a href="https://redirect.github.com/gin-gonic/gin/pull/3260">#3260</a></li>
<li>docs(readme): Add the TOML rendering example <a href="https://redirect.github.com/gin-gonic/gin/pull/3400">#3400</a></li>
<li>docs(readme): move more example to docs/doc.md <a href="https://redirect.github.com/gin-gonic/gin/pull/3449">#3449</a></li>
<li>docs: update markdown format <a href="https://redirect.github.com/gin-gonic/gin/pull/3446">#3446</a></li>
</ul>
<h2>v1.8.2</h2>
<h2>Changelog</h2>
<h3>Bug fixes</h3>
<ul>
<li>0c2a691 fix(engine): missing route params for CreateTestContext (<a href="https://redirect.github.com/gin-gonic/gin/issues/2778">#2778</a>) (<a href="https://redirect.github.com/gin-gonic/gin/issues/2803">#2803</a>)</li>
<li>e305e21 fix(route): redirectSlash bug (<a href="https://redirect.github.com/gin-gonic/gin/issues/3227">#3227</a>)</li>
</ul>
<h3>Others</h3>
<ul>
<li>6a2a260 Fix the GO-2022-1144 vulnerability (<a href="https://redirect.github.com/gin-gonic/gin/issues/3432">#3432</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/gin-gonic/gin/blob/master/CHANGELOG.md">github.com/gin-gonic/gin's changelog</a>.</em></p>
<blockquote>
<h2>Gin v1.9.0</h2>
<h3>BREAK CHANGES</h3>
<ul>
<li>Stop useless panicking in context and render <a href="https://redirect.github.com/gin-gonic/gin/pull/2150">#2150</a></li>
</ul>
<h3>BUG FIXES</h3>
<ul>
<li>fix(router): tree bug where loop index is not decremented. <a href="https://redirect.github.com/gin-gonic/gin/pull/3460">#3460</a></li>
<li>fix(context): panic on NegotiateFormat - index out of range <a href="https://redirect.github.com/gin-gonic/gin/pull/3397">#3397</a></li>
<li>Add escape logic for header <a href="https://redirect.github.com/gin-gonic/gin/pull/3500">#3500</a> and <a href="https://redirect.github.com/gin-gonic/gin/pull/3503">#3503</a></li>
</ul>
<h3>SECURITY</h3>
<ul>
<li>Fix the GO-2022-0969 and GO-2022-0288 vulnerabilities <a href="https://redirect.github.com/gin-gonic/gin/pull/3333">#3333</a></li>
<li>fix(security): vulnerability GO-2023-1571 <a href="https://redirect.github.com/gin-gonic/gin/pull/3505">#3505</a></li>
</ul>
<h3>ENHANCEMENTS</h3>
<ul>
<li>feat: add sonic json support <a href="https://redirect.github.com/gin-gonic/gin/pull/3184">#3184</a></li>
<li>chore(file): Creates a directory named path <a href="https://redirect.github.com/gin-gonic/gin/pull/3316">#3316</a></li>
<li>fix: modify interface check way <a href="https://redirect.github.com/gin-gonic/gin/pull/3327">#3327</a></li>
<li>remove deprecated of package io/ioutil <a href="https://redirect.github.com/gin-gonic/gin/pull/3395">#3395</a></li>
<li>refactor: avoid calling strings.ToLower twice <a href="https://redirect.github.com/gin-gonic/gin/pull/3433">#3343</a></li>
<li>console logger HTTP status code bug fixed <a href="https://redirect.github.com/gin-gonic/gin/pull/3453">#3453</a></li>
<li>chore(yaml): upgrade dependency to v3 version <a href="https://redirect.github.com/gin-gonic/gin/pull/3456">#3456</a></li>
<li>chore(router): match method added to routergroup for multiple HTTP methods supporting <a href="https://redirect.github.com/gin-gonic/gin/pull/3464">#3464</a></li>
<li>chore(http): add support for go1.20 http.rwUnwrapper to gin.responseWriter <a href="https://redirect.github.com/gin-gonic/gin/pull/3489">#3489</a></li>
</ul>
<h3>DOCS</h3>
<ul>
<li>docs: update markdown format <a href="https://redirect.github.com/gin-gonic/gin/pull/3260">#3260</a></li>
<li>docs(readme): Add the TOML rendering example <a href="https://redirect.github.com/gin-gonic/gin/pull/3400">#3400</a></li>
<li>docs(readme): move more example to docs/doc.md <a href="https://redirect.github.com/gin-gonic/gin/pull/3449">#3449</a></li>
<li>docs: update markdown format <a href="https://redirect.github.com/gin-gonic/gin/pull/3446">#3446</a></li>
</ul>
<h2>Gin v1.8.2</h2>
<h3>BUG FIXES</h3>
<ul>
<li>fix(route): redirectSlash bug (<a href="(https://redirect.github.com/gin-gonic/gin/pull/3227)">#3227</a>)</li>
<li>fix(engine): missing route params for CreateTestContext (<a href="(https://redirect.github.com/gin-gonic/gin/pull/2778)">#2778</a>) (<a href="(https://redirect.github.com/gin-gonic/gin/pull/2803)">#2803</a>)</li>
</ul>
<h3>SECURITY</h3>
<ul>
<li>Fix the GO-2022-1144 vulnerability (<a href="(https://redirect.github.com/gin-gonic/gin/pull/3432)">#3432</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/gin-gonic/gin/commit/ea03e10384502e1baf6f560a2b0ea32c342ede5b"><code>ea03e10</code></a> docs(readme): release v1.9.0 version (<a href="https://redirect.github.com/gin-gonic/gin/issues/3474">#3474</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/4cee78f5382d5245c3652e6c15fee715eec505c3"><code>4cee78f</code></a> Fix <a href="https://redirect.github.com/gin-gonic/gin/issues/3500">#3500</a> Add escape logic for header (<a href="https://redirect.github.com/gin-gonic/gin/issues/3503">#3503</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/fc1c43298de675e5252d0b44f97dc5e204bd4e1e"><code>fc1c432</code></a> fix(security): vulnerability GO-2023-1571 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3505">#3505</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/81ac7d55a09e34013225db0aeac6e70c1ae68928"><code>81ac7d5</code></a> Add escape logic for header (<a href="https://redirect.github.com/gin-gonic/gin/issues/3500">#3500</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/d07db174acf44bfaf191ca2f6d7beafa2ff946da"><code>d07db17</code></a> chore(deps): bump golang.org/x/net from 0.5.0 to 0.6.0 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3498">#3498</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/c1d06e3d08692f9eddde381a5e277b41fff5a297"><code>c1d06e3</code></a> add supprt for go1.20 http.rwUnwrapper to gin.responseWriter (<a href="https://redirect.github.com/gin-gonic/gin/issues/3489">#3489</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/bd82c9e351be91e9e8267e5ce011627dd6c55d51"><code>bd82c9e</code></a> chore(go): Add support go 1.20 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3484">#3484</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/0c96a20209ca035964be126a745c167196fb6db3"><code>0c96a20</code></a> Stop useless panicking in context and render (<a href="https://redirect.github.com/gin-gonic/gin/issues/2150">#2150</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/153b229fcc6570bac0674d02ab1a629804f29072"><code>153b229</code></a> chore(deps): bump github.com/ugorji/go/codec from 1.2.8 to 1.2.9 (<a href="https://redirect.github.com/gin-gonic/gin/issues/3491">#3491</a>)</li>
<li><a href="https://github.com/gin-gonic/gin/commit/e02ae6ae61fada360379b5bdc7f23e46f21ce5de"><code>e02ae6a</code></a> chore(router): match method added to routergroup for multiple HTTP methods su...</li>
<li>Additional commits viewable in <a href="https://github.com/gin-gonic/gin/compare/v1.8.1...v1.9.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=github.com/gin-gonic/gin&package-manager=go_modules&previous-version=1.8.1&new-version=1.9.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-engine/network/alerts).

</details>

---

## Discussion

### Comment by @jira-linking on May 05, 2023 at 02:44 AM UTC

Commits missing Jira IDs:
e3d4cae4bf5450e42a2216d166e8c5a170296d15


### Comment by @codecov-commenter on May 05, 2023 at 02:55 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1218?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage: **`25.00`**% and project coverage change: **`-0.07`** :warning:
> Comparison is base [(`466f9e9`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/466f9e9177f0753208c841f318fc9de035f596cb?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.02% compared to head [(`e3d4cae`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1218?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 61.96%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1218      +/-   ##
==========================================
- Coverage   62.02%   61.96%   -0.07%     
==========================================
  Files         103      103              
  Lines        6286     6299      +13     
==========================================
+ Hits         3899     3903       +4     
- Misses       1887     1896       +9     
  Partials      500      500              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `61.96% <25.00%> (-0.07%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1218?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/authentication.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1218?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `22.95% <ø> (+2.66%)` | :arrow_up: |
| [manager/controllers/baseline\_systems.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1218?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9zeXN0ZW1zLmdv) | `48.45% <25.00%> (-8.13%)` | :arrow_down: |


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1218?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @MichaelMraka - Approved on May 18, 2023 at 09:33 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1218*
