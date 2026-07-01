---
type: pull_request
number: 802
title: "Build: Bump the go group across 1 directory with 5 updates"
state: closed
author: dependabot
created: 2024-09-06T14:09:01Z
updated: 2024-09-09T04:09:08Z
closed: 2024-09-09T04:09:07Z
base_branch: main
head_branch: dependabot/go_modules/go-1735a3e9a4
labels: ["dependencies"]
url: https://github.com/content-services/content-sources-backend/pull/802
---

# Pull Request #802: Build: Bump the go group across 1 directory with 5 updates

**Author**: @dependabot
**Created**: September 06, 2024 at 02:09 PM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-1735a3e9a4`

## Description

Bumps the go group with 5 updates in the / directory:

| Package | From | To |
| --- | --- | --- |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.20.2` | `1.20.3` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.30.4` | `1.30.5` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.29` | `1.17.32` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.37.5` | `1.39.0` |
| [github.com/RedHatInsights/insights-operator-utils](https://github.com/RedHatInsights/insights-operator-utils) | `1.25.9` | `1.25.10` |


Updates `github.com/prometheus/client_golang` from 1.20.2 to 1.20.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.20.3</h2>
<ul>
<li>[BUGFIX] histograms: Fix possible data race when appending exemplars. <a href="https://redirect.github.com/prometheus/client_golang/issues/1608">#1608</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/v1.20.3/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.20.3 / 2024-09-05</h2>
<ul>
<li>[BUGFIX] histograms: Fix possible data race when appending exemplars. <a href="https://redirect.github.com/prometheus/client_golang/issues/1608">#1608</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/ef2f87ea986252194ea960187b20b409180044dd"><code>ef2f87e</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1620">#1620</a> from prometheus/arthursens/prepare-1.20.3</li>
<li><a href="https://github.com/prometheus/client_golang/commit/937ac63d3d2dda83847f4ca842d62edabce4e743"><code>937ac63</code></a> Add changelog entry for 1.20.3</li>
<li><a href="https://github.com/prometheus/client_golang/commit/6e9914db5af255f5def17d54a7ca9c531771f4ca"><code>6e9914d</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1608">#1608</a> from krajorama/index-out-of-range-native-histogram-e...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/d6b8c8925bd16626cf168e642eb70724b17a0d61"><code>d6b8c89</code></a> Update comments with more explanations</li>
<li><a href="https://github.com/prometheus/client_golang/commit/504566f07c680f68743c3a5d239dede48538c7ec"><code>504566f</code></a> Use simplified solution from <a href="https://redirect.github.com/prometheus/client_golang/issues/1609">#1609</a> for the data race</li>
<li><a href="https://github.com/prometheus/client_golang/commit/dc8e9a4d8a4c7c64d5ae2c9d29a91bb1407d549b"><code>dc8e9a4</code></a> fix: native histogram: Simplify and fix addExemplar</li>
<li><a href="https://github.com/prometheus/client_golang/commit/dc819ceb1b0f906f1ab124f7492693970733a54d"><code>dc819ce</code></a> Use a trivial solution to <a href="https://redirect.github.com/prometheus/client_golang/issues/1605">#1605</a></li>
<li><a href="https://github.com/prometheus/client_golang/commit/e061dfae88c0dd63ff477a153096a1ba28f69f7e"><code>e061dfa</code></a> native histogram: use exemplars in concurrency test</li>
<li>See full diff in <a href="https://github.com/prometheus/client_golang/compare/v1.20.2...v1.20.3">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.30.4 to 1.30.5
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2b751d1ba71f59175a41f9cae5f159f1044360f"><code>a2b751d</code></a> Release 2024-09-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e22c2495bd38232c640776ef3c1a84fb3145a8b9"><code>e22c249</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff0cf6fbfa7c7a12388606d5568135f2beae6599"><code>ff0cf6f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3120376762853b3098fda7e9217fb39fe1bf5105"><code>3120376</code></a> refactoring of buildQuery to accept a list of maintained headers to l… (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2773">#2773</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ed838eab2a963cb16301501c8b8c3e29dac4c20"><code>4ed838e</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2768">#2768</a> from bhavya2109sharma/presignedurl-requestpayer-change</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/d4bd42fdc82c2954f429bd9eeac3096f5590eaac"><code>d4bd42f</code></a> Merge branch 'main' into presignedurl-requestpayer-change</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0353706229a89749afa93324432ece53da37822b"><code>0353706</code></a> Added Changelog</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/97e2d3fec04bc8bfd69eaf22ce476b12f8673810"><code>97e2d3f</code></a> Release 2024-08-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4cca52b4a81df57e91b1e5d0a65fd6df89606d02"><code>4cca52b</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a5146734eb1e6d59acba18b5967c78dc7a5e42"><code>c8a5146</code></a> Update endpoints model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.30.4...v1.30.5">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.29 to 1.17.32
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f1d71c59a1499981873f70db8c92d07242779670"><code>f1d71c5</code></a> Release 2024-09-04</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e4813e114f3c2ce8db07624f656073440adc1140"><code>e4813e1</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0e8bb90d8c0d03824fc8bf8656bd2de6346e48f8"><code>0e8bb90</code></a> Update partitions file</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/6a5875b13d01e8270c1cbdfae0c63ee479386917"><code>6a5875b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f7030de42b2a92ad4f66b153ae4d88ad28a6b2fe"><code>f7030de</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/a2b751d1ba71f59175a41f9cae5f159f1044360f"><code>a2b751d</code></a> Release 2024-09-03</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/e22c2495bd38232c640776ef3c1a84fb3145a8b9"><code>e22c249</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/ff0cf6fbfa7c7a12388606d5568135f2beae6599"><code>ff0cf6f</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3120376762853b3098fda7e9217fb39fe1bf5105"><code>3120376</code></a> refactoring of buildQuery to accept a list of maintained headers to l… (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2773">#2773</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/4ed838eab2a963cb16301501c8b8c3e29dac4c20"><code>4ed838e</code></a> Merge pull request <a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2768">#2768</a> from bhavya2109sharma/presignedurl-requestpayer-change</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.29...credentials/v1.17.32">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.37.5 to 1.39.0
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/807c921d812ada1f36e8e016035a375ee20b6379"><code>807c921</code></a> Release 2023-09-20</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/39b991de3a1abe48d06850b17381a53ab8c345e8"><code>39b991d</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b1d6dafaff6a5fe77514b0329508ed467404f282"><code>b1d6daf</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5771d2244ce2f85d536e4ee5c9f014b9b7cfff4a"><code>5771d22</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/7143703367f9e82e2add61e58fcf19dfe3d7cb43"><code>7143703</code></a> Release 2023-09-19</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/1803c873fb010d3991d36927eafc5a8683bc9ea4"><code>1803c87</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/228b865841787ce518d849ac7584225b65b22518"><code>228b865</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f416d6651c5eb1850565963502dca48f946ec1e0"><code>f416d66</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/f0b396e56da1e19abc0855af33cdfbf3c7531882"><code>f0b396e</code></a> Release 2023-09-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/53896255622f253e3273ffe2bada16ab94a199d0"><code>5389625</code></a> Regenerated Clients</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/ssm/v1.37.5...service/s3/v1.39.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/RedHatInsights/insights-operator-utils` from 1.25.9 to 1.25.10
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/737dd9f5137f7ad570a9643eeddb94cd69aae360"><code>737dd9f</code></a> [CCXDEV-14052][CCXDEV-14098] Avoid duplicates when using Sentry/Glitchtip (<a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/577">#577</a>)</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/4eb4ca0819aed30979b1cf334d0bec6a7bf56e86"><code>4eb4ca0</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/576">#576</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/87c044b561c57cef1a0365df7f2cb3ca3e6fa0c1"><code>87c044b</code></a> Bump github.com/prometheus/client_golang from 1.20.1 to 1.20.2</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/ea3fa7a427c07b900d8e1cad302ed00154c4d5cf"><code>ea3fa7a</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/575">#575</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/71d59cb4bd199ab3a9729e3a33435559781e54d5"><code>71d59cb</code></a> Bump github.com/prometheus/client_golang from 1.20.0 to 1.20.1</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/06947ff285eb5d5e6e470f48490f0b7e08e9d65c"><code>06947ff</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/insights-operator-utils/issues/574">#574</a> from RedHatInsights/dependabot/go_modules/github.com/...</li>
<li><a href="https://github.com/RedHatInsights/insights-operator-utils/commit/aeedde16e251144fec30ce3e583be18969d13af0"><code>aeedde1</code></a> Bump github.com/prometheus/client_golang from 1.19.1 to 1.20.0</li>
<li>See full diff in <a href="https://github.com/RedHatInsights/insights-operator-utils/compare/v1.25.9...v1.25.10">compare view</a></li>
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

### Comment by @app-sre-bot on September 06, 2024 at 02:09 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on September 06, 2024 at 07:57 PM UTC

[test]

### Comment by @dependabot on September 09, 2024 at 04:09 AM UTC

Looks like these dependencies are updatable in another way, so this is no longer needed.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/802*
