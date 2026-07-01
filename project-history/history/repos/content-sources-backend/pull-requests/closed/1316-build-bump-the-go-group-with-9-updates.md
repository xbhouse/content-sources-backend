---
type: pull_request
number: 1316
title: "Build: Bump the go group with 9 updates"
state: closed
author: dependabot
created: 2025-12-01T04:24:15Z
updated: 2025-12-01T20:05:04Z
closed: 2025-12-01T20:05:02Z
base_branch: main
head_branch: dependabot/go_modules/go-1e8d9ebae9
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1316
---

# Pull Request #1316: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: December 01, 2025 at 04:24 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-1e8d9ebae9`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.17` | `1.0.18` |
| [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) | `4.19.0` | `4.19.1` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.32.1` | `1.32.2` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.19.1` | `1.19.2` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.60.1` | `1.61.1` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.92.0` | `1.92.1` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.11.1763728709` | `2025.11.1764189739` |
| [github.com/getsentry/sentry-go](https://github.com/getsentry/sentry-go) | `0.38.0` | `0.40.0` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.16.0` | `9.17.1` |

Updates `github.com/content-services/yummy` from 1.0.17 to 1.0.18
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/yummy/releases">github.com/content-services/yummy's releases</a>.</em></p>
<blockquote>
<h2>v1.0.18</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump golang.org/x/crypto from 0.42.0 to 0.45.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/content-services/yummy/pull/40">content-services/yummy#40</a></li>
<li>HMS-8948: fix parsing stream versions by <a href="https://github.com/rverdile"><code>@​rverdile</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/39">content-services/yummy#39</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.17...v1.0.18">https://github.com/content-services/yummy/compare/v1.0.17...v1.0.18</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/0bf48832464df5e7dbe8fe1dfcc17ec5054bf9fe"><code>0bf4883</code></a> HMS-8948: fix parsing stream versions (<a href="https://redirect.github.com/content-services/yummy/issues/39">#39</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/47f22f9a71389bc07b53acb820c22dbc5ae5b7dd"><code>47f22f9</code></a> Bump golang.org/x/crypto from 0.42.0 to 0.45.0 (<a href="https://redirect.github.com/content-services/yummy/issues/40">#40</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.17...v1.0.18">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/golang-migrate/migrate/v4` from 4.19.0 to 4.19.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.19.1</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/dvsekhvalnov/jose2go from 1.6.0 to 1.7.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1334">golang-migrate/migrate#1334</a></li>
<li>Update docker image to use go 1.25, addressing vulnerabilities  by <a href="https://github.com/matsoob"><code>@​matsoob</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1333">golang-migrate/migrate#1333</a></li>
<li>Bump golang.org/x/crypto from 0.36.0 to 0.45.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1335">golang-migrate/migrate#1335</a></li>
<li>fix linter issues by <a href="https://github.com/HaraldNordgren"><code>@​HaraldNordgren</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1325">golang-migrate/migrate#1325</a></li>
<li>Rename databaseName to databaseDriverName by <a href="https://github.com/iamonah"><code>@​iamonah</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1304">golang-migrate/migrate#1304</a></li>
<li>chore: Update cloud.google.com/go/spanner version by <a href="https://github.com/jferrl"><code>@​jferrl</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1330">golang-migrate/migrate#1330</a></li>
<li>Bump golang.org/x/crypto from 0.43.0 to 0.45.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1336">golang-migrate/migrate#1336</a></li>
<li>chore: remove dependency on &quot;hashicorp/go-multierror&quot; by <a href="https://github.com/leonklingele"><code>@​leonklingele</code></a> in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1322">golang-migrate/migrate#1322</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/matsoob"><code>@​matsoob</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1333">golang-migrate/migrate#1333</a></li>
<li><a href="https://github.com/iamonah"><code>@​iamonah</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1304">golang-migrate/migrate#1304</a></li>
<li><a href="https://github.com/jferrl"><code>@​jferrl</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1330">golang-migrate/migrate#1330</a></li>
<li><a href="https://github.com/leonklingele"><code>@​leonklingele</code></a> made their first contribution in <a href="https://redirect.github.com/golang-migrate/migrate/pull/1322">golang-migrate/migrate#1322</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/golang-migrate/migrate/compare/v4.19.0...v4.19.1">https://github.com/golang-migrate/migrate/compare/v4.19.0...v4.19.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/89e308c340a2c069c3260f268307d38a7db88227"><code>89e308c</code></a> chore: remove dependency on &quot;hashicorp/go-multierror&quot; (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1322">#1322</a>)</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/472ef2e75cf6acfb1d5c4092c9e8e05715631f24"><code>472ef2e</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1336">#1336</a> from golang-migrate/dependabot/go_modules/golang.org...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/8d7625987a5bcc4f74000a4357c4535b60de2ef6"><code>8d76259</code></a> Bump golang.org/x/crypto from 0.43.0 to 0.45.0</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/9f9df7c897f5fcc4432d36e2c9049339d4010626"><code>9f9df7c</code></a> chore: Update cloud.google.com/go/spanner version (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1330">#1330</a>)</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/a371c8ea2b677e6b5fb0bb88a2f0a501748132a3"><code>a371c8e</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1304">#1304</a> from iamonah/fix/clarify-databaseName-meaning</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/43cc3b3ea50ac4385eab5f6a95b85034d9dbf7a9"><code>43cc3b3</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1325">#1325</a> from HaraldNordgren/linter_issues</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/f939a8992502b33eeb4d7cb1e37972e916ec3bf7"><code>f939a89</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1335">#1335</a> from golang-migrate/dependabot/go_modules/golang.org...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/6dd86e0ded6f854fb88a54907030077288fda5af"><code>6dd86e0</code></a> Bump golang.org/x/crypto from 0.36.0 to 0.45.0</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/70e6d6d57408f73e126285a976eaa9b70834f5c2"><code>70e6d6d</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1333">#1333</a> from matsoob/updateGoInBaseImage</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/a51d0da644e28288679c31aadf7ec00e90c9483f"><code>a51d0da</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1334">#1334</a> from golang-migrate/dependabot/go_modules/github.com...</li>
<li>Additional commits viewable in <a href="https://github.com/golang-migrate/migrate/compare/v4.19.0...v4.19.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.32.1 to 1.32.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0cbb5aa17f9078cb45dc0e82d3e1d0abee3744a9"><code>0cbb5aa</code></a> Release 2024-10-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54c1dd6c74185b0c7df78159ec4d5b2c27e9e280"><code>54c1dd6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2cde144eedda9f509141751c3011ca64a6b6528e"><code>2cde144</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67fbd35762ef8694839df209714d2ec2c33d3df9"><code>67fbd35</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa04330cb6978ccb6a7bb3e198b3fb21abbd6333"><code>aa04330</code></a> Allow non-nil but empty headers (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2826">#2826</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5a4e5bb42c08ff5a4e0e601a7461c8466565e44e"><code>5a4e5bb</code></a> add feature tracking for cbor protocol (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2821">#2821</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/183987cda0c2487a1b25c8e9cbf8dba510046c73"><code>183987c</code></a> add annotations to deprecated services and introduce codegen integration for ...</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.1...v1.32.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.19.1 to 1.19.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/31c2f3f89b98bd55ccb8833812d087baa9764e45"><code>31c2f3f</code></a> Release 2025-01-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ed70e6b14012a65186a1175d16bc60a665803d6e"><code>ed70e6b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5aef5b0eac63548e528afefe0431f9acd3b3a22d"><code>5aef5b0</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6e21e3f6b7c4adaa3db93457d03939c34b369ad8"><code>6e21e3f</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90178248e923ae9c90df9b592e6392878c07a7f4"><code>9017824</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ebb7c024c114fe872f65f8a1d58ba49f0cc1a376"><code>ebb7c02</code></a> retry net.ErrClosed by default (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2949">#2949</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/19d2a2833744bbc81aeca1896461c2d81e718d4e"><code>19d2a28</code></a> Release 2025-01-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e153a5993473dd666481855453b06dde48d66a7a"><code>e153a59</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/349cb942a7ae0cc5c42b764f39ac441115ccf7cf"><code>349cb94</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/740de30f11f73c77ef70aa722c0cf74fc876133d"><code>740de30</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.19.1...service/m2/v1.19.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.60.1 to 1.61.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2b751d1ba71f59175a41f9cae5f159f1044360f"><code>a2b751d</code></a> Release 2024-09-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e22c2495bd38232c640776ef3c1a84fb3145a8b9"><code>e22c249</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff0cf6fbfa7c7a12388606d5568135f2beae6599"><code>ff0cf6f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3120376762853b3098fda7e9217fb39fe1bf5105"><code>3120376</code></a> refactoring of buildQuery to accept a list of maintained headers to l… (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2773">#2773</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ed838eab2a963cb16301501c8b8c3e29dac4c20"><code>4ed838e</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2768">#2768</a> from bhavya2109sharma/presignedurl-requestpayer-change</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d4bd42fdc82c2954f429bd9eeac3096f5590eaac"><code>d4bd42f</code></a> Merge branch 'main' into presignedurl-requestpayer-change</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0353706229a89749afa93324432ece53da37822b"><code>0353706</code></a> Added Changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/97e2d3fec04bc8bfd69eaf22ce476b12f8673810"><code>97e2d3f</code></a> Release 2024-08-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4cca52b4a81df57e91b1e5d0a65fd6df89606d02"><code>4cca52b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a5146734eb1e6d59acba18b5967c78dc7a5e42"><code>c8a5146</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.60.1...service/s3/v1.61.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.92.0 to 1.92.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/34ec0a43cbf69475be3bbbb6329c49687cc018df"><code>34ec0a4</code></a> Release 2025-11-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/540441aa03ecb2b3b2b2f1eab1b71661b0ae2481"><code>540441a</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9869a2b39a8edb77fc98e0fc70d9bb059c6534a5"><code>9869a2b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a6b38c167b9ce07461c4f0af81be087055267121"><code>a6b38c1</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3953a0d499bf4a96a11e3377af3e5291831eed6c"><code>3953a0d</code></a> add explicit message deser to all s3 errors (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3237">#3237</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e7eec69683d04dbab33314fb6de43f8a43684163"><code>e7eec69</code></a> Fix panic during auth scheme resolution due to region validation (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3235">#3235</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d68b3a05c22b3bde751da6bb70e6fe01fd02407f"><code>d68b3a0</code></a> Release 2025-11-24</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/292a19869df57271d51b382018591d71f09f72d3"><code>292a198</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/dfeabc64ff80e2ee65951eb84d616072c8cd4b60"><code>dfeabc6</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a7a1be5d2c14e9270e927654b63272fdfbee1aa6"><code>a7a1be5</code></a> Release 2025-11-21</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.92.0...service/s3/v1.92.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.11.1763728709 to 2025.11.1764189739
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/63a349804bcfbb9a2b371a11aa9f5e790bbb7dbf"><code>63a3498</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a972869692294c57bd286834ede...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.11.1763728709...release/v2025.11.1764189739">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/getsentry/sentry-go` from 0.38.0 to 0.40.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/releases">github.com/getsentry/sentry-go's releases</a>.</em></p>
<blockquote>
<h2>0.40.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.40.0.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Disable <code>DisableTelemetryBuffer</code> flag and noop Telemetry Buffer, to prevent a panic at runtime (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1149">#1149</a>).</li>
</ul>
<h2>0.39.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.39.0.</p>
<h3>Features</h3>
<ul>
<li>Drop events from the telemetry buffer when rate-limited or transport is full, allowing the buffer queue to empty itself under load (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1138">#1138</a>).</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Fix scheduler's <code>hasWork()</code> method to check if buffers are ready to flush. The previous implementation was causing CPU spikes (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1143">#1143</a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-go/blob/master/CHANGELOG.md">github.com/getsentry/sentry-go's changelog</a>.</em></p>
<blockquote>
<h2>0.40.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.40.0.</p>
<h3>Bug Fixes</h3>
<ul>
<li>Disable <code>DisableTelemetryBuffer</code> flag and noop Telemetry Buffer, to prevent a panic at runtime (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1149">#1149</a>).</li>
</ul>
<h2>0.39.0</h2>
<p>The Sentry SDK team is happy to announce the immediate availability of Sentry Go SDK v0.39.0.</p>
<h3>Features</h3>
<ul>
<li>Drop events from the telemetry buffer when rate-limited or transport is full, allowing the buffer queue to empty itself under load (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1138">#1138</a>).</li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li>Fix scheduler's <code>hasWork()</code> method to check if buffers are ready to flush. The previous implementation was causing CPU spikes (<a href="https://redirect.github.com/getsentry/sentry-go/pull/1143">#1143</a>).</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-go/commit/f3fc5645c662d42363d3c47adee155a03a757f4b"><code>f3fc564</code></a> release: 0.40.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/9232f5d2f9727119fef5b9db5aad9e2906758f58"><code>9232f5d</code></a> chore: prepare 0.40.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1150">#1150</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/55ee13e338b87f7fb0c691f38158809ff9fb30f7"><code>55ee13e</code></a> feat: disable Telemetry Buffer (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1149">#1149</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/34261f3c64f1876aa6c891848108ffd88a299787"><code>34261f3</code></a> Merge branch 'release/0.39.0'</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a36268230302c750b38d4fe3a6f097a216fc355d"><code>a362682</code></a> release: 0.39.0</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/2153b5042be712bdbe1cba4f15c3e2aa7ac4ffe4"><code>2153b50</code></a> chore: prepare 0.39.0 (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1144">#1144</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/112f257bfa4d5b452a3f5467e72c4ee1adfbff47"><code>112f257</code></a> fix: scheduler should check for ready to flush buffers (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1143">#1143</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/03de096007be889e2fdec6bf5fc658419faa2fb7"><code>03de096</code></a> feat: drop rate-limited events (<a href="https://redirect.github.com/getsentry/sentry-go/issues/1138">#1138</a>)</li>
<li><a href="https://github.com/getsentry/sentry-go/commit/a2066ef178eeb4cb2808f565416283465b51b233"><code>a2066ef</code></a> Merge branch 'release/0.38.0'</li>
<li>See full diff in <a href="https://github.com/getsentry/sentry-go/compare/v0.38.0...v0.40.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.16.0 to 9.17.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>9.17.1</h2>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>add wait to keyless commands list (<a href="https://redirect.github.com/redis/go-redis/pull/3615">#3615</a>) by <a href="https://github.com/marcoferrer"><code>@​marcoferrer</code></a></li>
<li>fix(time): remove cached time optimization (<a href="https://redirect.github.com/redis/go-redis/pull/3611">#3611</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>chore(deps): bump golangci/golangci-lint-action from 9.0.0 to 9.1.0 (<a href="https://redirect.github.com/redis/go-redis/pull/3609">#3609</a>)</li>
<li>chore(deps): bump actions/checkout from 5 to 6 (<a href="https://redirect.github.com/redis/go-redis/pull/3610">#3610</a>)</li>
<li>chore(script): fix help call in tag.sh (<a href="https://redirect.github.com/redis/go-redis/pull/3606">#3606</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/marcoferrer"><code>@​marcoferrer</code></a> and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<h2>9.17.0</h2>
<h2>🚀 Highlights</h2>
<h3>Redis 8.4 Support</h3>
<p>Added support for Redis 8.4, including new commands and features (<a href="https://redirect.github.com/redis/go-redis/pull/3572">#3572</a>)</p>
<h3>Typed Errors</h3>
<p>Introduced typed errors for better error handling using <code>errors.As</code> instead of string checks. Errors can now be wrapped and set to commands in hooks without breaking library functionality (<a href="https://redirect.github.com/redis/go-redis/pull/3602">#3602</a>)</p>
<h3>New Commands</h3>
<ul>
<li><strong>CAS/CAD Commands</strong>: Added support for Compare-And-Set/Compare-And-Delete operations with conditional matching (<code>IFEQ</code>, <code>IFNE</code>, <code>IFDEQ</code>, <code>IFDNE</code>) (<a href="https://redirect.github.com/redis/go-redis/pull/3583">#3583</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3595">#3595</a>)</li>
<li><strong>MSETEX</strong>: Atomically set multiple key-value pairs with expiration options and conditional modes (<a href="https://redirect.github.com/redis/go-redis/pull/3580">#3580</a>)</li>
<li><strong>XReadGroup CLAIM</strong>: Consume both incoming and idle pending entries from streams in a single call (<a href="https://redirect.github.com/redis/go-redis/pull/3578">#3578</a>)</li>
<li><strong>ACL Commands</strong>: Added <code>ACLGenPass</code>, <code>ACLUsers</code>, and <code>ACLWhoAmI</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3576">#3576</a>)</li>
<li><strong>SLOWLOG Commands</strong>: Added <code>SLOWLOG LEN</code> and <code>SLOWLOG RESET</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3585">#3585</a>)</li>
<li><strong>LATENCY Commands</strong>: Added <code>LATENCY LATEST</code> and <code>LATENCY RESET</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3584">#3584</a>)</li>
</ul>
<h3>Search &amp; Vector Improvements</h3>
<ul>
<li><strong>Hybrid Search</strong>: Added  <strong>EXPERIMENTAL</strong> support for the new <code>FT.HYBRID</code> command (<a href="https://redirect.github.com/redis/go-redis/pull/3573">#3573</a>)</li>
<li><strong>Vector Range</strong>: Added <code>VRANGE</code> command for vector sets (<a href="https://redirect.github.com/redis/go-redis/pull/3543">#3543</a>)</li>
<li><strong>FT.INFO Enhancements</strong>: Added vector-specific attributes in FT.INFO response (<a href="https://redirect.github.com/redis/go-redis/pull/3596">#3596</a>)</li>
</ul>
<h3>Connection Pool Improvements</h3>
<ul>
<li><strong>Improved Connection Success Rate</strong>: Implemented FIFO queue-based fairness and context pattern for connection creation to prevent premature cancellation under high concurrency (<a href="https://redirect.github.com/redis/go-redis/pull/3518">#3518</a>)</li>
<li><strong>Connection State Machine</strong>: Resolved race conditions and improved pool performance with proper state tracking (<a href="https://redirect.github.com/redis/go-redis/pull/3559">#3559</a>)</li>
<li><strong>Pool Performance</strong>: Significant performance improvements with faster semaphores, lockless hook manager, and reduced allocations (47-67% faster Get/Put operations) (<a href="https://redirect.github.com/redis/go-redis/pull/3565">#3565</a>)</li>
</ul>
<h3>Metrics &amp; Observability</h3>
<ul>
<li><strong>Canceled Metric Attribute</strong>: Added 'canceled' metrics attribute to distinguish context cancellation errors from other errors (<a href="https://redirect.github.com/redis/go-redis/pull/3566">#3566</a>)</li>
</ul>
<h2>✨ New Features</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/blob/master/RELEASE-NOTES.md">github.com/redis/go-redis/v9's changelog</a>.</em></p>
<blockquote>
<h1>9.17.1 (2025-11-25)</h1>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>add wait to keyless commands list (<a href="https://redirect.github.com/redis/go-redis/pull/3615">#3615</a>) by <a href="https://github.com/marcoferrer"><code>@​marcoferrer</code></a></li>
<li>fix(time): remove cached time optimization (<a href="https://redirect.github.com/redis/go-redis/pull/3611">#3611</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<h2>🧰 Maintenance</h2>
<ul>
<li>chore(deps): bump golangci/golangci-lint-action from 9.0.0 to 9.1.0 (<a href="https://redirect.github.com/redis/go-redis/pull/3609">#3609</a>)</li>
<li>chore(deps): bump actions/checkout from 5 to 6 (<a href="https://redirect.github.com/redis/go-redis/pull/3610">#3610</a>)</li>
<li>chore(script): fix help call in tag.sh (<a href="https://redirect.github.com/redis/go-redis/pull/3606">#3606</a>) by <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/marcoferrer"><code>@​marcoferrer</code></a> and <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<hr />
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.17.0...v9.17.1">https://github.com/redis/go-redis/compare/v9.17.0...v9.17.1</a></p>
<h1>9.17.0 (2025-11-19)</h1>
<h2>🚀 Highlights</h2>
<h3>Redis 8.4 Support</h3>
<p>Added support for Redis 8.4, including new commands and features (<a href="https://redirect.github.com/redis/go-redis/pull/3572">#3572</a>)</p>
<h3>Typed Errors</h3>
<p>Introduced typed errors for better error handling using <code>errors.As</code> instead of string checks. Errors can now be wrapped and set to commands in hooks without breaking library functionality (<a href="https://redirect.github.com/redis/go-redis/pull/3602">#3602</a>)</p>
<h3>New Commands</h3>
<ul>
<li><strong>CAS/CAD Commands</strong>: Added support for Compare-And-Set/Compare-And-Delete operations with conditional matching (<code>IFEQ</code>, <code>IFNE</code>, <code>IFDEQ</code>, <code>IFDNE</code>) (<a href="https://redirect.github.com/redis/go-redis/pull/3583">#3583</a>, <a href="https://redirect.github.com/redis/go-redis/pull/3595">#3595</a>)</li>
<li><strong>MSETEX</strong>: Atomically set multiple key-value pairs with expiration options and conditional modes (<a href="https://redirect.github.com/redis/go-redis/pull/3580">#3580</a>)</li>
<li><strong>XReadGroup CLAIM</strong>: Consume both incoming and idle pending entries from streams in a single call (<a href="https://redirect.github.com/redis/go-redis/pull/3578">#3578</a>)</li>
<li><strong>ACL Commands</strong>: Added <code>ACLGenPass</code>, <code>ACLUsers</code>, and <code>ACLWhoAmI</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3576">#3576</a>)</li>
<li><strong>SLOWLOG Commands</strong>: Added <code>SLOWLOG LEN</code> and <code>SLOWLOG RESET</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3585">#3585</a>)</li>
<li><strong>LATENCY Commands</strong>: Added <code>LATENCY LATEST</code> and <code>LATENCY RESET</code> (<a href="https://redirect.github.com/redis/go-redis/pull/3584">#3584</a>)</li>
</ul>
<h3>Search &amp; Vector Improvements</h3>
<ul>
<li><strong>Hybrid Search</strong>: Added  <strong>EXPERIMENTAL</strong> support for the new <code>FT.HYBRID</code> command (<a href="https://redirect.github.com/redis/go-redis/pull/3573">#3573</a>)</li>
<li><strong>Vector Range</strong>: Added <code>VRANGE</code> command for vector sets (<a href="https://redirect.github.com/redis/go-redis/pull/3543">#3543</a>)</li>
<li><strong>FT.INFO Enhancements</strong>: Added vector-specific attributes in FT.INFO response (<a href="https://redirect.github.com/redis/go-redis/pull/3596">#3596</a>)</li>
</ul>
<h3>Connection Pool Improvements</h3>
<ul>
<li><strong>Improved Connection Success Rate</strong>: Implemented FIFO queue-based fairness and context pattern for connection creation to prevent premature cancellation under high concurrency (<a href="https://redirect.github.com/redis/go-redis/pull/3518">#3518</a>)</li>
<li><strong>Connection State Machine</strong>: Resolved race conditions and improved pool performance with proper state tracking (<a href="https://redirect.github.com/redis/go-redis/pull/3559">#3559</a>)</li>
<li><strong>Pool Performance</strong>: Significant performance improvements with faster semaphores, lockless hook manager, and reduced allocations (47-67% faster Get/Put operations) (<a href="https://redirect.github.com/redis/go-redis/pull/3565">#3565</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/68d8c5955713fb22642a0d44163c6f785ba7425e"><code>68d8c59</code></a> chore(release): v9.17.1 (<a href="https://redirect.github.com/redis/go-redis/issues/3617">#3617</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/932caa5f13e11433a75d95e052e862010b016a1a"><code>932caa5</code></a> chore(deps): bump actions/stale from 9 to 10 (<a href="https://redirect.github.com/redis/go-redis/issues/3505">#3505</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/b362eb7f7d1f4b8aea6759ceea22dcaba9992b4b"><code>b362eb7</code></a> fix(txpipeline) add wait to keyless commands list (<a href="https://redirect.github.com/redis/go-redis/issues/3615">#3615</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/fd437cea4fb569a15d2a2dcb8b4969e63c28c157"><code>fd437ce</code></a> chore(deps): bump golangci/golangci-lint-action from 9.0.0 to 9.1.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3609">#3609</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/8549116d8a7553e128a3de9c262b48990d80f403"><code>8549116</code></a> chore(deps): bump actions/checkout from 5 to 6 (<a href="https://redirect.github.com/redis/go-redis/issues/3610">#3610</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/63908223ca0e4919396e33cbb27b8fa823b67793"><code>6390822</code></a> fix(time): remove cached time optimization (<a href="https://redirect.github.com/redis/go-redis/issues/3611">#3611</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/1bb9e0d130f3c6acb602d6d9f1ca4acebbe96677"><code>1bb9e0d</code></a> chore(scripts: fix help call in tag.sh (<a href="https://redirect.github.com/redis/go-redis/issues/3606">#3606</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/5b0b228a37c83b7c7ace2fcd7cf230f61e2e3092"><code>5b0b228</code></a> chore(release): v9.17.0 (<a href="https://redirect.github.com/redis/go-redis/issues/3604">#3604</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/6c24f600de102270c5b59e561ced9539fd13c817"><code>6c24f60</code></a> feat(errors): Introduce typed errors (<a href="https://redirect.github.com/redis/go-redis/issues/3602">#3602</a>)</li>
<li><a href="https://github.com/redis/go-redis/commit/817e62455b598fbb50049b3d6f907ced1e309344"><code>817e624</code></a> chore(ci): official 8.4.0 image (<a href="https://redirect.github.com/redis/go-redis/issues/3603">#3603</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/redis/go-redis/compare/v9.16.0...v9.17.1">compare view</a></li>
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

### Comment by @TenSt on December 01, 2025 at 09:58 AM UTC

@dependabot rebase

### Comment by @dependabot on December 01, 2025 at 08:05 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1316*
