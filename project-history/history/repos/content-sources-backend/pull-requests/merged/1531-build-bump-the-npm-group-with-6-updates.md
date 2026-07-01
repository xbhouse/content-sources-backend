---
type: pull_request
number: 1531
title: "Build: Bump the npm group with 6 updates"
state: merged
author: dependabot
created: 2026-06-08T17:29:18Z
updated: 2026-06-08T18:36:49Z
closed: 2026-06-08T18:36:46Z
merged: 2026-06-08T18:36:46Z
base_branch: main
head_branch: dependabot/npm_and_yarn/npm-2298bd91cc
labels: ["dependencies", "javascript"]
url: https://github.com/content-services/content-sources-backend/pull/1531
---

# Pull Request #1531: Build: Bump the npm group with 6 updates

**Author**: @dependabot
**Created**: June 08, 2026 at 05:29 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `javascript`
**Base**: `main` ← **Head**: `dependabot/npm_and_yarn/npm-2298bd91cc`

## Description

Bumps the npm group with 6 updates:

| Package | From | To |
| --- | --- | --- |
| [@currents/playwright](https://github.com/currents-dev/currents-playwright-changelog) | `1.24.0` | `2.2.0` |
| [@eslint/js](https://github.com/eslint/eslint/tree/HEAD/packages/js) | `9.39.2` | `10.0.1` |
| [@typescript-eslint/eslint-plugin](https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/eslint-plugin) | `8.49.0` | `8.60.1` |
| [@typescript-eslint/parser](https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/parser) | `8.49.0` | `8.60.1` |
| [eslint](https://github.com/eslint/eslint) | `9.39.2` | `10.4.1` |
| [prettier](https://github.com/prettier/prettier) | `3.7.4` | `3.8.3` |

Updates `@currents/playwright` from 1.24.0 to 2.2.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/currents-dev/currents-playwright-changelog/blob/main/CHANGELOG.md">@​currents/playwright's changelog</a>.</em></p>
<blockquote>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.4...v2.2.0">2.2.0</a> (2026-06-02)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>bump jiti from 2.6.1 to 2.7.0 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/863">#863</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/7c074b241f6982a57f146ab85ae1d730bd5e1987">7c074b2</a>)</li>
<li>bump tmp from 0.2.6 to 0.2.7 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/865">#865</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/293915d1a47a526ffe06ac95dc341eb59afe8be7">293915d</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>use unidici and drop axios [ENG-568] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/836">#836</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/a3d7f72df8ebd7e851687cce6f72b668b0867a26">a3d7f72</a>), closes <a href="https://redirect.github.com/currents-dev/currents-playwright/issues/813">#813</a></li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.3...v2.1.4">2.1.4</a> (2026-05-29)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>restore git-preferred commit merge and clean CI branch names [ENG-563] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/859">#859</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/a6610719e7c0e807a4677add101284b5ec6e029b">a661071</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.2...v2.1.3">2.1.3</a> (2026-05-29)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>adopt CURRENTS_CI_URL as ciUrl config option [ENG-561] (<a href="https://github.com/currents-dev/currents-playwright/commit/ef9326b3dd3728834d3bb86422ee915afd45aa6a">ef9326b</a>)</li>
<li>bump tmp from 0.2.5 to 0.2.6 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/853">#853</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/73a765d42c7599c251289457753d5689075405b6">73a765d</a>)</li>
<li>bump tmp from 0.2.5 to 0.2.7 in /examples (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/852">#852</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/8beb700ca00b087bdee1882842965363373765d1">8beb700</a>)</li>
<li>bump tmp from 0.2.5 to 0.2.7 in /examples/imported-tests (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/855">#855</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/ba99ef2cb6b415177e76f708fd128f16dc7d1910">ba99ef2</a>)</li>
<li>update dependencies in package-lock.json and package.json (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/858">#858</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/52955430f2c09b0f59d82ac81907f1f300370bd7">5295543</a>)</li>
</ul>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.1...v2.1.2">2.1.2</a> (2026-05-26)</h2>
<h2><a href="https://github.com/currents-dev/currents-playwright/compare/v2.1.0...v2.1.1">2.1.1</a> (2026-05-22)</h2>
<h3>Bug Fixes</h3>
<ul>
<li>include the fully parallel flag into the accepted flags for pwc-p run. (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/840">#840</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/34615d4da4151d33c27e2efda0dfb3d5a9ec42a8">34615d4</a>)</li>
</ul>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v2.0.0...v2.1.0">2.1.0</a> (2026-05-22)</h1>
<h3>Features</h3>
<ul>
<li>add --pwc-environment option and discovery environment field [CSR-3613] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/833">#833</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/3164c335e3bb82db1794a65a8223442b5dae4718">3164c33</a>)</li>
<li>azure ci data [CSR-4184] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/816">#816</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/d452c276f53b21fc7b920ae199d10a27e96a4fc4">d452c27</a>)</li>
</ul>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v1.23.2...v2.0.0">2.0.0</a> (2026-05-19)</h1>
<h3>Breaking changes</h3>
<p>Breaking changes for orchestration (<code>pwc-p</code>) only.
See the upgrade guide at <a href="https://docs.currents.dev/resources/reporters/currents-playwright/migration-to-playwright-1.60">https://docs.currents.dev/resources/reporters/currents-playwright/migration-to-playwright-1.60</a></p>
<ul>
<li>Update uses of <code>pwc-p</code> with the run subcommand <code>pwc-p run</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/currents-dev/currents-playwright-changelog/commits">compare view</a></li>
</ul>
</details>
<br />

Updates `@eslint/js` from 9.39.2 to 10.0.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">@​eslint/js's releases</a>.</em></p>
<blockquote>
<h2>v10.0.1</h2>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/c87d5bded54c5cf491eb04c24c9d09bbbd42c23e"><code>c87d5bd</code></a> fix: update eslint (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20531">#20531</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/d84100115c14691691058f00779c94e74fca946a"><code>d841001</code></a> fix: update <code>minimatch</code> to <code>10.2.1</code> to address security vulnerabilities (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20519">#20519</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/04c21475b3004904948f02049f2888b401d82c78"><code>04c2147</code></a> fix: update error message for unused suppressions (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20496">#20496</a>) (fnx)</li>
<li><a href="https://github.com/eslint/eslint/commit/38b089c1726feac0e31a31d47941bd99e29ce003"><code>38b089c</code></a> fix: update dependency <code>@​eslint/config-array</code> to ^0.23.1 (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20484">#20484</a>) (renovate[bot])</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/5b3dbce50a1404a9f118afe810cefeee79388a2a"><code>5b3dbce</code></a> docs: add AI acknowledgement section to templates (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20431">#20431</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/6f23076037d5879f20fb3be2ef094293b1e8d38c"><code>6f23076</code></a> docs: toggle nav in no-JS mode (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20476">#20476</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/b69cfb32a16c5d5e9986390d484fae1d21e406f9"><code>b69cfb3</code></a> docs: Update README (GitHub Actions Bot)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/e5c281ffd038a3a7a3e5364db0b9378e0ad83020"><code>e5c281f</code></a> chore: updates for v9.39.3 release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/8c3832adb77cd993b4a24891900d5eeaaf093cdc"><code>8c3832a</code></a> chore: update <code>@​typescript-eslint/parser</code> to ^8.56.0 (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20514">#20514</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/8330d238ae6adb68bb6a1c9381e38cfedd990d94"><code>8330d23</code></a> test: add tests for config-api (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20493">#20493</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/37d6e91e88fa6a2ca6d8726679096acff21ba6cc"><code>37d6e91</code></a> chore: remove eslint v10 prereleases from eslint-config-eslint deps (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20494">#20494</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/da7cd0e79197ad16e17052eef99df141de6dbfb1"><code>da7cd0e</code></a> refactor: cleanup error message templates (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20479">#20479</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/84fb885d49ac810e79a9491276b4828b53d913e5"><code>84fb885</code></a> chore: package.json update for <code>@​eslint/js</code> release (Jenkins)</li>
<li><a href="https://github.com/eslint/eslint/commit/1f667344b57c4c09b548d94bcfac1f91b6e5c63d"><code>1f66734</code></a> chore: add <code>eslint</code> to <code>peerDependencies</code> of <code>@eslint/js</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20467">#20467</a>) (Milos Djermanovic)</li>
</ul>
<h2>v10.0.0</h2>
<h2>Breaking Changes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/f9e54f43a5e497cdfa179338b431093245cb787b"><code>f9e54f4</code></a> feat!: estimate rule-tester failure location (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20420">#20420</a>) (ST-DDT)</li>
<li><a href="https://github.com/eslint/eslint/commit/a176319d8ade1a7d9b2d7fb8f038f55a2662325f"><code>a176319</code></a> feat!: replace <code>chalk</code> with <code>styleText</code> and add <code>color</code> to <code>ResultsMeta</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20227">#20227</a>) (루밀LuMir)</li>
<li><a href="https://github.com/eslint/eslint/commit/c7046e6c1e03c4ca0eee4888a1f2eba4c6454f84"><code>c7046e6</code></a> feat!: enable JSX reference tracking (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20152">#20152</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/fa31a608901684fbcd9906d1907e66561d16e5aa"><code>fa31a60</code></a> feat!: add <code>name</code> to configs (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20015">#20015</a>) (Kirk Waiblinger)</li>
<li><a href="https://github.com/eslint/eslint/commit/3383e7ec9028166cafc8ea7986c2f7498d0049f0"><code>3383e7e</code></a> fix!: remove deprecated <code>SourceCode</code> methods (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20137">#20137</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/501abd0e916a35554c58b7c0365537f1fa3880ce"><code>501abd0</code></a> feat!: update dependency minimatch to v10 (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20246">#20246</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/ca4d3b40085de47561f89656a2207d09946ed45e"><code>ca4d3b4</code></a> fix!: stricter rule tester assertions for valid test cases (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20125">#20125</a>) (唯然)</li>
<li><a href="https://github.com/eslint/eslint/commit/96512a66c86402fb0538cdcb6cd30b9073f6bf3b"><code>96512a6</code></a> fix!: Remove deprecated rule context methods (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20086">#20086</a>) (Nicholas C. Zakas)</li>
<li><a href="https://github.com/eslint/eslint/commit/c69fdacdb2e886b9d965568a397aa8220db3fe90"><code>c69fdac</code></a> feat!: remove eslintrc support (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20037">#20037</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/208b5cc34a8374ff81412b5bec2e0800eebfbd04"><code>208b5cc</code></a> feat!: Use <code>ScopeManager#addGlobals()</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20132">#20132</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/a2ee188ea7a38a0c6155f3d39e2b00e1d0f36e14"><code>a2ee188</code></a> fix!: add <code>uniqueItems: true</code> in <code>no-invalid-regexp</code> option (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20155">#20155</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/a89059dbf2832d417dd493ee81483227ec44e4ab"><code>a89059d</code></a> feat!: Program range span entire source text (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20133">#20133</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/39a6424373d915fa9de0d7b0caba9a4dc3da9b53"><code>39a6424</code></a> fix!: assert 'text' is a string across all RuleFixer methods (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20082">#20082</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/f28fbf846244e043c92b355b224d121b06140b44"><code>f28fbf8</code></a> fix!: Deprecate <code>&quot;always&quot;</code> and <code>&quot;as-needed&quot;</code> options of the <code>radix</code> rule (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20223">#20223</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/aa3fb2b233e929b37220be940575f42c280e0b98"><code>aa3fb2b</code></a> fix!: tighten <code>func-names</code> schema (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20119">#20119</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/f6c0ed0311dcfee853367d5068c765d066e6b756"><code>f6c0ed0</code></a> feat!: report <code>eslint-env</code> comments as errors (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20128">#20128</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/4bf739fb533e59f7f0a66b65f7bc80be0f37d8db"><code>4bf739f</code></a> fix!: remove deprecated <code>LintMessage#nodeType</code> and <code>TestCaseError#type</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20096">#20096</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/523c076866400670fb2192a3f55dbf7ad3469247"><code>523c076</code></a> feat!: drop support for jiti &lt; 2.2.0 (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20016">#20016</a>) (michael faith)</li>
<li><a href="https://github.com/eslint/eslint/commit/454a292c95f34dad232411ddac06408e6383bb64"><code>454a292</code></a> feat!: update <code>eslint:recommended</code> configuration (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20210">#20210</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/4f880ee02992e1bf0e96ebaba679985e2d1295f1"><code>4f880ee</code></a> feat!: remove <code>v10_*</code> and inactive <code>unstable_*</code> flags (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20225">#20225</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/f18115c363a4ac7671a4c7f30ee13d57ebba330f"><code>f18115c</code></a> feat!: <code>no-shadow-restricted-names</code> report <code>globalThis</code> by default (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20027">#20027</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/c6358c31fbd3937b92d89be2618ffdf5a774604e"><code>c6358c3</code></a> feat!: Require Node.js <code>^20.19.0 || ^22.13.0 || &gt;=24</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20160">#20160</a>) (Milos Djermanovic)</li>
</ul>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/bff9091927811497dbf066b0e3b85ecb37d43822"><code>bff9091</code></a> feat: handle <code>Array.fromAsync</code> in <code>array-callback-return</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20457">#20457</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/290c594bb50c439fb71bc75521ee5360daa8c222"><code>290c594</code></a> feat: add <code>self</code> to <code>no-implied-eval</code> rule (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20468">#20468</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/43677de07ebd6e14bfac40a46ad749ba783c45f2"><code>43677de</code></a> feat: fix handling of function and class expression names in <code>no-shadow</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20432">#20432</a>) (Milos Djermanovic)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/84fb885d49ac810e79a9491276b4828b53d913e5"><code>84fb885</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/1f667344b57c4c09b548d94bcfac1f91b6e5c63d"><code>1f66734</code></a> chore: add <code>eslint</code> to <code>peerDependencies</code> of <code>@eslint/js</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20467">#20467</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/f3fbc2f60cbe2c718364feb8c3fc0452c0df3c56"><code>f3fbc2f</code></a> chore: set <code>@eslint/js</code> version to 10.0.0 to skip releasing it (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20466">#20466</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/b4b3127f8542c599ce2dea804b6582ebc40c993d"><code>b4b3127</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/0b14059491d830a49b3577931f4f68fbcfce6be5"><code>0b14059</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/fa31a608901684fbcd9906d1907e66561d16e5aa"><code>fa31a60</code></a> feat!: add <code>name</code> to configs (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20015">#20015</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/1e2cad5f6fa47ed6ed89d2a29798dda926d50990"><code>1e2cad5</code></a> chore: package.json update for <code>@​eslint/js</code> release</li>
<li><a href="https://github.com/eslint/eslint/commit/454a292c95f34dad232411ddac06408e6383bb64"><code>454a292</code></a> feat!: update <code>eslint:recommended</code> configuration (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20210">#20210</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/c6358c31fbd3937b92d89be2618ffdf5a774604e"><code>c6358c3</code></a> feat!: Require Node.js <code>^20.19.0 || ^22.13.0 || &gt;=24</code> (<a href="https://github.com/eslint/eslint/tree/HEAD/packages/js/issues/20160">#20160</a>)</li>
<li>See full diff in <a href="https://github.com/eslint/eslint/commits/v10.0.1/packages/js">compare view</a></li>
</ul>
</details>
<br />

Updates `@typescript-eslint/eslint-plugin` from 8.49.0 to 8.60.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/releases">@​typescript-eslint/eslint-plugin's releases</a>.</em></p>
<blockquote>
<h2>v8.60.1</h2>
<h2>8.60.1 (2026-06-01)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> respect ECMAScript line terminators in ts-comment rules (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12352">#12352</a>)</li>
<li><strong>eslint-plugin:</strong> [no-shadow] correct rule to match ESLint v10 handling (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12182">#12182</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>lumir</li>
<li>Nevette Bailey <a href="https://github.com/nevette-bailey"><code>@​nevette-bailey</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>v8.60.0</h2>
<h2>8.60.0 (2026-05-25)</h2>
<h3>🚀 Features</h3>
<ul>
<li><strong>rule-tester:</strong> added updates of RuleTester from upstream (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12291">#12291</a>)</li>
</ul>
<h3>🩹 Fixes</h3>
<ul>
<li>playground TS version selector is not working (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12326">#12326</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/12325">#12325</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Evyatar Daud <a href="https://github.com/StyleShit"><code>@​StyleShit</code></a></li>
<li>Vinccool96</li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.0">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>v8.59.4</h2>
<h2>8.59.4 (2026-05-18)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [no-floating-promises] stack overflow when using recursive types (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12294">#12294</a>)</li>
<li><strong>project-service:</strong> throw error cause in <code>getParsedConfigFileFromTSServer</code> (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12321">#12321</a>)</li>
<li><strong>typescript-eslint:</strong> export Compatible* types from typescript-eslint to resolve pnpm TS error (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12340">#12340</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Evyatar Daud <a href="https://github.com/StyleShit"><code>@​StyleShit</code></a></li>
<li>Kirk Waiblinger <a href="https://github.com/kirkwaiblinger"><code>@​kirkwaiblinger</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/eslint-plugin/CHANGELOG.md">@​typescript-eslint/eslint-plugin's changelog</a>.</em></p>
<blockquote>
<h2>8.60.1 (2026-06-01)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [no-shadow] correct rule to match ESLint v10 handling (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12182">#12182</a>)</li>
<li><strong>eslint-plugin:</strong> respect ECMAScript line terminators in ts-comment rules (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12352">#12352</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>lumir</li>
<li>Nevette Bailey <a href="https://github.com/nevette-bailey"><code>@​nevette-bailey</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.60.0 (2026-05-25)</h2>
<p>This was a version bump only for eslint-plugin to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.0">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.4 (2026-05-18)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [no-floating-promises] stack overflow when using recursive types (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12294">#12294</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Evyatar Daud <a href="https://github.com/StyleShit"><code>@​StyleShit</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.59.4">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.3 (2026-05-11)</h2>
<p>This was a version bump only for eslint-plugin to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.59.3">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.2 (2026-05-04)</h2>
<h3>🩹 Fixes</h3>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/4f84a697aedc436559c3ae09b5b357d98b448d68"><code>4f84a69</code></a> chore(release): publish 8.60.1</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/598af564db50593277ba46c7fdea3648e4425391"><code>598af56</code></a> docs(eslint-plugin): clarify no-redeclare type-value collision not covered by...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/1849b530c254fb4f89d7270160f3a998e4acd964"><code>1849b53</code></a> chore: typecheck using tsgo (<a href="https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/eslint-plugin/issues/12139">#12139</a>)</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/5341d59dd3c21fc4e2bf3bce55cf35d8f84e5216"><code>5341d59</code></a> chore: fix lint issues (<a href="https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/eslint-plugin/issues/12369">#12369</a>)</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/f525814f01766487ab34a54d56de20ea5c4fb576"><code>f525814</code></a> fix(eslint-plugin): [no-shadow] correct rule to match ESLint v10 handling (<a href="https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/eslint-plugin/issues/1">#1</a>...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/2df540cd8bbeb3e2c56d516912f69bf63c1e9450"><code>2df540c</code></a> chore(eslint-plugin): defer type checks to improve rules performance (<a href="https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/eslint-plugin/issues/12296">#12296</a>)</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/1ab4284789b36cf482a4d9924719162a02d54243"><code>1ab4284</code></a> fix(eslint-plugin): respect ECMAScript line terminators in ts-comment rules (...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/2f49df599b2db5f2937caf975d3c63e5cdeb0ea1"><code>2f49df5</code></a> docs: update references to <code>@stylistic/eslint-plugin</code> rules in documentation ...</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/f891c29de5f3e23f3d8c59cc599d3196e54e9b58"><code>f891c29</code></a> chore(release): publish 8.60.0</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/ca6ca1431b6d18235297a7e29feb5d98f012dff2"><code>ca6ca14</code></a> chore(release): publish 8.59.4</li>
<li>Additional commits viewable in <a href="https://github.com/typescript-eslint/typescript-eslint/commits/v8.60.1/packages/eslint-plugin">compare view</a></li>
</ul>
</details>
<br />

Updates `@typescript-eslint/parser` from 8.49.0 to 8.60.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/releases">@​typescript-eslint/parser's releases</a>.</em></p>
<blockquote>
<h2>v8.60.1</h2>
<h2>8.60.1 (2026-06-01)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> respect ECMAScript line terminators in ts-comment rules (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12352">#12352</a>)</li>
<li><strong>eslint-plugin:</strong> [no-shadow] correct rule to match ESLint v10 handling (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12182">#12182</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>lumir</li>
<li>Nevette Bailey <a href="https://github.com/nevette-bailey"><code>@​nevette-bailey</code></a></li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>v8.60.0</h2>
<h2>8.60.0 (2026-05-25)</h2>
<h3>🚀 Features</h3>
<ul>
<li><strong>rule-tester:</strong> added updates of RuleTester from upstream (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12291">#12291</a>)</li>
</ul>
<h3>🩹 Fixes</h3>
<ul>
<li>playground TS version selector is not working (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12326">#12326</a>, <a href="https://redirect.github.com/typescript-eslint/typescript-eslint/issues/12325">#12325</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Evyatar Daud <a href="https://github.com/StyleShit"><code>@​StyleShit</code></a></li>
<li>Vinccool96</li>
</ul>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.0">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>v8.59.4</h2>
<h2>8.59.4 (2026-05-18)</h2>
<h3>🩹 Fixes</h3>
<ul>
<li><strong>eslint-plugin:</strong> [no-floating-promises] stack overflow when using recursive types (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12294">#12294</a>)</li>
<li><strong>project-service:</strong> throw error cause in <code>getParsedConfigFileFromTSServer</code> (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12321">#12321</a>)</li>
<li><strong>typescript-eslint:</strong> export Compatible* types from typescript-eslint to resolve pnpm TS error (<a href="https://redirect.github.com/typescript-eslint/typescript-eslint/pull/12340">#12340</a>)</li>
</ul>
<h3>❤️ Thank You</h3>
<ul>
<li>Evyatar Daud <a href="https://github.com/StyleShit"><code>@​StyleShit</code></a></li>
<li>Kirk Waiblinger <a href="https://github.com/kirkwaiblinger"><code>@​kirkwaiblinger</code></a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/typescript-eslint/typescript-eslint/blob/main/packages/parser/CHANGELOG.md">@​typescript-eslint/parser's changelog</a>.</em></p>
<blockquote>
<h2>8.60.1 (2026-06-01)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.60.0 (2026-05-25)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.60.0">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.4 (2026-05-18)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.59.4">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.3 (2026-05-11)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.59.3">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.2 (2026-05-04)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.59.2">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.1 (2026-04-27)</h2>
<p>This was a version bump only for parser to align it with other projects, there were no code changes.</p>
<p>See <a href="https://github.com/typescript-eslint/typescript-eslint/releases/tag/v8.59.1">GitHub Releases</a> for more information.</p>
<p>You can read about our <a href="https://typescript-eslint.io/users/versioning">versioning strategy</a> and <a href="https://typescript-eslint.io/users/releases">releases</a> on our website.</p>
<h2>8.59.0 (2026-04-20)</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/4f84a697aedc436559c3ae09b5b357d98b448d68"><code>4f84a69</code></a> chore(release): publish 8.60.1</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/1849b530c254fb4f89d7270160f3a998e4acd964"><code>1849b53</code></a> chore: typecheck using tsgo (<a href="https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/parser/issues/12139">#12139</a>)</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/f891c29de5f3e23f3d8c59cc599d3196e54e9b58"><code>f891c29</code></a> chore(release): publish 8.60.0</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/ca6ca1431b6d18235297a7e29feb5d98f012dff2"><code>ca6ca14</code></a> chore(release): publish 8.59.4</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/48e13c0261e3cb1bf4f4dfaa462cdb3a56ef7383"><code>48e13c0</code></a> chore(release): publish 8.59.3</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/44f9625336841a8ee3eb01a9e02e49b1d7b12648"><code>44f9625</code></a> chore(deps): update vitest monorepo to v4.1.5 (<a href="https://github.com/typescript-eslint/typescript-eslint/tree/HEAD/packages/parser/issues/12307">#12307</a>)</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/2ec35f1760aade4df4c631d76d78c7ed5e136333"><code>2ec35f1</code></a> chore(release): publish 8.59.2</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/52457932e5507b5ca01e720a541f3f8d01e09b9d"><code>5245793</code></a> chore(release): publish 8.59.1</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/ea9ae4f8817873480e3501145059f63e39e8d8a1"><code>ea9ae4f</code></a> chore(release): publish 8.59.0</li>
<li><a href="https://github.com/typescript-eslint/typescript-eslint/commit/90c2803a4c250e0343598d41e973f95e743bf4ce"><code>90c2803</code></a> chore(release): publish 8.58.2</li>
<li>Additional commits viewable in <a href="https://github.com/typescript-eslint/typescript-eslint/commits/v8.60.1/packages/parser">compare view</a></li>
</ul>
</details>
<br />

Updates `eslint` from 9.39.2 to 10.4.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/eslint/eslint/releases">eslint's releases</a>.</em></p>
<blockquote>
<h2>v10.4.1</h2>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/e557467db7496220eebcbe2ac5ea6d38c12bb1ec"><code>e557467</code></a> fix: update <code>@eslint/plugin-kit</code> version to 0.7.2 (<a href="https://redirect.github.com/eslint/eslint/issues/20930">#20930</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/d4ce898796ca22c3b96aa70d3014cb85f4bac1cd"><code>d4ce898</code></a> fix: propagate failures from delegated commands (<a href="https://redirect.github.com/eslint/eslint/issues/20917">#20917</a>) (Minh Vu)</li>
<li><a href="https://github.com/eslint/eslint/commit/f4f3507460bc016b5be979c05d2969793f570cbf"><code>f4f3507</code></a> fix: prefer-arrow-callback invalid autofix with newline after <code>async</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20916">#20916</a>) (kuldeep kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/c5bc78b37e08b9054a11f0cc2d81808bb24acb85"><code>c5bc78b</code></a> fix: false positive for reference in <code>finally</code> block (<a href="https://redirect.github.com/eslint/eslint/issues/20655">#20655</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/27538c01f5df4e9306f6f4ba867b2dd6307fae59"><code>27538c0</code></a> fix: add missing CodePath and CodePathSegment types (<a href="https://redirect.github.com/eslint/eslint/issues/20853">#20853</a>) (Pixel998)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/61b0add61ffc52665562be7bb96f526690a78b30"><code>61b0add</code></a> docs: remove deprecated rule from related rules of <code>max-params</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20921">#20921</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/305d5b91aeac24d36fde42f75625a8f183d4ce43"><code>305d5b9</code></a> docs: remove deprecated rules from related rules section (<a href="https://redirect.github.com/eslint/eslint/issues/20911">#20911</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/49b0202d01918b8061720d586dffd7c68047090c"><code>49b0202</code></a> docs: fix <code>display: none</code> of ad (<a href="https://redirect.github.com/eslint/eslint/issues/20901">#20901</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/9067f9492ec998afc5b4f057a477ecf6ebd45e44"><code>9067f94</code></a> docs: switch build to Node.js 24 (<a href="https://redirect.github.com/eslint/eslint/issues/20893">#20893</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/c91b0417e3420c76807ce1fa2aea76e2de87ab86"><code>c91b041</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/e349265cb37f3ebc837e178e48a725bb782bd870"><code>e349265</code></a> docs: clarify semver strings in rule deprecation objects (<a href="https://redirect.github.com/eslint/eslint/issues/20885">#20885</a>) (Milos Djermanovic)</li>
</ul>
<h2>Chores</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/b0e466b6ab47bfc7de43d8de0c315d8ee83aa584"><code>b0e466b</code></a> test: add <code>data</code> property to invalid tests cases for rules (<a href="https://redirect.github.com/eslint/eslint/issues/20924">#20924</a>) (Tanuj Kanti)</li>
<li><a href="https://github.com/eslint/eslint/commit/f78838bc4c86d487e1bcc7cede260c4467721c46"><code>f78838b</code></a> test: add CodePath type coverage (<a href="https://redirect.github.com/eslint/eslint/issues/20904">#20904</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/1daa4bd734b79a62e317d0394394a6b38cff49f9"><code>1daa4bd</code></a> chore: update <code>eslint-plugin-eslint-comments</code> test data to latest commit (<a href="https://redirect.github.com/eslint/eslint/issues/20922">#20922</a>) (Francesco Trotta)</li>
<li><a href="https://github.com/eslint/eslint/commit/002942ce988ea28b78e0a2f3b074081e638b552c"><code>002942c</code></a> ci: declare contents:read on update-readme workflow (<a href="https://redirect.github.com/eslint/eslint/issues/20919">#20919</a>) (Arpit Jain)</li>
<li><a href="https://github.com/eslint/eslint/commit/64bca24e7bed35bc3c864fc625cb2d89eca87d5b"><code>64bca24</code></a> chore: update ecosystem plugins (<a href="https://redirect.github.com/eslint/eslint/issues/20912">#20912</a>) (ESLint Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/6d7c832950d5e92499d88e504080661f888f8f56"><code>6d7c832</code></a> chore: ignore fflate updates in renovate (<a href="https://redirect.github.com/eslint/eslint/issues/20908">#20908</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/b2c86382164d87c6203b78d52068cd6a2a6ffe30"><code>b2c8638</code></a> ci: bump pnpm/action-setup from 6.0.7 to 6.0.8 (<a href="https://redirect.github.com/eslint/eslint/issues/20889">#20889</a>) (dependabot[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/a9b8d7f74c50211701cfc49710fa541fd91b2aa5"><code>a9b8d7f</code></a> chore: increase maxBuffer for ecosystem tests (<a href="https://redirect.github.com/eslint/eslint/issues/20881">#20881</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/b702ead5e1ed7cb9f28238a454797662efb37396"><code>b702ead</code></a> chore: update ecosystem update PR settings (<a href="https://redirect.github.com/eslint/eslint/issues/20884">#20884</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/507f60e9a78c9a902bc8759f066ae17a1ea6cd81"><code>507f60e</code></a> chore: update ecosystem plugins (<a href="https://redirect.github.com/eslint/eslint/issues/20882">#20882</a>) (ESLint Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/92f5c5bb6bf3a5d167c8ee53a430833410295c6d"><code>92f5c5b</code></a> test: add unit test for message-count (<a href="https://redirect.github.com/eslint/eslint/issues/20878">#20878</a>) (kuldeep kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/df321080af5758b1fa25e4b9a40e26135642dd6e"><code>df32108</code></a> chore: add <code>@​eslint/markdown</code> and typescript-eslint ecosystem tests (<a href="https://redirect.github.com/eslint/eslint/issues/20837">#20837</a>) (sethamus)</li>
<li><a href="https://github.com/eslint/eslint/commit/327f91d36aa49f2a50ded931d841a16374fd875f"><code>327f91d</code></a> chore: use includeIgnoreFile internally (<a href="https://redirect.github.com/eslint/eslint/issues/20876">#20876</a>) (Kirk Waiblinger)</li>
<li><a href="https://github.com/eslint/eslint/commit/f0dc4bd893fb3a9f44e4ddc3ad7063ffb0beacd3"><code>f0dc4bd</code></a> chore: pin fflate@0.8.2 (<a href="https://redirect.github.com/eslint/eslint/issues/20877">#20877</a>) (Milos Djermanovic)</li>
<li><a href="https://github.com/eslint/eslint/commit/0f4bd257a67a082b756de746d9e0c4842ab764ca"><code>0f4bd25</code></a> ci: run Discord alert for ecosystem test failures (<a href="https://redirect.github.com/eslint/eslint/issues/20873">#20873</a>) (Copilot)</li>
</ul>
<h2>v10.4.0</h2>
<h2>Features</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/1a45ec596af1dd5f880e6874cb8f24dafb6a7ecf"><code>1a45ec5</code></a> feat: check sequence expressions in <code>for-direction</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20701">#20701</a>) (kuldeep kumar)</li>
<li><a href="https://github.com/eslint/eslint/commit/450040bd89b989b3531824c6be45feb5fe3d936b"><code>450040b</code></a> feat: add <code>includeIgnoreFile()</code> to <code>eslint/config</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20735">#20735</a>) (Kirk Waiblinger)</li>
</ul>
<h2>Bug Fixes</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/544c0c3da589166ad8e5d634f35d3d06701c57be"><code>544c0c3</code></a> fix: escape code path DOT labels in debug output (<a href="https://redirect.github.com/eslint/eslint/issues/20866">#20866</a>) (Pixel998)</li>
<li><a href="https://github.com/eslint/eslint/commit/6799431203f2579632d0870f98ba132067f4040c"><code>6799431</code></a> fix: update dependency <code>@​eslint/config-helpers</code> to ^0.6.0 (<a href="https://redirect.github.com/eslint/eslint/issues/20850">#20850</a>) (renovate[bot])</li>
<li><a href="https://github.com/eslint/eslint/commit/f078fef5005dceb14fc162aab7c7200e027688dd"><code>f078fef</code></a> fix: handle non-array deprecated rule replacements (<a href="https://redirect.github.com/eslint/eslint/issues/20825">#20825</a>) (xbinaryx)</li>
</ul>
<h2>Documentation</h2>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/7e52a7151fb92eec0e0f67fe4e5ddbd1ccce796f"><code>7e52a71</code></a> docs: add mention of <code>@eslint-react/eslint-plugin</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20869">#20869</a>) (Pavel)</li>
<li><a href="https://github.com/eslint/eslint/commit/db3468ba746407d7f286f18f7ea9db6df0e3bc08"><code>db3468b</code></a> docs: tweak wording around ambiguous CJS-vs-ESM config (<a href="https://redirect.github.com/eslint/eslint/issues/20865">#20865</a>) (Kirk Waiblinger)</li>
<li><a href="https://github.com/eslint/eslint/commit/90846643ec6e97d447ae0d831fabe6d17b0a998a"><code>9084664</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/9cc73875046e3c4b8313644cbb1e99e26b36bd3f"><code>9cc7387</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/3d7b5484407403817aa9071a394d336d8ea96eb5"><code>3d7b548</code></a> docs: Update README (GitHub Actions Bot)</li>
<li><a href="https://github.com/eslint/eslint/commit/191ec3c0a3f94ce0f110df761f0b2b8949011ccb"><code>191ec3c</code></a> docs: Update README (GitHub Actions Bot)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/eslint/eslint/commit/4a3d15a99c452c4db2fd56b577fa7597e98ab0c2"><code>4a3d15a</code></a> 10.4.1</li>
<li><a href="https://github.com/eslint/eslint/commit/43e7e2bdc5c6cacc535446b7d23c10f780384ba8"><code>43e7e2b</code></a> Build: changelog update for 10.4.1</li>
<li><a href="https://github.com/eslint/eslint/commit/e557467db7496220eebcbe2ac5ea6d38c12bb1ec"><code>e557467</code></a> fix: update <code>@eslint/plugin-kit</code> version to 0.7.2 (<a href="https://redirect.github.com/eslint/eslint/issues/20930">#20930</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/b0e466b6ab47bfc7de43d8de0c315d8ee83aa584"><code>b0e466b</code></a> test: add <code>data</code> property to invalid tests cases for rules (<a href="https://redirect.github.com/eslint/eslint/issues/20924">#20924</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/d4ce898796ca22c3b96aa70d3014cb85f4bac1cd"><code>d4ce898</code></a> fix: propagate failures from delegated commands (<a href="https://redirect.github.com/eslint/eslint/issues/20917">#20917</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/f4f3507460bc016b5be979c05d2969793f570cbf"><code>f4f3507</code></a> fix: prefer-arrow-callback invalid autofix with newline after <code>async</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20916">#20916</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/f78838bc4c86d487e1bcc7cede260c4467721c46"><code>f78838b</code></a> test: add CodePath type coverage (<a href="https://redirect.github.com/eslint/eslint/issues/20904">#20904</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/61b0add61ffc52665562be7bb96f526690a78b30"><code>61b0add</code></a> docs: remove deprecated rule from related rules of <code>max-params</code> (<a href="https://redirect.github.com/eslint/eslint/issues/20921">#20921</a>)</li>
<li><a href="https://github.com/eslint/eslint/commit/1daa4bd734b79a62e317d0394394a6b38cff49f9"><code>1daa4bd</code></a> chore: update <code>eslint-plugin-eslint-comments</code> test data to latest commit (<a href="https://redirect.github.com/eslint/eslint/issues/20">#20</a>...</li>
<li><a href="https://github.com/eslint/eslint/commit/002942ce988ea28b78e0a2f3b074081e638b552c"><code>002942c</code></a> ci: declare contents:read on update-readme workflow (<a href="https://redirect.github.com/eslint/eslint/issues/20919">#20919</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/eslint/eslint/compare/v9.39.2...v10.4.1">compare view</a></li>
</ul>
</details>
<br />

Updates `prettier` from 3.7.4 to 3.8.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prettier/prettier/releases">prettier's releases</a>.</em></p>
<blockquote>
<h2>3.8.3</h2>
<ul>
<li>SCSS: Prevent trailing comma in <code>if()</code> function (<a href="https://redirect.github.com/prettier/prettier/pull/18471">prettier/prettier#18471</a> by <a href="https://github.com/kovsu"><code>@​kovsu</code></a>)</li>
</ul>
<p>🔗 <a href="https://github.com/prettier/prettier/blob/3.8.3/CHANGELOG.md#383">Changelog</a></p>
<h2>3.8.2</h2>
<ul>
<li>Support Angular v21.2</li>
</ul>
<p>🔗 <a href="https://github.com/prettier/prettier/blob/main/CHANGELOG.md#382">Changelog</a></p>
<h2>3.8.1</h2>
<ul>
<li>Include available <code>printers</code> in plugin type declarations (<a href="https://redirect.github.com/prettier/prettier/pull/18706">#18706</a> by <a href="https://github.com/porada"><code>@​porada</code></a>)</li>
</ul>
<p>🔗 <a href="https://github.com/prettier/prettier/blob/main/CHANGELOG.md#381">Changelog</a></p>
<h2>3.8.0</h2>
<ul>
<li>Support Angular v21.1</li>
</ul>
<p><a href="https://github.com/prettier/prettier/compare/3.7.4...3.8.0">diff</a></p>
<p>🔗 <a href="https://prettier.io/blog/2026/01/14/3.8.0">Release note &quot;Prettier 3.8: Support for Angular v21.1&quot;</a></p>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prettier/prettier/blob/main/CHANGELOG.md">prettier's changelog</a>.</em></p>
<blockquote>
<h1>3.8.3</h1>
<p><a href="https://github.com/prettier/prettier/compare/3.8.2...3.8.3">diff</a></p>
<h4>SCSS: Prevent trailing comma in <code>if()</code> function (<a href="https://redirect.github.com/prettier/prettier/pull/18471">#18471</a> by <a href="https://github.com/kovsu"><code>@​kovsu</code></a>)</h4>
<!-- raw HTML omitted -->
<pre lang="scss"><code>// Input
$value: if(sass(false): 1; else: -1);
<p>// Prettier 3.8.2
$value: if(
sass(false): 1; else: -1,
);</p>
<p>// Prettier 3.8.3
$value: if(sass(false): 1; else: -1);
</code></pre></p>
<h1>3.8.2</h1>
<p><a href="https://github.com/prettier/prettier/compare/3.8.1...3.8.2">diff</a></p>
<h4>Angular: Support Angular v21.2 (<a href="https://redirect.github.com/prettier/prettier/pull/18722">#18722</a>, <a href="https://redirect.github.com/prettier/prettier/pull/19034">#19034</a> by <a href="https://github.com/fisker"><code>@​fisker</code></a>)</h4>
<p>Exhaustive typechecking with <code>@default never;</code></p>
<!-- raw HTML omitted -->
<pre lang="html"><code>&lt;!-- Input --&gt;
@switch (foo) {
  @case (1) {}
  @default never;
}
<p>&lt;!-- Prettier 3.8.1 --&gt;
SyntaxError: Incomplete block &quot;default never&quot;. If you meant to write the @ character, you should use the &quot;&amp;<a href="https://redirect.github.com/prettier/prettier/issues/64">#64</a>;&quot; HTML entity instead. (3:3)</p>
<p>&lt;!-- Prettier 3.8.2 --&gt;
<a href="https://github.com/switch"><code>@​switch</code></a> (foo) {
<a href="https://github.com/case"><code>@​case</code></a> (1) {}
<a href="https://github.com/default"><code>@​default</code></a> never;
}
</code></pre></p>
<p><code>arrow function</code> and <code>instanceof</code> expressions.</p>
<!-- raw HTML omitted -->
<pre lang="html"><code>&lt;/tr&gt;&lt;/table&gt; 
</code></pre>
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prettier/prettier/commit/d7108a79ec745c04292aabf22c4c1adbd690b191"><code>d7108a7</code></a> Release 3.8.3</li>
<li><a href="https://github.com/prettier/prettier/commit/177f90898170d363ef64fde663e4d13170688bfe"><code>177f908</code></a> Prevent trailing comma in SCSS <code>if()</code> function (<a href="https://redirect.github.com/prettier/prettier/issues/18471">#18471</a>)</li>
<li><a href="https://github.com/prettier/prettier/commit/1cd40668c3d6f2f4cf9d87bbc9096d92361b2606"><code>1cd4066</code></a> Release <code>@​prettier/plugin-oxc</code><a href="https://github.com/0"><code>@​0</code></a>.1.4</li>
<li><a href="https://github.com/prettier/prettier/commit/a8700e245038cd8cc0cf28ef06ffedbcb3fc2dfc"><code>a8700e2</code></a> Update oxc-parser to v0.125.0</li>
<li><a href="https://github.com/prettier/prettier/commit/752157c78eca6f0a30e5d5cb513b682c5ecfa01e"><code>752157c</code></a> Fix tests</li>
<li><a href="https://github.com/prettier/prettier/commit/053fd418e180b12fa2014260212fae831f5fc5ec"><code>053fd41</code></a> Bump Prettier dependency to 3.8.2</li>
<li><a href="https://github.com/prettier/prettier/commit/904c6365ec46726fd0e21021c52ae934b7e5abc6"><code>904c636</code></a> Clean changelog_unreleased</li>
<li><a href="https://github.com/prettier/prettier/commit/dc1f7fcc508d116cbf1644d69a1f0eb93e40d4a4"><code>dc1f7fc</code></a> Update dependents count</li>
<li><a href="https://github.com/prettier/prettier/commit/b31557cf331a02acf83e7e29d1001b070189a0d9"><code>b31557c</code></a> Release 3.8.2</li>
<li><a href="https://github.com/prettier/prettier/commit/96bbaeda0525bf758e464aed2f939d739a85c315"><code>96bbaed</code></a> Support Angular v21.2 (<a href="https://redirect.github.com/prettier/prettier/issues/18722">#18722</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prettier/prettier/compare/3.7.4...3.8.3">compare view</a></li>
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
- `@dependabot show <dependency name> ignore conditions` will show all of the ignore conditions of the specified dependency
- `@dependabot ignore <dependency name> major version` will close this group update PR and stop Dependabot creating any more for the specific dependency's major version (unless you unignore this specific dependency's major version or upgrade to it yourself)
- `@dependabot ignore <dependency name> minor version` will close this group update PR and stop Dependabot creating any more for the specific dependency's minor version (unless you unignore this specific dependency's minor version or upgrade to it yourself)
- `@dependabot ignore <dependency name>` will close this group update PR and stop Dependabot creating any more for the specific dependency (unless you unignore this specific dependency or upgrade to it yourself)
- `@dependabot unignore <dependency name>` will remove all of the ignore conditions of the specified dependency
- `@dependabot unignore <dependency name> <ignore condition>` will remove the ignore condition of the specified dependency and ignore conditions


</details>

---

## Reviews

### Review by @swadeley - Approved on June 08, 2026 at 06:23 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1531*
