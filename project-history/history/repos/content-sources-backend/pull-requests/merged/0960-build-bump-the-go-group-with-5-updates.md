---
type: pull_request
number: 960
title: "Build: Bump the go group with 5 updates"
state: merged
author: dependabot
created: 2025-02-03T04:49:54Z
updated: 2025-02-03T13:06:15Z
closed: 2025-02-03T13:06:06Z
merged: 2025-02-03T13:06:06Z
base_branch: main
head_branch: dependabot/go_modules/go-66fbf66651
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/960
---

# Pull Request #960: Build: Bump the go group with 5 updates

**Author**: @dependabot
**Created**: February 03, 2025 at 04:49 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-66fbf66651`

## Description

Bumps the go group with 5 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.127.1-0.20240830113739-c606b5546b12` | `0.129.0` |
| [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) | `4.18.1` | `4.18.2` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.34.0` | `1.36.0` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.55` | `1.17.57` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.45.8` | `1.45.10` |

Updates `github.com/getkin/kin-openapi` from 0.127.1-0.20240830113739-c606b5546b12 to 0.129.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.129.0</h2>
<h2>What's Changed</h2>
<ul>
<li>README: add Fuego to dependents by <a href="https://github.com/EwenQuim"><code>@​EwenQuim</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1017">getkin/kin-openapi#1017</a></li>
<li>openapi3: skip a test in CI to avoid 403s from some remote server by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1019">getkin/kin-openapi#1019</a></li>
<li>openapi3: introduce StringMap type to enable unmarshalling of maps with Origin by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1018">getkin/kin-openapi#1018</a></li>
<li>openapi3: reference originating locations in YAML specs - step 1 by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1007">getkin/kin-openapi#1007</a></li>
<li>openapi3: reference originating locations in YAML specs - step 2 by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1024">getkin/kin-openapi#1024</a></li>
<li>openapi3: process discriminator mapping values as refs by <a href="https://github.com/jgresty"><code>@​jgresty</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1022">getkin/kin-openapi#1022</a></li>
<li>openapi3filter: register decoder for other JSON content types by <a href="https://github.com/oliverli"><code>@​oliverli</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1026">getkin/kin-openapi#1026</a></li>
<li>Revert &quot;openapi3: process discriminator mapping values as refs&quot; by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1029">getkin/kin-openapi#1029</a></li>
<li>openapi3: fail to load spec because of schema names in mapping  by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1027">getkin/kin-openapi#1027</a></li>
<li>openapi2conv: convert schemaRef for additional props by <a href="https://github.com/jayanth-tatina-groww"><code>@​jayanth-tatina-groww</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1030">getkin/kin-openapi#1030</a></li>
<li>openapi3: simplify by replacing math.Min with min by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1032">getkin/kin-openapi#1032</a></li>
<li>openapi3: fix deprecation comments by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1034">getkin/kin-openapi#1034</a></li>
<li>test: fix expected-actual parameters in require.Equal by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1035">getkin/kin-openapi#1035</a></li>
<li>use forked yaml modules without &quot;replace&quot; by <a href="https://github.com/reuvenharrison"><code>@​reuvenharrison</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1038">getkin/kin-openapi#1038</a></li>
<li>openapi3: update date schema formats to not match months or days of '00' by <a href="https://github.com/RulerOfTheQueendom"><code>@​RulerOfTheQueendom</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1042">getkin/kin-openapi#1042</a></li>
<li>openapi3,openapi3filter: replace interface{} with any by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1040">getkin/kin-openapi#1040</a></li>
<li>openapi3filter: Simplify ValidateRequest implementation by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1041">getkin/kin-openapi#1041</a></li>
<li>openapi3filter: validation of <code>x-www-form-urlencoded</code> with arbitrary nested allOf by <a href="https://github.com/mikhalytch"><code>@​mikhalytch</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1046">getkin/kin-openapi#1046</a></li>
<li>openapi2conv: convert references in nested additionalProperties schemas by <a href="https://github.com/travisnewhouse"><code>@​travisnewhouse</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1047">getkin/kin-openapi#1047</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/EwenQuim"><code>@​EwenQuim</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1017">getkin/kin-openapi#1017</a></li>
<li><a href="https://github.com/jgresty"><code>@​jgresty</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1022">getkin/kin-openapi#1022</a></li>
<li><a href="https://github.com/oliverli"><code>@​oliverli</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1026">getkin/kin-openapi#1026</a></li>
<li><a href="https://github.com/jayanth-tatina-groww"><code>@​jayanth-tatina-groww</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1030">getkin/kin-openapi#1030</a></li>
<li><a href="https://github.com/RulerOfTheQueendom"><code>@​RulerOfTheQueendom</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1042">getkin/kin-openapi#1042</a></li>
<li><a href="https://github.com/mikhalytch"><code>@​mikhalytch</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1046">getkin/kin-openapi#1046</a></li>
<li><a href="https://github.com/travisnewhouse"><code>@​travisnewhouse</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1047">getkin/kin-openapi#1047</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.128.0...v0.129.0">https://github.com/getkin/kin-openapi/compare/v0.128.0...v0.129.0</a></p>
<h2>v0.128.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3filter: Fix default value for array in for query param by <a href="https://github.com/Tommy-42"><code>@​Tommy-42</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1000">getkin/kin-openapi#1000</a></li>
<li>Add github.com/pb33f/libopenapi by <a href="https://github.com/Jille"><code>@​Jille</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1004">getkin/kin-openapi#1004</a></li>
<li>Introduce an option to override the regex implementation by <a href="https://github.com/alexbakker"><code>@​alexbakker</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1006">getkin/kin-openapi#1006</a></li>
<li>make form required field order deterministic by <a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1008">getkin/kin-openapi#1008</a></li>
<li>openapi2: fix un/marshalling discriminator field by <a href="https://github.com/reversearrow"><code>@​reversearrow</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1011">getkin/kin-openapi#1011</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Tommy-42"><code>@​Tommy-42</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1000">getkin/kin-openapi#1000</a></li>
<li><a href="https://github.com/Jille"><code>@​Jille</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1004">getkin/kin-openapi#1004</a></li>
<li><a href="https://github.com/alexbakker"><code>@​alexbakker</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1006">getkin/kin-openapi#1006</a></li>
<li><a href="https://github.com/jlsherrill"><code>@​jlsherrill</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/1008">getkin/kin-openapi#1008</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.127.0...v0.128.0">https://github.com/getkin/kin-openapi/compare/v0.127.0...v0.128.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/getkin/kin-openapi/commits/v0.129.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/golang-migrate/migrate/v4` from 4.18.1 to 4.18.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/golang-migrate/migrate/releases">github.com/golang-migrate/migrate/v4's releases</a>.</em></p>
<blockquote>
<h2>v4.18.2</h2>
<h2>Changelog</h2>
<ul>
<li>e145cde Bump github.com/golang-jwt/jwt/v4 from 4.4.2 to 4.5.1</li>
<li>e22d012 Don't output sensitive information on error</li>
<li>e5a152b Drop support for Azure SQL Edge</li>
<li>12c619e Fix CI (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1222">#1222</a>)</li>
<li>bc06922 Update dktest from v0.4.3 to v0.4.4</li>
<li>7651c8a linter fixes</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/golang-migrate/migrate/commit/d477553dcd9ae33e2698c09f186e988767d71b79"><code>d477553</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1195">#1195</a> from golang-migrate/dependabot/go_modules/github.com...</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/e145cded4a5df0f5a8c4fb3aef34aa66f15ad2b9"><code>e145cde</code></a> Bump github.com/golang-jwt/jwt/v4 from 4.4.2 to 4.5.1</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/7651c8a37ed0b473ce5d686830956a455a20a290"><code>7651c8a</code></a> linter fixes</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/e5a152b9f6499b0534300967f914856119931ed6"><code>e5a152b</code></a> Drop support for Azure SQL Edge</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/12c619ee47b713bf4855f3dd6ea6e9a68b719a07"><code>12c619e</code></a> Fix CI (<a href="https://redirect.github.com/golang-migrate/migrate/issues/1222">#1222</a>)</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/bc06922b4de8631b553f99600e564b161e79e798"><code>bc06922</code></a> Update dktest from v0.4.3 to v0.4.4</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/c378583d782e026f472dff657bfd088bf2510038"><code>c378583</code></a> Merge pull request <a href="https://redirect.github.com/golang-migrate/migrate/issues/1162">#1162</a> from rselbach/rselbach/no-sensitive-info</li>
<li><a href="https://github.com/golang-migrate/migrate/commit/e22d0128ccf15813fa40e3720010e5e6710ee3bb"><code>e22d012</code></a> Don't output sensitive information on error</li>
<li>See full diff in <a href="https://github.com/golang-migrate/migrate/compare/v4.18.1...v4.18.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.34.0 to 1.36.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2e9697d8ebe330a7435716c2f31b1cea4dff3c0"><code>e2e9697</code></a> Release 2025-01-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6576a0939a79d5f31eef10164750faedd78a45d4"><code>6576a09</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f762573ab5d9286d9751d49091f6a240c12c0742"><code>f762573</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c94df29ecd457e8ec40931fd2fe54d8da2f93ce2"><code>c94df29</code></a> add transfer manager doc header (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2990">#2990</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/880543ce2034570eb3b93c4811289c3b0e55600f"><code>880543c</code></a> revert the revert on the transfer manager beta (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2993">#2993</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8da49e527e317a77ef0f1b2f52b4dc72a4fbd720"><code>8da49e5</code></a> switch to code-generated waiters for remaining services (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2994">#2994</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c7c68659ce67e5b7e18f31bc66068cec9e3d790d"><code>c7c6865</code></a> Release 2025-01-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70f736c5dc0b8652c5fe5c387b2165c3b9beddb1"><code>70f736c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/28731c2bdef3c2555a95632396b6d4936e58099d"><code>28731c2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3505e4b255c327a1fa38f870612c327b93302dc0"><code>3505e4b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.34.0...v1.36.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.55 to 1.17.57
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2e9697d8ebe330a7435716c2f31b1cea4dff3c0"><code>e2e9697</code></a> Release 2025-01-31</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6576a0939a79d5f31eef10164750faedd78a45d4"><code>6576a09</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f762573ab5d9286d9751d49091f6a240c12c0742"><code>f762573</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c94df29ecd457e8ec40931fd2fe54d8da2f93ce2"><code>c94df29</code></a> add transfer manager doc header (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2990">#2990</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/880543ce2034570eb3b93c4811289c3b0e55600f"><code>880543c</code></a> revert the revert on the transfer manager beta (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2993">#2993</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8da49e527e317a77ef0f1b2f52b4dc72a4fbd720"><code>8da49e5</code></a> switch to code-generated waiters for remaining services (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2994">#2994</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c7c68659ce67e5b7e18f31bc66068cec9e3d790d"><code>c7c6865</code></a> Release 2025-01-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70f736c5dc0b8652c5fe5c387b2165c3b9beddb1"><code>70f736c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/28731c2bdef3c2555a95632396b6d4936e58099d"><code>28731c2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3505e4b255c327a1fa38f870612c327b93302dc0"><code>3505e4b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.55...credentials/v1.17.57">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.8 to 1.45.10
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
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/polly/v1.45.8...service/polly/v1.45.10">compare view</a></li>
</ul>
</details>
<br />

<details>
<summary>Most Recent Ignore Conditions Applied to This Pull Request</summary>

| Dependency Name | Ignore Conditions |
| --- | --- |
| github.com/getkin/kin-openapi | [>= 0.128.a, < 0.129] |
</details>


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

### Review by @jlsherrill - Approved on February 03, 2025 at 01:06 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/960*
