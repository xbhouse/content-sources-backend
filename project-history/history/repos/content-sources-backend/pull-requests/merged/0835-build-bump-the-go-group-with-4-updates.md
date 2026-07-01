---
type: pull_request
number: 835
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2024-09-30T04:49:14Z
updated: 2024-10-01T14:10:54Z
closed: 2024-10-01T14:10:46Z
merged: 2024-10-01T14:10:46Z
base_branch: main
head_branch: dependabot/go_modules/go-5eb127e07b
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/835
---

# Pull Request #835: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: September 30, 2024 at 04:49 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-5eb127e07b`

## Description

Bumps the go group with 4 updates: [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2), [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.34 to 1.17.37
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/22efd7712b3738e45ecab4429df34d60caad0310"><code>22efd77</code></a> Release 2024-09-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ce669cae191a2f9abaf6335ddd70c406ec73d8de"><code>ce669ca</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b0084471ffefb9962d2fef39013b9c86b029431c"><code>b008447</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a912d17bf44dff5ea0cd750a062f4c88ce27b939"><code>a912d17</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9306dbe4f90db9893b7fd9dbc3aa10a451059c9f"><code>9306dbe</code></a> remove deprecated worklink service (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2814">#2814</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d795d86c05c4c80bf513b171ab24d46d01f17062"><code>d795d86</code></a> Release 2024-09-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03c35535f7e52ce38e54249b3267a1b9aa4e1ab4"><code>03c3553</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b1c6a1f231eda85d6fe8342f1547571c68dd135b"><code>b1c6a1f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a51f8d95851407cc4369f559cb1eab7b04afab3c"><code>a51f8d9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3a96a42cb9478b39c0cb9c0395b1cf4ed785fb9f"><code>3a96a42</code></a> update changelog (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2805">#2805</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.34...credentials/v1.17.37">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.40.0 to 1.40.3
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-02-15)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/accessanalyzer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/accessanalyzer/CHANGELOG.md#v1193-2023-02-15">v1.19.3</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/account</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/account/CHANGELOG.md#v191-2023-02-15">v1.9.1</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/acm</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/acm/CHANGELOG.md#v1173-2023-02-15">v1.17.3</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/acmpca</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/acmpca/CHANGELOG.md#v1212-2023-02-15">v1.21.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/alexaforbusiness</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/alexaforbusiness/CHANGELOG.md#v1152-2023-02-15">v1.15.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amp</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/amp/CHANGELOG.md#v1162-2023-02-15">v1.16.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplify</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/amplify/CHANGELOG.md#v1132-2023-02-15">v1.13.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplifybackend</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/amplifybackend/CHANGELOG.md#v1142-2023-02-15">v1.14.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/amplifyuibuilder/CHANGELOG.md#v192-2023-02-15">v1.9.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apigateway</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/apigateway/CHANGELOG.md#v1163-2023-02-15">v1.16.3</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/apigatewaymanagementapi/CHANGELOG.md#v1112-2023-02-15">v1.11.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apigatewayv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/apigatewayv2/CHANGELOG.md#v1133-2023-02-15">v1.13.3</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appconfig</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/appconfig/CHANGELOG.md#v1171-2023-02-15">v1.17.1</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appconfigdata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/appconfigdata/CHANGELOG.md#v161-2023-02-15">v1.6.1</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
<li><strong>Bug Fix</strong>: Correct error type parsing for restJson services.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/rds/v1.40.3/service/appflow/CHANGELOG.md#v1242-2023-02-15">v1.24.2</a>
<ul>
<li><strong>Announcement</strong>: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2012">#2012</a> tracked in issue <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1910">#1910</a>.</li>
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
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ac28b42bc2600bc4e6fde4532951e97fbe78023e"><code>ac28b42</code></a> Release 2023-02-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6da6aab277744138ab50aeaf5016ce9ca0370983"><code>6da6aab</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/106a028d56a5261d1d56d553defe6905ee334b4f"><code>106a028</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6795d673e660460d7504e8b127ec2cd40086bb68"><code>6795d67</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1f9c50e881aca24ab2bf8ee0e8d15be7f19a056a"><code>1f9c50e</code></a> bugfix: correct error type parsing for restJson services</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04aa6ff2415f52e9d52637d9ea91fd449e6c92b7"><code>04aa6ff</code></a> Release 2023-02-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e102656d5c90490c8335a355604d06b873ee8d01"><code>e102656</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2851996c1db4dbb68620a6a36ba749d2bf29b923"><code>2851996</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1dcf2438044ea8437011c9464b6c256db742d418"><code>1dcf243</code></a> Fix v4a tests broken in Go 1.20 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2017">#2017</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/81ac480b6b7f4ff7fb890785db891249fe41f3f4"><code>81ac480</code></a> regenerate</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.40.0...service/rds/v1.40.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.4.16 to 4.4.17
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/b004074caff97ecf6fd1d0f957ff8f2ec9c73fff"><code>b004074</code></a> Update bindings to release/v4.4.17</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.4.16...release/v4.4.17">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.9.1726834826 to 2024.9.1727266727
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/41189b0b9b0df4c9a77885b42a65ae94a92c1abf"><code>41189b0</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a84748995869a037d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/952a0c7416a7deb63e6e178fb5a6190d4238a1f7"><code>952a0c7</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a8474544d2add037d356e3e82bb4...</li>
<li><a href="https://github.com/content-services/zest/commit/61e89766022e1723953783fa9a98b4cbc869280a"><code>61e8976</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b7358943a39f37d2e4353b86e2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.9.1726834826...release/v2024.9.1727266727">compare view</a></li>
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

## Discussion

### Comment by @app-sre-bot on September 30, 2024 at 04:51 AM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on September 30, 2024 at 06:18 AM UTC

[test]

---

## Reviews

### Review by @rverdile - Approved on October 01, 2024 at 01:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/835*
