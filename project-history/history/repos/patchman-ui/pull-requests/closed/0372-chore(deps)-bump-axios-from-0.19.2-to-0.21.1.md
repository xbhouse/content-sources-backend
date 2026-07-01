---
type: pull_request
number: 372
title: "chore(deps): bump axios from 0.19.2 to 0.21.1"
state: closed
author: dependabot
created: 2021-01-05T11:51:47Z
updated: 2021-01-25T13:13:45Z
closed: 2021-01-25T13:13:43Z
base_branch: master
head_branch: dependabot/npm_and_yarn/axios-0.21.1
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-ui/pull/372
---

# Pull Request #372: chore(deps): bump axios from 0.19.2 to 0.21.1

**Author**: @dependabot
**Created**: January 05, 2021 at 11:51 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/axios-0.21.1`

## Description

Bumps [axios](https://github.com/axios/axios) from 0.19.2 to 0.21.1.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
<h2>v0.21.1</h2>
<h3>0.21.1 (December 21, 2020)</h3>
<p>Fixes and Functionality:</p>
<ul>
<li>Hotfix: Prevent SSRF (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3410">#3410</a>)</li>
<li>Protocol not parsed when setting proxy config from env vars (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3070">#3070</a>)</li>
<li>Updating axios in types to be lower case (<a href="https://github-redirect.dependabot.com/axios/axios/issues/2797">#2797</a>)</li>
<li>Adding a type guard for <code>AxiosError</code> (<a href="https://github-redirect.dependabot.com/axios/axios/issues/2949">#2949</a>)</li>
</ul>
<p>Internal and Tests:</p>
<ul>
<li>Remove the skipping of the <code>socket</code> http test (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3364">#3364</a>)</li>
<li>Use different socket for Win32 test (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3375">#3375</a>)</li>
</ul>
<p>Huge thanks to everyone who contributed to this release via code (authors listed below) or via reviews and triaging on GitHub:</p>
<ul>
<li>Daniel Lopretto <a href="mailto:timemachine3030@users.noreply.github.com">timemachine3030@users.noreply.github.com</a></li>
<li>Jason Kwok <a href="mailto:JasonHK@users.noreply.github.com">JasonHK@users.noreply.github.com</a></li>
<li>Jay <a href="mailto:jasonsaayman@gmail.com">jasonsaayman@gmail.com</a></li>
<li>Jonathan Foster <a href="mailto:jonathan@jonathanfoster.io">jonathan@jonathanfoster.io</a></li>
<li>Remco Haszing <a href="mailto:remcohaszing@gmail.com">remcohaszing@gmail.com</a></li>
<li>Xianming Zhong <a href="mailto:chinesedfan@qq.com">chinesedfan@qq.com</a></li>
</ul>
<h2>v0.21.0</h2>
<h3>0.21.0 (October 23, 2020)</h3>
<p>Fixes and Functionality:</p>
<ul>
<li>Fixing requestHeaders.Authorization (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3287">#3287</a>)</li>
<li>Fixing node types (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3237">#3237</a>)</li>
<li>Fixing axios.delete ignores config.data (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3282">#3282</a>)</li>
<li>Revert &quot;Fixing overwrite Blob/File type as Content-Type in browser. (<a href="https://github-redirect.dependabot.com/axios/axios/issues/1773">#1773</a>)&quot; (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3289">#3289</a>)</li>
<li>Fixing an issue that type 'null' and 'undefined' is not assignable to validateStatus when typescript strict option is enabled (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3200">#3200</a>)</li>
</ul>
<p>Internal and Tests:</p>
<ul>
<li>Lock travis to not use node v15 (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3361">#3361</a>)</li>
</ul>
<p>Documentation:</p>
<ul>
<li>Fixing simple typo, existant -&gt; existent (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3252">#3252</a>)</li>
<li>Fixing typos (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3309">#3309</a>)</li>
</ul>
<p>Huge thanks to everyone who contributed to this release via code (authors listed below) or via reviews and triaging on GitHub:</p>
<ul>
<li>Allan Cruz <a href="mailto:57270969+Allanbcruz@users.noreply.github.com">57270969+Allanbcruz@users.noreply.github.com</a></li>
<li>George Cheng <a href="mailto:Gerhut@GMail.com">Gerhut@GMail.com</a></li>
<li>Jay <a href="mailto:jasonsaayman@gmail.com">jasonsaayman@gmail.com</a></li>
<li>Kevin Kirsche <a href="mailto:Kev.Kirsche+GitHub@gmail.com">Kev.Kirsche+GitHub@gmail.com</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v0.21.1/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
<h3>0.21.1 (December 21, 2020)</h3>
<p>Fixes and Functionality:</p>
<ul>
<li>Hotfix: Prevent SSRF (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3410">#3410</a>)</li>
<li>Protocol not parsed when setting proxy config from env vars (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3070">#3070</a>)</li>
<li>Updating axios in types to be lower case (<a href="https://github-redirect.dependabot.com/axios/axios/issues/2797">#2797</a>)</li>
<li>Adding a type guard for <code>AxiosError</code> (<a href="https://github-redirect.dependabot.com/axios/axios/issues/2949">#2949</a>)</li>
</ul>
<p>Internal and Tests:</p>
<ul>
<li>Remove the skipping of the <code>socket</code> http test (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3364">#3364</a>)</li>
<li>Use different socket for Win32 test (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3375">#3375</a>)</li>
</ul>
<p>Huge thanks to everyone who contributed to this release via code (authors listed below) or via reviews and triaging on GitHub:</p>
<ul>
<li>Daniel Lopretto <a href="mailto:timemachine3030@users.noreply.github.com">timemachine3030@users.noreply.github.com</a></li>
<li>Jason Kwok <a href="mailto:JasonHK@users.noreply.github.com">JasonHK@users.noreply.github.com</a></li>
<li>Jay <a href="mailto:jasonsaayman@gmail.com">jasonsaayman@gmail.com</a></li>
<li>Jonathan Foster <a href="mailto:jonathan@jonathanfoster.io">jonathan@jonathanfoster.io</a></li>
<li>Remco Haszing <a href="mailto:remcohaszing@gmail.com">remcohaszing@gmail.com</a></li>
<li>Xianming Zhong <a href="mailto:chinesedfan@qq.com">chinesedfan@qq.com</a></li>
</ul>
<h3>0.21.0 (October 23, 2020)</h3>
<p>Fixes and Functionality:</p>
<ul>
<li>Fixing requestHeaders.Authorization (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3287">#3287</a>)</li>
<li>Fixing node types (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3237">#3237</a>)</li>
<li>Fixing axios.delete ignores config.data (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3282">#3282</a>)</li>
<li>Revert &quot;Fixing overwrite Blob/File type as Content-Type in browser. (<a href="https://github-redirect.dependabot.com/axios/axios/issues/1773">#1773</a>)&quot; (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3289">#3289</a>)</li>
<li>Fixing an issue that type 'null' and 'undefined' is not assignable to validateStatus when typescript strict option is enabled (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3200">#3200</a>)</li>
</ul>
<p>Internal and Tests:</p>
<ul>
<li>Lock travis to not use node v15 (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3361">#3361</a>)</li>
</ul>
<p>Documentation:</p>
<ul>
<li>Fixing simple typo, existant -&gt; existent (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3252">#3252</a>)</li>
<li>Fixing typos (<a href="https://github-redirect.dependabot.com/axios/axios/pull/3309">#3309</a>)</li>
</ul>
<p>Huge thanks to everyone who contributed to this release via code (authors listed below) or via reviews and triaging on GitHub:</p>
<ul>
<li>Allan Cruz <a href="mailto:57270969+Allanbcruz@users.noreply.github.com">57270969+Allanbcruz@users.noreply.github.com</a></li>
<li>George Cheng <a href="mailto:Gerhut@GMail.com">Gerhut@GMail.com</a></li>
<li>Jay <a href="mailto:jasonsaayman@gmail.com">jasonsaayman@gmail.com</a></li>
<li>Kevin Kirsche <a href="mailto:Kev.Kirsche+GitHub@gmail.com">Kev.Kirsche+GitHub@gmail.com</a></li>
<li>Remco Haszing <a href="mailto:remcohaszing@gmail.com">remcohaszing@gmail.com</a></li>
<li>Taemin Shin <a href="mailto:cprayer13@gmail.com">cprayer13@gmail.com</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/a64050a6cfbcc708a55a7dc8030d85b1c78cdf38"><code>a64050a</code></a> Releasing 0.21.1</li>
<li><a href="https://github.com/axios/axios/commit/d57cd976f3cc0f1c5bb1f0681660e50004781db5"><code>d57cd97</code></a> Updating changelog for 0.21.1 release</li>
<li><a href="https://github.com/axios/axios/commit/8b0f373df0574b7cb3c6b531b4092cd670dac6e3"><code>8b0f373</code></a> Use different socket for Win32 test (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3375">#3375</a>)</li>
<li><a href="https://github.com/axios/axios/commit/e426910be7c417bdbcde9c18cb184ead826fc0e1"><code>e426910</code></a> Protocol not parsed when setting proxy config from env vars (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3070">#3070</a>)</li>
<li><a href="https://github.com/axios/axios/commit/c7329fefc890050edd51e40e469a154d0117fc55"><code>c7329fe</code></a> Hotfix: Prevent SSRF (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3410">#3410</a>)</li>
<li><a href="https://github.com/axios/axios/commit/f472e5da5fe76c72db703d6a0f5190e4ad31e642"><code>f472e5d</code></a> Adding a type guard for <code>AxiosError</code> (<a href="https://github-redirect.dependabot.com/axios/axios/issues/2949">#2949</a>)</li>
<li><a href="https://github.com/axios/axios/commit/768825589fd0d36b64a66717ca6df2efd8fb7844"><code>7688255</code></a> Remove the skipping of the <code>socket</code> http test (<a href="https://github-redirect.dependabot.com/axios/axios/issues/3364">#3364</a>)</li>
<li><a href="https://github.com/axios/axios/commit/820fe6e41a96f05fb4781673ce07486f1b37515d"><code>820fe6e</code></a> Updating axios in types to be lower case (<a href="https://github-redirect.dependabot.com/axios/axios/issues/2797">#2797</a>)</li>
<li><a href="https://github.com/axios/axios/commit/94ca24b5b23f343769a15f325693246e07c177d2"><code>94ca24b</code></a> Releasing 0.21.0</li>
<li><a href="https://github.com/axios/axios/commit/2130a0c8acc588c72b53dfef31a11442043ffb06"><code>2130a0c</code></a> Updating changelog for 0.21.0 release</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v0.19.2...v0.21.1">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=axios&package-manager=npm_and_yarn&previous-version=0.19.2&new-version=0.21.1)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot use these labels` will set the current labels as the default for future PRs for this repo and language
- `@dependabot use these reviewers` will set the current reviewers as the default for future PRs for this repo and language
- `@dependabot use these assignees` will set the current assignees as the default for future PRs for this repo and language
- `@dependabot use this milestone` will set the current milestone as the default for future PRs for this repo and language

You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @codecov-io on January 05, 2021 at 11:54 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372?src=pr&el=h1) Report
> Merging [#372](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372?src=pr&el=desc) (fe118db) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/1bb8b088aabdf01efbda0a184a6e7b0a3a27117d?el=desc) (1bb8b08) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #372   +/-   ##
=======================================
  Coverage   72.96%   72.96%           
=======================================
  Files          69       69           
  Lines        1239     1239           
  Branches      161      161           
=======================================
  Hits          904      904           
  Misses        280      280           
  Partials       55       55           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372?src=pr&el=footer). Last update [1bb8b08...fe118db](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/372?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @rvsia on January 20, 2021 at 10:16 AM UTC

@jiridostal Hello, can we please get this merged and release a new version?

There is a security vulnerability in axios and it's being shown in inventory-frontend.

### Comment by @dependabot on January 25, 2021 at 01:13 PM UTC

Looks like axios is up-to-date now, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/372*
