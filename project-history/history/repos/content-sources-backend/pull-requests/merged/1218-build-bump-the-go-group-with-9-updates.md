---
type: pull_request
number: 1218
title: "Build: Bump the go group with 9 updates"
state: merged
author: dependabot
created: 2025-09-29T04:26:49Z
updated: 2025-09-29T07:44:10Z
closed: 2025-09-29T07:44:01Z
merged: 2025-09-29T07:44:01Z
base_branch: main
head_branch: dependabot/go_modules/go-5562d27cc2
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/1218
---

# Pull Request #1218: Build: Bump the go group with 9 updates

**Author**: @dependabot
**Created**: September 29, 2025 at 04:26 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-5562d27cc2`

## Description

Bumps the go group with 9 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/yummy](https://github.com/content-services/yummy) | `1.0.15` | `1.0.17` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.39.0` | `1.39.2` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.31.8` | `1.31.11` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.18.12` | `1.18.15` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.58.0` | `1.58.2` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.88.1` | `1.88.3` |
| [github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2](https://github.com/cloudevents/sdk-go) | `2.16.1` | `2.16.2` |
| [github.com/cloudevents/sdk-go/v2](https://github.com/cloudevents/sdk-go) | `2.16.1` | `2.16.2` |
| [github.com/content-services/zest/release/v2025](https://github.com/content-services/zest) | `2025.6.1750263718` | `2025.9.1758831219` |

Updates `github.com/content-services/yummy` from 1.0.15 to 1.0.17
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/yummy/releases">github.com/content-services/yummy's releases</a>.</em></p>
<blockquote>
<h2>v1.0.17</h2>
<h2>What's Changed</h2>
<ul>
<li>HMS-9183: replace yaml parser with goccy/go-yaml by <a href="https://github.com/Dugowitch"><code>@​Dugowitch</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/37">content-services/yummy#37</a></li>
<li>remove toolchain version from go.mod by <a href="https://github.com/rverdile"><code>@​rverdile</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/38">content-services/yummy#38</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Dugowitch"><code>@​Dugowitch</code></a> made their first contribution in <a href="https://redirect.github.com/content-services/yummy/pull/37">content-services/yummy#37</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.16...v1.0.17">https://github.com/content-services/yummy/compare/v1.0.16...v1.0.17</a></p>
<h2>v1.0.16</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/cloudflare/circl from 1.3.9 to 1.6.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/content-services/yummy/pull/33">content-services/yummy#33</a></li>
<li>Bump golang.org/x/crypto from 0.31.0 to 0.35.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/content-services/yummy/pull/32">content-services/yummy#32</a></li>
<li>Bump github.com/ulikunitz/xz from 0.5.12 to 0.5.14 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a>[bot] in <a href="https://redirect.github.com/content-services/yummy/pull/35">content-services/yummy#35</a></li>
<li>HMS-9066: Update go version to 1.24 by <a href="https://github.com/mayurilahane"><code>@​mayurilahane</code></a> in <a href="https://redirect.github.com/content-services/yummy/pull/36">content-services/yummy#36</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/mayurilahane"><code>@​mayurilahane</code></a> made their first contribution in <a href="https://redirect.github.com/content-services/yummy/pull/36">content-services/yummy#36</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/yummy/compare/v1.0.15...v1.0.16">https://github.com/content-services/yummy/compare/v1.0.15...v1.0.16</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/yummy/commit/a70df1df4d66fd8c97bfc31edb9d32475ef5421c"><code>a70df1d</code></a> remove toolchain version from go.mod (<a href="https://redirect.github.com/content-services/yummy/issues/38">#38</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/2c83b1b3949cd00fcd1610360b22a3d2cd1afeaf"><code>2c83b1b</code></a> HMS-9183: replace yaml parser with goccy/go-yaml (<a href="https://redirect.github.com/content-services/yummy/issues/37">#37</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/d90aefe1a7d18bfeaf10eb20c0d282c2ed933dca"><code>d90aefe</code></a> HMS-9066: Update go version to 1.24 (<a href="https://redirect.github.com/content-services/yummy/issues/36">#36</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/ca96d3d1c6f86fc27a5e48c6757e0ee21e38a102"><code>ca96d3d</code></a> Bump github.com/ulikunitz/xz from 0.5.12 to 0.5.14 (<a href="https://redirect.github.com/content-services/yummy/issues/35">#35</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/06580d2a01ce984c8c46637bca614114ce97c210"><code>06580d2</code></a> Bump golang.org/x/crypto from 0.31.0 to 0.35.0 (<a href="https://redirect.github.com/content-services/yummy/issues/32">#32</a>)</li>
<li><a href="https://github.com/content-services/yummy/commit/a482947dcbd9c0b15b4c9e7a79f7ef42dfce2273"><code>a482947</code></a> Bump github.com/cloudflare/circl from 1.3.9 to 1.6.1 (<a href="https://redirect.github.com/content-services/yummy/issues/33">#33</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/yummy/compare/v1.0.15...v1.0.17">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.39.0 to 1.39.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67db6904b816b95073883b7ad378384c4839b28c"><code>67db690</code></a> Release 2025-09-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32ee1b5d75fc303c0626a6f5e769f4e08cc491a8"><code>32ee1b5</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b431223309a815cffc048072556aa651ee1455f"><code>0b43122</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44786d920f3627b73a99e81c7b6399dbfcf7ab42"><code>44786d9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c98edb73809256823906d7e307ecf3c9abc16700"><code>c98edb7</code></a> update internal endpts comment that was wrong (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3194">#3194</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/88da3c8c5569dece0e99802dab638faa047a0db0"><code>88da3c8</code></a> Release 2025-09-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74a74fc179f8bbd879383cc75fa29a1937266dcc"><code>74a74fc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e6f7ae6139ca69044bb706664b4dbdc31227a32"><code>5e6f7ae</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e722ab42ff6bc6bb810c2937b8e1b41937e17c3"><code>0e722ab</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/41a7d004b9ff794f6007d30168afc825031f2c61"><code>41a7d00</code></a> Release 2025-09-24</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.39.0...v1.39.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.31.8 to 1.31.11
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67db6904b816b95073883b7ad378384c4839b28c"><code>67db690</code></a> Release 2025-09-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32ee1b5d75fc303c0626a6f5e769f4e08cc491a8"><code>32ee1b5</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b431223309a815cffc048072556aa651ee1455f"><code>0b43122</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44786d920f3627b73a99e81c7b6399dbfcf7ab42"><code>44786d9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c98edb73809256823906d7e307ecf3c9abc16700"><code>c98edb7</code></a> update internal endpts comment that was wrong (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3194">#3194</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/88da3c8c5569dece0e99802dab638faa047a0db0"><code>88da3c8</code></a> Release 2025-09-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74a74fc179f8bbd879383cc75fa29a1937266dcc"><code>74a74fc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e6f7ae6139ca69044bb706664b4dbdc31227a32"><code>5e6f7ae</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e722ab42ff6bc6bb810c2937b8e1b41937e17c3"><code>0e722ab</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/41a7d004b9ff794f6007d30168afc825031f2c61"><code>41a7d00</code></a> Release 2025-09-24</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.31.8...config/v1.31.11">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.18.12 to 1.18.15
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/89a5b69d4f68ca391b004269fd1f6a824e7bc936"><code>89a5b69</code></a> Release 2023-02-22</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c863429153eccf73a2afa12d92596d8b9a9a0e19"><code>c863429</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aa8a947960be4a8de7885d091c1d88f560825b16"><code>aa8a947</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e845552dbba693b68976db0cc1943ef20920d215"><code>e845552</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2037">#2037</a> from aws/fix-struct-errorcode</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cbeebabfb761a91719fa7ca41112a4e612ec8d56"><code>cbeebab</code></a> cl</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5c057fde8410c337ac739649de2fb6ae09dcf72c"><code>5c057fd</code></a> fix: prevent nil pointer deref on struct ErrorCode method</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0f8b6e7659bbc072af4a0d9dad7ef7b73e0cc9ba"><code>0f8b6e7</code></a> Release 2023-02-21</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e5cb2ee630838d582e97b5f6939d53b3f7e86bd2"><code>e5cb2ee</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6120ee681e6b8d7f215ed756ac26583e124ed500"><code>6120ee6</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/42acd6af1db8d13bf9c58fafba52a0167b2b1d82"><code>42acd6a</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.18.12...config/v1.18.15">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.58.0 to 1.58.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0599931efcf551dc63f84ec669d7fb5cfe14f64c"><code>0599931</code></a> Release 2024-07-10.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ca2a30ec8d823300ec948b784c491904eeba8929"><code>ca2a30e</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2dff5f9bddaf1f23ea57e1e04d80f6ce2d54221"><code>e2dff5f</code></a> temporarily re-add jmespath dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2712">#2712</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/383fd26928547348efed0ee1f8914d9e7a1b0287"><code>383fd26</code></a> Release 2024-07-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4a055f9d9e17eb7ef2206e8e37ba98d661d13e6a"><code>4a055f9</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e3457953351e6993f57f5ab8e373af8af70be45b"><code>e345795</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/52a10ac239cc1aab2d070fbb004f6445ca0ddcc0"><code>52a10ac</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/adab0de9e4ae75aee4f894c99fd1e2ce2de0035e"><code>adab0de</code></a> remove unused jmespath dependency from main module (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2707">#2707</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e07cc82b25692dce8f68e0b5bd0d0c5cdbcd279"><code>0e07cc8</code></a> Release 2024-07-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e3583451cd8a91dbca2cc22672c2cfa0c7860cf"><code>5e35834</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.58.0...service/s3/v1.58.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.88.1 to 1.88.3
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/67db6904b816b95073883b7ad378384c4839b28c"><code>67db690</code></a> Release 2025-09-26</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/32ee1b5d75fc303c0626a6f5e769f4e08cc491a8"><code>32ee1b5</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b431223309a815cffc048072556aa651ee1455f"><code>0b43122</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/44786d920f3627b73a99e81c7b6399dbfcf7ab42"><code>44786d9</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c98edb73809256823906d7e307ecf3c9abc16700"><code>c98edb7</code></a> update internal endpts comment that was wrong (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3194">#3194</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/88da3c8c5569dece0e99802dab638faa047a0db0"><code>88da3c8</code></a> Release 2025-09-25</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/74a74fc179f8bbd879383cc75fa29a1937266dcc"><code>74a74fc</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5e6f7ae6139ca69044bb706664b4dbdc31227a32"><code>5e6f7ae</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e722ab42ff6bc6bb810c2937b8e1b41937e17c3"><code>0e722ab</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/41a7d004b9ff794f6007d30168afc825031f2c61"><code>41a7d00</code></a> Release 2025-09-24</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.88.1...service/s3/v1.88.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2` from 2.16.1 to 2.16.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/af3e8599b3316ab6b4b73ff69aa8ec0efddbb5bb"><code>af3e859</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1184">#1184</a> from cloudevents/dependabot/github_actions/actions/s...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/bf5b8a6066d7650f84f862445ff60395bfa0aec1"><code>bf5b8a6</code></a> chore(deps): Bump actions/setup-go from 5.5.0 to 6.0.0</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/945d930a12f23f090bc739a21970afc1f9eb7d99"><code>945d930</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1180">#1180</a> from philicious/refactor-pubsubv2</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/35623f2b949d5d13c24103a3d7d1f3f9fbe74e78"><code>35623f2</code></a> chore: Replace deprecated grpc.Dial calls</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f0bd406957d7eb4d0ddf029c505bf45498de8b0a"><code>f0bd406</code></a> fix: Properly test that supplied ReceiveSettings are being used</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/ae89eba4471ea18a19fbed89e0fce720145e21bc"><code>ae89eba</code></a> Refactor pubsub protocol to use new upstream v2 library</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/34f7f6ffcdad49d7118bbd7603624bf2bbf11b37"><code>34f7f6f</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1183">#1183</a> from embano1/go-version</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/cafbb12d94d012b767efdf7b1bca5a8ce293115e"><code>cafbb12</code></a> chore: update go version in workflows</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/0a1147cb5ec2870334309b473df8308e28adb1cb"><code>0a1147c</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1182">#1182</a> from cloudevents/automated-dependency-updates</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/835c4867e90c93c13b2d244af210a646977df0f4"><code>835c486</code></a> chore: update dependencies</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.16.1...v2.16.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/cloudevents/sdk-go/v2` from 2.16.1 to 2.16.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cloudevents/sdk-go/commit/af3e8599b3316ab6b4b73ff69aa8ec0efddbb5bb"><code>af3e859</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1184">#1184</a> from cloudevents/dependabot/github_actions/actions/s...</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/bf5b8a6066d7650f84f862445ff60395bfa0aec1"><code>bf5b8a6</code></a> chore(deps): Bump actions/setup-go from 5.5.0 to 6.0.0</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/945d930a12f23f090bc739a21970afc1f9eb7d99"><code>945d930</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1180">#1180</a> from philicious/refactor-pubsubv2</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/35623f2b949d5d13c24103a3d7d1f3f9fbe74e78"><code>35623f2</code></a> chore: Replace deprecated grpc.Dial calls</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/f0bd406957d7eb4d0ddf029c505bf45498de8b0a"><code>f0bd406</code></a> fix: Properly test that supplied ReceiveSettings are being used</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/ae89eba4471ea18a19fbed89e0fce720145e21bc"><code>ae89eba</code></a> Refactor pubsub protocol to use new upstream v2 library</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/34f7f6ffcdad49d7118bbd7603624bf2bbf11b37"><code>34f7f6f</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1183">#1183</a> from embano1/go-version</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/cafbb12d94d012b767efdf7b1bca5a8ce293115e"><code>cafbb12</code></a> chore: update go version in workflows</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/0a1147cb5ec2870334309b473df8308e28adb1cb"><code>0a1147c</code></a> Merge pull request <a href="https://redirect.github.com/cloudevents/sdk-go/issues/1182">#1182</a> from cloudevents/automated-dependency-updates</li>
<li><a href="https://github.com/cloudevents/sdk-go/commit/835c4867e90c93c13b2d244af210a646977df0f4"><code>835c486</code></a> chore: update dependencies</li>
<li>Additional commits viewable in <a href="https://github.com/cloudevents/sdk-go/compare/v2.16.1...v2.16.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2025` from 2025.6.1750263718 to 2025.9.1758831219
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/41eefb1365353bc727584ce8f9031cf430336668"><code>41eefb1</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7e4488333e4087e8639db952b...</li>
<li><a href="https://github.com/content-services/zest/commit/2508f3bab675296018c403127962812d328549a9"><code>2508f3b</code></a> Switch back to konflux build</li>
<li><a href="https://github.com/content-services/zest/commit/471a4a9b30b780047c2efccc451674c2e1ba618e"><code>471a4a9</code></a> bump: go version to 1.24 (<a href="https://redirect.github.com/content-services/zest/issues/26">#26</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2025.6.1750263718...release/v2025.9.1758831219">compare view</a></li>
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

### Review by @swadeley - Approved on September 29, 2025 at 07:43 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1218*
