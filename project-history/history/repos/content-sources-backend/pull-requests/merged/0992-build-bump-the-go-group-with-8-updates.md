---
type: pull_request
number: 992
title: "Build: Bump the go group with 8 updates"
state: merged
author: dependabot
created: 2025-02-24T04:22:53Z
updated: 2025-02-24T09:44:07Z
closed: 2025-02-24T09:43:58Z
merged: 2025-02-24T09:43:58Z
base_branch: main
head_branch: dependabot/go_modules/go-68ea6ef347
labels: ["dependencies", "go"]
url: https://github.com/content-services/content-sources-backend/pull/992
---

# Pull Request #992: Build: Bump the go group with 8 updates

**Author**: @dependabot
**Created**: February 24, 2025 at 04:22 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `go`
**Base**: `main` ← **Head**: `dependabot/go_modules/go-68ea6ef347`

## Description

Bumps the go group with 8 updates:

| Package | From | To |
| --- | --- | --- |
| [github.com/content-services/tang](https://github.com/content-services/tang) | `0.0.11` | `0.0.12` |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | `1.20.5` | `1.21.0` |
| [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | `1.36.1` | `1.36.2` |
| [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2) | `1.29.6` | `1.29.7` |
| [github.com/aws/aws-sdk-go-v2/credentials](https://github.com/aws/aws-sdk-go-v2) | `1.17.59` | `1.17.60` |
| [github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs](https://github.com/aws/aws-sdk-go-v2) | `1.45.13` | `1.45.14` |
| [github.com/aws/aws-sdk-go-v2/service/s3](https://github.com/aws/aws-sdk-go-v2) | `1.77.0` | `1.77.1` |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | `9.7.0` | `9.7.1` |

Updates `github.com/content-services/tang` from 0.0.11 to 0.0.12
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/content-services/tang/releases">github.com/content-services/tang's releases</a>.</em></p>
<blockquote>
<h2>v0.0.12</h2>
<h2>What's Changed</h2>
<ul>
<li>HMS 5162: add new repos for testing by <a href="https://github.com/rverdile"><code>@​rverdile</code></a> in <a href="https://redirect.github.com/content-services/tang/pull/17">content-services/tang#17</a></li>
</ul>
<p><strong>Full Changelog</strong>: <a href="https://github.com/content-services/tang/compare/v0.0.11...v0.0.12">https://github.com/content-services/tang/compare/v0.0.11...v0.0.12</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/content-services/tang/commit/079f9fcfffe1b5f4ae9551d43c19ede628d7e9f2"><code>079f9fc</code></a> HMS 5162: add new repos for testing (<a href="https://redirect.github.com/content-services/tang/issues/17">#17</a>)</li>
<li>See full diff in <a href="https://github.com/content-services/tang/compare/v0.0.11...v0.0.12">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/prometheus/client_golang` from 1.20.5 to 1.21.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/releases">github.com/prometheus/client_golang's releases</a>.</em></p>
<blockquote>
<h2>v1.21.0 / 2025-02-19</h2>
<p>:warning: This release contains potential breaking change if you upgrade <code>github.com/prometheus/common</code> to 0.62+ together with client_golang (and depend on the strict, legacy validation for the label names). New common version <a href="https://redirect.github.com/prometheus/common/pull/724">changes <code>model.NameValidationScheme</code> global variable</a>, which relaxes the validation of label names and metric name, allowing all UTF-8 characters. Typically, this should not break any user, unless your test or usage expects strict certain names to panic/fail on client_golang metric registration, gathering or scrape. In case of problems change <code>model.NameValidationScheme</code> to old <code>model.LegacyValidation</code> value in your project <code>init</code> function. :warning:</p>
<ul>
<li>[BUGFIX] gocollector: Fix help message for runtime/metric metrics. <a href="https://redirect.github.com/prometheus/client_golang/issues/1583">#1583</a></li>
<li>[BUGFIX] prometheus: Fix <code>Desc.String()</code> method for no labels case. <a href="https://redirect.github.com/prometheus/client_golang/issues/1687">#1687</a></li>
<li>[PERF] prometheus: Optimize popular <code>prometheus.BuildFQName</code> function; now up to 30% faster. <a href="https://redirect.github.com/prometheus/client_golang/issues/1665">#1665</a></li>
<li>[PERF] prometheus: Optimize <code>Inc</code>, <code>Add</code> and <code>Observe</code> cumulative metrics; now up to 50% faster under high concurrent contention. <a href="https://redirect.github.com/prometheus/client_golang/issues/1661">#1661</a></li>
<li>[CHANGE] Upgrade prometheus/common to 0.62.0 which changes <code>model.NameValidationScheme</code> global variable. <a href="https://redirect.github.com/prometheus/client_golang/issues/1712">#1712</a></li>
<li>[CHANGE] Add support for Go 1.23. <a href="https://redirect.github.com/prometheus/client_golang/issues/1602">#1602</a></li>
<li>[FEATURE] process_collector: Add support for Darwin systems. <a href="https://redirect.github.com/prometheus/client_golang/issues/1600">#1600</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1616">#1616</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1625">#1625</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1675">#1675</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1715">#1715</a></li>
<li>[FEATURE] api: Add ability to invoke <code>CloseIdleConnections</code> on api.Client using <code>api.Client.(CloseIdler).CloseIdleConnections()</code> casting. <a href="https://redirect.github.com/prometheus/client_golang/issues/1513">#1513</a></li>
<li>[FEATURE] promhttp: Add <code>promhttp.HandlerOpts.EnableOpenMetricsTextCreatedSamples</code> option to create OpenMetrics _created lines. Not recommended unless you want to use opt-in Created Timestamp feature. Community works on OpenMetrics 2.0 format that should make those lines obsolete (they increase cardinality significantly). <a href="https://redirect.github.com/prometheus/client_golang/issues/1408">#1408</a></li>
<li>[FEATURE] prometheus: Add <code>NewConstNativeHistogram</code> function. <a href="https://redirect.github.com/prometheus/client_golang/issues/1654">#1654</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/prometheus/client_golang/blob/main/CHANGELOG.md">github.com/prometheus/client_golang's changelog</a>.</em></p>
<blockquote>
<h2>1.21.0 / 2025-02-17</h2>
<p>:warning: This release contains potential breaking change if you upgrade <code>github.com/prometheus/common</code> to 0.62+ together with client_golang. :warning:</p>
<p>New common version <a href="https://redirect.github.com/prometheus/common/pull/724">changes <code>model.NameValidationScheme</code> global variable</a>, which relaxes the validation of label names and metric name, allowing all UTF-8 characters. Typically, this should not break any user, unless your test or usage expects strict certain names to panic/fail on client_golang metric registration, gathering or scrape. In case of problems change <code>model.NameValidationScheme</code> to old <code>model.LegacyValidation</code> value in your project <code>init</code> function.</p>
<ul>
<li>[BUGFIX] gocollector: Fix help message for runtime/metric metrics. <a href="https://redirect.github.com/prometheus/client_golang/issues/1583">#1583</a></li>
<li>[BUGFIX] prometheus: Fix <code>Desc.String()</code> method for no labels case. <a href="https://redirect.github.com/prometheus/client_golang/issues/1687">#1687</a></li>
<li>[ENHANCEMENT] prometheus: Optimize popular <code>prometheus.BuildFQName</code> function; now up to 30% faster. <a href="https://redirect.github.com/prometheus/client_golang/issues/1665">#1665</a></li>
<li>[ENHANCEMENT] prometheus: Optimize <code>Inc</code>, <code>Add</code> and <code>Observe</code> cumulative metrics; now up to 50% faster under high concurrent contention. <a href="https://redirect.github.com/prometheus/client_golang/issues/1661">#1661</a></li>
<li>[CHANGE] Upgrade prometheus/common to 0.62.0 which changes <code>model.NameValidationScheme</code> global variable. <a href="https://redirect.github.com/prometheus/client_golang/issues/1712">#1712</a></li>
<li>[CHANGE] Add support for Go 1.23. <a href="https://redirect.github.com/prometheus/client_golang/issues/1602">#1602</a></li>
<li>[FEATURE] process_collector: Add support for Darwin systems. <a href="https://redirect.github.com/prometheus/client_golang/issues/1600">#1600</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1616">#1616</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1625">#1625</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1675">#1675</a> <a href="https://redirect.github.com/prometheus/client_golang/issues/1715">#1715</a></li>
<li>[FEATURE] api: Add ability to invoke <code>CloseIdleConnections</code> on api.Client using <code>api.Client.(CloseIdler).CloseIdleConnections()</code> casting. <a href="https://redirect.github.com/prometheus/client_golang/issues/1513">#1513</a></li>
<li>[FEATURE] promhttp: Add <code>promhttp.HandlerOpts.EnableOpenMetricsTextCreatedSamples</code> option to create OpenMetrics _created lines. Not recommended unless you want to use opt-in Created Timestamp feature. Community works on OpenMetrics 2.0 format that should make those lines obsolete (they increase cardinality significantly). <a href="https://redirect.github.com/prometheus/client_golang/issues/1408">#1408</a></li>
<li>[FEATURE] prometheus: Add <code>NewConstNativeHistogram</code> function. <a href="https://redirect.github.com/prometheus/client_golang/issues/1654">#1654</a></li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/prometheus/client_golang/commit/eaf03ef9509cf7e0e56a7d0eda1f11a05506f045"><code>eaf03ef</code></a> Cut 1.21.0 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1737">#1737</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/f1f89dc6c527ddf1e80b49c4f56b2c52b164105c"><code>f1f89dc</code></a> Cut 1.21.0-rc.0 (<a href="https://redirect.github.com/prometheus/client_golang/issues/1718">#1718</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/c923f7c8e40301ccf857e28f655a241695c470d7"><code>c923f7c</code></a> Revert &quot;ci: daggerize test and lint pipelines (<a href="https://redirect.github.com/prometheus/client_golang/issues/1534">#1534</a>)&quot; (<a href="https://redirect.github.com/prometheus/client_golang/issues/1717">#1717</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/1bcda802c13d6334110e088bbef96d32c1c05db7"><code>1bcda80</code></a> process collector: Fixed pedantic registry failures on darwin with cgo. (<a href="https://redirect.github.com/prometheus/client_golang/issues/1715">#1715</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/038b37aea518190a66e2d4df5d231e549eed7759"><code>038b37a</code></a> tutorials/whatsup: Updated deps (<a href="https://redirect.github.com/prometheus/client_golang/issues/1716">#1716</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/56a24311d5ef75b4c198516ed1c4555318ec729a"><code>56a2431</code></a> docs: Add RELEASE.md for the release process (<a href="https://redirect.github.com/prometheus/client_golang/issues/1690">#1690</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/cbd9526e6ddc36b3cec0407b70e86e8249edf4ed"><code>cbd9526</code></a> Merge pull request <a href="https://redirect.github.com/prometheus/client_golang/issues/1713">#1713</a> from prometheus/dependabot/go_modules/tutorials/what...</li>
<li><a href="https://github.com/prometheus/client_golang/commit/80b5a2a705c6cf39bf08260d06fd130024affbd5"><code>80b5a2a</code></a> build(deps): bump golang.org/x/net in /tutorials/whatsup</li>
<li><a href="https://github.com/prometheus/client_golang/commit/1a822a841f0ae8c1b93d6cbd3748d881e9023e05"><code>1a822a8</code></a> Upgrade to prometheus/common 0.62.0 with breaking change (<a href="https://redirect.github.com/prometheus/client_golang/issues/1712">#1712</a>)</li>
<li><a href="https://github.com/prometheus/client_golang/commit/7b39d0144166aa94cc8ce4125bcb3b0da89aad5e"><code>7b39d01</code></a> Update common Prometheus files (<a href="https://redirect.github.com/prometheus/client_golang/issues/1708">#1708</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/prometheus/client_golang/compare/v1.20.5...v1.21.0">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2` from 1.36.1 to 1.36.2
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54aed732316b5162e5c4382a1f2d3891175d0254"><code>54aed73</code></a> Release 2025-02-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/082781faee876f9d612fa7c113b4304a29766b14"><code>082781f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ed185b604684a86547e679154975f1914f97312"><code>3ed185b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03da7378d668622cd880ec741d57e93cc370efa1"><code>03da737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a8ccb619ffbfe00e99a83e99729b948f20be29"><code>c8a8ccb</code></a> Bump go version to 1.22 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3010">#3010</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8b7c7bf6d9a1c63d0c5262724ae8a15a44e366a6"><code>8b7c7bf</code></a> fix missing AccountIDEndpointMode binding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3013">#3013</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90f9d1081a37acaf792ccda5bfb07e2ee7590a9e"><code>90f9d10</code></a> Release 2025-02-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40dd351c61c016749a3f4105cca0c965e7c66d7b"><code>40dd351</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06352dfafe9067da1956229d6925efed328d5ff6"><code>06352df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/833566b553122ebd5bfa1237ee7c905a8db0d687"><code>833566b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/v1.36.1...v1.36.2">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/config` from 1.29.6 to 1.29.7
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54aed732316b5162e5c4382a1f2d3891175d0254"><code>54aed73</code></a> Release 2025-02-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/082781faee876f9d612fa7c113b4304a29766b14"><code>082781f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ed185b604684a86547e679154975f1914f97312"><code>3ed185b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03da7378d668622cd880ec741d57e93cc370efa1"><code>03da737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a8ccb619ffbfe00e99a83e99729b948f20be29"><code>c8a8ccb</code></a> Bump go version to 1.22 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3010">#3010</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8b7c7bf6d9a1c63d0c5262724ae8a15a44e366a6"><code>8b7c7bf</code></a> fix missing AccountIDEndpointMode binding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3013">#3013</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90f9d1081a37acaf792ccda5bfb07e2ee7590a9e"><code>90f9d10</code></a> Release 2025-02-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40dd351c61c016749a3f4105cca0c965e7c66d7b"><code>40dd351</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06352dfafe9067da1956229d6925efed328d5ff6"><code>06352df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/833566b553122ebd5bfa1237ee7c905a8db0d687"><code>833566b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/config/v1.29.6...config/v1.29.7">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/credentials` from 1.17.59 to 1.17.60
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54aed732316b5162e5c4382a1f2d3891175d0254"><code>54aed73</code></a> Release 2025-02-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/082781faee876f9d612fa7c113b4304a29766b14"><code>082781f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ed185b604684a86547e679154975f1914f97312"><code>3ed185b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03da7378d668622cd880ec741d57e93cc370efa1"><code>03da737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a8ccb619ffbfe00e99a83e99729b948f20be29"><code>c8a8ccb</code></a> Bump go version to 1.22 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3010">#3010</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8b7c7bf6d9a1c63d0c5262724ae8a15a44e366a6"><code>8b7c7bf</code></a> fix missing AccountIDEndpointMode binding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3013">#3013</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90f9d1081a37acaf792ccda5bfb07e2ee7590a9e"><code>90f9d10</code></a> Release 2025-02-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40dd351c61c016749a3f4105cca0c965e7c66d7b"><code>40dd351</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06352dfafe9067da1956229d6925efed328d5ff6"><code>06352df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/833566b553122ebd5bfa1237ee7c905a8db0d687"><code>833566b</code></a> Update API model</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.59...credentials/v1.17.60">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs` from 1.45.13 to 1.45.14
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c7c68659ce67e5b7e18f31bc66068cec9e3d790d"><code>c7c6865</code></a> Release 2025-01-30</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/70f736c5dc0b8652c5fe5c387b2165c3b9beddb1"><code>70f736c</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/28731c2bdef3c2555a95632396b6d4936e58099d"><code>28731c2</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3505e4b255c327a1fa38f870612c327b93302dc0"><code>3505e4b</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/0b4f6362bb9017615ed38124cdd20df7714bf98f"><code>0b4f636</code></a> don't sign transfer-encoding header (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2991">#2991</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/b820c5731d3bff3ae303166ea4379c73fb18a8d4"><code>b820c57</code></a> Release 2025-01-29</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40ddb76e611944643404e79a6b92ac335f2921f2"><code>40ddb76</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/2b4adae0bf0cb97fd910110c1f567d5bbc1f348b"><code>2b4adae</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/78c2be4ca0b69cddb5825b20309e35d4f54c6cf4"><code>78c2be4</code></a> Revert &quot;beta: feature/s3/transfermanager (S3 transfer manager v2) (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2988">#2988</a>)&quot;</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/5c9d67c15ecc49b769f7ddab05f840e723e146ac"><code>5c9d67c</code></a> beta: feature/s3/transfermanager (S3 transfer manager v2) (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/2988">#2988</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/polly/v1.45.13...service/polly/v1.45.14">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/aws/aws-sdk-go-v2/service/s3` from 1.77.0 to 1.77.1
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/54aed732316b5162e5c4382a1f2d3891175d0254"><code>54aed73</code></a> Release 2025-02-18</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/082781faee876f9d612fa7c113b4304a29766b14"><code>082781f</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/3ed185b604684a86547e679154975f1914f97312"><code>3ed185b</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/03da7378d668622cd880ec741d57e93cc370efa1"><code>03da737</code></a> Update API model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/c8a8ccb619ffbfe00e99a83e99729b948f20be29"><code>c8a8ccb</code></a> Bump go version to 1.22 (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3010">#3010</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/8b7c7bf6d9a1c63d0c5262724ae8a15a44e366a6"><code>8b7c7bf</code></a> fix missing AccountIDEndpointMode binding (<a href="https://redirect.github.com/aws/aws-sdk-go-v2/issues/3013">#3013</a>)</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/90f9d1081a37acaf792ccda5bfb07e2ee7590a9e"><code>90f9d10</code></a> Release 2025-02-17</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/40dd351c61c016749a3f4105cca0c965e7c66d7b"><code>40dd351</code></a> Regenerated Clients</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/06352dfafe9067da1956229d6925efed328d5ff6"><code>06352df</code></a> Update endpoints model</li>
<li><a href="https://github.com/aws/aws-sdk-go-v2/commit/833566b553122ebd5bfa1237ee7c905a8db0d687"><code>833566b</code></a> Update API model</li>
<li>See full diff in <a href="https://github.com/aws/aws-sdk-go-v2/compare/service/s3/v1.77.0...service/s3/v1.77.1">compare view</a></li>
</ul>
</details>
<br />

Updates `github.com/redis/go-redis/v9` from 9.7.0 to 9.7.1
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/redis/go-redis/releases">github.com/redis/go-redis/v9's releases</a>.</em></p>
<blockquote>
<h2>v9.7.1</h2>
<h1>Changes</h1>
<ul>
<li>Recognize byte slice for key argument in cluster client hash slot computation (<a href="https://redirect.github.com/redis/go-redis/issues/3049">#3049</a>)</li>
<li>fix(search&amp;aggregate):fix error overwrite and typo  <a href="https://redirect.github.com/redis/go-redis/issues/3220">#3220</a> (<a href="https://redirect.github.com/redis/go-redis/issues/3224">#3224</a>)</li>
<li>fix: linter configuration (<a href="https://redirect.github.com/redis/go-redis/issues/3279">#3279</a>)</li>
<li>fix(search): if ft.aggregate use limit when limitoffset is zero (<a href="https://redirect.github.com/redis/go-redis/issues/3275">#3275</a>)</li>
<li>Reinstate read-only lock on hooks access in dialHook to fix data race (<a href="https://redirect.github.com/redis/go-redis/issues/3225">#3225</a>)</li>
<li>fix: flaky ClientKillByFilter test (<a href="https://redirect.github.com/redis/go-redis/issues/3268">#3268</a>)</li>
<li>chore: fix some comments (<a href="https://redirect.github.com/redis/go-redis/issues/3226">#3226</a>)</li>
<li>fix(aggregate, search): ft.aggregate bugfixes (<a href="https://redirect.github.com/redis/go-redis/issues/3263">#3263</a>)</li>
<li>fix: add unstableresp3 to cluster client (<a href="https://redirect.github.com/redis/go-redis/issues/3266">#3266</a>)</li>
<li>Fix race condition in clusterNodes.Addrs() (<a href="https://redirect.github.com/redis/go-redis/issues/3219">#3219</a>)</li>
<li>SortByWithCount FTSearchOptions fix (<a href="https://redirect.github.com/redis/go-redis/issues/3201">#3201</a>)</li>
<li>Eliminate redundant dial mutex causing unbounded connection queue contention (<a href="https://redirect.github.com/redis/go-redis/issues/3088">#3088</a>)</li>
<li>Add guidance on unstable RESP3 support for RediSearch commands to README (<a href="https://redirect.github.com/redis/go-redis/issues/3177">#3177</a>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li>Add guidance on unstable RESP3 support for RediSearch commands to README (<a href="https://redirect.github.com/redis/go-redis/issues/3177">#3177</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li>fix(search): if ft.aggregate use limit when limitoffset is zero (<a href="https://redirect.github.com/redis/go-redis/issues/3275">#3275</a>)</li>
<li>fix: add unstableresp3 to cluster client (<a href="https://redirect.github.com/redis/go-redis/issues/3266">#3266</a>)</li>
<li>fix(aggregate, search): ft.aggregate bugfixes (<a href="https://redirect.github.com/redis/go-redis/issues/3263">#3263</a>)</li>
<li>SortByWithCount FTSearchOptions fix (<a href="https://redirect.github.com/redis/go-redis/issues/3201">#3201</a>)</li>
<li>Recognize byte slice for key argument in cluster client hash slot computation (<a href="https://redirect.github.com/redis/go-redis/issues/3049">#3049</a>)</li>
</ul>
<h2>Contributors</h2>
<p>We'd like to thank all the contributors who worked on this release!</p>
<p><a href="https://github.com/ofekshenawa"><code>@​ofekshenawa</code></a>, <a href="https://github.com/Cgol9"><code>@​Cgol9</code></a>, <a href="https://github.com/LINKIWI"><code>@​LINKIWI</code></a>, <a href="https://github.com/shawnwgit"><code>@​shawnwgit</code></a>, <a href="https://github.com/zhuhaicity"><code>@​zhuhaicity</code></a>, <a href="https://github.com/bitsark"><code>@​bitsark</code></a>, <a href="https://github.com/vladvildanov"><code>@​vladvildanov</code></a>, <a href="https://github.com/ndyakov"><code>@​ndyakov</code></a></p>
<p><strong>Full Changelog</strong>: <a href="https://github.com/redis/go-redis/compare/v9.7.0...v9.7.1">https://github.com/redis/go-redis/compare/v9.7.0...v9.7.1</a></p>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/redis/go-redis/commit/3d041a1dd6974412c95803384965d73a50bf07ab"><code>3d041a1</code></a> release: 9.7.1 patch (<a href="https://redirect.github.com/redis/go-redis/issues/3278">#3278</a>)</li>
<li>See full diff in <a href="https://github.com/redis/go-redis/compare/v9.7.0...v9.7.1">compare view</a></li>
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

## Reviews

### Review by @swadeley - Approved on February 24, 2025 at 09:43 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/992*
