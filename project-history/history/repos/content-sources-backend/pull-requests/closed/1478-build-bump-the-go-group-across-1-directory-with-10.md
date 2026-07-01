---
type: pull_request
number: 1478
title: "Build: Bump the go group across 1 directory with 10 updates"
state: closed
author: dependabot
created: 2026-04-27T04:59:26Z
updated: 2026-04-27T15:42:21Z
closed: 2026-04-27T15:42:19Z
base_branch: main
head_branch: dependabot/go_modules/go-cd980accfd
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1478
---

# Pull Request #1478: Build: Bump the go group across 1 directory with 10 updates

**Author**: @dependabot
**Created**: April 27, 2026 at 04:59 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-cd980accfd`

## Description

[//]: # (dependabot-start)
⚠️  **Dependabot is rebasing this PR** ⚠️ 

Rebasing might not happen immediately, so don't worry if this takes some time.

Note: if you make any changes to this PR yourself, they will take precedence over the rebase.

---

[//]: # (dependabot-end)

Bumps the go group with 10 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.135.0` | `0.136.0` |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.35.0` | `1.35.1` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.47.0` | `1.48.0` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.69.1` | `1.70.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.99.1` | `1.100.0` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.7.5` | `4.8.0` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.4.1775836578` | `2026.4.1777255680` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.45.1` | `0.46.0` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.7.0` | `1.8.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.27.0` | `1.28.0` |


Updates `github.com/getkin/kin-openapi` from 0.135.0 to 0.136.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.136.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3: stop injecting contentless default in NewResponses() by <a href="https://github.com/0-don"><code>@​0-don</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1148">getkin/kin-openapi#1148</a></li>
<li>openapi3: standardize Origin json tag to &quot;-&quot; across all types by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1149">getkin/kin-openapi#1149</a></li>
<li>Update usage message in cmd/validate by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1150">getkin/kin-openapi#1150</a></li>
<li>openapi3: fix determinism when handling discriminator mappings by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1151">getkin/kin-openapi#1151</a></li>
<li>feat: bump Go to 1.26 by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1152">getkin/kin-openapi#1152</a></li>
<li>openapi3: use componentNames for deterministic visitings by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1153">getkin/kin-openapi#1153</a></li>
<li>feat: add OpenAPI 3.1 support by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1125">getkin/kin-openapi#1125</a></li>
<li>openapi3: add JoinFunc for custom $ref path resolution by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1154">getkin/kin-openapi#1154</a></li>
<li>Add many many tests from ApisGuruOpenapiDirectory by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1155">getkin/kin-openapi#1155</a></li>
<li>openapi3: remove map-iteration order leaks causing flaky tests by <a href="https://github.com/cloudnativeninja"><code>@​cloudnativeninja</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1158">getkin/kin-openapi#1158</a></li>
<li>openapi2conv: nil-guard components lookup in FromV3SchemaRef by <a href="https://github.com/SAY-5"><code>@​SAY-5</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1156">getkin/kin-openapi#1156</a></li>
<li>Address various lint errors by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1157">getkin/kin-openapi#1157</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/0-don"><code>@​0-don</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1148">getkin/kin-openapi#1148</a></li>
<li><a href="https://github.com/cloudnativeninja"><code>@​cloudnativeninja</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1158">getkin/kin-openapi#1158</a></li>
<li><a href="https://github.com/SAY-5"><code>@​SAY-5</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1156">getkin/kin-openapi#1156</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.135.0...v0.136.0">https://github.com/getkin/kin-openapi/compare/v0.135.0...v0.136.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/ff4bce724cfa3fdf6fac4470c76a8511662b21f6"><code>ff4bce7</code></a> fix and upgrade goimports-reviser</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/028df2a310efbbd795fe0d9a20bbe681b3d551e4"><code>028df2a</code></a> refacto(tests): use t.Context instead of context.Background</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/cc4f8d9903cdec697ca5291bb27040396050dc9f"><code>cc4f8d9</code></a> refacto: replace <code>openapi3.*Ptr(..)</code> funcs with <code>new(..)</code></li>
<li><a href="https://github.com/getkin/kin-openapi/commit/df95b871780abc8cd8cf6a694f7565dfdd056eeb"><code>df95b87</code></a> address various lint errors</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/355692920adc95a261f8afba887dce636a0e2d67"><code>3556929</code></a> openapi2conv: nil-guard components lookup in FromV3SchemaRef (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1156">#1156</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/5a0a3373598fa8184f7214e0af51164533130c7d"><code>5a0a337</code></a> openapi3: remove map-iteration order leaks causing flaky tests (<a href="https://redirect.github.com/getkin/kin-openapi/issues/1158">#1158</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/3489553bbab5fb309aa0dce6bf4bc077d2fa6f70"><code>3489553</code></a> openapi3: skip v3.1 load/validation flaky tests</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/3aa08cdb36fb98a2d2433cd581b1adf59bf97e35"><code>3aa08cd</code></a> openapi3: record v3.1 load/validation test failures</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/31797753134d75aa411c00d2e8ffcb3f669a0187"><code>3179775</code></a> openapi3: enable testing for 3.1 documents</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/77e911a361248f421c22e4478a6cf3950c3ab694"><code>77e911a</code></a> openapi3: record failures after rebasing on top of <a href="https://redirect.github.com/getkin/kin-openapi/issues/1125">#1125</a></li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.135.0...v0.136.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/rs/zerolog` from 1.35.0 to 1.35.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/rs/zerolog/commit/116c8060e034e8d46855354d22db2acbc8df9e1e"><code>116c806</code></a> event: restore Err() logging when ErrorStackMarshaler returns nil (<a href="https://redirect.github.com/rs/zerolog/issues/763">#763</a>)</li>
<li>See full diff in <a href="https://github.com/rs/zerolog/compare/v1.35.0...v1.35.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.47.0 to 1.48.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.48.0 (2026-04-24)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:tada: New Features / Improvements</h3>
<ul>
<li>feat(producer): partition muting for msg ordering by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3422">IBM/sarama#3422</a></li>
</ul>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: handle nullable metadata in OffsetFetchResponse by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3473">IBM/sarama#3473</a></li>
<li>fix: nil response/done channels after SASLv1 failure by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3474">IBM/sarama#3474</a></li>
<li>fix(protocol): handle ElectLeaders V1 non-flexible headers by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3478">IBM/sarama#3478</a></li>
<li>fix: correct a number of goroutine leaks by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3476">IBM/sarama#3476</a></li>
<li>fix: resolve deadlock in concurrent offset commits by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3477">IBM/sarama#3477</a></li>
<li>fix(consumer): avoid broker race in response feeder by <a href="https://github.com/DCjanus"><code>@​DCjanus</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3486">IBM/sarama#3486</a></li>
<li>fix: stop dispatcher for dying children in brokerConsumer.abort() by <a href="https://github.com/lizthegrey"><code>@​lizthegrey</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3492">IBM/sarama#3492</a></li>
<li>fix: close broken tcp connections by <a href="https://github.com/Asphaltt"><code>@​Asphaltt</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3384">IBM/sarama#3384</a></li>
<li>fix: add Unwrap() to DescribeConfigError and AlterConfigError by <a href="https://github.com/ShinThirty"><code>@​ShinThirty</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3487">IBM/sarama#3487</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): update dependency golangci/golangci-lint to v2.11.1 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3462">IBM/sarama#3462</a></li>
<li>chore(deps): bump github.com/pierrec/lz4/v4 from 4.1.25 to 4.1.26 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3461">IBM/sarama#3461</a></li>
<li>chore(deps): bump golang.org/x/sync from 0.19.0 to 0.20.0 in the golang-x group across 1 directory by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3466">IBM/sarama#3466</a></li>
<li>chore(deps): update module golang.org/x/crypto to v0.49.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3468">IBM/sarama#3468</a></li>
<li>chore(deps): update dependency golangci/golangci-lint to v2.11.3 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3464">IBM/sarama#3464</a></li>
<li>fix(deps): update module golang.org/x/net to v0.52.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3472">IBM/sarama#3472</a></li>
<li>fix(deps): update module github.com/klauspost/compress to v1.18.5 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3480">IBM/sarama#3480</a></li>
<li>chore(deps): update module golang.org/x/crypto to v0.50.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3489">IBM/sarama#3489</a></li>
<li>fix(deps): update module golang.org/x/net to v0.53.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3493">IBM/sarama#3493</a></li>
<li>chore(deps): update docker/setup-buildx-action action to v4 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3458">IBM/sarama#3458</a></li>
<li>chore(deps): update docker/bake-action action to v7.1.0 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3459">IBM/sarama#3459</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore: add kafka versions 3.9.2 and 4.2.0 by <a href="https://github.com/edoardocomar"><code>@​edoardocomar</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3471">IBM/sarama#3471</a></li>
</ul>
<h3>:memo: Documentation</h3>
<ul>
<li>Update the Kakfa Protocol Specification Link by <a href="https://github.com/MohishKhadse55"><code>@​MohishKhadse55</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3463">IBM/sarama#3463</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>chore(deps): update dependency golangci/golangci-lint to v2.11.4 by <a href="https://github.com/renovate"><code>@​renovate</code></a>[bot] in <a href="https://redirect.github.com/IBM/sarama/pull/3482">IBM/sarama#3482</a></li>
<li>fix: update API version URL as previous link was not working by <a href="https://github.com/MohishKhadse55"><code>@​MohishKhadse55</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/3485">IBM/sarama#3485</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/MohishKhadse55"><code>@​MohishKhadse55</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3463">IBM/sarama#3463</a></li>
<li><a href="https://github.com/Asphaltt"><code>@​Asphaltt</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3384">IBM/sarama#3384</a></li>
<li><a href="https://github.com/ShinThirty"><code>@​ShinThirty</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/3487">IBM/sarama#3487</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.47.0...v1.48.0">https://github.com/IBM/sarama/compare/v1.47.0...v1.48.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/1c048928d4d904e85a270e0ffa68ca905d4de4b2"><code>1c04892</code></a> fix: add Unwrap() to DescribeConfigError and AlterConfigError (<a href="https://redirect.github.com/IBM/sarama/issues/3487">#3487</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/e025a7344a07c9490667c5c4dc1b87911a050077"><code>e025a73</code></a> fix: remove return from checkSeedBrokersHealth</li>
<li><a href="https://github.com/IBM/sarama/commit/0e6dc1b59908373bfb9f571ded3b1d85e9adf438"><code>0e6dc1b</code></a> fix: skip broker health tests without socket probing</li>
<li><a href="https://github.com/IBM/sarama/commit/ea568bd58746b4ba891042b6c557b9b9d6715182"><code>ea568bd</code></a> fix: keep bootstrap brokers and unblock async shutdown</li>
<li><a href="https://github.com/IBM/sarama/commit/8eafc46589bdf53b897391c827e77e6ac6958790"><code>8eafc46</code></a> fix: close broken tcp connections</li>
<li><a href="https://github.com/IBM/sarama/commit/188e5a2fd51e02ea21b4de4ab814b5d6ecc2ddbb"><code>188e5a2</code></a> chore(ci): Update registry.access.redhat.com/ubi9/ubi-minimal:9.7 Docker dige...</li>
<li><a href="https://github.com/IBM/sarama/commit/a0475058074e4359eb1d01066da986835bd384b4"><code>a047505</code></a> chore(ci): bump the actions group with 3 updates (<a href="https://redirect.github.com/IBM/sarama/issues/3499">#3499</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/374261bf7bb77d7af9f446d40657aa0819ce5c7e"><code>374261b</code></a> chore(ci): bump github/codeql-action from 4.34.1 to 4.35.2 (<a href="https://redirect.github.com/IBM/sarama/issues/3500">#3500</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/f1cd6c3c1b81de592bf4d3963d25113f4356b091"><code>f1cd6c3</code></a> chore(deps): update docker/bake-action action to v7.1.0 (<a href="https://redirect.github.com/IBM/sarama/issues/3459">#3459</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3d590ad72f31064b687f6c461fc4d680b356a5c1"><code>3d590ad</code></a> chore(deps): update docker/setup-buildx-action action to v4 (<a href="https://redirect.github.com/IBM/sarama/issues/3458">#3458</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.47.0...v1.48.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.69.1 to 1.70.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/776903f3dd7208803912e19d3aa25006a7fbdeee"><code>776903f</code></a> Release 2024-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/170b13cac4658e0909b13468d3959f94c358faf3"><code>170b13c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c1a02e760211a1d0762f664973ea4d896376a621"><code>c1a02e7</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2835f7bb8d4925acb36a5ab7813cca6ac8977cb7"><code>2835f7b</code></a> Fix user agent to add business metrics at the end instead of prepend them (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2">#2</a>...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ba4965dd291e68f9345c5eae4dec8a0f63bda436"><code>ba4965d</code></a> Release 2024-11-27</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55149b05036afa39e3b7c6dd945543faee4069e6"><code>55149b0</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fd6bb8b363cce171fba74b72d2bb4142b9306b7c"><code>fd6bb8b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/65ab4f88b048175d8fc13fab6f72f31b739a2455"><code>65ab4f8</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8939ed049c8b3490f9a3dbd4e4d3d56cf22a27f2"><code>8939ed0</code></a> Release 2024-11-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23cf36b2492a3caf859a6235c6911f1c2fb715ca"><code>23cf36b</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.69.1...service/s3/v1.70.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.99.1 to 1.100.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/18fee87f2c0615b0e5c3f28f1b95af810a9e77b5"><code>18fee87</code></a> Release 2026-04-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0ce157cd6b5427066ddb682a4967a3047230de4b"><code>0ce157c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/68cdb4a02d56b9d932fbfed22660cfa64e815ef6"><code>68cdb4a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f399972fb759f6ffd315963de511da785ef408e6"><code>f399972</code></a> Bump Smithy from 1.67.0 to 1.69.0 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3394">#3394</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3c19a9b3eb17f48f6b9fc1cad644859d062775ad"><code>3c19a9b</code></a> Release 2026-04-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1dfac3ec5a6c560e6fee389d0f1f364eb57ae616"><code>1dfac3e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d842df6fcef5ecc94e2133cb1d96fac4926aadb4"><code>d842df6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ffeb52a25097274529920d8f91b4ca27f7b636a9"><code>ffeb52a</code></a> bump to latest smithy-go codegen (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3393">#3393</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e839880e8d30897bd8a7684d83626d81f82c565a"><code>e839880</code></a> Release 2026-04-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4c898c88577f4102793d3e15de05ad3b5b1e2ef5"><code>4c898c8</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.99.1...service/s3/v1.100.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.7.5 to 4.8.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/c4ff4a69712419c5bc9b004bdec9cf3ca86f863b"><code>c4ff4a6</code></a> Update candlepin bindings to release/v4.8.0</li>
<li><a href="https://github.com/content-services/caliri/commit/8f9a5115e256b07d4bd5842b91fce3faae9c3145"><code>8f9a511</code></a> Update candlepin bindings to release/v4.7.5</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.7.5...release/v4.8.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2026` from 2026.4.1775836578 to 2026.4.1777255680
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/2f59c6a80e3e0a22e930de37d89d3c7bf7b7df7f"><code>2f59c6a</code></a> Update pulp bindings to e69db48356e528a464be3da896237b3d5abd83b1a7d54eb5892a9...</li>
<li><a href="https://github.com/content-services/zest/commit/f4103510bd095f6917a8250ff2678b873784fdcb"><code>f410351</code></a> Update pulp bindings to e69db48356e528a464be3da896237b386dadea51a7d54eb5892a9...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2026.4.1775836578...release/v2026.4.1777255680">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.45.1 to 0.46.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.46.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Remove SetExtra by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1274">#1274</a></li>
<li>Update compatibility policy to align with Go, supporting only the last two major Go versions by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1264">#1264</a></li>
<li>Drop support for Go 1.24 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1264">#1264</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add internal_sdk_error client report on serialization fail by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1273">#1273</a></li>
<li>Add grpc integration support by <a href="https://github.com/ribice"><code>@​ribice</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/938">#938</a></li>
<li>Re-enable Telemetry Processor by default. To disable the behavior use the <code>DisableTelemetryBuffer</code> flag by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1254">#1254</a></li>
<li>Simplify client DSN storage to <code>internal/protocol.Dsn</code> and make it safe to access by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1254">#1254</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<ul>
<li>Bump github.com/labstack/echo/v5 from 5.0.0 to 5.0.3 in /echo by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1253">#1253</a></li>
<li>Bump github.com/labstack/echo/v5 from 5.0.0 to 5.0.3 in /crosstest by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1272">#1272</a></li>
<li>Bump golangci-lint action from 2.1.1 to 2.11.4 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1265">#1265</a></li>
<li>Bump go.opentelemetry.io/otel/sdk from 1.40.0 to 1.43.0 in /otel by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1256">#1256</a></li>
<li>Bump go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp from 1.40.0 to 1.43.0 in /otel/otlp by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1255">#1255</a></li>
</ul>
<h4>Other</h4>
<ul>
<li>Improve ci by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1271">#1271</a></li>
<li>Add crosstest package by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1269">#1269</a></li>
<li>Add sentrytest package by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1267">#1267</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.46.0</h2>
<h3>Breaking Changes 🛠</h3>
<ul>
<li>Remove SetExtra by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1274">#1274</a></li>
<li>Update compatibility policy to align with Go, supporting only the last two major Go versions by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1264">#1264</a></li>
<li>Drop support for Go 1.24 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1264">#1264</a></li>
</ul>
<h3>New Features ✨</h3>
<ul>
<li>Add internal_sdk_error client report on serialization fail by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1273">#1273</a></li>
<li>Add grpc integration support by <a href="https://github.com/ribice"><code>@​ribice</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/938">#938</a></li>
<li>Re-enable Telemetry Processor by default. To disable the behavior use the <code>DisableTelemetryBuffer</code> flag by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1254">#1254</a></li>
<li>Simplify client DSN storage to <code>internal/protocol.Dsn</code> and make it safe to access by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1254">#1254</a></li>
</ul>
<h3>Internal Changes 🔧</h3>
<h4>Deps</h4>
<ul>
<li>Bump github.com/labstack/echo/v5 from 5.0.0 to 5.0.3 in /echo by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1253">#1253</a></li>
<li>Bump github.com/labstack/echo/v5 from 5.0.0 to 5.0.3 in /crosstest by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1272">#1272</a></li>
<li>Bump golangci-lint action from 2.1.1 to 2.11.4 by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1265">#1265</a></li>
<li>Bump go.opentelemetry.io/otel/sdk from 1.40.0 to 1.43.0 in /otel by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1256">#1256</a></li>
<li>Bump go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp from 1.40.0 to 1.43.0 in /otel/otlp by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1255">#1255</a></li>
</ul>
<h4>Other</h4>
<ul>
<li>Improve ci by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1271">#1271</a></li>
<li>Add crosstest package by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1269">#1269</a></li>
<li>Add sentrytest package by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1267">#1267</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/7e1f9265d7119244f29cecbeb913dc3eefe4014b"><code>7e1f926</code></a> release: 0.46.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/55688a9ed1a91c150f6584414f1ade15929bc30f"><code>55688a9</code></a> fix: keep replace directives (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1275">#1275</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/edb9172ff7f59e5815e24defae8916aeeb9b5391"><code>edb9172</code></a> feat!: remove SetExtra (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1274">#1274</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/dbb964c7ac89f6b879cb8551cd9aa6901a1f70ae"><code>dbb964c</code></a> feat: add internal_sdk_error client report on serialization fail (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1273">#1273</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/1fee8951f583ce2d2a5cff30dbeb50240f97c8de"><code>1fee895</code></a> feat: Add grpc integration support (<a href="https://redirect.github.com/getsentry/sentry-go/issues/938">#938</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/f4eabe880fcbfdfdf81df93f0d88d5f71c14af1e"><code>f4eabe8</code></a> feat: re-enable Telemetry Processor and simplify setup (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1254">#1254</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/aaef0ffcbf4d26e2d59d3f68b255161f369e97eb"><code>aaef0ff</code></a> chore: improve ci (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1271">#1271</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/533d76265cd3a103bebfd7b1e6a676784ade1215"><code>533d762</code></a> chore: simplify makefile with go work (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1270">#1270</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/1ecdc2861ffc2d1c87c2877b096b0f5a12b3bbba"><code>1ecdc28</code></a> build(deps): bump github.com/labstack/echo/v5 from 5.0.0 to 5.0.3 in /echo (#...</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/00dcbf0d836f3d13d71ec592d7ef7399c22463ac"><code>00dcbf0</code></a> build(deps): bump github.com/labstack/echo/v5 from 5.0.0 to 5.0.3 in /crosste...</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-go/compare/v0.45.1...v0.46.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/project-kessel/kessel-sdk-go` from 1.7.0 to 1.8.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/project-kessel/kessel-sdk-go/releases">github.com/project-kessel/kessel-sdk-go's releases</a>.</em></p>
<blockquote>
<h2>v1.8.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Add cursor skill for release process by <a href="https://github.com/lennysgarage"><code>@​lennysgarage</code></a> in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/41">project-kessel/kessel-sdk-go#41</a></li>
<li>Update generated protobuf files by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/43">project-kessel/kessel-sdk-go#43</a></li>
<li>Update generated protobuf files by <a href="https://github.com/github-actions"><code>@​github-actions</code></a>[bot] in <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/pull/57">project-kessel/kessel-sdk-go#57</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.7.0...v1.8.0">https://github.com/project-kessel/kessel-sdk-go/compare/v1.7.0...v1.8.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/f8bb3ddcb8c062daa4f4e713ed0f049415e26829"><code>f8bb3dd</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/57">#57</a> from project-kessel/buf-generate-update</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/5f6d7a35184cec629fb0332bd3f831a46ca0c393"><code>5f6d7a3</code></a> Regenerate protobuf files from buf</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/cc804934e27664d328b4119a23938e00d7cfe618"><code>cc80493</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/43">#43</a> from project-kessel/buf-generate-update</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/9d36f186d6a56581f28d139a30a827dd64e4fdb7"><code>9d36f18</code></a> Regenerate protobuf files from buf</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/a1d8ac9b681bd70e80ceae9952fed1a2e749eeb6"><code>a1d8ac9</code></a> Merge pull request <a href="https://redirect.github.com/project-kessel/kessel-sdk-go/issues/41">#41</a> from project-kessel/add-cursor-skill</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/4ae00acf44480759b1782cbae7aef14945f55f06"><code>4ae00ac</code></a> Handle release changes in different PR</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/9230325556e5c9877db240248d4c133900be8bf8"><code>9230325</code></a> Add preflight check</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/4d8033f6d22a2505dc552d7a14a64bd429aa31e2"><code>4d8033f</code></a> Add confirmation before committing changes</li>
<li><a href="https://github.com/project-kessel/kessel-sdk-go/commit/2712fd8b8a5c8600c14976bd18c35ace1cc1edef"><code>2712fd8</code></a> Add cursor skill for release process</li>
<li>See full diff in <a href="https://github.com/project-kessel/kessel-sdk-go/compare/v1.7.0...v1.8.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.27.0 to 1.28.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>Release v1.28.0</h2>
<ul>
<li>Removes archived zerolog, replaces it with built-in go-sentry/zerolog</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/0ca1dfcbb93394c6fac2a82a47abf61c465ed3e3"><code>0ca1dfc</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/842">#842</a> from lenasolarova/fix/replace-archived-zerolog-sentry...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/7ceecca88c25e1ba6a4dcfe94285aa898d0a6d26"><code>7ceecca</code></a> fix: replace archived zerolog-sentry with sentry-go/zerolog</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/2ce51919b25758c652f905122f4760afecded0d0"><code>2ce5191</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/841">#841</a> from RedHatInsights/repo-sync/processing-tools/default</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/e6af653857f082befa18f9a171bf3ad068c14756"><code>e6af653</code></a> 🔄 synced local 'renovate.json' with remote 'renovate.json'</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/c20b9a6b24428b817f14083801146b658635c669"><code>c20b9a6</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/840">#840</a> from joselsegura/remove-old-bots-automerge</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/69d907fe9d1454a4ea0c3633cfdb6b9d607f3e39"><code>69d907f</code></a> Remove old bots auto-merge workflow</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/7fdaf4bdaa1d1e515fb3c80ce66481904f2d616d"><code>7fdaf4b</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/838">#838</a> from RedHatInsights/edit_codecov</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/32b98706575dba020c60f1f01f8d4a4f6e25872e"><code>32b9870</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/839">#839</a> from RedHatInsights/repo-sync/processing-tools/default</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/7675c681185746290eb382133861ee2dbfa7647d"><code>7675c68</code></a> 🔄 created local '.github/workflows/bots-auto-merge.yaml' from remote 'workflo...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/cd6711e4a2d5f08d613a881d1180efd9f5d82407"><code>cd6711e</code></a> 🔄 synced local '.github/workflows/linters.yaml' with remote 'workflows_exampl...</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.27.0...v1.28.0">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


---

## Discussion

### Comment by @swadeley on April 27, 2026 at 09:45 AM UTC

Hi

Will try recreate after this is merged:
https://github.com/content-services/content-sources-backend/pull/1480

### Comment by @swadeley on April 27, 2026 at 11:13 AM UTC

@dependabot recreate

### Comment by @dependabot on April 27, 2026 at 03:42 PM UTC

Looks like these dependencies are no longer being updated by Dependabot, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1478*
