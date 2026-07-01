---
type: pull_request
number: 1575
title: "chore(deps): bump axios and @currents/playwright"
state: merged
author: dependabot
created: 2026-04-12T11:12:24Z
updated: 2026-04-14T17:05:28Z
closed: 2026-04-14T17:05:19Z
merged: 2026-04-14T17:05:19Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-69fcc311fd
labels: ["dependencies", "minor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1575
---

# Pull Request #1575: chore(deps): bump axios and @currents/playwright

**Author**: @dependabot
**Created**: April 12, 2026 at 11:12 AM UTC
**Status**: Merged
**Labels**: `dependencies`, `minor`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-69fcc311fd`

## Description

Bumps [axios](https://github.com/axios/axios) and [@currents/playwright](https://github.com/currents-dev/currents-playwright-changelog). These dependencies needed to be updated together.
Updates `axios` from 1.13.6 to 1.15.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
<h2>v1.15.0</h2>
<p>This release delivers two critical security patches, adds runtime support for Deno and Bun, and includes significant CI hardening, documentation improvements, and routine dependency updates.</p>
<h2>⚠️ Important Changes</h2>
<ul>
<li><strong>Deprecation:</strong> <code>url.parse()</code> usage has been replaced to address Node.js deprecation warnings. If you are on a recent version of Node.js, this resolves console warnings you may have been seeing. (<strong><a href="https://redirect.github.com/axios/axios/issues/10625">#10625</a></strong>)</li>
</ul>
<h2>🔒 Security Fixes</h2>
<ul>
<li><strong>Proxy Handling:</strong> Fixed a <code>no_proxy</code> hostname normalisation bypass that could lead to Server-Side Request Forgery (SSRF). (<strong><a href="https://redirect.github.com/axios/axios/issues/10661">#10661</a></strong>)</li>
<li><strong>Header Injection:</strong> Fixed an unrestricted cloud metadata exfiltration vulnerability via a header injection chain. (<strong><a href="https://redirect.github.com/axios/axios/issues/10660">#10660</a></strong>)</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li><strong>Runtime Support:</strong> Added compatibility checks and documentation for Deno and Bun environments. (<strong><a href="https://redirect.github.com/axios/axios/issues/10652">#10652</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10653">#10653</a></strong>)</li>
</ul>
<h2>🔧 Maintenance &amp; Chores</h2>
<ul>
<li><strong>CI Security:</strong> Hardened workflow permissions to least privilege, added the <code>zizmor</code> security scanner, pinned action versions, and gated npm publishing with OIDC and environment protection. (<strong><a href="https://redirect.github.com/axios/axios/issues/10618">#10618</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10619">#10619</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10627">#10627</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10637">#10637</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10666">#10666</a></strong>)</li>
<li><strong>Dependencies:</strong> Bumped <code>serialize-javascript</code>, <code>handlebars</code>, <code>picomatch</code>, <code>vite</code>, and <code>denoland/setup-deno</code> to latest versions. Added a 7-day Dependabot cooldown period. (<strong><a href="https://redirect.github.com/axios/axios/issues/10574">#10574</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10572">#10572</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10568">#10568</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10663">#10663</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10664">#10664</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10665">#10665</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10669">#10669</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10670">#10670</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10616">#10616</a></strong>)</li>
<li><strong>Documentation:</strong> Unified docs, improved <code>beforeRedirect</code> credential leakage example, clarified <code>withCredentials</code>/<code>withXSRFToken</code> behaviour, HTTP/2 support notes, async/await timeout error handling, header case preservation, and various typo fixes. (<strong><a href="https://redirect.github.com/axios/axios/issues/10649">#10649</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10624">#10624</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7452">#7452</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/7471">#7471</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10654">#10654</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10644">#10644</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10589">#10589</a></strong>)</li>
<li><strong>Housekeeping:</strong> Removed stale files, regenerated lockfile, and updated sponsor scripts and blocks. (<strong><a href="https://redirect.github.com/axios/axios/issues/10584">#10584</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10650">#10650</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10582">#10582</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10640">#10640</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10659">#10659</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10668">#10668</a></strong>)</li>
<li><strong>Tests:</strong> Added regression coverage for urlencoded <code>Content-Type</code> casing. (<strong><a href="https://redirect.github.com/axios/axios/issues/10573">#10573</a></strong>)</li>
</ul>
<h2>🌟 New Contributors</h2>
<p>We are thrilled to welcome our new contributors. Thank you for helping improve Axios:</p>
<ul>
<li><strong><a href="https://github.com/raashish1601"><code>@​raashish1601</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10573">#10573</a></strong>)</li>
<li><strong><a href="https://github.com/Kilros0817"><code>@​Kilros0817</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10625">#10625</a></strong>)</li>
<li><strong><a href="https://github.com/ashstrc"><code>@​ashstrc</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10624">#10624</a></strong>)</li>
<li><strong><a href="https://github.com/Abhi3975"><code>@​Abhi3975</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10589">#10589</a></strong>)</li>
<li><strong><a href="https://github.com/theamodhshetty"><code>@​theamodhshetty</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/7452">#7452</a></strong>)</li>
</ul>
<h2>v1.14.0</h2>
<p>This release focuses on compatibility fixes, adapter stability improvements, and test/tooling modernisation.</p>
<h2>⚠️ Important Changes</h2>
<ul>
<li><strong>Breaking Changes:</strong> None identified in this release.</li>
<li><strong>Action Required:</strong> If you rely on env-based proxy behaviour or CJS resolution edge-cases, validate your integration after upgrade (notably <code>proxy-from-env</code> v2 alignment and <code>main</code> entry compatibility fix).</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li><strong>Runtime Features:</strong> No new end-user features were introduced in this release.</li>
<li><strong>Test Coverage Expansion:</strong> Added broader smoke/module test coverage for CJS and ESM package usage. (<a href="https://redirect.github.com/axios/axios/pull/7510">#7510</a>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Headers:</strong> Trim trailing CRLF in normalised header values. (<a href="https://redirect.github.com/axios/axios/pull/7456">#7456</a>)</li>
<li><strong>HTTP/2:</strong> Close detached HTTP/2 sessions on timeout to avoid lingering sessions. (<a href="https://redirect.github.com/axios/axios/pull/7457">#7457</a>)</li>
<li><strong>Fetch Adapter:</strong> Cancel <code>ReadableStream</code> created during request-stream capability probing to prevent async resource leaks. (<a href="https://redirect.github.com/axios/axios/pull/7515">#7515</a>)</li>
<li><strong>Proxy Handling:</strong> Fixed env proxy behavior with <code>proxy-from-env</code> v2 usage. (<a href="https://redirect.github.com/axios/axios/pull/7499">#7499</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v1.x/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
<h2>v1.15.0 — April 7, 2026</h2>
<p>This release delivers two critical security patches targeting header injection and SSRF via proxy bypass, adds official runtime support for Deno and Bun, and includes significant CI security hardening.</p>
<h2>🔒 Security Fixes</h2>
<ul>
<li>
<p><strong>Header Injection (CRLF):</strong> Rejects any header value containing <code>\r</code> or <code>\n</code> characters to block CRLF injection chains that could be used to exfiltrate cloud metadata (IMDS). Behavior change: headers with CR/LF now throw <code>&quot;Invalid character in header content&quot;</code>. (<strong><a href="https://redirect.github.com/axios/axios/issues/10660">#10660</a></strong>)</p>
</li>
<li>
<p><strong>SSRF via <code>no_proxy</code> Bypass:</strong> Introduces a <code>shouldBypassProxy</code> helper that normalises hostnames (strips trailing dots, handles bracketed IPv6) before evaluating <code>no_proxy</code>/<code>NO_PROXY</code> rules, closing a gap that could cause loopback or internal hosts to be inadvertently proxied. (<strong><a href="https://redirect.github.com/axios/axios/issues/10661">#10661</a></strong>)</p>
</li>
</ul>
<h2>🚀 New Features</h2>
<ul>
<li><strong>Deno &amp; Bun Runtime Support:</strong> Added full smoke test suites for Deno and Bun, with CI workflows that run both runtimes before any release is cut. (<strong><a href="https://redirect.github.com/axios/axios/issues/10652">#10652</a></strong>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<ul>
<li><strong>Node.js v22 Compatibility:</strong> Replaced deprecated <code>url.parse()</code> calls with the WHATWG <code>URL</code>/<code>URLSearchParams</code> API across examples, sandbox, and tests, eliminating <code>DEP0169</code> deprecation warnings on Node.js v22+. (<strong><a href="https://redirect.github.com/axios/axios/issues/10625">#10625</a></strong>)</li>
</ul>
<h2>🔧 Maintenance &amp; Chores</h2>
<ul>
<li>
<p><strong>CI Security Hardening:</strong> Added <a href="https://github.com/zizmorcore/zizmor">zizmor</a> GitHub Actions security scanner; switched npm publish to OIDC Trusted Publishing (removing the long-lived <code>NODE_AUTH_TOKEN</code>); pinned all action references to full commit SHAs; narrowed workflow permissions to least privilege; gated the publish step behind a dedicated <code>npm-publish</code> environment; and blocked the sponsor-block workflow from running on forks. (<strong><a href="https://redirect.github.com/axios/axios/issues/10618">#10618</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10619">#10619</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10627">#10627</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10637">#10637</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10641">#10641</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10666">#10666</a></strong>)</p>
</li>
<li>
<p><strong>Docs:</strong> Clarified HTTP/2 support and the unsupported <code>httpVersion</code> option; added documentation for header case preservation; improved the <code>beforeRedirect</code> example to prevent accidental credential leakage. (<strong><a href="https://redirect.github.com/axios/axios/issues/10644">#10644</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10654">#10654</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10624">#10624</a></strong>)</p>
</li>
<li>
<p><strong>Dependencies:</strong> Bumped <code>picomatch</code>, <code>handlebars</code>, <code>serialize-javascript</code>, <code>vite</code> (×3), <code>denoland/setup-deno</code>, and 4 additional dev dependencies to latest versions. (<strong><a href="https://redirect.github.com/axios/axios/issues/10564">#10564</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10565">#10565</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10567">#10567</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10568">#10568</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10572">#10572</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10574">#10574</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10663">#10663</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10664">#10664</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10665">#10665</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10669">#10669</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10670">#10670</a></strong>)</p>
</li>
</ul>
<h2>🌟 New Contributors</h2>
<p>We are thrilled to welcome our new contributors. Thank you for helping improve axios:</p>
<ul>
<li><strong><a href="https://github.com/Kilros0817"><code>@​Kilros0817</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10625">#10625</a></strong>)</li>
<li><strong><a href="https://github.com/shaanmajid"><code>@​shaanmajid</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10616">#10616</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10617">#10617</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10618">#10618</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10619">#10619</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10637">#10637</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10641">#10641</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10666">#10666</a></strong>)</li>
<li><strong><a href="https://github.com/ashstrc"><code>@​ashstrc</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10624">#10624</a></strong>, <strong><a href="https://redirect.github.com/axios/axios/issues/10644">#10644</a></strong>)</li>
<li><strong><a href="https://github.com/Abhi3975"><code>@​Abhi3975</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10589">#10589</a></strong>)</li>
<li><strong><a href="https://github.com/raashish1601"><code>@​raashish1601</code></a></strong> (<strong><a href="https://redirect.github.com/axios/axios/issues/10573">#10573</a></strong>)</li>
</ul>
<p><a href="https://github.com/axios/axios/compare/v1.14.0...v1.15.0">Full Changelog</a></p>
<hr />
<h2>v1.14.0 — March 27, 2026</h2>
<p>This release fixes a security vulnerability in the <code>formidable</code> dependency, resolves a CommonJS compatibility regression, hardens proxy and HTTP/2 handling, and modernises the build and test toolchain.</p>
<h2>🔒 Security Fixes</h2>
<ul>
<li><strong>Formidable Vulnerability:</strong> Upgraded <code>formidable</code> from v2 to v3 to address a reported arbitrary-file vulnerability. Updated test server and assertions to align with the v3 API. (<strong><a href="https://redirect.github.com/axios/axios/issues/7533">#7533</a></strong>)</li>
</ul>
<h2>🐛 Bug Fixes</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/772a4e54ecc4cc2421e2b746daff0aca10f359d7"><code>772a4e5</code></a> chore(release): prepare release 1.15.0 (<a href="https://redirect.github.com/axios/axios/issues/10671">#10671</a>)</li>
<li><a href="https://github.com/axios/axios/commit/4b071371be2f810b4bc7797a13838e0f806ebb22"><code>4b07137</code></a> chore(deps-dev): bump vite from 8.0.0 to 8.0.5 in /tests/smoke/esm (<a href="https://redirect.github.com/axios/axios/issues/10663">#10663</a>)</li>
<li><a href="https://github.com/axios/axios/commit/51e57b39db251bfe3d34af5c943dfea18e06c8b6"><code>51e57b3</code></a> chore(deps-dev): bump vite from 8.0.2 to 8.0.5 (<a href="https://redirect.github.com/axios/axios/issues/10664">#10664</a>)</li>
<li><a href="https://github.com/axios/axios/commit/fba1a77930f0c459677b729161627234b88c90aa"><code>fba1a77</code></a> chore(deps-dev): bump vite from 8.0.2 to 8.0.5 in /tests/module/esm (<a href="https://redirect.github.com/axios/axios/issues/10665">#10665</a>)</li>
<li><a href="https://github.com/axios/axios/commit/0bf6e28eac86e87da2b60bbf5ea4237910e1a08e"><code>0bf6e28</code></a> chore(deps): bump denoland/setup-deno in the github-actions group (<a href="https://redirect.github.com/axios/axios/issues/10669">#10669</a>)</li>
<li><a href="https://github.com/axios/axios/commit/8107157c572ee4a54cb28c01ab7f7f3d895ba661"><code>8107157</code></a> chore(deps-dev): bump the development_dependencies group with 4 updates (<a href="https://redirect.github.com/axios/axios/issues/10670">#10670</a>)</li>
<li><a href="https://github.com/axios/axios/commit/e66530e3302d56176befd0778155dafea2487542"><code>e66530e</code></a> ci: require npm-publish environment for releases (<a href="https://redirect.github.com/axios/axios/issues/10666">#10666</a>)</li>
<li><a href="https://github.com/axios/axios/commit/49f23cbfe4d308a075281c5f798d4c68f648cbe2"><code>49f23cb</code></a> chore(sponsor): update sponsor block (<a href="https://redirect.github.com/axios/axios/issues/10668">#10668</a>)</li>
<li><a href="https://github.com/axios/axios/commit/363185461b90b1b78845dc8a99a1f103d9b122a1"><code>3631854</code></a> fix: unrestricted cloud metadata exfiltration via header injection chain (<a href="https://redirect.github.com/axios/axios/issues/10">#10</a>...</li>
<li><a href="https://github.com/axios/axios/commit/fb3befb6daac6cad26b2e54094d0f2d9e47f24df"><code>fb3befb</code></a> fix: no_proxy hostname normalization bypass leads to ssrf (<a href="https://redirect.github.com/axios/axios/issues/10661">#10661</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v1.13.6...v1.15.0">compare view</a></li>
</ul>
</details>
<details>
<summary>Install script changes</summary>
<p>This version modifies <code>prepare</code> script that runs during installation. Review the package contents before updating.</p>
</details>
<br />

Updates `@currents/playwright` from 1.22.3 to 1.23.0
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/currents-dev/currents-playwright-changelog/blob/main/CHANGELOG.md"><code>@​currents/playwright</code>'s changelog</a>.</em></p>
<blockquote>
<h1><a href="https://github.com/currents-dev/currents-playwright/compare/v1.22.3...v1.23.0">1.23.0</a> (2026-04-09)</h1>
<h3>Bug Fixes</h3>
<ul>
<li>bump axios from 1.13.6 to 1.15.0 in /packages/reporter [CSR-0] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/768">#768</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/ae427638542cf8a1f47e4a700fd25c48f7fb97f7">ae42763</a>)</li>
<li>bump lodash from 4.17.23 to 4.18.1 (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/767">#767</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/017e881d245c6fd80acfa9abe4164b48a063b358">017e881</a>)</li>
<li>enhance CI provider detection and PR env capture [CSR-4046] (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/765">#765</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/d2bd16e272eaf37bfb181168eafc7f672df4d13b">d2bd16e</a>)</li>
</ul>
<h3>Features</h3>
<ul>
<li>add visual testing example (<a href="https://redirect.github.com/currents-dev/currents-playwright/issues/738">#738</a>) (<a href="https://github.com/currents-dev/currents-playwright/commit/5b044829d6c7c1adad471f5a1795f904d0866450">5b04482</a>)</li>
</ul>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/currents-dev/currents-playwright-changelog/commits">compare view</a></li>
</ul>
</details>
<br />


---

## Discussion

### Comment by @codecov-commenter on April 12, 2026 at 11:15 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1575?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 75.88%. Comparing base ([`9b58e2a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9b58e2adbd50ec1afe0dfab5cf4132d80a28472b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`662cc69`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/662cc69aad2139658c609edd6f1a1505ca5b3b79?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1575   +/-   ##
=======================================
  Coverage   75.88%   75.88%           
=======================================
  Files         103      103           
  Lines        3164     3164           
  Branches      686      686           
=======================================
  Hits         2401     2401           
  Misses        684      684           
  Partials       79       79           
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1575?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @swadeley - Approved on April 14, 2026 at 03:55 PM UTC

ACK

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1575*
