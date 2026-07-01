---
type: pull_request
number: 1212
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2025-09-22T04:07:05Z
updated: 2025-09-22T09:44:58Z
closed: 2025-09-22T09:44:51Z
merged: 2025-09-22T09:44:51Z
base_branch: main
head_branch: dependabot/go_modules/go-4bdc3daebb
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1212
---

# Pull Request #1212: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: September 22, 2025 at 04:07 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-4bdc3daebb`

## Description

[//]: # (dependabot-start)
⚠️  **Dependabot is rebasing this PR** ⚠️ 

Rebasing might not happen immediately, so don't worry if this takes some time.

Note: if you make any changes to this PR yourself, they will take precedence over the rebase.

---

[//]: # (dependabot-end)

Bumps the go group with 4 updates: [github.com/content-services/tang](https://github.com/content-services/tang), [github.com/IBM/sarama](https://github.com/IBM/sarama), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) and [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go).

Updates `github.com/content-services/tang` from 0.0.13 to 0.0.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/ef0dcfe9a00bec4b61b3fb3ad6ab0e1c99e52002"><code>ef0dcfe</code></a> HMS-8902: use content_ids for version queries (<a href="https://redirect.github.com/content-services/tang/issues/20">#20</a>)</li>
<li><a href="https://github.com/content-services/tang/commit/d934153339e5c4db7c659461f54c2a040ea365f5"><code>d934153</code></a> HMS-8916: fix volume mount to persist database (<a href="https://redirect.github.com/content-services/tang/issues/19">#19</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.13...v0.0.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.46.0 to 1.46.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.46.1 (2025-09-18)</h2>
<blockquote>
<p>[!NOTE]<br />
The go.mod directive has been bumped to 1.24.0 as the minimum version of Go required for the module. This was necessary to continue to receive updates from some of the third party dependencies that Sarama makes use of.</p>
</blockquote>
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat: support more describe log dirs versions (V2-V4) by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3293">IBM/sarama#3293</a></li>
<li>feat: support V5 ListConsumerGroups protocol by <a href="https://github.com/hindessm"><code>@​hindessm</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3292">IBM/sarama#3292</a></li>
<li>feat: add SASLv1 support for Kerberos by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3279">IBM/sarama#3279</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: add read deadline to tls write by <a href="https://github.com/bvalente"><code>@​bvalente</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3283">IBM/sarama#3283</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump go directive to 1.24.0 and golang.org/x/{crypto,net,sync} by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3288">IBM/sarama#3288</a></li>
<li>chore(deps): bump the golang-x group across 6 directories with 1 update by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3291">IBM/sarama#3291</a></li>
<li>chore(deps): bump github.com/stretchr/testify from 1.11.0 to 1.11.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3274">IBM/sarama#3274</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: refactor to use modern atomic types by <a href="https://github.com/Sahil-4555"><code>@​Sahil-4555</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3277">IBM/sarama#3277</a></li>
<li>chore: pre-commit autoupdate to latest by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3278">IBM/sarama#3278</a></li>
<li>chore: apply modernize fixes from gopls by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3297">IBM/sarama#3297</a></li>
<li>chore(config): update comments of sarama.Config.Metadata.SingleFlight by <a href="https://github.com/gunli"><code>@​gunli</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3296">IBM/sarama#3296</a></li>
<li>chore(client): update comments of client methods by <a href="https://github.com/gunli"><code>@​gunli</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3295">IBM/sarama#3295</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Sahil-4555"><code>@​Sahil-4555</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3277">IBM/sarama#3277</a></li>
<li><a href="https://github.com/bvalente"><code>@​bvalente</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3283">IBM/sarama#3283</a></li>
<li><a href="https://github.com/gunli"><code>@​gunli</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3296">IBM/sarama#3296</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.46.0...v1.46.1">https://github.com/IBM/sarama/compare/v1.46.0...v1.46.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/9bc3d146bc3cb1adb31d11a1447867be813beb24"><code>9bc3d14</code></a> chore(client): update comments of client methods (<a href="https://redirect.github.com/IBM/sarama/issues/3295">#3295</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/096b846ae88c545d58ca255ef4130abf2c272640"><code>096b846</code></a> Merge pull request <a href="https://redirect.github.com/IBM/sarama/issues/3296">#3296</a> from gunli/update-single-comments</li>
<li><a href="https://github.com/IBM/sarama/commit/dedff7aa24ade6c793275e88380968c6c3d23561"><code>dedff7a</code></a> Merge pull request <a href="https://redirect.github.com/IBM/sarama/issues/3297">#3297</a> from IBM/dnwe/sarama-modernize</li>
<li><a href="https://github.com/IBM/sarama/commit/5648bd97ddd066c9949af4868185885d2c1cae31"><code>5648bd9</code></a> chore: inline strsContains</li>
<li><a href="https://github.com/IBM/sarama/commit/30d5d02a04724b64feeb9587e163d66088003878"><code>30d5d02</code></a> chore: apply stringsseq from modernize</li>
<li><a href="https://github.com/IBM/sarama/commit/9cba592012de10e2e54f92c1404f090c412d72f9"><code>9cba592</code></a> chore: apply bloop from modernize</li>
<li><a href="https://github.com/IBM/sarama/commit/7b353e936cf867d7a69924c5d2c028968b26f1fb"><code>7b353e9</code></a> chore: apply fmtappendf from modernize</li>
<li><a href="https://github.com/IBM/sarama/commit/7ce589719b01d07216d2bddf8fc8bdb9ec8bd667"><code>7ce5897</code></a> chore: apply mapsloop from modernize</li>
<li><a href="https://github.com/IBM/sarama/commit/c85f6fba2f0ad7c98ff377a252828844e097bf57"><code>c85f6fb</code></a> chore: apply sortslice from modernize</li>
<li><a href="https://github.com/IBM/sarama/commit/bccb0fe09dd46633d3d5408fa6d11fdf0c04e99d"><code>bccb0fe</code></a> chore: apply slicescontains from modernize</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.46.0...v1.46.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.57.4 to 1.58.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/66f75b6c7fba726f68fbd28790b68ce1e6cbe050"><code>66f75b6</code></a> Release 2024-07-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/607ebbf2c1801759227d8ff8560d493962b91d77"><code>607ebbf</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23079711e886b1d295c9f9dcc9a64a76a1b0ec0f"><code>2307971</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59bb031cb9935169618cc8237150f2154df079e4"><code>59bb031</code></a> Release 2024-07-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdced9df4acc57b453ecba3747ec553fc62c9ad6"><code>bdced9d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ee8f0446c7418eed12862575f47b334dc9c9fdc6"><code>ee8f044</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0eeba41ba755676d769c442c8fe21f80641f439b"><code>0eeba41</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4509a600408280c8dcdbc6825ba750cf1628423d"><code>4509a60</code></a> Release 2024-06-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0c61504d74dd81214542aae8a68993166935fa2a"><code>0c61504</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7ca59bf2c0fc11f80a59f7f0e5c7f4d7805444e3"><code>7ca59bf</code></a> Update SDK's smithy-go dependency to v1.20.3</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.57.4...service/s3/v1.58.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.35.2 to 0.35.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.35.3</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.3.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Add missing rate limit categories (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1082">#1082</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.35.3</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.35.3.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Add missing rate limit categories (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1082">#1082</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/5feaeafa7162d152509b9452ce4dc1417725336f"><code>5feaeaf</code></a> release: 0.35.3</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/3bcd0d3e603d46b34bc78f2eaad8d1a570eee720"><code>3bcd0d3</code></a> Prepare 0.35.3 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1084">#1084</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a43141fba26bda18ea84f84e076cfb175ea89617"><code>a43141f</code></a> add missing rate limit categories (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1082">#1082</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/54a69e05ea609d3fc32fb1393770258dde6796c1"><code>54a69e0</code></a> Update ISSUE_TEMPLATE (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1079">#1079</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/5a5b81c60029039d6c81c0aab14ca9f95a7503bc"><code>5a5b81c</code></a> Merge branch 'release/0.35.2'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.35.2...v0.35.3">compare view</a></li>
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

### Review by @swadeley - Approved on September 22, 2025 at 09:44 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1212*
