---
type: pull_request
number: 1263
title: "chore(deps): bump @sentry/browser, @redhat-cloud-services/frontend-components-utilities and @redhat-cloud-services/frontend-components-remediations"
state: closed
author: dependabot
created: 2025-01-28T11:59:17Z
updated: 2025-01-30T14:49:58Z
closed: 2025-01-30T14:49:56Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-bb2367b2e9
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1263
---

# Pull Request #1263: chore(deps): bump @sentry/browser, @redhat-cloud-services/frontend-components-utilities and @redhat-cloud-services/frontend-components-remediations

**Author**: @dependabot
**Created**: January 28, 2025 at 11:59 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-bb2367b2e9`

## Description

Bumps [@sentry/browser](https://github.com/getsentry/sentry-javascript) to 7.120.3 and updates ancestor dependencies [@sentry/browser](https://github.com/getsentry/sentry-javascript), [@redhat-cloud-services/frontend-components-utilities](https://github.com/RedHatInsights/frontend-components) and [@redhat-cloud-services/frontend-components-remediations](https://github.com/RedHatInsights/frontend-components). These dependencies need to be updated together.

Updates `@sentry/browser` from 5.30.0 to 7.120.3
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-javascript/releases"><code>@​sentry/browser</code>'s releases</a>.</em></p>
<blockquote>
<h2>7.120.3</h2>
<ul>
<li>fix(v7/publish): Ensure discontinued packages are published with <code>latest</code> tag (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14926">#14926</a>)</li>
</ul>
<h2>Bundle size 📦</h2>
<table>
<thead>
<tr>
<th>Path</th>
<th>Size</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay, Feedback) - Webpack (gzipped)</td>
<td>80.96 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - Webpack (gzipped)</td>
<td>71.9 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay with Canvas) - Webpack (gzipped)</td>
<td>76.15 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - Webpack with treeshaking flags (gzipped)</td>
<td>65.53 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing) - Webpack (gzipped)</td>
<td>35.78 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. browserTracingIntegration) - Webpack (gzipped)</td>
<td>35.67 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Feedback) - Webpack (gzipped)</td>
<td>31.71 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. sendFeedback) - Webpack (gzipped)</td>
<td>31.73 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> - Webpack (gzipped)</td>
<td>22.92 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay, Feedback) - ES6 CDN Bundle (gzipped)</td>
<td>79.18 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - ES6 CDN Bundle (gzipped)</td>
<td>70.5 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing) - ES6 CDN Bundle (gzipped)</td>
<td>36.17 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> - ES6 CDN Bundle (gzipped)</td>
<td>25.42 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - ES6 CDN Bundle (minified &amp; uncompressed)</td>
<td>221.94 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing) - ES6 CDN Bundle (minified &amp; uncompressed)</td>
<td>109.54 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> - ES6 CDN Bundle (minified &amp; uncompressed)</td>
<td>76.26 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing) - ES5 CDN Bundle (gzipped)</td>
<td>39.45 KB</td>
</tr>
<tr>
<td><code>@​sentry/react</code> (incl. Tracing, Replay) - Webpack (gzipped)</td>
<td>72.4 KB</td>
</tr>
<tr>
<td><code>@​sentry/react</code> - Webpack (gzipped)</td>
<td>22.95 KB</td>
</tr>
<tr>
<td><code>@​sentry/nextjs</code> Client (incl. Tracing, Replay) - Webpack (gzipped)</td>
<td>90.16 KB</td>
</tr>
<tr>
<td><code>@​sentry/nextjs</code> Client - Webpack (gzipped)</td>
<td>54.28 KB</td>
</tr>
<tr>
<td><code>@​sentry-internal/feedback</code> - Webpack (gzipped)</td>
<td>17.34 KB</td>
</tr>
</tbody>
</table>
<h2>7.120.2</h2>
<ul>
<li>fix(tracing-internal): Fix case when lrp keys offset is 0 (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14615">#14615</a>)</li>
</ul>
<p>Work in this release contributed by <a href="https://github.com/LubomirIgonda1"><code>@​LubomirIgonda1</code></a>. Thank you for your contribution!</p>
<h2>Bundle size 📦</h2>
<table>
<thead>
<tr>
<th>Path</th>
<th>Size</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay, Feedback) - Webpack (gzipped)</td>
<td>80.96 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - Webpack (gzipped)</td>
<td>71.9 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay with Canvas) - Webpack (gzipped)</td>
<td>76.15 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - Webpack with treeshaking flags (gzipped)</td>
<td>65.52 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing) - Webpack (gzipped)</td>
<td>35.78 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. browserTracingIntegration) - Webpack (gzipped)</td>
<td>35.67 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Feedback) - Webpack (gzipped)</td>
<td>31.71 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. sendFeedback) - Webpack (gzipped)</td>
<td>31.73 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> - Webpack (gzipped)</td>
<td>22.92 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay, Feedback) - ES6 CDN Bundle (gzipped)</td>
<td>79.18 KB</td>
</tr>
<tr>
<td><code>@​sentry/browser</code> (incl. Tracing, Replay) - ES6 CDN Bundle (gzipped)</td>
<td>70.5 KB</td>
</tr>
</tbody>
</table>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/getsentry/sentry-javascript/blob/7.120.3/CHANGELOG.md"><code>@​sentry/browser</code>'s changelog</a>.</em></p>
<blockquote>
<h2>7.120.3</h2>
<ul>
<li>fix(v7/publish): Ensure discontinued packages are published with <code>latest</code> tag (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14926">#14926</a>)</li>
</ul>
<h2>7.120.2</h2>
<ul>
<li>fix(tracing-internal): Fix case when lrp keys offset is 0 (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14615">#14615</a>)</li>
</ul>
<p>Work in this release contributed by <a href="https://github.com/LubomirIgonda1"><code>@​LubomirIgonda1</code></a>. Thank you for your contribution!</p>
<h2>7.120.1</h2>
<ul>
<li>fix(v7/cdn): Ensure <code>_sentryModuleMetadata</code> is not mangled (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14357">#14357</a>)</li>
</ul>
<p>Work in this release contributed by <a href="https://github.com/gilisho"><code>@​gilisho</code></a>. Thank you for your contribution!</p>
<h2>7.120.0</h2>
<ul>
<li>feat(v7/browser): Add moduleMetadataIntegration lazy loading support (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/13822">#13822</a>)</li>
</ul>
<p>Work in this release contributed by <a href="https://github.com/gilisho"><code>@​gilisho</code></a>. Thank you for your contribution!</p>
<h2>7.119.2</h2>
<ul>
<li>chore(nextjs/v7): Bump rollup to 2.79.2</li>
</ul>
<h2>7.119.1</h2>
<ul>
<li>fix(browser/v7): Ensure wrap() only returns functions (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/13838">#13838</a> backport)</li>
</ul>
<p>Work in this release contributed by <a href="https://github.com/legobeat"><code>@​legobeat</code></a>. Thank you for your contribution!</p>
<h2>7.119.0</h2>
<ul>
<li>backport(tracing): Report dropped spans for transactions (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/13343">#13343</a>)</li>
</ul>
<h2>7.118.0</h2>
<ul>
<li>fix(v7/bundle): Ensure CDN bundles do not overwrite <code>window.Sentry</code> (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/12579">#12579</a>)</li>
</ul>
<h2>7.117.0</h2>
<ul>
<li>feat(browser/v7): Publish browserprofling CDN bundle (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/12224">#12224</a>)</li>
<li>fix(v7/publish): Add <code>v7</code> tag to <code>@sentry/replay</code> (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/12304">#12304</a>)</li>
</ul>
<h2>7.116.0</h2>
<ul>
<li>build(craft): Publish lambda layer under its own name for v7 (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/12098">#12098</a>) (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/12099">#12099</a>)</li>
</ul>
<p>This release publishes a new AWS Lambda layer under the name <code>SentryNodeServerlessSDKv7</code> that users still running v7 can</p>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/5a833b498a37dc925ba5a909e28908c2105cb10c"><code>5a833b4</code></a> release: 7.120.3</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/6972e9a66e0608b7196017cec6c85bbc34f121e9"><code>6972e9a</code></a> meta: Update Changelog for 7.120.3 (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14960">#14960</a>)</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/e779b43e092feb40d09fa36c3319fda50147aa73"><code>e779b43</code></a> fix(v7/publish): Ensure discontinued packages are published with <code>latest</code> tag...</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/bc7dfd64961cc85405ca0d2ae91ea6545847a1de"><code>bc7dfd6</code></a> Merge branch 'release/7.120.2' into v7</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/85a3cacde4523c39c734b72346b7cdb39125d30b"><code>85a3cac</code></a> release: 7.120.2</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/5908e0acc73b0ae05df8267d92172be4636928c6"><code>5908e0a</code></a> meta(v7/changelog): Add changelog for 1.120.2 (<a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14673">#14673</a>)</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/4690bf34cc06840af59902ecdd70e1bd05424796"><code>4690bf3</code></a> Merge pull request <a href="https://redirect.github.com/getsentry/sentry-javascript/issues/14615">#14615</a> from LubomirIgonda1/fix-express-route-params</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/0e9c7f6cd537670b4c8d816b455d3e272980f6c4"><code>0e9c7f6</code></a> fix(tracing-internal): fix code formatting</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/78892adb1b3a62e757ee9d60e6f13965715c0bcd"><code>78892ad</code></a> fix(tracing-internal): Fix case when lrp keys offset is 0</li>
<li><a href="https://github.com/getsentry/sentry-javascript/commit/fde81a635771adbeaa5937a0f36cf836d1305f23"><code>fde81a6</code></a> Merge branch 'release/7.120.1' into v7</li>
<li>Additional commits viewable in <a href="https://github.com/getsentry/sentry-javascript/compare/5.30.0...7.120.3">compare view</a></li>
</ul>
</details>
<br />

Updates `@redhat-cloud-services/frontend-components-utilities` from 4.0.17 to 4.0.19
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/RedHatInsights/frontend-components/commits">compare view</a></li>
</ul>
</details>
<br />

Updates `@redhat-cloud-services/frontend-components-remediations` from 3.2.12 to 3.2.24
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/frontend-components/releases"><code>@​redhat-cloud-services/frontend-components-remediations</code>'s releases</a>.</em></p>
<blockquote>
<h2><code>@​redhat-cloud-services/frontend-components-remediations-3</code>.2.24</h2>
<h1>Changelog</h1>
<p>This file was generated using <a href="https://github.com/jscutlery/semver"><code>@​jscutlery/semver</code></a>.</p>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.23...@redhat-cloud-services/frontend-components-remediations-3.2.24">3.2.24</a> (2025-01-28)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.2.2</code></li>
<li><code>@redhat-cloud-services/frontend-components-utilities</code> updated to version <code>5.0.9</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.22...@redhat-cloud-services/frontend-components-remediations-3.2.23">3.2.23</a> (2025-01-27)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.2.1</code></li>
<li><code>@redhat-cloud-services/frontend-components-utilities</code> updated to version <code>5.0.8</code></li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li><strong>build:</strong> fix release postTarget nested dependencies (<a href="https://github.com/RedHatInsights/frontend-components/commit/4895cd2eba32336a220ddec442916858400ebb3e">4895cd2</a>)</li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.21...@redhat-cloud-services/frontend-components-remediations-3.2.22">3.2.22</a> (2025-01-23)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.2.0</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.20...@redhat-cloud-services/frontend-components-remediations-3.2.21">3.2.21</a> (2025-01-16)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.1.3</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.19...@redhat-cloud-services/frontend-components-remediations-3.2.20">3.2.20</a> (2025-01-15)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.1.2</code></li>
<li><code>@redhat-cloud-services/frontend-components-utilities</code> updated to version <code>5.0.7</code></li>
</ul>
<h3>Bug Fixes</h3>
<ul>
<li><strong>deps:</strong> use latest version of fec packages (<a href="https://github.com/RedHatInsights/frontend-components/commit/922dcf6795942109d75c77273b546ca7f726b2a8">922dcf6</a>)</li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.18...@redhat-cloud-services/frontend-components-remediations-3.2.19">3.2.19</a> (2025-01-14)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.1.1</code></li>
<li><code>@redhat-cloud-services/frontend-components-utilities</code> updated to version <code>5.0.6</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.17...@redhat-cloud-services/frontend-components-remediations-3.2.18">3.2.18</a> (2025-01-14)</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/174ffcd1eac4de459ba44973796eda405d9714a8"><code>174ffcd</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-remediations</code> to 3.2.24...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/d120169b2fc78222c320fa2d8459b298ba0bff22"><code>d120169</code></a> chore: bump <code>@​redhat-cloud-services/rule-components</code> to 3.2.21 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/01f0f1fc1a020df4135a14e9f2fcb899a88da47b"><code>01f0f1f</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-translations</code> to 3.2.17...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/f2593c540d07e90c8133a46b01411d4f9b9f12e8"><code>f2593c5</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-testing</code> to 0.1.5 [skip...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/4f0fb61ab2932faae57132f5f1668f7a7807ba4b"><code>4f0fb61</code></a> chore: bump <code>@​redhat-cloud-services/chrome</code> to 1.0.17 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/b62796cf4070d34f2a4edd35fc5eb69c596e526b"><code>b62796c</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-advisor-components</code> to ...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/1d9a8b53d0ea749fa7422c41f51a025b4287b663"><code>1d9a8b5</code></a> chore: bump <code>@​redhat-cloud-services/eslint-config-redhat-cloud-services</code> to 2.0...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/b9aca2d5609a39bbc3cebd55791849d1a263ad09"><code>b9aca2d</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-notifications</code> to 4.1.1...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/893b216bf1c0b81d011a5a4cb05435f3382b2c34"><code>893b216</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-config</code> to 6.4.2 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/2d3e20e611852f7c2110c8c00be320ffbb87350b"><code>2d3e20e</code></a> chore: bump <code>@​redhat-cloud-services/tsc-transform-imports</code> to 1.0.22 [skip ci]</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/rule-components-3.2.12...@redhat-cloud-services/frontend-components-remediations-3.2.24">compare view</a></li>
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
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-ui/network/alerts).

</details>

---

## Discussion

### Comment by @dependabot on January 30, 2025 at 02:49 PM UTC

Superseded by #1271.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1263*
