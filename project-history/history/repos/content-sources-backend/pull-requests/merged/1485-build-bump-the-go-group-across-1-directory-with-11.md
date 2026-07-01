---
type: pull_request
number: 1485
title: "Build: Bump the go group across 1 directory with 11 updates"
state: merged
author: dependabot
created: 2026-04-29T15:05:35Z
updated: 2026-04-29T19:09:13Z
closed: 2026-04-29T19:09:10Z
merged: 2026-04-29T19:09:10Z
base_branch: main
head_branch: dependabot/go_modules/go-bde8d2cbda
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1485
---

# Pull Request #1485: Build: Bump the go group across 1 directory with 11 updates

**Author**: @dependabot
**Created**: April 29, 2026 at 03:05 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-bde8d2cbda`

## Description

Bumps the go group with 11 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | `1.35.0` | `1.35.1` |
| [github.com/IBM/sarama](https://github.com/IBM/sarama) | `1.47.0` | `1.48.0` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.69.1` | `1.71.0` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.99.1` | `1.100.0` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.7.5` | `4.8.0` |
| [github.com/content-services/zest/release/v2026](https://github.com/content-services/zest) | `2026.4.1775836578` | `2026.4.1777255680` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.46.0` | `0.46.1` |
| [github.com/getsentry/sentry-go/zerolog](https://github.com/getsentry/sentry-go) | `0.46.0` | `0.46.1` |
| [github.com/project-kessel/kessel-sdk-go](https://github.com/project-kessel/kessel-sdk-go) | `1.7.0` | `1.8.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.18.0` | `9.19.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.27.0` | `1.28.0` |


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

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.69.1 to 1.71.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9a211f9a8eb9923a47740503892928d702a2ac63"><code>9a211f9</code></a> Release 2024-12-03.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c357d6ce9cf35ab9fc343d1d8ea1147885a1276b"><code>c357d6c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e6160c9846d4f72e5d09a2c6160b1cd1770b5bad"><code>e6160c9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cfd99f86832af2f033fcd34148bef6f32ab36211"><code>cfd99f8</code></a> Merge customizations for DSQL</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3418b7488a6dda849fa2b314aad7d8a24af435e0"><code>3418b74</code></a> Release 2024-12-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bc3406170ac25a0cdb61aaa1b7c25415fdee17d3"><code>bc34061</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4eb6ed801453f7b5f45423d59f39b27409dd517b"><code>4eb6ed8</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b9b73ec7647e6ea6e5862d906f616e8c2d2b3e4b"><code>b9b73ec</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/776903f3dd7208803912e19d3aa25006a7fbdeee"><code>776903f</code></a> Release 2024-12-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/170b13cac4658e0909b13468d3959f94c358faf3"><code>170b13c</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ecs/v1.69.1...service/s3/v1.71.0">compare view</a></li>
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

Updates `github.com/getsentry/sentry-go` from 0.46.0 to 0.46.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.46.1</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Correctly capture request body for fasthttp and fiber by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1284">#1284</a></li>
<li>(http) Avoid async transport shutdown panics by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1288">#1288</a></li>
<li>(httpclient) Clone request before adding trace headers by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1290">#1290</a></li>
<li>(scope) Use scoped client for request PII by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1289">#1289</a></li>
<li>Safe concurrent access for span and scope by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1285">#1285</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.46.1</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Correctly capture request body for fasthttp and fiber. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1284">#1284</a></li>
<li>(http) Avoid async transport shutdown panics by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1288">#1288</a></li>
<li>(httpclient) Clone request before adding trace headers by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1290">#1290</a></li>
<li>(scope) Use scoped client for request PII by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1289">#1289</a></li>
<li>Safe concurrent access for span and scope by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1285">#1285</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/e972183b89e882147beae49a1ec8bf98ba1c3298"><code>e972183</code></a> release: 0.46.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/6b9885c0372193f8dfb7895f61d2354ef2e51502"><code>6b9885c</code></a> fix(http): avoid async transport shutdown panics (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1288">#1288</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/79947a7ad33239d1849ba619af2cb8922b074eb3"><code>79947a7</code></a> fix: safe concurrent access for span and scope (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1285">#1285</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c8ea578dfc589f9b3ca06b7a9c13019ac96325b5"><code>c8ea578</code></a> fix(scope): use scoped client for request PII (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1289">#1289</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/0bb583ea2b4292f2204468e09b465314048b03e1"><code>0bb583e</code></a> fix(httpclient): clone request before adding trace headers (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1290">#1290</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bd20df0d91c5d258394e0d52c732e18f0009d6d5"><code>bd20df0</code></a> fix(fasthttp,fiber): correctly capture request body on scope (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1284">#1284</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/41f09e115efec73c14961253fbffac6e55fef552"><code>41f09e1</code></a> fix(lint): Resolve workspace submodule lint issues (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1276">#1276</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7a480f73b74f6232e45a1f460ae2f88b2c07f086"><code>7a480f7</code></a> Merge branch 'release/0.46.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.46.0...v0.46.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go/zerolog` from 0.46.0 to 0.46.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go/zerolog's releases</a>.</em></p>
<blockquote>
<h2>0.46.1</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Correctly capture request body for fasthttp and fiber by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1284">#1284</a></li>
<li>(http) Avoid async transport shutdown panics by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1288">#1288</a></li>
<li>(httpclient) Clone request before adding trace headers by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1290">#1290</a></li>
<li>(scope) Use scoped client for request PII by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1289">#1289</a></li>
<li>Safe concurrent access for span and scope by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1285">#1285</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go/zerolog's changelog</a>.</em></p>
<blockquote>
<h2>0.46.1</h2>
<h3>Bug Fixes 🐛</h3>
<ul>
<li>Correctly capture request body for fasthttp and fiber. by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1284">#1284</a></li>
<li>(http) Avoid async transport shutdown panics by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1288">#1288</a></li>
<li>(httpclient) Clone request before adding trace headers by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1290">#1290</a></li>
<li>(scope) Use scoped client for request PII by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1289">#1289</a></li>
<li>Safe concurrent access for span and scope by <a href="https://github.com/giortzisg"><code>@​giortzisg</code></a> in <a href="https://redirect.github.com/getsentry/sentry-go/pull/1285">#1285</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/e972183b89e882147beae49a1ec8bf98ba1c3298"><code>e972183</code></a> release: 0.46.1</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/6b9885c0372193f8dfb7895f61d2354ef2e51502"><code>6b9885c</code></a> fix(http): avoid async transport shutdown panics (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1288">#1288</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/79947a7ad33239d1849ba619af2cb8922b074eb3"><code>79947a7</code></a> fix: safe concurrent access for span and scope (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1285">#1285</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/c8ea578dfc589f9b3ca06b7a9c13019ac96325b5"><code>c8ea578</code></a> fix(scope): use scoped client for request PII (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1289">#1289</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/0bb583ea2b4292f2204468e09b465314048b03e1"><code>0bb583e</code></a> fix(httpclient): clone request before adding trace headers (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1290">#1290</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/bd20df0d91c5d258394e0d52c732e18f0009d6d5"><code>bd20df0</code></a> fix(fasthttp,fiber): correctly capture request body on scope (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1284">#1284</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/41f09e115efec73c14961253fbffac6e55fef552"><code>41f09e1</code></a> fix(lint): Resolve workspace submodule lint issues (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1276">#1276</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/7a480f73b74f6232e45a1f460ae2f88b2c07f086"><code>7a480f7</code></a> Merge branch 'release/0.46.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.46.0...v0.46.1">compare view</a></li>
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

Updates `github.com/redis/go-redis/v9` from 9.18.0 to 9.19.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.19.0</h2>
<h2>🚀 Highlights</h2>
<h3>FIPS-Compatible Script Helper</h3>
<p><code>Script</code> now supports a FIPS-safe execution mode that avoids client-side SHA-1 computation, which is blocked in strict FIPS environments. A new <code>NewScriptServerSHA</code> constructor uses <code>SCRIPT LOAD</code> to obtain and cache the digest from the server, then runs commands via <code>EVALSHA</code>/<code>EVALSHA_RO</code>. Falls back to <code>EVAL</code>/<code>EVALRO</code> if loading fails, and transparently retries once on <code>NOSCRIPT</code>. The default behavior is unchanged for existing users.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3700">#3700</a>) by <a href="https://github.com/chaitanyabodlapati"><code>@​chaitanyabodlapati</code></a></p>
<h3>FT.AGGREGATE Step-Based Pipeline Builder</h3>
<p>Added a new step-based <code>FT.AGGREGATE</code> pipeline API via <code>FTAggregateOptions.Steps</code>, allowing <code>LOAD</code>, <code>APPLY</code>, <code>GROUPBY</code>, and <code>SORTBY</code> (with per-step <code>MAX</code>) to be repeated and interleaved in arbitrary order — matching Redis's native multi-stage aggregation semantics. The legacy <code>Load</code>/<code>Apply</code>/<code>GroupBy</code>/<code>SortBy</code>/<code>SortByMax</code> fields are now deprecated.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3782">#3782</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h3>Raw RESP Protocol Access</h3>
<p>Added <code>DoRaw</code> and <code>DoRawWriteTo</code> methods for executing arbitrary commands and reading the raw RESP response. Useful for proxying, custom protocol inspection, and working with commands not yet wrapped by go-redis.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3713">#3713</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
<h3>Configurable Dial Retry Backoff</h3>
<p>Added <code>DialerRetryBackoff</code> option (plumbed through <code>Options</code>, <code>ClusterOptions</code>, <code>RingOptions</code>, <code>FailoverOptions</code>) to let callers customize the delay between failed dial attempts. Helpers <code>DialRetryBackoffConstant</code> and <code>DialRetryBackoffExponential</code> (with jitter and cap) are provided out of the box. Dial timeout is now also applied <strong>per attempt</strong> rather than across all retries.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3706">#3706</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3705">#3705</a>) by <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a></p>
<h2>✨ New Features</h2>
<ul>
<li><strong>FT.AGGREGATE Steps</strong>: Step-based pipeline builder for <code>FT.AGGREGATE</code> with support for repeated/interleaved <code>LOAD</code>, <code>APPLY</code>, <code>GROUPBY</code>, and <code>SORTBY</code> stages (<a href="https://redirect.github.com/redis/go-redis/pull/3782">#3782</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>VectorSet commands</strong>: Added <code>VISMEMBER</code> and <code>WITHATTRIBS</code> support (<a href="https://redirect.github.com/redis/go-redis/pull/3753">#3753</a>) by <a href="https://github.com/romanpovol"><code>@​romanpovol</code></a></li>
<li><strong>FIPS-safe Script</strong>: <code>NewScriptServerSHA</code> uses <code>SCRIPT LOAD</code> to obtain the digest from the server, avoiding client-side SHA-1 (<a href="https://redirect.github.com/redis/go-redis/pull/3700">#3700</a>) by <a href="https://github.com/chaitanyabodlapati"><code>@​chaitanyabodlapati</code></a></li>
<li><strong>Raw RESP access</strong>: <code>DoRaw</code> and <code>DoRawWriteTo</code> for raw RESP protocol access (<a href="https://redirect.github.com/redis/go-redis/pull/3713">#3713</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>Dial retry backoff</strong>: <code>DialerRetryBackoff</code> function option with constant and exponential helpers (<a href="https://redirect.github.com/redis/go-redis/pull/3706">#3706</a>) by <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a></li>
<li><strong>Typed NOSCRIPT error</strong>: Redis <code>NOSCRIPT</code> replies are now surfaced as a typed error for easier handling (<a href="https://redirect.github.com/redis/go-redis/pull/3738">#3738</a>) by <a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a></li>
<li><strong>PubSub ClientSetName</strong>: Added <code>ClientSetName</code> method to <code>PubSub</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3727">#3727</a>) by <a href="https://github.com/Flack74"><code>@​Flack74</code></a></li>
<li><strong>ReplicaOf</strong>: New <code>ReplicaOf</code> method replaces the deprecated <code>SlaveOf</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3720">#3720</a>) by <a href="https://github.com/apps/copilot-swe-agent"><code>@​Copilot</code></a></li>
<li><strong>HSCAN BinaryUnmarshaler</strong>: <code>HScan</code> now supports types implementing <code>encoding.BinaryUnmarshaler</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3768">#3768</a>) by <a href="https://github.com/Aaditya-dubey1"><code>@​Aaditya-dubey1</code></a></li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Auto hostname type detection</strong>: Improved endpoint type detection for maintenance notifications using DNS-based classification; handles empty hosts and expanded private-IP ranges (<a href="https://redirect.github.com/redis/go-redis/pull/3789">#3789</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>HELLO fallback</strong>: Don't send <code>CLIENT MAINT_NOTIFICATIONS</code> handshake when <code>HELLO</code> fails and connection falls back to RESP2; fail fast when explicitly enabled with RESP3 (<a href="https://redirect.github.com/redis/go-redis/pull/3788">#3788</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Dial TCP retry</strong>: <code>ShouldRetry</code> now treats <code>net.OpError</code> with <code>Op == &quot;dial&quot;</code> timeout errors as safe to retry since no command was sent (<a href="https://redirect.github.com/redis/go-redis/pull/3787">#3787</a>) by <a href="https://github.com/vladisa88"><code>@​vladisa88</code></a></li>
<li><strong>wrappedOnClose leak</strong>: Fixed resource leak caused by repeatedly wrapping <code>baseClient</code> close logic; replaced with a bounded, concurrency-safe named-hook registry (<a href="https://redirect.github.com/redis/go-redis/pull/3785">#3785</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Pool Close() on stale connections</strong>: Suppress close errors (e.g., TLS <code>closeNotify</code> timeouts) for connections already dropped by the server due to idle timeout (<a href="https://redirect.github.com/redis/go-redis/pull/3778">#3778</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>FIFO waiter ordering</strong>: Fixed race in <code>ConnStateMachine.notifyWaiters</code> that could wake multiple waiters under a single mutex hold and violate FIFO ordering (<a href="https://redirect.github.com/redis/go-redis/pull/3777">#3777</a>) by <a href="https://github.com/0x48core"><code>@​0x48core</code></a></li>
<li><strong>Lua READONLY detection</strong>: Detect <code>READONLY</code> errors embedded in Lua script error messages on read-only replicas so commands are correctly retried (<a href="https://redirect.github.com/redis/go-redis/pull/3769">#3769</a>) by <a href="https://github.com/zhengjilei"><code>@​zhengjilei</code></a></li>
<li><strong>VectorScoreSliceCmd RESP2</strong>: Fixed <code>VSimWithScores</code>, <code>VSimWithArgsWithScores</code>, and <code>VLinksWithScores</code> which were broken on RESP2 connections returning flat arrays instead of maps (<a href="https://redirect.github.com/redis/go-redis/pull/3767">#3767</a>) by <a href="https://github.com/apps/copilot-swe-agent"><code>@​Copilot</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.19.0 (2026-04-27)</h1>
<h2>🚀 Highlights</h2>
<h3>FIPS-Compatible Script Helper</h3>
<p><code>Script</code> now supports a FIPS-safe execution mode that avoids client-side SHA-1 computation, which is blocked in strict FIPS environments. A new <code>NewScriptServerSHA</code> constructor uses <code>SCRIPT LOAD</code> to obtain and cache the digest from the server, then runs commands via <code>EVALSHA</code>/<code>EVALSHA_RO</code>. Falls back to <code>EVAL</code>/<code>EVALRO</code> if loading fails, and transparently retries once on <code>NOSCRIPT</code>. The default behavior is unchanged for existing users.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3700">#3700</a>) by <a href="https://github.com/chaitanyabodlapati"><code>@​chaitanyabodlapati</code></a></p>
<h3>FT.AGGREGATE Step-Based Pipeline Builder</h3>
<p>Added a new step-based <code>FT.AGGREGATE</code> pipeline API via <code>FTAggregateOptions.Steps</code>, allowing <code>LOAD</code>, <code>APPLY</code>, <code>GROUPBY</code>, and <code>SORTBY</code> (with per-step <code>MAX</code>) to be repeated and interleaved in arbitrary order — matching Redis's native multi-stage aggregation semantics. The legacy <code>Load</code>/<code>Apply</code>/<code>GroupBy</code>/<code>SortBy</code>/<code>SortByMax</code> fields are now deprecated.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3782">#3782</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h3>Raw RESP Protocol Access</h3>
<p>Added <code>DoRaw</code> and <code>DoRawWriteTo</code> methods for executing arbitrary commands and reading the raw RESP response. Useful for proxying, custom protocol inspection, and working with commands not yet wrapped by go-redis.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3713">#3713</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></p>
<h3>Configurable Dial Retry Backoff</h3>
<p>Added <code>DialerRetryBackoff</code> option (plumbed through <code>Options</code>, <code>ClusterOptions</code>, <code>RingOptions</code>, <code>FailoverOptions</code>) to let callers customize the delay between failed dial attempts. Helpers <code>DialRetryBackoffConstant</code> and <code>DialRetryBackoffExponential</code> (with jitter and cap) are provided out of the box. Dial timeout is now also applied <strong>per attempt</strong> rather than across all retries.</p>
<p>(<a href="https://redirect.github.com/redis/go-redis/pull/3706">#3706</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3705">#3705</a>) by <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a></p>
<h2>✨ New Features</h2>
<ul>
<li><strong>FT.AGGREGATE Steps</strong>: Step-based pipeline builder for <code>FT.AGGREGATE</code> with support for repeated/interleaved <code>LOAD</code>, <code>APPLY</code>, <code>GROUPBY</code>, and <code>SORTBY</code> stages (<a href="https://redirect.github.com/redis/go-redis/pull/3782">#3782</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>VectorSet commands</strong>: Added <code>VISMEMBER</code> and <code>WITHATTRIBS</code> support (<a href="https://redirect.github.com/redis/go-redis/pull/3753">#3753</a>) by <a href="https://github.com/romanpovol"><code>@​romanpovol</code></a></li>
<li><strong>FIPS-safe Script</strong>: <code>NewScriptServerSHA</code> uses <code>SCRIPT LOAD</code> to obtain the digest from the server, avoiding client-side SHA-1 (<a href="https://redirect.github.com/redis/go-redis/pull/3700">#3700</a>) by <a href="https://github.com/chaitanyabodlapati"><code>@​chaitanyabodlapati</code></a></li>
<li><strong>Raw RESP access</strong>: <code>DoRaw</code> and <code>DoRawWriteTo</code> for raw RESP protocol access (<a href="https://redirect.github.com/redis/go-redis/pull/3713">#3713</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>Dial retry backoff</strong>: <code>DialerRetryBackoff</code> function option with constant and exponential helpers (<a href="https://redirect.github.com/redis/go-redis/pull/3706">#3706</a>) by <a href="https://github.com/mwhooker"><code>@​mwhooker</code></a></li>
<li><strong>Typed NOSCRIPT error</strong>: Redis <code>NOSCRIPT</code> replies are now surfaced as a typed error for easier handling (<a href="https://redirect.github.com/redis/go-redis/pull/3738">#3738</a>) by <a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a></li>
<li><strong>PubSub ClientSetName</strong>: Added <code>ClientSetName</code> method to <code>PubSub</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3727">#3727</a>) by <a href="https://github.com/Flack74"><code>@​Flack74</code></a></li>
<li><strong>ReplicaOf</strong>: New <code>ReplicaOf</code> method replaces the deprecated <code>SlaveOf</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3720">#3720</a>) by <a href="https://github.com/apps/copilot-swe-agent"><code>@​Copilot</code></a></li>
<li><strong>HSCAN BinaryUnmarshaler</strong>: <code>HScan</code> now supports types implementing <code>encoding.BinaryUnmarshaler</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3768">#3768</a>) by <a href="https://github.com/Aaditya-dubey1"><code>@​Aaditya-dubey1</code></a></li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Auto hostname type detection</strong>: Improved endpoint type detection for maintenance notifications using DNS-based classification; handles empty hosts and expanded private-IP ranges (<a href="https://redirect.github.com/redis/go-redis/pull/3789">#3789</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>HELLO fallback</strong>: Don't send <code>CLIENT MAINT_NOTIFICATIONS</code> handshake when <code>HELLO</code> fails and connection falls back to RESP2; fail fast when explicitly enabled with RESP3 (<a href="https://redirect.github.com/redis/go-redis/pull/3788">#3788</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Dial TCP retry</strong>: <code>ShouldRetry</code> now treats <code>net.OpError</code> with <code>Op == &quot;dial&quot;</code> timeout errors as safe to retry since no command was sent (<a href="https://redirect.github.com/redis/go-redis/pull/3787">#3787</a>) by <a href="https://github.com/vladisa88"><code>@​vladisa88</code></a></li>
<li><strong>wrappedOnClose leak</strong>: Fixed resource leak caused by repeatedly wrapping <code>baseClient</code> close logic; replaced with a bounded, concurrency-safe named-hook registry (<a href="https://redirect.github.com/redis/go-redis/pull/3785">#3785</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
<li><strong>Pool Close() on stale connections</strong>: Suppress close errors (e.g., TLS <code>closeNotify</code> timeouts) for connections already dropped by the server due to idle timeout (<a href="https://redirect.github.com/redis/go-redis/pull/3778">#3778</a>) by <a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a></li>
<li><strong>FIFO waiter ordering</strong>: Fixed race in <code>ConnStateMachine.notifyWaiters</code> that could wake multiple waiters under a single mutex hold and violate FIFO ordering (<a href="https://redirect.github.com/redis/go-redis/pull/3777">#3777</a>) by <a href="https://github.com/0x48core"><code>@​0x48core</code></a></li>
<li><strong>Lua READONLY detection</strong>: Detect <code>READONLY</code> errors embedded in Lua script error messages on read-only replicas so commands are correctly retried (<a href="https://redirect.github.com/redis/go-redis/pull/3769">#3769</a>) by <a href="https://github.com/zhengjilei"><code>@​zhengjilei</code></a></li>
<li><strong>VectorScoreSliceCmd RESP2</strong>: Fixed <code>VSimWithScores</code>, <code>VSimWithArgsWithScores</code>, and <code>VLinksWithScores</code> which were broken on RESP2 connections returning flat arrays instead of maps (<a href="https://redirect.github.com/redis/go-redis/pull/3767">#3767</a>) by <a href="https://github.com/apps/copilot-swe-agent"><code>@​Copilot</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/e7e9866e54f4423addca48ec804044cfca1b30d6"><code>e7e9866</code></a> chore(release): v9.19.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3796">#3796</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/22b26f4b3ea5c2ab58605092e0937b78a8673267"><code>22b26f4</code></a> feat(ft.aggregate): Add Steps for query building (<a href="https://redirect.github.com/redis/go-redis/issues/3782">#3782</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/d9d769491154d4da6d683b73e9dfff5c878c1890"><code>d9d7694</code></a> fix(pool): two fixes for closed connection handling (<a href="https://redirect.github.com/redis/go-redis/issues/3764">#3764</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/44e8b7309cf547f620b9849a5a4dd9013cc3660a"><code>44e8b73</code></a> fix(sch): auto hostname type detection (<a href="https://redirect.github.com/redis/go-redis/issues/3789">#3789</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/ad21622ce8dba26c4ac7dc0fbf74a0bd278d2c19"><code>ad21622</code></a> fix(hello): do not send maintnotifications handshake when hello fails (<a href="https://redirect.github.com/redis/go-redis/issues/3788">#3788</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/1a7ac744c3f98b415852b4c9842b97991e3e4d91"><code>1a7ac74</code></a> fix(pool): suppress pool Close() errors for stale connections (<a href="https://redirect.github.com/redis/go-redis/issues/3778">#3778</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/903d6bd360af9ea53d4af4ca652d18b915398c7d"><code>903d6bd</code></a> fix(retry): make dial tcp error redirectable (<a href="https://redirect.github.com/redis/go-redis/issues/3786">#3786</a>) (<a href="https://redirect.github.com/redis/go-redis/issues/3787">#3787</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/00a551b3b468687e5e06b295b419292786bcdc99"><code>00a551b</code></a> fix(credentials): leak in wrappedOnClose (<a href="https://redirect.github.com/redis/go-redis/issues/3785">#3785</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b5a6f99b7a9d496bf2e3d68ad9c33e43c23c4df9"><code>b5a6f99</code></a> refactor(pool): remove redundant Conn.closed atomic field (<a href="https://redirect.github.com/redis/go-redis/issues/3783">#3783</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/928f27aef850b3508b590e41a050f27d29fb14a2"><code>928f27a</code></a> feat(hscan): add support for encoding.BinaryUnmarshaler (<a href="https://redirect.github.com/redis/go-redis/issues/3768">#3768</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.18.0...v9.19.0">compare view</a></li>
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

### Review by @swadeley - Approved on April 29, 2026 at 06:50 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1485*
