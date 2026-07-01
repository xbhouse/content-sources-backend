---
type: pull_request
number: 904
title: "Build: Bump the go group across 1 directory with 6 updates"
state: merged
author: dependabot
created: 2024-11-25T04:29:54Z
updated: 2024-11-25T20:24:06Z
closed: 2024-11-25T20:24:03Z
merged: 2024-11-25T20:24:03Z
base_branch: main
head_branch: dependabot/go_modules/go-9e8306995f
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/904
---

# Pull Request #904: Build: Bump the go group across 1 directory with 6 updates

**Author**: @dependabot
**Created**: November 25, 2024 at 04:29 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-9e8306995f`

## Description

Bumps the go group with 6 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | `1.9.0` | `1.10.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.32.4` | `1.32.5` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.45` | `1.17.46` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.43.2` | `1.44.0` |
| [github.com/content-services/caliri/release/v4](https://github.com/content-services/caliri) | `4.4.19` | `4.5.0` |
| [github.com/content-services/zest/release/v2024](https://github.com/content-services/zest) | `2024.11.1731697894` | `2024.11.1732298036` |


Updates `github.com/stretchr/testify` from 1.9.0 to 1.10.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/stretchr/testify/releases">github.com/stretchr/testify's releases</a>.</em></p>
<blockquote>
<h2>v1.10.0</h2>
<h2>What's Changed</h2>
<h3>Functional Changes</h3>
<ul>
<li>Add PanicAssertionFunc by <a href="https://github.com/fahimbagar"><code>@​fahimbagar</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1337">stretchr/testify#1337</a></li>
<li>assert: deprecate CompareType by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1566">stretchr/testify#1566</a></li>
<li>assert: make YAML dependency pluggable via build tags by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1579">stretchr/testify#1579</a></li>
<li>assert: new assertion NotElementsMatch by <a href="https://github.com/hendrywiranto"><code>@​hendrywiranto</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1600">stretchr/testify#1600</a></li>
<li>mock: in order mock calls by <a href="https://github.com/ReyOrtiz"><code>@​ReyOrtiz</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1637">stretchr/testify#1637</a></li>
<li>Add assertion for NotErrorAs by <a href="https://github.com/palsivertsen"><code>@​palsivertsen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1129">stretchr/testify#1129</a></li>
<li>Record Return Arguments of a Call by <a href="https://github.com/jayd3e"><code>@​jayd3e</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1636">stretchr/testify#1636</a></li>
<li>assert.EqualExportedValues: accepts everything by <a href="https://github.com/redachl"><code>@​redachl</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1586">stretchr/testify#1586</a></li>
</ul>
<h3>Fixes</h3>
<ul>
<li>assert: make tHelper a type alias by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1562">stretchr/testify#1562</a></li>
<li>Do not get argument again unnecessarily in Arguments.Error() by <a href="https://github.com/TomWright"><code>@​TomWright</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/820">stretchr/testify#820</a></li>
<li>Fix time.Time compare by <a href="https://github.com/myxo"><code>@​myxo</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1582">stretchr/testify#1582</a></li>
<li>assert.Regexp: handle []byte array properly by <a href="https://github.com/kevinburkesegment"><code>@​kevinburkesegment</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1587">stretchr/testify#1587</a></li>
<li>assert: collect.FailNow() should not panic by <a href="https://github.com/marshall-lee"><code>@​marshall-lee</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1481">stretchr/testify#1481</a></li>
<li>mock: simplify implementation of FunctionalOptions by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1571">stretchr/testify#1571</a></li>
<li>mock: caller information for unexpected method call by <a href="https://github.com/spirin"><code>@​spirin</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1644">stretchr/testify#1644</a></li>
<li>suite: fix test failures by <a href="https://github.com/stevenh"><code>@​stevenh</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1421">stretchr/testify#1421</a></li>
<li>Fix issue <a href="https://redirect.github.com/stretchr/testify/issues/1662">#1662</a> (comparing infs should fail) by <a href="https://github.com/ybrustin"><code>@​ybrustin</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1663">stretchr/testify#1663</a></li>
<li>NotSame should fail if args are not pointers <a href="https://redirect.github.com/stretchr/testify/issues/1661">#1661</a> by <a href="https://github.com/sikehish"><code>@​sikehish</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1664">stretchr/testify#1664</a></li>
<li>Increase timeouts in Test_Mock_Called_blocks to reduce flakiness in CI by <a href="https://github.com/sikehish"><code>@​sikehish</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1667">stretchr/testify#1667</a></li>
<li>fix: compare functional option names for indirect calls by <a href="https://github.com/arjun-1"><code>@​arjun-1</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1626">stretchr/testify#1626</a></li>
</ul>
<h3>Documantation, Build &amp; CI</h3>
<ul>
<li>.gitignore: ignore &quot;go test -c&quot; binaries by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1565">stretchr/testify#1565</a></li>
<li>mock: improve doc by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1570">stretchr/testify#1570</a></li>
<li>mock: fix FunctionalOptions docs by <a href="https://github.com/snirye"><code>@​snirye</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1433">stretchr/testify#1433</a></li>
<li>README: link out to the excellent testifylint by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1568">stretchr/testify#1568</a></li>
<li>assert: fix typo in comment by <a href="https://github.com/JohnEndson"><code>@​JohnEndson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1580">stretchr/testify#1580</a></li>
<li>Correct the EventuallyWithT and EventuallyWithTf example by <a href="https://github.com/JonCrowther"><code>@​JonCrowther</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1588">stretchr/testify#1588</a></li>
<li>CI: bump softprops/action-gh-release from 1 to 2 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1575">stretchr/testify#1575</a></li>
<li>mock: document more alternatives to deprecated AnythingOfTypeArgument by <a href="https://github.com/dolmen"><code>@​dolmen</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1569">stretchr/testify#1569</a></li>
<li>assert: Correctly document EqualValues behavior by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1593">stretchr/testify#1593</a></li>
<li>fix: grammar in godoc by <a href="https://github.com/miparnisari"><code>@​miparnisari</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1607">stretchr/testify#1607</a></li>
<li>.github/workflows: Run tests for Go 1.22 by <a href="https://github.com/HaraldNordgren"><code>@​HaraldNordgren</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1629">stretchr/testify#1629</a></li>
<li>Document suite's lack of support for t.Parallel by <a href="https://github.com/brackendawson"><code>@​brackendawson</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1645">stretchr/testify#1645</a></li>
<li>assert: fix typos in comments by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1650">stretchr/testify#1650</a></li>
<li>mock: fix doc comment for NotBefore by <a href="https://github.com/alexandear"><code>@​alexandear</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1651">stretchr/testify#1651</a></li>
<li>Generate better comments for require package by <a href="https://github.com/Neokil"><code>@​Neokil</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1610">stretchr/testify#1610</a></li>
<li>README: replace Testify V2 notice with <a href="https://github.com/dolmen"><code>@​dolmen</code></a>'s V2 manifesto by <a href="https://github.com/hendrywiranto"><code>@​hendrywiranto</code></a> in <a href="https://redirect.github.com/stretchr/testify/pull/1518">stretchr/testify#1518</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/fahimbagar"><code>@​fahimbagar</code></a> made their first contribution in <a href="https://redirect.github.com/stretchr/testify/pull/1337">stretchr/testify#1337</a></li>
<li><a href="https://github.com/TomWright"><code>@​TomWright</code></a> made their first contribution in <a href="https://redirect.github.com/stretchr/testify/pull/820">stretchr/testify#820</a></li>
<li><a href="https://github.com/snirye"><code>@​snirye</code></a> made their first contribution in <a href="https://redirect.github.com/stretchr/testify/pull/1433">stretchr/testify#1433</a></li>
<li><a href="https://github.com/myxo"><code>@​myxo</code></a> made their first contribution in <a href="https://redirect.github.com/stretchr/testify/pull/1582">stretchr/testify#1582</a></li>
<li><a href="https://github.com/JohnEndson"><code>@​JohnEndson</code></a> made their first contribution in <a href="https://redirect.github.com/stretchr/testify/pull/1580">stretchr/testify#1580</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/stretchr/testify/commit/89cbdd9e7b39eb58896d316a7495597d3aba4371"><code>89cbdd9</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1626">#1626</a> from arjun-1/fix-functional-options-diff-indirect-calls</li>
<li><a href="https://github.com/stretchr/testify/commit/07bac606be2191ae89a281195e58d01e0de7b5f8"><code>07bac60</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1667">#1667</a> from sikehish/flaky</li>
<li><a href="https://github.com/stretchr/testify/commit/716de8dff46ed7ae3c6ebb7a6124db741ba7c018"><code>716de8d</code></a> Increase timeouts in Test_Mock_Called_blocks to reduce flakiness in CI</li>
<li><a href="https://github.com/stretchr/testify/commit/118fb8346630c192421c8914848381af9d4412a7"><code>118fb83</code></a> NotSame should fail if args are not pointers <a href="https://redirect.github.com/stretchr/testify/issues/1661">#1661</a> (<a href="https://redirect.github.com/stretchr/testify/issues/1664">#1664</a>)</li>
<li><a href="https://github.com/stretchr/testify/commit/7d99b2b43d8f60a8982a78cde6e8bd287dea5da0"><code>7d99b2b</code></a> attempt 2</li>
<li><a href="https://github.com/stretchr/testify/commit/05f87c016035811e6d8371f1887ec360c318f53f"><code>05f87c0</code></a> more similar</li>
<li><a href="https://github.com/stretchr/testify/commit/ea7129e00694592e20cb34c58a6b8a251418b9da"><code>ea7129e</code></a> better fmt</li>
<li><a href="https://github.com/stretchr/testify/commit/a1b9c9efe3c25c50678b1e492045164b914e255f"><code>a1b9c9e</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1663">#1663</a> from ybrustin/master</li>
<li><a href="https://github.com/stretchr/testify/commit/8302de98b17649445fc1f1992fc3fecdb40c59ba"><code>8302de9</code></a> Merge branch 'master' into master</li>
<li><a href="https://github.com/stretchr/testify/commit/89352f7958086841c72425ccd6f43ab299e1309c"><code>89352f7</code></a> Merge pull request <a href="https://redirect.github.com/stretchr/testify/issues/1518">#1518</a> from hendrywiranto/adjust-readme-remove-v2</li>
<li>Additional commits viewable in <a href="https://github.com/stretchr/testify/compare/v1.9.0...v1.10.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.32.4 to 1.32.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d125de3792b20980da07dd1424afa90285d50c90"><code>d125de3</code></a> Release 2024-11-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fec51f3fff1b0f1a2cc913f4f7874977a3d47d0d"><code>fec51f3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fba59970453e163b476913a363a126412641b5fb"><code>fba5997</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b8e5c842fc161b15aa5e13c8a0caf1357b5e9c7"><code>0b8e5c8</code></a> Bump smithy-go dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2902">#2902</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50ba45ce162cd2458c65f5da799fd907ad826561"><code>50ba45c</code></a> Release 2024-11-15.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/444bdffccd6dce18f60ae626b74c087641c3d052"><code>444bdff</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55ab381b20235964ab0c670a29d096821e158e90"><code>55ab381</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94c083768b80bbd372c0e9feb45f02511442b896"><code>94c0837</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2398a7903ca141692f98da65b8537a4a53e9707e"><code>2398a79</code></a> Remove elastictranscoder service's integration test (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2901">#2901</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93e0f294f6c692e39adf1b8ec2c8681ba9ee5f01"><code>93e0f29</code></a> Release 2024-11-15</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.32.4...v1.32.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.45 to 1.17.46
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d125de3792b20980da07dd1424afa90285d50c90"><code>d125de3</code></a> Release 2024-11-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fec51f3fff1b0f1a2cc913f4f7874977a3d47d0d"><code>fec51f3</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/fba59970453e163b476913a363a126412641b5fb"><code>fba5997</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b8e5c842fc161b15aa5e13c8a0caf1357b5e9c7"><code>0b8e5c8</code></a> Bump smithy-go dependency (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2902">#2902</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/50ba45ce162cd2458c65f5da799fd907ad826561"><code>50ba45c</code></a> Release 2024-11-15.2</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/444bdffccd6dce18f60ae626b74c087641c3d052"><code>444bdff</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/55ab381b20235964ab0c670a29d096821e158e90"><code>55ab381</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/94c083768b80bbd372c0e9feb45f02511442b896"><code>94c0837</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2398a7903ca141692f98da65b8537a4a53e9707e"><code>2398a79</code></a> Remove elastictranscoder service's integration test (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2901">#2901</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/93e0f294f6c692e39adf1b8ec2c8681ba9ee5f01"><code>93e0f29</code></a> Release 2024-11-15</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.45...credentials/v1.17.46">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.43.2 to 1.44.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7720f87d0ec2f3609e2912c864b6f9ae550e0792"><code>7720f87</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4bd06b93f8627928c17604383bd8639f2cb23739"><code>4bd06b9</code></a> Merge customizations for service s3</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/24bd51460e581c55d4a2c05832f8d73da5de0833"><code>24bd514</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2373">#2373</a> from aws/deprecate-macie</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ce87b8413004cd20963ffc8f069963ae2e738e32"><code>ce87b84</code></a> Add changelog for last commit</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ea520b62d3cb5c276a39212b73bb2d2eb7f66cdd"><code>ea520b6</code></a> Remove macie service</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2a621e148ba47120f4829e8f5c5edbe97a3a9a2"><code>a2a621e</code></a> Release 2023-11-16</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5590a8748ea3f2733d883dae0d316ef3e1c2aec4"><code>5590a87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e2a7ad742f9e3b25fff66485cf9e5e74c088c1d5"><code>e2a7ad7</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b97a73ebbc56983a1897b25a784e66849952257d"><code>b97a73e</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/21c00906a39a7c8b5be5bc698591b508ae8c7f8d"><code>21c0090</code></a> fix trailing comment parse in properties (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2371">#2371</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.43.2...service/s3/v1.44.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/caliri/release/v4` from 4.4.19 to 4.5.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/caliri/commit/a4286373081ee5e61f697723443daaf6a62e43cf"><code>a428637</code></a> Update bindings to release/v4.5.0</li>
<li>See full diff in <a href="https://github.com/content-services/caliri/compare/release/v4.4.19...release/v4.5.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/content-services/zest/release/v2024` from 2024.11.1731697894 to 2024.11.1732298036
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/zest/commit/1562d6d20acd5a4db68f6076ed15b47a01ce23e7"><code>1562d6d</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b97933dd6233137ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/b1ff179d97c5d14db0ec83d110ac5329f05b4150"><code>b1ff179</code></a> Update pulp bindings to 9d5beb82a436bd89a92e4d8a965b73996e39be037d2e4353b86e2...</li>
<li><a href="https://github.com/content-services/zest/commit/927c9ed8cd7de76f716ed8efe7d02d0eb2e8f684"><code>927c9ed</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b979332ee326c37ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/56d9a0c562138d6e552495da2e2360d4754ec370"><code>56d9a0c</code></a> Update pulp bindings to 9258456e9d3b4a85d3e54836d2b9793a9955a3c37ae5486b92ad2...</li>
<li><a href="https://github.com/content-services/zest/commit/8f1498184b6f5dfd716226f5228e6b49c76a8e8f"><code>8f14981</code></a> Update pulp bindings to 689d598a363629d8a5469d9b2a847485d8bed2c37d356e3e82bb4...</li>
<li>See full diff in <a href="https://github.com/content-services/zest/compare/release/v2024.11.1731697894...release/v2024.11.1732298036">compare view</a></li>
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

### Comment by @jlsherrill on November 25, 2024 at 02:06 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on November 25, 2024 at 08:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/904*
