---
type: pull_request
number: 1251
title: "chore(deps): bump the lint group across 1 directory with 5 updates"
state: closed
author: dependabot
created: 2025-01-15T10:14:02Z
updated: 2025-01-17T12:48:36Z
closed: 2025-01-17T12:48:34Z
base_branch: master
head_branch: dependabot/npm_and_yarn/lint-72e20c6a74
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1251
---

# Pull Request #1251: chore(deps): bump the lint group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: January 15, 2025 at 10:14 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/lint-72e20c6a74`

## Description

Bumps the lint group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [eslint](https://github.com/eslint/eslint) | `7.32.0` | `9.18.0` |
| [eslint-config-prettier](https://github.com/prettier/eslint-config-prettier) | `7.2.0` | `10.0.1` |
| [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) | `7.35.0` | `7.37.4` |
| [stylelint](https://github.com/stylelint/stylelint) | `16.10.0` | `16.13.2` |
| [stylelint-scss](https://github.com/stylelint-scss/stylelint-scss) | `6.8.0` | `6.10.0` |


Updates `eslint` from 7.32.0 to 9.18.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v9.18.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/e84e6e269c4aefc84952e17a1f967697b02b7ad2"><code>e84e6e2</code></a> feat: Report allowed methods for <code>no-console</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19306">#19306</a>) (Anna Bocharova)</li>
<li><a href="https://github.com/eslint/eslint/commit/8efc2d0c92dab6099f34c1479cd80bdc5cd1b07b"><code>8efc2d0</code></a> feat: unflag TypeScript config files (<a href="https://redirect.github.com/eslint/eslint/issues/19266">#19266</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/87a9352c621e7cd1d5bb77b3c08df7837363ea12"><code>87a9352</code></a> feat: check imports and class names in <code>no-shadow-restricted-names</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19272">#19272</a>) (Milos Djermanovic)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/da768d4541c4c30bfc33640a07a8d8a485520b18"><code>da768d4</code></a> fix: correct <code>overrideConfigFile</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19289">#19289</a>) (Francesco Trotta)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/d9c23c55be52a431141f38561c14140ee8b15686"><code>d9c23c5</code></a> docs: replace <code>var</code> with <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19325">#19325</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/8e1a898411fd16c73332d7a2dd28aff9bac8da01"><code>8e1a898</code></a> docs: add tabs to cli code blocks (<a href="https://redirect.github.com/eslint/eslint/issues/18784">#18784</a>) (Jay)</li>
<li><a href="https://github.com/eslint/eslint/commit/f3aeefbd6547c25d78819ab7e77cf36a2c26611c"><code>f3aeefb</code></a> docs: rewrite using let and const in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19320">#19320</a>) (PoloSpark)</li>
<li><a href="https://github.com/eslint/eslint/commit/0b680b3cc19c1e8d79ab94e7160051177c4adfe7"><code>0b680b3</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/98c86a99f7657a2f15ea30a251523446b10a7cad"><code>98c86a9</code></a> docs: <code>Edit this page</code> button link to different branches (<a href="https://redirect.github.com/eslint/eslint/issues/19228">#19228</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/6947901d14b18dbb2db259c9769bd8ac4cd04c3c"><code>6947901</code></a> docs: remove hardcoded edit link (<a href="https://redirect.github.com/eslint/eslint/issues/19323">#19323</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/03f2f442a9a8bec15e89786980c07be5980cdac5"><code>03f2f44</code></a> docs: rewrite var with const in rules examples (<a href="https://redirect.github.com/eslint/eslint/issues/19317">#19317</a>) (Thiago)</li>
<li><a href="https://github.com/eslint/eslint/commit/26c3003bfca2f7d98950446fdf5b3978d17a3a60"><code>26c3003</code></a> docs: Clarify dangers of eslint:all (<a href="https://redirect.github.com/eslint/eslint/issues/19318">#19318</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/c03825730d277405c357388d62ed48b3973083ba"><code>c038257</code></a> docs: add <code>eqeqeq</code> in related rules to <code>no-eq-null</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19310">#19310</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/89c8fc54c977ac457d3b5525a87cec1c51e72e23"><code>89c8fc5</code></a> docs: rewrite examples with var using let and const (<a href="https://redirect.github.com/eslint/eslint/issues/19315">#19315</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/db574c4d380e2d25b6111a06bd15caa83f75bb2d"><code>db574c4</code></a> docs: add missing backticks to <code>no-void</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19313">#19313</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/8d943c335c528a6a6a631dcbd98506238240ecfb"><code>8d943c3</code></a> docs: add missing backticks to <code>default-case-last</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19311">#19311</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/36ef8bbeab495ef2598a4b1f52e32b4cb50be5e2"><code>36ef8bb</code></a> docs: rewrite examples with var using let and const (<a href="https://redirect.github.com/eslint/eslint/issues/19298">#19298</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/1610c9ee1479f23b1bc5a6853d0b42b83dacdb7f"><code>1610c9e</code></a> docs: add missing backticks to <code>no-else-return</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19309">#19309</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/df409d8f76555c7baa4353d678d5fc460454a4d7"><code>df409d8</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/2e842138e689ee5623552e885c3a5ac1b0c2bfcf"><code>2e84213</code></a> docs: Fix Horizontal Scroll Overflow in Rule Description on Mobile View (<a href="https://redirect.github.com/eslint/eslint/issues/19304">#19304</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/6e7361bb6ae93c87fccdf2219379c7793517f17a"><code>6e7361b</code></a> docs: replace <code>var</code> with <code>let</code> and <code>const</code> in rule example (<a href="https://redirect.github.com/eslint/eslint/issues/19302">#19302</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/069af5e9ac43c7f33bd2a30abce3d5d94f504465"><code>069af5e</code></a> docs: rewrite <code>var</code> using <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19303">#19303</a>) (Kim GyeonWon)</li>
<li><a href="https://github.com/eslint/eslint/commit/064e35de95339cfedcad467c3c9871d5ff70c1a7"><code>064e35d</code></a> docs: remove 'I hope to' comments from scope-manager-interface (<a href="https://redirect.github.com/eslint/eslint/issues/19300">#19300</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/8e003056a805468b07bcf4edba83a90a932fb520"><code>8e00305</code></a> docs: replace <code>var</code> with <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19299">#19299</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/a559009f51ad9f081bae5252bb2b7a6e23c54767"><code>a559009</code></a> docs: Add warning about extending core rules (<a href="https://redirect.github.com/eslint/eslint/issues/19295">#19295</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/0bfdf6caaf3e1553c67a77da900245879c730ad3"><code>0bfdf6c</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/ce0b9ff04242f61c8c49fc1ce164eb45eb3c459a"><code>ce0b9ff</code></a> docs: add navigation link for <code>code explorer</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19285">#19285</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/e255cc98abef202929112378bfe133f260f2ac9d"><code>e255cc9</code></a> docs: add bluesky icon to footer (<a href="https://redirect.github.com/eslint/eslint/issues/19290">#19290</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/5d64851955f410f31c159a7097f6cc7d4a01d6a1"><code>5d64851</code></a> docs: remove outdated info about environments (<a href="https://redirect.github.com/eslint/eslint/issues/19296">#19296</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/eec01f04ae1c44f7c9a8c6afec59dd72f5a57600"><code>eec01f0</code></a> docs: switch rule examples config format to <code>languageOptions</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19277">#19277</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/b36ca0a490829c579358ec7193bde35275000e04"><code>b36ca0a</code></a> docs: Fixing Focus Order by Rearranging Element Sequence (<a href="https://redirect.github.com/eslint/eslint/issues/19241">#19241</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/d122c8a756bb8e232ef7c25cca6dcae645094835"><code>d122c8a</code></a> docs: add missing backticks to <code>sort-imports</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19282">#19282</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/0367a70a43346f1b9df8be75d38f98f9cfe4007c"><code>0367a70</code></a> docs: update custom parser docs (<a href="https://redirect.github.com/eslint/eslint/issues/19288">#19288</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/8c07ebb9004309f8691f972d554e8bbb3eb517bc"><code>8c07ebb</code></a> docs: add <code>border-radius</code> to <code>hX:target</code> selector styles (<a href="https://redirect.github.com/eslint/eslint/issues/19270">#19270</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/eff7c5721c101975a03e7906905f1fe2c9538df0"><code>eff7c57</code></a> docs: add limitation section in <code>no-loop-func</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19287">#19287</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/5db226f4da9ad7d53a4505a90290b68d4036c082"><code>5db226f</code></a> docs: add missing backticks in various parts of the documentation (<a href="https://redirect.github.com/eslint/eslint/issues/19269">#19269</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/789edbbae5aeeefc8fee94cd653b0b5f3e2ae3eb"><code>789edbb</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/613c06a2c341758739473409a2331074884ec7f8"><code>613c06a</code></a> docs: mark rules that are frozen with ❄️ (<a href="https://redirect.github.com/eslint/eslint/issues/19231">#19231</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/43172ecbd449c13a503cb39539e31106179f5d80"><code>43172ec</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/ac8b3c4ca9f7b84f84356137cf23a1ba6dfecf11"><code>ac8b3c4</code></a> docs: fix description of <code>overrideConfigFile</code> option (<a href="https://redirect.github.com/eslint/eslint/issues/19262">#19262</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/bbb9b46c20662019e98df85dedde9b68719afa1f"><code>bbb9b46</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/995b49231a3f0ccddb941663175ce4fead9c9432"><code>995b492</code></a> docs: fix inconsistent divider in rule categories box (<a href="https://redirect.github.com/eslint/eslint/issues/19249">#19249</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/f76d05da6e745adbea574c32b334638c7ba3c0c8"><code>f76d05d</code></a> docs: Refactor search result handling with better event listener cleanup (<a href="https://redirect.github.com/eslint/eslint/issues/19252">#19252</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/c5f3d7dab303468ae33ccfec61bba75a816f832c"><code>c5f3d7d</code></a> docs: Update README (GitHub Actions Bot)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/blob/main/CHANGELOG.md">eslint's changelog</a>.</em></p>
<blockquote>
<p>v9.18.0 - January 10, 2025</p>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/c52be85c4a916f70807377e1a486adb3a5857347"><code>c52be85</code></a> chore: upgrade to <code>@eslint/js@9.18.0</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19330">#19330</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/362099c580992b2602316fc417ce3e595b96f28c"><code>362099c</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/94861418f1573e4e1cbdd0174413d19054553294"><code>9486141</code></a> deps: upgrade <code>@eslint/core</code> and <code>@eslint/plugin-kit</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19329">#19329</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/d9c23c55be52a431141f38561c14140ee8b15686"><code>d9c23c5</code></a> docs: replace <code>var</code> with <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19325">#19325</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/8e1a898411fd16c73332d7a2dd28aff9bac8da01"><code>8e1a898</code></a> docs: add tabs to cli code blocks (<a href="https://redirect.github.com/eslint/eslint/issues/18784">#18784</a>) (Jay)</li>
<li><a href="https://github.com/eslint/eslint/commit/f3aeefbd6547c25d78819ab7e77cf36a2c26611c"><code>f3aeefb</code></a> docs: rewrite using let and const in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19320">#19320</a>) (PoloSpark)</li>
<li><a href="https://github.com/eslint/eslint/commit/0b680b3cc19c1e8d79ab94e7160051177c4adfe7"><code>0b680b3</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/98c86a99f7657a2f15ea30a251523446b10a7cad"><code>98c86a9</code></a> docs: <code>Edit this page</code> button link to different branches (<a href="https://redirect.github.com/eslint/eslint/issues/19228">#19228</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/6947901d14b18dbb2db259c9769bd8ac4cd04c3c"><code>6947901</code></a> docs: remove hardcoded edit link (<a href="https://redirect.github.com/eslint/eslint/issues/19323">#19323</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/03f2f442a9a8bec15e89786980c07be5980cdac5"><code>03f2f44</code></a> docs: rewrite var with const in rules examples (<a href="https://redirect.github.com/eslint/eslint/issues/19317">#19317</a>) (Thiago)</li>
<li><a href="https://github.com/eslint/eslint/commit/26c3003bfca2f7d98950446fdf5b3978d17a3a60"><code>26c3003</code></a> docs: Clarify dangers of eslint:all (<a href="https://redirect.github.com/eslint/eslint/issues/19318">#19318</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/c03825730d277405c357388d62ed48b3973083ba"><code>c038257</code></a> docs: add <code>eqeqeq</code> in related rules to <code>no-eq-null</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19310">#19310</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/89c8fc54c977ac457d3b5525a87cec1c51e72e23"><code>89c8fc5</code></a> docs: rewrite examples with var using let and const (<a href="https://redirect.github.com/eslint/eslint/issues/19315">#19315</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/495aa499a7390f99b763cba8f2b8312e3eecfe0d"><code>495aa49</code></a> chore: extract package <code>name</code> from <code>package.json</code> for public interface (<a href="https://redirect.github.com/eslint/eslint/issues/19314">#19314</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/db574c4d380e2d25b6111a06bd15caa83f75bb2d"><code>db574c4</code></a> docs: add missing backticks to <code>no-void</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19313">#19313</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/8d943c335c528a6a6a631dcbd98506238240ecfb"><code>8d943c3</code></a> docs: add missing backticks to <code>default-case-last</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19311">#19311</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/36ef8bbeab495ef2598a4b1f52e32b4cb50be5e2"><code>36ef8bb</code></a> docs: rewrite examples with var using let and const (<a href="https://redirect.github.com/eslint/eslint/issues/19298">#19298</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/1610c9ee1479f23b1bc5a6853d0b42b83dacdb7f"><code>1610c9e</code></a> docs: add missing backticks to <code>no-else-return</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19309">#19309</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/df409d8f76555c7baa4353d678d5fc460454a4d7"><code>df409d8</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/e84e6e269c4aefc84952e17a1f967697b02b7ad2"><code>e84e6e2</code></a> feat: Report allowed methods for <code>no-console</code> rule (<a href="https://redirect.github.com/eslint/eslint/issues/19306">#19306</a>) (Anna Bocharova)</li>
<li><a href="https://github.com/eslint/eslint/commit/2e842138e689ee5623552e885c3a5ac1b0c2bfcf"><code>2e84213</code></a> docs: Fix Horizontal Scroll Overflow in Rule Description on Mobile View (<a href="https://redirect.github.com/eslint/eslint/issues/19304">#19304</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/6e7361bb6ae93c87fccdf2219379c7793517f17a"><code>6e7361b</code></a> docs: replace <code>var</code> with <code>let</code> and <code>const</code> in rule example (<a href="https://redirect.github.com/eslint/eslint/issues/19302">#19302</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/069af5e9ac43c7f33bd2a30abce3d5d94f504465"><code>069af5e</code></a> docs: rewrite <code>var</code> using <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19303">#19303</a>) (Kim GyeonWon)</li>
<li><a href="https://github.com/eslint/eslint/commit/064e35de95339cfedcad467c3c9871d5ff70c1a7"><code>064e35d</code></a> docs: remove 'I hope to' comments from scope-manager-interface (<a href="https://redirect.github.com/eslint/eslint/issues/19300">#19300</a>) (Josh Goldberg ✨)</li>
<li><a href="https://github.com/eslint/eslint/commit/8e003056a805468b07bcf4edba83a90a932fb520"><code>8e00305</code></a> docs: replace <code>var</code> with <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19299">#19299</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/a559009f51ad9f081bae5252bb2b7a6e23c54767"><code>a559009</code></a> docs: Add warning about extending core rules (<a href="https://redirect.github.com/eslint/eslint/issues/19295">#19295</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/0bfdf6caaf3e1553c67a77da900245879c730ad3"><code>0bfdf6c</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/ce0b9ff04242f61c8c49fc1ce164eb45eb3c459a"><code>ce0b9ff</code></a> docs: add navigation link for <code>code explorer</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19285">#19285</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/e255cc98abef202929112378bfe133f260f2ac9d"><code>e255cc9</code></a> docs: add bluesky icon to footer (<a href="https://redirect.github.com/eslint/eslint/issues/19290">#19290</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/5d64851955f410f31c159a7097f6cc7d4a01d6a1"><code>5d64851</code></a> docs: remove outdated info about environments (<a href="https://redirect.github.com/eslint/eslint/issues/19296">#19296</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/eec01f04ae1c44f7c9a8c6afec59dd72f5a57600"><code>eec01f0</code></a> docs: switch rule examples config format to <code>languageOptions</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19277">#19277</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/b36ca0a490829c579358ec7193bde35275000e04"><code>b36ca0a</code></a> docs: Fixing Focus Order by Rearranging Element Sequence (<a href="https://redirect.github.com/eslint/eslint/issues/19241">#19241</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/d122c8a756bb8e232ef7c25cca6dcae645094835"><code>d122c8a</code></a> docs: add missing backticks to <code>sort-imports</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19282">#19282</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/0367a70a43346f1b9df8be75d38f98f9cfe4007c"><code>0367a70</code></a> docs: update custom parser docs (<a href="https://redirect.github.com/eslint/eslint/issues/19288">#19288</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/da768d4541c4c30bfc33640a07a8d8a485520b18"><code>da768d4</code></a> fix: correct <code>overrideConfigFile</code> type (<a href="https://redirect.github.com/eslint/eslint/issues/19289">#19289</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/8c07ebb9004309f8691f972d554e8bbb3eb517bc"><code>8c07ebb</code></a> docs: add <code>border-radius</code> to <code>hX:target</code> selector styles (<a href="https://redirect.github.com/eslint/eslint/issues/19270">#19270</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/eff7c5721c101975a03e7906905f1fe2c9538df0"><code>eff7c57</code></a> docs: add limitation section in <code>no-loop-func</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19287">#19287</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/8efc2d0c92dab6099f34c1479cd80bdc5cd1b07b"><code>8efc2d0</code></a> feat: unflag TypeScript config files (<a href="https://redirect.github.com/eslint/eslint/issues/19266">#19266</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/87a9352c621e7cd1d5bb77b3c08df7837363ea12"><code>87a9352</code></a> feat: check imports and class names in <code>no-shadow-restricted-names</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19272">#19272</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/5db226f4da9ad7d53a4505a90290b68d4036c082"><code>5db226f</code></a> docs: add missing backticks in various parts of the documentation (<a href="https://redirect.github.com/eslint/eslint/issues/19269">#19269</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/789edbbae5aeeefc8fee94cd653b0b5f3e2ae3eb"><code>789edbb</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/613c06a2c341758739473409a2331074884ec7f8"><code>613c06a</code></a> docs: mark rules that are frozen with ❄️ (<a href="https://redirect.github.com/eslint/eslint/issues/19231">#19231</a>) (Amaresh  S M)</li>
<li><a href="https://github.com/eslint/eslint/commit/43172ecbd449c13a503cb39539e31106179f5d80"><code>43172ec</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/ac8b3c4ca9f7b84f84356137cf23a1ba6dfecf11"><code>ac8b3c4</code></a> docs: fix description of <code>overrideConfigFile</code> option (<a href="https://redirect.github.com/eslint/eslint/issues/19262">#19262</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/6fe0e7244a7e88458ea7fdcebc43794c03793c4b"><code>6fe0e72</code></a> chore: update dependency <code>@​eslint/json</code> to ^0.9.0 (<a href="https://redirect.github.com/eslint/eslint/issues/19263">#19263</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/bbb9b46c20662019e98df85dedde9b68719afa1f"><code>bbb9b46</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/995b49231a3f0ccddb941663175ce4fead9c9432"><code>995b492</code></a> docs: fix inconsistent divider in rule categories box (<a href="https://redirect.github.com/eslint/eslint/issues/19249">#19249</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/f76d05da6e745adbea574c32b334638c7ba3c0c8"><code>f76d05d</code></a> docs: Refactor search result handling with better event listener cleanup (<a href="https://redirect.github.com/eslint/eslint/issues/19252">#19252</a>) (Amaresh  S M)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/1c87b415313b4aacda496b3b26efc4e0b93dc13a"><code>1c87b41</code></a> 9.18.0</li>
<li><a href="https://github.com/eslint/eslint/commit/4df3132ae241313b57170edb4fb1354abae840ce"><code>4df3132</code></a> Build: changelog update for 9.18.0</li>
<li><a href="https://github.com/eslint/eslint/commit/c52be85c4a916f70807377e1a486adb3a5857347"><code>c52be85</code></a> chore: upgrade to <code>@eslint/js@9.18.0</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19330">#19330</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/362099c580992b2602316fc417ce3e595b96f28c"><code>362099c</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/94861418f1573e4e1cbdd0174413d19054553294"><code>9486141</code></a> deps: upgrade <code>@eslint/core</code> and <code>@eslint/plugin-kit</code> (<a href="https://redirect.github.com/eslint/eslint/issues/19329">#19329</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/d9c23c55be52a431141f38561c14140ee8b15686"><code>d9c23c5</code></a> docs: replace <code>var</code> with <code>const</code> in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19325">#19325</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/8e1a898411fd16c73332d7a2dd28aff9bac8da01"><code>8e1a898</code></a> docs: add tabs to cli code blocks (<a href="https://redirect.github.com/eslint/eslint/issues/18784">#18784</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/f3aeefbd6547c25d78819ab7e77cf36a2c26611c"><code>f3aeefb</code></a> docs: rewrite using let and const in rule examples (<a href="https://redirect.github.com/eslint/eslint/issues/19320">#19320</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/0b680b3cc19c1e8d79ab94e7160051177c4adfe7"><code>0b680b3</code></a> docs: Update README</li>
<li><a href="https://github.com/eslint/eslint/commit/98c86a99f7657a2f15ea30a251523446b10a7cad"><code>98c86a9</code></a> docs: <code>Edit this page</code> button link to different branches (<a href="https://redirect.github.com/eslint/eslint/issues/19228">#19228</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v7.32.0...v9.18.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~eslintbot">eslintbot</a>, a new releaser for eslint since your current version.</p>
</details>
<br />

Updates `eslint-config-prettier` from 7.2.0 to 10.0.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/releases">eslint-config-prettier's releases</a>.</em></p>
<blockquote>
<h2>v10.0.1</h2>
<h1>eslint-config-prettier</h1>
<h2>10.0.1</h2>
<h2>What's Changed</h2>
<ul>
<li>chore: migrate to changeset for automatically releasing by <a href="https://github.com/JounQin"><code>@​JounQin</code></a> in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/278">prettier/eslint-config-prettier#278</a></li>
<li>add support for <code>@stylistic/eslint-plugin</code> by <a href="https://github.com/abrahamguo"><code>@​abrahamguo</code></a> in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/272">prettier/eslint-config-prettier#272</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/JounQin"><code>@​JounQin</code></a> made their first contribution in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/278">prettier/eslint-config-prettier#278</a></li>
<li><a href="https://github.com/abrahamguo"><code>@​abrahamguo</code></a> made their first contribution in <a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/272">prettier/eslint-config-prettier#272</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/prettier/eslint-config-prettier/compare/v9.1.0...v10.0.1">https://github.com/prettier/eslint-config-prettier/compare/v9.1.0...v10.0.1</a></p>
<h2>v10.0.0</h2>
<h3>Major Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/272">#272</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/5be64bef68c3a9bf7202f591f54ffec02572e46b"><code>5be64be</code></a> Thanks <a href="https://github.com/abrahamguo"><code>@​abrahamguo</code></a>! - add support for <code>@stylistic</code> formatting rules</li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/eslint-config-prettier/blob/main/CHANGELOG.md">eslint-config-prettier's changelog</a>.</em></p>
<blockquote>
<h1>eslint-config-prettier</h1>
<h2>10.0.0</h2>
<h3>Major Changes</h3>
<ul>
<li><a href="https://redirect.github.com/prettier/eslint-config-prettier/pull/272">#272</a> <a href="https://github.com/prettier/eslint-config-prettier/commit/5be64bef68c3a9bf7202f591f54ffec02572e46b"><code>5be64be</code></a> Thanks <a href="https://github.com/abrahamguo"><code>@​abrahamguo</code></a>! - add support for <a href="https://github.com/stylistic"><code>@​stylistic</code></a> formatting rules</li>
</ul>
<h2>Versions before 10.0.0</h2>
<h3>Version 9.1.0 (2023-12-02)</h3>
<ul>
<li>Added: [unicorn/template-indent], (as a [special rule][unicorn/template-indent-special]). Thanks to Gürgün Dayıoğlu (<a href="https://github.com/gurgunday"><code>@​gurgunday</code></a>)!</li>
<li>Changed: All the [formatting rules that were deprecated in ESLint 8.53.0][deprecated-8.53.0] are now excluded if you set the <code>ESLINT_CONFIG_PRETTIER_NO_DEPRECATED</code> environment variable.</li>
</ul>
<h3>Version 9.0.0 (2023-08-05)</h3>
<ul>
<li>Added: The CLI helper tool now works with eslint.config.js (flat config). Just like ESLint itself, the CLI tool automatically first tries eslint.config.js and then eslintrc, and you can force which one to use by setting the [ESLINT_USE_FLAT_CONFIG] environment variable. Note that the <em>config</em> of eslint-config-prettier has always been compatible with eslint.config.js (flat config) – it was just the CLI tool that needed updating. On top of that, the docs have been updated to mention how to use both eslint.config.js (flat config) and eslintrc, and the tests now test both config systems.</li>
<li>Changed: [unicode-bom] is no longer turned off. Prettier preserves the BOM if you have one, and does not add one if missing. It was wrong of eslint-config-prettier to disable that rule. If you get ESLint errors after upgrading, either add <code>&quot;unicode-bom&quot;: &quot;off&quot;</code> to your config to disable it again, or run ESLint with <code>--fix</code> to fix all files according to the rule (add or remove BOM). Thanks to Nicolas Stepien (<a href="https://github.com/nstepien"><code>@​nstepien</code></a>)!</li>
</ul>
<h3>Version 8.10.0 (2023-08-03)</h3>
<ul>
<li>Added: [max-statements-per-line]. Thanks to <a href="https://github.com/Zamiell"><code>@​Zamiell</code></a>!</li>
</ul>
<h3>Version 8.9.0 (2023-07-27)</h3>
<ul>
<li>Added: [vue/array-element-newline]. Thanks to <a href="https://github.com/xcatliu"><code>@​xcatliu</code></a>!</li>
</ul>
<h3>Version 8.8.0 (2023-03-20)</h3>
<ul>
<li>Added: [<code>@​typescript-eslint/lines-around-comment</code>]. Thanks to <a href="https://github.com/ttionya"><code>@​ttionya</code></a>!</li>
</ul>
<h3>Version 8.7.0 (2023-03-06)</h3>
<ul>
<li>Added: [<code>@​typescript-eslint/block-spacing</code>]. Thanks to <a href="https://github.com/ttionya"><code>@​ttionya</code></a>!</li>
<li>Added: [<code>@​typescript-eslint/key-spacing</code>]. Thanks to <a href="https://github.com/ttionya"><code>@​ttionya</code></a>!</li>
</ul>
<h3>Version 8.6.0 (2023-01-02)</h3>
<ul>
<li>Added: [vue/multiline-ternary]. Thanks to <a href="https://github.com/xcatliu"><code>@​xcatliu</code></a>!</li>
</ul>
<h3>Version 8.5.0 (2022-03-02)</h3>
<ul>
<li>Added: [<code>@​typescript-eslint/space-before-blocks</code>]. Thanks to Masafumi Koba (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>)!</li>
</ul>
<h3>Version 8.4.0 (2022-02-19)</h3>
<ul>
<li>Added: [vue/quote-props]. Thanks to <a href="https://github.com/xcatliu"><code>@​xcatliu</code></a>!</li>
</ul>
<h3>Version 8.3.0 (2021-04-24)</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/c5170f0281a4a439acb2e31fa2136bf6dcd54de0"><code>c5170f0</code></a> fix: add main field</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/e814e707b19ca96737218f5d014737ca7cd939d0"><code>e814e70</code></a> chore: change release folders</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/9ebedc7c7dafe099237d4606839cb71bca45c04c"><code>9ebedc7</code></a> chore: release eslint-config-prettier (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/296">#296</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/5be64bef68c3a9bf7202f591f54ffec02572e46b"><code>5be64be</code></a> feat: add support for <code>@stylistic/eslint-plugin</code> (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/272">#272</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/5687e7de69ac132f053b4bc5474d5963aa3bf6d6"><code>5687e7d</code></a> chore: migrate to changeset for automatically releasing (<a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/278">#278</a>)</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4f3bbb445fe2bbf39efb5fc4404ed5159baab78d"><code>4f3bbb4</code></a> Remove unused eslint-disable-next-line comment</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/40c7f3d0624129934bc3b40ab13a8ed938c6313b"><code>40c7f3d</code></a> eslint-config-prettier v9.1.0</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4110dff0c5b258be37506ecee9578cd7ff8e4759"><code>4110dff</code></a> Merge pull request <a href="https://redirect.github.com/prettier/eslint-config-prettier/issues/271">#271</a> from prettier/deprecated</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/6d0bd9294aeeea34cf9004bde2e6cb79883141fa"><code>6d0bd92</code></a> Update tests to handle newly deprecated rules</li>
<li><a href="https://github.com/prettier/eslint-config-prettier/commit/4c876b9424f38e52cee1118ef51ff19ce041cee1"><code>4c876b9</code></a> Move rules deprecated in ESLint 8.53.0 to the deprecated section</li>
<li>Additional commits viewable in <a href="https://github.com/prettier/eslint-config-prettier/compare/v7.2.0...v10.0.1">compare view</a></li>
</ul>
</details>
<details>
<summary>Maintainer changes</summary>
<p>This version was pushed to npm by <a href="https://www.npmjs.com/~jounqin">jounqin</a>, a new releaser for eslint-config-prettier since your current version.</p>
</details>
<br />

Updates `eslint-plugin-react` from 7.35.0 to 7.37.4
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/jsx-eslint/eslint-plugin-react/releases">eslint-plugin-react's releases</a>.</em></p>
<blockquote>
<h2>v7.37.4</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-unknown-property</code>]: support <code>onBeforeToggle</code>, <code>popoverTarget</code>, <code>popoverTargetAction</code> attributes (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3865">#3865</a>[] <a href="https://github.com/acusti"><code>@​acusti</code></a>)</li>
<li>[types] fix types of flat configs (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3874">#3874</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1000">#1000</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1000">jsx-eslint/eslint-plugin-react#1000</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1002">#1002</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1002">jsx-eslint/eslint-plugin-react#1002</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1005">#1005</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1005">jsx-eslint/eslint-plugin-react#1005</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/100">#100</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/100">jsx-eslint/eslint-plugin-react#100</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1010">#1010</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1010">jsx-eslint/eslint-plugin-react#1010</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1013">#1013</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1013">jsx-eslint/eslint-plugin-react#1013</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1022">#1022</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1022">jsx-eslint/eslint-plugin-react#1022</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1029">#1029</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1029">jsx-eslint/eslint-plugin-react#1029</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/102">#102</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/102">jsx-eslint/eslint-plugin-react#102</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1034">#1034</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1034">jsx-eslint/eslint-plugin-react#1034</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1038">#1038</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1038">jsx-eslint/eslint-plugin-react#1038</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1041">#1041</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1041">jsx-eslint/eslint-plugin-react#1041</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1043">#1043</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1043">jsx-eslint/eslint-plugin-react#1043</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1046">#1046</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1046">jsx-eslint/eslint-plugin-react#1046</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1047">#1047</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1047">jsx-eslint/eslint-plugin-react#1047</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1050">#1050</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1050">jsx-eslint/eslint-plugin-react#1050</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1053">#1053</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1053">jsx-eslint/eslint-plugin-react#1053</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1057">#1057</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1057">jsx-eslint/eslint-plugin-react#1057</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/105">#105</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/105">jsx-eslint/eslint-plugin-react#105</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1061">#1061</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1061">jsx-eslint/eslint-plugin-react#1061</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1062">#1062</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1062">jsx-eslint/eslint-plugin-react#1062</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1070">#1070</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1070">jsx-eslint/eslint-plugin-react#1070</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1071">#1071</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1071">jsx-eslint/eslint-plugin-react#1071</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1073">#1073</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1073">jsx-eslint/eslint-plugin-react#1073</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1076">#1076</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1076">jsx-eslint/eslint-plugin-react#1076</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1079">#1079</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1079">jsx-eslint/eslint-plugin-react#1079</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1088">#1088</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1088">jsx-eslint/eslint-plugin-react#1088</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1098">#1098</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1098">jsx-eslint/eslint-plugin-react#1098</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1101">#1101</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1101">jsx-eslint/eslint-plugin-react#1101</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1103">#1103</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1103">jsx-eslint/eslint-plugin-react#1103</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/110">#110</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/110">jsx-eslint/eslint-plugin-react#110</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1116">#1116</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1116">jsx-eslint/eslint-plugin-react#1116</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1117">#1117</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1117">jsx-eslint/eslint-plugin-react#1117</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1119">#1119</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1119">jsx-eslint/eslint-plugin-react#1119</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1121">#1121</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1121">jsx-eslint/eslint-plugin-react#1121</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1122">#1122</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1122">jsx-eslint/eslint-plugin-react#1122</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1123">#1123</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1123">jsx-eslint/eslint-plugin-react#1123</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1130">#1130</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1130">jsx-eslint/eslint-plugin-react#1130</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1131">#1131</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1131">jsx-eslint/eslint-plugin-react#1131</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1132">#1132</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1132">jsx-eslint/eslint-plugin-react#1132</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1134">#1134</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1134">jsx-eslint/eslint-plugin-react#1134</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1135">#1135</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1135">jsx-eslint/eslint-plugin-react#1135</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1139">#1139</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1139">jsx-eslint/eslint-plugin-react#1139</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1148">#1148</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1148">jsx-eslint/eslint-plugin-react#1148</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/1149">#1149</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/1149">jsx-eslint/eslint-plugin-react#1149</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/jsx-eslint/eslint-plugin-react/blob/master/CHANGELOG.md">eslint-plugin-react's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.3...v7.37.4">7.37.4</a> - 2025.01.12</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-unknown-property</code>]: support <code>onBeforeToggle</code>, <code>popoverTarget</code>, <code>popoverTargetAction</code> attributes (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3865">#3865</a>[] <a href="https://github.com/acusti"><code>@​acusti</code></a>)</li>
<li>[types] fix types of flat configs (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3874">#3874</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3874">#3874</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3874">jsx-eslint/eslint-plugin-react#3874</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3865">#3865</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3865">jsx-eslint/eslint-plugin-react#3865</a></p>
<h2><a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.2...v7.37.3">7.37.3</a> - 2024.12.23</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>no-danger</code>]: avoid a crash on a nested component name (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3833">#3833</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
<li>[Fix] types: correct generated type declaration (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3840">#3840</a>[] <a href="https://github.com/ocavue"><code>@​ocavue</code></a>)</li>
<li>[<code>no-unknown-property</code>]: support <code>precedence</code> prop in react 19 (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3829">#3829</a>[] <a href="https://github.com/acusti"><code>@​acusti</code></a>)</li>
<li>[<code>prop-types</code>]: props missing in validation when using generic types from a namespace import (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3859">#3859</a>[] <a href="https://github.com/rbondoc96"><code>@​rbondoc96</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[Tests] [<code>jsx-no-script-url</code>]: Improve tests (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3849">#3849</a>[] <a href="https://github.com/radu2147"><code>@​radu2147</code></a>)</li>
<li>[Docs] fix broken links: [<code>default-props-match-prop-types</code>], [<code>jsx-boolean-value</code>], [<code>jsx-curly-brace-presence</code>], [<code>jsx-no-bind</code>], [<code>no-array-index-key</code>], [<code>no-is-mounted</code>], [<code>no-render-return-value</code>], [<code>require-default-props</code>] (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3841">#3841</a>[] <a href="https://github.com/bastiendmt"><code>@​bastiendmt</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3859">#3859</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3859">jsx-eslint/eslint-plugin-react#3859</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3849">#3849</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3849">jsx-eslint/eslint-plugin-react#3849</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3841">#3841</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3841">jsx-eslint/eslint-plugin-react#3841</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3840">#3840</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3840">jsx-eslint/eslint-plugin-react#3840</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3833">#3833</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3833">jsx-eslint/eslint-plugin-react#3833</a>
<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3829">#3829</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3829">jsx-eslint/eslint-plugin-react#3829</a></p>
<h2><a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.1...v7.37.2">7.37.2</a> - 2024.10.22</h2>
<h3>Fixed</h3>
<ul>
<li>[<code>destructuring-assignment</code>]: fix false negative when using <code>typeof props.a</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3835">#3835</a>[] <a href="https://github.com/golopot"><code>@​golopot</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[Refactor] [<code>destructuring-assignment</code>]: use <code>getParentStatelessComponent</code> (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3835">#3835</a>[] <a href="https://github.com/golopot"><code>@​golopot</code></a>)</li>
</ul>
<p><a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3835">#3835</a>: <a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/pull/3835">jsx-eslint/eslint-plugin-react#3835</a></p>
<h2><a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.37.0...v7.37.1">7.37.1</a> - 2024.10.01</h2>
<h3>Fixed</h3>
<ul>
<li>[meta] do not npmignore <code>d.ts</code> files (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3836">#3836</a>[] <a href="https://github.com/ljharb"><code>@​ljharb</code></a>)</li>
</ul>
<h3>Changed</h3>
<ul>
<li>[readme] Fix shared settings link (<a href="https://redirect.github.com/jsx-eslint/eslint-plugin-react/issues/3834">#3834</a>[] <a href="https://github.com/MgenGlder"><code>@​MgenGlder</code></a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/e6b5b41191690ee166d0cca1e9db27092b910f03"><code>e6b5b41</code></a> Update CHANGELOG and bump version</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/cfd5edd4f66905ea3300d3ee96289b9fefbbfb5f"><code>cfd5edd</code></a> [Dev Deps] update <code>@babel/eslint-parser</code></li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/efc021fbece010b3f4a658cd1b05753e9e2776f1"><code>efc021f</code></a> [types] fix types of flat configs</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/82a196a0b2c7c5c91e57f320428b58900cb4a8cf"><code>82a196a</code></a> [Fix] <code>no-unknown-property</code>: support <code>onBeforeToggle</code>, <code>popoverTarget</code>, `popo...</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/5c816edc1d146767143ce894707bd7284aca400c"><code>5c816ed</code></a> [actions] publish action: allow additional URL</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/244743a7a974787f60d0afb05dd2cae8f7392064"><code>244743a</code></a> Update CHANGELOG and bump version</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/bc5b9dda28e37ee2a5dc9daa98868ba4a339839e"><code>bc5b9dd</code></a> [actions] release workflow needs some new domains</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/ed3b1cf8f398a284353e9f8d41a4ab7b1c31a543"><code>ed3b1cf</code></a> [Tests] <code>jsx-uses-vars</code>, <code>jsx-uses-react</code>: fix <code>no-unused-vars</code> tests in esli...</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/033ad19c6afd4cc6a0777f6c174e7f2dadbd7853"><code>033ad19</code></a> [Deps] update <code>array.prototype.flatmap</code>, <code>es-iterator-helpers</code>, `object.value...</li>
<li><a href="https://github.com/jsx-eslint/eslint-plugin-react/commit/b4a14f4cdbac6e6ad78f79bd3cec167416ddbc90"><code>b4a14f4</code></a> [meta] add <code>directories.test</code></li>
<li>Additional commits viewable in <a href="https://github.com/jsx-eslint/eslint-plugin-react/compare/v7.35.0...v7.37.4">compare view</a></li>
</ul>
</details>
<br />

Updates `stylelint` from 16.10.0 to 16.13.2
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stylelint/stylelint/releases">stylelint's releases</a>.</em></p>
<blockquote>
<h2>16.13.2</h2>
<ul>
<li>Fixed: <code>--fix</code> CLI flag raising unknown value error (<a href="https://redirect.github.com/stylelint/stylelint/issues/8313">#8313</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
</ul>
<h2>16.13.1</h2>
<ul>
<li>Fixed: <code>ignore.default is not a function</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8305">#8305</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
</ul>
<h2>16.13.0</h2>
<ul>
<li>Deprecated: ambiguous position arguments passed to <code>utils.report()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8244">#8244</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Added: <code>lax</code>/<code>strict</code> values to the <code>fix</code> Node.js API option and CLI flag (<a href="https://redirect.github.com/stylelint/stylelint/issues/8106">#8106</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: support for profiling rule performance via the <code>TIMING</code> environment variable (<a href="https://redirect.github.com/stylelint/stylelint/issues/8108">#8108</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-descriptor-no-unknown</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8197">#8197</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-descriptor-value-no-unknown</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8211">#8211</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>at-rule-no-deprecated</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8251">#8251</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Added: <code>at-rule-prelude-no-invalid</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8268">#8268</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>declaration-property-value-keyword-no-deprecated</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8223">#8223</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Added: <code>&quot;ignore&quot;: [&quot;at-rule-preludes&quot;, &quot;declaration-values&quot;]</code> to <code>string-no-newline</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8214">#8214</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>messageArgs</code> to <code>declaration-property-value-no-unknown</code>, <code>font-family-name-quotes</code>, <code>font-family-no-duplicate-names</code>, <code>function-calc-no-unspaced-operator</code>, <code>import-notation</code>, <code>media-feature-name-unit-allowed-list</code>, <code>selector-attribute-quotes</code> and <code>selector-pseudo-element-colon-notation</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8285">#8285</a> &amp; <a href="https://redirect.github.com/stylelint/stylelint/issues/8252">#8252</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: deprecation warnings to only display once per (custom) rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8265">#8265</a>) (<a href="https://github.com/romainmenke"><code>@​romainmenke</code></a>).</li>
<li>Fixed: <code>*-no-vendor-prefix</code> message ambiguity (<a href="https://redirect.github.com/stylelint/stylelint/issues/8239">#8239</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>at-rule-(dis)allowed-list</code>, <code>at-rule-no-vendor-prefix</code>, <code>at-rule-property-required-list</code> message argument (<a href="https://redirect.github.com/stylelint/stylelint/issues/8277">#8277</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>at-rule-property-required-list</code> message for inclusion of properties and descriptors (<a href="https://redirect.github.com/stylelint/stylelint/issues/8207">#8207</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>at-rule-*</code> false positives and negatives for <code>@charset</code> rule (<a href="https://redirect.github.com/stylelint/stylelint/issues/8215">#8215</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>declaration-property-value-no-unknown</code> false positives for descriptors (<a href="https://redirect.github.com/stylelint/stylelint/issues/8240">#8240</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>property-(dis)allowed-list</code> false negatives for custom properties, use <code>/^--/</code> to (dis)allow them (<a href="https://redirect.github.com/stylelint/stylelint/issues/8209">#8209</a>) (<a href="https://github.com/fbasmaison-lucca"><code>@​fbasmaison-lucca</code></a>).</li>
<li>Fixed: <code>property-no-unknown</code> false positives for descriptors (<a href="https://redirect.github.com/stylelint/stylelint/issues/8203">#8203</a>) (<a href="https://github.com/jeddy3"><code>@​jeddy3</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for deprecated pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8264">#8264</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-type-case</code> false positives for <code>hatchPath</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8264">#8264</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-type-no-unknown</code> false positives for <code>shadow</code>, <code>hatch</code> and <code>hatchpath</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8264">#8264</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
<h2>16.12.0</h2>
<ul>
<li>Added: <code>selector-pseudo-class-allowed-list</code> now checks <code>@page</code> pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8176">#8176</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Added: <code>selector-pseudo-class-disallowed-list</code> now checks <code>@page</code> pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8171">#8171</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>at-rule-property-required-list</code> message to use &quot;descriptor&quot; for accuracy (<a href="https://redirect.github.com/stylelint/stylelint/issues/8186">#8186</a>) (<a href="https://github.com/ybiquitous"><code>@​ybiquitous</code></a>).</li>
<li>Fixed: <code>custom-property-no-missing-var-function</code> false positives for <code>container-name</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8157">#8157</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>custom-property-no-missing-var-function</code> false positives for custom properties passed to <code>running()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8172">#8172</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>function-no-unknown</code> false positives for <code>running()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8172">#8172</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for <code>:open</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8184">#8184</a>) (<a href="https://github.com/lukewarlow"><code>@​lukewarlow</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for <code>:recto</code>, <code>:verso</code> and <code>:nth()</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8170">#8170</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-pseudo-class-no-unknown</code> false positives for some <code>moz-*</code> vendor-prefixed pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8188">#8188</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>selector-pseudo-element-no-unknown</code> false positives for <code>::details-content</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8174">#8174</a>) (<a href="https://github.com/lukewarlow"><code>@​lukewarlow</code></a>).</li>
<li>Fixed: <code>selector-type-no-unknown</code> false positives for idents in functional pseudo-classes (<a href="https://redirect.github.com/stylelint/stylelint/issues/8191">#8191</a>) (<a href="https://github.com/elskhn"><code>@​elskhn</code></a>).</li>
<li>Fixed: <code>value-keyword-case</code> false negatives (<a href="https://redirect.github.com/stylelint/stylelint/issues/8158">#8158</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Fixed: <code>value-keyword-case</code> false positives for vendor-prefixed system colors (<a href="https://redirect.github.com/stylelint/stylelint/issues/8146">#8146</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
</ul>
<h2>16.11.0</h2>
<ul>
<li>Added: <code>--report-unscoped-disables</code> CLI flag and <code>reportUnscopedDisables</code> option to Node.js API and configuration object (<a href="https://redirect.github.com/stylelint/stylelint/issues/8024">#8024</a>) (<a href="https://github.com/Mouvedia"><code>@​Mouvedia</code></a>).</li>
<li>Added: <code>ignoreFunctions: []</code> to <code>media-query-no-invalid</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8060">#8060</a>) (<a href="https://github.com/ryo-manba"><code>@​ryo-manba</code></a>).</li>
<li>Added: <code>name</code> configuration property under <code>overrides</code> (<a href="https://redirect.github.com/stylelint/stylelint/issues/8095">#8095</a>) (<a href="https://github.com/ryo-man...

_Description has been truncated_

---

## Discussion

### Comment by @dependabot on January 17, 2025 at 12:48 PM UTC

Superseded by #1257.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1251*
