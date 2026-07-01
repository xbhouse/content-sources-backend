---
type: pull_request
number: 1250
title: "chore(deps): bump axios, @redhat-cloud-services/rbac-client, @redhat-cloud-services/frontend-components-remediations and @redhat-cloud-services/host-inventory-client"
state: closed
author: dependabot
created: 2025-01-15T10:13:25Z
updated: 2025-01-17T12:48:16Z
closed: 2025-01-17T12:48:14Z
base_branch: master
head_branch: dependabot/npm_and_yarn/multi-2061adb26d
labels: ["dependencies", "major"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1250
---

# Pull Request #1250: chore(deps): bump axios, @redhat-cloud-services/rbac-client, @redhat-cloud-services/frontend-components-remediations and @redhat-cloud-services/host-inventory-client

**Author**: @dependabot
**Created**: January 15, 2025 at 10:13 AM UTC
**Status**: Closed
**Labels**: `dependencies`, `major`
**Base**: `master` ← **Head**: `dependabot/npm_and_yarn/multi-2061adb26d`

## Description

Bumps [axios](https://github.com/axios/axios) to 1.7.7 and updates ancestor dependencies [axios](https://github.com/axios/axios), [@redhat-cloud-services/rbac-client](https://github.com/RedHatInsights/javascript-clients), [@redhat-cloud-services/frontend-components-remediations](https://github.com/RedHatInsights/frontend-components) and [@redhat-cloud-services/host-inventory-client](https://github.com/RedHatInsights/javascript-clients). These dependencies need to be updated together.

Updates `axios` from 0.27.2 to 1.7.7
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/releases">axios's releases</a>.</em></p>
<blockquote>
<h2>Release v1.7.7</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> fix stream handling in Safari by fallback to using a stream reader instead of an async iterator; (<a href="https://redirect.github.com/axios/axios/issues/6584">#6584</a>) (<a href="https://github.com/axios/axios/commit/d1980854fee1765cd02fa0787adf5d6e34dd9dcf">d198085</a>)</li>
<li><strong>http:</strong> fixed support for IPv6 literal strings in url (<a href="https://redirect.github.com/axios/axios/issues/5731">#5731</a>) (<a href="https://github.com/axios/axios/commit/364993f0d8bc6e0e06f76b8a35d2d0a35cab054c">364993f</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/Rishi556" title="+39/-1 ([#5731](https://github.com/axios/axios/issues/5731) )">Rishi556</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+27/-7 ([#6584](https://github.com/axios/axios/issues/6584) )">Dmitriy Mozgovoy</a></li>
</ul>
<h2>Release v1.7.6</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> fix content length calculation for FormData payload; (<a href="https://redirect.github.com/axios/axios/issues/6524">#6524</a>) (<a href="https://github.com/axios/axios/commit/085f56861a83e9ac02c140ad9d68dac540dfeeaa">085f568</a>)</li>
<li><strong>fetch:</strong> optimize signals composing logic; (<a href="https://redirect.github.com/axios/axios/issues/6582">#6582</a>) (<a href="https://github.com/axios/axios/commit/df9889b83c2cc37e9e6189675a73ab70c60f031f">df9889b</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+98/-46 ([#6582](https://github.com/axios/axios/issues/6582) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/jacquesg" title="+5/-1 ([#6524](https://github.com/axios/axios/issues/6524) )">Jacques Germishuys</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/kuroino721" title="+3/-1 ([#6575](https://github.com/axios/axios/issues/6575) )">kuroino721</a></li>
</ul>
<h2>Release v1.7.5</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>adapter:</strong> fix undefined reference to hasBrowserEnv (<a href="https://redirect.github.com/axios/axios/issues/6572">#6572</a>) (<a href="https://github.com/axios/axios/commit/7004707c4180b416341863bd86913fe4fc2f1df1">7004707</a>)</li>
<li><strong>core:</strong> add the missed implementation of AxiosError#status property; (<a href="https://redirect.github.com/axios/axios/issues/6573">#6573</a>) (<a href="https://github.com/axios/axios/commit/6700a8adac06942205f6a7a21421ecb36c4e0852">6700a8a</a>)</li>
<li><strong>core:</strong> fix <code>ReferenceError: navigator is not defined</code> for custom environments; (<a href="https://redirect.github.com/axios/axios/issues/6567">#6567</a>) (<a href="https://github.com/axios/axios/commit/fed1a4b2d78ed4a588c84e09d32749ed01dc2794">fed1a4b</a>)</li>
<li><strong>fetch:</strong> fix credentials handling in Cloudflare workers (<a href="https://redirect.github.com/axios/axios/issues/6533">#6533</a>) (<a href="https://github.com/axios/axios/commit/550d885eb90fd156add7b93bbdc54d30d2f9a98d">550d885</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+187/-83 ([#6573](https://github.com/axios/axios/issues/6573) [#6567](https://github.com/axios/axios/issues/6567) [#6566](https://github.com/axios/axios/issues/6566) [#6564](https://github.com/axios/axios/issues/6564) [#6563](https://github.com/axios/axios/issues/6563) [#6557](https://github.com/axios/axios/issues/6557) [#6556](https://github.com/axios/axios/issues/6556) [#6555](https://github.com/axios/axios/issues/6555) [#6554](https://github.com/axios/axios/issues/6554) [#6552](https://github.com/axios/axios/issues/6552) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/antoninbas" title="+6/-6 ([#6572](https://github.com/axios/axios/issues/6572) )">Antonin Bas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/hansottowirtz" title="+4/-1 ([#6533](https://github.com/axios/axios/issues/6533) )">Hans Otto Wirtz</a></li>
</ul>
<h2>Release v1.7.4</h2>
<h2>Release notes:</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>sec:</strong> CVE-2024-39338 (<a href="https://redirect.github.com/axios/axios/issues/6539">#6539</a>) (<a href="https://redirect.github.com/axios/axios/issues/6543">#6543</a>) (<a href="https://github.com/axios/axios/commit/6b6b605eaf73852fb2dae033f1e786155959de3a">6b6b605</a>)</li>
<li><strong>sec:</strong> disregard protocol-relative URL to remediate SSRF (<a href="https://redirect.github.com/axios/axios/issues/6539">#6539</a>) (<a href="https://github.com/axios/axios/commit/07a661a2a6b9092c4aa640dcc7f724ec5e65bdda">07a661a</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/levpachmanov" title="+47/-11 ([#6543](https://github.com/axios/axios/issues/6543) )">Lev Pachmanov</a></li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Changelog</summary>
<p><em>Sourced from <a href="https://github.com/axios/axios/blob/v1.x/CHANGELOG.md">axios's changelog</a>.</em></p>
<blockquote>
<h2><a href="https://github.com/axios/axios/compare/v1.7.6...v1.7.7">1.7.7</a> (2024-08-31)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> fix stream handling in Safari by fallback to using a stream reader instead of an async iterator; (<a href="https://redirect.github.com/axios/axios/issues/6584">#6584</a>) (<a href="https://github.com/axios/axios/commit/d1980854fee1765cd02fa0787adf5d6e34dd9dcf">d198085</a>)</li>
<li><strong>http:</strong> fixed support for IPv6 literal strings in url (<a href="https://redirect.github.com/axios/axios/issues/5731">#5731</a>) (<a href="https://github.com/axios/axios/commit/364993f0d8bc6e0e06f76b8a35d2d0a35cab054c">364993f</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/Rishi556" title="+39/-1 ([#5731](https://github.com/axios/axios/issues/5731) )">Rishi556</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+27/-7 ([#6584](https://github.com/axios/axios/issues/6584) )">Dmitriy Mozgovoy</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.7.5...v1.7.6">1.7.6</a> (2024-08-30)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>fetch:</strong> fix content length calculation for FormData payload; (<a href="https://redirect.github.com/axios/axios/issues/6524">#6524</a>) (<a href="https://github.com/axios/axios/commit/085f56861a83e9ac02c140ad9d68dac540dfeeaa">085f568</a>)</li>
<li><strong>fetch:</strong> optimize signals composing logic; (<a href="https://redirect.github.com/axios/axios/issues/6582">#6582</a>) (<a href="https://github.com/axios/axios/commit/df9889b83c2cc37e9e6189675a73ab70c60f031f">df9889b</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+98/-46 ([#6582](https://github.com/axios/axios/issues/6582) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/jacquesg" title="+5/-1 ([#6524](https://github.com/axios/axios/issues/6524) )">Jacques Germishuys</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/kuroino721" title="+3/-1 ([#6575](https://github.com/axios/axios/issues/6575) )">kuroino721</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.7.4...v1.7.5">1.7.5</a> (2024-08-23)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>adapter:</strong> fix undefined reference to hasBrowserEnv (<a href="https://redirect.github.com/axios/axios/issues/6572">#6572</a>) (<a href="https://github.com/axios/axios/commit/7004707c4180b416341863bd86913fe4fc2f1df1">7004707</a>)</li>
<li><strong>core:</strong> add the missed implementation of AxiosError#status property; (<a href="https://redirect.github.com/axios/axios/issues/6573">#6573</a>) (<a href="https://github.com/axios/axios/commit/6700a8adac06942205f6a7a21421ecb36c4e0852">6700a8a</a>)</li>
<li><strong>core:</strong> fix <code>ReferenceError: navigator is not defined</code> for custom environments; (<a href="https://redirect.github.com/axios/axios/issues/6567">#6567</a>) (<a href="https://github.com/axios/axios/commit/fed1a4b2d78ed4a588c84e09d32749ed01dc2794">fed1a4b</a>)</li>
<li><strong>fetch:</strong> fix credentials handling in Cloudflare workers (<a href="https://redirect.github.com/axios/axios/issues/6533">#6533</a>) (<a href="https://github.com/axios/axios/commit/550d885eb90fd156add7b93bbdc54d30d2f9a98d">550d885</a>)</li>
</ul>
<h3>Contributors to this release</h3>
<ul>
<li><!-- raw HTML omitted --> <a href="https://github.com/DigitalBrainJS" title="+187/-83 ([#6573](https://github.com/axios/axios/issues/6573) [#6567](https://github.com/axios/axios/issues/6567) [#6566](https://github.com/axios/axios/issues/6566) [#6564](https://github.com/axios/axios/issues/6564) [#6563](https://github.com/axios/axios/issues/6563) [#6557](https://github.com/axios/axios/issues/6557) [#6556](https://github.com/axios/axios/issues/6556) [#6555](https://github.com/axios/axios/issues/6555) [#6554](https://github.com/axios/axios/issues/6554) [#6552](https://github.com/axios/axios/issues/6552) )">Dmitriy Mozgovoy</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/antoninbas" title="+6/-6 ([#6572](https://github.com/axios/axios/issues/6572) )">Antonin Bas</a></li>
<li><!-- raw HTML omitted --> <a href="https://github.com/hansottowirtz" title="+4/-1 ([#6533](https://github.com/axios/axios/issues/6533) )">Hans Otto Wirtz</a></li>
</ul>
<h2><a href="https://github.com/axios/axios/compare/v1.7.3...v1.7.4">1.7.4</a> (2024-08-13)</h2>
<h3>Bug Fixes</h3>
<ul>
<li><strong>sec:</strong> CVE-2024-39338 (<a href="https://redirect.github.com/axios/axios/issues/6539">#6539</a>) (<a href="https://redirect.github.com/axios/axios/issues/6543">#6543</a>) (<a href="https://github.com/axios/axios/commit/6b6b605eaf73852fb2dae033f1e786155959de3a">6b6b605</a>)</li>
<li><strong>sec:</strong> disregard protocol-relative URL to remediate SSRF (<a href="https://redirect.github.com/axios/axios/issues/6539">#6539</a>) (<a href="https://github.com/axios/axios/commit/07a661a2a6b9092c4aa640dcc7f724ec5e65bdda">07a661a</a>)</li>
</ul>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/axios/axios/commit/5b8a826771b77ab30081d033fdba9ef3b90e439a"><code>5b8a826</code></a> chore(release): v1.7.7 (<a href="https://redirect.github.com/axios/axios/issues/6585">#6585</a>)</li>
<li><a href="https://github.com/axios/axios/commit/364993f0d8bc6e0e06f76b8a35d2d0a35cab054c"><code>364993f</code></a> fix(http): fixed support for IPv6 literal strings in url (<a href="https://redirect.github.com/axios/axios/issues/5731">#5731</a>)</li>
<li><a href="https://github.com/axios/axios/commit/d1980854fee1765cd02fa0787adf5d6e34dd9dcf"><code>d198085</code></a> fix(fetch): fix stream handling in Safari by fallback to using a stream reade...</li>
<li><a href="https://github.com/axios/axios/commit/d584fcfa62ba5217baf2be0748b7c5eda6da16ad"><code>d584fcf</code></a> chore(release): v1.7.6 (<a href="https://redirect.github.com/axios/axios/issues/6583">#6583</a>)</li>
<li><a href="https://github.com/axios/axios/commit/bc03c6cbc41eb8449daa2f4b6b8048671a05bded"><code>bc03c6c</code></a> chore(examples): fix module import (<a href="https://redirect.github.com/axios/axios/issues/6575">#6575</a>)</li>
<li><a href="https://github.com/axios/axios/commit/df9889b83c2cc37e9e6189675a73ab70c60f031f"><code>df9889b</code></a> fix(fetch): optimize signals composing logic; (<a href="https://redirect.github.com/axios/axios/issues/6582">#6582</a>)</li>
<li><a href="https://github.com/axios/axios/commit/ee208cfcae8770ac579f26e7278867567886e026"><code>ee208cf</code></a> chore(sponsor): update sponsor block (<a href="https://redirect.github.com/axios/axios/issues/6576">#6576</a>)</li>
<li><a href="https://github.com/axios/axios/commit/085f56861a83e9ac02c140ad9d68dac540dfeeaa"><code>085f568</code></a> fix(fetch): fix content length calculation for FormData payload; (<a href="https://redirect.github.com/axios/axios/issues/6524">#6524</a>)</li>
<li><a href="https://github.com/axios/axios/commit/59cd6b0dece4050b190717a7c5cdf77906ce2104"><code>59cd6b0</code></a> chore(release): v1.7.5 (<a href="https://redirect.github.com/axios/axios/issues/6574">#6574</a>)</li>
<li><a href="https://github.com/axios/axios/commit/6700a8adac06942205f6a7a21421ecb36c4e0852"><code>6700a8a</code></a> fix(core): add the missed implementation of AxiosError#status property; (<a href="https://redirect.github.com/axios/axios/issues/6573">#6573</a>)</li>
<li>Additional commits viewable in <a href="https://github.com/axios/axios/compare/v0.27.2...v1.7.7">compare view</a></li>
</ul>
</details>
<br />

Updates `@redhat-cloud-services/rbac-client` from 1.4.5 to 2.2.11
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/javascript-clients/releases"><code>@​redhat-cloud-services/rbac-client</code>'s releases</a>.</em></p>
<blockquote>
<h2><code>@​redhat-cloud-services/rbac-client-2</code>.2.11</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/rbac-client-2.2.10...@redhat-cloud-services/rbac-client-2.2.11">2.2.11</a> (2024-11-18)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/javascript-clients-shared</code> updated to version <code>1.2.6</code></li>
</ul>
<h2><code>@​redhat-cloud-services/rbac-client-2</code>.2.10</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/rbac-client-2.2.9...@redhat-cloud-services/rbac-client-2.2.10">2.2.10</a> (2024-11-18)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/javascript-clients-shared</code> updated to version <code>1.2.5</code></li>
</ul>
<h2><code>@​redhat-cloud-services/rbac-client-2</code>.2.9</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/rbac-client-2.2.8...@redhat-cloud-services/rbac-client-2.2.9">2.2.9</a> (2024-11-06)</h2>
<h2><code>@​redhat-cloud-services/rbac-client-2</code>.2.8</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/rbac-client-2.2.7...@redhat-cloud-services/rbac-client-2.2.8">2.2.8</a> (2024-10-31)</h2>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/000285399923403d335e0a5f4a50de3e111b1e17"><code>0002853</code></a> release: bump <code>@​redhat-cloud-services/rbac-client</code> to 2.2.11 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/af2edbc4fb4509e0e1f0adcf42ab7d68b8dd9476"><code>af2edbc</code></a> release: bump <code>@​redhat-cloud-services/javascript-clients-shared</code> to 1.2.6 [skip...</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/3a4039b7f0a8752f2c00dd6ff8c9f1694ff70028"><code>3a4039b</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/javascript-clients/issues/304">#304</a> from catastrophe-brandon/btweed/integrate-e2e-test</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/34ea7fbe2fe3ad04f7016037c3bb4f358e8eb0a8"><code>34ea7fb</code></a> Improve test:integration cmd</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/845da37da7851eba7d1a894efb151b7db37232c8"><code>845da37</code></a> Interactive code review clean up</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/2568b2ad1822f52af419a8e73ba58ef59cf43ef9"><code>2568b2a</code></a> Add Maven target folder to .gitignore</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/dde21b73e0cfcaca9896a0144e99b7bf73146b4e"><code>dde21b7</code></a> release: bump <code>@​redhat-cloud-services/rbac-client</code> to 2.2.10 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/e4acf7bb9cbf0ac415b024752209d06dcd575852"><code>e4acf7b</code></a> release: bump <code>@​redhat-cloud-services/compliance-client</code> to 1.1.1 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/606529ad7fe6991bf78a887ebcf853629b8c6f55"><code>606529a</code></a> release: bump <code>@​redhat-cloud-services/integrations-client</code> to 2.5.5 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/5df79bea9019681030fdc12d708ebb17a9b57692"><code>5df79be</code></a> release: bump <code>@​redhat-cloud-services/entitlements-client</code> to 2.1.1 [skip ci]</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/rbac-client-1.4.5...@redhat-cloud-services/rbac-client-2.2.11">compare view</a></li>
</ul>
</details>
<br />

Updates `@redhat-cloud-services/frontend-components-remediations` from 3.2.12 to 3.2.19
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/frontend-components/releases"><code>@​redhat-cloud-services/frontend-components-remediations</code>'s releases</a>.</em></p>
<blockquote>
<h2><code>@​redhat-cloud-services/frontend-components-remediations-3</code>.2.19</h2>
<h1>Changelog</h1>
<p>This file was generated using <a href="https://github.com/jscutlery/semver"><code>@​jscutlery/semver</code></a>.</p>
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
<h2><code>@​redhat-cloud-services/frontend-components-remediations-3</code>.2.18</h2>
<h1>Changelog</h1>
<p>This file was generated using <a href="https://github.com/jscutlery/semver"><code>@​jscutlery/semver</code></a>.</p>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.17...@redhat-cloud-services/frontend-components-remediations-3.2.18">3.2.18</a> (2025-01-14)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/frontend-components</code> updated to version <code>5.1.0</code></li>
</ul>
<h2><a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/frontend-components-remediations-3.2.16...@redhat-cloud-services/frontend-components-remediations-3.2.17">3.2.17</a> (2025-01-13)</h2>
<!-- raw HTML omitted -->
</blockquote>
<p>... (truncated)</p>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/9b2d07f778435130f678eff2a175d7d53d08b170"><code>9b2d07f</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-remediations</code> to 3.2.19...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/a3195a8dbe0428936cb5d106bb82a55b49eca561"><code>a3195a8</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-notifications</code> to 4.1.7...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/c13d77890d497c709bc4e66886a7063da9700c19"><code>c13d778</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components</code> to 5.1.1 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/1a06f260d6bc12897f6b2809f5fe004cf53b6fa8"><code>1a06f26</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-utilities</code> to 5.0.6 [sk...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/abd2542b77a2758ef9e7499da4c02a33773eb3b6"><code>abd2542</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/frontend-components/issues/2144">#2144</a> from gkarat/fix-use-imports</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/42a9b954fea88cf11db2b5e0d9eaf6cc1afde176"><code>42a9b95</code></a> chore: bump <code>@​redhat-cloud-services/eslint-config-redhat-cloud-services</code> to 2.0...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/88b1259af35a307eb9aa4ddc8e3dff69166cb4c9"><code>88b1259</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/frontend-components/issues/2142">#2142</a> from Hyperkid123/emit-eslint-definitions</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/e332a571a6d4d159dddd5088a06c23f681562a44"><code>e332a57</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-advisor-components</code> to ...</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/869dd6f24feac484d05ff9e7f54135f320854d94"><code>869dd6f</code></a> chore: bump <code>@​redhat-cloud-services/rule-components</code> to 3.2.15 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/frontend-components/commit/49b4e0660cf3e5159b8c7c85d211336bb3e24163"><code>49b4e06</code></a> chore: bump <code>@​redhat-cloud-services/frontend-components-remediations</code> to 3.2.18...</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/frontend-components/compare/@redhat-cloud-services/rule-components-3.2.12...@redhat-cloud-services/frontend-components-remediations-3.2.19">compare view</a></li>
</ul>
</details>
<br />

Updates `@redhat-cloud-services/host-inventory-client` from 1.5.2 to 4.0.0
<details>
<summary>Release notes</summary>
<p><em>Sourced from <a href="https://github.com/RedHatInsights/javascript-clients/releases"><code>@​redhat-cloud-services/host-inventory-client</code>'s releases</a>.</em></p>
<blockquote>
<h2><code>@​redhat-cloud-services/host-inventory-client-4</code>.0.0</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-3.0.2...@redhat-cloud-services/host-inventory-client-4.0.0">4.0.0</a> (2025-01-10)</h2>
<h3>⚠ BREAKING CHANGES</h3>
<ul>
<li>Update host-inventory to use new builder</li>
</ul>
<h3>Features</h3>
<ul>
<li>Update host-inventory to use new builder (<a href="https://github.com/RedHatInsights/javascript-clients/commit/5aa3fec1d4947622587f6a6099ba8c4ca14ca8ca">5aa3fec</a>)</li>
</ul>
<h2><code>@​redhat-cloud-services/host-inventory-client-3</code>.0.2</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-3.0.1...@redhat-cloud-services/host-inventory-client-3.0.2">3.0.2</a> (2024-12-09)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/javascript-clients-shared</code> updated to version <code>1.2.7</code></li>
</ul>
<h2><code>@​redhat-cloud-services/host-inventory-client-3</code>.0.1</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-3.0.0...@redhat-cloud-services/host-inventory-client-3.0.1">3.0.1</a> (2024-11-22)</h2>
<h2><code>@​redhat-cloud-services/host-inventory-client-3</code>.0.0</h2>
<p>No release notes provided.</p>
<h2><code>@​redhat-cloud-services/host-inventory-client-2</code>.0.3</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-2.0.2...@redhat-cloud-services/host-inventory-client-2.0.3">2.0.3</a> (2024-11-18)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/javascript-clients-shared</code> updated to version <code>1.2.6</code></li>
</ul>
<h2><code>@​redhat-cloud-services/host-inventory-client-2</code>.0.2</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-2.0.1...@redhat-cloud-services/host-inventory-client-2.0.2">2.0.2</a> (2024-11-18)</h2>
<h3>Dependency Updates</h3>
<ul>
<li><code>@redhat-cloud-services/javascript-clients-shared</code> updated to version <code>1.2.5</code></li>
</ul>
<h2><code>@​redhat-cloud-services/host-inventory-client-2</code>.0.1</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-2.0.0...@redhat-cloud-services/host-inventory-client-2.0.1">2.0.1</a> (2024-11-07)</h2>
<h2><code>@​redhat-cloud-services/host-inventory-client-2</code>.0.0</h2>
<p>No release notes provided.</p>
<h2><code>@​redhat-cloud-services/host-inventory-client-1</code>.5.4</h2>
<h2><a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-1.5.3...@redhat-cloud-services/host-inventory-client-1.5.4">1.5.4</a> (2024-11-06)</h2>
</blockquote>
</details>
<details>
<summary>Commits</summary>
<ul>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/6911b634c7ed78e781bf3e40951cb19292e2ed61"><code>6911b63</code></a> release: bump <code>@​redhat-cloud-services/host-inventory-client</code> to 4.0.0 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/dd928176071baff71664f326b65aeda9e2082769"><code>dd92817</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/javascript-clients/issues/352">#352</a> from RedHatInsights/switch-host-inventory-new-builder</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/b0ea48ae2e47384efcfdb52abee019d5f47572a1"><code>b0ea48a</code></a> release: bump <code>@​redhat-cloud-services/notifications-client</code> to 4.0.0 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/c23dceecd68ef91022084769b4f67c027cb4aa75"><code>c23dcee</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/javascript-clients/issues/351">#351</a> from RedHatInsights/switch-notifications-new-builder</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/62bc1bb2820719a33d78062fe0c83650693d93d6"><code>62bc1bb</code></a> release: bump <code>@​redhat-cloud-services/integrations-client</code> to 4.0.0 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/0158002299c09f7800ce78ee7277da1bc7e88d4d"><code>0158002</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/javascript-clients/issues/350">#350</a> from RedHatInsights/switch-integrations-new-builder</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/12b18acc6425cfd2a3a0df1838cedcdb3951e063"><code>12b18ac</code></a> release: bump <code>@​redhat-cloud-services/entitlements-client</code> to 4.0.0 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/1356736919b046d6e76c55e3a01a56694886917c"><code>1356736</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/javascript-clients/issues/349">#349</a> from RedHatInsights/switch-entitlements-new-builder</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/7bb72ee105944bfb2320e4f7ffc655ef05ca025b"><code>7bb72ee</code></a> release: bump <code>@​redhat-cloud-services/config-manager-client</code> to 4.0.0 [skip ci]</li>
<li><a href="https://github.com/RedHatInsights/javascript-clients/commit/502da3179cc8f2ae87b22fe3969dd07d7e3c92b3"><code>502da31</code></a> Merge pull request <a href="https://redirect.github.com/RedHatInsights/javascript-clients/issues/348">#348</a> from RedHatInsights/switch-config-manager-new-builder</li>
<li>Additional commits viewable in <a href="https://github.com/RedHatInsights/javascript-clients/compare/@redhat-cloud-services/host-inventory-client-1.5.2...@redhat-cloud-services/host-inventory-client-4.0.0">compare view</a></li>
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

### Comment by @dependabot on January 17, 2025 at 12:48 PM UTC

Superseded by #1255.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1250*
