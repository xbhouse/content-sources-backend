---
type: pull_request
number: 1345
title: "chore(deps): bump axios from 1.7.9 to 1.10.0"
state: closed
author: dependabot
created: 2025-06-23T20:03:18Z
updated: 2025-08-12T01:52:07Z
closed: 2025-08-12T01:52:06Z
base_branch: master
head_branch: dependabot/npm_and_yarn/axios-1.10.0
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1345
---

# Pull Request #1345: chore(deps): bump axios from 1.7.9 to 1.10.0

**Author**: @dependabot
**Created**: June 23, 2025 at 08:03 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/axios-1.10.0`

## Description

Bumps [axios](https://github.com/axios/axios) from 1.7.9 to 1.10.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
<h2>Release v1.10.0</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>adapter:</strong> pass fetchOptions to fetch function (<a href="https://redirect.github.com/axios/axios/issues/6883">#6883</a>) (<a href="https://github.com/axios/axios/commit/0f50af8e076b7fb403844789bd5e812dedcaf4ed">0f50af8</a>)</li>
<li><strong>form-data:</strong> convert boolean values to strings in FormData serialization (<a href="https://redirect.github.com/axios/axios/issues/6917">#6917</a>) (<a href="https://github.com/axios/axios/commit/5064b108de336ff34862650709761b8a96d26be0">5064b10</a>)</li>
<li><strong>package:</strong> add module entry point for React Native; (<a href="https://redirect.github.com/axios/axios/issues/6933">#6933</a>) (<a href="https://github.com/axios/axios/commit/3d343b86dc4fd0eea0987059c5af04327c7ae304">3d343b8</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>types:</strong> improved fetchOptions interface (<a href="https://redirect.github.com/axios/axios/issues/6867">#6867</a>) (<a href="https://github.com/axios/axios/commit/63f1fce233009f5db1abf2586c145825ac98c3d7">63f1fce</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+30/-19 ([#6933](https://github.com/axios/axios/issues/6933) [#6920](https://github.com/axios/axios/issues/6920) [#6893](https://github.com/axios/axios/issues/6893) [#6892](https://github.com/axios/axios/issues/6892) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/noritaka1166" title="+2/-6 ([#6922](https://github.com/axios/axios/issues/6922) [#6923](https://github.com/axios/axios/issues/6923) )">Noritaka Kobayashi</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/dimitry-lzs" title="+4/-0 ([#6917](https://github.com/axios/axios/issues/6917) )">Dimitrios Lazanas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/AdrianKnapp" title="+2/-2 ([#6867](https://github.com/axios/axios/issues/6867) )">Adrian Knapp</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/howiezhao" title="+3/-1 ([#6872](https://github.com/axios/axios/issues/6872) )">Howie Zhao</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/warpdev" title="+1/-1 ([#6883](https://github.com/axios/axios/issues/6883) )">Uhyeon Park</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/stscoundrel" title="+1/-1 ([#6913](https://github.com/axios/axios/issues/6913) )">Sampo Silvennoinen</a></li>
</ul>
<h2>Release v1.9.0</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>core:</strong> fix the Axios constructor implementation to treat the config argument as optional; (<a href="https://redirect.github.com/axios/axios/issues/6881">#6881</a>) (<a href="https://github.com/axios/axios/commit/6c5d4cd69286868059c5e52d45085cb9a894a983">6c5d4cd</a>)</li>
<li><strong>fetch:</strong> fixed ERR_NETWORK mapping for Safari browsers; (<a href="https://redirect.github.com/axios/axios/issues/6767">#6767</a>) (<a href="https://github.com/axios/axios/commit/dfe8411c9a082c3d068bdd1f8d6e73054f387f45">dfe8411</a>)</li>
<li><strong>headers:</strong> allow iterable objects to be a data source for the set method; (<a href="https://redirect.github.com/axios/axios/issues/6873">#6873</a>) (<a href="https://github.com/axios/axios/commit/1b1f9ccdc15f1ea745160ec9a5223de9db4673bc">1b1f9cc</a>)</li>
<li><strong>headers:</strong> fix <code>getSetCookie</code> by using 'get' method for caseless access; (<a href="https://redirect.github.com/axios/axios/issues/6874">#6874</a>) (<a href="https://github.com/axios/axios/commit/d4f7df4b304af8b373488fdf8e830793ff843eb9">d4f7df4</a>)</li>
<li><strong>headers:</strong> fixed support for setting multiple header values from an iterated source; (<a href="https://redirect.github.com/axios/axios/issues/6885">#6885</a>) (<a href="https://github.com/axios/axios/commit/f7a3b5e0f7e5e127b97defa92a132fbf1b55cf15">f7a3b5e</a>)</li>
<li><strong>http:</strong> send minimal end multipart boundary (<a href="https://redirect.github.com/axios/axios/issues/6661">#6661</a>) (<a href="https://github.com/axios/axios/commit/987d2e2dd3b362757550f36eab875e60640b6ddc">987d2e2</a>)</li>
<li><strong>types:</strong> fix autocomplete for adapter config (<a href="https://redirect.github.com/axios/axios/issues/6855">#6855</a>) (<a href="https://github.com/axios/axios/commit/e61a8934d8f94dd429a2f309b48c67307c700df0">e61a893</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>AxiosHeaders:</strong> add getSetCookie method to retrieve set-cookie headers values (<a href="https://redirect.github.com/axios/axios/issues/5707">#5707</a>) (<a href="https://github.com/axios/axios/commit/80ea756e72bcf53110fa792f5d7ab76e8b11c996">80ea756</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+200/-34 ([#6890](https://github.com/axios/axios/issues/6890) [#6889](https://github.com/axios/axios/issues/6889) [#6888](https://github.com/axios/axios/issues/6888) [#6885](https://github.com/axios/axios/issues/6885) [#6881](https://github.com/axios/axios/issues/6881) [#6767](https://github.com/axios/axios/issues/6767) [#6874](https://github.com/axios/axios/issues/6874) [#6873](https://github.com/axios/axios/issues/6873) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/jasonsaayman" title="+26/-1 ()">Jay</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/WillianAgostini" title="+21/-0 ([#5707](https://github.com/axios/axios/issues/5707) )">Willian Agostini</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/Gerhut" title="+3/-3 ([#5096](https://github.com/axios/axios/issues/5096) )">George Cheng</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/FatahChan" title="+2/-2 ([#6855](https://github.com/axios/axios/issues/6855) )">FatahChan</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/igstan" title="+1/-1 ([#6661](https://github.com/axios/axios/issues/6661) )">Ionuț G. Stan</a></li>
</ul>
<h2>Release v1.8.4</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v1.x/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
<h1><a href="https://github.com/axios/axios/compare/v1.9.0...v1.10.0">1.10.0</a> (2025-06-14)</h1>
<h3>Bug Fixes</h3>
<ul>
<li><strong>adapter:</strong> pass fetchOptions to fetch function (<a href="https://redirect.github.com/axios/axios/issues/6883">#6883</a>) (<a href="https://github.com/axios/axios/commit/0f50af8e076b7fb403844789bd5e812dedcaf4ed">0f50af8</a>)</li>
<li><strong>form-data:</strong> convert boolean values to strings in FormData serialization (<a href="https://redirect.github.com/axios/axios/issues/6917">#6917</a>) (<a href="https://github.com/axios/axios/commit/5064b108de336ff34862650709761b8a96d26be0">5064b10</a>)</li>
<li><strong>package:</strong> add module entry point for React Native; (<a href="https://redirect.github.com/axios/axios/issues/6933">#6933</a>) (<a href="https://github.com/axios/axios/commit/3d343b86dc4fd0eea0987059c5af04327c7ae304">3d343b8</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>types:</strong> improved fetchOptions interface (<a href="https://redirect.github.com/axios/axios/issues/6867">#6867</a>) (<a href="https://github.com/axios/axios/commit/63f1fce233009f5db1abf2586c145825ac98c3d7">63f1fce</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+30/-19 ([#6933](https://github.com/axios/axios/issues/6933) [#6920](https://github.com/axios/axios/issues/6920) [#6893](https://github.com/axios/axios/issues/6893) [#6892](https://github.com/axios/axios/issues/6892) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/noritaka1166" title="+2/-6 ([#6922](https://github.com/axios/axios/issues/6922) [#6923](https://github.com/axios/axios/issues/6923) )">Noritaka Kobayashi</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/dimitry-lzs" title="+4/-0 ([#6917](https://github.com/axios/axios/issues/6917) )">Dimitrios Lazanas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/AdrianKnapp" title="+2/-2 ([#6867](https://github.com/axios/axios/issues/6867) )">Adrian Knapp</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/howiezhao" title="+3/-1 ([#6872](https://github.com/axios/axios/issues/6872) )">Howie Zhao</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/warpdev" title="+1/-1 ([#6883](https://github.com/axios/axios/issues/6883) )">Uhyeon Park</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/stscoundrel" title="+1/-1 ([#6913](https://github.com/axios/axios/issues/6913) )">Sampo Silvennoinen</a></li>
</ul>
<h1><a href="https://github.com/axios/axios/compare/v1.8.4...v1.9.0">1.9.0</a> (2025-04-24)</h1>
<h3>Bug Fixes</h3>
<ul>
<li><strong>core:</strong> fix the Axios constructor implementation to treat the config argument as optional; (<a href="https://redirect.github.com/axios/axios/issues/6881">#6881</a>) (<a href="https://github.com/axios/axios/commit/6c5d4cd69286868059c5e52d45085cb9a894a983">6c5d4cd</a>)</li>
<li><strong>fetch:</strong> fixed ERR_NETWORK mapping for Safari browsers; (<a href="https://redirect.github.com/axios/axios/issues/6767">#6767</a>) (<a href="https://github.com/axios/axios/commit/dfe8411c9a082c3d068bdd1f8d6e73054f387f45">dfe8411</a>)</li>
<li><strong>headers:</strong> allow iterable objects to be a data source for the set method; (<a href="https://redirect.github.com/axios/axios/issues/6873">#6873</a>) (<a href="https://github.com/axios/axios/commit/1b1f9ccdc15f1ea745160ec9a5223de9db4673bc">1b1f9cc</a>)</li>
<li><strong>headers:</strong> fix <code>getSetCookie</code> by using 'get' method for caseless access; (<a href="https://redirect.github.com/axios/axios/issues/6874">#6874</a>) (<a href="https://github.com/axios/axios/commit/d4f7df4b304af8b373488fdf8e830793ff843eb9">d4f7df4</a>)</li>
<li><strong>headers:</strong> fixed support for setting multiple header values from an iterated source; (<a href="https://redirect.github.com/axios/axios/issues/6885">#6885</a>) (<a href="https://github.com/axios/axios/commit/f7a3b5e0f7e5e127b97defa92a132fbf1b55cf15">f7a3b5e</a>)</li>
<li><strong>http:</strong> send minimal end multipart boundary (<a href="https://redirect.github.com/axios/axios/issues/6661">#6661</a>) (<a href="https://github.com/axios/axios/commit/987d2e2dd3b362757550f36eab875e60640b6ddc">987d2e2</a>)</li>
<li><strong>types:</strong> fix autocomplete for adapter config (<a href="https://redirect.github.com/axios/axios/issues/6855">#6855</a>) (<a href="https://github.com/axios/axios/commit/e61a8934d8f94dd429a2f309b48c67307c700df0">e61a893</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li><strong>AxiosHeaders:</strong> add getSetCookie method to retrieve set-cookie headers values (<a href="https://redirect.github.com/axios/axios/issues/5707">#5707</a>) (<a href="https://github.com/axios/axios/commit/80ea756e72bcf53110fa792f5d7ab76e8b11c996">80ea756</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+200/-34 ([#6890](https://github.com/axios/axios/issues/6890) [#6889](https://github.com/axios/axios/issues/6889) [#6888](https://github.com/axios/axios/issues/6888) [#6885](https://github.com/axios/axios/issues/6885) [#6881](https://github.com/axios/axios/issues/6881) [#6767](https://github.com/axios/axios/issues/6767) [#6874](https://github.com/axios/axios/issues/6874) [#6873](https://github.com/axios/axios/issues/6873) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/jasonsaayman" title="+26/-1 ()">Jay</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/WillianAgostini" title="+21/-0 ([#5707](https://github.com/axios/axios/issues/5707) )">Willian Agostini</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/Gerhut" title="+3/-3 ([#5096](https://github.com/axios/axios/issues/5096) )">George Cheng</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/FatahChan" title="+2/-2 ([#6855](https://github.com/axios/axios/issues/6855) )">FatahChan</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/igstan" title="+1/-1 ([#6661](https://github.com/axios/axios/issues/6661) )">Ionuț G. Stan</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/73a836dae75f06055c24561d83cf4ca1c43e2854"><code>73a836d</code></a> chore(release): v1.10.0 (<a href="https://redirect.github.com/axios/axios/issues/6943">#6943</a>)</li>
<li><a href="https://github.com/axios/axios/commit/3d343b86dc4fd0eea0987059c5af04327c7ae304"><code>3d343b8</code></a> fix(package): add module entry point for React Native; (<a href="https://redirect.github.com/axios/axios/issues/6933">#6933</a>)</li>
<li><a href="https://github.com/axios/axios/commit/0f50af8e076b7fb403844789bd5e812dedcaf4ed"><code>0f50af8</code></a> fix(adapter): pass fetchOptions to fetch function (<a href="https://redirect.github.com/axios/axios/issues/6883">#6883</a>)</li>
<li><a href="https://github.com/axios/axios/commit/ee7799e13c0783c0fdfa656613bb1af6f5e53ccd"><code>ee7799e</code></a> refactor: remove unused import in test (<a href="https://redirect.github.com/axios/axios/issues/6922">#6922</a>)</li>
<li><a href="https://github.com/axios/axios/commit/eb0a2db04beda089e6bdcb2820f193ed2faecbc3"><code>eb0a2db</code></a> chore: fix typos in test (<a href="https://redirect.github.com/axios/axios/issues/6923">#6923</a>)</li>
<li><a href="https://github.com/axios/axios/commit/7d551393c384e58058e04ae954c4cfd929afcd64"><code>7d55139</code></a> docs(readme): improve error descriptions; (<a href="https://redirect.github.com/axios/axios/issues/6920">#6920</a>)</li>
<li><a href="https://github.com/axios/axios/commit/f4fc6b8564ab794e67b4d1147167f2ecfc3557a3"><code>f4fc6b8</code></a> chore(sponsor): update sponsor block (<a href="https://redirect.github.com/axios/axios/issues/6921">#6921</a>)</li>
<li><a href="https://github.com/axios/axios/commit/5064b108de336ff34862650709761b8a96d26be0"><code>5064b10</code></a> fix(form-data): convert boolean values to strings in FormData serialization (...</li>
<li><a href="https://github.com/axios/axios/commit/c7e0fea78716e86694d5023f8f17d174bf064e8a"><code>c7e0fea</code></a> CI: add Node 24 (<a href="https://redirect.github.com/axios/axios/issues/6913">#6913</a>)</li>
<li><a href="https://github.com/axios/axios/commit/7ba895c8874f3fdc4e9da992b2b9e34fe5a25b55"><code>7ba895c</code></a> chore(sponsor): update sponsor block (<a href="https://redirect.github.com/axios/axios/issues/6907">#6907</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v1.7.9...v1.10.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=axios&package-manager=npm_and_yarn&previous-version=1.7.9&new-version=1.10.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)


</details>

---

## Discussion

### Comment by @dependabot on August 12, 2025 at 01:52 AM UTC

Superseded by #1365.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1345*
