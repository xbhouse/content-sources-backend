---
type: pull_request
number: 1158
title: "Build: Bump the go group with 3 updates"
state: merged
author: dependabot
created: 2025-07-28T06:08:14Z
updated: 2025-07-28T08:06:01Z
closed: 2025-07-28T08:05:54Z
merged: 2025-07-28T08:05:54Z
base_branch: main
head_branch: dependabot/go_modules/go-eca36e2f3e
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1158
---

# Pull Request #1158: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: July 28, 2025 at 06:08 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-eca36e2f3e`

## Description

Bumps the go group with 3 updates: [gorm.io/gorm](https://github.com/go-gorm/gorm), [github.com/lzap/cloudwatchwriter2](https://github.com/lzap/cloudwatchwriter2) and [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils).

Updates `gorm.io/gorm` from 1.30.0 to 1.30.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/go-gorm/gorm/releases">gorm.io/gorm's releases</a>.</em></p>
<blockquote>
<h2>Release v1.30.1</h2>
<h2>Changes</h2>
<ul>
<li>optimize: field.ReflectValueOf <a href="https://github.com/liov"><code>@​liov</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7530">#7530</a>)</li>
<li>optimize: performance optimization <a href="https://github.com/liov"><code>@​liov</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7526">#7526</a>)</li>
<li>fix(schema): check the hook function parameter type <a href="https://github.com/demoManito"><code>@​demoManito</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7468">#7468</a>)</li>
<li>Fix: Unexpected OR Conditions force converted to AND <a href="https://github.com/Riseif"><code>@​Riseif</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7512">#7512</a>)</li>
<li>Add GaussDB Database Support <a href="https://github.com/moseszane168"><code>@​moseszane168</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7508">#7508</a>)</li>
<li>Call after initialize for gorm.Config <a href="https://github.com/jinzhu"><code>@​jinzhu</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7518">#7518</a>)</li>
<li>A little optimization for filed.ValueOf <a href="https://github.com/liov"><code>@​liov</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7499">#7499</a>)</li>
<li>fixes <a href="https://redirect.github.com/go-gorm/gorm/issues/7486">#7486</a> <a href="https://github.com/Eshan-Jogwar"><code>@​Eshan-Jogwar</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7492">#7492</a>)</li>
<li>fix decimal migrate error. <a href="https://github.com/Chise1"><code>@​Chise1</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7450">#7450</a>)</li>
<li>test: update MySQL test matrix to use official images and add 9.0, 8.4 versions <a href="https://github.com/enomotodev"><code>@​enomotodev</code></a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7476">#7476</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/49eaeacb892c8bb2de4a1ea70607eaaf4c15ab9b"><code>49eaeac</code></a> optimize: field.ReflectValueOf (<a href="https://redirect.github.com/go-gorm/gorm/issues/7530">#7530</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/52b4744410b4c60fcbcebe7839f42a6f92be30c8"><code>52b4744</code></a> optimize: performance optimization (<a href="https://redirect.github.com/go-gorm/gorm/issues/7526">#7526</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/9af6d510b58c7691fa33694004c404032208f84f"><code>9af6d51</code></a> Fix query when map keys include table-qualified column names, close <a href="https://redirect.github.com/go-gorm/gorm/issues/7507">#7507</a></li>
<li><a href="https://github.com/go-gorm/gorm/commit/c63374f5d154709f99de902cb4d90a1aa8279103"><code>c63374f</code></a> Don't request LastInsertID from database if not necessary, close <a href="https://redirect.github.com/go-gorm/gorm/issues/7469">#7469</a></li>
<li><a href="https://github.com/go-gorm/gorm/commit/b9c7e562b0e99beb93824c5c24d8bae47696c537"><code>b9c7e56</code></a> fix(schema): check the hook function parameter type (<a href="https://redirect.github.com/go-gorm/gorm/issues/7468">#7468</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/985940f0d855dcefa2a1a847a1d37795335d4579"><code>985940f</code></a> should check inner condition length (<a href="https://redirect.github.com/go-gorm/gorm/issues/7512">#7512</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/991c2d4891ddf17eccc2758c84c8061ef6d0b7b5"><code>991c2d4</code></a> Add GaussDB Database Support (<a href="https://redirect.github.com/go-gorm/gorm/issues/7508">#7508</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/751a6dde7a4510c105bb3fb21ea8cc0f26d3aebd"><code>751a6dd</code></a> Call after initialize for gorm.Config (<a href="https://redirect.github.com/go-gorm/gorm/issues/7518">#7518</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/2f4925e017f020909c061c3feb8ba54a0c5d002c"><code>2f4925e</code></a> A little optimization for filed.ValueOf (<a href="https://redirect.github.com/go-gorm/gorm/issues/7499">#7499</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/1e8baf545953dd58e2f301f4dfef5febbc12da0f"><code>1e8baf5</code></a> fixes <a href="https://redirect.github.com/go-gorm/gorm/issues/7486">#7486</a> (<a href="https://redirect.github.com/go-gorm/gorm/issues/7492">#7492</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/go-gorm/gorm/compare/v1.30.0...v1.30.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/lzap/cloudwatchwriter2` from 1.4.2 to 1.5.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/3e33bb08b7f19440ee767d284bdf775e729db6c4"><code>3e33bb0</code></a> writer: cloudwatch: fix close timeout handling</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/f2ded8bb91e411289146a3b5afd54a5872ff8e0c"><code>f2ded8b</code></a> writer: improve race condition test</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/076996e820b9e8349bcf470c9628dc87ad0d8d9a"><code>076996e</code></a> writer: protect Flush/Close with mutex</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/b2627b498fa9c8ccf38f688f3880497b3ec730ee"><code>b2627b4</code></a> handler: add missing CloseWithTimeout method</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/7c999ce09eabafb24bb1300b967ea2b2e437f542"><code>7c999ce</code></a> writer: add CloseWithTimeout method</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/c25fdb8b82cea1bdf8f0c317ce2a3106e63fb4a7"><code>c25fdb8</code></a> deps: update AWS SDK</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/45a722c8031a34623262debb64a12f7b4fbfeb5c"><code>45a722c</code></a> writer: add race condition test for Flush and Close</li>
<li><a href="https://github.com/lzap/cloudwatchwriter2/commit/adb1989de9cd0eb5c26ea9d0a6d16d65baf5fd0d"><code>adb1989</code></a> handler: add missing Flush method</li>
<li>See full diff in <a href="https://github.com/lzap/cloudwatchwriter2/compare/v1.4.2...v1.5.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.15 to 1.26.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/2bb36276873da5676b177f85b24b55b568a04dde"><code>2bb3627</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/624">#624</a> from joselsegura/remove_kafkazerolog</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4f676d9a7bb416ee2ffca58779e45263638dc4bc"><code>4f676d9</code></a> [CCXDEV-15328] Remove KafkaZerolog dependency</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/a60f4d0f18fe09a25081d34ae9f290a5901264f0"><code>a60f4d0</code></a> refactor: remove unused file manifest.txt (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/623">#623</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/d2d41606e4c4b17afe906e7365cd02c31b59630e"><code>d2d4160</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/622">#622</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/19380261b540407eeb6c1d9d0fd759ba17228765"><code>1938026</code></a> Bump github.com/getsentry/sentry-go from 0.34.0 to 0.34.1</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/259d54fb96b64d1fef2e772da91c26a683aaf52c"><code>259d54f</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/621">#621</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/57090955d12663895f6fa246d23afd6394a0f6c6"><code>5709095</code></a> Bump github.com/redis/go-redis/v9 from 9.10.0 to 9.11.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/64fa22d210ae2d8911362e8a24f87e162cdc1f43"><code>64fa22d</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/620">#620</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4348907f257a7f72081bb096c99f8681feffdd0b"><code>4348907</code></a> Bump github.com/getsentry/sentry-go from 0.33.0 to 0.34.0</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.15...v1.26.0">compare view</a></li>
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
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on July 28, 2025 at 08:05 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1158*
