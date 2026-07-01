---
type: pull_request
number: 544
title: "Build: Bump the go group with 2 updates"
state: closed
author: dependabot
created: 2024-01-22T15:35:27Z
updated: 2024-01-23T15:55:08Z
closed: 2024-01-23T15:55:06Z
base_branch: main
head_branch: dependabot/go_modules/go-cef819f8f1
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/544
---

# Pull Request #544: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: January 22, 2024 at 03:35 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-cef819f8f1`

## Description

Bumps the go group with 2 updates: [github.com/content-services/tang](https://github.com/content-services/tang) and [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go).

Updates `github.com/content-services/tang` from 0.0.1 to 0.0.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/tang/releases">github.com/content-services/tang's releases</a>.</em></p>
<blockquote>
<h2>v0.0.2</h2>
<h2>What's Changed</h2>
<ul>
<li>add comps groups searching by <a href="https://github.com/rverdile"><code>@​rverdile</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/3">content-services/tang#3</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.1...v0.0.2">https://github.com/content-services/tang/compare/v0.0.1...v0.0.2</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/cf719bd1a9c90f383d6c915fe10e126c9c1ba55d"><code>cf719bd</code></a> add comps groups searching (<a href="https://redirect.github.com/content-services/tang/issues/3">#3</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.1...v0.0.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2` from 2.14.0 to 2.15.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.15.0</h2>
<h1>Highlights 💫</h1>
<p>This release includes various updates and improvements such as README enhancements, dependency bumps, bug fixes, race condition resolutions, and protocol-related adjustments. Notable changes involve upgrading dependencies like grpc and go.opentelemetry, addressing race conditions, fixing Kafka test issues, and introducing new features like binary content mode for NATS and JetStream protocols. Additionally, there are governance documentation updates, link corrections, and improvements in error handling and documentation across different modules.</p>
<h1>Breaking 🚨</h1>
<p>The Kafka Sarama protocol now uses the <code>&quot;github.com/IBM/sarama&quot;</code> Go module import path.</p>
<h1>Commits 📄</h1>
<p>896e1d0 Update README.md
75ec0f2 Bump actions/setup-go from 4 to 5
41e80f7 fixed couple issues
9ccd339 bugfix_value_type_of_dataschema
c8cbca9 adds unique package name for import
f1bca09 relative .pb.go generation, go_package set to package name
c20eef2 bump the pahao mqtt to v0.12
ed7be6b Add WithCustomAttributes for PubSub
be31358 returning the error when doing a nack in the message
ecead5c Make a few comments a bit clearer
57be3cd Try to make sure the Receiver starts before we send events
f5c7061 Try to fix race again - don't reuse clients for sender/receiver
8bea925 Bump google.golang.org/grpc from 1.56.1 to 1.56.3 in /samples/http
fa6be00 Bump google.golang.org/grpc from 1.56.1 to 1.56.3 in /protocol/pubsub/v2
7e05ecd Bump google.golang.org/grpc from 1.56.1 to 1.56.3 in /samples/pubsub
13825ba Sleep less to avoid timeouts
3162d69 Bump github.com/nats-io/nats-server/v2 in /protocol/stan/v2
ec8b0f9 deps: update nats dependencies
dae9f6c Bump go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
1d6360b Bump go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
06658a2 Bump go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
7c1a3b1 fix race
6f5984b Move to go 1.18 Had to run gofmt and fix some weird typos due to tabs in the comments
0a006bb Fix race condition in kafka tests
510b002 issue 814 - Add binary content mode for NATS and JetStream protocols
ac3d30c add link to our security mailing list
9405398 Bump golang.org/x/net in /observability/opencensus/v2
3cbfae0 Bump golang.org/x/net from 0.9.0 to 0.17.0 in /protocol/pubsub/v2
65eb52e Bump golang.org/x/net from 0.12.0 to 0.17.0 in /protocol/kafka_sarama/v2
d25d6e4 Bump golang.org/x/net from 0.9.0 to 0.17.0 in /samples/pubsub
e4653a8 Bump golang.org/x/net from 0.12.0 to 0.17.0 in /test/conformance
6ed9f79 Bump golang.org/x/net from 0.9.0 to 0.17.0 in /samples/http
6a3393c Bump golang.org/x/net from 0.7.0 to 0.17.0 in /test/benchmark
806ef35 Bump golang.org/x/net from 0.12.0 to 0.17.0 in /samples/kafka
de13f1b Bump golang.org/x/net from 0.12.0 to 0.17.0 in /test/integration
3eefeb1 Governance docs per CE PR 1226
1bcaa28 Update links to cloudevents spec
6aa2742 context.Done() may never reach if waiting on r.incoming &lt;- msgErr
4bcddda move it to write message</p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/998919d6876d9561e43a1f55787aeabe92105e76"><code>998919d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1005">#1005</a> from Vedadiyan/main</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/896e1d07dda9f2988642de347c666b2433282b8e"><code>896e1d0</code></a> Update README.md</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/9d3236f13daa57378e54419c597dfd18e631e020"><code>9d3236f</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/989">#989</a> from cloudevents/dependabot/github_actions/actions/se...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/75ec0f2e1f1b1d0a765956ca4c1d07b847d05a44"><code>75ec0f2</code></a> Bump actions/setup-go from 4 to 5</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/e620cf71daa685ac84e318afeaf96ac3a17e686e"><code>e620cf7</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/986">#986</a> from timmyb32r/bugfix_value_type_of_dataschema2</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/41e80f74b16532322a1f38cb4df7424c4c6bbc99"><code>41e80f7</code></a> fixed couple issues</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/9ccd339056990f8fb4357e59da54b3f6e4782a86"><code>9ccd339</code></a> bugfix_value_type_of_dataschema</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/39fe13d4806f5815a524a916879b3a760fbdfd51"><code>39fe13d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/969">#969</a> from jnorman-us/fix/proto</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/c8cbca955d22362e9f313ffb2a87ecf26142714e"><code>c8cbca9</code></a> adds unique package name for import</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f1bca09cbeaba30a7a72735d839c99266bc6e0fc"><code>f1bca09</code></a> relative .pb.go generation, go_package set to package name</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.14.0...v2.15.0">compare view</a></li>
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

### Comment by @app-sre-bot on January 22, 2024 at 03:36 PM UTC

Can one of the admins verify this patch?

### Comment by @swadeley on January 22, 2024 at 06:51 PM UTC

/retest

### Comment by @dependabot on January 23, 2024 at 03:55 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/544*
