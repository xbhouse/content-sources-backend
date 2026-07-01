---
type: pull_request
number: 731
title: "Build: Bump the go group with 3 updates"
state: closed
author: dependabot
created: 2024-07-08T04:33:58Z
updated: 2024-07-09T13:10:36Z
closed: 2024-07-09T13:10:35Z
base_branch: main
head_branch: dependabot/go_modules/go-d2b4da0cb4
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/731
---

# Pull Request #731: Build: Bump the go group with 3 updates

**Author**: @dependabot
**Created**: July 08, 2024 at 04:33 AM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-d2b4da0cb4`

## Description

Bumps the go group with 3 updates: [github.com/getkin/kin-openapi](https://github.com/getkin/kin-openapi), [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) and [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest).

Updates `github.com/getkin/kin-openapi` from 0.125.0 to 0.126.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getkin/kin-openapi/releases">github.com/getkin/kin-openapi's releases</a>.</em></p>
<blockquote>
<h2>v0.126.0</h2>
<h2>What's Changed</h2>
<ul>
<li>openapi3: document v0.124.0 breaking API changes by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/964">getkin/kin-openapi#964</a></li>
<li>openapi3: introduce <code>ReferencesComponentInRootDocument(doc *T, ref componentRef) (string, bool)</code> by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/945">getkin/kin-openapi#945</a></li>
<li>Update Go module dependencies by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/965">getkin/kin-openapi#965</a></li>
<li>Move paragraph back to its correct section by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/967">getkin/kin-openapi#967</a></li>
<li>Replace interface{} with any by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/966">getkin/kin-openapi#966</a></li>
<li>openapi3: allow Extensions next to $ref in SchemaRef by <a href="https://github.com/paulmach"><code>@​paulmach</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/901">getkin/kin-openapi#901</a></li>
<li>openapi3: implement circular reference backtracking by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/970">getkin/kin-openapi#970</a></li>
<li>openapi3: improve ipv6 validation by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/971">getkin/kin-openapi#971</a></li>
<li>openapi3: resolve recursive file references by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/974">getkin/kin-openapi#974</a></li>
<li>openapi3: improve internalization ref naming to avoid collisions by <a href="https://github.com/percivalalb"><code>@​percivalalb</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/955">getkin/kin-openapi#955</a></li>
<li>openapi3: add a test for additionalProperties: false validation by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/975">getkin/kin-openapi#975</a></li>
<li>openapi3: add support for number and integer format validators by <a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/976">getkin/kin-openapi#976</a></li>
<li>openapi3: allow YAML-marshaling invalid specs by <a href="https://github.com/fenollp"><code>@​fenollp</code></a> in <a href="https://redirect.github.com/getkin/kin-openapi/pull/977">getkin/kin-openapi#977</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/paulmach"><code>@​paulmach</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/901">getkin/kin-openapi#901</a></li>
<li><a href="https://github.com/AnatolyRugalev"><code>@​AnatolyRugalev</code></a> made their first contribution in <a href="https://redirect.github.com/getkin/kin-openapi/pull/970">getkin/kin-openapi#970</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/getkin/kin-openapi/compare/v0.125.0...v0.126.0">https://github.com/getkin/kin-openapi/compare/v0.125.0...v0.126.0</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getkin/kin-openapi/commit/2692f43ba21c89366b2a221a86be520b87539352"><code>2692f43</code></a> openapi3: allow YAML-marshaling invalid specs (<a href="https://redirect.github.com/getkin/kin-openapi/issues/977">#977</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/4144c56505ce18a834a42de224bf9527a5b883fe"><code>4144c56</code></a> openapi3: add support for number and integer format validators (<a href="https://redirect.github.com/getkin/kin-openapi/issues/976">#976</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/4b53bf6e0fe58d9b426fa72fe61a8f3c027e2462"><code>4b53bf6</code></a> openapi3: add a test for additionalProperties: false validation (<a href="https://redirect.github.com/getkin/kin-openapi/issues/975">#975</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/0ed9f5d122299b343138b169ee343d81b09b5a9a"><code>0ed9f5d</code></a> openapi3: improve internalization ref naming to avoid collisions (<a href="https://redirect.github.com/getkin/kin-openapi/issues/955">#955</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/a27c9e791d943bf6c94bb3f750c250aa5cfc133b"><code>a27c9e7</code></a> openapi3: resolve recursive file references (<a href="https://redirect.github.com/getkin/kin-openapi/issues/974">#974</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/42a2d80973ff0d2b1254f327e1ddac9d6cff623c"><code>42a2d80</code></a> openapi3: improve ipv6 validation (<a href="https://redirect.github.com/getkin/kin-openapi/issues/971">#971</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/fe47dca093cf6e92f3dc43f8a500f67427c7bf4a"><code>fe47dca</code></a> openapi3: implement circular reference backtracking (<a href="https://redirect.github.com/getkin/kin-openapi/issues/970">#970</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/57624b323cc511f2c41b6f2721a9df0c6be78501"><code>57624b3</code></a> openapi3: allow Extensions next to $ref in SchemaRef (<a href="https://redirect.github.com/getkin/kin-openapi/issues/901">#901</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/7ec31299a5eef6d8d5b7e7bfcb591c1e053057ea"><code>7ec3129</code></a> Replace interface{} with any (<a href="https://redirect.github.com/getkin/kin-openapi/issues/966">#966</a>)</li>
<li><a href="https://github.com/getkin/kin-openapi/commit/3aa9b4e83f2e210777c69cfd6e2f71bb18484dae"><code>3aa9b4e</code></a> Move paragraph back to its correct section (<a href="https://redirect.github.com/getkin/kin-openapi/issues/967">#967</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/getkin/kin-openapi/compare/v0.125.0...v0.126.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.23 to 1.17.24
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/9005edfbb1194c8f340b9bb7288698b58fc7274a"><code>9005edf</code></a> Release 2024-07-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4d2473c78acb422be1c24a901930a8fb16349694"><code>4d2473c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a04e14633d631a9ebf78d277232ce02b714ca15a"><code>a04e146</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1dfec25b923b5e903d95b4fe34a0a5074a29e9de"><code>1dfec25</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/66f75b6c7fba726f68fbd28790b68ce1e6cbe050"><code>66f75b6</code></a> Release 2024-07-02</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/607ebbf2c1801759227d8ff8560d493962b91d77"><code>607ebbf</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/23079711e886b1d295c9f9dcc9a64a76a1b0ec0f"><code>2307971</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/59bb031cb9935169618cc8237150f2154df079e4"><code>59bb031</code></a> Release 2024-07-01</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/bdced9df4acc57b453ecba3747ec553fc62c9ad6"><code>bdced9d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ee8f0446c7418eed12862575f47b334dc9c9fdc6"><code>ee8f044</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.23...credentials/v1.17.24">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.6.1719602258 to 2024.7.1720031367
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/d1fe98b22b48dfdf732e50f80d53d6fa2e76dc8f"><code>d1fe98b</code></a> Update pulp bindings to 4ead5b9864b4839e4e5d2a6348bd7d4b655d42f57952eab83695d...</li>
<li><a href="https://github.com/content-services/zest/commit/54b07151680a06192b08fc1fdbd7b626c700ff02"><code>54b0715</code></a> Update pulp bindings to e69db48356e528a464be3da896237b855d966bca7d54eb5892a9d...</li>
<li><a href="https://github.com/content-services/zest/commit/36d7bc8790abf56723f705b4634035ce574dc9be"><code>36d7bc8</code></a> Update pulp bindings to b96da963e282496dea8346bd425a7eb3e9b332c87e8639db952b3...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.6.1719602258...release/v2024.7.1720031367">compare view</a></li>
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

### Comment by @app-sre-bot on July 08, 2024 at 04:34 AM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on July 09, 2024 at 01:10 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/731*
