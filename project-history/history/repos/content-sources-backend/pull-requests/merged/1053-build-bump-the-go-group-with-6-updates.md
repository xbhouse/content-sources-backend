---
type: pull_request
number: 1053
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2025-03-31T04:06:49Z
updated: 2025-04-04T15:51:36Z
closed: 2025-04-04T15:51:29Z
merged: 2025-04-04T15:51:28Z
base_branch: main
head_branch: dependabot/go_modules/go-69775e0d40
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1053
---

# Pull Request #1053: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: March 31, 2025 at 04:06 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-69775e0d40`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/spf13/viper](https://github.com/spf13/viper) | `1.20.0` | `1.20.1` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.9` | `1.29.12` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.62` | `1.17.65` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.15.2` | `2.16.0` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.3.1741891889` | `2025.3.1743021229` |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | `5.7.3` | `5.7.4` |

Updates `github.com/spf13/viper` from 1.20.0 to 1.20.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/spf13/viper/releases">github.com/spf13/viper's releases</a>.</em></p>
<blockquote>
<h2>v1.20.1</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Backport config type fixes to 1.20.x by <a href="https://github.com/sagikazarmark"><code>@​sagikazarmark</code></a> in <a href="https://redirect.github.com/spf13/viper/pull/2005">spf13/viper#2005</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/spf13/viper/compare/v1.20.0...v1.20.1">https://github.com/spf13/viper/compare/v1.20.0...v1.20.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/spf13/viper/commit/9568cfcfd660a1c1c6c762f335ae79f370488417"><code>9568cfc</code></a> fix: config type check when loading any config</li>
<li><a href="https://github.com/spf13/viper/commit/fd05140cd675b256b5076759e38403874fc5826d"><code>fd05140</code></a> fix(config): get config type from v.configType or config file ext</li>
<li>See full diff in <a href="https://github.com/spf13/viper/compare/v1.20.0...v1.20.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.9 to 1.29.12
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5686f5441d6cf11a87ca8f623cce20948572326"><code>a5686f5</code></a> Release 2025-03-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78e6f022956ee16f32164ad5a0bb01f3d99a5ff2"><code>78e6f02</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54113564fe3e694e58e4956680a6d343a871542f"><code>5411356</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c9ca62d168d26b375f07a48b07e4cc9c83ce5ef1"><code>c9ca62d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad4fc4c28b5e589872821d34c8e3a25cc6bfe6f3"><code>ad4fc4c</code></a> Release 2025-03-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e181e119e8faa478dd41ee2d5f82f4d00e82cc80"><code>e181e11</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/486b615bf2c2f56fdf0e67a08a0be56efe54f55a"><code>486b615</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04fed2fac6507648d379818fe363124f3462cb98"><code>04fed2f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6fa167adb5d1a2618a2f9bbe8f2b885c0d7e2893"><code>6fa167a</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2996">#2996</a> from aws/feat-tmv2-getobject</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ba02d233a98ac6f6438f13d7abb3238ce26859d"><code>8ba02d2</code></a> Release 2025-03-25</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.9...config/v1.29.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.62 to 1.17.65
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a5686f5441d6cf11a87ca8f623cce20948572326"><code>a5686f5</code></a> Release 2025-03-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78e6f022956ee16f32164ad5a0bb01f3d99a5ff2"><code>78e6f02</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54113564fe3e694e58e4956680a6d343a871542f"><code>5411356</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c9ca62d168d26b375f07a48b07e4cc9c83ce5ef1"><code>c9ca62d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ad4fc4c28b5e589872821d34c8e3a25cc6bfe6f3"><code>ad4fc4c</code></a> Release 2025-03-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e181e119e8faa478dd41ee2d5f82f4d00e82cc80"><code>e181e11</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/486b615bf2c2f56fdf0e67a08a0be56efe54f55a"><code>486b615</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/04fed2fac6507648d379818fe363124f3462cb98"><code>04fed2f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6fa167adb5d1a2618a2f9bbe8f2b885c0d7e2893"><code>6fa167a</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2996">#2996</a> from aws/feat-tmv2-getobject</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ba02d233a98ac6f6438f13d7abb3238ce26859d"><code>8ba02d2</code></a> Release 2025-03-25</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.62...credentials/v1.17.65">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/v2` from 2.15.2 to 2.16.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cloudevents/sdk-go/releases">github.com/cloudevents/sdk-go/v2's releases</a>.</em></p>
<blockquote>
<h2>Release v2.16.0</h2>
<h3>✨ Features &amp; Enhancements</h3>
<ul>
<li>
<p><strong>Confluent Kafka binding</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/988">#988</a> by <a href="https://github.com/yanmxa"><code>@​yanmxa</code></a><br />
Added a new Confluent Kafka protocol binding for CloudEvents, supporting modern Kafka client features.</p>
</li>
<li>
<p><strong>Producer report channel for Confluent Kafka</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1031">#1031</a> by <a href="https://github.com/yanmxa"><code>@​yanmxa</code></a><br />
Exposed a producer report channel via <code>Events()</code> to allow users to track delivery status of Kafka messages.</p>
</li>
<li>
<p><strong>Support structured content type suffixes</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1007">#1007</a> by <a href="https://github.com/dan-j"><code>@​dan-j</code></a><br />
Improved content type handling by recognizing structured syntax suffixes like <code>+json</code>, increasing compatibility with various encodings.</p>
</li>
<li>
<p><strong>Default timeout via context</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/992">#992</a> by <a href="https://github.com/nkreiger"><code>@​nkreiger</code></a><br />
Introduced support for configuring protocol default timeouts using <code>context.Context</code>.</p>
</li>
<li>
<p><strong>Benchmarks for CESQL</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1050">#1050</a> by <a href="https://github.com/Cali0707"><code>@​Cali0707</code></a><br />
Added benchmark tests to measure CESQL query performance.</p>
</li>
<li>
<p><strong>Optimized CESQL <code>LIKE</code> matching</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1049">#1049</a> by <a href="https://github.com/Cali0707"><code>@​Cali0707</code></a><br />
Improved the performance of CESQL's <code>LIKE</code> pattern matching logic.</p>
</li>
<li>
<p><strong>Expose AddFunction API for CESQL Parser</strong> [#1047 / <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1051">#1051</a>] by <a href="https://github.com/dgeorgievski"><code>@​dgeorgievski</code></a><br />
Enabled users to register custom functions in CESQL via the newly exposed <code>AddFunction</code> API.</p>
</li>
<li>
<p><strong>Flexible subject matching for NATS JetStream</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1084">#1084</a> by <a href="https://github.com/evankanderson"><code>@​evankanderson</code></a><br />
Added support for flexible subject pattern matching in NATS JetStream subscriptions.</p>
</li>
<li>
<p><strong>Add v3 version of NATS JetStream protocol</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1095">#1095</a> by <a href="https://github.com/stephen-totty-hpe"><code>@​stephen-totty-hpe</code></a><br />
Introduced a new version of the NATS JetStream protocol (v3) with enhanced features and forward compatibility.</p>
</li>
<li>
<p><strong>Expose <code>WithHost</code> option</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1070">#1070</a> by <a href="https://github.com/jaxtonw"><code>@​jaxtonw</code></a><br />
Added a configurable <code>WithHost</code> option for improved protocol initialization flexibility.</p>
</li>
<li>
<p><strong>Support <code>dataref</code> cloud event extension</strong> [[Dataref Extension <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1018">#1018</a>](https://redirect.github.com/cloudevents/sdk-go/pull/1018)] by <a href="https://github.com/matzew"><code>@​matzew</code></a>
Implements the Dataref (Claim Check Pattern) as specified by the CloudEvent Extension Attributes <a href="https://github.com/cloudevents/spec/blob/main/cloudevents/extensions/dataref.md">spec</a>.</p>
</li>
</ul>
<hr />
<h3>🛠 Fixes</h3>
<ul>
<li>
<p><strong>Handle multiple AMQP data fields correctly</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1034">#1034</a> by <a href="https://github.com/embano1"><code>@​embano1</code></a><br />
Fixed parsing of AMQP messages containing multiple data fields to conform with spec expectations.</p>
</li>
<li>
<p><strong>Fix invalid <code>ce-</code> prefix in Confluent binding</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1059">#1059</a> by <a href="https://github.com/embano1"><code>@​embano1</code></a><br />
Corrected an issue where CloudEvents extensions were incorrectly prefixed in the Confluent Kafka binding.</p>
</li>
<li>
<p><strong>Fix LIKE expression error handling</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1046">#1046</a> by <a href="https://github.com/Cali0707"><code>@​Cali0707</code></a><br />
Prevented panics on malformed <code>LIKE</code> expressions in CESQL; now returns a parse error instead.</p>
</li>
<li>
<p><strong>Fix MQTT content-type issue</strong> <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1063">#1063</a> by <a href="https://github.com/yanmxa"><code>@​yanmxa</code></a></p>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/6de37de32b83aee0bc53fe31b6358d4e5f1e9ede"><code>6de37de</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1139">#1139</a> from duglin/upgradeLint</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/25cdf948a23a5ac98fd8db61d537b0d6770dc34f"><code>25cdf94</code></a> upgrade-lint</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/240eb029b2003a4a998c1495469a1167dfe8eb63"><code>240eb02</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1137">#1137</a> from duglin/fix-automerge</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/7b6701c5c1dc330f10714cec441cb203312a1e83"><code>7b6701c</code></a> fix automerge by giving better names</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/d3e82a74062596ccfb210315f48527d47a4fe036"><code>d3e82a7</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1136">#1136</a> from cloudevents/dependabot/github_actions/golangci/...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/00e3d7cf14fc4c944d844b5202484e4ff59783a6"><code>00e3d7c</code></a> chore(deps): Bump golangci/golangci-lint-action from 6.5.2 to 7.0.0</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/69cfc2dfa3beb4632630748017e6f8c0f682a28a"><code>69cfc2d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1018">#1018</a> from matzew/add_dataref_extension</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/535da92997c5a762db09111ad6777e5baecf68f3"><code>535da92</code></a> Adding a simple dataref extension, similar to the java sdk</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/74ac76d56f0d899827f940b1a19079014f90e749"><code>74ac76d</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1134">#1134</a> from cloudevents/dependabot/go_modules/observability...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/6da042f56d387e81193b5e27692f9abb3fcf1950"><code>6da042f</code></a> chore: run go mod tidy</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.15.2...v2.16.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.3.1741891889 to 2025.3.1743021229
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/6f9dd5fc1033aaef0b8d3a3de19fca69246d2648"><code>6f9dd5f</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7dd9e4545d157952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/36c09a729c7290739c89122bf8ac507160a9ccb8"><code>36c09a7</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276ebd548d8027b8e3dae95843...</li>
<li><a href="https://github.com/content-services/zest/commit/fbc8426462e9b1948421027b1c20409117b372c5"><code>fbc8426</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e47568eed35a0b7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/1a8878884468f42a6ae63613754c3eebfd7fccf5"><code>1a88788</code></a> Update pulp bindings to b59d5a8be2d6a9584b32d696a45276eb6238dbc27b8e3dae95843...</li>
<li><a href="https://github.com/content-services/zest/commit/94b5fdfaf9b7f8187223d27ce83659b59e1cd830"><code>94b5fdf</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a45e9ddb30b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/a962ca6fb6e5cc1792888ab78dbf1e0c47bdc43d"><code>a962ca6</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27a45585545cb7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/34457802360c4094c0c4db5db7c081d1582856cd"><code>3445780</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a972adaa946e057bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/5710f44ba1517866585788908a7429476d8ffd7a"><code>5710f44</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27ab43a26d31b7a8d53bd49838...</li>
<li><a href="https://github.com/content-services/zest/commit/27419a2a99edcc929aef2aabcb8f06c6e780b894"><code>27419a2</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97994b89a5df37ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/a9c2ad67208ca1ef0fca90779b1b43d1cbb3d13e"><code>a9c2ad6</code></a> Update pulp bindings to d6a984592e569b42a26e536b94e27ab4245ea8cb7a8d53bd49838...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.3.1741891889...release/v2025.3.1743021229">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/jackc/pgx/v5` from 5.7.3 to 5.7.4
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jackc/pgx/blob/master/CHANGELOG.md">github.com/jackc/pgx/v5's changelog</a>.</em></p>
<blockquote>
<h1>5.7.4 (March 24, 2025)</h1>
<ul>
<li>Fix / revert change to scanning JSON <code>null</code> (Felix Röhrich)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jackc/pgx/commit/04bcc0219dc3acf67f27e68decd6dffe97334779"><code>04bcc02</code></a> Add v5.7.4 to changelog</li>
<li><a href="https://github.com/jackc/pgx/commit/0e0a7d8344b035b96be53eda85e11c0aff14f212"><code>0e0a7d8</code></a> Merge pull request <a href="https://redirect.github.com/jackc/pgx/issues/2288">#2288</a> from felix-roehrich/fr/fix-plan-scan</li>
<li><a href="https://github.com/jackc/pgx/commit/63422c7d6cfe092af402f48e16729acd1e3bae1c"><code>63422c7</code></a> revert change in if</li>
<li>See full diff in <a href="https://github.com/jackc/pgx/compare/v5.7.3...v5.7.4">compare view</a></li>
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

### Comment by @jlsherrill on March 31, 2025 at 06:53 PM UTC

/retest

### Comment by @jlsherrill on March 31, 2025 at 09:16 PM UTC

/retest

### Comment by @jlsherrill on April 01, 2025 at 12:47 PM UTC

/retest

### Comment by @rverdile on April 02, 2025 at 07:35 PM UTC

@dependabot rebase

### Comment by @dependabot on April 02, 2025 at 07:35 PM UTC

Looks like this PR has been edited by someone other than Dependabot. That means Dependabot can't rebase it - sorry!

If you're happy for Dependabot to recreate it from scratch, overwriting any edits, you can request `@dependabot recreate`.


---

## Reviews

### Review by @rverdile - Approved on April 04, 2025 at 03:47 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1053*
