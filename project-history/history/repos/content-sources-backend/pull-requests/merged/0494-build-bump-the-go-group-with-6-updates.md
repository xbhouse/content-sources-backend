---
type: pull_request
number: 494
title: "Build: Bump the go group with 6 updates"
state: merged
author: dependabot
created: 2023-12-04T04:42:44Z
updated: 2023-12-04T16:08:10Z
closed: 2023-12-04T16:08:05Z
merged: 2023-12-04T16:08:05Z
base_branch: main
head_branch: dependabot/go_modules/go-1bd25d55a2
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/494
---

# Pull Request #494: Build: Bump the go group with 6 updates

**Author**: @dependabot
**Created**: December 04, 2023 at 04:42 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-1bd25d55a2`

## Description

Bumps the go group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi) | `0.120.0` | `0.122.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.23.1` | `1.23.5` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.16.4` | `1.16.9` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.28.0` | `1.29.2` |
| [github.com/content-services/zest/release/v2023](https://github.com/content-services/zest) | `2023.11.1700250630` | `2023.11.1701177874` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.24.11` | `1.24.12` |

Updates `github.com/getkin/kin-openapi` from 0.120.0 to 0.122.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.122.0</h2>
<h2>What's Changed</h2>
<ul>
<li>docs.sh: fix narrow docs checks spectrum by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/877">getkin/kin-openapi#877</a></li>
<li>fix after <a href="https://redirect.github.com/getkin/kin-openapi/issues/870">#870</a>: make sure Bis does not surface up by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/878">getkin/kin-openapi#878</a></li>
<li>openapi3: add support for extensions on the few types left by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/763">getkin/kin-openapi#763</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.121.0...v0.122.0">https://github.com/getkin/kin-openapi/compare/v0.121.0...v0.122.0</a></p>
<h2>v0.121.0</h2>
<h2>What's Changed</h2>
<ul>
<li>Replace deprecated io/ioutil with io by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/840">getkin/kin-openapi#840</a></li>
<li>fix compile errors in README example code by <a href="https://github.com/Jumpaku"><code>@​Jumpaku</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/842">getkin/kin-openapi#842</a></li>
<li>openapi3: add length of circular reference to error message by <a href="https://github.com/jamietanna"><code>@​jamietanna</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/839">getkin/kin-openapi#839</a></li>
<li>openapi2conv: fix that removes an extra <code>required</code> array from <code>formDataBody</code> by <a href="https://github.com/liankui"><code>@​liankui</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/848">getkin/kin-openapi#848</a></li>
<li>openapi3: fix schema validation for <code>anyOf</code>, <code>allOf</code>, <code>oneOf</code> operations by <a href="https://github.com/AmadeusK525"><code>@​AmadeusK525</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/850">getkin/kin-openapi#850</a></li>
<li>reproduce incorrect allOf + nullable behaviour by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/253">getkin/kin-openapi#253</a></li>
<li>reproduce issue <a href="https://redirect.github.com/getkin/kin-openapi/issues/423">#423</a> by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/432">getkin/kin-openapi#432</a></li>
<li>openapi2conv: expose loader and location to ToV3 by <a href="https://github.com/AriehSchneier"><code>@​AriehSchneier</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/851">getkin/kin-openapi#851</a></li>
<li>openAPI3 template T class add server helper function by <a href="https://github.com/spdrcd"><code>@​spdrcd</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/852">getkin/kin-openapi#852</a></li>
<li>openapi3: handle tangible *nil in Schema.IsEmpty impl by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/858">getkin/kin-openapi#858</a></li>
<li>openapi3: no longer error when drilling into a silenced yaml tag by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/859">getkin/kin-openapi#859</a></li>
<li>openapi3: shave some tests by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/862">getkin/kin-openapi#862</a></li>
<li>openapi3: refacto drill bit by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/865">getkin/kin-openapi#865</a></li>
<li>openapi3: unify Loader.resolve*Ref impls by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/863">getkin/kin-openapi#863</a></li>
<li>openapi3: refacto path resolving, remove a couple error cases by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/864">getkin/kin-openapi#864</a></li>
<li>openapi3: fast path drilling into UFO-type refs by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/866">getkin/kin-openapi#866</a></li>
<li>cmd/validate: read v2 documents that dont mention swagger as well by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/868">getkin/kin-openapi#868</a></li>
<li>openapi3: save some allocs when checking extras in refs by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/869">getkin/kin-openapi#869</a></li>
<li>openapi3: fix validating extra siblings in remote refs by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/872">getkin/kin-openapi#872</a></li>
<li>openapi3: reclaims Extensions unused mem post-unmarshal by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/871">getkin/kin-openapi#871</a></li>
<li>openapi3: handle refs missing fragment by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/511">getkin/kin-openapi#511</a></li>
<li>openapi3: support \uC4FE codepoint syntax in Schema.Pattern by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/873">getkin/kin-openapi#873</a></li>
<li>close <a href="https://redirect.github.com/getkin/kin-openapi/issues/594">#594</a>: yaml &quot;control characters are not allowed&quot; no longer reproducible by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/597">getkin/kin-openapi#597</a></li>
<li>openapi{2,3}: simplify unmarshal errors by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/870">getkin/kin-openapi#870</a></li>
<li>openapi3: refacto ref-resolving end conditions by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/874">getkin/kin-openapi#874</a></li>
<li>openapi3: rename type of Components.Responses to ResponseBodies (from Responses) by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/875">getkin/kin-openapi#875</a></li>
<li>openapi3: correct implementations of JSONLookup by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/876">getkin/kin-openapi#876</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/Jumpaku"><code>@​Jumpaku</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/842">getkin/kin-openapi#842</a></li>
<li><a href="https://github.com/jamietanna"><code>@​jamietanna</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/839">getkin/kin-openapi#839</a></li>
<li><a href="https://github.com/liankui"><code>@​liankui</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/848">getkin/kin-openapi#848</a></li>
<li><a href="https://github.com/AmadeusK525"><code>@​AmadeusK525</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/850">getkin/kin-openapi#850</a></li>
<li><a href="https://github.com/AriehSchneier"><code>@​AriehSchneier</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/851">getkin/kin-openapi#851</a></li>
<li><a href="https://github.com/spdrcd"><code>@​spdrcd</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/852">getkin/kin-openapi#852</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.120.0...v0.121.0">https://github.com/getkin/kin-openapi/compare/v0.120.0...v0.121.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/6740cd2b25853647fafc2e40c33d60f6f78b4fd2"><code>6740cd2</code></a> openapi3: add support for extensions on the few types left (<a href="https://redirect.github.com/getkin/kin-openapi/issues/763">#763</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/f6d1b8c132b1ba26fe4054238278b1c0c016248f"><code>f6d1b8c</code></a> fix after <a href="https://redirect.github.com/getkin/kin-openapi/issues/870">#870</a>: make sure Bis does not surface up (<a href="https://redirect.github.com/getkin/kin-openapi/issues/878">#878</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/0cc5e2214e16fe574ac2c4867fa8bbabbddfa6e6"><code>0cc5e22</code></a> docs.sh: fix narrow docs checks spectrum (<a href="https://redirect.github.com/getkin/kin-openapi/issues/877">#877</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/4e7d03195bde822fdfe798fd04dbee5b0e3109d9"><code>4e7d031</code></a> openapi3: correct implementations of JSONLookup (<a href="https://redirect.github.com/getkin/kin-openapi/issues/876">#876</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/c1681a9dadb2a5e9e579ce54b6a026bb902cd976"><code>c1681a9</code></a> openapi3: rename type of Components.Responses to ResponseBodies (from Respons...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/663b0ddf417355c6c9bd3a412c90bedc55e4e584"><code>663b0dd</code></a> openapi3: refacto ref-resolving end conditions (<a href="https://redirect.github.com/getkin/kin-openapi/issues/874">#874</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/377bb40b6e2b46cd93cbd0a6ee2065f0999e489d"><code>377bb40</code></a> openapi{2,3}: simplify unmarshal errors (<a href="https://redirect.github.com/getkin/kin-openapi/issues/870">#870</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/fe1c5f5e84a3ed4b7c4ff55b3a5d3ac53d569257"><code>fe1c5f5</code></a> close <a href="https://redirect.github.com/getkin/kin-openapi/issues/594">#594</a>: yaml &quot;control characters are not allowed&quot; no longer reproducible ...</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/e7a726a9f8e85c51f997ab5b15f70c571b968533"><code>e7a726a</code></a> openapi3: support \uC4FE codepoint syntax in Schema.Pattern (<a href="https://redirect.github.com/getkin/kin-openapi/issues/873">#873</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/582e6d06120fa34d3a4bac4123d1924218b4c678"><code>582e6d0</code></a> openapi3: handle refs missing fragment (<a href="https://redirect.github.com/getkin/kin-openapi/issues/511">#511</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.120.0...v0.122.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.23.1 to 1.23.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ce842a7ed4981d404542e66d3ccd71f8bd7d8b2c"><code>ce842a7</code></a> Release 2023-12-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d40f9232ab713206c290db84ded240da03bdd490"><code>d40f923</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3a00ef501de9570bd0ae9764bc16594ffe033e65"><code>3a00ef5</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1f61a067923798373d8a7b782139f5824c470acc"><code>1f61a06</code></a> fix: correct recognition and zeroing of cache-wrapped AnonymousCredentials (#...</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/623d430e6204c294fb0f24c71e1f71878baa498d"><code>623d430</code></a> fix: correct wrapping of errors in authentication workflow (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2403">#2403</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/cb676e7c25486166bfc40d2758057613e1ea1d66"><code>cb676e7</code></a> Release 2023-11-30.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21db1a2c403944cd167f11f1debcebf3edf6e65f"><code>21db1a2</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/833655d15336d67d8633acb5f99edb8def18e735"><code>833655d</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fcfb7acd506e14bbb98ada6db0a8498ec17b4d53"><code>fcfb7ac</code></a> fix: use region overrides in endpoint discovery (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2393">#2393</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f3ea5b8aca81cfc644b822ffc9910f2382e474c2"><code>f3ea5b8</code></a> Release 2023-11-30</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.23.1...v1.23.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.16.4 to 1.16.9
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/CHANGELOG.md">github.com/aws/aws-sdk-go-v2/credentials's changelog</a>.</em></p>
<blockquote>
<h1>Release (2022-08-08)</h1>
<h2>General Highlights</h2>
<ul>
<li><strong>Dependency Update</strong>: Updated to the latest SDK module versions</li>
</ul>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2</code>: v1.16.9
<ul>
<li><strong>Bug Fix</strong>: aws/signer/v4: Fixes a panic in SDK's handling of endpoint URLs with ports by correcting how URL path is parsed from opaque URLs. Fixes <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1294">#1294</a>.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/glue</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/glue/CHANGELOG.md#v1290-2022-08-08">v1.29.0</a>
<ul>
<li><strong>Feature</strong>: Add an option to run non-urgent or non-time sensitive Glue Jobs on spare capacity</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/identitystore</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/identitystore/CHANGELOG.md#v11410-2022-08-08">v1.14.10</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates to reflect service rename - AWS IAM Identity Center (successor to AWS Single Sign-On)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iotwireless</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/iotwireless/CHANGELOG.md#v1210-2022-08-08">v1.21.0</a>
<ul>
<li><strong>Feature</strong>: AWS IoT Wireless release support for sidewalk data reliability.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/pinpoint</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/pinpoint/CHANGELOG.md#v1170-2022-08-08">v1.17.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for Advance Quiet Time in Journeys. Adds RefreshOnSegmentUpdate and WaitForQuietTime to JourneyResponse.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/quicksight</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/quicksight/CHANGELOG.md#v1232-2022-08-08">v1.23.2</a>
<ul>
<li><strong>Documentation</strong>: A series of documentation updates to the QuickSight API reference.</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/sso</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/sso/CHANGELOG.md#v11114-2022-08-08">v1.11.14</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates to reflect service rename - AWS IAM Identity Center (successor to AWS Single Sign-On)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ssoadmin</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/ssoadmin/CHANGELOG.md#v1152-2022-08-08">v1.15.2</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates to reflect service rename - AWS IAM Identity Center (successor to AWS Single Sign-On)</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/ssooidc</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/ssooidc/CHANGELOG.md#v11212-2022-08-08">v1.12.12</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates to reflect service rename - AWS IAM Identity Center (successor to AWS Single Sign-On)</li>
</ul>
</li>
</ul>
<h1>Release (2022-08-04)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/chimesdkmeetings/CHANGELOG.md#v1130-2022-08-04">v1.13.0</a>
<ul>
<li><strong>Feature</strong>: Adds support for Tags on Amazon Chime SDK WebRTC sessions</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/configservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/configservice/CHANGELOG.md#v1240-2022-08-04">v1.24.0</a>
<ul>
<li><strong>Feature</strong>: Add resourceType enums for Athena, GlobalAccelerator, Detective and EC2 types</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/databasemigrationservice</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/databasemigrationservice/CHANGELOG.md#v1213-2022-08-04">v1.21.3</a>
<ul>
<li><strong>Documentation</strong>: Documentation updates for Database Migration Service (DMS).</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/iot</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/iot/CHANGELOG.md#v1280-2022-08-04">v1.28.0</a>
<ul>
<li><strong>Feature</strong>: The release is to support attach a provisioning template to CACert for JITP function,  Customer now doesn't have to hardcode a roleArn and templateBody during register a CACert to enable JITP.</li>
</ul>
</li>
</ul>
<h1>Release (2022-08-03)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/cognitoidentityprovider/CHANGELOG.md#v1180-2022-08-03">v1.18.0</a>
<ul>
<li><strong>Feature</strong>: Add a new exception type, ForbiddenException, that is returned when request is not allowed</li>
</ul>
</li>
<li><code>github.com/aws/aws-sdk-go-v2/service/wafv2</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/wafv2/CHANGELOG.md#v1220-2022-08-03">v1.22.0</a>
<ul>
<li><strong>Feature</strong>: You can now associate an AWS WAF web ACL with an Amazon Cognito user pool.</li>
</ul>
</li>
</ul>
<h1>Release (2022-08-02)</h1>
<h2>Module Highlights</h2>
<ul>
<li><code>github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions</code>: <a href="https://github.com/aws/aws-sdk-go-v2/blob/v1.16.9/service/licensemanagerusersubscriptions/CHANGELOG.md#v100-2022-08-02">v1.0.0</a>
<ul>
<li><strong>Release</strong>: New AWS service client module</li>
</ul>
</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/de69d46aee14e480dcbe52af836fe6ac027868dd"><code>de69d46</code></a> Release 2022-08-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d8c68cb2a00fbc9babfb6be7f26c02fd2c64a37d"><code>d8c68cb</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4d93d564249a327c44554b3d67e0c33dd1d303f8"><code>4d93d56</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/76877dee9ae29011d64d30f48d5ab9cc4de78feb"><code>76877de</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/72cfbc9ceb7b4398ba8497542400853edd6426bc"><code>72cfbc9</code></a> Update CI to include Go1.19 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1785">#1785</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2e91bf1f3c05f4c16e6c4ed1085e7194bbd01488"><code>2e91bf1</code></a> aws/signer/v4: Fix panic handling of endpoint URL without schemes (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1754">#1754</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8d6ca65be706ddf54eeab9d272a61dd31a3f8e25"><code>8d6ca65</code></a> feature/s3/manager: Suppress WriterAt which can bypass Windows optimizations....</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/876fd54a091dcb7fcd7eb058ace62d21869bed89"><code>876fd54</code></a> Batch sync API models from upstream (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/1752">#1752</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a01a5c47baf41b33e2065c28a5c0f2a110ceae61"><code>a01a5c4</code></a> Release 2022-08-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2552ca4934f10f7b67eb10e780b42745407a0194"><code>2552ca4</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.16.4...v1.16.9">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.28.0 to 1.29.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0966539129860e2582d8cecf4e640eff8414b377"><code>0966539</code></a> Release 2022-11-10</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/aec7ab65913d4bc2a4318ff06c814fb5896e6508"><code>aec7ab6</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d008171cc1ff86cdb15e40ac6d1612c41c1b132b"><code>d008171</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/29d44ebdd76da48cb6324e9aebb6246b97c186da"><code>29d44eb</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/01cee3e9ffa2bc6ad2c1646da4e2dff037d169f3"><code>01cee3e</code></a> Release 2022-11-09</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a94734177174d3f45de40328b41affce82cdef73"><code>a947341</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7aa742ab4933522f224e5757dc3254dc7ac16749"><code>7aa742a</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/89b64d9742d7d33aa8b511f995438171c8889acd"><code>89b64d9</code></a> Release 2022-11-08</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/312cdea3a3b1ea9b7c44eb9d8837b884d4790ac3"><code>312cdea</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8ee470880abfd9e97618058bb88c58fd20d992f0"><code>8ee4708</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.28.0...service/s3/v1.29.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2023` from 2023.11.1700250630 to 2023.11.1701177874
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/0a129e63a240552f297f58f85d8a5c765e48f07f"><code>0a129e6</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d6eebd3b2c57952eab83695d...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2023.11.1700250630...release/v2023.11.1701177874">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.24.11 to 1.24.12
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/insights-operator-utils/releases">github.com/RedHatInsights/insights-operator-utils's releases</a>.</em></p>
<blockquote>
<h2>v1.24.12</h2>
<h2>What's Changed</h2>
<ul>
<li>Create dependabot.yml by <a href="https://github.com/tisnik"><code>@​tisnik</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/311">RedHatInsights/insights-operator-utils#311</a></li>
<li>Bump github.com/rs/zerolog from 1.29.1 to 1.31.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/315">RedHatInsights/insights-operator-utils#315</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.35.7 to 1.45.19 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/317">RedHatInsights/insights-operator-utils#317</a></li>
<li>Bump github.com/google/uuid from 1.3.0 to 1.3.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/313">RedHatInsights/insights-operator-utils#313</a></li>
<li>Bump github.com/archdx/zerolog-sentry from 0.0.1 to 1.5.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/316">RedHatInsights/insights-operator-utils#316</a></li>
<li>Bump github.com/prometheus/client_model from 0.4.0 to 0.5.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/318">RedHatInsights/insights-operator-utils#318</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.45.19 to 1.45.24 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/319">RedHatInsights/insights-operator-utils#319</a></li>
<li>Bump github.com/prometheus/client_golang from 1.16.0 to 1.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/320">RedHatInsights/insights-operator-utils#320</a></li>
<li>Bump golang.org/x/net from 0.10.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/321">RedHatInsights/insights-operator-utils#321</a></li>
<li>GRPC dependency to fix HTTP/2 issue by <a href="https://github.com/tisnik"><code>@​tisnik</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/322">RedHatInsights/insights-operator-utils#322</a></li>
<li>Update dependabot.yml by <a href="https://github.com/tisnik"><code>@​tisnik</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/323">RedHatInsights/insights-operator-utils#323</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.45.24 to 1.45.25 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/324">RedHatInsights/insights-operator-utils#324</a></li>
<li>Bump github.com/go-redis/redismock/v9 from 9.0.3 to 9.2.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/326">RedHatInsights/insights-operator-utils#326</a></li>
<li>Bump github.com/archdx/zerolog-sentry from 1.5.0 to 1.6.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/325">RedHatInsights/insights-operator-utils#325</a></li>
<li>Bump github.com/redis/go-redis/v9 from 9.0.5 to 9.2.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/314">RedHatInsights/insights-operator-utils#314</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.45.25 to 1.45.26 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/327">RedHatInsights/insights-operator-utils#327</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.45.26 to 1.45.27 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/328">RedHatInsights/insights-operator-utils#328</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.45.27 to 1.45.28 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/329">RedHatInsights/insights-operator-utils#329</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.45.28 to 1.46.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/330">RedHatInsights/insights-operator-utils#330</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.0 to 1.46.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/331">RedHatInsights/insights-operator-utils#331</a></li>
<li>GH actions by <a href="https://github.com/tisnik"><code>@​tisnik</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/333">RedHatInsights/insights-operator-utils#333</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.1 to 1.46.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/332">RedHatInsights/insights-operator-utils#332</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.2 to 1.46.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/334">RedHatInsights/insights-operator-utils#334</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.3 to 1.46.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/335">RedHatInsights/insights-operator-utils#335</a></li>
<li>Bump github.com/google/uuid from 1.3.1 to 1.4.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/336">RedHatInsights/insights-operator-utils#336</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.4 to 1.46.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/337">RedHatInsights/insights-operator-utils#337</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.5 to 1.46.6 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/338">RedHatInsights/insights-operator-utils#338</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.6 to 1.46.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/339">RedHatInsights/insights-operator-utils#339</a></li>
<li>Bump github.com/redis/go-redis/v9 from 9.2.1 to 9.3.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/340">RedHatInsights/insights-operator-utils#340</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.46.7 to 1.47.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/341">RedHatInsights/insights-operator-utils#341</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.0 to 1.47.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/342">RedHatInsights/insights-operator-utils#342</a></li>
<li>Bump github.com/gorilla/mux from 1.8.0 to 1.8.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/345">RedHatInsights/insights-operator-utils#345</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.1 to 1.47.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/344">RedHatInsights/insights-operator-utils#344</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.3 to 1.47.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/346">RedHatInsights/insights-operator-utils#346</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.4 to 1.47.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/347">RedHatInsights/insights-operator-utils#347</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.5 to 1.47.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/348">RedHatInsights/insights-operator-utils#348</a></li>
<li>Bump github.com/archdx/zerolog-sentry from 1.6.1 to 1.7.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/350">RedHatInsights/insights-operator-utils#350</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.7 to 1.47.9 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/351">RedHatInsights/insights-operator-utils#351</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.9 to 1.47.10 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/352">RedHatInsights/insights-operator-utils#352</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.10 to 1.47.11 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/353">RedHatInsights/insights-operator-utils#353</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.11 to 1.47.12 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/354">RedHatInsights/insights-operator-utils#354</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.47.12 to 1.48.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/356">RedHatInsights/insights-operator-utils#356</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.0 to 1.48.1 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/357">RedHatInsights/insights-operator-utils#357</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.1 to 1.48.2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/358">RedHatInsights/insights-operator-utils#358</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.2 to 1.48.3 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/359">RedHatInsights/insights-operator-utils#359</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.3 to 1.48.4 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/360">RedHatInsights/insights-operator-utils#360</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.4 to 1.48.5 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/362">RedHatInsights/insights-operator-utils#362</a></li>
<li>Bump github.com/aws/aws-sdk-go from 1.48.5 to 1.48.7 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/pull/363">RedHatInsights/insights-operator-utils#363</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/a279b23b9b98417e1c0456f5b9d355bf8ee2add5"><code>a279b23</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/366">#366</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/29dc603e0b0f942be558239d26030f08c54d1c7e"><code>29dc603</code></a> Bump github.com/aws/aws-sdk-go from 1.48.9 to 1.48.10</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/e71652e32fcc7d7888d4f815d9ad523074cef3ef"><code>e71652e</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/365">#365</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/bcbb6f75163f22974515bc033d2608d7cd2dddd1"><code>bcbb6f7</code></a> Bump github.com/aws/aws-sdk-go from 1.48.7 to 1.48.9</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/b9e668ce8cd912cb89144db79e190f7e212e7592"><code>b9e668c</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/363">#363</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/53b1b5bea6a254f90e0854b270c2c7237a4e3591"><code>53b1b5b</code></a> Bump github.com/aws/aws-sdk-go from 1.48.5 to 1.48.7</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/775dabdcf38ef94538ebe108b22c81689be73121"><code>775dabd</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/362">#362</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/d9ab605eab4fffa85b88accb074b408240b1667c"><code>d9ab605</code></a> Bump github.com/aws/aws-sdk-go from 1.48.4 to 1.48.5</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/76f59441362d4d4a68f1305ad028842ecd9a3d10"><code>76f5944</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/360">#360</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/5b2d274188eb7fd15c659a30c6b1309ce6b66360"><code>5b2d274</code></a> Bump github.com/aws/aws-sdk-go from 1.48.3 to 1.48.4</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.24.11...v1.24.12">compare view</a></li>
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

### Comment by @app-sre-bot on December 04, 2023 at 04:43 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on December 04, 2023 at 03:20 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on December 04, 2023 at 04:07 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/494*
