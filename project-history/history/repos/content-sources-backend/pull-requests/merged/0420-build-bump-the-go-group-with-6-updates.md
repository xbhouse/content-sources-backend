---
type: pull_request
number: 420
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2023-10-09T04:40:24Z
updated: 2023-10-13T14:26:17Z
closed: 2023-10-13T14:26:07Z
merged: 2023-10-13T14:26:07Z
base_branch: main
head_branch: dependabot/go_modules/go-c0a4e99dc0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/420
---

# Pull Request #420: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: October 09, 2023 at 04:40 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-c0a4e99dc0`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/spf13/viper](https://github.com/spf13/viper) | `1.16.0` | `1.17.0` |
| [github.com/archdx/zerolog-sentry](https://github.com/archdx/zerolog-sentry) | `1.5.0` | `1.6.1` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.21.0` | `1.21.1` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.13.40` | `1.13.42` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.24.0` | `1.24.1` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.24.1` | `0.25.0` |

Updates `github.com/spf13/viper` from 1.16.0 to 1.17.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/spf13/viper/releases">github.com/spf13/viper's releases</a>.</em></p>
<blockquote>
<h2>v1.17.0</h2>
<h2>Major changes</h2>
<p>Highlighting some of the changes for better visibility.</p>
<p>Please share your feedback in the Discussion forum. Thanks! ❤️</p>
<h3>Minimum Go version: 1.19</h3>
<p>Viper now requires Go 1.19</p>
<p>This change ensures we can stay up to date with modern practices and dependencies.</p>
<h3><code>log/slog</code> support <strong>[BREAKING]</strong></h3>
<p>Viper <a href="https://github.com/spf13/viper/releases/tag/v1.11.0">v1.11.0</a> added an experimental <code>Logger</code> interface to allow custom implementations (besides <a href="https://github.com/spf13/jwalterweatherman">jwalterweatherman</a>).</p>
<p>In addition, it also exposed an experimental <code>WithLogger</code> function allowing to set a custom logger.</p>
<p>This release deprecates that interface in favor of <a href="https://pkg.go.dev/log/slog">log/slog</a> released in Go 1.21.</p>
<blockquote>
<p>[!WARNING]
<code>WithLogger</code> accepts an <a href="https://pkg.go.dev/log/slog#Logger">*slog.Logger</a> from now on.</p>
</blockquote>
<p>To preserve backwards compatibility with older Go versions, prior to Go 1.21 Viper accepts a <a href="https://pkg.go.dev/golang.org/x/exp/slog#Logger">*golang.org/x/exp/slog.Logger</a>.</p>
<p>The experimental flag is removed.</p>
<h3>New finder implementation <strong>[BREAKING]</strong></h3>
<p>As of this release, Viper uses a new library to look for files, called <a href="https://github.com/sagikazarmark/locafero">locafero</a>.</p>
<p>The new library is better covered by tests and has been built from scratch as a general purpose file finder library.</p>
<p>The implementation is experimental and is hidden behind a <code>finder</code> build tag.</p>
<blockquote>
<p>[!WARNING]
The <code>io/fs</code> based implementation (that used to be hidden behind a <code>finder</code> build tag) has been removed.</p>
</blockquote>
<h2>What's Changed</h2>
<h3>Exciting New Features 🎉</h3>
<ul>
<li>Add NATS support by <a href="https://github.com/hooksie1"><code>@​hooksie1</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1590">spf13/viper#1590</a></li>
<li>Add slog support by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1627">spf13/viper#1627</a></li>
</ul>
<h3>Enhancements 🚀</h3>
<ul>
<li>chore: add local development environment using nix by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1572">spf13/viper#1572</a></li>
<li>feat: add func GetEnvPrefix by <a href="https://github.com/baruchiro"><code>@​baruchiro</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1565">spf13/viper#1565</a></li>
<li>Improve dev env by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1575">spf13/viper#1575</a></li>
<li>fix: code optimization by <a href="https://github.com/testwill"><code>@​testwill</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1557">spf13/viper#1557</a></li>
<li>test: remove not needed testutil.Setenv by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1610">spf13/viper#1610</a></li>
<li>new finder library based on afero by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/1625">spf13/viper#1625</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/spf13/viper/commit/f62f86a84b8395051efe0e490a29f3f89830a3ed"><code>f62f86a</code></a> refactor: make use of <code>strings.Cut</code></li>
<li><a href="https://github.com/spf13/viper/commit/94632fa21e01819de78bf0c931eb3cbdd1d426b4"><code>94632fa</code></a> chore: Use pip3 explicitly to install yamllint</li>
<li><a href="https://github.com/spf13/viper/commit/3f6cadcbeb9f193f8483c9edf6c6114ca7a20056"><code>3f6cadc</code></a> chore: Fix copy-paste error for yamllint target</li>
<li><a href="https://github.com/spf13/viper/commit/287507c0b5a13320f9b616e458ab7f3aadab1603"><code>287507c</code></a> docs: add set subset KV example</li>
<li><a href="https://github.com/spf13/viper/commit/f1cb2262bbda4cc549e617f41ba963409f23871c"><code>f1cb226</code></a> chore(deps): update crypt</li>
<li><a href="https://github.com/spf13/viper/commit/c292b55050aadcf08f9094033eb14d241f5002c5"><code>c292b55</code></a> test: refactor asserts</li>
<li><a href="https://github.com/spf13/viper/commit/3d006fe361ca4ea1ec973e39ce40f98aa3d2f2c9"><code>3d006fe</code></a> refactor: replace interface{} with any</li>
<li><a href="https://github.com/spf13/viper/commit/8a6dc5d43ce8b58288ec5b1db9ce83ab32580549"><code>8a6dc5d</code></a> build(deps): bump github/codeql-action from 2.21.8 to 2.21.9</li>
<li><a href="https://github.com/spf13/viper/commit/96c5c0083fb4a0e42282fbf04b626ca6792d8885"><code>96c5c00</code></a> chore: remove deprecated build tags</li>
<li><a href="https://github.com/spf13/viper/commit/44911d2cacff57448929c13f7346a1ea6ab13591"><code>44911d2</code></a> build(deps): bump github/codeql-action from 2.21.7 to 2.21.8</li>
<li>Additional commits viewable in <a href="https://github.com/spf13/viper/compare/v1.16.0...v1.17.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/archdx/zerolog-sentry` from 1.5.0 to 1.6.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/archdx/zerolog-sentry/releases">github.com/archdx/zerolog-sentry's releases</a>.</em></p>
<blockquote>
<h2>Close ErrFlushTimeout error</h2>
<p>ErrFlushTimeout error is returned from Close() method in case of flush error</p>
<h2>Sentry HTTP Client / MaxErrorDepth options</h2>
<p>No release notes provided.</p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/76f37725d6811ed602972033fd129a56d377d77c"><code>76f3772</code></a> Merge pull request <a href="https://redirect.github.com/archdx/zerolog-sentry/issues/12">#12</a> from Constructor-io/pe-671/error-on-flush-fail</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/520f730557c908ee256f96846e61b18820f76494"><code>520f730</code></a> fix: return error if flush fails</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/02582d2f681b5a8d6343cf5362c5f9be6a103d8a"><code>02582d2</code></a> Merge pull request <a href="https://redirect.github.com/archdx/zerolog-sentry/issues/11">#11</a> from Constructor-io/pe-657/http-cli-max-err-depth-cfg</li>
<li><a href="https://github.com/archdx/zerolog-sentry/commit/57dfb95a0010435f4b14575aff419fa650fc5290"><code>57dfb95</code></a> feat: http client &amp; max error depth configs</li>
<li>See full diff in <a href="https://github.com/archdx/zerolog-sentry/compare/v1.5.0...v1.6.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.21.0 to 1.21.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e155bb72a2ec20ec61db50fc3d4568e373fa4b63"><code>e155bb7</code></a> Release 2023-10-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9d342ba33937c562d215f317a37dea121ee9763d"><code>9d342ba</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1df99141a143a38570d64a182ed972ce9e3dba65"><code>1df9914</code></a> Update SDK's smithy-go dependency to v1.15.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32ada3a191ac770b1b24164b667692183fc77ed9"><code>32ada3a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/12ba4ac01f9e9be475ba983d7d185499acc658dc"><code>12ba4ac</code></a> Release 2023-10-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be8a8e0b31969b67b82bf0d045b71c8ffb32e6a6"><code>be8a8e0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dc38adb291c694a5e357c8b783f26e1c211a76b4"><code>dc38adb</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a52086ea4c9c3805aa5a07a7da9a8469d348dedb"><code>a52086e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1ed22c11618e32686959e583cb0c5a750ddab921"><code>1ed22c1</code></a> Release 2023-10-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e007bcdc649192a5515622ba2409ce54a9ef4f57"><code>e007bcd</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.21.0...v1.21.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.13.40 to 1.13.42
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e155bb72a2ec20ec61db50fc3d4568e373fa4b63"><code>e155bb7</code></a> Release 2023-10-06</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9d342ba33937c562d215f317a37dea121ee9763d"><code>9d342ba</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1df99141a143a38570d64a182ed972ce9e3dba65"><code>1df9914</code></a> Update SDK's smithy-go dependency to v1.15.0</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32ada3a191ac770b1b24164b667692183fc77ed9"><code>32ada3a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/12ba4ac01f9e9be475ba983d7d185499acc658dc"><code>12ba4ac</code></a> Release 2023-10-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/be8a8e0b31969b67b82bf0d045b71c8ffb32e6a6"><code>be8a8e0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dc38adb291c694a5e357c8b783f26e1c211a76b4"><code>dc38adb</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a52086ea4c9c3805aa5a07a7da9a8469d348dedb"><code>a52086e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1ed22c11618e32686959e583cb0c5a750ddab921"><code>1ed22c1</code></a> Release 2023-10-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e007bcdc649192a5515622ba2409ce54a9ef4f57"><code>e007bcd</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.13.40...credentials/v1.13.42">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.24.0 to 1.24.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/56f7cd4685f22d474411b33be826718a0ce9c56c"><code>56f7cd4</code></a> Release 2022-01-28</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/acd36849acb005e715a01553d1ca8873419cfff0"><code>acd3684</code></a> Update API Models (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1574">#1574</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/14f597921b4e948439857a039f422be0cd9d232e"><code>14f5979</code></a> config: Fix bug in SDK's merging of duration_seconds shared config (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1568">#1568</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a1bf7dd9733bb564b5b1a6c9e7f6b6e1ed5179c2"><code>a1bf7dd</code></a> Pre-allocate response body using Content-Length (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1565">#1565</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/452ee5e2e893ee7083116ac734e9e19f2c40db1b"><code>452ee5e</code></a> config: Update shared config loading to use os.UserHomeDir() (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1563">#1563</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/484478f8b7899157089e7da71126935c1c76d091"><code>484478f</code></a> config: Return error from optFns in LoadDefaultConfig (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1562">#1562</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/66251136b55d3a64355b88883f226f1d24a40bc7"><code>6625113</code></a> Remove negative smoke test for WAF (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1564">#1564</a>)</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.24.0...service/s3/v1.24.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.24.1 to 0.25.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.25.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.25.0.</p>
<h3>Breaking Changes</h3>
<p>As previously announced, this release removes two global constants from the SDK.</p>
<ul>
<li><code>sentry.Version</code> was removed. Use <code>sentry.SDKVersion</code> instead (<a href="https://redirect.github.com/getsentry/sentry-go/pull/727">#727</a>)</li>
<li><code>sentry.SDKIdentifier</code> was removed. Use <code>Client.GetSDKIdentifier()</code> instead (<a href="https://redirect.github.com/getsentry/sentry-go/pull/727">#727</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>Add <code>ClientOptions.IgnoreTransactions</code>, which allows you to ignore specific transactions based on their name (<a href="https://redirect.github.com/getsentry/sentry-go/pull/717">#717</a>)</li>
<li>Add <code>ClientOptions.Tags</code>, which allows you to set global tags that are applied to all events. You can also define tags by setting <code>SENTRY_TAGS_</code> environment variables (<a href="https://redirect.github.com/getsentry/sentry-go/pull/718">#718</a>)</li>
</ul>
<h3>Bug fixes</h3>
<ul>
<li>Fix an issue in the profiler that would cause an infinite loop if the duration of a transaction is longer than 30 seconds (<a href="https://redirect.github.com/getsentry/sentry-go/issues/724">#724</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li><code>dsn.RequestHeaders()</code> is not to be removed, though it is still considered deprecated and should only be used when using a custom transport that sends events to the <code>/store</code> endpoint (<a href="https://redirect.github.com/getsentry/sentry-go/pull/720">#720</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.25.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.25.0.</p>
<h3>Breaking Changes</h3>
<p>As previously announced, this release removes two global constants from the SDK.</p>
<ul>
<li><code>sentry.Version</code> was removed. Use <code>sentry.SDKVersion</code> instead (<a href="https://redirect.github.com/getsentry/sentry-go/pull/727">#727</a>)</li>
<li><code>sentry.SDKIdentifier</code> was removed. Use <code>Client.GetSDKIdentifier()</code> instead (<a href="https://redirect.github.com/getsentry/sentry-go/pull/727">#727</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>Add <code>ClientOptions.IgnoreTransactions</code>, which allows you to ignore specific transactions based on their name (<a href="https://redirect.github.com/getsentry/sentry-go/pull/717">#717</a>)</li>
<li>Add <code>ClientOptions.Tags</code>, which allows you to set global tags that are applied to all events. You can also define tags by setting <code>SENTRY_TAGS_</code> environment variables (<a href="https://redirect.github.com/getsentry/sentry-go/pull/718">#718</a>)</li>
</ul>
<h3>Bug fixes</h3>
<ul>
<li>Fix an issue in the profiler that would cause an infinite loop if the duration of a transaction is longer than 30 seconds (<a href="https://redirect.github.com/getsentry/sentry-go/issues/724">#724</a>)</li>
</ul>
<h3>Misc</h3>
<ul>
<li><code>dsn.RequestHeaders()</code> is not to be removed, though it is still considered deprecated and should only be used when using a custom transport that sends events to the <code>/store</code> endpoint (<a href="https://redirect.github.com/getsentry/sentry-go/pull/720">#720</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/03fdca5324bff5c699a3ca0ee0dfdadda70f6fc2"><code>03fdca5</code></a> release: 0.25.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/382a4959d464e654cdd45eddccf80c8b973fb424"><code>382a495</code></a> Prepare 0.25.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/728">#728</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2a763487b0f5ee0243a45140753ab1a97c6963a9"><code>2a76348</code></a> Remove deprecated code (<a href="https://redirect.github.com/getsentry/sentry-go/issues/727">#727</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f1f017663e9cec36c09d26bb2aa4dcf17b22b975"><code>f1f0176</code></a> fix: profiler infinite loop after a long transaction (<a href="https://redirect.github.com/getsentry/sentry-go/issues/725">#725</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7e37edeb9589c5e8966c5c80822aa4ddaa0fbfd2"><code>7e37ede</code></a> Add the SDK version when functions will be removed to <code>Deprecated:</code> comments ...</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/3b3ce90268e898b7b5253f720c8d195d73435a27"><code>3b3ce90</code></a> Clarify the usage of <code>dsn.RequestHeaders()</code> (<a href="https://redirect.github.com/getsentry/sentry-go/issues/720">#720</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/68e2c0f78294ef8230303a14373abd212308ea3a"><code>68e2c0f</code></a> feat: global tags (<a href="https://redirect.github.com/getsentry/sentry-go/issues/718">#718</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/8bf9a1582defccf747644e0155deac65a855ce21"><code>8bf9a15</code></a> feat: support IgnoreTransactions client option (<a href="https://redirect.github.com/getsentry/sentry-go/issues/717">#717</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a83590750c3d7ec01592a6cb427bfb03bb1c8c4a"><code>a835907</code></a> Merge branch 'release/0.24.1'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.24.1...v0.25.0">compare view</a></li>
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

### Comment by @app-sre-bot on October 09, 2023 at 04:40 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 10, 2023 at 05:26 PM UTC

[test]

### Comment by @jlsherrill on October 11, 2023 at 03:07 PM UTC

/retest

### Comment by @jlsherrill on October 13, 2023 at 12:33 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on October 13, 2023 at 02:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/420*
