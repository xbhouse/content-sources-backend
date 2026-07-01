---
type: pull_request
number: 1338
title: "chore(deps): bump axios from 1.7.9 to 1.9.0"
state: closed
author: dependabot
created: 2025-05-19T18:26:14Z
updated: 2025-06-23T20:03:23Z
closed: 2025-06-23T20:03:21Z
base_branch: master
head_branch: dependabot/npm_and_yarn/axios-1.9.0
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1338
---

# Pull Request #1338: chore(deps): bump axios from 1.7.9 to 1.9.0

**Author**: @dependabot
**Created**: May 19, 2025 at 06:26 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/axios-1.9.0`

## Description

Bumps [axios](https://github.com/axios/axios) from 1.7.9 to 1.9.0.
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
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
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>buildFullPath:</strong> handle <code>allowAbsoluteUrls: false</code> without <code>baseURL</code> (<a href="https://redirect.github.com/axios/axios/issues/6833">#6833</a>) (<a href="https://github.com/axios/axios/commit/f10c2e0de7fde0051f848609a29c2906d0caa1d9">f10c2e0</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+5/-1 ([#6833](https://github.com/axios/axios/issues/6833) )">Marc Hassan</a></li>
</ul>
<h2>Release v1.8.3</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li>add missing type for allowAbsoluteUrls (<a href="https://redirect.github.com/axios/axios/issues/6818">#6818</a>) (<a href="https://github.com/axios/axios/commit/10fa70ef14fe39558b15a179f0e82f5f5e5d11b2">10fa70e</a>)</li>
<li><strong>xhr/fetch:</strong> pass <code>allowAbsoluteUrls</code> to <code>buildFullPath</code> in <code>xhr</code> and <code>fetch</code> adapters (<a href="https://redirect.github.com/axios/axios/issues/6814">#6814</a>) (<a href="https://github.com/axios/axios/commit/ec159e507bdf08c04ba1a10fe7710094e9e50ec9">ec159e5</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/Electroid" title="+6/-0 ([#6811](https://github.com/axios/axios/issues/6811) )">Ashcon Partovi</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/StefanBRas" title="+4/-0 ([#6818](https://github.com/axios/axios/issues/6818) )">StefanBRas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+2/-2 ([#6814](https://github.com/axios/axios/issues/6814) )">Marc Hassan</a></li>
</ul>
<h2>Release v1.8.2</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v1.x/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
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
<h2><a href="https://github.com/axios/axios/compare/v1.8.3...v1.8.4">1.8.4</a> (2025-03-19)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>buildFullPath:</strong> handle <code>allowAbsoluteUrls: false</code> without <code>baseURL</code> (<a href="https://redirect.github.com/axios/axios/issues/6833">#6833</a>) (<a href="https://github.com/axios/axios/commit/f10c2e0de7fde0051f848609a29c2906d0caa1d9">f10c2e0</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/mhassan1" title="+5/-1 ([#6833](https://github.com/axios/axios/issues/6833) )">Marc Hassan</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.8.2...v1.8.3">1.8.3</a> (2025-03-10)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>add missing type for allowAbsoluteUrls (<a href="https://redirect.github.com/axios/axios/issues/6818">#6818</a>) (<a href="https://github.com/axios/axios/commit/10fa70ef14fe39558b15a179f0e82f5f5e5d11b2">10fa70e</a>)</li>
<li><strong>xhr/fetch:</strong> pass <code>allowAbsoluteUrls</code> to <code>buildFullPath</code> in <code>xhr</code> and <code>fetch</code> adapters (<a href="https://redirect.github.com/axios/axios/issues/6814">#6814</a>) (<a href="https://github.com/axios/axios/commit/ec159e507bdf08c04ba1a10fe7710094e9e50ec9">ec159e5</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/Electroid" title="+6/-0 ([#6811](https://github.com/axios/axios/issues/6811) )">Ashcon Partovi</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/StefanBRas" title="+4/-0 ([#6818](https://github.com/axios/axios/issues/6818) )">StefanBRas</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/cdcfd214c169a1acba8e267ab8e77ff4dfec3105"><code>cdcfd21</code></a> chore(release): v1.9.0 (<a href="https://redirect.github.com/axios/axios/issues/6891">#6891</a>)</li>
<li><a href="https://github.com/axios/axios/commit/987d2e2dd3b362757550f36eab875e60640b6ddc"><code>987d2e2</code></a> fix(http): send minimal end multipart boundary (<a href="https://redirect.github.com/axios/axios/issues/6661">#6661</a>)</li>
<li><a href="https://github.com/axios/axios/commit/f112edfdfacb5857ac4f91156fe8b116c456dfca"><code>f112edf</code></a> chore(ci): add PR files guard action; (<a href="https://redirect.github.com/axios/axios/issues/6890">#6890</a>)</li>
<li><a href="https://github.com/axios/axios/commit/61de4c05cc01e2ccd7705bbc87c4a49e6dea7928"><code>61de4c0</code></a> chore(ci): update github actions; (<a href="https://redirect.github.com/axios/axios/issues/6889">#6889</a>)</li>
<li><a href="https://github.com/axios/axios/commit/c3aba3d8f700337651d906b1a1ad3ecfe88a075e"><code>c3aba3d</code></a> chore(ci): add labeler github action; (<a href="https://redirect.github.com/axios/axios/issues/6888">#6888</a>)</li>
<li><a href="https://github.com/axios/axios/commit/f7a3b5e0f7e5e127b97defa92a132fbf1b55cf15"><code>f7a3b5e</code></a> fix(headers): fixed support for setting multiple header values from an iterat...</li>
<li><a href="https://github.com/axios/axios/commit/e61a8934d8f94dd429a2f309b48c67307c700df0"><code>e61a893</code></a> fix(types): fix autocomplete for adapter config (<a href="https://redirect.github.com/axios/axios/issues/6855">#6855</a>)</li>
<li><a href="https://github.com/axios/axios/commit/6c5d4cd69286868059c5e52d45085cb9a894a983"><code>6c5d4cd</code></a> fix(core): fix the Axios constructor implementation to treat the config argum...</li>
<li><a href="https://github.com/axios/axios/commit/dfe8411c9a082c3d068bdd1f8d6e73054f387f45"><code>dfe8411</code></a> fix(fetch): fixed ERR_NETWORK mapping for Safari browsers; (<a href="https://redirect.github.com/axios/axios/issues/6767">#6767</a>)</li>
<li><a href="https://github.com/axios/axios/commit/d4f7df4b304af8b373488fdf8e830793ff843eb9"><code>d4f7df4</code></a> fix(headers): fix <code>getSetCookie</code> by using 'get' method for caseless access; (...</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v1.7.9...v1.9.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=axios&package-manager=npm_and_yarn&previous-version=1.7.9&new-version=1.9.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

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

## Summary by Sourcery

Bump axios dependency from v1.7.9 to v1.9.0

New Features:
- Add AxiosHeaders.getSetCookie method to retrieve Set-Cookie header values

Bug Fixes:
- Fix Axios constructor to treat the config argument as optional
- Improve ERR_NETWORK error mapping in Safari fetch adapter
- Allow iterable objects as a data source for header set method
- Fix caseless access for getSetCookie in headers
- Support setting multiple header values from an iterable source
- Send minimal end boundary in multipart HTTP requests
- Fix autocomplete for adapter config types

---

## Discussion

### Comment by @sourcery-ai on May 19, 2025 at 06:26 PM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

This PR upgrades the axios dependency from v1.7.x to v1.9.0 by updating the version in package.json and regenerating the lockfile, pulling in the latest bug fixes and feature enhancements.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Upgrade axios to v1.9.0 | <ul><li>Bump axios version specifier in package.json</li><li>Regenerate package-lock.json to incorporate the new release</li></ul> | `package.json`<br/>`package-lock.json` |

---

<details>
<summary>Tips and commands</summary>

#### Interacting with Sourcery

- **Trigger a new review:** Comment `@sourcery-ai review` on the pull request.
- **Continue discussions:** Reply directly to Sourcery's review comments.
- **Generate a GitHub issue from a review comment:** Ask Sourcery to create an
  issue from a review comment by replying to it. You can also reply to a
  review comment with `@sourcery-ai issue` to create an issue from it.
- **Generate a pull request title:** Write `@sourcery-ai` anywhere in the pull
  request title to generate a title at any time. You can also comment
  `@sourcery-ai title` on the pull request to (re-)generate the title at any time.
- **Generate a pull request summary:** Write `@sourcery-ai summary` anywhere in
  the pull request body to generate a PR summary at any time exactly where you
  want it. You can also comment `@sourcery-ai summary` on the pull request to
  (re-)generate the summary at any time.
- **Generate reviewer's guide:** Comment `@sourcery-ai guide` on the pull
  request to (re-)generate the reviewer's guide at any time.
- **Resolve all Sourcery comments:** Comment `@sourcery-ai resolve` on the
  pull request to resolve all Sourcery comments. Useful if you've already
  addressed all the comments and don't want to see them anymore.
- **Dismiss all Sourcery reviews:** Comment `@sourcery-ai dismiss` on the pull
  request to dismiss all existing Sourcery reviews. Especially useful if you
  want to start fresh with a new review - don't forget to comment
  `@sourcery-ai review` to trigger a new review!

#### Customizing Your Experience

Access your [dashboard](https://app.sourcery.ai) to:
- Enable or disable review features such as the Sourcery-generated pull request
  summary, the reviewer's guide, and others.
- Change the review language.
- Add, remove or edit custom review instructions.
- Adjust other review settings.

#### Getting Help

- [Contact our support team](mailto:support@sourcery.ai) for questions or feedback.
- Visit our [documentation](https://docs.sourcery.ai) for detailed guides and information.
- Keep in touch with the Sourcery team by following us on [X/Twitter](https://x.com/SourceryAI), [LinkedIn](https://www.linkedin.com/company/sourcery-ai/) or [GitHub](https://github.com/sourcery-ai).

</details>

<!-- Generated by sourcery-ai[bot]: end review_guide -->

### Comment by @dependabot on June 23, 2025 at 08:03 PM UTC

Superseded by #1345.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1338*
