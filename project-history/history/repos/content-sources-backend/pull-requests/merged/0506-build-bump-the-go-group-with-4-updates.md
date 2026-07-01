---
type: pull_request
number: 506
title: "Build: Bump the go group with 4 updates"
state: merged
author: dependabot
created: 2023-12-18T04:24:46Z
updated: 2023-12-18T18:55:28Z
closed: 2023-12-18T18:55:19Z
merged: 2023-12-18T18:55:19Z
base_branch: main
head_branch: dependabot/go_modules/go-30af6c8bb0
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/506
---

# Pull Request #506: Build: Bump the go group with 4 updates

**Author**: @dependabot
**Created**: December 18, 2023 at 04:24 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-30af6c8bb0`

## Description

Bumps the go group with 4 updates: [github.com/google/uuid](https://github.com/google/uuid), [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2), [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) and [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils).

Updates `github.com/google/uuid` from 1.4.0 to 1.5.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/google/uuid/releases">github.com/google/uuid's releases</a>.</em></p>
<blockquote>
<h2>v1.5.0</h2>
<h2><a href="https://github.com/google/uuid/compare/v1.4.0...v1.5.0">1.5.0</a> (2023-12-12)</h2>
<h3>Features</h3>
<ul>
<li>Validate UUID without creating new UUID (<a href="https://redirect.github.com/google/uuid/issues/141">#141</a>) (<a href="https://github.com/google/uuid/commit/9ee7366e66c9ad96bab89139418a713dc584ae29">9ee7366</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/google/uuid/blob/master/CHANGELOG.md">github.com/google/uuid's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/google/uuid/compare/v1.4.0...v1.5.0">1.5.0</a> (2023-12-12)</h2>
<h3>Features</h3>
<ul>
<li>Validate UUID without creating new UUID (<a href="https://redirect.github.com/google/uuid/issues/141">#141</a>) (<a href="https://github.com/google/uuid/commit/9ee7366e66c9ad96bab89139418a713dc584ae29">9ee7366</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/google/uuid/commit/4d47f8eb066f43cfaedd728a543479d9c9dfa8f6"><code>4d47f8e</code></a> chore(master): release 1.5.0 (<a href="https://redirect.github.com/google/uuid/issues/145">#145</a>)</li>
<li><a href="https://github.com/google/uuid/commit/9ee7366e66c9ad96bab89139418a713dc584ae29"><code>9ee7366</code></a> feat: Validate UUID without creating new UUID (<a href="https://redirect.github.com/google/uuid/issues/141">#141</a>)</li>
<li><a href="https://github.com/google/uuid/commit/b35aa6a595277504b1ec94c520d4091ec050b9d5"><code>b35aa6a</code></a> add uuid version 6 and 7 (<a href="https://redirect.github.com/google/uuid/issues/139">#139</a>)</li>
<li>See full diff in <a href="https://github.com/google/uuid/compare/v1.4.0...v1.5.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.29.5 to 1.30.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs's changelog</a>.</em></p>
<blockquote>
<h1>Release (2023-01-05)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/accessanalyzer</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/accessanalyzer/CHANGELOG.md#v1190-2023-01-05">v1.19.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/account</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/account/CHANGELOG.md#v180-2023-01-05">v1.8.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/acm</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/acm/CHANGELOG.md#v1170-2023-01-05">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/acmpca</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/acmpca/CHANGELOG.md#v1200-2023-01-05">v1.20.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/alexaforbusiness</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/alexaforbusiness/CHANGELOG.md#v1150-2023-01-05">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amp</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/amp/CHANGELOG.md#v1160-2023-01-05">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplify</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/amplify/CHANGELOG.md#v1130-2023-01-05">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplifybackend</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/amplifybackend/CHANGELOG.md#v1140-2023-01-05">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
<li><strong>Feature</strong>: Updated GetBackendAPIModels response to include ModelIntrospectionSchema json string</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/amplifyuibuilder/CHANGELOG.md#v190-2023-01-05">v1.9.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apigateway</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/apigateway/CHANGELOG.md#v1160-2023-01-05">v1.16.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/apigatewaymanagementapi/CHANGELOG.md#v1110-2023-01-05">v1.11.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apigatewayv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/apigatewayv2/CHANGELOG.md#v1130-2023-01-05">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appconfig</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/appconfig/CHANGELOG.md#v1150-2023-01-05">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appconfigdata</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/appconfigdata/CHANGELOG.md#v150-2023-01-05">v1.5.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appflow</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/appflow/CHANGELOG.md#v1230-2023-01-05">v1.23.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appintegrations</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/appintegrations/CHANGELOG.md#v1140-2023-01-05">v1.14.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationautoscaling</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/applicationautoscaling/CHANGELOG.md#v1170-2023-01-05">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/applicationcostprofiler/CHANGELOG.md#v1100-2023-01-05">v1.10.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/applicationdiscoveryservice/CHANGELOG.md#v1150-2023-01-05">v1.15.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/applicationinsights</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/applicationinsights/CHANGELOG.md#v1170-2023-01-05">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/appmesh</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/appmesh/CHANGELOG.md#v1170-2023-01-05">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Add <code>ErrorCodeOverride</code><code>aws/smithy-go#401</code></li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/apprunner</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/service/s3/v1.30.0/service/apprunner/CHANGELOG.md#v1160-2023-01-05">v1.16.0</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/116a622a26e2b4c0a47749c132faaee81ff76b7e"><code>116a622</code></a> Release 2023-01-05</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ce039452b6a7e2046ee8e6292260cc9d1e78af7e"><code>ce03945</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/095bbfff591b1bccef0d87f315171cf9b2a1db2f"><code>095bbff</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2998a9800a670c002b68787fa37b059f1bd57674"><code>2998a98</code></a> Regenerate clients with <code>ErrorCodeOverride</code> (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1969">#1969</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1b0a07d93d9e22efb8b426e4025bca10b0d9c296"><code>1b0a07d</code></a> Release 2023-01-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff5b1c7a27c061145c37e8169a285a91a2aabc6c"><code>ff5b1c7</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cabea36bb4efa8711271d6dac95e4666efe8a4b4"><code>cabea36</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cd385dc3b81a32b8f75a499259b4139a7634058c"><code>cd385dc</code></a> Update links to point to smithy.io</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4dd79b8978186cc289c7e2eb70b14d40ec179051"><code>4dd79b8</code></a> Rename SyntheticClone to Synthetic</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b302f0a86c80ee4c17b0a3d1f30871d7a9cb8a93"><code>b302f0a</code></a> Release 2023-01-03</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.29.5...service/s3/v1.30.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2023` from 2023.12.1701890198 to 2023.12.1702603647
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/c5e9446bb7c9b5f02e1f3abb3c1a2feb97b442fd"><code>c5e9446</code></a> Update pulp bindings to de24b382d968bd5ed3a4252b36e475e6664d66cb7a98d596b8e32...</li>
<li><a href="https://github.com/content-services/zest/commit/554464a62ee45bc1818336f9eb14082b2a6dbede"><code>554464a</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a972eda53626057bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/c919c960cc8f02f772048749d04ee4b17105cd51"><code>c919c96</code></a> Update pulp bindings to de5624ba39358ba4ba962535e4a972ed25366ac57bd286834eded...</li>
<li><a href="https://github.com/content-services/zest/commit/3a6d6c1c26f2fa8e4b13d63d2048ef97f9373af7"><code>3a6d6c1</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73e8588dd8037d2e4353b86e2...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2023.12.1701890198...release/v2023.12.1702603647">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.24.12 to 1.24.13
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>v1.24.13</h2>
<h2>What's Changed</h2>
<ul>
<li>Bump github.com/aws/aws-sdk-go from 1.48.10 to 1.48.12 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/368">RedHatInsights/insights-operator-utils#368</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.12 to 1.48.15 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/372">RedHatInsights/insights-operator-utils#372</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.15 to 1.48.16 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/373">RedHatInsights/insights-operator-utils#373</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.16 to 1.49.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/374">RedHatInsights/insights-operator-utils#374</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.24.12...v1.24.13">https://github.com/RedHatInsights/insights-operator-utils/compare/v1.24.12...v1.24.13</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/aa7338bed02523f859b1c51023a72b584a6e1953"><code>aa7338b</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/374">#374</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4f66cba7fdc93c009cf32d9f3be8a760ef51b756"><code>4f66cba</code></a> Bump github.com/aws/aws-sdk-go from 1.48.16 to 1.49.0</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/311c2f4874e43763f04d49895aac6411a5a6f946"><code>311c2f4</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/373">#373</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/7981983bc7e16df176afd2dc903f51941447950e"><code>7981983</code></a> Bump github.com/aws/aws-sdk-go from 1.48.15 to 1.48.16</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/f0d0494ee064a0a2fbd5be2ab2bb19608a24d806"><code>f0d0494</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/372">#372</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/0a8ae025d5c01ca4d0cb0b9901a5df54f99747cd"><code>0a8ae02</code></a> Bump github.com/aws/aws-sdk-go from 1.48.12 to 1.48.15</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/78f18a34196a3dccc98d2a0287ea2956c3c3545a"><code>78f18a3</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/368">#368</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/83b381bf10d54acaef2f85ea5e9b20f6c9e38218"><code>83b381b</code></a> Bump github.com/aws/aws-sdk-go from 1.48.10 to 1.48.12</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.24.12...v1.24.13">compare view</a></li>
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

### Comment by @app-sre-bot on December 18, 2023 at 04:25 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on December 18, 2023 at 02:34 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on December 18, 2023 at 06:55 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/506*
