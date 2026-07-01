---
type: pull_request
number: 1256
title: "chore(deps): bump @sentry/browser, @redhat-cloud-services/frontend-components-utilities and @redhat-cloud-services/frontend-components-remediations"
state: closed
author: dependabot
created: 2025-01-17T12:48:14Z
updated: 2025-01-28T11:59:22Z
closed: 2025-01-28T11:59:20Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-e5a381f095
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1256
---

# Pull Request #1256: chore(deps): bump @sentry/browser, @redhat-cloud-services/frontend-components-utilities and @redhat-cloud-services/frontend-components-remediations

**Author**: @dependabot
**Created**: January 17, 2025 at 12:48 PM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-e5a381f095`

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

Updates `@redhat-cloud-services/frontend-components-remediations` from 3.2.12 to 3.2.20
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/frontend-components/releases"><code>@​redhat-cloud-services/frontend-components-remediations</code>'s releases</a>.</em></p>
<blockquote>
<h2><code>@​redhat-cloud-services/frontend-components-remediations-3</code>.2.20</h2>
<h1>Changelog</h1>
<p>This file was generated using <a href="https://github.com/jscutlery/semver"><code>@​jscutlery/semver</code></a>.</p>
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
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.1.0</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.16...@redhat-cloud-services/frontend-components-remediations-3.2.17">3.2.17</a> (2025-01-13)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.0.5</code></li>
<li><code>@redhat-cloud-services/frontend-components-utilities</code> updated to version <code>5.0.5</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.15...@redhat-cloud-services/frontend-components-remediations-3.2.16">3.2.16</a> (2024-12-17)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.0.4</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.14...@redhat-cloud-services/frontend-components-remediations-3.2.15">3.2.15</a> (2024-12-13)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.0.3</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.13...@redhat-cloud-services/frontend-components-remediations-3.2.14">3.2.14</a> (2024-12-02)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.0.2</code></li>
<li><code>@redhat-cloud-services/frontend-components-utilities</code> updated to version <code>5.0.4</code></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/8b25fd957cb8decf21e60b437851cf0189178fd6"><code>8b25fd9</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-remediations</code> to 3.2.20...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/162f1e7a63a59163b5920bdf4c71378662d2c04a"><code>162f1e7</code></a> chore: bump <code>@​redhat-cloud-services/rule-components</code> to 3.2.17 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/beb5f6591f9c641c5a870cc929a4e4e2ca39285f"><code>beb5f65</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-translations</code> to 3.2.15...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/ef16821a241541754493a33da287803498bc32ef"><code>ef16821</code></a> chore: bump <code>@​redhat-cloud-services/chrome</code> to 1.0.15 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/c40f7160cd897a96e34bb1a420338ef007666957"><code>c40f716</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-advisor-components</code> to ...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/8613f47b508d378cc99b0484acb4460ab6ea64ba"><code>8613f47</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-notifications</code> to 4.1.9...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/88b2372d3716fbb9158a3fe159436215d1757912"><code>88b2372</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-config</code> to 6.3.6 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/35139fe4564933b8bc920ec5f2369e8880e078bb"><code>35139fe</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components</code> to 5.1.2 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/f5bde1aed4397284eb4e3f6b8798ef6805d22c30"><code>f5bde1a</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-utilities</code> to 5.0.7 [sk...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/9d8ed84577db9edf7f053b4a7b91792068b26168"><code>9d8ed84</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/frontend-components/issues/2146">#2146</a> from bastilian/bump_components_utils_dep</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/rule-components-3.2.12...@redhat-cloud-services/frontend-components-remediations-3.2.20">compare view</a></li>
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

### Comment by @dependabot on January 28, 2025 at 11:59 AM UTC

Superseded by #1263.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1256*
