---
type: pull_request
number: 559
title: "Build: Bump the go group with 2 updates"
state: closed
author: dependabot
created: 2024-02-02T16:01:19Z
updated: 2024-02-05T04:10:23Z
closed: 2024-02-05T04:10:22Z
base_branch: main
head_branch: dependabot/go_modules/go-20e5fc5382
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/559
---

# Pull Request #559: Build: Bump the go group with 2 updates

**Author**: @dependabot
**Created**: February 02, 2024 at 04:01 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-20e5fc5382`

## Description

Bumps the go group with 2 updates: [gorm.io/gorm](https://github.com/go-gorm/gorm) and [github.com/IBM/sarama](https://github.com/IBM/sarama).

Updates `gorm.io/gorm` from 1.25.5 to 1.25.6
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/go-gorm/gorm/commit/0123dd45094295fade41e13550cd305eb5e3a848"><code>0123dd4</code></a> fix: ignore .gen.go suffix in logger to get the real caller when using gen <a href="https://redirect.github.com/go-gorm/gorm/issues/6">#6</a>...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/940358e0ddd1e6d5e893184943a4fa71a0c28521"><code>940358e</code></a> Fix tests doesn't follow <a href="https://gorm.io/docs/method_chaining.html">https://gorm.io/docs/method_chaining.html</a> convention</li>
<li><a href="https://github.com/go-gorm/gorm/commit/87decced23be0ce21929fe393fc4fa3a936b1ec8"><code>87decce</code></a> fix: ExplainSQL using consecutive pairs of escaper in SQL string represents a...</li>
<li><a href="https://github.com/go-gorm/gorm/commit/436cca753cd784969a19477f022db4eb3d84f2ec"><code>436cca7</code></a> fix: join and select mytable.* not working (<a href="https://redirect.github.com/go-gorm/gorm/issues/6761">#6761</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/a2cac75218c9844d0b43832e2d2a3c35f9700406"><code>a2cac75</code></a> feature: bring custom type and id column name to polymorphism (<a href="https://redirect.github.com/go-gorm/gorm/issues/6716">#6716</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/b9ebdb13c777029c03aa5c6c1581f579cdf1f79a"><code>b9ebdb1</code></a> Making locking parameters more intuitive (<a href="https://redirect.github.com/go-gorm/gorm/issues/6719">#6719</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/2fb4928aa873bef89d016e0af2c0724b44725043"><code>2fb4928</code></a> refactor: Resolve implicit memory aliasing in for loop (<a href="https://redirect.github.com/go-gorm/gorm/issues/6730">#6730</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/f0af94cd167c03af78b80586f7628bc8fb470577"><code>f0af94c</code></a> add test to show that update from works</li>
<li><a href="https://github.com/go-gorm/gorm/commit/3207ad6033aad5e76c6c9d578ef663032765e484"><code>3207ad6</code></a> map insert support return increment id (<a href="https://redirect.github.com/go-gorm/gorm/issues/6662">#6662</a>)</li>
<li><a href="https://github.com/go-gorm/gorm/commit/c1e911f6ed8d3d929aebbd39985a33c9ebe3bad7"><code>c1e911f</code></a> Update tests/go.mod</li>
<li>Additional commits viewable in <a href="https://github.com/go-gorm/gorm/compare/v1.25.5...v1.25.6">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/IBM/sarama` from 1.40.1 to 1.42.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/releases">github.com/IBM/sarama's releases</a>.</em></p>
<blockquote>
<h2>Version 1.42.1 (2023-11-07)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: make fetchInitialOffset use correct protocol by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2705">IBM/sarama#2705</a></li>
<li>fix(config): relax ClientID validation after 1.0.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2706">IBM/sarama#2706</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1">https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1</a></p>
<h2>Version 1.42.0 (2023-11-02)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>Asynchronously close brokers during a RefreshBrokers by <a href="https://github.com/bmassemin"><code>@​bmassemin</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2693">IBM/sarama#2693</a></li>
<li>Fix data race on Broker.done channel by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2698">IBM/sarama#2698</a></li>
<li>fix: data race in Broker.AsyncProduce by <a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2678">IBM/sarama#2678</a></li>
<li>Fix default retention time value in offset commit by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2700">IBM/sarama#2700</a></li>
<li>fix(txmgr): ErrOffsetsLoadInProgress is retriable by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2701">IBM/sarama#2701</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(ci): improve ossf scorecard result by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2685">IBM/sarama#2685</a></li>
<li>chore(ci): add kafka 3.6.0 to FVT and versions by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2692">IBM/sarama#2692</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>chore(ci): ossf scorecard.yml by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2683">IBM/sarama#2683</a></li>
<li>fix(ci): always run CodeQL on every commit by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2689">IBM/sarama#2689</a></li>
<li>chore(doc): add OpenSSF Scorecard badge by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2691">IBM/sarama#2691</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/bmassemin"><code>@​bmassemin</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2693">IBM/sarama#2693</a></li>
<li><a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2678">IBM/sarama#2678</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.3...v1.42.0">https://github.com/IBM/sarama/compare/v1.41.3...v1.42.0</a></p>
<h2>Version 1.41.3 (2023-10-17)</h2>
<!-- raw HTML omitted -->
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: pre-compile regex for parsing kafka version by <a href="https://github.com/qshuai"><code>@​qshuai</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2663">IBM/sarama#2663</a></li>
<li>fix(client): ignore empty Metadata responses when refreshing by <a href="https://github.com/HaoSunUber"><code>@​HaoSunUber</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2672">IBM/sarama#2672</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump the golang-org-x group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2661">IBM/sarama#2661</a></li>
<li>chore(deps): bump golang.org/x/net from 0.16.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2671">IBM/sarama#2671</a></li>
</ul>
<h3>:memo: Documentation</h3>
<ul>
<li>fix(docs): correct topic name in rebalancing strategy example by <a href="https://github.com/maksadbek"><code>@​maksadbek</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2657">IBM/sarama#2657</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/maksadbek"><code>@​maksadbek</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2657">IBM/sarama#2657</a></li>
<li><a href="https://github.com/qshuai"><code>@​qshuai</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2663">IBM/sarama#2663</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3">https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/IBM/sarama/blob/main/CHANGELOG.md">github.com/IBM/sarama's changelog</a>.</em></p>
<blockquote>
<h2>Version 1.42.1 (2023-11-07)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: make fetchInitialOffset use correct protocol by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2705">IBM/sarama#2705</a></li>
<li>fix(config): relax ClientID validation after 1.0.0 by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2706">IBM/sarama#2706</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1">https://github.com/IBM/sarama/compare/v1.42.0...v1.42.1</a></p>
<h2>Version 1.42.0 (2023-11-02)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>Asynchronously close brokers during a RefreshBrokers by <a href="https://github.com/bmassemin"><code>@​bmassemin</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2693">IBM/sarama#2693</a></li>
<li>Fix data race on Broker.done channel by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2698">IBM/sarama#2698</a></li>
<li>fix: data race in Broker.AsyncProduce by <a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2678">IBM/sarama#2678</a></li>
<li>Fix default retention time value in offset commit by <a href="https://github.com/prestona"><code>@​prestona</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2700">IBM/sarama#2700</a></li>
<li>fix(txmgr): ErrOffsetsLoadInProgress is retriable by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2701">IBM/sarama#2701</a></li>
</ul>
<h3>:wrench: Maintenance</h3>
<ul>
<li>chore(ci): improve ossf scorecard result by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2685">IBM/sarama#2685</a></li>
<li>chore(ci): add kafka 3.6.0 to FVT and versions by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2692">IBM/sarama#2692</a></li>
</ul>
<h3>:heavy_plus_sign: Other Changes</h3>
<ul>
<li>chore(ci): ossf scorecard.yml by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2683">IBM/sarama#2683</a></li>
<li>fix(ci): always run CodeQL on every commit by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2689">IBM/sarama#2689</a></li>
<li>chore(doc): add OpenSSF Scorecard badge by <a href="https://github.com/dnwe"><code>@​dnwe</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2691">IBM/sarama#2691</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/bmassemin"><code>@​bmassemin</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2693">IBM/sarama#2693</a></li>
<li><a href="https://github.com/lzakharov"><code>@​lzakharov</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2678">IBM/sarama#2678</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.3...v1.42.0">https://github.com/IBM/sarama/compare/v1.41.3...v1.42.0</a></p>
<h2>Version 1.41.3 (2023-10-17)</h2>
<h2>What's Changed</h2>
<h3>:bug: Fixes</h3>
<ul>
<li>fix: pre-compile regex for parsing kafka version by <a href="https://github.com/qshuai"><code>@​qshuai</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2663">IBM/sarama#2663</a></li>
<li>fix(client): ignore empty Metadata responses when refreshing by <a href="https://github.com/HaoSunUber"><code>@​HaoSunUber</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2672">IBM/sarama#2672</a></li>
</ul>
<h3>:package: Dependency updates</h3>
<ul>
<li>chore(deps): bump the golang-org-x group with 2 updates by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2661">IBM/sarama#2661</a></li>
<li>chore(deps): bump golang.org/x/net from 0.16.0 to 0.17.0 by <a href="https://github.com/dependabot"><code>@​dependabot</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2671">IBM/sarama#2671</a></li>
</ul>
<h3>:memo: Documentation</h3>
<ul>
<li>fix(docs): correct topic name in rebalancing strategy example by <a href="https://github.com/maksadbek"><code>@​maksadbek</code></a> in <a href="https://redirect.github.com/IBM/sarama/pull/2657">IBM/sarama#2657</a></li>
</ul>
<h2>New Contributors</h2>
<ul>
<li><a href="https://github.com/maksadbek"><code>@​maksadbek</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2657">IBM/sarama#2657</a></li>
<li><a href="https://github.com/qshuai"><code>@​qshuai</code></a> made their first contribution in <a href="https://redirect.github.com/IBM/sarama/pull/2663">IBM/sarama#2663</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3">https://github.com/IBM/sarama/compare/v1.41.2...v1.41.3</a></p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/IBM/sarama/commit/d88a48a6a6fd6cf1bd89e76192216378379fe799"><code>d88a48a</code></a> chore: update CHANGELOG.md to v1.42.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2711">#2711</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/385b3b435e85e6e3d1e29299b49330a0e9954f98"><code>385b3b4</code></a> fix(config): relax ClientID validation after 1.0.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2706">#2706</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/3364ff02d7e986ba0eb75d1c2dbb9a62a614dd15"><code>3364ff0</code></a> chore(doc): add CODE_OF_CONDUCT.md</li>
<li><a href="https://github.com/IBM/sarama/commit/768496e6e6f2e9de5cfb2d1fca3534a6bb737f92"><code>768496e</code></a> chore(ci): bump actions/dependency-review-action from 3.1.0 to 3.1.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2707">#2707</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/27710af1531a2f9a57dc21d4291441be6df274c6"><code>27710af</code></a> fix: make fetchInitialOffset use correct protocol (<a href="https://redirect.github.com/IBM/sarama/issues/2705">#2705</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/a46917f5b87dea431c327b5725e37bd843a2dc7a"><code>a46917f</code></a> chore(ci): bump actions/dependency-review-action from 2.5.1 to 3.1.0 (<a href="https://redirect.github.com/IBM/sarama/issues/2702">#2702</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/4168f7ccba85976c3b03fe01d3d693c876c617a2"><code>4168f7c</code></a> chore(ci): bump ossf/scorecard-action from 2.1.2 to 2.3.1 (<a href="https://redirect.github.com/IBM/sarama/issues/2703">#2703</a>)</li>
<li><a href="https://github.com/IBM/sarama/commit/7155d51d84883025253f297233f95ba13683de33"><code>7155d51</code></a> chore(ci): add kafka 3.6.0 to FVT and versions</li>
<li><a href="https://github.com/IBM/sarama/commit/e0c3c627b80ae4468f9e1dee60306dc2ce3fb284"><code>e0c3c62</code></a> fix(txmgr): ErrOffsetsLoadInProgress is retriable</li>
<li><a href="https://github.com/IBM/sarama/commit/2e077cf8d86cc06bffc2614b80a699caa868ad2b"><code>2e077cf</code></a> Fix default retention time value in offset commit (<a href="https://redirect.github.com/IBM/sarama/issues/2700">#2700</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/IBM/sarama/compare/v1.40.1...v1.42.1">compare view</a></li>
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

### Comment by @app-sre-bot on February 02, 2024 at 04:03 PM UTC

Can one of the admins verify this patch?

### Comment by @dependabot on February 05, 2024 at 04:10 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/559*
