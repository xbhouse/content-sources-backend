---
type: pull_request
number: 1336
title: "chore(deps): bump the lint group across 1 directory with 7 updates"
state: closed
author: dependabot
created: 2025-05-12T18:33:40Z
updated: 2025-09-22T18:02:13Z
closed: 2025-09-22T18:02:12Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-dede607b67
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1336
---

# Pull Request #1336: chore(deps): bump the lint group across 1 directory with 7 updates

**Author**: @dependabot
**Created**: May 12, 2025 at 06:33 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-dede607b67`

## Description

Bumps the lint group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [eslint](https://github.com/eslint/eslint) | `7.32.0` | `9.26.0` |
| [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier) | `7.2.0` | `10.1.5` |
| [eslint-plugin-cypress](https://github.com/cypress-io/eslint-plugin-cypress) | `2.15.2` | `4.3.0` |
| [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) | `7.37.4` | `7.37.5` |
| [stylelint](https://github.com/stylelint/stylelint) | `16.14.1` | `16.19.1` |
| [stylelint-config-recommended-scss](https://github.com/stylelint-scss/stylelint-config-recommended-scss) | `14.1.0` | `15.0.0` |


Updates `eslint` from 7.32.0 to 9.26.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v9.26.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/e9754e7433edf665602ceba4f7f8fbca559c974f"><code>e9754e7</code></a> feat: add reportGlobalThis to no-shadow-restricted-names (<a href="https://redirect.github.com/eslint/eslint/issues/19670">#19670</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/0fa2b7a3666f1eedcc091446dc860037c9bafa5c"><code>0fa2b7a</code></a> feat: add suggestions for <code>eqeqeq</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19640">#19640</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/dcbdcc9c6be628240269b41f7bb576dde1e6f5b3"><code>dcbdcc9</code></a> feat: Add MCP server (<a href="https://redirect.github.com/eslint/eslint/issues/19592">#19592</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/2dfd83ef4ee054f748732581c422508c45d6f1bf"><code>2dfd83e</code></a> feat: add <code>ignoreDirectives</code> option in <code>no-unused-expressions</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19645">#19645</a>) (sethamus)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/96e84de55ad17c96e5b6f2dece75145542505469"><code>96e84de</code></a> fix: check cache file existence before deletion (<a href="https://redirect.github.com/eslint/eslint/issues/19648">#19648</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/d683aebc8e0792e4f80bd1488c705c90f22c317e"><code>d683aeb</code></a> fix: don't crash on tests with circular references in <code>RuleTester</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19664">#19664</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/9736d5d15870c9185da7d140becb9a15aa69057d"><code>9736d5d</code></a> fix: add <code>namespace</code> to <code>Plugin.meta</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19661">#19661</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/17bae69e02fff6f26487a3cbd9c3c3218088949c"><code>17bae69</code></a> fix: update <code>RuleTester.run()</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19634">#19634</a>) (Nitin Kumar)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/dd98d63f09c9324124734206d904d31d433a7c92"><code>dd98d63</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/c25e858d2d7e9bd3e53dcb32c9af5251d6f0569e"><code>c25e858</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/b2397e9bef5ca7faf7e100ecebc20e457bf0b588"><code>b2397e9</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/addd0a6a62d1b89dc7ab49cbd08c5a6af3e7da29"><code>addd0a6</code></a> docs: fix formatting of unordered lists in Markdown (<a href="https://redirect.github.com/eslint/eslint/issues/19660">#19660</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/a21b38db0276ab3373c95ebc7b1ef1910b79dfe6"><code>a21b38d</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/c0721a7f34264da0a32ade8432511eeda4a2c1b9"><code>c0721a7</code></a> docs: fix double space in command (<a href="https://redirect.github.com/eslint/eslint/issues/19657">#19657</a>) (CamWass)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/5b247c859f1b653297a9b9135d92a59742a669cc"><code>5b247c8</code></a> chore: upgrade to <code>@eslint/js@9.26.0</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19681">#19681</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/d6fa4ac031c2fe24fb778e84940393fbda3ddf77"><code>d6fa4ac</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/09586905be394c05839996a5ea812adfac44d320"><code>0958690</code></a> chore: disambiguate internal types <code>LanguageOptions</code> and <code>Rule</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19669">#19669</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/f1c858e3c1e9712ef398588bf5ed68bc19fad3f2"><code>f1c858e</code></a> chore: fix internal type references to <code>Plugin</code> and <code>Rule</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19665">#19665</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/40dd2998cedddb75e0514b2c5cc855293c85da41"><code>40dd299</code></a> refactor: One-shot ESQuery selector analysis (<a href="https://redirect.github.com/eslint/eslint/issues/19652">#19652</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/1cfd7024226cd9c42ceb75732f79e3bc36e8305c"><code>1cfd702</code></a> chore: update dependency <code>@​eslint/json</code> to ^0.12.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19656">#19656</a>) (renovate[bot])</li>
</ul>
<h2>v9.25.1</h2>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/cdc8e8c950ddfe1f9d462ea138ad7866da0394da"><code>cdc8e8c</code></a> fix: revert directive detection in no-unused-expressions (<a href="https://redirect.github.com/eslint/eslint/issues/19639">#19639</a>) (sethamus)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/1f2b057ddcbef4340f78d1314456935054b8d93f"><code>1f2b057</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.25.1 (<a href="https://redirect.github.com/eslint/eslint/issues/19642">#19642</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/771317fa937a07277201f7155e9b835e6a5658f9"><code>771317f</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
</ul>
<h2>v9.25.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/dcd95aafa33a95c8102834af85129f6f398fe394"><code>dcd95aa</code></a> feat: support TypeScript syntax in no-empty-function rule (<a href="https://redirect.github.com/eslint/eslint/issues/19551">#19551</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/77d6d5bc4923012aee34b0a7c3d971f017d65555"><code>77d6d5b</code></a> feat: support TS syntax in <code>no-unused-expressions</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19564">#19564</a>) (Sweta Tanwar)</li>
<li><a href="https://github.com/eslint/eslint/commit/90228e5d57672579cf82bede29880532c2cb8ca9"><code>90228e5</code></a> feat: support <code>JSRuleDefinition</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19604">#19604</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/59ba6b73789835813ab3002c651a7217dd30a8cc"><code>59ba6b7</code></a> feat: add allowObjects option to no-restricted-properties (<a href="https://redirect.github.com/eslint/eslint/issues/19607">#19607</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/db650a036baf502c7366a7da633d4cd00719394e"><code>db650a0</code></a> feat: support TypeScript syntax in <code>no-invalid-this</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19532">#19532</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/9535cffe7b66abe850d90258e702279afabce7fe"><code>9535cff</code></a> feat: support TS syntax in <code>no-loop-func</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19559">#19559</a>) (Nitin Kumar)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/910bd13c4cb49001f2a9f172229360771b857585"><code>910bd13</code></a> fix: <code>nodeTypeKey</code> not being used in <code>NodeEventGenerator</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19631">#19631</a>) (StyleShit)</li>
</ul>
<h2>Documentation</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/blob/main/CHANGELOG.md">eslint's changelog</a>.</em></p>
<blockquote>
<p>v9.26.0 - May 2, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/5b247c859f1b653297a9b9135d92a59742a669cc"><code>5b247c8</code></a> chore: upgrade to <code>@eslint/js@9.26.0</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19681">#19681</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/d6fa4ac031c2fe24fb778e84940393fbda3ddf77"><code>d6fa4ac</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/e9754e7433edf665602ceba4f7f8fbca559c974f"><code>e9754e7</code></a> feat: add reportGlobalThis to no-shadow-restricted-names (<a href="https://redirect.github.com/eslint/eslint/issues/19670">#19670</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/0fa2b7a3666f1eedcc091446dc860037c9bafa5c"><code>0fa2b7a</code></a> feat: add suggestions for <code>eqeqeq</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19640">#19640</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/dd98d63f09c9324124734206d904d31d433a7c92"><code>dd98d63</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/96e84de55ad17c96e5b6f2dece75145542505469"><code>96e84de</code></a> fix: check cache file existence before deletion (<a href="https://redirect.github.com/eslint/eslint/issues/19648">#19648</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/c25e858d2d7e9bd3e53dcb32c9af5251d6f0569e"><code>c25e858</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/09586905be394c05839996a5ea812adfac44d320"><code>0958690</code></a> chore: disambiguate internal types <code>LanguageOptions</code> and <code>Rule</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19669">#19669</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/dcbdcc9c6be628240269b41f7bb576dde1e6f5b3"><code>dcbdcc9</code></a> feat: Add MCP server (<a href="https://redirect.github.com/eslint/eslint/issues/19592">#19592</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/b2397e9bef5ca7faf7e100ecebc20e457bf0b588"><code>b2397e9</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/d683aebc8e0792e4f80bd1488c705c90f22c317e"><code>d683aeb</code></a> fix: don't crash on tests with circular references in <code>RuleTester</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19664">#19664</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/f1c858e3c1e9712ef398588bf5ed68bc19fad3f2"><code>f1c858e</code></a> chore: fix internal type references to <code>Plugin</code> and <code>Rule</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19665">#19665</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/9736d5d15870c9185da7d140becb9a15aa69057d"><code>9736d5d</code></a> fix: add <code>namespace</code> to <code>Plugin.meta</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19661">#19661</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/40dd2998cedddb75e0514b2c5cc855293c85da41"><code>40dd299</code></a> refactor: One-shot ESQuery selector analysis (<a href="https://redirect.github.com/eslint/eslint/issues/19652">#19652</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/addd0a6a62d1b89dc7ab49cbd08c5a6af3e7da29"><code>addd0a6</code></a> docs: fix formatting of unordered lists in Markdown (<a href="https://redirect.github.com/eslint/eslint/issues/19660">#19660</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/a21b38db0276ab3373c95ebc7b1ef1910b79dfe6"><code>a21b38d</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/c0721a7f34264da0a32ade8432511eeda4a2c1b9"><code>c0721a7</code></a> docs: fix double space in command (<a href="https://redirect.github.com/eslint/eslint/issues/19657">#19657</a>) (CamWass)</li>
<li><a href="https://github.com/eslint/eslint/commit/1cfd7024226cd9c42ceb75732f79e3bc36e8305c"><code>1cfd702</code></a> chore: update dependency <code>@​eslint/json</code> to ^0.12.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19656">#19656</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/2dfd83ef4ee054f748732581c422508c45d6f1bf"><code>2dfd83e</code></a> feat: add <code>ignoreDirectives</code> option in <code>no-unused-expressions</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19645">#19645</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/17bae69e02fff6f26487a3cbd9c3c3218088949c"><code>17bae69</code></a> fix: update <code>RuleTester.run()</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19634">#19634</a>) (Nitin Kumar)</li>
</ul>
<p>v9.25.1 - April 21, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/1f2b057ddcbef4340f78d1314456935054b8d93f"><code>1f2b057</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.25.1 (<a href="https://redirect.github.com/eslint/eslint/issues/19642">#19642</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/771317fa937a07277201f7155e9b835e6a5658f9"><code>771317f</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/cdc8e8c950ddfe1f9d462ea138ad7866da0394da"><code>cdc8e8c</code></a> fix: revert directive detection in no-unused-expressions (<a href="https://redirect.github.com/eslint/eslint/issues/19639">#19639</a>) (sethamus)</li>
</ul>
<p>v9.25.0 - April 18, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/88dc1965a4f53babec36e0f5bd450dd02467acde"><code>88dc196</code></a> chore: upgrade <code>@​eslint/js</code><a href="https://github.com/9"><code>@​9</code></a>.25.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19636">#19636</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/345288d7b270e8c122e922bfa31f219aedc4e63b"><code>345288d</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/910bd13c4cb49001f2a9f172229360771b857585"><code>910bd13</code></a> fix: <code>nodeTypeKey</code> not being used in <code>NodeEventGenerator</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19631">#19631</a>) (StyleShit)</li>
<li><a href="https://github.com/eslint/eslint/commit/ca7a735dde44120111d56e36ce93ba750b3c3c86"><code>ca7a735</code></a> docs: update <code>no-undef-init</code> when not to use section (<a href="https://redirect.github.com/eslint/eslint/issues/19624">#19624</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/affe6be0181422a51875a2ad40eb5152d94fc254"><code>affe6be</code></a> chore: upgrade trunk (<a href="https://redirect.github.com/eslint/eslint/issues/19628">#19628</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/1b870c9da4b3aa28f4a6f4f62e0437b743344994"><code>1b870c9</code></a> docs: use <code>eslint-config-xo</code> in the getting started guide (<a href="https://redirect.github.com/eslint/eslint/issues/19629">#19629</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/5d4af16ab170306862dd0c33894044e59e03d041"><code>5d4af16</code></a> docs: add types for multiple rule options (<a href="https://redirect.github.com/eslint/eslint/issues/19616">#19616</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/e8f8d57bd6c0d95f9f25db8c5b3ff72de42488b7"><code>e8f8d57</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/a40348f1f67a6c3da320682d683589f91d7e6f7b"><code>a40348f</code></a> docs: no-use-before-define tweaks (<a href="https://redirect.github.com/eslint/eslint/issues/19622">#19622</a>) (Kirk Waiblinger)</li>
<li><a href="https://github.com/eslint/eslint/commit/0ba3ae3e5a2425560baf771c05e7c69c63a1983c"><code>0ba3ae3</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/865dbfed6cbade8a22756965be256da317801937"><code>865dbfe</code></a> docs: ensure &quot;learn more&quot; deprecation links point to useful resource (<a href="https://redirect.github.com/eslint/eslint/issues/19590">#19590</a>) (Kirk Waiblinger)</li>
<li><a href="https://github.com/eslint/eslint/commit/dcd95aafa33a95c8102834af85129f6f398fe394"><code>dcd95aa</code></a> feat: support TypeScript syntax in no-empty-function rule (<a href="https://redirect.github.com/eslint/eslint/issues/19551">#19551</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/77d6d5bc4923012aee34b0a7c3d971f017d65555"><code>77d6d5b</code></a> feat: support TS syntax in <code>no-unused-expressions</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19564">#19564</a>) (Sweta Tanwar)</li>
<li><a href="https://github.com/eslint/eslint/commit/90228e5d57672579cf82bede29880532c2cb8ca9"><code>90228e5</code></a> feat: support <code>JSRuleDefinition</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19604">#19604</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/f80b746d850021d253c01bb0eae466a701e63055"><code>f80b746</code></a> docs: add known limitations for no-self-compare (<a href="https://redirect.github.com/eslint/eslint/issues/19612">#19612</a>) (Nitin Kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/59ba6b73789835813ab3002c651a7217dd30a8cc"><code>59ba6b7</code></a> feat: add allowObjects option to no-restricted-properties (<a href="https://redirect.github.com/eslint/eslint/issues/19607">#19607</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/db650a036baf502c7366a7da633d4cd00719394e"><code>db650a0</code></a> feat: support TypeScript syntax in <code>no-invalid-this</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19532">#19532</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/dd20cf274e285f09f230638184c997c44912485f"><code>dd20cf2</code></a> test: fix <code>no-loop-func</code> test with duplicate variable reports (<a href="https://redirect.github.com/eslint/eslint/issues/19610">#19610</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/9535cffe7b66abe850d90258e702279afabce7fe"><code>9535cff</code></a> feat: support TS syntax in <code>no-loop-func</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19559">#19559</a>) (Nitin Kumar)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/8bbabc4691d97733a422180c71eba6c097b35475"><code>8bbabc4</code></a> 9.26.0</li>
<li><a href="https://github.com/eslint/eslint/commit/16f5ff799122737e2c4b853b441e86f224878942"><code>16f5ff7</code></a> Build: changelog update for 9.26.0</li>
<li><a href="https://github.com/eslint/eslint/commit/5b247c859f1b653297a9b9135d92a59742a669cc"><code>5b247c8</code></a> chore: upgrade to <code>@eslint/js@9.26.0</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19681">#19681</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/d6fa4ac031c2fe24fb778e84940393fbda3ddf77"><code>d6fa4ac</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/e9754e7433edf665602ceba4f7f8fbca559c974f"><code>e9754e7</code></a> feat: add reportGlobalThis to no-shadow-restricted-names (<a href="https://redirect.github.com/eslint/eslint/issues/19670">#19670</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/0fa2b7a3666f1eedcc091446dc860037c9bafa5c"><code>0fa2b7a</code></a> feat: add suggestions for <code>eqeqeq</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19640">#19640</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/dd98d63f09c9324124734206d904d31d433a7c92"><code>dd98d63</code></a> docs: Update README</li>
<li><a href="https://github.com/eslint/eslint/commit/96e84de55ad17c96e5b6f2dece75145542505469"><code>96e84de</code></a> fix: check cache file existence before deletion (<a href="https://redirect.github.com/eslint/eslint/issues/19648">#19648</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/c25e858d2d7e9bd3e53dcb32c9af5251d6f0569e"><code>c25e858</code></a> docs: Update README</li>
<li><a href="https://github.com/eslint/eslint/commit/09586905be394c05839996a5ea812adfac44d320"><code>0958690</code></a> chore: disambiguate internal types <code>LanguageOptions</code> and <code>Rule</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19669">#19669</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v7.32.0...v9.26.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~eslintbot">eslintbot</a>, a new releaser for eslint since your current version.</p>
</details>
<br />

Updates `eslint-config-prettier` from 7.2.0 to 10.1.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/releases">eslint-config-prettier's releases</a>.</em></p>
<blockquote>
<h2>v10.1.5</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/332">#332</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4"><code>60fef02</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - chore: add <code>funding</code> field into <code>package.json</code></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.4...v10.1.5">https://github.com/prettier/eslint-config-prettier/compare/v10.1.4...v10.1.5</a></p>
<h2>v10.1.4</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/328">#328</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850"><code>94b4799</code></a> Thanks <a href="https://github.com/silvenon"><code>@​silvenon</code></a>! - fix(cli): do not crash on no rules configured</li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.3...v10.1.4">https://github.com/prettier/eslint-config-prettier/compare/v10.1.3...v10.1.4</a></p>
<h2>v10.1.3</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/325">#325</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df"><code>4e95a1d</code></a> Thanks <a href="https://github.com/pilikan"><code>@​pilikan</code></a>! - fix: this package is <code>commonjs</code>, align its types correctly</li>
</ul>
<h3>New Contributors</h3>
<ul>
<li><a href="https://github.com/pilikan"><code>@​pilikan</code></a> made their first contribution in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/325">prettier/eslint-config-prettier#325</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v10.1.2...v10.1.3">https://github.com/prettier/eslint-config-prettier/compare/v10.1.2...v10.1.3</a></p>
<h2>v10.1.2</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> Thanks <a href="https://github.com/Fdawgs"><code>@​Fdawgs</code></a>! - chore(package): add homepage for some 3rd-party registry - see <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> for more details</li>
</ul>
<h2>v10.1.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/309">#309</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - fix: separate the <code>/flat</code> entry for compatibility</p>
<p>For flat config users, the previous <code>&quot;eslint-config-prettier&quot;</code> entry still works, but <code>&quot;eslint-config-prettier/flat&quot;</code> adds a new <code>name</code> property for <a href="https://eslint.org/blog/2024/04/eslint-config-inspector/">config-inspector</a>, we just can't add it for the default entry for compatibility.</p>
<p>See also <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/308">prettier/eslint-config-prettier#308</a></p>
<pre lang="ts"><code>// before
import eslintConfigPrettier from &quot;eslint-config-prettier&quot;;
<p>// after<br />
import eslintConfigPrettier from &quot;eslint-config-prettier/flat&quot;;<br />
</code></pre></p>
</li>
</ul>
<h2>v10.1.0</h2>
<h3>Minor Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/306">#306</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/56e2e3466391d0fdfc200e42130309c687aaab53"><code>56e2e34</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - feat: migrate to exports field</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/blob/main/CHANGELOG.md">eslint-config-prettier's changelog</a>.</em></p>
<blockquote>
<h2>10.1.5</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/332">#332</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4"><code>60fef02</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - chore: add <code>funding</code> field into <code>package.json</code></li>
</ul>
<h2>10.1.4</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/328">#328</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850"><code>94b4799</code></a> Thanks <a href="https://github.com/silvenon"><code>@​silvenon</code></a>! - fix(cli): do not crash on no rules configured</li>
</ul>
<h2>10.1.3</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/325">#325</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df"><code>4e95a1d</code></a> Thanks <a href="https://github.com/pilikan"><code>@​pilikan</code></a>! - fix: this package is <code>commonjs</code>, align its types correctly</li>
</ul>
<h2>10.1.2</h2>
<h3>Patch Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> Thanks <a href="https://github.com/Fdawgs"><code>@​Fdawgs</code></a>! - chore(package): add homepage for some 3rd-party registry - see <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/321">#321</a> for more details</li>
</ul>
<h2>10.1.1</h2>
<h3>Patch Changes</h3>
<ul>
<li>
<p><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/309">#309</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/eb56a5e09964e49045bccde3c616275eb0a0902d"><code>eb56a5e</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - fix: separate the <code>/flat</code> entry for compatibility</p>
<p>For flat config users, the previous <code>&quot;eslint-config-prettier&quot;</code> entry still works, but <code>&quot;eslint-config-prettier/flat&quot;</code> adds a new <code>name</code> property for <a href="https://eslint.org/blog/2024/04/eslint-config-inspector/">config-inspector</a>, we just can't add it for the default entry for compatibility.</p>
<p>See also <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/308">prettier/eslint-config-prettier#308</a></p>
<pre lang="ts"><code>// before
import eslintConfigPrettier from &quot;eslint-config-prettier&quot;;
<p>// after<br />
import eslintConfigPrettier from &quot;eslint-config-prettier/flat&quot;;<br />
</code></pre></p>
</li>
</ul>
<h2>10.1.0</h2>
<h3>Minor Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/306">#306</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/56e2e3466391d0fdfc200e42130309c687aaab53"><code>56e2e34</code></a> Thanks <a href="https://github.com/JounQin"><code>@​JounQin</code></a>! - feat: migrate to exports field</li>
</ul>
<h2>10.0.3</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4c9489339d37bf96d31e0596e64bb8d4cb4308ef"><code>4c94893</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/333">#333</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/60fef02574467d31d10ff47ecb567d378483c9d4"><code>60fef02</code></a> chore: add <code>funding</code> field into <code>package.json</code> (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/332">#332</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/f55501ffe9be65fc9a8ec7d788459fd3a9cb6095"><code>f55501f</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/329">#329</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/50a8a22b0468e3469b7a177e6c81e843bd5cb73e"><code>50a8a22</code></a> chore(deps): update all dependencies (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/330">#330</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/94b47999e7eb13b703835729331376cef598b850"><code>94b4799</code></a> fix(cli): do not crash on no rules configured (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/328">#328</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/cdc4a5c7e39e7f2d5760c60ea39cecb028fb34dc"><code>cdc4a5c</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/326">#326</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4e95a1d50073f1a24f004239ad6e1a4ffa8476df"><code>4e95a1d</code></a> fix: this package is <code>commonjs</code>, align its types correctly (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/325">#325</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/8911369cbc66f1f859e19751eaefdea687129de5"><code>8911369</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/322">#322</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/a8768bfe54a91d08f0cef8705f91de2883436bb0"><code>a8768bf</code></a> chore(package): add homepage url (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/321">#321</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4ae04c0dea72dd7d950bb575a8d87a90ab5ea787"><code>4ae04c0</code></a> chore(deps): update yarn to v4.8.1 (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/320">#320</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prettier/eslint-config-prettier/compare/v7.2.0...v10.1.5">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~jounqin">jounqin</a>, a new releaser for eslint-config-prettier since your current version.</p>
</details>
<br />

Updates `eslint-plugin-cypress` from 2.15.2 to 4.3.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/cypress-io/eslint-plugin-cypress/releases">eslint-plugin-cypress's releases</a>.</em></p>
<blockquote>
<h2>v4.3.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.2.1...v4.3.0">4.3.0</a> (2025-04-22)</h1>
<h3>Features</h3>
<ul>
<li>add no-chained-get rule (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/249">#249</a>) (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/2c911ac76dad352fba1cd5e24fb6decb63118072">2c911ac</a>)</li>
</ul>
<h2>v4.2.1</h2>
<h2><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.2.0...v4.2.1">4.2.1</a> (2025-04-09)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>deps:</strong> update all dependencies (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/5b827eb700b2dabc1946722f7a41b7551d4dc439">5b827eb</a>)</li>
</ul>
<h2>v4.2.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.1.0...v4.2.0">4.2.0</a> (2025-03-07)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>address comments in PR (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/64eaba0a572e15fd7025d6b426c38887c3ceb53c">64eaba0</a>)</li>
<li>doc title (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1fbedaca91e35715648091a3387f46b67bb56157">1fbedac</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>add rule disallow usage of cypress-xpath (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/1ae902a3907c984820fbda2010e8c078d00fe503">1ae902a</a>)</li>
</ul>
<h2>v4.1.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v4.0.0...v4.1.0">4.1.0</a> (2024-10-30)</h1>
<h3>Features</h3>
<ul>
<li><strong>docs:</strong> publish updated readme (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/2bc8d5a82208f7da463f250573d493e6e6c287c4">2bc8d5a</a>)</li>
</ul>
<h2>v4.0.0</h2>
<h1><a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v3.6.0...v4.0.0">4.0.0</a> (2024-10-11)</h1>
<h3>Features</h3>
<ul>
<li>minimum version eslint v9 (<a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/b50181ad06d423dd61e611a0e052a074758bfa8f">b50181a</a>)</li>
</ul>
<h3>BREAKING CHANGES</h3>
<ul>
<li>Support ESLint v9 and above only</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/2c911ac76dad352fba1cd5e24fb6decb63118072"><code>2c911ac</code></a> feat: add no-chained-get rule (<a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/249">#249</a>)</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/4f062acfaa2fbecf3092c4d3f68e904a77aa5e36"><code>4f062ac</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/254">#254</a> from MikeMcC399/docs/merge-guidelines</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/84e03d0b290377c12e068b8d9ec1c477c743588b"><code>84e03d0</code></a> docs: add PR merge guidelines for maintainers</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/de3ef37b0265a9f489e4c6fa2511115995a34a1e"><code>de3ef37</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/253">#253</a> from MikeMcC399/update/deps</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/5b827eb700b2dabc1946722f7a41b7551d4dc439"><code>5b827eb</code></a> fix(deps): update all dependencies</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/83af80b0bfdbd3e81dde6b7b8970af0f5c01dae4"><code>83af80b</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/252">#252</a> from MikeMcC399/add/project-test</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/9060087dca4342fe66d0cdbae93340fd8d8266b9"><code>9060087</code></a> Update test-project/README.md</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/75607fdae2949adf6f2c19d55c20ce6f8ad6afb1"><code>75607fd</code></a> ci: add test project</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/cd4b9affe9e67b86f18ed8ca0e46bf56cb15a7ea"><code>cd4b9af</code></a> Merge pull request <a href="https://redirect.github.com/cypress-io/eslint-plugin-cypress/issues/251">#251</a> from MikeMcC399/update/cimg</li>
<li><a href="https://github.com/cypress-io/eslint-plugin-cypress/commit/11733ecb8f7b0d3eef0d7e72bd0d2acd0d026399"><code>11733ec</code></a> ci(deps): update to cimg/node:22.14.0</li>
<li>Additional commits viewable in <a href="https://github.com/cypress-io/eslint-plugin-cypress/compare/v2.15.2...v4.3.0">compare view</a></li>
</ul>
</details>
<br />

Updates `eslint-plugin-react` from 7.37.4 to 7.37.5
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jsx-eslint/eslint-plugin-react/releases">eslint-plugin-react's releases</a>.</em></p>
<blockquote>
<h2>v7.37.5</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-unknown-property</code>]: allow shadow root attrs on <code>\&lt;template&gt;</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
<li>[<code>prop-types</code>]: support <code>ComponentPropsWithRef</code> from a namespace import (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>[] <a href="https://github.com/corydeppen"><code>@​corydeppen</code></a>)</li>
<li>[<code>jsx-no-constructed-context-values</code>]: detect constructed context values in React 19 <code>&lt;Context&gt;</code> usage (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>[] <a href="https://github.com/TildaDares"><code>@​TildaDares</code></a>)</li>
<li>[<code>no-unknown-property</code>]: allow <code>transform-origin</code> on <code>rect</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[Docs] [<code>button-has-type</code>]: clean up phrasing (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>[] <a href="https://github.com/hamirmahal"><code>@​hamirmahal</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3651">jsx-eslint/eslint-plugin-react#3651</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3909">jsx-eslint/eslint-plugin-react#3909</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3910">jsx-eslint/eslint-plugin-react#3910</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">jsx-eslint/eslint-plugin-react#3912</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">jsx-eslint/eslint-plugin-react#3914</a>
[<code>button-has-type</code>]: docs/rules/button-has-type.md
[<code>jsx-no-constructed-context-values</code>]: docs/rules/jsx-no-constructed-context-values.md
[<code>no-unknown-property</code>]: docs/rules/no-unknown-property.md
[<code>prop-types</code>]: docs/rules/prop-types.md</p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jsx-eslint/eslint-plugin-react/blob/master/CHANGELOG.md">eslint-plugin-react's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.4...v7.37.5">7.37.5</a> - 2025.04.03</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-unknown-property</code>]: allow shadow root attrs on <code>\&lt;template&gt;</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
<li>[<code>prop-types</code>]: support <code>ComponentPropsWithRef</code> from a namespace import (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>[] <a href="https://github.com/corydeppen"><code>@​corydeppen</code></a>)</li>
<li>[<code>jsx-no-constructed-context-values</code>]: detect constructed context values in React 19 <code>&lt;Context&gt;</code> usage (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>[] <a href="https://github.com/TildaDares"><code>@​TildaDares</code></a>)</li>
<li>[<code>no-unknown-property</code>]: allow <code>transform-origin</code> on <code>rect</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[Docs] [<code>button-has-type</code>]: clean up phrasing (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>[] <a href="https://github.com/hamirmahal"><code>@​hamirmahal</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">#3914</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3914">jsx-eslint/eslint-plugin-react#3914</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">#3912</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3912">jsx-eslint/eslint-plugin-react#3912</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3910">#3910</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3910">jsx-eslint/eslint-plugin-react#3910</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3909">#3909</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3909">jsx-eslint/eslint-plugin-react#3909</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3651">#3651</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3651">jsx-eslint/eslint-plugin-react#3651</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/2c98b83c451a4297edf1787d9a616e50687e27e8"><code>2c98b83</code></a> Update CHANGELOG and bump version</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/2f64deadac51b42fc1a8660fad026ac4c68b92f3"><code>2f64dea</code></a> [Fix] <code>no-unknown-property</code>: allow <code>transform-origin</code> on <code>rect</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/2428618b5a7334b96b7b7eb9629212d07b6fd510"><code>2428618</code></a> [Fix] <code>jsx-no-constructed-context-values</code>: detect constructed context values ...</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/60b731621c98b8d3f6c8c5339a50dc54bf3fd068"><code>60b7316</code></a> [Tests] <code>prop-types</code>: use proper spacing/semis, button type</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/90a00b9318374b402114a4136c6f118b48d9346e"><code>90a00b9</code></a> [Fix] <code>prop-types</code>: support <code>ComponentPropsWithRef</code> from a namespace import</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/3fd9b9223e3f4fc6b34eb6f3ab734a7e2c73743d"><code>3fd9b92</code></a> [Fix] <code>no-unknown-property</code>: allow shadow root attrs on <code>\&lt;template&gt;</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/addad4687b710c022f868ea17f6cabfaaddd8b44"><code>addad46</code></a> [Deps] update <code>object.entries</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/98a31f8e76a4d8aa52caeeb55940f35682b18b2f"><code>98a31f8</code></a> [Dev Deps] update <code>@babel/core</code>, <code>@babel/eslint-parser</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/7eb6ca9144333c828f24abdc98154a45aec46d54"><code>7eb6ca9</code></a> [Docs] <code>button-has-type</code>: clean up phrasing</li>
<li>See full diff in <a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.4...v7.37.5">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 16.14.1 to 16.19.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>16.19.1</h2>
<ul>
<li>Fixed: <code>no-empty-source</code> false positives for non-standard syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8548">#8548</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
</ul>
<h2>16.19.0</h2>
<p>It adds 2 options to 2 rules and fixes 3 bugs.</p>
<ul>
<li>Added: <code>exceptWithoutPropertyFallback: []</code> to <code>function-allowed-list</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8488">#8488</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignore: [&quot;four-into-three-edge-values&quot;]</code> to <code>shorthand-property-no-redundant-values</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8527">#8527</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>compact</code> formatter with pnpm to newline the exit code (<a href="https://redirect.github.com/stylelint/stylelint/issues/8534">#8534</a>) (<a href="https://github.com/konomae"><code>@​konomae</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> range and message for invalid syntax within known functions (<a href="https://redirect.github.com/stylelint/stylelint/issues/8528">#8528</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>no-empty-source</code> false positives for <code>--report-needless-disables</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8536">#8536</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
</ul>
<h2>16.18.0</h2>
<p>It adds 2 new rules and fixes 2 bugs. We've turned on these rules, and the <code>syntax-string-no-invalid</code> and <code>layer-name-pattern</code> ones from recent releases, in our <a href="https://www.npmjs.com/package/stylelint-config-standard">standard config</a>.</p>
<ul>
<li>Added: <code>color-function-alias-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8499">#8499</a>) (<a href="https://github.com/EduardAkhmetshin"><code>@​EduardAkhmetshin</code></a>).</li>
<li>Added: <code>container-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8498">#8498</a>) (<a href="https://github.com/nate10j"><code>@​nate10j</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>math</code> of <code>font-size</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8495">#8495</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
<li>Fixed: <code>font-family-no-missing-generic-family-keyword</code> false positives for <code>math</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8489">#8489</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
</ul>
<h2>16.17.0</h2>
<p>It adds 1 new rule, support for <code>languageOptions</code> to 2 rules, 1 option to a rule, the <code>--compute-edit-info</code> CLI flag (along with support for <code>EditInfo</code> in 3 rules), and fixes 1 bug. <code>EditInfo</code> is useful for automated fixing tools and editor integrations.</p>
<ul>
<li>Added: <code>layer-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8474">#8474</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>--compute-edit-info</code> CLI flag (<a href="https://redirect.github.com/stylelint/stylelint/issues/8473">#8473</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignorePreludeOfAtRules: []</code> to <code>length-zero-no-unit</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8472">#8472</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8475">#8475</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>property-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8476">#8476</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>declaration-block-no-redundant-longhand-properties</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8482">#8482</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>function-url-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8483">#8483</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-attribute-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8484">#8484</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Fixed: <code>custom-property-pattern</code> false negatives for <code>@property</code> preludes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8468">#8468</a>) (<a href="https://github.com/rohitgs28"><code>@​rohitgs28</code></a>).</li>
</ul>
<h2>16.16.0</h2>
<p>It adds support for computing <code>EditInfo</code> to 22 more rules and reverts a change that added <code>context.lexer</code> to our public API in the previous release.</p>
<ul>
<li>Added: <code>at-rule-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8425">#8425</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-deprecated</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8426">#8426</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8427">#8427</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>color-function-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8437">#8437</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8443">#8443</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-property-value-keyword-no-deprecated</code> support for computing <code>EditInfo</code>. (<a href="https://redirect.github.com/stylelint/stylelint/issues/8439">#8439</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>font-family-name-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8419">#8419</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>font-weight-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8420">#8420</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>function-calc-no-unspaced-operator</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8440">#8440</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>function-name-case</code> support for support for computing <code>EditInfo</code>.&quot; (<a href="https://redirect.github.com/stylelint/stylelint/issues/8442">#8442</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>hue-degree-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8444">#8444</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>import-notation</code> support for computing <code>EditInfo</code>. (<a href="https://redirect.github.com/stylelint/stylelint/issues/8445">#8445</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>keyframe-selector-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8457">#8457</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>length-zero-no-unit</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8459">#8459</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/blob/main/CHANGELOG.md">stylelint's changelog</a>.</em></p>
<blockquote>
<h2>16.19.1 - 2025-04-25</h2>
<ul>
<li>Fixed: <code>no-empty-source</code> false positives for non-standard syntaxes (<a href="https://redirect.github.com/stylelint/stylelint/pull/8548">#8548</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
</ul>
<h2>16.19.0 - 2025-04-23</h2>
<p>It adds 2 options to 2 rules and fixes 3 bugs.</p>
<ul>
<li>Added: <code>exceptWithoutPropertyFallback: []</code> to <code>function-allowed-list</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8488">#8488</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignore: [&quot;four-into-three-edge-values&quot;]</code> to <code>shorthand-property-no-redundant-values</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8527">#8527</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>compact</code> formatter with pnpm to newline the exit code (<a href="https://redirect.github.com/stylelint/stylelint/pull/8534">#8534</a>) (<a href="https://github.com/konomae"><code>@​konomae</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> range and message for invalid syntax within known functions (<a href="https://redirect.github.com/stylelint/stylelint/pull/8528">#8528</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Fixed: <code>no-empty-source</code> false positives for <code>--report-needless-disables</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8536">#8536</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
</ul>
<h2>16.18.0 - 2025-04-06</h2>
<p>It adds 2 new rules and fixes 2 bugs. We've turned on these rules, and the <code>syntax-string-no-invalid</code> and <code>layer-name-pattern</code> ones from recent releases, in our <a href="https://www.npmjs.com/package/stylelint-config-standard">standard config</a>.</p>
<ul>
<li>Added: <code>color-function-alias-notation</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8499">#8499</a>) (<a href="https://github.com/EduardAkhmetshin"><code>@​EduardAkhmetshin</code></a>).</li>
<li>Added: <code>container-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8498">#8498</a>) (<a href="https://github.com/nate10j"><code>@​nate10j</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for <code>math</code> of <code>font-size</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8495">#8495</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
<li>Fixed: <code>font-family-no-missing-generic-family-keyword</code> false positives for <code>math</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8489">#8489</a>) (<a href="https://github.com/otomad"><code>@​otomad</code></a>).</li>
</ul>
<h2>16.17.0 - 2025-03-26</h2>
<p>It adds 1 new rule, support for <code>languageOptions</code> to 2 rules, 1 option to a rule, the <code>--compute-edit-info</code> CLI flag (along with support for <code>EditInfo</code> in 3 rules), and fixes 1 bug. <code>EditInfo</code> is useful for automated fixing tools and editor integrations.</p>
<ul>
<li>Added: <code>layer-name-pattern</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/pull/8474">#8474</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>--compute-edit-info</code> CLI flag (<a href="https://redirect.github.com/stylelint/stylelint/pull/8473">#8473</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>ignorePreludeOfAtRules: []</code> to <code>length-zero-no-unit</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8472">#8472</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8475">#8475</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>property-no-unknown</code> support for <code>languageOptions</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8476">#8476</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>declaration-block-no-redundant-longhand-properties</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8482">#8482</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>function-url-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8483">#8483</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>selector-attribute-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8484">#8484</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Fixed: <code>custom-property-pattern</code> false negatives for <code>@property</code> preludes (<a href="https://redirect.github.com/stylelint/stylelint/pull/8468">#8468</a>) (<a href="https://github.com/rohitgs28"><code>@​rohitgs28</code></a>).</li>
</ul>
<h2>16.16.0 - 2025-03-14</h2>
<p>It adds support for computing <code>EditInfo</code> to 22 more rules and reverts a change that added <code>context.lexer</code> to our public API in the previous release.</p>
<ul>
<li>Added: <code>at-rule-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8425">#8425</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-deprecated</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8426">#8426</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>at-rule-no-vendor-prefix</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8427">#8427</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>color-function-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8437">#8437</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-empty-line-before</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8443">#8443</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>declaration-property-value-keyword-no-deprecated</code> support for computing <code>EditInfo</code>. (<a href="https://redirect.github.com/stylelint/stylelint/pull/8439">#8439</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
<li>Added: <code>font-family-name-quotes</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8419">#8419</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>font-weight-notation</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8420">#8420</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>function-calc-no-unspaced-operator</code> support for computing <code>EditInfo</code> (<a href="https://redirect.github.com/stylelint/stylelint/pull/8440">#8440</a>) (<a href="https://github.com/pamelalozano16"><code>@​pamelalozano16</code></a>).</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stylelint/stylelint/commit/25968c6066d0a747b6522252c7ea8c35b36883ea"><code>25968c6</code></a> 16.19.1</li>
<li><a href="https://github.com/stylelint/stylelint/commit/015042b5d20e51818ccf17acfe342db6afb94fd5"><code>015042b</code></a> Prepare 16.19.1 (<a href="https://redirect.github.com/stylelint/stylelint/issues/8551">#8551</a>)</li>
<li><a href="https://github.com/stylelint/stylelint/commit/d4c45d33a4888b0e9c6db87d0091ea00fd4b017d"><code>d4c45d3</code></a> Refactor to allow node for <code>isConfigurationComment()</code> private utility (<a href="https://red...

_Description has been truncated_

> **Note**
> Automatic rebases have been disabled on this pull request as it has been open for over 30 days.

---

## Discussion

### Comment by @dependabot on August 11, 2025 at 07:19 PM UTC

Dependabot tried to update this pull request, but something went wrong. We're looking into it, but in the meantime you can retry the update by commenting `@dependabot rebase`.

### Comment by @dependabot on September 15, 2025 at 06:11 PM UTC

Dependabot tried to update this pull request, but something went wrong. We're looking into it, but in the meantime you can retry the update by commenting `@dependabot rebase`.

### Comment by @dependabot on September 22, 2025 at 06:02 PM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1336*
