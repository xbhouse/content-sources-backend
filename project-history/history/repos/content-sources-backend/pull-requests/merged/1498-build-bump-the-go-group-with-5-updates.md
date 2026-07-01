---
type: pull_request
number: 1498
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2026-05-18T05:56:46Z
updated: 2026-05-18T09:53:15Z
closed: 2026-05-18T09:53:13Z
merged: 2026-05-18T09:53:13Z
base_branch: main
head_branch: dependabot/go_modules/go-c3c8c60276
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1498
---

# Pull Request #1498: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: May 18, 2026 at 05:56 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-c3c8c60276`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.48.1` | `1.48.2` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.73.0` | `1.74.0` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.8.0` | `4.8.1` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.4.1777574211` | `2026.5.1778263552` |
| [google.golang.org/grpc](https://github.com/grpc/grpc-go) | `1.81.0` | `1.81.1` |

Updates `github.com/IBM/sarama` from 1.48.1 to 1.48.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.48.2 (2026-05-13)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(admin): add KIP-396 list/alter offsets APIs by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3419">IBM/sarama#3419</a></li>
<li>feat: add SubscriptionUserDataProvider hook for BalanceStrategy by <a href="https://github.com/lizthegrey"><code>@​lizthegrey</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3506">IBM/sarama#3506</a></li>
<li>perf(zstd): scale idle zstd encoder cap to GOMAXPROCS by <a href="https://github.com/lizthegrey"><code>@​lizthegrey</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3507">IBM/sarama#3507</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: retry ListTopics on transient transport errors by <a href="https://github.com/huynhanx03"><code>@​huynhanx03</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3497">IBM/sarama#3497</a></li>
<li>test(fvt): only safeClose if we created by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3530">IBM/sarama#3530</a></li>
<li>fix(client): scope metadata refresh errors to requested topics by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3532">IBM/sarama#3532</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>test(fvt): speedup functional test runs by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3528">IBM/sarama#3528</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.48.1...v1.48.2">https://github.com/IBM/sarama/compare/v1.48.1...v1.48.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/b8cdb3e7f27b556a74585442c0f1cf4328873d4b"><code>b8cdb3e</code></a> perf(zstd): scale idle zstd encoder cap to GOMAXPROCS (<a href="https://redirect.github.com/IBM/sarama/issues/3507">#3507</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/fa624c04a746c7682770a708df0f9244bd971106"><code>fa624c0</code></a> fix(client): scope metadata refresh errors to requested topics (<a href="https://redirect.github.com/IBM/sarama/issues/3532">#3532</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/6677d9814f6b274266e778eb2d11913ed58b2426"><code>6677d98</code></a> feat: add SubscriptionUserDataProvider hook for BalanceStrategy (<a href="https://redirect.github.com/IBM/sarama/issues/3506">#3506</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/0fdab5b08a667cc36429c39698c2aed994a43989"><code>0fdab5b</code></a> chore(ci): bump actions/dependency-review-action from 4.9.0 to 5.0.0 in the a...</li>
<li><a href="https://github.com/IBM/sarama/commit/22fad66496adc79028a3ca874fa3f47c118c070c"><code>22fad66</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.7 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/5a758233a81b9b8efb9f15c0dfa31c59ba64ba42"><code>5a75823</code></a> test(fvt): only safeClose if we created</li>
<li><a href="https://github.com/IBM/sarama/commit/f24d38d20cef6b16082d954ed538755f1896237a"><code>f24d38d</code></a> test(fvt): tighten docker bootstrap and topic prep polling</li>
<li><a href="https://github.com/IBM/sarama/commit/eeb273b446f6172d5a1417cf37b6b4177144c775"><code>eeb273b</code></a> test(fvt): trim TestFuncConsumerGroupFuzzy volume</li>
<li><a href="https://github.com/IBM/sarama/commit/5ed4e14914c17b3c4e265893b845367357750525"><code>5ed4e14</code></a> test(fvt): run independent FVT tests in parallel</li>
<li><a href="https://github.com/IBM/sarama/commit/5b0e5f0530630c448386832b23aee1e26593e72a"><code>5b0e5f0</code></a> test(fvt): tighten consumer group default timeouts</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.48.1...v1.48.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.73.0 to 1.74.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ad9d5996fd752f0756be2dbbdd4f8a4841fe362"><code>4ad9d59</code></a> Release 2025-01-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5ad8437a282f45f876255b6410b00ad74b512796"><code>5ad8437</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/903ab00324d065403132ea5e19abc51c7b5bd0ac"><code>903ab00</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/847821e55367b6185db008b95c72887767949b07"><code>847821e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d2c7a1901f52750bad242315cef97902048b8fdf"><code>d2c7a19</code></a> Update README.md (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2980">#2980</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e5689c9ef90df930b7634fb952411db6a6c00cab"><code>e5689c9</code></a> add client config to disable logging skipped output checksum validation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2976">#2976</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/560c44d97a1f92bf73a6d5e54b6f458edb92080a"><code>560c44d</code></a> Release 2025-01-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c45f0a58aac89ca19f2a8801b63346b8c2c61d70"><code>c45f0a5</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54e903f52e99876a6c131e2eb314ae83c85c9c3c"><code>54e903f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3464887a31a8d47392c7c9ee9e1b2ce70fa7a27d"><code>3464887</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.73.0...service/s3/v1.74.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.8.0 to 4.8.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/f281e2f1df9dc5445499b3764a6e5055edc2aaa8"><code>f281e2f</code></a> Update candlepin bindings to release/v4.8.1</li>
<li><a href="https://github.com/content-services/caliri/commit/8e76d319e00996c1b2d61fdf37d3ac974b76987c"><code>8e76d31</code></a> Update candlepin bindings to release/v4.8.0</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.8.0...release/v4.8.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.4.1777574211 to 2026.5.1778263552
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/8f6ff5104fd5795cb2b54079ed12d904f7ac3498"><code>8f6ff51</code></a> Update pulp bindings to e69db48356e528a464be3da896237b3d294bb531a7d54eb5892a9...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.4.1777574211...release/v2026.5.1778263552">compare view</a></li>
</ul>
</details>
<br />

Updates `google.golang.org/grpc` from 1.81.0 to 1.81.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/grpc/grpc-go/releases">google.golang.org/grpc's releases</a>.</em></p>
<blockquote>
<h2>Release 1.81.1</h2>
<h1>Security</h1>
<ul>
<li>xds/rbac: Fix a potential authorization bypass caused by incorrectly falling through URI/DNS SANs to Subject Distinguished Name (DN) when matching the authenticated principal name. With this fix, only the first non-empty identity source will be used, as per <a href="https://github.com/grpc/proposal/blob/master/A41-xds-rbac.md">gRFC A41</a>. (<a href="https://redirect.github.com/grpc/grpc-go/issues/9111">#9111</a>)
<ul>
<li>Special Thanks: <a href="https://github.com/al4an444"><code>@​al4an444</code></a></li>
</ul>
</li>
</ul>
<h1>Bug Fixes</h1>
<ul>
<li>otel: Segregate client and server RPC information used for metrics and traces, to avoid one overwriting the other. (<a href="https://redirect.github.com/grpc/grpc-go/issues/9081">#9081</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/grpc/grpc-go/commit/caf0772c2bcb8bc15d43eb53448e921f34f0b7e8"><code>caf0772</code></a> Change version from 1.81.1-dev to 1.81.1 (<a href="https://redirect.github.com/grpc/grpc-go/issues/9122">#9122</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/6ccbeebf058ede71e43a5ac28fada2a736573215"><code>6ccbeeb</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/9111">#9111</a> into v1.81.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/9121">#9121</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/b33c29e41b438e371c8504de9bdf64a80098cc29"><code>b33c29e</code></a> Cherry-pick <a href="https://redirect.github.com/grpc/grpc-go/issues/9081">#9081</a> into v1.81.x (<a href="https://redirect.github.com/grpc/grpc-go/issues/9102">#9102</a>)</li>
<li><a href="https://github.com/grpc/grpc-go/commit/c45fae6d06a5c192b7b96418a2bc26a96b856834"><code>c45fae6</code></a> Change version to 1.81.1-dev (<a href="https://redirect.github.com/grpc/grpc-go/issues/9063">#9063</a>)</li>
<li>See full diff in <a href="https://github.com/grpc/grpc-go/compare/v1.81.0...v1.81.1">compare view</a></li>
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
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on May 18, 2026 at 09:53 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1498*
