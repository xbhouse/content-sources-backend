---
type: pull_request
number: 1243
title: "Build: Bump the go group with 9 updates"
state: merged
author: dependabot
created: 2025-10-20T04:12:05Z
updated: 2025-10-20T09:53:49Z
closed: 2025-10-20T09:53:42Z
merged: 2025-10-20T09:53:42Z
base_branch: main
head_branch: dependabot/go_modules/go-a94c11f518
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1243
---

# Pull Request #1243: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: October 20, 2025 at 04:12 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-a94c11f518`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/tang](https://github.com/content-services/tang) | `0.0.15` | `0.0.16` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.39.2` | `1.39.3` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.12` | `1.31.13` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.16` | `1.18.17` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.58.2` | `1.58.3` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.88.4` | `1.88.5` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.10.1760102143` | `2025.10.1760721429` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.35.3` | `0.36.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.14.0` | `9.14.1` |

Updates `github.com/content-services/tang` from 0.0.15 to 0.0.16
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/tang/releases">github.com/content-services/tang's releases</a>.</em></p>
<blockquote>
<h2>v0.0.16</h2>
<h2>What's Changed</h2>
<ul>
<li>HMS-9436: update mockery config to v3 by <a href="https://github.com/Dugowitch"><code>@​Dugowitch</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/22">content-services/tang#22</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Dugowitch"><code>@​Dugowitch</code></a> made their first contribution in <a href="https://redirect.github.com/content-services/tang/pull/22">content-services/tang#22</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.15...v0.0.16">https://github.com/content-services/tang/compare/v0.0.15...v0.0.16</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/f2e092470a0410a8b4054dbc2294a16a7927626a"><code>f2e0924</code></a> HMS-9436: update mockery config to v3 (<a href="https://redirect.github.com/content-services/tang/issues/22">#22</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.15...v0.0.16">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.39.2 to 1.39.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9b0c6553b80f99603b4f8356b88f5baf1328deb"><code>b9b0c65</code></a> Release 2025-10-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2bc8a0ec6f430876fc7de4432ea9cc89c9568f8"><code>e2bc8a0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8691ee380a96c49351e4b5ab8a70bc5d4d100724"><code>8691ee3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51e8a3fe032fc566d31b389f492ab58475a98398"><code>51e8a3f</code></a> bump to go1.23 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3211">#3211</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad2d36cba7c5772b4e8e4caf96939dc41b95c65c"><code>ad2d36c</code></a> Release 2025-10-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19a35d639f969ee328553e632e8cf8b83d324106"><code>19a35d6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35cb02fd50fb125601b9c3b33feb72f3a2bcaa56"><code>35cb02f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f673a1b0a80e666c0128ec606ff053dace9771f1"><code>f673a1b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48421fd812d8592a4eb2b32d11ae07e228969012"><code>48421fd</code></a> Release 2025-10-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fedcba778c21b451a91b4e4bcdd5d6c1554c6a5a"><code>fedcba7</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.39.2...v1.39.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.12 to 1.31.13
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9b0c6553b80f99603b4f8356b88f5baf1328deb"><code>b9b0c65</code></a> Release 2025-10-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2bc8a0ec6f430876fc7de4432ea9cc89c9568f8"><code>e2bc8a0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8691ee380a96c49351e4b5ab8a70bc5d4d100724"><code>8691ee3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51e8a3fe032fc566d31b389f492ab58475a98398"><code>51e8a3f</code></a> bump to go1.23 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3211">#3211</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad2d36cba7c5772b4e8e4caf96939dc41b95c65c"><code>ad2d36c</code></a> Release 2025-10-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19a35d639f969ee328553e632e8cf8b83d324106"><code>19a35d6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35cb02fd50fb125601b9c3b33feb72f3a2bcaa56"><code>35cb02f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f673a1b0a80e666c0128ec606ff053dace9771f1"><code>f673a1b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48421fd812d8592a4eb2b32d11ae07e228969012"><code>48421fd</code></a> Release 2025-10-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fedcba778c21b451a91b4e4bcdd5d6c1554c6a5a"><code>fedcba7</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.31.12...config/v1.31.13">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.16 to 1.18.17
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/908071452b2e0d4d44636d15eab8ad6f04d87d3a"><code>9080714</code></a> Release 2023-03-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ee27f3d7777e4860fc89f48cfb2d72654de28d4"><code>2ee27f3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ae563fea7a3356751200e9f15b7e57d4d8d5d16"><code>0ae563f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/af8cc4f8ac5190cc1dcd8abc0b4a17d4e4551cbe"><code>af8cc4f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c2533c6f582626648e4879af07a645bb98123dca"><code>c2533c6</code></a> feat: add flag to disable IMDSv1 fallback (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2048">#2048</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fe59fdbb04e9978e6237e4e3b44d320224dade24"><code>fe59fdb</code></a> breakfix: sparse trait missing on support service (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2050">#2050</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2ea4bf85fd27a93d9c94f9271c95cb9c1a07e0f7"><code>2ea4bf8</code></a> Release 2023-03-13</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b6f2b7c1becfc032b7066e7605e4c980a9eaf58c"><code>b6f2b7c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3e21c2a2dfe1ae0b9ba6d615c51ef5c1277ab329"><code>3e21c2a</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/79b0f7fd0b5ce6fc85310a3fd70b1b131683aadb"><code>79b0f7f</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.16...config/v1.18.17">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.58.2 to 1.58.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5102c2e4b1a28439ef06579615323ea351ff1d4a"><code>5102c2e</code></a> Release 2024-08-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/001b1fdb33bd5c41ac2e0c3108107096f46f58a5"><code>001b1fd</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0897137f76caba2d0733c709dddcdc803d314d72"><code>0897137</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2747951bd66c0b687a0a483195843b73f6448bb"><code>a274795</code></a> add assurance tests for auth scheme select (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2730">#2730</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2d43bf8bc37b409cec3488d61167a19157bd8080"><code>2d43bf8</code></a> Release 2024-08-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5bd193dfea5e730e322beacaed7b1697c7c4f731"><code>5bd193d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5954703703ed37543b746529f37714bb9c6d3e9e"><code>5954703</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0c73d2b5dab5b71f49e363b6f0415734b4a1f600"><code>0c73d2b</code></a> Release 2024-07-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3102b1fd84cd8797ad3684bc4b0da0ec98cc2426"><code>3102b1f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/348b939a21893b9199ecaafe917c538112a6f6b3"><code>348b939</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.58.2...service/s3/v1.58.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.88.4 to 1.88.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9b0c6553b80f99603b4f8356b88f5baf1328deb"><code>b9b0c65</code></a> Release 2025-10-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2bc8a0ec6f430876fc7de4432ea9cc89c9568f8"><code>e2bc8a0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8691ee380a96c49351e4b5ab8a70bc5d4d100724"><code>8691ee3</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/51e8a3fe032fc566d31b389f492ab58475a98398"><code>51e8a3f</code></a> bump to go1.23 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3211">#3211</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad2d36cba7c5772b4e8e4caf96939dc41b95c65c"><code>ad2d36c</code></a> Release 2025-10-15</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19a35d639f969ee328553e632e8cf8b83d324106"><code>19a35d6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/35cb02fd50fb125601b9c3b33feb72f3a2bcaa56"><code>35cb02f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f673a1b0a80e666c0128ec606ff053dace9771f1"><code>f673a1b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/48421fd812d8592a4eb2b32d11ae07e228969012"><code>48421fd</code></a> Release 2025-10-14</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fedcba778c21b451a91b4e4bcdd5d6c1554c6a5a"><code>fedcba7</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.88.4...service/s3/v1.88.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.10.1760102143 to 2025.10.1760721429
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/1f03b9da6d83666a172b027792507105df1ee479"><code>1f03b9d</code></a> Update pulp bindings to e69db48356e528a464be3da896237b3e92a3585fa7d54eb5892a9...</li>
<li><a href="https://github.com/content-services/zest/commit/9638a6fc0f0c9bca57f62ec32a9f0b777613a83f"><code>9638a6f</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73d62e56235137d2e4353b86e...</li>
<li><a href="https://github.com/content-services/zest/commit/b98ff58477ab34f11abe9d2c5ab58a3be5082f6c"><code>b98ff58</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e4754b3664529cb7a98d596b8e3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.10.1760102143...release/v2025.10.1760721429">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.35.3 to 0.36.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.36.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.36.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>
<p>Behavioral change for the <code>MaxBreadcrumbs</code> client option. Removed the hard limit of 100 breadcrumbs, allowing users to set a larger limit and also changed the default limit from 30 to 100 (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1106">#1106</a>))</p>
</li>
<li>
<p>The changes to error handling (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1075">#1075</a>) will affect issue grouping. It is expected that any wrapped and complex errors will be grouped under a new issue group.</p>
</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add support for improved issue grouping with enhanced error chain handling (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1075">#1075</a>)</p>
<p>The SDK now provides better handling of complex error scenarios, particularly when dealing with multiple related errors or error chains. This feature automatically detects and properly structures errors created with Go's <code>errors.Join()</code> function and other multi-error patterns.</p>
<pre lang="go"><code>// Multiple errors are now properly grouped and displayed in Sentry
err1 := errors.New(&quot;err1&quot;)
err2 := errors.New(&quot;err2&quot;) 
combinedErr := errors.Join(err1, err2)
<p>// When captured, these will be shown as related exceptions in Sentry
sentry.CaptureException(combinedErr)
</code></pre></p>
</li>
<li>
<p>Add <code>TraceIgnoreStatusCodes</code> option to allow filtering of HTTP transactions based on status codes (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1089">#1089</a>)</p>
<ul>
<li>Configure which HTTP status codes should not be traced by providing single codes or ranges</li>
<li>Example: <code>TraceIgnoreStatusCodes: [][]int{{404}, {500, 599}}</code> ignores 404 and server errors 500-599</li>
</ul>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Fix logs being incorrectly filtered by <code>BeforeSend</code> callback (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1109">#1109</a>)
<ul>
<li>Logs now bypass the <code>processEvent</code> method and are sent directly to the transport</li>
<li>This ensures logs are only filtered by <code>BeforeSendLog</code>, not by the error/message <code>BeforeSend</code> callback</li>
</ul>
</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Add support for Go 1.25 and drop support for Go 1.22 (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1103">#1103</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.36.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.36.0.</p>
<h3>Breaking Changes</h3>
<ul>
<li>
<p>Behavioral change for the <code>MaxBreadcrumbs</code> client option. Removed the hard limit of 100 breadcrumbs, allowing users to set a larger limit and also changed the default limit from 30 to 100 (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1106">#1106</a>))</p>
</li>
<li>
<p>The changes to error handling (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1075">#1075</a>) will affect issue grouping. It is expected that any wrapped and complex errors will be grouped under a new issue group.</p>
</li>
</ul>
<h3>Features</h3>
<ul>
<li>
<p>Add support for improved issue grouping with enhanced error chain handling (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1075">#1075</a>)</p>
<p>The SDK now provides better handling of complex error scenarios, particularly when dealing with multiple related errors or error chains. This feature automatically detects and properly structures errors created with Go's <code>errors.Join()</code> function and other multi-error patterns.</p>
<pre lang="go"><code>// Multiple errors are now properly grouped and displayed in Sentry
err1 := errors.New(&quot;err1&quot;)
err2 := errors.New(&quot;err2&quot;) 
combinedErr := errors.Join(err1, err2)
<p>// When captured, these will be shown as related exceptions in Sentry
sentry.CaptureException(combinedErr)
</code></pre></p>
</li>
<li>
<p>Add <code>TraceIgnoreStatusCodes</code> option to allow filtering of HTTP transactions based on status codes (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1089">#1089</a>)</p>
<ul>
<li>Configure which HTTP status codes should not be traced by providing single codes or ranges</li>
<li>Example: <code>TraceIgnoreStatusCodes: [][]int{{404}, {500, 599}}</code> ignores 404 and server errors 500-599</li>
</ul>
</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Fix logs being incorrectly filtered by <code>BeforeSend</code> callback (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1109">#1109</a>)
<ul>
<li>Logs now bypass the <code>processEvent</code> method and are sent directly to the transport</li>
<li>This ensures logs are only filtered by <code>BeforeSendLog</code>, not by the error/message <code>BeforeSend</code> callback</li>
</ul>
</li>
</ul>
<h3>Misc</h3>
<ul>
<li>Add support for Go 1.25 and drop support for Go 1.22 (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1103">#1103</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/fe877a9b588baf3485270860b69185d9a8080982"><code>fe877a9</code></a> release: 0.36.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/586e92e4aea2b31aefd29eb16c9e9815eda2142f"><code>586e92e</code></a> Prepare 0.36.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1110">#1110</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/99cebd57a973b21ea8ef1c1462eb7d31f2e4b29f"><code>99cebd5</code></a> feat: properly support error capturing (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1075">#1075</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/49d29a6c8c7956de141b07ca1cdd57f60c93d45c"><code>49d29a6</code></a> fix: logs should not pass from processEvent (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1109">#1109</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2dc3b97fd5d4776ee4aeeacf0701d9161160a786"><code>2dc3b97</code></a> feat!: remove MaxBreadcrumbs hard limit and change default to 100 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1106">#1106</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c24b7484373c2a3afc81e7cffd33f24a6bcc8ece"><code>c24b748</code></a> ref: create debuglog package (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1105">#1105</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/4edee8a70278d8df24be29e92fe650a553b540cf"><code>4edee8a</code></a> Run tests on Go 1.25 and remove Go 1.22 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1103">#1103</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/476385356a59acbacc385dffac9bd37b5b648deb"><code>4763853</code></a> build(deps): bump actions/create-github-app-token from 2.0.6 to 2.1.4 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1098">#1098</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/194af72085956d975ee0c144494a0d34bade6d49"><code>194af72</code></a> build(deps): bump actions/setup-go from 5 to 6 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1097">#1097</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/db1e716c63f98dad6520340f386ff3689491559c"><code>db1e716</code></a> build(deps): bump actions/checkout from 4 to 5 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1099">#1099</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.35.3...v0.36.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.14.0 to 9.14.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.14.1</h2>
<h2>Changes</h2>
<ul>
<li>fix(otel): Add support for filtering traces for certain commands (<a href="https://redirect.github.com/redis/go-redis/pull/3519">#3519</a>)</li>
<li>fix(pool): remove conn from idleConns if present (<a href="https://redirect.github.com/redis/go-redis/pull/3546">#3546</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/Sovietaced"><code>@​Sovietaced</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/v9.14.1/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.14.1 (2025-10-15)</h1>
<h2>Changes</h2>
<ul>
<li>fix(otel): Add support for filtering traces for certain commands (<a href="https://redirect.github.com/redis/go-redis/pull/3519">#3519</a>)</li>
<li>fix(pool): remove conn from idleConns if present (<a href="https://redirect.github.com/redis/go-redis/pull/3546">#3546</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/Sovietaced"><code>@​Sovietaced</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/512579f2e7582d5a248a5c221006e814a2c83e0b"><code>512579f</code></a> chore(release): bump version to 9.14.1 (<a href="https://redirect.github.com/redis/go-redis/issues/3548">#3548</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/c699a458bc0752854027679265cf857d4b3ae647"><code>c699a45</code></a> fix(pool): remove conn from idle conns if present (<a href="https://redirect.github.com/redis/go-redis/issues/3546">#3546</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/33c70b2f4ce800e72ea61f55ff3a0ccc10a6cb10"><code>33c70b2</code></a> fix: pipeline repeatedly sets the error (<a href="https://redirect.github.com/redis/go-redis/issues/3525">#3525</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/7706f88cea3c18c59fd3067a10c1f2a7e157703b"><code>7706f88</code></a> Add support for filtering traces for certain commands (<a href="https://redirect.github.com/redis/go-redis/issues/3519">#3519</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.14.0...v9.14.1">compare view</a></li>
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

### Review by @TenSt - Approved on October 20, 2025 at 09:52 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1243*
